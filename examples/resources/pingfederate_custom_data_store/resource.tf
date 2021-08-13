resource "pingfederate_custom_data_store" "example" {
  name = "example"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.datastore.other.RestDataSourceDriver"
  }
  configuration {
    tables {
      name = "Base URLs and Tags"
      rows {
        fields {
          name  = "Base URL"
          value = "https://example.com"
        }
        fields {
          name  = "Tags"
          value = ""
        }
        default_row = true
      }
    }
    tables {
      name = "HTTP Request Headers"
    }
    tables {
      name = "Attributes"
      rows {
        fields {
          name  = "Local Attribute"
          value = "foo"
        }
        fields {
          name  = "JSON Response Attribute Path"
          value = "/bar"
        }
      }
    }
    fields {
      name  = "Authentication Method"
      value = "None"
    }
    fields {
      name  = "Username"
      value = ""
    }
    fields {
      name  = "Password"
      value = ""
    }
    fields {
      name  = "OAuth Token Endpoint"
      value = ""
    }
    fields {
      name  = "OAuth Scope"
      value = ""
    }
    fields {
      name  = "Client ID"
      value = ""
    }
    fields {
      name  = "Client Secret"
      value = ""
    }
    fields {
      name  = "Enable HTTPS Hostname Verification"
      value = "true"
    }
    fields {
      name  = "Read Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Connection Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Max Payload Size (KB)"
      value = "1024"
    }
    fields {
      name  = "Retry Request"
      value = "true"
    }
    fields {
      name  = "Maximum Retries Limit"
      value = "5"
    }
    fields {
      name  = "Retry Error Codes"
      value = "429"
    }
    fields {
      name = "Test Connection URL"
    }
  }
}
