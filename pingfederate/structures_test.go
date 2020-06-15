package pingfederate

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func Test_weCanFlattenScopes(t *testing.T) {
	initialScopes := []*pf.ScopeEntry{
		&pf.ScopeEntry{Name: String("mail"), Description: String("mail")},
		&pf.ScopeEntry{Name: String("profile"), Description: String("profile")},
		&pf.ScopeEntry{Name: String("openid"), Description: String("openid")},
		&pf.ScopeEntry{Name: String("address"), Description: String("address")},
		&pf.ScopeEntry{Name: String("phone"), Description: String("phone")},
	}

	output := []map[string]interface{}{
		map[string]interface{}{"name": "mail", "description": "mail"},
		map[string]interface{}{"name": "profile", "description": "profile"},
		map[string]interface{}{"name": "openid", "description": "openid"},
		map[string]interface{}{"name": "address", "description": "address"},
		map[string]interface{}{"name": "phone", "description": "phone"}}

	flattened := flattenScopes(initialScopes)

	equals(t, output, flattened)
}

func Test_weCanFlattenScopeGroups(t *testing.T) {
	initialScopeGroups := []*pf.ScopeGroupEntry{
		&pf.ScopeGroupEntry{Name: String("mail"), Description: String("mail"), Scopes: &[]*string{String("mail"), String("profile"), String("openid"), String("address"), String("phone")}},
	}

	output := []map[string]interface{}{
		map[string]interface{}{"name": "mail", "description": "mail", "scopes": []interface{}{
			"mail", "profile", "openid", "address", "phone",
		}}}

	flattened := flattenScopeGroups(initialScopeGroups)

	equals(t, output, flattened)
}

func Test_weCanFlattenPersistentGrantContract(t *testing.T) {
	initialPersistentGrantContract := &pf.PersistentGrantContract{
		ExtendedAttributes: &[]*pf.PersistentGrantAttribute{
			&pf.PersistentGrantAttribute{
				Name: String("woot"),
			},
		},
	}

	output := []map[string]interface{}{map[string]interface{}{"extended_attributes": []interface{}{"woot"}}}

	flattened := flattenPersistentGrantContract(initialPersistentGrantContract)

	equals(t, output, flattened)
}

func Test_weCanFlattenClientAuth(t *testing.T) {
	initialClientAuth := &pf.ClientAuth{
		Type:                String("CERTIFICATE"),
		ClientCertIssuerDn:  String(""),
		ClientCertSubjectDn: String(""),
	}

	output := []map[string]interface{}{map[string]interface{}{"type": "CERTIFICATE", "client_cert_issuer_dn": "", "client_cert_subject_dn": ""}}

	flattened := flattenClientAuth(initialClientAuth, initialClientAuth)

	equals(t, output, flattened)
}

func Test_weCanFlattenJwksSettings(t *testing.T) {
	initialJwksSettings := &pf.JwksSettings{
		JwksUrl: String("https://foo/bar.jwks"),
	}

	output := []map[string]interface{}{map[string]interface{}{"jwks_url": "https://foo/bar.jwks"}}

	flattened := flattenJwksSettings(initialJwksSettings)

	equals(t, output, flattened)
}

func Test_weCanFlattenResourceLink(t *testing.T) {
	initialResourceLink := &pf.ResourceLink{
		Id: String("atat"),
	}

	output := []map[string]interface{}{map[string]interface{}{"id": "atat"}}

	flattened := flattenResourceLink(initialResourceLink)

	equals(t, output, flattened)
}

func Test_weCanFlattenClientOIDCPolicy(t *testing.T) {
	initialClientOIDCPolicy := &pf.ClientOIDCPolicy{
		GrantAccessSessionRevocationApi: Bool(true),
		LogoutUris:                      &[]*string{String("https://logout")},
		PingAccessLogoutCapable:         Bool(true),
		PolicyGroup: &pf.ResourceLink{
			Id: String("atat"),
		},
	}

	output := []map[string]interface{}{map[string]interface{}{"grant_access_session_revocation_api": true, "logout_uris": []interface{}{"https://logout"}, "ping_access_logout_capable": true, "policy_group": []map[string]interface{}{map[string]interface{}{"id": "atat"}}}}

	flattened := flattenClientOIDCPolicy(initialClientOIDCPolicy)

	equals(t, output, flattened)
}

func testPluginConfiguration() []interface{} {
	return []interface{}{map[string]interface{}{"fields": schema.NewSet(configFieldHash, []interface{}{map[string]interface{}{"name": "Token Length", "value": "28", "inherited": false}})}}
}

func Test_weCanFlattenPluginConfiguration(t *testing.T) {
	initialPluginConfiguration := &pf.PluginConfiguration{
		Tables: &[]*pf.ConfigTable{
			&pf.ConfigTable{
				Name: String("Users"),
				Rows: &[]*pf.ConfigRow{
					&pf.ConfigRow{
						Fields: &[]*pf.ConfigField{
							&pf.ConfigField{
								Name:  String("Username"),
								Value: String("test"),
							},
						},
					},
				},
			},
		},
		Fields: &[]*pf.ConfigField{
			&pf.ConfigField{
				Name:      String("Token Length"),
				Value:     String("28"),
				Inherited: Bool(false),
			},
		},
	}

	// output := []interface{}{map[string]interface{}{"fields": schema.NewSet(configFieldHash, []interface{}{map[string]interface{}{"name": "Token Length", "value": "28", "inherited": false}})}}
	output := testPluginConfiguration()

	flattened := flattenPluginConfiguration(initialPluginConfiguration)

	assert(t, output[0].(map[string]interface{})["fields"].(*schema.Set).Equal(flattened[0].(map[string]interface{})["fields"].(*schema.Set)), "")
}

func Test_expandPluginConfiguration(t *testing.T) {
	// expanded := flatmap.Expand(testPluginConfiguration(), "configuration").([]interface{})
	expandPluginConfiguration := expandPluginConfiguration(testPluginConfiguration())

	equals(t, 1, len(*(*expandPluginConfiguration).Fields))
}

func Test_weCanFlattenAuthenticationPolicyContractAttribute(t *testing.T) {
	attributes := &[]*pf.AuthenticationPolicyContractAttribute{
		&pf.AuthenticationPolicyContractAttribute{
			Name: String("woot"),
		},
	}

	output := []interface{}{"woot"}

	flattened := flattenAuthenticationPolicyContractAttribute(*attributes)

	equals(t, output, flattened)
}

func Test_maskPluginConfigurationFromDescriptor(t *testing.T) {
	type args struct {
		desc     []byte
		origConf *pf.PluginConfiguration
		conf     *pf.PluginConfiguration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "spAdapter",
			args: args{
				desc: []byte(spDescripter),
				origConf: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:  String("Username"),
							Value: String("bob"),
						},
						{
							Name:  String("Password"),
							Value: String("demo"),
						},
						{
							Name:  String("Confirm Password"),
							Value: String("demo"),
						},
						{
							Name:  String("Relax Password Requirements"),
							Value: String("true"),
						},
					},
				},
				conf: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{
						{
							Name:  String("Username"),
							Value: String("bob"),
						},
						{
							Name: String("Password"),
						},
						{
							Name: String("Confirm Password"),
						},
						{
							Name:  String("Relax Password Requirements"),
							Value: String("true"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var descripter pf.PluginConfigDescriptor
			json.Unmarshal(tt.args.desc, &descripter)
			maskPluginConfigurationFromDescriptor(&descripter, tt.args.origConf, tt.args.conf)

			if got := tt.args.conf; !reflect.DeepEqual(got, tt.args.origConf) {
				t.Errorf("maskPluginConfigurationFromDescriptor() = %v, want %v", got, tt.args.origConf)
			}
		})
	}
}

var spDescripter = `{
	"description": "OpenToken Adapter 2.5.8",
	"fields": [
	  {
		"type": "TEXT",
		"name": "Password",
		"label": "Password",
		"description": "Password to use for generating the encryption key.",
		"defaultValue": "",
		"advanced": false,
		"required": true,
		"encrypted": true,
		"size": 10
	  },
	  {
		"type": "TEXT",
		"name": "Confirm Password",
		"label": "Confirm Password",
		"description": "Must match password field.",
		"defaultValue": "",
		"advanced": false,
		"required": true,
		"encrypted": true,
		"size": 10
	  },
	  {
		"type": "RADIO_GROUP",
		"name": "Transport Mode",
		"label": "Transport Mode",
		"description": "How the token is transported to/from the application, either via a query parameter, a cookie, or as a form POST.",
		"defaultValue": "2",
		"advanced": true,
		"required": false,
		"optionValues": [
		  {
			"name": "Query Parameter",
			"value": "0"
		  },
		  {
			"name": "Cookie",
			"value": "1"
		  },
		  {
			"name": "Form POST",
			"value": "2"
		  }
		]
	  },
	  {
		"type": "TEXT",
		"name": "Token Name",
		"label": "Token Name",
		"description": "The name of the cookie or query parameter that contains the token. This name must be unique for each adapter instance.",
		"defaultValue": "opentoken",
		"advanced": true,
		"required": true,
		"encrypted": false,
		"size": 15
	  },
	  {
		"type": "RADIO_GROUP",
		"name": "Cipher Suite",
		"label": "Cipher Suite",
		"description": "The algorithm, cipher mode, and key size that should be used for encrypting the token.",
		"defaultValue": "2",
		"advanced": true,
		"required": false,
		"optionValues": [
		  {
			"name": "Null",
			"value": "0"
		  },
		  {
			"name": "AES-256/CBC",
			"value": "1"
		  },
		  {
			"name": "AES-128/CBC",
			"value": "2"
		  },
		  {
			"name": "3DES-168/CBC",
			"value": "3"
		  }
		]
	  },
	  {
		"type": "TEXT",
		"name": "Authentication Service",
		"label": "Authentication Service",
		"description": "The URL to which the user is redirected for an SSO event. This URL overrides the Target Resource which is sent as a parameter to the Authentication Service.",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 40
	  },
	  {
		"type": "TEXT",
		"name": "Account Link Service",
		"label": "Account Link Service",
		"description": "The URL to which the user is redirected for Account Linking. This URL is part of an external SP application. This external application performs user authentication and returns the local user ID inside the token.",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 40
	  },
	  {
		"type": "TEXT",
		"name": "Logout Service",
		"label": "Logout Service",
		"description": "The URL to which the user is redirected for a single-logout event. This URL is part of an external application, which terminates the user session.",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 40
	  },
	  {
		"type": "TEXT",
		"name": "Cookie Domain",
		"label": "Cookie Domain",
		"description": "The server domain should be in the format of example.com. If no domain is specified, the value is obtained from the request.",
		"defaultValue": "",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 15
	  },
	  {
		"type": "TEXT",
		"name": "Cookie Path",
		"label": "Cookie Path",
		"description": "The path for the cookie that contains the token.",
		"defaultValue": "/",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 10
	  },
	  {
		"type": "TEXT",
		"name": "Token Lifetime",
		"label": "Token Lifetime",
		"description": "The duration (in seconds) for which the token is valid. Valid range is 1 to 28800.",
		"defaultValue": "300",
		"advanced": true,
		"required": true,
		"encrypted": false,
		"size": 8
	  },
	  {
		"type": "TEXT",
		"name": "Session Lifetime",
		"label": "Session Lifetime",
		"description": "The duration (in seconds) for which the token may be re-issued without authentication. Valid range is 1 to 259200.",
		"defaultValue": "43200",
		"advanced": true,
		"required": true,
		"encrypted": false,
		"size": 8
	  },
	  {
		"type": "TEXT",
		"name": "Not Before Tolerance",
		"label": "Not Before Tolerance",
		"description": "The amount of time (in seconds) to allow for clock skew between servers. Valid range is 0 to 3600.",
		"defaultValue": "0",
		"advanced": true,
		"required": true,
		"encrypted": false,
		"size": 8
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Force SunJCE Provider",
		"label": "Force SunJCE Provider",
		"description": "If checked, the SunJCE provider will be forced for encryption/decryption.",
		"defaultValue": "false",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Use Verbose Error Messages",
		"label": "Use Verbose Error Messages",
		"description": "If checked, use verbose TokenException messages",
		"defaultValue": "false",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Obfuscate Password",
		"label": "Obfuscate Password",
		"description": "If checked, the password will be obfuscated and password-strength validation will be applied. Clearing the checkbox allows backward compatibility with previous OpenToken agents.",
		"defaultValue": "true",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Session Cookie",
		"label": "Session Cookie",
		"description": "If checked, OpenToken will be set as a session cookie (rather than a persistent cookie). Applies only if Transport Mode is set as 'Cookie'.",
		"defaultValue": "false",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Secure Cookie",
		"label": "Secure Cookie",
		"description": "If checked, the OpenToken cookie will be set only if the request is on a secure channel (https). Applies only if Transport Mode is set as 'Cookie'.",
		"defaultValue": "false",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "HTTP Only Flag",
		"label": "HTTP Only Flag",
		"description": "Sets a flag for the cookie that it can only be read via http requests and disallows Javascript to access the cookie. Note: not all browsers respect the HTTP Only flag. ",
		"defaultValue": "true",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Send Subject as Query Parameter",
		"label": "Send Subject as Query Parameter",
		"description": "Checking this box will send the Subject ID as a clear-text query parameter, if Transport Mode is set to \"Query Parameter\". If Transport Mode is set to \"Form POST\", the Subject ID is sent as POST data.",
		"defaultValue": "",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "TEXT",
		"name": "Subject Query Parameter                 ",
		"label": "Subject Query Parameter                 ",
		"description": "The parameter name used for the Subject ID when the \"Send Subject ID as Query Parameter\" box is checked.",
		"defaultValue": "",
		"advanced": true,
		"required": false,
		"encrypted": false,
		"size": 15
	  },
	  {
		"type": "SELECT",
		"name": "Send Extended Attributes",
		"label": "Send Extended Attributes",
		"description": "Extended Attributes are typically sent only within the token, but this option overrides the normal behavior and allows the attributes to be included in browser cookies or query parameters.",
		"defaultValue": "",
		"advanced": true,
		"required": false,
		"optionValues": [
		  {
			"name": "None",
			"value": "0"
		  },
		  {
			"name": "Cookies",
			"value": "1"
		  },
		  {
			"name": "Query Parameters",
			"value": "2"
		  }
		]
	  },
	  {
		"type": "CHECK_BOX",
		"name": "Skip Trimming of Trailing Backslashes",
		"label": "Skip Trimming of Trailing Backslashes",
		"description": "If not checked, it prevents insecure content from affecting the security of your application/agent. We recommend to update your applications with the latest version of the agent. We recommend not to change the value of this flag.",
		"defaultValue": "false",
		"advanced": true,
		"required": false
	  },
	  {
		"type": "CHECK_BOX",
		"name": "URL Encode Cookie Values",
		"label": "URL Encode Cookie Values",
		"description": "If checked, the extended attribute cookie value will be URL encoded.",
		"defaultValue": "true",
		"advanced": true,
		"required": false
	  }
	],
	"tables": [],
	"actionDescriptors": [
	  {
		"name": "Download",
		"description": "Download the configuration file for the agent.",
		"downloadContentType": "text",
		"downloadFileName": "agent-config.txt",
		"download": true
	  }
	]
  }`
