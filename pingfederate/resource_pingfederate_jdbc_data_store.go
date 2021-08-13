package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"regexp"

	"github.com/hashicorp/go-cty/cty"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateJdbcDataStoreResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for Jdbc Data Stores within PingFederate.",
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
		"data_store_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.",
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				v := value.(string)
				r, _ := regexp.Compile(`^[a-zA-Z0-9._-]+$`)
				if !r.MatchString(v) {
					return diag.Errorf("the data_store_id can only contain alphanumeric characters, dash, dot and underscore.")
				}
				return nil
			},
		},
		"mask_attribute_values": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether attribute values should be masked in the log.",
		},
		"connection_url_tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The set of connection URLs and associated tags for this JDBC data store.",
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
			Type:        schema.TypeString,
			Required:    true,
			Description: "The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags are specified.",
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the connection url and the username.",
		},
		"driver_class": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the driver class used to communicate with the source database.",
		},
		"user_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name that identifies the user when connecting to the database.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The password needed to access the database. GETs will not return this attribute. To update this field, specify the new value in this attribute.",
		},
		"encrypted_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The encrypted password needed to access the database. If you do not want to update the stored value, this attribute should be passed back unchanged.",
		},
		"validate_connection_sql": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed.",
		},
		"allow_multi_value_attributes": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute.",
		},
		"min_pool_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     10,
			Description: "The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.",
		},
		"max_pool_size": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     100,
			Description: "The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.",
		},
		"blocking_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     5000,
			Description: "The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default.",
		},
		"idle_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     5,
			Description: "The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default.",
		},
	}
}

func resourcePingFederateJdbcDataStoreResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateJdbcDataStoreResourceReadData(d)
	input := dataStores.CreateJdbcDataStoreInput{
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.CreateJdbcDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create JdbcDataStores: %s", err)
	}
	d.SetId(*store.Id)
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, store)
}

func resourcePingFederateJdbcDataStoreResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	input := dataStores.GetJdbcDataStoreInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetJdbcDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read JdbcDataStores: %s", err)
	}
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, result)
}

func resourcePingFederateJdbcDataStoreResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateJdbcDataStoreResourceReadData(d)
	input := dataStores.UpdateJdbcDataStoreInput{
		Id:                       d.Id(),
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.UpdateJdbcDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update JdbcDataStores: %s", err)
	}
	return resourcePingFederateJdbcDataStoreResourceReadResult(d, store)
}

func resourcePingFederateJdbcDataStoreResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).DataStores
	input := dataStores.DeleteDataStoreInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete JdbcDataStores: %s", err)
	}
	return nil
}

func resourcePingFederateJdbcDataStoreResourceReadResult(d *schema.ResourceData, rv *pf.JdbcDataStore) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "mask_attribute_values", rv.MaskAttributeValues, &diags)
	setResourceDataStringWithDiagnostic(d, "data_store_id", rv.Id, &diags)
	if rv.ConnectionUrlTags != nil && len(*rv.ConnectionUrlTags) != 0 {
		if err := d.Set("connection_url_tags", flattenJdbcTagConfigs(rv.ConnectionUrlTags)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	setResourceDataStringWithDiagnostic(d, "connection_url", rv.ConnectionUrl, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "driver_class", rv.DriverClass, &diags)
	setResourceDataStringWithDiagnostic(d, "user_name", rv.UserName, &diags)
	//TODO i need to handle this not being set
	setResourceDataStringWithDiagnostic(d, "password", rv.Password, &diags)
	setResourceDataStringWithDiagnostic(d, "encrypted_password", rv.EncryptedPassword, &diags)
	setResourceDataStringWithDiagnostic(d, "validate_connection_sql", rv.ValidateConnectionSql, &diags)
	setResourceDataBoolWithDiagnostic(d, "allow_multi_value_attributes", rv.AllowMultiValueAttributes, &diags)
	setResourceDataIntWithDiagnostic(d, "min_pool_size", rv.MinPoolSize, &diags)
	setResourceDataIntWithDiagnostic(d, "max_pool_size", rv.MaxPoolSize, &diags)
	setResourceDataIntWithDiagnostic(d, "blocking_timeout", rv.BlockingTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "idle_timeout", rv.IdleTimeout, &diags)

	return diags
}

func resourcePingFederateJdbcDataStoreResourceReadData(d *schema.ResourceData) *pf.JdbcDataStore {
	ds := &pf.JdbcDataStore{}
	if v, ok := d.GetOk("data_store_id"); ok {
		ds.Id = String(v.(string))
	}
	if v, ok := d.GetOkExists("mask_attribute_values"); ok {
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
	if val, ok := d.GetOkExists("allow_multi_value_attributes"); ok {
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
