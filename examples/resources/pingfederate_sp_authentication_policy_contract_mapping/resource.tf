resource "pingfederate_sp_authentication_policy_contract_mapping" "example" {
  source_id = pingfederate_authentication_policy_contract.example.id
  target_id = pingfederate_sp_adapter.example.id
  attribute_contract_fulfillment {
    key_name = "subject"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
  default_target_resource = "https://example.com"
}
