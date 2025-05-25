data "technitium_dhcp_scopes" "list" {}

output "scope_list" {
  value = data.technitium_dhcp_scopes.list.scopes
}