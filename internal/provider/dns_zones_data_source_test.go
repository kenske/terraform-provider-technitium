package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"net/http"
	"terraform-provider-technitium/internal/test"
	"testing"
	"time"
)

func TestDnsZonesDataSource(t *testing.T) {

	scenario := test.GetMockScenarioFromFile(t, "../test/mocks/dns_zones_response.json", http.StatusOK)
	server := test.NewTestServer(scenario)
	defer server.Close()

	config := fmt.Sprintf(`
provider "technitium" {
  host = "%s"
  token = "test"
}
`, server.URL)

	resource.Test(t, resource.TestCase{
		IsUnitTest:               true,
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config + `data "technitium_dns_zones" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.name", "0.in-addr.arpa"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.type", "Primary"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.catalog", ""),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.disabled", "false"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.dnssec_status", "Unsigned"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.soa_serial", "1"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.expiry", ""),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.is_expired", "false"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.0.internal", "true"),
					resource.TestCheckResourceAttr("data.technitium_dns_zones.test", "zones.#", "6"),

					resource.TestCheckResourceAttrWith(
						"data.technitium_dns_zones.test",
						"zones.0.last_modified",
						func(value string) error {
							_, err := time.Parse(time.RFC3339Nano, value)
							if err != nil {
								return fmt.Errorf("last_modified is not a valid date: %s", value)
							}
							return nil
						},
					),
				),
			},
		},
	})
}
