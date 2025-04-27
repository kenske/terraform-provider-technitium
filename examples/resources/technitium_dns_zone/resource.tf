resource "technitium_dns_zone" "example" {
  name                       = "example3.com"
  type                       = "Primary"
  use_soa_serial_date_scheme = true
}