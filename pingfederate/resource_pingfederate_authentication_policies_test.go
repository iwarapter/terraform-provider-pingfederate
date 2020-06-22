package pingfederate

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateAuthenticationPoliciesResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateAuthenticationPoliciesResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPoliciesResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesResourceExists("pingfederate_authentication_policies.demo"),
					// testAccCheckPingFederateAuthenticationPoliciesResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederateAuthenticationPoliciesResourceConfigTearDown(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesResourceExists("pingfederate_authentication_policies.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateAuthenticationPoliciesResourceDestroy(s *terraform.State) error {
	return nil
}
func testAccPingFederateAuthenticationPoliciesResourceConfigTearDown() string {
	return `
resource "pingfederate_authentication_policies" "demo" {
  fail_if_no_selection = false
  tracked_http_parameters = ["foo"]
}

resource "pingfederate_idp_adapter" "demo" {
  name = "testing"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"
  }

  configuration {
    tables {
      name = "Credential Validators"
      rows {
        fields {
          name  = "Password Credential Validator Instance"
          value = pingfederate_password_credential_validator.demo.name
        }
      }
    }
    fields {
      name  = "Realm"
      value = "foo"
    }

    fields {
      name  = "Challenge Retries"
      value = "3"
    }
  }

  attribute_contract {
    core_attributes {
      name      = "username"
      pseudonym = true
    }
    extended_attributes {
      name = "sub"
    }
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "ADAPTER"
      }
      value = "sub"
    }
    attribute_contract_fulfillment {
      key_name = "username"
      source {
        type = "ADAPTER"
      }
      value = "username"
    }
  }
}

resource "pingfederate_password_credential_validator" "demo" {
  name = "demo"
  plugin_descriptor_ref {
    id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
  }

  configuration {
    tables {
      name = "Users"
      rows {
        fields {
          name  = "Username"
          value = "example"
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
`
}

func testAccPingFederateAuthenticationPoliciesResourceConfig() string {
	return `
resource "pingfederate_authentication_policies" "demo" {
  fail_if_no_selection = false
  tracked_http_parameters = ["foo"]
  default_authentication_sources {
    type = "IDP_ADAPTER"
    source_ref {
      id = pingfederate_idp_adapter.demo.id
    }
  }
  authn_selection_trees {
    name = "bar"
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
          type = "RESTART"
          context = "Fail"
        }
      }
      children {
        action {
          type = "DONE"
          context = "Success"
        }
      }
    }
  }
}

resource "pingfederate_idp_adapter" "demo" {
  name = "testing"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"
  }

  configuration {
    tables {
      name = "Credential Validators"
      rows {
        fields {
          name  = "Password Credential Validator Instance"
          value = pingfederate_password_credential_validator.demo.name
        }
      }
    }
    fields {
      name  = "Realm"
      value = "foo"
    }

    fields {
      name  = "Challenge Retries"
      value = "3"
    }
  }

  attribute_contract {
    core_attributes {
      name      = "username"
      pseudonym = true
    }
    extended_attributes {
      name = "sub"
    }
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "ADAPTER"
      }
      value = "sub"
    }
    attribute_contract_fulfillment {
      key_name = "username"
      source {
        type = "ADAPTER"
      }
      value = "username"
    }
  }
}

resource "pingfederate_password_credential_validator" "demo" {
  name = "demo"
  plugin_descriptor_ref {
    id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
  }

  configuration {
    tables {
      name = "Users"
      rows {
        fields {
          name  = "Username"
          value = "example"
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
`
}

func testAccCheckPingFederateAuthenticationPoliciesResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationPolicies
		result, _, err := conn.GetDefaultAuthenticationPolicy()

		if err != nil {
			return fmt.Errorf("Error: AuthenticationPolicyContract (%s) not found", n)
		}

		if *(*result.TrackedHttpParameters)[0] != rs.Primary.Attributes["tracked_http_parameters.0"] {
			return fmt.Errorf("Error: AuthenticationPolicyContract response (%s) didnt match state (%s)", *(*result.TrackedHttpParameters)[0], rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateAuthenticationPoliciesResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AuthenticationPolicy
	}{
		{
			Resource: pf.AuthenticationPolicy{
				FailIfNoSelection: Bool(false),
				TrackedHttpParameters: &[]*string{
					String("foo"),
				},
				DefaultAuthenticationSources: &[]*pf.AuthenticationSource{
					{
						Type: String("IDP_ADAPTER"),
						SourceRef: &pf.ResourceLink{
							Id:       String("bar"),
							Location: String("https://foo.bar"),
						},
					},
				},
				AuthnSelectionTrees: &[]*pf.AuthenticationPolicyTree{
					{
						AuthenticationApiApplicationRef: &pf.ResourceLink{
							Id:       String("bar"),
							Location: String("https://foo.bar"),
						},
						Description: String("foo"),
						Enabled:     Bool(true),
						Name:        String("foo"),
						RootNode: &pf.AuthenticationPolicyTreeNode{
							Action: &pf.PolicyAction{
								ApcMappingPolicyAction: pf.ApcMappingPolicyAction{
									AttributeMapping: &pf.AttributeMapping{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Value:  String("bar"),
												Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
											},
										},
										AttributeSources: &[]*pf.AttributeSource{},
										IssuanceCriteria: &pf.IssuanceCriteria{},
									},
									AuthenticationPolicyContractRef: &pf.ResourceLink{Id: String("foo")},
								},
								LocalIdentityMappingPolicyAction: pf.LocalIdentityMappingPolicyAction{
									LocalIdentityRef: &pf.ResourceLink{Id: String("foo")},
									InboundMapping: &pf.AttributeMapping{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Value:  String("bar"),
												Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
											},
										},
										AttributeSources: &[]*pf.AttributeSource{},
										IssuanceCriteria: &pf.IssuanceCriteria{},
									},
									OutboundAttributeMapping: &pf.AttributeMapping{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Value:  String("bar"),
												Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
											},
										},
										AttributeSources: &[]*pf.AttributeSource{},
										IssuanceCriteria: &pf.IssuanceCriteria{},
									},
								},
								AuthnSelectorPolicyAction: pf.AuthnSelectorPolicyAction{
									AuthenticationSelectorRef: &pf.ResourceLink{Id: String("foo")},
								},
								AuthnSourcePolicyAction: pf.AuthnSourcePolicyAction{
									AuthenticationSource: &pf.AuthenticationSource{
										SourceRef: &pf.ResourceLink{Id: String("foo")},
										Type:      String("foo"),
									},
									InputUserIdMapping: &pf.AttributeFulfillmentValue{
										Value:  String("bar"),
										Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
									},
									AttributeRules: &pf.AttributeRules{
										FallbackToSuccess: Bool(true),
										Items: &[]*pf.AttributeRule{
											{
												AttributeName: String("foo"),
												Condition:     String("foo"),
												ExpectedValue: String("foo"),
												Result:        String("foo"),
											},
										},
									},
								},
								Context: String("Success"),
								Type:    String("DONE"),
							},
							Children: &[]*pf.AuthenticationPolicyTreeNode{
								{
									Action: &pf.PolicyAction{
										Type:    String("RESTART"),
										Context: String("Fail"),
									},
									Children: &[]*pf.AuthenticationPolicyTreeNode{
										{
											Action: &pf.PolicyAction{
												Type:    String("DONE"),
												Context: String("Fail"),
											},
										},
										{
											Action: &pf.PolicyAction{
												Type:    String("CONTINUE"),
												Context: String("Fail"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateAuthenticationPoliciesResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateAuthenticationPoliciesResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateAuthenticationPoliciesResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateAuthenticationPoliciesResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
