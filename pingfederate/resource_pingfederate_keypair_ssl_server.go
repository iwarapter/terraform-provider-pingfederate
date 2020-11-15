package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
)

func resourcePingFederateKeypairSslServerResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateKeypairSslServerResourceCreate,
		ReadContext:   resourcePingFederateKeypairSslServerResourceRead,
		DeleteContext: resourcePingFederateKeypairSslServerResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceKeypairResourceSchema(),
	}
}

func resourcePingFederateKeypairSslServerResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	if _, ok := d.GetOk("file_data"); ok {
		input := keyPairsSslServer.ImportKeyPairInput{
			Body: pf.PKCS12File{
				FileData: String(d.Get("file_data").(string)),
				Password: String(d.Get("password").(string)),
			},
		}
		if val, ok := d.GetOk("keypair_id"); ok {
			input.Body.Id = String(val.(string))
		}
		result, _, err := svc.ImportKeyPair(&input)
		if err != nil {
			return diag.Errorf("unable to create SslServerKeypair: %s", err)
		}

		d.SetId(*result.Id)
		return resourceKeypairResourceReadResult(d, result)
	}

	input := keyPairsSslServer.CreateKeyPairInput{
		Body: *resourcePingFederateKeypairResourceReadData(d),
	}
	result, _, err := svc.CreateKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to generate SslServerKeypair: %s", err)
	}

	d.SetId(*result.Id)
	return resourceKeypairResourceReadResult(d, result)

}

func resourcePingFederateKeypairSslServerResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to read SslServerKeypair: %s", err)
	}
	return resourceKeypairResourceReadResult(d, result)
}

func resourcePingFederateKeypairSslServerResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.DeleteKeyPairInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to delete SslServerKeypair: %s", err)
	}
	return nil
}