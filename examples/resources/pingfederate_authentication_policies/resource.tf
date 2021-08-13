resource "pingfederate_authentication_policies" "demo" {
  fail_if_no_selection    = false
  tracked_http_parameters = ["foo"]
  authn_selection_trees {
    name = "one"
    root_node {
      action {
        type = "AUTHN_SELECTOR"
        authentication_selector_ref {
          id = pingfederate_authentication_selector.demo.id
        }
      }
      children {
        action {
          type    = "AUTHN_SELECTOR"
          context = "No"
          authentication_selector_ref {
            id = pingfederate_authentication_selector.demo.id
          }
        }
        children {
          action {
            type    = "CONTINUE"
            context = "No"
          }
        }
        children {
          action {
            type    = "CONTINUE"
            context = "Yes"
          }
        }
      }
      children {
        action {
          type    = "CONTINUE"
          context = "Yes"
        }
      }
    }
  }
  authn_selection_trees {
    name = "foo"
    root_node {
      action {
        type = "AUTHN_SOURCE"
        authentication_source {
          type = "IDP_ADAPTER"
          source_ref {
            id = pingfederate_idp_adapter.demo.id
          }
        }
      }
      children {
        action {
          type    = "RESTART"
          context = "Fail"
        }
      }
      children {
        action {
          type    = "DONE"
          context = "Success"
        }
      }
    }
  }
  dynamic "authn_selection_trees" {
    for_each = local.isPF10_2 ? [1] : []
    content {
      enabled = true
      name    = "frag"
      root_node {
        action {
          type = "FRAGMENT"
          fragment {
            id = pingfederate_authentication_policy_fragment.demo[0].id
          }
          fragment_mapping {
            attribute_contract_fulfillment {
              key_name = "one"

              source {
                type = "NO_MAPPING"
              }
            }
            attribute_contract_fulfillment {
              key_name = "subject"

              source {
                type = "NO_MAPPING"
              }
            }
            attribute_contract_fulfillment {
              key_name = "two"

              source {
                type = "NO_MAPPING"
              }
            }
          }
        }

        children {
          action {
            context = "Fail"
            type    = "DONE"
          }
        }
        children {
          action {
            context = "Success"
            type    = "DONE"
          }
        }
      }
    }
  }
}
