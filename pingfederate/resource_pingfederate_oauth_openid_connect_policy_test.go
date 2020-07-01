package pingfederate

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateOauthOpenIdConnectPolicy(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthOpenIdConnectPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthOpenIdConnectPolicyConfig("ClientId"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthOpenIdConnectPolicyExists("pingfederate_oauth_openid_connect_policy.demo"),
				),
			},
			{
				Config: testAccPingFederateOauthOpenIdConnectPolicyConfig("ClientId"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthOpenIdConnectPolicyExists("pingfederate_oauth_openid_connect_policy.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthOpenIdConnectPolicyDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthOpenIdConnectPolicyConfig(configUpdate string) string {
	return `
resource "pingfederate_oauth_openid_connect_policy" "demo" {
  policy_id = "foo"
  name      = "foo"
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.demo.id
  }
  attribute_contract {
    extended_attributes {
      name = "name"
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
      key_name = "name"
      source {
        type = "NO_MAPPING"
      }
    }
  }

  scope_attribute_mappings {
    key_name = "address"
	values = ["name"]
  }
}

resource "pingfederate_oauth_access_token_manager" "demo" {
  instance_id = "foo"
  name = "foo"
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
      value = "NONE"
    }

    fields {
      name  = "Maximum Token Lifetime"
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
      name = "Expand Scope Groups"
      value = "false"
    }
  }

  attribute_contract {
    extended_attributes = ["name"]
  }
}`
}

func testAccCheckPingFederateOauthOpenIdConnectPolicyExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		//conn := testAccProvider.Meta().(pfClient).OauthOpenIdConnect
		//result, _, err := conn.GetPolicy(&pf.GetPolicyInput{Id: rs.Primary.ID})
		//
		//if err != nil {
		//  return fmt.Errorf("Error: OauthOpenIdConnectPolicy (%s) not found", n)
		//}

		//if *result.AttributeMapping. != rs.Primary.Attributes["context.0.type"] {
		//  return fmt.Errorf("Error: OauthOpenIdConnectPolicy response (%s) didnt match state (%s)", *result.Context.Type, rs.Primary.Attributes["context.0.type"])
		//}

		return nil
	}
}

func Test_resourcePingFederateOauthOpenIdConnectPolicyResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.OpenIdConnectPolicy
	}{
		{
			Resource: pf.OpenIdConnectPolicy{
				Id:   String("foo"),
				Name: String("foo"),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeContract: &pf.OpenIdConnectAttributeContract{
					CoreAttributes: &[]*pf.OpenIdConnectAttribute{
						{
							Name:              String("sub"),
							IncludeInIdToken:  Bool(false),
							IncludeInUserInfo: Bool(true),
						},
					},
					ExtendedAttributes: &[]*pf.OpenIdConnectAttribute{
						{
							Name:              String("name"),
							IncludeInIdToken:  Bool(false),
							IncludeInUserInfo: Bool(true),
						},
					},
				},
				IdTokenLifetime: Int(5),
				AttributeMapping: &pf.AttributeMapping{
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
						"foo": {
							Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
							Value:  String("foo"),
						},
						"sub": {
							Source: &pf.SourceTypeIdKey{Type: String("foo")},
						},
					},
					IssuanceCriteria: &pf.IssuanceCriteria{
						ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
							{
								AttributeName: String("foo"),
								Condition:     String("foo"),
								ErrorResult:   String("foo"),
								Source: &pf.SourceTypeIdKey{
									Id:   String("foo"),
									Type: String("foo"),
								},
								Value: String("foo"),
							},
						},
						ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
							{
								ErrorResult: String("foo"),
								Expression:  String("foo"),
							},
						},
					},
					AttributeSources: &[]*pf.AttributeSource{
						{
							LdapAttributeSource: pf.LdapAttributeSource{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
								BaseDn:                       String("foo"),
								BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
									"foo": {BinaryEncoding: String("foo")},
								},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description:         String("foo"),
								Id:                  String("foo"),
								MemberOfNestedGroup: Bool(true),
								SearchFilter:        String("foo"),
								SearchScope:         String("foo"),
								Type:                String("LDAP"),
							},
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
							DataStoreRef: &pf.ResourceLink{
								Id: String("foo"),
							},
							Description: String("foo"),
							Id:          String("foo"),
							Type:        String("LDAP"),
						},
						{
							JdbcAttributeSource: pf.JdbcAttributeSource{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description: String("foo"),
								Filter:      String("foo"),
								Id:          String("foo"),
								Schema:      String("foo"),
								Table:       String("foo"),
								Type:        String("JDBC"),
							},
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
							DataStoreRef: &pf.ResourceLink{
								Id: String("foo"),
							},
							Description: String("foo"),
							Id:          String("foo"),
							Type:        String("JDBC"),
						},
						{
							CustomAttributeSource: pf.CustomAttributeSource{
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description: String("foo"),
								FilterFields: &[]*pf.FieldEntry{
									{
										Name:  String("foo"),
										Value: String("foo"),
									},
								},
								Id:   String("foo"),
								Type: String("CUSTOM"),
							},
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
							DataStoreRef: &pf.ResourceLink{
								Id: String("foo"),
							},
							Description: String("foo"),
							Id:          String("foo"),
							Type:        String("CUSTOM"),
						},
					},
				},
				ScopeAttributeMappings: map[string]*pf.ParameterValues{
					"Content-Type": {Values: &[]*string{
						String("charset=utf-8"),
						String("text/html"),
					}},
					"abc": {Values: &[]*string{
						String("456"),
						String("123"),
					}},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOpenIdConnectPolicyResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOpenIdConnectPolicyResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateOpenIdConnectPolicyResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
