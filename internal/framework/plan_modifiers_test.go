package framework

import (
	"context"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

func TestDefaultValue(t *testing.T) {
	t.Parallel()
	t.Skip("")

	type testCase struct {
		plannedValue  attr.Value
		currentValue  attr.Value
		defaultValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		"non-default non-Null string": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringValue("beta"),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringValue("gamma"),
		},
		"non-default non-Null string, current Null": {
			plannedValue:  types.StringValue("gamma"),
			currentValue:  types.StringNull(),
			defaultValue:  types.StringValue("alpha"),
			expectedValue: types.StringValue("gamma"),
		},
		//"non-default Null string, current Null": {
		//	plannedValue:  types.String{Null: true},
		//	currentValue:  types.StringValue: "beta"},
		//	defaultValue:  types.StringValue: "alpha"},
		//	expectedValue: types.String{Null: true},
		//},
		//"default string": {
		//	plannedValue:  types.String{Null: true},
		//	currentValue:  types.StringValue: "alpha"},
		//	defaultValue:  types.StringValue: "alpha"},
		//	expectedValue: types.StringValue: "alpha"},
		//},
		//"non-default non-Null number": {
		//	plannedValue:  types.Number{Value: big.NewFloat(30)},
		//	currentValue:  types.Number{Value: big.NewFloat(10)},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Value: big.NewFloat(30)},
		//},
		//"non-default non-Null number, current Null": {
		//	plannedValue:  types.Number{Value: big.NewFloat(30)},
		//	currentValue:  types.Number{Null: true},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Value: big.NewFloat(30)},
		//},
		//"non-default Null number, current Null": {
		//	plannedValue:  types.Number{Null: true},
		//	currentValue:  types.Number{Value: big.NewFloat(10)},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Null: true},
		//},
		//"default number": {
		//	plannedValue:  types.Number{Null: true},
		//	currentValue:  types.Number{Value: big.NewFloat(-10)},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Value: big.NewFloat(-10)},
		//},
		//"non-default string list": {
		//	plannedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "POST"},
		//	}},
		//	currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "PUT"},
		//	}},
		//	defaultValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "POST"},
		//	}},
		//},
		//"non-default string list, current out of order": {
		//	plannedValue: types.List{ElemType: types.StringType, Null: true},
		//	currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "HEAD"},
		//		types.StringValue: "GET"},
		//	}},
		//	defaultValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.List{ElemType: types.StringType, Null: true},
		//},
		//"default string list": {
		//	plannedValue: types.List{ElemType: types.StringType, Null: true},
		//	currentValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	defaultValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.List{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//},
		//"non-default string set": {
		//	plannedValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "POST"},
		//	}},
		//	currentValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "PUT"},
		//	}},
		//	defaultValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "POST"},
		//	}},
		//},
		//"default string set, current out of order": {
		//	plannedValue: types.Set{ElemType: types.StringType, Null: true},
		//	currentValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "HEAD"},
		//		types.StringValue: "GET"},
		//	}},
		//	defaultValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "HEAD"},
		//		types.StringValue: "GET"},
		//	}},
		//},
		//"default string set": {
		//	plannedValue: types.Set{ElemType: types.StringType, Null: true},
		//	currentValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	defaultValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//	expectedValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{
		//		types.StringValue: "GET"},
		//		types.StringValue: "HEAD"},
		//	}},
		//},
		//"non-default object": {
		//	plannedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "gamma"},
		//		},
		//	},
		//	currentValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "beta"},
		//		},
		//	},
		//	defaultValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "alpha"},
		//		},
		//	},
		//	expectedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "gamma"},
		//		},
		//	},
		//},
		//"non-default object, different value": {
		//	plannedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Null: true,
		//	},
		//	currentValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "beta"},
		//		},
		//	},
		//	defaultValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "alpha"},
		//		},
		//	},
		//	expectedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Null: true,
		//	},
		//},
		//"default object": {
		//	plannedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Null: true,
		//	},
		//	currentValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "alpha"},
		//		},
		//	},
		//	defaultValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "alpha"},
		//		},
		//	},
		//	expectedValue: types.Object{
		//		AttrTypes: map[string]attr.Type{
		//			"value": types.StringType,
		//		},
		//		Attrs: map[string]attr.Value{
		//			"value": types.StringValue: "alpha"},
		//		},
		//	},
		//},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ModifyAttributePlanRequest{
				AttributePath:  path.Root("test"),
				AttributePlan:  test.plannedValue,
				AttributeState: test.currentValue,
			}
			response := tfsdk.ModifyAttributePlanResponse{}
			Default(test.defaultValue).Modify(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}

			if diff := cmp.Diff(response.AttributePlan, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func Test_setStringDefaultModifier_Modify(t *testing.T) {
	t.Parallel()

	type testCase struct {
		plannedValue  attr.Value
		currentValue  attr.Value
		defaultValue  attr.Value
		expectedValue attr.Value
		expectError   bool
	}
	tests := map[string]testCase{
		//"non-default non-Null number": {
		//	plannedValue:  types.Number{Value: big.NewFloat(30)},
		//	currentValue:  types.Number{Value: big.NewFloat(10)},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Value: big.NewFloat(30)},
		//},
		//"non-default non-Null number, current Null": {
		//	plannedValue:  types.Number{Value: big.NewFloat(30)},
		//	currentValue:  types.Number{Null: true},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Value: big.NewFloat(30)},
		//},
		//"non-default Null number, current Null": {
		//	plannedValue:  types.Number{Null: true},
		//	currentValue:  types.Number{Value: big.NewFloat(10)},
		//	defaultValue:  types.Number{Value: big.NewFloat(-10)},
		//	expectedValue: types.Number{Null: true},
		//},
		"default number": {
			plannedValue:  types.NumberNull(),
			currentValue:  types.NumberNull(),
			defaultValue:  types.NumberValue(big.NewFloat(-10)),
			expectedValue: types.NumberValue(big.NewFloat(-10)),
		},
		//"a default with no config": {
		//	plannedValue:  types.Set{ElemType: types.StringType, Elems: []attr.Value{types.StringValue: "bar"}}},
		//	currentValue:  types.Set{ElemType: types.StringType, Elems: []attr.Value{types.StringValue: "subject"}}},
		//	defaultValue:  types.Set{ElemType: types.StringType, Elems: []attr.Value{types.StringValue: "subject"}}},
		//	expectedValue: types.Set{ElemType: types.StringType, Elems: []attr.Value{types.StringValue: "bar"}}},
		//},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := tfsdk.ModifyAttributePlanRequest{
				AttributePath:   path.Root("test"),
				AttributePlan:   test.plannedValue,
				AttributeState:  test.currentValue,
				AttributeConfig: test.currentValue,
			}
			response := tfsdk.ModifyAttributePlanResponse{
				AttributePlan: test.plannedValue,
			}
			Default(test.defaultValue).Modify(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}

			assert.Equal(t, test.expectedValue, response.AttributePlan)
		})
	}
}
