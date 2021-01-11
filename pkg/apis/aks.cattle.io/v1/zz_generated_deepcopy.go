// +build !ignore_autogenerated

/*
Copyright 2019 Wrangler Sample Controller Authors

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfig) DeepCopyInto(out *AKSClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfig.
func (in *AKSClusterConfig) DeepCopy() *AKSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigList) DeepCopyInto(out *AKSClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKSClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigList.
func (in *AKSClusterConfigList) DeepCopy() *AKSClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigSpec) DeepCopyInto(out *AKSClusterConfigSpec) {
	*out = *in
	if in.DNSPrefix != nil {
		in, out := &in.DNSPrefix, &out.DNSPrefix
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]AKSNodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateCluster != nil {
		in, out := &in.PrivateCluster, &out.PrivateCluster
		*out = new(bool)
		**out = **in
	}
	if in.AuthorizedIPRanges != nil {
		in, out := &in.AuthorizedIPRanges, &out.AuthorizedIPRanges
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigSpec.
func (in *AKSClusterConfigSpec) DeepCopy() *AKSClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigStatus) DeepCopyInto(out *AKSClusterConfigStatus) {
	*out = *in
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigStatus.
func (in *AKSClusterConfigStatus) DeepCopy() *AKSClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSNodePool) DeepCopyInto(out *AKSNodePool) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int32)
		**out = **in
	}
	if in.MaxPods != nil {
		in, out := &in.MaxPods, &out.MaxPods
		*out = new(int32)
		**out = **in
	}
	if in.OsDiskSizeGB != nil {
		in, out := &in.OsDiskSizeGB, &out.OsDiskSizeGB
		*out = new(int32)
		**out = **in
	}
	if in.OrchestratorVersion != nil {
		in, out := &in.OrchestratorVersion, &out.OrchestratorVersion
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int32)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int32)
		**out = **in
	}
	if in.EnableAutoScaling != nil {
		in, out := &in.EnableAutoScaling, &out.EnableAutoScaling
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSNodePool.
func (in *AKSNodePool) DeepCopy() *AKSNodePool {
	if in == nil {
		return nil
	}
	out := new(AKSNodePool)
	in.DeepCopyInto(out)
	return out
}
