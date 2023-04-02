package framework

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPingFederateOauthAuthServerSettingsDatasource(t *testing.T) {
	os.Setenv("PINGFEDERATE_PASSWORD", "2FederateM0re")
	datasourceName := "data.pingfederate_oauth_auth_server_settings.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOauthAuthServerSettingsDatasourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthAuthServerSettingsDatasourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "activation_code_check_mode", "AFTER_AUTHENTICATION"),
					resource.TestCheckNoResourceAttr(datasourceName, "admin_web_service_pcv_ref"),
					resource.TestCheckResourceAttr(datasourceName, "allow_unidentified_client_extension_grants", "false"),
					resource.TestCheckResourceAttr(datasourceName, "allow_unidentified_client_ro_creds", "false"),
					resource.TestCheckResourceAttr(datasourceName, "allowed_origins.#", "0"),
					resource.TestCheckNoResourceAttr(datasourceName, "approved_scopes_attribute"),
					resource.TestCheckResourceAttr(datasourceName, "atm_id_for_o_auth_grant_management", ""),
					resource.TestCheckResourceAttr(datasourceName, "authorization_code_entropy", "30"),
					resource.TestCheckResourceAttr(datasourceName, "authorization_code_timeout", "60"),
					resource.TestCheckResourceAttr(datasourceName, "bypass_activation_code_confirmation", "false"),
					resource.TestCheckResourceAttr(datasourceName, "bypass_authorization_for_approved_grants", "false"),
					resource.TestCheckResourceAttr(datasourceName, "client_secret_retention_period", "0"),
					resource.TestCheckResourceAttr(datasourceName, "default_scope_description", ""),
					resource.TestCheckResourceAttr(datasourceName, "device_polling_interval", "5"),
					resource.TestCheckResourceAttr(datasourceName, "disallow_plain_pkce", "false"),
					resource.TestCheckResourceAttr(datasourceName, "include_issuer_in_authorization_response", "false"),
					resource.TestCheckResourceAttr(datasourceName, "jwt_secured_authorization_response_mode_lifetime", "600"),
					resource.TestCheckResourceAttr(datasourceName, "par_reference_length", "24"),
					resource.TestCheckResourceAttr(datasourceName, "par_reference_timeout", "60"),
					resource.TestCheckResourceAttr(datasourceName, "par_status", "ENABLED"),
					resource.TestCheckResourceAttr(datasourceName, "pending_authorization_timeout", "600"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_idle_timeout", "30"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_idle_timeout_time_unit", "DAYS"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_lifetime", "-1"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_lifetime_unit", "DAYS"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_reuse_grant_types.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_reuse_grant_types.0", "IMPLICIT"),
					resource.TestCheckResourceAttr(datasourceName, "exclude_scopes.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "exclude_scope_groups.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_contract.core_attributes.#", "2"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_contract.core_attributes.0.name", "USER_KEY"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_contract.core_attributes.1.name", "USER_NAME"),
					resource.TestCheckResourceAttr(datasourceName, "persistent_grant_contract.extended_attributes.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "refresh_rolling_interval", "0"),
					resource.TestCheckResourceAttr(datasourceName, "refresh_token_length", "42"),
					resource.TestCheckResourceAttr(datasourceName, "refresh_token_rolling_grace_period", "0"),
					resource.TestCheckResourceAttr(datasourceName, "registered_authorization_path", ""),
					resource.TestCheckResourceAttr(datasourceName, "roll_refresh_token_values", "false"),
					resource.TestCheckResourceAttr(datasourceName, "scope_for_o_auth_grant_management", ""),
					resource.TestCheckResourceAttr(datasourceName, "scope_groups.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "scopes.#", "0"),
					resource.TestCheckResourceAttr(datasourceName, "token_endpoint_base_url", ""),
					resource.TestCheckResourceAttr(datasourceName, "track_user_sessions_for_logout", "false"),
					resource.TestCheckNoResourceAttr(datasourceName, "user_authorization_consent_adapter"),
					resource.TestCheckResourceAttr(datasourceName, "user_authorization_consent_page_setting", "INTERNAL"),
					resource.TestCheckResourceAttr(datasourceName, "user_authorization_url", ""),
				),
			},
		},
	})
}

func testAccCheckPingFederateOauthAuthServerSettingsDatasourceDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthAuthServerSettingsDatasourceConfig() string {
	return `data "pingfederate_oauth_auth_server_settings" "example" {}`
}
