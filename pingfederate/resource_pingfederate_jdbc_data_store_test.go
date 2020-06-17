package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateJdbcDataStoreResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateJdbcDataStoreResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateJdbcDataStoreResourceConfig("1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateJdbcDataStoreResourceExists("pingfederate_jdbc_data_store.demo"),
					// testAccCheckPingFederateJdbcDataStoreResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederateJdbcDataStoreResourceConfig("1001"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateJdbcDataStoreResourceExists("pingfederate_jdbc_data_store.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateJdbcDataStoreResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateJdbcDataStoreResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`resource "pingfederate_jdbc_data_store" "demo" {
		name = "terraform"
		driver_class = "org.hsqldb.jdbcDriver"
		user_name = "sa"
		password = ""
		max_pool_size = %s
		connection_url = "jdbc:hsqldb:mem:mymemdb"
		connection_url_tags {
		  connection_url = "jdbc:hsqldb:mem:mymemdb"
		  default_source = true
		}
	}`, configUpdate)
}

func testAccCheckPingFederateJdbcDataStoreResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).DataStores
		result, _, err := conn.GetDataStore(&dataStores.GetDataStoreInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: DataStore (%s) not found", n)
		}

		if *result.DriverClass != rs.Primary.Attributes["driver_class"] {
			return fmt.Errorf("Error: DataStore response (%s) didnt match state (%s)", *result.DriverClass, rs.Primary.Attributes["driver_class"])
		}

		return nil
	}
}

func Test_resourcePingFederateJdbcDataStoreResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.JdbcDataStore
	}{
		{
			Resource: pf.JdbcDataStore{
				ConnectionUrl: String("jdbc:hsqldb:mem:mymemdb"),
				//ConnectionUrlTags: &[]*pf.JdbcTagConfig{},
				DriverClass:               String("org.hsqldb.jdbcDriver"),
				Name:                      String("terraform"),
				Password:                  String("foo"),
				UserName:                  String("sa"),
				AllowMultiValueAttributes: Bool(true),
				MaskAttributeValues:       Bool(true),
				ValidateConnectionSql:     String("select * from dual;"),
				Type:                      String("JDBC"),

				//TODO Only set because terraform sdk defaults int's to 0 :/
				BlockingTimeout:   Int(5),
				EncryptedPassword: String("foo"),
				IdleTimeout:       Int(5),
				MaxPoolSize:       Int(10),
				MinPoolSize:       Int(100),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateJdbcDataStoreResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateJdbcDataStoreResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateJdbcDataStoreResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateJdbcDataStoreResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
