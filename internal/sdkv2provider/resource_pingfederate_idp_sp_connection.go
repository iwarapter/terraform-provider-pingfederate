package sdkv2provider

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpSpConnections"
)

func resourcePingFederateIdpSpConnectionResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateIdpSpConnectionResourceCreate,
		ReadContext:   resourcePingFederateIdpSpConnectionResourceRead,
		UpdateContext: resourcePingFederateIdpSpConnectionResourceUpdate,
		DeleteContext: resourcePingFederateIdpSpConnectionResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateIdpSpConnectionResourceSchema(),
	}
}

func resourcePingFederateIdpSpConnectionResourceSchema() map[string]*schema.Schema {
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
			Default:     false,
		},
		"additional_allowed_entities_configuration": {
			Type:        schema.TypeList,
			Elem:        resourceAdditionalAllowedEntitiesConfiguration(),
			Description: "Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.",
			Optional:    true,
		},
		"application_icon_url": {
			Type:        schema.TypeString,
			Description: "The application icon url.",
			Optional:    true,
		},
		"application_name": {
			Type:        schema.TypeString,
			Description: "The application name.",
			Optional:    true,
		},
		"attribute_query": {
			Type:        schema.TypeList,
			Elem:        resourceSpAttributeQuery(),
			Description: "The attribute query settings for supporting SPs in requesting user attributes.",
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
		"extended_properties": {
			Type:        schema.TypeSet,
			Elem:        resourceParameterValues(),
			Description: "Extended Properties allows to store additional information for IdP/SP Connections. The names of these extended properties should be defined in /extendedProperties.",
			Optional:    true,
		},
		"license_connection_group": {
			Type:        schema.TypeString,
			Description: "The license connection group. If your PingFederate license is based on connection groups, each connection must be assigned to a group before it can be used.",
			Optional:    true,
			Default:     "",
		},
		"logging_mode": {
			Type:        schema.TypeString,
			Description: "The level of transaction logging applicable for this connection. Default is STANDARD.",
			Optional:    true,
			Default:     "STANDARD",
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
		"outbound_provision": {
			Type:        schema.TypeList,
			Elem:        resourceOutboundProvision(),
			Description: "The Outbound Provision settings.",
			Optional:    true,
		},
		"sp_browser_sso": {
			Type:        schema.TypeList,
			Elem:        resourceSpBrowserSso(),
			Description: "The browser-based SSO settings used to communicate with your SP.",
			Optional:    true,
		},
		"type": {
			Type:        schema.TypeString,
			Description: "The type of this connection. This must be set to 'SP'.",
			Optional:    true,
			Default:     "SP",
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
			Elem:        resourceSpWsTrust(),
			Description: "The Ws-Trust settings.",
			Optional:    true,
		},
	}
}

func resourcePingFederateIdpSpConnectionResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	cli := m.(pfClient)
	input := idpSpConnections.CreateConnectionInput{
		Body:                     *spConnectionVersionSpecificConfig(cli, resourcePingFederateIdpSpConnectionResourceReadData(d)),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreateConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create IdpSpConnections: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.GetConnectionInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read IdpSpConnections: %s", err)
	}
	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	cli := m.(pfClient)
	input := idpSpConnections.UpdateConnectionInput{
		Id:                       d.Id(),
		Body:                     *spConnectionVersionSpecificConfig(cli, resourcePingFederateIdpSpConnectionResourceReadData(d)),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update IdpSpConnections: %s", err)
	}

	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.DeleteConnectionInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete IdpSpConnections: %s", err)
	}
	return nil
}

func resourcePingFederateIdpSpConnectionResourceReadResult(d *schema.ResourceData, rv *pf.SpConnection) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "connection_id", rv.Id, &diags)
	setResourceDataBoolWithDiagnostic(d, "active", rv.Active, &diags)
	if rv.AdditionalAllowedEntitiesConfiguration != nil {
		if err := d.Set("additional_allowed_entities_configuration", flattenAdditionalAllowedEntitiesConfiguration(rv.AdditionalAllowedEntitiesConfiguration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "application_icon_url", rv.ApplicationIconUrl, &diags)
	setResourceDataStringWithDiagnostic(d, "application_name", rv.ApplicationName, &diags)
	if rv.AttributeQuery != nil {
		if err := d.Set("attribute_query", flattenSpAttributeQuery(rv.AttributeQuery)); err != nil {
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
	setResourceDataStringWithDiagnostic(d, "license_connection_group", rv.LicenseConnectionGroup, &diags)
	setResourceDataStringWithDiagnostic(d, "logging_mode", rv.LoggingMode, &diags)
	if rv.MetadataReloadSettings != nil {
		if err := d.Set("metadata_reload_settings", flattenConnectionMetadataUrl(rv.MetadataReloadSettings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	if rv.OutboundProvision != nil {
		var state []map[string]interface{}
		if val := d.Get("outbound_provision"); len(val.([]interface{})) == 1 {
			orig := expandOutboundProvision(d.Get("outbound_provision").([]interface{})[0].(map[string]interface{}))
			state = maskOutboundProvision(orig, rv.OutboundProvision)
		} else {
			state = flattenOutboundProvision(rv.OutboundProvision)
		}
		if err := d.Set("outbound_provision", state); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.SpBrowserSso != nil {
		if err := d.Set("sp_browser_sso", flattenSpBrowserSso(rv.SpBrowserSso)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "type", rv.Type, &diags)
	if rv.VirtualEntityIds != nil {
		if err := d.Set("virtual_entity_ids", rv.VirtualEntityIds); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.WsTrust != nil {
		if err := d.Set("ws_trust", flattenSpWsTrust(rv.WsTrust)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateIdpSpConnectionResourceReadData(d *schema.ResourceData) *pf.SpConnection {
	result := pf.SpConnection{
		VirtualEntityIds: &[]*string{},
	}
	if val, ok := d.GetOk("extended_properties"); ok {
		result.ExtendedProperties = expandMapOfParameterValues(val.(*schema.Set).List())
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
	if val, ok := d.GetOk("sp_browser_sso"); ok && len(val.([]interface{})) > 0 {
		result.SpBrowserSso = expandSpBrowserSso(val.([]interface{})[0].(map[string]interface{}))
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
		if val.([]interface{})[0] == nil {
			result.ContactInfo = &pf.ContactInfo{}
		} else {
			result.ContactInfo = expandContactInfo(val.([]interface{})[0].(map[string]interface{}))
		}
	}
	if val, ok := d.GetOk("license_connection_group"); ok {
		result.LicenseConnectionGroup = String(val.(string))
	}
	if val, ok := d.GetOk("application_name"); ok {
		result.ApplicationName = String(val.(string))
	}
	if val, ok := d.GetOk("additional_allowed_entities_configuration"); ok && len(val.([]interface{})) > 0 {
		result.AdditionalAllowedEntitiesConfiguration = expandAdditionalAllowedEntitiesConfiguration(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("type"); ok {
		result.Type = String(val.(string))
	}
	if val, ok := d.GetOk("attribute_query"); ok && len(val.([]interface{})) > 0 {
		result.AttributeQuery = expandSpAttributeQuery(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("application_icon_url"); ok {
		result.ApplicationIconUrl = String(val.(string))
	}
	if val, ok := d.GetOk("outbound_provision"); ok && len(val.([]interface{})) > 0 {
		result.OutboundProvision = expandOutboundProvision(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("ws_trust"); ok && len(val.([]interface{})) > 0 {
		result.WsTrust = expandSpWsTrust(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := d.GetOk("logging_mode"); ok {
		result.LoggingMode = String(val.(string))
	}

	return &result
}

func spConnectionVersionSpecificConfig(cli pfClient, in *pf.SpConnection) *pf.SpConnection {
	if in.Credentials != nil {
		if in.Credentials.SigningSettings != nil {
			if cli.IsVersionLessEqThan(11, 0) {
				in.Credentials.SigningSettings.AlternativeSigningKeyPairRefs = nil
			}
		}
	}
	return in
}
