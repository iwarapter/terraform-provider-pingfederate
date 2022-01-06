package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesCa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("certificates_ca", &resource.Sweeper{
		Name:         "certificates_ca",
		Dependencies: []string{},
		F: func(r string) error {
			svc := certificatesCa.New(cfg)
			results, _, err := svc.GetTrustedCAs()
			if err != nil {
				return fmt.Errorf("unable to list certificate ca %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteTrustedCA(&certificatesCa.DeleteTrustedCAInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep certificate ca %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateCertificatesCa(t *testing.T) {
	resourceName := "pingfederate_certificates_ca.demo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateCertificatesCaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateCertificatesCaConfig("3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateCertificatesCaExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "sha1_fingerprint", "8DA7F965EC5EFC37910F1C6E59FDC1CC6A6EDE16"),
					resource.TestCheckResourceAttr(resourceName, "sha256_fingerprint", "8ECDE6884F3D87B1125BA31AC3FCB13D7016DE7F57CC904FE1CB97C6AE98196E"),
					resource.TestCheckResourceAttr(resourceName, "expires", "2038-01-17T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "issuer_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "143266978916655856878034712317230054538369994"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "subject_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "valid_from", "2015-05-26T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "version", "3"),
					resource.TestCheckResourceAttr(resourceName, "certificate_id", "3"),
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

func testAccCheckPingFederateCertificatesCaDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateCertificatesCaConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_certificates_ca" "demo" {
  certificate_id = "%s"
  file_data = base64encode(file("test_cases/amazon_root_ca2.pem"))
}
`, configUpdate)
}

func testAccCheckPingFederateCertificatesCaExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).CertificatesCa
		result, _, err := conn.GetTrustedCert(&certificatesCa.GetTrustedCertInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: CertificatesCa (%s) not found", n)
		}

		if *result.Id != rs.Primary.Attributes["id"] {
			return fmt.Errorf("Error: CertificatesCa response (%s) didnt match state (%s)", *result.Id, rs.Primary.Attributes["id"])
		}

		return nil
	}
}
