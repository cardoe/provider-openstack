/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ZoneV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ZoneV2Parameters struct {

	// Attributes for the DNS Service scheduler.
	// Changing this creates a new zone.
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// A description of the zone.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Disable wait for zone to reach ACTIVE
	// status. The check is enabled by default. If this argument is true, zone
	// will be considered as created/updated if OpenStack request returned success.
	// +kubebuilder:validation:Optional
	DisableStatusCheck *bool `json:"disableStatusCheck,omitempty" tf:"disable_status_check,omitempty"`

	// The email contact for the zone record.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// An array of master DNS servers. For when type is
	// SECONDARY.
	// +kubebuilder:validation:Optional
	Masters []*string `json:"masters,omitempty" tf:"masters,omitempty"`

	// The name of the zone. Note the . at the end of the name.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project DNS zone is created
	// for, sets X-Auth-Sudo-Tenant-ID header (requires an assigned
	// user role in target project)
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new DNS zone.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The time to live (TTL) of the zone.
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`

	// The type of zone. Can either be PRIMARY or SECONDARY.
	// Changing this creates a new zone.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Map of additional options. Changing this creates a
	// new zone.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ZoneV2Spec defines the desired state of ZoneV2
type ZoneV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneV2Parameters `json:"forProvider"`
}

// ZoneV2Status defines the observed state of ZoneV2.
type ZoneV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2 is the Schema for the ZoneV2s API. Manages a DNS zone in the OpenStack DNS Service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ZoneV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneV2Spec   `json:"spec"`
	Status            ZoneV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneV2List contains a list of ZoneV2s
type ZoneV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneV2 `json:"items"`
}

// Repository type metadata.
var (
	ZoneV2_Kind             = "ZoneV2"
	ZoneV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ZoneV2_Kind}.String()
	ZoneV2_KindAPIVersion   = ZoneV2_Kind + "." + CRDGroupVersion.String()
	ZoneV2_GroupVersionKind = CRDGroupVersion.WithKind(ZoneV2_Kind)
)

func init() {
	SchemeBuilder.Register(&ZoneV2{}, &ZoneV2List{})
}
