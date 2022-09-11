package sdkv2provider

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientSettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOauthClientSettings(t *testing.T) {
	resourceName := "pingfederate_oauth_client_settings.settings"

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthClientSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthClientSettingsConfig("true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientSettingsExists(t, resourceName),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.allow_client_delete", "true"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.allowed_exclusive_scopes", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.bypass_activation_code_confirmation_override", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.ciba_polling_interval", "3"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.ciba_require_signed_requests", "false"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.client_cert_issuer_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.client_cert_issuer_type", "NONE"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.default_access_token_manager_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.device_flow_setting_type", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.device_polling_interval_override", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.disable_registration_access_tokens", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.enforce_replay_prevention", "true"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.initial_access_token_scope", "openid"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.oidc_policy", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.pending_authorization_timeout_override", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_time", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_time_unit", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_type", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout_time_unit", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout_type", "SERVER_DEFAULT"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.policy_refs", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_rolling", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_token_rolling_interval", "0"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_token_rolling_interval_type", "SERVER_DEFAULT"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.request_policy_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.require_proof_key_for_code_exchange", "true"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.require_signed_requests", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restrict_common_scopes", "true"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restrict_to_default_access_token_manager", "false"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restricted_common_scopes", ""),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.rotate_client_secret", "true"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.rotate_registration_access_token", "true"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.token_exchange_processor_policy_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.user_authorization_url_override", ""),
				),
			},
			{
				Config: testAccPingFederateOauthClientSettingsConfig("false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientSettingsExists(t, resourceName),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.allow_client_delete", "false"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.allowed_exclusive_scopes", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.bypass_activation_code_confirmation_override", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.ciba_polling_interval", "3"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.ciba_require_signed_requests", "false"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.client_cert_issuer_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.client_cert_issuer_type", "NONE"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.default_access_token_manager_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.device_flow_setting_type", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.device_polling_interval_override", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.disable_registration_access_tokens", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.enforce_replay_prevention", "true"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.initial_access_token_scope", "openid"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.oidc_policy", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.pending_authorization_timeout_override", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_time", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_time_unit", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_expiration_type", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout", "0"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout_time_unit", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.persistent_grant_idle_timeout_type", "SERVER_DEFAULT"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.policy_refs", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_rolling", "SERVER_DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_token_rolling_interval", "0"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.refresh_token_rolling_interval_type", "SERVER_DEFAULT"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.request_policy_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.require_proof_key_for_code_exchange", "true"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.require_signed_requests", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restrict_common_scopes", "false"),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restrict_to_default_access_token_manager", "false"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.restricted_common_scopes", ""),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.rotate_client_secret", "true"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.rotate_registration_access_token", "true"),
					//resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.token_exchange_processor_policy_ref", ""),
					resource.TestCheckResourceAttr(resourceName, "dynamic_client_registration.0.user_authorization_url_override", ""),
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

func testAccCheckPingFederateOauthClientSettingsDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(pfClient).OauthClientSettings
	_, _, err := conn.UpdateClientSettings(&oauthClientSettings.UpdateClientSettingsInput{Body: pf.ClientSettings{}})
	return err
}

func testAccPingFederateOauthClientSettingsConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_client_settings" "settings" {
  client_metadata {
    parameter = "example1"
  }
  client_metadata {
    parameter   = "example2"
    description = "example2"
  }
  dynamic_client_registration {
    restrict_common_scopes              = %s
    initial_access_token_scope          = "openid"
    enforce_replay_prevention           = true
    require_proof_key_for_code_exchange = true
  }
}`, configUpdate)
}

func testAccCheckPingFederateOauthClientSettingsExists(t *testing.T, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthClientSettings
		result, _, err := conn.GetClientSettings()

		assert.NoError(t, err)
		assert.Equal(t, strconv.FormatBool(*result.DynamicClientRegistration.RestrictCommonScopes), rs.Primary.Attributes["dynamic_client_registration.0.restrict_common_scopes"])
		//assert.Equal(t, *result.Description, rs.Primary.Attributes["description"])

		return nil
	}
}

func Test_resourcePingFederateOauthClientSettingsResourceReadResult(t *testing.T) {
	cases := []struct {
		Version  string
		Resource pf.ClientSettings
	}{
		{
			Version: "10.3.1",
			Resource: pf.ClientSettings{
				ClientMetadata: &[]*pf.ClientMetadata{
					{
						Parameter:   String("example1"),
						Description: String(""),
						MultiValued: Bool(false),
					},
					{
						Parameter:   String("example2"),
						Description: String("example2"),
						MultiValued: Bool(false),
					},
				},
				DynamicClientRegistration: &pf.DynamicClientRegistration{
					AllowClientDelete:                        Bool(true),
					AllowedExclusiveScopes:                   &[]*string{String("foo")},
					BypassActivationCodeConfirmationOverride: Bool(true),
					CibaPollingInterval:                      Int(1),
					CibaRequireSignedRequests:                Bool(true),
					ClientCertIssuerRef:                      &pf.ResourceLink{Id: String("bar")},
					ClientCertIssuerType:                     String("bar1"),
					DefaultAccessTokenManagerRef:             &pf.ResourceLink{Id: String("bar")},
					DeviceFlowSettingType:                    String("bar2"),
					DevicePollingIntervalOverride:            Int(1),
					DisableRegistrationAccessTokens:          Bool(true),
					EnforceReplayPrevention:                  Bool(true),
					InitialAccessTokenScope:                  String("bar3"),
					OidcPolicy: &pf.ClientRegistrationOIDCPolicy{
						IdTokenContentEncryptionAlgorithm: String("bar1"),
						IdTokenEncryptionAlgorithm:        String("bar2"),
						IdTokenSigningAlgorithm:           String("bar3"),
						PolicyGroup:                       &pf.ResourceLink{Id: String("bar")},
					},
					PendingAuthorizationTimeoutOverride: Int(1),
					PersistentGrantExpirationTime:       Int(1),
					PersistentGrantExpirationTimeUnit:   String("bar4"),
					PersistentGrantExpirationType:       String("bar5"),
					PersistentGrantIdleTimeout:          Int(1),
					PersistentGrantIdleTimeoutTimeUnit:  String("bar6"),
					PersistentGrantIdleTimeoutType:      String("bar7"),
					PolicyRefs: &[]*pf.ResourceLink{
						{Id: String("bar")},
						{Id: String("foo")},
					},
					RefreshRolling:                      String("bar8"),
					RefreshTokenRollingInterval:         Int(1),
					RefreshTokenRollingIntervalType:     String("bar0"),
					RequestPolicyRef:                    &pf.ResourceLink{Id: String("bar")},
					RequireProofKeyForCodeExchange:      Bool(true),
					RequireSignedRequests:               Bool(true),
					RestrictCommonScopes:                Bool(true),
					RestrictToDefaultAccessTokenManager: Bool(true),
					RestrictedCommonScopes:              &[]*string{String("foo")},
					RotateClientSecret:                  Bool(true),
					RotateRegistrationAccessToken:       Bool(true),
					TokenExchangeProcessorPolicyRef:     &pf.ResourceLink{Id: String("bar")},
					UserAuthorizationUrlOverride:        String("bara"),
				},
			},
		},
		{
			Version:  "10.3.1",
			Resource: pf.ClientSettings{},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOauthClientSettingsResource().Schema
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateOauthClientSettingsResourceReadResult(resourceLocalData, &tc.Resource, tc.Version)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthClientSettingsResourceReadData(resourceLocalData, tc.Version))
		})
	}
}
