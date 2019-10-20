provider "pingfederate" {
  password = "2FederateM0re"
}

// resource "pingfederate_oauth_auth_server_settings" "settings" {
//   scopes {
//     name        = "address"
//     description = "address"
//   }

//   scopes {
//     name        = "mail"
//     description = "mail"
//   }

//   scopes {
//     name        = "openid"
//     description = "openid"
//   }

//   scopes {
//     name        = "phone"
//     description = "phone"
//   }

//   scopes {
//     name        = "profile"
//     description = "profile"
//   }

//   scope_groups {
//     name        = "group1"
//     description = "group1"

//     scopes = [
//       "address",
//       "mail",
//       "phone",
//       "openid",
//       "profile",
//     ]
//   }

//   persistent_grant_contract {
//     extended_attributes = ["woot"]
//   }

//   allowed_origins = [
//     "http://localhost",
//   ]

//   default_scope_description  = ""
//   authorization_code_timeout = 60
//   authorization_code_entropy = 30
//   refresh_token_length       = 42
//   refresh_rolling_interval   = 0
// }

// resource "pingfederate_oauth_client" "woot" {
//   client_id = "woot"
//   name      = "woot"

//   grant_types = [
//     "EXTENSION",
//   ]

//   client_auth {
//     // type                      = "CERTIFICATE"
//     // client_cert_issuer_dn     = ""
//     // client_cert_subject_dn    = ""
//     enforce_replay_prevention = false

//     secret = "super_top_secret"
//     type   = "SECRET"
//   }

//   // jwks_settings {
//   //   jwks = "https://stuff"
//   // }
//   default_access_token_manager_ref {
//     id = "atat"
//   }

//   oidc_policy {
//     grant_access_session_revocation_api = false

//     logout_uris = [
//       "https://logout",
//       "https://foo",
//     ]

//     ping_access_logout_capable = true
//   }
// }

resource "pingfederate_oauth_access_token_manager" "my_atm" {
  instance_id = "myatat"
  name        = "my_atat"

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

#resource "pingfederate_authentication_policy_contract" "apc_foo" {
#  name = "wee"
#  core_attributes = ["subject"]
#  extended_attributes = ["foo", "bar"]
#}

resource "pingfederate_password_credential_validator" "demo" {
  name = "foo"
  plugin_descriptor_ref {
    id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
  }

  configuration {
    tables {
      name = "Users"
      rows {
        fields {
          name  = "Username"
          value = "bob"
        }

        sensitive_fields {
          name  = "Password"
          value = "demo"
        }

        sensitive_fields {
          name  = "Confirm Password"
          value = "demo"
        }

        fields {
          name  = "Relax Password Requirements"
          value = "true"
        }
      }
    }
  }
  attribute_contract {
    core_attributes = ["username"]
  }
}