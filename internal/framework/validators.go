package framework

//
//type FakeEnumValidator struct{}
//
//func (f FakeEnumValidator) Description(ctx context.Context) string {
//	return "we need to validate some enums"
//}
//
//func (f FakeEnumValidator) MarkdownDescription(ctx context.Context) string {
//	return f.Description(ctx)
//}
//
//func (f FakeEnumValidator) Validate(ctx context.Context, request tfsdk.ValidateAttributeRequest, response *tfsdk.ValidateAttributeResponse) {
//	var str string
//	response.Diagnostics.Append(request.Config.GetAttribute(ctx, request.AttributePath, &str)...)
//	if response.Diagnostics.HasError() {
//		return
//	}
//	if !request.AttributeConfig.Equal(types.String{Value: "cheese"}) {
//		response.Diagnostics.AddAttributeError(request.AttributePath, "bad", "should be cheese")
//	}
//}
