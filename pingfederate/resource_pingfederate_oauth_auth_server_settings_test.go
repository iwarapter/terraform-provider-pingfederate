package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateOauthAuthServerSettings(t *testing.T) {
	var out pf.AuthorizationServerSettings

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAuthServerSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "404"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists("pingfederate_oauth_auth_server_settings.settings", 3, &out),
				),
			},
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "403"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists("pingfederate_oauth_auth_server_settings.settings", 6, &out),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthAuthServerSettingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthServerSettingsConfig(name, configUpdate string) string {
	return `
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

		scope_groups {
			name        = "group1"
			description = "group1"
	
			scopes = [
				"address",
				"mail",
				"phone",
				"openid",
				"profile",
			]
		}

		persistent_grant_contract {
			extended_attributes = ["woot"]
		}

		allowed_origins = [
			"http://localhost"
		]

		default_scope_description  = ""
		authorization_code_timeout = 60
		authorization_code_entropy = 30
		refresh_token_length       = 42
		refresh_rolling_interval   = 0
	}`
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

		conn := testAccProvider.Meta().(pfClient).OauthAuthServerSettings
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
