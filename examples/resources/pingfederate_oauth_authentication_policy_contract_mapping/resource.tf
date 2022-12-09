resource "pingfederate_oauth_authentication_policy_contract_mapping" "example" {
  authentication_policy_contract_ref = pingfederate_authentication_policy_contract.example.id
  attribute_contract_fulfillment = {
    "USER_NAME" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    },
    "USER_KEY" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    },
    "example" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    }
  }
}

resource "pingfederate_authentication_policy_contract" "example" {
  name                = "example"
  extended_attributes = ["foo", "email"]
}
