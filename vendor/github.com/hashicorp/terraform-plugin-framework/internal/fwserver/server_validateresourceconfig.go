package fwserver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// ValidateResourceConfigRequest is the framework server request for the
// ValidateResourceConfig RPC.
type ValidateResourceConfigRequest struct {
	Config       *tfsdk.Config
	ResourceType provider.ResourceType
}

// ValidateResourceConfigResponse is the framework server response for the
// ValidateResourceConfig RPC.
type ValidateResourceConfigResponse struct {
	Diagnostics diag.Diagnostics
}

// ValidateResourceConfig implements the framework server ValidateResourceConfig RPC.
func (s *Server) ValidateResourceConfig(ctx context.Context, req *ValidateResourceConfigRequest, resp *ValidateResourceConfigResponse) {
	if req == nil || req.Config == nil {
		return
	}

	// Always instantiate new Resource instances.
	logging.FrameworkDebug(ctx, "Calling provider defined ResourceType NewResource")
	resourceImpl, diags := req.ResourceType.NewResource(ctx, s.Provider)
	logging.FrameworkDebug(ctx, "Called provider defined ResourceType NewResource")

	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	vdscReq := resource.ValidateConfigRequest{
		Config: *req.Config,
	}

	if resourceWithConfigValidators, ok := resourceImpl.(resource.ResourceWithConfigValidators); ok {
		logging.FrameworkTrace(ctx, "Resource implements ResourceWithConfigValidators")

		for _, configValidator := range resourceWithConfigValidators.ConfigValidators(ctx) {
			vdscResp := &resource.ValidateConfigResponse{
				Diagnostics: resp.Diagnostics,
			}

			logging.FrameworkDebug(
				ctx,
				"Calling provider defined ResourceConfigValidator",
				map[string]interface{}{
					logging.KeyDescription: configValidator.Description(ctx),
				},
			)
			configValidator.ValidateResource(ctx, vdscReq, vdscResp)
			logging.FrameworkDebug(
				ctx,
				"Called provider defined ResourceConfigValidator",
				map[string]interface{}{
					logging.KeyDescription: configValidator.Description(ctx),
				},
			)

			resp.Diagnostics = vdscResp.Diagnostics
		}
	}

	if resourceWithValidateConfig, ok := resourceImpl.(resource.ResourceWithValidateConfig); ok {
		logging.FrameworkTrace(ctx, "Resource implements ResourceWithValidateConfig")

		vdscResp := &resource.ValidateConfigResponse{
			Diagnostics: resp.Diagnostics,
		}

		logging.FrameworkDebug(ctx, "Calling provider defined Resource ValidateConfig")
		resourceWithValidateConfig.ValidateConfig(ctx, vdscReq, vdscResp)
		logging.FrameworkDebug(ctx, "Called provider defined Resource ValidateConfig")

		resp.Diagnostics = vdscResp.Diagnostics
	}

	validateSchemaReq := ValidateSchemaRequest{
		Config: *req.Config,
	}
	validateSchemaResp := ValidateSchemaResponse{
		Diagnostics: resp.Diagnostics,
	}

	SchemaValidate(ctx, req.Config.Schema, validateSchemaReq, &validateSchemaResp)

	resp.Diagnostics = validateSchemaResp.Diagnostics
}