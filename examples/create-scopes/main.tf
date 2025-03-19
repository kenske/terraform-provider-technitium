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
  name = "Test1"
  starting_address = "10.0.0.2"
  ending_address = "10.0.0.254"
  subnet_mask = "255.255.255.0"
  router_address = "10.0.0.1"
  domain_name = "kenskes"
  use_this_dns_server = false
  dns_servers = [
    "1.1.1.1",
    "8.8.8.8"
  ]
  exclusions = [
    {
      starting_address = "10.0.0.10"
      ending_address   = "10.0.0.40"
    },
    {
      starting_address = "10.0.0.240"
      ending_address   = "10.0.0.250"
    }
  ]
}



output "scopes" {
 value = data.technitium_dhcp_scopes.list.scopes
}

