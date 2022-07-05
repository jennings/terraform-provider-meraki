package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

type networksDataSourceType struct{}

func (t networksDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "Use this data source to retrieve a list of networks in an organization.",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "Organization ID",
				Computed:            true,
				Type:                types.StringType,
			},
			"organization_id": {
				MarkdownDescription: "Organization ID",
				Required:            true,
				Type:                types.StringType,
			},
			"values": {
				MarkdownDescription: "Networks in the organization",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						MarkdownDescription: "Network ID",
						Computed:            true,
						Type:                types.StringType,
					},
					"name": {
						MarkdownDescription: "Device name",
						Computed:            true,
						Type:                types.StringType,
					},
				}),
			},
		},
	}, nil
}

func (t networksDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return networksDataSource{
		provider: provider,
	}, diags
}

type network struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type networksDataSourceData struct {
	ID             types.String `tfsdk:"id"`
	OrganizationID types.String `tfsdk:"organization_id"`
	Values         []network    `tfsdk:"values"`
}

type networksDataSource struct {
	*provider
}

func (d networksDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data networksDataSourceData

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// example, err := d.provider.client.ReadExample(...)
	// if err != nil {
	//     resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read example, got error: %s", err))
	//     return
	// }
	params := organizations.GetOrganizationNetworksParams{
		OrganizationID: fmt.Sprint(data.OrganizationID.Value),
		Context:        ctx,
	}

	rsp, err := d.provider.client.Organizations.GetOrganizationNetworks(&params, nil)
	if err != nil {
		resp.Diagnostics.AddError("unable to fetch organization devices statuses", err.Error())
		return
	}

	data.Values, err = mapNetworksToState(rsp, resp.Diagnostics)
	if err != nil {
		resp.Diagnostics.AddError("unable to map device statuses", err.Error())
		return
	}

	// Resources are required to have an "id" attribute
	data.ID = data.OrganizationID

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func mapNetworksToState(resp *organizations.GetOrganizationNetworksOK, diags diag.Diagnostics) ([]network, error) {
	l := make([]network, 0, len(resp.Payload))

	for _, device := range resp.Payload {
		n := device.(map[string]interface{})
		s, err := mapNetworkToState(n, diags)
		if err != nil {
			return nil, err
		}
		l = append(l, s)
	}

	return l, nil
}

func mapNetworkToState(status map[string]interface{}, diags diag.Diagnostics) (network, error) {
	state := network{}

	for key, value := range status {
		switch key {
		case "id":
			v, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.ID = types.String{Value: v}

		case "name":
			v, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.Name = types.String{Value: v}
		}
	}

	return state, nil
}
