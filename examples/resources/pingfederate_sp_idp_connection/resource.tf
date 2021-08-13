resource "pingfederate_sp_idp_connection" "example" {
  name         = "example"
  entity_id    = "example"
  active       = true
  logging_mode = "STANDARD"
  credentials {
    outbound_back_channel_auth {
      type                  = "OUTBOUND"
      digital_signature     = false
      validate_partner_cert = false
    }
  }
  attribute_query {
    url = "https://example.com"
    policy {
      sign_attribute_query        = false
      encrypt_name_id             = false
      require_signed_response     = false
      require_signed_assertion    = false
      require_encrypted_assertion = false
      mask_attribute_values       = false
    }
  }
}
