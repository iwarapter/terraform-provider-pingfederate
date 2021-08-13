resource "pingfederate_oauth_client" "example" {
  client_id   = "example"
  name        = "example"
  grant_types = ["CLIENT_CREDENTIALS"]

  client_auth {
    secret = "super_top_secret"
    type   = "SECRET"
  }

  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }

  oidc_policy {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout", "https://foo"]
    ping_access_logout_capable          = true
  }
}
