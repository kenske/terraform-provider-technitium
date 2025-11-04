terraform {
  required_providers {
    technitium = {
      source = "kenske/technitium"
    }
  }
}

provider "technitium" {
  host  = "http://localhost:5380"
  username = "admin"
    password = "password"
}

resource "technitium_dns_zone" "example" {
  name              = "example3.com"
  type              = "Forwarder"
  forwarder         = "this-server"
  protocol          = "Https"
  dnssec_validation = true
}

resource "technitium_dns_zone_record" "cname"  {
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
  comments    = "This is a test A record"
  ttl        = "420"

  depends_on = [technitium_dns_zone.example]

}

resource "technitium_dns_zone_record" "app" {
  zone     = technitium_dns_zone.example.name
  domain     = "test.${technitium_dns_zone.example.name}"
  type       = "APP"
  app_name   = "NO DATA"
  class_path  = "NoData.App"
  record_data = <<EOT
  {
    "blockedTypes": [
      "HTTPS"
    ]
  }
  EOT
  depends_on = [technitium_dns_zone.example]
}


# PTR
resource "technitium_dns_zone" "example_ptr" {
  name              = "1.168.192.in-addr.arpa"
  type              = "Primary"
}

resource "technitium_dns_zone_record" "ptr" {
  domain   = "1.${technitium_dns_zone.example_ptr.name}"
  ptr_name = "ptr.${technitium_dns_zone.example.name}"
  type     = "PTR"
  comments = "test"
  ttl      = "3600"
}