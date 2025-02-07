terraform {
  required_providers {
    hashicups = {
      source = "kenske.com/terraform/technitium-dns"
    }
  }
}

provider "technitium" {}

data "technitium_dhcp_scopes" "scopes" {}


output "scopes" {
  value = data.technitium_dhcp_scopes.scopes
}