package sdkv2provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateCustomDataStoreDataSource(t *testing.T) {

	resourceName := "data.pingfederate_custom_data_store.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateCustomDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateCustomDataStoreDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "plugin_descriptor_ref.0.id", "com.pingidentity.pf.datastore.other.RestDataSourceDriver"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.name", "Base URLs and Tags"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.name", "Base URL"),
					resource.TestCheckResourceAttr(resourceName, "configuration.0.tables.0.rows.0.fields.0.value", "https://example.com"),
				),
			},
		},
	})
}

func TestAccPingFederateCustomDataStoreDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateCustomDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingFederateCustomDataStoreDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find custom data store with name 'junk'`),
			},
		},
	})
}

func testAccCheckPingFederateCustomDataStoreDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateCustomDataStoreDataSourceConfig() string {
	return `
provider "pingfederate" {
  bypass_external_validation = true
}
data "pingfederate_version" "instance" {}

locals {
  isSupported = length(regexall("(11).[0-9]", data.pingfederate_version.instance.version)) > 0
}

data "pingfederate_custom_data_store" "test" {
  name = pingfederate_custom_data_store.example.name
}

resource "pingfederate_custom_data_store" "example" {
  name = "customterraform"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.datastore.other.RestDataSourceDriver"
  }
  configuration {
    tables {
      name = "Base URLs and Tags"
      rows {
        fields {
          name  = "Base URL"
          value = "https://example.com"
        }
        fields {
          name  = "Tags"
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
          name  = "Local Attribute"
          value = "foo"
        }
        fields {
          name  = "JSON Response Attribute Path"
          value = "/bar"
        }
      }
    }
    fields {
      name  = "Authentication Method"
      value = "None"
    }
    fields {
      name  = "Username"
      value = ""
    }
    fields {
      name  = "Password"
      value = ""
    }
    fields {
      name  = "OAuth Token Endpoint"
      value = ""
    }
    fields {
      name  = "OAuth Scope"
      value = ""
    }
    fields {
      name  = "Client ID"
      value = ""
    }
    fields {
      name  = "Client Secret"
      value = ""
    }
    fields {
      name  = "Enable HTTPS Hostname Verification"
      value = "true"
    }
    fields {
      name  = "Read Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Connection Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Max Payload Size (KB)"
      value = "1024"
    }
    fields {
      name  = "Retry Request"
      value = "true"
    }
    fields {
      name  = "Maximum Retries Limit"
      value = "5"
    }
    fields {
      name  = "Retry Error Codes"
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
}`
}

func testAccPingFederateCustomDataStoreDataSourceConfigNonExistent() string {
	return `
data "pingfederate_custom_data_store" "test" {
  name = "junk"
}`
}
