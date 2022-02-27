package framework

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	// client can contain the upstream provider SDK or HTTP client used to
	// communicate with the upstream service. Resource and DataSource
	// implementations can then make calls using this client.
	//
	// TODO: If appropriate, implement upstream provider SDK or HTTP client.
	client *pfClient

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
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

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	tflog.Debug(ctx, "configuring provider")
	var data providerData
	data.Username.Value = "Administrator"
	data.Password.Value = "2Federate"
	data.Context.Value = "/pf-admin-api/v1"
	data.BaseUrl.Value = "https://localhost:9999"
	data.BypassExternalValidation.Value = false

	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Username.Null {
		if s, ok := os.LookupEnv("PINGFEDERATE_USERNAME"); ok {
			data.Username.Value = s
		} else {
			data.Username.Value = "Administrator"
		}
	}
	if data.Password.Null {
		if s, ok := os.LookupEnv("PINGFEDERATE_PASSWORD"); ok {
			data.Password.Value = s
		} else {
			data.Password.Value = "2Federate"
		}
	}
	if data.Context.Null {
		if s, ok := os.LookupEnv("PINGFEDERATE_CONTEXT"); ok {
			data.Context.Value = s
		} else {
			data.Context.Value = "/pf-admin-api/v1"
		}
	}
	if data.BaseUrl.Null {
		if s, ok := os.LookupEnv("PINGFEDERATE_BASEURL"); ok {
			data.BaseUrl.Value = s
		} else {
			data.BaseUrl.Value = "https://localhost:9999"
		}
	}

	config := &pfConfig{
		Username:                 data.Username.Value,
		Password:                 data.Password.Value,
		BaseURL:                  data.BaseUrl.Value,
		Context:                  data.Context.Value,
		BypassExternalValidation: data.BypassExternalValidation.Value,
	}

	cli, errs := config.Client()
	if errs.HasError() {
		resp.Diagnostics.Append(errs...)
		return
	}
	p.client = cli
	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"pingfederate_authentication_policy_contract":               pingfederateAuthenticationPolicyContractType{},
		"pingfederate_oauth_authentication_policy_contract_mapping": pingfederateOauthAuthenticationPolicyContractMappingType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		//"scaffolding_example": exampleDataSourceType{},
	}, nil
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The username for pingfederate API.",
			},
			"password": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The password for pingfederate API.",
			},
			"base_url": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The base url of the pingfederate API.",
			},
			"context": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The context path of the pingfederate API.",
			},
			"bypass_external_validation": {
				Optional:    true,
				Type:        types.BoolType,
				Description: "External validation will be bypassed when set to true. Default to false.",
			},
		},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
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
