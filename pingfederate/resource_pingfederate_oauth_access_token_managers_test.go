package pingfederate

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateOauthAccessTokenManager(t *testing.T) {
	var out pf.Client

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAccessTokenManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAccessTokenManagerConfig("acc", "120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenManagerExists("pingfederate_oauth_access_token_manager.my_atm", &out),
				),
			},
			{
				Config: testAccPingFederateOauthAccessTokenManagerConfig("acc", "180"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenManagerExists("pingfederate_oauth_access_token_manager.my_atm", &out),
				),
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
		name = "%s"
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
	name = "foo"
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

func testAccCheckPingFederateOauthAccessTokenManagerExists(n string, out *pf.Client) resource.TestCheckFunc {
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
				Id:                  String(""),
				Name:                String(""),
				PluginDescriptorRef: &pf.ResourceLink{Id: String("org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin")},
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
				AttributeContract: &pf.AccessTokenAttributeContract{ExtendedAttributes: &[]*pf.AccessTokenAttribute{
					{
						Name: String("foo"),
					},
				}},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateOauthAccessTokenManagersResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthAccessTokenManagersResourceReadResult(resourceLocalData, &tc.Resource, m)

			if got := *resourcePingFederateOauthAccessTokenManagersResourceReadData(resourceLocalData, m); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateOauthAccessTokenManagerResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
