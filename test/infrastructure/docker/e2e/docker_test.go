// +build e2e

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	bootstrapv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha3"
	"sigs.k8s.io/cluster-api/bootstrap/kubeadm/types/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"
	"sigs.k8s.io/cluster-api/test/framework"
	infrav1 "sigs.k8s.io/cluster-api/test/infrastructure/docker/api/v1alpha3"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("Docker", func() {
	Describe("Cluster Creation", func() {
		var (
			namespace  string
			client     ctrlclient.Client
			clusterGen = &ClusterGenerator{}
			cluster    *clusterv1.Cluster
		)
		SetDefaultEventuallyTimeout(3 * time.Minute)
		SetDefaultEventuallyPollingInterval(10 * time.Second)

		BeforeEach(func() {
			namespace = "default"
		})

		AfterEach(func() {
			By("cleaning up the test cluster")
			// Dump cluster API and docker related resources to artifacts before deleting them.
			Expect(framework.DumpResources(mgmt, resourcesPath, GinkgoWriter)).To(Succeed())
			resources := map[string]runtime.Object{
				"DockerCluster":         &infrav1.DockerClusterList{},
				"DockerMachine":         &infrav1.DockerMachineList{},
				"DockerMachineTemplate": &infrav1.DockerMachineTemplateList{},
			}
			Expect(framework.DumpProviderResources(mgmt, resources, resourcesPath, GinkgoWriter)).To(Succeed())

			deleteClusterInput := framework.DeleteClusterInput{
				Deleter: client,
				Cluster: cluster,
			}
			framework.DeleteCluster(ctx, deleteClusterInput)

			waitForClusterDeletedInput := framework.WaitForClusterDeletedInput{
				Getter:  client,
				Cluster: cluster,
			}
			framework.WaitForClusterDeleted(ctx, waitForClusterDeletedInput)

			assertAllClusterAPIResourcesAreGoneInput := framework.AssertAllClusterAPIResourcesAreGoneInput{
				Lister:  client,
				Cluster: cluster,
			}
			framework.AssertAllClusterAPIResourcesAreGone(ctx, assertAllClusterAPIResourcesAreGoneInput)

			ensureDockerDeletedInput := ensureDockerArtifactsDeletedInput{
				Lister:  client,
				Cluster: cluster,
			}
			ensureDockerArtifactsDeleted(ensureDockerDeletedInput)
		})

		Describe("Multi-node controlplane cluster", func() {
			var controlPlane *controlplanev1.KubeadmControlPlane

			Specify("Basic create", func() {
				replicas := 3
				var (
					infraCluster *infrav1.DockerCluster
					template     *infrav1.DockerMachineTemplate
					err          error
				)
				cluster, infraCluster, controlPlane, template = clusterGen.GenerateCluster(namespace, int32(replicas))
				// Set failure domains here
				infraCluster.Spec.FailureDomains = clusterv1.FailureDomains{
					"domain-one":   {ControlPlane: true},
					"domain-two":   {ControlPlane: true},
					"domain-three": {ControlPlane: true},
					"domain-four":  {ControlPlane: false},
				}

				md, infraTemplate, bootstrapTemplate := GenerateMachineDeployment(cluster, 1)

				// Set up the client to the management cluster
				client, err = mgmt.GetClient()
				Expect(err).NotTo(HaveOccurred())

				// Set up the cluster object
				createClusterInput := framework.CreateClusterInput{
					Creator:      client,
					Cluster:      cluster,
					InfraCluster: infraCluster,
				}
				framework.CreateCluster(ctx, createClusterInput)

				// Set up the KubeadmControlPlane
				createKubeadmControlPlaneInput := framework.CreateKubeadmControlPlaneInput{
					Creator:         client,
					ControlPlane:    controlPlane,
					MachineTemplate: template,
				}
				framework.CreateKubeadmControlPlane(ctx, createKubeadmControlPlaneInput)

				// Wait for the cluster to provision.
				assertClusterProvisionsInput := framework.WaitForClusterToProvisionInput{
					Getter:  client,
					Cluster: cluster,
				}
				framework.WaitForClusterToProvision(ctx, assertClusterProvisionsInput)

				// Wait for at least one control plane node to be ready
				waitForOneKubeadmControlPlaneMachineToExistInput := framework.WaitForOneKubeadmControlPlaneMachineToExistInput{
					Lister:       client,
					Cluster:      cluster,
					ControlPlane: controlPlane,
				}
				framework.WaitForOneKubeadmControlPlaneMachineToExist(ctx, waitForOneKubeadmControlPlaneMachineToExistInput)

				// Install a networking solution on the workload cluster
				workloadClient, err := mgmt.GetWorkloadClient(ctx, cluster.Namespace, cluster.Name)
				Expect(err).ToNot(HaveOccurred())
				applyYAMLURLInput := framework.ApplyYAMLURLInput{
					Client:        workloadClient,
					HTTPGetter:    http.DefaultClient,
					NetworkingURL: "https://docs.projectcalico.org/manifests/calico.yaml",
					Scheme:        mgmt.Scheme,
				}
				framework.ApplyYAMLURL(ctx, applyYAMLURLInput)

				// Wait for the controlplane nodes to exist
				assertKubeadmControlPlaneNodesExistInput := framework.WaitForKubeadmControlPlaneMachinesToExistInput{
					Lister:       client,
					Cluster:      cluster,
					ControlPlane: controlPlane,
				}
				framework.WaitForKubeadmControlPlaneMachinesToExist(ctx, assertKubeadmControlPlaneNodesExistInput, "10m", "10s")

				// Create the workload nodes
				createMachineDeploymentinput := framework.CreateMachineDeploymentInput{
					Creator:                 client,
					MachineDeployment:       md,
					BootstrapConfigTemplate: bootstrapTemplate,
					InfraMachineTemplate:    infraTemplate,
				}
				framework.CreateMachineDeployment(ctx, createMachineDeploymentinput)

				// Wait for the workload nodes to exist
				waitForMachineDeploymentNodesToExistInput := framework.WaitForMachineDeploymentNodesToExistInput{
					Lister:            client,
					Cluster:           cluster,
					MachineDeployment: md,
				}
				framework.WaitForMachineDeploymentNodesToExist(ctx, waitForMachineDeploymentNodesToExistInput)

				// Wait for the control plane to be ready
				waitForControlPlaneToBeReadyInput := framework.WaitForControlPlaneToBeReadyInput{
					Getter:       client,
					ControlPlane: controlPlane,
				}
				framework.WaitForControlPlaneToBeReady(ctx, waitForControlPlaneToBeReadyInput)

				// Assert failure domain is working as expected
				assertControlPlaneFailureDomainInput := framework.AssertControlPlaneFailureDomainsInput{
					GetLister:  client,
					ClusterKey: util.ObjectKey(cluster),
					ExpectedFailureDomains: map[string]int{
						"domain-one":   1,
						"domain-two":   1,
						"domain-three": 1,
						"domain-four":  0,
					},
				}
				framework.AssertControlPlaneFailureDomains(ctx, assertControlPlaneFailureDomainInput)
			})

			Specify("Full upgrade", func() {
				By("upgrading the control plane object to a new version")
				patchHelper, err := patch.NewHelper(controlPlane, client)
				Expect(err).ToNot(HaveOccurred())
				controlPlane.Spec.Version = "1.17.2"
				Expect(patchHelper.Patch(ctx, controlPlane)).To(Succeed())
				By("waiting for all control plane nodes to exist")
				inClustersNamespaceListOption := ctrlclient.InNamespace(cluster.Namespace)
				// ControlPlane labels
				matchClusterListOption := ctrlclient.MatchingLabels{
					clusterv1.MachineControlPlaneLabelName: "",
					clusterv1.ClusterLabelName:             cluster.Name,
				}

				Eventually(func() (int, error) {
					machineList := &clusterv1.MachineList{}
					if err := client.List(ctx, machineList, inClustersNamespaceListOption, matchClusterListOption); err != nil {
						fmt.Println(err)
						return 0, err
					}
					upgraded := 0
					for _, machine := range machineList.Items {
						if *machine.Spec.Version == controlPlane.Spec.Version {
							upgraded++
						}
					}
					if len(machineList.Items) > upgraded {
						return 0, errors.New("old nodes remain")
					}
					return upgraded, nil
				}, "10m", "30s").Should(Equal(int(*controlPlane.Spec.Replicas)))
			})
		})

		Describe("Controlplane Adoption", func() {
			Specify("KubeadmControlPlane adopts up-to-date control plane Machines without modification", func() {
				var (
					controlPlane *controlplanev1.KubeadmControlPlane
					infraCluster *infrav1.DockerCluster
					template     *infrav1.DockerMachineTemplate
					err          error
				)
				replicas := 1 /* TODO: can't seem to get CAPD to bootstrap a cluster with more than one control plane machine */
				cluster, infraCluster, controlPlane, template = clusterGen.GenerateCluster(namespace, int32(replicas))
				controlPlaneRef := cluster.Spec.ControlPlaneRef
				cluster.Spec.ControlPlaneRef = nil

				// Set up the client to the management cluster
				client, err = mgmt.GetClient()
				Expect(err).NotTo(HaveOccurred())

				// Set up the cluster object
				createClusterInput := framework.CreateClusterInput{
					Creator:      client,
					Cluster:      cluster,
					InfraCluster: infraCluster,
				}
				framework.CreateCluster(ctx, createClusterInput)

				version := "1.16.3"

				// Wait for the cluster to provision.
				assertClusterProvisionsInput := framework.WaitForClusterToProvisionInput{
					Getter:  client,
					Cluster: cluster,
				}
				framework.WaitForClusterToProvision(ctx, assertClusterProvisionsInput)

				initMachines, bootstrap, infra := generateControlPlaneMachines(cluster, namespace, version, replicas)
				for i := 0; i < len(initMachines); i++ {
					// we have to go one at a time, otherwise weird things start to happen
					By("initializing control plane machines")
					createMachineInput := framework.CreateMachineInput{
						Creator:         client,
						BootstrapConfig: bootstrap[i],
						InfraMachine:    infra[i],
						Machine:         initMachines[i],
					}
					framework.CreateMachine(ctx, createMachineInput)

					// Wait for the first control plane machine to boot
					assertMachinesProvisionInput := framework.WaitForMachineNodesToExistInput{
						Getter:   client,
						Machines: initMachines[i : i+1],
					}
					framework.WaitForMachineNodesToExist(ctx, assertMachinesProvisionInput)

					if i == 0 {
						// Install a networking solution on the workload cluster
						workloadClient, err := mgmt.GetWorkloadClient(ctx, cluster.Namespace, cluster.Name)
						Expect(err).ToNot(HaveOccurred())
						applyYAMLURLInput := framework.ApplyYAMLURLInput{
							Client:        workloadClient,
							HTTPGetter:    http.DefaultClient,
							NetworkingURL: "https://docs.projectcalico.org/manifests/calico.yaml",
							Scheme:        mgmt.Scheme,
						}
						framework.ApplyYAMLURL(ctx, applyYAMLURLInput)
					}
				}

				// Set up the KubeadmControlPlane
				createKubeadmControlPlaneInput := framework.CreateKubeadmControlPlaneInput{
					Creator:         client,
					ControlPlane:    controlPlane,
					MachineTemplate: template,
				}
				framework.CreateKubeadmControlPlane(ctx, createKubeadmControlPlaneInput)

				// We have to set the control plane ref on the cluster as well
				cl := &clusterv1.Cluster{}
				client.Get(ctx, ctrlclient.ObjectKey{Namespace: cluster.Namespace, Name: cluster.Name}, cl)
				cl.Spec.ControlPlaneRef = controlPlaneRef
				Expect(client.Update(ctx, cl)).To(Succeed())

				waitForControlPlaneToBeUpToDateInput := framework.WaitForControlPlaneToBeUpToDateInput{
					Getter:       client,
					ControlPlane: controlPlane,
				}
				framework.WaitForControlPlaneToBeUpToDate(ctx, waitForControlPlaneToBeUpToDateInput)

				machines := clusterv1.MachineList{}
				Expect(client.List(ctx, &machines,
					ctrlclient.InNamespace(namespace),
					ctrlclient.HasLabels{
						clusterv1.MachineControlPlaneLabelName,
					})).To(Succeed())

				By("taking stable ownership of the Machines")
				for _, m := range machines.Items {
					Expect(&m).To(HaveControllerRef(framework.TypeToKind(controlPlane), controlPlane))
					Expect(m.CreationTimestamp.Time).To(BeTemporally("<", controlPlane.CreationTimestamp.Time))
				}
				Expect(machines.Items).To(HaveLen(1))

				By("taking ownership of the cluster's PKI material")
				secrets := corev1.SecretList{}
				Expect(client.List(ctx, &secrets, ctrlclient.InNamespace(namespace), ctrlclient.MatchingLabels{
					clusterv1.ClusterLabelName: cluster.Name,
				})).To(Succeed())

				for _, s := range secrets.Items {
					// We don't check the data, and removing it from the object makes assertions much easier to read
					s.Data = nil

					// The bootstrap secret should still be owned by the bootstrap config so its cleaned up properly,
					// but the cluster PKI materials should have their ownership transferred.
					switch {
					case strings.HasSuffix(s.Name, "-kubeconfig"):
						// Do nothing
					case strings.HasPrefix(s.Name, "bootstrap-"):
						fi := -1
						for i, b := range bootstrap {
							if s.Name == b.Name {
								fi = i
							}
						}
						Expect(fi).To(BeNumerically(">=", 0), "could not find matching bootstrap object for Secret %s", s.Name)
						Expect(&s).To(HaveControllerRef(framework.TypeToKind(bootstrap[fi]), bootstrap[fi]))
					default:
						Expect(&s).To(HaveControllerRef(framework.TypeToKind(controlPlane), controlPlane))
					}
				}
				Expect(secrets.Items).To(HaveLen(4 /* pki */ + 1 /* kubeconfig */ + int(replicas)))

				By("ensuring we can still join machines after the adoption")
				md, infraTemplate, bootstrapTemplate := GenerateMachineDeployment(cluster, 1)

				// Create the workload nodes
				createMachineDeploymentinput := framework.CreateMachineDeploymentInput{
					Creator:                 client,
					MachineDeployment:       md,
					BootstrapConfigTemplate: bootstrapTemplate,
					InfraMachineTemplate:    infraTemplate,
				}
				framework.CreateMachineDeployment(ctx, createMachineDeploymentinput)

				// Wait for the workload nodes to exist
				waitForMachineDeploymentNodesToExistInput := framework.WaitForMachineDeploymentNodesToExistInput{
					Lister:            client,
					Cluster:           cluster,
					MachineDeployment: md,
				}
				framework.WaitForMachineDeploymentNodesToExist(ctx, waitForMachineDeploymentNodesToExistInput)
			})
		})
	})
})

func GenerateMachineDeployment(cluster *clusterv1.Cluster, replicas int32) (*clusterv1.MachineDeployment, *infrav1.DockerMachineTemplate, *bootstrapv1.KubeadmConfigTemplate) {
	namespace := cluster.GetNamespace()
	generatedName := fmt.Sprintf("%s-md", cluster.GetName())
	version := "1.16.3"

	infraTemplate := &infrav1.DockerMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: infrav1.DockerMachineTemplateSpec{},
	}

	bootstrap := &bootstrapv1.KubeadmConfigTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
	}

	template := clusterv1.MachineTemplateSpec{
		ObjectMeta: clusterv1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: clusterv1.MachineSpec{
			ClusterName: cluster.GetName(),
			Bootstrap: clusterv1.Bootstrap{
				ConfigRef: &corev1.ObjectReference{
					APIVersion: bootstrapv1.GroupVersion.String(),
					Kind:       framework.TypeToKind(bootstrap),
					Namespace:  bootstrap.GetNamespace(),
					Name:       bootstrap.GetName(),
				},
			},
			InfrastructureRef: corev1.ObjectReference{
				APIVersion: infrav1.GroupVersion.String(),
				Kind:       framework.TypeToKind(infraTemplate),
				Namespace:  infraTemplate.GetNamespace(),
				Name:       infraTemplate.GetName(),
			},
			Version: &version,
		},
	}

	machineDeployment := &clusterv1.MachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: clusterv1.MachineDeploymentSpec{
			ClusterName:             cluster.GetName(),
			Replicas:                &replicas,
			Template:                template,
			Strategy:                nil,
			MinReadySeconds:         nil,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
	}
	return machineDeployment, infraTemplate, bootstrap
}

type ClusterGenerator struct {
	counter int
}

func (c *ClusterGenerator) GenerateCluster(namespace string, replicas int32) (*clusterv1.Cluster, *infrav1.DockerCluster, *controlplanev1.KubeadmControlPlane, *infrav1.DockerMachineTemplate) {
	generatedName := fmt.Sprintf("test-%d", c.counter)
	c.counter++
	version := "1.16.3"

	infraCluster := &infrav1.DockerCluster{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
	}

	template := &infrav1.DockerMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: infrav1.DockerMachineTemplateSpec{},
	}

	kcp := &controlplanev1.KubeadmControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: controlplanev1.KubeadmControlPlaneSpec{
			Replicas: &replicas,
			Version:  version,
			InfrastructureTemplate: corev1.ObjectReference{
				Kind:       framework.TypeToKind(template),
				Namespace:  template.GetNamespace(),
				Name:       template.GetName(),
				APIVersion: infrav1.GroupVersion.String(),
			},
			KubeadmConfigSpec: bootstrapv1.KubeadmConfigSpec{
				ClusterConfiguration: &v1beta1.ClusterConfiguration{
					APIServer: v1beta1.APIServer{
						// Darwin support
						CertSANs: []string{"127.0.0.1"},
					},
				},
				InitConfiguration: &v1beta1.InitConfiguration{},
				JoinConfiguration: &v1beta1.JoinConfiguration{},
			},
		},
	}

	cluster := &clusterv1.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      generatedName,
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				Services: &clusterv1.NetworkRanges{CIDRBlocks: []string{}},
				Pods:     &clusterv1.NetworkRanges{CIDRBlocks: []string{"192.168.0.0/16"}},
			},
			InfrastructureRef: &corev1.ObjectReference{
				APIVersion: infrav1.GroupVersion.String(),
				Kind:       framework.TypeToKind(infraCluster),
				Namespace:  infraCluster.GetNamespace(),
				Name:       infraCluster.GetName(),
			},
			ControlPlaneRef: &corev1.ObjectReference{
				APIVersion: controlplanev1.GroupVersion.String(),
				Kind:       framework.TypeToKind(kcp),
				Namespace:  kcp.GetNamespace(),
				Name:       kcp.GetName(),
			},
		},
	}
	return cluster, infraCluster, kcp, template
}

func generateControlPlaneMachines(cluster *clusterv1.Cluster, namespace, version string, replicas int) ([]*clusterv1.Machine, []*bootstrapv1.KubeadmConfig, []*infrav1.DockerMachine) {
	machines := make([]*clusterv1.Machine, 0, replicas)
	bootstrap := make([]*bootstrapv1.KubeadmConfig, 0, replicas)
	infra := make([]*infrav1.DockerMachine, 0, replicas)
	for i := 0; i < replicas; i++ {
		bootstrap = append(bootstrap, &bootstrapv1.KubeadmConfig{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      fmt.Sprintf("bootstrap-controlplane-%d", i),
			},
			Spec: bootstrapv1.KubeadmConfigSpec{
				ClusterConfiguration: &v1beta1.ClusterConfiguration{
					APIServer: v1beta1.APIServer{
						// Darwin support
						CertSANs: []string{"127.0.0.1"},
					},
				},
				InitConfiguration: &v1beta1.InitConfiguration{},
				JoinConfiguration: &v1beta1.JoinConfiguration{},
			},
		})

		infra = append(infra, &infrav1.DockerMachine{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      fmt.Sprintf("controlplane-%d-infra", i),
			},
			Spec: infrav1.DockerMachineSpec{},
		})

		machines = append(machines, &clusterv1.Machine{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      fmt.Sprintf("controlplane-%d", i),
				Labels: map[string]string{
					clusterv1.MachineControlPlaneLabelName: "true",
				},
			},
			Spec: clusterv1.MachineSpec{
				ClusterName: cluster.GetName(),
				Bootstrap: clusterv1.Bootstrap{
					ConfigRef: &corev1.ObjectReference{
						APIVersion: bootstrapv1.GroupVersion.String(),
						Kind:       framework.TypeToKind(bootstrap[i]),
						Namespace:  bootstrap[i].GetNamespace(),
						Name:       bootstrap[i].GetName(),
					},
				},
				InfrastructureRef: corev1.ObjectReference{
					APIVersion: infrav1.GroupVersion.String(),
					Kind:       framework.TypeToKind(infra[i]),
					Namespace:  infra[i].GetNamespace(),
					Name:       infra[i].GetName(),
				},
				Version: &version,
			},
		})
	}

	return machines, bootstrap, infra
}
