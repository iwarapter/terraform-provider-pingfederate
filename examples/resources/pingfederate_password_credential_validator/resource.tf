resource "pingfederate_password_credential_validator" "example" {
  name = "example"
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
          value = "demo2"
        }

        sensitive_fields {
          name  = "Confirm Password"
          value = "demo2"
        }

        fields {
          name  = "Relax Password Requirements"
          value = "true"
        }
      }
    }
  }
}
