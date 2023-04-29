package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/metadataUrls"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("metadata_url", &resource.Sweeper{
		Name:         "metadata_url",
		Dependencies: []string{},
		F: func(r string) error {
			results, _, err := pfc.MetadataUrls.GetMetadataUrls()
			if err != nil {
				return fmt.Errorf("unable to list metadata urls %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := pfc.MetadataUrls.DeleteMetadataUrl(&metadataUrls.DeleteMetadataUrlInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep metadata url %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateMetadataUrlResource(t *testing.T) {
	resourceName := "pingfederate_metadata_url.demo"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateMetadataUrlResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateMetadataUrlResourceConfig("acc_test_one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://sptest.iamshowcase.com/testsp_metadata.xml"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "true"),
				),
			},
			{
				Config: testAccPingFederateMetadataUrlResourceConfig("acc_test_one_changed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one_changed"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://sptest.iamshowcase.com/testsp_metadata.xml"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "true"),
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

func TestAccPingFederateMetadataUrlResource_SignedMetadata(t *testing.T) {
	resourceName := "pingfederate_metadata_url.demo"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateMetadataUrlResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateMetadataUrlResourceSignedMetadataConfig("acc_test_two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_two"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://localhost:9031/pf/federation_metadata.ping?PartnerIdpId=foobar"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "true"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.expires", "2084-01-29T16:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.issuer_dn", "O=Acme Co"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.serial_number", "97129276724337570813249812937731361303"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.sha1fingerprint", "15DBD260C7465ECCA6DE2C0B2181187F66EE0D1A"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.sha256fingerprint", "AB601914436E58BABB17B9166155CAF97BD7E5F8DEB9B659BCDB66C58B49F323"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.subject_dn", "O=Acme Co"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.valid_from", "1970-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.version", "3"),
				),
			},
			{
				Config: testAccPingFederateMetadataUrlResourceSignedMetadataConfig("acc_test_two_changed"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_two_changed"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://localhost:9031/pf/federation_metadata.ping?PartnerIdpId=foobar"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "true"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.expires", "2084-01-29T16:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.issuer_dn", "O=Acme Co"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.serial_number", "97129276724337570813249812937731361303"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.sha1fingerprint", "15DBD260C7465ECCA6DE2C0B2181187F66EE0D1A"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.sha256fingerprint", "AB601914436E58BABB17B9166155CAF97BD7E5F8DEB9B659BCDB66C58B49F323"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.subject_dn", "O=Acme Co"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.valid_from", "1970-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "cert_view.version", "3"),
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

func testAccCheckPingFederateMetadataUrlResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateMetadataUrlResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_metadata_url" "demo" {
  id                 = "acc_test_one"
  name               = "%s"
  url                = "https://sptest.iamshowcase.com/testsp_metadata.xml"
  validate_signature = true
}`, configUpdate)
}

func testAccPingFederateMetadataUrlResourceSignedMetadataConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_sp_idp_connection" "test" {
  active             = true
  base_url           = "https://localhost:9031"
  entity_id          = "foobar"
  error_page_msg_id  = "errorDetail.spSsoFailure"
  logging_mode       = "STANDARD"
  name               = "foobar"
  virtual_entity_ids = []

  credentials {
    certs {
      active_verification_cert    = true
      encryption_cert             = false
      primary_verification_cert   = true
      secondary_verification_cert = false

      x509_file {
        file_data = <<-EOT
                        -----BEGIN CERTIFICATE-----
                        MIIDOTCCAiGgAwIBAgIQSRJrEpBGFc7tNb1fb5pKFzANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQK
                        EwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYwMDAwWjASMRAwDgYDVQQKEwdB
                        Y21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6Gba5tHV1dAKouAaXO3/ebDU
                        U4rvwCUg/CNaJ2PT5xLD4N1Vcb8rbFSW2HXKq+MPfVdwIKR/1DczEoAGf/JWQTW7EgzlXrCd3rla
                        jEX2D73faWJekD0UaUgz5vtrTXZ90BQL7WvRICd7FlEZ6FPOcPlumiyNmzUqtwGhO+9ad1W5BqJa
                        RI6PYfouNkwR6Na4TzSj5BrqUfP0FwDizKSJ0XXmh8g8G9mtwxOSN3Ru1QFc61XyelukPOGKBV/q
                        6RBNklTNe0gI8usUMlYyoC7ytppNMW7X2vodAelSu25jgx2anj9fDVZuh7AXF5+4nJS4AAt0n1lN
                        Y7nGSsdZas8PbQIDAQABo4GIMIGFMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcD
                        ATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBStsdjh3/JCXXYlQryOrL4Sh7BW5TAuBgNVHREE
                        JzAlggtleGFtcGxlLmNvbYcEfwAAAYcQAAAAAAAAAAAAAAAAAAAAATANBgkqhkiG9w0BAQsFAAOC
                        AQEAxWGI5NhpF3nwwy/4yB4i/CwwSpLrWUa70NyhvprUBC50PxiXav1TeDzwzLx/o5HyNwsvcxv3
                        HdkLW59i/0SlJSrNnWdfZ19oTcS+6PtLoVyISgtyN6DpkKpdG1cOkW3Cy2P2+tK/tKHRP1Y/Ra0R
                        iDpOAmqn0gCOFGz8+lqDIor/T7MTpibL3IxqWfPrvfVRHL3Bgrw/ZQTTIVjjh4JBSW3WyWgNo/ik
                        C1lrVxzl4iPUGptxT36Cr7Zk2Bsg0XqwbOvK5d+NTDREkSnUbie4GeutujmX3Dsx88UiV6UY/4lH
                        Ja6I5leHUNOHahRbpbWeOfs/WkBKOclmOV2xlTVuPw==
                        -----END CERTIFICATE-----
                    EOT
        id        = "example"
      }
    }
  }

  idp_browser_sso {
    assertions_signed    = false
    enabled_profiles     = ["IDP_INITIATED_SSO"]
    idp_identity_mapping = "NONE"
    incoming_bindings    = ["POST"]
    protocol             = "SAML20"
    sign_authn_requests  = false

    attribute_contract {
      core_attributes {
        masked = false
        name   = "SAML_SUBJECT"
      }
    }

    decryption_policy {
      assertion_encrypted           = false
      attributes_encrypted          = false
      slo_encrypt_subject_name_id   = false
      slo_subject_name_id_encrypted = false
      subject_name_id_encrypted     = false
    }
  }
}

resource "pingfederate_metadata_url" "demo" {
  name               = "%s"
  url                = "https://localhost:9031/pf/federation_metadata.ping?PartnerIdpId=${pingfederate_sp_idp_connection.test.entity_id}"
  validate_signature = true
  x509file = {
    file_data = <<-EOT
                        -----BEGIN CERTIFICATE-----
                        MIIDOTCCAiGgAwIBAgIQSRJrEpBGFc7tNb1fb5pKFzANBgkqhkiG9w0BAQsFADASMRAwDgYDVQQK
                        EwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYwMDAwWjASMRAwDgYDVQQKEwdB
                        Y21lIENvMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6Gba5tHV1dAKouAaXO3/ebDU
                        U4rvwCUg/CNaJ2PT5xLD4N1Vcb8rbFSW2HXKq+MPfVdwIKR/1DczEoAGf/JWQTW7EgzlXrCd3rla
                        jEX2D73faWJekD0UaUgz5vtrTXZ90BQL7WvRICd7FlEZ6FPOcPlumiyNmzUqtwGhO+9ad1W5BqJa
                        RI6PYfouNkwR6Na4TzSj5BrqUfP0FwDizKSJ0XXmh8g8G9mtwxOSN3Ru1QFc61XyelukPOGKBV/q
                        6RBNklTNe0gI8usUMlYyoC7ytppNMW7X2vodAelSu25jgx2anj9fDVZuh7AXF5+4nJS4AAt0n1lN
                        Y7nGSsdZas8PbQIDAQABo4GIMIGFMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcD
                        ATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBStsdjh3/JCXXYlQryOrL4Sh7BW5TAuBgNVHREE
                        JzAlggtleGFtcGxlLmNvbYcEfwAAAYcQAAAAAAAAAAAAAAAAAAAAATANBgkqhkiG9w0BAQsFAAOC
                        AQEAxWGI5NhpF3nwwy/4yB4i/CwwSpLrWUa70NyhvprUBC50PxiXav1TeDzwzLx/o5HyNwsvcxv3
                        HdkLW59i/0SlJSrNnWdfZ19oTcS+6PtLoVyISgtyN6DpkKpdG1cOkW3Cy2P2+tK/tKHRP1Y/Ra0R
                        iDpOAmqn0gCOFGz8+lqDIor/T7MTpibL3IxqWfPrvfVRHL3Bgrw/ZQTTIVjjh4JBSW3WyWgNo/ik
                        C1lrVxzl4iPUGptxT36Cr7Zk2Bsg0XqwbOvK5d+NTDREkSnUbie4GeutujmX3Dsx88UiV6UY/4lH
                        Ja6I5leHUNOHahRbpbWeOfs/WkBKOclmOV2xlTVuPw==
                        -----END CERTIFICATE-----
                    EOT
  }
}`, configUpdate)
}

func testAccCheckPingFederateMetadataUrlResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := pfc.MetadataUrls
		result, _, err := conn.GetMetadataUrl(&metadataUrls.GetMetadataUrlInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: MetadataUrl (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: MetadataUrl response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateMetadataUrlResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.MetadataUrl
	}{
		{
			Resource: pf.MetadataUrl{
				CertView: &pf.CertView{
					CryptoProvider:          String("CryptoProvider"),
					Expires:                 String("Expires"),
					Id:                      String("Id"),
					IssuerDN:                String("IssuerDN"),
					KeyAlgorithm:            String("KeyAlgorithm"),
					KeySize:                 Int(1),
					SerialNumber:            String("SerialNumber"),
					Sha1Fingerprint:         String("Sha1Fingerprint"),
					Sha256Fingerprint:       String("Sha256Fingerprint"),
					SignatureAlgorithm:      String("SignatureAlgorithm"),
					Status:                  String("Status"),
					SubjectAlternativeNames: &[]*string{String("SubjectAlternativeNames")},
					SubjectDN:               String("SubjectDN"),
					ValidFrom:               String("ValidFrom"),
					Version:                 Int(2),
				},
				Id:                String("Id"),
				Name:              String("Name"),
				Url:               String("Url"),
				ValidateSignature: Bool(true),
				X509File: &pf.X509File{
					CryptoProvider: String("CryptoProvider"),
					FileData:       String("FileData"),
					Id:             String("X509File"),
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateMetadataUrlResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenMetadataUrl(&tc.Resource)).HasError())

			check := MetadataUrlData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandMetadataUrl(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
