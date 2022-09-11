package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
)

type pingfederateOAuthClientType struct{}

func (p pingfederateOAuthClientType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceClient(), nil
}

func (p pingfederateOAuthClientType) NewResource(ctx context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	pf, diags := convertProviderType(in)

	return pingfederateOAuthClientResource{
		provider: pf,
	}, diags
}

type pingfederateOAuthClientResource struct {
	provider pfprovider
}

func (r pingfederateOAuthClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.provider.client.OauthClients.CreateClientWithContext(ctx, &oauthClients.CreateClientInput{
		Body: *expandClient(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OAuthClient, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenClient(body))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.Null {
		var originalSecret string
		//resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		//resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r pingfederateOAuthClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.provider.client.OauthClients.GetClientWithContext(ctx, &oauthClients.GetClientInput{Id: data.ClientId.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenClient(body))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.Null {
		var originalSecret string
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r pingfederateOAuthClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.provider.client.OauthClients.UpdateClientWithContext(ctx, &oauthClients.UpdateClientInput{
		Body: *expandClient(data),
		Id:   data.ClientId.Value,
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenClient(body))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.Null {
		var originalSecret string
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r pingfederateOAuthClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.provider.client.OauthClients.DeleteClientWithContext(ctx, &oauthClients.DeleteClientInput{Id: data.ClientId.Value})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete OAuthClient, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r pingfederateOAuthClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("client_id"), req, resp)
}

//
//type pingfederateOAuthClientTypeVersion0Data struct {
//	CoreAttributes     []types.String `tfsdk:"core_attributes"`
//	ExtendedAttributes []types.String `tfsdk:"extended_attributes"`
//	Id                 types.String   `tfsdk:"policy_contract_id"`
//	Name               types.String   `tfsdk:"name"`
//}
//
//func (p pingfederateOAuthClientType) UpgradeState(ctx context.Context) map[int64]tfsdk.ResourceStateUpgrader {
//	return map[int64]tfsdk.ResourceStateUpgrader{
//		0: {
//			PriorSchema: &tfsdk.Schema{
//				Attributes: map[string]tfsdk.Attribute{
//					"policy_contract_id": {
//						Type: types.StringType,
//					},
//					"name": {
//						Type: types.StringType,
//					},
//					"core_attributes": {
//						Type: types.SetType{
//							ElemType: types.StringType,
//						},
//					},
//					"extended_attributes": {
//						Type: types.SetType{
//							ElemType: types.StringType,
//						},
//					},
//				},
//			},
//			StateUpgrader: func(ctx context.Context, request tfsdk.UpgradeResourceStateRequest, response *tfsdk.UpgradeResourceStateResponse) {
//				request.State
//			},
//		},
//	}
//}
