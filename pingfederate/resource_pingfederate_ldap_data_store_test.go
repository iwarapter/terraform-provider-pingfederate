package pingfederate

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("ldap_data_store", &resource.Sweeper{
		Name:         "ldap_data_store",
		Dependencies: []string{"idp_sp_connection"},
		F: func(r string) error {
			svc := dataStores.New(cfg)
			results, _, err := svc.GetDataStores()
			if err != nil {
				return fmt.Errorf("unable to list data stores %s", err)
			}
			for _, item := range *results.Items {
				switch v := item.(type) {
				case *pf.LdapDataStore:
					_, _, err := svc.DeleteDataStore(&dataStores.DeleteDataStoreInput{Id: *v.Id})
					if err != nil {
						return fmt.Errorf("unable to sweep data store %s because %s", *v.Id, err)
					}
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateLdapDataStoreResource(t *testing.T) {
	re := regexp.MustCompile(`^((10|11)\.[0-9])`)
	if !re.MatchString(pfVersion) {
		t.Skipf("This test only runs against PingFederate 10.0 and above, not: %s", pfVersion)
	}
	resourceName := "pingfederate_ldap_data_store.demo"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateLdapDataStoreResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateLdapDataStoreResourceConfig("1000"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateLdapDataStoreResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "terraform_ldap"),
					resource.TestCheckResourceAttr(resourceName, "ldap_type", "PING_DIRECTORY"),
					resource.TestCheckResourceAttr(resourceName, "hostnames.0", "host.docker.internal:1389"),
					resource.TestCheckResourceAttr(resourceName, "user_dn", "test"),
					resource.TestCheckResourceAttr(resourceName, "password", "secret"),
					resource.TestCheckResourceAttrSet(resourceName, "encrypted_password"),
					resource.TestCheckResourceAttr(resourceName, "bind_anonymously", "false"),
					resource.TestCheckResourceAttr(resourceName, "min_connections", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "1000"),
				),
			},
			{
				Config: testAccPingFederateLdapDataStoreResourceConfig("1001"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateLdapDataStoreResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "terraform_ldap"),
					resource.TestCheckResourceAttr(resourceName, "ldap_type", "PING_DIRECTORY"),
					resource.TestCheckResourceAttr(resourceName, "hostnames.0", "host.docker.internal:1389"),
					resource.TestCheckResourceAttr(resourceName, "user_dn", "test"),
					resource.TestCheckResourceAttr(resourceName, "password", "secret"),
					resource.TestCheckResourceAttrSet(resourceName, "encrypted_password"),
					resource.TestCheckResourceAttr(resourceName, "bind_anonymously", "false"),
					resource.TestCheckResourceAttr(resourceName, "min_connections", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_connections", "1001"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"password",
					"encrypted_password",
				},
			},
		},
	})
}

func testAccCheckPingFederateLdapDataStoreResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateLdapDataStoreResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
provider "pingfederate" {
  bypass_external_validation = true
}

resource "pingfederate_ldap_data_store" "demo" {
    data_store_id = "ldapexample"
	name             = "terraform_ldap"
	ldap_type        = "PING_DIRECTORY"
	hostnames        = ["host.docker.internal:1389"]
	user_dn          = "test"
	password         = "secret"
	bind_anonymously = false
	min_connections  = 1
	max_connections  = %s
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
				Id:                String("example"),
				BindAnonymously:   Bool(false),
				ConnectionTimeout: Int(3000),
				CreateIfNecessary: Bool(false),
				DnsTtl:            Int(60000),
				//TODO Only set because terraform sdk defaults int's to 0 :/
				HostnamesTags: &[]*pf.LdapTagConfig{
					{
						DefaultSource: Bool(true),
						Hostnames:     &[]*string{String("ldaps://foo")},
						Tags:          String("tags"),
					},
				},
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

			assert.Equal(t, tc.Resource, *resourcePingFederateLdapDataStoreResourceReadData(resourceLocalData))
		})
	}
}
