package pingfederate
//
//import (
//	"fmt"
//	"testing"
//
//	"github.com/hashicorp/terraform/helper/resource"
//	"github.com/hashicorp/terraform/terraform"
//	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
//)
//
//func TestAccPingFederateIdpAdapter(t *testing.T) {
//	var out pf.Client
//
//	resource.ParallelTest(t, resource.TestCase{
//		// PreCheck:     func() { testAccPreCheck(t) },
//		Providers:    testAccProviders,
//		CheckDestroy: testAccCheckPingFederateIdpAdapterDestroy,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccPingFederateIdpAdapterConfig("acc", "120"),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckPingFederateIdpAdapterExists("pingfederate_idp_adapter.my_atm", &out),
//					// testAccCheckPingFederateIdpAdapterAttributes(),
//				),
//			},
//			{
//				Config: testAccPingFederateIdpAdapterConfig("acc", "180"),
//				Check: resource.ComposeTestCheckFunc(
//					testAccCheckPingFederateIdpAdapterExists("pingfederate_idp_adapter.my_atm", &out),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckPingFederateIdpAdapterDestroy(s *terraform.State) error {
//	return nil
//}
//
//func testAccPingFederateIdpAdapterConfig(name, configUpdate string) string {
//	return fmt.Sprintf(`
//	resource "pingfederate_idp_adapter" "my_atm" {
//		name = "%s"
//		plugin_descriptor_ref {
//			id = "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"
//		}
//	
//		configuration {
//			tables {
//				name = "Credential Validators"
//				rows {
//					fields {
//						name = "Password Credential Validator Instance"
//						value = "foo"
//					}
//				}
//			}
//			fields {
//				name  = "Realm"
//				value = "foo"
//			}
//	
//			fields {
//				name  = "Challenge Retries"
//				value = "3"
//			}
//
//		}
//	
//		attribute_contract {
//			core_attributes = []
//			extended_attributes = ["sub"]
//		}
//	}`, configUpdate)
//}
//
//func testAccCheckPingFederateIdpAdapterExists(n string, out *pf.Client) resource.TestCheckFunc {
//	return func(s *terraform.State) error {
//		rs, ok := s.RootModule().Resources[n]
//		if !ok {
//			return fmt.Errorf("Not found: %s", n)
//		}
//
//		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
//			return fmt.Errorf("No rule ID is set")
//		}
//
//		//conn := testAccProvider.Meta().(*pf.PfClient).IdpAdapters
//		//result, _, err := conn.GetTokenManager(&pf.GetTokenManagerInput{Id: rs.Primary.ID})
//		//
//		//if err != nil {
//		//	return fmt.Errorf("Error: IdpAdapter (%s) not found", n)
//		//}
//		//
//		//if *result.Name != rs.Primary.Attributes["name"] {
//		//	return fmt.Errorf("Error: IdpAdapter response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
//		//}
//
//		return nil
//	}
//}
