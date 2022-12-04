package sdkv2provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateCertificatesRevocationSettingsResource(t *testing.T) {
	resourceName := "pingfederate_certificates_revocation_settings.demo"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateCertificatesRevocationSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateCertificatesRevocationSettingsResourceConfig(true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateCertificatesRevocationSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.next_retry_mins_when_next_update_in_past", "60"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.next_retry_mins_when_resolve_failed", "1440"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.treat_non_retrievable_crl_as_revoked", "false"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.verify_crl_signature", "true"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.action_on_responder_unavailable", "CONTINUE"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.action_on_status_unknown", "FAILOVER"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.action_on_unsuccessful_response", "FAIL"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.current_update_grace_period", "5"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.next_update_grace_period", "5"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.requester_add_nonce", "false"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.responder_timeout", "5"),
					resource.TestCheckResourceAttr(resourceName, "ocsp_settings.0.response_cache_period", "48"),
				),
			},
			{
				Config: testAccPingFederateCertificatesRevocationSettingsResourceConfig(false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateCertificatesRevocationSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.next_retry_mins_when_next_update_in_past", "60"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.next_retry_mins_when_resolve_failed", "1440"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.treat_non_retrievable_crl_as_revoked", "false"),
					resource.TestCheckResourceAttr(resourceName, "crl_settings.0.verify_crl_signature", "true"),
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

func testAccCheckPingFederateCertificatesRevocationSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateCertificatesRevocationSettingsResourceConfig(enableOcsp bool) string {
	if enableOcsp {
		return `resource "pingfederate_certificates_revocation_settings" "demo" {
  crl_settings {}
  ocsp_settings {
    responder_url            = "http://somewhere.foo"
    action_on_status_unknown = "FAILOVER"
  }
}
`
	}
	return `
resource "pingfederate_certificates_revocation_settings" "demo" {
  crl_settings {}
}
`
}

func testAccCheckPingFederateCertificatesRevocationSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).CertificatesRevocation
		result, _, err := conn.GetRevocationSettings()

		if err != nil {
			return fmt.Errorf("error: CertificatesRevocationSettings (%s) not found", n)
		}

		if strconv.FormatBool(*result.CrlSettings.VerifyCrlSignature) != rs.Primary.Attributes["crl_settings.0.verify_crl_signature"] {
			return fmt.Errorf("error: CertificatesRevocationSettings response (%s) didnt match state (%s)", strconv.FormatBool(*result.CrlSettings.VerifyCrlSignature), rs.Primary.Attributes["crl_settings.0.verify_crl_signature"])
		}
		return nil
	}
}

func Test_resourcePingFederateCertificatesRevocationSettingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.CertificateRevocationSettings
	}{
		{
			pf.CertificateRevocationSettings{
				CrlSettings: &pf.CrlSettings{
					NextRetryMinsWhenNextUpdateInPast: Int(1),
					NextRetryMinsWhenResolveFailed:    Int(1),
					TreatNonRetrievableCrlAsRevoked:   Bool(true),
					VerifyCrlSignature:                Bool(true),
				},
				OcspSettings: &pf.OcspSettings{
					ActionOnResponderUnavailable: String("1"),
					ActionOnStatusUnknown:        String("2"),
					ActionOnUnsuccessfulResponse: String("3"),
					CurrentUpdateGracePeriod:     Int(1),
					NextUpdateGracePeriod:        Int(2),
					RequesterAddNonce:            Bool(true),
					ResponderCertReference:       &pf.ResourceLink{Id: String("id")},
					ResponderTimeout:             Int(3),
					ResponderUrl:                 String("4"),
					ResponseCachePeriod:          Int(4),
				},
				ProxySettings: &pf.ProxySettings{
					Host: String("1"),
					Port: Int(1),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateCertificatesRevocationSettingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateCertificatesRevocationSettingsResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateCertificatesRevocationSettingsResourceReadData(resourceLocalData))
		})
	}
}
