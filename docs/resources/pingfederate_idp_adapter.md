# Resource: pingfederate_idp_adapter

Provides an Idp Adapter.

## Example Usage
```hcl
resource "pingfederate_idp_adapter" "example" {
  name = "example"
  plugin_descriptor_ref {
    id = "com.pingidentity.adapters.httpbasic.idp.HttpBasicIdpAuthnAdapter"
  }

  configuration {
    tables {
      name = "Credential Validators"
      rows {
        fields {
          name = "Password Credential Validator Instance"
          value = pingfederate_password_credential_validator.demo.name
        }
      }
    }
    fields {
      name = "Realm"
      value = "foo"
    }

    fields {
      name = "Challenge Retries"
      value = "3"
    }
  }

  attribute_contract {
    core_attributes {
      name = "username"
      pseudonym = true
    }
    extended_attributes {
      name = "sub"
    }
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "ADAPTER"
      }
      value = "sub"
    }
    attribute_contract_fulfillment {
      key_name = "username"
      source {
        type = "ADAPTER"
      }
      value = "username"
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- `name` - (Required) The plugin instance name. The name cannot be modified once the instance is created.
- `plugin_descriptor_ref` - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.
- `attribute_contract` - (Optional) The list of attributes that the IdP adapter provides.
- `attribute_mapping` - (Optional) The attributes mapping from attribute sources to attribute targets.
- `authn_ctx_class_ref` - (Optional) The fixed value that indicates how the user was authenticated.
- `configuration` - (Required) Plugin instance configuration.
- `parent_ref` - (Optional) The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances.<br>Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)


### Parent Ref

The `parent_ref` block allows you to specify the `id` and query the attribute `location`

- [`id`](#parent_ref-id) - The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances.

### Plugin Descriptor Ref

The `plugin_descriptor_ref` block allows you to specify the `id` and query the attribute `location`

- [`id`](#plugin_descriptor_ref-id) - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.

### Configuration

The `configuration` block allows you to specify the `fields` and `tables` for the plugin configuration

- [`fields`](#configuration-field) - List of configuration fields blocks.
- [`sensitive_fields`](#configuration-sensitive_fields) - List of sensitive configuration fields blocks which are masked by terraform.
- [`tables`](#configuration-tables) - List of configuration tables blocks.

### Configuration Fields

The `fields`/`sensitive_fields` block the configuration fields for a root configuration or configuration row.

- [`name`](#configuration_field-name) - The name of the configuration field.
- [`value`](#configuration_field-value) - The value for the configuration field.
- [`inherited`](#configuration_field-inherited) - Whether this field is inherited from its parent instance. If true, the value value properties become read-only. The default value is false.

### Configuration Tables

The `tables` block for the root configuration.

- [`name`](#configuration_table-name) - The name of the table.
- [`rows`](#configuration_table-rows) - List of table rows.
- [`inherited`](#configuration_table-inherited) - Whether this table is inherited from its parent instance. If true, the rows become read-only. The default value is false.

### Configuration Rows

The `rows` block for each of the `tables` blocks.

- [`fields`](#configuration-field) - List of configuration fields blocks.
- [`sensitive_fields`](#configuration-sensitive_fields) - List of sensitive configuration fields blocks which are masked by terraform.

### Attribute Contract

The `attribute_contract` block - A set of attributes exposed by an IdP adapter.

- `core_attributes` - (Required) A list of IdP adapter attributes that correspond to the attributes exposed by the IdP adapter type.
- `extended_attributes` - (Optional) A list of additional attributes that can be returned by the IdP adapter. The extended attributes are only used if the adapter supports them.
- `inherited` - (Optional) Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.
- `mask_ognl_values` - (Optional) Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false.
- `unique_user_key_attribute` - (Optional) The attribute to use for uniquely identify a user's authentication sessions.

### Core / Extended Attributes

The `core_attributes` / `extended_attributes` block - An attribute for the IdP adapter attribute contract.

- `name` - (Required) The name of this attribute.
- `masked` - (Optional) Specifies whether this attribute is masked in PingFederate logs. Defaults to false.
- `pseudonym` - (Optional) Specifies whether this attribute is used to construct a pseudonym for the SP. Defaults to false.


### Contract Mapping

The `contract_mapping` block -

- `attribute_contract_fulfillment` - (Required) A list of mappings from attribute names to their fulfillment values.
- `attribute_sources` - (Optional) A list of configured data stores to look up attributes from.
- `inherited` - (Optional) Whether this attribute mapping is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.
- `issuance_criteria` - (Optional) The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.

### JDBC Attribute Source

The `jdbc_attribute_source` block - The configured settings used to look up attributes from a JDBC data store.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings
- `data_store_ref` - (Required) Reference to the associated data store.
- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings
- `filter` - (Required) The JDBC WHERE clause used to query your data store to locate a user record.
- `schema` - (Optional) Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.
- `table` - (Required) The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.
- `type` - (Required) The data store type of this attribute source.

### LDAP Attribute Source

The `ldap_attribute_source` block - The configured settings used to look up attributes from a LDAP data store.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings
- `base_dn` - (Optional) The base DN to search from. If not specified, the search will start at the LDAP's root.
- `binary_attribute_settings` - (Optional) The advanced settings for binary LDAP attributes.
- `data_store_ref` - (Required) Reference to the associated data store.
- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings
- `member_of_nested_group` - (Optional) Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.
- `search_filter` - (Required) The LDAP filter that will be used to lookup the objects from the directory.
- `search_scope` - (Required) Determines the node depth of the query.
- `type` - (Required) The data store type of this attribute source.

### Attribute Fulfillment

The `attribute_contract_fulfillment` block - Defines how an attribute in an attribute contract should be populated.

- `key_name` - (Required) The attribute key mapping.
- `source` - (Required) The attribute value source.
- `value` - (Required) The value for this attribute.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the plugin instance (`instance_id`).

## Import

IDP Adapters can be imported using the id, e.g.

```
terraform import pingfederate_idp_adapter.demo 123
```
