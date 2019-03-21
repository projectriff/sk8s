// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AdmissionHookClientConfig).DeepCopyInto(out.(*AdmissionHookClientConfig))
			return nil
		}, InType: reflect.TypeOf(&AdmissionHookClientConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHook).DeepCopyInto(out.(*ExternalAdmissionHook))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHook{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHookConfiguration).DeepCopyInto(out.(*ExternalAdmissionHookConfiguration))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHookConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExternalAdmissionHookConfigurationList).DeepCopyInto(out.(*ExternalAdmissionHookConfigurationList))
			return nil
		}, InType: reflect.TypeOf(&ExternalAdmissionHookConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Initializer).DeepCopyInto(out.(*Initializer))
			return nil
		}, InType: reflect.TypeOf(&Initializer{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InitializerConfiguration).DeepCopyInto(out.(*InitializerConfiguration))
			return nil
		}, InType: reflect.TypeOf(&InitializerConfiguration{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InitializerConfigurationList).DeepCopyInto(out.(*InitializerConfigurationList))
			return nil
		}, InType: reflect.TypeOf(&InitializerConfigurationList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Rule).DeepCopyInto(out.(*Rule))
			return nil
		}, InType: reflect.TypeOf(&Rule{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RuleWithOperations).DeepCopyInto(out.(*RuleWithOperations))
			return nil
		}, InType: reflect.TypeOf(&RuleWithOperations{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ServiceReference).DeepCopyInto(out.(*ServiceReference))
			return nil
		}, InType: reflect.TypeOf(&ServiceReference{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionHookClientConfig) DeepCopyInto(out *AdmissionHookClientConfig) {
	*out = *in
	out.Service = in.Service
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionHookClientConfig.
func (in *AdmissionHookClientConfig) DeepCopy() *AdmissionHookClientConfig {
	if in == nil {
		return nil
	}
	out := new(AdmissionHookClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHook) DeepCopyInto(out *ExternalAdmissionHook) {
	*out = *in
	in.ClientConfig.DeepCopyInto(&out.ClientConfig)
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]RuleWithOperations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(FailurePolicyType)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHook.
func (in *ExternalAdmissionHook) DeepCopy() *ExternalAdmissionHook {
	if in == nil {
		return nil
	}
	out := new(ExternalAdmissionHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHookConfiguration) DeepCopyInto(out *ExternalAdmissionHookConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ExternalAdmissionHooks != nil {
		in, out := &in.ExternalAdmissionHooks, &out.ExternalAdmissionHooks
		*out = make([]ExternalAdmissionHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHookConfiguration.
func (in *ExternalAdmissionHookConfiguration) DeepCopy() *ExternalAdmissionHookConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExternalAdmissionHookConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalAdmissionHookConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAdmissionHookConfigurationList) DeepCopyInto(out *ExternalAdmissionHookConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalAdmissionHookConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAdmissionHookConfigurationList.
func (in *ExternalAdmissionHookConfigurationList) DeepCopy() *ExternalAdmissionHookConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ExternalAdmissionHookConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalAdmissionHookConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Initializer) DeepCopyInto(out *Initializer) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Initializer.
func (in *Initializer) DeepCopy() *Initializer {
	if in == nil {
		return nil
	}
	out := new(Initializer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerConfiguration) DeepCopyInto(out *InitializerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Initializers != nil {
		in, out := &in.Initializers, &out.Initializers
		*out = make([]Initializer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerConfiguration.
func (in *InitializerConfiguration) DeepCopy() *InitializerConfiguration {
	if in == nil {
		return nil
	}
	out := new(InitializerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitializerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitializerConfigurationList) DeepCopyInto(out *InitializerConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InitializerConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitializerConfigurationList.
func (in *InitializerConfigurationList) DeepCopy() *InitializerConfigurationList {
	if in == nil {
		return nil
	}
	out := new(InitializerConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InitializerConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIVersions != nil {
		in, out := &in.APIVersions, &out.APIVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleWithOperations) DeepCopyInto(out *RuleWithOperations) {
	*out = *in
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]OperationType, len(*in))
		copy(*out, *in)
	}
	in.Rule.DeepCopyInto(&out.Rule)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleWithOperations.
func (in *RuleWithOperations) DeepCopy() *RuleWithOperations {
	if in == nil {
		return nil
	}
	out := new(RuleWithOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceReference) DeepCopyInto(out *ServiceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceReference.
func (in *ServiceReference) DeepCopy() *ServiceReference {
	if in == nil {
		return nil
	}
	out := new(ServiceReference)
	in.DeepCopyInto(out)
	return out
}
