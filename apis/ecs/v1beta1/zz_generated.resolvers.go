/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/autoscaling/v1beta1"
	v1beta13 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	v1beta12 "github.com/upbound/provider-aws/apis/elbv2/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this CapacityProvider.
func (mg *CapacityProvider) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AutoScalingGroupProvider); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArnRef,
			Selector:     mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArnSelector,
			To: reference.To{
				List:    &v1beta1.AutoscalingGroupList{},
				Managed: &v1beta1.AutoscalingGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArn")
		}
		mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AutoScalingGroupProvider[i3].AutoScalingGroupArnRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this ClusterCapacityProviders.
func (mg *ClusterCapacityProviders) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterNameRef,
		Selector:     mg.Spec.ForProvider.ClusterNameSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterName")
	}
	mg.Spec.ForProvider.ClusterName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterRef,
		Selector:     mg.Spec.ForProvider.ClusterSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IAMRole),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.IAMRoleRef,
		Selector:     mg.Spec.ForProvider.IAMRoleSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IAMRole")
	}
	mg.Spec.ForProvider.IAMRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IAMRoleRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LoadBalancer); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnRef,
			Selector:     mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnSelector,
			To: reference.To{
				List:    &v1beta12.LBTargetGroupList{},
				Managed: &v1beta12.LBTargetGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LoadBalancer[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfiguration); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroupRefs,
			Selector:      mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroupSelector,
			To: reference.To{
				List:    &v1beta13.SecurityGroupList{},
				Managed: &v1beta13.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkConfiguration[i3].SecurityGroupRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfiguration); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkConfiguration[i3].Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.NetworkConfiguration[i3].SubnetRefs,
			Selector:      mg.Spec.ForProvider.NetworkConfiguration[i3].SubnetSelector,
			To: reference.To{
				List:    &v1beta13.SubnetList{},
				Managed: &v1beta13.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration[i3].Subnets")
		}
		mg.Spec.ForProvider.NetworkConfiguration[i3].Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkConfiguration[i3].SubnetRefs = mrsp.ResolvedReferences

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TaskDefinition),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TaskDefinitionRef,
		Selector:     mg.Spec.ForProvider.TaskDefinitionSelector,
		To: reference.To{
			List:    &TaskDefinitionList{},
			Managed: &TaskDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TaskDefinition")
	}
	mg.Spec.ForProvider.TaskDefinition = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TaskDefinitionRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TaskDefinition.
func (mg *TaskDefinition) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExecutionRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ExecutionRoleArnRef,
		Selector:     mg.Spec.ForProvider.ExecutionRoleArnSelector,
		To: reference.To{
			List:    &v1beta11.RoleList{},
			Managed: &v1beta11.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExecutionRoleArn")
	}
	mg.Spec.ForProvider.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExecutionRoleArnRef = rsp.ResolvedReference

	return nil
}
