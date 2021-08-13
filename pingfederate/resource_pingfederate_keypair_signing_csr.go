package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
)

func resourcePingFederateKeypairSigningCsrResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for importing Signing KeyPair CSR Responses within PingFederate.",
		CreateContext: resourcePingFederateKeypairSigningCsrResourceCreate,
		ReadContext:   resourcePingFederateKeypairSigningCsrResourceRead,
		DeleteContext: resourcePingFederateKeypairSigningCsrResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceKeypairCsrResourceSchema(),
	}
}

func resourcePingFederateKeypairSigningCsrResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.ImportCsrResponseInput{
		Body: pf.CSRResponse{
			FileData: String(d.Get("file_data").(string)),
		},
		Id: d.Get("keypair_id").(string),
	}
	result, _, err := svc.ImportCsrResponseWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create SigningCSR: %s", err)
	}

	d.SetId(*result.Id)
	return nil

}

func resourcePingFederateKeypairSigningCsrResourceRead(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateKeypairSigningCsrResourceDelete(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
