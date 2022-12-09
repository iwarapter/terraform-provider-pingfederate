resource "pingfederate_oauth_access_token_manager" "example" {
  instance_id = "example"
  name        = "example"
  plugin_descriptor_ref {
    id = "org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin"
  }
  configuration {
    fields {
      name  = "Token Length"
      value = "28"
    }
    fields {
      name  = "Token Lifetime"
      value = "120"
    }
    fields {
      name  = "Lifetime Extension Policy"
      value = "ALL"
    }
    fields {
      name  = "Maximum Token Lifetime"
      value = ""
    }
    fields {
      name  = "Lifetime Extension Threshold Percentage"
      value = "30"
    }
    fields {
      name  = "Mode for Synchronous RPC"
      value = "3"
    }
    fields {
      name  = "RPC Timeout"
      value = "500"
    }
    fields {
      name  = "Expand Scope Groups"
      value = "false"
    }
  }

  attribute_contract {
    extended_attributes = ["sub"]
  }
}
