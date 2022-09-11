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

func TestAccPingFederateKeyPairSslServerCsrDataSource(t *testing.T) {
	resourceName := "data.pingfederate_keypair_ssl_server_csr.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKeyPairSslServerCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeyPairSslServerCsrDataSourceConfig(),
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

func TestAccPingFederateKeyPairSslServerCsrDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKeyPairSslServerCsrDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingFederateKeyPairSslServerCsrDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`The resource with ID 'junk' is not found. Please specify a recognized resource ID.`),
			},
		},
	})
}

func testAccCheckPingFederateKeyPairSslServerCsrDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeyPairSslServerCsrDataSourceConfig() string {
	return `
resource "pingfederate_keypair_ssl_server" "ds_test" {
  city              = "Test"
  common_name       = "data store test ssl server"
  country           = "GB"
  key_algorithm     = "RSA"
  key_size          = 2048
  organization      = "Test"
  organization_unit = "Test"
  state             = "Test"
  valid_days        = 365
}

data "pingfederate_keypair_ssl_server_csr" "test" {
  id = pingfederate_keypair_ssl_server.ds_test.id
}`
}

func testAccPingFederateKeyPairSslServerCsrDataSourceConfigNonExistent() string {
	return `
data "pingfederate_keypair_ssl_server_csr" "test" {
  id = "junk"
}`
}
