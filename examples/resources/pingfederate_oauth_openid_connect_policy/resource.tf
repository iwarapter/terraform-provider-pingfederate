resource "pingfederate_oauth_openid_connect_policy" "example" {
  policy_id = "example"
  name      = "example"
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }
  attribute_contract {
    extended_attributes {
      name                = "name"
      include_in_id_token = true
    }
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "NO_MAPPING"
      }
    }
    attribute_contract_fulfillment {
      key_name = "name"
      source {
        type = "NO_MAPPING"
      }
    }
  }

  scope_attribute_mappings {
    key_name = "address"
    values   = ["name"]
  }
}
