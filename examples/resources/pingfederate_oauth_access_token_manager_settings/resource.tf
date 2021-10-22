resource "pingfederate_oauth_access_token_manager_settings" "demo" {
  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }
}
