package pingfederate

import (
	"fmt"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"
	"log"

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
		Create: resourcePingFederateOauthAuthServerSettingsResourceCreate,
		Read:   resourcePingFederateOauthAuthServerSettingsResourceRead,
		Update: resourcePingFederateOauthAuthServerSettingsResourceUpdate,
		Delete: resourcePingFederateOauthAuthServerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			State: resourcePingFederateOauthAuthServerSettingsResourceImport,
		},

		Schema: map[string]*schema.Schema{
			"default_scope_description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"scopes":                 scopes,
			"scope_groups":           scopeGroups,
			"exclusive_scopes":       scopes,
			"exclusive_scope_groups": scopeGroups,
			"authorization_code_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"authorization_code_entropy": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"track_user_sessions_for_logout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"token_endpoint_base_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistent_grant_lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"persistent_grant_lifetime_unit": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validatePersistentGrantLifetimeUnit,
			},
			"refresh_token_length": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"roll_refresh_token_values": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"refresh_rolling_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"persistent_grant_reuse_grant_types": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateGrantTypes,
				},
			},
			"persistent_grant_contract": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extended_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"bypass_authorization_for_approved_grants": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allow_unidentified_client_ro_creds": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"allowed_origins": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allow_unidentified_client_extension_grants": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"admin_web_service_pcv_ref": resourceLinkSchema(),
		},
	}
}

func resourcePingFederateOauthAuthServerSettingsResourceCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("OauthAuthServerSettings")
	return resourcePingFederateOauthAuthServerSettingsResourceUpdate(d, m)
}

func resourcePingFederateOauthAuthServerSettingsResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(pfClient).OauthAuthServerSettings
	result, _, err := svc.GetAuthorizationServerSettings()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceUpdate(d *schema.ResourceData, m interface{}) error {
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

	if _, ok := d.GetOk("allowed_origins"); ok {
		strs := expandStringList(d.Get("allowed_origins").(*schema.Set).List())
		authSettings.AllowedOrigins = &strs
	}

	if _, ok := d.GetOk("allow_unidentified_client_extension_grants"); ok {
		authSettings.AllowUnidentifiedClientExtensionGrants = Bool(d.Get("allow_unidentified_client_extension_grants").(bool))
	}

	if _, ok := d.GetOk("allow_unidentified_client_ro_creds"); ok {
		authSettings.AllowUnidentifiedClientROCreds = Bool(d.Get("allow_unidentified_client_ro_creds").(bool))
	}

	if _, ok := d.GetOk("bypass_authorization_for_approved_grants"); ok {
		authSettings.BypassAuthorizationForApprovedGrants = Bool(d.Get("bypass_authorization_for_approved_grants").(bool))
	}

	if _, ok := d.GetOk("exclusive_scope_groups"); ok {
		authSettings.ExclusiveScopeGroups = expandScopeGroups(d.Get("exclusive_scope_groups").(*schema.Set).List())
	}

	if _, ok := d.GetOk("exclusive_scopes"); ok {
		authSettings.ExclusiveScopes = expandScopes(d.Get("exclusive_scopes").(*schema.Set).List())
	}

	if _, ok := d.GetOk("persistent_grant_contract"); ok {
		authSettings.PersistentGrantContract = expandPersistentGrantContract(d.Get("persistent_grant_contract").(*schema.Set).List())
	}

	if _, ok := d.GetOk("persistent_grant_lifetime"); ok {
		authSettings.PersistentGrantLifetime = Int(d.Get("persistent_grant_lifetime").(int))
	}

	if _, ok := d.GetOk("persistent_grant_lifetime_unit"); ok {
		authSettings.PersistentGrantLifetimeUnit = String(d.Get("persistent_grant_lifetime_unit").(string))
	}

	if _, ok := d.GetOk("persistent_grant_reuse_grant_types"); ok {
		authSettings.PersistentGrantReuseGrantTypes = expandStringList(d.Get("persistent_grant_reuse_grant_types").(*schema.Set).List())
	}

	if _, ok := d.GetOk("scopes"); ok {
		authSettings.Scopes = expandScopes(d.Get("scopes").(*schema.Set).List())
	}

	if _, ok := d.GetOk("scope_groups"); ok {
		authSettings.ScopeGroups = expandScopeGroups(d.Get("scope_groups").(*schema.Set).List())
	}

	if _, ok := d.GetOk("roll_refresh_token_values"); ok {
		authSettings.RollRefreshTokenValues = Bool(d.Get("roll_refresh_token_values").(bool))
	}

	if _, ok := d.GetOk("token_endpoint_base_url"); ok {
		authSettings.TokenEndpointBaseUrl = String(d.Get("token_endpoint_base_url").(string))
	}

	if _, ok := d.GetOk("track_user_sessions_for_logout"); ok {
		authSettings.TrackUserSessionsForLogout = Bool(d.Get("track_user_sessions_for_logout").(bool))
	}

	svc := m.(pfClient).OauthAuthServerSettings
	input := &oauthAuthServerSettings.UpdateAuthorizationServerSettingsInput{
		Body: *authSettings,
	}

	result, _, err := svc.UpdateAuthorizationServerSettings(input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceDelete(d *schema.ResourceData, m interface{}) error {
	//resource cannot be deleted just not tracked by terraform anymore
	return nil
}

func resourcePingFederateOauthAuthServerSettingsResourceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).OauthAuthServerSettings
	result, _, err := svc.GetAuthorizationServerSettings()
	if err != nil {
		return []*schema.ResourceData{d}, fmt.Errorf(err.Error())
	}
	resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateOauthAuthServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthorizationServerSettings) (err error) {
	log.Printf("[INFO] ")
	setResourceDataString(d, "default_scope_description", rv.DefaultScopeDescription)
	setResourceDataInt(d, "authorization_code_timeout", rv.AuthorizationCodeTimeout)
	setResourceDataInt(d, "authorization_code_entropy", rv.AuthorizationCodeEntropy)
	setResourceDataInt(d, "refresh_token_length", rv.RefreshTokenLength)
	setResourceDataInt(d, "refresh_rolling_interval", rv.RefreshRollingInterval)
	setResourceDataBool(d, "allow_unidentified_client_extension_grants", rv.AllowUnidentifiedClientExtensionGrants)
	setResourceDataBool(d, "track_user_sessions_for_logout", rv.TrackUserSessionsForLogout)
	setResourceDataString(d, "token_endpoint_base_url", rv.TokenEndpointBaseUrl)
	setResourceDataInt(d, "persistent_grant_lifetime", rv.PersistentGrantLifetime)
	setResourceDataString(d, "persistent_grant_lifetime_unit", rv.PersistentGrantLifetimeUnit)
	setResourceDataBool(d, "roll_refresh_token_values", rv.RollRefreshTokenValues)
	setResourceDataBool(d, "bypass_authorization_for_approved_grants", rv.BypassAuthorizationForApprovedGrants)
	setResourceDataBool(d, "allow_unidentified_client_ro_creds", rv.AllowUnidentifiedClientROCreds)

	// "admin_web_service_pcv_ref"

	if rv.PersistentGrantReuseGrantTypes != nil && len(rv.PersistentGrantReuseGrantTypes) > 0 {
		if err = d.Set("persistent_grant_reuse_grant_types", rv.PersistentGrantReuseGrantTypes); err != nil {
			return err
		}
	}

	if rv.AllowedOrigins != nil && len(*rv.AllowedOrigins) > 0 {
		if err = d.Set("allowed_origins", *rv.AllowedOrigins); err != nil {
			return err
		}
	}

	if rv.PersistentGrantContract != nil {
		if err = d.Set("persistent_grant_contract", flattenPersistentGrantContract(rv.PersistentGrantContract)); err != nil {
			return err
		}
	}
	if rv.Scopes != nil && len(*rv.Scopes) > 0 {
		if err = d.Set("scopes", flattenScopes(*rv.Scopes)); err != nil {
			return err
		}
	}
	if rv.ScopeGroups != nil && len(*rv.ScopeGroups) > 0 {
		if err = d.Set("scope_groups", flattenScopeGroups(*rv.ScopeGroups)); err != nil {
			return err
		}
	}
	if rv.ExclusiveScopes != nil && len(*rv.ExclusiveScopes) > 0 {
		if err = d.Set("exclusive_scopes", flattenScopes(*rv.ExclusiveScopes)); err != nil {
			return err
		}
	}
	if rv.ExclusiveScopeGroups != nil && len(*rv.ExclusiveScopeGroups) > 0 {
		if err = d.Set("exclusive_scope_groups", flattenScopeGroups(*rv.ExclusiveScopeGroups)); err != nil {
			return err
		}
	}

	return nil
}
