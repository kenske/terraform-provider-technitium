resource "technitium_dns_zone" "example" {
  name                       = "example3.com"
  type                       = "Primary"
  use_soa_serial_date_scheme = true
}

resource "technitium_dns_zone_record" "cname" {
  zone     = technitium_dns_zone.example.name
  domain   = "cname.${technitium_dns_zone.example.name}"
  type     = "CNAME"
  cname    = "other.example3.com"
  comments = "This is a test CNAME record"
  ttl      = "420"

  depends_on = [technitium_dns_zone.example]

}

resource "technitium_dns_zone_record" "a" {
  domain     = "testa.example3.com"
  type       = "A"
  ip_address = "1.1.1.1"
  comments   = "This is a test A record"
  ttl        = "420"

  depends_on = [technitium_dns_zone.example]

}