package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateExtendedPropertiesResource(t *testing.T) {
	resourceName := "pingfederate_extended_properties.demo"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateExtendedPropertiesResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateExtendedPropertiesResourceConfig("testing1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateExtendedPropertiesResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "property.0.name", "testing1"),
					resource.TestCheckResourceAttr(resourceName, "property.0.description", "my custom client property"),
				),
			},
			{
				Config: testAccPingFederateExtendedPropertiesResourceConfig("testing2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateExtendedPropertiesResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "property.0.name", "testing2"),
					resource.TestCheckResourceAttr(resourceName, "property.0.description", "my custom client property"),
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

func testAccCheckPingFederateExtendedPropertiesResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateExtendedPropertiesResourceConfig(first string) string {
	return fmt.Sprintf(`
resource "pingfederate_extended_properties" "demo" {
  property {
    name        = "%s"
    description = "my custom client property"
  }
}
`, first)
}

func testAccCheckPingFederateExtendedPropertiesResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).ExtendedProperties
		result, _, err := conn.GetExtendedProperties()

		if err != nil {
			return fmt.Errorf("error: ExtendedProperties (%s) not found", n)
		}

		if *(*result.Items)[0].Description != rs.Primary.Attributes["property.0.description"] {
			return fmt.Errorf("error: ExtendedProperties response (%s) didnt match state (%s)", *(*result.Items)[0].Description, rs.Primary.Attributes["property.0.description"])
		}

		return nil
	}
}
