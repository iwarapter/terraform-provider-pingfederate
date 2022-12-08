resource "pingfederate_oauth_client" "example" {
  client_id                        = "example"
  name                             = "example"
  grant_types                      = ["CLIENT_CREDENTIALS"]
  default_access_token_manager_ref = pingfederate_oauth_access_token_manager.example.id
  client_auth = {
    secret = "super_top_secret"
    type   = "SECRET"
  }
  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout", "https://foo"]
    ping_access_logout_capable          = true
  }
}
