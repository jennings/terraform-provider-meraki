package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wan1_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wan1_ip_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wan2_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"wan2_ip_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"using_cellular_failover": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

	resp, err := pc.MerakiDashboardAPI.Organizations.GetOrganizationDevicesStatuses(&req, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	statuses := make([]map[string]interface{}, 0)
	for _, device := range resp.Payload {
		device := device.(map[string]interface{})
		status := make(map[string]interface{})
		copyValueIfExists(device, status, "name", "name")
		copyValueIfExists(device, status, "productType", "product_type")
		copyValueIfExists(device, status, "mac", "mac")
		copyValueIfExists(device, status, "publicIp", "public_ip")
		copyValueIfExists(device, status, "networkId", "network_id")
		copyValueIfExists(device, status, "wan1Ip", "wan1_ip")
		copyValueIfExists(device, status, "wan2Ip", "wan2_ip")
		copyValueIfExists(device, status, "wan1IpType", "wan1_ip_type")
		copyValueIfExists(device, status, "wan2IpType", "wan2_ip_type")
		copyValueIfExists(device, status, "usingCellularFailover", "using_cellular_failover")
		copyValueIfExists(device, status, "tags", "tags")
		statuses = append(statuses, status)
	}

	if err := d.Set("statuses", statuses); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func copyValueIfExists(from, to map[string]interface{}, fromKey, toKey string) {
	if val, ok := from[fromKey]; ok {
		to[toKey] = val
	}
}
