package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOauthAccessTokenMappings(t *testing.T) {
	resource.Test(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAccessTokenMappingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAccessTokenMappingsConfig("ClientId"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenMappingsExists("pingfederate_oauth_access_token_mappings.demo"),
				),
			},
			{
				Config: testAccPingFederateOauthAccessTokenMappingsConfig("ClientId"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenMappingsExists("pingfederate_oauth_access_token_mappings.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthAccessTokenMappingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAccessTokenMappingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_oauth_access_token_mappings" "demo" {
		access_token_manager_ref {
			id = pingfederate_oauth_access_token_manager.demo.id
		}

		context {
      		type = "CLIENT_CREDENTIALS"
    	}
		attribute_contract_fulfillment {
		  key_name = "sub"
		  source {
			type = "CONTEXT"
		  }
		  value = "%s"
		}
	}

resource "pingfederate_oauth_access_token_manager" "demo" {
		instance_id = "demo2"
		name = "demo2"
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
				value = "300"
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

func testAccCheckPingFederateOauthAccessTokenMappingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAccessTokenMappings
		result, _, err := conn.GetMapping(&oauthAccessTokenMappings.GetMappingInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthAccessTokenMappings (%s) not found", n)
		}

		if *result.Context.Type != rs.Primary.Attributes["context.0.type"] {
			return fmt.Errorf("Error: OauthAccessTokenMappings response (%s) didnt match state (%s)", *result.Context.Type, rs.Primary.Attributes["context.0.type"])
		}

		return nil
	}
}
