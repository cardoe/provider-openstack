// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
Copyright 2023 Jakob Schlagenhaufer, Jan Dittrich
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type QuotasetV3InitParameters struct {

	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes *float64 `json:"backupGigabytes,omitempty" tf:"backup_gigabytes,omitempty"`

	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups *float64 `json:"backups,omitempty" tf:"backups,omitempty"`

	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes *float64 `json:"gigabytes,omitempty" tf:"gigabytes,omitempty"`

	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups *float64 `json:"groups,omitempty" tf:"groups,omitempty"`

	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes *float64 `json:"perVolumeGigabytes,omitempty" tf:"per_volume_gigabytes,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots *float64 `json:"snapshots,omitempty" tf:"snapshots,omitempty"`

	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are snapshots_<volume_type_name>,
	// volumes_<volume_type_name> and gigabytes_<volume_type_name>.
	VolumeTypeQuota map[string]*string `json:"volumeTypeQuota,omitempty" tf:"volume_type_quota,omitempty"`

	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes *float64 `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type QuotasetV3Observation struct {

	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	BackupGigabytes *float64 `json:"backupGigabytes,omitempty" tf:"backup_gigabytes,omitempty"`

	// Quota value for backups. Changing this updates the
	// existing quotaset.
	Backups *float64 `json:"backups,omitempty" tf:"backups,omitempty"`

	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	Gigabytes *float64 `json:"gigabytes,omitempty" tf:"gigabytes,omitempty"`

	// Quota value for groups. Changing this updates the
	// existing quotaset.
	Groups *float64 `json:"groups,omitempty" tf:"groups,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	PerVolumeGigabytes *float64 `json:"perVolumeGigabytes,omitempty" tf:"per_volume_gigabytes,omitempty"`

	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	Snapshots *float64 `json:"snapshots,omitempty" tf:"snapshots,omitempty"`

	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are snapshots_<volume_type_name>,
	// volumes_<volume_type_name> and gigabytes_<volume_type_name>.
	VolumeTypeQuota map[string]*string `json:"volumeTypeQuota,omitempty" tf:"volume_type_quota,omitempty"`

	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	Volumes *float64 `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type QuotasetV3Parameters struct {

	// Quota value for backup gigabytes. Changing
	// this updates the existing quotaset.
	// +kubebuilder:validation:Optional
	BackupGigabytes *float64 `json:"backupGigabytes,omitempty" tf:"backup_gigabytes,omitempty"`

	// Quota value for backups. Changing this updates the
	// existing quotaset.
	// +kubebuilder:validation:Optional
	Backups *float64 `json:"backups,omitempty" tf:"backups,omitempty"`

	// Quota value for gigabytes. Changing this updates the
	// existing quotaset.
	// +kubebuilder:validation:Optional
	Gigabytes *float64 `json:"gigabytes,omitempty" tf:"gigabytes,omitempty"`

	// Quota value for groups. Changing this updates the
	// existing quotaset.
	// +kubebuilder:validation:Optional
	Groups *float64 `json:"groups,omitempty" tf:"groups,omitempty"`

	// Quota value for gigabytes per volume .
	// Changing this updates the existing quotaset.
	// +kubebuilder:validation:Optional
	PerVolumeGigabytes *float64 `json:"perVolumeGigabytes,omitempty" tf:"per_volume_gigabytes,omitempty"`

	// ID of the project to manage quotas. Changing this
	// creates a new quotaset.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-openstack/apis/identity/v1alpha1.ProjectV3
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a ProjectV3 in identity to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// The region in which to create the volume. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new quotaset.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Quota value for snapshots. Changing this updates the
	// existing quotaset.
	// +kubebuilder:validation:Optional
	Snapshots *float64 `json:"snapshots,omitempty" tf:"snapshots,omitempty"`

	// Key/Value pairs for setting quota for
	// volumes types. Possible keys are snapshots_<volume_type_name>,
	// volumes_<volume_type_name> and gigabytes_<volume_type_name>.
	// +kubebuilder:validation:Optional
	VolumeTypeQuota map[string]*string `json:"volumeTypeQuota,omitempty" tf:"volume_type_quota,omitempty"`

	// Quota value for volumes. Changing this updates the
	// existing quotaset.
	// +kubebuilder:validation:Optional
	Volumes *float64 `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

// QuotasetV3Spec defines the desired state of QuotasetV3
type QuotasetV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QuotasetV3Parameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider QuotasetV3InitParameters `json:"initProvider,omitempty"`
}

// QuotasetV3Status defines the observed state of QuotasetV3.
type QuotasetV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QuotasetV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QuotasetV3 is the Schema for the QuotasetV3s API. Manages a V3 quotaset resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type QuotasetV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QuotasetV3Spec   `json:"spec"`
	Status            QuotasetV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QuotasetV3List contains a list of QuotasetV3s
type QuotasetV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QuotasetV3 `json:"items"`
}

// Repository type metadata.
var (
	QuotasetV3_Kind             = "QuotasetV3"
	QuotasetV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QuotasetV3_Kind}.String()
	QuotasetV3_KindAPIVersion   = QuotasetV3_Kind + "." + CRDGroupVersion.String()
	QuotasetV3_GroupVersionKind = CRDGroupVersion.WithKind(QuotasetV3_Kind)
)

func init() {
	SchemeBuilder.Register(&QuotasetV3{}, &QuotasetV3List{})
}
