// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceComputeEnvironment,
			TypeName: "aws_batch_compute_environment",
		},
		{
			Factory:  DataSourceJobQueue,
			TypeName: "aws_batch_job_queue",
		},
		{
			Factory:  DataSourceSchedulingPolicy,
			TypeName: "aws_batch_scheduling_policy",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceComputeEnvironment,
			TypeName: "aws_batch_compute_environment",
			Name:     "Compute Environment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceJobDefinition,
			TypeName: "aws_batch_job_definition",
			Name:     "Job Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceJobQueue,
			TypeName: "aws_batch_job_queue",
			Name:     "Job Queue",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSchedulingPolicy,
			TypeName: "aws_batch_scheduling_policy",
			Name:     "Job Definition",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Batch
}

var ServicePackage = &servicePackage{}
