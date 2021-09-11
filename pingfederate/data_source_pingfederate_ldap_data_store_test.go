package pingfederate

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateLdapDataStoreDataSource(t *testing.T) {
	re := regexp.MustCompile(`^((10)\.[0-9])`)
	if !re.MatchString(pfVersion) {
		t.Skipf("This test only runs against PingFederate 10.0 and above, not: %s", pfVersion)
	}
	resourceName := "data.pingfederate_ldap_data_store.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateLdapDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateLdapDataStoreDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "name"),
					resource.TestCheckResourceAttr(resourceName, "ldap_type", "PING_DIRECTORY"),
					resource.TestCheckResourceAttr(resourceName, "hostnames.0", "host.docker.internal:1636"),
					resource.TestCheckResourceAttr(resourceName, "user_dn", "foo"),
					resource.TestCheckResourceAttrSet(resourceName, "encrypted_password"),
					resource.TestCheckResourceAttr(resourceName, "bind_anonymously", "false"),
					resource.TestCheckResourceAttr(resourceName, "min_connections", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "2"),
				),
			},
		},
	})
}

func TestAccPingFederateLdapDataStoreDataSource_NotFound(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateLdapDataStoreDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccPingFederateLdapDataStoreDataSourceConfigNonExistent(),
				ExpectError: regexp.MustCompile(`unable to find ldap data store with name 'junk'`),
			},
		},
	})
}

func testAccCheckPingFederateLdapDataStoreDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateLdapDataStoreDataSourceConfig() string {
	return `
provider "pingfederate" {
  bypass_external_validation = true
}

resource "pingfederate_ldap_data_store" "example" {
	ldap_type        = "PING_DIRECTORY"
	hostnames        = ["host.docker.internal:1636"]
	user_dn          = "foo"
	password         = "secret"
	bind_anonymously = false
	min_connections  = 1
	max_connections  = 2
}

data "pingfederate_ldap_data_store" "test" {
	name = pingfederate_ldap_data_store.example.name
}`
}

func testAccPingFederateLdapDataStoreDataSourceConfigNonExistent() string {
	return `
data "pingfederate_ldap_data_store" "test" {
	name = "junk"
}`
}
