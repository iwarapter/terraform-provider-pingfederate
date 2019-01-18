package pingfederate

import (
	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourceLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
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
			ca.PolicyGroup = expandResourceLink(val.(*schema.Set).List())
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
