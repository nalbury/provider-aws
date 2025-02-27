/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Vault.
func (mg *Vault) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Notification); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Notification[i3].SnsTopic),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.Notification[i3].SnsTopicRef,
			Selector:     mg.Spec.ForProvider.Notification[i3].SnsTopicSelector,
			To: reference.To{
				List:    &v1beta1.TopicList{},
				Managed: &v1beta1.Topic{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Notification[i3].SnsTopic")
		}
		mg.Spec.ForProvider.Notification[i3].SnsTopic = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Notification[i3].SnsTopicRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this VaultLock.
func (mg *VaultLock) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VaultName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VaultNameRef,
		Selector:     mg.Spec.ForProvider.VaultNameSelector,
		To: reference.To{
			List:    &VaultList{},
			Managed: &Vault{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VaultName")
	}
	mg.Spec.ForProvider.VaultName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VaultNameRef = rsp.ResolvedReference

	return nil
}
