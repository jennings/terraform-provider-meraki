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
				Config: testAccNetworkWebhookHttpserverResourceConfig("one", vars.NetworkID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "name", "one"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "network_id", vars.NetworkID),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "id", "example-id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "meraki_network_webhook_httpserver.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccNetworkWebhookHttpserverResourceConfig("two", vars.NetworkID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "name", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccNetworkWebhookHttpserverResourceConfig(name, networkID string) string {
	return fmt.Sprintf(`
resource "meraki_network_webhook_httpserver" "test" {
  name       = %q
  network_id = %q
  url        = "https://example.com/%[1]s"
}
`, name, networkID)
}
