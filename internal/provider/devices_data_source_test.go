package provider

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

func TestAccDataSourceOrganizationDevicesWithFilters(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: `
				 	data "meraki_organizations" "o" {}
					data "meraki_organization_devices" "test" {
						organization_id = tolist(data.meraki_organizations.o)[0].id
						product_types = ["appliance"]
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.meraki_organization_devices.test", "id", "example-id"),
				),
			},
		},
	})
}

func TestAccDataSourceOrganizationDevices(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: `
					data "meraki_organization_devices" "test" {
						organization_id = 123
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.meraki_device_statuses.test", "id", "example-id"),
				),
			},
		},
	})
}

func TestMapOrganizationDevicesToState(t *testing.T) {
	jsonResponse := `[
		{
			"name": "NAME",
			"mac": "MAC",
			"networkId": "NETWORKID",
			"productType": "PRODUCTTYPE",
			"tags": [
				"tag1",
				"tag2",
				123
			],
			"unknownString": "UNKNOWNSTRING",
			"unknownBoolean": true,
			"unknownNumber": 123
		}
	]`
	var payload []interface{}
	err := json.Unmarshal([]byte(jsonResponse), &payload)
	if err != nil {
		t.Fatal(err)
	}

	resp := organizations.GetOrganizationDevicesOK{Payload: payload}
	var diags diag.Diagnostics

	statuses, err := mapOrganizationDevicesToState(&resp, diags)
	if err != nil {
		t.Fatal(err)
	}

	if len(statuses) != 1 {
		t.Fatalf("expected 1 device status, got: %v value: %+#v", len(statuses), statuses)
	}
	status := statuses[0]
	assertStringEqual(t, status.Name, "NAME", "name")
	assertStringEqual(t, status.ProductType, "PRODUCTTYPE", "product_type")
}
