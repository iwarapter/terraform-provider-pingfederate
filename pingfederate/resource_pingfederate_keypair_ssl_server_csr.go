package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateKeypairSslServerCsrResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for importing Ssl Server KeyPair CSR Responses within PingFederate.",
		CreateContext: resourcePingFederateKeypairSslServerCsrResourceCreate,
		ReadContext:   resourcePingFederateKeypairSslServerCsrResourceRead,
		DeleteContext: resourcePingFederateKeypairSslServerCsrResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceKeypairCsrResourceSchema(),
	}
}

func resourcePingFederateKeypairSslServerCsrResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.ImportCsrResponseInput{
		Body: pf.CSRResponse{
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Get("keypair_id").(string),
	}
	result, _, err := svc.ImportCsrResponseWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create SslServerCsr: %s", err)
	}

	d.SetId(*result.Id)
	return nil

}

func resourcePingFederateKeypairSslServerCsrResourceRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateKeypairSslServerCsrResourceDelete(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
