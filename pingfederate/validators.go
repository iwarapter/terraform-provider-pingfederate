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

func validatePersistentGrantReuseGrantTypes(value interface{}, field string) (warns []string, errs []error) {
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
