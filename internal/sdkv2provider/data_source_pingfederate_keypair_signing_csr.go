package sdkv2provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
)

func dataSourcePingFederateKeyPairSigningCsr() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the CSR of a signing keypair in Ping Federate.",
		ReadContext: dataSourcePingFederateKeyPairSigningCsrRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource ID of the signing keypair to retrieve the CSR for.",
			},
			"cert_request_pem": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PEM-encoded CSR of the signing keypair.",
			},
		},
	}
}

func dataSourcePingFederateKeyPairSigningCsrRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	svc := m.(pfClient).KeyPairsSigning
	input := &keyPairsSigning.ExportCsrInput{
		Id: d.Get("id").(string),
	}
	result, _, err := svc.ExportCsrWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to read KeyPairSigningCsr: %s", err)

	}
	d.SetId(d.Get("id").(string))
	*result = strings.ReplaceAll(*result, " NEW ", " ")
	setResourceDataStringWithDiagnostic(d, "cert_request_pem", result, &diags)
	return diags
}
