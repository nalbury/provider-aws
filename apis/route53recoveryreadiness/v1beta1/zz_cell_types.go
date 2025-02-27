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

type CellInitParameters struct {

	// List of cell arns to add as nested fault domains within this cell.
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CellObservation struct {

	// ARN of the cell
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// List of cell arns to add as nested fault domains within this cell.
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of readiness scopes (recovery groups or cells) that contain this cell.
	ParentReadinessScopes []*string `json:"parentReadinessScopes,omitempty" tf:"parent_readiness_scopes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CellParameters struct {

	// List of cell arns to add as nested fault domains within this cell.
	// +kubebuilder:validation:Optional
	Cells []*string `json:"cells,omitempty" tf:"cells,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CellSpec defines the desired state of Cell
type CellSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CellParameters `json:"forProvider"`
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
	InitProvider CellInitParameters `json:"initProvider,omitempty"`
}

// CellStatus defines the observed state of Cell.
type CellStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CellObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cell is the Schema for the Cells API. Provides an AWS Route 53 Recovery Readiness Cell
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cell struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CellSpec   `json:"spec"`
	Status            CellStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CellList contains a list of Cells
type CellList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cell `json:"items"`
}

// Repository type metadata.
var (
	Cell_Kind             = "Cell"
	Cell_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cell_Kind}.String()
	Cell_KindAPIVersion   = Cell_Kind + "." + CRDGroupVersion.String()
	Cell_GroupVersionKind = CRDGroupVersion.WithKind(Cell_Kind)
)

func init() {
	SchemeBuilder.Register(&Cell{}, &CellList{})
}
