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