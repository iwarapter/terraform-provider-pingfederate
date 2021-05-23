package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func dataSourcePingFederateLdapDataStore() *schema.Resource {
	var dsSchema = resourcePingFederateLdapDataStoreResourceSchema()
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
		ReadContext: dataSourcePingFederateLdapDataStoreRead,
		Schema:      dsSchema,
	}
}

func dataSourcePingFederateLdapDataStoreRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	result, _, err := svc.GetDataStoresWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read LdapDataStore: %s", err)

	}
	for _, item := range *result.Items {
		switch v := item.(type) {
		case *pf.LdapDataStore:
			if d.Get("name").(string) == *v.Name {
				d.SetId(*v.Id)
				return resourcePingFederateLdapDataStoreResourceReadResult(d, v)
			}
		}
	}
	return diag.Errorf("unable to find ldap data store with name '%s' found '%d' results", d.Get("name").(string), len(*result.Items))
}
