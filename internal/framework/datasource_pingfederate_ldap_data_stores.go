package framework

import (
	"context"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &pingfederateLdapDataStoresDataSource{}
	_ datasource.DataSourceWithConfigure = &pingfederateLdapDataStoresDataSource{}
)

type pingfederateLdapDataStoresDataSource struct {
	client *pfClient
}

func NewLdapDataStoresDataSource() datasource.DataSource {
	return &pingfederateLdapDataStoresDataSource{}
}

func (p *pingfederateLdapDataStoresDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ldap_data_stores"
}

func (p *pingfederateLdapDataStoresDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	p.client = req.ProviderData.(*pfClient)
}

func (p *pingfederateLdapDataStoresDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasourceLdapDataStores()
}

func (p *pingfederateLdapDataStoresDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data LdapDataStoresData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	ds, _, err := p.client.DataStores.GetDataStoresWithContext(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to read LDAP data stores", err.Error())
		return
	}
	stores := make([]*pf.LdapDataStore, 0)
	for _, item := range *ds.Items {
		switch v := item.(type) {
		case *pf.LdapDataStore:
			stores = append(stores, v)
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, flattenLdapDataStores(stores))...)
}
