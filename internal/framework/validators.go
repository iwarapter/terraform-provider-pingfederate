package framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func EnumValidation(enums ...string) tfsdk.AttributeValidator {
	return enumValidator{enums: enums}
}

type enumValidator struct {
	enums []string
}

func (v enumValidator) Description(ctx context.Context) string {
	return "we need to validate some enums"
}

func (v enumValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v enumValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
	var str string
	response.Diagnostics.Append(request.Config.GetAttribute(ctx, request.AttributePath, &str)...)
	if response.Diagnostics.HasError() {
		return
	}
	for _, enum := range v.enums {
		if request.AttributeConfig.Equal(types.StringValue(enum)) {
			return
		}
	}
	response.Diagnostics.AddAttributeError(request.AttributePath, "Enum Validation Failure", fmt.Sprintf("should be one of: %s", strings.Join(v.enums, ", ")))
}
