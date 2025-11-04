resource "technitium_dns_zone" "example" {
  name              = "example.com"
  type              = "Primary"
}

resource "technitium_dns_zone" "example_ptr" {
  name              = "1.168.192.in-addr.arpa"
  type              = "Primary"
}

resource "technitium_dns_zone_record" "ptr" {
  domain   = "1.${technitium_dns_zone.example_ptr.name}"
  ptr_name = "ptr.${technitium_dns_zone.example.name}"
  type     = "PTR"
  comments = "test-ptr"
  ttl      = "3600"
}