// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DiskInitParameters struct {

	// The Availability Zone in which to create your disk.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The instance port the load balancer will connect.
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DiskObservation struct {

	// The ARN of the Lightsail load balancer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Availability Zone in which to create your disk.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The timestamp when the load balancer was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The name of the disk  (matches name).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instance port the load balancer will connect.
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// The support code for the disk. Include this code in your email to support when you have questions about a disk in Lightsail. This code enables our support team to look up your Lightsail information more easily.
	SupportCode *string `json:"supportCode,omitempty" tf:"support_code,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DiskParameters struct {

	// The Availability Zone in which to create your disk.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The instance port the load balancer will connect.
	// +kubebuilder:validation:Optional
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DiskSpec defines the desired state of Disk
type DiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskParameters `json:"forProvider"`
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
	InitProvider DiskInitParameters `json:"initProvider,omitempty"`
}

// DiskStatus defines the observed state of Disk.
type DiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Disk is the Schema for the Disks API. Provides a Lightsail Disk resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Disk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sizeInGb) || (has(self.initProvider) && has(self.initProvider.sizeInGb))",message="spec.forProvider.sizeInGb is a required parameter"
	Spec   DiskSpec   `json:"spec"`
	Status DiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskList contains a list of Disks
type DiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Disk `json:"items"`
}

// Repository type metadata.
var (
	Disk_Kind             = "Disk"
	Disk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Disk_Kind}.String()
	Disk_KindAPIVersion   = Disk_Kind + "." + CRDGroupVersion.String()
	Disk_GroupVersionKind = CRDGroupVersion.WithKind(Disk_Kind)
)

func init() {
	SchemeBuilder.Register(&Disk{}, &DiskList{})
}
