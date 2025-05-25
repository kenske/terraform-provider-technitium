data "technitium_dns_zone" "example" {
  name = "example.com"
}

data "technitium_dns_zone_records" "example" {
  domain = data.technitium_dns_zone.example.name
}




