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
	_ resource.Resource                = &pingfederateSessionSettingsResource{}
	_ resource.ResourceWithSchema      = &pingfederateSessionSettingsResource{}
	_ resource.ResourceWithConfigure   = &pingfederateSessionSettingsResource{}
	_ resource.ResourceWithImportState = &pingfederateSessionSettingsResource{}
)

type pingfederateSessionSettingsResource struct {
	client *pfClient
}

func NewSessionSettingsResource() resource.Resource {
	return &pingfederateSessionSettingsResource{}
}

func (r *pingfederateSessionSettingsResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceSessionSettings()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateSessionSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateSessionSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_session_settings"
}

func (r *pingfederateSessionSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SessionSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.Session.UpdateSessionSettingsWithContext(ctx, &session.UpdateSessionSettingsInput{
		Body: *expandSessionSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create SessionSettings, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addSessionSettingsId(body))...)
}

func (r *pingfederateSessionSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SessionSettingsData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.GetSessionSettingsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get SessionSettings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addSessionSettingsId(body))...)
}

func (r *pingfederateSessionSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SessionSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.Session.UpdateSessionSettingsWithContext(ctx, &session.UpdateSessionSettingsInput{
		Body: *expandSessionSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update SessionSettings, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, addSessionSettingsId(body))...)
}

func (r *pingfederateSessionSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *pingfederateSessionSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func addSessionSettingsId(body *pf.SessionSettings) SessionSettingsData {
	save := *flattenSessionSettings(body)
	save.Id = types.StringValue("settings")
	return save
}
