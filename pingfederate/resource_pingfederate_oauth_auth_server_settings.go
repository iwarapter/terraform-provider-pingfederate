package pingfederate

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
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
				"scopes": scopes,
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
					ValidateFunc: validatePersistentGrantReuseGrantTypes,
				},
			},
			"persistent_grant_contract": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extended_attributes": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
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
			"admin_web_service_pcv_ref": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourcePingFederateOauthAuthServerSettingsResourceCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("OauthAuthServerSettings")
	return resourcePingFederateOauthAuthServerSettingsResourceUpdate(d, m)
}

func resourcePingFederateOauthAuthServerSettingsResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAuthServerSettings
	result, _, _ := svc.GetAuthorizationServerSettings()
	//TODO handle the resp/error
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceUpdate(d *schema.ResourceData, m interface{}) error {
	defaultScopeDescription := d.Get("default_scope_description").(string)
	authorizationCodeTimeout := d.Get("authorization_code_timeout").(int)
	authorizationCodeEntropy := d.Get("authorization_code_entropy").(int)
	refreshTokenLength := d.Get("refresh_token_length").(int)
	refreshRollingInterval := d.Get("refresh_rolling_interval").(int)

	authSettings := &pf.AuthorizationServerSettings{
		DefaultScopeDescription:  String(defaultScopeDescription),
		AuthorizationCodeTimeout: Int(authorizationCodeTimeout),
		AuthorizationCodeEntropy: Int(authorizationCodeEntropy),
		RefreshTokenLength:       Int(refreshTokenLength),
		RefreshRollingInterval:   Int(refreshRollingInterval),
	}

	if _, ok := d.GetOk("scopes"); ok {
		authSettings.Scopes = expandScopes(d.Get("scopes").(*schema.Set).List())
	}

	svc := m.(*pf.PfClient).OauthAuthServerSettings
	input := &pf.UpdateAuthorizationServerSettingsInput{
		Body: *authSettings,
	}

	result, _, _ := svc.UpdateAuthorizationServerSettings(input)
	//TODO handle the resp/error
	return resourcePingFederateOauthAuthServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthServerSettingsResourceDelete(d *schema.ResourceData, m interface{}) error {
	//resource cannot be deleted just not tracked by terraform anymore
	return nil
}

func resourcePingFederateOauthAuthServerSettingsResourceImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	// idParts := strings.SplitN(d.Id(), "/", 2)
	// if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
	// 	return nil, fmt.Errorf("unexpected format of ID (%q), expected <application_id>/<resource_id>", d.Id())
	// }
	// application_id := idParts[0]
	// resource_id := idParts[1]
	// d.Set("application_id", application_id)
	// d.SetId(resource_id)

	// svc := m.(*pingfederate.PfClient).OauthAuthServerSettings
	// result, resp, err := svc.GetAuthorizationServerSettings()
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateOauthAuthServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthorizationServerSettings) (err error) {
	log.Printf("[INFO] ")
	setResourceDataString(d, "default_scope_description", rv.DefaultScopeDescription)
	setResourceDataInt(d, "authorization_code_timeout", rv.AuthorizationCodeTimeout)
	setResourceDataInt(d, "authorization_code_entropy", rv.AuthorizationCodeEntropy)
	setResourceDataInt(d, "refresh_token_length", rv.RefreshTokenLength)
	setResourceDataInt(d, "refresh_rolling_interval", rv.RefreshRollingInterval)

	if rv.Scopes != nil && len(*rv.Scopes) > 0 {
		if err = d.Set("scopes", flattenScopes(*rv.Scopes)); err != nil {
			return err
		}
	}
	return nil
}
