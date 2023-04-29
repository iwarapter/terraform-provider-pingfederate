package framework

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccPingFederateJdbcDataStoresDatasource(t *testing.T) {
	resourceName := "data.pingfederate_jdbc_data_stores.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateJdbcDataStoresDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "items.0.id"),
					resource.TestCheckResourceAttr(resourceName, "items.0.type", "JDBC"),
				),
			},
		},
	})
}

func testAccPingFederateJdbcDataStoresDatasourceConfig() string {
	return `data "pingfederate_jdbc_data_stores" "test" {}`
}
