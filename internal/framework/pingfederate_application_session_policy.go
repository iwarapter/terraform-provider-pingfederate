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
	_ resource.Resource                = &pingfederateApplicationSessionPolicyResource{}
	_ resource.ResourceWithConfigure   = &pingfederateApplicationSessionPolicyResource{}
	_ resource.ResourceWithImportState = &pingfederateApplicationSessionPolicyResource{}
)

type pingfederateApplicationSessionPolicyResource struct {
	client *pfClient
}

func NewApplicationSessionPolicyResource() resource.Resource {
	return &pingfederateApplicationSessionPolicyResource{}
}

func (r *pingfederateApplicationSessionPolicyResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceApplicationSessionPolicy()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateApplicationSessionPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateApplicationSessionPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_application_session_policy"
}

func (r *pingfederateApplicationSessionPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApplicationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.Session.UpdateApplicationPolicyWithContext(ctx, &session.UpdateApplicationPolicyInput{
		Body: *expandApplicationSessionPolicy(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ApplicationSessionPolicy, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addApplicationSessionPolicyId(body))...)
}

func (r *pingfederateApplicationSessionPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApplicationSessionPolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.GetApplicationPolicyWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get ApplicationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addApplicationSessionPolicyId(body))...)
}

func (r *pingfederateApplicationSessionPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApplicationSessionPolicyData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.UpdateApplicationPolicyWithContext(ctx, &session.UpdateApplicationPolicyInput{
		Body: *expandApplicationSessionPolicy(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ApplicationSessionPolicy, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addApplicationSessionPolicyId(body))...)
}

func (r *pingfederateApplicationSessionPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *pingfederateApplicationSessionPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func addApplicationSessionPolicyId(body *pf.ApplicationSessionPolicy) ApplicationSessionPolicyData {
	save := *flattenApplicationSessionPolicy(body)
	save.Id = types.StringValue("settings")
	return save
}
