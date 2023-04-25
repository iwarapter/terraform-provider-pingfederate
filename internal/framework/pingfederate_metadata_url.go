package framework

import (
	"context"
	"fmt"

	"github.com/iwarapter/pingfederate-sdk-go/services/metadataUrls"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateMetadataUrlResource{}
	_ resource.ResourceWithConfigure   = &pingfederateMetadataUrlResource{}
	_ resource.ResourceWithImportState = &pingfederateMetadataUrlResource{}
)

type pingfederateMetadataUrlResource struct {
	client *pfClient
}

func NewMetadataUrlResource() resource.Resource {
	return &pingfederateMetadataUrlResource{}
}

func (r *pingfederateMetadataUrlResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceMetadataUrl()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateMetadataUrlResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateMetadataUrlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_metadata_url"
}

func (r *pingfederateMetadataUrlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MetadataUrlData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.MetadataUrls.AddMetadataUrlWithContext(ctx, &metadataUrls.AddMetadataUrlInput{
		Body: *expandMetadataUrl(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create MetadataUrl, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenMetadataUrl(body))...)
}

func (r *pingfederateMetadataUrlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MetadataUrlData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.MetadataUrls.GetMetadataUrlWithContext(ctx, &metadataUrls.GetMetadataUrlInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get MetadataUrl, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenMetadataUrl(body))...)
}

func (r *pingfederateMetadataUrlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MetadataUrlData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.MetadataUrls.UpdateMetadataUrlWithContext(ctx, &metadataUrls.UpdateMetadataUrlInput{
		Body: *expandMetadataUrl(data),
		Id:   data.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update MetadataUrl, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, *flattenMetadataUrl(body))...)
}

func (r *pingfederateMetadataUrlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MetadataUrlData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.MetadataUrls.DeleteMetadataUrlWithContext(ctx, &metadataUrls.DeleteMetadataUrlInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete MetadataUrl, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateMetadataUrlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
