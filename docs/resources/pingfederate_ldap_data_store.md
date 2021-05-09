# Resource: pingfederate_ldap_data_store

Provides a Ldap Data Store.

## Example Usage
```hcl
resource "pingfederate_ldap_data_store" "demo" {
	name             = "terraform_ldap"
	ldap_type        = "PING_DIRECTORY"
	hostnames        = ["host.docker.internal:1389"]
	bind_anonymously = true
	min_connections  = 1
	max_connections  = 10
}
```

## Argument Attributes

The following arguments are supported:

- [`data_store_id`](#data_store_id) - (Optional)  The Data Store ID. This property is system-assigned if not specified.
- `binary_attributes` - (Optional) The list of LDAP attributes to be handled as binary data.
- `bind_anonymously` - (Optional) Whether username and password are required. The default value is false.
- `connection_timeout` - (Optional) The maximum number of milliseconds that a connection attempt should be allowed to continue before returning an error. A value of -1 causes the pool to wait indefinitely. Omitting this attribute will set the value to the default value.
- `create_if_necessary` - (Optional) Indicates whether temporary connections can be created when the Maximum Connections threshold is reached.
- `dns_ttl` - (Optional) The maximum time in milliseconds that DNS information are cached. Omitting this attribute will set the value to the default value.
- `encrypted_password` - (Optional) The encrypted password credential required to access the data store.  If you do not want to update the stored value, this attribute should be passed back unchanged.
- `follow_ldap_referrals` - (Optional) Follow LDAP Referrals in the domain tree. The default value is false. This property does not apply to PingDirectory as this functionality is configured in PingDirectory.
- `hostnames` - (Optional) The default LDAP host names. This field is required if no mapping for host names and tags are specified.
- `hostnames_tags` - (Optional) The set of host names and associated tags for this LDAP data store.
- `ldap_dns_srv_prefix` - (Optional) The prefix value used to discover LDAP DNS SRV record. Omitting this attribute will set the value to the default value.
- `ldap_type` - (Required) A type that allows PingFederate to configure many provisioning settings automatically. The 'UNBOUNDID_DS' type has been deprecated, please use the 'PING_DIRECTORY' type instead.
- `ldaps_dns_srv_prefix` - (Optional) The prefix value used to discover LDAPs DNS SRV record. Omitting this attribute will set the value to the default value.
- `mask_attribute_values` - (Optional) Whether attribute values should be masked in the log.
- `max_connections` - (Optional) The largest number of active connections that can remain in each pool without releasing extra ones. Omitting this attribute will set the value to the default value.
- `max_wait` - (Optional) The maximum number of milliseconds the pool waits for a connection to become available when trying to obtain a connection from the pool. Omitting this attribute or setting a value of -1 causes the pool not to wait at all and to either create a new connection or produce an error (when no connections are available).
- `min_connections` - (Optional) The smallest number of connections that can remain in each pool, without creating extra ones. Omitting this attribute will set the value to the default value.
- `name` - (Optional) The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the hostname(s) and the principal.
- `password` - (Optional) The password credential required to access the data store. GETs will not return this attribute. To update this field, specify the new value in this attribute.
- `read_timeout` - (Optional) The maximum number of milliseconds a connection waits for a response to be returned before producing an error. A value of -1 causes the connection to wait indefinitely. Omitting this attribute will set the value to the default value.
- `test_on_borrow` - (Optional) Indicates whether objects are validated before being borrowed from the pool.
- `test_on_return` - (Optional) Indicates whether objects are validated before being returned to the pool.
- `time_between_evictions` - (Optional) The frequency, in milliseconds, that the evictor cleans up the connections in the pool. A value of -1 disables the evictor. Omitting this attribute will set the value to the default value.
- `type` - (Required) The data store type.
- `use_dns_srv_records` - (Optional) Use DNS SRV Records to discover LDAP server information. The default value is false.
- `use_ssl` - (Optional) Connects to the LDAP data store using secure SSL/TLS encryption (LDAPS). The default value is false.
- `user_dn` - (Optional) The username credential required to access the data store.
- `verify_host` - (Optional) Verifies that the presented server certificate includes the address to which the client intended to establish a connection. Omitting this attribute will set the value to true.

### Hostname Tags

The `hostnames_tags` block - An LDAP data store's hostnames and tags configuration. This is required if no default hostname is specified.

- `default_source` - (Optional) Whether this is the default connection. Defaults to false if not specified.

- `hostnames` - (Required) The LDAP host names.

- `tags` - (Optional) Tags associated with this data source.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the ldap data store.

## Import

Ldap Data Stores can be imported using the id, e.g.

```
terraform import pingfederate_ldap_data_store.example 123
```
