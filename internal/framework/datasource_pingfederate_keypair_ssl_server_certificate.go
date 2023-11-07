package framework

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &pingfederateKeyPairSslServerCertificateDataSource{}
	_ datasource.DataSourceWithConfigure = &pingfederateKeyPairSslServerCertificateDataSource{}
)

type pingfederateKeyPairSslServerCertificateDataSource struct {
	client *pfClient
}

func NewKeyPairSslServerCertificateDataSource() datasource.DataSource {
	return &pingfederateKeyPairSslServerCertificateDataSource{}
}

func (p *pingfederateKeyPairSslServerCertificateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_keypair_ssl_server_certificate"
}

func (p *pingfederateKeyPairSslServerCertificateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*pfClient)
}

func (p *pingfederateKeyPairSslServerCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"key_pair_id": schema.StringAttribute{
				Description: "ID of the key pair.",
				Required:    true,
			},
			"certificate": schema.StringAttribute{
				Description: "PEM-encoded CSR of the ssl server keypair.",
				Computed:    true,
			},
		},
	}
}

func (p *pingfederateKeyPairSslServerCertificateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	type KeyPairIdData struct {
		Id          types.String `tfsdk:"id"`
		KeyPairId   types.String `tfsdk:"key_pair_id"`
		Certificate types.String `tfsdk:"certificate"`
	}

	var data KeyPairIdData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	cert, _, err := p.client.KeyPairsSslServer.ExportCertificateFileWithContext(ctx, &keyPairsSslServer.ExportCertificateFileInput{Id: data.KeyPairId.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Unable to read keypair ssl server settings certificate", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), data.KeyPairId.ValueString())...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("key_pair_id"), data.KeyPairId.ValueString())...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("certificate"), *cert)...)
}
