package framework

import (
	"context"
	"math/big"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

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
