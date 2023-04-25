package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
	"github.com/stretchr/testify/require"

	fresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("pingfederate_oauth_client", &resource.Sweeper{
		Name:         "pingfederate_oauth_client",
		Dependencies: []string{},
		F: func(r string) error {
			results, _, err := pfc.OauthClients.GetClients(&oauthClients.GetClientsInput{Filter: "acc_test"})
			if err != nil {
				return fmt.Errorf("unable to list oauth clients %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := pfc.OauthClients.DeleteClient(&oauthClients.DeleteClientInput{Id: *item.ClientId})
				if err != nil {
					return fmt.Errorf("unable to sweep oauth client %s because %s", *item.ClientId, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOAuthClientResource(t *testing.T) {
	resourceName := "pingfederate_oauth_client.example1"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOAuthClientResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthClientResourceConfig("https://demo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "my-client-name"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_session_management_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.1", "https://demo"),
				),
			},
			{
				Config: testAccPingFederateOAuthClientResourceConfig("https://update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "my-client-name"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_session_management_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.1", "https://update"),
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

func TestAccPingFederateOAuthClientWithClientSecretResource(t *testing.T) {
	resourceName := "pingfederate_oauth_client.example1"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOAuthClientResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthClientWithClientSecretResourceConfig("https://demo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "my-client-with-secret"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_session_management_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.1", "https://demo"),
				),
			},
			{
				Config: testAccPingFederateOAuthClientWithClientSecretResourceConfig("https://update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "my-client-with-secret"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_session_management_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.1", "https://update"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_auth.secret", "client_auth.encrypted_secret"},
			},
		},
	})
}

func TestAccPingFederateOAuthClientWithClientWithoutOidcPolicy(t *testing.T) {
	resourceName := "pingfederate_oauth_client.example"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPingFederateOAuthClientResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: `
resource "pingfederate_oauth_client" "example" {
  client_id         = "acc_test_withoutoidc"
  name              = "acc_test_withoutoidc"
  grant_types       = ["ACCESS_TOKEN_VALIDATION"]
  restrict_scopes   = true
  restricted_scopes = ["openid"]
}`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_withoutoidc"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "ACCESS_TOKEN_VALIDATION"),
					resource.TestCheckResourceAttr(resourceName, "restrict_scopes", "true"),
					resource.TestCheckResourceAttr(resourceName, "restricted_scopes.0", "openid"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"client_auth.secret", "client_auth.encrypted_secret"},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_withoutoidc"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "ACCESS_TOKEN_VALIDATION"),
					resource.TestCheckResourceAttr(resourceName, "restrict_scopes", "true"),
					resource.TestCheckResourceAttr(resourceName, "restricted_scopes.0", "openid"),
				),
			},
		},
	})
}

func TestAccPingFederateOAuthClientResourceSdkUpgradeV0toV1(t *testing.T) {
	resourceName := "pingfederate_oauth_client.my_client"
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"pingfederate": {
						VersionConstraint: "0.0.24",
						Source:            "iwarapter/pingfederate",
					},
				},
				Config: testAccPingFederateOAuthClientResourceSdkUpgradeV0config(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref.0.id", "testme"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PlanOnly:                 true,
				Config:                   testAccPingFederateOAuthClientResourceSdkUpgradeV1config(),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPingFederateOAuthClientResourceSdkUpgradeV1config(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.ping_access_logout_capable", "true"),
				),
			},
		},
	})
}

func TestAccPingFederateOAuthClientResourceSdkUpgradeV1checkListToSetHandles(t *testing.T) {
	resourceName := "pingfederate_oauth_client.my_client"
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"pingfederate": {
						VersionConstraint: "0.1.0",
						Source:            "iwarapter/pingfederate",
					},
				},
				Config: testAccPingFederateOAuthClientResourceSdkUpgradeV1configWithRedirects(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade2"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.ping_access_logout_capable", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_uris.0", "https://foo.com"),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PlanOnly:                 true,
				Config:                   testAccPingFederateOAuthClientResourceSdkUpgradeV1configWithRedirects(),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPingFederateOAuthClientResourceSdkUpgradeV1configWithRedirects(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "name", "acc_test_upgrade2"),
					resource.TestCheckResourceAttr(resourceName, "grant_types.0", "EXTENSION"),
					resource.TestCheckResourceAttr(resourceName, "default_access_token_manager_ref", "testme"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.grant_access_session_revocation_api", "false"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.logout_uris.0", "https://logout"),
					resource.TestCheckResourceAttr(resourceName, "oidc_policy.ping_access_logout_capable", "true"),
					resource.TestCheckResourceAttr(resourceName, "redirect_uris.0", "https://foo.com"),
				),
			},
		},
	})
}

func TestAccPingFederateOAuthClientResourceIssue263(t *testing.T) {
	resourceName := "pingfederate_oauth_client.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthClientResourceIssue263(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
				),
			},
		},
	})
}

func TestAccPingFederateOAuthClientResourceIssue275(t *testing.T) {
	resourceName := "pingfederate_oauth_client.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOAuthClientResourceIssue275("secret1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "client_auth.secret", "secret1"),
				),
			},
			{
				Config: testAccPingFederateOAuthClientResourceIssue275("secret2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOAuthClientResourceExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "client_auth.secret", "secret2"),
				),
			},
		},
	})
}

func testAccCheckPingFederateOAuthClientResourceDestroy(_ *terraform.State) error {
	return nil
}

func testAccPingFederateOAuthClientResourceSdkUpgradeV0config() string {
	return `
resource "pingfederate_oauth_client" "my_client" {
  client_id   = "acc_test_upgrade"
  name        = "acc_test_upgrade"
  grant_types = ["EXTENSION"]
  default_access_token_manager_ref {
    id = "testme"
  }
  client_auth {
    type   = "SECRET"
    secret = "top_secret"
  }

  oidc_policy {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout"]
    ping_access_logout_capable          = true
  }
}
`
}

func testAccPingFederateOAuthClientResourceSdkUpgradeV1config() string {
	return `
resource "pingfederate_oauth_client" "my_client" {
  client_id                        = "acc_test_upgrade"
  name                             = "acc_test_upgrade"
  grant_types                      = ["EXTENSION"]
  default_access_token_manager_ref = "testme"
  client_auth = {
    type   = "SECRET"
    secret = "top_secret"
  }
  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout"]
    ping_access_logout_capable          = true
  }
}`
}

func testAccPingFederateOAuthClientResourceSdkUpgradeV1configWithRedirects() string {
	return `
resource "pingfederate_oauth_client" "my_client" {
  client_id                        = "acc_test_upgrade2"
  name                             = "acc_test_upgrade2"
  grant_types                      = ["EXTENSION"]
  default_access_token_manager_ref = "testme"
  client_auth = {
    type   = "SECRET"
    secret = "top_secret"
  }
  redirect_uris = ["https://foo.com"]
  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout"]
    ping_access_logout_capable          = true
  }
}`
}

func testAccPingFederateOAuthClientResourceIssue263() string {
	return `
resource "pingfederate_oauth_client" "test" {
  client_id   = "test"
  name        = "test"
  description = "This client is a test client"

  grant_types = [
    "ACCESS_TOKEN_VALIDATION"
  ]

  client_auth = {
    type                      = "PRIVATE_KEY_JWT"
    enforce_replay_prevention = false
  }

  jwks_settings = {
    jwks_url = "https://localhost/"
  }

  validate_using_all_eligible_atms = true

  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = null
    ping_access_logout_capable          = false
  }

  ## Enabling the following three lines "solves" the issue (but it is more like a workaround though)
  # persistent_grant_expiration_time = 90
  # persistent_grant_expiration_time_unit = "DAYS"
  # persistent_grant_expiration_type = "OVERRIDE_SERVER_DEFAULT"
}`
}

func testAccPingFederateOAuthClientResourceIssue275(secret string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_client" "test" {
  client_id   = "test2"
  name        = "test2"
  description = "This client is a test client"

  grant_types = [
    "ACCESS_TOKEN_VALIDATION"
  ]
  client_auth = {
    secret = "%s"
    type   = "SECRET"
  }
}`, secret)
}

func testAccPingFederateOAuthClientResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_client" "example1" {
  client_id                        = "tf-acc-woot"
  name                             = "my-client-name"
  grant_types                      = ["EXTENSION"]
  default_access_token_manager_ref = "testme"

  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout", "%s"]
    ping_access_logout_capable          = true
  }
}`, configUpdate)
}

func testAccPingFederateOAuthClientWithClientSecretResourceConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_oauth_client" "example1" {
  client_id                        = "tf-acc-secret"
  name                             = "my-client-with-secret"
  grant_types                      = ["EXTENSION"]
  default_access_token_manager_ref = "testme"

  client_auth = {
    type   = "SECRET"
    secret = "top_secret"
  }

  oidc_policy = {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://logout", "%s"]
    ping_access_logout_capable          = true
  }
}`, configUpdate)
}

func testAccCheckPingFederateOAuthClientResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no ID is set")
		}

		conn := pfc.OauthClients
		result, _, err := conn.GetClient(&oauthClients.GetClientInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("oauthClient (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("oauthClient response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateOAuthClientResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.Client
	}{
		{
			Resource: pf.Client{
				AllowAuthenticationApiInit:               Bool(true),
				BypassActivationCodeConfirmationOverride: Bool(true),
				BypassApprovalPage:                       Bool(true),
				CibaDeliveryMode:                         String("CibaDeliveryMode"),
				CibaNotificationEndpoint:                 String("CibaNotificationEndpoint"),
				CibaPollingInterval:                      Int(1),
				CibaRequestObjectSigningAlgorithm:        String("CibaRequestObjectSigningAlgorithm"),
				CibaRequireSignedRequests:                Bool(true),
				CibaUserCodeSupported:                    Bool(true),
				ClientAuth: &pf.ClientAuth{
					ClientCertIssuerDn:                String("ClientCertIssuerDn"),
					ClientCertSubjectDn:               String("ClientCertSubjectDn"),
					EnforceReplayPrevention:           Bool(true),
					Secret:                            String("Secret"),
					TokenEndpointAuthSigningAlgorithm: String("TokenEndpointAuthSigningAlgorithm"),
					Type:                              String("Type"),
				},
				ClientId: String("ClientId"),
				DefaultAccessTokenManagerRef: &pf.ResourceLink{
					Id: String("DefaultAccessTokenManagerRef"),
				},
				Description:                   String("Description"),
				DeviceFlowSettingType:         String("DeviceFlowSettingType"),
				DevicePollingIntervalOverride: Int(2),
				Enabled:                       Bool(true),
				ExclusiveScopes:               &[]*string{String("ExclusiveScopes")},
				ExtendedParameters: map[string]*pf.ParameterValues{
					"example": {
						Values: &[]*string{String("ExtendedParameters")},
					},
				},
				GrantTypes: &[]*string{String("GrantTypes")},
				JwksSettings: &pf.JwksSettings{
					Jwks:    String("Jwks"),
					JwksUrl: String("JwksUrl"),
				},
				LogoUrl: String("LogoUrl"),
				Name:    String("Name"),
				OidcPolicy: &pf.ClientOIDCPolicy{
					GrantAccessSessionRevocationApi:   Bool(true),
					IdTokenContentEncryptionAlgorithm: String("IdTokenContentEncryptionAlgorithm"),
					IdTokenEncryptionAlgorithm:        String("IdTokenEncryptionAlgorithm"),
					IdTokenSigningAlgorithm:           String("IdTokenSigningAlgorithm"),
					LogoutUris:                        &[]*string{String("LogoutUris")},
					PairwiseIdentifierUserType:        Bool(true),
					PingAccessLogoutCapable:           Bool(true),
					PolicyGroup: &pf.ResourceLink{
						Id: String("PolicyGroup"),
					},
					SectorIdentifierUri: String("SectorIdentifierUri"),
				},
				PendingAuthorizationTimeoutOverride: Int(3),
				PersistentGrantExpirationTime:       Int(4),
				PersistentGrantExpirationTimeUnit:   String("PersistentGrantExpirationTimeUnit"),
				PersistentGrantExpirationType:       String("PersistentGrantExpirationType"),
				PersistentGrantIdleTimeout:          Int(5),
				PersistentGrantIdleTimeoutTimeUnit:  String("PersistentGrantIdleTimeoutTimeUnit"),
				PersistentGrantIdleTimeoutType:      String("PersistentGrantIdleTimeoutType"),
				PersistentGrantReuseGrantTypes:      &[]*string{String("PersistentGrantReuseGrantTypes")},
				PersistentGrantReuseType:            String("PersistentGrantReuseType"),
				RedirectUris:                        &[]*string{String("RedirectUris")},
				RefreshRolling:                      String("RefreshRolling"),
				RefreshTokenRollingInterval:         Int(1),
				RefreshTokenRollingIntervalType:     String("RefreshTokenRollingIntervalType"),
				RequestObjectSigningAlgorithm:       String("RequestObjectSigningAlgorithm"),
				RequestPolicyRef:                    &pf.ResourceLink{Id: String("RequestPolicyRef")},
				RequireProofKeyForCodeExchange:      Bool(true),
				RequirePushedAuthorizationRequests:  Bool(true),
				RequireSignedRequests:               Bool(true),
				RestrictScopes:                      Bool(true),
				RestrictToDefaultAccessTokenManager: Bool(true),
				RestrictedResponseTypes:             &[]*string{String("RestrictedResponseTypes")},
				RestrictedScopes:                    &[]*string{String("RestrictedScopes")},
				TokenExchangeProcessorPolicyRef:     &pf.ResourceLink{Id: String("TokenExchangeProcessorPolicyRef")},
				UserAuthorizationUrlOverride:        String("UserAuthorizationUrlOverride"),
				ValidateUsingAllEligibleAtms:        Bool(true),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {
			res := &pingfederateOAuthClientResource{}
			ctx := context.Background()
			schResp := &fresource.SchemaResponse{}
			res.Schema(ctx, fresource.SchemaRequest{}, schResp)
			require.False(t, schResp.Diagnostics.HasError())

			state := tfsdk.State{Schema: schResp.Schema}
			require.False(t, state.Set(ctx, flattenClient(&tc.Resource)).HasError())

			check := ClientData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandClient(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
}

func Test_resourcePingFederateOAuthClientResourceUpgradeData(t *testing.T) {
	res := &pingfederateOAuthClientResource{}
	ctx := context.Background()
	schemaV0 := resourceClientV0()

	dataV0 := ClientDataV0{
		ID:                                  types.StringValue("id"),
		Name:                                types.StringValue("name"),
		ClientId:                            types.StringValue("client_id"),
		Enabled:                             types.BoolValue(true),
		GrantTypes:                          []types.String{types.StringValue("grant_type")},
		PersistentGrantExpirationTime:       types.NumberNull(),
		RequireProofKeyForCodeExchange:      types.BoolValue(true),
		RestrictedScopes:                    []types.String{types.StringValue("RestrictedScopes")},
		CibaDeliveryMode:                    types.StringNull(),
		PendingAuthorizationTimeoutOverride: types.NumberNull(),
		RequestObjectSigningAlgorithm:       types.StringNull(),
		RestrictScopes:                      types.BoolValue(true),
		TokenExchangeProcessorPolicyRef:     []ResourceLink{{ID: types.StringValue("id")}},
		CibaPollingInterval:                 types.NumberNull(),
		RedirectUris:                        []types.String{types.StringValue("RedirectUris")},
		PersistentGrantIdleTimeout:          types.NumberNull(),
		DevicePollingIntervalOverride:       types.NumberNull(),
		DeviceFlowSettingType:               types.StringNull(),
		LogoUrl:                             types.StringNull(),
		OidcPolicy: []ClientOIDCPolicyDataV0{
			{
				GrantAccessSessionRevocationApi:   types.BoolValue(true),
				IdTokenContentEncryptionAlgorithm: types.StringValue("IdTokenContentEncryptionAlgorithm"),
				IdTokenEncryptionAlgorithm:        types.StringValue("IdTokenEncryptionAlgorithm"),
				IdTokenSigningAlgorithm:           types.StringValue("IdTokenSigningAlgorithm"),
				LogoutUris:                        []types.String{types.StringValue("LogoutUris")},
				PairwiseIdentifierUserType:        types.BoolValue(true),
				PingAccessLogoutCapable:           types.BoolValue(true),
				PolicyGroup:                       []ResourceLink{{ID: types.StringValue("id")}},
				SectorIdentifierUri:               types.StringValue("SectorIdentifierUri"),
			},
		},
		PersistentGrantExpirationTimeUnit:        types.StringNull(),
		RequirePushedAuthorizationRequests:       types.BoolValue(true),
		UserAuthorizationUrlOverride:             types.StringNull(),
		BypassActivationCodeConfirmationOverride: types.BoolValue(true),
		ValidateUsingAllEligibleAtms:             types.BoolValue(true),
		CibaUserCodeSupported:                    types.BoolValue(true),
		Description:                              types.StringNull(),
		JwksSettings: []JwksSettingsData{
			{
				Jwks:    types.StringValue("Jwks"),
				JwksUrl: types.StringValue("JwksUrl"),
			},
		},
		RefreshRolling:                      types.StringValue("refresh_rolling"),
		RequestPolicyRef:                    []ResourceLink{{ID: types.StringValue("id")}},
		RequireSignedRequests:               types.BoolValue(true),
		CibaNotificationEndpoint:            types.StringNull(),
		CibaRequireSignedRequests:           types.BoolValue(true),
		RestrictToDefaultAccessTokenManager: types.BoolValue(true),
		PersistentGrantExpirationType:       types.StringNull(),
		PersistentGrantIdleTimeoutTimeUnit:  types.StringNull(),
		DefaultAccessTokenManagerRef:        []ResourceLink{{ID: types.StringValue("id")}},
		ExclusiveScopes:                     []types.String{types.StringValue("ExclusiveScopes")},
		ClientAuth: []ClientAuthDataV0{
			{
				ClientCertIssuerDn:                types.StringValue("ClientCertIssuerDn"),
				ClientCertSubjectDn:               types.StringValue("ClientCertSubjectDn"),
				EnforceReplayPrevention:           types.BoolValue(true),
				Secret:                            types.StringValue("Secret"),
				TokenEndpointAuthSigningAlgorithm: types.StringValue("TokenEndpointAuthSigningAlgorithm"),
				Type:                              types.StringValue("Type"),
			},
		},
		PersistentGrantIdleTimeoutType:    types.StringNull(),
		RestrictedResponseTypes:           []types.String{types.StringValue("RestrictedResponseTypes")},
		BypassApprovalPage:                types.BoolValue(true),
		CibaRequestObjectSigningAlgorithm: types.StringNull(),
		ExtendedProperties: []ExtendedPropertiesDataV0{
			{
				KeyName: types.StringValue("foo"),
				Values:  []types.String{types.StringValue("bar")},
			},
		},
	}
	state := tfsdk.State{Schema: schemaV0}
	require.False(t, state.Set(ctx, &dataV0).HasError())

	upgr := res.UpgradeState(ctx)
	resp := &fresource.UpgradeStateResponse{State: tfsdk.State{Schema: resourceClient()}}
	upgr[0].StateUpgrader(ctx, fresource.UpgradeStateRequest{State: &state}, resp)
	require.False(t, resp.Diagnostics.HasError())

	check := ClientData{}
	require.False(t, resp.State.Get(ctx, &check).HasError())

	expected := ClientData{
		Id:                                  types.StringValue("client_id"),
		Name:                                types.StringValue("name"),
		ClientId:                            types.StringValue("client_id"),
		Enabled:                             types.BoolValue(true),
		GrantTypes:                          types.ListValueMust(types.StringType, []attr.Value{types.StringValue("grant_type")}),
		PersistentGrantExpirationTime:       types.NumberNull(),
		RequireProofKeyForCodeExchange:      types.BoolValue(true),
		RestrictedScopes:                    types.SetValueMust(types.StringType, []attr.Value{types.StringValue("RestrictedScopes")}),
		CibaDeliveryMode:                    types.StringNull(),
		PendingAuthorizationTimeoutOverride: types.NumberNull(),
		RequestObjectSigningAlgorithm:       types.StringNull(),
		RestrictScopes:                      types.BoolValue(true),
		PersistentGrantReuseGrantTypes:      types.ListNull(types.StringType),
		TokenExchangeProcessorPolicyRef:     types.StringValue("id"),
		CibaPollingInterval:                 types.NumberNull(),
		RedirectUris:                        types.SetValueMust(types.StringType, []attr.Value{types.StringValue("RedirectUris")}),
		PersistentGrantIdleTimeout:          types.NumberNull(),
		DevicePollingIntervalOverride:       types.NumberNull(),
		DeviceFlowSettingType:               types.StringNull(),
		LogoUrl:                             types.StringNull(),
		OidcPolicy: &ClientOIDCPolicyData{
			GrantAccessSessionRevocationApi:   types.BoolValue(true),
			IdTokenContentEncryptionAlgorithm: types.StringValue("IdTokenContentEncryptionAlgorithm"),
			IdTokenEncryptionAlgorithm:        types.StringValue("IdTokenEncryptionAlgorithm"),
			IdTokenSigningAlgorithm:           types.StringValue("IdTokenSigningAlgorithm"),
			LogoutUris:                        types.ListValueMust(types.StringType, []attr.Value{types.StringValue("LogoutUris")}),
			PairwiseIdentifierUserType:        types.BoolValue(true),
			PingAccessLogoutCapable:           types.BoolValue(true),
			PolicyGroup:                       types.StringValue("id"),
			SectorIdentifierUri:               types.StringValue("SectorIdentifierUri"),
		},
		PersistentGrantExpirationTimeUnit:        types.StringNull(),
		RequirePushedAuthorizationRequests:       types.BoolValue(true),
		UserAuthorizationUrlOverride:             types.StringNull(),
		BypassActivationCodeConfirmationOverride: types.BoolValue(true),
		ValidateUsingAllEligibleAtms:             types.BoolValue(true),
		CibaUserCodeSupported:                    types.BoolValue(true),
		Description:                              types.StringNull(),
		JwksSettings: &JwksSettingsData{
			Jwks:    types.StringValue("Jwks"),
			JwksUrl: types.StringValue("JwksUrl"),
		},
		RefreshRolling:                      types.StringValue("refresh_rolling"),
		RequestPolicyRef:                    types.StringValue("id"),
		RequireSignedRequests:               types.BoolValue(true),
		CibaNotificationEndpoint:            types.StringNull(),
		CibaRequireSignedRequests:           types.BoolValue(true),
		RestrictToDefaultAccessTokenManager: types.BoolValue(true),
		PersistentGrantExpirationType:       types.StringNull(),
		PersistentGrantIdleTimeoutTimeUnit:  types.StringNull(),
		DefaultAccessTokenManagerRef:        types.StringValue("id"),
		ExclusiveScopes:                     types.ListValueMust(types.StringType, []attr.Value{types.StringValue("ExclusiveScopes")}),
		ClientAuth: &ClientAuthData{
			ClientCertIssuerDn:                types.StringValue("ClientCertIssuerDn"),
			ClientCertSubjectDn:               types.StringValue("ClientCertSubjectDn"),
			EnforceReplayPrevention:           types.BoolValue(true),
			Secret:                            types.StringValue("Secret"),
			TokenEndpointAuthSigningAlgorithm: types.StringValue("TokenEndpointAuthSigningAlgorithm"),
			Type:                              types.StringValue("Type"),
		},
		PersistentGrantIdleTimeoutType:    types.StringNull(),
		RestrictedResponseTypes:           types.ListValueMust(types.StringType, []attr.Value{types.StringValue("RestrictedResponseTypes")}),
		BypassApprovalPage:                types.BoolValue(true),
		CibaRequestObjectSigningAlgorithm: types.StringNull(),
		ExtendedParameters: map[string]*ParameterValuesData{
			"foo": {
				Values: types.ListValueMust(types.StringType, []attr.Value{types.StringValue("bar")}),
			},
		},
	}
	assert.Equal(t, expected, check)
}

func Test_resourcePingFederateOAuthClientResourceVersionModifications(t *testing.T) {
	cli := pfClient{}
	res := &pingfederateOAuthClientResource{client: &cli}

	defaults := ClientData{
		ClientSecretRetentionPeriodType:            types.StringValue("SERVER_DEFAULT"),
		PersistentGrantReuseType:                   types.StringValue("SERVER_DEFAULT"),
		RequireJwtSecuredAuthorizationResponseMode: types.BoolValue(false),
		RefreshTokenRollingGracePeriodType:         types.StringValue("SERVER_DEFAULT"),
		RefreshTokenRollingIntervalType:            types.StringValue("SERVER_DEFAULT"),
	}
	tests := []struct {
		version string
		after   ClientData
	}{
		{"9.3", ClientData{}},
		{"10.0", ClientData{}},
		{"10.1", ClientData{}},
		{"10.2", ClientData{}},
		{"10.3", ClientData{RefreshTokenRollingIntervalType: types.StringValue("SERVER_DEFAULT")}},
		{"11.0", ClientData{PersistentGrantReuseType: types.StringValue("SERVER_DEFAULT"), RefreshTokenRollingGracePeriodType: types.StringValue("SERVER_DEFAULT"), RefreshTokenRollingIntervalType: types.StringValue("SERVER_DEFAULT")}},
		{"11.1", ClientData{ClientSecretRetentionPeriodType: types.StringValue("SERVER_DEFAULT"), PersistentGrantReuseType: types.StringValue("SERVER_DEFAULT"), RequireJwtSecuredAuthorizationResponseMode: types.BoolValue(false), RefreshTokenRollingGracePeriodType: types.StringValue("SERVER_DEFAULT"), RefreshTokenRollingIntervalType: types.StringValue("SERVER_DEFAULT")}},
		{"11.2", ClientData{ClientSecretRetentionPeriodType: types.StringValue("SERVER_DEFAULT"), PersistentGrantReuseType: types.StringValue("SERVER_DEFAULT"), RequireJwtSecuredAuthorizationResponseMode: types.BoolValue(false), RefreshTokenRollingGracePeriodType: types.StringValue("SERVER_DEFAULT"), RefreshTokenRollingIntervalType: types.StringValue("SERVER_DEFAULT")}},
		{"11.3", ClientData{ClientSecretRetentionPeriodType: types.StringValue("SERVER_DEFAULT"), PersistentGrantReuseType: types.StringValue("SERVER_DEFAULT"), RequireJwtSecuredAuthorizationResponseMode: types.BoolValue(false), RefreshTokenRollingGracePeriodType: types.StringValue("SERVER_DEFAULT"), RefreshTokenRollingIntervalType: types.StringValue("SERVER_DEFAULT")}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("we handle %s", tt.version), func(t *testing.T) {
			cli.apiVersion = tt.version
			var err error
			cli.major, cli.minor, err = parseVersion(tt.version)
			require.NoError(t, err)
			before := defaults
			res.versionRequestModifier(&before)
			assert.Equal(t, tt.after, before)

			assert.Equal(t, defaults, *res.versionResponseModifier(&before))
		})
	}
}
