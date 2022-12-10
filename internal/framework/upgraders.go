package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func legacyResourceLinkSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Optional: true,
		},
		"location": schema.StringAttribute{
			Computed: true,
		},
	}
}

func legacyResourceParameterValues() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key_name": schema.StringAttribute{Optional: true},
				"values":   schema.SetAttribute{Optional: true, ElementType: types.StringType},
			},
		},
	}
}
