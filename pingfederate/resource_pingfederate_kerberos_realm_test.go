package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/kerberosRealms"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateKerberosRealmResource(t *testing.T) {
	resourceName := "pingfederate_kerberos_realm.demo"

	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKerberosRealmResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKerberosRealmResourceConfig("bar.foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKerberosRealmResourceExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateKerberosRealmResourceConfig("foo.foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKerberosRealmResourceExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"kerberos_password",
					"kerberos_encrypted_password",
				},
			},
		},
	})
}

func testAccCheckPingFederateKerberosRealmResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKerberosRealmResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_kerberos_realm" "demo" {
	kerberos_realm_name      = "test"
	key_distribution_centers = ["foo.com", "%s"]
	kerberos_username 		 = "user"
	kerberos_password 		 = "secret"
}`, configUpdate)
}

func testAccCheckPingFederateKerberosRealmResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).KerberosRealms
		result, _, err := conn.GetKerberosRealm(&kerberosRealms.GetKerberosRealmInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: KerberosRealm (%s) not found", n)
		}

		if *result.KerberosRealmName != rs.Primary.Attributes["kerberos_realm_name"] {
			return fmt.Errorf("error: KerberosRealm response (%s) didnt match state (%s)", *result.KerberosRealmName, rs.Primary.Attributes["kerberos_realm_name"])
		}

		return nil
	}
}

func Test_resourcePingFederateKerberosRealmResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.KerberosRealm
	}{
		{
			Resource: pf.KerberosRealm{
				Id:                              String("foo"),
				KerberosEncryptedPassword:       String("foo"),
				KerberosPassword:                String("foo"),
				KerberosRealmName:               String("foo"),
				KerberosUsername:                String("foo"),
				KeyDistributionCenters:          &[]*string{String("foo")},
				SuppressDomainNameConcatenation: Bool(true),
			},
		},
		{
			Resource: pf.KerberosRealm{
				KerberosRealmName:               String("foo"),
				KerberosUsername:                String("foo"),
				SuppressDomainNameConcatenation: Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateKerberosRealmResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateKerberosRealmResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateKerberosRealmResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateKerberosRealmResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
