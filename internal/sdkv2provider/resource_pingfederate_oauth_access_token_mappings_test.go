package sdkv2provider

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("oauth_access_token_mappings", &resource.Sweeper{
		Name:         "oauth_access_token_mappings",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthAccessTokenMappings.New(cfg)
			results, _, err := svc.GetMappings()
			if err != nil {
				return fmt.Errorf("unable to list oauth access token mappings %s", err)
			}
			for _, item := range *results {
				if strings.Contains(*item.Id, "acc_test") {
					_, _, err := svc.DeleteMapping(&oauthAccessTokenMappings.DeleteMappingInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep oauth access token mappings %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthAccessTokenMappings(t *testing.T) {
	resourceName := "pingfederate_oauth_access_token_mappings.demo"

	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOauthAccessTokenMappingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAccessTokenMappingsConfig("ClientId"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenMappingsExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateOauthAccessTokenMappingsConfig("ClientIp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenMappingsExists(resourceName),
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

func testAccCheckPingFederateOauthAccessTokenMappingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAccessTokenMappingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_access_token_mappings" "demo" {
  access_token_manager_ref {
    id = "testme"
  }

  context {
    type = "CLIENT_CREDENTIALS"
  }
  attribute_contract_fulfillment {
    key_name = "name"
    source {
      type = "CONTEXT"
    }
    value = "%s"
  }
}`, configUpdate)
}

func testAccCheckPingFederateOauthAccessTokenMappingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAccessTokenMappings
		result, _, err := conn.GetMapping(&oauthAccessTokenMappings.GetMappingInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAccessTokenMappings (%s) not found", n)
		}

		if *result.Context.Type != rs.Primary.Attributes["context.0.type"] {
			return fmt.Errorf("Error: OauthAccessTokenMappings response (%s) didnt match state (%s)", *result.Context.Type, rs.Primary.Attributes["context.0.type"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthAccessTokenMappingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AccessTokenMapping
	}{
		{
			Resource: pf.AccessTokenMapping{
				AccessTokenManagerRef: &pf.ResourceLink{
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
							SearchAttributes: &[]*string{
								String("attr"),
							},
							Type: String("LDAP"),
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
				Context: &pf.AccessTokenMappingContext{
					ContextRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					Type: String("foo"),
				},
				Id: String(""),
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
			Resource: pf.AccessTokenMapping{
				Id: String(""),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id:       String("foo"),
					Location: String("foo"),
				},
				Context: &pf.AccessTokenMappingContext{
					ContextRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					Type: String("foo"),
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
			Resource: pf.AccessTokenMapping{
				Id: String(""),
				AccessTokenManagerRef: &pf.ResourceLink{
					Id:       String("foo"),
					Location: String("foo"),
				},
				Context: &pf.AccessTokenMappingContext{
					ContextRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					Type: String("foo"),
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

			resourceSchema := resourcePingFederateOauthAccessTokenMappingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthAccessTokenMappingsResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthAccessTokenMappingsResourceReadData(resourceLocalData))

		})
	}
}
