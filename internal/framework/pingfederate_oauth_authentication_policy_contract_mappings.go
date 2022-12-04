package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	_ resource.ResourceWithConfigure   = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	_ resource.ResourceWithImportState = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
)

type pingfederateOauthAuthenticationPolicyContractMappingResource struct {
	client *pfClient
}

func NewOauthAuthenticationPolicyContractMappingResource() resource.Resource {
	return &pingfederateOauthAuthenticationPolicyContractMappingResource{}
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceApcToPersistentGrantMapping(), nil
}

// Configure adds the client configured client to the resource.
func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_authentication_policy_contract_mapping"
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.CreateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.CreateApcMappingInput{
		BypassExternalValidation: Bool(r.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create oauthAccessTokenMapping, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.GetApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.GetApcMappingInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.UpdateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.UpdateApcMappingInput{
		BypassExternalValidation: Bool(r.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
		Id:                       data.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenApcToPersistentGrantMapping(body))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.OauthAuthenticationPolicyContractMappings.DeleteApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
