resource "pingfederate_keypairs_oauth_openid_connect" "example" {
  static_jwks_enabled = true
  rsa_active_cert_ref {
    id = pingfederate_keypair_signing.new.id
  }
  rsa_previous_cert_ref {
    id = pingfederate_keypair_signing.old.id
  }
  rsa_publish_x5c_parameter = true
}
