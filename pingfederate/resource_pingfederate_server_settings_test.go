package pingfederate

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateServerSettingsResource(t *testing.T) {
	resourceName := "pingfederate_server_settings.demo"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateServerSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateServerSettingsResourceConfig("testing1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateServerSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "federation_info.0.saml2_entity_id", "testing1"),
					resource.TestCheckResourceAttr(resourceName, "roles_and_protocols.0.sp_role.0.enable", "true"),
				),
			},
			{
				Config: testAccPingFederateServerSettingsResourceConfig("testing2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateServerSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "federation_info.0.saml2_entity_id", "testing2"),
					resource.TestCheckResourceAttr(resourceName, "roles_and_protocols.0.sp_role.0.enable", "true"),
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

func testAccCheckPingFederateServerSettingsResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateServerSettingsResourceConfig(first string) string {
	return fmt.Sprintf(`
resource "pingfederate_server_settings" "demo" {
  federation_info {
	base_url = "https://localhost:9031"
	saml2_entity_id = "%s"
  }
  roles_and_protocols {
    idp_role {
      enable = true
      saml20_profile {
        enable = true
      }
    }
    oauth_role {
      enable_oauth          = true
      enable_openid_connect = true
    }
    sp_role {
      enable = true
	  saml20_profile {
        enable = true
      }
    }
  }
}
`, first)
}

func testAccCheckPingFederateServerSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).ServerSettings
		result, _, err := conn.GetServerSettings()

		if err != nil {
			return fmt.Errorf("error: AuthenticationPolicyContract (%s) not found", n)
		}

		if strconv.FormatBool(*result.RolesAndProtocols.IdpRole.Enable) != rs.Primary.Attributes["roles_and_protocols.0.idp_role.0.enable"] {
			return fmt.Errorf("error: ServerSettings response (%s) didnt match state (%s)", strconv.FormatBool(*result.RolesAndProtocols.IdpRole.Enable), rs.Primary.Attributes["roles_and_protocols.0.idp_role.0.enable"])
		}
		if strconv.FormatBool(*result.RolesAndProtocols.SpRole.Enable) != rs.Primary.Attributes["roles_and_protocols.0.sp_role.0.enable"] {
			return fmt.Errorf("error: ServerSettings response (%s) didnt match state (%s)", strconv.FormatBool(*result.RolesAndProtocols.SpRole.Enable), rs.Primary.Attributes["roles_and_protocols.0.sp_role.0.enable"])
		}

		return nil
	}
}

func Test_resourcePingFederateServerSettinsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.ServerSettings
	}{
		{
			Resource: pf.ServerSettings{
				FederationInfo: &pf.FederationInfo{
					BaseUrl:        String("foo"),
					Saml1xIssuerId: String("foo"),
					Saml1xSourceId: String("foo"),
					Saml2EntityId:  String("foo"),
					WsfedRealm:     String("foo"),
				},
				RolesAndProtocols: &pf.RolesAndProtocols{
					EnableIdpDiscovery: Bool(true),
					OauthRole: &pf.OAuthRole{
						EnableOauth:         Bool(true),
						EnableOpenIdConnect: Bool(true),
					},
					IdpRole: &pf.IdpRole{
						Enable:                     Bool(true),
						EnableOutboundProvisioning: Bool(true),
						EnableSaml10:               Bool(true),
						EnableSaml11:               Bool(true),
						EnableWsFed:                Bool(true),
						EnableWsTrust:              Bool(true),
						Saml20Profile: &pf.SAML20Profile{
							Enable:            Bool(true),
							EnableAutoConnect: Bool(true),
						},
					},
					SpRole: &pf.SpRole{
						Enable:                    Bool(true),
						EnableInboundProvisioning: Bool(true),
						EnableOpenIDConnect:       Bool(true),
						EnableSaml10:              Bool(true),
						EnableSaml11:              Bool(true),
						EnableWsFed:               Bool(true),
						EnableWsTrust:             Bool(true),
						Saml20Profile: &pf.SpSAML20Profile{
							Enable:            Bool(true),
							EnableAutoConnect: Bool(true),
							EnableXASP:        Bool(true),
						},
					},
				},
			},
		},
		{
			Resource: pf.ServerSettings{
				FederationInfo: &pf.FederationInfo{
					BaseUrl:        String("foo"),
					Saml1xIssuerId: String("foo"),
					Saml1xSourceId: String("foo"),
					Saml2EntityId:  String("foo"),
					WsfedRealm:     String("foo"),
				},
				RolesAndProtocols: &pf.RolesAndProtocols{
					EnableIdpDiscovery: Bool(false),
					OauthRole: &pf.OAuthRole{
						EnableOauth:         Bool(false),
						EnableOpenIdConnect: Bool(false),
					},
					IdpRole: &pf.IdpRole{
						Enable:                     Bool(false),
						EnableOutboundProvisioning: Bool(false),
						EnableSaml10:               Bool(false),
						EnableSaml11:               Bool(false),
						EnableWsFed:                Bool(false),
						EnableWsTrust:              Bool(false),
						Saml20Profile: &pf.SAML20Profile{
							Enable:            Bool(false),
							EnableAutoConnect: Bool(false),
						},
					},
					SpRole: &pf.SpRole{
						Enable:                    Bool(false),
						EnableInboundProvisioning: Bool(false),
						EnableOpenIDConnect:       Bool(false),
						EnableSaml10:              Bool(false),
						EnableSaml11:              Bool(false),
						EnableWsFed:               Bool(false),
						EnableWsTrust:             Bool(false),
						Saml20Profile: &pf.SpSAML20Profile{
							Enable:            Bool(false),
							EnableAutoConnect: Bool(false),
							EnableXASP:        Bool(false),
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateServerSettingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateServerSettingsResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateServerSettingsResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateSpAdapterResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
