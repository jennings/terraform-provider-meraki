package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkWebhookHttpserverResource(t *testing.T) {
	vars := readTestVars(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: fmt.Sprintf(`
					resource "meraki_network_webhook_httpserver" "test" {
						name          = "one"
						network_id    = %q
						url           = "https://example.com/test"
						shared_secret = "one"
					}`, vars.NetworkID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("meraki_network_webhook_httpserver.test", "id"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "network_id", vars.NetworkID),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "name", "one"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "shared_secret", "one"),
				),
			},
			// // ImportState testing
			// {
			// 	ResourceName:      "meraki_network_webhook_httpserver.test",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// 	ImportStateIdFunc: func(s *terraform.State) (string, error) {
			// 		r := s.Children("root").Resources["meraki_network_webhook_httpserver"]
			// 		return "", nil
			// 	},
			// 	ImportStateVerifyIgnore: []string{"shared_secret"},
			// },
			// Update and Read testing
			{
				Config: fmt.Sprintf(`
					resource "meraki_network_webhook_httpserver" "test" {
						name          = "two"
						network_id    = %q
						url           = "https://example.com/test"
						shared_secret = "two"
					}`, vars.NetworkID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("meraki_network_webhook_httpserver.test", "id"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "network_id", vars.NetworkID),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "name", "two"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "shared_secret", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
