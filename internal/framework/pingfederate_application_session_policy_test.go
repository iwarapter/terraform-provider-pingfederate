package framework

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateApplicationSessionPolicyResource(t *testing.T) {
	resourceName := "pingfederate_application_session_policy.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateApplicationSessionPolicyResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateApplicationSessionPolicyResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateApplicationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "10"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "12"),
				),
			},
			{
				Config: testAccPingFederateApplicationSessionPolicyResourceConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateApplicationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "15"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "20"),
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

func testAccCheckPingFederateApplicationSessionPolicyResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateApplicationSessionPolicyResourceConfig() string {
	return `
resource "pingfederate_application_session_policy" "example" {
  idle_timeout_mins = 10
  max_timeout_mins  = 12
}`
}

func testAccPingFederateApplicationSessionPolicyResourceConfigUpdate() string {
	return `
resource "pingfederate_application_session_policy" "example" {
  idle_timeout_mins = 15
  max_timeout_mins  = 20
}`
}

func testAccCheckPingFederateApplicationSessionPolicyResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := pfc.Session
		result, _, err := conn.GetApplicationPolicy()

		if err != nil {
			return fmt.Errorf("Error: ApplicationSessionPolicy (%s) not found", n)
		}
		if strconv.Itoa(*result.IdleTimeoutMins) != rs.Primary.Attributes["idle_timeout_mins"] {
			return fmt.Errorf("Error: ApplicationSessionPolicy response (%s) didnt match state (%s)", strconv.Itoa(*result.IdleTimeoutMins), rs.Primary.Attributes["idle_timeout_mins"])
		}

		return nil
	}
}

func Test_resourcePingFederateApplicationSessionPolicyResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.ApplicationSessionPolicy
	}{
		{
			Resource: pf.ApplicationSessionPolicy{
				IdleTimeoutMins: Int(5),
				MaxTimeoutMins:  Int(6),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateApplicationSessionPolicyResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenApplicationSessionPolicy(&tc.Resource)).HasError())

			check := ApplicationSessionPolicyData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandApplicationSessionPolicy(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
