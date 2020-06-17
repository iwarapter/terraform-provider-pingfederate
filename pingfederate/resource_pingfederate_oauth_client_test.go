package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateOauthClient(t *testing.T) {
	var out pf.Client

	resource.Test(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthClientDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthClientConfig("https://demo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientExists("pingfederate_oauth_client.my_client", &out),
					// testAccCheckPingFederateOauthClientAttributes(),
				),
			},
			{
				Config: testAccPingFederateOauthClientConfig("https://update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientExists("pingfederate_oauth_client.my_client", &out),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthClientDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthClientConfig(configUpdate string) string {
	return fmt.Sprintf(`
	%s

	resource "pingfederate_oauth_client" "my_client" {
		client_id = "tf-acc-woot"
		name      = "tf-acc-woot"
		
		grant_types = [
			"EXTENSION",
		]

		default_access_token_manager_ref {
			id = "${pingfederate_oauth_access_token_manager.my_atm.name}"
		}
	
		oidc_policy {
			grant_access_session_revocation_api = false
			logout_uris = [
				"https://logout",
				"%s"
			]
			ping_access_logout_capable = true
		}
	}`, testAccPingFederateOauthAccessTokenManagerConfig("client", "180"), configUpdate)
}

func testAccCheckPingFederateOauthClientExists(n string, out *pf.Client) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthClients
		result, _, err := conn.GetClient(&oauthClients.GetClientInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthClient (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: OauthClient response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}
