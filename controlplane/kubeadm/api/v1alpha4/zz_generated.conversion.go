//go:build !ignore_autogenerated_kubeadm_controlplane
// +build !ignore_autogenerated_kubeadm_controlplane

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha4

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"

	apiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
	kubeadmapiv1alpha4 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1alpha4"
	v1beta1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
	errors "sigs.k8s.io/cluster-api/errors"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlane)(nil), (*v1beta1.KubeadmControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(a.(*KubeadmControlPlane), b.(*v1beta1.KubeadmControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeadmControlPlane)(nil), (*KubeadmControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(a.(*v1beta1.KubeadmControlPlane), b.(*KubeadmControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneList)(nil), (*v1beta1.KubeadmControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList(a.(*KubeadmControlPlaneList), b.(*v1beta1.KubeadmControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeadmControlPlaneList)(nil), (*KubeadmControlPlaneList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(a.(*v1beta1.KubeadmControlPlaneList), b.(*KubeadmControlPlaneList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneMachineTemplate)(nil), (*v1beta1.KubeadmControlPlaneMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate(a.(*KubeadmControlPlaneMachineTemplate), b.(*v1beta1.KubeadmControlPlaneMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneSpec)(nil), (*v1beta1.KubeadmControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec(a.(*KubeadmControlPlaneSpec), b.(*v1beta1.KubeadmControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneStatus)(nil), (*v1beta1.KubeadmControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus(a.(*KubeadmControlPlaneStatus), b.(*v1beta1.KubeadmControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneTemplate)(nil), (*v1beta1.KubeadmControlPlaneTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(a.(*KubeadmControlPlaneTemplate), b.(*v1beta1.KubeadmControlPlaneTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeadmControlPlaneTemplate)(nil), (*KubeadmControlPlaneTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(a.(*v1beta1.KubeadmControlPlaneTemplate), b.(*KubeadmControlPlaneTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneTemplateList)(nil), (*v1beta1.KubeadmControlPlaneTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList(a.(*KubeadmControlPlaneTemplateList), b.(*v1beta1.KubeadmControlPlaneTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeadmControlPlaneTemplateList)(nil), (*KubeadmControlPlaneTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList(a.(*v1beta1.KubeadmControlPlaneTemplateList), b.(*KubeadmControlPlaneTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneTemplateResource)(nil), (*v1beta1.KubeadmControlPlaneTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource(a.(*KubeadmControlPlaneTemplateResource), b.(*v1beta1.KubeadmControlPlaneTemplateResource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*KubeadmControlPlaneTemplateSpec)(nil), (*v1beta1.KubeadmControlPlaneTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec(a.(*KubeadmControlPlaneTemplateSpec), b.(*v1beta1.KubeadmControlPlaneTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.KubeadmControlPlaneTemplateSpec)(nil), (*KubeadmControlPlaneTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec(a.(*v1beta1.KubeadmControlPlaneTemplateSpec), b.(*KubeadmControlPlaneTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RollingUpdate)(nil), (*v1beta1.RollingUpdate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_RollingUpdate_To_v1beta1_RollingUpdate(a.(*RollingUpdate), b.(*v1beta1.RollingUpdate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RollingUpdate)(nil), (*RollingUpdate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RollingUpdate_To_v1alpha4_RollingUpdate(a.(*v1beta1.RollingUpdate), b.(*RollingUpdate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RolloutStrategy)(nil), (*v1beta1.RolloutStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_RolloutStrategy_To_v1beta1_RolloutStrategy(a.(*RolloutStrategy), b.(*v1beta1.RolloutStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1beta1.RolloutStrategy)(nil), (*RolloutStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RolloutStrategy_To_v1alpha4_RolloutStrategy(a.(*v1beta1.RolloutStrategy), b.(*RolloutStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*KubeadmControlPlaneSpec)(nil), (*v1beta1.KubeadmControlPlaneTemplateResourceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneTemplateResourceSpec(a.(*KubeadmControlPlaneSpec), b.(*v1beta1.KubeadmControlPlaneTemplateResourceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.KubeadmControlPlaneMachineTemplate)(nil), (*KubeadmControlPlaneMachineTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneMachineTemplate_To_v1alpha4_KubeadmControlPlaneMachineTemplate(a.(*v1beta1.KubeadmControlPlaneMachineTemplate), b.(*KubeadmControlPlaneMachineTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.KubeadmControlPlaneSpec)(nil), (*KubeadmControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(a.(*v1beta1.KubeadmControlPlaneSpec), b.(*KubeadmControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.KubeadmControlPlaneStatus)(nil), (*KubeadmControlPlaneStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(a.(*v1beta1.KubeadmControlPlaneStatus), b.(*KubeadmControlPlaneStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.KubeadmControlPlaneTemplateResourceSpec)(nil), (*KubeadmControlPlaneSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneTemplateResourceSpec_To_v1alpha4_KubeadmControlPlaneSpec(a.(*v1beta1.KubeadmControlPlaneTemplateResourceSpec), b.(*KubeadmControlPlaneSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.KubeadmControlPlaneTemplateResource)(nil), (*KubeadmControlPlaneTemplateResource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_KubeadmControlPlaneTemplateResource_To_v1alpha4_KubeadmControlPlaneTemplateResource(a.(*v1beta1.KubeadmControlPlaneTemplateResource), b.(*KubeadmControlPlaneTemplateResource), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(in *KubeadmControlPlane, out *v1beta1.KubeadmControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(in *KubeadmControlPlane, out *v1beta1.KubeadmControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in *v1beta1.KubeadmControlPlane, out *KubeadmControlPlane, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane is an autogenerated conversion function.
func Convert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in *v1beta1.KubeadmControlPlane, out *KubeadmControlPlane, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList(in *KubeadmControlPlaneList, out *v1beta1.KubeadmControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.KubeadmControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList(in *KubeadmControlPlaneList, out *v1beta1.KubeadmControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in *v1beta1.KubeadmControlPlaneList, out *KubeadmControlPlaneList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmControlPlane, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList is an autogenerated conversion function.
func Convert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in *v1beta1.KubeadmControlPlaneList, out *KubeadmControlPlaneList, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate(in *KubeadmControlPlaneMachineTemplate, out *v1beta1.KubeadmControlPlaneMachineTemplate, s conversion.Scope) error {
	if err := apiv1alpha4.Convert_v1alpha4_ObjectMeta_To_v1beta1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.InfrastructureRef = in.InfrastructureRef
	out.NodeDrainTimeout = (*v1.Duration)(unsafe.Pointer(in.NodeDrainTimeout))
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate(in *KubeadmControlPlaneMachineTemplate, out *v1beta1.KubeadmControlPlaneMachineTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneMachineTemplate_To_v1alpha4_KubeadmControlPlaneMachineTemplate(in *v1beta1.KubeadmControlPlaneMachineTemplate, out *KubeadmControlPlaneMachineTemplate, s conversion.Scope) error {
	if err := apiv1alpha4.Convert_v1beta1_ObjectMeta_To_v1alpha4_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, s); err != nil {
		return err
	}
	out.InfrastructureRef = in.InfrastructureRef
	out.NodeDrainTimeout = (*v1.Duration)(unsafe.Pointer(in.NodeDrainTimeout))
	// WARNING: in.NodeVolumeDetachTimeout requires manual conversion: does not exist in peer-type
	// WARNING: in.NodeDeletionTimeout requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec(in *KubeadmControlPlaneSpec, out *v1beta1.KubeadmControlPlaneSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Version = in.Version
	if err := Convert_v1alpha4_KubeadmControlPlaneMachineTemplate_To_v1beta1_KubeadmControlPlaneMachineTemplate(&in.MachineTemplate, &out.MachineTemplate, s); err != nil {
		return err
	}
	if err := kubeadmapiv1alpha4.Convert_v1alpha4_KubeadmConfigSpec_To_v1beta1_KubeadmConfigSpec(&in.KubeadmConfigSpec, &out.KubeadmConfigSpec, s); err != nil {
		return err
	}
	out.RolloutAfter = (*v1.Time)(unsafe.Pointer(in.RolloutAfter))
	out.RolloutStrategy = (*v1beta1.RolloutStrategy)(unsafe.Pointer(in.RolloutStrategy))
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec(in *KubeadmControlPlaneSpec, out *v1beta1.KubeadmControlPlaneSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneSpec(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneSpec_To_v1alpha4_KubeadmControlPlaneSpec(in *v1beta1.KubeadmControlPlaneSpec, out *KubeadmControlPlaneSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Version = in.Version
	if err := Convert_v1beta1_KubeadmControlPlaneMachineTemplate_To_v1alpha4_KubeadmControlPlaneMachineTemplate(&in.MachineTemplate, &out.MachineTemplate, s); err != nil {
		return err
	}
	if err := kubeadmapiv1alpha4.Convert_v1beta1_KubeadmConfigSpec_To_v1alpha4_KubeadmConfigSpec(&in.KubeadmConfigSpec, &out.KubeadmConfigSpec, s); err != nil {
		return err
	}
	// WARNING: in.RolloutBefore requires manual conversion: does not exist in peer-type
	out.RolloutAfter = (*v1.Time)(unsafe.Pointer(in.RolloutAfter))
	out.RolloutStrategy = (*RolloutStrategy)(unsafe.Pointer(in.RolloutStrategy))
	// WARNING: in.RemediationStrategy requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus(in *KubeadmControlPlaneStatus, out *v1beta1.KubeadmControlPlaneStatus, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Replicas = in.Replicas
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.UpdatedReplicas = in.UpdatedReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureReason = errors.KubeadmControlPlaneStatusError(in.FailureReason)
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1beta1.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha4.Convert_v1alpha4_Condition_To_v1beta1_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus(in *KubeadmControlPlaneStatus, out *v1beta1.KubeadmControlPlaneStatus, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneStatus_To_v1beta1_KubeadmControlPlaneStatus(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneStatus_To_v1alpha4_KubeadmControlPlaneStatus(in *v1beta1.KubeadmControlPlaneStatus, out *KubeadmControlPlaneStatus, s conversion.Scope) error {
	out.Selector = in.Selector
	out.Replicas = in.Replicas
	out.Version = (*string)(unsafe.Pointer(in.Version))
	out.UpdatedReplicas = in.UpdatedReplicas
	out.ReadyReplicas = in.ReadyReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	out.Initialized = in.Initialized
	out.Ready = in.Ready
	out.FailureReason = errors.KubeadmControlPlaneStatusError(in.FailureReason)
	out.FailureMessage = (*string)(unsafe.Pointer(in.FailureMessage))
	out.ObservedGeneration = in.ObservedGeneration
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(apiv1alpha4.Conditions, len(*in))
		for i := range *in {
			if err := apiv1alpha4.Convert_v1beta1_Condition_To_v1alpha4_Condition(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	// WARNING: in.LastRemediation requires manual conversion: does not exist in peer-type
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(in *KubeadmControlPlaneTemplate, out *v1beta1.KubeadmControlPlaneTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(in *KubeadmControlPlaneTemplate, out *v1beta1.KubeadmControlPlaneTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(in *v1beta1.KubeadmControlPlaneTemplate, out *KubeadmControlPlaneTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate is an autogenerated conversion function.
func Convert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(in *v1beta1.KubeadmControlPlaneTemplate, out *KubeadmControlPlaneTemplate, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList(in *KubeadmControlPlaneTemplateList, out *v1beta1.KubeadmControlPlaneTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]v1beta1.KubeadmControlPlaneTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList(in *KubeadmControlPlaneTemplateList, out *v1beta1.KubeadmControlPlaneTemplateList, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList(in *v1beta1.KubeadmControlPlaneTemplateList, out *KubeadmControlPlaneTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeadmControlPlaneTemplate, len(*in))
		for i := range *in {
			if err := Convert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList is an autogenerated conversion function.
func Convert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList(in *v1beta1.KubeadmControlPlaneTemplateList, out *KubeadmControlPlaneTemplateList, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList(in, out, s)
}

func autoConvert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource(in *KubeadmControlPlaneTemplateResource, out *v1beta1.KubeadmControlPlaneTemplateResource, s conversion.Scope) error {
	if err := Convert_v1alpha4_KubeadmControlPlaneSpec_To_v1beta1_KubeadmControlPlaneTemplateResourceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource(in *KubeadmControlPlaneTemplateResource, out *v1beta1.KubeadmControlPlaneTemplateResource, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneTemplateResource_To_v1alpha4_KubeadmControlPlaneTemplateResource(in *v1beta1.KubeadmControlPlaneTemplateResource, out *KubeadmControlPlaneTemplateResource, s conversion.Scope) error {
	// WARNING: in.ObjectMeta requires manual conversion: does not exist in peer-type
	if err := Convert_v1beta1_KubeadmControlPlaneTemplateResourceSpec_To_v1alpha4_KubeadmControlPlaneSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec(in *KubeadmControlPlaneTemplateSpec, out *v1beta1.KubeadmControlPlaneTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1alpha4_KubeadmControlPlaneTemplateResource_To_v1beta1_KubeadmControlPlaneTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec is an autogenerated conversion function.
func Convert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec(in *KubeadmControlPlaneTemplateSpec, out *v1beta1.KubeadmControlPlaneTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1alpha4_KubeadmControlPlaneTemplateSpec_To_v1beta1_KubeadmControlPlaneTemplateSpec(in, out, s)
}

func autoConvert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec(in *v1beta1.KubeadmControlPlaneTemplateSpec, out *KubeadmControlPlaneTemplateSpec, s conversion.Scope) error {
	if err := Convert_v1beta1_KubeadmControlPlaneTemplateResource_To_v1alpha4_KubeadmControlPlaneTemplateResource(&in.Template, &out.Template, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec is an autogenerated conversion function.
func Convert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec(in *v1beta1.KubeadmControlPlaneTemplateSpec, out *KubeadmControlPlaneTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_KubeadmControlPlaneTemplateSpec_To_v1alpha4_KubeadmControlPlaneTemplateSpec(in, out, s)
}

func autoConvert_v1alpha4_RollingUpdate_To_v1beta1_RollingUpdate(in *RollingUpdate, out *v1beta1.RollingUpdate, s conversion.Scope) error {
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	return nil
}

// Convert_v1alpha4_RollingUpdate_To_v1beta1_RollingUpdate is an autogenerated conversion function.
func Convert_v1alpha4_RollingUpdate_To_v1beta1_RollingUpdate(in *RollingUpdate, out *v1beta1.RollingUpdate, s conversion.Scope) error {
	return autoConvert_v1alpha4_RollingUpdate_To_v1beta1_RollingUpdate(in, out, s)
}

func autoConvert_v1beta1_RollingUpdate_To_v1alpha4_RollingUpdate(in *v1beta1.RollingUpdate, out *RollingUpdate, s conversion.Scope) error {
	out.MaxSurge = (*intstr.IntOrString)(unsafe.Pointer(in.MaxSurge))
	return nil
}

// Convert_v1beta1_RollingUpdate_To_v1alpha4_RollingUpdate is an autogenerated conversion function.
func Convert_v1beta1_RollingUpdate_To_v1alpha4_RollingUpdate(in *v1beta1.RollingUpdate, out *RollingUpdate, s conversion.Scope) error {
	return autoConvert_v1beta1_RollingUpdate_To_v1alpha4_RollingUpdate(in, out, s)
}

func autoConvert_v1alpha4_RolloutStrategy_To_v1beta1_RolloutStrategy(in *RolloutStrategy, out *v1beta1.RolloutStrategy, s conversion.Scope) error {
	out.Type = v1beta1.RolloutStrategyType(in.Type)
	out.RollingUpdate = (*v1beta1.RollingUpdate)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_v1alpha4_RolloutStrategy_To_v1beta1_RolloutStrategy is an autogenerated conversion function.
func Convert_v1alpha4_RolloutStrategy_To_v1beta1_RolloutStrategy(in *RolloutStrategy, out *v1beta1.RolloutStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha4_RolloutStrategy_To_v1beta1_RolloutStrategy(in, out, s)
}

func autoConvert_v1beta1_RolloutStrategy_To_v1alpha4_RolloutStrategy(in *v1beta1.RolloutStrategy, out *RolloutStrategy, s conversion.Scope) error {
	out.Type = RolloutStrategyType(in.Type)
	out.RollingUpdate = (*RollingUpdate)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_v1beta1_RolloutStrategy_To_v1alpha4_RolloutStrategy is an autogenerated conversion function.
func Convert_v1beta1_RolloutStrategy_To_v1alpha4_RolloutStrategy(in *v1beta1.RolloutStrategy, out *RolloutStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_RolloutStrategy_To_v1alpha4_RolloutStrategy(in, out, s)
}
