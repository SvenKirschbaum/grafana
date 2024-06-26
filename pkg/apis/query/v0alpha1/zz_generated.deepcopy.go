//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by deepcopy-gen. DO NOT EDIT.

package v0alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourceApiServer) DeepCopyInto(out *DataSourceApiServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AliasIDs != nil {
		in, out := &in.AliasIDs, &out.AliasIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourceApiServer.
func (in *DataSourceApiServer) DeepCopy() *DataSourceApiServer {
	if in == nil {
		return nil
	}
	out := new(DataSourceApiServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataSourceApiServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataSourceApiServerList) DeepCopyInto(out *DataSourceApiServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataSourceApiServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataSourceApiServerList.
func (in *DataSourceApiServerList) DeepCopy() *DataSourceApiServerList {
	if in == nil {
		return nil
	}
	out := new(DataSourceApiServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataSourceApiServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryDataRequest) DeepCopyInto(out *QueryDataRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.QueryDataRequest.DeepCopyInto(&out.QueryDataRequest)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryDataRequest.
func (in *QueryDataRequest) DeepCopy() *QueryDataRequest {
	if in == nil {
		return nil
	}
	out := new(QueryDataRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryDataRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryDataResponse) DeepCopyInto(out *QueryDataResponse) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.QueryDataResponse.DeepCopyInto(&out.QueryDataResponse)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryDataResponse.
func (in *QueryDataResponse) DeepCopy() *QueryDataResponse {
	if in == nil {
		return nil
	}
	out := new(QueryDataResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryDataResponse) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryTypeDefinition) DeepCopyInto(out *QueryTypeDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryTypeDefinition.
func (in *QueryTypeDefinition) DeepCopy() *QueryTypeDefinition {
	if in == nil {
		return nil
	}
	out := new(QueryTypeDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryTypeDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryTypeDefinitionList) DeepCopyInto(out *QueryTypeDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QueryTypeDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryTypeDefinitionList.
func (in *QueryTypeDefinitionList) DeepCopy() *QueryTypeDefinitionList {
	if in == nil {
		return nil
	}
	out := new(QueryTypeDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueryTypeDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
