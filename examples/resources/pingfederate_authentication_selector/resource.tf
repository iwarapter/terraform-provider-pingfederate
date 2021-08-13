resource "pingfederate_authentication_selector" "example" {
  name = "example"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"
  }

  configuration {
    fields {
      name  = "Result Attribute Name"
      value = ""
    }
    tables {
      name = "Networks"
      rows {
        fields {
          name  = "Network Range (CIDR notation)"
          value = "127.0.0.1/32"
        }
      }
    }
  }
}
