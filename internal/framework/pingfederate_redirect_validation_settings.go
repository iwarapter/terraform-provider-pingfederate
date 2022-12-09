package framework

import (
	"context"
	"fmt"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/iwarapter/pingfederate-sdk-go/services/redirectValidation"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateRedirectValidationSettingsResource{}
	_ resource.ResourceWithSchema      = &pingfederateRedirectValidationSettingsResource{}
	_ resource.ResourceWithConfigure   = &pingfederateRedirectValidationSettingsResource{}
	_ resource.ResourceWithImportState = &pingfederateRedirectValidationSettingsResource{}
)

type pingfederateRedirectValidationSettingsResource struct {
	client *pfClient
}

func NewRedirectValidationResource() resource.Resource {
	return &pingfederateRedirectValidationSettingsResource{}
}

func (r *pingfederateRedirectValidationSettingsResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceRedirectValidationSettings()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateRedirectValidationSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateRedirectValidationSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_redirect_validation_settings"
}
func (r *pingfederateRedirectValidationSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.RedirectValidation.UpdateRedirectValidationSettingsWithContext(ctx, &redirectValidation.UpdateRedirectValidationSettingsInput{
		Body: *expandRedirectValidationSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addRedirectValidationSettingsId(body))...)
}

func (r *pingfederateRedirectValidationSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.RedirectValidation.GetRedirectValidationSettingsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addRedirectValidationSettingsId(body))...)
}

func (r *pingfederateRedirectValidationSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.RedirectValidation.UpdateRedirectValidationSettingsWithContext(ctx, &redirectValidation.UpdateRedirectValidationSettingsInput{
		Body: *expandRedirectValidationSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, addRedirectValidationSettingsId(body))...)
}

func (r *pingfederateRedirectValidationSettingsResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r *pingfederateRedirectValidationSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func addRedirectValidationSettingsId(body *pf.RedirectValidationSettings) RedirectValidationSettingsData {
	save := *flattenRedirectValidationSettings(body)
	save.Id = types.StringValue("settings")
	return save
}
