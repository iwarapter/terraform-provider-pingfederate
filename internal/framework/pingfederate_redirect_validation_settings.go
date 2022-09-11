package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/iwarapter/pingfederate-sdk-go/services/redirectValidation"
)

type pingfederateRedirectValidationSettingsType struct{}

func (p pingfederateRedirectValidationSettingsType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceRedirectValidationSettings(), nil
}

func (p pingfederateRedirectValidationSettingsType) NewResource(_ context.Context, in provider.Provider) (resource.Resource, diag.Diagnostics) {
	pf, diags := convertProviderType(in)

	return pingfederateRedirectValidationResource{
		provider: pf,
	}, diags
}

type pingfederateRedirectValidationResource struct {
	provider pfprovider
}

func (r pingfederateRedirectValidationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.provider.client.RedirectValidation.UpdateRedirectValidationSettingsWithContext(ctx, &redirectValidation.UpdateRedirectValidationSettingsInput{
		Body: *expandRedirectValidationSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenRedirectValidationSettings(body))...)
}

func (r pingfederateRedirectValidationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.provider.client.RedirectValidation.GetRedirectValidationSettingsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenRedirectValidationSettings(body))...)
}

func (r pingfederateRedirectValidationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data RedirectValidationSettingsData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.provider.client.RedirectValidation.UpdateRedirectValidationSettingsWithContext(ctx, &redirectValidation.UpdateRedirectValidationSettingsInput{
		Body: *expandRedirectValidationSettings(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update RedirectValidation, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenRedirectValidationSettings(body))...)
}

func (r pingfederateRedirectValidationResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

func (r pingfederateRedirectValidationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
