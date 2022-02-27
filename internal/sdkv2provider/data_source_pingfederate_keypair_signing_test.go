package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateKeyPairSigningDataSource(t *testing.T) {
	resourceName := "data.pingfederate_keypair_signing.test"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKeyPairSigningDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeyPairSigningDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "keys.0.sha1_fingerprint"),
					resource.TestCheckResourceAttrSet(resourceName, "keys.0.sha256_fingerprint"),
					resource.TestCheckResourceAttrSet(resourceName, "keys.0.expires"),
					resource.TestCheckResourceAttrSet(resourceName, "keys.0.serial_number"),
					resource.TestCheckResourceAttrSet(resourceName, "keys.0.valid_from"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.subject_dn", "CN=data source test, OU=OrgUnit, O=Org, L=City, ST=State, C=GB"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.issuer_dn", "CN=data source test, OU=OrgUnit, O=Org, L=City, ST=State, C=GB"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.subject_alternative_names.0", "bar"),
					resource.TestCheckResourceAttr(resourceName, "keys.0.subject_alternative_names.1", "foo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateKeyPairSigningDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeyPairSigningDataSourceConfig() string {
	return `
        resource "pingfederate_keypair_signing" "ds_test" {
			city = "City"
			common_name = "data source test"
			country = "GB"
			key_algorithm = "RSA"
			key_size = 2048
			organization = "Org"
			organization_unit = "OrgUnit"
			state = "State"
			valid_days = 365
			subject_alternative_names = ["foo", "bar"]
        }

        data "pingfederate_keypair_signing" "test" {
			depends_on = [pingfederate_keypair_signing.ds_test]
		}`
}
