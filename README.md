# Terraform Provider - Technitium DNS Server

Works with Technitium v13.x [API](https://github.com/TechnitiumSoftware/DnsServer/blob/master/APIDOCS.md) 

Provider documentation is available at [registry.terraform.io/kenske/technitium](https://registry.terraform.io/providers/kenske/technitium/latest/docs).

## Getting Started

Add the provider to your Terraform configuration:

```hcl
terraform {
  required_providers {
    technitium = {
      version = "~> 0.0.7"
      source = "registry.terraform.io/kenske/technitium"
    }
  }
}

```

Then, configure the provider with the host and token:

```hcl
provider "technitium" {
  host  = var.TECHNITIUM_HOST
  token = var.TECHNITIUM_TOKEN
}
```

Alternatively, the values can be set using the `TECHNITIUM_HOST` and 
`TECHNITIUM_TOKEN` environment variables.

Complete example:

```hcl  
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

resource "technitium_dhcp_scope" "test" {
  name                = "Test1"
  starting_address    = "10.0.0.2"
  ending_address      = "10.0.0.254"
  subnet_mask         = "255.255.255.0"
  router_address      = "10.0.0.1"
  domain_name         = "local"
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


resource "technitium_dns_zone" "example" {
  name                       = "example.com"
  type                       = "Primary"
  use_soa_serial_date_scheme = true
}

resource "technitium_dns_zone_record" "cname" {
  zone     = technitium_dns_zone.example.name
  domain   = "cname.${technitium_dns_zone.example.name}"
  type     = "CNAME"
  cname    = "other.example.com"
  comments = "This is a test CNAME record"
  ttl      = "420"

  depends_on = [technitium_dns_zone.example]

}

resource "technitium_dns_zone_record" "a" {
  domain     = "testa.example3.com"
  type       = "A"
  ip_address = "1.1.1.1"
  comments   = "This is a test A record"
  ttl        = "420"

  depends_on = [technitium_dns_zone.example]

}


```

## Provider Development

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:
1. Setup the override in your `~/.terraformrc` file as described in the Terraform provider [documentation](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install).
1. Use `TF_LOG_PROVIDER=debug` to enable debug logging for the provider.

```bash

## To Do
- [ ] Add support for resource imports
- [ ] More/better unit tests

