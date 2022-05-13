package provider

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

func TestAccDataSourceOrganizationDevicesWithFilters(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: fmt.Sprintf(`
					data "meraki_devices" "test" {
						organization_id = "%v"
						product_types = ["appliance"]
					}`, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.meraki_devices.test", "values"),
				),
			},
		},
	})
}

func TestAccDataSourceOrganizationDevices(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: fmt.Sprintf(`
					data "meraki_devices" "test" {
						organization_id = "%v"
					}`, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.meraki_device_statuses.test", "values"),
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
