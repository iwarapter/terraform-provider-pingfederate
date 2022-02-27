package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationApi"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("authentication_api_applications", &resource.Sweeper{
		Name: "authentication_api_applications",
		F: func(r string) error {
			svc := authenticationApi.New(cfg)
			results, _, err := svc.GetAuthenticationApiApplications()
			if err != nil {
				return fmt.Errorf("unable to list authentication api applications to sweep %s", err)
			}
			for _, item := range *results.Items {
				_, _, err = svc.DeleteApplication(&authenticationApi.DeleteApplicationInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep authentication api application %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateAuthnApiApplicationResource(t *testing.T) {
	resourceName := "pingfederate_authentication_api_application.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateAuthnApiApplicationResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthnApiApplicationResourceConfig("bar.foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthnApiApplicationResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://foo"),
					resource.TestCheckResourceAttr(resourceName, "description", "this is words"),
					resource.TestCheckResourceAttr(resourceName, "additional_allowed_origins.0", "https://bar.foo"),
					resource.TestCheckResourceAttr(resourceName, "additional_allowed_origins.1", "https://foo.com"),
				),
			},
			{
				Config: testAccPingFederateAuthnApiApplicationResourceConfig("foo.foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthnApiApplicationResourceExists("pingfederate_authentication_api_application.demo"),
					resource.TestCheckResourceAttr(resourceName, "name", "test"),
					resource.TestCheckResourceAttr(resourceName, "url", "https://foo"),
					resource.TestCheckResourceAttr(resourceName, "description", "this is words"),
					resource.TestCheckResourceAttr(resourceName, "additional_allowed_origins.0", "https://foo.com"),
					resource.TestCheckResourceAttr(resourceName, "additional_allowed_origins.1", "https://foo.foo"),
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

func testAccCheckPingFederateAuthnApiApplicationResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthnApiApplicationResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_authentication_api_application" "demo" {
	name      				   = "test"
	url				 		   = "https://foo"
	description 			   = "this is words"
	additional_allowed_origins = ["https://foo.com", "https://%s"]
}`, configUpdate)
}

func testAccCheckPingFederateAuthnApiApplicationResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationApi
		result, _, err := conn.GetApplication(&authenticationApi.GetApplicationInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: AuthnApiApplication (%s) not found", n)
		}

		if *result.Description != rs.Primary.Attributes["description"] {
			return fmt.Errorf("error: AuthnApiApplication response (%s) didnt match state (%s)", *result.Description, rs.Primary.Attributes["description"])
		}

		return nil
	}
}

func Test_resourcePingFederateAuthnApiApplicationResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AuthnApiApplication
	}{
		{
			Resource: pf.AuthnApiApplication{
				AdditionalAllowedOrigins: &[]*string{String("foo")},
				Description:              String("foo"),
				Id:                       String("foo"),
				Name:                     String("foo"),
				Url:                      String("foo"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateAuthnApiApplicationResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateAuthnApiApplicationResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateAuthnApiApplicationResourceReadData(resourceLocalData))
		})
	}
}
