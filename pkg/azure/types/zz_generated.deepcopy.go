// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package types

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureInterface) DeepCopyInto(out *AzureInterface) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]AzureAddress, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureInterface.
func (in *AzureInterface) DeepCopy() *AzureInterface {
	if in == nil {
		return nil
	}
	out := new(AzureInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSpec) DeepCopyInto(out *AzureSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSpec.
func (in *AzureSpec) DeepCopy() *AzureSpec {
	if in == nil {
		return nil
	}
	out := new(AzureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureStatus) DeepCopyInto(out *AzureStatus) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]AzureInterface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureStatus.
func (in *AzureStatus) DeepCopy() *AzureStatus {
	if in == nil {
		return nil
	}
	out := new(AzureStatus)
	in.DeepCopyInto(out)
	return out
}
