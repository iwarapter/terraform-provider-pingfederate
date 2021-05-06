# Resource: pingfederate_custom_data_store

Provides a Custom Data Store.

## Example Usage
```hcl
resource "pingfederate_custom_data_store" "example" {
  name = "example"
  plugin_descriptor_ref {
    id = "com.pingidentity.pf.datastore.other.RestDataSourceDriver"
  }
  configuration {
    tables {
      name = "Base URLs and Tags"
      rows {
        fields {
          name  = "Base URL"
          value = "https://example.com"
        }
        fields {
          name  = "Tags"
          value = ""
        }
        default_row = true
      }
    }
    tables {
      name = "HTTP Request Headers"
    }
    tables {
      name = "Attributes"
      rows {
        fields {
          name  = "Local Attribute"
          value = "foo"
        }
        fields {
          name  = "JSON Response Attribute Path"
          value = "/bar"
        }
      }
    }
    fields {
      name  = "Authentication Method"
      value = "None"
    }
    fields {
      name  = "Username"
      value = ""
    }
    fields {
      name  = "Password"
      value = ""
    }
    fields {
      name  = "OAuth Token Endpoint"
      value = ""
    }
    fields {
      name  = "OAuth Scope"
      value = ""
    }
    fields {
      name  = "Client ID"
      value = ""
    }
    fields {
      name  = "Client Secret"
      value = ""
    }
    fields {
      name  = "Enable HTTPS Hostname Verification"
      value = "true"
    }
    fields {
      name  = "Read Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Connection Timeout (ms)"
      value = "10000"
    }
    fields {
      name  = "Max Payload Size (KB)"
      value = "1024"
    }
    fields {
      name  = "Retry Request"
      value = "true"
    }
    fields {
      name  = "Maximum Retries Limit"
      value = "5"
    }
    fields {
      name  = "Retry Error Codes"
      value = "429"
    }
    fields {
      name = "Test Connection URL"
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The plugin instance name. The name cannot be modified once the instance is created.
- `allow_multi_value_attributes` - (Optional) Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute.

### Plugin Descriptor Ref

The `plugin_descriptor_ref` block allows you to specify the `id` and query the attribute `location`

- [`id`](#plugin_descriptor_ref-id) - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.

### Parent Ref

The `parent_ref` block allows you to specify the `id` and query the attribute `location`

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

- [`id`](#id) - The ID of the custom data store.


## Import

Custom Data Stores can be imported using the id, e.g.

```
terraform import pingfederate_custom_data_store.example 123
```
