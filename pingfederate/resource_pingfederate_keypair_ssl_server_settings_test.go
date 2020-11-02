package pingfederate

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateKeypairSslServerSettingsResource(t *testing.T) {
	svc := keyPairsSslServer.New(cfg)
	settings, _, err := svc.GetSettings()
	if err != nil {
		t.Fatalf("unable to retrieve ssl server settings")
	}
	resourceName := "pingfederate_keypair_ssl_server_settings.test"
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateKeypairSslServerSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateKeypairSslServerSettingsResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairSslServerSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "admin_server_cert"),
					resource.TestCheckResourceAttrSet(resourceName, "runtime_server_cert"),
					resource.TestCheckResourceAttrSet(resourceName, "active_runtime_server_certs.0"),
					resource.TestCheckResourceAttrSet(resourceName, "active_admin_server_certs.0"),
				),
			},
			{
				Config: testAccPingFederateKeypairSslServerSettingsResourceResetConfig(*settings.AdminConsoleCertRef.Id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateKeypairSslServerSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "admin_server_cert", *settings.AdminConsoleCertRef.Id),
					resource.TestCheckResourceAttr(resourceName, "runtime_server_cert", *settings.RuntimeServerCertRef.Id),
					resource.TestCheckResourceAttr(resourceName, "active_runtime_server_certs.0", *settings.AdminConsoleCertRef.Id),
					resource.TestCheckResourceAttr(resourceName, "active_admin_server_certs.0", *settings.RuntimeServerCertRef.Id),
				),
			},
		},
	})
}

func testAccCheckPingFederateKeypairSslServerSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateKeypairSslServerSettingsResourceConfig() string {
	return `
resource "pingfederate_keypair_ssl_server" "demo" {
	city = "Test"
	common_name = "Settings Test"
	country = "GB"
	key_algorithm = "RSA"
	key_size = 2048
	organization = "Test"
	organization_unit = "Test"
	state = "Test"
	valid_days = 365
}

resource "pingfederate_keypair_ssl_server_settings" "test" {
	admin_server_cert = pingfederate_keypair_ssl_server.demo.id
	runtime_server_cert = pingfederate_keypair_ssl_server.demo.id
	active_runtime_server_certs = [pingfederate_keypair_ssl_server.demo.id]
	active_admin_server_certs = [pingfederate_keypair_ssl_server.demo.id]
}`
}

func testAccPingFederateKeypairSslServerSettingsResourceResetConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_keypair_ssl_server" "demo" {
	city = "Test"
	common_name = "Settings Test"
	country = "GB"
	key_algorithm = "RSA"
	key_size = 2048
	organization = "Test"
	organization_unit = "Test"
	state = "Test"
	valid_days = 365
}

resource "pingfederate_keypair_ssl_server_settings" "test" {
	admin_server_cert = "%s"
	runtime_server_cert = "%s"
	active_runtime_server_certs = ["%s"]
	active_admin_server_certs = ["%s"]
}`, configUpdate, configUpdate, configUpdate, configUpdate)
}

func testAccCheckPingFederateKeypairSslServerSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}
		conn := testAccProvider.Meta().(pfClient).KeyPairsSslServer
		result, _, err := conn.GetSettings()

		if err != nil {
			return fmt.Errorf("error: SslServerSettings (%s) not found", n)
		}
		if *result.AdminConsoleCertRef.Id != rs.Primary.Attributes["admin_server_cert"] {
			return fmt.Errorf("error: SslServerSettings response (%s) didnt match state (%s)", *result.AdminConsoleCertRef.Id, rs.Primary.Attributes["admin_server_cert"])
		}
		return nil
	}
}

func Test_resourcePingFederateKeypairSslServerSettingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.SslServerSettings
	}{
		{
			Resource: pf.SslServerSettings{
				ActiveAdminConsoleCerts:  &[]*pf.ResourceLink{{Id: String("xyz")}, {Id: String("zyx")}},
				ActiveRuntimeServerCerts: &[]*pf.ResourceLink{{Id: String("def")}, {Id: String("abc")}},
				AdminConsoleCertRef:      &pf.ResourceLink{Id: String("foo")},
				RuntimeServerCertRef:     &pf.ResourceLink{Id: String("bar")},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateKeypairSslServerSettingsResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateKeypairSslServerSettingsResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateKeypairSslServerSettingsResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateKeypairSslServerSettingsResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
