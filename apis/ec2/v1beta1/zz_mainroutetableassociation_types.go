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

type MainRouteTableAssociationInitParameters struct {
}

type MainRouteTableAssociationObservation struct {

	// The ID of the Route Table Association
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Used internally, see Notes below
	OriginalRouteTableID *string `json:"originalRouteTableId,omitempty" tf:"original_route_table_id,omitempty"`

	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// The ID of the VPC whose main route table should be set
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type MainRouteTableAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ID of the Route Table to set as the new
	// main route table for the target VPC
	// +crossplane:generate:reference:type=RouteTable
	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// Reference to a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDRef *v1.Reference `json:"routeTableIdRef,omitempty" tf:"-"`

	// Selector for a RouteTable to populate routeTableId.
	// +kubebuilder:validation:Optional
	RouteTableIDSelector *v1.Selector `json:"routeTableIdSelector,omitempty" tf:"-"`

	// The ID of the VPC whose main route table should be set
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// MainRouteTableAssociationSpec defines the desired state of MainRouteTableAssociation
type MainRouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MainRouteTableAssociationParameters `json:"forProvider"`
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
	InitProvider MainRouteTableAssociationInitParameters `json:"initProvider,omitempty"`
}

// MainRouteTableAssociationStatus defines the observed state of MainRouteTableAssociation.
type MainRouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MainRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MainRouteTableAssociation is the Schema for the MainRouteTableAssociations API. Provides a resource for managing the main routing table of a VPC.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MainRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MainRouteTableAssociationSpec   `json:"spec"`
	Status            MainRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MainRouteTableAssociationList contains a list of MainRouteTableAssociations
type MainRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MainRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	MainRouteTableAssociation_Kind             = "MainRouteTableAssociation"
	MainRouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MainRouteTableAssociation_Kind}.String()
	MainRouteTableAssociation_KindAPIVersion   = MainRouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	MainRouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(MainRouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&MainRouteTableAssociation{}, &MainRouteTableAssociationList{})
}
