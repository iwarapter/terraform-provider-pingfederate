package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/session"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("authentication_session_policy", &resource.Sweeper{
		Name:         "authentication_session_policy",
		Dependencies: []string{},
		F: func(r string) error {
			results, _, err := pfc.Session.GetSourcePolicies()
			if err != nil {
				return fmt.Errorf("unable to list authentication policy contracts %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := pfc.Session.DeleteSourcePolicy(&session.DeleteSourcePolicyInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep authentication policy contract %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateAuthenticationSessionPolicyResource(t *testing.T) {
	resourceName := "pingfederate_authentication_session_policy.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateAuthenticationSessionPolicyResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateAuthenticationSessionPolicyResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "id", "acc_test1"),
					resource.TestCheckResourceAttr(resourceName, "authentication_source.type", "IDP_ADAPTER"),
					resource.TestCheckResourceAttr(resourceName, "authentication_source.source_ref", "idptestme"),
					resource.TestCheckResourceAttr(resourceName, "enable_sessions", "true"),
					resource.TestCheckResourceAttr(resourceName, "persistent", "true"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "300"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "300"),
					resource.TestCheckResourceAttr(resourceName, "timeout_display_unit", "HOURS"),
					resource.TestCheckResourceAttr(resourceName, "authn_context_sensitive", "true"),
				),
			},
			{
				Config: testAccPingFederateAuthenticationSessionPolicyResourceConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateAuthenticationSessionPolicyResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "id", "acc_test1"),
					resource.TestCheckResourceAttr(resourceName, "authentication_source.type", "IDP_ADAPTER"),
					resource.TestCheckResourceAttr(resourceName, "authentication_source.source_ref", "idptestme"),
					resource.TestCheckResourceAttr(resourceName, "enable_sessions", "false"),
					resource.TestCheckResourceAttr(resourceName, "persistent", "false"),
					resource.TestCheckResourceAttr(resourceName, "idle_timeout_mins", "240"),
					resource.TestCheckResourceAttr(resourceName, "max_timeout_mins", "240"),
					resource.TestCheckResourceAttr(resourceName, "timeout_display_unit", "MINUTES"),
					resource.TestCheckResourceAttr(resourceName, "authn_context_sensitive", "false"),
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

func testAccCheckPingFederateAuthenticationSessionPolicyResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateAuthenticationSessionPolicyResourceConfig() string {
	return `
resource "pingfederate_authentication_session_policy" "example" {
  id = "acc_test1"
  authentication_source = {
    type       = "IDP_ADAPTER"
    source_ref = "idptestme"
  }
  enable_sessions         = true
  persistent              = true
  idle_timeout_mins       = 300
  max_timeout_mins        = 300
  timeout_display_unit    = "HOURS"
  authn_context_sensitive = true
}`
}

func testAccPingFederateAuthenticationSessionPolicyResourceConfigUpdate() string {
	return `
resource "pingfederate_authentication_session_policy" "example" {
  id = "acc_test1"
  authentication_source = {
    type       = "IDP_ADAPTER"
    source_ref = "idptestme"
  }
  enable_sessions         = false
  persistent              = false
  idle_timeout_mins       = 240
  max_timeout_mins        = 240
  timeout_display_unit    = "MINUTES"
  authn_context_sensitive = false
}`
}

func testAccCheckPingFederateAuthenticationSessionPolicyResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := pfc.Session
		result, _, err := conn.GetSourcePolicy(&session.GetSourcePolicyInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: AuthenticationSessionPolicy (%s) not found", n)
		}
		if *result.TimeoutDisplayUnit != rs.Primary.Attributes["timeout_display_unit"] {
			return fmt.Errorf("Error: AuthenticationSessionPolicy response (%s) didnt match state (%s)", *result.TimeoutDisplayUnit, rs.Primary.Attributes["timeout_display_unit"])
		}

		return nil
	}
}

func Test_resourcePingFederateAuthenticationSessionPolicyResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.AuthenticationSessionPolicy
	}{
		{
			Resource: pf.AuthenticationSessionPolicy{
				AuthenticationSource: &pf.AuthenticationSource{
					SourceRef: &pf.ResourceLink{
						Id: String("auth_session_source_id"),
					},
					Type: String("auth_session_type"),
				},
				AuthnContextSensitive: Bool(true),
				EnableSessions:        Bool(true),
				Id:                    String("some+id"),
				IdleTimeoutMins:       Int(5),
				MaxTimeoutMins:        Int(5),
				Persistent:            Bool(true),
				TimeoutDisplayUnit:    String("MINUTES"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateAuthenticationSessionPolicyResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenAuthenticationSessionPolicy(&tc.Resource)).HasError())

			check := AuthenticationSessionPolicyData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandAuthenticationSessionPolicy(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
