package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

type organizationsDataSourceType struct{}

func (t organizationsDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Organizations available using the given API key.",

		Attributes: map[string]tfsdk.Attribute{
			"values": {
				MarkdownDescription: "List of organizations",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						MarkdownDescription: "Organization ID",
						Computed:            true,
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
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t organizationsDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return organizationsDataSource{
		provider: provider,
	}, diags
}

type organizationsDataSource struct {
	*provider
}

type organizationsDataSourceData struct {
	Values []organization `tfsdk:"values"`
}

func (d organizationsDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	res, err := d.provider.client.Organizations.GetOrganizations(&organizations.GetOrganizationsParams{
		Context: ctx,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("%v", err))
		return
	}

	var data organizationsDataSourceData
	data.Values = buildOrganizationsFromPayload(res.Payload)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func buildOrganizationsFromPayload(payload []*organizations.GetOrganizationsOKBodyItems0) []organization {
	orgs := make([]organization, 0, len(payload))
	for _, value := range payload {
		org := organization{
			ID:   value.ID,
			Name: value.Name,
			URL:  value.URL,
		}
		orgs = append(orgs, org)
	}

	return orgs
}
