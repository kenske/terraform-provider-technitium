resource "technitium_dns_zone" "example" {
  name              = "example.com"
  type              = "Primary"
}

resource "technitium_dns_zone" "example_ptr" {
  name              = "1.168.192.in-addr.arpa"
  type              = "Primary"
}