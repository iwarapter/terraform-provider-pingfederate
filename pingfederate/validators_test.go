package pingfederate

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func Test_validatePersistentGrantLifetimeUnit(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "HOURS passes",
			value:         "HOURS",
			expectedDiags: nil,
		},
		{
			name:          "MINUTES passes",
			value:         "MINUTES",
			expectedDiags: nil,
		},
		{
			name:          "DAYS passes",
			value:         "DAYS",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'MINUTES' or 'DAYS' or 'HOURS' not %s", "other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validatePersistentGrantLifetimeUnit(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateGrantTypes(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "IMPLICIT passes",
			value:         "IMPLICIT",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' or 'CLIENT_CREDENTIALS' or 'REFRESH_TOKEN' or 'EXTENSION' or 'ACCESS_TOKEN_VALIDATION' not %s", "other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateGrantTypes(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateTokenSigningAlgorithm(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "HS256 passes",
			value:         "HS256",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'NONE' or 'HS256' or 'HS384' or 'HS512' or 'RS256' or 'RS384' or 'RS512' or 'ES256' or 'ES384' or 'ES512' or 'PS256' or 'PS384' or 'PS512' not %s", "other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateTokenSigningAlgorithm(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}

func Test_validateClientAuthType(t *testing.T) {
	tests := []struct {
		name          string
		value         interface{}
		expectedDiags diag.Diagnostics
	}{
		{
			name:          "SECRET passes",
			value:         "SECRET",
			expectedDiags: nil,
		},
		{
			name:          "junk does not pass",
			value:         "other",
			expectedDiags: diag.Errorf("must be either 'SECRET' or 'CERTIFICATE' or 'PRIVATE_KEY_JWT' not %s", "other"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := validateClientAuthType(tc.value, cty.Path{})
			if len(diags) != len(tc.expectedDiags) {
				t.Fatalf("%s: wrong number of diags, expected %d, got %d", tc.name, len(tc.expectedDiags), len(diags))
			}
			for j := range diags {
				if diags[j].Severity != tc.expectedDiags[j].Severity {
					t.Fatalf("%s: expected severity %v, got %v", tc.name, tc.expectedDiags[j].Severity, diags[j].Severity)
				}
				if !diags[j].AttributePath.Equals(tc.expectedDiags[j].AttributePath) {
					t.Fatalf("%s: attribute paths do not match expected: %v, got %v", tc.name, tc.expectedDiags[j].AttributePath, diags[j].AttributePath)
				}
				if diags[j].Summary != tc.expectedDiags[j].Summary {
					t.Fatalf("%s: summary does not match expected: %v, got %v", tc.name, tc.expectedDiags[j].Summary, diags[j].Summary)
				}
			}
		})
	}
}
