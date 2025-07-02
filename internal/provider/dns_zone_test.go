package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"terraform-provider-technitium/internal/technitium"
	"testing"
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

func TestAccDnsZoneRecord_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create initial resource
			{
				Config: GetFileConfig(t, "dns_zone_record.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_dns_zone_record.test", "ip_address", "192.168.1.10"),
					resource.TestCheckResourceAttr("technitium_dns_zone_record.test", "ttl", "3600"),
				),
			},
			// Update the resource with a different TTL and IP
			{
				Config: GetFileConfig(t, "dns_zone_record_updated.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						// This checks that the resource is updated in-place rather than replaced
						plancheck.ExpectResourceAction("technitium_dns_zone_record.test", plancheck.ResourceActionUpdate),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_dns_zone_record.test", "ip_address", "192.168.1.20"),
					resource.TestCheckResourceAttr("technitium_dns_zone_record.test", "ttl", "7200"),
				),
			},
			{
				Config: GetFileConfig(t, "dns_zone_record_updated.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func TestAccDnsZoneRecord_App(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create initial resource
			{
				Config: GetFileConfig(t, "dns_zone_record_app.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_dns_zone_record.app", "app_name", "NO DATA"),
					resource.TestCheckResourceAttr("technitium_dns_zone_record.app", "ttl", "0"),
				),
			},
			{
				// Make sure the TTL got updated
				Config: GetFileConfig(t, "dns_zone_record_app.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_dns_zone_record.app", "ttl", "0"),
				),
			},
		},
	})
}

func TestAccDnsZoneRecord_drift(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create initial resource
			{
				Config: GetFileConfig(t, "dns_zone_record.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("technitium_dns_zone_record.test", "comments", "test"),
				),
			},
			{
				// Simulate drift by updating the record outside Terraform
				PreConfig: func() {
					err := updateZoneRecord()
					if err != nil {
						t.Fatalf("Error updating DNS zone record: %v", err)
					}
				},
				Config: GetFileConfig(t, "dns_zone_record.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction("technitium_dns_zone_record.test", plancheck.ResourceActionUpdate),
						ExpectDrift("technitium_dns_zone_record.test", "comments", "external update"),
					},
				},
			},
			{
				Config: GetFileConfig(t, "dns_zone_record.tf"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
			},
		},
	})
}

func updateZoneRecord() error {
	host := "http://localhost:5380"
	username := "admin"
	password := "password"

	token, err := technitium.GetToken(host, username, password)

	if err != nil {
		return fmt.Errorf("Error getting token: %v\n", err)
	}

	ctx := context.Background()
	client, err := technitium.NewClient(host, token, ctx)

	if err != nil {
		return fmt.Errorf("Error creating client: %v\n", err)
	}

	update := technitium.DnsZoneRecordUpdate{
		DnsZoneRecordCreate: technitium.DnsZoneRecordCreate{
			Domain:    "test.example.com",
			Type:      "A",
			Zone:      "example.com",
			IPAddress: "192.168.1.10",
			Comments:  "external update",
		},
	}

	err = client.UpdateDnsZoneRecord(update, ctx)
	if err != nil {
		return fmt.Errorf("Error updating DNS zone record: %v\n", err)
	}

	return nil

}
