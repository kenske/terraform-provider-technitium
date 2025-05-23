resource "technitium_dhcp_reserved_lease" "test" {
  name             = "example"
  hardware_address = "00:11:22:33:44:56"
  ip_address       = "10.0.0.9"
  host_name        = "test-host"
  comments         = "This is a test reserved lease"
}
