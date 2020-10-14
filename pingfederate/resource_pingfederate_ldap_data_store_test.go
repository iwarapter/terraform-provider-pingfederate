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

func TestAccPingFederateLdapDataStoreResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateLdapDataStoreResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateLdapDataStoreResourceConfig("1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateLdapDataStoreResourceExists("pingfederate_ldap_data_store.demo"),
					// testAccCheckPingFederateLdapDataStoreResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederateLdapDataStoreResourceConfig("1001"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateLdapDataStoreResourceExists("pingfederate_ldap_data_store.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateLdapDataStoreResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateLdapDataStoreResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_ldap_data_store" "demo" {
	bypass_external_validation = true
	name      = "terraform_ldap"
	ldap_type = "PING_DIRECTORY"
	hostnames = ["host.docker.internal:1389"]
	bind_anonymously = true
	min_connections = 1
	max_connections = %s
}`, configUpdate)
}

func testAccCheckPingFederateLdapDataStoreResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).DataStores
		result, _, err := conn.GetLdapDataStore(&dataStores.GetLdapDataStoreInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: DataStore (%s) not found", n)
		}

		if *result.LdapType != rs.Primary.Attributes["ldap_type"] {
			return fmt.Errorf("Error: DataStore response (%s) didnt match state (%s)", *result.LdapType, rs.Primary.Attributes["ldap_type"])
		}

		return nil
	}
}

func Test_resourcePingFederateLdapDataStoreResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.LdapDataStore
	}{
		{
			Resource: pf.LdapDataStore{
				BindAnonymously:   Bool(false),
				ConnectionTimeout: Int(3000),
				CreateIfNecessary: Bool(false),
				DnsTtl:            Int(60000),
				//TODO Only set because terraform sdk defaults int's to 0 :/
				EncryptedPassword:    String("foo"),
				FollowLDAPReferrals:  Bool(false),
				LdapDnsSrvPrefix:     String("_ldap._tcp"),
				LdapsDnsSrvPrefix:    String("_ldaps._tcp"),
				MaskAttributeValues:  Bool(false),
				MaxWait:              Int(-1),
				Name:                 String("terraform"),
				Password:             String("foo"),
				ReadTimeout:          Int(3000),
				TestOnBorrow:         Bool(false),
				TestOnReturn:         Bool(false),
				TimeBetweenEvictions: Int(60000),
				Type:                 String("LDAP"),
				UseDnsSrvRecords:     Bool(false),
				UseSsl:               Bool(false),
				VerifyHost:           Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateLdapDataStoreResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateLdapDataStoreResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateLdapDataStoreResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateLdapDataStoreResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
