package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

func dataSourceDeviceStatuses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDeviceStatusesRead,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"product_types": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"models": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"statuses": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"attributes": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     &schema.Schema{},
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceStatusesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pc := m.(*providerContext)
	var diags diag.Diagnostics

	req := organizations.GetOrganizationDevicesStatusesParams{
		OrganizationID: fmt.Sprint(d.Get("organization_id")),
		Context:        ctx,
	}
	pts, ok := d.GetOk("product_types")
	if ok {
		pts := pts.([]interface{})
		req.ProductTypes = make([]string, len(pts))
		for i, pt := range pts {
			req.ProductTypes[i] = pt.(string)
		}
	}
	ms, ok := d.GetOk("models")
	if ok {
		ms := ms.([]interface{})
		req.Models = make([]string, len(ms))
		for i, m := range ms {
			req.Models[i] = m.(string)
		}
	}

	tflog.Trace(ctx, "requesting organization devices statuses",
		"organization_id", req.OrganizationID)

	resp, err := pc.MerakiDashboardAPI.Organizations.GetOrganizationDevicesStatuses(&req, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Trace(ctx, "request succeeded",
		"count", len(resp.Payload))

	statuses := mapDeviceStatusesToState(resp)

	if err := d.Set("statuses", statuses); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func mapDeviceStatusesToState(resp *organizations.GetOrganizationDevicesStatusesOK) []map[string]interface{} {
	statuses := make([]map[string]interface{}, 0)
	for _, device := range resp.Payload {
		device := device.(map[string]interface{})
		s := mapDeviceStatusToState(device)
		statuses = append(statuses, s)
	}

	return statuses
}

func mapDeviceStatusToState(status map[string]interface{}) map[string]interface{} {
	state := make(map[string]interface{})
	attributes := make(map[string]interface{})
	state["attributes"] = attributes

	for key, value := range status {
		isAttr := false

		switch key {
		case "name":
		case "mac":
		case "serial":
		case "tags":
			// already in snake_case

		// format these as snake_case
		case "productType":
			key = "product_type"
		case "networkId":
			key = "network_id"

		// everything else goes in the attributes map
		default:
			isAttr = true
		}

		if isAttr {
			attributes[key] = value
		} else {
			state[key] = value
		}
	}

	return state
}
