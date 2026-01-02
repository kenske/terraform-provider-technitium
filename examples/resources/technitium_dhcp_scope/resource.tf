resource "technitium_dhcp_scope" "test" {
  name             = "Test1"
  starting_address = "10.0.0.2"
  ending_address   = "10.0.0.254"
  subnet_mask      = "255.255.255.0"

  # Lease time configuration
  lease_time_days    = 7
  lease_time_hours   = 0
  lease_time_minutes = 0

  # Ping check to ensure IP is available
  ping_check_enabled = true
  ping_check_timeout = 1000
  ping_check_retries = 2

  # Network configuration
  router_address      = "10.0.0.1"
  use_this_dns_server = false

  dns_servers = [
    "1.1.1.1",
    "8.8.8.8"
  ]

  ntp_servers = [
    "time.nist.gov",
    "us.pool.ntp.org",
  ]

  # Domain configuration
  domain_name = "example.home"
  domain_search_list = [
    "example.home",
    "internal.example.home",
    "example.local",
  ]

  # Static routes for clients
  static_routes = [
    {
      destination = "192.168.100.0"
      subnet_mask = "255.255.255.0"
      router      = "10.0.0.1"
    }
  ]

  # Vendor specific information
  vendor_info = [
    {
      identifier  = "MSFT 5.0"
      information = "010203"
    }
  ]

  # TFTP server for network booting
  tftp_server_addresses = ["10.0.0.5"]
  boot_file_name        = "pxelinux.0"
  next_server_address   = "10.0.0.5"

  # Generic DHCP options (example: option 252 for WPAD)
  generic_options = [
    {
      code  = 252
      value = "687474703a2f2f7770616400" # hex encoded "http://wpad"
    }
  ]

  # IP address exclusions
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

  # Access control
  allow_only_reserved_leases               = false
  block_locally_administered_mac_addresses = false
  ignore_client_identifier_option          = false
}
