package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOAuthAccessTokenManagerSettingsResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages OAuth Access Token Manager Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateOAuthAccessTokenManagerSettingsResourceCreate,
		ReadContext:   resourcePingFederateOAuthAccessTokenManagerSettingsResourceRead,
		UpdateContext: resourcePingFederateOAuthAccessTokenManagerSettingsResourceUpdate,
		DeleteContext: resourcePingFederateOAuthAccessTokenManagerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateOAuthAccessTokenManagerSettingsResourceSchema(),
	}
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_access_token_manager_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Elem:        resourceLinkResource(),
			Description: "Reference to the default access token manager, if one is defined.",
		},
	}
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.UpdateSettingsInput{
		Body: *resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create OAuthAccessTokenManagerSettings: %s", err)
	}
	d.SetId("default_oauth_access_token_manager_settings")
	return resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	result, _, err := svc.GetSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read OAuthAccessTokenManagerSettings: %s", err)
	}
	return resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.UpdateSettingsInput{
		Body: *resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OAuthAccessTokenManagerSettings: %s", err)
	}
	return resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AccessTokenManagementSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	if rv.DefaultAccessTokenManagerRef != nil {
		if err := d.Set("default_access_token_manager_ref", flattenResourceLink(rv.DefaultAccessTokenManagerRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateOAuthAccessTokenManagerSettingsResourceReadData(d *schema.ResourceData) *pf.AccessTokenManagementSettings {
	result := pf.AccessTokenManagementSettings{}
	if v, ok := d.GetOk("default_access_token_manager_ref"); ok && len(v.([]interface{})) > 0 {
		result.DefaultAccessTokenManagerRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}
