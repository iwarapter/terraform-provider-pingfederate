package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func TestAccPingFederatePasswordCredentialValidatorResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederatePasswordCredentialValidatorResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederatePasswordCredentialValidatorResourceConfig("demo1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederatePasswordCredentialValidatorResourceExists("pingfederate_password_credential_validator.demo"),
					// testAccCheckPingFederatePasswordCredentialValidatorResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederatePasswordCredentialValidatorResourceConfig("demo2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederatePasswordCredentialValidatorResourceExists("pingfederate_password_credential_validator.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederatePasswordCredentialValidatorResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederatePasswordCredentialValidatorResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`resource "pingfederate_password_credential_validator" "demo" {
	  name = "foo"
	  plugin_descriptor_ref {
		id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
	  }
	
	  configuration {
		tables {
		  name = "Users"
		  rows {
			fields {
			  name  = "Username"
			  value = "%s"
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
	}`, configUpdate)
}

func testAccCheckPingFederatePasswordCredentialValidatorResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(*pf.PfClient).PasswordCredentialValidators
		result, _, err := conn.GetPasswordCredentialValidator(&pf.GetPasswordCredentialValidatorInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: PasswordCredentialValidator (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: PasswordCredentialValidator response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

//func Test_resourcePingFederatePasswordCredentialValidatorResourceReadData(t *testing.T) {
//	cases := []struct {
//		Resource pf.PasswordCredentialValidator
//	}{
//		{
//			Resource: pf.PasswordCredentialValidator{
//				Name: String("foo"),
//				PluginDescriptorRef: &pf.ResourceLink{
//					Id: String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
//				},
//				Configuration: &pf.PluginConfiguration{
//					Tables: &[]*pf.ConfigTable{
//						{
//							Name: String("Users"),
//							Rows: &[]*pf.ConfigRow{
//								{
//									Fields: &[]*pf.ConfigField{
//										{
//											Name:      String("Confirm Password"),
//											Value:     String("secret"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Password"),
//											Value:     String("secret"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Username"),
//											Value:     String("test"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Relax Password Requirements"),
//											Value:     String("true"),
//											Inherited: Bool(false),
//										},
//									},
//								},
//							},
//							Inherited: Bool(false),
//						},
//					},
//					Fields: &[]*pf.ConfigField{
//						{
//							Name:      String("Confirm Password"),
//							Value:     String("secret"),
//							Inherited: Bool(false),
//						},
//					},
//				},
//				AttributeContract: &pf.PasswordCredentialValidatorAttributeContract{
//					CoreAttributes: &[]*pf.PasswordCredentialValidatorAttribute{
//						{
//							Name: String("username"),
//						},
//					},
//				},
//			},
//		},
//	}
//	for i, tc := range cases {
//		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
//
//			descs := pf.PluginConfigDescriptor{
//				ActionDescriptors: nil,
//				Description:       nil,
//				Fields:            &[]*pf.FieldDescriptor{},
//				Tables: &[]*pf.TableDescriptor{
//					{
//						Columns: &[]*pf.FieldDescriptor{
//							{
//								Type: String("TEXT"),
//								Name: String("Username"),
//							},
//							{
//								Type: String("HASHED_TEXT"),
//								Name: String("Password"),
//							},
//							{
//								Type: String("HASHED_TEXT"),
//								Name: String("Confirm Password"),
//							},
//							{
//								Type: String("CHECK_BOX"),
//								Name: String("Relax Password Requirements"),
//							},
//						},
//						Description:       nil,
//						Label:             nil,
//						Name:              nil,
//						RequireDefaultRow: nil,
//					},
//				},
//			}
//
//			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
//				// Test request parameters
//				equals(t, req.URL.String(), "/passwordCredentialValidators/descriptors/org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator")
//				// Send response to be tested
//				b, _ := json.Marshal(pf.PasswordCredentialValidatorDescriptor{
//					AttributeContract:        nil,
//					ClassName:                String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
//					ConfigDescriptor:         &descs,
//					Id:                       String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
//					Name:                     String("Simple Username Password Credential Validator"),
//					SupportsExtendedContract: nil,
//				})
//				rw.Write(b)
//			}))
//			// Close the server when test finishes
//			defer server.Close()
//
//			// Use Client & URL from our local test server
//			url, _ := url.Parse(server.URL)
//			c := pf.NewClient("", "", url, "", server.Client())
//
//			resourceSchema := resourcePingFederatePasswordCredentialValidatorResourceSchema()
//			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
//			resourcePingFederatePasswordCredentialValidatorResourceReadResult(resourceLocalData, &pf.PasswordCredentialValidator{
//				Name: String("foo"),
//				PluginDescriptorRef: &pf.ResourceLink{
//					Id: String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"),
//				},
//				Configuration: &pf.PluginConfiguration{
//					Tables: &[]*pf.ConfigTable{
//						{
//							Name: String("Users"),
//							Rows: &[]*pf.ConfigRow{
//								{
//									Fields: &[]*pf.ConfigField{
//										{
//											Name:      String("Confirm Password"),
//											Value:     String("secret"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Password"),
//											Value:     String("secret"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Username"),
//											Value:     String("test"),
//											Inherited: Bool(false),
//										},
//										{
//											Name:      String("Relax Password Requirements"),
//											Value:     String("true"),
//											Inherited: Bool(false),
//										},
//									},
//								},
//							},
//							Inherited: Bool(false),
//						},
//					},
//					Fields: &[]*pf.ConfigField{
//						{
//							Name:           String("Confirm Password"),
//							EncryptedValue: String("secret"),
//							Inherited:      Bool(false),
//						},
//					},
//				},
//				AttributeContract: &pf.PasswordCredentialValidatorAttributeContract{
//					CoreAttributes: &[]*pf.PasswordCredentialValidatorAttribute{
//						{
//							Name: String("username"),
//						},
//					},
//				},
//			}, c.PasswordCredentialValidators)
//			resourcePingFederatePasswordCredentialValidatorResourceReadResult(resourceLocalData, &tc.Resource, c.PasswordCredentialValidators)
//
//			//tc.Resource.Configuration.Fields = &[]*pf.ConfigField{
//			//	&pf.ConfigField{
//			//		Name:           String("Confirm Password"),
//			//		EncryptedValue: String("secret"),
//			//		Inherited:      Bool(false),
//			//	},
//			//}
//			//resourcePingFederatePasswordCredentialValidatorResourceReadResult(resourceLocalData, &tc.Resource)
//
//			//TODO fix me
//			//if got := *resourcePingFederatePasswordCredentialValidatorResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
//			//	t.Errorf("resourcePingFederatePasswordCredentialValidatorResourceReadData() = %v", cmp.Diff(got, tc.Resource))
//			//}
//		})
//	}
//}
