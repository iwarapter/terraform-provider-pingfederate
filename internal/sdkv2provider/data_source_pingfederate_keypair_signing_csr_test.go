package sdkv2provider

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateKeyPairSigningCsrDataSource(t *testing.T) {
	resourceName := "data.pingfederate_keypair_signing_csr.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateKeyPairSigningCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeyPairSigningCsrDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "cert_request_pem"),
					func(s *terraform.State) error {
						rs := s.RootModule().Resources[resourceName]
						atr := rs.Primary.Attributes["cert_request_pem"]
						block, _ := pem.Decode([]byte(atr))
						_, err := x509.ParseCertificateRequest(block.Bytes)
						if err != nil {
							return fmt.Errorf("unable to parse csr: %s", err)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccPingFederateKeyPairSigningCsrDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateKeyPairSigningCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingFederateKeyPairSigningCsrDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`The resource with ID 'junk' is not found. Please specify a recognized resource ID.`),
			},
		},
	})
}

func testAccCheckPingFederateKeyPairSigningCsrDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeyPairSigningCsrDataSourceConfig() string {
	return `
resource "pingfederate_keypair_signing" "ds_test" {
  city              = "Test"
  common_name       = "data store test signing"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Test"
  state             = "Test"
  valid_days        = 365
}

data "pingfederate_keypair_signing_csr" "test" {
  id = pingfederate_keypair_signing.ds_test.id
}`
}

func testAccPingFederateKeyPairSigningCsrDataSourceConfigNonExistent() string {
	return `
data "pingfederate_keypair_signing_csr" "test" {
  id = "junk"
}`
}
