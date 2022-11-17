package sdkv2provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateAuthnApiSettingsResource(t *testing.T) {
	resourceName := "pingfederate_authentication_api_settings.demo"

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateAuthnApiSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthnApiSettingsResourceConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthnApiSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "api_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_api_descriptions", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "default_application_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "default_application_ref.0.location"),
				),
			},
			{
				Config: testAccPingFederateAuthnApiSettingsResourceConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthnApiSettingsResourceExists("pingfederate_authentication_api_settings.demo"),
					resource.TestCheckResourceAttr(resourceName, "api_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "enable_api_descriptions", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "default_application_ref.0.id"),
					resource.TestCheckNoResourceAttr(resourceName, "default_application_ref.0.location"),
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

func testAccCheckPingFederateAuthnApiSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthnApiSettingsResourceConfig(configUpdate string) string {
	ref := `default_application_ref {
		id = pingfederate_authentication_api_application.demo.id
	}`
	if configUpdate == "false" {
		ref = ""
	}

	return fmt.Sprintf(`
resource "pingfederate_authentication_api_settings" "demo" {
  api_enabled             = true
  enable_api_descriptions = %s
	%s
}

resource "pingfederate_authentication_api_application" "demo" {
  name        = "settings"
  url         = "https://bar"
  description = "this is words"
}`, configUpdate, ref)
}

func testAccCheckPingFederateAuthnApiSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationApi
		result, _, err := conn.GetAuthenticationApiSettings()

		if err != nil {
			return fmt.Errorf("error: AuthnApiSettings (%s) not found", n)
		}

		stored, err := strconv.ParseBool(rs.Primary.Attributes["enable_api_descriptions"])
		if err != nil {
			return fmt.Errorf("error: unable to parse attribute 'enable_api_descriptions' as bool %s", err.Error())
		}
		if *result.EnableApiDescriptions != stored {
			return fmt.Errorf("error: AuthnApiSettings response (%v) didnt match state (%s)", stored, rs.Primary.Attributes["description"])
		}

		return nil
	}
}
