# Resource: pingfederate_authentication_selector

Provides an authentication selector.

## Example Usage
```hcl
resource "pingfederate_authentication_selector" "example" {
  name = "example"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.selectors.cidr.CIDRAdapterSelector"
  }

  configuration {
    fields {
      name = "Result Attribute Name"
      value = ""
    }
    tables {
      name = "Networks"
      rows {
        fields {
          name  = "Network Range (CIDR notation)"
          value = "127.0.0.1/32"
        }
      }
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The plugin instance name. The name cannot be modified once the instance is created.

### Plugin Descriptor Ref

The `plugin_descriptor_ref` block allows you to specify the `id` and query the attribute `location`

- [`id`](#plugin_descriptor_ref-id) - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.

### Attribute Contract

the `attribute_contract` block allows you to specify the list of attributes that the Authentication Selector provides.

- [`extended_fields`](#attribute_contract-extended_fields) -  A list of additional attributes that can be returned by the Authentication Selector. The extended attributes are only used if the Authentication Selector supports them.

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


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the selector.

## Import

Authentication API Applications can be imported using the id, e.g.

```
terraform import pingfederate_authentication_selector.example 123
```
