package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"
)

type pingfederateAuthenticationPolicyContractType struct{}

func (p pingfederateAuthenticationPolicyContractType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	sch := resourceAuthenticationPolicyContract()
	attribute := sch.Attributes["id"]
	attribute.Optional = true
	attribute.Required = false
	attribute.Computed = true
	attribute.PlanModifiers = tfsdk.AttributePlanModifiers{
		resource.UseStateForUnknown(),
	}
	sch.Attributes["id"] = attribute
	return sch, nil
}

func (p pingfederateAuthenticationPolicyContractType) NewResource(ctx context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	pf, diags := convertProviderType(in)

	return pingfederateAuthenticationPolicyContractResource{
		provider: pf,
	}, diags
}

type pingfederateAuthenticationPolicyContractResource struct {
	provider pfprovider
}

func (p pingfederateAuthenticationPolicyContractResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := p.provider.client.AuthenticationPolicyContracts.CreateAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.CreateAuthenticationPolicyContractInput{
		Body: *expandAuthenticationPolicyContract(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AuthenticationPolicyContract, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (p pingfederateAuthenticationPolicyContractResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := p.provider.client.AuthenticationPolicyContracts.GetAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.GetAuthenticationPolicyContractInput{Id: data.Id.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (p pingfederateAuthenticationPolicyContractResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationPolicyContractData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := p.provider.client.AuthenticationPolicyContracts.UpdateAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.UpdateAuthenticationPolicyContractInput{
		Body: *expandAuthenticationPolicyContract(data),
		Id:   data.Id.Value,
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationPolicyContract(body))...)
}

func (p pingfederateAuthenticationPolicyContractResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := p.provider.client.AuthenticationPolicyContracts.DeleteAuthenticationPolicyContractWithContext(ctx, &authenticationPolicyContracts.DeleteAuthenticationPolicyContractInput{Id: data.Id.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AuthenticationPolicyContract, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (p pingfederateAuthenticationPolicyContractResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
