package framework

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &pfprovider{}
)

// client satisfies the provider.Provider interface and usually is included
// with all Resource and DataSource implementations.
type pfprovider struct {
	// client can contain the upstream client SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	// TODO: If appropriate, implement upstream client SDK or HTTP client.
	client *pfClient

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the client was previously configured.
	configured bool

	// version is set to the client version on release, "dev" when the
	// client is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username                 types.String `tfsdk:"username"`
	Password                 types.String `tfsdk:"password"`
	Context                  types.String `tfsdk:"context"`
	BaseUrl                  types.String `tfsdk:"base_url"`
	BypassExternalValidation types.Bool   `tfsdk:"bypass_external_validation"`
}

// Metadata returns the client type name.
func (p *pfprovider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pingfederate"
}

func (p *pfprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "configuring client")
	var data providerData
	username := "Administrator"
	if s, ok := os.LookupEnv("PINGFEDERATE_USERNAME"); ok {
		username = s
	}
	password := "2Federate"
	if s, ok := os.LookupEnv("PINGFEDERATE_PASSWORD"); ok {
		password = s
	}
	contextPath := "/pf-admin-api/v1"
	if s, ok := os.LookupEnv("PINGFEDERATE_CONTEXT"); ok {
		contextPath = s
	}
	baseUrl := "https://localhost:9999"
	if s, ok := os.LookupEnv("PINGFEDERATE_BASEURL"); ok {
		baseUrl = s
	}
	bypassExternalValidation := false

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Username.ValueString() != "" {
		username = data.Username.ValueString()
	}
	if data.Password.ValueString() != "" {
		password = data.Password.ValueString()
	}
	if data.Context.ValueString() != "" {
		contextPath = data.Context.ValueString()
	}
	if data.BaseUrl.ValueString() != "" {
		baseUrl = data.BaseUrl.ValueString()
	}
	if !data.BypassExternalValidation.IsUnknown() {
		bypassExternalValidation = data.BypassExternalValidation.ValueBool()
	}

	config := &pfConfig{
		Username:                 username,
		Password:                 password,
		BaseURL:                  baseUrl,
		Context:                  contextPath,
		BypassExternalValidation: bypassExternalValidation,
	}

	cli, errs := config.Client()
	if errs.HasError() {
		resp.Diagnostics.Append(errs...)
		return
	}
	p.client = cli
	p.configured = true

	resp.ResourceData = cli
	resp.DataSourceData = cli

	tflog.Info(ctx, "Configured PingFederate client", map[string]any{"success": true})
}

func (p *pfprovider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewApplicationSessionPolicyResource,
		NewAuthenticationPolicyContractResource,
		NewAuthenticationSessionPolicyResource,
		NewGlobalAuthenticationSessionPolicyResource,
		NewMetadataUrlResource,
		NewOauthAuthenticationPolicyContractMappingResource,
		NewOAuthClientResource,
		NewRedirectValidationResource,
		NewSessionSettingsResource,
	}
}

func (p *pfprovider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewOauthAuthServerSettingsDataSource,
	}
}

func (p *pfprovider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "The username for pingfederate API.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Description: "The password for pingfederate API.",
			},
			"base_url": schema.StringAttribute{
				Optional:    true,
				Description: "The base url of the pingfederate API.",
			},
			"context": schema.StringAttribute{
				Optional:    true,
				Description: "The context path of the pingfederate API.",
			},
			"bypass_external_validation": schema.BoolAttribute{
				Optional:    true,
				Description: "External validation will be bypassed when set to true. Default to false.",
			},
		},
	}
}

func New(version string) provider.Provider {
	return &pfprovider{
		version: version,
	}
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Float is a helper routine that allocates a new float value
// to store v and returns a pointer to it.
func Float(v float32) *float32 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
