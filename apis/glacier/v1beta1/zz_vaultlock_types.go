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

type VaultLockInitParameters struct {

	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to false, the Glacier Lock Policy remains in a testing mode for 24 hours. Changing this from false to true will show as resource recreation, which is expected. Changing this from true to false is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock *bool `json:"completeLock,omitempty" tf:"complete_lock,omitempty"`

	// This should only be used in conjunction with complete_lock being set to true.
	IgnoreDeletionError *bool `json:"ignoreDeletionError,omitempty" tf:"ignore_deletion_error,omitempty"`

	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type VaultLockObservation struct {

	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to false, the Glacier Lock Policy remains in a testing mode for 24 hours. Changing this from false to true will show as resource recreation, which is expected. Changing this from true to false is not possible unless the Glacier Vault is recreated at the same time.
	CompleteLock *bool `json:"completeLock,omitempty" tf:"complete_lock,omitempty"`

	// Glacier Vault name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// This should only be used in conjunction with complete_lock being set to true.
	IgnoreDeletionError *bool `json:"ignoreDeletionError,omitempty" tf:"ignore_deletion_error,omitempty"`

	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The name of the Glacier Vault.
	VaultName *string `json:"vaultName,omitempty" tf:"vault_name,omitempty"`
}

type VaultLockParameters struct {

	// Boolean whether to permanently apply this Glacier Lock Policy. Once completed, this cannot be undone. If set to false, the Glacier Lock Policy remains in a testing mode for 24 hours. Changing this from false to true will show as resource recreation, which is expected. Changing this from true to false is not possible unless the Glacier Vault is recreated at the same time.
	// +kubebuilder:validation:Optional
	CompleteLock *bool `json:"completeLock,omitempty" tf:"complete_lock,omitempty"`

	// This should only be used in conjunction with complete_lock being set to true.
	// +kubebuilder:validation:Optional
	IgnoreDeletionError *bool `json:"ignoreDeletionError,omitempty" tf:"ignore_deletion_error,omitempty"`

	// JSON string containing the IAM policy to apply as the Glacier Vault Lock policy.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the Glacier Vault.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/glacier/v1beta1.Vault
	// +kubebuilder:validation:Optional
	VaultName *string `json:"vaultName,omitempty" tf:"vault_name,omitempty"`

	// Reference to a Vault in glacier to populate vaultName.
	// +kubebuilder:validation:Optional
	VaultNameRef *v1.Reference `json:"vaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in glacier to populate vaultName.
	// +kubebuilder:validation:Optional
	VaultNameSelector *v1.Selector `json:"vaultNameSelector,omitempty" tf:"-"`
}

// VaultLockSpec defines the desired state of VaultLock
type VaultLockSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultLockParameters `json:"forProvider"`
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
	InitProvider VaultLockInitParameters `json:"initProvider,omitempty"`
}

// VaultLockStatus defines the observed state of VaultLock.
type VaultLockStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultLockObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultLock is the Schema for the VaultLocks API. Manages a Glacier Vault Lock.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VaultLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.completeLock) || (has(self.initProvider) && has(self.initProvider.completeLock))",message="spec.forProvider.completeLock is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   VaultLockSpec   `json:"spec"`
	Status VaultLockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultLockList contains a list of VaultLocks
type VaultLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultLock `json:"items"`
}

// Repository type metadata.
var (
	VaultLock_Kind             = "VaultLock"
	VaultLock_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultLock_Kind}.String()
	VaultLock_KindAPIVersion   = VaultLock_Kind + "." + CRDGroupVersion.String()
	VaultLock_GroupVersionKind = CRDGroupVersion.WithKind(VaultLock_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultLock{}, &VaultLockList{})
}
