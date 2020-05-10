package pingfederate

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func setOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func requiredListOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

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
				"location": {
					Type:     schema.TypeString,
					Computed: true,
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
					Type:     schema.TypeList,
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
			"rows": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceConfigRow(),
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceConfigRow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceConfigField(),
			},
			"sensitive_fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceSensitiveConfigField(),
			},
		},
	}
}

func resourceSensitiveConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
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
				Required: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourcePasswordCredentialValidatorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"core_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extended_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIdpAdapterAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": &schema.Schema{
				Type:    schema.TypeBool,
				Default: false,
			},
			"core_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem:     resourceIdpAdapterAttribute(),
			},
			"mask_ognl_values": &schema.Schema{
				Type:    schema.TypeBool,
				Default: false,
			},
			"extended_attributes": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem:     resourceIdpAdapterAttribute(),
			},
		},
	}
}

func resourceIdpAdapterAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of this attribute.",
				Required:    true,
			},
			"pseudonym": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this attribute is used to construct a pseudonym for the SP. Defaults to false.",
				Default:     false,
			},
			"masked": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this attribute is masked in PingFederate logs. Defaults to false.",
				Default:     false,
			},
		},
	}
}

func resourceIdpAdapterAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			//TODO
			"attribute_sources": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ldap_attribute_source": {
							Type: schema.TypeList,
							Elem: resourceLdapAttributeSource(),
						},
						"fields": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem:     resourceConfigField(),
						},
					},
				},
			},
			//"core_attributes": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Required: true,
			//	MinItems: 1,
			//	Elem:     resourceIdpAdapterAttribute(),
			//},
			//"mask_ognl_values": &schema.Schema{
			//	Type:    schema.TypeBool,
			//	Default: false,
			//},
			//"extended_attributes": &schema.Schema{
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	MinItems: 1,
			//	Elem:     resourceIdpAdapterAttribute(),
			//},
		},
	}
}

func resourceLdapAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_store_ref": resourceLinkSchema(),
			"base_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			//"attribute_contract_fulfillment": {
			//	Type: schema.TypeMap,
			//	Optional: true,
			//},
			//"binary_attribute_settings": {
			//	Type: schema.TypeMap,
			//	Optional: true,
			//},
			"member_of_nested_group": {
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

func flattenStringSet(list []*string) *schema.Set {
	return schema.NewSet(schema.HashString, flattenStringList(list))
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
		// if val, ok := l["location"]; ok {
		// 	ca.Location = String(val.(string))
		// }
	}
	return ca
}

func flattenResourceLink(in *pf.ResourceLink) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	// if in.Location != nil {
	// 	s["location"] = *in.Location
	// }
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
		if v.EncryptedValue != nil {
			continue
		}
		s := make(map[string]interface{})
		s["name"] = *v.Name
		//We check if the Encrypted value is set, if its not we can update the value as a normal password field
		//will not return the value so we need to not overwrite it, which unfortunely means we cannot track password changes
		//this is a limitation of ping federate.
		if v.Value != nil && v.EncryptedValue == nil {
			s["value"] = *v.Value
		}
		// if v.EncryptedValue != nil && *v.EncryptedValue != "" {
		// 	s["encrypted_value"] = *v.EncryptedValue
		// }
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return schema.NewSet(configFieldHash, m)
}

func flattenSensitiveConfigField(in []*pf.ConfigField) *schema.Set {
	m := []interface{}{}
	for _, v := range in {
		if v.EncryptedValue == nil {
			continue
		}
		s := make(map[string]interface{})
		s["name"] = *v.Name
		//We check if the Encrypted value is set, if its not we can update the value as a normal password field
		//will not return the value so we need to not overwrite it, which unfortunely means we cannot track password changes
		//this is a limitation of ping federate.
		//if v.Value != nil && v.EncryptedValue == nil {
		if v.Value != nil {
			s["value"] = *v.Value
		}
		// if v.EncryptedValue != nil && *v.EncryptedValue != "" {
		// 	s["encrypted_value"] = *v.EncryptedValue
		// }
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
	// if d, ok := m["encrypted_value"]; ok && d.(string) != "" {
	// 	buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	// }
	if d, ok := m["inherited"]; ok {
		buf.WriteString(fmt.Sprintf("%t-", d.(bool)))
	}
	return hashcode.String(buf.String())
}

func expandConfigFields(in []interface{}) *[]*pf.ConfigField {
	configFields := []*pf.ConfigField{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if _, ok := l["encrypted_value"]; !ok {
			s := &pf.ConfigField{
				Name: String(l["name"].(string)),
			}
			if val, ok := l["value"]; ok {
				s.Value = String(val.(string))
			}
			if val, ok := l["inherited"]; ok {
				s.Inherited = Bool(val.(bool))
			}
			configFields = append(configFields, s)
		}
	}
	return &configFields
}

func expandSensitiveConfigFields(in []interface{}) *[]*pf.ConfigField {
	configFields := []*pf.ConfigField{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		log.Printf("[DEBUG] HELPER: expand sensitive fields: %v", l)
		if val, ok := l["value"]; ok && val.(string) != "" {
			s := &pf.ConfigField{
				Name: String(l["name"].(string)),
			}
			if val, ok := l["value"]; ok {
				s.Value = String(val.(string))
			}
			if val, ok := l["inherited"]; ok {
				s.Inherited = Bool(val.(bool))
			}
			configFields = append(configFields, s)
		}
	}
	return &configFields
}

func flattenConfigRow(in []*pf.ConfigRow) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["fields"] = flattenConfigField(*v.Fields)
		s["sensitive_fields"] = flattenSensitiveConfigField(*v.Fields)
		m = append(m, s)
	}
	return m
}

func expandConfigRow(in []interface{}) *[]*pf.ConfigRow {
	configRows := []*pf.ConfigRow{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		row := &pf.ConfigRow{}
		if val, ok := l["fields"]; ok {
			row.Fields = expandConfigFields(val.(*schema.Set).List())
		}
		if val, ok := l["sensitive_fields"]; ok {
			fields := expandSensitiveConfigFields(val.(*schema.Set).List())
			*row.Fields = append(*row.Fields, *fields...)
		}
		configRows = append(configRows, row)
	}
	return &configRows
}

func flattenConfigTable(in []*pf.ConfigTable) []interface{} {
	var m []interface{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		if v.Rows != nil {
			s["rows"] = flattenConfigRow(*v.Rows)
		}
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return m
	//return schema.NewSet(configTableHash, m)
}

func configTableHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["name"].(string))
	return hashcode.String(buf.String())
}

func expandConfigTable(in []interface{}) *[]*pf.ConfigTable {
	configTables := []*pf.ConfigTable{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ConfigTable{
			Name: String(l["name"].(string)),
		}
		if val, ok := l["rows"]; ok {
			s.Rows = expandConfigRow(val.([]interface{}))
		}
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
	if in.Tables != nil {
		s["tables"] = flattenConfigTable(*in.Tables)
	}
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
			config.Tables = expandConfigTable(val.([]interface{}))
		}
		//if val, ok := l["fields"]; ok {
		//	config.Fields = expandConfigFields(val.(*schema.Set).List())
		//	log.Printf("[INFO] ConfigFields: %v", len(*config.Fields))
		//	*config.Fields = append(*config.Fields, *expandSensitiveConfigFields(val.(*schema.Set).List())...)
		//	log.Printf("[INFO] ConfigFieldsWithSensitive: %v", len(*config.Fields))
		//}
		if val, ok := l["fields"]; ok {
			config.Fields = expandConfigFields(val.(*schema.Set).List())
		}
		if val, ok := l["sensitive_fields"]; ok {
			fields := expandSensitiveConfigFields(val.(*schema.Set).List())
			*config.Fields = append(*config.Fields, *fields...)
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

func flattenAuthenticationPolicyContractAttribute(in []*pf.AuthenticationPolicyContractAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func expandAuthenticationPolicyContractAttribute(in []interface{}) *[]*pf.AuthenticationPolicyContractAttribute {
	contractList := []*pf.AuthenticationPolicyContractAttribute{}
	for _, raw := range in {
		c := &pf.AuthenticationPolicyContractAttribute{
			Name: String(raw.(string)),
		}
		contractList = append(contractList, c)
	}
	return &contractList
}

func flattenPasswordCredentialValidatorAttribute(in []*pf.PasswordCredentialValidatorAttribute) *schema.Set {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return schema.NewSet(schema.HashString, m)
}

func expandPasswordCredentialValidatorAttribute(in []interface{}) *[]*pf.PasswordCredentialValidatorAttribute {
	contractList := []*pf.PasswordCredentialValidatorAttribute{}
	for _, raw := range in {
		c := &pf.PasswordCredentialValidatorAttribute{
			Name: String(raw.(string)),
		}
		contractList = append(contractList, c)
	}
	return &contractList
}

func expandPasswordCredentialValidatorAttributeContract(in []interface{}) *pf.PasswordCredentialValidatorAttributeContract {
	pgc := &pf.PasswordCredentialValidatorAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["extended_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			pgc.ExtendedAttributes = expandPasswordCredentialValidatorAttribute(v.(*schema.Set).List())
		}
		if v, ok := l["core_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			pgc.CoreAttributes = expandPasswordCredentialValidatorAttribute(v.(*schema.Set).List())
		}
	}
	return pgc
}

func flattenPasswordCredentialValidatorAttributeContract(in *pf.PasswordCredentialValidatorAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenPasswordCredentialValidatorAttribute(*in.ExtendedAttributes)
	}
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenPasswordCredentialValidatorAttribute(*in.CoreAttributes)
	}
	m = append(m, s)
	return m
}

func flattenJdbcDataStore(in *pf.JdbcDataStore) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.MaskAttributeValues != nil {
		s["mask_attribute_values"] = *in.MaskAttributeValues
	}
	if in.ConnectionUrlTags != nil && len(*in.ConnectionUrlTags) != 0 {
		//connection_url_tags
	}
	if in.ConnectionUrl != nil {
		s["connection_url"] = *in.ConnectionUrl
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.DriverClass != nil {
		s["driver_class"] = *in.DriverClass
	}
	if in.UserName != nil {
		s["user_name"] = *in.UserName
	}
	if in.Password != nil {
		//TODO i need to handle this not being set
		s["password"] = *in.Password
	}
	if in.EncryptedPassword != nil {
		s["encrypted_password"] = *in.EncryptedPassword
	}
	if in.ValidateConnectionSql != nil {
		s["validate_connection_sql"] = *in.ValidateConnectionSql
	}
	if in.AllowMultiValueAttributes != nil {
		s["allow_multi_value_attributes"] = *in.AllowMultiValueAttributes
	}
	if in.MinPoolSize != nil {
		s["min_pool_size"] = *in.MinPoolSize
	}
	if in.MaxPoolSize != nil {
		s["max_pool_size"] = *in.MaxPoolSize
	}
	if in.BlockingTimeout != nil {
		s["blocking_timeout"] = *in.BlockingTimeout
	}
	if in.IdleTimeout != nil {
		s["idle_timeout"] = *in.IdleTimeout
	}
	//s["type"] = String("JDBC")
	m = append(m, s)
	return m
}

func expandJdbcDataStore(in []interface{}) *pf.JdbcDataStore {
	ca := &pf.JdbcDataStore{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["mask_attribute_values"]; ok {
			ca.MaskAttributeValues = Bool(val.(bool))
		}
		//TODO connection_url_tags
		if val, ok := l["connection_url"]; ok {
			ca.ConnectionUrl = String(val.(string))
		}
		if val, ok := l["name"]; ok {
			ca.Name = String(val.(string))
		}
		if val, ok := l["driver_class"]; ok {
			ca.DriverClass = String(val.(string))
		}
		if val, ok := l["user_name"]; ok {
			ca.UserName = String(val.(string))
		}
		if val, ok := l["password"]; ok {
			ca.Password = String(val.(string))
		}
		if val, ok := l["encrypted_password"]; ok {
			ca.EncryptedPassword = String(val.(string))
		}
		if val, ok := l["validate_connection_sql"]; ok {
			ca.ValidateConnectionSql = String(val.(string))
		}
		if val, ok := l["allow_multi_value_attributes"]; ok {
			ca.AllowMultiValueAttributes = Bool(val.(bool))
		}
		if val, ok := l["min_pool_size"]; ok {
			ca.MinPoolSize = Int(val.(int))
		}
		if val, ok := l["max_pool_size"]; ok {
			ca.MaxPoolSize = Int(val.(int))
		}
		if val, ok := l["blocking_timeout"]; ok {
			ca.BlockingTimeout = Int(val.(int))
		}
		if val, ok := l["idle_timeout"]; ok {
			ca.IdleTimeout = Int(val.(int))
		}
		ca.Type = String("JDBC")
	}
	return ca
}

func flattenLdapDataStore(in *pf.LdapDataStore) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.MaskAttributeValues != nil {
		s["mask_attribute_values"] = *in.MaskAttributeValues
	}
	if in.HostnamesTags != nil && len(*in.HostnamesTags) != 0 {
		//TODO connection_url_tags
	}
	if in.Hostnames != nil {
		s["hostnames"] = flattenStringList(*in.Hostnames)
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.LdapType != nil {
		s["ldap_type"] = *in.LdapType
	}
	if in.BindAnonymously != nil {
		s["bind_anonymously"] = *in.BindAnonymously
	}
	if in.UserDN != nil {
		s["user_dn"] = *in.UserDN
	}
	if in.Password != nil {
		s["password"] = *in.Password
	}
	if in.EncryptedPassword != nil {
		s["encrypted_password"] = *in.EncryptedPassword
	}
	if in.UseSsl != nil {
		s["use_ssl"] = *in.UseSsl
	}
	if in.UseDnsSrvRecords != nil {
		s["use_dns_srv_records"] = *in.UseDnsSrvRecords
	}
	if in.FollowLDAPReferrals != nil {
		s["follow_ldap_referrals"] = *in.FollowLDAPReferrals
	}
	if in.TestOnBorrow != nil {
		s["test_on_borrow"] = *in.TestOnBorrow
	}
	if in.TestOnReturn != nil {
		s["test_on_return"] = *in.TestOnReturn
	}
	if in.CreateIfNecessary != nil {
		s["create_if_necessary"] = *in.CreateIfNecessary
	}
	if in.VerifyHost != nil {
		s["verify_host"] = *in.VerifyHost
	}
	if in.MinConnections != nil {
		s["min_connections"] = *in.MinConnections
	}
	if in.MaxConnections != nil {
		s["max_connections"] = *in.MaxConnections
	}
	if in.MaxWait != nil {
		s["max_wait"] = *in.MaxWait
	}
	if in.TimeBetweenEvictions != nil {
		s["time_between_evictions"] = *in.TimeBetweenEvictions
	}
	if in.ReadTimeout != nil {
		s["read_timeout"] = *in.ReadTimeout
	}
	if in.ConnectionTimeout != nil {
		s["connection_timeout"] = *in.ConnectionTimeout
	}
	if in.DnsTtl != nil {
		s["dns_ttl"] = *in.DnsTtl
	}
	if in.LdapDnsSrvPrefix != nil {
		s["ldap_dns_srv_prefix"] = *in.LdapDnsSrvPrefix
	}
	if in.LdapsDnsSrvPrefix != nil {
		s["ldaps_dns_srv_prefix"] = *in.LdapsDnsSrvPrefix
	}
	if in.BinaryAttributes != nil {
		s["binary_attributes"] = flattenStringList(*in.BinaryAttributes)
	}
	//s["type"] = String("LDAP")
	m = append(m, s)
	return m
}

func expandLdapDataStore(in []interface{}) *pf.LdapDataStore {
	ca := &pf.LdapDataStore{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		//log.Printf("[DEBUG] HELPER expandLdapDataStore: %v", l)
		if val, ok := l["mask_attribute_values"]; ok {
			ca.MaskAttributeValues = Bool(val.(bool))
		}
		//TODO hostnames_tags
		if val, ok := l["hostnames"]; ok {
			strs := expandStringList(val.(*schema.Set).List())
			ca.Hostnames = &strs
		}
		if val, ok := l["name"]; ok {
			ca.Name = String(val.(string))
		}
		if val, ok := l["ldap_type"]; ok {
			ca.LdapType = String(val.(string))
		}
		if val, ok := l["bind_anonymously"]; ok {
			ca.BindAnonymously = Bool(val.(bool))
		}
		if val, ok := l["user_dn"]; ok {
			ca.UserDN = String(val.(string))
		}
		if val, ok := l["password"]; ok {
			ca.Password = String(val.(string))
		}
		if val, ok := l["encrypted_password"]; ok {
			ca.EncryptedPassword = String(val.(string))
		}
		if val, ok := l["use_ssl"]; ok {
			ca.UseSsl = Bool(val.(bool))
		}
		if val, ok := l["use_dns_srv_records"]; ok {
			ca.UseDnsSrvRecords = Bool(val.(bool))
		}
		if val, ok := l["follow_ldap_referrals"]; ok {
			ca.FollowLDAPReferrals = Bool(val.(bool))
		}
		if val, ok := l["test_on_borrow"]; ok {
			ca.TestOnBorrow = Bool(val.(bool))
		}
		if val, ok := l["test_on_return"]; ok {
			ca.TestOnReturn = Bool(val.(bool))
		}
		if val, ok := l["create_if_necessary"]; ok {
			ca.CreateIfNecessary = Bool(val.(bool))
		}
		if val, ok := l["verify_host"]; ok {
			ca.VerifyHost = Bool(val.(bool))
		}
		if val, ok := l["min_connections"]; ok {
			ca.MinConnections = Int(val.(int))
		}
		if val, ok := l["max_connections"]; ok {
			ca.MaxConnections = Int(val.(int))
		}
		if val, ok := l["max_wait"]; ok {
			ca.MaxWait = Int(val.(int))
		}
		if val, ok := l["time_between_evictions"]; ok {
			ca.TimeBetweenEvictions = Int(val.(int))
		}
		if val, ok := l["read_timeout"]; ok {
			ca.ReadTimeout = Int(val.(int))
		}
		if val, ok := l["connection_timeout"]; ok {
			ca.ConnectionTimeout = Int(val.(int))
		}
		if val, ok := l["dns_ttl"]; ok {
			ca.DnsTtl = Int(val.(int))
		}
		if val, ok := l["ldap_dns_srv_prefix"]; ok {
			ca.LdapDnsSrvPrefix = String(val.(string))
		}
		if val, ok := l["ldaps_dns_srv_prefix"]; ok {
			ca.LdapsDnsSrvPrefix = String(val.(string))
		}
		if val, ok := l["binary_attributes"]; ok {
			strs := expandStringList(val.(*schema.Set).List())
			ca.BinaryAttributes = &strs
		}
		ca.Type = String("LDAP")
	}
	return ca
}

func maskPluginConfigurationFromDescriptor(desc *pf.PluginConfigDescriptor, origConf, conf *pf.PluginConfiguration) []interface{} {
	//printPluginConfig("originConf",origConf)
	//printPluginConfig("conf",conf)

	for _, f := range *desc.Fields {
		if *f.Type == "HASHED_TEXT" {
			for _, i := range *conf.Fields {
				if *i.Name == *f.Name {
					*i.Value, _ = getConfigFieldValueByName(*i.Name, origConf.Fields)
				}
			}
		}
	}

	for _, dt := range *desc.Tables {
		for _, dc := range *dt.Columns {
			if *dc.Type == "HASHED_TEXT" {
				for ctIndex, ct := range *conf.Tables {
					for crIndex, cr := range *ct.Rows {
						for _, f := range *cr.Fields {
							if *f.Name == *dc.Name {
								//log.Printf("[DEBUG] HELPER: getConfigFieldValueByName ctIndex: %d crIndex: %d", ctIndex, crIndex)
								//log.Printf("[DEBUG] HELPER: %s", *(*origConf.Tables)[ctIndex].Name)
								//log.Printf("[DEBUG] HELPER: %v", (*(*origConf.Tables)[ctIndex].Rows)[crIndex])
								//log.Printf("[DEBUG] HELPER: %v", *(*(*(*origConf.Tables)[ctIndex].Rows)[crIndex].Fields)[0].Name)
								val, _ := getConfigFieldValueByName(*f.Name, (*(*origConf.Tables)[ctIndex].Rows)[crIndex].Fields)
								f.Value = &val
								//*f.Value, _ = getConfigFieldValueByName(*f.Name, (*(*origConf.Tables)[ctIndex].Rows)[crIndex].Fields)
							}
						}
					}
				}
			}
		}
	}
	return flattenPluginConfiguration(conf)
}

func getConfigFieldValueByName(name string, fields *[]*pf.ConfigField) (string, error) {
	for _, f := range *fields {
		if *f.Name == name {
			return *f.Value, nil
		}
	}
	return "", nil
}

func printPluginConfig(name string, conf *pf.PluginConfiguration) {
	log.Printf("[DEBUG] Helper: %s", name)
	log.Printf("[DEBUG] Helper: %s Fields: %d", name, len(*conf.Fields))
	for _, f := range *conf.Fields {
		log.Printf("[DEBUG] Helper: %s Field: %s: Value: %s EncryptedValue: %s", name, *f.Name, *f.Value, *f.EncryptedValue)
	}
	log.Printf("[DEBUG] Helper: %s Tables: %d", name, len(*conf.Tables))
	for _, t := range *conf.Tables {
		log.Printf("[DEBUG] Helper: %s Table: %s", name, *t.Name)
		log.Printf("[DEBUG] Helper: %s Table: %s Rows: %d", name, *t.Name, len(*t.Rows))
		for _, r := range *t.Rows {
			for i, f := range *r.Fields {
				if f.Value != nil {
					log.Printf("[DEBUG] Helper: %s Table: %s Row: %d Field: %s Value: %s", name, *t.Name, i, *f.Name, *f.Value)
				}
				if f.EncryptedValue != nil {
					log.Printf("[DEBUG] Helper: %s Table: %s Row: %d Field: %s EncryptedValue: %s", name, *t.Name, i, *f.Name, *f.EncryptedValue)
				}
			}
		}

	}
}

func resourceAuthenticationSelectorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}


func flattenAuthenticationSelectorAttributeContract(in *pf.AuthenticationSelectorAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["extended_attributes"] = flattenAuthenticationSelectorAttributes(*in.ExtendedAttributes)
	m = append(m, s)
	return m
}

func expandAuthenticationSelectorAttributeContract(in []interface{}) *pf.AuthenticationSelectorAttributeContract {
	pgc := &pf.AuthenticationSelectorAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		atr := []*pf.AuthenticationSelectorAttribute{}
		for _, exAtr := range l["extended_attributes"].([]interface{}) {
			atr = append(atr, &pf.AuthenticationSelectorAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func flattenAuthenticationSelectorAttributes(in []*pf.AuthenticationSelectorAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}