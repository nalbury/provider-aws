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

type CustomDataIdentifierInitParameters struct {

	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords []*string `json:"ignoreWords,omitempty" tf:"ignore_words,omitempty"`

	// An array that lists specific character sequences (keywords), one of which must be within proximity (maximum_match_distance) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`

	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance *float64 `json:"maximumMatchDistance,omitempty" tf:"maximum_match_distance,omitempty"`

	// A custom name for the custom data identifier. The name can contain as many as 128 characters. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomDataIdentifierObservation struct {

	// The Amazon Resource Name (ARN) of the custom data identifier.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time, in UTC and extended RFC 3339 format, when the Amazon Macie account was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique identifier (ID) of the macie custom data identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords []*string `json:"ignoreWords,omitempty" tf:"ignore_words,omitempty"`

	// An array that lists specific character sequences (keywords), one of which must be within proximity (maximum_match_distance) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`

	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance *float64 `json:"maximumMatchDistance,omitempty" tf:"maximum_match_distance,omitempty"`

	// A custom name for the custom data identifier. The name can contain as many as 128 characters. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type CustomDataIdentifierParameters struct {

	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	// +kubebuilder:validation:Optional
	IgnoreWords []*string `json:"ignoreWords,omitempty" tf:"ignore_words,omitempty"`

	// An array that lists specific character sequences (keywords), one of which must be within proximity (maximum_match_distance) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	// +kubebuilder:validation:Optional
	Keywords []*string `json:"keywords,omitempty" tf:"keywords,omitempty"`

	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	// +kubebuilder:validation:Optional
	MaximumMatchDistance *float64 `json:"maximumMatchDistance,omitempty" tf:"maximum_match_distance,omitempty"`

	// A custom name for the custom data identifier. The name can contain as many as 128 characters. Conflicts with name_prefix.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	// +kubebuilder:validation:Optional
	Regex *string `json:"regex,omitempty" tf:"regex,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// CustomDataIdentifierSpec defines the desired state of CustomDataIdentifier
type CustomDataIdentifierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDataIdentifierParameters `json:"forProvider"`
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
	InitProvider CustomDataIdentifierInitParameters `json:"initProvider,omitempty"`
}

// CustomDataIdentifierStatus defines the observed state of CustomDataIdentifier.
type CustomDataIdentifierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDataIdentifierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDataIdentifier is the Schema for the CustomDataIdentifiers API. Provides a resource to manage an AWS Macie Custom Data Identifier.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CustomDataIdentifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDataIdentifierSpec   `json:"spec"`
	Status            CustomDataIdentifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDataIdentifierList contains a list of CustomDataIdentifiers
type CustomDataIdentifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDataIdentifier `json:"items"`
}

// Repository type metadata.
var (
	CustomDataIdentifier_Kind             = "CustomDataIdentifier"
	CustomDataIdentifier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDataIdentifier_Kind}.String()
	CustomDataIdentifier_KindAPIVersion   = CustomDataIdentifier_Kind + "." + CRDGroupVersion.String()
	CustomDataIdentifier_GroupVersionKind = CRDGroupVersion.WithKind(CustomDataIdentifier_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDataIdentifier{}, &CustomDataIdentifierList{})
}
