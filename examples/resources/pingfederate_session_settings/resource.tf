resource "pingfederate_session_settings" "example" {
  revoke_user_session_on_logout     = true
  session_revocation_lifetime       = 60
  track_adapter_sessions_for_logout = true
}
