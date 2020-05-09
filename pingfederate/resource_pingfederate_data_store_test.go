package pingfederate

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func TestAccPingFederateDataStoreResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateDataStoreResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateDataStoreResourceConfig("1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateDataStoreResourceExists("pingfederate_data_store.demo"),
					// testAccCheckPingFederateDataStoreResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederateDataStoreResourceConfig("2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateDataStoreResourceExists("pingfederate_data_store.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateDataStoreResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateDataStoreResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`resource "pingfederate_data_store" "demo" {
	  jdbc_data_store {
		name = "terraform"
		driver_class = "org.hsqldb.jdbcDriver"
		user_name = "sa"
		password = ""
		max_pool_size = %s
		connection_url = "jdbc:hsqldb:mem:mymemdb"
	  }
	}`, configUpdate)
}

func testAccCheckPingFederateDataStoreResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(*pf.PfClient).DataStores
		result, _, err := conn.GetDataStore(&pf.GetDataStoreInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: DataStore (%s) not found", n)
		}

		if *result.DriverClass != rs.Primary.Attributes["jdbc_data_store.0.driver_class"] {
			return fmt.Errorf("Error: DataStore response (%s) didnt match state (%s)", *result.DriverClass, rs.Primary.Attributes["jdbc_data_store.0.driver_class"])
		}

		return nil
	}
}

func Test_resourcePingFederateDataStoreResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.DataStore
	}{
		{
			Resource: pf.DataStore{
				JdbcDataStore: pf.JdbcDataStore{
					ConnectionUrl:             String("jdbc:hsqldb:mem:mymemdb"),
					DriverClass:               String("org.hsqldb.jdbcDriver"),
					Name:                      String("terraform"),
					Password:                  String(""),
					UserName:                  String("sa"),
					AllowMultiValueAttributes: Bool(false),
					MaskAttributeValues:       Bool(false),
					ValidateConnectionSql:     String(""),
					Type:                      String("JDBC"),

					//TODO Only set because terraform sdk defaults int's to 0 :/
					BlockingTimeout:   Int(0),
					EncryptedPassword: String(""),
					IdleTimeout:       Int(0),
					MaxPoolSize:       Int(0),
					MinPoolSize:       Int(0),
				},
				Type: String("JDBC"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateDataStoreResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateDataStoreResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateDataStoreResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateDataStoreResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
