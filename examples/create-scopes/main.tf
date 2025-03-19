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

data "technitium_dhcp_scopes" "list" {
  depends_on = [technitium_dhcp_scope.test]
}

resource "technitium_dhcp_scope" "test" {
  name = "Test"
  starting_address = "10.1.0.100"
  ending_address = "10.1.0.157"
  subnet_mask = "255.255.255.0"
}


output "scopes" {
 value = data.technitium_dhcp_scopes.list.scopes
}

