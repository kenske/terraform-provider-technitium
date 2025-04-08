terraform {
  required_providers {
    technitium = {
      source = "kenske.com/terraform/technitium-dns"
    }
  }
}

provider "technitium" {
  host  = var.TECHNITIUM_HOST
  token = var.TECHNITIUM_TOKEN
}

# data "technitium_dns_zone" "localhost" {
#   name = "localhost"
# }

data "technitium_dns_zones" "list" {
}


# output "default_zone" {
#   value = data.technitium_dns_zone.localhost
# }

output "all_zones" {
  value = data.technitium_dns_zones.list.zones
}

