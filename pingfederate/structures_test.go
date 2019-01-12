package pingfederate

import (
	"testing"

	"github.com/hashicorp/terraform/flatmap"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

// Returns test configuration
func testConf() map[string]string {
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
	expanded := flatmap.Expand(testConf(), "scopes").([]interface{})
	expandScopes := expandScopes(expanded)

	equals(t, 5, len(*expandScopes))
}
