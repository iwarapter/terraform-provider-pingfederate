package sdkv2provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("oauth_client", &resource.Sweeper{
		Name:         "oauth_client",
		Dependencies: []string{},
		F: func(r string) error {
			svc := oauthClients.New(cfg)
			results, _, err := svc.GetClients(&oauthClients.GetClientsInput{Filter: "acc_test"})
			if err != nil {
				return fmt.Errorf("unable to list oauth client %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteClient(&oauthClients.DeleteClientInput{Id: *item.ClientId})
				if err != nil {
					return fmt.Errorf("unable to sweep oauth client %s because %s", *item.ClientId, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateOauthClient(t *testing.T) {
	resourceName := "pingfederate_oauth_client.my_client"
	resource.ParallelTest(t, resource.TestCase{
		// PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateOauthClientDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateOauthClientConfig("https://demo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientExists(resourceName),
					// testAccCheckPingFederateOauthClientAttributes(),
				),
			},
			{
				Config: testAccPingFederateOauthClientConfig("https://update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateOauthClientExists(resourceName),
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

func testAccCheckPingFederateOauthClientDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateOauthClientConfig(configUpdate string) string {
	return fmt.Sprintf(`
	resource "pingfederate_oauth_client" "my_client" {
		client_id = "tf-acc-woot"
		name      = "tf-acc-woot"

		grant_types = [
			"EXTENSION",
		]

		default_access_token_manager_ref {
			id = "testme"
		}

		oidc_policy {
			grant_access_session_revocation_api = false
			logout_uris = [
				"https://logout",
				"%s"
			]
			ping_access_logout_capable = true
		}
	}

	resource "pingfederate_oauth_client" "my_client_2" {
		client_id = "tf-acc-woot-2"
		name      = "tf-acc-woot-2"

		grant_types = [
			"CLIENT_CREDENTIALS",
		]

		default_access_token_manager_ref {
			id = "testme"
		}

		client_auth {
			type = "SECRET"
			secret = "Secret"
		}

		oidc_policy {
			grant_access_session_revocation_api = false
			logout_uris = []
			ping_access_logout_capable = false
			pairwise_identifier_user_type = false
		}
	}`, configUpdate)
}

func testAccCheckPingFederateOauthClientExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).OauthClients
		result, _, err := conn.GetClient(&oauthClients.GetClientInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: OauthClient (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: OauthClient response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateOauthClientResourceReadData(t *testing.T) {

	cases := []struct {
		Resource pf.Client
	}{
		{
			Resource: pf.Client{
				BypassActivationCodeConfirmationOverride: Bool(false),
				BypassApprovalPage:                       Bool(false),
				CibaDeliveryMode:                         String("CibaDeliveryMode"),
				CibaNotificationEndpoint:                 String("CibaNotificationEndpoint"),
				CibaPollingInterval:                      Int(1),
				CibaRequestObjectSigningAlgorithm:        String("CibaRequestObjectSigningAlgorithm"),
				CibaRequireSignedRequests:                Bool(false),
				CibaUserCodeSupported:                    Bool(false),
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
				Enabled:                       Bool(false),
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
				RedirectUris:                        &[]*string{String("RedirectUris")},
				RefreshRolling:                      String("RefreshRolling"),
				RequestObjectSigningAlgorithm:       String("RequestObjectSigningAlgorithm"),
				RequestPolicyRef:                    &pf.ResourceLink{Id: String("RequestPolicyRef")},
				RequireProofKeyForCodeExchange:      Bool(false),
				RequireSignedRequests:               Bool(false),
				RestrictToDefaultAccessTokenManager: Bool(false),
				RestrictScopes:                      Bool(false),
				RestrictedResponseTypes:             &[]*string{String("RestrictedResponseTypes")},
				RestrictedScopes:                    &[]*string{String("RestrictedScopes")},
				TokenExchangeProcessorPolicyRef:     &pf.ResourceLink{Id: String("TokenExchangeProcessorPolicyRef")},
				UserAuthorizationUrlOverride:        String("UserAuthorizationUrlOverride"),
				ValidateUsingAllEligibleAtms:        Bool(false),
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateOauthClientResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			if resourceLocalData.Set("client_auth", []map[string]string{{"secret": "Secret"}}) != nil {
				t.Errorf("unable to set test data for Test_resourcePingFederateOauthClientResourceReadData")
			}

			resourcePingFederateOauthClientResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateOauthClientResourceReadData(resourceLocalData))
		})
	}
}
