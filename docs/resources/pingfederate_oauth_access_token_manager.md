# Resource: pingfederate_oauth_access_token_manager

Provides a OAuth Access Token Manager.

## Example Usage
```hcl
resource "pingfederate_oauth_access_token_manager" "example" {
  instance_id = "example"
  name        = "example"
  plugin_descriptor_ref {
    id = "org.sourceid.oauth20.token.plugin.impl.ReferenceBearerAccessTokenManagementPlugin"
  }
  configuration {
    fields {
      name  = "Token Length"
      value = "28"
    }
    fields {
      name  = "Token Lifetime"
      value = "120"
    }
    fields {
      name  = "Lifetime Extension Policy"
      value = "ALL"
    }
    fields {
      name  = "Maximum Token Lifetime"
      value = ""
    }
    fields {
      name  = "Lifetime Extension Threshold Percentage"
      value = "30"
    }
    fields {
      name  = "Mode for Synchronous RPC"
      value = "3"
    }
    fields {
      name  = "RPC Timeout"
      value = "500"
    }
    fields {
      name  = "Expand Scope Groups"
      value = "false"
    }
  }

  attribute_contract {
    extended_attributes = ["sub"]
  }
}
```

## Argument Attributes

The following arguments are supported:

- `name` - (Required) The plugin instance name. The name cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
- `publisher_id` - (Required) The ID of the plugin instance. The ID cannot be modified once the instance is created.
- `plugin_descriptor_ref` - (Required) Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.
- `configuration` - (Required) Plugin instance configuration.
- `attribute_contract` - (Optional) The list of attributes that will be added to an access token.
- `parent_ref` - (Optional) The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances.<br>Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)
- `selection_settings` - (Optional) Settings which determine how this token manager can be selected for use by an OAuth request.
- `session_validation_settings` - (Optional) Settings which determine how the user session is associated with the access token.
- `access_control_settings` - (Optional) Settings which determine which clients may access this token manager.

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

The `attribute_contract` block - A set of attributes exposed by an Access Token Manager.

- `core_attributes` - (Optional) A list of core token attributes that are associated with the access token management plugin type. This field is read-only and is ignored on POST/PUT.

- `default_subject_attribute` - (Optional) Default subject attribute to use for audit logging when validating the access token. Blank value means to use USER_KEY attribute value after grant lookup.

- `extended_attributes` - (Optional) A list of additional token attributes that are associated with this access token management plugin instance.

- `inherited` - (Optional) Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.

### Selection Settings

The `selection_settings` block - Selection settings for an access token management plugin instance.

- `inherited` - (Optional) If this token manager has a parent, this flag determines whether selection settings, such as resource URI's, are inherited from the parent. When set to true, the other fields in this model become read-only. The default value is false.

- `resource_uris` - (Optional) The list of base resource URI's which map to this token manager. A resource URI, specified via the 'aud' parameter, can be used to select a specific token manager for an OAuth request.

### Session Validation Settings

The `session_validation_settings` block - Session validation settings for an access token management plugin instance.

- `check_session_revocation_status` - (Optional) Check the session revocation status when validating the access token.

- `check_valid_authn_session` - (Optional) Check for a valid authentication session when validating the access token.

- `include_session_id` - (Optional) Include the session identifier in the access token. Note that if any of the session validation features is enabled, the session identifier will already be included in the access tokens.

- `inherited` - (Optional) If this token manager has a parent, this flag determines whether session validation settings, such as checkValidAuthnSession, are inherited from the parent. When set to true, the other fields in this model become read-only. The default value is false.

- `update_authn_session_activity` - (Optional) Update authentication session activity when validating the access token.

### Access Control Settings

The `access_control_settings` block - Access control settings for an access token management plugin instance.

- `allowed_clients` - (Optional) If 'restrictClients' is true, this field defines the list of OAuth clients that are allowed to access the token manager.

- `inherited` - (Optional) If this token manager has a parent, this flag determines whether access control settings are inherited from the parent. When set to true, the other fields in this model become read-only. The default value is false.

- `restrict_clients` - (Optional) Determines whether access to this token manager is restricted to specific OAuth clients. If false, the 'allowedClients' field is ignored. The default value is false.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the plugin instance (`instance_id`).

## Import

OAuth Access Token Managers can be imported using the id, e.g.

```
terraform import pingfederate_oauth_access_token_manager.demo 123
```
