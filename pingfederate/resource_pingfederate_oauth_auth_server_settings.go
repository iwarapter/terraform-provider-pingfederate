package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthAuthServerSettingsResource() *schema.Resource {

	scopes := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}

	scopeGroups := &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"description": {
					Type:     schema.TypeString,
					Required: true,
				},
				"scopes": {
					Type:     schema.TypeList,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	return &schema.Resource{
		CreateContext: resourcePingFederateOauthAuthServerSettingsResourceCreate,
		ReadContext:   resourcePingFederateOauthAuthServerSettingsResourceRead,
		UpdateContext: resourcePingFederateOauthAuthServerSettingsResourceUpdate,
		DeleteContext: resourcePingFederateOauthAuthServerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateOauthAuthServerSettingsResourceImport,
		},

		Schema: map[string]*schema.Schema{
			"default_scope_description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scopes":                 scopes,
			"scope_groups":           scopeGroups,
			"exclusive_scopes":       scopes,
			"exclusive_scope_groups": scopeGroups,
			"authorization_code_timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"authorization_code_entropy": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"track_user_sessions_for_logout": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token_endpoint_base_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistent_grant_lifetime": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"persistent_grant_lifetime_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validatePersistentGrantLifetimeUnit,
			},
			"refresh_token_length": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"roll_refresh_token_values": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"refresh_rolling_interval": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"persistent_grant_reuse_grant_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateGrantTypes,
				},
			},
			"persistent_grant_contract": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extended_attributes": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"core_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"bypass_authorization_for_approved_grants": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_unidentified_client_ro_creds": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allowed_origins": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allow_unidentified_client_extension_grants": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"admin_web_service_pcv_ref": resourceLinkSchema(),
		},
	}
}

func resourcePingFederateOauthAuthServerSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("OauthAuthServerSettings")
	return resourcePingFederateOauthAuthServerSettingsResourceUpdate(ctx, d, m)
}

func resourcePingFederateOauthAuthServerSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthServerSettings
	result, _, err := svc.GetAuthorizationServerSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read OauthAuthServerSettings: %s", err)
	}
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
}

func resourcePingFederateOauthAuthServerSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	defaultScopeDescription := d.Get("default_scope_description").(string)
	authorizationCodeTimeout := d.Get("authorization_code_timeout").(int)
	authorizationCodeEntropy := d.Get("authorization_code_entropy").(int)
	refreshTokenLength := d.Get("refresh_token_length").(int)
	refreshRollingInterval := d.Get("refresh_rolling_interval").(int)

	authSettings := &pf.AuthorizationServerSettings{
		// AdminWebServicePcvRef:             *ResourceLink
		AuthorizationCodeEntropy: Int(authorizationCodeEntropy),
		AuthorizationCodeTimeout: Int(authorizationCodeTimeout),
		DefaultScopeDescription:  String(defaultScopeDescription),
		RefreshRollingInterval:   Int(refreshRollingInterval),
		RefreshTokenLength:       Int(refreshTokenLength),
	}

	//TODO
	//if _, ok := d.GetOk("admin_web_service_pcv_ref"); ok {
	//	// authSettings.AdminWebServicePcvRef = //expand
	//}
	isPF10 := m.(pfClient).IsPF10()

	if v, ok := d.GetOk("allowed_origins"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		authSettings.AllowedOrigins = &strs
	}

	if v, ok := d.GetOkExists("allow_unidentified_client_extension_grants"); ok {
		authSettings.AllowUnidentifiedClientExtensionGrants = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("allow_unidentified_client_ro_creds"); ok {
		authSettings.AllowUnidentifiedClientROCreds = Bool(v.(bool))
	}

	if v, ok := d.GetOk("bypass_authorization_for_approved_grants"); ok {
		authSettings.BypassAuthorizationForApprovedGrants = Bool(v.(bool))
	}

	if v, ok := d.GetOk("exclusive_scope_groups"); ok {
		authSettings.ExclusiveScopeGroups = expandScopeGroups(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("exclusive_scopes"); ok {
		authSettings.ExclusiveScopes = expandScopes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("persistent_grant_contract"); ok && v.(*schema.Set).Len() > 0 {
		authSettings.PersistentGrantContract = expandPersistentGrantContract(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("persistent_grant_lifetime"); ok {
		authSettings.PersistentGrantLifetime = Int(v.(int))
	} else if isPF10 {
		authSettings.PersistentGrantLifetime = Int(-1)
	}

	if v, ok := d.GetOk("persistent_grant_lifetime_unit"); ok {
		authSettings.PersistentGrantLifetimeUnit = String(v.(string))
	} else if isPF10 {
		authSettings.PersistentGrantLifetimeUnit = String("DAYS")
	}

	if v, ok := d.GetOk("persistent_grant_reuse_grant_types"); ok {
		authSettings.PersistentGrantReuseGrantTypes = expandStringList(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("scopes"); ok {
		authSettings.Scopes = expandScopes(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("scope_groups"); ok {
		authSettings.ScopeGroups = expandScopeGroups(v.(*schema.Set).List())
	}

	if v, ok := d.GetOkExists("roll_refresh_token_values"); ok {
		authSettings.RollRefreshTokenValues = Bool(v.(bool))
	}

	if v, ok := d.GetOk("token_endpoint_base_url"); ok {
		authSettings.TokenEndpointBaseUrl = String(v.(string))
	}

	if v, ok := d.GetOkExists("track_user_sessions_for_logout"); ok {
		authSettings.TrackUserSessionsForLogout = Bool(v.(bool))
	}

	svc := m.(pfClient).OauthAuthServerSettings
	input := &oauthAuthServerSettings.UpdateAuthorizationServerSettingsInput{
		Body: *authSettings,
	}

	result, _, err := svc.UpdateAuthorizationServerSettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to update OauthAuthServerSettings: %s", err)
	}
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
}

func resourcePingFederateOauthAuthServerSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	//resource cannot be deleted just not tracked by terraform anymore
	return nil
}

func resourcePingFederateOauthAuthServerSettingsResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).OauthAuthServerSettings
	result, _, err := svc.GetAuthorizationServerSettingsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateOauthAuthServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthorizationServerSettings, isPF10 bool) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "default_scope_description", rv.DefaultScopeDescription, &diags)
	setResourceDataIntWithDiagnostic(d, "authorization_code_timeout", rv.AuthorizationCodeTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "authorization_code_entropy", rv.AuthorizationCodeEntropy, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_token_length", rv.RefreshTokenLength, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_rolling_interval", rv.RefreshRollingInterval, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_unidentified_client_extension_grants", rv.AllowUnidentifiedClientExtensionGrants, &diags)
	setResourceDataBoolWithDiagnostic(d, "track_user_sessions_for_logout", rv.TrackUserSessionsForLogout, &diags)
	setResourceDataStringWithDiagnostic(d, "token_endpoint_base_url", rv.TokenEndpointBaseUrl, &diags)
	if rv.PersistentGrantLifetime != nil && *rv.PersistentGrantLifetime == -1 && isPF10 {
		rv.PersistentGrantLifetime = nil
	}
	setResourceDataIntWithDiagnostic(d, "persistent_grant_lifetime", rv.PersistentGrantLifetime, &diags)
	if rv.PersistentGrantLifetimeUnit != nil && *rv.PersistentGrantLifetimeUnit == "DAYS" && isPF10 {
		rv.PersistentGrantLifetimeUnit = nil
	}
	setResourceDataStringWithDiagnostic(d, "persistent_grant_lifetime_unit", rv.PersistentGrantLifetimeUnit, &diags)
	setResourceDataBoolWithDiagnostic(d, "roll_refresh_token_values", rv.RollRefreshTokenValues, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_authorization_for_approved_grants", rv.BypassAuthorizationForApprovedGrants, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_unidentified_client_ro_creds", rv.AllowUnidentifiedClientROCreds, &diags)

	// "admin_web_service_pcv_ref"

	if rv.PersistentGrantReuseGrantTypes != nil && len(rv.PersistentGrantReuseGrantTypes) > 0 {
		if err := d.Set("persistent_grant_reuse_grant_types", rv.PersistentGrantReuseGrantTypes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AllowedOrigins != nil && len(*rv.AllowedOrigins) > 0 {
		if err := d.Set("allowed_origins", *rv.AllowedOrigins); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.PersistentGrantContract != nil {
		if err := d.Set("persistent_grant_contract", flattenPersistentGrantContract(rv.PersistentGrantContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Scopes != nil && len(*rv.Scopes) > 0 {
		if err := d.Set("scopes", flattenScopes(*rv.Scopes)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ScopeGroups != nil && len(*rv.ScopeGroups) > 0 {
		if err := d.Set("scope_groups", flattenScopeGroups(*rv.ScopeGroups)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ExclusiveScopes != nil && len(*rv.ExclusiveScopes) > 0 {
		if err := d.Set("exclusive_scopes", flattenScopes(*rv.ExclusiveScopes)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ExclusiveScopeGroups != nil && len(*rv.ExclusiveScopeGroups) > 0 {
		if err := d.Set("exclusive_scope_groups", flattenScopeGroups(*rv.ExclusiveScopeGroups)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}
