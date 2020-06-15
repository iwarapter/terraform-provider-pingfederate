package pingfederate

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
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

func testAccCheckPingFederateOauthAccessTokenManagerExists(n string, out *pf.Client) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(*pf.PfClient).OauthAccessTokenManagers
		result, _, err := conn.GetTokenManager(&pf.GetTokenManagerInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAccessTokenManager (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: OauthAccessTokenManager response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthAccessTokenManagerResourceReadData(t *testing.T) {
	descs := pf.PluginConfigDescriptor{
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
	}
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		equals(t, req.URL.String(), "/oauth/accessTokenManagers/descriptors/org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin")
		// Send response to be tested
		b, _ := json.Marshal(pf.AuthenticationSelectorDescriptor{
			ConfigDescriptor: &descs,
		})
		rw.Write(b)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	url, _ := url.Parse(server.URL)
	c := pf.NewClient("", "", url, "", server.Client())

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
			resourcePingFederateOauthAccessTokenManagersResourceReadResult(resourceLocalData, &tc.Resource, c.OauthAccessTokenManagers)

			if got := *resourcePingFederateOauthAccessTokenManagersResourceReadData(resourceLocalData, c.OauthAccessTokenManagers); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateOauthAccessTokenManagerResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
