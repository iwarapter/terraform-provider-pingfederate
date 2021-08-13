resource "pingfederate_oauth_authentication_policy_contract_mapping" "example" {
  authentication_policy_contract_ref {
    id = pingfederate_authentication_policy_contract.example.id
  }
  attribute_contract_fulfillment {
    key_name = "USER_NAME"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
  attribute_contract_fulfillment {
    key_name = "USER_KEY"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
  attribute_contract_fulfillment {
    key_name = "example"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
}

resource "pingfederate_authentication_policy_contract" "example" {
  name                = "example"
  extended_attributes = ["foo", "email"]
}
