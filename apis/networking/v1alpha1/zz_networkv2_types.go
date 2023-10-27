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

type NetworkV2Observation struct {

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The collection of tags assigned on the network, which have been
	// explicitly and implicitly added.
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	AvailabilityZoneHints []*string `json:"availabilityZoneHints,omitempty" tf:"availability_zone_hints,omitempty"`

	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The dns_domain of a network in conjunction with the
	// dns_name attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	External *bool `json:"external,omitempty" tf:"external,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The network MTU. Available for read-only, when Neutron
	// net-mtu extension is enabled. Available for the modification, when
	// Neutron net-mtu-writable extension is enabled.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the network. Changing this updates the name of
	// the existing network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to false will disable port security. Valid values are true and
	// false.
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// Reference to the associated QoS policy.
	QosPolicyID *string `json:"qosPolicyId,omitempty" tf:"qos_policy_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// network.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of one or more provider segment objects.
	// Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
	// updating any provider related segments attributes. Check your plug-in whether
	// it supports updating.
	Segments []SegmentsObservation `json:"segments,omitempty" tf:"segments,omitempty"`

	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A set of string tags for the network.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the transparent_vlan attribute of the existing
	// network.
	TransparentVlan *bool `json:"transparentVlan,omitempty" tf:"transparent_vlan,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type NetworkV2Parameters struct {

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value updates the
	// state of the existing network.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new network.
	// +kubebuilder:validation:Optional
	AvailabilityZoneHints []*string `json:"availabilityZoneHints,omitempty" tf:"availability_zone_hints,omitempty"`

	// The network DNS domain. Available, when Neutron DNS
	// extension is enabled. The dns_domain of a network in conjunction with the
	// dns_name attribute of its ports will be published in an external DNS
	// service when Neutron is configured to integrate with such a service.
	// +kubebuilder:validation:Optional
	DNSDomain *string `json:"dnsDomain,omitempty" tf:"dns_domain,omitempty"`

	// Human-readable description of the network. Changing this
	// updates the name of the existing network.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the network resource has the
	// external routing facility. Valid values are true and false. Defaults to
	// false. Changing this updates the external attribute of the existing network.
	// +kubebuilder:validation:Optional
	External *bool `json:"external,omitempty" tf:"external,omitempty"`

	// The network MTU. Available for read-only, when Neutron
	// net-mtu extension is enabled. Available for the modification, when
	// Neutron net-mtu-writable extension is enabled.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the network. Changing this updates the name of
	// the existing network.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Whether to explicitly enable or disable
	// port security on the network. Port Security is usually enabled by default, so
	// omitting this argument will usually result in a value of "true". Setting this
	// explicitly to false will disable port security. Valid values are true and
	// false.
	// +kubebuilder:validation:Optional
	PortSecurityEnabled *bool `json:"portSecurityEnabled,omitempty" tf:"port_security_enabled,omitempty"`

	// Reference to the associated QoS policy.
	// +kubebuilder:validation:Optional
	QosPolicyID *string `json:"qosPolicyId,omitempty" tf:"qos_policy_id,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron network. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// network.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// An array of one or more provider segment objects.
	// Note: most Networking plug-ins (e.g. ML2 Plugin) and drivers do not support
	// updating any provider related segments attributes. Check your plug-in whether
	// it supports updating.
	// +kubebuilder:validation:Optional
	Segments []SegmentsParameters `json:"segments,omitempty" tf:"segments,omitempty"`

	// Specifies whether the network resource can be accessed
	// by any tenant or not. Changing this updates the sharing capabilities of the
	// existing network.
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// A set of string tags for the network.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the network. Required if admin wants to
	// create a network for another tenant. Changing this creates a new network.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies whether the network resource has the
	// VLAN transparent attribute set. Valid values are true and false. Defaults to
	// false. Changing this updates the transparent_vlan attribute of the existing
	// network.
	// +kubebuilder:validation:Optional
	TransparentVlan *bool `json:"transparentVlan,omitempty" tf:"transparent_vlan,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type SegmentsObservation struct {

	// The type of physical network.
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The physical network where this network is implemented.
	PhysicalNetwork *string `json:"physicalNetwork,omitempty" tf:"physical_network,omitempty"`

	// An isolated segment on the physical network.
	SegmentationID *float64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`
}

type SegmentsParameters struct {

	// The type of physical network.
	// +kubebuilder:validation:Optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type,omitempty"`

	// The physical network where this network is implemented.
	// +kubebuilder:validation:Optional
	PhysicalNetwork *string `json:"physicalNetwork,omitempty" tf:"physical_network,omitempty"`

	// An isolated segment on the physical network.
	// +kubebuilder:validation:Optional
	SegmentationID *float64 `json:"segmentationId,omitempty" tf:"segmentation_id,omitempty"`
}

// NetworkV2Spec defines the desired state of NetworkV2
type NetworkV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkV2Parameters `json:"forProvider"`
}

// NetworkV2Status defines the observed state of NetworkV2.
type NetworkV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkV2 is the Schema for the NetworkV2s API. Manages a V2 Neutron network resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type NetworkV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkV2Spec   `json:"spec"`
	Status            NetworkV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkV2List contains a list of NetworkV2s
type NetworkV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkV2 `json:"items"`
}

// Repository type metadata.
var (
	NetworkV2_Kind             = "NetworkV2"
	NetworkV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkV2_Kind}.String()
	NetworkV2_KindAPIVersion   = NetworkV2_Kind + "." + CRDGroupVersion.String()
	NetworkV2_GroupVersionKind = CRDGroupVersion.WithKind(NetworkV2_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkV2{}, &NetworkV2List{})
}
