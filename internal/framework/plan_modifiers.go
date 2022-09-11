package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// defaultValueModifier is a plan modifier that sets a default value for a
// types.StringType attribute when it is not configured. The attribute must be
// marked as Optional and Computed. When setting the state during the resource
// Create, Read, or Update methods, this default value must also be included or
// the Terraform CLI will generate an error.
type defaultValueModifier struct {
	value attr.Value
}

// Description returns a plain text description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m defaultValueModifier) Description(ctx context.Context) string {
	return "Sets the default value for this attribute."
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior, suitable for a practitioner to understand its impact.
func (m defaultValueModifier) MarkdownDescription(ctx context.Context) string {
	return "Sets the default value for this attribute."
}

// Modify runs the logic of the plan modifier.
// Access to the configuration, plan, and state is available in `req`, while
// `resp` contains fields for updating the planned value, triggering resource
// replacement, and returning diagnostics.
func (m defaultValueModifier) Modify(
	ctx context.Context,
	req tfsdk.ModifyAttributePlanRequest,
	resp *tfsdk.ModifyAttributePlanResponse,
) {
	if req.AttributeConfig == nil || resp.AttributePlan == nil {
		return
	}

	// if configuration was provided, then don't use the default
	if !req.AttributeConfig.IsNull() {
		return
	}

	// If the plan is known and not null (for example due to another plan modifier),
	// don't set the default value
	if !resp.AttributePlan.IsUnknown() && !resp.AttributePlan.IsNull() {
		return
	}

	resp.AttributePlan = m.value
}

func Default(defaultValue attr.Value) defaultValueModifier {
	return defaultValueModifier{
		value: defaultValue,
	}
}
