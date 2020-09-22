package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthClientResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateOauthClientResourceCreate,
		ReadContext:   resourcePingFederateOauthClientResourceRead,
		UpdateContext: resourcePingFederateOauthClientResourceUpdate,
		DeleteContext: resourcePingFederateOauthClientResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateOauthClientResourceSchema(),
	}
}

func resourcePingFederateOauthClientResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"grant_types": {
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validateGrantTypes,
			},
		},
		"bypass_approval_page": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"exclusive_scopes": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"logo_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"persistent_grant_expiration_time": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"persistent_grant_expiration_time_unit": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DAYS",
		},
		"persistent_grant_expiration_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SERVER_DEFAULT",
		},
		"redirect_uris": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"refresh_rolling": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SERVER_DEFAULT",
		},
		"require_signed_requests": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"restrict_scopes": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"restricted_response_types": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"restricted_scopes": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"validate_using_all_eligible_atms": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"client_auth": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"client_cert_issuer_dn": {
						Type:          schema.TypeString,
						Optional:      true,
						ConflictsWith: []string{"client_auth.0.secret"},
					},
					"client_cert_subject_dn": {
						Type:          schema.TypeString,
						Optional:      true,
						ConflictsWith: []string{"client_auth.0.secret"},
					},
					"enforce_replay_prevention": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					//TODO do we enable Secret/EncryptedSecret??
					"secret": {
						Type:          schema.TypeString,
						Optional:      true,
						Sensitive:     true,
						ConflictsWith: []string{"client_auth.0.client_cert_issuer_dn", "client_auth.0.client_cert_subject_dn"},
					},
					"type": {
						Type:             schema.TypeString,
						Required:         true,
						ValidateDiagFunc: validateClientAuthType,
					},
					"token_endpoint_auth_signing_algorithm": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"default_access_token_manager_ref": resourceLinkSchema(),
		"oidc_policy": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"grant_access_session_revocation_api": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"pairwise_identifier_user_type": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"id_token_signing_algorithm": {
						Type:             schema.TypeString,
						Optional:         true,
						Default:          "RS256",
						ValidateDiagFunc: validateTokenSigningAlgorithm,
					},
					"id_token_encryption_algorithm": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"id_token_content_encryption_algorithm": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"sector_identifier_uri": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"logout_uris": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"ping_access_logout_capable": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"policy_group": resourceLinkSchema(),
				},
			},
		},
		"jwks_settings": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"jwks": {
						Type:          schema.TypeString,
						Optional:      true,
						ConflictsWith: []string{"jwks_settings.0.jwks_url"},
					},
					"jwks_url": {
						Type:          schema.TypeString,
						Optional:      true,
						ConflictsWith: []string{"jwks_settings.0.jwks"},
					},
				},
			},
		},
		"ciba_delivery_mode": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ciba_notification_endpoint": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ciba_polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"ciba_request_object_signing_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ciba_require_signed_requests": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"ciba_user_code_supported": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"bypass_activation_code_confirmation_override": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"device_flow_setting_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SERVER_DEFAULT",
		},
		"device_polling_interval_override": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"extended_properties": {
			Type:     schema.TypeSet,
			Elem:     resourceParameterValues(),
			Optional: true,
		},
		"pending_authorization_timeout_override": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"persistent_grant_idle_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"persistent_grant_idle_timeout_time_unit": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DAYS",
		},
		"persistent_grant_idle_timeout_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SERVER_DEFAULT",
		},
		"request_object_signing_algorithm": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"request_policy_ref": resourceLinkSchema(),
		"require_proof_key_for_code_exchange": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"token_exchange_processor_policy_ref": resourceLinkSchema(),
		"user_authorization_url_override": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourcePingFederateOauthClientResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.CreateClientInput{
		Body: *resourcePingFederateOauthClientResourceReadData(d),
	}
	result, _, err := svc.CreateClient(&input)
	if err != nil {
		return diag.Errorf("unable to create OauthClients: %s", err)
	}
	d.SetId(*result.ClientId)
	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.GetClientInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetClient(&input)
	if err != nil {
		return diag.Errorf("unable to read OauthClients: %s", err)
	}
	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.UpdateClientInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthClientResourceReadData(d),
	}
	result, _, err := svc.UpdateClient(&input)
	if err != nil {
		return diag.Errorf("unable to update OauthClients: %s", err)
	}

	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.DeleteClientInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteClient(&input)
	if err != nil {
		return diag.Errorf("unable to delete OauthClients: %s", err)
	}
	return nil
}

func resourcePingFederateOauthClientResourceReadResult(d *schema.ResourceData, rv *pf.Client) diag.Diagnostics {
	//TODO
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "client_id", rv.ClientId, &diags)
	setResourceDataBoolWithDiagnostic(d, "enabled", rv.Enabled, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_approval_page", rv.BypassApprovalPage, &diags)
	setResourceDataStringWithDiagnostic(d, "description", rv.Description, &diags)
	setResourceDataStringWithDiagnostic(d, "logo_url", rv.LogoUrl, &diags)
	setResourceDataIntWithDiagnostic(d, "persistent_grant_expiration_time", rv.PersistentGrantExpirationTime, &diags)
	setResourceDataStringWithDiagnostic(d, "persistent_grant_expiration_time_unit", rv.PersistentGrantExpirationTimeUnit, &diags)
	setResourceDataStringWithDiagnostic(d, "persistent_grant_expiration_type", rv.PersistentGrantExpirationType, &diags)
	setResourceDataStringWithDiagnostic(d, "refresh_rolling", rv.RefreshRolling, &diags)
	setResourceDataBoolWithDiagnostic(d, "require_signed_requests", rv.RequireSignedRequests, &diags)
	setResourceDataBoolWithDiagnostic(d, "restrict_scopes", rv.RestrictScopes, &diags)
	setResourceDataBoolWithDiagnostic(d, "validate_using_all_eligible_atms", rv.ValidateUsingAllEligibleAtms, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_activation_code_confirmation_override", rv.BypassActivationCodeConfirmationOverride, &diags)
	setResourceDataStringWithDiagnostic(d, "ciba_delivery_mode", rv.CibaDeliveryMode, &diags)
	setResourceDataStringWithDiagnostic(d, "ciba_notification_endpoint", rv.CibaNotificationEndpoint, &diags)
	setResourceDataIntWithDiagnostic(d, "ciba_polling_interval", rv.CibaPollingInterval, &diags)
	setResourceDataStringWithDiagnostic(d, "ciba_request_object_signing_algorithm", rv.CibaRequestObjectSigningAlgorithm, &diags)
	setResourceDataBoolWithDiagnostic(d, "ciba_require_signed_requests", rv.CibaRequireSignedRequests, &diags)
	setResourceDataBoolWithDiagnostic(d, "ciba_user_code_supported", rv.CibaUserCodeSupported, &diags)

	setResourceDataStringWithDiagnostic(d, "device_flow_setting_type", rv.DeviceFlowSettingType, &diags)
	setResourceDataIntWithDiagnostic(d, "device_polling_interval_override", rv.DevicePollingIntervalOverride, &diags)
	setResourceDataIntWithDiagnostic(d, "pending_authorization_timeout_override", rv.PendingAuthorizationTimeoutOverride, &diags)
	setResourceDataIntWithDiagnostic(d, "persistent_grant_idle_timeout", rv.PersistentGrantIdleTimeout, &diags)
	setResourceDataStringWithDiagnostic(d, "persistent_grant_idle_timeout_time_unit", rv.PersistentGrantIdleTimeoutTimeUnit, &diags)
	setResourceDataStringWithDiagnostic(d, "persistent_grant_idle_timeout_type", rv.PersistentGrantIdleTimeoutType, &diags)
	setResourceDataStringWithDiagnostic(d, "request_object_signing_algorithm", rv.RequestObjectSigningAlgorithm, &diags)
	setResourceDataBoolWithDiagnostic(d, "require_proof_key_for_code_exchange", rv.RequireProofKeyForCodeExchange, &diags)
	setResourceDataStringWithDiagnostic(d, "user_authorization_url_override", rv.UserAuthorizationUrlOverride, &diags)

	if rv.RequestPolicyRef != nil {
		if err := d.Set("request_policy_ref", flattenResourceLink(rv.RequestPolicyRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.TokenExchangeProcessorPolicyRef != nil {
		if err := d.Set("token_exchange_processor_policy_ref", flattenResourceLink(rv.TokenExchangeProcessorPolicyRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.ExtendedParameters != nil {
		if err := d.Set("extended_properties", flattenMapOfParameterValues(rv.ExtendedParameters)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.GrantTypes != nil && len(*rv.GrantTypes) > 0 {
		if err := d.Set("grant_types", *rv.GrantTypes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ExclusiveScopes != nil && len(*rv.ExclusiveScopes) > 0 {
		if err := d.Set("exclusive_scopes", *rv.ExclusiveScopes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RedirectUris != nil && len(*rv.RedirectUris) > 0 {
		if err := d.Set("redirect_uris", *rv.RedirectUris); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RestrictedResponseTypes != nil && len(*rv.RestrictedResponseTypes) > 0 {
		if err := d.Set("restricted_response_types", *rv.RestrictedResponseTypes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RestrictedScopes != nil && len(*rv.RestrictedScopes) > 0 {
		if err := d.Set("restricted_scopes", *rv.RestrictedScopes); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ClientAuth != nil && *rv.ClientAuth.Type != "NONE" {
		orig := expandClientAuth(d.Get("client_auth").([]interface{}))

		if err := d.Set("client_auth", flattenClientAuth(orig, rv.ClientAuth)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.JwksSettings != nil {
		if err := d.Set("jwks_settings", flattenJwksSettings(rv.JwksSettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.DefaultAccessTokenManagerRef != nil {
		if err := d.Set("default_access_token_manager_ref", flattenResourceLink(rv.DefaultAccessTokenManagerRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.OidcPolicy != nil {
		if err := d.Set("oidc_policy", flattenClientOIDCPolicy(rv.OidcPolicy)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateOauthClientResourceReadData(d *schema.ResourceData) *pf.Client {
	grants := expandStringList(d.Get("grant_types").(*schema.Set).List())
	client := &pf.Client{
		Name:       String(d.Get("name").(string)),
		ClientId:   String(d.Get("client_id").(string)),
		Enabled:    Bool(d.Get("enabled").(bool)),
		GrantTypes: &grants,
	}

	if v, ok := d.GetOk("bypass_approval_page"); ok {
		client.BypassApprovalPage = Bool(v.(bool))
	}

	if v, ok := d.GetOk("description"); ok {
		client.Description = String(v.(string))
	}

	if v, ok := d.GetOk("exclusive_scopes"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		client.ExclusiveScopes = &strs
	}

	if v, ok := d.GetOk("logo_url"); ok {
		client.LogoUrl = String(v.(string))
	}

	if v, ok := d.GetOk("persistent_grant_expiration_time"); ok {
		client.PersistentGrantExpirationTime = Int(v.(int))
	}

	if v, ok := d.GetOk("persistent_grant_expiration_time_unit"); ok {
		client.PersistentGrantExpirationTimeUnit = String(v.(string))
	}

	if v, ok := d.GetOk("persistent_grant_expiration_type"); ok {
		client.PersistentGrantExpirationType = String(v.(string))
	}

	if v, ok := d.GetOk("redirect_uris"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		client.RedirectUris = &strs
	}

	if v, ok := d.GetOk("refresh_rolling"); ok {
		client.RefreshRolling = String(v.(string))
	}

	if v, ok := d.GetOk("require_signed_requests"); ok {
		client.RequireSignedRequests = Bool(v.(bool))
	}

	if v, ok := d.GetOk("restrict_scopes"); ok {
		client.RestrictScopes = Bool(v.(bool))
	}

	if v, ok := d.GetOk("restricted_response_types"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		client.RestrictedResponseTypes = &strs
	}

	if v, ok := d.GetOk("restricted_scopes"); ok {
		strs := expandStringList(v.(*schema.Set).List())
		client.RestrictedScopes = &strs
	}

	if v, ok := d.GetOk("validate_using_all_eligible_atms"); ok {
		client.ValidateUsingAllEligibleAtms = Bool(v.(bool))
	}
	if v, ok := d.GetOk("client_auth"); ok && len(v.([]interface{})) > 0 {
		client.ClientAuth = expandClientAuth(v.([]interface{}))
	}
	if v, ok := d.GetOk("jwks_settings"); ok && len(v.([]interface{})) > 0 {
		client.JwksSettings = expandJwksSettings(v.([]interface{}))
	}
	if v, ok := d.GetOk("default_access_token_manager_ref"); ok && len(v.([]interface{})) > 0 {
		client.DefaultAccessTokenManagerRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("oidc_policy"); ok && len(v.([]interface{})) > 0 {
		client.OidcPolicy = expandClientOIDCPolicy(v.([]interface{}))
	}
	if v, ok := d.GetOk("request_policy_ref"); ok && len(v.([]interface{})) > 0 {
		client.RequestPolicyRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("ciba_delivery_mode"); ok {
		client.CibaDeliveryMode = String(v.(string))
	}
	if v, ok := d.GetOk("ciba_notification_endpoint"); ok {
		client.CibaNotificationEndpoint = String(v.(string))
	}
	if v, ok := d.GetOk("ciba_polling_interval"); ok {
		client.CibaPollingInterval = Int(v.(int))
	}
	if v, ok := d.GetOk("ciba_request_object_signing_algorithm"); ok {
		client.CibaRequestObjectSigningAlgorithm = String(v.(string))
	}
	if v, ok := d.GetOk("ciba_require_signed_requests"); ok {
		client.CibaRequireSignedRequests = Bool(v.(bool))
	}
	if v, ok := d.GetOk("ciba_user_code_supported"); ok {
		client.CibaUserCodeSupported = Bool(v.(bool))
	}
	if v, ok := d.GetOk("bypass_activation_code_confirmation_override"); ok {
		client.BypassActivationCodeConfirmationOverride = Bool(v.(bool))
	}
	if v, ok := d.GetOk("device_flow_setting_type"); ok {
		client.DeviceFlowSettingType = String(v.(string))
	}
	if v, ok := d.GetOk("device_polling_interval_override"); ok {
		client.DevicePollingIntervalOverride = Int(v.(int))
	}
	if v, ok := d.GetOk("extended_properties"); ok {
		client.ExtendedParameters = expandMapOfParameterValues(v.(*schema.Set).List())
	}
	if v, ok := d.GetOk("pending_authorization_timeout_override"); ok {
		client.PendingAuthorizationTimeoutOverride = Int(v.(int))
	}
	if v, ok := d.GetOk("persistent_grant_idle_timeout"); ok {
		client.PersistentGrantIdleTimeout = Int(v.(int))
	}
	if v, ok := d.GetOk("persistent_grant_idle_timeout_time_unit"); ok {
		client.PersistentGrantIdleTimeoutTimeUnit = String(v.(string))
	}
	if v, ok := d.GetOk("persistent_grant_idle_timeout_type"); ok {
		client.PersistentGrantIdleTimeoutType = String(v.(string))
	}
	if v, ok := d.GetOk("request_object_signing_algorithm"); ok {
		client.RequestObjectSigningAlgorithm = String(v.(string))
	}
	if v, ok := d.GetOk("require_proof_key_for_code_exchange"); ok {
		client.RequireProofKeyForCodeExchange = Bool(v.(bool))
	}
	if v, ok := d.GetOk("token_exchange_processor_policy_ref"); ok && len(v.([]interface{})) > 0 {
		client.TokenExchangeProcessorPolicyRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("user_authorization_url_override"); ok {
		client.UserAuthorizationUrlOverride = String(v.(string))
	}
	return client
}
