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
		Schema: resourcePingFederateKeypairSigningResourceSchema(),
	}
}

func resourcePingFederateKeypairSigningResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"crypto_provider": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			ValidateDiagFunc: validateCryptoProvider,
		},
		"file_data": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"password": {
			Type:          schema.TypeString,
			Sensitive:     true,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"city": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"common_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"country": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_algorithm": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Computed:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_size": {
			Type:          schema.TypeInt,
			Optional:      true,
			ForceNew:      true,
			Computed:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization_unit": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"state": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"subject_alternative_names": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"valid_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"expires": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha256_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha1_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature_algorithm": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_cn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"valid_from": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
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
		return resourcePingFederateKeypairSigningResourceReadResult(d, result)
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

	result, _, err := svc.CreateKeyPair(&input)
	if err != nil {
		return diag.Errorf("unable to generate SigningKeypair: %s", err)
	}

	d.SetId(*result.Id)
	return resourcePingFederateKeypairSigningResourceReadResult(d, result)

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
	return resourcePingFederateKeypairSigningResourceReadResult(d, result)
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

func resourcePingFederateKeypairSigningResourceReadResult(d *schema.ResourceData, rv *pf.KeyPairView) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "keypair_id", rv.Id, &diags)
	setResourceDataStringWithDiagnostic(d, "crypto_provider", rv.CryptoProvider, &diags)
	setResourceDataStringWithDiagnostic(d, "expires", rv.Expires, &diags)
	setResourceDataStringWithDiagnostic(d, "issuer_dn", rv.IssuerDN, &diags)
	setResourceDataStringWithDiagnostic(d, "key_algorithm", rv.KeyAlgorithm, &diags)
	setResourceDataIntWithDiagnostic(d, "key_size", rv.KeySize, &diags)
	setResourceDataStringWithDiagnostic(d, "serial_number", rv.SerialNumber, &diags)
	setResourceDataStringWithDiagnostic(d, "sha1_fingerprint", rv.Sha1Fingerprint, &diags)
	setResourceDataStringWithDiagnostic(d, "sha256_fingerprint", rv.Sha256Fingerprint, &diags)
	setResourceDataStringWithDiagnostic(d, "signature_algorithm", rv.SignatureAlgorithm, &diags)
	setResourceDataStringWithDiagnostic(d, "status", rv.Status, &diags)
	if rv.SubjectAlternativeNames != nil && len(*rv.SubjectAlternativeNames) > 0 {
		if err := d.Set("subject_alternative_names", *rv.SubjectAlternativeNames); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "subject_dn", rv.SubjectDN, &diags)
	setResourceDataStringWithDiagnostic(d, "valid_from", rv.ValidFrom, &diags)
	setResourceDataIntWithDiagnostic(d, "version", rv.Version, &diags)
	return diags
}
