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

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
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
						},
						"id_token_signing_algorithm": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "RS256",
							ValidateDiagFunc: validateTokenSigningAlgorithm,
						},
						"logout_uris": {
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ping_access_logout_capable": {
							Type:     schema.TypeBool,
							Optional: true,
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
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringithDiagnostic(d, "client_id", rv.ClientId, &diags)
	setResourceDataBoolWithDiagnostic(d, "bypass_approval_page", rv.BypassApprovalPage, &diags)
	setResourceDataStringithDiagnostic(d, "description", rv.Description, &diags)
	setResourceDataStringithDiagnostic(d, "logo_url", rv.LogoUrl, &diags)
	setResourceDataIntWithDiagnostic(d, "persistent_grant_expiration_time", rv.PersistentGrantExpirationTime, &diags)
	setResourceDataStringithDiagnostic(d, "persistent_grant_expiration_time_unit", rv.PersistentGrantExpirationTimeUnit, &diags)
	setResourceDataStringithDiagnostic(d, "persistent_grant_expiration_type", rv.PersistentGrantExpirationType, &diags)
	setResourceDataStringithDiagnostic(d, "refresh_rolling", rv.RefreshRolling, &diags)
	setResourceDataBoolWithDiagnostic(d, "require_signed_requests", rv.RequireSignedRequests, &diags)
	setResourceDataBoolWithDiagnostic(d, "restrict_scopes", rv.RestrictScopes, &diags)
	setResourceDataBoolWithDiagnostic(d, "validate_using_all_eligible_atms", rv.ValidateUsingAllEligibleAtms, &diags)
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
		GrantTypes: &grants,
	}

	if _, ok := d.GetOk("bypass_approval_page"); ok {
		client.BypassApprovalPage = Bool(d.Get("bypass_approval_page").(bool))
	}

	if _, ok := d.GetOk("description"); ok {
		client.Description = String(d.Get("description").(string))
	}

	if _, ok := d.GetOk("exclusive_scopes"); ok {
		strs := expandStringList(d.Get("exclusive_scopes").(*schema.Set).List())
		client.ExclusiveScopes = &strs
	}

	if _, ok := d.GetOk("logo_url"); ok {
		client.LogoUrl = String(d.Get("logo_url").(string))
	}

	if _, ok := d.GetOk("persistent_grant_expiration_time"); ok {
		client.PersistentGrantExpirationTime = Int(d.Get("persistent_grant_expiration_time").(int))
	}

	if _, ok := d.GetOk("persistent_grant_expiration_time_unit"); ok {
		client.PersistentGrantExpirationTimeUnit = String(d.Get("persistent_grant_expiration_time_unit").(string))
	}

	if _, ok := d.GetOk("persistent_grant_expiration_type"); ok {
		client.PersistentGrantExpirationType = String(d.Get("persistent_grant_expiration_type").(string))
	}

	if _, ok := d.GetOk("redirect_uris"); ok {
		strs := expandStringList(d.Get("redirect_uris").(*schema.Set).List())
		client.RedirectUris = &strs
	}

	if _, ok := d.GetOk("refresh_rolling"); ok {
		client.RefreshRolling = String(d.Get("refresh_rolling").(string))
	}

	if _, ok := d.GetOk("require_signed_requests"); ok {
		client.RequireSignedRequests = Bool(d.Get("require_signed_requests").(bool))
	}

	if _, ok := d.GetOk("restrict_scopes"); ok {
		client.RestrictScopes = Bool(d.Get("restrict_scopes").(bool))
	}

	if _, ok := d.GetOk("restricted_response_types"); ok {
		strs := expandStringList(d.Get("restricted_response_types").(*schema.Set).List())
		client.RestrictedResponseTypes = &strs
	}

	if _, ok := d.GetOk("restricted_scopes"); ok {
		strs := expandStringList(d.Get("restricted_scopes").(*schema.Set).List())
		client.RestrictedScopes = &strs
	}

	if _, ok := d.GetOk("validate_using_all_eligible_atms"); ok {
		client.ValidateUsingAllEligibleAtms = Bool(d.Get("validate_using_all_eligible_atms").(bool))
	}

	if v, ok := d.GetOk("client_auth"); ok && len(v.([]interface{})) > 0 {
		client.ClientAuth = expandClientAuth(v.([]interface{}))
	}

	if _, ok := d.GetOk("jwks_settings"); ok {
		client.JwksSettings = expandJwksSettings(d.Get("jwks_settings").(*schema.Set).List())
	}

	if v, ok := d.GetOk("default_access_token_manager_ref"); ok {
		client.DefaultAccessTokenManagerRef = expandResourceLink(v.([]interface{}))
	}

	if v, ok := d.GetOk("oidc_policy"); ok && len(v.([]interface{})) > 0 {
		client.OidcPolicy = expandClientOIDCPolicy(v.([]interface{}))
	}

	return client
}
