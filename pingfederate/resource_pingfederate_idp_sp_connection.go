package pingfederate

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
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
		"connection_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"additional_allowed_entities_configuration": {
			Type:     schema.TypeList,
			Elem:     resourceAdditionalAllowedEntitiesConfiguration(),
			Optional: true,
		},
		"application_icon_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"application_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"attribute_query": {
			Type:     schema.TypeList,
			Elem:     resourceSpAttributeQuery(),
			Optional: true,
		},
		"base_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"contact_info": {
			Type:     schema.TypeList,
			Elem:     resourceContactInfo(),
			Optional: true,
		},
		"credentials": {
			Type:     schema.TypeList,
			Elem:     resourceConnectionCredentials(),
			Optional: true,
		},
		"default_virtual_entity_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"entity_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"extended_properties": {
			Type:     schema.TypeSet,
			Elem:     resourceParameterValues(),
			Optional: true,
		},
		"license_connection_group": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"logging_mode": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"metadata_reload_settings": {
			Type:     schema.TypeList,
			Elem:     resourceConnectionMetadataUrl(),
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"outbound_provision": {
			Type:     schema.TypeList,
			Elem:     resourceOutboundProvision(),
			Optional: true,
		},
		"sp_browser_sso": {
			Type:     schema.TypeList,
			Elem:     resourceSpBrowserSso(),
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SP",
		},
		"virtual_entity_ids": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Optional: true,
		},
		"ws_trust": {
			Type:     schema.TypeList,
			Elem:     resourceSpWsTrust(),
			Optional: true,
		},
	}
}

func resourcePingFederateIdpSpConnectionResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.CreateConnectionInput{
		Body:                     *resourcePingFederateIdpSpConnectionResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.CreateConnection(&input)
	if err != nil {
		return diag.Errorf("unable to create IdpSpConnections: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.GetConnectionInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetConnection(&input)
	if err != nil {
		return diag.Errorf("unable to read IdpSpConnections: %s", err)
	}
	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.UpdateConnectionInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateIdpSpConnectionResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.UpdateConnection(&input)
	if err != nil {
		return diag.Errorf("unable to update IdpSpConnections: %s", err)
	}

	return resourcePingFederateIdpSpConnectionResourceReadResult(d, result)
}

func resourcePingFederateIdpSpConnectionResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpSpConnections
	input := idpSpConnections.DeleteConnectionInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteConnection(&input)
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
		if err := d.Set("outbound_provision", flattenOutboundProvision(rv.OutboundProvision)); err != nil {
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
	if val, ok := d.GetOk("active"); ok {
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
