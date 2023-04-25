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
				Config: testAccPingFederateMetadataUrlResourceConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://sptest.iamshowcase.com/testsp_metadata.xml"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "true"),
				),
			},
			{
				Config: testAccPingFederateMetadataUrlResourceConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateMetadataUrlResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://sptest.iamshowcase.com/testsp_metadata.xml"),
					resource.TestCheckResourceAttr(resourceName, "validate_signature", "false"),
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
  name               = "acc_test_one"
  url                = "https://sptest.iamshowcase.com/testsp_metadata.xml"
  validate_signature = %s
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
