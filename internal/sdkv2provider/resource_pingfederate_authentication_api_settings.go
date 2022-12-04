package sdkv2provider

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationApi"
)

func resourcePingFederateAuthnApiSettingsResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages Authentication API Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateAuthnApiSettingsResourceCreate,
		ReadContext:   resourcePingFederateAuthnApiSettingsResourceRead,
		UpdateContext: resourcePingFederateAuthnApiSettingsResourceUpdate,
		DeleteContext: resourcePingFederateAuthnApiSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthnApiSettingsResourceSchema(),
	}
}

func resourcePingFederateAuthnApiSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"api_enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Specifies whether the authentication API is enabled. The default value is false.",
		},
		"enable_api_descriptions": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Enable the API Descriptions endpoint.",
		},
		"default_application_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem:        resourceLinkResource(),
			Description: "Application for non authentication policy use cases.",
		},
	}
}

func resourcePingFederateAuthnApiSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.UpdateAuthenticationApiSettingsInput{
		Body: *resourcePingFederateAuthnApiSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationApiSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create AuthnApiSettings: %s", err)
	}
	d.SetId("default_authentication_api_settings")
	return resourcePingFederateAuthnApiSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	result, _, err := svc.GetAuthenticationApiSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read AuthnApiSettings: %s", err)
	}
	return resourcePingFederateAuthnApiSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.UpdateAuthenticationApiSettingsInput{
		Body: *resourcePingFederateAuthnApiSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationApiSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update AuthnApiSettings: %s", err)
	}
	return resourcePingFederateAuthnApiSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateAuthnApiSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthnApiSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "api_enabled", rv.ApiEnabled, &diags)
	setResourceDataBoolWithDiagnostic(d, "enable_api_descriptions", rv.EnableApiDescriptions, &diags)
	if rv.DefaultApplicationRef != nil {
		if err := d.Set("default_application_ref", flattenResourceLink(rv.DefaultApplicationRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateAuthnApiSettingsResourceReadData(d *schema.ResourceData) *pf.AuthnApiSettings {
	result := pf.AuthnApiSettings{
		ApiEnabled: Bool(d.Get("api_enabled").(bool)),
	}
	if val, ok := d.GetOkExists("enable_api_descriptions"); ok {
		result.EnableApiDescriptions = Bool(val.(bool))
	}
	if v, ok := d.GetOk("default_application_ref"); ok && len(v.([]interface{})) > 0 {
		result.DefaultApplicationRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}
