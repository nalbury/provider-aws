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

type IdentityProviderInitParameters struct {

	// The map of attribute mapping of user pool attributes. AttributeMapping in AWS API documentation
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	// The list of identity providers.
	IdpIdentifiers []*string `json:"idpIdentifiers,omitempty" tf:"idp_identifiers,omitempty"`

	// The map of identity details, such as access token
	ProviderDetails map[string]*string `json:"providerDetails,omitempty" tf:"provider_details,omitempty"`

	// The provider name
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The provider type.  See AWS API for valid values
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`
}

type IdentityProviderObservation struct {

	// The map of attribute mapping of user pool attributes. AttributeMapping in AWS API documentation
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of identity providers.
	IdpIdentifiers []*string `json:"idpIdentifiers,omitempty" tf:"idp_identifiers,omitempty"`

	// The map of identity details, such as access token
	ProviderDetails map[string]*string `json:"providerDetails,omitempty" tf:"provider_details,omitempty"`

	// The provider name
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The provider type.  See AWS API for valid values
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// The user pool id
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`
}

type IdentityProviderParameters struct {

	// The map of attribute mapping of user pool attributes. AttributeMapping in AWS API documentation
	// +kubebuilder:validation:Optional
	AttributeMapping map[string]*string `json:"attributeMapping,omitempty" tf:"attribute_mapping,omitempty"`

	// The list of identity providers.
	// +kubebuilder:validation:Optional
	IdpIdentifiers []*string `json:"idpIdentifiers,omitempty" tf:"idp_identifiers,omitempty"`

	// The map of identity details, such as access token
	// +kubebuilder:validation:Optional
	ProviderDetails map[string]*string `json:"providerDetails,omitempty" tf:"provider_details,omitempty"`

	// The provider name
	// +kubebuilder:validation:Optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name,omitempty"`

	// The provider type.  See AWS API for valid values
	// +kubebuilder:validation:Optional
	ProviderType *string `json:"providerType,omitempty" tf:"provider_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The user pool id
	// +crossplane:generate:reference:type=UserPool
	// +kubebuilder:validation:Optional
	UserPoolID *string `json:"userPoolId,omitempty" tf:"user_pool_id,omitempty"`

	// Reference to a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDRef *v1.Reference `json:"userPoolIdRef,omitempty" tf:"-"`

	// Selector for a UserPool to populate userPoolId.
	// +kubebuilder:validation:Optional
	UserPoolIDSelector *v1.Selector `json:"userPoolIdSelector,omitempty" tf:"-"`
}

// IdentityProviderSpec defines the desired state of IdentityProvider
type IdentityProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderParameters `json:"forProvider"`
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
	InitProvider IdentityProviderInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderStatus defines the observed state of IdentityProvider.
type IdentityProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProvider is the Schema for the IdentityProviders API. Provides a Cognito User Identity Provider resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerDetails) || (has(self.initProvider) && has(self.initProvider.providerDetails))",message="spec.forProvider.providerDetails is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerName) || (has(self.initProvider) && has(self.initProvider.providerName))",message="spec.forProvider.providerName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.providerType) || (has(self.initProvider) && has(self.initProvider.providerType))",message="spec.forProvider.providerType is a required parameter"
	Spec   IdentityProviderSpec   `json:"spec"`
	Status IdentityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderList contains a list of IdentityProviders
type IdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProvider `json:"items"`
}

// Repository type metadata.
var (
	IdentityProvider_Kind             = "IdentityProvider"
	IdentityProvider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProvider_Kind}.String()
	IdentityProvider_KindAPIVersion   = IdentityProvider_Kind + "." + CRDGroupVersion.String()
	IdentityProvider_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProvider_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProvider{}, &IdentityProviderList{})
}
