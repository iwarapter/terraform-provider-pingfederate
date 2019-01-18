package pingfederate

import "fmt"

func validatePersistentGrantLifetimeUnit(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	switch v {
	case
		"MINUTES",
		"DAYS",
		"HOURS":
		return
	}
	errs = append(errs, fmt.Errorf("%q must be either 'MINUTES' or 'DAYS' or 'HOURS' not %s", field, v))
	return
}

func validateGrantTypes(value interface{}, field string) (warns []string, errs []error) {
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
		return
	}
	errs = append(errs, fmt.Errorf("%q must be either 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' or 'CLIENT_CREDENTIALS' or 'REFRESH_TOKEN' or 'EXTENSION' or 'ACCESS_TOKEN_VALIDATION' not %s", field, v))
	return
}

func validateClientAuthType(value interface{}, field string) (warns []string, errs []error) {
	v := value.(string)
	switch v {
	case
		"SECRET",
		"CERTIFICATE",
		"PRIVATE_KEY_JWT":
		return
	}
	errs = append(errs, fmt.Errorf("%q must be either 'SECRET' or 'CERTIFICATE' or 'PRIVATE_KEY_JWT' not %s", field, v))
	return
}

func validateTokenSigningAlgorithm(value interface{}, field string) (warns []string, errs []error) {
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
		return
	}
	errs = append(errs, fmt.Errorf("%q must be either 'NONE' or 'HS256' or 'HS384' or 'HS512' or 'RS256' or 'RS384' or 'RS512' or 'ES256' or 'ES384' or 'ES512' or 'PS256' or 'PS384' or 'PS512' not %s", field, v))
	return
}
