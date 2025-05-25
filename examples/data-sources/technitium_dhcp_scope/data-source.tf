data "technitium_dhcp_scope" "default" {
  name = "Defaults"
}

output "default_scope" {
  value = data.technitium_dhcp_scope.default
}