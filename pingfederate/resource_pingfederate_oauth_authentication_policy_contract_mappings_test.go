package pingfederate

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_authentication_policy_contract_mapping", &resource.Sweeper{
		Name:         "oauth_authentication_policy_contract_mapping",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthAuthenticationPolicyContractMappings.New(cfg)
			results, _, err := svc.GetApcMappings()
			if err != nil {
				return fmt.Errorf("unable to list oauth authentication policy contract mappings %s", err)
			}
			for _, item := range *results.Items {
				if strings.Contains(*item.Id, "acc_test") {
					_, _, err := svc.DeleteApcMapping(&oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep oauth authentication policy contract mapping %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthAuthenticationPolicyContractMapping(t *testing.T) {
	resourceName := "pingfederate_oauth_authentication_policy_contract_mapping.demo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAuthenticationPolicyContractMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthenticationPolicyContractMappingConfig("subject"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthenticationPolicyContractMappingExists(resourceName),
					//testAccCheckPingFederateOauthAuthenticationPolicyContractMappingAttributes(),
				),
			},
			{
				Config: testAccPingFederateOauthAuthenticationPolicyContractMappingConfig("subject"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthenticationPolicyContractMappingExists(resourceName),
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

func testAccCheckPingFederateOauthAuthenticationPolicyContractMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthenticationPolicyContractMappingConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_authentication_policy_contract_mapping" "demo" {
  authentication_policy_contract_ref {
    id = pingfederate_authentication_policy_contract.demo.id
  }
  attribute_contract_fulfillment {
    key_name = "USER_NAME"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
  attribute_contract_fulfillment {
    key_name = "USER_KEY"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "subject"
  }
  attribute_contract_fulfillment {
    key_name = "woot"
    source {
      type = "AUTHENTICATION_POLICY_CONTRACT"
    }
    value = "%s"
  }
}

resource "pingfederate_authentication_policy_contract" "demo" {
  name = "acc_test_test3"
  extended_attributes = ["foo", "email"]
}`, configUpdate)
}

func testAccCheckPingFederateOauthAuthenticationPolicyContractMappingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAuthenticationPolicyContractMappings
		result, _, err := conn.GetApcMapping(&oauthAuthenticationPolicyContractMappings.GetApcMappingInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAuthenticationPolicyContractMapping (%s) not found", n)
		}

		if *result.AuthenticationPolicyContractRef.Id != rs.Primary.Attributes["authentication_policy_contract_ref.0.id"] {
			return fmt.Errorf("Error: OauthAuthenticationPolicyContractMapping response (%s) didnt match state (%s)", *result.AuthenticationPolicyContractRef.Id, rs.Primary.Attributes["authentication_policy_contract_ref.0.id"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthAuthenticationPolicyContractMappingResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.ApcToPersistentGrantMapping
	}{
		{
			Resource: pf.ApcToPersistentGrantMapping{
				Id: String(""),
				AuthenticationPolicyContractRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
					"foo": {
						Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
						Value:  String("foo"),
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
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
									Value:  String("foo"),
								},
							},
							BaseDn: String("foo"),
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
						AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
							"foo": {
								Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
								Value:  String("foo"),
							},
						},
						DataStoreRef: &pf.ResourceLink{
							Id: String("foo"),
						},
						Description: String("foo"),
						Id:          String("foo"),
						Type:        String("LDAP"),
					},
					{
						JdbcAttributeSource: pf.JdbcAttributeSource{
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
									Value:  String("foo"),
								},
							},
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
						AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
							"foo": {
								Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
								Value:  String("foo"),
							},
						},
						DataStoreRef: &pf.ResourceLink{
							Id: String("foo"),
						},
						Description: String("foo"),
						Id:          String("foo"),
						Type:        String("JDBC"),
					},
					{
						CustomAttributeSource: pf.CustomAttributeSource{
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
									Value:  String("foo"),
								},
							},
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
						AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
							"foo": {
								Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
								Value:  String("foo"),
							},
						},
						DataStoreRef: &pf.ResourceLink{
							Id: String("foo"),
						},
						Description: String("foo"),
						Id:          String("foo"),
						Type:        String("CUSTOM"),
					},
				},
			},
		},
		{
			Resource: pf.ApcToPersistentGrantMapping{
				Id: String(""),
				AuthenticationPolicyContractRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
				AttributeSources:             &[]*pf.AttributeSource{},
				IssuanceCriteria: &pf.IssuanceCriteria{
					ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
						{
							ErrorResult: String("foo"),
							Expression:  String("foo"),
						},
					},
				},
			},
		},
		{
			Resource: pf.ApcToPersistentGrantMapping{
				Id: String(""),
				AuthenticationPolicyContractRef: &pf.ResourceLink{
					Id: String("foo"),
				},
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
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateOauthAuthenticationPolicyContractMappingResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
