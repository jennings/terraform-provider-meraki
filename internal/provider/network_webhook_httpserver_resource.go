package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/networks"
)

type networkWebhookHttpserverResourceType struct{}

func (t networkWebhookHttpserverResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "An HTTP Server registered to receive webhooks.",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed:            true,
				MarkdownDescription: "ID of the HTTP server",
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
				Type: types.StringType,
			},
			"network_id": {
				MarkdownDescription: "Network ID to add the HTTP server.",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"name": {
				MarkdownDescription: "Name of the HTTP server.",
				Required:            true,
				Type:                types.StringType,
			},
			"url": {
				MarkdownDescription: "URL that will receive the webhook.",
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"shared_secret": {
				MarkdownDescription: "Optional shared secret that can be used in the payload template to authenticate the request.",
				Optional:            true,
				Type:                types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"payload_template_id": {
				MarkdownDescription: "Payload template ID to send. Defaults to wpt_00001",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
		},
	}, nil
}

func (t networkWebhookHttpserverResourceType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return networkWebhookHttpserverResource{
		provider: provider,
	}, diags
}

type networkWebhookHttpserverResourceData struct {
	ID                types.String `tfsdk:"id"`
	NetworkID         types.String `tfsdk:"network_id"`
	Name              types.String `tfsdk:"name"`
	URL               types.String `tfsdk:"url"`
	SharedSecret      types.String `tfsdk:"shared_secret"`
	PayloadTemplateID types.String `tfsdk:"payload_template_id"`
}

type networkWebhookHttpserverResource struct {
	provider provider
}

func (r networkWebhookHttpserverResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan networkWebhookHttpserverResourceData

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.CreateNetworkWebhooksHTTPServerParams{
		Context:   ctx,
		NetworkID: plan.NetworkID.Value,
		CreateNetworkWebhooksHTTPServer: networks.CreateNetworkWebhooksHTTPServerBody{
			Name:         &plan.Name.Value,
			URL:          &plan.URL.Value,
			SharedSecret: plan.SharedSecret.Value,
		},
	}
	if !plan.PayloadTemplateID.Unknown {
		params.CreateNetworkWebhooksHTTPServer.PayloadTemplate = &networks.CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate{
			PayloadTemplateID: plan.PayloadTemplateID.Value,
		}
	}
	res, err := r.provider.client.Networks.CreateNetworkWebhooksHTTPServer(&params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("unable to create HTTP server: %v", err))
		return
	}

	err = buildNetworkWebhookHttpserverResourceData(&plan, res)
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

func (r networkWebhookHttpserverResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state networkWebhookHttpserverResourceData

	diags := req.State.Get(ctx, &state)
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
	if state.ID.Null || state.ID.Unknown || state.ID.Value == "" {
		resp.Diagnostics.AddError("State error", fmt.Sprintf("Expected state to have ID but had value: %v", state.ID))
		return
	}

	res, err := r.provider.client.Networks.GetNetworkWebhooksHTTPServer(&networks.GetNetworkWebhooksHTTPServerParams{
		Context:      ctx,
		NetworkID:    state.NetworkID.Value,
		HTTPServerID: state.ID.Value,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("request failed: %v", err))
		return
	}

	err = buildNetworkWebhookHttpserverResourceData(&state, res)
	if err != nil {
		resp.Diagnostics.AddError("Parse error", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r networkWebhookHttpserverResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan networkWebhookHttpserverResourceData
	var state networkWebhookHttpserverResourceData

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// build update request
	params := networks.UpdateNetworkWebhooksHTTPServerParams{
		Context:      ctx,
		NetworkID:    plan.NetworkID.Value,
		HTTPServerID: plan.ID.Value,
	}
	update := false
	sharedSecretChanged := false
	if plan.Name.Value != state.Name.Value {
		params.UpdateNetworkWebhooksHTTPServer.Name = plan.Name.Value
		update = true
	}
	if !plan.SharedSecret.Null && plan.SharedSecret.Value != state.SharedSecret.Value {
		params.UpdateNetworkWebhooksHTTPServer.SharedSecret = plan.SharedSecret.Value
		update = true
		sharedSecretChanged = true
	}
	if !plan.PayloadTemplateID.Unknown && !plan.PayloadTemplateID.Null && plan.PayloadTemplateID.Value != state.PayloadTemplateID.Value {
		params.UpdateNetworkWebhooksHTTPServer.PayloadTemplate = &networks.UpdateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate{
			PayloadTemplateID: plan.PayloadTemplateID.Value,
		}
		update = true
	}
	if !update {
		return
	}

	res, err := r.provider.client.Networks.UpdateNetworkWebhooksHTTPServer(&params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("unable to update HTTP server: %v", err))
	}

	err = buildNetworkWebhookHttpserverResourceData(&state, res)
	if err != nil {
		resp.Diagnostics.AddError("Parse error", err.Error())
		return
	}
	if sharedSecretChanged {
		state.SharedSecret = types.String{Value: params.UpdateNetworkWebhooksHTTPServer.SharedSecret}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r networkWebhookHttpserverResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state networkWebhookHttpserverResourceData

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.provider.client.Networks.DeleteNetworkWebhooksHTTPServer(&networks.DeleteNetworkWebhooksHTTPServerParams{
		Context:      ctx,
		NetworkID:    state.NetworkID.Value,
		HTTPServerID: state.ID.Value,
	}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client error", fmt.Sprintf("error with network request: %v", err))
	}
}

func (r networkWebhookHttpserverResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	parts := strings.Split(req.ID, ",")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected import identifier",
			fmt.Sprintf("Expected identifier with format 'network_id,id', got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("network_id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("id"), parts[1])...)
}

func buildNetworkWebhookHttpserverResourceData(data *networkWebhookHttpserverResourceData, res merakiResponse) error {
	p := res.GetPayload()
	payload, ok := p.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected response %T to have payload type map[string]interface{} but got: %T", res, p)
	}

	data.ID = types.String{Value: payload["id"].(string)}
	data.Name = types.String{Value: payload["name"].(string)}
	data.NetworkID = types.String{Value: payload["networkId"].(string)}
	data.URL = types.String{Value: payload["url"].(string)}
	data.PayloadTemplateID = types.String{
		Value: payload["payloadTemplate"].(map[string]interface{})["payloadTemplateId"].(string),
	}
	return nil
}
