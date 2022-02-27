package sdkv2provider

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_access_token_manager", &resource.Sweeper{
		Name:         "oauth_access_token_manager",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthAccessTokenManagers.New(cfg)
			results, _, err := svc.GetTokenManagers()
			if err != nil {
				return fmt.Errorf("unable to list oauth access token managers %s", err)
			}
			for _, item := range *results.Items {
				if strings.Contains(*item.Name, "acc_test") {
					_, _, err := svc.DeleteTokenManager(&oauthAccessTokenManagers.DeleteTokenManagerInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep oauth access token manager %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthAccessTokenManager(t *testing.T) {
	resourceName := "pingfederate_oauth_access_token_manager.my_atm"

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAccessTokenManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAccessTokenManagerConfig("acc", "120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenManagerExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateOauthAccessTokenManagerConfig("acc", "180"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenManagerExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:      testAccPingFederateOauthAccessTokenManagerConfigWrongPlugin(),
				ExpectError: regexp.MustCompile(`unable to find plugin_descriptor for org\.sourceid\.oauth20\.token\.plugin\.impl\.wrong available plugins:`),
			},
		},
	})
}

func testAccCheckPingFederateOauthAccessTokenManagerDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAccessTokenManagerConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_oauth_access_token_manager" "my_atm" {
		instance_id = "%s"
		name = "acc_test_%s"
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
				value = "%s"
			}

			fields {
				name  = "Lifetime Extension Policy"
				value = "ALL"
			}

			fields {
				name  = "Maximum Token Lifetime"
				value = "3000"
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
			extended_attributes = ["sub"]
		}
	}`, name, name, configUpdate)
}

func testAccPingFederateOauthAccessTokenManagerConfigWrongPlugin() string {
	return `
resource "pingfederate_oauth_access_token_manager" "my_atm" {
	instance_id = "foo"
	name = "acc_test_foo"
	plugin_descriptor_ref {
		id = "org.sourceid.oauth20.token.plugin.impl.wrong"
	}

	configuration {
		fields {
			name  = "Token Length"
			value = "28"
		}

		fields {
			name  = "Token Lifetime"
			value = "%s"
		}
	}

	attribute_contract {
		extended_attributes = ["sub"]
	}
}`
}

func testAccCheckPingFederateOauthAccessTokenManagerExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAccessTokenManagers
		result, _, err := conn.GetTokenManager(&oauthAccessTokenManagers.GetTokenManagerInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAccessTokenManager (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: OauthAccessTokenManager response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

type oauthAccessTokenManagersMock struct {
	oauthAccessTokenManagers.OauthAccessTokenManagersAPI
}

func (m oauthAccessTokenManagersMock) GetTokenManagerDescriptor(input *oauthAccessTokenManagers.GetTokenManagerDescriptorInput) (output *pf.AccessTokenManagerDescriptor, resp *http.Response, err error) {
	return &pf.AccessTokenManagerDescriptor{
		ConfigDescriptor: &pf.PluginConfigDescriptor{
			ActionDescriptors: nil,
			Description:       nil,
			Fields: &[]*pf.FieldDescriptor{
				{
					DefaultValue: String("28"),
					Label:        String("Token Length"),
					Name:         String("Token Length"),
					Required:     Bool(true),
					Type:         String("TEXT"),
				},
			},
			Tables: &[]*pf.TableDescriptor{
				{
					Columns: &[]*pf.FieldDescriptor{
						{
							Type: String("TEXT"),
							Name: String("Username"),
						},
					},
					Description:       nil,
					Label:             nil,
					Name:              String("Networks"),
					RequireDefaultRow: nil,
				},
			},
		},
	}, nil, nil
}

func Test_resourcePingFederateOauthAccessTokenManagerResourceReadData(t *testing.T) {
	m := &oauthAccessTokenManagersMock{}
	cases := []struct {
		Resource pf.AccessTokenManager
	}{
		{
			Resource: pf.AccessTokenManager{
				AttributeContract: &pf.AccessTokenAttributeContract{ExtendedAttributes: &[]*pf.AccessTokenAttribute{
					{
						Name: String("foo"),
					},
				}},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("foo"),
							Value:     String("bar"),
							Inherited: Bool(false),
						},
					},
					Tables: nil,
				},
				Id:                  String(""),
				Name:                String(""),
				PluginDescriptorRef: &pf.ResourceLink{Id: String("org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin")},
				ParentRef:           &pf.ResourceLink{Id: String("example")},
				AccessControlSettings: &pf.AtmAccessControlSettings{
					AllowedClients: &[]*pf.ResourceLink{
						{
							Id: String("one"),
						},
						{
							Id: String("two"),
						},
					},
					Inherited:       Bool(true),
					RestrictClients: Bool(true),
				},
				SelectionSettings: &pf.AtmSelectionSettings{
					Inherited:    Bool(true),
					ResourceUris: &[]*string{String("one"), String("two")},
				},
				SessionValidationSettings: &pf.SessionValidationSettings{
					CheckSessionRevocationStatus: Bool(true),
					CheckValidAuthnSession:       Bool(true),
					IncludeSessionId:             Bool(true),
					Inherited:                    Bool(true),
					UpdateAuthnSessionActivity:   Bool(true),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateOauthAccessTokenManagersResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthAccessTokenManagersResourceReadResult(resourceLocalData, &tc.Resource, m)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthAccessTokenManagersResourceReadData(resourceLocalData, m))
		})
	}
}
