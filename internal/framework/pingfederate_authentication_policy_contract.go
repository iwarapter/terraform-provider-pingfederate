package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
		tfsdk.UseStateForUnknown(),
	}
	sch.Attributes["id"] = attribute
	return sch, nil
}

func (p pingfederateAuthenticationPolicyContractType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return pingfederateAuthenticationPolicyContractResource{
		provider: provider,
	}, diags
}

type pingfederateAuthenticationPolicyContractResource struct {
	provider provider
}

func (p pingfederateAuthenticationPolicyContractResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
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

func (p pingfederateAuthenticationPolicyContractResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
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

func (p pingfederateAuthenticationPolicyContractResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
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

func (p pingfederateAuthenticationPolicyContractResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
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

func (p pingfederateAuthenticationPolicyContractResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
