resource "pingfederate_authentication_session_policy" "example" {
  id = "acc_test1"
  authentication_source = {
    type       = "IDP_ADAPTER"
    source_ref = pingfederate_idp_adapter.example.id
  }
  enable_sessions         = true
  persistent              = true
  idle_timeout_mins       = 300
  max_timeout_mins        = 300
  timeout_display_unit    = "HOURS"
  authn_context_sensitive = true
}
