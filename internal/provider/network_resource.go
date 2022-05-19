package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/networks"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

type networkResourceType struct{}

func (t networkResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "A Meraki network",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "ID of the network",
				Computed:            true,
				Type:                types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"organization_id": {
				MarkdownDescription: "Organization ID to add the HTTP server.",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"name": {
				MarkdownDescription: "Name of the network.",
				Required:            true,
				Type:                types.StringType,
			},
			"product_types": {
				MarkdownDescription: "Product types. Possible values are: wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, environmental.",
				Required:            true,
				Type: types.SetType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"tags": {
				MarkdownDescription: "Tags",
				Optional:            true,
				Type: types.SetType{
					ElemType: types.StringType,
				},
			},
		},
	}, nil
}

func (t networkResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return networkResource{
		provider: provider,
	}, diags
}

type networkResourceData struct {
	ID             types.String `tfsdk:"id"`
	OrganizationID types.String `tfsdk:"organization_id"`
	Name           types.String `tfsdk:"name"`
	ProductTypes   types.Set    `tfsdk:"product_types"`
	Tags           types.Set    `tfsdk:"tags"`
}

type networkResource struct {
	provider provider
}

func (r networkResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan networkResourceData

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := organizations.CreateOrganizationNetworkParams{
		Context:        ctx,
		OrganizationID: plan.OrganizationID.Value,
		CreateOrganizationNetwork: organizations.CreateOrganizationNetworkBody{
			Name: &plan.Name.Value,
		},
	}
	if !plan.ProductTypes.Unknown {
		params.CreateOrganizationNetwork.ProductTypes = setToStringArray(plan.ProductTypes)
	}
	if !plan.Tags.Unknown {
		params.CreateOrganizationNetwork.Tags = setToStringArray(plan.Tags)
	}
	res, err := r.provider.client.Organizations.CreateOrganizationNetwork(&params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("unable to create network: %v", err))
		return
	}

	var state networkResourceData
	buildNetworkFromPayload(&plan, &state, res.Payload)
	if err != nil {
		resp.Diagnostics.AddError("Parse error", err.Error())
		return
	}

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func setToStringArray(set types.Set) []string {
	ss := make([]string, len(set.Elems))
	for i, v := range set.Elems {
		ss[i] = v.(types.String).Value
	}
	return ss
}

func buildNetworkFromPayload(plan, data *networkResourceData, payload interface{}) {
	p := payload.(map[string]interface{})
	data.ID = types.String{Value: p["id"].(string)}
	data.Name = types.String{Value: p["name"].(string)}
	data.OrganizationID = types.String{Value: p["organizationId"].(string)}

	pts := p["productTypes"].([]interface{})
	data.ProductTypes = types.Set{ElemType: types.StringType, Elems: make([]attr.Value, len(pts))}
	for i, pt := range pts {
		data.ProductTypes.Elems[i] = types.String{Value: pt.(string)}
	}

	if plan != nil && !plan.Tags.Null {
		tags := p["tags"].([]interface{})
		data.Tags = types.Set{ElemType: types.StringType, Elems: make([]attr.Value, len(tags))}
		for i, t := range tags {
			data.Tags.Elems[i] = types.String{Value: t.(string)}
		}
	}
}

func (r networkResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state networkResourceData

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if state.ID.Null || state.ID.Unknown || state.ID.Value == "" {
		resp.Diagnostics.AddError("State error", fmt.Sprintf("Expected state to have ID but had value: %v", state.ID))
		return
	}

	res, err := r.provider.client.Networks.GetNetwork(&networks.GetNetworkParams{
		Context:   ctx,
		NetworkID: state.ID.Value,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("request failed: %v", err))
		return
	}

	buildNetworkFromPayload(nil, &state, res.Payload)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r networkResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan networkResourceData
	var state networkResourceData

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// build update request
	params := networks.UpdateNetworkParams{
		Context:       ctx,
		NetworkID:     plan.ID.Value,
		UpdateNetwork: networks.UpdateNetworkBody{},
	}
	if !plan.Name.Equal(state.Name) {
		params.UpdateNetwork.Name = plan.Name.Value
	}
	if !plan.Tags.Null && !plan.Tags.Equal(state.Tags) {
		params.UpdateNetwork.Tags = setToStringArray(plan.Tags)
	}

	res, err := r.provider.client.Networks.UpdateNetwork(&params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("unable to update HTTP server: %v", err))
	}

	buildNetworkFromPayload(&plan, &state, res.Payload)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r networkResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state networkResourceData

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.provider.client.Networks.DeleteNetwork(&networks.DeleteNetworkParams{
		Context:   ctx,
		NetworkID: state.ID.Value,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("error with network request: %v", err))
	}
}

func (r networkResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	parts := strings.Split(req.ID, ",")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected import identifier",
			fmt.Sprintf("Expected identifier with format 'organization_id,id', got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("organization_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), parts[1])...)
}
