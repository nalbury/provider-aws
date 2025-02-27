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
	v1beta12 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/kms/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/sns/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.IAMRoles),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.IAMRoleRefs,
		Selector:      mg.Spec.ForProvider.IAMRoleSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRoles")
	}
	mg.Spec.ForProvider.IAMRoles = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.IAMRoleRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSKeyArn),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.KMSKeyArnRef,
		Selector:     mg.Spec.ForProvider.KMSKeyArnSelector,
		To: reference.To{
			List:    &v1beta11.KeyList{},
			Managed: &v1beta11.Key{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KMSKeyArn")
	}
	mg.Spec.ForProvider.KMSKeyArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KMSKeyArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneClusterParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneClusterParameterGroupNameSelector,
		To: reference.To{
			List:    &ClusterParameterGroupList{},
			Managed: &ClusterParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneClusterParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneClusterParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneClusterParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ReplicationSourceIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ReplicationSourceIdentifierRef,
		Selector:     mg.Spec.ForProvider.ReplicationSourceIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ReplicationSourceIdentifier")
	}
	mg.Spec.ForProvider.ReplicationSourceIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ReplicationSourceIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnapshotIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SnapshotIdentifierRef,
		Selector:     mg.Spec.ForProvider.SnapshotIdentifierSelector,
		To: reference.To{
			List:    &ClusterSnapshotList{},
			Managed: &ClusterSnapshot{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnapshotIdentifier")
	}
	mg.Spec.ForProvider.SnapshotIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnapshotIdentifierRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCSecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.VPCSecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta12.SecurityGroupList{},
			Managed: &v1beta12.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCSecurityGroupIds")
	}
	mg.Spec.ForProvider.VPCSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCSecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this ClusterEndpoint.
func (mg *ClusterEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterInstance.
func (mg *ClusterInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.ClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterIdentifier")
	}
	mg.Spec.ForProvider.ClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIdentifierRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneParameterGroupNameSelector,
		To: reference.To{
			List:    &ParameterGroupList{},
			Managed: &ParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneParameterGroupName")
	}
	mg.Spec.ForProvider.NeptuneParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NeptuneSubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NeptuneSubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.NeptuneSubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NeptuneSubnetGroupName")
	}
	mg.Spec.ForProvider.NeptuneSubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NeptuneSubnetGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ClusterSnapshot.
func (mg *ClusterSnapshot) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DBClusterIdentifier),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.DBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DBClusterIdentifier")
	}
	mg.Spec.ForProvider.DBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventSubscription.
func (mg *EventSubscription) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SnsTopicArn),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SnsTopicArnRef,
		Selector:     mg.Spec.ForProvider.SnsTopicArnSelector,
		To: reference.To{
			List:    &v1beta13.TopicList{},
			Managed: &v1beta13.Topic{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SnsTopicArn")
	}
	mg.Spec.ForProvider.SnsTopicArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SnsTopicArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GlobalCluster.
func (mg *GlobalCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceDBClusterIdentifier),
		Extract:      resource.ExtractParamPath("arn", true),
		Reference:    mg.Spec.ForProvider.SourceDBClusterIdentifierRef,
		Selector:     mg.Spec.ForProvider.SourceDBClusterIdentifierSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SourceDBClusterIdentifier")
	}
	mg.Spec.ForProvider.SourceDBClusterIdentifier = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SourceDBClusterIdentifierRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta12.SubnetList{},
			Managed: &v1beta12.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}
