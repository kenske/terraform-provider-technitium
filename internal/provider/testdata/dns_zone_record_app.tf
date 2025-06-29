resource "technitium_dns_zone" "example" {
  name = "example.com"
  type = "Primary"
}


resource "technitium_dns_zone_record" "app" {
  zone        = technitium_dns_zone.example.name
  domain      = "test.${technitium_dns_zone.example.name}"
  type        = "APP"
  app_name    = "NO DATA"
  class_path  = "NoData.App"
  record_data = <<EOT
  {
    "blockedTypes": [
      "HTTPS"
    ]
  }
  EOT
  depends_on  = [technitium_dns_zone.example]
}
