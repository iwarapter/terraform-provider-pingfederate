package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func TestAccPingFederateOauthAuthServerSettings(t *testing.T) {
	var out pf.AuthorizationServerSettings

	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAuthServerSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "404"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists("pingfederate_oauth_auth_server_settings.settings", 3, &out),
					// testAccCheckPingFederateOauthAuthServerSettingsAttributes(),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName}, []string{roleName}, []string{groupName}, &out),
				),
			},
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "403"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists("pingfederate_oauth_auth_server_settings.settings", 6, &out),
					// testAccCheckAWSPolicyAttachmentAttributes([]string{userName2, userName3},
					// 	[]string{roleName2, roleName3}, []string{groupName2, groupName3}, &out),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthAuthServerSettingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthServerSettingsConfig(name, configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_oauth_auth_server_settings" "settings" {
		scopes {
			name        = "address"
			description = "address"
		}
	
		scopes {
			name        = "mail"
			description = "mail"
		}
	
		scopes {
			name        = "openid"
			description = "openid"
		}
	
		scopes {
			name        = "phone"
			description = "phone"
		}
	
		scopes {
			name        = "profile"
			description = "profile"
		}
	
		default_scope_description  = ""
		authorization_code_timeout = 60
		authorization_code_entropy = 30
		refresh_token_length       = 42
		refresh_rolling_interval   = 0
	}`) //, name, name, configUpdate)
}

func testAccCheckPingFederateOauthAuthServerSettingsExists(n string, c int64, out *pf.AuthorizationServerSettings) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(*pf.PfClient).OauthAuthServerSettings
		result, _, err := conn.GetAuthorizationServerSettings()

		if err != nil {
			return fmt.Errorf("Error: OauthAuthServerSettings (%s) not found", n)
		}

		if *result.DefaultScopeDescription != rs.Primary.Attributes["default_scope_description"] {
			return fmt.Errorf("Error: OauthAuthServerSettings response (%s) didnt match state (%s)", *result.DefaultScopeDescription, rs.Primary.Attributes["default_scope_description"])
		}

		return nil
	}
}
