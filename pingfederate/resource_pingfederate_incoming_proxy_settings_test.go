package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateIncomingProxySettings(t *testing.T) {
	resourceName := "pingfederate_incoming_proxy_settings.settings"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateIncomingProxySettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateIncomingProxySettingsConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIncomingProxySettingsExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateIncomingProxySettingsConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIncomingProxySettingsExists(resourceName),
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

func testAccCheckPingFederateIncomingProxySettingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateIncomingProxySettingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_incoming_proxy_settings" "settings" {
		client_cert_chain_sslheader_name = "foo"
		client_cert_sslheader_name = "foo"
		forwarded_host_header_index = "LAST"
		forwarded_host_header_name = "foo"
		forwarded_ip_address_header_index = "LAST"
		forwarded_ip_address_header_name = "%s"
		proxy_terminates_https_conns = false
	}`, configUpdate)
}

func testAccCheckPingFederateIncomingProxySettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).IncomingProxySettings
		_, _, err := conn.GetIncomingProxySettings()

		if err != nil {
			return fmt.Errorf("error: IncomingProxySettings (%s) not found", n)
		}

		//if *result.ForwardedIpAddressHeaderIndex != rs.Primary.Attributes["forwarded_ip_address_header_index"] {
		//	return fmt.Errorf("error: IncomingProxySettings response (%s) didnt match state (%s)", *result.ForwardedIpAddressHeaderIndex, rs.Primary.Attributes["forwarded_ip_address_header_index"])
		//}

		return nil
	}
}
