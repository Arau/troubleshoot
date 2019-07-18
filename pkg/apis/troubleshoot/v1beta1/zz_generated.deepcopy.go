// +build !ignore_autogenerated

/*
Copyright 2019 Replicated, Inc..

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

// autogenerated by controller-gen object, do not modify manually

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analyze) DeepCopyInto(out *Analyze) {
	*out = *in
	if in.ClusterVersion != nil {
		in, out := &in.ClusterVersion, &out.ClusterVersion
		*out = new(ClusterVersion)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(StorageClass)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomResourceDefinition != nil {
		in, out := &in.CustomResourceDefinition, &out.CustomResourceDefinition
		*out = new(CustomResourceDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(Ingress)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(AnalyzeSecret)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analyze.
func (in *Analyze) DeepCopy() *Analyze {
	if in == nil {
		return nil
	}
	out := new(Analyze)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzeMeta) DeepCopyInto(out *AnalyzeMeta) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzeMeta.
func (in *AnalyzeMeta) DeepCopy() *AnalyzeMeta {
	if in == nil {
		return nil
	}
	out := new(AnalyzeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzeSecret) DeepCopyInto(out *AnalyzeSecret) {
	*out = *in
	out.AnalyzeMeta = in.AnalyzeMeta
	if in.Outcomes != nil {
		in, out := &in.Outcomes, &out.Outcomes
		*out = make([]*Outcome, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Outcome)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzeSecret.
func (in *AnalyzeSecret) DeepCopy() *AnalyzeSecret {
	if in == nil {
		return nil
	}
	out := new(AnalyzeSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Analyzer) DeepCopyInto(out *Analyzer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Analyzer.
func (in *Analyzer) DeepCopy() *Analyzer {
	if in == nil {
		return nil
	}
	out := new(Analyzer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Analyzer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJob) DeepCopyInto(out *AnalyzerJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJob.
func (in *AnalyzerJob) DeepCopy() *AnalyzerJob {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobList) DeepCopyInto(out *AnalyzerJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalyzerJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobList.
func (in *AnalyzerJobList) DeepCopy() *AnalyzerJobList {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobSpec) DeepCopyInto(out *AnalyzerJobSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobSpec.
func (in *AnalyzerJobSpec) DeepCopy() *AnalyzerJobSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerJobStatus) DeepCopyInto(out *AnalyzerJobStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerJobStatus.
func (in *AnalyzerJobStatus) DeepCopy() *AnalyzerJobStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyzerJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerList) DeepCopyInto(out *AnalyzerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Analyzer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerList.
func (in *AnalyzerList) DeepCopy() *AnalyzerList {
	if in == nil {
		return nil
	}
	out := new(AnalyzerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyzerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerSpec) DeepCopyInto(out *AnalyzerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerSpec.
func (in *AnalyzerSpec) DeepCopy() *AnalyzerSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyzerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyzerStatus) DeepCopyInto(out *AnalyzerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyzerStatus.
func (in *AnalyzerStatus) DeepCopy() *AnalyzerStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyzerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterResources) DeepCopyInto(out *ClusterResources) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterResources.
func (in *ClusterResources) DeepCopy() *ClusterResources {
	if in == nil {
		return nil
	}
	out := new(ClusterResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVersion) DeepCopyInto(out *ClusterVersion) {
	*out = *in
	out.AnalyzeMeta = in.AnalyzeMeta
	if in.Outcomes != nil {
		in, out := &in.Outcomes, &out.Outcomes
		*out = make([]*Outcome, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Outcome)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVersion.
func (in *ClusterVersion) DeepCopy() *ClusterVersion {
	if in == nil {
		return nil
	}
	out := new(ClusterVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collect) DeepCopyInto(out *Collect) {
	*out = *in
	if in.ClusterInfo != nil {
		in, out := &in.ClusterInfo, &out.ClusterInfo
		*out = new(ClusterInfo)
		**out = **in
	}
	if in.ClusterResources != nil {
		in, out := &in.ClusterResources, &out.ClusterResources
		*out = new(ClusterResources)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(Secret)
		**out = **in
	}
	if in.Logs != nil {
		in, out := &in.Logs, &out.Logs
		*out = new(Logs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collect.
func (in *Collect) DeepCopy() *Collect {
	if in == nil {
		return nil
	}
	out := new(Collect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Collector) DeepCopyInto(out *Collector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]*Collect, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Collect)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Collector.
func (in *Collector) DeepCopy() *Collector {
	if in == nil {
		return nil
	}
	out := new(Collector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Collector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJob) DeepCopyInto(out *CollectorJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJob.
func (in *CollectorJob) DeepCopy() *CollectorJob {
	if in == nil {
		return nil
	}
	out := new(CollectorJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobList) DeepCopyInto(out *CollectorJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CollectorJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobList.
func (in *CollectorJobList) DeepCopy() *CollectorJobList {
	if in == nil {
		return nil
	}
	out := new(CollectorJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobSpec) DeepCopyInto(out *CollectorJobSpec) {
	*out = *in
	out.Collector = in.Collector
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobSpec.
func (in *CollectorJobSpec) DeepCopy() *CollectorJobSpec {
	if in == nil {
		return nil
	}
	out := new(CollectorJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorJobStatus) DeepCopyInto(out *CollectorJobStatus) {
	*out = *in
	if in.Running != nil {
		in, out := &in.Running, &out.Running
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorJobStatus.
func (in *CollectorJobStatus) DeepCopy() *CollectorJobStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorList) DeepCopyInto(out *CollectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Collector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorList.
func (in *CollectorList) DeepCopy() *CollectorList {
	if in == nil {
		return nil
	}
	out := new(CollectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CollectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorRef) DeepCopyInto(out *CollectorRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorRef.
func (in *CollectorRef) DeepCopy() *CollectorRef {
	if in == nil {
		return nil
	}
	out := new(CollectorRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectorStatus) DeepCopyInto(out *CollectorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectorStatus.
func (in *CollectorStatus) DeepCopy() *CollectorStatus {
	if in == nil {
		return nil
	}
	out := new(CollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinition) DeepCopyInto(out *CustomResourceDefinition) {
	*out = *in
	out.AnalyzeMeta = in.AnalyzeMeta
	if in.Outcomes != nil {
		in, out := &in.Outcomes, &out.Outcomes
		*out = make([]*Outcome, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Outcome)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinition.
func (in *CustomResourceDefinition) DeepCopy() *CustomResourceDefinition {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	out.AnalyzeMeta = in.AnalyzeMeta
	if in.Outcomes != nil {
		in, out := &in.Outcomes, &out.Outcomes
		*out = make([]*Outcome, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Outcome)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogLimits) DeepCopyInto(out *LogLimits) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogLimits.
func (in *LogLimits) DeepCopy() *LogLimits {
	if in == nil {
		return nil
	}
	out := new(LogLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logs) DeepCopyInto(out *Logs) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(LogLimits)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logs.
func (in *Logs) DeepCopy() *Logs {
	if in == nil {
		return nil
	}
	out := new(Logs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Outcome) DeepCopyInto(out *Outcome) {
	*out = *in
	if in.Fail != nil {
		in, out := &in.Fail, &out.Fail
		*out = new(SingleOutcome)
		**out = **in
	}
	if in.Warn != nil {
		in, out := &in.Warn, &out.Warn
		*out = new(SingleOutcome)
		**out = **in
	}
	if in.Pass != nil {
		in, out := &in.Pass, &out.Pass
		*out = new(SingleOutcome)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Outcome.
func (in *Outcome) DeepCopy() *Outcome {
	if in == nil {
		return nil
	}
	out := new(Outcome)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preflight) DeepCopyInto(out *Preflight) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preflight.
func (in *Preflight) DeepCopy() *Preflight {
	if in == nil {
		return nil
	}
	out := new(Preflight)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Preflight) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJob) DeepCopyInto(out *PreflightJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJob.
func (in *PreflightJob) DeepCopy() *PreflightJob {
	if in == nil {
		return nil
	}
	out := new(PreflightJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobList) DeepCopyInto(out *PreflightJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreflightJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobList.
func (in *PreflightJobList) DeepCopy() *PreflightJobList {
	if in == nil {
		return nil
	}
	out := new(PreflightJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobSpec) DeepCopyInto(out *PreflightJobSpec) {
	*out = *in
	out.Preflight = in.Preflight
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobSpec.
func (in *PreflightJobSpec) DeepCopy() *PreflightJobSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightJobStatus) DeepCopyInto(out *PreflightJobStatus) {
	*out = *in
	if in.CollectorsRunning != nil {
		in, out := &in.CollectorsRunning, &out.CollectorsRunning
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CollectorsSuccessful != nil {
		in, out := &in.CollectorsSuccessful, &out.CollectorsSuccessful
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CollectorsFailed != nil {
		in, out := &in.CollectorsFailed, &out.CollectorsFailed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AnalyzersRunning != nil {
		in, out := &in.AnalyzersRunning, &out.AnalyzersRunning
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AnalyzersSuccessful != nil {
		in, out := &in.AnalyzersSuccessful, &out.AnalyzersSuccessful
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AnalyzersFailed != nil {
		in, out := &in.AnalyzersFailed, &out.AnalyzersFailed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightJobStatus.
func (in *PreflightJobStatus) DeepCopy() *PreflightJobStatus {
	if in == nil {
		return nil
	}
	out := new(PreflightJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightList) DeepCopyInto(out *PreflightList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preflight, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightList.
func (in *PreflightList) DeepCopy() *PreflightList {
	if in == nil {
		return nil
	}
	out := new(PreflightList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightRef) DeepCopyInto(out *PreflightRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightRef.
func (in *PreflightRef) DeepCopy() *PreflightRef {
	if in == nil {
		return nil
	}
	out := new(PreflightRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightSpec) DeepCopyInto(out *PreflightSpec) {
	*out = *in
	if in.Collectors != nil {
		in, out := &in.Collectors, &out.Collectors
		*out = make([]*Collect, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Collect)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Analyzers != nil {
		in, out := &in.Analyzers, &out.Analyzers
		*out = make([]*Analyze, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Analyze)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightSpec.
func (in *PreflightSpec) DeepCopy() *PreflightSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightStatus) DeepCopyInto(out *PreflightStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightStatus.
func (in *PreflightStatus) DeepCopy() *PreflightStatus {
	if in == nil {
		return nil
	}
	out := new(PreflightStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SingleOutcome) DeepCopyInto(out *SingleOutcome) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SingleOutcome.
func (in *SingleOutcome) DeepCopy() *SingleOutcome {
	if in == nil {
		return nil
	}
	out := new(SingleOutcome)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
	out.AnalyzeMeta = in.AnalyzeMeta
	if in.Outcomes != nil {
		in, out := &in.Outcomes, &out.Outcomes
		*out = make([]*Outcome, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Outcome)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}
