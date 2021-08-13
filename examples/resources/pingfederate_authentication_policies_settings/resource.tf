resource "pingfederate_authentication_policies_settings" "settings" {
  enable_idp_authn_selection = true
  enable_sp_authn_selection  = true
}
