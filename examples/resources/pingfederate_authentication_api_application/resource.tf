resource "pingfederate_authentication_api_application" "example" {
  name                       = "myapp"
  url                        = "https://example.com"
  description                = "This is an auth api app"
  additional_allowed_origins = ["https://foo.bar", "https://bar.foo"]
}
