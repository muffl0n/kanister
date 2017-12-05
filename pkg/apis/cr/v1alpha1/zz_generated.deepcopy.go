// +build !ignore_autogenerated

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
			in.(*ActionSet).DeepCopyInto(out.(*ActionSet))
			return nil
		}, InType: reflect.TypeOf(&ActionSet{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActionSetList).DeepCopyInto(out.(*ActionSetList))
			return nil
		}, InType: reflect.TypeOf(&ActionSetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActionSetSpec).DeepCopyInto(out.(*ActionSetSpec))
			return nil
		}, InType: reflect.TypeOf(&ActionSetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActionSetStatus).DeepCopyInto(out.(*ActionSetStatus))
			return nil
		}, InType: reflect.TypeOf(&ActionSetStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActionSpec).DeepCopyInto(out.(*ActionSpec))
			return nil
		}, InType: reflect.TypeOf(&ActionSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActionStatus).DeepCopyInto(out.(*ActionStatus))
			return nil
		}, InType: reflect.TypeOf(&ActionStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Blueprint).DeepCopyInto(out.(*Blueprint))
			return nil
		}, InType: reflect.TypeOf(&Blueprint{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BlueprintAction).DeepCopyInto(out.(*BlueprintAction))
			return nil
		}, InType: reflect.TypeOf(&BlueprintAction{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BlueprintList).DeepCopyInto(out.(*BlueprintList))
			return nil
		}, InType: reflect.TypeOf(&BlueprintList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*BlueprintPhase).DeepCopyInto(out.(*BlueprintPhase))
			return nil
		}, InType: reflect.TypeOf(&BlueprintPhase{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ObjectReference).DeepCopyInto(out.(*ObjectReference))
			return nil
		}, InType: reflect.TypeOf(&ObjectReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Phase).DeepCopyInto(out.(*Phase))
			return nil
		}, InType: reflect.TypeOf(&Phase{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSet) DeepCopyInto(out *ActionSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		if *in == nil {
			*out = nil
		} else {
			*out = new(ActionSetSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(ActionSetStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSet.
func (in *ActionSet) DeepCopy() *ActionSet {
	if in == nil {
		return nil
	}
	out := new(ActionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSetList) DeepCopyInto(out *ActionSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*ActionSet, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(ActionSet)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSetList.
func (in *ActionSetList) DeepCopy() *ActionSetList {
	if in == nil {
		return nil
	}
	out := new(ActionSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActionSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSetSpec) DeepCopyInto(out *ActionSetSpec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ActionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSetSpec.
func (in *ActionSetSpec) DeepCopy() *ActionSetSpec {
	if in == nil {
		return nil
	}
	out := new(ActionSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSetStatus) DeepCopyInto(out *ActionSetStatus) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ActionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSetStatus.
func (in *ActionSetStatus) DeepCopy() *ActionSetStatus {
	if in == nil {
		return nil
	}
	out := new(ActionSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSpec) DeepCopyInto(out *ActionSpec) {
	*out = *in
	out.Object = in.Object
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make(map[string]Artifact, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make(map[string]ObjectReference, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]ObjectReference, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSpec.
func (in *ActionSpec) DeepCopy() *ActionSpec {
	if in == nil {
		return nil
	}
	out := new(ActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionStatus) DeepCopyInto(out *ActionStatus) {
	*out = *in
	out.Object = in.Object
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]Phase, len(*in))
		copy(*out, *in)
	}
	if in.Artifacts != nil {
		in, out := &in.Artifacts, &out.Artifacts
		*out = make(map[string]Artifact, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionStatus.
func (in *ActionStatus) DeepCopy() *ActionStatus {
	if in == nil {
		return nil
	}
	out := new(ActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blueprint) DeepCopyInto(out *Blueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make(map[string]*BlueprintAction, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(BlueprintAction)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blueprint.
func (in *Blueprint) DeepCopy() *Blueprint {
	if in == nil {
		return nil
	}
	out := new(Blueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Blueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintAction) DeepCopyInto(out *BlueprintAction) {
	*out = *in
	if in.ConfigMapNames != nil {
		in, out := &in.ConfigMapNames, &out.ConfigMapNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecretNames != nil {
		in, out := &in.SecretNames, &out.SecretNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InputArtifactNames != nil {
		in, out := &in.InputArtifactNames, &out.InputArtifactNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OutputArtifacts != nil {
		in, out := &in.OutputArtifacts, &out.OutputArtifacts
		*out = make(map[string]Artifact, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]BlueprintPhase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintAction.
func (in *BlueprintAction) DeepCopy() *BlueprintAction {
	if in == nil {
		return nil
	}
	out := new(BlueprintAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintList) DeepCopyInto(out *BlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]*Blueprint, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Blueprint)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintList.
func (in *BlueprintList) DeepCopy() *BlueprintList {
	if in == nil {
		return nil
	}
	out := new(BlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintPhase) DeepCopyInto(out *BlueprintPhase) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintPhase.
func (in *BlueprintPhase) DeepCopy() *BlueprintPhase {
	if in == nil {
		return nil
	}
	out := new(BlueprintPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Phase) DeepCopyInto(out *Phase) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Phase.
func (in *Phase) DeepCopy() *Phase {
	if in == nil {
		return nil
	}
	out := new(Phase)
	in.DeepCopyInto(out)
	return out
}
