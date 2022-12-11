package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/session"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateGlobalAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithSchema      = &pingfederateGlobalAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithConfigure   = &pingfederateGlobalAuthenticationSessionPolicyResource{}
	_ resource.ResourceWithImportState = &pingfederateGlobalAuthenticationSessionPolicyResource{}
)

type pingfederateGlobalAuthenticationSessionPolicyResource struct {
	client *pfClient
}

func NewGlobalAuthenticationSessionPolicyResource() resource.Resource {
	return &pingfederateGlobalAuthenticationSessionPolicyResource{}
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceGlobalAuthenticationSessionPolicy()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_global_authentication_session_policy"
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GlobalAuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.Session.UpdateGlobalPolicyWithContext(ctx, &session.UpdateGlobalPolicyInput{
		Body: *expandGlobalAuthenticationSessionPolicy(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create GlobalAuthenticationSessionPolicy, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addGlobalAuthenticationSessionPolicyId(body))...)
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GlobalAuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.GetGlobalPolicyWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get GlobalAuthenticationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addGlobalAuthenticationSessionPolicyId(body))...)
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GlobalAuthenticationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.UpdateGlobalPolicyWithContext(ctx, &session.UpdateGlobalPolicyInput{
		Body: *expandGlobalAuthenticationSessionPolicy(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update GlobalAuthenticationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addGlobalAuthenticationSessionPolicyId(body))...)
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *pingfederateGlobalAuthenticationSessionPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func addGlobalAuthenticationSessionPolicyId(body *pf.GlobalAuthenticationSessionPolicy) GlobalAuthenticationSessionPolicyData {
	save := *flattenGlobalAuthenticationSessionPolicy(body)
	save.Id = types.StringValue("settings")
	return save
}
