resource "pingfederate_application_session_policy" "example" {
  idle_timeout_mins = 10
  max_timeout_mins  = 12
}
