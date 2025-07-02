resource "technitium_dns_zone" "test" {
  name = "example.com"
  type = "Primary"
}

resource "technitium_dns_zone_record" "test" {
  domain    = "test.example.com"
  zone      = technitium_dns_zone.test.name
  type      = "A"
  ip_address = "192.168.1.10"
  comments = "test"
  ttl       = 3600
}