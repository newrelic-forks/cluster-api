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

package machinefilters

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1alpha3"
)

type Func func(machine *clusterv1.Machine) bool

// And returns a filter that returns true if all of the given filters returns true.
func And(filters ...Func) Func {
	return func(machine *clusterv1.Machine) bool {
		for _, f := range filters {
			if !f(machine) {
				return false
			}
		}
		return true
	}
}

// Or returns a filter that returns true if any of the given filters returns true.
func Or(filters ...Func) Func {
	return func(machine *clusterv1.Machine) bool {
		for _, f := range filters {
			if f(machine) {
				return true
			}
		}
		return false
	}
}

// Not returns a filter that returns the opposite of the given filter.
func Not(mf Func) Func {
	return func(machine *clusterv1.Machine) bool {
		return !mf(machine)
	}
}

// IsNil is a MachineFilter that returns true if the machine is nil
var IsNil = func(machine *clusterv1.Machine) bool {
	return machine == nil
}

// HasControllerRef is a MachineFilter that returns true if the machine has a controller ref
var HasControllerRef = func(machine *clusterv1.Machine) bool {
	return metav1.GetControllerOf(machine) != nil
}

// InFailureDomains returns a MachineFilter function to find all machines
// in any of the given failure domains
func InFailureDomains(failureDomains ...*string) Func {
	return func(machine *clusterv1.Machine) bool {
		if machine == nil {
			return false
		}
		for i := range failureDomains {
			fd := failureDomains[i]
			if fd == nil {
				if fd == machine.Spec.FailureDomain {
					return true
				}
				continue
			}
			if machine.Spec.FailureDomain == nil {
				continue
			}
			if *fd == *machine.Spec.FailureDomain {
				return true
			}
		}
		return false
	}
}

// OwnedMachines returns a MachineFilter function to find all owned control plane machines.
// Usage: managementCluster.GetMachinesForCluster(ctx, cluster, OwnedMachines(controlPlane))
func OwnedMachines(owner metav1.Object) func(machine *clusterv1.Machine) bool {
	return func(machine *clusterv1.Machine) bool {
		if machine == nil {
			return false
		}
		controllerRef := metav1.GetControllerOf(machine)
		if controllerRef == nil {
			return false
		}
		return controllerRef.Kind == "KubeadmControlPlane" && controllerRef.Name == owner.GetName() && controllerRef.UID == owner.GetUID()
	}
}

func ControlPlaneMachines(clusterName string) func(machine *clusterv1.Machine) bool {
	selector := ControlPlaneSelectorForCluster(clusterName)
	return func(machine *clusterv1.Machine) bool {
		if machine == nil {
			return false
		}
		return selector.Matches(labels.Set(machine.Labels))
	}
}

// AdoptableControlPlaneMachines returns a MachineFilter function to find all un-controlled control plane machines.
// Usage: managementCluster.GetMachinesForCluster(ctx, cluster, AdoptableControlPlaneMachines(cluster.Name, controlPlane))
func AdoptableControlPlaneMachines(clusterName string) func(machine *clusterv1.Machine) bool {
	return And(
		Not(IsNil),
		ControlPlaneMachines(clusterName),
		Not(HasControllerRef),
	)
}

// HasDeletionTimestamp is a MachineFilter to find all machines
// that have a deletion timestamp.
func HasDeletionTimestamp(machine *clusterv1.Machine) bool {
	if machine == nil {
		return false
	}
	return !machine.DeletionTimestamp.IsZero()
}

// MatchesConfigurationHash returns a filter to find all machines
// that match a given KubeadmControlPlane configuration hash.
func MatchesConfigurationHash(configHash string) Func {
	return func(machine *clusterv1.Machine) bool {
		if machine == nil {
			return false
		}
		if hash, ok := machine.Labels[controlplanev1.KubeadmControlPlaneHashLabelKey]; ok {
			return hash == configHash
		}
		return false
	}
}

// OlderThan returns a filter to find all machines
// that have a CreationTimestamp earlier than the given time.
func OlderThan(t *metav1.Time) Func {
	return func(machine *clusterv1.Machine) bool {
		if machine == nil {
			return false
		}
		return machine.CreationTimestamp.Before(t)
	}
}

// HasAnnotationKey returns a filter to find all machines that have the
// specified Annotation key present
func HasAnnotationKey(key string) Func {
	return func(machine *clusterv1.Machine) bool {
		if machine == nil || machine.Annotations == nil {
			return false
		}
		if _, ok := machine.Annotations[key]; ok {
			return true
		}
		return false
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
