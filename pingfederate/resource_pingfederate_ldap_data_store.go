package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateLdapDataStoreResource() *schema.Resource {
	return &schema.Resource{
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
		"mask_attribute_values": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"hostnames_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"hostnames": requiredSetOfString(),
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
		"hostnames": setOfString(),
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ldap_type": {
			Type:     schema.TypeString,
			Required: true,
			//TODO Add validator
			//['ACTIVE_DIRECTORY' or 'ORACLE_DIRECTORY_SERVER' or 'ORACLE_UNIFIED_DIRECTORY' or 'UNBOUNDID_DS' or 'PING_DIRECTORY' or 'GENERIC']: A type that allows PingFederate to configure many provisioning settings automatically. The 'UNBOUNDID_DS' type has been deprecated, please use the 'PING_DIRECTORY' type instead.
		},
		"bind_anonymously": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"user_dn": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"encrypted_password": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_ssl": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"use_dns_srv_records": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"follow_ldap_referrals": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"test_on_borrow": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"test_on_return": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"create_if_necessary": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"verify_host": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"min_connections": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_connections": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"max_wait": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  -1,
		},
		"time_between_evictions": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60000,
		},
		"read_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3000,
		},
		"connection_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3000,
		},
		"dns_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60000,
		},
		"ldap_dns_srv_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "_ldap._tcp",
		},
		"ldaps_dns_srv_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "_ldaps._tcp",
		},
		"binary_attributes": setOfString(),
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
