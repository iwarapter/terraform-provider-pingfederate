package pingfederate

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("custom_data_store", &resource.Sweeper{
		Name:         "custom_data_store",
		Dependencies: []string{},
		F: func(r string) error {
			svc := dataStores.New(cfg)
			results, _, err := svc.GetDataStores()
			if err != nil {
				return fmt.Errorf("unable to list data stores %s", err)
			}
			for _, item := range *results.Items {
				switch v := item.(type) {
				case pf.CustomDataStore:
					_, _, err := svc.DeleteDataStore(&dataStores.DeleteDataStoreInput{Id: *v.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep data store %s because %s", *v.Id, err)
					}
				}

			}
			return nil
		},
	})
}

func TestAccPingFederateCustomDataStoreResource(t *testing.T) {
	resourceName := "pingfederate_custom_data_store.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateCustomDataStoreResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateCustomDataStoreResourceConfig("https://example.foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateCustomDataStoreResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.name", "Base URLs and Tags"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.name", "Base URL"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.value", "https://example.foo"),
				),
			},
			{
				Config: testAccPingFederateCustomDataStoreResourceConfig("https://example.bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateCustomDataStoreResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.name", "Base URLs and Tags"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.name", "Base URL"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.value", "https://example.bar"),
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

func testAccCheckPingFederateCustomDataStoreResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateCustomDataStoreResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
provider "pingfederate" {
  bypass_external_validation = true
}
data "pingfederate_version" "instance" {}

locals {
  isSupported = length(regexall("(11).[0-9]", data.pingfederate_version.instance.version)) > 0
}

resource "pingfederate_custom_data_store" "demo" {
  data_store_id = "customexample"
  name = "customterra"
  plugin_descriptor_ref {
	id = "com.pingidentity.pf.datastore.other.RestDataSourceDriver"
  }
  configuration {
	tables {
	  name = "Base URLs and Tags"
	  rows {
		fields {
		  name = "Base URL"
		  value = "%s"
		}
		fields {
		  name = "Tags"
		  value = ""
		}
		default_row = true
	  }
	}
	tables {
	  name = "HTTP Request Headers"
	}
	tables {
	  name = "Attributes"
	  rows {
		fields {
		  name = "Local Attribute"
		  value = "foo"
		}
		fields {
		  name = "JSON Response Attribute Path"
		  value = "/bar"
		}
	  }
	}
	fields {
	  name = "Authentication Method"
	  value = "None"
	}
	fields {
	  name = "Username"
	  value = ""
	}
	fields {
	  name = "Password"
	  value = ""
	}
	fields {
	  name = "OAuth Token Endpoint"
	  value = ""
	}
	fields {
	  name = "OAuth Scope"
	  value = ""
	}
	fields {
	  name = "Client ID"
	  value = ""
	}
	fields {
	  name = "Client Secret"
	  value = ""
	}
	fields {
	  name = "Enable HTTPS Hostname Verification"
	  value = "true"
	}
	fields {
	  name = "Read Timeout (ms)"
	  value = "10000"
	}
	fields {
	  name = "Connection Timeout (ms)"
	  value = "10000"
	}
	fields {
	  name = "Max Payload Size (KB)"
	  value = "1024"
	}
	fields {
	  name = "Retry Request"
	  value = "true"
	}
	fields {
	  name = "Maximum Retries Limit"
	  value = "5"
	}
	fields {
	  name = "Retry Error Codes"
	  value = "429"
	}
	fields {
	  name = "Test Connection URL"
	}
	dynamic "fields" {
      for_each = local.isSupported ? [1] : []
      content {
        name  = "Client Secret Reference"
        value = null
      }
    }
	dynamic "fields" {
      for_each = local.isSupported ? [1] : []
      content {
        name  = "HTTP Method"
        value = "GET"
      }
    }
	dynamic "fields" {
      for_each = local.isSupported ? [1] : []
      content {
        name  = "Password Reference"
        value = null
      }
    }
	dynamic "fields" {
      for_each = local.isSupported ? [1] : []
      content {
        name  = "Test Connection Body"
        value = null
      }
    }
  }
}`, configUpdate)
}

func testAccCheckPingFederateCustomDataStoreResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).DataStores
		result, _, err := conn.GetDataStore(&dataStores.GetDataStoreInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: DataStore (%s) not found", n)
		}

		if *result.PluginDescriptorRef.Id != rs.Primary.Attributes["plugin_descriptor_ref.0.id"] {
			return fmt.Errorf("error: DataStore response (%s) didnt match state (%s)", *result.PluginDescriptorRef.Id, rs.Primary.Attributes["plugin_descriptor_ref.0.id"])
		}

		return nil
	}
}

type dataStoresMock struct {
	dataStores.DataStoresAPI
}

func (s dataStoresMock) GetCustomDataStoreDescriptor(_ *dataStores.GetCustomDataStoreDescriptorInput) (output *pf.CustomDataStoreDescriptor, resp *http.Response, err error) {
	return &pf.CustomDataStoreDescriptor{
		AttributeContract: nil,
		ClassName:         String("com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
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
		Id:                       String("com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
		Name:                     String("REST API"),
		SupportsExtendedContract: nil,
	}, nil, nil
}

func Test_resourcePingFederateCustomDataStoreResourceReadData(t *testing.T) {
	svc := &dataStoresMock{}
	cases := []struct {
		Resource pf.CustomDataStore
	}{
		{
			Resource: pf.CustomDataStore{
				Id:                  String("example"),
				MaskAttributeValues: Bool(false),
				Name:                String("terraform"),
				Type:                String("CUSTOM"),
				PluginDescriptorRef: &pf.ResourceLink{
					Id: String("com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
				},
				ParentRef: &pf.ResourceLink{
					Id: String("com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
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
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateCustomDataStoreResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateCustomDataStoreResourceReadResult(resourceLocalData, &tc.Resource, svc)

			assert.Equal(t, tc.Resource, *resourcePingFederateCustomDataStoreResourceReadData(resourceLocalData))
		})
	}
}
