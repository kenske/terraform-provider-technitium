resource "technitium_dns_zone" "example" {
  name                       = "example.com"
  type                       = "Primary"
  use_soa_serial_date_scheme = true
}

resource "technitium_dns_zone" "internal" {
  name      = "internal-example.com"
  type      = "Forwarder"
  forwarder = "this-server"
}

resource "technitium_dns_zone" "ptr" {
  name = "1.168.192.in-addr.arpa"
  type = "Primary"
}