package pingfederate

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_validatePersistentGrantLifetimeUnit(t *testing.T) {
	type args struct {
		value interface{}
		field string
	}
	tests := []struct {
		name      string
		args      args
		wantWarns []string
		wantErrs  []error
	}{
		{
			name: "HOURS passes",
			args: args{
				value: "HOURS",
				field: "persistent_grant_lifetime_unit",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "MINUTES passes",
			args: args{
				value: "MINUTES",
				field: "persistent_grant_lifetime_unit",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "DAYS passes",
			args: args{
				value: "DAYS",
				field: "persistent_grant_lifetime_unit",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "persistent_grant_lifetime_unit",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'MINUTES' or 'DAYS' or 'HOURS' not %s", "persistent_grant_lifetime_unit", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validatePersistentGrantLifetimeUnit(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validatePersistentGrantLifetimeUnit() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validatePersistentGrantLifetimeUnit() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateGrantTypes(t *testing.T) {
	type args struct {
		value interface{}
		field string
	}
	tests := []struct {
		name      string
		args      args
		wantWarns []string
		wantErrs  []error
	}{
		{
			name: "IMPLICIT passes",
			args: args{
				value: "IMPLICIT",
				field: "persistent_grant_reuse_grant_types",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "persistent_grant_reuse_grant_types",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' or 'CLIENT_CREDENTIALS' or 'REFRESH_TOKEN' or 'EXTENSION' or 'ACCESS_TOKEN_VALIDATION' not %s", "persistent_grant_reuse_grant_types", "other")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateGrantTypes(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validatePersistentGrantReuseGrantTypes() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validatePersistentGrantReuseGrantTypes() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func Test_validateTokenSigningAlgorithm(t *testing.T) {
	type args struct {
		value interface{}
		field string
	}
	tests := []struct {
		name      string
		args      args
		wantWarns []string
		wantErrs  []error
	}{
		{
			name: "HS256 passes",
			args: args{
				value: "HS256",
				field: "id_token_signing_algorithm",
			},
			wantWarns: nil,
			wantErrs:  nil,
		},
		{
			name: "junk does not pass",
			args: args{
				value: "other",
				field: "id_token_signing_algorithm",
			},
			wantWarns: nil,
			wantErrs:  []error{fmt.Errorf("%q must be either 'NONE' or 'HS256' or 'HS384' or 'HS512' or 'RS256' or 'RS384' or 'RS512' or 'ES256' or 'ES384' or 'ES512' or 'PS256' or 'PS384' or 'PS512' not %s", "id_token_signing_algorithm", "other")},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWarns, gotErrs := validateTokenSigningAlgorithm(tt.args.value, tt.args.field)
			if !reflect.DeepEqual(gotWarns, tt.wantWarns) {
				t.Errorf("validateTokenSigningAlgorithm() gotWarns = %v, want %v", gotWarns, tt.wantWarns)
			}
			if !reflect.DeepEqual(gotErrs, tt.wantErrs) {
				t.Errorf("validateTokenSigningAlgorithm() gotErrs = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}
