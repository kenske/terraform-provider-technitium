terraform {
  required_providers {
    technitium = {
      source = "kenske.com/terraform/technitium-dns"
    }
  }
}

provider "technitium" {
  host  = "http://localhost:19090"
  token = "test"
}

