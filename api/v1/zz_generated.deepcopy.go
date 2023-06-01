//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
(c) Copyright IBM Corp.
(c) Copyright Instana Inc.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPodSpec) DeepCopyInto(out *AgentPodSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
	in.ResourceRequirements.DeepCopyInto(&out.ResourceRequirements)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPodSpec.
func (in *AgentPodSpec) DeepCopy() *AgentPodSpec {
	if in == nil {
		return nil
	}
	out := new(AgentPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseAgentSpec) DeepCopyInto(out *BaseAgentSpec) {
	*out = *in
	if in.AdditionalBackends != nil {
		in, out := &in.AdditionalBackends, &out.AdditionalBackends
		*out = make([]BackendSpec, len(*in))
		copy(*out, *in)
	}
	in.TlsSpec.DeepCopyInto(&out.TlsSpec)
	in.ExtendedImageSpec.DeepCopyInto(&out.ExtendedImageSpec)
	in.UpdateStrategy.DeepCopyInto(&out.UpdateStrategy)
	in.Pod.DeepCopyInto(&out.Pod)
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.Configuration = in.Configuration
	out.Host = in.Host
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseAgentSpec.
func (in *BaseAgentSpec) DeepCopy() *BaseAgentSpec {
	if in == nil {
		return nil
	}
	out := new(BaseAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Create) DeepCopyInto(out *Create) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Create.
func (in *Create) DeepCopy() *Create {
	if in == nil {
		return nil
	}
	out := new(Create)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Enabled) DeepCopyInto(out *Enabled) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Enabled.
func (in *Enabled) DeepCopy() *Enabled {
	if in == nil {
		return nil
	}
	out := new(Enabled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExtendedImageSpec) DeepCopyInto(out *ExtendedImageSpec) {
	*out = *in
	out.ImageSpec = in.ImageSpec
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtendedImageSpec.
func (in *ExtendedImageSpec) DeepCopy() *ExtendedImageSpec {
	if in == nil {
		return nil
	}
	out := new(ExtendedImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostSpec) DeepCopyInto(out *HostSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostSpec.
func (in *HostSpec) DeepCopy() *HostSpec {
	if in == nil {
		return nil
	}
	out := new(HostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanaAgent) DeepCopyInto(out *InstanaAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanaAgent.
func (in *InstanaAgent) DeepCopy() *InstanaAgent {
	if in == nil {
		return nil
	}
	out := new(InstanaAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanaAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanaAgentList) DeepCopyInto(out *InstanaAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InstanaAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanaAgentList.
func (in *InstanaAgentList) DeepCopy() *InstanaAgentList {
	if in == nil {
		return nil
	}
	out := new(InstanaAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanaAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanaAgentSpec) DeepCopyInto(out *InstanaAgentSpec) {
	*out = *in
	in.Agent.DeepCopyInto(&out.Agent)
	out.Cluster = in.Cluster
	out.Zone = in.Zone
	if in.OpenShift != nil {
		in, out := &in.OpenShift, &out.OpenShift
		*out = new(bool)
		**out = **in
	}
	out.Rbac = in.Rbac
	out.Service = in.Service
	in.OpenTelemetry.DeepCopyInto(&out.OpenTelemetry)
	in.Prometheus.DeepCopyInto(&out.Prometheus)
	out.ServiceAccountSpec = in.ServiceAccountSpec
	in.PodSecurityPolicySpec.DeepCopyInto(&out.PodSecurityPolicySpec)
	in.KubernetesSpec.DeepCopyInto(&out.KubernetesSpec)
	in.K8sSensor.DeepCopyInto(&out.K8sSensor)
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]Zone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanaAgentSpec.
func (in *InstanaAgentSpec) DeepCopy() *InstanaAgentSpec {
	if in == nil {
		return nil
	}
	out := new(InstanaAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanaAgentStatus) DeepCopyInto(out *InstanaAgentStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	out.ConfigMap = in.ConfigMap
	out.DaemonSet = in.DaemonSet
	if in.LeadingAgentPod != nil {
		in, out := &in.LeadingAgentPod, &out.LeadingAgentPod
		*out = make(map[string]ResourceInfo, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanaAgentStatus.
func (in *InstanaAgentStatus) DeepCopy() *InstanaAgentStatus {
	if in == nil {
		return nil
	}
	out := new(InstanaAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sSpec) DeepCopyInto(out *K8sSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
	out.ImageSpec = in.ImageSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sSpec.
func (in *K8sSpec) DeepCopy() *K8sSpec {
	if in == nil {
		return nil
	}
	out := new(K8sSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesDeploymentSpec) DeepCopyInto(out *KubernetesDeploymentSpec) {
	*out = *in
	in.Enabled.DeepCopyInto(&out.Enabled)
	in.Pod.DeepCopyInto(&out.Pod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesDeploymentSpec.
func (in *KubernetesDeploymentSpec) DeepCopy() *KubernetesDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesSpec) DeepCopyInto(out *KubernetesSpec) {
	*out = *in
	in.DeploymentSpec.DeepCopyInto(&out.DeploymentSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesSpec.
func (in *KubernetesSpec) DeepCopy() *KubernetesSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Name) DeepCopyInto(out *Name) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Name.
func (in *Name) DeepCopy() *Name {
	if in == nil {
		return nil
	}
	out := new(Name)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenTelemetry) DeepCopyInto(out *OpenTelemetry) {
	*out = *in
	in.Enabled.DeepCopyInto(&out.Enabled)
	if in.GRPC != nil {
		in, out := &in.GRPC, &out.GRPC
		*out = new(Enabled)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = new(Enabled)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenTelemetry.
func (in *OpenTelemetry) DeepCopy() *OpenTelemetry {
	if in == nil {
		return nil
	}
	out := new(OpenTelemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSecurityPolicySpec) DeepCopyInto(out *PodSecurityPolicySpec) {
	*out = *in
	in.Enabled.DeepCopyInto(&out.Enabled)
	out.Name = in.Name
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSecurityPolicySpec.
func (in *PodSecurityPolicySpec) DeepCopy() *PodSecurityPolicySpec {
	if in == nil {
		return nil
	}
	out := new(PodSecurityPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
	in.RemoteWrite.DeepCopyInto(&out.RemoteWrite)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInfo) DeepCopyInto(out *ResourceInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInfo.
func (in *ResourceInfo) DeepCopy() *ResourceInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSpec) DeepCopyInto(out *ServiceAccountSpec) {
	*out = *in
	out.Create = in.Create
	out.Name = in.Name
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSpec.
func (in *ServiceAccountSpec) DeepCopy() *ServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsSpec) DeepCopyInto(out *TlsSpec) {
	*out = *in
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsSpec.
func (in *TlsSpec) DeepCopy() *TlsSpec {
	if in == nil {
		return nil
	}
	out := new(TlsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	out.Name = in.Name
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}
