// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceGraph,
			TypeName: "aws_detective_graph",
			Name:     "Graph",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceInvitationAccepter,
			TypeName: "aws_detective_invitation_accepter",
			Name:     "Invitation Accepter",
		},
		{
			Factory:  ResourceMember,
			TypeName: "aws_detective_member",
			Name:     "Member",
		},
		{
			Factory:  ResourceOrganizationAdminAccount,
			TypeName: "aws_detective_organization_admin_account",
			Name:     "Organization Admin Account",
		},
		{
			Factory:  ResourceOrganizationConfiguration,
			TypeName: "aws_detective_organization_configuration",
			Name:     "Organization Configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Detective
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*detective.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws.Config))
	optFns := []func(*detective.Options){
		detective.WithEndpointResolverV2(newEndpointResolverV2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
		withExtraOptions(ctx, p, config),
	}

	return detective.NewFromConfig(cfg, optFns...), nil
}

// withExtraOptions returns a functional option that allows this service package to specify extra API client options.
// This option is always called after any generated options.
func withExtraOptions(ctx context.Context, sp conns.ServicePackage, config map[string]any) func(*detective.Options) {
	if v, ok := sp.(interface {
		withExtraOptions(context.Context, map[string]any) []func(*detective.Options)
	}); ok {
		optFns := v.withExtraOptions(ctx, config)

		return func(o *detective.Options) {
			for _, optFn := range optFns {
				optFn(o)
			}
		}
	}

	return func(*detective.Options) {}
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
