package pingfederate

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func validatePersistentGrantLifetimeUnit(value interface{}, _ cty.Path) diag.Diagnostics {
	v := value.(string)
	switch v {
	case
		"MINUTES",
		"DAYS",
		"HOURS":
		return nil
	}
	return diag.Errorf("must be either 'MINUTES' or 'DAYS' or 'HOURS' not %s", v)
}

func validateGrantTypes(value interface{}, _ cty.Path) diag.Diagnostics {
	v := value.(string)
	switch v {
	case
		"IMPLICIT",
		"AUTHORIZATION_CODE",
		"RESOURCE_OWNER_CREDENTIALS",
		"CLIENT_CREDENTIALS",
		"REFRESH_TOKEN",
		"EXTENSION",
		"ACCESS_TOKEN_VALIDATION":
		return nil
	}
	return diag.Errorf("must be either 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' or 'CLIENT_CREDENTIALS' or 'REFRESH_TOKEN' or 'EXTENSION' or 'ACCESS_TOKEN_VALIDATION' not %s", v)
}

func validateClientAuthType(value interface{}, _ cty.Path) diag.Diagnostics {
	v := value.(string)
	switch v {
	case
		"SECRET",
		"CERTIFICATE",
		"PRIVATE_KEY_JWT":
		return nil
	}
	return diag.Errorf("must be either 'SECRET' or 'CERTIFICATE' or 'PRIVATE_KEY_JWT' not %s", v)
}

func validateTokenSigningAlgorithm(value interface{}, _ cty.Path) diag.Diagnostics {
	v := value.(string)
	switch v {
	case
		"NONE",
		"HS256",
		"HS384",
		"HS512",
		"RS256",
		"RS384",
		"RS512",
		"ES256",
		"ES384",
		"ES512",
		"PS256",
		"PS384",
		"PS512":
		return nil
	}
	return diag.Errorf("must be either 'NONE' or 'HS256' or 'HS384' or 'HS512' or 'RS256' or 'RS384' or 'RS512' or 'ES256' or 'ES384' or 'ES512' or 'PS256' or 'PS384' or 'PS512' not %s", v)
}
