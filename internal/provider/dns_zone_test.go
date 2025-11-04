package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestDnsZones(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: GetFileConfig(t, "dns_zone.tf"),
				Check: resource.TestCheckResourceAttr(
					"technitium_dns_zone.test", "name", "example.com",
				),
			},
		},
	})
}
