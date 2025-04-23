package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"
	"time"
)

func TestDnsZonesDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `data "technitium_dns_zones" "test" {}`,
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
