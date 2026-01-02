resource "technitium_dhcp_scope" "test" {
  name                = "Test1"
  starting_address    = "10.0.0.2"
  ending_address      = "10.0.0.254"
  subnet_mask         = "255.255.255.0"
  router_address      = "10.0.0.1"
  domain_name         = "example"
  use_this_dns_server = false
  dns_servers = [
    "1.1.1.1",
    "8.8.8.8"
  ]
  ntp_servers = [
    "time.nist.gov",
    "us.pool.ntp.org"
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
