package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

type organizationDataSourceType struct{}

func (t organizationDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "Organization data",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "Organization ID",
				Required:            true,
				Type:                types.StringType,
			},
			"name": {
				MarkdownDescription: "Organization name",
				Computed:            true,
				Type:                types.StringType,
			},
			"url": {
				MarkdownDescription: "URL",
				Computed:            true,
				Type:                types.StringType,
			},
		},
	}, nil
}

func (t organizationDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return organizationDataSource{
		provider: provider,
	}, diags
}

type organizationDataSource struct {
	*provider
}

type organization struct {
	ID   string `tfsdk:"id"`
	Name string `tfsdk:"name"`
	URL  string `tfsdk:"url"`
}

func (d organizationDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var id string
	diags := req.Config.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	res, err := d.provider.client.Organizations.GetOrganization(&organizations.GetOrganizationParams{
		Context:        ctx,
		OrganizationID: id,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("%v", err))
		return
	}

	state := organization{
		ID:   res.Payload.ID,
		Name: res.Payload.Name,
		URL:  res.Payload.URL,
	}
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
