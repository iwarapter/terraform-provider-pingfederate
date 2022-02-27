package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func dataSourcePingFederateJdbcDataStore() *schema.Resource {
	var dsSchema = resourcePingFederateJdbcDataStoreResourceSchema()
	for s := range dsSchema {
		dsSchema[s].Computed = true
		dsSchema[s].Required = false
		dsSchema[s].Optional = false
		dsSchema[s].Default = nil
		dsSchema[s].ValidateDiagFunc = nil
	}
	dsSchema["name"].Required = true
	dsSchema["name"].Computed = false
	return &schema.Resource{
		Description: "Use this data source to get a jdbc data store in Ping Federate by its name.",
		ReadContext: dataSourcePingFederateJdbcDataStoreRead,
		Schema:      dsSchema,
	}
}

func dataSourcePingFederateJdbcDataStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	result, _, err := svc.GetDataStoresWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read JdbcDataStore: %s", err)

	}
	for _, item := range *result.Items {
		switch v := item.(type) {
		case *pf.JdbcDataStore:
			if d.Get("name").(string) == *v.Name {
				d.SetId(*v.Id)
				return resourcePingFederateJdbcDataStoreResourceReadResult(d, v)
			}
		}
	}
	return diag.Errorf("unable to find jdbc data store with name '%s' found '%d' results", d.Get("name").(string), len(*result.Items))
}
