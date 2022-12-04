package sdkv2provider

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"
)

func resourcePingFederateSpIdpConnectionResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for SP IDP Connections within PingFederate.",
		CreateContext: resourcePingFederateSpIdpConnectionResourceCreate,
		ReadContext:   resourcePingFederateSpIdpConnectionResourceRead,
		UpdateContext: resourcePingFederateSpIdpConnectionResourceUpdate,
		DeleteContext: resourcePingFederateSpIdpConnectionResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateSpIdpConnectionResourceSchema(),
	}
}

func resourcePingFederateSpIdpConnectionResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the connection. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.",
		},
		"active": {
			Type:        schema.TypeBool,
			Description: "Specifies whether the connection is active and ready to process incoming requests. The default value is false.",
			Optional:    true,
		},
		"additional_allowed_entities_configuration": {
			Type:        schema.TypeList,
			Elem:        resourceAdditionalAllowedEntitiesConfiguration(),
			Description: "Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.",
			Optional:    true,
		},
		"attribute_query": {
			Type:        schema.TypeList,
			Elem:        resourceIdpAttributeQuery(),
			Description: "The attribute query settings for requesting user attributes from an attribute authority.",
			Optional:    true,
		},
		"base_url": {
			Type:        schema.TypeString,
			Description: "The fully-qualified hostname and port on which your partner's federation deployment runs.",
			Optional:    true,
		},
		"contact_info": {
			Type:        schema.TypeList,
			Elem:        resourceContactInfo(),
			Description: "The contact information for this partner.",
			Optional:    true,
		},
		"credentials": {
			Type:        schema.TypeList,
			Elem:        resourceConnectionCredentials(),
			Description: "The certificates and settings for encryption, signing, and signature verification. It is required for  SAMLx.x and WS-Fed Connections.",
			Optional:    true,
		},
		"default_virtual_entity_id": {
			Type:        schema.TypeString,
			Description: "The default alternate entity ID that identifies the local server to this partner. It is required when virtualEntityIds is not empty and must be included in that list.",
			Optional:    true,
		},
		"entity_id": {
			Type:        schema.TypeString,
			Description: "The partner's entity ID (connection ID) or issuer value (for OIDC Connections).",
			Required:    true,
		},
		"error_page_msg_id": {
			Type:        schema.TypeString,
			Description: "Identifier that specifies the message displayed on a user-facing error page.",
			Optional:    true,
		},
		"extended_properties": {
			Type:        schema.TypeSet,
			Elem:        resourceParameterValues(),
			Description: "Extended Properties allows to store additional information for IdP/SP Connections. The names of these extended properties should be defined in /extendedProperties.",
			Optional:    true,
		},
		"idp_browser_sso": {
			Type:        schema.TypeList,
			Elem:        resourceIdpBrowserSso(),
			Description: "The browser-based SSO settings used to communicate with your IdP.",
			Optional:    true,
		},
		"idp_oauth_grant_attribute_mapping": {
			Type:        schema.TypeList,
			Elem:        resourceIdpOAuthGrantAttributeMapping(),
			Description: "The OAuth Assertion Grant settings used to map from your IdP.",
			Optional:    true,
		},
		"license_connection_group": {
			Type:        schema.TypeString,
			Description: "The license connection group. If your PingFederate license is based on connection groups, each connection must be assigned to a group before it can be used.",
			Optional:    true,
		},
		"logging_mode": {
			Type:        schema.TypeString,
			Description: "The level of transaction logging applicable for this connection. Default is STANDARD.",
			Optional:    true,
		},
		"metadata_reload_settings": {
			Type:        schema.TypeList,
			Elem:        resourceConnectionMetadataUrl(),
			Description: "Connection metadata automatic reload settings.",
			Optional:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The connection name.",
			Required:    true,
		},
		"oidc_client_credentials": {
			Type:        schema.TypeList,
			Elem:        resourceOIDCClientCredentials(),
			Description: "The OIDC client credentials. This is required for an OIDC connection.",
			Optional:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "The type of this connection. Default is 'IDP'.",
			Optional:    true,
			Default:     "IDP",
		},
		"virtual_entity_ids": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "List of alternate entity IDs that identifies the local server to this partner.",
			Optional:    true,
		},
		"ws_trust": {
			Type:        schema.TypeList,
			Elem:        resourceIdpWsTrust(),
			Description: "The Ws-Trust settings.",
			Optional:    true,
		},
	}
}

func resourcePingFederateSpIdpConnectionResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpIdpConnections
	input := spIdpConnections.CreateConnectionInput{
		Body:                     *resourcePingFederateSpIdpConnectionResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreateConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create SpIdpConnections: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateSpIdpConnectionResourceReadResult(d, result)
}

func resourcePingFederateSpIdpConnectionResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpIdpConnections
	input := spIdpConnections.GetConnectionInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read SpIdpConnections: %s", err)
	}
	return resourcePingFederateSpIdpConnectionResourceReadResult(d, result)
}

func resourcePingFederateSpIdpConnectionResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpIdpConnections
	input := spIdpConnections.UpdateConnectionInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateSpIdpConnectionResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update SpIdpConnections: %s", err)
	}

	return resourcePingFederateSpIdpConnectionResourceReadResult(d, result)
}

func resourcePingFederateSpIdpConnectionResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpIdpConnections
	input := spIdpConnections.DeleteConnectionInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete SpIdpConnections: %s", err)
	}
	return nil
}

func resourcePingFederateSpIdpConnectionResourceReadResult(d *schema.ResourceData, rv *pf.IdpConnection) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "connection_id", rv.Id, &diags)
	setResourceDataBoolWithDiagnostic(d, "active", rv.Active, &diags)
	setResourceDataStringWithDiagnostic(d, "error_page_msg_id", rv.ErrorPageMsgId, &diags)
	if rv.AdditionalAllowedEntitiesConfiguration != nil {
		if err := d.Set("additional_allowed_entities_configuration", flattenAdditionalAllowedEntitiesConfiguration(rv.AdditionalAllowedEntitiesConfiguration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.AttributeQuery != nil {
		if err := d.Set("attribute_query", flattenIdpAttributeQuery(rv.AttributeQuery)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "base_url", rv.BaseUrl, &diags)
	if rv.ContactInfo != nil {
		if err := d.Set("contact_info", flattenContactInfo(rv.ContactInfo)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Credentials != nil {
		if err := d.Set("credentials", flattenConnectionCredentials(rv.Credentials)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "default_virtual_entity_id", rv.DefaultVirtualEntityId, &diags)
	setResourceDataStringWithDiagnostic(d, "entity_id", rv.EntityId, &diags)
	if rv.ExtendedProperties != nil {
		if err := d.Set("extended_properties", flattenMapOfParameterValues(rv.ExtendedProperties)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.IdpBrowserSso != nil {
		if err := d.Set("idp_browser_sso", flattenIdpBrowserSso(rv.IdpBrowserSso)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.IdpOAuthGrantAttributeMapping != nil {
		if err := d.Set("idp_oauth_grant_attribute_mapping", flattenIdpOAuthGrantAttributeMapping(rv.IdpOAuthGrantAttributeMapping)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.OidcClientCredentials != nil {
		if err := d.Set("oidc_client_credentials", flattenOIDCClientCredentials(rv.OidcClientCredentials)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "license_connection_group", rv.LicenseConnectionGroup, &diags)
	setResourceDataStringWithDiagnostic(d, "logging_mode", rv.LoggingMode, &diags)
	if rv.MetadataReloadSettings != nil {
		if err := d.Set("metadata_reload_settings", flattenConnectionMetadataUrl(rv.MetadataReloadSettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "type", rv.Type, &diags)
	if rv.VirtualEntityIds != nil {
		if err := d.Set("virtual_entity_ids", rv.VirtualEntityIds); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.WsTrust != nil {
		if err := d.Set("ws_trust", flattenIdpWsTrust(rv.WsTrust)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateSpIdpConnectionResourceReadData(d *schema.ResourceData) *pf.IdpConnection {
	result := pf.IdpConnection{
		VirtualEntityIds: &[]*string{},
	}
	if val, ok := d.GetOk("error_page_msg_id"); ok {
		result.ErrorPageMsgId = String(val.(string))
	}
	if val, ok := d.GetOk("idp_browser_sso"); ok && len(val.([]interface{})) > 0 {
		result.IdpBrowserSso = expandIdpBrowserSso(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("idp_oauth_grant_attribute_mapping"); ok && len(val.([]interface{})) > 0 {
		result.IdpOAuthGrantAttributeMapping = expandIdpOAuthGrantAttributeMapping(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("extended_properties"); ok {
		result.ExtendedProperties = expandMapOfParameterValues(val.(*schema.Set).List())
	}
	if val, ok := d.GetOk("oidc_client_credentials"); ok && len(val.([]interface{})) > 0 {
		result.OidcClientCredentials = expandOIDCClientCredentials(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("connection_id"); ok {
		result.Id = String(val.(string))
	}
	if val, ok := d.GetOk("entity_id"); ok {
		result.EntityId = String(val.(string))
	}
	if val, ok := d.GetOk("base_url"); ok {
		result.BaseUrl = String(val.(string))
	}
	if val, ok := d.GetOk("default_virtual_entity_id"); ok {
		result.DefaultVirtualEntityId = String(val.(string))
	}
	if val, ok := d.GetOk("metadata_reload_settings"); ok && len(val.([]interface{})) > 0 {
		result.MetadataReloadSettings = expandConnectionMetadataUrl(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("credentials"); ok && len(val.([]interface{})) > 0 {
		result.Credentials = expandConnectionCredentials(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("name"); ok {
		result.Name = String(val.(string))
	}
	if val, ok := d.GetOkExists("active"); ok {
		result.Active = Bool(val.(bool))
	}
	if val, ok := d.GetOk("virtual_entity_ids"); ok {
		strs := expandStringList(val.([]interface{}))
		result.VirtualEntityIds = &strs
	}
	if val, ok := d.GetOk("contact_info"); ok && len(val.([]interface{})) > 0 {
		result.ContactInfo = expandContactInfo(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("license_connection_group"); ok {
		result.LicenseConnectionGroup = String(val.(string))
	}
	if val, ok := d.GetOk("additional_allowed_entities_configuration"); ok && len(val.([]interface{})) > 0 {
		result.AdditionalAllowedEntitiesConfiguration = expandAdditionalAllowedEntitiesConfiguration(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("type"); ok {
		result.Type = String(val.(string))
	}
	if val, ok := d.GetOk("attribute_query"); ok && len(val.([]interface{})) > 0 {
		result.AttributeQuery = expandIdpAttributeQuery(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("ws_trust"); ok && len(val.([]interface{})) > 0 {
		result.WsTrust = expandIdpWsTrust(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("logging_mode"); ok {
		result.LoggingMode = String(val.(string))
	}

	return &result
}
