package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// PingSpec defines the desired state of Ping.
type PingSpec struct {
	// Message is the payload emitted on each ping cycle.
	// +kubebuilder:validation:MaxLength=140
	Message string `json:"message,omitempty"`

	// IntervalSeconds is how often the operator logs a ping.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=30
	IntervalSeconds int32 `json:"intervalSeconds,omitempty"`
}

// PingStatus defines the observed state of Ping.
type PingStatus struct {
	// LastPinged is the RFC3339 time of the most recent ping.
	LastPinged string `json:"lastPinged,omitempty"`
	// PingCount is the number of pings emitted since creation.
	PingCount int64 `json:"pingCount,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Message",type=string,JSONPath=`.spec.message`
// +kubebuilder:printcolumn:name="LastPinged",type=string,JSONPath=`.status.lastPinged`

// Ping is the Schema for the pings API.
type Ping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PingSpec   `json:"spec,omitempty"`
	Status PingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PingList contains a list of Ping.
type PingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ping{}, &PingList{})
}

// DeepCopyInto copies the receiver into out.
func (in *PingSpec) DeepCopyInto(out *PingSpec) {
	*out = *in
}

// DeepCopy creates a deep copy of PingSpec.
func (in *PingSpec) DeepCopy() *PingSpec {
	if in == nil {
		return nil
	}
	out := new(PingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *PingStatus) DeepCopyInto(out *PingStatus) {
	*out = *in
}

// DeepCopy creates a deep copy of PingStatus.
func (in *PingStatus) DeepCopy() *PingStatus {
	if in == nil {
		return nil
	}
	out := new(PingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto copies the receiver into out.
func (in *Ping) DeepCopyInto(out *Ping) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy creates a deep copy of Ping.
func (in *Ping) DeepCopy() *Ping {
	if in == nil {
		return nil
	}
	out := new(Ping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject implements runtime.Object.
func (in *Ping) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto copies the receiver into out.
func (in *PingList) DeepCopyInto(out *PingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		l := make([]Ping, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&l[i])
		}
		out.Items = l
	}
}

// DeepCopy creates a deep copy of PingList.
func (in *PingList) DeepCopy() *PingList {
	if in == nil {
		return nil
	}
	out := new(PingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject implements runtime.Object.
func (in *PingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
