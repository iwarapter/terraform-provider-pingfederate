package pingfederate

import (
	"context"

	"github.com/iwarapter/terraform-provider-pingfederate/internal/mutexkv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//Provider does stuff
//
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["username"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGFEDERATE_USERNAME"}, "Administrator"),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["password"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGFEDERATE_PASSWORD"}, "2Federate"),
			},
			"context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["context"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGFEDERATE_CONTEXT"}, "/pf-admin-api/v1"),
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["base_url"],
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"PINGFEDERATE_BASEURL"}, "https://localhost:9999"),
			},
			"bypass_external_validation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "External validation will be bypassed when set to true. Default to false.",
				Default:     false,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"pingfederate_custom_data_store":      dataSourcePingFederateCustomDataStore(),
			"pingfederate_jdbc_data_store":        dataSourcePingFederateJdbcDataStore(),
			"pingfederate_ldap_data_store":        dataSourcePingFederateLdapDataStore(),
			"pingfederate_keypair_signing_csr":    dataSourcePingFederateKeyPairSigningCsr(),
			"pingfederate_keypair_ssl_server_csr": dataSourcePingFederateKeyPairSslServerCsr(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"pingfederate_authentication_policies_settings":             resourcePingFederateAuthenticationPoliciesSettingsResource(),
			"pingfederate_authentication_policies":                      resourcePingFederateAuthenticationPoliciesResource(),
			"pingfederate_authentication_api_application":               resourcePingFederateAuthnApiApplicationResource(),
			"pingfederate_authentication_api_settings":                  resourcePingFederateAuthnApiSettingsResource(),
			"pingfederate_authentication_policy_contract":               resourcePingFederateAuthenticationPolicyContractResource(),
			"pingfederate_authentication_selector":                      resourcePingFederateAuthenticationSelectorResource(),
			"pingfederate_certificates_ca":                              resourcePingFederateCertificatesCaResource(),
			"pingfederate_custom_data_store":                            resourcePingFederateCustomDataStoreResource(),
			"pingfederate_jdbc_data_store":                              resourcePingFederateJdbcDataStoreResource(),
			"pingfederate_ldap_data_store":                              resourcePingFederateLdapDataStoreResource(),
			"pingfederate_idp_sp_connection":                            resourcePingFederateIdpSpConnectionResource(),
			"pingfederate_idp_adapter":                                  resourcePingFederateIdpAdapterResource(),
			"pingfederate_kerberos_realm":                               resourcePingFederateKerberosRealmResource(),
			"pingfederate_keypair_signing":                              resourcePingFederateKeypairSigningResource(),
			"pingfederate_keypair_signing_csr":                          resourcePingFederateKeypairSigningCsrResource(),
			"pingfederate_keypair_ssl_server":                           resourcePingFederateKeypairSslServerResource(),
			"pingfederate_keypair_ssl_server_csr":                       resourcePingFederateKeypairSslServerCsrResource(),
			"pingfederate_keypair_ssl_server_settings":                  resourcePingFederateKeypairSslServerSettingsResource(),
			"pingfederate_notification_publisher":                       resourcePingFederateNotificationPublisherResource(),
			"pingfederate_oauth_auth_server_settings":                   resourcePingFederateOauthAuthServerSettingsResource(),
			"pingfederate_oauth_authentication_policy_contract_mapping": resourcePingFederateOauthAuthenticationPolicyContractMappingsResource(),
			"pingfederate_oauth_client":                                 resourcePingFederateOauthClientResource(),
			"pingfederate_oauth_access_token_manager":                   resourcePingFederateOauthAccessTokenManagersResource(),
			"pingfederate_oauth_access_token_mappings":                  resourcePingFederateOauthAccessTokenMappingsResource(),
			"pingfederate_oauth_openid_connect_policy":                  resourcePingFederateOpenIdConnectPolicyResource(),
			"pingfederate_oauth_resource_owner_credentials_mappings":    resourcePingFederateOauthResourceOwnerCredentialsMappingsResource(),
			"pingfederate_server_settings":                              resourcePingFederateServerSettingsResource(),
			"pingfederate_sp_adapter":                                   resourcePingFederateSpAdapterResource(),
			"pingfederate_sp_authentication_policy_contract_mapping":    resourcePingFederateSpAuthenticationPolicyContractMappingResource(),
			"pingfederate_sp_idp_connection":                            resourcePingFederateSpIdpConnectionResource(),
			"pingfederate_password_credential_validator":                resourcePingFederatePasswordCredentialValidatorResource(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// This is a global MutexKV for use within this plugin.
var awsMutexKV = mutexkv.NewMutexKV()

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"username":                   "The username for pingfederate API.",
		"password":                   "The password for pingfederate API.",
		"base_url":                   "The base url of the pingfederate API.",
		"context":                    "The context path of the pingfederate API.",
		"bypass_external_validation": "External validation will be bypassed when set to true.",
	}
}

func providerConfigure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := &pfConfig{
		Username:                 d.Get("username").(string),
		Password:                 d.Get("password").(string),
		BaseURL:                  d.Get("base_url").(string),
		Context:                  d.Get("context").(string),
		BypassExternalValidation: d.Get("bypass_external_validation").(bool),
	}

	return config.Client()
}

// Takes the result of flatmap.Expand for an array of strings
// and returns a []*string
func expandStringList(configured []interface{}) []*string {
	vs := make([]*string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, String(v.(string)))
		}
	}
	return vs
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

func setResourceDataStringWithDiagnostic(d *schema.ResourceData, name string, data *string, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}

func setResourceDataIntWithDiagnostic(d *schema.ResourceData, name string, data *int, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}

func setResourceDataBoolWithDiagnostic(d *schema.ResourceData, name string, data *bool, diags *diag.Diagnostics) {
	if data != nil {
		if err := d.Set(name, *data); err != nil {
			*diags = append(*diags, diag.FromErr(err)...)
		}
	}
}
