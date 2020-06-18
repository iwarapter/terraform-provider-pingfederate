package pingfederate

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateIdpAdapter(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateIdpAdapterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateIdpAdapterConfig("3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpAdapterExists("pingfederate_idp_adapter.demo"),
					//testAccCheckPingFederateIdpAdapterAttributes(),
				),
			},
			{
				Config: testAccPingFederateIdpAdapterConfig("4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpAdapterExists("pingfederate_idp_adapter.demo"),
				),
			},
			{
				Config:      testAccPingFederateIdpAdapterConfigWrongPlugin(),
				ExpectError: regexp.MustCompile(`unable to find plugin_descriptor for com\.pingidentity\.adapters\.httpbasic\.idp\.wrong available plugins:`),
			},
		},
	})
}

func testAccCheckPingFederateIdpAdapterDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateIdpAdapterConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_idp_adapter" "demo" {
  name = "barrr"
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
      value = "%s"
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
    jdbc_attribute_source {
      filter      = "\"\""
      description = "foo"
      schema      = "INFORMATION_SCHEMA"
      table       = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
      data_store_ref {
        id = "ProvisionerDS"
      }
    }
  }
}

resource "pingfederate_password_credential_validator" "demo" {
  name = "barrrrrr"
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
`, configUpdate)
}

func testAccPingFederateIdpAdapterConfigWrongPlugin() string {
	return `
resource "pingfederate_idp_adapter" "demo" {
  name = "barrr"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.httpbasic.idp.wrong"
  }

  configuration {
    fields {
      name  = "Realm"
      value = "foo"
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
`
}

func testAccCheckPingFederateIdpAdapterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).IdpAdapters
		result, _, err := conn.GetIdpAdapter(&idpAdapters.GetIdpAdapterInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: IdpAdapter (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdpAdapter response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateIdpAdapterResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.IdpAdapter
	}{
		{
			Resource: pf.IdpAdapter{
				Name: String("foo"),
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"),
				},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("Result Attribute Name"),
							Value:     String(""),
							Inherited: Bool(false),
						},
					},
					Tables: &[]*pf.ConfigTable{
						{
							Name:      String("Networks"),
							Inherited: Bool(false),
							Rows: &[]*pf.ConfigRow{
								{
									//DefaultRow: Bool(false),
									Fields: &[]*pf.ConfigField{
										{
											Name:      String("Network Range (CIDR notation)"),
											Value:     String("0.0.0.0/0"),
											Inherited: Bool(false),
										},
									},
								},
							},
						},
					},
				},
				AttributeMapping: &pf.IdpAdapterContractMapping{
					Inherited: Bool(false),
					AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
						"username": {
							Source: &pf.SourceTypeIdKey{
								Type: String("ADAPTER"),
							},
							Value: String("username"),
						},
						"sub": {
							Source: &pf.SourceTypeIdKey{
								Type: String("ADAPTER"),
							},
							Value: String("sub"),
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
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateIdpAdapterResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateIdpAdapterResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateIdpAdapterResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateIdpAdapterResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
