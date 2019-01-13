package pingfederate

import (
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

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
