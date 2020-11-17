# Resource: pingfederate_notification_publisher

Provides a notification publisher.

## Example Usage
```hcl
resource "pingfederate_notification_publisher" "demo" {
  name         = "bar"
  publisher_id = "foo1"
  plugin_descriptor_ref {
    id = "com.pingidentity.email.SmtpNotificationPlugin"
  }

  configuration {
    fields {
      name  = "From Address"
      value = "help@foo.org"
    }
    fields {
      name  = "Email Server"
      value = "foo"
    }
    fields {
      name  = "SMTP Port"
      value = "25"
    }
    fields {
      name  = "Encryption Method"
      value = "NONE"
    }
    fields {
      name  = "SMTPS Port"
      value = "465"
    }
    fields {
      name  = "Verify Hostname"
      value = "true"
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
      name  = "Test Address"
      value = ""
    }
    fields {
      name  = "Connection Timeout"
      value = "30"
    }
    fields {
      name  = "Retry Attempt"
      value = "2"
    }
    fields {
      name  = "Retry Delay"
      value = "2"
    }
    fields {
      name  = "Enable SMTP Debugging Messages"
      value = "false"
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- [`name`](#name) - (Required) The plugin instance name. The name cannot be modified once the instance is created.

- [`publisher_id`](#publisher_id) - (Required) The ID of the plugin instance. The ID cannot be modified once the instance is created.

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

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the plugin instance (`publisher_id`).

## Import

Notification Publisher can be imported using the id, e.g.

```
terraform import pingfederate_notification_publisher.demo 123
```
