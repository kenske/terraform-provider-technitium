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

resource "technitium_dns_zone" "example" {
  name                       = "example3.com"
  type                       = "Primary"
  use_soa_serial_date_scheme = true
}

resource "technitium_dns_zone_record" "example" {
  domain   = "test.example3.com"
  type     = "CNAME"
  cname    = "other.example.com"
  comments = "This is a test record"
  ttl      = "420"

  depends_on = [technitium_dns_zone.example]

}
