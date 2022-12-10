package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateAuthenticationPolicyContractResource{}
	_ resource.ResourceWithSchema      = &pingfederateAuthenticationPolicyContractResource{}
	_ resource.ResourceWithConfigure   = &pingfederateAuthenticationPolicyContractResource{}
	_ resource.ResourceWithImportState = &pingfederateAuthenticationPolicyContractResource{}
)

type pingfederateAuthenticationPolicyContractResource struct {
	client *pfClient
}

func NewAuthenticationPolicyContractResource() resource.Resource {
	return &pingfederateAuthenticationPolicyContractResource{}
}

func (r *pingfederateAuthenticationPolicyContractResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceAuthenticationPolicyContract()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateAuthenticationPolicyContractResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateAuthenticationPolicyContractResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_policy_contract"
}

func (r *pingfederateAuthenticationPolicyContractResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.AuthenticationPolicyContracts.CreateAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.CreateAuthenticationPolicyContractInput{
		Body: *expandAuthenticationPolicyContract(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AuthenticationPolicyContract, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (r *pingfederateAuthenticationPolicyContractResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.AuthenticationPolicyContracts.GetAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.GetAuthenticationPolicyContractInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (r *pingfederateAuthenticationPolicyContractResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.AuthenticationPolicyContracts.UpdateAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.UpdateAuthenticationPolicyContractInput{
		Body: *expandAuthenticationPolicyContract(data),
		Id:   data.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (r *pingfederateAuthenticationPolicyContractResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.AuthenticationPolicyContracts.DeleteAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.DeleteAuthenticationPolicyContractInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateAuthenticationPolicyContractResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
