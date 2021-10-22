resource "pingfederate_oauth_client_registration_policy" "demo" {
  policy_id = "mypolicy"
  name      = "mypolicy"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.client.registration.ResponseTypesConstraintsPlugin"
  }
  configuration {
    fields {
      name  = "code"
      value = "true"
    }
    fields {
      name  = "code id_token"
      value = "false"
    }
    fields {
      name  = "code id_token token"
      value = "false"
    }
    fields {
      name  = "code token"
      value = "false"
    }
    fields {
      name  = "id_token"
      value = "false"
    }
    fields {
      name  = "id_token token"
      value = "false"
    }
    fields {
      name  = "token"
      value = "true"
    }
  }
}
