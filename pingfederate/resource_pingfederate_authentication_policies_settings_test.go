package pingfederate

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateAuthenticationPoliciesSettingsResource(t *testing.T) {
	resourceName := "pingfederate_authentication_policies_settings.demo"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateAuthenticationPoliciesSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPoliciesSettingsResourceConfig("false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_idp_authn_selection", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_sp_authn_selection", "false"),
				),
			},
			{
				Config: testAccPingFederateAuthenticationPoliciesSettingsResourceConfig("true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPoliciesSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_idp_authn_selection", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_sp_authn_selection", "true"),
				),
			},
		},
	})
}

func testAccCheckPingFederateAuthenticationPoliciesSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthenticationPoliciesSettingsResourceConfig(first, second string) string {
	return fmt.Sprintf(`
resource "pingfederate_authentication_policies_settings" "demo" {
  enable_idp_authn_selection = %s
  enable_sp_authn_selection  = %s
}
`, first, second)
}

func testAccCheckPingFederateAuthenticationPoliciesSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationPolicies
		result, _, err := conn.GetSettings()

		if err != nil {
			return fmt.Errorf("error: AuthenticationPolicyContract (%s) not found", n)
		}

		if strconv.FormatBool(*result.EnableSpAuthnSelection) != rs.Primary.Attributes["enable_idp_authn_selection"] {
			return fmt.Errorf("error: AuthenticationPoliciesSettings response (%s) didnt match state (%s)", strconv.FormatBool(*result.EnableSpAuthnSelection), rs.Primary.Attributes["enable_idp_authn_selection"])
		}
		if strconv.FormatBool(*result.EnableSpAuthnSelection) != rs.Primary.Attributes["enable_sp_authn_selection"] {
			return fmt.Errorf("error: AuthenticationPoliciesSettings response (%s) didnt match state (%s)", strconv.FormatBool(*result.EnableSpAuthnSelection), rs.Primary.Attributes["enable_sp_authn_selection"])
		}

		return nil
	}
}
