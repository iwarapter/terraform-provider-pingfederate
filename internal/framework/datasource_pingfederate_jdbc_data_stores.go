package framework

import (
	"context"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &pingfederateJdbcDataStoresDataSource{}
	_ datasource.DataSourceWithConfigure = &pingfederateJdbcDataStoresDataSource{}
)

type pingfederateJdbcDataStoresDataSource struct {
	client *pfClient
}

func NewJdbcDataStoresDataSource() datasource.DataSource {
	return &pingfederateJdbcDataStoresDataSource{}
}

func (p *pingfederateJdbcDataStoresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_jdbc_data_stores"
}

func (p *pingfederateJdbcDataStoresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*pfClient)
}

func (p *pingfederateJdbcDataStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasourceJdbcDataStores()
}

func (p *pingfederateJdbcDataStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data JdbcDataStoresData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	ds, _, err := p.client.DataStores.GetDataStoresWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read keypair ssl server settings certificate", err.Error())
		return
	}
	stores := make([]*pf.JdbcDataStore, 0)
	for _, item := range *ds.Items {
		switch v := item.(type) {
		case *pf.JdbcDataStore:
			stores = append(stores, v)
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, flattenJdbcDataStores(stores))...)
}
