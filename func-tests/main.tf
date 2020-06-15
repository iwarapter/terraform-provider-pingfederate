provider "pingfederate" {
  password = "2Federate"
}
//
resource "pingfederate_oauth_auth_server_settings" "settings" {
  scopes {
    name        = "address"
    description = "address"
  }

  scopes {
    name        = "mail"
    description = "mail"
  }

  scopes {
    name        = "openid"
    description = "openid"
  }

  scopes {
    name        = "phone"
    description = "phone"
  }

  scopes {
    name        = "profile"
    description = "profile"
  }

  scope_groups {
    name        = "group1"
    description = "group1"

    scopes = [
      "address",
      "mail",
      "phone",
      "openid",
      "profile",
    ]
  }

  persistent_grant_contract {
    extended_attributes = ["woot"]
  }

  allowed_origins = [
    "http://localhost",
  ]

  default_scope_description  = ""
  authorization_code_timeout = 60
  authorization_code_entropy = 30
  refresh_token_length       = 42
  refresh_rolling_interval   = 0
}

resource "pingfederate_oauth_client" "woot" {
  client_id = "woot"
  name      = "woot"

  grant_types = [
    "EXTENSION",
  ]

  client_auth {
    // type                      = "CERTIFICATE"
    // client_cert_issuer_dn     = ""
    // client_cert_subject_dn    = ""
    enforce_replay_prevention = false

    secret = "super_top_secret"
    type   = "SECRET"
  }

  // jwks_settings {
  //   jwks = "https://stuff"
  // }
  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.my_atm.id
  }

  oidc_policy {
    grant_access_session_revocation_api = false

    logout_uris = [
      "https://logout",
      "https://foo",
    ]

    ping_access_logout_capable = true
  }
}

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

//#resource "pingfederate_authentication_policy_contract" "apc_foo" {
//#  name = "wee"
//#  core_attributes = ["subject"]
//#  extended_attributes = ["foo", "bar"]
//#}
//
//resource "pingfederate_password_credential_validator" "demo" {
//  name = "foo"
//  plugin_descriptor_ref {
//    id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
//  }
//
//  configuration {
//    tables {
//      name = "Users"
//      rows {
//        fields {
//          name  = "Username"
//          value = "bob"
//        }
//
//        sensitive_fields {
//          name  = "Password"
//          value = "demo2"
//        }
//
//        sensitive_fields {
//          name  = "Confirm Password"
//          value = "demo2"
//        }
//
//        fields {
//          name  = "Relax Password Requirements"
//          value = "true"
//        }
//      }
//    }
//  }
//  attribute_contract {
//    core_attributes = ["username"]
//  }
//}
//
//resource "pingfederate_data_store" "demo" {
//  jdbc_data_store {
//    name         = "terraform"
//    driver_class = "org.hsqldb.jdbcDriver"
//    user_name    = "sa"
//    //    password = ""
//    connection_url = "jdbc:hsqldb:mem:mymemdb"
//
//    connection_url_tags {
//      connection_url = "jdbc:hsqldb:mem:mymemdb"
//      default_source = true
//    }
//  }
//}
////
//resource "pingfederate_data_store" "demo_ldap" {
//  bypass_external_validation = true
//  ldap_data_store {
//    name      = "terraform_ldap"
//    ldap_type = "PING_DIRECTORY"
//    hostnames = [
//    "host.docker.internal:1389"]
//    bind_anonymously = true
//
//
//    min_connections = 1
//    max_connections = 1
//  }
//}
//
//resource "pingfederate_idp_adapter" "demo" {
//  name = "bart"
//  plugin_descriptor_ref {
//    id = "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"
//  }
//
//  configuration {
//    tables {
//      name = "Credential Validators"
//      rows {
//        fields {
//          name  = "Password Credential Validator Instance"
//          value = pingfederate_password_credential_validator.demo.name
//        }
//      }
//    }
//    fields {
//      name  = "Realm"
//      value = "foo"
//    }
//
//    fields {
//      name  = "Challenge Retries"
//      value = "3"
//    }
//
//  }
//
//  attribute_contract {
//    core_attributes {
//      name      = "username"
//      pseudonym = true
//    }
//    extended_attributes {
//      name = "sub"
//    }
//  }
//  attribute_mapping {
//    attribute_contract_fulfillment {
//      key_name = "sub"
//      source {
//        type = "ADAPTER"
//      }
//      value = "sub"
//    }
//    attribute_contract_fulfillment {
//      key_name = "username"
//      source {
//        type = "ADAPTER"
//      }
//      value = "username"
//    }
//    jdbc_attribute_source {
//      filter      = "\"\""
//      description = "foo"
//      schema      = "INFORMATION_SCHEMA"
//      table       = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
//      data_store_ref {
//        id = "ProvisionerDS"
//      }
//    }
//  }
//}

resource "pingfederate_sp_adapter" "demo" {
  name          = "bar"
  sp_adapter_id = "bar"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.opentoken.SpAuthnAdapter"
  }

  configuration {

    sensitive_fields {
      name  = "Password"
      value = "Secret123"
    }
    sensitive_fields {
      name  = "Confirm Password"
      value = "Secret123"
    }
    fields {
      name  = "Account Link Service"
      value = ""
    }
    fields {
      name  = "Authentication Service"
      value = ""
    }
    fields {
      name  = "Cipher Suite"
      value = "2"
    }
    fields {
      name  = "Cookie Domain"
      value = ""
    }
    fields {
      name  = "Cookie Path"
      value = "/"
    }
    fields {
      name  = "Force SunJCE Provider"
      value = "false"
    }
    fields {
      name  = "HTTP Only Flag"
      value = "true"
    }
    fields {
      name  = "Logout Service"
      value = ""
    }
    fields {
      name  = "Not Before Tolerance"
      value = "0"
    }
    fields {
      name  = "Obfuscate Password"
      value = "true"
    }
    fields {
      name  = "Secure Cookie"
      value = "false"
    }
    fields {
      name  = "Send Extended Attributes"
      value = ""
    }
    fields {
      name  = "Send Subject as Query Parameter"
      value = "false"
    }
    fields {
      name  = "Session Cookie"
      value = "false"
    }
    fields {
      name  = "Session Lifetime"
      value = "43200"
    }
    fields {
      name  = "Skip Trimming of Trailing Backslashes"
      value = "false"
    }
    fields {
      name  = "Subject Query Parameter                 "
      value = ""
    }
    fields {
      name  = "Token Lifetime"
      value = "300"
    }
    fields {
      name  = "Token Name"
      value = "opentoken"
    }
    fields {
      name  = "Transport Mode"
      value = "2"
    }
    fields {
      name  = "URL Encode Cookie Values"
      value = "true"
    }
    fields {
      name  = "Use Verbose Error Messages"
      value = "false"
    }


  }

  attribute_contract {
    core_attributes = ["subject"]
  }

  target_application_info {
    application_name     = "foo"
    application_icon_url = "https://foo"
  }
}

resource "pingfederate_oauth_access_token_mappings" "demo" {
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.my_atm.id
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

resource "pingfederate_oauth_openid_connect_policy" "demo" {
  policy_id = "foo"
  name      = "foo"
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.my_atm.id
  }
  attribute_contract {
    core_attributes {
      name = "sub"
    }
    extended_attributes {
      name                 = "email"
      include_in_user_info = true
    }
    extended_attributes {
      name                 = "email_verified"
      include_in_user_info = true
    }
    extended_attributes {
      name                 = "family_name"
      include_in_user_info = true
    }
    extended_attributes {
      name                 = "name"
      include_in_user_info = true
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
      key_name = "email"
      source {
        type = "NO_MAPPING"
      }
    }
    attribute_contract_fulfillment {
      key_name = "email_verified"
      source {
        type = "NO_MAPPING"
      }
    }
    attribute_contract_fulfillment {
      key_name = "family_name"
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

//  scope_attribute_mappings = { //TODO hoping the new TF 2.0.0 SDK will finally support sensible maps
//    address = ["foo", "bar"]
//  }
}