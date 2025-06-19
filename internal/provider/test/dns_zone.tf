provider "technitium" {
  host = "http://localhost:5380"
  username = "admin"
  password = "password"
}

resource "technitium_dns_zone" "test" {
  name = "example.com"
  type = "Primary"
}
