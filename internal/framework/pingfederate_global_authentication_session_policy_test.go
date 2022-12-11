package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateGlobalAuthenticationSessionPolicyResource(t *testing.T) {
	resourceName := "pingfederate_global_authentication_session_policy.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateGlobalAuthenticationSessionPolicyResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateGlobalAuthenticationSessionPolicyResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateGlobalAuthenticationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_sessions", "true"),
					resource.TestCheckResourceAttr(resourceName, "hash_unique_user_key_attribute", "true"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_display_unit", "MINUTES"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "5"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_display_unit", "MINUTES"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "6"),
					resource.TestCheckResourceAttr(resourceName, "persistent_sessions", "true"),
				),
			},
			{
				Config: testAccPingFederateGlobalAuthenticationSessionPolicyResourceConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateGlobalAuthenticationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "enable_sessions", "false"),
					resource.TestCheckResourceAttr(resourceName, "hash_unique_user_key_attribute", "false"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_display_unit", "HOURS"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "60"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_display_unit", "HOURS"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "120"),
					resource.TestCheckResourceAttr(resourceName, "persistent_sessions", "false"),
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

func testAccCheckPingFederateGlobalAuthenticationSessionPolicyResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateGlobalAuthenticationSessionPolicyResourceConfig() string {
	return `
resource "pingfederate_global_authentication_session_policy" "example" {
  enable_sessions                = true
  hash_unique_user_key_attribute = true
  idle_timeout_display_unit      = "MINUTES"
  idle_timeout_mins              = 5
  max_timeout_display_unit       = "MINUTES"
  max_timeout_mins               = 6
  persistent_sessions            = true
}`
}

func testAccPingFederateGlobalAuthenticationSessionPolicyResourceConfigUpdate() string {
	return `
resource "pingfederate_global_authentication_session_policy" "example" {
  enable_sessions                = false
  hash_unique_user_key_attribute = false
  idle_timeout_display_unit      = "HOURS"
  idle_timeout_mins              = 60
  max_timeout_display_unit       = "HOURS"
  max_timeout_mins               = 120
  persistent_sessions            = false
}`
}

func testAccCheckPingFederateGlobalAuthenticationSessionPolicyResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := pfc.Session
		result, _, err := conn.GetGlobalPolicy()

		if err != nil {
			return fmt.Errorf("Error: GlobalAuthenticationSessionPolicy (%s) not found", n)
		}
		if *result.IdleTimeoutDisplayUnit != rs.Primary.Attributes["idle_timeout_display_unit"] {
			return fmt.Errorf("Error: GlobalAuthenticationSessionPolicy response (%s) didnt match state (%s)", *result.IdleTimeoutDisplayUnit, rs.Primary.Attributes["idle_timeout_display_unit"])
		}

		return nil
	}
}

func Test_resourcePingFederateGlobalAuthenticationSessionPolicyResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.GlobalAuthenticationSessionPolicy
	}{
		{
			Resource: pf.GlobalAuthenticationSessionPolicy{
				EnableSessions:             Bool(true),
				HashUniqueUserKeyAttribute: Bool(true),
				IdleTimeoutDisplayUnit:     String("IdleTimeoutDisplayUnit"),
				IdleTimeoutMins:            Int(5),
				MaxTimeoutDisplayUnit:      String("MaxTimeoutDisplayUnit"),
				MaxTimeoutMins:             Int(6),
				PersistentSessions:         Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateGlobalAuthenticationSessionPolicyResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenGlobalAuthenticationSessionPolicy(&tc.Resource)).HasError())

			check := GlobalAuthenticationSessionPolicyData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandGlobalAuthenticationSessionPolicy(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
