package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateDataStoreResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateDataStoreResourceCreate,
		Read:   resourcePingFederateDataStoreResourceRead,
		Update: resourcePingFederateDataStoreResourceUpdate,
		Delete: resourcePingFederateDataStoreResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingFederateDataStoreResourceSchema(),
	}
}

func resourcePingFederateDataStoreResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
		"jdbc_data_store": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{"ldap_data_store", "custom_data_store"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
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
						//Default: 10,
					},
					"max_pool_size": {
						Type:     schema.TypeInt,
						Optional: true,
						//Default: 100,
					},
					"blocking_timeout": {
						Type:     schema.TypeInt,
						Optional: true,
						//Default: 5000,
					},
					"idle_timeout": {
						Type:     schema.TypeInt,
						Optional: true,
						//Default: 5,
					},
				},
			},
		},
		"ldap_data_store": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{"jdbc_data_store", "custom_data_store"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mask_attribute_values": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"hostnames_tags": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"hostnames": {
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
					},
					"verify_host": {
						Type:     schema.TypeBool,
						Optional: true,
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
					},
					"time_between_evictions": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"read_timeout": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"connection_timeout": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"dns_ttl": {
						Type:     schema.TypeInt,
						Optional: true,
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
				},
			},
		},
		"custom_data_store": {
			Type:          schema.TypeList,
			Optional:      true,
			MaxItems:      1,
			ConflictsWith: []string{"ldap_data_store", "jdbc_data_store"},
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"mask_attribute_values": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"name": {
						Type:     schema.TypeString,
						Required: true,
					},
					"plugin_descriptor_ref": resourceLinkSchema(),
					"parent_ref":            resourceLinkSchema(),
					"configuration":         resourcePluginConfiguration(),
				},
			},
		},
	}
}

func resourcePingFederateDataStoreResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).DataStores
	ds := *resourcePingFederateDataStoreResourceReadData(d)
	var result pf.DataStore
	result.Type = ds.Type
	switch *ds.Type {
	case "LDAP":
		input := pf.CreateLdapDataStoreInput{Body: ds.LdapDataStore, BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool))}
		store, _, err := svc.CreateLdapDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.LdapDataStore = *store
		result.Id = store.Id
		break
	case "JDBC":
		input := pf.CreateJdbcDataStoreInput{Body: ds.JdbcDataStore, BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool))}
		store, _, err := svc.CreateJdbcDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.JdbcDataStore = *store
		result.Id = store.Id
		break
	case "CUSTOM":
		input := pf.CreateCustomDataStoreInput{Body: ds.CustomDataStore, BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool))}
		store, _, err := svc.CreateCustomDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.CustomDataStore = *store
		result.Id = store.Id
		break
	}
	d.SetId(*result.Id)
	return resourcePingFederateDataStoreResourceReadResult(d, &result)
}

func resourcePingFederateDataStoreResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).DataStores
	input := pf.GetDataStoreInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetDataStore(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateDataStoreResourceReadResult(d, result)
}

func resourcePingFederateDataStoreResourceUpdate(d *schema.ResourceData, m interface{}) (err error) {
	svc := m.(*pf.PfClient).DataStores
	ds := *resourcePingFederateDataStoreResourceReadData(d)
	var result pf.DataStore
	result.Type = ds.Type
	switch *ds.Type {
	case "LDAP":
		input := pf.UpdateLdapDataStoreInput{Id: d.Id(), Body: ds.LdapDataStore}
		store, _, err := svc.UpdateLdapDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.LdapDataStore = *store
		break
	case "JDBC":
		input := pf.UpdateJdbcDataStoreInput{Id: d.Id(), Body: ds.JdbcDataStore}
		store, _, err := svc.UpdateJdbcDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.JdbcDataStore = *store
		break
	case "CUSTOM":
		input := pf.UpdateCustomDataStoreInput{Id: d.Id(), Body: ds.CustomDataStore}
		store, _, err := svc.UpdateCustomDataStore(&input)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		result.CustomDataStore = *store
		break
	}

	return resourcePingFederateDataStoreResourceReadResult(d, &result)
}

func resourcePingFederateDataStoreResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).DataStores
	input := pf.DeleteDataStoreInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDataStore(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateDataStoreResourceReadResult(d *schema.ResourceData, rv *pf.DataStore) (err error) {
	switch *rv.Type {
	case "LDAP":
		if rv.Name != nil {
			rv.LdapDataStore.Name = rv.Name
		}
		if err := d.Set("ldap_data_store", flattenLdapDataStore(&rv.LdapDataStore)); err != nil {
			return err
		}
		break
	case "JDBC":
		if rv.Name != nil {
			rv.JdbcDataStore.Name = rv.Name
		}
		if err := d.Set("jdbc_data_store", flattenJdbcDataStore(&rv.JdbcDataStore)); err != nil {
			return err
		}
		break
	case "CUSTOM":
		if rv.Name != nil {
			rv.CustomDataStore.Name = rv.Name
		}
		//input.Body.CustomDataStore = &ds.CustomDataStore
		break
	}

	return nil
}

func resourcePingFederateDataStoreResourceReadData(d *schema.ResourceData) *pf.DataStore {
	ds := &pf.DataStore{}
	if v, ok := d.GetOkExists("jdbc_data_store"); ok && len(v.([]interface{})) != 0 {
		ds.Type = String("JDBC")
		ds.JdbcDataStore = *expandJdbcDataStore(v.([]interface{}))
	}
	if v, ok := d.GetOkExists("ldap_data_store"); ok && len(v.([]interface{})) != 0 {
		ds.Type = String("LDAP")
		ds.LdapDataStore = *expandLdapDataStore(v.([]interface{}))
	}
	//if v, ok := d.GetOkExists("custom_data_store"); ok && len(v.([]interface{})) != 0 {
	//	ds.Type = String("CUSTOM")
	//	ds.CustomDataStore = *expandCustomDataStore(v.([]interface{}))
	//}

	return ds
}
