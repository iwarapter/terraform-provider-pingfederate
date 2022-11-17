package sdkv2provider

//lint:file-ignore U1000 Ignore report - no current alternative

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/pingOneConnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("pingone_connection", &resource.Sweeper{
		Name:         "pingone_connection",
		Dependencies: []string{},
		F: func(r string) error {
			svc := pingOneConnections.New(cfg)
			results, _, err := svc.GetPingOneConnections()
			if err != nil {
				return fmt.Errorf("unable to list pingone connection %s", err)
			}
			for _, item := range *results.Items {
				if strings.Contains(*item.Name, "acctest") {
					_, _, err := svc.DeletePingOneConnection(&pingOneConnections.DeletePingOneConnectionInput{Id: *item.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep pingone connection %s because %s", *item.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederatePingOneConnectionResource(t *testing.T) {
	t.Skipf("unable to create without valid pingone credential")
	resourceName := "pingfederate_pingone_connection.demo"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederatePingOneConnectionResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederatePingOneConnectionResourceConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederatePingOneConnectionResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "active", "true"),
				),
			},
			{
				Config: testAccPingFederatePingOneConnectionResourceConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederatePingOneConnectionResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "active", "false"),
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

func testAccCheckPingFederatePingOneConnectionResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederatePingOneConnectionResourceConfig(config string) string {
	return fmt.Sprintf(`
resource "pingfederate_pingone_connection" "demo" {
  name       = "acctestfoo"
  active     = %s
  credential = "secret"
}`, config)
}

func testAccCheckPingFederatePingOneConnectionResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).PingOneConnections
		result, _, err := conn.GetPingOneConnection(&pingOneConnections.GetPingOneConnectionInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: PingOneConnection (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: PingOneConnection response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederatePingOneConnectionResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.PingOneConnection
	}{
		{
			Resource: pf.PingOneConnection{
				Name:       String("foo"),
				Active:     Bool(true),
				Credential: String("secret"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederatePingOneConnectionResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederatePingOneConnectionResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederatePingOneConnectionResourceReadData(resourceLocalData))
		})
	}
}
