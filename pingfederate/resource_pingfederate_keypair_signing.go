package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
)

func resourcePingFederateKeypairSigningResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateKeypairSigningResourceCreate,
		ReadContext:   resourcePingFederateKeypairSigningResourceRead,
		DeleteContext: resourcePingFederateKeypairSigningResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourceKeypairResourceSchema(),
	}
}

func resourcePingFederateKeypairSigningResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	if _, ok := d.GetOk("file_data"); ok {
		input := keyPairsSigning.ImportKeyPairInput{
			Body: pf.PKCS12File{
				CryptoProvider:    nil,
				EncryptedPassword: nil,
				FileData:          String(d.Get("file_data").(string)),
				Id:                nil,
				Password:          String(d.Get("password").(string)),
			},
		}
		result, _, err := svc.ImportKeyPair(&input)
		if err != nil {
			return diag.Errorf("unable to create SigningKeypair: %s", err)
		}

		d.SetId(*result.Id)
		return resourceKeypairResourceReadResult(d, result)
	}
	input := keyPairsSigning.CreateKeyPairInput{
		Body: pf.NewKeyPairSettings{
			CommonName:   String(d.Get("common_name").(string)),
			Country:      String(d.Get("country").(string)),
			KeyAlgorithm: String(d.Get("key_algorithm").(string)),
			KeySize:      Int(d.Get("key_size").(int)),
			Organization: String(d.Get("organization").(string)),
			ValidDays:    Int(d.Get("valid_days").(int)),
		},
	}
	if val, ok := d.GetOk("city"); ok {
		input.Body.City = String(val.(string))
	}
	if val, ok := d.GetOk("organization_unit"); ok {
		input.Body.OrganizationUnit = String(val.(string))
	}
	if val, ok := d.GetOk("state"); ok {
		input.Body.State = String(val.(string))
	}
	if val, ok := d.GetOk("crypto_provider"); ok {
		input.Body.CryptoProvider = String(val.(string))
	}
	if val, ok := d.GetOk("subject_alternative_names"); ok {
		sans := expandStringList(val.(*schema.Set).List())
		input.Body.SubjectAlternativeNames = &sans
	}

	result, _, err := svc.CreateKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to generate SigningKeypair: %s", err)
	}

	d.SetId(*result.Id)
	return resourceKeypairResourceReadResult(d, result)

}

func resourcePingFederateKeypairSigningResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to read SigningKeypair: %s", err)
	}
	return resourceKeypairResourceReadResult(d, result)
}

func resourcePingFederateKeypairSigningResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.DeleteKeyPairInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to delete SigningKeypair: %s", err)
	}
	return nil
}
