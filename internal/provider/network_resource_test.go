package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkResource(t *testing.T) {
	vars := readTestVars(t)
	name1 := generateName("one")
	name2 := generateName("two")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: fmt.Sprintf(`
					resource "meraki_network" "test" {
						name            = %q
						organization_id = %q
						product_types   = ["appliance"]
					}`, name1, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("meraki_network.test", "id"),
					resource.TestCheckResourceAttr("meraki_network.test", "organization_id", vars.OrganizationID),
					resource.TestCheckResourceAttr("meraki_network.test", "name", name1),
				),
			},
			// // ImportState testing
			// {
			// 	ResourceName:      "meraki_network.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateIdFunc: func(s *terraform.State) (string, error) {
			// 		r := s.Children("root").Resources["meraki_network"]
			// 		return "", nil
			// 	},
			// },
			// Update and Read testing
			{
				Config: fmt.Sprintf(`
					resource "meraki_network" "test" {
						name            = %q
						organization_id = %q
						product_types   = ["appliance"]
					}`, name2, vars.OrganizationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("meraki_network.test", "id"),
					resource.TestCheckResourceAttr("meraki_network.test", "organization_id", vars.OrganizationID),
					resource.TestCheckResourceAttr("meraki_network.test", "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
