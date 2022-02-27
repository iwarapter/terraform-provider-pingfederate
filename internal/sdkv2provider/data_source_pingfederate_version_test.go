package sdkv2provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateVersionDataSource(t *testing.T) {
	resourceName := "data.pingfederate_version.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateVersionDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateVersionDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "version"),
				),
			},
		},
	})
}

func testAccCheckPingFederateVersionDataSourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateVersionDataSourceConfig() string {
	return `data "pingfederate_version" "test" {}`
}
