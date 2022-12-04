package sdkv2provider

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesCa"
)

func resourcePingFederateCertificatesCaResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for Certificate CAs within PingFederate.",
		CreateContext: resourcePingFederateCertificatesCaResourceCreate,
		ReadContext:   resourcePingFederateCertificatesCaResourceRead,
		DeleteContext: resourcePingFederateCertificatesCaResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateCertificatesCaResourceImport,
		},
		Schema: resourcePingFederateCertificatesCaResourceSchema(),
	}
}

func resourcePingFederateCertificatesCaResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the certificate.",
		},
		"file_data": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringIsBase64,
			Description:  "Base64 Encoded certificate data.",
		},
		"crypto_provider": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Description:      "Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.",
			ValidateDiagFunc: validateCryptoProvider,
		},
		"expires": {
			Type:        schema.TypeString,
			Description: "The end date up until which the item is valid, in ISO 8601 format (UTC).",
			Computed:    true,
		},
		"issuer_dn": {
			Type:        schema.TypeString,
			Description: "The issuer's distinguished name.",
			Computed:    true,
		},
		"key_algorithm": {
			Type:        schema.TypeString,
			Description: "The public key algorithm.",
			Computed:    true,
		},
		"key_size": {
			Type:        schema.TypeInt,
			Description: "The public key size.",
			Computed:    true,
		},
		"serial_number": {
			Type:        schema.TypeString,
			Description: "The serial number assigned by the CA.",
			Computed:    true,
		},
		"sha1_fingerprint": {
			Type:        schema.TypeString,
			Description: "SHA-1 fingerprint in Hex encoding.",
			Computed:    true,
		},
		"sha256_fingerprint": {
			Type:        schema.TypeString,
			Description: "SHA-256 fingerprint in Hex encoding.",
			Computed:    true,
		},
		"signature_algorithm": {
			Type:        schema.TypeString,
			Description: "The signature algorithm.",
			Computed:    true,
		},
		"status": {
			Type:        schema.TypeString,
			Description: "Status of the item.",
			Computed:    true,
		},
		"subject_alternative_names": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The subject alternative names (SAN).",
		},
		"subject_dn": {
			Type:        schema.TypeString,
			Description: "The subject's distinguished name.",
			Computed:    true,
		},
		"valid_from": {
			Type:        schema.TypeString,
			Description: "The start date from which the item is valid, in ISO 8601 format (UTC).",
			Computed:    true,
		},
		"version": {
			Type:        schema.TypeInt,
			Description: "The X.509 version to which the item conforms.",
			Computed:    true,
		},
	}
}

func resourcePingFederateCertificatesCaResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("certificates_ca")
	defer awsMutexKV.Unlock("certificates_ca")

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
	result, _, err := svc.ImportTrustedCAWithContext(ctx, &input)

	if err != nil {
		return diag.Errorf("unable to create CertificatesCa: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateCertificatesCaResourceReadResult(d, result)
}

func resourcePingFederateCertificatesCaResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.GetTrustedCertInput{
		Id: d.Id(),
	}

	result, _, err := svc.GetTrustedCertWithContext(ctx, &input)

	if err != nil {
		return diag.Errorf("unable to read CertificatesCa: %s", err)
	}
	return resourcePingFederateCertificatesCaResourceReadResult(d, result)
}

func resourcePingFederateCertificatesCaResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("certificates_ca")
	defer awsMutexKV.Unlock("certificates_ca")

	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.DeleteTrustedCAInput{
		Id: d.Id(),
	}

	_, _, err := svc.DeleteTrustedCAWithContext(ctx, &input)

	if err != nil {
		return diag.Errorf("unable to delete CertificatesCa: %s", err)
	}
	return nil
}

func resourcePingFederateCertificatesCaResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	diags := resourcePingFederateCertificatesCaResourceRead(ctx, d, m)
	if diags.HasError() {
		msg := []string{}
		for _, diagnostic := range diags {
			msg = append(msg, diagnostic.Summary)
		}
		return nil, fmt.Errorf("unable to retrieve certifcate information:\n%s", strings.Join(msg, "\n"))
	}

	svc := m.(pfClient).CertificatesCa
	input := certificatesCa.ExportCertificateFileInput{
		Id: d.Id(),
	}
	result, _, err := svc.ExportCertificateFileWithContext(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve certifcate file data %s", err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(*result))
	setResourceDataStringWithDiagnostic(d, "certificate_id", String(d.Id()), &diags)
	setResourceDataStringWithDiagnostic(d, "file_data", String(encoded), &diags)

	return []*schema.ResourceData{d}, nil
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

	if rv.SubjectAlternativeNames != nil {
		if err := d.Set("subject_alternative_names", *rv.SubjectAlternativeNames); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags

}
