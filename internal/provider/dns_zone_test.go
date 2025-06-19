package provider

import (
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"os"
	"testing"
)

func TestDnsZones(t *testing.T) {

	configBytes, err := os.ReadFile("test/dns_zone.tf")
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}
	config := string(configBytes)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.TestCheckResourceAttr(
					"technitium_dns_zone.test", "name", "example.com",
				),
			},
		},
	})
}
