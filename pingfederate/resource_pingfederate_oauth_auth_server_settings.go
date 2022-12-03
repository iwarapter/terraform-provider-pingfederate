package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthAuthServerSettingsResource() *schema.Resource {

	scopes := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dynamic": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}

	scopeGroups := &schema.Resource{
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
	}

	return &schema.Resource{
		Description: `Provides a OAuth Authorization Server Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateOauthAuthServerSettingsResourceCreate,
		ReadContext:   resourcePingFederateOauthAuthServerSettingsResourceRead,
		UpdateContext: resourcePingFederateOauthAuthServerSettingsResourceUpdate,
		DeleteContext: resourcePingFederateOauthAuthServerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateOauthAuthServerSettingsResourceImport,
		},

		Schema: map[string]*schema.Schema{
			"admin_web_service_pcv_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The password credential validator reference that is used for authenticating access to the OAuth Administrative Web Service.",
				Elem:        resourceLinkResource(),
			},
			"allow_unidentified_client_extension_grants": {
				Type:        schema.TypeBool,
				Description: "Allow unidentified clients to request extension grants. The default value is false.",
				Optional:    true,
				Default:     false,
			},
			"allow_unidentified_client_ro_creds": {
				Type:        schema.TypeBool,
				Description: "Allow unidentified clients to request resource owner password credentials grants. The default value is false.",
				Optional:    true,
				Default:     false,
			},
			"allowed_origins": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of allowed origins.",
			},
			"approved_scope_attribute": { //TODO deprecated
				Type:          schema.TypeString,
				Deprecated:    "This attribute is incorrectly named and will be removed in future releases, please use approved_scopes_attribute",
				Optional:      true,
				ConflictsWith: []string{"approved_scopes_attribute"},
			},
			"approved_scopes_attribute": {
				Type:          schema.TypeString,
				Description:   "Attribute from the external consent adapter's contract, intended for storing approved scopes returned by the external consent page.",
				Optional:      true,
				ConflictsWith: []string{"approved_scope_attribute"},
			},
			"atm_id_for_oauth_grant_management": {
				Type:        schema.TypeString,
				Description: "The ID of the Access Token Manager used for OAuth enabled grant management.",
				Optional:    true,
			},
			"authorization_code_entropy": {
				Type:        schema.TypeInt,
				Description: "The authorization code entropy, in bytes.",
				Required:    true,
			},
			"authorization_code_timeout": {
				Type:        schema.TypeInt,
				Description: "The authorization code timeout, in seconds.",
				Required:    true,
			},
			"bypass_activation_code_confirmation": {
				Type:        schema.TypeBool,
				Description: "Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device.",
				Optional:    true,
			},
			"bypass_authorization_for_approved_grants": {
				Type:        schema.TypeBool,
				Description: "Bypass authorization for previously approved persistent grants. The default value is false.",
				Optional:    true,
				Default:     false,
			},
			"default_scope_description": {
				Type:        schema.TypeString,
				Description: "The default scope description.",
				Required:    true,
			},
			"device_polling_interval": {
				Type:        schema.TypeInt,
				Description: "The amount of time client should wait between polling requests, in seconds.",
				Optional:    true,
				Default:     5,
			},
			"exclusive_scope_groups": {
				Type:        schema.TypeSet,
				Elem:        scopeGroups,
				Description: "The list of exclusive scope groups.",
				Optional:    true,
			},
			"exclusive_scopes": {
				Type:        schema.TypeSet,
				Elem:        scopes,
				Description: "The list of exclusive scopes.",
				Optional:    true,
			},
			"par_reference_length": {
				Type:        schema.TypeInt,
				Description: "The entropy of pushed authorization request references, in bytes. The default value is 24.",
				Optional:    true,
				Default:     24,
			},
			"par_reference_timeout": {
				Type:        schema.TypeInt,
				Description: "The timeout, in seconds, of the pushed authorization request reference. The default value is 60.",
				Optional:    true,
				Default:     60,
			},
			"par_status": {
				Type:        schema.TypeString,
				Description: "The status of pushed authorization request support. The default value is ENABLED.",
				Optional:    true,
				Default:     "ENABLED",
				ValidateFunc: validation.StringInSlice([]string{
					"ENABLED", "DISABLED", "REQUIRED",
				}, false),
			},
			"pending_authorization_timeout": {
				Type:        schema.TypeInt,
				Description: "The 'device_code' and 'user_code' timeout, in seconds.",
				Optional:    true,
				Default:     600,
			},
			"persistent_grant_contract": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "The persistent grant contract defines attributes that are associated with OAuth persistent grants.",
				MaxItems:    1,
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
			"persistent_grant_idle_timeout": {
				Type:        schema.TypeInt,
				Description: "The persistent grant idle timeout. The default value is 30 (days). -1 indicates an indefinite amount of time.",
				Optional:    true,
				Default:     30,
			},
			"persistent_grant_idle_timeout_time_unit": {
				Type:             schema.TypeString,
				Description:      "The persistent grant idle timeout time unit.",
				Optional:         true,
				Default:          "DAYS",
				ValidateDiagFunc: validatePersistentGrantLifetimeUnit,
			},
			"persistent_grant_lifetime": {
				Type:        schema.TypeInt,
				Description: "The persistent grant lifetime. The default value is indefinite. -1 indicates an indefinite amount of time.",
				Optional:    true,
				Default:     -1,
			},
			"persistent_grant_lifetime_unit": {
				Type:             schema.TypeString,
				Description:      "The persistent grant lifetime unit.",
				Optional:         true,
				ValidateDiagFunc: validatePersistentGrantLifetimeUnit,
			},
			"persistent_grant_reuse_grant_types": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateGrantTypes,
				},
				Description: "The grant types that the OAuth AS can reuse rather than creating a new grant for each request.",
			},
			"refresh_rolling_interval": {
				Type:        schema.TypeInt,
				Description: "The minimum interval to roll refresh tokens, in hours.",
				Required:    true,
			},
			"refresh_token_length": {
				Type:        schema.TypeInt,
				Description: "The refresh token length in number of characters.",
				Required:    true,
			},
			"registered_authorization_path": {
				Type:        schema.TypeString,
				Description: "The Registered Authorization Path is concatenated to PingFederate base URL to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request. PingFederate listens to this path if specified",
				Optional:    true,
			},
			"roll_refresh_token_values": {
				Type:        schema.TypeBool,
				Description: "The roll refresh token values default policy. The default value is true.",
				Optional:    true,
				Default:     true,
			},
			"scope_for_oauth_grant_management": {
				Type:        schema.TypeString,
				Description: "The OAuth scope to validate when accessing grant management service.",
				Optional:    true,
			},
			"scope_groups": {
				Type:        schema.TypeSet,
				Elem:        scopeGroups,
				Description: "The list of common scope groups.",
				Optional:    true,
			},
			"scopes": {
				Type:        schema.TypeSet,
				Elem:        scopes,
				Description: "The list of common scopes.",
				Optional:    true,
			},
			"token_endpoint_base_url": {
				Type:        schema.TypeString,
				Description: "The token endpoint base URL used to validate the 'aud' claim during Private Key JWT Client Authentication.",
				Optional:    true,
			},
			"track_user_sessions_for_logout": {
				Type:        schema.TypeBool,
				Description: "Determines whether user sessions are tracked for logout. If this property is not provided on a PUT, the setting is left unchanged.",
				Optional:    true,
			},
			"user_authorization_consent_adapter": {
				Type:        schema.TypeString,
				Description: "Adapter ID of the external consent adapter to be used for the consent page user interface.",
				Optional:    true,
			},
			"user_authorization_consent_page_setting": {
				Type:         schema.TypeString,
				Description:  "User Authorization Consent Page setting to use PingFederate's internal consent page or an external system",
				Optional:     true,
				Default:      "INTERNAL",
				ValidateFunc: validation.StringInSlice([]string{"INTERNAL", "ADAPTER"}, false),
			},
			"user_authorization_url": {
				Type:        schema.TypeString,
				Description: "The URL used to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request",
				Optional:    true,
			},
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
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	authSettings := resourcePingFederateOauthAuthServerSettingsResourceReadData(d, m.(pfClient).PfVersion())

	svc := m.(pfClient).OauthAuthServerSettings
	input := &oauthAuthServerSettings.UpdateAuthorizationServerSettingsInput{
		Body: *authSettings,
	}

	result, _, err := svc.UpdateAuthorizationServerSettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to update OauthAuthServerSettings: %s", err)
	}
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceReadData(d *schema.ResourceData, pfVersion string) *pf.AuthorizationServerSettings {
	re := regexp.MustCompile(`^(10\.[0-9])`)
	isPF10 := re.MatchString(pfVersion)
	re = regexp.MustCompile(`^(10\.[2-9])`)
	isPF10_2 := re.MatchString(pfVersion)

	defaultScopeDescription := d.Get("default_scope_description").(string)
	authorizationCodeTimeout := d.Get("authorization_code_timeout").(int)
	authorizationCodeEntropy := d.Get("authorization_code_entropy").(int)
	refreshTokenLength := d.Get("refresh_token_length").(int)
	refreshRollingInterval := d.Get("refresh_rolling_interval").(int)

	authSettings := &pf.AuthorizationServerSettings{
		AuthorizationCodeEntropy: Int(authorizationCodeEntropy),
		AuthorizationCodeTimeout: Int(authorizationCodeTimeout),
		DefaultScopeDescription:  String(defaultScopeDescription),
		RefreshRollingInterval:   Int(refreshRollingInterval),
		RefreshTokenLength:       Int(refreshTokenLength),
	}

	if v, ok := d.GetOk("admin_web_service_pcv_ref"); ok {
		authSettings.AdminWebServicePcvRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}

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

	if v, ok := d.GetOk("persistent_grant_reuse_grant_types"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		authSettings.PersistentGrantReuseGrantTypes = &strs
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
	if v, ok := d.GetOk("atm_id_for_oauth_grant_management"); ok {
		authSettings.AtmIdForOAuthGrantManagement = String(v.(string))
	}
	if v, ok := d.GetOk("approved_scope_attribute"); ok { //TODO deprecated
		authSettings.ApprovedScopesAttribute = String(v.(string))
	}
	if v, ok := d.GetOk("approved_scopes_attribute"); ok {
		authSettings.ApprovedScopesAttribute = String(v.(string))
	}
	if v, ok := d.GetOkExists("bypass_activation_code_confirmation"); ok {
		authSettings.BypassActivationCodeConfirmation = Bool(v.(bool))
	}
	if v, ok := d.GetOk("device_polling_interval"); ok {
		authSettings.DevicePollingInterval = Int(v.(int))
	}

	if isPF10 {
		if v, ok := d.GetOk("persistent_grant_lifetime"); ok {
			authSettings.PersistentGrantLifetime = Int(v.(int))
		}
		if v, ok := d.GetOk("persistent_grant_lifetime_unit"); ok {
			authSettings.PersistentGrantLifetimeUnit = String(v.(string))
		}
	} else {
		_ = d.Set("persistent_grant_lifetime", -1)
		_ = d.Set("persistent_grant_lifetime_unit", "DAYS")
	}

	if isPF10_2 {
		if v, ok := d.GetOk("par_reference_length"); ok {
			authSettings.ParReferenceLength = Int(v.(int))
		}
		if v, ok := d.GetOk("par_reference_timeout"); ok {
			authSettings.ParReferenceTimeout = Int(v.(int))
		}
		if v, ok := d.GetOk("par_status"); ok {
			authSettings.ParStatus = String(v.(string))
		}
	} else {
		_ = d.Set("par_reference_length", 24)
		_ = d.Set("par_reference_timeout", 60)
		_ = d.Set("par_status", "ENABLED")
	}

	if v, ok := d.GetOk("pending_authorization_timeout"); ok {
		authSettings.PendingAuthorizationTimeout = Int(v.(int))
	}
	if v, ok := d.GetOk("persistent_grant_idle_timeout"); ok {
		authSettings.PersistentGrantIdleTimeout = Int(v.(int))
	}
	if v, ok := d.GetOk("persistent_grant_idle_timeout_time_unit"); ok {
		authSettings.PersistentGrantIdleTimeoutTimeUnit = String(v.(string))
	}
	if v, ok := d.GetOk("registered_authorization_path"); ok {
		authSettings.RegisteredAuthorizationPath = String(v.(string))
	}
	if v, ok := d.GetOk("scope_for_oauth_grant_management"); ok {
		authSettings.ScopeForOAuthGrantManagement = String(v.(string))
	}
	if v, ok := d.GetOk("user_authorization_consent_adapter"); ok {
		authSettings.UserAuthorizationConsentAdapter = String(v.(string))
	}
	if v, ok := d.GetOk("user_authorization_consent_page_setting"); ok {
		authSettings.UserAuthorizationConsentPageSetting = String(v.(string))
	}
	if v, ok := d.GetOk("user_authorization_url"); ok {
		authSettings.UserAuthorizationUrl = String(v.(string))
	}
	return authSettings
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
	resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateOauthAuthServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthorizationServerSettings) diag.Diagnostics {
	var diags diag.Diagnostics

	setResourceDataStringWithDiagnostic(d, "default_scope_description", rv.DefaultScopeDescription, &diags)
	setResourceDataIntWithDiagnostic(d, "authorization_code_timeout", rv.AuthorizationCodeTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "authorization_code_entropy", rv.AuthorizationCodeEntropy, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_token_length", rv.RefreshTokenLength, &diags)
	setResourceDataIntWithDiagnostic(d, "refresh_rolling_interval", rv.RefreshRollingInterval, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_unidentified_client_extension_grants", rv.AllowUnidentifiedClientExtensionGrants, &diags)
	setResourceDataBoolWithDiagnostic(d, "track_user_sessions_for_logout", rv.TrackUserSessionsForLogout, &diags)
	setResourceDataStringWithDiagnostic(d, "token_endpoint_base_url", rv.TokenEndpointBaseUrl, &diags)
	if rv.PersistentGrantLifetime == nil {
		rv.PersistentGrantLifetime = Int(-1)
	}
	setResourceDataIntWithDiagnostic(d, "persistent_grant_lifetime", rv.PersistentGrantLifetime, &diags)
	if rv.PersistentGrantLifetimeUnit == nil {
		rv.PersistentGrantLifetimeUnit = String("DAYS")
	}
	setResourceDataStringWithDiagnostic(d, "persistent_grant_lifetime_unit", rv.PersistentGrantLifetimeUnit, &diags)
	setResourceDataBoolWithDiagnostic(d, "roll_refresh_token_values", rv.RollRefreshTokenValues, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_authorization_for_approved_grants", rv.BypassAuthorizationForApprovedGrants, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_unidentified_client_ro_creds", rv.AllowUnidentifiedClientROCreds, &diags)
	setResourceDataStringWithDiagnostic(d, "atm_id_for_oauth_grant_management", rv.AtmIdForOAuthGrantManagement, &diags)
	setResourceDataStringWithDiagnostic(d, "approved_scope_attribute", rv.ApprovedScopesAttribute, &diags) //TODO deprecated
	setResourceDataStringWithDiagnostic(d, "approved_scopes_attribute", rv.ApprovedScopesAttribute, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_activation_code_confirmation", rv.BypassActivationCodeConfirmation, &diags)
	setResourceDataIntWithDiagnostic(d, "device_polling_interval", rv.DevicePollingInterval, &diags)
	setResourceDataIntWithDiagnostic(d, "pending_authorization_timeout", rv.PendingAuthorizationTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "persistent_grant_idle_timeout", rv.PersistentGrantIdleTimeout, &diags)
	setResourceDataStringWithDiagnostic(d, "persistent_grant_idle_timeout_time_unit", rv.PersistentGrantIdleTimeoutTimeUnit, &diags)
	setResourceDataStringWithDiagnostic(d, "registered_authorization_path", rv.RegisteredAuthorizationPath, &diags)
	setResourceDataStringWithDiagnostic(d, "scope_for_oauth_grant_management", rv.ScopeForOAuthGrantManagement, &diags)
	setResourceDataStringWithDiagnostic(d, "user_authorization_consent_adapter", rv.UserAuthorizationConsentAdapter, &diags)
	setResourceDataStringWithDiagnostic(d, "user_authorization_consent_page_setting", rv.UserAuthorizationConsentPageSetting, &diags)
	setResourceDataStringWithDiagnostic(d, "user_authorization_url", rv.UserAuthorizationUrl, &diags)

	if rv.ParReferenceLength == nil {
		rv.ParReferenceLength = Int(24)
	}
	setResourceDataIntWithDiagnostic(d, "par_reference_length", rv.ParReferenceLength, &diags)
	if rv.ParReferenceTimeout == nil {
		rv.ParReferenceTimeout = Int(60)
	}
	setResourceDataIntWithDiagnostic(d, "par_reference_timeout", rv.ParReferenceTimeout, &diags)
	if rv.ParStatus == nil {
		rv.ParStatus = String("ENABLED")
	}
	setResourceDataStringWithDiagnostic(d, "par_status", rv.ParStatus, &diags)

	if rv.AdminWebServicePcvRef != nil {
		if err := d.Set("admin_web_service_pcv_ref", flattenResourceLink(rv.AdminWebServicePcvRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.PersistentGrantReuseGrantTypes != nil && len(*rv.PersistentGrantReuseGrantTypes) > 0 {
		if err := d.Set("persistent_grant_reuse_grant_types", rv.PersistentGrantReuseGrantTypes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AllowedOrigins != nil && len(*rv.AllowedOrigins) > 0 {
		if err := d.Set("allowed_origins", *rv.AllowedOrigins); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.PersistentGrantContract != nil && persistentGrantContractShouldFlatten(rv.PersistentGrantContract) {
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
