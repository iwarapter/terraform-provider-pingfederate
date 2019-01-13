package pingfederate

import (
	"testing"

	"github.com/hashicorp/terraform/flatmap"
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
