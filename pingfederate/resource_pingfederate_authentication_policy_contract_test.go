package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateAuthenticationPolicyContractResource(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateAuthenticationPolicyContractResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPolicyContractResourceConfig("email"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists("pingfederate_authentication_policy_contract.demo"),
					// testAccCheckPingFederateAuthenticationPolicyContractResourceAttributes(),
				),
			},
			{
				Config: testAccPingFederateAuthenticationPolicyContractResourceConfig("address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists("pingfederate_authentication_policy_contract.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateAuthenticationPolicyContractResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthenticationPolicyContractResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_authentication_policy_contract" "demo" {
		name = "testing"
		core_attributes = ["subject"]
		extended_attributes = ["foo", "%s"]
	}`, configUpdate)
}

func testAccCheckPingFederateAuthenticationPolicyContractResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).AuthenticationPolicyContracts
		result, _, err := conn.GetAuthenticationPolicyContract(&authenticationPolicyContracts.GetAuthenticationPolicyContractInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: AuthenticationPolicyContract (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: AuthenticationPolicyContract response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateAuthenticationPolicyContractResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AuthenticationPolicyContract
	}{
		{
			Resource: pf.AuthenticationPolicyContract{
				Name: String("foo"),
				CoreAttributes: &[]*pf.AuthenticationPolicyContractAttribute{
					&pf.AuthenticationPolicyContractAttribute{
						Name: String("subject"),
					},
				},
				ExtendedAttributes: &[]*pf.AuthenticationPolicyContractAttribute{
					&pf.AuthenticationPolicyContractAttribute{
						Name: String("bar"),
					},
					&pf.AuthenticationPolicyContractAttribute{
						Name: String("foo"),
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateAuthenticationPolicyContractResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateAuthenticationPolicyContractResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateAuthenticationPolicyContractResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateAuthenticationPolicyContractResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
