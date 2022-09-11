package sdkv2provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthResourceOwnerCredentialsMappings"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("oauth_resource_owner_credentials_mappings", &resource.Sweeper{
		Name:         "oauth_resource_owner_credentials_mappings",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthResourceOwnerCredentialsMappings.New(cfg)
			results, _, err := svc.GetResourceOwnerCredentialsMappings()
			if err != nil {
				return fmt.Errorf("unable to list oauth resource owner credentials mappings %s", err)
			}
			for _, item := range *results.Items {
				if strings.Contains(*item.Id, "acc_test") {
					_, _, err := svc.DeleteResourceOwnerCredentialsMapping(&oauthResourceOwnerCredentialsMappings.DeleteResourceOwnerCredentialsMappingInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep oauth resource owner credentials mappings %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthResourceOwnerCredentialsMappings(t *testing.T) {
	resourceName := "pingfederate_oauth_resource_owner_credentials_mappings.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthResourceOwnerCredentialsMappingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthResourceOwnerCredentialsMappingsConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthResourceOwnerCredentialsMappingsExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.key_name", "USER_KEY"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.source.0.type", "NO_MAPPING"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.1.key_name", "woot"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.1.source.0.type", "NO_MAPPING"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.attribute_name", "username"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.condition", "EQUALS"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.error_result", "deny"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.value", "foo"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.source.0.type", "PASSWORD_CREDENTIAL_VALIDATOR"),
				),
			},
			{
				Config: testAccPingFederateOauthResourceOwnerCredentialsMappingsConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthResourceOwnerCredentialsMappingsExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.key_name", "USER_KEY"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.source.0.type", "NO_MAPPING"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.1.key_name", "woot"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.1.source.0.type", "NO_MAPPING"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.attribute_name", "username"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.condition", "EQUALS"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.error_result", "deny"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.value", "bar"),
					resource.TestCheckResourceAttr(resourceName, "issuance_criteria.0.conditional_criteria.0.source.0.type", "PASSWORD_CREDENTIAL_VALIDATOR"),
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

func testAccCheckPingFederateOauthResourceOwnerCredentialsMappingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthResourceOwnerCredentialsMappingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_resource_owner_credentials_mappings" "demo" {
  password_validator_ref {
    id = "pcvtestme"
  }

  attribute_contract_fulfillment {
    key_name = "USER_KEY"
    source {
      type = "NO_MAPPING"
    }
  }

  attribute_contract_fulfillment {
    key_name = "woot"
    source {
      type = "NO_MAPPING"
    }
  }
  issuance_criteria {
    conditional_criteria {
      attribute_name = "username"
      condition      = "EQUALS"
      error_result   = "deny"
      value          = "%s"

      source {
        type = "PASSWORD_CREDENTIAL_VALIDATOR"
      }
    }
  }
}`, configUpdate)
}

func testAccCheckPingFederateOauthResourceOwnerCredentialsMappingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthResourceOwnerCredentialsMappings
		result, _, err := conn.GetResourceOwnerCredentialsMapping(&oauthResourceOwnerCredentialsMappings.GetResourceOwnerCredentialsMappingInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: OauthResourceOwnerCredentialsMappings (%s) not found", n)
		}

		if *result.PasswordValidatorRef.Id != rs.Primary.Attributes["password_validator_ref.0.id"] {
			return fmt.Errorf("Error: OauthResourceOwnerCredentialsMappings response (%s) didnt match state (%s)", *result.PasswordValidatorRef.Id, rs.Primary.Attributes["password_validator_ref.0.id"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.ResourceOwnerCredentialsMapping
	}{
		{
			Resource: pf.ResourceOwnerCredentialsMapping{
				PasswordValidatorRef: &pf.ResourceLink{
					Id:       String("foo"),
					Location: String("foo"),
				},
				AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
					"foo": {
						Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
						Value:  String("foo"),
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
				Id: String("foo"),
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
			},
		},
		{
			Resource: pf.ResourceOwnerCredentialsMapping{
				Id: String("foo"),
				PasswordValidatorRef: &pf.ResourceLink{
					Id:       String("foo"),
					Location: String("foo"),
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
			Resource: pf.ResourceOwnerCredentialsMapping{
				Id: String("foo"),
				PasswordValidatorRef: &pf.ResourceLink{
					Id:       String("foo"),
					Location: String("foo"),
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

			resourceSchema := resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadData(resourceLocalData))

		})
	}
}
