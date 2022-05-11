package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetworkWebhookHttpserverResource(t *testing.T) {
	networkId := os.Getenv("MERAKI_NETWORK_ID")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccNetworkWebhookHttpserverResourceConfig("one", networkId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "name", "one"),
					resource.TestCheckResourceAttr("meraki_network_webhook_httpserver.test", "network_id", networkId),
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
				Config: testAccNetworkWebhookHttpserverResourceConfig("two", networkId),
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
  name = %[1]q
  network_id = %[1]q
}
`, name, networkID)
}
