package framework

import (
	"context"
	"fmt"

	"github.com/iwarapter/pingfederate-sdk-go/services/session"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithSchema      = &pingfederateAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithConfigure   = &pingfederateAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithImportState = &pingfederateAuthenticationSessionPolicyResource{}
)

type pingfederateAuthenticationSessionPolicyResource struct {
	client *pfClient
}

func NewAuthenticationSessionPolicyResource() resource.Resource {
	return &pingfederateAuthenticationSessionPolicyResource{}
}

func (r *pingfederateAuthenticationSessionPolicyResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceAuthenticationSessionPolicy()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateAuthenticationSessionPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateAuthenticationSessionPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authentication_session_policy"
}

func (r *pingfederateAuthenticationSessionPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.Session.CreateSourcePolicyWithContext(ctx, &session.CreateSourcePolicyInput{
		Body: *expandAuthenticationSessionPolicy(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AuthenticationSessionPolicy, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationSessionPolicy(body))...)
}

func (r *pingfederateAuthenticationSessionPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.GetSourcePolicyWithContext(ctx, &session.GetSourcePolicyInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get AuthenticationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationSessionPolicy(body))...)
}

func (r *pingfederateAuthenticationSessionPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.UpdateSourcePolicyWithContext(ctx, &session.UpdateSourcePolicyInput{
		Body: *expandAuthenticationSessionPolicy(data),
		Id:   data.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AuthenticationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenAuthenticationSessionPolicy(body))...)
}

func (r *pingfederateAuthenticationSessionPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.Session.DeleteSourcePolicyWithContext(ctx, &session.DeleteSourcePolicyInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AuthenticationSessionPolicy, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateAuthenticationSessionPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
