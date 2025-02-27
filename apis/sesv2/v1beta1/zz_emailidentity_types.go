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

type DKIMSigningAttributesInitParameters struct {

	// [Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
	DomainSigningPrivateKey *string `json:"domainSigningPrivateKey,omitempty" tf:"domain_signing_private_key,omitempty"`

	// [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
	DomainSigningSelector *string `json:"domainSigningSelector,omitempty" tf:"domain_signing_selector,omitempty"`

	// [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day. Valid values: RSA_1024_BIT, RSA_2048_BIT.
	NextSigningKeyLength *string `json:"nextSigningKeyLength,omitempty" tf:"next_signing_key_length,omitempty"`
}

type DKIMSigningAttributesObservation struct {

	// [Easy DKIM] The key length of the DKIM key pair in use.
	CurrentSigningKeyLength *string `json:"currentSigningKeyLength,omitempty" tf:"current_signing_key_length,omitempty"`

	// [Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
	DomainSigningPrivateKey *string `json:"domainSigningPrivateKey,omitempty" tf:"domain_signing_private_key,omitempty"`

	// [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
	DomainSigningSelector *string `json:"domainSigningSelector,omitempty" tf:"domain_signing_selector,omitempty"`

	// [Easy DKIM] The last time a key pair was generated for this identity.
	LastKeyGenerationTimestamp *string `json:"lastKeyGenerationTimestamp,omitempty" tf:"last_key_generation_timestamp,omitempty"`

	// [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day. Valid values: RSA_1024_BIT, RSA_2048_BIT.
	NextSigningKeyLength *string `json:"nextSigningKeyLength,omitempty" tf:"next_signing_key_length,omitempty"`

	// A string that indicates how DKIM was configured for the identity. AWS_SES indicates that DKIM was configured for the identity by using Easy DKIM. EXTERNAL indicates that DKIM was configured for the identity by using Bring Your Own DKIM (BYODKIM).
	SigningAttributesOrigin *string `json:"signingAttributesOrigin,omitempty" tf:"signing_attributes_origin,omitempty"`

	// Describes whether or not Amazon SES has successfully located the DKIM records in the DNS records for the domain. See the AWS SES API v2 Reference for supported statuses.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// If you used Easy DKIM to configure DKIM authentication for the domain, then this object contains a set of unique strings that you use to create a set of CNAME records that you add to the DNS configuration for your domain. When Amazon SES detects these records in the DNS configuration for your domain, the DKIM authentication process is complete. If you configured DKIM authentication for the domain by providing your own public-private key pair, then this object contains the selector for the public key.
	Tokens []*string `json:"tokens,omitempty" tf:"tokens,omitempty"`
}

type DKIMSigningAttributesParameters struct {

	// [Bring Your Own DKIM] A private key that's used to generate a DKIM signature. The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
	// +kubebuilder:validation:Optional
	DomainSigningPrivateKey *string `json:"domainSigningPrivateKey,omitempty" tf:"domain_signing_private_key,omitempty"`

	// [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
	// +kubebuilder:validation:Optional
	DomainSigningSelector *string `json:"domainSigningSelector,omitempty" tf:"domain_signing_selector,omitempty"`

	// [Easy DKIM] The key length of the future DKIM key pair to be generated. This can be changed at most once per day. Valid values: RSA_1024_BIT, RSA_2048_BIT.
	// +kubebuilder:validation:Optional
	NextSigningKeyLength *string `json:"nextSigningKeyLength,omitempty" tf:"next_signing_key_length,omitempty"`
}

type EmailIdentityInitParameters struct {

	// The configuration of the DKIM authentication settings for an email domain identity.
	DKIMSigningAttributes []DKIMSigningAttributesInitParameters `json:"dkimSigningAttributes,omitempty" tf:"dkim_signing_attributes,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EmailIdentityObservation struct {

	// ARN of the Email Identity.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// The configuration of the DKIM authentication settings for an email domain identity.
	DKIMSigningAttributes []DKIMSigningAttributesObservation `json:"dkimSigningAttributes,omitempty" tf:"dkim_signing_attributes,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The email identity type. Valid values: EMAIL_ADDRESS, DOMAIN.
	IdentityType *string `json:"identityType,omitempty" tf:"identity_type,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies whether or not the identity is verified.
	VerifiedForSendingStatus *bool `json:"verifiedForSendingStatus,omitempty" tf:"verified_for_sending_status,omitempty"`
}

type EmailIdentityParameters struct {

	// The configuration set to use by default when sending from this identity. Note that any configuration set defined in the email sending request takes precedence.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sesv2/v1beta1.ConfigurationSet
	// +kubebuilder:validation:Optional
	ConfigurationSetName *string `json:"configurationSetName,omitempty" tf:"configuration_set_name,omitempty"`

	// Reference to a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameRef *v1.Reference `json:"configurationSetNameRef,omitempty" tf:"-"`

	// Selector for a ConfigurationSet in sesv2 to populate configurationSetName.
	// +kubebuilder:validation:Optional
	ConfigurationSetNameSelector *v1.Selector `json:"configurationSetNameSelector,omitempty" tf:"-"`

	// The configuration of the DKIM authentication settings for an email domain identity.
	// +kubebuilder:validation:Optional
	DKIMSigningAttributes []DKIMSigningAttributesParameters `json:"dkimSigningAttributes,omitempty" tf:"dkim_signing_attributes,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// EmailIdentitySpec defines the desired state of EmailIdentity
type EmailIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EmailIdentityParameters `json:"forProvider"`
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
	InitProvider EmailIdentityInitParameters `json:"initProvider,omitempty"`
}

// EmailIdentityStatus defines the observed state of EmailIdentity.
type EmailIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EmailIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentity is the Schema for the EmailIdentitys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EmailIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailIdentitySpec   `json:"spec"`
	Status            EmailIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmailIdentityList contains a list of EmailIdentitys
type EmailIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailIdentity `json:"items"`
}

// Repository type metadata.
var (
	EmailIdentity_Kind             = "EmailIdentity"
	EmailIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EmailIdentity_Kind}.String()
	EmailIdentity_KindAPIVersion   = EmailIdentity_Kind + "." + CRDGroupVersion.String()
	EmailIdentity_GroupVersionKind = CRDGroupVersion.WithKind(EmailIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&EmailIdentity{}, &EmailIdentityList{})
}
