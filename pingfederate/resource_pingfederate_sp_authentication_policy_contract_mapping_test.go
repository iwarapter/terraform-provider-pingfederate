package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/spAuthenticationPolicyContractMappings"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateSpAuthenticationPolicyContractMapping(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateSpAuthenticationPolicyContractMappingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateSpAuthenticationPolicyContractMappingConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpAuthenticationPolicyContractMappingExists("pingfederate_sp_authentication_policy_contract_mapping.demo"),
					//testAccCheckPingFederateSpAuthenticationPolicyContractMappingAttributes(),
				),
			},
			{
				Config: testAccPingFederateSpAuthenticationPolicyContractMappingConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpAuthenticationPolicyContractMappingExists("pingfederate_sp_authentication_policy_contract_mapping.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateSpAuthenticationPolicyContractMappingDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateSpAuthenticationPolicyContractMappingConfig(configUpdate string) string {
	return fmt.Sprintf(`resource "pingfederate_sp_authentication_policy_contract_mapping" "demo" {
	source_id = pingfederate_authentication_policy_contract.demo.id
	target_id = pingfederate_sp_adapter.demo.id
    attribute_contract_fulfillment {
      key_name = "subject"
      source {
        type = "AUTHENTICATION_POLICY_CONTRACT"
      }
      value = "subject"
    }
	default_target_resource = "https://%s"
}

resource "pingfederate_authentication_policy_contract" "demo" {
  name = "spadaptertest2"
  extended_attributes = ["foo", "email"]
}

resource "pingfederate_sp_adapter" "demo" {
  name = "spadaptertest2"
  sp_adapter_id = "spadaptertest2"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.opentoken.SpAuthnAdapter"
  }

  configuration {
    sensitive_fields {
      name  = "Password"
      value = "Secret123"
    }
    sensitive_fields {
      name  = "Confirm Password"
      value = "Secret123"
    }
    fields {
      name  = "Account Link Service"
      value = ""
    }
    fields {
      name  = "Authentication Service"
      value = ""
    }
    fields {
      name  = "Cipher Suite"
      value = "2"
    }
    fields {
      name  = "Cookie Domain"
      value = ""
    }
    fields {
      name  = "Cookie Path"
      value = "/"
    }
    fields {
      name  = "Force SunJCE Provider"
      value = "false"
    }
    fields {
      name  = "HTTP Only Flag"
      value = "true"
    }
    fields {
      name  = "Logout Service"
      value = ""
    }
    fields {
      name  = "Not Before Tolerance"
      value = "0"
    }
    fields {
      name  = "Obfuscate Password"
      value = "true"
    }
    fields {
      name  = "Secure Cookie"
      value = "false"
    }
    fields {
      name  = "Send Extended Attributes"
      value = ""
    }
    fields {
      name  = "Send Subject as Query Parameter"
      value = "false"
    }
    fields {
      name  = "Session Cookie"
      value = "false"
    }
    fields {
      name  = "Session Lifetime"
      value = "43200"
    }
    fields {
      name  = "Skip Trimming of Trailing Backslashes"
      value = "false"
    }
    fields {
      name  = "Subject Query Parameter                 "
      value = ""
    }
    fields {
      name  = "Token Lifetime"
      value = "300"
    }
    fields {
      name  = "Token Name"
      value = "opentoken"
    }
    fields {
      name  = "Transport Mode"
      value = "2"
    }
    fields {
      name  = "URL Encode Cookie Values"
      value = "true"
    }
    fields {
      name  = "Use Verbose Error Messages"
      value = "false"
    }

  }

  attribute_contract {
    core_attributes = [ "subject" ]
  }
  
  target_application_info {
	application_name = "foo"
	application_icon_url = "https://bar"
  }
}`, configUpdate)
}

func testAccCheckPingFederateSpAuthenticationPolicyContractMappingExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).SpAuthenticationPolicyContractMappings
		result, _, err := conn.GetApcToSpAdapterMappingById(&spAuthenticationPolicyContractMappings.GetApcToSpAdapterMappingByIdInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: SpAuthenticationPolicyContractMapping (%s) not found", n)
		}

		if *result.SourceId != rs.Primary.Attributes["source_id"] {
			return fmt.Errorf("Error: SpAuthenticationPolicyContractMapping response (%s) didnt match state (%s)", *result.SourceId, rs.Primary.Attributes["source_id"])
		}

		return nil
	}
}

func Test_resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.ApcToSpAdapterMapping
	}{
		{
			Resource: pf.ApcToSpAdapterMapping{
				Id:                               String(""),
				SourceId:                         String("foo"),
				TargetId:                         String("foo"),
				LicenseConnectionGroupAssignment: String("foo"),
				DefaultTargetResource:            String("foo"),
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
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateSpAuthenticationPolicyContractMappingResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
