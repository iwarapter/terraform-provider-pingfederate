package framework

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// provider satisfies the provider.Provider interface and usually is included
// with all Resource and DataSource implementations.
type pfprovider struct {
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

func (p *pfprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "configuring provider")
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

	//var data providerData
	////data.Username.Value = "Administrator"
	////data.Password.Value = "2Federate"
	////data.Context.Value = "/pf-admin-api/v1"
	////data.BaseUrl.Value = "https://localhost:9999"
	////data.BypassExternalValidation.Value = false
	//
	//diags := req.Config.Get(ctx, &data)
	//resp.Diagnostics.Append(diags...)
	//
	//if resp.Diagnostics.HasError() {
	//	return
	//}
	//
	//var username string
	//if data.Username.Null {
	//	if s, ok := os.LookupEnv("PINGFEDERATE_USERNAME"); ok {
	//		username = s
	//	} else {
	//		username = "Administrator"
	//	}
	//} else {
	//	username = data.Username.Value
	//}
	//var password string
	//if data.Password.Null {
	//	if s, ok := os.LookupEnv("PINGFEDERATE_PASSWORD"); ok {
	//		password = s
	//	} else {
	//		password = "2Federate"
	//	}
	//} else {
	//	password = data.Password.Value
	//}
	//var contextPath string
	//if data.Context.Null {
	//	if s, ok := os.LookupEnv("PINGFEDERATE_CONTEXT"); ok {
	//		contextPath = s
	//	} else {
	//		contextPath = "/pf-admin-api/v1"
	//	}
	//} else {
	//	contextPath = data.Context.Value
	//}
	//var baseUrl string
	//if data.BaseUrl.Null {
	//	if s, ok := os.LookupEnv("PINGFEDERATE_BASEURL"); ok {
	//		baseUrl = s
	//	} else {
	//		baseUrl = "https://localhost:9999"
	//	}
	//} else {
	//	baseUrl = data.BaseUrl.Value
	//}
	//var bypassExternalValidation bool
	//if data.BypassExternalValidation.Null {
	//	bypassExternalValidation = false
	//} else {
	//	bypassExternalValidation = data.BypassExternalValidation.Value
	//}
	//
	//config := &pfConfig{
	//	Username:                 username,
	//	Password:                 password,
	//	BaseURL:                  baseUrl,
	//	Context:                  contextPath,
	//	BypassExternalValidation: bypassExternalValidation,
	//}

	cli, errs := config.Client()
	if errs.HasError() {
		resp.Diagnostics.Append(errs...)
		return
	}
	p.client = cli
	p.configured = true
}

func (p *pfprovider) GetResources(_ context.Context) (map[string]provider.ResourceType, diag.Diagnostics) {
	return map[string]provider.ResourceType{
		"pingfederate_authentication_policy_contract":               pingfederateAuthenticationPolicyContractType{},
		"pingfederate_oauth_authentication_policy_contract_mapping": pingfederateOauthAuthenticationPolicyContractMappingType{},
		"pingfederate_oauth_client":                                 pingfederateOAuthClientType{},
		"pingfederate_redirect_validation_settings":                 pingfederateRedirectValidationSettingsType{},
	}, nil
}

func (p *pfprovider) GetDataSources(_ context.Context) (map[string]provider.DataSourceType, diag.Diagnostics) {
	return map[string]provider.DataSourceType{}, nil
}

func (p *pfprovider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The username for pingfederate API.",
				PlanModifiers: []tfsdk.AttributePlanModifier{
					Default(types.String{Value: "Administrator"}),
				},
			},
			"password": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The password for pingfederate API.",
				PlanModifiers: []tfsdk.AttributePlanModifier{
					Default(types.String{Value: "2FederateM0re"}),
				},
			},
			"base_url": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The base url of the pingfederate API.",
				PlanModifiers: []tfsdk.AttributePlanModifier{
					Default(types.String{Value: "https://localhost:9999"}),
				},
			},
			"context": {
				Optional:    true,
				Type:        types.StringType,
				Description: "The context path of the pingfederate API.",
				PlanModifiers: []tfsdk.AttributePlanModifier{
					Default(types.String{Value: "/pf-admin-api/v1"}),
				},
			},
			"bypass_external_validation": {
				Optional:    true,
				Type:        types.BoolType,
				Description: "External validation will be bypassed when set to true. Default to false.",
			},
		},
	}, nil
}

func New(version string) provider.Provider {
	return &pfprovider{
		version: version,
	}
}

//func New(version string) func() tfsdk.Provider {
//	return func() tfsdk.Provider {
//		return &provider{
//			version: version,
//		}
//	}
//}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in provider.Provider) (pfprovider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*pfprovider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return pfprovider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return pfprovider{}, diags
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
