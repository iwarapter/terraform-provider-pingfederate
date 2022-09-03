# Resource: pingfederate_authentication_policies

Provides an Authentication Policies tree.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply clears all policies.

## Example Usage

```hcl-terraform
resource "pingfederate_authentication_policies" "demo" {
  fail_if_no_selection    = false
  tracked_http_parameters = ["foo"]
  default_authentication_sources {
    type = "IDP_ADAPTER"
    source_ref {
      id = pingfederate_idp_adapter.demo.id
    }
  }
  authn_selection_trees {
    name = "bar"
    root_node {
      action {
        type = "AUTHN_SOURCE"
        authentication_source {
          type = "IDP_ADAPTER"
          source_ref {
            id = pingfederate_idp_adapter.demo.id
          }
        }
      }
      children {
        action {
          type    = "RESTART"
          context = "Fail"
        }
      }
      children {
        action {
          type    = "DONE"
          context = "Success"
        }
      }
    }
  }
  authn_selection_trees {
    name = "foo"
    root_node {
      action {
        type = "AUTHN_SOURCE"
        authentication_source {
          type = "IDP_ADAPTER"
          source_ref {
            id = pingfederate_idp_adapter.demo.id
          }
        }
      }
      children {
        action {
          type    = "RESTART"
          context = "Fail"
        }
      }
      children {
        action {
          type    = "DONE"
          context = "Success"
        }
      }
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- `authn_selection_trees` - (Optional) The list of authentication policy trees.
- `default_authentication_sources` - (Optional) The default authentication sources.
- `fail_if_no_selection` - (Optional) Fail if policy finds no authentication source.
- `tracked_http_parameters` - (Optional) The HTTP request parameters to track and make available to authentication sources, selectors, and contract mappings throughout the authentication policy.

### authn_selection_trees

The `authn_selection_trees` block - An authentication policy tree.

- `authentication_api_application_ref` - (Optional) Authentication API Application Id to be used in this policy branch. If the value is not specified, no Authentication API Application will be used.
- `description` - (Optional) A description for the authentication policy.
- `enabled` - (Optional) Whether or not this authentication policy tree is enabled. Default is true.
- `name` - (Optional) The authentication policy name. Name is unique.
- `root_node` - (Optional) A node inside the authentication policy tree.

### root_node / children

The `root_node` and `children` block - An authentication policy tree node.

- `action` - (Required) The result action.
- `children` - (Optional) The nodes inside the authentication policy tree node.


### action

The `action` block - An authentication policy selection action. This can be configured for one of the following policy types.

- ApcMappingPolicyAction
- LocalIdentityMappingPolicyAction
- AuthnSelectorPolicyAction
- AuthnSourcePolicyAction
- ContinuePolicyAction
- RestartPolicyAction
- DonePolicyAction
- FragmentPolicyAction

#### APC Mapping Policy Action

The `action` block

- `attribute_mapping` - (Required) Contract fulfillment with the authentication policy contract's default values, and additional attributes retrieved from local data stores.
- `authentication_policy_contract_ref` - (Required) Reference to the associated authentication policy contract.
- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.

#### LocalIdentityMappingPolicyAction

The `action` block

- `context` - (Optional) The result context.
- `inbound_mapping` - (Optional) Inbound mappings into the local identity profile fields.
- `local_identity_ref` - (Required) Reference to the associated local identity profile.
- `outbound_attribute_mapping` - (Required) Authentication policy contract mappings associated with this local Identity profile.
- `type` - (Required) The authentication selection type.

#### AuthnSelectorPolicyAction

The `action` block

- `authentication_selector_ref` - (Required) Reference to the associated authentication selector.
- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.

#### AuthnSourcePolicyAction

The `action` block

- `attribute_rules` - (Optional) The authentication policy rules.
- `authentication_source` - (Required) The associated authentication source.
- `context` - (Optional) The result context.
- `input_user_id_mapping` - (Optional) The input user id mapping.
- `type` - (Required) The authentication selection type.
- `user_id_authenticated` (Boolean) Indicates whether the user ID obtained by the user ID mapping is authenticated.

#### ContinuePolicyAction

The `action` block

- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.

#### DonePolicyAction

The `action` block

- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.

#### RestartPolicyAction

The `action` block

- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.

#### FragmentPolicyAction

The `action` block

- `context` - (Optional) The result context.
- `type` - (Required) The authentication selection type.
- `fragment` - (Required) Reference to the associated authentication fragment.
- `fragment_mapping` - (Optional) The fragment mapping for attributes to be passed into the authentication fragment.
- `attribute_rules` - (Optional) The authentication policy rules.

### attribute_contract_fulfillment

The `attribute_contract_fulfillment` block - Defines how an attribute in an attribute contract should be populated.

- `key_name` - (Required) The attribute key mapping.
- `source` - (Required) The attribute value source.
- `value` - (Required) The value for this attribute.

### attribute_mapping

The `attribute_mapping` block - A list of mappings from attribute sources to attribute targets.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values.
- `ldap_attribute_source` - (Optional) A list of configured ldap data stores to look up attributes from.
- `jdbc_attribute_source` - (Optional) A list of configured jdbc data stores to look up attributes from.
- `custom_attribute_source` - (Optional) A list of configured custom data stores to look up attributes from.
- `issuance_criteria` - (Optional) The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.

### attribute_rule

The `attribute_rule` block - Authentication policy rules using attributes from the previous authentication source. Each rule is evaluated to determine the next action in the policy.

- `attribute_name` - (Required) The name of the attribute to use in this attribute rule.
- `condition` - (Required) The condition that will be applied to the attribute's expected value.
- `expected_value` - (Required) The expected value of this attribute rule.
- `result` - (Required) The result of this attribute rule.

### attribute_rules

The `attribute_rules` block - A collection of attribute rules

- `fallback_to_success` - (Optional) When all the rules fail, you may choose to default to the general success action or fail. Default to success.
- `items` - (Optional) The actual list of attribute rules.

### authentication_source

The `authentication_source` block - An authentication source (IdP adapter or IdP connection).

- `source_ref` - (Required) A reference to the authentication source.
- `type` - (Required) The type of this authentication source.

### conditional_criteria

The `conditional_criteria` block - An issuance criterion that checks a source attribute against a particular condition and the expected value. If the condition is true then this issuance criterion passes, otherwise the criterion fails.

- `attribute_name` - (Required) The name of the attribute to use in this issuance criterion.
- `condition` - (Required) The condition that will be applied to the source attribute's value and the expected value.
- `error_result` - (Optional) The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.
- `source` - (Required) The source of the attribute.
- `value` - (Required) The expected value of this issuance criterion.

### custom_attribute_source

The `custom_attribute_source` block - The configured settings used to look up attributes from a custom data store.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings
- `data_store_ref` - (Required) Reference to the associated data store.
- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings
- `filter_fields` - (Optional) The list of fields that can be used to filter a request to the custom data store.

### expression_criteria

The `expression_criteria` block - An issuance criterion that uses a Boolean return value from an OGNL expression to determine whether or not it passes.

- `error_result` - (Optional) The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.
- `expression` - (Required) The OGNL expression to evaluate.

### filter_fields

The `filter_fields` block - A simple name value pair to represent a field entry.

- `name` - (Required) The name of this field.
- `value` - (Optional) The value of this field. Whether or not the value is required will be determined by plugin validation checks.

### issuance_criteria

The `issuance_criteria` block - A list of criteria that determines whether a transaction (usually a SSO transaction) is continued. All criteria must pass in order for the transaction to continue.

- `conditional_criteria` - (Optional) A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue.
- `expression_criteria` - (Optional) A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue.

### jdbc_attribute_source

The `jdbc_attribute_source` block - The configured settings used to look up attributes from a JDBC data store.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings
- `data_store_ref` - (Required) Reference to the associated data store.
- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings
- `filter` - (Required) The JDBC WHERE clause used to query your data store to locate a user record.
- `schema` - (Optional) Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.
- `table` - (Required) The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.

### ldap_attribute_source

The `ldap_attribute_source` block - The configured settings used to look up attributes from a LDAP data store.

- `attribute_contract_fulfillment` - (Optional) A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings.
- `base_dn` - (Optional) The base DN to search from. If not specified, the search will start at the LDAP's root.
- `binary_attribute_settings` - (Optional) The advanced settings for binary LDAP attributes.
- `data_store_ref` - (Required) Reference to the associated data store.
- `description` - (Optional) The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings
- `member_of_nested_group` - (Optional) Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.
- `search_filter` - (Required) The LDAP filter that will be used to lookup the objects from the directory.
- `search_scope` - (Required) Determines the node depth of the query.

### source

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

## Import

-> The resource ID is fixed as `default_authentication_policies` because this is a singleton resource.

Authentication Policies can be imported using the id, e.g.

```
terraform import pingfederate_authentication_policies.demo default_authentication_policies
```
