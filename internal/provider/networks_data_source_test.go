package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworksDataSource(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: fmt.Sprintf(`
					data "meraki_networks" "test" {
						organization_id = "%v"
					}`, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.meraki_networks.test", "values"),
				),
			},
		},
	})
}
