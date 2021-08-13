resource "pingfederate_authentication_api_settings" "settings" {
  api_enabled             = true
  enable_api_descriptions = false
  default_application_ref {
    id = pingfederate_authentication_api_application.example.id
  }
}
