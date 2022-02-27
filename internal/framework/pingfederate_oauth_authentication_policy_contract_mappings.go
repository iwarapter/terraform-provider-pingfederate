package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
)

type pingfederateOauthAuthenticationPolicyContractMappingType struct{}

func (p pingfederateOauthAuthenticationPolicyContractMappingType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	sch := resourceApcToPersistentGrantMapping()
	attribute := sch.Attributes["id"]
	attribute.Optional = false
	attribute.Required = false
	attribute.Computed = true
	attribute.PlanModifiers = tfsdk.AttributePlanModifiers{
		tfsdk.UseStateForUnknown(),
	}
	sch.Attributes["id"] = attribute
	return sch, nil
}

func (p pingfederateOauthAuthenticationPolicyContractMappingType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return pingfederateOauthAuthenticationPolicyContractMappingResource{
		provider: provider,
	}, diags
}

type pingfederateOauthAuthenticationPolicyContractMappingResource struct {
	provider provider
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := p.provider.client.OauthAuthenticationPolicyContractMappings.CreateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.CreateApcMappingInput{
		BypassExternalValidation: Bool(p.provider.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create oauthAccessTokenMapping, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := p.provider.client.OauthAuthenticationPolicyContractMappings.GetApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.GetApcMappingInput{Id: data.Id.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := p.provider.client.OauthAuthenticationPolicyContractMappings.UpdateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.UpdateApcMappingInput{
		BypassExternalValidation: Bool(p.provider.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
		Id:                       data.Id.Value,
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := p.provider.client.OauthAuthenticationPolicyContractMappings.DeleteApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{Id: data.Id.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
