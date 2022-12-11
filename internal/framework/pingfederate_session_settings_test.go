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

func TestAccPingFederateSessionSettingsResource(t *testing.T) {
	resourceName := "pingfederate_session_settings.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateSessionSettingsResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateSessionSettingsResourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSessionSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "revoke_user_session_on_logout", "true"),
					resource.TestCheckResourceAttr(resourceName, "session_revocation_lifetime", "60"),
					resource.TestCheckResourceAttr(resourceName, "track_adapter_sessions_for_logout", "true"),
				),
			},
			{
				Config: testAccPingFederateSessionSettingsResourceConfigUpdate(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSessionSettingsResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "revoke_user_session_on_logout", "false"),
					resource.TestCheckResourceAttr(resourceName, "session_revocation_lifetime", "120"),
					resource.TestCheckResourceAttr(resourceName, "track_adapter_sessions_for_logout", "false"),
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

func testAccCheckPingFederateSessionSettingsResourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateSessionSettingsResourceConfig() string {
	return `
resource "pingfederate_session_settings" "example" {
  revoke_user_session_on_logout     = true
  session_revocation_lifetime       = 60
  track_adapter_sessions_for_logout = true
}`
}

func testAccPingFederateSessionSettingsResourceConfigUpdate() string {
	return `
resource "pingfederate_session_settings" "example" {
  revoke_user_session_on_logout     = false
  session_revocation_lifetime       = 120
  track_adapter_sessions_for_logout = false
}`
}

func testAccCheckPingFederateSessionSettingsResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := pfc.Session
		result, _, err := conn.GetSessionSettings()

		if err != nil {
			return fmt.Errorf("Error: SessionSettings (%s) not found", n)
		}
		if strconv.Itoa(*result.SessionRevocationLifetime) != rs.Primary.Attributes["session_revocation_lifetime"] {
			return fmt.Errorf("Error: SessionSettings response (%s) didnt match state (%s)", strconv.Itoa(*result.SessionRevocationLifetime), rs.Primary.Attributes["session_revocation_lifetime"])
		}

		return nil
	}
}

func Test_resourcePingFederateSessionSettingsResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.SessionSettings
	}{
		{
			Resource: pf.SessionSettings{
				RevokeUserSessionOnLogout:     Bool(true),
				SessionRevocationLifetime:     Int(60),
				TrackAdapterSessionsForLogout: Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateSessionSettingsResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenSessionSettings(&tc.Resource)).HasError())

			check := SessionSettingsData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandSessionSettings(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}
