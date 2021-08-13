package pingfederate

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
)

func dataSourcePingFederateKeyPairSslServerCsr() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the CSR of a ssl server keypair in Ping Federate.",
		ReadContext: dataSourcePingFederateKeyPairSslServerCsrRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource ID of the ssl server keypair to retrieve the CSR for.",
			},
			"cert_request_pem": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PEM-encoded CSR of the ssl server keypair.",
			},
		},
	}
}

func dataSourcePingFederateKeyPairSslServerCsrRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	svc := m.(pfClient).KeyPairsSslServer
	input := &keyPairsSslServer.ExportCsrInput{
		Id: d.Get("id").(string),
	}
	result, _, err := svc.ExportCsrWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to read KeyPairSslServerCsr: %s", err)

	}
	d.SetId(d.Get("id").(string))
	*result = strings.ReplaceAll(*result, " NEW ", " ")
	setResourceDataStringWithDiagnostic(d, "cert_request_pem", result, &diags)
	return diags
}
