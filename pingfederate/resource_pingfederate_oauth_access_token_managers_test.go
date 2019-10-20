package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func TestAccPingFederateOauthAccessTokenManager(t *testing.T) {
	var out pf.Client

	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAccessTokenManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAccessTokenManagerConfig("acc", "120"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAccessTokenManagerExists("pingfederate_oauth_access_token_manager.my_atm", &out),
					// testAccCheckPingFederateOauthAccessTokenManagerAttributes(),
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
