terraform {
  required_providers {
    technitium = {
      source = "registry.terraform.io/kenske/technitium"
    }
  }
}

provider "technitium" {
  host  = "http://localhost:5380"
  username = "admin"
  password = "password"
}

data "technitium_dns_zone" "example" {
  name = "example3.com"
}

data "technitium_dns_zones" "list" {
}

data "technitium_dns_zone_records" "example" {
  domain = data.technitium_dns_zone.example.name
}


data "technitium_dns_zone_record" "example" {
  name = "testa.example3.com"
  type =  "A"
}



output "all_records" {
  value = data.technitium_dns_zone_records.example
}

output "single_record" {
  value = data.technitium_dns_zone_record.example
}
