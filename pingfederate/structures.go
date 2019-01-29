package pingfederate

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourceLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func resourcePluginConfiguration() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"tables": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem:     resourceConfigTable(),
				},
				"fields": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem:     resourceConfigField(),
				},
			},
		},
	}
}

func resourceConfigTable() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// "rows:": {
			// 	Type: schema.TypeSet,
			// 	// Optional: true,
			// 	Elem: resourceConfigRow(),
			// },
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceConfigRow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fields": {
				Type: schema.TypeSet,
				// Optional: true,
				Elem: resourceConfigField(),
			},
		},
	}
}

func resourceConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypted_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// Takes the result of schema.Set of strings and returns a []*string
func expandStringSet(configured *schema.Set) []*string {
	return expandStringList(configured.List())
}

// Takes list of pointers to strings. Expand to an array
// of raw strings and returns a []interface{}
// to keep compatibility w/ schema.NewSetschema.NewSet
func flattenStringList(list []*string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, *v)
	}
	return vs
}

func flattenScopes(in []*pf.ScopeEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		s["description"] = *v.Description
		m = append(m, s)
	}
	return m
}

func expandScopes(in []interface{}) *[]*pf.ScopeEntry {
	scopeList := []*pf.ScopeEntry{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ScopeEntry{
			Name:        String(l["name"].(string)),
			Description: String(l["description"].(string)),
		}
		scopeList = append(scopeList, s)
	}
	return &scopeList
}

func flattenScopeGroups(in []*pf.ScopeGroupEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		s["description"] = *v.Description
		s["scopes"] = flattenStringList(*v.Scopes)
		m = append(m, s)
	}
	return m
}

func expandScopeGroups(in []interface{}) *[]*pf.ScopeGroupEntry {
	scopeGroupList := []*pf.ScopeGroupEntry{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ScopeGroupEntry{
			Name:        String(l["name"].(string)),
			Description: String(l["description"].(string)),
		}
		scopes := []*string{}
		for _, scope := range l["scopes"].([]interface{}) {
			scopes = append(scopes, String(scope.(string)))
		}
		s.Scopes = &scopes
		scopeGroupList = append(scopeGroupList, s)
	}
	return &scopeGroupList
}

func flattenPersistentGrantContract(in *pf.PersistentGrantContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["extended_attributes"] = flattenPersistentGrantAttributes(*in.ExtendedAttributes)
	m = append(m, s)
	return m
}

func expandPersistentGrantContract(in []interface{}) *pf.PersistentGrantContract {
	pgc := &pf.PersistentGrantContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		atr := []*pf.PersistentGrantAttribute{}
		for _, exAtr := range l["extended_attributes"].([]interface{}) {
			atr = append(atr, &pf.PersistentGrantAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func flattenPersistentGrantAttributes(in []*pf.PersistentGrantAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func expandClientAuth(in []interface{}) *pf.ClientAuth {
	ca := &pf.ClientAuth{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["client_cert_issuer_dn"]; ok {
			ca.ClientCertIssuerDn = String(val.(string))
		}
		if val, ok := l["client_cert_subject_dn"]; ok {
			ca.ClientCertSubjectDn = String(val.(string))
		}
		if val, ok := l["enforce_replay_prevention"]; ok {
			ca.EnforceReplayPrevention = Bool(val.(bool))
		}
		if val, ok := l["secret"]; ok {
			ca.Secret = String(val.(string))
		}
		ca.Type = String(l["type"].(string))
	}
	return ca
}

func flattenClientAuth(in *pf.ClientAuth) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ClientCertIssuerDn != nil {
		s["client_cert_issuer_dn"] = *in.ClientCertIssuerDn
	}
	if in.ClientCertSubjectDn != nil {
		s["client_cert_subject_dn"] = *in.ClientCertSubjectDn
	}
	if in.EnforceReplayPrevention != nil {
		s["enforce_replay_prevention"] = *in.EnforceReplayPrevention
	}
	// if in.Secret != nil {
	// 	s["secret"] = *in.Secret
	// }
	s["type"] = *in.Type
	m = append(m, s)
	return m
}

func expandJwksSettings(in []interface{}) *pf.JwksSettings {
	ca := &pf.JwksSettings{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["jwks"]; ok {
			ca.Jwks = String(val.(string))
		}
		if val, ok := l["jwks_url"]; ok {
			ca.JwksUrl = String(val.(string))
		}
	}
	return ca
}

func flattenJwksSettings(in *pf.JwksSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Jwks != nil {
		s["jwks"] = *in.Jwks
	}
	if in.JwksUrl != nil {
		s["jwks_url"] = *in.JwksUrl
	}
	m = append(m, s)
	return m
}

func expandResourceLink(in []interface{}) *pf.ResourceLink {
	ca := &pf.ResourceLink{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["id"]; ok {
			ca.Id = String(val.(string))
		}
	}
	return ca
}

func flattenResourceLink(in *pf.ResourceLink) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	m = append(m, s)
	return m
}

func expandClientOIDCPolicy(in []interface{}) *pf.ClientOIDCPolicy {
	ca := &pf.ClientOIDCPolicy{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["grant_access_session_revocation_api"]; ok {
			ca.GrantAccessSessionRevocationApi = Bool(val.(bool))
		}
		if val, ok := l["id_token_signing_algorithm"]; ok {
			ca.IdTokenSigningAlgorithm = String(val.(string))
		}
		if val, ok := l["logout_uris"]; ok {
			str := expandStringList(val.([]interface{}))
			ca.LogoutUris = &str
		}
		if val, ok := l["ping_access_logout_capable"]; ok {
			ca.PingAccessLogoutCapable = Bool(val.(bool))
		}
		if val, ok := l["policy_group"]; ok {
			ca.PolicyGroup = expandResourceLink(val.([]interface{}))
		}
	}
	return ca
}

func flattenClientOIDCPolicy(in *pf.ClientOIDCPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.GrantAccessSessionRevocationApi != nil {
		s["grant_access_session_revocation_api"] = *in.GrantAccessSessionRevocationApi
	}
	if in.IdTokenSigningAlgorithm != nil {
		s["id_token_signing_algorithm"] = *in.IdTokenSigningAlgorithm
	}
	if in.LogoutUris != nil && len(*in.LogoutUris) > 0 {
		s["logout_uris"] = flattenStringList(*in.LogoutUris)
	}
	if in.PingAccessLogoutCapable != nil {
		s["ping_access_logout_capable"] = *in.PingAccessLogoutCapable
	}
	if in.PolicyGroup != nil {
		s["policy_group"] = flattenResourceLink(in.PolicyGroup)
	}
	m = append(m, s)
	return m
}

func flattenConfigField(in []*pf.ConfigField) *schema.Set {
	m := []interface{}{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		if v.Value != nil {
			s["value"] = *v.Value
		}
		if v.EncryptedValue != nil && *v.EncryptedValue != "" {
			s["encrypted_value"] = *v.EncryptedValue
		}
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return schema.NewSet(configFieldHash, m)
}

func configFieldHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["name"].(string))
	if d, ok := m["value"]; ok && d.(string) != "" {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["encrypted_value"]; ok && d.(string) != "" {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["inherited"]; ok {
		buf.WriteString(fmt.Sprintf("%t-", d.(bool)))
	}
	return hashcode.String(buf.String())
}

func expandConfigField(in []interface{}) *[]*pf.ConfigField {
	configFields := []*pf.ConfigField{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ConfigField{
			Name: String(l["name"].(string)),
		}
		if val, ok := l["value"]; ok {
			s.Value = String(val.(string))
		}
		if val, ok := l["encrypted_value"]; ok {
			s.EncryptedValue = String(val.(string))
		}
		if val, ok := l["inherited"]; ok {
			s.Inherited = Bool(val.(bool))
		}
		configFields = append(configFields, s)
	}
	return &configFields
}

func flattenConfigRow(in []*pf.ConfigRow) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["fields"] = flattenConfigField(*v.Fields)
		m = append(m, s)
	}
	return m
}

func expandConfigRow(in []interface{}) *[]*pf.ConfigRow {
	configRows := []*pf.ConfigRow{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ConfigRow{
			Fields: expandConfigField(l["name"].([]interface{})),
		}
		configRows = append(configRows, s)
	}
	return &configRows
}

func flattenConfigTable(in []*pf.ConfigTable) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		// if v.Rows != nil {
		// 	s["rows"] = flattenConfigRow(*v.Rows)
		// }
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return m
}

func expandConfigTable(in []interface{}) *[]*pf.ConfigTable {
	configTables := []*pf.ConfigTable{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ConfigTable{
			Name: String(l["name"].(string)),
		}
		// if val, ok := l["rows"]; ok {
		// 	s.Rows = expandConfigRow(val.([]interface{}))
		// }
		if val, ok := l["inherited"]; ok {
			s.Inherited = Bool(val.(bool))
		}
		configTables = append(configTables, s)
	}
	return &configTables
}

func flattenPluginConfiguration(in *pf.PluginConfiguration) []interface{} {
	// m := []interface{}{}
	s := make(map[string]interface{})
	// if in.Tables != nil {
	// 	s["tables"] = flattenConfigTable(*in.Tables)
	// }
	if in.Fields != nil {
		s["fields"] = flattenConfigField(*in.Fields)
	}
	// for _, v := range cbs.Items {
	// 	s = append(s, flattenCacheBehaviorDeprecated(v))
	// }
	// return schema.NewSet(, []interface{}{s})
	// m := make([]map[string]interface{}, 0, 1)
	// s := make(map[string]interface{})
	// if in.Tables != nil {
	// 	s["tables"] = flattenConfigTable(*in.Tables)
	// }
	// if in.Fields != nil {
	// 	s["fields"] = flattenConfigField(*in.Fields)
	// }
	// m = append(m, s)
	// log.Printf("[INFO] PluginConfig: %s", m)
	return []interface{}{s}
}

func expandPluginConfiguration(in []interface{}) *pf.PluginConfiguration {
	config := &pf.PluginConfiguration{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["tables"]; ok {
			config.Tables = expandConfigTable(val.(*schema.Set).List())
		}
		if val, ok := l["fields"]; ok {
			config.Fields = expandConfigField(val.(*schema.Set).List())
		}
	}
	return config
}

func flattenAccessTokenAttributeContract(in *pf.AccessTokenAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["extended_attributes"] = flattenAccessTokenAttributes(*in.ExtendedAttributes)
	m = append(m, s)
	return m
}

func expandAccessTokenAttributeContract(in []interface{}) *pf.AccessTokenAttributeContract {
	pgc := &pf.AccessTokenAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		atr := []*pf.AccessTokenAttribute{}
		for _, exAtr := range l["extended_attributes"].([]interface{}) {
			atr = append(atr, &pf.AccessTokenAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func flattenAccessTokenAttributes(in []*pf.AccessTokenAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}
