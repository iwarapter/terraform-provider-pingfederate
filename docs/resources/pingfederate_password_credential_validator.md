# Resource: pingfederate_password_credential_validator

Provides a Password Credential Validator.

## Example Usage
```hcl
resource "pingfederate_password_credential_validator" "demo" {
  name = "demo"
  plugin_descriptor_ref {
    id = "org.sourceid.saml20.domain.SimpleUsernamePasswordCredentialValidator"
  }

  configuration {
    tables {
      name = "Users"
      rows {
        fields {
          name  = "Username"
          value = "bob"
        }

        sensitive_fields {
          name  = "Password"
          value = "demo2"
        }

        sensitive_fields {
          name  = "Confirm Password"
          value = "demo2"
        }

        fields {
          name  = "Relax Password Requirements"
          value = "true"
        }
      }
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- `name` - (Required) The plugin instance name. The name cannot be modified once the instance is created. Note: Ignored when specifying a connection's adapter override.
- `plugin_descriptor_ref` - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created. Note: Ignored when specifying a connection's adapter override.
- `parent_ref` - (Optional) The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances. Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)
- `configuration` - (Required) Plugin instance configuration.
- `attribute_contract` - (Optional) The list of attributes that the password credential validator provides.

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

The `attribute_contract` block - A set of attributes exposed by an Access Token Manager.

- [`extended_attributes`](#extended_attributes) - (Optional) A list of additional attributes that can be returned by the password credential validator. The extended attributes are only used if the adapter supports them.
- [`inherited`](#inherited) - (Optional) Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the plugin instance.

## Import

Password Credential Validator can be imported using the id, e.g.

```
terraform import pingfederate_password_credential_validator.demo 123
```
