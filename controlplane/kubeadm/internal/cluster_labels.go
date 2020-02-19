/*
Copyright 2020 The Kubernetes Authors.

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

package internal

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"
)

// ControlPlaneLabelsForClusterWithHash returns a set of labels to add to a control plane machine for this specific
// cluster and configuration hash
func ControlPlaneLabelsForClusterWithHash(clusterName string, hash string) map[string]string {
	labels := ControlPlaneLabelsForCluster(clusterName)
	labels[controlplanev1.KubeadmControlPlaneHashLabelKey] = hash
	return labels
}

// ControlPlaneLabelsForCluster returns a set of labels to add to a control plane machine for this specific cluster.
func ControlPlaneLabelsForCluster(clusterName string) map[string]string {
	return map[string]string{
		clusterv1.ClusterLabelName:             clusterName,
		clusterv1.MachineControlPlaneLabelName: "",
	}
}

// ControlPlaneSelectorForCluster returns the label selector necessary to get control plane machines for a given cluster.
func ControlPlaneSelectorForCluster(clusterName string) labels.Selector {
	must := func(r *labels.Requirement, err error) *labels.Requirement {
		if err != nil {
			panic(err)
		}
		return r
	}
	return labels.NewSelector().Add(
		*must(labels.NewRequirement(clusterv1.ClusterLabelName, selection.Equals, []string{clusterName})),
		*must(labels.NewRequirement(clusterv1.MachineControlPlaneLabelName, selection.Exists, []string{})),
	)
}
