resource "pingfederate_oauth_resource_owner_credentials_mappings" "example" {
  password_validator_ref {
    id = pingfederate_password_credential_validator.example.id
  }

  attribute_contract_fulfillment {
    key_name = "USER_KEY"
    source {
      type = "NO_MAPPING"
    }
  }

  attribute_contract_fulfillment {
    key_name = "example"
    source {
      type = "NO_MAPPING"
    }
  }
  issuance_criteria {
    conditional_criteria {
      attribute_name = "username"
      condition      = "EQUALS"
      error_result   = "deny"
      value          = "example"

      source {
        type = "PASSWORD_CREDENTIAL_VALIDATOR"
      }
    }
  }
}
