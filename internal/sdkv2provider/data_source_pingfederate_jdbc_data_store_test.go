package sdkv2provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateJdbcDataStoreDataSource(t *testing.T) {

	resourceName := "data.pingfederate_jdbc_data_store.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateJdbcDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateJdbcDataStoreDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "connection_url_tags.0.connection_url", "jdbc:hsqldb:mem:mymemdb"),
					resource.TestCheckResourceAttr(resourceName, "user_name", "test"),
					resource.TestCheckResourceAttrSet(resourceName, "encrypted_password"),
					resource.TestCheckResourceAttr(resourceName, "max_pool_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "min_pool_size", "10"),
					resource.TestCheckResourceAttr(resourceName, "mask_attribute_values", "false"),
				),
			},
		},
	})
}

func TestAccPingFederateJdbcDataStoreDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateJdbcDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingFederateJdbcDataStoreDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find jdbc data store with name 'junk'`),
			},
		},
	})
}

func testAccCheckPingFederateJdbcDataStoreDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateJdbcDataStoreDataSourceConfig() string {
	return `
provider "pingfederate" {
  bypass_external_validation = true
  alias                      = "bypass"
}

resource "pingfederate_jdbc_data_store" "example" {
  driver_class   = "org.hsqldb.jdbcDriver"
  user_name      = "test"
  password       = "example"
  max_pool_size  = 50
  connection_url = "jdbc:hsqldb:mem:mymemdb"
  connection_url_tags {
    connection_url = "jdbc:hsqldb:mem:mymemdb"
    default_source = true
  }
  provider = pingfederate.bypass
}

data "pingfederate_jdbc_data_store" "test" {
  name     = pingfederate_jdbc_data_store.example.name
  provider = pingfederate.bypass
}`
}

func testAccPingFederateJdbcDataStoreDataSourceConfigNonExistent() string {
	return `
data "pingfederate_jdbc_data_store" "test" {
  name = "junk"
}`
}
