package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOauthAuthServerSettings(t *testing.T) {
	resourceName := "pingfederate_oauth_auth_server_settings.settings"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthAuthServerSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists(resourceName),
				),
			},
			{
				Config: testAccPingFederateOauthAuthServerSettingsConfig("scopes {\nname = \"example\"\ndescription = \"example\"\n}\nscopes {\nname = \"example:*\"\ndescription = \"example dynamic\"\ndynamic = true\n}"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthAuthServerSettingsExists(resourceName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				//ImportStateVerifyIgnore: []string{"persistent_grant_lifetime", "persistent_grant_lifetime_unit"},
			},
		},
	})
}

func testAccCheckPingFederateOauthAuthServerSettingsDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthServerSettingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_oauth_auth_server_settings" "settings" {
		scopes {
			name        = "address"
			description = "address"
		}

		scopes {
			name        = "mail"
			description = "mail"
		}

		scopes {
			name        = "openid"
			description = "openid"
		}

		scopes {
			name        = "phone"
			description = "phone"
		}

		scopes {
			name        = "profile"
			description = "profile"
		}

		%s
		scope_groups {
			name        = "group1"
			description = "group1"

			scopes = [
				"address",
				"mail",
				"phone",
				"openid",
				"profile",
			]
		}

		persistent_grant_contract {
			extended_attributes = ["woot"]
		}

		allowed_origins = [
			"http://localhost"
		]

		persistent_grant_lifetime      = -1
        persistent_grant_lifetime_unit = "DAYS"
		default_scope_description  = ""
		authorization_code_timeout = 60
		authorization_code_entropy = 30
		refresh_token_length       = 42
		refresh_rolling_interval   = 0
	}`, configUpdate)
}

func testAccCheckPingFederateOauthAuthServerSettingsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthAuthServerSettings
		result, _, err := conn.GetAuthorizationServerSettings()

		if err != nil {
			return fmt.Errorf("Error: OauthAuthServerSettings (%s) not found", n)
		}

		if *result.DefaultScopeDescription != rs.Primary.Attributes["default_scope_description"] {
			return fmt.Errorf("Error: OauthAuthServerSettings response (%s) didnt match state (%s)", *result.DefaultScopeDescription, rs.Primary.Attributes["default_scope_description"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthAuthServerSettingsResourceReadResult(t *testing.T) {
	cases := []struct {
		Resource pf.AuthorizationServerSettings
		Version  string
	}{
		{
			Version: "10.2",
			Resource: pf.AuthorizationServerSettings{
				AdminWebServicePcvRef:                  &pf.ResourceLink{Id: String("1")},
				AllowUnidentifiedClientExtensionGrants: Bool(true),
				AllowUnidentifiedClientROCreds:         Bool(true),
				AllowedOrigins:                         &[]*string{String("foo")},
				ApprovedScopesAttribute:                String("foo1"),
				AtmIdForOAuthGrantManagement:           String("foo2"),
				AuthorizationCodeEntropy:               Int(1),
				AuthorizationCodeTimeout:               Int(2),
				BypassActivationCodeConfirmation:       Bool(true),
				BypassAuthorizationForApprovedGrants:   Bool(true),
				DefaultScopeDescription:                String("foo3"),
				DevicePollingInterval:                  Int(3),
				ExclusiveScopeGroups: &[]*pf.ScopeGroupEntry{
					{
						Description: String("foo"),
						Scopes:      &[]*string{String("foo")},
						Name:        String("foo"),
					}},
				ExclusiveScopes: &[]*pf.ScopeEntry{
					{
						Description: String("foo"),
						Dynamic:     Bool(true),
						Name:        String("foo"),
					},
				},
				ParReferenceLength:                 Int(4),
				ParReferenceTimeout:                Int(5),
				ParStatus:                          String("foo4"),
				PendingAuthorizationTimeout:        Int(6),
				PersistentGrantContract:            nil,
				PersistentGrantIdleTimeout:         Int(7),
				PersistentGrantIdleTimeoutTimeUnit: String("foo5"),
				PersistentGrantLifetime:            Int(8),
				PersistentGrantLifetimeUnit:        String("foo6"),
				PersistentGrantReuseGrantTypes:     []*string{String("foo7")},
				RefreshRollingInterval:             Int(9),
				RefreshTokenLength:                 Int(10),
				RegisteredAuthorizationPath:        String("foo8"),
				RollRefreshTokenValues:             Bool(true),
				ScopeForOAuthGrantManagement:       String("foo9"),
				ScopeGroups: &[]*pf.ScopeGroupEntry{
					{
						Description: String("bar1"),
						Scopes:      &[]*string{String("bar2")},
						Name:        String("bar3"),
					}},
				Scopes: &[]*pf.ScopeEntry{
					{
						Description: String("foo1"),
						Dynamic:     Bool(true),
						Name:        String("foo2"),
					},
				},
				TokenEndpointBaseUrl:                String("foo10"),
				TrackUserSessionsForLogout:          Bool(true),
				UserAuthorizationConsentAdapter:     String("foo11"),
				UserAuthorizationConsentPageSetting: String("foo12"),
				UserAuthorizationUrl:                String("foo13"),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOauthAuthServerSettingsResource().Schema
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthAuthServerSettingsResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthAuthServerSettingsResourceReadData(resourceLocalData, tc.Version))
		})
	}
}
