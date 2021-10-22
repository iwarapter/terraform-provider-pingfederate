package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"
	"regexp"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientSettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthClientSettingsResource() *schema.Resource {

	return &schema.Resource{
		Description: `Provides a OAuth Client Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.

!> This resource cannot be used together with ` + "`pingfederate_extended_properties` as both API's configure the same client metadata.",
		CreateContext: resourcePingFederateOauthClientSettingsResourceCreate,
		ReadContext:   resourcePingFederateOauthClientSettingsResourceRead,
		UpdateContext: resourcePingFederateOauthClientSettingsResourceUpdate,
		DeleteContext: resourcePingFederateOauthClientSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateOauthClientSettingsResourceImport,
		},
		Schema: map[string]*schema.Schema{
			"client_metadata": {
				Type:        schema.TypeSet,
				Elem:        resourceClientMetadata(),
				Description: "The client metadata.",
				Optional:    true,
			},
			"dynamic_client_registration": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        resourceDynamicClientRegistration(),
				Description: "Dynamic client registration settings.",
				Optional:    true,
			},
		},
	}
}

func resourcePingFederateOauthClientSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("OauthClientSettings")
	return resourcePingFederateOauthClientSettingsResourceUpdate(ctx, d, m)
}

func resourcePingFederateOauthClientSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClientSettings
	result, _, err := svc.GetClientSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read OauthClientSettings: %s", err)
	}
	return resourcePingFederateOauthClientSettingsResourceReadResult(d, result, m.(pfClient).PfVersion())
}

func resourcePingFederateOauthClientSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	authSettings := resourcePingFederateOauthClientSettingsResourceReadData(d, m.(pfClient).PfVersion())

	svc := m.(pfClient).OauthClientSettings
	input := &oauthClientSettings.UpdateClientSettingsInput{
		Body: *authSettings,
	}

	result, _, err := svc.UpdateClientSettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to update OauthClientSettings: %s", err)
	}
	return resourcePingFederateOauthClientSettingsResourceReadResult(d, result, m.(pfClient).PfVersion())
}

func resourcePingFederateOauthClientSettingsResourceReadData(d *schema.ResourceData, pfVersion string) *pf.ClientSettings {
	re := regexp.MustCompile(`^(10\.[1-9])`)
	isPF10_1 := re.MatchString(pfVersion)
	re = regexp.MustCompile(`^(10\.[3-9])`)
	isPF10_3 := re.MatchString(pfVersion)

	result := &pf.ClientSettings{}
	if val, ok := d.GetOk("client_metadata"); ok {
		result.ClientMetadata = expandClientMetadataList(val.(*schema.Set).List())
	}
	if _, ok := d.GetOk("dynamic_client_registration"); ok {
		result.DynamicClientRegistration = &pf.DynamicClientRegistration{}
		if val, ok := d.GetOk("dynamic_client_registration.0.require_proof_key_for_code_exchange"); ok {
			result.DynamicClientRegistration.RequireProofKeyForCodeExchange = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.token_exchange_processor_policy_ref"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.TokenExchangeProcessorPolicyRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.allowed_exclusive_scopes"); ok {
			strs := expandStringList(val.([]interface{}))
			result.DynamicClientRegistration.AllowedExclusiveScopes = &strs
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.enforce_replay_prevention"); ok {
			result.DynamicClientRegistration.EnforceReplayPrevention = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.restrict_to_default_access_token_manager"); ok {
			result.DynamicClientRegistration.RestrictToDefaultAccessTokenManager = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_idle_timeout"); ok {
			result.DynamicClientRegistration.PersistentGrantIdleTimeout = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.refresh_token_rolling_interval"); ok {
			result.DynamicClientRegistration.RefreshTokenRollingInterval = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.user_authorization_url_override"); ok {
			result.DynamicClientRegistration.UserAuthorizationUrlOverride = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.initial_access_token_scope"); ok {
			result.DynamicClientRegistration.InitialAccessTokenScope = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.restricted_common_scopes"); ok {
			strs := expandStringList(val.([]interface{}))
			result.DynamicClientRegistration.RestrictedCommonScopes = &strs
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.client_cert_issuer_ref"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.ClientCertIssuerRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.refresh_rolling"); ok {
			result.DynamicClientRegistration.RefreshRolling = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.request_policy_ref"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.RequestPolicyRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_idle_timeout_time_unit"); ok {
			result.DynamicClientRegistration.PersistentGrantIdleTimeoutTimeUnit = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.oidc_policy"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.OidcPolicy = expandClientRegistrationOIDCPolicy(val.([]interface{})[0].(map[string]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_idle_timeout_type"); ok {
			result.DynamicClientRegistration.PersistentGrantIdleTimeoutType = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.device_polling_interval_override"); ok {
			result.DynamicClientRegistration.DevicePollingIntervalOverride = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.ciba_polling_interval"); ok {
			result.DynamicClientRegistration.CibaPollingInterval = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.device_flow_setting_type"); ok {
			result.DynamicClientRegistration.DeviceFlowSettingType = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.require_signed_requests"); ok {
			result.DynamicClientRegistration.RequireSignedRequests = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.default_access_token_manager_ref"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.DefaultAccessTokenManagerRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.client_cert_issuer_type"); ok {
			result.DynamicClientRegistration.ClientCertIssuerType = String(val.(string))
		}
		if isPF10_3 {
			if val, ok := d.GetOk("dynamic_client_registration.0.refresh_token_rolling_interval_type"); ok {
				result.DynamicClientRegistration.RefreshTokenRollingIntervalType = String(val.(string))
			}
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.restrict_common_scopes"); ok {
			result.DynamicClientRegistration.RestrictCommonScopes = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_expiration_type"); ok {
			result.DynamicClientRegistration.PersistentGrantExpirationType = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_expiration_time_unit"); ok {
			result.DynamicClientRegistration.PersistentGrantExpirationTimeUnit = String(val.(string))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.disable_registration_access_tokens"); ok {
			result.DynamicClientRegistration.DisableRegistrationAccessTokens = Bool(val.(bool))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.ciba_require_signed_requests"); ok {
			result.DynamicClientRegistration.CibaRequireSignedRequests = Bool(val.(bool))
		}
		if isPF10_1 {
			if val, ok := d.GetOk("dynamic_client_registration.0.allow_client_delete"); ok {
				result.DynamicClientRegistration.AllowClientDelete = Bool(val.(bool))
			}
			if val, ok := d.GetOk("dynamic_client_registration.0.rotate_client_secret"); ok {
				result.DynamicClientRegistration.RotateClientSecret = Bool(val.(bool))
			}
			if val, ok := d.GetOk("dynamic_client_registration.0.rotate_registration_access_token"); ok {
				result.DynamicClientRegistration.RotateRegistrationAccessToken = Bool(val.(bool))
			}
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.persistent_grant_expiration_time"); ok {
			result.DynamicClientRegistration.PersistentGrantExpirationTime = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.policy_refs"); ok && len(val.([]interface{})) > 0 {
			result.DynamicClientRegistration.PolicyRefs = expandResourceLinkList(val.([]interface{}))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.pending_authorization_timeout_override"); ok {
			result.DynamicClientRegistration.PendingAuthorizationTimeoutOverride = Int(val.(int))
		}
		if val, ok := d.GetOk("dynamic_client_registration.0.bypass_activation_code_confirmation_override"); ok {
			result.DynamicClientRegistration.BypassActivationCodeConfirmationOverride = Bool(val.(bool))
		}
	}
	return result
}

func resourcePingFederateOauthClientSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	//resource cannot be deleted just not tracked by terraform anymore
	return nil
}

func resourcePingFederateOauthClientSettingsResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).OauthClientSettings
	result, _, err := svc.GetClientSettingsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	resourcePingFederateOauthClientSettingsResourceReadResult(d, result, m.(pfClient).PfVersion())
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateOauthClientSettingsResourceReadResult(d *schema.ResourceData, rv *pf.ClientSettings, pfVersion string) diag.Diagnostics {
	var diags diag.Diagnostics

	if rv.ClientMetadata != nil && len(*rv.ClientMetadata) > 0 {
		if err := d.Set("client_metadata", flattenClientMetadatas(rv.ClientMetadata)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.DynamicClientRegistration != nil {
		if err := d.Set("dynamic_client_registration", flattenDynamicClientRegistration(rv.DynamicClientRegistration, pfVersion)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}
