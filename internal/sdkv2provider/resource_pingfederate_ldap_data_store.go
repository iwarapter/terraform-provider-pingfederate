package sdkv2provider

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/go-cty/cty"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateLdapDataStoreResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for Ldap Data Stores within PingFederate.",
		CreateContext: resourcePingFederateLdapDataStoreResourceCreate,
		ReadContext:   resourcePingFederateLdapDataStoreResourceRead,
		UpdateContext: resourcePingFederateLdapDataStoreResourceUpdate,
		DeleteContext: resourcePingFederateLdapDataStoreResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateLdapDataStoreResourceSchema(),
	}
}

func resourcePingFederateLdapDataStoreResourceSchema() map[string]*schema.Schema {
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
			Description: "Whether attribute values should be masked in the log.",
		},
		"hostnames_tags": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			Description: "The set of host names and associated tags for this LDAP data store.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"hostnames": {
						Type:        schema.TypeSet,
						Required:    true,
						Description: "The LDAP host names.",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"tags": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The LDAP host names.",
					},
					"default_source": {
						Type:        schema.TypeBool,
						Optional:    true,
						Default:     false,
						Description: "Whether this is the default connection. Defaults to false if not specified.",
					},
				},
			},
		},
		"hostnames": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The default LDAP host names. This field is required if no mapping for host names and tags are specified.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"name": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal.",
		},
		"ldap_type": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "A type that allows PingFederate to configure many provisioning settings automatically. The 'UNBOUNDID_DS' type has been deprecated, please use the 'PING_DIRECTORY' type instead.",
			ValidateFunc: validation.StringInSlice([]string{
				"ACTIVE_DIRECTORY", "ORACLE_DIRECTORY_SERVER", "ORACLE_UNIFIED_DIRECTORY", "UNBOUNDID_DS", "PING_DIRECTORY", "GENERIC",
			}, false),
		},
		"bind_anonymously": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Whether username and password are required. The default value is false.",
		},
		"user_dn": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The username credential required to access the data store.",
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The password credential required to access the data store. GETs will not return this attribute. To update this field, specify the new value in this attribute.",
		},
		"encrypted_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "The encrypted password credential required to access the data store.  If you do not want to update the stored value, this attribute should be passed back unchanged.",
		},
		"use_ssl": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false.",
		},
		"use_dns_srv_records": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Use DNS SRV Records to discover LDAP server information. The default value is false.",
		},
		"follow_ldap_referrals": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Follow LDAP Referrals in the domain tree. The default value is false. This property does not apply to PingDirectory as this functionality is configured in PingDirectory.",
		},
		"test_on_borrow": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicates whether objects are validated before being borrowed from the pool.",
		},
		"test_on_return": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Indicates whether objects are validated before being returned to the pool.",
		},
		"create_if_necessary": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Indicates whether temporary connections can be created when the Maximum Connections threshold is reached.",
		},
		"verify_host": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Verifies that the presented server certificate includes the address to which the client intended to establish a connection. Omitting this attribute will set the value to true.",
		},
		"min_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The smallest number of connections that can remain in each pool, without creating extra ones. Omitting this attribute will set the value to the default value.",
		},
		"max_connections": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "The largest number of active connections that can remain in each pool without releasing extra ones. Omitting this attribute will set the value to the default value.",
		},
		"max_wait": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     -1,
			Description: "The maximum number of milliseconds the pool waits for a connection to become available when trying to obtain a connection from the pool. Omitting this attribute or setting a value of -1 causes the pool not to wait at all and to either create a new connection or produce an error (when no connections are available).",
		},
		"time_between_evictions": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60000,
			Description: "The frequency, in milliseconds, that the evictor cleans up the connections in the pool. A value of -1 disables the evictor. Omitting this attribute will set the value to the default value.",
		},
		"read_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     3000,
			Description: "The maximum number of milliseconds a connection waits for a response to be returned before producing an error. A value of -1 causes the connection to wait indefinitely. Omitting this attribute will set the value to the default value.",
		},
		"connection_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     3000,
			Description: "The maximum number of milliseconds that a connection attempt should be allowed to continue before returning an error. A value of -1 causes the pool to wait indefinitely. Omitting this attribute will set the value to the default value.",
		},
		"dns_ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Default:     60000,
			Description: "The maximum time in milliseconds that DNS information are cached. Omitting this attribute will set the value to the default value.",
		},
		"ldap_dns_srv_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "_ldap._tcp",
			Description: "The prefix value used to discover LDAP DNS SRV record. Omitting this attribute will set the value to the default value.",
		},
		"ldaps_dns_srv_prefix": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "_ldaps._tcp",
			Description: "The prefix value used to discover LDAPs DNS SRV record. Omitting this attribute will set the value to the default value.",
		},
		"binary_attributes": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The list of LDAP attributes to be handled as binary data.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourcePingFederateLdapDataStoreResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateLdapDataStoreResourceReadData(d)
	input := dataStores.CreateLdapDataStoreInput{
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.CreateLdapDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create LdapDataStores: %s", err)
	}
	d.SetId(*store.Id)
	return resourcePingFederateLdapDataStoreResourceReadResult(d, store)
}

func resourcePingFederateLdapDataStoreResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	input := dataStores.GetLdapDataStoreInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetLdapDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read LdapDataStores: %s", err)
	}
	return resourcePingFederateLdapDataStoreResourceReadResult(d, result)
}

func resourcePingFederateLdapDataStoreResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).DataStores
	ds := resourcePingFederateLdapDataStoreResourceReadData(d)
	input := dataStores.UpdateLdapDataStoreInput{
		Id:                       d.Id(),
		Body:                     *ds,
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	store, _, err := svc.UpdateLdapDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update LdapDataStores: %s", err)
	}
	return resourcePingFederateLdapDataStoreResourceReadResult(d, store)
}

func resourcePingFederateLdapDataStoreResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).DataStores
	input := dataStores.DeleteDataStoreInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDataStoreWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete LdapDataStores: %s", err)
	}
	return nil
}

func resourcePingFederateLdapDataStoreResourceReadResult(d *schema.ResourceData, rv *pf.LdapDataStore) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "data_store_id", rv.Id, &diags)
	setResourceDataBoolWithDiagnostic(d, "mask_attribute_values", rv.MaskAttributeValues, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "ldap_type", rv.LdapType, &diags)
	setResourceDataBoolWithDiagnostic(d, "bind_anonymously", rv.BindAnonymously, &diags)
	setResourceDataStringWithDiagnostic(d, "user_dn", rv.UserDN, &diags)
	setResourceDataStringWithDiagnostic(d, "password", rv.Password, &diags)
	setResourceDataStringWithDiagnostic(d, "encrypted_password", rv.EncryptedPassword, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_ssl", rv.UseSsl, &diags)
	setResourceDataBoolWithDiagnostic(d, "use_dns_srv_records", rv.UseDnsSrvRecords, &diags)
	setResourceDataBoolWithDiagnostic(d, "follow_ldap_referrals", rv.FollowLDAPReferrals, &diags)
	setResourceDataBoolWithDiagnostic(d, "test_on_borrow", rv.TestOnBorrow, &diags)
	setResourceDataBoolWithDiagnostic(d, "test_on_return", rv.TestOnReturn, &diags)
	setResourceDataBoolWithDiagnostic(d, "create_if_necessary", rv.CreateIfNecessary, &diags)
	setResourceDataBoolWithDiagnostic(d, "verify_host", rv.VerifyHost, &diags)
	setResourceDataIntWithDiagnostic(d, "min_connections", rv.MinConnections, &diags)
	setResourceDataIntWithDiagnostic(d, "max_connections", rv.MaxConnections, &diags)
	setResourceDataIntWithDiagnostic(d, "max_wait", rv.MaxWait, &diags)
	setResourceDataIntWithDiagnostic(d, "time_between_evictions", rv.TimeBetweenEvictions, &diags)
	setResourceDataIntWithDiagnostic(d, "read_timeout", rv.ReadTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "connection_timeout", rv.ConnectionTimeout, &diags)
	setResourceDataIntWithDiagnostic(d, "dns_ttl", rv.DnsTtl, &diags)
	setResourceDataStringWithDiagnostic(d, "ldap_dns_srv_prefix", rv.LdapDnsSrvPrefix, &diags)
	setResourceDataStringWithDiagnostic(d, "ldaps_dns_srv_prefix", rv.LdapsDnsSrvPrefix, &diags)

	if rv.HostnamesTags != nil && len(*rv.HostnamesTags) != 0 {
		if err := d.Set("hostnames_tags", flattenLdapTagConfigs(rv.HostnamesTags)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Hostnames != nil && len(*rv.Hostnames) > 0 {
		if err := d.Set("hostnames", flattenStringList(*rv.Hostnames)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.BinaryAttributes != nil && len(*rv.BinaryAttributes) > 0 {
		if err := d.Set("binary_attributes", flattenStringList(*rv.BinaryAttributes)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return diags
}

func resourcePingFederateLdapDataStoreResourceReadData(d *schema.ResourceData) *pf.LdapDataStore {
	ds := &pf.LdapDataStore{
		CreateIfNecessary: Bool(d.Get("create_if_necessary").(bool)),
		VerifyHost:        Bool(d.Get("verify_host").(bool)),
	}
	if v, ok := d.GetOk("data_store_id"); ok {
		ds.Id = String(v.(string))
	}
	if val, ok := d.GetOkExists("mask_attribute_values"); ok {
		ds.MaskAttributeValues = Bool(val.(bool))
	}
	if val, ok := d.GetOk("hostnames_tags"); ok {
		ds.HostnamesTags = expandLdapTagConfigList(val.([]interface{}))
	}
	if val, ok := d.GetOk("hostnames"); ok {
		strs := expandStringList(val.(*schema.Set).List())
		ds.Hostnames = &strs
	}
	if val, ok := d.GetOk("name"); ok {
		ds.Name = String(val.(string))
	}
	if val, ok := d.GetOk("ldap_type"); ok {
		ds.LdapType = String(val.(string))
	}
	if val, ok := d.GetOkExists("bind_anonymously"); ok {
		ds.BindAnonymously = Bool(val.(bool))
	}
	if val, ok := d.GetOk("user_dn"); ok {
		ds.UserDN = String(val.(string))
	}
	if val, ok := d.GetOk("password"); ok {
		ds.Password = String(val.(string))
	}
	if val, ok := d.GetOk("encrypted_password"); ok {
		ds.EncryptedPassword = String(val.(string))
	}
	if val, ok := d.GetOkExists("use_ssl"); ok {
		ds.UseSsl = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("use_dns_srv_records"); ok {
		ds.UseDnsSrvRecords = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("follow_ldap_referrals"); ok {
		ds.FollowLDAPReferrals = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("test_on_borrow"); ok {
		ds.TestOnBorrow = Bool(val.(bool))
	}
	if val, ok := d.GetOkExists("test_on_return"); ok {
		ds.TestOnReturn = Bool(val.(bool))
	}
	if val, ok := d.GetOk("min_connections"); ok {
		ds.MinConnections = Int(val.(int))
	}
	if val, ok := d.GetOk("max_connections"); ok {
		ds.MaxConnections = Int(val.(int))
	}
	if val, ok := d.GetOk("max_wait"); ok {
		ds.MaxWait = Int(val.(int))
	}
	if val, ok := d.GetOk("time_between_evictions"); ok {
		ds.TimeBetweenEvictions = Int(val.(int))
	}
	if val, ok := d.GetOk("read_timeout"); ok {
		ds.ReadTimeout = Int(val.(int))
	}
	if val, ok := d.GetOk("connection_timeout"); ok {
		ds.ConnectionTimeout = Int(val.(int))
	}
	if val, ok := d.GetOk("dns_ttl"); ok {
		ds.DnsTtl = Int(val.(int))
	}
	if val, ok := d.GetOk("ldap_dns_srv_prefix"); ok {
		ds.LdapDnsSrvPrefix = String(val.(string))
	}
	if val, ok := d.GetOk("ldaps_dns_srv_prefix"); ok {
		ds.LdapsDnsSrvPrefix = String(val.(string))
	}
	if val, ok := d.GetOk("binary_attributes"); ok {
		strs := expandStringList(val.(*schema.Set).List())
		ds.BinaryAttributes = &strs
	}
	ds.Type = String("LDAP")

	return ds
}
