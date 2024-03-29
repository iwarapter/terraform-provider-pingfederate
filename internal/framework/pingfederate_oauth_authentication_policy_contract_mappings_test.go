package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_authentication_policy_contract_mapping", &resource.Sweeper{
		Name:         "oauth_authentication_policy_contract_mapping",
		Dependencies: []string{},
		F: func(r string) error {
			results, _, err := pfc.OauthAuthenticationPolicyContractMappings.GetApcMappings()
			if err != nil {
				return fmt.Errorf("unable to list oauth authentication policy contract mappings %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := pfc.OauthAuthenticationPolicyContractMappings.DeleteApcMapping(&oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep oauth authentication policy contract mapping %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthAuthenticationPolicyContractMapping(t *testing.T) {
	resourceName := "pingfederate_oauth_authentication_policy_contract_mapping.demo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOauthAuthenticationPolicyContractMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthenticationPolicyContractMappingConfig(`jdbc_attribute_sources = [
  	{
	  data_store_ref = "ProvisionerDS"
	  filter       	 = "*"
	  description  	 = "JDBC"
	  schema       	 = "INFORMATION_SCHEMA"
	  table        	 = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
	}
  ]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthenticationPolicyContractMappingExists(resourceName),
					//testAccCheckPingFederateOauthAuthenticationPolicyContractMappingAttributes(),
				),
			},
			{
				Config: testAccPingFederateOauthAuthenticationPolicyContractMappingConfig(""),
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

func TestAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV0toV1(t *testing.T) {
	resourceName := "pingfederate_oauth_authentication_policy_contract_mapping.test"
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"pingfederate": {
						VersionConstraint: "0.0.24",
						Source:            "iwarapter/pingfederate",
					},
				},
				Config: testAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV0config(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.key_name", "USER_KEY"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.0.value", "subject"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PlanOnly:                 true,
				Config:                   testAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV1config(),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV1config(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthenticationPolicyContractMappingExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.USER_KEY.value", "subject"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.USER_KEY.source.type", "AUTHENTICATION_POLICY_CONTRACT"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.USER_NAME.value", "subject"),
					resource.TestCheckResourceAttr(resourceName, "attribute_contract_fulfillment.USER_NAME.source.type", "AUTHENTICATION_POLICY_CONTRACT"),
				),
			},
		},
	})
}

func testAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV0config() string {
	return `
resource "pingfederate_ldap_data_store" "example" {
  ldap_type       = "ACTIVE_DIRECTORY"
  hostnames       = ["ldap:389"]
  max_connections = 100
  min_connections = 10
  name            = "ldap"
  user_dn         = "cn=admin,dc=example,dc=org"
  password        = "admin"
}

resource "pingfederate_authentication_policy_contract" "demo" {
  name                = "acc_test_upgrade2"
  extended_attributes = ["foo", "email"]
}

resource "pingfederate_oauth_authentication_policy_contract_mapping" "test" {
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
  ldap_attribute_source {
    description            = "desc"
    id                     = "ldap"
    member_of_nested_group = false
    search_filter          = "uid=$${subject}"
    search_scope           = "SUBTREE"

    data_store_ref {
      id = pingfederate_ldap_data_store.example.id
    }
  }
  jdbc_attribute_source {
    description = "jdbc"
    filter      = "uid=$${email}"
    id          = "jdbc"
    schema      = "INFORMATION_SCHEMA"
    table       = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"

    data_store_ref {
      id = "ProvisionerDS"
    }
  }

  issuance_criteria {
    conditional_criteria {
      attribute_name = "Subject DN"
      condition      = "EQUALS"
      value          = "foo"

      source {
        id   = "ldap"
        type = "LDAP_DATA_STORE"
      }
    }
    expression_criteria {
      expression = "far"
    }
    expression_criteria {
      error_result = "woot"
      expression   = "bar"
    }
  }
}
`
}
func testAccPingFederateOauthAuthenticationPolicyContractMappingResourceSdkUpgradeV1config() string {
	return `
resource "pingfederate_ldap_data_store" "example" {
  ldap_type       = "ACTIVE_DIRECTORY"
  hostnames       = ["ldap:389"]
  max_connections = 100
  min_connections = 10
  name            = "ldap"
  user_dn         = "cn=admin,dc=example,dc=org"
  password        = "admin"
}

resource "pingfederate_authentication_policy_contract" "demo" {
  name                = "acc_test_upgrade2"
  extended_attributes = ["foo", "email"]
}

resource "pingfederate_oauth_authentication_policy_contract_mapping" "test" {
  authentication_policy_contract_ref = pingfederate_authentication_policy_contract.demo.id
  attribute_contract_fulfillment = {
    "USER_NAME" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    },
    "USER_KEY" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    }
  }
  jdbc_attribute_sources = [
    {
      description    = "jdbc"
      filter         = "uid=$${email}"
      id             = "jdbc"
      schema         = "INFORMATION_SCHEMA"
      table          = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
      data_store_ref = "ProvisionerDS"
    }
  ]
  ldap_attribute_sources = [
    {
      description            = "desc"
      id                     = "ldap"
      member_of_nested_group = false
      search_filter          = "uid=$${subject}"
      search_scope           = "SUBTREE"
      data_store_ref         = pingfederate_ldap_data_store.example.id
    }
  ]

  issuance_criteria = {
    conditional_criteria = [
      {
        attribute_name = "Subject DN"
        condition      = "EQUALS"
        value          = "foo"

        source = {
          id   = "ldap"
          type = "LDAP_DATA_STORE"
        }
      }
    ]
    expression_criteria = [
      {
        expression = "far"
      },
      {
        error_result = "woot"
        expression   = "bar"
      }
    ]
  }
}
`
}

func testAccCheckPingFederateOauthAuthenticationPolicyContractMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthenticationPolicyContractMappingConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_authentication_policy_contract" "demo" {
  name                = "acc_test_test3"
  extended_attributes = ["foo", "email"]
}

resource "pingfederate_oauth_authentication_policy_contract_mapping" "demo" {
  authentication_policy_contract_ref = pingfederate_authentication_policy_contract.demo.id
  attribute_contract_fulfillment = {
    "USER_NAME" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    },
    "USER_KEY" = {
      source = {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    }
  }
  %s
}
`, configUpdate)
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

		conn := pfc.OauthAuthenticationPolicyContractMappings
		result, _, err := conn.GetApcMapping(&oauthAuthenticationPolicyContractMappings.GetApcMappingInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAuthenticationPolicyContractMapping (%s) not found", n)
		}

		if *result.AuthenticationPolicyContractRef.Id != rs.Primary.Attributes["authentication_policy_contract_ref"] {
			return fmt.Errorf("Error: OauthAuthenticationPolicyContractMapping response (%s) didnt match state (%s)", *result.AuthenticationPolicyContractRef.Id, rs.Primary.Attributes["authentication_policy_contract_ref"])
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
							ColumnNames: &[]*string{
								String("column"),
								String("name"),
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
			res := &pingfederateOauthAuthenticationPolicyContractMappingResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenApcToPersistentGrantMapping(&tc.Resource)).HasError())

			check := ApcToPersistentGrantMappingData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandApcToPersistentGrantMapping(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}

func Test_resourceWithExtraReturnedDataDoesntError(t *testing.T) {
	p := &pfClient{OauthAuthenticationPolicyContractMappings: oauthAuthenticationPolicyContractMappingsMock{}}

	model := &pf.ApcToPersistentGrantMapping{
		Id: String("foo"),
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
	}

	res := &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	ctx := context.Background()
	schResp := &fresource.SchemaResponse{}
	res.Schema(ctx, fresource.SchemaRequest{}, schResp)
	require.False(t, schResp.Diagnostics.HasError())

	r := pingfederateOauthAuthenticationPolicyContractMappingResource{p}

	resp := &fresource.CreateResponse{}
	r.Create(ctx, fresource.CreateRequest{
		Config: tfsdk.Config{
			Schema: schResp.Schema,
		},
		Plan: tfsdk.Plan{Schema: schResp.Schema},
	}, resp)

	//model := &pf.ApcToPersistentGrantMapping{
	//	Id: String("foo"),
	//	AuthenticationPolicyContractRef: &pf.ResourceLink{
	//		Id: String("foo"),
	//	},
	//	AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{},
	//	AttributeSources:             &[]*pf.AttributeSource{},
	//	IssuanceCriteria: &pf.IssuanceCriteria{
	//		ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
	//			{
	//				AttributeName: String("foo"),
	//				Condition:     String("foo"),
	//				ErrorResult:   String("foo"),
	//				Source: &pf.SourceTypeIdKey{
	//					Id:   String("foo"),
	//					Type: String("foo"),
	//				},
	//				Value: String("foo"),
	//			},
	//		},
	//	},
	//}

	state := tfsdk.State{Schema: schResp.Schema}

	require.False(t, state.Set(ctx, flattenApcToPersistentGrantMapping(model)).HasError())

	check := ApcToPersistentGrantMappingData{}
	require.False(t, state.Get(ctx, &check).HasError())

	//resp := *expandApcToPersistentGrantMapping(check)
	//assert.Equal(t, *model, resp)
}
