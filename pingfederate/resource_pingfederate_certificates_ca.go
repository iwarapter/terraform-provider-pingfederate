package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesCa"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateCertificatesCaResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateCertificatesCaResourceCreate,
		ReadContext:   resourcePingFederateCertificatesCaResourceRead,
		DeleteContext: resourcePingFederateCertificatesCaResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateCertificatesCaResourceSchema(),
	}
}

func resourcePingFederateCertificatesCaResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"file_data": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"crypto_provider": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"expires": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_algorithm": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_size": {
			Type:     schema.TypeInt,
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
		"sha256_fingerprint": {
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
		"subject_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_alternative_names": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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

func resourcePingFederateCertificatesCaResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.ImportTrustedCAInput{
		Body: pf.X509File{
			FileData: String(d.Get("file_data").(string)),
		},
	}
	if v, ok := d.GetOk("certificate_id"); ok {
		input.Body.Id = String(v.(string))
	}
	if v, ok := d.GetOk("crypto_provider"); ok {
		input.Body.CryptoProvider = String(v.(string))
	}
	result, _, err := svc.ImportTrustedCA(&input)

	if err != nil {
		return diag.Errorf("unable to create CertificatesCa: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateCertificatesCaResourceReadResult(d, result)
}

func resourcePingFederateCertificatesCaResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.GetTrustedCertInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetTrustedCert(&input)

	if err != nil {
		return diag.Errorf("unable to read CertificatesCa: %s", err)
	}
	return resourcePingFederateCertificatesCaResourceReadResult(d, result)
}

func resourcePingFederateCertificatesCaResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.DeleteTrustedCAInput{
		Id: d.Id(),
	}

	_, _, err := svc.DeleteTrustedCA(&input)

	if err != nil {
		return diag.Errorf("unable to delete CertificatesCa: %s", err)
	}
	return nil
}

func resourcePingFederateCertificatesCaResourceReadResult(d *schema.ResourceData, rv *pf.CertView) diag.Diagnostics {
	var diags diag.Diagnostics
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
	setResourceDataStringWithDiagnostic(d, "subject_dn", rv.SubjectDN, &diags)
	setResourceDataStringWithDiagnostic(d, "valid_from", rv.ValidFrom, &diags)
	setResourceDataIntWithDiagnostic(d, "version", rv.Version, &diags)

	if rv.SubjectAlternativeNames != nil && len(*rv.SubjectAlternativeNames) > 0 {
		if err := d.Set("subject_alternative_names", *rv.SubjectAlternativeNames); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags

}
