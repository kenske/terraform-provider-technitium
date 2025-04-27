terraform {
  required_providers {
    technitium = {
      source = "registry.terraform.io/kenske/technitium"
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
  name                = "Test1"
  starting_address    = "10.0.0.2"
  ending_address      = "10.0.0.254"
  subnet_mask         = "255.255.255.0"
  router_address      = "10.0.0.1"
  domain_name         = "kenskes"
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

resource "technitium_dhcp_reserved_lease" "test" {
  name             = technitium_dhcp_scope.test.name
  hardware_address = "00:11:22:33:44:56"
  ip_address       = "10.0.0.9"
  host_name        = "test-host"
  comments         = "This is a test reserved lease"
}



output "scopes" {
  value = data.technitium_dhcp_scopes.list.scopes
}

