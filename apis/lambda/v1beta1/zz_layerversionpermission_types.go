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

type LayerVersionPermissionInitParameters struct {

	// Action, which will be allowed. lambda:GetLayerVersion value is suggested by AWS documantation.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name or ARN of the Lambda Layer, which you want to grant access to.
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// An identifier of AWS Organization, which should be able to use your Lambda Layer. principal should be equal to * if organization_id provided.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// AWS account ID which should be able to use your Lambda Layer. * can be used here, if you want to share your Lambda Layer widely.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// The name of Lambda Layer Permission, for example dev-account - human readable note about what is this permission for.
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// Version of Lambda Layer, which you want to grant access to. Note: permissions only apply to a single version of a layer.
	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

type LayerVersionPermissionObservation struct {

	// Action, which will be allowed. lambda:GetLayerVersion value is suggested by AWS documantation.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The layer_name and version_number, separated by a comma (,).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or ARN of the Lambda Layer, which you want to grant access to.
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// An identifier of AWS Organization, which should be able to use your Lambda Layer. principal should be equal to * if organization_id provided.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Full Lambda Layer Permission policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// AWS account ID which should be able to use your Lambda Layer. * can be used here, if you want to share your Lambda Layer widely.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// A unique identifier for the current revision of the policy.
	RevisionID *string `json:"revisionId,omitempty" tf:"revision_id,omitempty"`

	// The name of Lambda Layer Permission, for example dev-account - human readable note about what is this permission for.
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// Version of Lambda Layer, which you want to grant access to. Note: permissions only apply to a single version of a layer.
	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

type LayerVersionPermissionParameters struct {

	// Action, which will be allowed. lambda:GetLayerVersion value is suggested by AWS documantation.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// The name or ARN of the Lambda Layer, which you want to grant access to.
	// +kubebuilder:validation:Optional
	LayerName *string `json:"layerName,omitempty" tf:"layer_name,omitempty"`

	// An identifier of AWS Organization, which should be able to use your Lambda Layer. principal should be equal to * if organization_id provided.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// AWS account ID which should be able to use your Lambda Layer. * can be used here, if you want to share your Lambda Layer widely.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of Lambda Layer Permission, for example dev-account - human readable note about what is this permission for.
	// +kubebuilder:validation:Optional
	StatementID *string `json:"statementId,omitempty" tf:"statement_id,omitempty"`

	// Version of Lambda Layer, which you want to grant access to. Note: permissions only apply to a single version of a layer.
	// +kubebuilder:validation:Optional
	VersionNumber *float64 `json:"versionNumber,omitempty" tf:"version_number,omitempty"`
}

// LayerVersionPermissionSpec defines the desired state of LayerVersionPermission
type LayerVersionPermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LayerVersionPermissionParameters `json:"forProvider"`
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
	InitProvider LayerVersionPermissionInitParameters `json:"initProvider,omitempty"`
}

// LayerVersionPermissionStatus defines the observed state of LayerVersionPermission.
type LayerVersionPermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LayerVersionPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersionPermission is the Schema for the LayerVersionPermissions API. Provides a Lambda Layer Version Permission resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LayerVersionPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.action) || (has(self.initProvider) && has(self.initProvider.action))",message="spec.forProvider.action is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.layerName) || (has(self.initProvider) && has(self.initProvider.layerName))",message="spec.forProvider.layerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || (has(self.initProvider) && has(self.initProvider.principal))",message="spec.forProvider.principal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.statementId) || (has(self.initProvider) && has(self.initProvider.statementId))",message="spec.forProvider.statementId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.versionNumber) || (has(self.initProvider) && has(self.initProvider.versionNumber))",message="spec.forProvider.versionNumber is a required parameter"
	Spec   LayerVersionPermissionSpec   `json:"spec"`
	Status LayerVersionPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LayerVersionPermissionList contains a list of LayerVersionPermissions
type LayerVersionPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LayerVersionPermission `json:"items"`
}

// Repository type metadata.
var (
	LayerVersionPermission_Kind             = "LayerVersionPermission"
	LayerVersionPermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LayerVersionPermission_Kind}.String()
	LayerVersionPermission_KindAPIVersion   = LayerVersionPermission_Kind + "." + CRDGroupVersion.String()
	LayerVersionPermission_GroupVersionKind = CRDGroupVersion.WithKind(LayerVersionPermission_Kind)
)

func init() {
	SchemeBuilder.Register(&LayerVersionPermission{}, &LayerVersionPermissionList{})
}
