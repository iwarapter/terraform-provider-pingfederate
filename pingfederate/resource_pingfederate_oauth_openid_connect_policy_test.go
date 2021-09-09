package pingfederate

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthOpenIdConnect"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_openid_connect_policy", &resource.Sweeper{
		Name:         "oauth_openid_connect_policy",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthOpenIdConnect.New(cfg)
			settings, _, err := serverSettings.New(cfg).GetServerSettings()
			if err != nil {
				return fmt.Errorf("unable to check server settings %s", err)
			}
			if !*settings.RolesAndProtocols.OauthRole.EnableOpenIdConnect {
				return nil
			}
			results, _, err := svc.GetPolicies()
			if err != nil {
				return fmt.Errorf("unable to list oauth openid connect policy %s", err)
			}
			for _, item := range *results.Items {
				if strings.Contains(*item.Name, "acc_test") {
					_, _, err := svc.DeletePolicy(&oauthOpenIdConnect.DeletePolicyInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep oauth openid connect policy %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthOpenIdConnectPolicy(t *testing.T) {
	re := regexp.MustCompile(`^((10)\.[0-9])`)
	if !re.MatchString(pfVersion) {
		t.Skipf("This test only runs against PingFederate 10.0 and above, not: %s", pfVersion)
	}
	resourceName := "pingfederate_oauth_openid_connect_policy.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthOpenIdConnectPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthOpenIdConnectPolicyConfig("name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthOpenIdConnectPolicyExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateOauthOpenIdConnectPolicyConfig("name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthOpenIdConnectPolicyExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingFederateOauthOpenIdConnectPolicyDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthOpenIdConnectPolicyConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_openid_connect_policy" "demo" {
  policy_id = "acc_test_foo"
  name      = "acc_test_foo"
  access_token_manager_ref {
    id = "testme"
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
	values = ["%s"]
  }
}

resource "pingfederate_oauth_openid_connect_policy" "demo_two" {
  policy_id = "acc_test_bar"
  name      = "acc_test_bar"
  access_token_manager_ref {
    id = "testme"
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "NO_MAPPING"
      }
    }
  }
}`, configUpdate)
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
				IncludeSriInIdToken:         Bool(false),
				IncludeUserInfoInIdToken:    Bool(false),
				IncludeSHashInIdToken:       Bool(false),
				ReturnIdTokenOnRefreshGrant: Bool(false),
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
		{
			Resource: pf.OpenIdConnectPolicy{
				Id:              String("test two"),
				Name:            String("test two"),
				IdTokenLifetime: Int(5),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeMapping: &pf.AttributeMapping{
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
					AttributeSources:             &[]*pf.AttributeSource{},
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
					},
				},
				AttributeContract: &pf.OpenIdConnectAttributeContract{
					ExtendedAttributes: &[]*pf.OpenIdConnectAttribute{
						{
							IncludeInIdToken:  nil,
							IncludeInUserInfo: nil,
							Name:              String("foo"),
						},
					},
				},
			},
		},
		{
			Resource: pf.OpenIdConnectPolicy{
				Id:              String("test two"),
				Name:            String("test two"),
				IdTokenLifetime: Int(5),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeMapping: &pf.AttributeMapping{
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
					AttributeSources:             &[]*pf.AttributeSource{},
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
					},
				},
				AttributeContract: &pf.OpenIdConnectAttributeContract{
					ExtendedAttributes: &[]*pf.OpenIdConnectAttribute{
						{
							IncludeInIdToken:  Bool(true),
							IncludeInUserInfo: Bool(false),
							Name:              String("foo"),
						},
					},
				},
			},
		},
		{
			Resource: pf.OpenIdConnectPolicy{
				Id:              String("test two"),
				Name:            String("test two"),
				IdTokenLifetime: Int(5),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeMapping: &pf.AttributeMapping{
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
					AttributeSources:             &[]*pf.AttributeSource{},
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
					},
				},
				AttributeContract: &pf.OpenIdConnectAttributeContract{
					ExtendedAttributes: &[]*pf.OpenIdConnectAttribute{
						{
							IncludeInIdToken:  Bool(false),
							IncludeInUserInfo: Bool(true),
							Name:              String("foo"),
						},
					},
				},
			},
		},
		{
			Resource: pf.OpenIdConnectPolicy{
				Id:              String("test two"),
				Name:            String("test two"),
				IdTokenLifetime: Int(5),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeMapping: &pf.AttributeMapping{
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
					AttributeSources:             &[]*pf.AttributeSource{},
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
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOpenIdConnectPolicyResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOpenIdConnectPolicyResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateOpenIdConnectPolicyResourceReadData(resourceLocalData))
		})
	}
}

func Test_resourcePingFederateOauthOpenIdConnectPolicy_resourceOpenIdConnectAttribute_flatten_expand(t *testing.T) {
	atrs := []*pf.OpenIdConnectAttribute{
		{
			Name:              String("one"),
			IncludeInIdToken:  Bool(false),
			IncludeInUserInfo: Bool(true),
		},
		{
			Name:              String("two"),
			IncludeInIdToken:  Bool(true),
			IncludeInUserInfo: Bool(false),
		},
		{
			Name: String("three"),
			//IncludeInIdToken:  Bool(false),
			//IncludeInUserInfo: Bool(false),
		},
	}
	resourceSchema := resourceOpenIdConnectAttributeContract().Schema
	resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
	require.Nil(t, resourceLocalData.Set("extended_attributes", flattenOpenIdConnectAttributes(atrs)))
	assert.ElementsMatch(t, atrs, *expandOpenIdConnectAttributes(resourceLocalData.Get("extended_attributes").(*schema.Set).List()))

	//lets change our false example to true to enable
	atrs[2].IncludeInIdToken = Bool(true)
	atrs[2].IncludeInUserInfo = Bool(false)
	require.Nil(t, resourceLocalData.Set("extended_attributes", flattenOpenIdConnectAttributes(atrs)))
	assert.ElementsMatch(t, atrs, *expandOpenIdConnectAttributes(resourceLocalData.Get("extended_attributes").(*schema.Set).List()))
}
