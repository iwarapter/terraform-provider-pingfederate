package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthClientResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for OAuth Clients within PingFederate.",
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
			Type:        schema.TypeString,
			Required:    true,
			Description: "A descriptive name for the client instance. This name appears when the user is prompted for authorization.",
		},
		"client_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "A unique identifier the client provides to the Resource Server to identify itself. This identifier is included with every request the client makes.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Specifies whether the client is enabled. The default value is true.",
		},
		"grant_types": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "The grant types allowed for this client. The EXTENSION grant type applies to SAML/JWT assertion grants.",
			Elem: &schema.Schema{
				Type:             schema.TypeString,
				ValidateDiagFunc: validateGrantTypes,
			},
		},
		"bypass_approval_page": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use this setting, for example, when you want to deploy a trusted application and authenticate end users via an IdP adapter or IdP connection.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A description of what the client application does. This description appears when the user is prompted for authorization.",
		},
		"exclusive_scopes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The exclusive scopes available for this client.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"logo_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The location of the logo used on user-facing OAuth grant authorization and revocation pages.",
		},
		"persistent_grant_expiration_time": {
			Type:        schema.TypeInt,
			Description: "The persistent grant expiration time. -1 indicates an indefinite amount of time.",
			Optional:    true,
		},
		"persistent_grant_expiration_time_unit": {
			Type:        schema.TypeString,
			Description: "The persistent grant expiration time unit.",
			Optional:    true,
			Default:     "DAYS",
		},
		"persistent_grant_expiration_type": {
			Type:        schema.TypeString,
			Description: "Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
			Optional:    true,
			Default:     "SERVER_DEFAULT",
		},
		"redirect_uris": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "URIs to which the OAuth AS may redirect the resource owner's user agent after authorization is obtained. A redirection URI is used with the Authorization Code and Implicit grant types. Wildcards are allowed. However, for security reasons, make the URL as restrictive as possible.For example: https://*.company.com/* Important: If more than one URI is added or if a single URI uses wildcards, then Authorization Code grant and token requests must contain a specific matching redirect uri parameter.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"refresh_rolling": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT.",
			Default:     "SERVER_DEFAULT",
		},
		"require_signed_requests": {
			Type:        schema.TypeBool,
			Description: "Determines whether signed requests are required for this client",
			Optional:    true,
		},
		"restrict_scopes": {
			Type:        schema.TypeBool,
			Description: "Restricts this client's access to specific scopes.",
			Optional:    true,
		},
		"restricted_response_types": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The response types allowed for this client. If omitted all response types are available to the client.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"restricted_scopes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The scopes available for this client.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"validate_using_all_eligible_atms": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Validates token using all eligible access token managers for the client. This setting is ignored if 'restrict_to_default_access_token_manager' is set to true.",
		},
		"client_auth": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			MinItems:    1,
			Description: "Client authentication settings.  If this model is null, it indicates that no client authentication will be used.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"client_cert_issuer_dn": {
						Type:          schema.TypeString,
						Optional:      true,
						Description:   "Client TLS Certificate Issuer DN.",
						ConflictsWith: []string{"client_auth.0.secret"},
					},
					"client_cert_subject_dn": {
						Type:          schema.TypeString,
						Optional:      true,
						Description:   "Client TLS Certificate Subject DN.",
						ConflictsWith: []string{"client_auth.0.secret"},
					},
					"enforce_replay_prevention": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication.",
					},
					//TODO do we enable Secret/EncryptedSecret??
					"secret": {
						Type:          schema.TypeString,
						Optional:      true,
						Sensitive:     true,
						Description:   "Client secret for Basic Authentication. To update the client secret, specify the plaintext value in this field. This field will not be populated for GET requests.",
						ConflictsWith: []string{"client_auth.0.client_cert_issuer_dn", "client_auth.0.client_cert_subject_dn"},
					},
					"type": {
						Type:             schema.TypeString,
						Required:         true,
						Description:      "Client authentication type.",
						ValidateDiagFunc: validateClientAuthType,
					},
					"token_endpoint_auth_signing_algorithm": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present.",
						ValidateFunc: validation.StringInSlice([]string{
							"RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512",
						}, false),
					},
				},
			},
		},
		"default_access_token_manager_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The default access token manager for this client.",
			Elem:        resourceLinkResource(),
		},
		"oidc_policy": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Open ID Connect Policy settings.  This is included in the message only when OIDC is enabled.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"grant_access_session_revocation_api": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Determines whether this client is allowed to access the Session Revocation API.",
					},
					"pairwise_identifier_user_type": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Determines whether the subject identifier type is pairwise.",
					},
					"id_token_signing_algorithm": {
						Type:             schema.TypeString,
						Optional:         true,
						Default:          "RS256",
						Description:      "The JSON Web Signature [JWS] algorithm required for the ID Token.",
						ValidateDiagFunc: validateTokenSigningAlgorithm,
					},
					"id_token_encryption_algorithm": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.",
					},
					"id_token_content_encryption_algorithm": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.",
					},
					"sector_identifier_uri": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The URI references a file with a single JSON array of Redirect URI and JWKS URL values.",
					},
					"logout_uris": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "A list of client logout URI's which will be invoked when a user logs out through one of PingFederate's SLO endpoints.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"ping_access_logout_capable": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention.",
					},
					"policy_group": {
						Type:        schema.TypeList,
						Optional:    true,
						MaxItems:    1,
						Description: "The Open ID Connect policy. A null value will represent the default policy group.",
						Elem:        resourceLinkResource(),
					},
				},
			},
		},
		"jwks_settings": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "JSON Web Key Set Settings of the OAuth client. Required if private key JWT client authentication or signed requests is enabled.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"jwks": {
						Type:          schema.TypeString,
						Optional:      true,
						Description:   "JSON Web Key Set (JWKS) document of the OAuth client. Either 'jwks' or 'jwks_url' must be provided if private key JWT client authentication or signed requests is enabled. If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.",
						ConflictsWith: []string{"jwks_settings.0.jwks_url"},
					},
					"jwks_url": {
						Type:          schema.TypeString,
						Optional:      true,
						Description:   "JSON Web Key Set (JWKS) URL of the OAuth client. Either 'jwks' or 'jwks_url' must be provided if private key JWT client authentication or signed requests is enabled. If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.",
						ConflictsWith: []string{"jwks_settings.0.jwks"},
					},
				},
			},
		},
		"ciba_delivery_mode": {
			Type:        schema.TypeString,
			Description: "The token delivery mode for the client.  The default value is 'POLL'.",
			Optional:    true,
		},
		"ciba_notification_endpoint": {
			Type:        schema.TypeString,
			Description: "The endpoint the OP will call after a successful or failed end-user authentication.",
			Optional:    true,
		},
		"ciba_polling_interval": {
			Type:        schema.TypeInt,
			Description: "The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds.",
			Optional:    true,
		},
		"ciba_request_object_signing_algorithm": {
			Type:        schema.TypeString,
			Description: "The JSON Web Signature [JWS] algorithm that must be used to sign the CIBA Request Object. All signing algorithms are allowed if value is not present.",
			Optional:    true,
			ValidateFunc: validation.StringInSlice([]string{
				"RS256", "RS384", "RS512", "ES256", "ES384", "ES512", "PS256", "PS384", "PS512",
			}, false),
		},
		"ciba_require_signed_requests": {
			Type:        schema.TypeBool,
			Description: "Determines whether CIBA signed requests are required for this client.",
			Optional:    true,
		},
		"ciba_user_code_supported": {
			Type:        schema.TypeBool,
			Description: "Determines whether CIBA user code is supported for this client.",
			Optional:    true,
		},
		"bypass_activation_code_confirmation_override": {
			Type:        schema.TypeBool,
			Description: "Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device. This overrides the 'bypassUseCodeConfirmation' value present in Authorization Server Settings.",
			Optional:    true,
		},
		"device_flow_setting_type": {
			Type:        schema.TypeString,
			Description: "Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
			Optional:    true,
			Default:     "SERVER_DEFAULT",
		},
		"device_polling_interval_override": {
			Type:        schema.TypeInt,
			Description: "The amount of time client should wait between polling requests, in seconds. This overrides the 'devicePollingInterval' value present in Authorization Server Settings.",
			Optional:    true,
		},
		"extended_properties": {
			Type:        schema.TypeSet,
			Description: "OAuth Client Metadata can be extended to use custom Client Metadata Parameters. The names of these custom parameters should be defined in /extendedProperties.",
			Elem:        resourceParameterValues(),
			Optional:    true,
		},
		"pending_authorization_timeout_override": {
			Type:        schema.TypeInt,
			Description: "The 'device_code' and 'user_code' timeout, in seconds. This overrides the 'pendingAuthorizationTimeout' value present in Authorization Server Settings.",
			Optional:    true,
		},
		"persistent_grant_idle_timeout": {
			Type:        schema.TypeInt,
			Description: "The persistent grant idle timeout.",
			Optional:    true,
		},
		"persistent_grant_idle_timeout_time_unit": {
			Type:        schema.TypeString,
			Description: "The persistent grant idle timeout time unit.",
			Optional:    true,
			Default:     "DAYS",
		},
		"persistent_grant_idle_timeout_type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
			Default:     "SERVER_DEFAULT",
		},
		"request_object_signing_algorithm": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The JSON Web Signature [JWS] algorithm that must be used to sign the Request Object. All signing algorithms are allowed if value is not present.",
		},
		"request_policy_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The CIBA request policy.",
			Elem:        resourceLinkResource(),
		},
		"require_proof_key_for_code_exchange": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Determines whether Proof Key for Code Exchange (PKCE) is required for this client.",
		},
		"require_pushed_authorization_requests": {
			Type:        schema.TypeBool,
			Description: "Determines whether pushed authorization requests are required when initiating an authorization request. The default is false.",
			Optional:    true,
		},
		"token_exchange_processor_policy_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The Token Exchange Processor policy.",
			Elem:        resourceLinkResource(),
		},
		"user_authorization_url_override": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The URL used as 'verification_url' and 'verification_url_complete' values in a Device Authorization request. This property overrides the 'userAuthorizationUrl' value present in Authorization Server Settings.",
		},
		"restrict_to_default_access_token_manager": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Determines whether the client is restricted to using only its default access token manager. The default is false.",
		},
	}
}

func resourcePingFederateOauthClientResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.CreateClientInput{
		Body: *resourcePingFederateOauthClientResourceReadData(d),
	}
	var result *pf.Client
	err := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		var err error
		var resp *http.Response
		result, resp, err = svc.CreateClientWithContext(ctx, &input)
		if resp != nil && resp.StatusCode == http.StatusUnprocessableEntity {
			return resource.RetryableError(fmt.Errorf("unable to create with retry OauthClients: %s", err))
		}
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("unable to create OauthClients: %s", err))
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*result.ClientId)
	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.GetClientInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetClientWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OauthClients: %s", err)
	}
	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.UpdateClientInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthClientResourceReadData(d),
	}
	result, _, err := svc.UpdateClientWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OauthClients: %s", err)
	}

	return resourcePingFederateOauthClientResourceReadResult(d, result)
}

func resourcePingFederateOauthClientResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClients
	input := oauthClients.DeleteClientInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteClientWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete OauthClients: %s", err)
	}
	return nil
}

func resourcePingFederateOauthClientResourceReadResult(d *schema.ResourceData, rv *pf.Client) diag.Diagnostics {
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
	setResourceDataBoolWithDiagnostic(d, "require_pushed_authorization_requests", rv.RequirePushedAuthorizationRequests, &diags)
	setResourceDataBoolWithDiagnostic(d, "require_proof_key_for_code_exchange", rv.RequireProofKeyForCodeExchange, &diags)
	setResourceDataStringWithDiagnostic(d, "user_authorization_url_override", rv.UserAuthorizationUrlOverride, &diags)
	setResourceDataBoolWithDiagnostic(d, "restrict_to_default_access_token_manager", rv.RestrictToDefaultAccessTokenManager, &diags)

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

	if v, ok := d.GetOkExists("bypass_approval_page"); ok {
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

	if v, ok := d.GetOkExists("require_signed_requests"); ok {
		client.RequireSignedRequests = Bool(v.(bool))
	}

	if v, ok := d.GetOkExists("restrict_scopes"); ok {
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

	if v, ok := d.GetOkExists("validate_using_all_eligible_atms"); ok {
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
	if v, ok := d.GetOkExists("ciba_require_signed_requests"); ok {
		client.CibaRequireSignedRequests = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("ciba_user_code_supported"); ok {
		client.CibaUserCodeSupported = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("bypass_activation_code_confirmation_override"); ok {
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
	if v, ok := d.GetOkExists("require_proof_key_for_code_exchange"); ok {
		client.RequireProofKeyForCodeExchange = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("require_pushed_authorization_requests"); ok {
		client.RequirePushedAuthorizationRequests = Bool(v.(bool))
	}
	if v, ok := d.GetOk("token_exchange_processor_policy_ref"); ok && len(v.([]interface{})) > 0 {
		client.TokenExchangeProcessorPolicyRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("user_authorization_url_override"); ok {
		client.UserAuthorizationUrlOverride = String(v.(string))
	}
	if v, ok := d.GetOkExists("restrict_to_default_access_token_manager"); ok {
		client.RestrictToDefaultAccessTokenManager = Bool(v.(bool))
	}
	return client
}
