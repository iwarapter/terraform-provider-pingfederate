package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingFederateKeyPairSslServerCertificateDatasource(t *testing.T) {
	resourceName := "data.pingfederate_keypair_ssl_server_certificate.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeyPairSslServerCertificateDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "certificate"),
				),
			},
		},
	})
}

func testAccPingFederateKeyPairSslServerCertificateDatasourceConfig() string {
	return `
resource "pingfederate_keypair_ssl_server" "example" {
  common_name   = "localhost"
  country       = "GB"
  key_algorithm = "RSA"
  key_size      = 2048
  organization  = "Test"
  valid_days    = 365

  lifecycle {
    create_before_destroy = true
  }
}

data "pingfederate_keypair_ssl_server_certificate" "test" {
  key_pair_id = pingfederate_keypair_ssl_server.example.id
}

resource "pingfederate_certificates_ca" "demo" {
  certificate_id = "local"
  file_data      = base64encode(data.pingfederate_keypair_ssl_server_certificate.test.certificate)
}
`
}
