package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingFederateLdapDataStoresDatasource(t *testing.T) {
	resourceName := "data.pingfederate_ldap_data_stores.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateLdapDataStoresDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "items.0.type", "LDAP"),
				),
			},
		},
	})
}

func testAccPingFederateLdapDataStoresDatasourceConfig() string {
	return `
provider "pingfederate" {
  bypass_external_validation = true
}

resource "pingfederate_ldap_data_store" "demo_ldap" {
  name             = "acc_test_datasource"
  ldap_type        = "PING_DIRECTORY"
  hostnames        = ["host.docker.internal:1389"]
  bind_anonymously = true
  min_connections  = 1
  max_connections  = 1
}

data "pingfederate_ldap_data_stores" "test" {
  depends_on = [pingfederate_ldap_data_store.demo_ldap]
}
`
}
