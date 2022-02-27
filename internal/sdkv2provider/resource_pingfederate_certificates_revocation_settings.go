package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesRevocation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateCertificatesRevocationSettingsResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages Certificate Revocation Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateCertificatesRevocationSettingsResourceCreate,
		ReadContext:   resourcePingFederateCertificatesRevocationSettingsResourceRead,
		UpdateContext: resourcePingFederateCertificatesRevocationSettingsResourceUpdate,
		DeleteContext: resourcePingFederateCertificatesRevocationSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateCertificatesRevocationSettingsResourceSchema(),
	}
}

func resourcePingFederateCertificatesRevocationSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"crl_settings": {
			Type:        schema.TypeList,
			Elem:        resourceCrlSettings(),
			Description: "Certificate revocation CRL settings. CRL revocation is enabled by default. It will be disabled if this attribute is not defined in the request body.",
			Optional:    true,
		},
		"ocsp_settings": {
			Type:        schema.TypeList,
			Elem:        resourceOcspSettings(),
			Description: "Certificate revocation OCSP settings. OCSP revocation is disabled by default. It will be enabled if this attribute is defined in the request body.",
			Optional:    true,
		},
		"proxy_settings": {
			Type:        schema.TypeList,
			Elem:        resourceProxySettings(),
			Description: "If OCSP messaging is routed through a proxy server, specify the server's host (DNS name or IP address) and the port number. The same proxy information applies to CRL checking, when CRL is enabled for failover.",
			Optional:    true,
		},
	}
}

func resourcePingFederateCertificatesRevocationSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesRevocation
	input := certificatesRevocation.UpdateRevocationSettingsInput{
		Body: *resourcePingFederateCertificatesRevocationSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateRevocationSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to set CertificatesRevocationSettings: %s", err)
	}
	d.SetId("certificates_revocation_settings")
	return resourcePingFederateCertificatesRevocationSettingsResourceReadResult(d, result)
}

func resourcePingFederateCertificatesRevocationSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesRevocation
	result, _, err := svc.GetRevocationSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read CertificatesRevocationSettings: %s", err)
	}
	return resourcePingFederateCertificatesRevocationSettingsResourceReadResult(d, result)
}

func resourcePingFederateCertificatesRevocationSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).CertificatesRevocation
	input := certificatesRevocation.UpdateRevocationSettingsInput{
		Body: *resourcePingFederateCertificatesRevocationSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateRevocationSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update CertificatesRevocation: %s", err)
	}

	return resourcePingFederateCertificatesRevocationSettingsResourceReadResult(d, result)
}

func resourcePingFederateCertificatesRevocationSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateCertificatesRevocationSettingsResourceReadResult(d *schema.ResourceData, rv *pf.CertificateRevocationSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	if rv.OcspSettings != nil {
		if err := d.Set("ocsp_settings", flattenOcspSettings(rv.OcspSettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.CrlSettings != nil {
		if err := d.Set("crl_settings", flattenCrlSettings(rv.CrlSettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ProxySettings != nil {
		if err := d.Set("proxy_settings", flattenProxySettings(rv.ProxySettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateCertificatesRevocationSettingsResourceReadData(d *schema.ResourceData) *pf.CertificateRevocationSettings {
	var result pf.CertificateRevocationSettings
	if val, ok := d.GetOk("ocsp_settings"); ok {
		result.OcspSettings = expandOcspSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("crl_settings"); ok {
		result.CrlSettings = expandCrlSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("proxy_settings"); ok {
		result.ProxySettings = expandProxySettings(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result

}
