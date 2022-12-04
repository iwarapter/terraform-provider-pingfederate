package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("keypair_signing", &resource.Sweeper{
		Name:         "keypair_signing",
		Dependencies: []string{},
		F: func(r string) error {
			svc := keyPairsSigning.New(cfg)
			results, _, err := svc.GetKeyPairs()
			if err != nil {
				return fmt.Errorf("unable to list signing keypairs %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteKeyPair(&keyPairsSigning.DeleteKeyPairInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep signing keypair %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateKeyPairSigning(t *testing.T) {
	resourceName := "pingfederate_keypair_signing.demo"
	resourceNameGen := "pingfederate_keypair_signing.test_generate"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateKeypairSigningDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeypairSigningConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairSigningExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "sha1_fingerprint", "596FC7AA20CEA185DA02AFEFBD239677D19BE43B"),
					resource.TestCheckResourceAttr(resourceName, "sha256_fingerprint", "C9D96118E2040126DE72E3DB1FC16930019047ED1032ED797E9C2F19E7028AD5"),
					resource.TestCheckResourceAttr(resourceName, "expires", "2024-08-13T20:21:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "issuer_dn", "CN=(LOCAL) CA, OU=LOCAL, O=ORGANISATION, L=LOCALITY, ST=STATE, C=--"),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "290604757668711433306844570316048963720924439315"),
					resource.TestCheckResourceAttr(resourceName, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "subject_dn", "CN=localhost.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "subject_alternative_names.0", "localhost"),
					resource.TestCheckResourceAttr(resourceName, "subject_alternative_names.1", "localhost.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "valid_from", "2019-08-15T20:21:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "version", "3"),
				),
			},
			//{
			//	ResourceName:      resourceName,
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
			{
				Config: testAccPingFederateKeypairSigningConfigGenerate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairSigningExists(resourceNameGen),
					resource.TestCheckResourceAttrSet(resourceNameGen, "sha1_fingerprint"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "sha256_fingerprint"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "expires"),
					resource.TestCheckResourceAttr(resourceNameGen, "issuer_dn", "CN=Test, OU=Test, O=Test, L=Test, ST=Test, C=GB"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "serial_number"),
					resource.TestCheckResourceAttr(resourceNameGen, "signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceNameGen, "status", "VALID"),
					resource.TestCheckResourceAttr(resourceNameGen, "subject_dn", "CN=Test, OU=Test, O=Test, L=Test, ST=Test, C=GB"),
					resource.TestCheckResourceAttrSet(resourceNameGen, "valid_from"),
					resource.TestCheckResourceAttr(resourceNameGen, "subject_alternative_names.0", "bar"),
					resource.TestCheckResourceAttr(resourceNameGen, "subject_alternative_names.1", "foo"),
				),
			},
			{
				ResourceName:      resourceNameGen,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingFederateKeypairSigningDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeypairSigningConfig() string {
	return `
resource "pingfederate_keypair_signing" "demo" {
  file_data = filebase64("test_cases/provider.p12")
  password  = "password"
}
`
}

func testAccPingFederateKeypairSigningConfigGenerate() string {
	return `
resource "pingfederate_keypair_signing" "test_generate" {
  city                      = "Test"
  common_name               = "Test"
  country                   = "GB"
  key_algorithm             = "RSA"
  key_size                  = 2048
  organization              = "Test"
  organization_unit         = "Test"
  state                     = "Test"
  valid_days                = 365
  subject_alternative_names = ["foo", "bar"]
}`
}

func testAccCheckPingFederateKeypairSigningExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).KeyPairsSigning
		result, _, err := conn.GetKeyPair(&keyPairsSigning.GetKeyPairInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: KeypairSigning (%s) not found", n)
		}

		if *result.Id != rs.Primary.Attributes["id"] {
			return fmt.Errorf("error: KeypairSigning response (%s) didnt match state (%s)", *result.Id, rs.Primary.Attributes["id"])
		}

		return nil
	}
}
