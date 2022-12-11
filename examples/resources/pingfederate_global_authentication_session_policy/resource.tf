resource "pingfederate_global_authentication_session_policy" "example" {
  enable_sessions                = true
  hash_unique_user_key_attribute = true
  idle_timeout_display_unit      = "MINUTES"
  idle_timeout_mins              = 5
  max_timeout_display_unit       = "MINUTES"
  max_timeout_mins               = 6
  persistent_sessions            = true
}
