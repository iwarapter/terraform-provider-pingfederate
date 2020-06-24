package pingfederate

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/spAdapters"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateSpAdapter(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateSpAdapterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateSpAdapterConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpAdapterExists("pingfederate_sp_adapter.demo"),
					//testAccCheckPingFederateSpAdapterAttributes(),
				),
			},
			{
				Config: testAccPingFederateSpAdapterConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpAdapterExists("pingfederate_sp_adapter.demo"),
				),
			},
			{
				Config:      testAccPingFederateSpAdapterConfigWrongPlugin(),
				ExpectError: regexp.MustCompile(`unable to find plugin_descriptor for com\.pingidentity\.adapters\.opentoken\.wrong available plugins:`),
			},
		},
	})
}

func testAccCheckPingFederateSpAdapterDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateSpAdapterConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_sp_adapter" "demo" {
  name = "bar"
  sp_adapter_id = "spadaptertest1"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.opentoken.SpAuthnAdapter"
  }

  configuration {
    sensitive_fields {
      name  = "Password"
      value = "Secret123"
    }
    sensitive_fields {
      name  = "Confirm Password"
      value = "Secret123"
    }
    fields {
      name  = "Account Link Service"
      value = ""
    }
    fields {
      name  = "Authentication Service"
      value = ""
    }
    fields {
      name  = "Cipher Suite"
      value = "2"
    }
    fields {
      name  = "Cookie Domain"
      value = ""
    }
    fields {
      name  = "Cookie Path"
      value = "/"
    }
    fields {
      name  = "Force SunJCE Provider"
      value = "false"
    }
    fields {
      name  = "HTTP Only Flag"
      value = "true"
    }
    fields {
      name  = "Logout Service"
      value = ""
    }
    fields {
      name  = "Not Before Tolerance"
      value = "0"
    }
    fields {
      name  = "Obfuscate Password"
      value = "true"
    }
    fields {
      name  = "Secure Cookie"
      value = "false"
    }
    fields {
      name  = "Send Extended Attributes"
      value = ""
    }
    fields {
      name  = "Send Subject as Query Parameter"
      value = "false"
    }
    fields {
      name  = "Session Cookie"
      value = "false"
    }
    fields {
      name  = "Session Lifetime"
      value = "43200"
    }
    fields {
      name  = "Skip Trimming of Trailing Backslashes"
      value = "false"
    }
    fields {
      name  = "Subject Query Parameter                 "
      value = ""
    }
    fields {
      name  = "Token Lifetime"
      value = "300"
    }
    fields {
      name  = "Token Name"
      value = "opentoken"
    }
    fields {
      name  = "Transport Mode"
      value = "2"
    }
    fields {
      name  = "URL Encode Cookie Values"
      value = "true"
    }
    fields {
      name  = "Use Verbose Error Messages"
      value = "false"
    }

  }

  attribute_contract {
    core_attributes = [ "subject" ]
  }

  target_application_info {
	application_name = "foo"
	application_icon_url = "https://%s"
  }
}
`, configUpdate)
}

func testAccPingFederateSpAdapterConfigWrongPlugin() string {
	return `
resource "pingfederate_sp_adapter" "demo" {
  name = "bar"
  sp_adapter_id = "spadaptertest1"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.opentoken.wrong"
  }

  configuration {
    fields {
      name  = "Use Verbose Error Messages"
      value = "false"
    }
  }

  attribute_contract {
    core_attributes = [ "subject" ]
  }

  target_application_info {
	application_name = "foo"
	application_icon_url = "https://%s"
  }
}
`
}

func testAccCheckPingFederateSpAdapterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).SpAdapters
		result, _, err := conn.GetSpAdapter(&spAdapters.GetSpAdapterInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: SpAdapter (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: SpAdapter response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

type spAdaptersMock struct {
	spAdapters.SpAdaptersAPI
}

func (s spAdaptersMock) GetSpAdapterDescriptorsById(input *spAdapters.GetSpAdapterDescriptorsByIdInput) (output *pf.SpAdapterDescriptor, resp *http.Response, err error) {
	return &pf.SpAdapterDescriptor{
		AttributeContract: nil,
		ClassName:         String("com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"),
		ConfigDescriptor: &pf.PluginConfigDescriptor{
			ActionDescriptors: nil,
			Description:       nil,
			Fields:            &[]*pf.FieldDescriptor{},
			Tables: &[]*pf.TableDescriptor{
				{
					Columns: &[]*pf.FieldDescriptor{
						{
							Type: String("TEXT"),
							Name: String("Username"),
						},
					},
					Description:       nil,
					Label:             nil,
					Name:              String("Networks"),
					RequireDefaultRow: nil,
				},
			},
		},
		Id:                       String("com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"),
		Name:                     String("CIDR Authentication Selector"),
		SupportsExtendedContract: nil,
	}, nil, nil
}

func Test_resourcePingFederateSpAdapterResourceReadData(t *testing.T) {
	m := &spAdaptersMock{}
	cases := []struct {
		Resource pf.SpAdapter
	}{
		{
			Resource: pf.SpAdapter{
				Name: String("foo"),
				Id:   String("foo"),
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("com.pingidentity.adapters.httpbasic.idp.HttpBasicSpAuthnAdapter"),
				},
				Configuration: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:      String("Result Attribute Name"),
							Value:     String(""),
							Inherited: Bool(false),
						},
					},
					Tables: &[]*pf.ConfigTable{
						{
							Name:      String("Networks"),
							Inherited: Bool(false),
							Rows: &[]*pf.ConfigRow{
								{
									//DefaultRow: Bool(false),
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
				AttributeContract: &pf.SpAdapterAttributeContract{
					Inherited: Bool(false),
					CoreAttributes: &[]*pf.SpAdapterAttribute{
						{Name: String("subject")},
					},
					ExtendedAttributes: nil,
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			resourceSchema := resourcePingFederateSpAdapterResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateSpAdapterResourceReadResult(resourceLocalData, &tc.Resource, m)

			if got := *resourcePingFederateSpAdapterResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateSpAdapterResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
