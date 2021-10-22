resource "pingfederate_oauth_client_settings" "settings" {
  client_metadata {
    parameter = "example1"
  }
  client_metadata {
    parameter   = "example2"
    description = "example2"
  }
  dynamic_client_registration {
    restrict_common_scopes              = false
    initial_access_token_scope          = "openid"
    enforce_replay_prevention           = true
    require_proof_key_for_code_exchange = true
  }
}
