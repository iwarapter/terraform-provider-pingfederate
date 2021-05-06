package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateCustomDataStoreResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateCustomDataStoreResourceCreate,
		ReadContext:   resourcePingFederateCustomDataStoreResourceRead,
		UpdateContext: resourcePingFederateCustomDataStoreResourceUpdate,
		DeleteContext: resourcePingFederateCustomDataStoreResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateCustomDataStoreResourceSchema(),
	}
}

func resourcePingFederateCustomDataStoreResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mask_attribute_values": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
		"parent_ref":            resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
	}
}

func resourcePingFederateCustomDataStoreResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateCustomDataStoreResourceReadData(d)
	input := dataStores.CreateCustomDataStoreInput{
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.CreateCustomDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create CustomDataStores: %s", err)
	}
	d.SetId(*store.Id)
	return resourcePingFederateCustomDataStoreResourceReadResult(d, store, svc)
}

func resourcePingFederateCustomDataStoreResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	input := dataStores.GetCustomDataStoreInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetCustomDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read CustomDataStores: %s", err)
	}
	return resourcePingFederateCustomDataStoreResourceReadResult(d, result, svc)
}

func resourcePingFederateCustomDataStoreResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateCustomDataStoreResourceReadData(d)
	input := dataStores.UpdateCustomDataStoreInput{
		Id:                       d.Id(),
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.UpdateCustomDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update CustomDataStores: %s", err)
	}
	return resourcePingFederateCustomDataStoreResourceReadResult(d, store, svc)
}

func resourcePingFederateCustomDataStoreResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).DataStores
	input := dataStores.DeleteDataStoreInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete CustomDataStores: %s", err)
	}
	return nil
}

func resourcePingFederateCustomDataStoreResourceReadResult(d *schema.ResourceData, rv *pf.CustomDataStore, svc dataStores.DataStoresAPI) diag.Diagnostics {
	desc, _, err := svc.GetCustomDataStoreDescriptor(&dataStores.GetCustomDataStoreDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve IdpAdapters descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "mask_attribute_values", rv.MaskAttributeValues, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	if rv.PluginDescriptorRef != nil {
		if err := d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.ParentRef != nil {
		if err := d.Set("parent_ref", flattenResourceLink(rv.ParentRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.Configuration != nil {
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err := d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}

func resourcePingFederateCustomDataStoreResourceReadData(d *schema.ResourceData) *pf.CustomDataStore {
	ds := &pf.CustomDataStore{
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})[0].(map[string]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}
	if v, ok := d.GetOk("parent_ref"); ok && len(v.([]interface{})) > 0 {
		ds.ParentRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOkExists("mask_attribute_values"); ok {
		ds.MaskAttributeValues = Bool(v.(bool))
	}
	ds.Type = String("CUSTOM")

	return ds
}
