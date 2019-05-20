// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Grafana) DeepCopyInto(out *Grafana) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Grafana.
func (in *Grafana) DeepCopy() *Grafana {
	if in == nil {
		return nil
	}
	out := new(Grafana)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Grafana) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboard) DeepCopyInto(out *GrafanaDashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboard.
func (in *GrafanaDashboard) DeepCopy() *GrafanaDashboard {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardList) DeepCopyInto(out *GrafanaDashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardList.
func (in *GrafanaDashboardList) DeepCopy() *GrafanaDashboardList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardSpec) DeepCopyInto(out *GrafanaDashboardSpec) {
	*out = *in
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardSpec.
func (in *GrafanaDashboardSpec) DeepCopy() *GrafanaDashboardSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardStatus) DeepCopyInto(out *GrafanaDashboardStatus) {
	*out = *in
	if in.Messages != nil {
		in, out := &in.Messages, &out.Messages
		*out = make([]GrafanaDashboardStatusMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardStatus.
func (in *GrafanaDashboardStatus) DeepCopy() *GrafanaDashboardStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDashboardStatusMessage) DeepCopyInto(out *GrafanaDashboardStatusMessage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDashboardStatusMessage.
func (in *GrafanaDashboardStatusMessage) DeepCopy() *GrafanaDashboardStatusMessage {
	if in == nil {
		return nil
	}
	out := new(GrafanaDashboardStatusMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSource) DeepCopyInto(out *GrafanaDataSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSource.
func (in *GrafanaDataSource) DeepCopy() *GrafanaDataSource {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDataSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceFields) DeepCopyInto(out *GrafanaDataSourceFields) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceFields.
func (in *GrafanaDataSourceFields) DeepCopy() *GrafanaDataSourceFields {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceList) DeepCopyInto(out *GrafanaDataSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GrafanaDataSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceList.
func (in *GrafanaDataSourceList) DeepCopy() *GrafanaDataSourceList {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaDataSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceSpec) DeepCopyInto(out *GrafanaDataSourceSpec) {
	*out = *in
	if in.Datasources != nil {
		in, out := &in.Datasources, &out.Datasources
		*out = make([]GrafanaDataSourceFields, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceSpec.
func (in *GrafanaDataSourceSpec) DeepCopy() *GrafanaDataSourceSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaDataSourceStatus) DeepCopyInto(out *GrafanaDataSourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaDataSourceStatus.
func (in *GrafanaDataSourceStatus) DeepCopy() *GrafanaDataSourceStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaDataSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaList) DeepCopyInto(out *GrafanaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Grafana, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaList.
func (in *GrafanaList) DeepCopy() *GrafanaList {
	if in == nil {
		return nil
	}
	out := new(GrafanaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GrafanaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaPlugin) DeepCopyInto(out *GrafanaPlugin) {
	*out = *in
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(GrafanaDashboard)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaPlugin.
func (in *GrafanaPlugin) DeepCopy() *GrafanaPlugin {
	if in == nil {
		return nil
	}
	out := new(GrafanaPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaSpec) DeepCopyInto(out *GrafanaSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DashboardLabelSelectors != nil {
		in, out := &in.DashboardLabelSelectors, &out.DashboardLabelSelectors
		*out = make([]*metav1.LabelSelector, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(metav1.LabelSelector)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaSpec.
func (in *GrafanaSpec) DeepCopy() *GrafanaSpec {
	if in == nil {
		return nil
	}
	out := new(GrafanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrafanaStatus) DeepCopyInto(out *GrafanaStatus) {
	*out = *in
	if in.InstalledPlugins != nil {
		in, out := &in.InstalledPlugins, &out.InstalledPlugins
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrafanaStatus.
func (in *GrafanaStatus) DeepCopy() *GrafanaStatus {
	if in == nil {
		return nil
	}
	out := new(GrafanaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PluginList) DeepCopyInto(out *PluginList) {
	{
		in := &in
		*out = make(PluginList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginList.
func (in PluginList) DeepCopy() PluginList {
	if in == nil {
		return nil
	}
	out := new(PluginList)
	in.DeepCopyInto(out)
	return *out
}
