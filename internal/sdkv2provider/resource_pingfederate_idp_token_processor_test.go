package sdkv2provider

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpTokenProcessors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("idp_token_processor", &resource.Sweeper{
		Name:         "idp_token_processor",
		Dependencies: []string{},
		F: func(r string) error {
			svc := idpTokenProcessors.New(cfg)
			results, _, err := svc.GetTokenProcessors()
			if err != nil {
				return fmt.Errorf("unable to list IdpTokenProcessor %s", err)
			}
			for _, item := range *results.Items {
				if strings.HasPrefix(*item.Name, "acctest") {
					_, _, err := svc.DeleteTokenProcessor(&idpTokenProcessors.DeleteTokenProcessorInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep IdpTokenProcessor %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateIdpTokenProcessor(t *testing.T) {
	resourceName := "pingfederate_idp_token_processor.demo"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateIdpTokenProcessorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateIdpTokenProcessorConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpTokenProcessorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo"),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "org.sourceid.wstrust.processor.jwt.JWTTokenProcessor"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.0.name", "Expiry Tolerance"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.0.value", "0"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.1.name", "Issuer"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.1.value", "example"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.2.name", "JWKS Endpoint URI"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.2.value", "https://foo"),
				),
			},
			{
				Config: testAccPingFederateIdpTokenProcessorConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpTokenProcessorExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acctest_demo"),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "org.sourceid.wstrust.processor.jwt.JWTTokenProcessor"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.0.name", "Expiry Tolerance"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.0.value", "0"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.1.name", "Issuer"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.1.value", "example"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.2.name", "JWKS Endpoint URI"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.fields.2.value", "https://bar"),
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

func testAccCheckPingFederateIdpTokenProcessorDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateIdpTokenProcessorConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_idp_token_processor" "demo" {
  processor_id = "test123"
  name         = "acctest_demo"
  plugin_descriptor_ref {
    id = "org.sourceid.wstrust.processor.jwt.JWTTokenProcessor"
  }
  configuration {
    fields {
      name  = "JWKS Endpoint URI"
      value = "https://%s"
    }
    fields {
      name  = "Issuer"
      value = "example"
    }
    fields {
      name  = "Expiry Tolerance"
      value = "0"
    }
  }
}
`, configUpdate)
}

func testAccCheckPingFederateIdpTokenProcessorExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).IdpTokenProcessors
		result, _, err := conn.GetTokenProcessor(&idpTokenProcessors.GetTokenProcessorInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: idpTokenProcessor (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: idpTokenProcessor response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

type tokenProcessorMock struct {
	idpTokenProcessors.IdpTokenProcessorsAPI
}

func (t tokenProcessorMock) GetTokenProcessorDescriptorsById(input *idpTokenProcessors.GetTokenProcessorDescriptorsByIdInput) (output *pf.TokenProcessorDescriptor, resp *http.Response, err error) {
	return &pf.TokenProcessorDescriptor{
		AttributeContract: &[]*string{String("example1")},
		ClassName:         String("example"),
		ConfigDescriptor: &pf.PluginConfigDescriptor{
			ActionDescriptors: nil,
			Description:       nil,
			Fields:            &[]*pf.FieldDescriptor{},
			Tables: &[]*pf.TableDescriptor{
				{
					Columns: &[]*pf.FieldDescriptor{
						{
							TextFieldDescriptor: pf.TextFieldDescriptor{
								Encrypted: Bool(true),
							},
							Name: String("Password"),
							Type: String("TEXT"),
						},
					},
					Name:              String("Networks"),
					RequireDefaultRow: nil,
				},
			},
		},
		Id:                       String("example"),
		Name:                     String("example"),
		SupportsExtendedContract: nil,
	}, nil, nil
}

func Test_resourcePingFederateIdpTokenProcessorResourceReadData(t *testing.T) {
	mock := tokenProcessorMock{}
	cases := []struct {
		Name     string
		Resource pf.TokenProcessor
	}{
		{
			Name: "we can marshal absolutely every field",
			Resource: pf.TokenProcessor{
				AttributeContract: &pf.TokenProcessorAttributeContract{
					CoreAttributes: &[]*pf.TokenProcessorAttribute{
						{
							Name:   String("example1"),
							Masked: Bool(true),
						},
					},
					ExtendedAttributes: &[]*pf.TokenProcessorAttribute{
						{
							Name:   String("example2"),
							Masked: Bool(true),
						},
						{
							Name:   String("example3"),
							Masked: Bool(true),
						},
					},
					Inherited:      Bool(true),
					MaskOgnlValues: Bool(false),
				},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("Password"),
							Value:     String("secret"),
							Inherited: Bool(false),
						},
					},
					Tables: &[]*pf.ConfigTable{
						{
							Name:      String("Networks"),
							Inherited: Bool(false),
							Rows: &[]*pf.ConfigRow{
								{
									DefaultRow: Bool(true),
									Fields: &[]*pf.ConfigField{
										{
											Name:      String("Network Range (CIDR notation)"),
											Value:     String("0.0.0.0/0"),
											Inherited: Bool(false),
										},
									},
								},
							},
						},
					},
				},
				Id:   String("bar"),
				Name: String("foo"),
				ParentRef: &pf.ResourceLink{
					Id: String("foo"),
				},
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("foo"),
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", tc.Name), func(t *testing.T) {

			resourceSchema := resourcePingFederateIdpTokenProcessorResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateIdpTokenProcessorResourceReadResult(resourceLocalData, &tc.Resource, mock)

			body, err := resourcePingFederateIdpTokenProcessorResourceReadData(resourceLocalData, mock)
			require.NoError(t, err)
			assert.Equal(t, tc.Resource, *body)
		})
	}
}
