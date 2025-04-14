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

data "technitium_dns_zones" "list" {
}

resource "technitium_dns_zone" "example" {
  name = "example3.com"
  type = "Primary"
  # use_soa_serial_date_scheme = true
}
