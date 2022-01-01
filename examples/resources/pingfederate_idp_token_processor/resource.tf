resource "pingfederate_idp_token_processor" "example" {
  processor_id = "example"
  name         = "example"
  plugin_descriptor_ref {
    id = "org.sourceid.wstrust.processor.jwt.JWTTokenProcessor"
  }
  configuration {
    fields {
      name  = "JWKS Endpoint URI"
      value = "https://example.com/jwks"
    }
    fields {
      name  = "Issuer"
      value = "example"
    }
    fields {
      name  = "Expiry Tolerance"
      value = "0"
    }
  }
}
