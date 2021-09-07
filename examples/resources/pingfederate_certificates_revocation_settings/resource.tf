resource "pingfederate_certificates_revocation_settings" "settings" {
  # enabled and default crl configuration
  crl_settings {}
  # ocsp responder with crl failover
  ocsp_settings {
    responder_url            = "http://my.ocspresponder.com"
    action_on_status_unknown = "FAILOVER"
  }
}
