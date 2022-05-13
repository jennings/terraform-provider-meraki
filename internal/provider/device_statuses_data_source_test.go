package provider

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/jennings/terraform-provider-meraki/internal/provider/client/organizations"
)

func TestAccDataSourceDeviceStatusesWithFilters(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: fmt.Sprintf(`
					data "meraki_device_statuses" "test" {
						organization_id = %q
						product_types = ["appliance"]
					}`, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.meraki_device_statuses.test", "values"),
				),
			},
		},
	})
}

func TestAccDataSourceDeviceStatuses(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: fmt.Sprintf(`
					data "meraki_device_statuses" "test" {
						organization_id = %q
					}`, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.meraki_device_statuses.test", "values"),
				),
			},
		},
	})
}

func TestMapDeviceStatusesToState(t *testing.T) {
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

	resp := organizations.GetOrganizationDevicesStatusesOK{Payload: payload}
	var diags diag.Diagnostics

	statuses, err := mapDeviceStatusesToState(&resp, diags)
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

func assertStringEqual(t *testing.T, actual types.String, expected, key string) {
	if actual.Value != expected {
		t.Errorf("expected %v %v, got: %v", key, expected, actual)
	}
}
