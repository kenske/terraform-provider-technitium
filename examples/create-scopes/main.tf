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

data "technitium_dhcp_scopes" "scopes" {}

resource "technitium_dhcp_scope" "test" {
  name = "Test"
  broadcast_address = "10.0.0.255"
  starting_address = "10.0.0.100"
  ending_address = "10.0.0.150"
  subnet_mask = "255.255.255.0"
}


output "scopes" {
 value = data.technitium_dhcp_scopes.scopes.scopes
}

