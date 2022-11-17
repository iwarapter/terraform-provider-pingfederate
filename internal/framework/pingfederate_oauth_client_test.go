package framework

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

//func init() {
//	resource.AddTestSweepers("authentication_policy_contract", &resource.Sweeper{
//		Name:         "authentication_policy_contract",
//		Dependencies: []string{},
//		F: func(r string) error {
//			svc := OAuthClients.New(cfg)
//			results, _, err := svc.GetOAuthClients(&OAuthClients.GetOAuthClientsInput{Filter: "acc_test"})
//			if err != nil {
//				return fmt.Errorf("unable to list authentication policy contracts %s", err)
//			}
//			for _, item := range *results.Items {
//				_, _, err := svc.DeleteOAuthClient(&OAuthClients.DeleteOAuthClientInput{Id: *item.Id})
//				if err != nil {
//					return fmt.Errorf("unable to sweep authentication policy contract %s because %s", *item.Id, err)
//				}
//			}
//			return nil
//		},
//	})
//}

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

func testAccCheckPingFederateOAuthClientResourceDestroy(_ *terraform.State) error {
	return nil
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
			resourceSchema, diags := res.GetSchema(ctx)
			require.False(t, diags.HasError())

			state := tfsdk.State{Schema: resourceSchema}
			require.False(t, state.Set(ctx, flattenClient(&tc.Resource)).HasError())

			check := ClientData{}
			require.False(t, state.Get(ctx, &check).HasError())

			resp := *expandClient(check)
			assert.Equal(t, tc.Resource, resp)
		})
	}
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
