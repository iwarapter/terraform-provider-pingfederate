# Resource: pingfederate_oauth_openid_connect_policy

Provides an OpenId Connect Policy.

## Example Usage
```hcl
resource "pingfederate_oauth_openid_connect_policy" "example" {
  policy_id = "example"
  name      = "example"
  access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.example.id
  }
  attribute_contract {
    extended_attributes {
      name = "name"
      include_in_id_token = true
    }
  }
  attribute_mapping {
    attribute_contract_fulfillment {
      key_name = "sub"
      source {
        type = "NO_MAPPING"
      }
    }
    attribute_contract_fulfillment {
      key_name = "name"
      source {
        type = "NO_MAPPING"
      }
    }
  }

  scope_attribute_mappings {
    key_name = "address"
    values = ["name"]
  }
}
```

## Argument Attributes

The following arguments are supported:

- `access_token_manager_ref` - (Required) The access token manager associated with this Open ID Connect policy.

- `attribute_contract` - (Required) The list of attributes that will be returned to OAuth clients in response to requests received at the PingFederate UserInfo endpoint.

- `attribute_mapping` - (Required) The attributes mapping from attribute sources to attribute targets.

- `id_token_lifetime` - (Optional) The ID Token Lifetime, in minutes. The default value is 5.

- `include_s_hash_in_id_token` - (Optional) Determines whether the State Hash should be included in the ID token.

- `include_sri_in_id_token` - (Optional) Determines whether a Session Reference Identifier is included in the ID token.

- `include_user_info_in_id_token` - (Optional) Determines whether the User Info is always included in the ID token.

- `name` - (Required) The name used for display in UI screens.

- `return_id_token_on_refresh_grant` - (Optional) Determines whether an ID Token should be returned when refresh grant is requested or not.

- `scope_attribute_mappings` - (Optional) The attribute scope mappings from scopes to attribute names.


### Attribute Contract

The `attribute_contract` block - A set of attributes that will be returned to OAuth clients in response to requests received at the PingFederate UserInfo endpoint.

- `extended_attributes` - (Optional) A list of additional attributes.


### Extended Attributes

The `extended_attributes` block - An attribute for the OpenID Connect returned to OAuth clients.

- `name` - (Required) The name of this attribute.

- `include_in_id_token` - (Optional) Attribute is included in the ID Token.

- `include_in_user_info` - (Optional) Attribute is included in the User Info


### Attribute Mapping

The `attribute_mapping` block - A list of mappings from attribute sources to attribute targets.

- `attribute_contract_fulfillment` - (Required) A list of mappings from attribute names to their fulfillment values.

- `ldap_attribute_source` - (Optional) A list of ldap configured data stores to look up attributes from.

- `jdbc_attribute_source` - (Optional) A list of jdbc configured data stores to look up attributes from.

- `custom_attribute_source` - (Optional) A list of custom configured data stores to look up attributes from.

- `issuance_criteria` - (Optional) The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.

### Attribute Contract Fulfillment

The `attribute_contract_fulfillment` block - Defines how an attribute in an attribute contract should be populated.

- `key_name` - (Required) The attribute name mapping.

- `source` - (Required) The attribute value source.

- `value` - (Optional) The value for this attribute.

### Ldap Attribute Source

The `ldap_attribute_source` block - The configured settings used to look up attributes from a LDAP data store.

- `id` - (Optional) - The ID that defines this attribute source. Only alphanumeric characters allowed.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings

- `base_dn` - (Optional) The base DN to search from. If not specified, the search will start at the LDAP's root.

- `binary_attribute_settings` - (Optional) The advanced settings for binary LDAP attributes.

- `data_store_ref` - (Required) Reference to the associated data store.

- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings

- `member_of_nested_group` - (Optional) Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.

- `search_filter` - (Required) The LDAP filter that will be used to lookup the objects from the directory.

- `search_scope` - (Required) Determines the node depth of the query.

### Jdbc Attribute Source

The `jdbc_attribute_source` block - The configured settings used to look up attributes from a JDBC data store.

- `id` - (Optional) - The ID that defines this attribute source. Only alphanumeric characters allowed.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings

- `data_store_ref` - (Required) Reference to the associated data store.

- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings

- `filter` - (Required) The JDBC WHERE clause used to query your data store to locate a user record.

- `schema` - (Optional) Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.

- `table` - (Required) The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.

### Custom Attribute Source

The `custom_attribute_source` block - The configured settings used to look up attributes from a custom data store.

- `id` - (Optional) - The ID that defines this attribute source. Only alphanumeric characters allowed.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings

- `data_store_ref` - (Required) Reference to the associated data store.

- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings

- `filter_fields` - (Optional) The list of fields that can be used to filter a request to the custom data store.

### Issuance Criteria

The `issuance_criteria` block - A list of criteria that determines whether a transaction (usually a SSO transaction) is continued. All criteria must pass in order for the transaction to continue.

- `conditional_criteria` - (Optional) A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue.

- `expression_criteria` - (Optional) A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue.

### Conditional Issuance Criteria

The `conditional_criteria` block - An issuance criterion that checks a source attribute against a particular condition and the expected value. If the condition is true then this issuance criterion passes, otherwise the criterion fails.

- `attribute_name` - (Required) The name of the attribute to use in this issuance criterion.

- `condition` - (Required) The condition that will be applied to the source attribute's value and the expected value.

- `error_result` - (Optional) The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.

- `source` - (Required) The source of the attribute.

- `value` - (Required) The expected value of this issuance criterion.

### Source Type ID Key

The `source` block - A key that is meant to reference a source from which an attribute can be retrieved.
This model is usually paired with a value which, depending on the SourceType, can be a hardcoded value or a reference to an attribute name specific to that SourceType.
Not all values are applicable - a validation error will be returned for incorrect values.

- `type` - (Required) The source type of this key.

For each SourceType, the value should be:

- `ACCOUNT_LINK` - If account linking was enabled for the browser SSO, the value must be 'Local User ID', unless it has been overridden in PingFederate's server configuration.
- `ADAPTER` - The value is one of the attributes of the IdP Adapter.
- `ASSERTION` - The value is one of the attributes coming from the SAML assertion.
- `AUTHENTICATION_POLICY_CONTRACT` - The value is one of the attributes coming from an authentication policy contract.
- `LOCAL_IDENTITY_PROFILE` - The value is one of the fields coming from a local identity profile.
- `CONTEXT` - The value must be one of the following ['TargetResource' or 'OAuthScopes' or 'ClientId' or 'AuthenticationCtx' or 'ClientIp' or 'Locale' or 'StsBasicAuthUsername' or 'StsSSLClientCertSubjectDN' or 'StsSSLClientCertChain' or 'VirtualServerId' or 'AuthenticatingAuthority' or 'DefaultPersistentGrantLifetime']
- `CLAIMS` - Attributes provided by the OIDC Provider.
- `CUSTOM_DATA_STORE` - The value is one of the attributes returned by this custom data store.
- `EXPRESSION` - The value is an OGNL expression.
- `EXTENDED_CLIENT_METADATA` - The value is from an OAuth extended client metadata parameter. This source type is deprecated and has been replaced by EXTENDED_PROPERTIES.
- `EXTENDED_PROPERTIES` - The value is from an OAuth Client's extended property.
- `IDP_CONNECTION` - The value is one of the attributes passed in by the IdP connection.
- `JDBC_DATA_STORE` - The value is one of the column names returned from the JDBC attribute source.
- `LDAP_DATA_STORE` - The value is one of the LDAP attributes supported by your LDAP data store.
- `MAPPED_ATTRIBUTES` - The value is the name of one of the mapped attributes that is defined in the associated attribute mapping.
- `OAUTH_PERSISTENT_GRANT` - The value is one of the attributes from the persistent grant.
- `PASSWORD_CREDENTIAL_VALIDATOR` - The value is one of the attributes of the PCV.
- `NO_MAPPING` - A placeholder value to indicate that an attribute currently has no mapped source.TEXT - A hardcoded value that is used to populate the corresponding attribute.
- `TOKEN` - The value is one of the token attributes.
- `REQUEST` - The value is from the request context such as the CIBA identity hint contract or the request contract for Ws-Trust.
- `TRACKED_HTTP_PARAMS` - The value is from the original request parameters.
- `SUBJECT_TOKEN` - The value is one of the OAuth 2.0 Token exchange subject_token attributes.
- `ACTOR_TOKEN` - The value is one of the OAuth 2.0 Token exchange actor_token attributes.
- `TOKEN_EXCHANGE_PROCESSOR_POLICY` - The value is one of the attributes coming from a Token Exchange Processor policy.

### Expression Issuance Criteria

The `expression_criteria` block - An issuance criterion that uses a Boolean return value from an OGNL expression to determine whether or not it passes.

- `error_result` - (Optional) The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.

- `expression` - (Required) The OGNL expression to evaluate.

### Scope Attribute Mappings

The `scope_attribute_mappings` block - Parameter Values.

- `key_name` - (Required) The mapping key.

- `values` - (Optional) A List of values


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the OpenId Connect Policy.

## Import

OpenId Connect Policy can be imported using the id, e.g.

```
terraform import pingfederate_oauth_openid_connect_policy.example 123
```
