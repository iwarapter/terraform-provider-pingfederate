resource "pingfederate_oauth_access_token_mappings" "example" {
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }

  context {
    type = "CLIENT_CREDENTIALS"
  }
  attribute_contract_fulfillment {
    key_name = "sub"
    source {
      type = "CONTEXT"
    }
    value = "ClientId"
  }
}
