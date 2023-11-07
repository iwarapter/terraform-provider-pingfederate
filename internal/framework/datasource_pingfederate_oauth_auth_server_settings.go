package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &pingfederateOauthAuthServerSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &pingfederateOauthAuthServerSettingsDataSource{}
)

type pingfederateOauthAuthServerSettingsDataSource struct {
	client *pfClient
}

func NewOauthAuthServerSettingsDataSource() datasource.DataSource {
	return &pingfederateOauthAuthServerSettingsDataSource{}
}

func (p *pingfederateOauthAuthServerSettingsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_auth_server_settings"
}

func (p *pingfederateOauthAuthServerSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*pfClient)
}

func (p *pingfederateOauthAuthServerSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasourceAuthorizationServerSettings()
}

func (p *pingfederateOauthAuthServerSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	settings, _, err := p.client.OauthAuthServerSettings.GetAuthorizationServerSettingsWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read oauth auth server settings", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, flattenAuthorizationServerSettings(settings))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "oauth_auth_server_settings")...)
}
