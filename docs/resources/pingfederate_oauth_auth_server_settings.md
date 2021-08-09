# Resource: pingfederate_oauth_auth_server_settings

Provides a OAuth Authorization Server Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.

## Example Usage

```hcl
resource "pingfederate_oauth_auth_server_settings" "settings" {
  scopes {
    name        = "address"
    description = "address"
  }

  scopes {
    name        = "mail"
    description = "mail"
  }

  scopes {
    name        = "openid"
    description = "openid"
  }

  scopes {
    name        = "phone"
    description = "phone"
  }

  scopes {
    name        = "profile"
    description = "profile"
  }

  scope_groups {
    name        = "group1"
    description = "group1"

    scopes = [
      "address",
      "mail",
      "phone",
      "openid",
      "profile",
    ]
  }

  persistent_grant_contract {
    extended_attributes = ["woot"]
  }

  allowed_origins = ["http://localhost"]

  default_scope_description  = ""
  authorization_code_timeout = 60
  authorization_code_entropy = 30
  refresh_token_length       = 42
  refresh_rolling_interval   = 0
}
```

## Argument Attributes

The following arguments are supported:

- `admin_web_service_pcv_ref` - (Optional) The password credential validator reference that is used for authenticating access to the OAuth Administrative Web Service.

- `allow_unidentified_client_extension_grants` - (Optional) Allow unidentified clients to request extension grants. The default value is false.

- `allow_unidentified_client_ro_creds` - (Optional) Allow unidentified clients to request resource owner password credentials grants. The default value is false.

- `allowed_origins` - (Optional) The list of allowed origins.

- `approved_scopes_attribute` - (Optional) Attribute from the external consent adapter's contract, intended for storing approved scopes returned by the external consent page.

- `atm_id_for_oauth_grant_management` - (Optional) The ID of the Access Token Manager used for OAuth enabled grant management.

- `authorization_code_entropy` - (Required) The authorization code entropy, in bytes.

- `authorization_code_timeout` - (Required) The authorization code timeout, in seconds.

- `bypass_activation_code_confirmation` - (Required) Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device.

- `bypass_authorization_for_approved_grants` - (Optional) Bypass authorization for previously approved persistent grants. The default value is false.

- `default_scope_description` - (Required) The default scope description.

- `device_polling_interval` - (Required) The amount of time client should wait between polling requests, in seconds.

- `exclusive_scope_groups` - (Optional) The list of exclusive scope groups.

- `exclusive_scopes` - (Optional) The list of exclusive scopes.

- `par_reference_length` - (Optional) The entropy of pushed authorization request references, in bytes. The default value is 24.

- `par_reference_timeout` - (Optional) The timeout, in seconds, of the pushed authorization request reference. The default value is 60.

- `par_status` - (Optional) The status of pushed authorization request support. The default value is ENABLED.

- `pending_authorization_timeout` - (Required) The 'device_code' and 'user_code' timeout, in seconds.

- `persistent_grant_contract` - (Optional) The persistent grant contract defines attributes that are associated with OAuth persistent grants.

- `persistent_grant_idle_timeout` - (Optional) The persistent grant idle timeout. The default value is 30 (days). -1 indicates an indefinite amount of time.

- `persistent_grant_idle_timeout_time_unit` - (Optional) The persistent grant idle timeout time unit.

- `persistent_grant_lifetime` - (Optional) The persistent grant lifetime. The default value is indefinite. -1 indicates an indefinite amount of time.

- `persistent_grant_lifetime_unit` - (Optional) The persistent grant lifetime unit.

- `persistent_grant_reuse_grant_types` - (Optional) The grant types that the OAuth AS can reuse rather than creating a new grant for each request.

- `refresh_rolling_interval` - (Required) The minimum interval to roll refresh tokens, in hours.

- `refresh_token_length` - (Required) The refresh token length in number of characters.

- `registered_authorization_path` - (Required) The Registered Authorization Path is concatenated to PingFederate base URL to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request. PingFederate listens to this path if specified

- `roll_refresh_token_values` - (Optional) The roll refresh token values default policy. The default value is true.

- `scope_for_oauth_grant_management` - (Optional) The OAuth scope to validate when accessing grant management service.

- `scope_groups` - (Optional) The list of common scope groups.

- `scopes` - (Optional) The list of common scopes.

- `token_endpoint_base_url` - (Optional) The token endpoint base URL used to validate the 'aud' claim during Private Key JWT Client Authentication.

- `track_user_sessions_for_logout` - (Optional) Determines whether user sessions are tracked for logout. If this property is not provided on a PUT, the setting is left unchanged.

- `user_authorization_consent_adapter` - (Optional) Adapter ID of the external consent adapter to be used for the consent page user interface.

- `user_authorization_consent_page_setting` - (Optional) User Authorization Consent Page setting to use PingFederate's internal consent page or an external system

- `user_authorization_url` - (Optional) The URL used to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request

### persistent_grant_contract

The `persistent_grant_contract` block -

- `core_attributes` - (Optional) This is a read-only list of persistent grant attributes and includes USER_KEY and USER_NAME. Changes to this field will be ignored.

- `extended_attributes` - (Optional) A list of additional attributes for the persistent grant contract.

### extended_attributes

The `extended_attributes` block - A persistent grant contract attribute.

- `name` - (Required) The name of this attribute.

### scopes / exclusive_scopes

The `scopes` / `exclusive_scopes` block - A scope name and its description.

- `description` - (Required) The description of the scope that appears when the user is prompted for authorization.

- `dynamic` - (Optional) True if the scope is dynamic. (Defaults to false)

- `name` - (Required) The name of the scope.

### scope_groups / exclusive_scope_groups

The `scope_group_entry` / `exclusive_scope_groups` block - A scope group name and its description.

- `description` - (Required) The description of the scope group.

- `name` - (Required) The name of the scope group.

- `scopes` - (Required) The set of scopes for this scope group.

## Import

-> The resource ID is fixed as `OauthAuthServerSettings` because this is a singleton resource.

OAuth Authorization Server Settings can be imported using the id, e.g.

```
terraform import pingfederate_oauth_auth_server_settings.demo OauthAuthServerSettings
```
