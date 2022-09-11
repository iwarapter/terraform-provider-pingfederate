package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
)

type pingfederateOauthAuthenticationPolicyContractMappingType struct{}

func (p pingfederateOauthAuthenticationPolicyContractMappingType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceApcToPersistentGrantMapping(), nil
}

func (p pingfederateOauthAuthenticationPolicyContractMappingType) NewResource(_ context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	pf, diags := convertProviderType(in)

	return pingfederateOauthAuthenticationPolicyContractMappingResource{
		provider: pf,
	}, diags
}

type pingfederateOauthAuthenticationPolicyContractMappingResource struct {
	provider pfprovider
}

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
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

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
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

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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

func (p pingfederateOauthAuthenticationPolicyContractMappingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
