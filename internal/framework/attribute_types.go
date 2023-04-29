package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var certViewAttrTypes = map[string]attr.Type{
	"crypto_provider":           types.StringType,
	"expires":                   types.StringType,
	"id":                        types.StringType,
	"issuer_dn":                 types.StringType,
	"key_algorithm":             types.StringType,
	"key_size":                  types.NumberType,
	"serial_number":             types.StringType,
	"sha1fingerprint":           types.StringType,
	"sha256fingerprint":         types.StringType,
	"signature_algorithm":       types.StringType,
	"status":                    types.StringType,
	"subject_alternative_names": types.ListType{ElemType: types.StringType},
	"subject_dn":                types.StringType,
	"valid_from":                types.StringType,
	"version":                   types.NumberType,
}
