package pingfederate

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateKeypairsOauthOpenIdConnectResource(t *testing.T) {
	resourceName := "pingfederate_keypairs_oauth_openid_connect.settings"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKeypairsOauthOpenIdConnectResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeypairsOauthOpenIdConnectResourceConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairsOauthOpenIdConnectResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttr(resourceName, "rsa_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p256_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p384_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p521_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p256_decryption_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p384_decryption_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "p521_decryption_publish_x5c_parameter", "true"),
					resource.TestCheckResourceAttr(resourceName, "rsa_decryption_publish_x5c_parameter", "true"),
				),
			},
			{
				Config: testAccPingFederateKeypairsOauthOpenIdConnectResourceConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairsOauthOpenIdConnectResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "rsa_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p256_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p384_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_active_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_active_cert_ref.0.location"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_previous_cert_ref.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "p521_decryption_previous_cert_ref.0.location"),
					resource.TestCheckResourceAttr(resourceName, "rsa_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p256_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p384_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p521_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p256_decryption_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p384_decryption_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "p521_decryption_publish_x5c_parameter", "false"),
					resource.TestCheckResourceAttr(resourceName, "rsa_decryption_publish_x5c_parameter", "false"),
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

func testAccCheckPingFederateKeypairsOauthOpenIdConnectResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateKeypairsOauthOpenIdConnectResourceConfig(first string) string {
	return fmt.Sprintf(`
resource "pingfederate_keypairs_oauth_openid_connect" "settings" {
  static_jwks_enabled = true
  rsa_active_cert_ref {
	id = pingfederate_keypair_signing.rsa[0].id
  }
  rsa_previous_cert_ref {
	id = pingfederate_keypair_signing.rsa[1].id
  }
  rsa_decryption_active_cert_ref {
	id = pingfederate_keypair_signing.rsa[0].id
  }
  rsa_decryption_previous_cert_ref {
	id = pingfederate_keypair_signing.rsa[1].id
  }

  p256_active_cert_ref {
	id = pingfederate_keypair_signing.ec256[0].id
  }
  p256_previous_cert_ref {
	id = pingfederate_keypair_signing.ec256[1].id
  }
  p256_decryption_active_cert_ref {
	id = pingfederate_keypair_signing.ec256[0].id
  }
  p256_decryption_previous_cert_ref {
	id = pingfederate_keypair_signing.ec256[1].id
  }

  p384_active_cert_ref {
	id = pingfederate_keypair_signing.ec384[0].id
  }
  p384_previous_cert_ref {
	id = pingfederate_keypair_signing.ec384[1].id
  }
  p384_decryption_active_cert_ref {
	id = pingfederate_keypair_signing.ec384[0].id
  }
  p384_decryption_previous_cert_ref {
	id = pingfederate_keypair_signing.ec384[1].id
  }

  p521_active_cert_ref {
	id = pingfederate_keypair_signing.ec521[0].id
  }
  p521_previous_cert_ref {
	id = pingfederate_keypair_signing.ec521[1].id
  }
  p521_decryption_active_cert_ref {
	id = pingfederate_keypair_signing.ec521[0].id
  }
  p521_decryption_previous_cert_ref {
	id = pingfederate_keypair_signing.ec521[1].id
  }
  rsa_publish_x5c_parameter = %s
  p256_publish_x5c_parameter = %s
  p384_publish_x5c_parameter = %s
  p521_publish_x5c_parameter = %s
  p256_decryption_publish_x5c_parameter = %s
  p384_decryption_publish_x5c_parameter = %s
  p521_decryption_publish_x5c_parameter = %s
  rsa_decryption_publish_x5c_parameter = %s
}

resource "pingfederate_keypair_signing" "rsa" {
	count = 2
	city = "Test"
	common_name = "example${count.index}"
	country = "GB"
	key_algorithm = "RSA"
	key_size = 2048
	organization = "Test"
	organization_unit = "Test"
	state = "Test"
	valid_days = 365
	subject_alternative_names = ["foo", "bar"]
}

resource "pingfederate_keypair_signing" "ec256" {
	count = 2
	common_name = "example${count.index}"
	country = "GB"
	key_algorithm = "EC"
	key_size = 256
	organization = "Test"
	valid_days = 365
	subject_alternative_names = ["foo", "bar"]
}

resource "pingfederate_keypair_signing" "ec384" {
	count = 2
	common_name = "example${count.index}"
	country = "GB"
	key_algorithm = "EC"
	key_size = 384
	organization = "Test"
	valid_days = 365
	subject_alternative_names = ["foo", "bar"]
}

resource "pingfederate_keypair_signing" "ec521" {
	count = 2
	common_name = "example${count.index}"
	country = "GB"
	key_algorithm = "EC"
	key_size = 521
	organization = "Test"
	valid_days = 365
	subject_alternative_names = ["foo", "bar"]
}
`, first, first, first, first, first, first, first, first)
}

func testAccCheckPingFederateKeypairsOauthOpenIdConnectResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).KeyPairsOauthOpenIdConnect
		result, _, err := conn.GetOauthOidcKeysSettings()

		if err != nil {
			return fmt.Errorf("error: AuthenticationPolicyContract (%s) not found", n)
		}

		if strconv.FormatBool(*result.StaticJwksEnabled) != rs.Primary.Attributes["static_jwks_enabled"] {
			return fmt.Errorf("error: KeypairsOauthOpenIdConnect response (%s) didnt match state (%s)", strconv.FormatBool(*result.StaticJwksEnabled), rs.Primary.Attributes["static_jwks_enabled"])
		}

		return nil
	}
}

func Test_resourcePingFederateKeypairsOauthOpenIdConnectResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.OAuthOidcKeysSettings
	}{
		{
			Resource: pf.OAuthOidcKeysSettings{
				P256ActiveCertRef: &pf.ResourceLink{
					Id: String("1"),
				},
				P256DecryptionActiveCertRef: &pf.ResourceLink{
					Id: String("2"),
				},
				P256DecryptionPreviousCertRef: &pf.ResourceLink{
					Id: String("3"),
				},
				P256DecryptionPublishX5cParameter: Bool(true),
				P256PreviousCertRef: &pf.ResourceLink{
					Id: String("4"),
				},
				P256PublishX5cParameter: Bool(true),
				P384ActiveCertRef: &pf.ResourceLink{
					Id: String("5"),
				},
				P384DecryptionActiveCertRef: &pf.ResourceLink{
					Id: String("6"),
				},
				P384DecryptionPreviousCertRef: &pf.ResourceLink{
					Id: String("7"),
				},
				P384DecryptionPublishX5cParameter: Bool(true),
				P384PreviousCertRef: &pf.ResourceLink{
					Id: String("8"),
				},
				P384PublishX5cParameter: Bool(true),
				P521ActiveCertRef: &pf.ResourceLink{
					Id: String("9"),
				},
				P521DecryptionActiveCertRef: &pf.ResourceLink{
					Id: String("10"),
				},
				P521DecryptionPreviousCertRef: &pf.ResourceLink{
					Id: String("11"),
				},
				P521DecryptionPublishX5cParameter: Bool(true),
				P521PreviousCertRef: &pf.ResourceLink{
					Id: String("12"),
				},
				P521PublishX5cParameter: Bool(true),
				RsaActiveCertRef: &pf.ResourceLink{
					Id: String("13"),
				},
				RsaDecryptionActiveCertRef: &pf.ResourceLink{
					Id: String("14"),
				},
				RsaDecryptionPreviousCertRef: &pf.ResourceLink{
					Id: String("15"),
				},
				RsaDecryptionPublishX5cParameter: Bool(true),
				RsaPreviousCertRef: &pf.ResourceLink{
					Id: String("16"),
				},
				RsaPublishX5cParameter: Bool(true),
				StaticJwksEnabled:      Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateKeypairsOauthOpenIdConnectResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateKeypairsOauthOpenIdConnectResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateKeypairsOauthOpenIdConnectResourceReadData(resourceLocalData))
		})
	}
}
