package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingFederateVersion() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the version of the PingFederate instance.",
		ReadContext: dataSourcePingFederateVersionRead,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server version.",
			},
		},
	}
}

func dataSourcePingFederateVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).Version
	result, _, err := svc.GetVersionWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read Version: %s", err)
	}
	var diags diag.Diagnostics
	d.SetId("version")
	setResourceDataStringWithDiagnostic(d, "version", result.Version, &diags)
	return diags
}
