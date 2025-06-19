terraform {
  required_providers {
    technitium = {
      source = "registry.terraform.io/kenske/technitium"
    }
  }
}


# Using API token
provider "technitium" {
  host  = "http://localhost:5380"
  token = var.TECHNITIUM_TOKEN
}

# Using API username/password
provider "technitium" {
  host     = "http://localhost:5380"
  username = "admin"
  password = "password"
}