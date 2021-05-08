package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsOauthOpenIdConnect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateKeypairsOauthOpenIdConnectResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateKeypairsOauthOpenIdConnectResourceCreate,
		ReadContext:   resourcePingFederateKeypairsOauthOpenIdConnectResourceRead,
		UpdateContext: resourcePingFederateKeypairsOauthOpenIdConnectResourceUpdate,
		DeleteContext: resourcePingFederateKeypairsOauthOpenIdConnectResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateKeypairsOauthOpenIdConnectResourceSchema(),
	}
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"static_jwks_enabled": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"p256_active_cert_ref":   resourceLinkSchema(),
		"p256_previous_cert_ref": resourceLinkSchema(),
		"p256_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"p384_active_cert_ref":   resourceLinkSchema(),
		"p384_previous_cert_ref": resourceLinkSchema(),
		"p384_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"p521_active_cert_ref":   resourceLinkSchema(),
		"p521_previous_cert_ref": resourceLinkSchema(),
		"p521_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"rsa_active_cert_ref":   resourceRequiredLinkSchema(),
		"rsa_previous_cert_ref": resourceLinkSchema(),
		"rsa_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"p256_decryption_active_cert_ref":   resourceLinkSchema(),
		"p256_decryption_previous_cert_ref": resourceLinkSchema(),
		"p256_decryption_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"p384_decryption_active_cert_ref":   resourceLinkSchema(),
		"p384_decryption_previous_cert_ref": resourceLinkSchema(),
		"p384_decryption_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"p521_decryption_active_cert_ref":   resourceLinkSchema(),
		"p521_decryption_previous_cert_ref": resourceLinkSchema(),
		"p521_decryption_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"rsa_decryption_active_cert_ref":   resourceLinkSchema(),
		"rsa_decryption_previous_cert_ref": resourceLinkSchema(),
		"rsa_decryption_publish_x5c_parameter": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsOauthOpenIdConnect
	input := keyPairsOauthOpenIdConnect.UpdateOAuthOidcKeysSettingsInput{
		Body: *resourcePingFederateKeypairsOauthOpenIdConnectResourceReadData(d),
	}
	result, _, err := svc.UpdateOAuthOidcKeysSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create KeypairsOauthOpenIdConnect: %s", err)
	}
	d.SetId("oidc_keypairs")
	return resourcePingFederateKeypairsOauthOpenIdConnectResourceReadResult(d, result)
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsOauthOpenIdConnect
	result, _, err := svc.GetOauthOidcKeysSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read KeypairsOauthOpenIdConnect: %s", err)
	}
	return resourcePingFederateKeypairsOauthOpenIdConnectResourceReadResult(d, result)
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsOauthOpenIdConnect
	input := keyPairsOauthOpenIdConnect.UpdateOAuthOidcKeysSettingsInput{
		Body: *resourcePingFederateKeypairsOauthOpenIdConnectResourceReadData(d),
	}
	result, _, err := svc.UpdateOAuthOidcKeysSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update KeypairsOauthOpenIdConnect: %s", err)
	}
	return resourcePingFederateKeypairsOauthOpenIdConnectResourceReadResult(d, result)
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsOauthOpenIdConnect
	input := keyPairsOauthOpenIdConnect.UpdateOAuthOidcKeysSettingsInput{
		Body: pf.OAuthOidcKeysSettings{
			StaticJwksEnabled: Bool(false),
		},
	}
	_, _, err := svc.UpdateOAuthOidcKeysSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to disable KeypairsOauthOpenIdConnect: %s", err)
	}
	return nil
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceReadResult(d *schema.ResourceData, rv *pf.OAuthOidcKeysSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "static_jwks_enabled", rv.StaticJwksEnabled, &diags)
	setResourceDataBoolWithDiagnostic(d, "p256_publish_x5c_parameter", rv.P256PublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "p384_publish_x5c_parameter", rv.P384PublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "p521_publish_x5c_parameter", rv.P521PublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "rsa_publish_x5c_parameter", rv.RsaPublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "p256_decryption_publish_x5c_parameter", rv.P256DecryptionPublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "p384_decryption_publish_x5c_parameter", rv.P384DecryptionPublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "p521_decryption_publish_x5c_parameter", rv.P521DecryptionPublishX5cParameter, &diags)
	setResourceDataBoolWithDiagnostic(d, "rsa_decryption_publish_x5c_parameter", rv.RsaDecryptionPublishX5cParameter, &diags)
	if rv.P256ActiveCertRef != nil {
		if err := d.Set("p256_active_cert_ref", flattenResourceLink(rv.P256ActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P256PreviousCertRef != nil {
		if err := d.Set("p256_previous_cert_ref", flattenResourceLink(rv.P256PreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P384ActiveCertRef != nil {
		if err := d.Set("p384_active_cert_ref", flattenResourceLink(rv.P384ActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P384PreviousCertRef != nil {
		if err := d.Set("p384_previous_cert_ref", flattenResourceLink(rv.P384PreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P521ActiveCertRef != nil {
		if err := d.Set("p521_active_cert_ref", flattenResourceLink(rv.P521ActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P521PreviousCertRef != nil {
		if err := d.Set("p521_previous_cert_ref", flattenResourceLink(rv.P521PreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RsaActiveCertRef != nil {
		if err := d.Set("rsa_active_cert_ref", flattenResourceLink(rv.RsaActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RsaPreviousCertRef != nil {
		if err := d.Set("rsa_previous_cert_ref", flattenResourceLink(rv.RsaPreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P256DecryptionActiveCertRef != nil {
		if err := d.Set("p256_decryption_active_cert_ref", flattenResourceLink(rv.P256DecryptionActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P256DecryptionPreviousCertRef != nil {
		if err := d.Set("p256_decryption_previous_cert_ref", flattenResourceLink(rv.P256DecryptionPreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P384DecryptionActiveCertRef != nil {
		if err := d.Set("p384_decryption_active_cert_ref", flattenResourceLink(rv.P384DecryptionActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P384DecryptionPreviousCertRef != nil {
		if err := d.Set("p384_decryption_previous_cert_ref", flattenResourceLink(rv.P384DecryptionPreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P521DecryptionActiveCertRef != nil {
		if err := d.Set("p521_decryption_active_cert_ref", flattenResourceLink(rv.P521DecryptionActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.P521DecryptionPreviousCertRef != nil {
		if err := d.Set("p521_decryption_previous_cert_ref", flattenResourceLink(rv.P521DecryptionPreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RsaDecryptionActiveCertRef != nil {
		if err := d.Set("rsa_decryption_active_cert_ref", flattenResourceLink(rv.RsaDecryptionActiveCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RsaDecryptionPreviousCertRef != nil {
		if err := d.Set("rsa_decryption_previous_cert_ref", flattenResourceLink(rv.RsaDecryptionPreviousCertRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateKeypairsOauthOpenIdConnectResourceReadData(d *schema.ResourceData) *pf.OAuthOidcKeysSettings {
	settings := &pf.OAuthOidcKeysSettings{
		StaticJwksEnabled: Bool(d.Get("static_jwks_enabled").(bool)),
		RsaActiveCertRef:  expandResourceLink(d.Get("rsa_active_cert_ref").([]interface{})[0].(map[string]interface{})),
	}
	if v, ok := d.GetOkExists("p256_publish_x5c_parameter"); ok {
		settings.P256PublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("p384_publish_x5c_parameter"); ok {
		settings.P384PublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("p521_publish_x5c_parameter"); ok {
		settings.P521PublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("rsa_publish_x5c_parameter"); ok {
		settings.RsaPublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("p256_decryption_publish_x5c_parameter"); ok {
		settings.P256DecryptionPublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("p384_decryption_publish_x5c_parameter"); ok {
		settings.P384DecryptionPublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("p521_decryption_publish_x5c_parameter"); ok {
		settings.P521DecryptionPublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("rsa_decryption_publish_x5c_parameter"); ok {
		settings.RsaDecryptionPublishX5cParameter = Bool(v.(bool))
	}
	if v, ok := d.GetOk("p256_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P256ActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p256_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P256PreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p384_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P384ActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p384_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P384PreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p521_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P521ActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p521_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P521PreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("rsa_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.RsaActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("rsa_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.RsaPreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p256_decryption_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P256DecryptionActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p256_decryption_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P256DecryptionPreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p384_decryption_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P384DecryptionActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p384_decryption_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P384DecryptionPreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p521_decryption_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P521DecryptionActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("p521_decryption_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.P521DecryptionPreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("rsa_decryption_active_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.RsaDecryptionActiveCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("rsa_decryption_previous_cert_ref"); ok && len(v.([]interface{})) > 0 {
		settings.RsaDecryptionPreviousCertRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	return settings
}
