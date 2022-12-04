package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePingFederateKeyPairSigning() *schema.Resource {
	sch := resourceKeypairResourceSchema()
	delete(sch, "file_data")
	delete(sch, "password")
	for s := range sch {
		sch[s].Computed = true
		sch[s].Required = false
		sch[s].Optional = false
		sch[s].Default = nil
		sch[s].ValidateDiagFunc = nil
		sch[s].ValidateFunc = nil
		sch[s].ConflictsWith = []string{}
		sch[s].RequiredWith = []string{}
	}
	sch["id"] = &schema.Schema{Type: schema.TypeString, Computed: true}
	return &schema.Resource{
		Description: "Use this data source to get all signing keypairs in Ping Federate.",
		ReadContext: dataSourcePingFederateKeyPairSigningRead,
		Schema: map[string]*schema.Schema{
			"keys": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: sch,
				},
			},
		},
	}
}

func dataSourcePingFederateKeyPairSigningRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	svc := m.(pfClient).KeyPairsSigning
	result, _, err := svc.GetKeyPairsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read KeyPairSigning: %s", err)

	}

	keys := make([]map[string]interface{}, 0, len(*result.Items))
	for _, v := range *result.Items {
		s := make(map[string]interface{})
		if v.Id != nil {
			s["id"] = *v.Id
		}
		if v.CryptoProvider != nil {
			s["crypto_provider"] = *v.CryptoProvider
		}
		if v.Expires != nil {
			s["expires"] = *v.Expires
		}
		if v.IssuerDN != nil {
			s["issuer_dn"] = *v.IssuerDN
		}
		if v.KeyAlgorithm != nil {
			s["key_algorithm"] = *v.KeyAlgorithm
		}
		if v.KeySize != nil {
			s["key_size"] = *v.KeySize
		}
		if v.SerialNumber != nil {
			s["serial_number"] = *v.SerialNumber
		}
		if v.Sha1Fingerprint != nil {
			s["sha1_fingerprint"] = *v.Sha1Fingerprint
		}
		if v.Sha256Fingerprint != nil {
			s["sha256_fingerprint"] = *v.Sha256Fingerprint
		}
		if v.SignatureAlgorithm != nil {
			s["signature_algorithm"] = *v.SignatureAlgorithm
		}
		if v.Status != nil {
			s["status"] = *v.Status
		}
		if v.SubjectAlternativeNames != nil {
			s["subject_alternative_names"] = *v.SubjectAlternativeNames
		}
		if v.SubjectDN != nil {
			s["subject_dn"] = *v.SubjectDN
		}
		if v.ValidFrom != nil {
			s["valid_from"] = *v.ValidFrom
		}
		if v.Version != nil {
			s["version"] = *v.Version
		}
		keys = append(keys, s)
	}
	if err := d.Set("keys", keys); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	d.SetId("keypairs_signing")
	return diags
}
