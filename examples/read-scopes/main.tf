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

data "technitium_dhcp_scopes" "list" {}

data "technitium_dhcp_scope" "default" {
  name = "Default"
}


output "scope_list" {
 value = data.technitium_dhcp_scopes.list.scopes
}


output "default_scope" {
  value = data.technitium_dhcp_scope.default
}

