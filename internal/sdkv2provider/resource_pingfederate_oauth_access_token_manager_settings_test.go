package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOAuthAccessTokenManagerSettingsResource(t *testing.T) {
	resourceName := "pingfederate_oauth_access_token_manager_settings.demo"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOAuthAccessTokenManagerSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthAccessTokenManagerSettingsResourceConfig("pingfederate_oauth_access_token_manager.example.id"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthAccessTokenManagerSettingsResourceExists(t, resourceName),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref.0.id", "settingsacctest"),
					resource.TestCheckResourceAttrSet(resourceName, "default_access_token_manager_ref.0.location"),
				),
			},
			{
				Config: testAccPingFederateOAuthAccessTokenManagerSettingsResourceConfig(`"testme"`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthAccessTokenManagerSettingsResourceExists(t, resourceName),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref.0.id", "testme"),
					resource.TestCheckResourceAttrSet(resourceName, "default_access_token_manager_ref.0.location"),
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

func testAccCheckPingFederateOAuthAccessTokenManagerSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOAuthAccessTokenManagerSettingsResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_access_token_manager_settings" "demo" {
	default_access_token_manager_ref {
		id = %s
	}
}

resource "pingfederate_oauth_access_token_manager" "example" {
	instance_id = "settingsacctest"
	name = "acc_test_settings"
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
			value = "28"
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
}`, configUpdate)
}

func testAccCheckPingFederateOAuthAccessTokenManagerSettingsResourceExists(t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAccessTokenManagers
		result, _, err := conn.GetSettings()

		if err != nil {
			return fmt.Errorf("error: OAuthAccessTokenManagerSettings (%s) not found", n)
		}

		assert.Equal(t, *result.DefaultAccessTokenManagerRef.Id, rs.Primary.Attributes["default_access_token_manager_ref.0.id"])

		return nil
	}
}
