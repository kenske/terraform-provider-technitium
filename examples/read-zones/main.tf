terraform {
  required_providers {
    technitium = {
      source = "kenske.com/terraform/technitium-dns"
    }
  }
}

provider "technitium" {
  host  = var.TECHNITIUM_HOST
  token = var.TECHNITIUM_TOKEN
}

data "technitium_dns_zone" "example" {
  name = "example.com"
}

data "technitium_dns_zones" "list" {
}

data "technitium_dns_zone_records" "example" {
  domain = data.technitium_dns_zone.example.name
}




output "example_zone" {
  value = data.technitium_dns_zone.example
}

output "all_zones" {
  value = data.technitium_dns_zones.list.zones
}

output "zone_records" {
  value = data.technitium_dns_zone_records.example.records
}



