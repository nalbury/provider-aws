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

type ServiceLinkedRoleInitParameters struct {

	// The AWS service to which this role is attached. You use a string similar to a URL but without the http:// in front. For example: elasticbeanstalk.amazonaws.com. To find the full list of services that support service-linked roles, check the docs.
	AwsServiceName *string `json:"awsServiceName,omitempty" tf:"aws_service_name,omitempty"`

	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix *string `json:"customSuffix,omitempty" tf:"custom_suffix,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServiceLinkedRoleObservation struct {

	// The Amazon Resource Name (ARN) specifying the role.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The AWS service to which this role is attached. You use a string similar to a URL but without the http:// in front. For example: elasticbeanstalk.amazonaws.com. To find the full list of services that support service-linked roles, check the docs.
	AwsServiceName *string `json:"awsServiceName,omitempty" tf:"aws_service_name,omitempty"`

	// The creation date of the IAM role.
	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	CustomSuffix *string `json:"customSuffix,omitempty" tf:"custom_suffix,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Amazon Resource Name (ARN) of the role.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the role.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The path of the role.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The stable and unique string identifying the role.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type ServiceLinkedRoleParameters struct {

	// The AWS service to which this role is attached. You use a string similar to a URL but without the http:// in front. For example: elasticbeanstalk.amazonaws.com. To find the full list of services that support service-linked roles, check the docs.
	// +kubebuilder:validation:Optional
	AwsServiceName *string `json:"awsServiceName,omitempty" tf:"aws_service_name,omitempty"`

	// Additional string appended to the role name. Not all AWS services support custom suffixes.
	// +kubebuilder:validation:Optional
	CustomSuffix *string `json:"customSuffix,omitempty" tf:"custom_suffix,omitempty"`

	// The description of the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServiceLinkedRoleSpec defines the desired state of ServiceLinkedRole
type ServiceLinkedRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceLinkedRoleParameters `json:"forProvider"`
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
	InitProvider ServiceLinkedRoleInitParameters `json:"initProvider,omitempty"`
}

// ServiceLinkedRoleStatus defines the observed state of ServiceLinkedRole.
type ServiceLinkedRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceLinkedRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceLinkedRole is the Schema for the ServiceLinkedRoles API. Provides an IAM service-linked role.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServiceLinkedRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.awsServiceName) || (has(self.initProvider) && has(self.initProvider.awsServiceName))",message="spec.forProvider.awsServiceName is a required parameter"
	Spec   ServiceLinkedRoleSpec   `json:"spec"`
	Status ServiceLinkedRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceLinkedRoleList contains a list of ServiceLinkedRoles
type ServiceLinkedRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceLinkedRole `json:"items"`
}

// Repository type metadata.
var (
	ServiceLinkedRole_Kind             = "ServiceLinkedRole"
	ServiceLinkedRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceLinkedRole_Kind}.String()
	ServiceLinkedRole_KindAPIVersion   = ServiceLinkedRole_Kind + "." + CRDGroupVersion.String()
	ServiceLinkedRole_GroupVersionKind = CRDGroupVersion.WithKind(ServiceLinkedRole_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceLinkedRole{}, &ServiceLinkedRoleList{})
}
