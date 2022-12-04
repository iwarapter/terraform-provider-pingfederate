package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func legacyResourceLinkSchema() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"id": {
			Type:     types.StringType,
			Optional: true,
		},
		"location": {
			Type:     types.StringType,
			Computed: true,
		},
	}
}

func legacyResourceParameterValues() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"key_name": {
			Type:     types.StringType,
			Optional: true,
		},
		"values": {
			Type:     types.SetType{ElemType: types.StringType},
			Optional: true,
		},
	}
}
