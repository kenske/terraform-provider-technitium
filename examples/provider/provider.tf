terraform {
  required_providers {
    technitium = {
      source = "registry.terraform.io/kenske/technitium"
    }
  }
}


provider "technitium" {
  host  = "http://localhost:5380"
  token = var.TECHNITIUM_TOKEN
}

