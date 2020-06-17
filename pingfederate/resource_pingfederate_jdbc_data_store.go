package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateJdbcDataStoreResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateJdbcDataStoreResourceCreate,
		ReadContext:   resourcePingFederateJdbcDataStoreResourceRead,
		UpdateContext: resourcePingFederateJdbcDataStoreResourceUpdate,
		DeleteContext: resourcePingFederateJdbcDataStoreResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateJdbcDataStoreResourceSchema(),
	}
}

func resourcePingFederateJdbcDataStoreResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
		"mask_attribute_values": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"connection_url_tags": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"connection_url": {
						Type:     schema.TypeString,
						Required: true,
					},
					"tags": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"default_source": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
				},
			},
		},
		"connection_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"driver_class": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"encrypted_password": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"validate_connection_sql": {
			Type:     schema.TypeString,
			Optional: true,
			//Default: "",
		},
		"allow_multi_value_attributes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"min_pool_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"max_pool_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"blocking_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"idle_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
	}
}

func resourcePingFederateJdbcDataStoreResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateJdbcDataStoreResourceReadData(d)
	input := dataStores.CreateJdbcDataStoreInput{
		Body:                     *ds,
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	store, _, err := svc.CreateJdbcDataStore(&input)
	if err != nil {
		return diag.Errorf("unable to create JdbcDataStores: %s", err)
	}
	d.SetId(*store.Id)
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, store)
}

func resourcePingFederateJdbcDataStoreResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	input := dataStores.GetJdbcDataStoreInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetJdbcDataStore(&input)
	if err != nil {
		return diag.Errorf("unable to read JdbcDataStores: %s", err)
	}
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, result)
}

func resourcePingFederateJdbcDataStoreResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateJdbcDataStoreResourceReadData(d)
	input := dataStores.UpdateJdbcDataStoreInput{
		Id:                       d.Id(),
		Body:                     *ds,
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	store, _, err := svc.UpdateJdbcDataStore(&input)
	if err != nil {
		return diag.Errorf("unable to update JdbcDataStores: %s", err)
	}
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, store)
}

func resourcePingFederateJdbcDataStoreResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	input := dataStores.DeleteDataStoreInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDataStore(&input)
	if err != nil {
		return diag.Errorf("unable to delete JdbcDataStores: %s", err)
	}
	return nil
}

func resourcePingFederateJdbcDataStoreResourceReadResult(d *schema.ResourceData, rv *pf.JdbcDataStore) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "mask_attribute_values", rv.MaskAttributeValues, &diags)

	if rv.ConnectionUrlTags != nil && len(*rv.ConnectionUrlTags) != 0 {
		if err := d.Set("connection_url_tags", flattenJdbcTagConfigs(rv.ConnectionUrlTags)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringithDiagnostic(d, "connection_url", rv.ConnectionUrl, &diags)
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringithDiagnostic(d, "driver_class", rv.DriverClass, &diags)
	setResourceDataStringithDiagnostic(d, "user_name", rv.UserName, &diags)
	//TODO i need to handle this not being set
	setResourceDataStringithDiagnostic(d, "password", rv.Password, &diags)
	setResourceDataStringithDiagnostic(d, "encrypted_password", rv.EncryptedPassword, &diags)
	setResourceDataStringithDiagnostic(d, "validate_connection_sql", rv.ValidateConnectionSql, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_multi_value_attributes", rv.AllowMultiValueAttributes, &diags)
	setResourceDataIntWithDiagnostic(d, "min_pool_size", rv.MinPoolSize, &diags)
	setResourceDataIntWithDiagnostic(d, "max_pool_size", rv.MaxPoolSize, &diags)
	setResourceDataIntWithDiagnostic(d, "blocking_timeout", rv.BlockingTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "idle_timeout", rv.IdleTimeout, &diags)

	return diags
}

func resourcePingFederateJdbcDataStoreResourceReadData(d *schema.ResourceData) *pf.JdbcDataStore {
	ds := &pf.JdbcDataStore{}
	if v, ok := d.GetOk("mask_attribute_values"); ok {
		ds.MaskAttributeValues = Bool(v.(bool))
	}
	if val, ok := d.GetOk("connection_url_tags"); ok && len(val.(*schema.Set).List()) > 0 {
		ds.ConnectionUrlTags = expandJdbcTagConfigs(val.(*schema.Set).List())
	}
	if val, ok := d.GetOk("connection_url"); ok {
		ds.ConnectionUrl = String(val.(string))
	}
	if val, ok := d.GetOk("name"); ok {
		ds.Name = String(val.(string))
	}
	if val, ok := d.GetOk("driver_class"); ok {
		ds.DriverClass = String(val.(string))
	}
	if val, ok := d.GetOk("user_name"); ok {
		ds.UserName = String(val.(string))
	}
	if val, ok := d.GetOk("password"); ok {
		ds.Password = String(val.(string))
	}
	if val, ok := d.GetOk("encrypted_password"); ok {
		ds.EncryptedPassword = String(val.(string))
	}
	if val, ok := d.GetOk("validate_connection_sql"); ok {
		ds.ValidateConnectionSql = String(val.(string))
	}
	if val, ok := d.GetOk("allow_multi_value_attributes"); ok {
		ds.AllowMultiValueAttributes = Bool(val.(bool))
	}
	if val, ok := d.GetOk("min_pool_size"); ok {
		ds.MinPoolSize = Int(val.(int))
	}
	if val, ok := d.GetOk("max_pool_size"); ok {
		ds.MaxPoolSize = Int(val.(int))
	}
	if val, ok := d.GetOk("blocking_timeout"); ok {
		ds.BlockingTimeout = Int(val.(int))
	}
	if val, ok := d.GetOk("idle_timeout"); ok {
		ds.IdleTimeout = Int(val.(int))
	}
	ds.Type = String("JDBC")

	return ds
}
