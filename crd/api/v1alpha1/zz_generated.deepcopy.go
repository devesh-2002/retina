//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capture) DeepCopyInto(out *Capture) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capture.
func (in *Capture) DeepCopy() *Capture {
	if in == nil {
		return nil
	}
	out := new(Capture)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Capture) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureConfiguration) DeepCopyInto(out *CaptureConfiguration) {
	*out = *in
	in.CaptureTarget.DeepCopyInto(&out.CaptureTarget)
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = new(CaptureConfigurationFilters)
		(*in).DeepCopyInto(*out)
	}
	if in.TcpdumpFilter != nil {
		in, out := &in.TcpdumpFilter, &out.TcpdumpFilter
		*out = new(string)
		**out = **in
	}
	in.CaptureOption.DeepCopyInto(&out.CaptureOption)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureConfiguration.
func (in *CaptureConfiguration) DeepCopy() *CaptureConfiguration {
	if in == nil {
		return nil
	}
	out := new(CaptureConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureConfigurationFilters) DeepCopyInto(out *CaptureConfigurationFilters) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureConfigurationFilters.
func (in *CaptureConfigurationFilters) DeepCopy() *CaptureConfigurationFilters {
	if in == nil {
		return nil
	}
	out := new(CaptureConfigurationFilters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureList) DeepCopyInto(out *CaptureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Capture, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureList.
func (in *CaptureList) DeepCopy() *CaptureList {
	if in == nil {
		return nil
	}
	out := new(CaptureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CaptureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureOption) DeepCopyInto(out *CaptureOption) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PacketSize != nil {
		in, out := &in.PacketSize, &out.PacketSize
		*out = new(int)
		**out = **in
	}
	if in.MaxCaptureSize != nil {
		in, out := &in.MaxCaptureSize, &out.MaxCaptureSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureOption.
func (in *CaptureOption) DeepCopy() *CaptureOption {
	if in == nil {
		return nil
	}
	out := new(CaptureOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureSpec) DeepCopyInto(out *CaptureSpec) {
	*out = *in
	in.CaptureConfiguration.DeepCopyInto(&out.CaptureConfiguration)
	in.OutputConfiguration.DeepCopyInto(&out.OutputConfiguration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureSpec.
func (in *CaptureSpec) DeepCopy() *CaptureSpec {
	if in == nil {
		return nil
	}
	out := new(CaptureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureStatus) DeepCopyInto(out *CaptureStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureStatus.
func (in *CaptureStatus) DeepCopy() *CaptureStatus {
	if in == nil {
		return nil
	}
	out := new(CaptureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaptureTarget) DeepCopyInto(out *CaptureTarget) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaptureTarget.
func (in *CaptureTarget) DeepCopy() *CaptureTarget {
	if in == nil {
		return nil
	}
	out := new(CaptureTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Containers) DeepCopyInto(out *Containers) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Containers.
func (in *Containers) DeepCopy() *Containers {
	if in == nil {
		return nil
	}
	out := new(Containers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPBlock) DeepCopyInto(out *IPBlock) {
	*out = *in
	if in.Except != nil {
		in, out := &in.Except, &out.Except
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPBlock.
func (in *IPBlock) DeepCopy() *IPBlock {
	if in == nil {
		return nil
	}
	out := new(IPBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfiguration) DeepCopyInto(out *MetricsConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfiguration.
func (in *MetricsConfiguration) DeepCopy() *MetricsConfiguration {
	if in == nil {
		return nil
	}
	out := new(MetricsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsConfigurationList) DeepCopyInto(out *MetricsConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricsConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsConfigurationList.
func (in *MetricsConfigurationList) DeepCopy() *MetricsConfigurationList {
	if in == nil {
		return nil
	}
	out := new(MetricsConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsContextOptions) DeepCopyInto(out *MetricsContextOptions) {
	*out = *in
	if in.SourceLabels != nil {
		in, out := &in.SourceLabels, &out.SourceLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DestinationLabels != nil {
		in, out := &in.DestinationLabels, &out.DestinationLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsContextOptions.
func (in *MetricsContextOptions) DeepCopy() *MetricsContextOptions {
	if in == nil {
		return nil
	}
	out := new(MetricsContextOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsNamespaces) DeepCopyInto(out *MetricsNamespaces) {
	*out = *in
	if in.Include != nil {
		in, out := &in.Include, &out.Include
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Exclude != nil {
		in, out := &in.Exclude, &out.Exclude
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsNamespaces.
func (in *MetricsNamespaces) DeepCopy() *MetricsNamespaces {
	if in == nil {
		return nil
	}
	out := new(MetricsNamespaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsSpec) DeepCopyInto(out *MetricsSpec) {
	*out = *in
	if in.ContextOptions != nil {
		in, out := &in.ContextOptions, &out.ContextOptions
		*out = make([]MetricsContextOptions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Namespaces.DeepCopyInto(&out.Namespaces)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsSpec.
func (in *MetricsSpec) DeepCopy() *MetricsSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsStatus) DeepCopyInto(out *MetricsStatus) {
	*out = *in
	if in.LastKnownSpec != nil {
		in, out := &in.LastKnownSpec, &out.LastKnownSpec
		*out = new(MetricsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsStatus.
func (in *MetricsStatus) DeepCopy() *MetricsStatus {
	if in == nil {
		return nil
	}
	out := new(MetricsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputConfiguration) DeepCopyInto(out *OutputConfiguration) {
	*out = *in
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(string)
		**out = **in
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(string)
		**out = **in
	}
	if in.BlobUpload != nil {
		in, out := &in.BlobUpload, &out.BlobUpload
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputConfiguration.
func (in *OutputConfiguration) DeepCopy() *OutputConfiguration {
	if in == nil {
		return nil
	}
	out := new(OutputConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OwnerReference) DeepCopyInto(out *OwnerReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OwnerReference.
func (in *OwnerReference) DeepCopy() *OwnerReference {
	if in == nil {
		return nil
	}
	out := new(OwnerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetinaEndpoint) DeepCopyInto(out *RetinaEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetinaEndpoint.
func (in *RetinaEndpoint) DeepCopy() *RetinaEndpoint {
	if in == nil {
		return nil
	}
	out := new(RetinaEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RetinaEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetinaEndpointList) DeepCopyInto(out *RetinaEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RetinaEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetinaEndpointList.
func (in *RetinaEndpointList) DeepCopy() *RetinaEndpointList {
	if in == nil {
		return nil
	}
	out := new(RetinaEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RetinaEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetinaEndpointSpec) DeepCopyInto(out *RetinaEndpointSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]RetinaEndpointStatusContainers, len(*in))
		copy(*out, *in)
	}
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]OwnerReference, len(*in))
		copy(*out, *in)
	}
	if in.PodIPs != nil {
		in, out := &in.PodIPs, &out.PodIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetinaEndpointSpec.
func (in *RetinaEndpointSpec) DeepCopy() *RetinaEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(RetinaEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetinaEndpointStatus) DeepCopyInto(out *RetinaEndpointStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetinaEndpointStatus.
func (in *RetinaEndpointStatus) DeepCopy() *RetinaEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(RetinaEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetinaEndpointStatusContainers) DeepCopyInto(out *RetinaEndpointStatusContainers) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetinaEndpointStatusContainers.
func (in *RetinaEndpointStatusContainers) DeepCopy() *RetinaEndpointStatusContainers {
	if in == nil {
		return nil
	}
	out := new(RetinaEndpointStatusContainers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceConfiguration) DeepCopyInto(out *TraceConfiguration) {
	*out = *in
	if in.TraceTargets != nil {
		in, out := &in.TraceTargets, &out.TraceTargets
		*out = make([]*TraceTargets, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TraceTargets)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceConfiguration.
func (in *TraceConfiguration) DeepCopy() *TraceConfiguration {
	if in == nil {
		return nil
	}
	out := new(TraceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceOutputConfiguration) DeepCopyInto(out *TraceOutputConfiguration) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceOutputConfiguration.
func (in *TraceOutputConfiguration) DeepCopy() *TraceOutputConfiguration {
	if in == nil {
		return nil
	}
	out := new(TraceOutputConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TracePoints) DeepCopyInto(out *TracePoints) {
	{
		in := &in
		*out = make(TracePoints, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracePoints.
func (in TracePoints) DeepCopy() TracePoints {
	if in == nil {
		return nil
	}
	out := new(TracePoints)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracePorts) DeepCopyInto(out *TracePorts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracePorts.
func (in *TracePorts) DeepCopy() *TracePorts {
	if in == nil {
		return nil
	}
	out := new(TracePorts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceTarget) DeepCopyInto(out *TraceTarget) {
	*out = *in
	in.IPBlock.DeepCopyInto(&out.IPBlock)
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceSelector != nil {
		in, out := &in.ServiceSelector, &out.ServiceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceTarget.
func (in *TraceTarget) DeepCopy() *TraceTarget {
	if in == nil {
		return nil
	}
	out := new(TraceTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceTargets) DeepCopyInto(out *TraceTargets) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(TraceTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Destination != nil {
		in, out := &in.Destination, &out.Destination
		*out = new(TraceTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*TracePorts, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TracePorts)
				**out = **in
			}
		}
	}
	if in.TracePoints != nil {
		in, out := &in.TracePoints, &out.TracePoints
		*out = make(TracePoints, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceTargets.
func (in *TraceTargets) DeepCopy() *TraceTargets {
	if in == nil {
		return nil
	}
	out := new(TraceTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracesConfiguration) DeepCopyInto(out *TracesConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(TracesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(TracesStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracesConfiguration.
func (in *TracesConfiguration) DeepCopy() *TracesConfiguration {
	if in == nil {
		return nil
	}
	out := new(TracesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracesConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracesConfigurationList) DeepCopyInto(out *TracesConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TracesConfigurationList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracesConfigurationList.
func (in *TracesConfigurationList) DeepCopy() *TracesConfigurationList {
	if in == nil {
		return nil
	}
	out := new(TracesConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TracesConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracesSpec) DeepCopyInto(out *TracesSpec) {
	*out = *in
	if in.TraceConfiguration != nil {
		in, out := &in.TraceConfiguration, &out.TraceConfiguration
		*out = make([]*TraceConfiguration, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TraceConfiguration)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TraceOutputConfiguration != nil {
		in, out := &in.TraceOutputConfiguration, &out.TraceOutputConfiguration
		*out = new(TraceOutputConfiguration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracesSpec.
func (in *TracesSpec) DeepCopy() *TracesSpec {
	if in == nil {
		return nil
	}
	out := new(TracesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TracesStatus) DeepCopyInto(out *TracesStatus) {
	*out = *in
	if in.LastKnownSpec != nil {
		in, out := &in.LastKnownSpec, &out.LastKnownSpec
		*out = new(TracesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TracesStatus.
func (in *TracesStatus) DeepCopy() *TracesStatus {
	if in == nil {
		return nil
	}
	out := new(TracesStatus)
	in.DeepCopyInto(out)
	return out
}
