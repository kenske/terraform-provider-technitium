terraform {
  required_providers {
    technitium = {
      source = "registry.terraform.io/kenske/technitium"
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




