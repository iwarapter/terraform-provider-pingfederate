package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOauthAuthServerSettings(t *testing.T) {
	resourceName := "pingfederate_oauth_auth_server_settings.settings"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAuthServerSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "404"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("bar", "403"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists(resourceName),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"persistent_grant_lifetime", "persistent_grant_lifetime_unit"},
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

		persistent_grant_lifetime      = -1
        persistent_grant_lifetime_unit = "DAYS"
		default_scope_description  = ""
		authorization_code_timeout = 60
		authorization_code_entropy = 30
		refresh_token_length       = 42
		refresh_rolling_interval   = 0
	}`
}

func testAccCheckPingFederateOauthAuthServerSettingsExists(n string) resource.TestCheckFunc {
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
