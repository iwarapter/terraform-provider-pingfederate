package pingfederate

import (
	"context"
	"fmt"

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
			StateContext: resourcePingFederateKeypairSslServerResourceImport,
		},
		Schema: resourceKeypairResourceSchema(),
	}
}

func resourcePingFederateKeypairSslServerResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		result, _, err := svc.ImportKeyPairWithContext(ctx, &input)
		if err != nil {
			return diag.Errorf("unable to create SslServerKeypair: %s", err)
		}

		d.SetId(*result.Id)
		return resourceKeypairResourceReadResult(d, result)
	}

	input := keyPairsSslServer.CreateKeyPairInput{
		Body: *resourcePingFederateKeypairResourceReadData(d),
	}
	result, _, err := svc.CreateKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to generate SslServerKeypair: %s", err)
	}

	d.SetId(*result.Id)
	return resourceKeypairResourceReadResult(d, result)

}

func resourcePingFederateKeypairSslServerResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read SslServerKeypair: %s", err)
	}
	return resourceKeypairResourceReadResult(d, result)
}

func resourcePingFederateKeypairSslServerResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.DeleteKeyPairInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete SslServerKeypair: %s", err)
	}
	return nil
}

func resourcePingFederateKeypairSslServerResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).KeyPairsSslServer
	input := keyPairsSslServer.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairWithContext(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve read ssl server keypair for import %s", err)
	}

	return importKeyPairView(d, result, err)
}
