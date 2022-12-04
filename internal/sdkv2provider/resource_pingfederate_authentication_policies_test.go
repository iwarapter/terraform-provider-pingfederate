package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("authentication_policies", &resource.Sweeper{
		Name:         "authentication_policies",
		Dependencies: []string{},
		F: func(r string) error {
			svc := authenticationPolicies.New(cfg)
			_, _, err := svc.UpdateDefaultAuthenticationPolicy(&authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{Body: pf.AuthenticationPolicy{
				FailIfNoSelection: Bool(false),
			}})
			if err != nil {
				return fmt.Errorf("unable to reset authentication policies %s", err)
			}
			return nil
		},
	})
}

func TestAccPingFederateAuthenticationPoliciesResource(t *testing.T) {
	resourceName := "pingfederate_authentication_policies.demo"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateAuthenticationPoliciesResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPoliciesResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "fail_if_no_selection", "false"),
					resource.TestCheckResourceAttr(resourceName, "tracked_http_parameters.0", "foo"),
					resource.TestCheckResourceAttr(resourceName, "default_authentication_sources.0.type", "IDP_ADAPTER"),
					resource.TestCheckResourceAttr(resourceName, "default_authentication_sources.0.source_ref.0.id", "idptestme"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.name", "bar"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.action.0.type", "AUTHN_SOURCE"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.action.0.authentication_source.0.type", "IDP_ADAPTER"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.action.0.authentication_source.0.source_ref.0.id", "idptestme"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.0.action.0.type", "RESTART"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.0.action.0.context", "Fail"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.1.action.0.type", "DONE"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.1.action.0.context", "Condition"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.2.action.0.type", "DONE"),
					resource.TestCheckResourceAttr(resourceName, "authn_selection_trees.0.root_node.0.children.2.action.0.context", "Success"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccPingFederateAuthenticationPoliciesResourceConfigTearDown(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesResourceExists(resourceName),
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
  fail_if_no_selection    = false
  tracked_http_parameters = ["foo"]
}`
}

func testAccPingFederateAuthenticationPoliciesResourceConfig() string {
	return `
resource "pingfederate_authentication_policies" "demo" {
  fail_if_no_selection    = false
  tracked_http_parameters = ["foo"]
  default_authentication_sources {
    type = "IDP_ADAPTER"
    source_ref {
      id = "idptestme"
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
            id = "idptestme"
          }
        }
        attribute_rules {
          items {
            attribute_name = "sub"
            expected_value = "boo"
            result         = "Condition"
            condition      = "EQUALS"
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
          context = "Condition"
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
  authn_selection_trees {
    name = "one"
    root_node {
      action {
        type = "AUTHN_SELECTOR"
        authentication_selector_ref {
          id = "authseltestme"
        }
      }
      children {
        action {
          type    = "AUTHN_SELECTOR"
          context = "No"
          authentication_selector_ref {
            id = "authseltestme"
          }
        }
        children {
          action {
            type    = "AUTHN_SELECTOR"
            context = "No"
            authentication_selector_ref {
              id = "authseltestme"
            }
          }
          children {
            action {
              type    = "AUTHN_SELECTOR"
              context = "No"
              authentication_selector_ref {
                id = "authseltestme"
              }
            }
            children {
              action {
                type    = "AUTHN_SELECTOR"
                context = "No"
                authentication_selector_ref {
                  id = "authseltestme"
                }
              }
              children {
                action {
                  type    = "AUTHN_SELECTOR"
                  context = "No"
                  authentication_selector_ref {
                    id = "authseltestme"
                  }
                }
                children {
                  action {
                    type    = "AUTHN_SELECTOR"
                    context = "No"
                    authentication_selector_ref {
                      id = "authseltestme"
                    }
                  }
                  children {
                    action {
                      type    = "AUTHN_SELECTOR"
                      context = "No"
                      authentication_selector_ref {
                        id = "authseltestme"
                      }
                    }
                    children {
                      action {
                        type    = "AUTHN_SELECTOR"
                        context = "No"
                        authentication_selector_ref {
                          id = "authseltestme"
                        }
                      }
                      children {
                        action {
                          type    = "AUTHN_SELECTOR"
                          context = "No"
                          authentication_selector_ref {
                            id = "authseltestme"
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
  //authn_selection_trees {
  //  name = "foo"
  //  root_node {
  //    action {
  //      type = "AUTHN_SOURCE"
  //      authentication_source {
  //        type = "IDP_ADAPTER"
  //        source_ref {
  //          id = "idptestme"
  //        }
  //      }
  //    }
  //    children {
  //      action {
  //        type = "RESTART"
  //        context = "Fail"
  //      }
  //    }
  //    children {
  //      action {
  //        type = "DONE"
  //        context = "Success"
  //      }
  //    }
  //}
  //}
}`
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
									},
									OutboundAttributeMapping: &pf.AttributeMapping{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Value:  String("bar"),
												Source: &pf.SourceTypeIdKey{Id: String("key"), Type: String("t")},
											},
										},
										AttributeSources: &[]*pf.AttributeSource{},
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

			assert.Equal(t, tc.Resource, *resourcePingFederateAuthenticationPoliciesResourceReadData(resourceLocalData))
		})
	}
}

func Test_inconsistentSetElementsPanicReproducer(t *testing.T) {
	policy1 := pf.AuthenticationPolicy{
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
				Name:    String("bar"),
				Enabled: Bool(true),
				RootNode: &pf.AuthenticationPolicyTreeNode{
					Action: &pf.PolicyAction{
						Context: String("Success"),
						Type:    String("DONE"),
					},
				},
			},
		},
	}
	policy2 := pf.AuthenticationPolicy{
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
				Name:    String("bar"),
				Enabled: Bool(true),
				RootNode: &pf.AuthenticationPolicyTreeNode{
					Action: &pf.PolicyAction{
						ApcMappingPolicyAction: pf.ApcMappingPolicyAction{
							AttributeMapping: &pf.AttributeMapping{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"subject": {
										Source: &pf.SourceTypeIdKey{
											Id:   String("basicadptr"),
											Type: String("ADAPTER"),
										},
										Value: String("username"),
									},
									"family_name": {
										Source: &pf.SourceTypeIdKey{
											Id:   String("basicadptr"),
											Type: String("ADAPTER"),
										},
										Value: String("family_name"),
									},
									"first_name": {
										Source: &pf.SourceTypeIdKey{
											Id:   String("basicadptr"),
											Type: String("ADAPTER"),
										},
										Value: String("first_name"),
									},
									"email": {
										Source: &pf.SourceTypeIdKey{
											Id:   String("basicadptr"),
											Type: String("ADAPTER"),
										},
										Value: String("email"),
									},
								},
								AttributeSources: &[]*pf.AttributeSource{},
							},
							AuthenticationPolicyContractRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
						},
						Context: String("Success"),
						Type:    String("APC_MAPPING"),
					},
				},
			},
		},
	}

	resourceSchema := resourcePingFederateAuthenticationPoliciesResourceSchema()
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
	d := resourcePingFederateAuthenticationPoliciesResourceReadResult(resourceLocalData, &policy1)

	if d.HasError() {
		t.Errorf("Unexpected error saving state %v", d)
	}

	d = resourcePingFederateAuthenticationPoliciesResourceReadResult(resourceLocalData, &policy2)

	if d.HasError() {
		t.Errorf("Unexpected error saving state %v", d)
	}
}
