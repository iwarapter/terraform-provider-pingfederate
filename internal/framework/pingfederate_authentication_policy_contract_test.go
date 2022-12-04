package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("authentication_policy_contract", &resource.Sweeper{
		Name:         "authentication_policy_contract",
		Dependencies: []string{},
		F: func(r string) error {
			results, _, err := pfc.AuthenticationPolicyContracts.GetAuthenticationPolicyContracts(&authenticationPolicyContracts.GetAuthenticationPolicyContractsInput{Filter: "acc_test"})
			if err != nil {
				return fmt.Errorf("unable to list authentication policy contracts %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := pfc.AuthenticationPolicyContracts.DeleteAuthenticationPolicyContract(&authenticationPolicyContracts.DeleteAuthenticationPolicyContractInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep authentication policy contract %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateAuthenticationPolicyContractResource(t *testing.T) {
	resourceName := "pingfederate_authentication_policy_contract.demo"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateAuthenticationPolicyContractResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationPolicyContractResourceConfig("email"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.0", "email"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.1", "foo"),
				),
			},
			{
				Config: testAccPingFederateAuthenticationPolicyContractResourceConfig("address"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_one"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.0", "address"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.1", "foo"),
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

func TestAccPingFederateAuthenticationPolicyContractResourceSdkUpgrade(t *testing.T) {
	resourceName := "pingfederate_authentication_policy_contract.demo"
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"pingfederate": {
						VersionConstraint: "0.0.24",
						Source:            "iwarapter/pingfederate",
					},
				},
				Config: `
resource "pingfederate_authentication_policy_contract" "demo" {
  policy_contract_id  = "acc_test_upgrade"
  name                = "acc_test_upgrade"
  extended_attributes = ["foo", "email"]
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "policy_contract_id", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.0", "email"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.1", "foo"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PlanOnly:                 true,
				Config: `
resource "pingfederate_authentication_policy_contract" "demo" {
  id                  = "acc_test_upgrade"
  name                = "acc_test_upgrade"
  extended_attributes = ["foo", "email"]
}`,
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: `
resource "pingfederate_authentication_policy_contract" "demo" {
  id                  = "acc_test_upgrade"
  name                = "acc_test_upgrade"
  extended_attributes = ["foo", "email"]
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationPolicyContractResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "id", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.0", "email"),
					resource.TestCheckResourceAttr(resourceName, "extended_attributes.1", "foo"),
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
  name                = "acc_test_one"
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

		conn := pfc.AuthenticationPolicyContracts
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
					{Name: String("subject")},
				},
				ExtendedAttributes: &[]*pf.AuthenticationPolicyContractAttribute{
					{Name: String("bar")},
					{Name: String("foo")},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateAuthenticationPolicyContractResource{}
			ctx := context.Background()
			resourceSchema, diags := res.GetSchema(ctx)
			require.False(t, diags.HasError())

			state := tfsdk.State{Schema: resourceSchema}
			require.False(t, state.Set(ctx, flattenAuthenticationPolicyContract(&tc.Resource)).HasError())

			check := AuthenticationPolicyContractData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandAuthenticationPolicyContract(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
