package pingfederate

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationSelectors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"
	"github.com/iwarapter/pingfederate-sdk-go/services/passwordCredentialValidators"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/services/version"
)

var cfg *config.Config
var pfc pfClient

func TestMain(m *testing.M) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cfg = config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1")
	os.Setenv("PINGFEDERATE_PASSWORD", "2FederateM0re")
	flag.Parse()
	sweep := flag.Lookup("sweep")
	_, acceptanceTesting := os.LookupEnv("TF_ACC")
	if (sweep != nil && sweep.Value.String() != "") || acceptanceTesting {
		client := version.New(cfg)
		v, r, err := client.GetVersion()
		if err != nil {
			if r != nil {
				b, _ := httputil.DumpResponse(r, true)
				log.Fatalf("unable to get pingfederate '%s'\n%s", err, b)
			}
			log.Fatalf("unable to get pingfederate '%s'", err)
		}
		pfc.apiVersion = *v.Version
		log.Printf("Connected to PingFederate %s", *v.Version)
	}
	if acceptanceTesting {
		if err := dataSetup(); err != nil {
			log.Fatalf("unable to setup test data\n%s", err)
		}
	}
	resource.TestMain(m)
}

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"pingfederate": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	err := testAccProvider.Configure(context.TODO(), terraform.NewResourceConfigRaw(nil))
	if err != nil {
		t.Fatal(err)
	}
}

func dataSetup() error {
	ssets := serverSettings.New(cfg)
	if _, _, err := ssets.UpdateServerSettings(&serverSettings.UpdateServerSettingsInput{
		Body: pf.ServerSettings{
			FederationInfo: &pf.FederationInfo{
				BaseUrl:        String("https://localhost:9031"),
				Saml2EntityId:  String("testing"),
				Saml1xIssuerId: String("foo"),
				WsfedRealm:     String("foo"),
			},
			RolesAndProtocols: &pf.RolesAndProtocols{
				IdpRole: &pf.IdpRole{
					Enable:                     Bool(true),
					EnableOutboundProvisioning: Bool(true),
					Saml20Profile: &pf.SAML20Profile{
						Enable: Bool(true),
					},
				},
				OauthRole: &pf.OAuthRole{
					EnableOauth:         Bool(true),
					EnableOpenIdConnect: Bool(true),
				},
				SpRole: &pf.SpRole{
					Enable: Bool(true),
					Saml20Profile: &pf.SpSAML20Profile{
						Enable: Bool(true),
					},
				},
			},
		},
	}); err != nil {
		return fmt.Errorf("unable to set server settings\n%s", err)
	}

	pcv := passwordCredentialValidators.New(cfg)
	if _, _, err := pcv.GetPasswordCredentialValidator(&passwordCredentialValidators.GetPasswordCredentialValidatorInput{Id: "pcvtestme"}); err != nil {
		if _, _, err := pcv.CreatePasswordCredentialValidator(&passwordCredentialValidators.CreatePasswordCredentialValidatorInput{
			Body: pf.PasswordCredentialValidator{
				Configuration: &pf.PluginConfiguration{
					Tables: &[]*pf.ConfigTable{
						{
							Name: String("Users"),
							Rows: &[]*pf.ConfigRow{
								{
									Fields: &[]*pf.ConfigField{
										{
											Name:  String("Username"),
											Value: String("example"),
										},
										{
											Name:  String("Password"),
											Value: String("example"),
										},
										{
											Name:  String("Confirm Password"),
											Value: String("example"),
										},
										{
											Name:  String("Relax Password Requirements"),
											Value: String("true"),
										},
									},
								},
							},
						},
					},
				},
				Id:                  String("pcvtestme"),
				Name:                String("pcvtestme"),
				PluginDescriptorRef: &pf.ResourceLink{Id: String("org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator")},
			},
		}); err != nil {
			return fmt.Errorf("unable to create test password credential validator\n%s", err)
		}
	}
	atv := oauthAccessTokenManagers.New(cfg)
	if _, _, err := atv.GetTokenManager(&oauthAccessTokenManagers.GetTokenManagerInput{Id: "testme"}); err != nil {
		if _, _, err := atv.CreateTokenManager(&oauthAccessTokenManagers.CreateTokenManagerInput{
			Body: pf.AccessTokenManager{
				AccessControlSettings: nil,
				AttributeContract: &pf.AccessTokenAttributeContract{
					ExtendedAttributes: &[]*pf.AccessTokenAttribute{
						{
							Name: String("name"),
						},
					},
				},
				Configuration:             &pf.PluginConfiguration{},
				Id:                        String("testme"),
				Name:                      String("testme"),
				PluginDescriptorRef:       &pf.ResourceLink{Id: String("org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin")},
				SelectionSettings:         nil,
				SessionValidationSettings: nil,
			},
		}); err != nil {
			return fmt.Errorf("unable to create test password credential validator\n%s", err)
		}
	}
	idp := idpAdapters.New(cfg)
	if _, _, err := idp.GetIdpAdapter(&idpAdapters.GetIdpAdapterInput{Id: "idptestme"}); err != nil {
		if _, _, err := idp.CreateIdpAdapter(&idpAdapters.CreateIdpAdapterInput{BypassExternalValidation: Bool(true),
			Body: pf.IdpAdapter{
				AttributeContract: &pf.IdpAdapterAttributeContract{
					CoreAttributes: &[]*pf.IdpAdapterAttribute{
						{
							Name:      String("username"),
							Pseudonym: Bool(true),
						},
					},
					ExtendedAttributes: &[]*pf.IdpAdapterAttribute{
						{
							Name: String("sub"),
						},
					},
				},
				AttributeMapping: nil,
				AuthnCtxClassRef: nil,
				Configuration: &pf.PluginConfiguration{
					Tables: &[]*pf.ConfigTable{
						{
							Name: String("Credential Validators"),
							Rows: &[]*pf.ConfigRow{
								{
									Fields: &[]*pf.ConfigField{
										{
											Name:  String("Password Credential Validator Instance"),
											Value: String("pcvtestme"),
										},
									},
								},
							},
						},
					},
					Fields: &[]*pf.ConfigField{
						{
							Name:  String("Realm"),
							Value: String("foo"),
						},
						{
							Name:  String("Challenge Retries"),
							Value: String("3"),
						},
					},
				},
				Id:                  String("idptestme"),
				Name:                String("idptestme"),
				ParentRef:           nil,
				PluginDescriptorRef: &pf.ResourceLink{Id: String("com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter")},
			},
		}); err != nil {
			return fmt.Errorf("unable to create test idp adapter\n%s", err)
		}
	}
	authnSel := authenticationSelectors.New(cfg)
	if _, _, err := authnSel.GetAuthenticationSelector(&authenticationSelectors.GetAuthenticationSelectorInput{Id: "authseltestme"}); err != nil {
		if _, _, err := authnSel.CreateAuthenticationSelector(&authenticationSelectors.CreateAuthenticationSelectorInput{
			Body: pf.AuthenticationSelector{
				Configuration: &pf.PluginConfiguration{
					Tables: &[]*pf.ConfigTable{
						{
							Name: String("Networks"),
							Rows: &[]*pf.ConfigRow{
								{
									Fields: &[]*pf.ConfigField{
										{
											Name:  String("Network Range (CIDR notation)"),
											Value: String("127.0.0.1/32"),
										},
									},
								},
							},
						},
					},
					Fields: &[]*pf.ConfigField{
						{
							Name:  String("Result Attribute Name"),
							Value: String(""),
						},
					},
				},
				Id:                  String("authseltestme"),
				Name:                String("authseltestme"),
				PluginDescriptorRef: &pf.ResourceLink{Id: String("com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector")},
			},
		}); err != nil {
			return fmt.Errorf("unable to create test authentication selector\n%s", err)
		}
	}
	return nil
}
