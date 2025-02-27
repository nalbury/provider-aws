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

type PolicyAttributeInitParameters struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PolicyAttributeObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PolicyAttributeParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.Policy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("policy_name",false)
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Reference to a Policy in elb to populate value.
	// +kubebuilder:validation:Optional
	ValueRef *v1.Reference `json:"valueRef,omitempty" tf:"-"`

	// Selector for a Policy in elb to populate value.
	// +kubebuilder:validation:Optional
	ValueSelector *v1.Selector `json:"valueSelector,omitempty" tf:"-"`
}

type PolicyInitParameters struct {

	// Policy attribute to apply to the policy.
	PolicyAttribute []PolicyAttributeInitParameters `json:"policyAttribute,omitempty" tf:"policy_attribute,omitempty"`

	// The name of the load balancer policy.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// The policy type.
	PolicyTypeName *string `json:"policyTypeName,omitempty" tf:"policy_type_name,omitempty"`
}

type PolicyObservation struct {

	// The ID of the policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load balancer on which the policy is defined.
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Policy attribute to apply to the policy.
	PolicyAttribute []PolicyAttributeObservation `json:"policyAttribute,omitempty" tf:"policy_attribute,omitempty"`

	// The name of the load balancer policy.
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// The policy type.
	PolicyTypeName *string `json:"policyTypeName,omitempty" tf:"policy_type_name,omitempty"`
}

type PolicyParameters struct {

	// The load balancer on which the policy is defined.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +kubebuilder:validation:Optional
	LoadBalancerName *string `json:"loadBalancerName,omitempty" tf:"load_balancer_name,omitempty"`

	// Reference to a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameRef *v1.Reference `json:"loadBalancerNameRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate loadBalancerName.
	// +kubebuilder:validation:Optional
	LoadBalancerNameSelector *v1.Selector `json:"loadBalancerNameSelector,omitempty" tf:"-"`

	// Policy attribute to apply to the policy.
	// +kubebuilder:validation:Optional
	PolicyAttribute []PolicyAttributeParameters `json:"policyAttribute,omitempty" tf:"policy_attribute,omitempty"`

	// The name of the load balancer policy.
	// +kubebuilder:validation:Optional
	PolicyName *string `json:"policyName,omitempty" tf:"policy_name,omitempty"`

	// The policy type.
	// +kubebuilder:validation:Optional
	PolicyTypeName *string `json:"policyTypeName,omitempty" tf:"policy_type_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
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
	InitProvider PolicyInitParameters `json:"initProvider,omitempty"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. Provides a load balancer policy, which can be attached to an ELB listener or backend server.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyName) || (has(self.initProvider) && has(self.initProvider.policyName))",message="spec.forProvider.policyName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyTypeName) || (has(self.initProvider) && has(self.initProvider.policyTypeName))",message="spec.forProvider.policyTypeName is a required parameter"
	Spec   PolicySpec   `json:"spec"`
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
