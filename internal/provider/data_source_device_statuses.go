package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

type deviceStatusesDataSourceType struct{}

func (t deviceStatusesDataSourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example data source",

		Attributes: map[string]tfsdk.Attribute{
			"organization_id": {
				MarkdownDescription: "Organization ID",
				Required:            true,
				Type:                types.Int64Type,
			},
			"product_types": {
				MarkdownDescription: "Filter for device product types",
				Optional:            true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
			},
			"models": {
				MarkdownDescription: "Filter for device models",
				Optional:            true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
			},
			"values": {
				MarkdownDescription: "List of device statuses returned from the Meraki API",
				Computed:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "Device name",
						Computed:            true,
						Type:                types.StringType,
					},
					"serial": {
						MarkdownDescription: "Serial number",
						Computed:            true,
						Type:                types.StringType,
					},
					"product_type": {
						MarkdownDescription: "Product type",
						Computed:            true,
						Type:                types.StringType,
					},
					"model": {
						MarkdownDescription: "Model name",
						Computed:            true,
						Type:                types.StringType,
					},
					"tags": {
						MarkdownDescription: "Tags",
						Computed:            true,
						Type: types.SetType{
							ElemType: types.StringType,
						},
					},
					"attributes": {
						MarkdownDescription: "All attributes as returned by the API. Keys are likely to be camelCase.",
						Computed:            true,
						Type: types.MapType{
							ElemType: types.StringType,
						},
					},
					"raw_json": {
						MarkdownDescription: "JSON encoded value returned by the Meraki API",
						Computed:            true,
						Type:                types.StringType,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (t deviceStatusesDataSourceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return deviceStatusesDataSource{
		provider: provider,
	}, diags
}

type deviceStatusesValuesDataSourceData struct {
	Name        types.String            `tfsdk:"name"`
	Serial      types.String            `tfsdk:"serial"`
	ProductType types.String            `tfsdk:"product_type"`
	Model       types.String            `tfsdk:"model"`
	Tags        []types.String          `tfsdk:"tags"`
	RawJson     types.String            `tfsdk:"raw_json"`
	Attributes  map[string]types.String `tfsdk:"attributes"`
}

type deviceStatusesDataSourceData struct {
	OrganizationID types.Int64                          `tfsdk:"organization_id"`
	ProductTypes   types.List                           `tfsdk:"product_types"`
	Models         types.List                           `tfsdk:"models"`
	Values         []deviceStatusesValuesDataSourceData `tfsdk:"values"`
}

type deviceStatusesDataSource struct {
	provider provider
}

func (d deviceStatusesDataSource) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var data deviceStatusesDataSourceData

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
	rq := organizations.GetOrganizationDevicesStatusesParams{
		OrganizationID: fmt.Sprint(data.OrganizationID.Value),
		Context:        ctx,
	}
	if !data.ProductTypes.Null {
		rq.ProductTypes = make([]string, 0, len(data.ProductTypes.Elems))
		for _, pt := range data.ProductTypes.Elems {
			pt, err := pt.ToTerraformValue(ctx)
			if err != nil {
				resp.Diagnostics.AddError("could not convert product_types value to Terraform value", err.Error())
				continue
			}
			var ptStr string
			err = pt.As(&ptStr)
			if err != nil {
				resp.Diagnostics.AddError("could not convert product_types value to string", err.Error())
				continue
			}
			rq.ProductTypes = append(rq.ProductTypes, ptStr)
		}
	}
	if !data.Models.Null {
		rq.Models = make([]string, 0, len(data.Models.Elems))
		for _, m := range data.Models.Elems {
			m, err := m.ToTerraformValue(ctx)
			if err != nil {
				resp.Diagnostics.AddError("could not convert models value to Terraform value", err.Error())
				continue
			}
			var mStr string
			err = m.As(&mStr)
			if err != nil {
				resp.Diagnostics.AddError("could not convert models value to string", err.Error())
				continue
			}
			rq.Models = append(rq.Models, mStr)
		}
	}

	rsp, err := d.provider.client.Organizations.GetOrganizationDevicesStatuses(&rq, nil)
	if err != nil {
		resp.Diagnostics.AddError("unable to fetch organization devices statuses", err.Error())
		return
	}

	data.Values, err = mapDeviceStatusesToState(rsp, resp.Diagnostics)
	if err != nil {
		resp.Diagnostics.AddError("unable to map device statuses", err.Error())
		return
	}

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func mapDeviceStatusesToState(resp *organizations.GetOrganizationDevicesStatusesOK, diags diag.Diagnostics) ([]deviceStatusesValuesDataSourceData, error) {
	l := make([]deviceStatusesValuesDataSourceData, 0, len(resp.Payload))

	for _, device := range resp.Payload {
		device := device.(map[string]interface{})
		s, err := mapDeviceStatusToState(device, diags)
		if err != nil {
			return nil, err
		}
		l = append(l, s)
	}

	return l, nil
}

func mapDeviceStatusToState(status map[string]interface{}, diags diag.Diagnostics) (deviceStatusesValuesDataSourceData, error) {
	state := deviceStatusesValuesDataSourceData{
		Attributes: make(map[string]types.String),
	}
	raw_json, err := json.Marshal(status)
	if err != nil {
		return deviceStatusesValuesDataSourceData{}, err
	}
	state.RawJson = types.String{Value: string(raw_json)}

	for key, value := range status {
		switch key {
		case "name":
			v, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.Name = types.String{Value: v}

		case "serial":
			v, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.Serial = types.String{Value: v}

		case "tags":
			value, ok := value.([]interface{})
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be []interface{}, got: %T", key, value))
				continue
			}
			elems := make([]types.String, 0, len(value))
			for _, tag := range value {
				tagStr, ok := tag.(string)
				if !ok {
					diags.AddWarning("unexpected data type", fmt.Sprintf("expected tag to be string, got: %T", tag))
					continue
				} else {
					elems = append(elems, types.String{Value: tagStr})
				}
			}
			state.Tags = elems

		case "model":
			model, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.Model = types.String{Value: model}
			continue

		case "productType":
			productType, ok := value.(string)
			if !ok {
				diags.AddWarning("unexpected data type", fmt.Sprintf("expected '%v' to be string, got: %T", key, value))
				continue
			}
			state.ProductType = types.String{Value: productType}
			continue

		// everything else goes in the attributes map
		default:
			if value == nil {
				state.Attributes[key] = types.String{Null: true}
			} else {
				state.Attributes[key] = types.String{Value: fmt.Sprintf("%v", value)}
			}
		}
	}

	return state, nil
}
