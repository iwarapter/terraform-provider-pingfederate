provider "pingfederate" {}

resource "pingfederate_oauth_auth_server_settings" "settings" {
  default_scope_description  = ""
  authorization_code_timeout = 60
  authorization_code_entropy = 30
  refresh_token_length       = 42
  refresh_rolling_interval   = 0
}
