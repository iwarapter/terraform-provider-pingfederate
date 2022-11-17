package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateOAuthClientResource{}
	_ resource.ResourceWithConfigure   = &pingfederateOAuthClientResource{}
	_ resource.ResourceWithImportState = &pingfederateOAuthClientResource{}
)

type pingfederateOAuthClientResource struct {
	client *pfClient
}

func NewOAuthClientResource() resource.Resource {
	return &pingfederateOAuthClientResource{}
}

func (r *pingfederateOAuthClientResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceClient(), nil
}

// Configure adds the client configured client to the resource.
func (r *pingfederateOAuthClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateOAuthClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_client"
}
func (r *pingfederateOAuthClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.versionRequestModifier(&data)
	body, _, err := r.client.OauthClients.CreateClientWithContext(ctx, &oauthClients.CreateClientInput{
		Body: *expandClient(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OAuthClient, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		//resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		//resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthClients.GetClientWithContext(ctx, &oauthClients.GetClientInput{Id: data.ClientId.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.versionRequestModifier(&data)
	body, _, err := r.client.OauthClients.UpdateClientWithContext(ctx, &oauthClients.UpdateClientInput{
		Body: *expandClient(data),
		Id:   data.ClientId.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.OauthClients.DeleteClientWithContext(ctx, &oauthClients.DeleteClientInput{Id: data.ClientId.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete OAuthClient, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateOAuthClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("client_id"), req, resp)
}

// old version of pingfederate dont handle the follow fields so we strip them before marshalling
func (r *pingfederateOAuthClientResource) versionRequestModifier(data *ClientData) {
	if !r.client.IsVersionGreaterEqThan(10, 3) {
		data.RefreshTokenRollingIntervalType = types.StringNull()
	}
	if !r.client.IsVersionGreaterEqThan(11, 0) {
		data.PersistentGrantReuseType = types.StringNull()
		data.RefreshTokenRollingGracePeriodType = types.StringNull()
	}
	if !r.client.IsVersionGreaterEqThan(11, 1) {
		data.ClientSecretRetentionPeriodType = types.StringNull()
		data.RequireJwtSecuredAuthorizationResponseMode = types.BoolNull()
	}
}

// old version of pingfederate dont handle the follow fields so add the defaults back to keep state happy!
func (r *pingfederateOAuthClientResource) versionResponseModifier(data *ClientData) *ClientData {
	if !r.client.IsVersionGreaterEqThan(10, 3) {
		data.RefreshTokenRollingIntervalType = types.StringValue("SERVER_DEFAULT")
	}
	if !r.client.IsVersionGreaterEqThan(11, 0) {
		data.PersistentGrantReuseType = types.StringValue("SERVER_DEFAULT")
		data.RefreshTokenRollingGracePeriodType = types.StringValue("SERVER_DEFAULT")
	}
	if !r.client.IsVersionGreaterEqThan(11, 1) {
		data.ClientSecretRetentionPeriodType = types.StringValue("SERVER_DEFAULT")
		data.RequireJwtSecuredAuthorizationResponseMode = types.BoolValue(false)
	}

	return data
}
