package pingfederate

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/flatmap"
	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

// Returns test configuration
func testScopeConf() map[string]string {
	return map[string]string{
		"scopes.#":                      "5",
		"scopes.1455792420.description": "mail",
		"scopes.1455792420.name":        "mail",
		"scopes.1963347957.description": "profile",
		"scopes.1963347957.name":        "profile",
		"scopes.296925214.description":  "openid",
		"scopes.296925214.name":         "openid",
		"scopes.3688904175.description": "address",
		"scopes.3688904175.name":        "address",
		"scopes.563431727.description":  "phone",
		"scopes.563431727.name":         "phone",
	}
}

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

func Test_expandScopes(t *testing.T) {
	expanded := flatmap.Expand(testScopeConf(), "scopes").([]interface{})
	expandScopes := expandScopes(expanded)

	equals(t, 5, len(*expandScopes))
}

func testScopeGroupConf() map[string]string {
	return map[string]string{
		"scope_groups.#":                      "1",
		"scope_groups.1867744217.description": "group1",
		"scope_groups.1867744217.name":        "group1",
		"scope_groups.1867744217.scopes.#":    "5",
		"scope_groups.1867744217.scopes.0":    "address",
		"scope_groups.1867744217.scopes.1":    "mail",
		"scope_groups.1867744217.scopes.2":    "phone",
		"scope_groups.1867744217.scopes.3":    "openid",
		"scope_groups.1867744217.scopes.4":    "profile",
	}
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

func Test_expandScopeGroups(t *testing.T) {
	expanded := flatmap.Expand(testScopeGroupConf(), "scope_groups").([]interface{})
	expandScopeGroups := expandScopeGroups(expanded)

	equals(t, 5, len(*(*expandScopeGroups)[0].Scopes))
}

func testPersistentGrantContractConf() map[string]string {
	return map[string]string{
		"persistent_grant_contract.#":                               "1",
		"persistent_grant_contract.454952399.extended_attributes.#": "1",
		"persistent_grant_contract.454952399.extended_attributes.0": "woot",
	}
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

func Test_expandPersistentGrantContract(t *testing.T) {
	expanded := flatmap.Expand(testPersistentGrantContractConf(), "persistent_grant_contract").([]interface{})
	expandPersistentGrantContract := expandPersistentGrantContract(expanded)

	equals(t, "woot", *(*(*expandPersistentGrantContract).ExtendedAttributes)[0].Name)
}

func testClientAuth() map[string]string {
	return map[string]string{
		"client_auth.#": "1",
		"client_auth.1602811133.client_cert_issuer_dn":     "",
		"client_auth.1602811133.client_cert_subject_dn":    "",
		"client_auth.1602811133.enforce_replay_prevention": "true",
		"client_auth.1602811133.type":                      "CERTIFICATE",
	}
}

func Test_weCanFlattenClientAuth(t *testing.T) {
	initialClientAuth := &pf.ClientAuth{
		Type:                String("CERTIFICATE"),
		ClientCertIssuerDn:  String(""),
		ClientCertSubjectDn: String(""),
	}

	output := []map[string]interface{}{map[string]interface{}{"type": "CERTIFICATE", "client_cert_issuer_dn": "", "client_cert_subject_dn": ""}}

	flattened := flattenClientAuth(initialClientAuth)

	equals(t, output, flattened)
}

func Test_expandClientAuth(t *testing.T) {
	expanded := flatmap.Expand(testClientAuth(), "client_auth").([]interface{})
	expandClientAuth := expandClientAuth(expanded)

	equals(t, "CERTIFICATE", *(*expandClientAuth).Type)
}

func testJwksSettings() map[string]string {
	return map[string]string{
		"jwks_settings.#":                   "1",
		"jwks_settings.3441057763.jwks_url": "https://foo/bar.jwks",
	}
}

func Test_weCanFlattenJwksSettings(t *testing.T) {
	initialJwksSettings := &pf.JwksSettings{
		JwksUrl: String("https://foo/bar.jwks"),
	}

	output := []map[string]interface{}{map[string]interface{}{"jwks_url": "https://foo/bar.jwks"}}

	flattened := flattenJwksSettings(initialJwksSettings)

	equals(t, output, flattened)
}

func Test_expandJwksSettings(t *testing.T) {
	expanded := flatmap.Expand(testJwksSettings(), "jwks_settings").([]interface{})
	expandJwksSettings := expandJwksSettings(expanded)

	equals(t, "https://foo/bar.jwks", *(*expandJwksSettings).JwksUrl)
}

func testResourceLink() map[string]string {
	return map[string]string{
		"default_access_token_manager_ref.#":             "1",
		"default_access_token_manager_ref.2212279603.id": "atat",
	}
}

func Test_weCanFlattenResourceLink(t *testing.T) {
	initialResourceLink := &pf.ResourceLink{
		Id: String("atat"),
	}

	output := []map[string]interface{}{map[string]interface{}{"id": "atat"}}

	flattened := flattenResourceLink(initialResourceLink)

	equals(t, output, flattened)
}

func Test_expandResourceLink(t *testing.T) {
	expanded := flatmap.Expand(testResourceLink(), "default_access_token_manager_ref").([]interface{})
	expandResourceLink := expandResourceLink(expanded)

	equals(t, "atat", *(*expandResourceLink).Id)
}

func testClientOIDCPolicy() map[string]string {
	return map[string]string{
		"oidc_policy.#": "1",
		"oidc_policy.2491599396.grant_access_session_revocation_api": "true",
		"oidc_policy.2491599396.id_token_signing_algorithm":          "",
		"oidc_policy.2491599396.logout_uris.#":                       "1",
		"oidc_policy.2491599396.logout_uris.3341685451":              "https://logout",
		"oidc_policy.2491599396.ping_access_logout_capable":          "true",
	}
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

func Test_expandClientOIDCPolicy(t *testing.T) {
	expanded := flatmap.Expand(testClientOIDCPolicy(), "oidc_policy").([]interface{})
	expandClientOIDCPolicy := expandClientOIDCPolicy(expanded)

	equals(t, true, *(*expandClientOIDCPolicy).PingAccessLogoutCapable)
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

func testAuthenticationPolicyContractAttribute() map[string]string {
	return map[string]string{
		"extended_attributes.#": "1",
		"extended_attributes.0": "woot",
	}
}

func Test_expandAuthenticationPolicyContractAttribute(t *testing.T) {
	expanded := flatmap.Expand(testAuthenticationPolicyContractAttribute(), "extended_attributes").([]interface{})
	attributes := expandAuthenticationPolicyContractAttribute(expanded)

	equals(t, 1, len(*attributes))
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
		desc     *pf.PluginConfigDescriptor
		origConf *pf.PluginConfiguration
		conf     *pf.PluginConfiguration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "HOURS passes",
			args: args{
				desc: &pf.PluginConfigDescriptor{
					ActionDescriptors: nil,
					Description:       nil,
					Fields:            &[]*pf.FieldDescriptor{},
					Tables: &[]*pf.TableDescriptor{
						{
							Columns: &[]*pf.FieldDescriptor{
								{
									Type: String("TEXT"),
									Name: String("Username"),
								},
								{
									Type: String("HASHED_TEXT"),
									Name: String("Password"),
								},
								{
									Type: String("HASHED_TEXT"),
									Name: String("Confirm Password"),
								},
								{
									Type: String("CHECK_BOX"),
									Name: String("Relax Password Requirements"),
								},
							},
							Description:       nil,
							Label:             nil,
							Name:              nil,
							RequireDefaultRow: nil,
						},
					},
				},
				origConf: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{},
					Tables: &[]*pf.ConfigTable{
						{
							Inherited: nil,
							Name:      String("foo"),
							Rows: &[]*pf.ConfigRow{
								{
									DefaultRow: nil,
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
							},
						},
					},
				},
				conf: &pf.PluginConfiguration{
					Fields: &[]*pf.ConfigField{},
					Tables: &[]*pf.ConfigTable{
						{
							Inherited: nil,
							Name:      String("foo"),
							Rows: &[]*pf.ConfigRow{
								{
									DefaultRow: nil,
									Fields: &[]*pf.ConfigField{
										{
											Name:  String("Username"),
											Value: String("bob"),
										},
										{
											Name:           String("Password"),
											EncryptedValue: String("OBF:JWE:eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2Iiwia2lkIjoicVByTTIzT2JreiIsInZlcnNpb24iOiIxMC4wLjIuMiJ9..eaoxz6IFWkCsg5om58lbSQ.1uRKxwIsB473vkP4KBY8yAUtet_Dt-ZCwDcAOkqJzGQ8sO19PfTZZQvTrmQMKZ7wTeFdKN0J5ipzSzt-MpRIuw.ViEz10Djqy9oy4XoGW2nRA"),
										},
										{
											Name:           String("Confirm Password"),
											EncryptedValue: String("OBF:JWE:eyJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2Iiwia2lkIjoicVByTTIzT2JreiIsInZlcnNpb24iOiIxMC4wLjIuMiJ9..eaoxz6IFWkCsg5om58lbSQ.1uRKxwIsB473vkP4KBY8yAUtet_Dt-ZCwDcAOkqJzGQ8sO19PfTZZQvTrmQMKZ7wTeFdKN0J5ipzSzt-MpRIuw.ViEz10Djqy9oy4XoGW2nRA"),
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maskPluginConfigurationFromDescriptor(tt.args.desc, tt.args.origConf, tt.args.conf)

			foo := flattenPluginConfiguration(tt.args.origConf)
			bar := flattenPluginConfiguration(tt.args.conf)
			printPluginConfig("foo", tt.args.origConf)
			printPluginConfig("foo", tt.args.conf)
			fmt.Printf("%v\n%v", foo, bar)
			//if got := ; !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("maskPluginConfigurationFromDescriptor() = %v, want %v", got, tt.want)
			//}
		})
	}
}
