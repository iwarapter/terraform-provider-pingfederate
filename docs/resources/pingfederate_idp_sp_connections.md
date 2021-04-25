# Resource: pingfederate_idp_sp_connection

Provides a idp sp connection.

## Example Usage
```hcl
resource "pingfederate_idp_sp_connection" "demo" {
  name         = "foo"
  entity_id    = "foo"
  active       = true
  logging_mode = "STANDARD"
  credentials {
    inbound_back_channel_auth {
      type                    = "INBOUND"
      digital_signature       = false
      require_ssl             = false
      verification_subject_dn = "CN=Example"
    }
  }
  attribute_query {
    jdbc_attribute_source {
      filter      = "*"
      description = "foo"
      schema      = "INFORMATION_SCHEMA"
      table       = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
      id          = "foo"
      data_store_ref {
        id = "ProvisionerDS"
      }
    }

    attribute_contract_fulfillment {
      key_name = "foo"
      source {
        type = "JDBC_DATA_STORE"
        id   = "foo"
      }
      value = "GRANTEE"
    }

    attributes = ["foo"]
    policy {
      sign_response                  = false
      sign_assertion                 = false
      encrypt_assertion              = false
      require_signed_attribute_query = false
      require_encrypted_name_id      = false
    }
  }
}
```

## Argument Attributes

The following arguments are supported:

- `connection_id` - (Optional) The persistent, unique ID for the connection. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.

- `active` - (Optional) Specifies whether the connection is active and ready to process incoming requests. The default value is false.

- `additional_allowed_entities_configuration` - (Optional) Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.

- `application_icon_url` - (Optional) The application icon url.

- `application_name` - (Optional) The application name.

- `attribute_query` - (Optional) The attribute query settings for supporting SPs in requesting user attributes.

- `base_url` - (Optional) The fully-qualified hostname and port on which your partner's federation deployment runs.

- `contact_info` - (Optional) The contact information for this partner.

- `credentials` - (Optional) The certificates and settings for encryption, signing, and signature verification. It is required for  SAMLx.x and WS-Fed Connections.

- `default_virtual_entity_id` - (Optional) The default alternate entity ID that identifies the local server to this partner. It is required when virtualEntityIds is not empty and must be included in that list.

- `entity_id` - (Required) The partner's entity ID (connection ID) or issuer value (for OIDC Connections).

- `extended_properties` - (Optional) Extended Properties allows to store additional information for IdP/SP Connections. The names of these extended properties should be defined in /extendedProperties.

- `license_connection_group` - (Optional) The license connection group. If your PingFederate license is based on connection groups, each connection must be assigned to a group before it can be used.

- `logging_mode` - (Optional) The level of transaction logging applicable for this connection. Default is STANDARD.

- `metadata_reload_settings` - (Optional) Connection metadata automatic reload settings.

- `name` - (Required) The connection name.

- `outbound_provision` - (Optional) The Outbound Provision settings.

- `sp_browser_sso` - (Optional) The browser-based SSO settings used to communicate with your SP.

- `type` - (Required) The type of this connection. This must be set to 'SP'.

- `virtual_entity_ids` - (Optional) List of alternate entity IDs that identifies the local server to this partner.

- `ws_trust` - (Optional) The Ws-Trust settings.

### AdditionalAllowedEntitiesConfiguration

The `additional_allowed_entities_configuration` block - Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.

- `additional_allowed_entities` - (Optional) An array of additional allowed entities or issuers to be accepted during entity or issuer validation.

- `allow_additional_entities` - (Optional) Set to true to configure additional entities or issuers to be accepted during entity or issuer validation.

- `allow_all_entities` - (Optional) Set to true to accept any entity or issuer during entity or issuer validation. (Not Recommended)

### SpAttributeQuery

The `sp_attribute_query` block - The attribute query profile supports SPs in requesting user attributes.

- `attribute_contract_fulfillment` - (Required) A list of mappings from attribute names to their fulfillment values.

- `attribute_sources` - (Required) A list of configured data stores to look up attributes from.

- `attributes` - (Required) The list of attributes that may be returned to the SP in the response to an attribute request.

- `issuance_criteria` - (Optional) The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.

- `policy` - (Optional) The attribute query profile's security policy.

### ContactInfo

The `contact_info` block - Contact information.

- `company` - (Optional) Company name.

- `email` - (Optional) Contact email address.

- `first_name` - (Optional) Contact first name.

- `last_name` - (Optional) Contact last name.

- `phone` - (Optional) Contact phone number.

### ConnectionCredentials

The `connection_credentials` block - The certificates and settings for encryption, signing, and signature verification.

- `block_encryption_algorithm` - (Optional) The algorithm used to encrypt assertions sent to this partner. AES_128, AES_256 and Triple_DES are also supported. Default is AES_128

- `certs` - (Optional) The certificates used for signature verification and XML encryption.

- `decryption_key_pair_ref` - (Optional) The ID of the primary key pair used to decrypt message content received from this partner. The ID of the key pair is also known as the alias and can be found by viewing the corresponding certificate under 'Signing & Decryption Keys & Certificates' in the PingFederate Administrative Console.

- `inbound_back_channel_auth` - (Optional) The SOAP authentication method(s) to use when you receive a message using SOAP back channel.

- `key_transport_algorithm` - (Optional) The algorithm used to transport keys to this partner. RSA_OAEP and RSA_v15 are supported. Default is RSA_OAEP

- `outbound_back_channel_auth` - (Optional) The SOAP authentication method(s) to use when you send a message using SOAP back channel.

- `secondary_decryption_key_pair_ref` - (Optional) The ID of the secondary key pair used to decrypt message content received from this partner.

- `signing_settings` - (Optional) Settings related to the manner in which messages sent to the partner are digitally signed. Required for SP Connections.

- `verification_issuer_dn` - (Optional) If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field.

- `verification_subject_dn` - (Optional) If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array.

### ConnectionMetadataUrl

The `connection_metadata_url` block - Configuration settings to enable automatic reload of partner's metadata.

- `enable_auto_metadata_update` - (Optional) Specifies whether the metadata of the connection will be automatically reloaded. The default value is true.

- `metadata_url_ref` - (Required) ID of the saved Metadata URL.

### OutboundProvision

The `outbound_provision` block - Outbound Provisioning allows an IdP to create and maintain user accounts at standards-based partner sites using SCIM as well as select-proprietary provisioning partner sites that are protocol-enabled.

- `channels` - (Required) Includes settings of a source data store, managing provisioning threads and mapping of attributes.

- `custom_schema` - (Optional) Custom SCIM attribute configuration.

- `target_settings` - (Required) Configuration fields that includes credentials to target SaaS application.

- `type` - (Required) The SaaS plugin type.

### SpBrowserSso

The `sp_browser_sso` block - The SAML settings used to enable secure browser-based SSO to resources at your partner's site.

- `adapter_mappings` - (Required) A list of adapters that map to outgoing assertions.

- `artifact` - (Optional) The settings for an artifact binding.

- `assertion_lifetime` - (Required) The timeframe of validity before and after the issuance of the assertion.

- `attribute_contract` - (Required) A set of user attributes that the IdP sends in the SAML assertion.

- `authentication_policy_contract_assertion_mappings` - (Optional) A list of authentication policy contracts that map to outgoing assertions.

- `default_target_url` - (Optional) Default Target URL for SAML1.x connections. For SP connections, this default URL represents the destination on the SP where the user will be directed. For IdP connections, entering a URL in the Default Target URL field overrides the SP Default URL SSO setting.

- `enabled_profiles` - (Optional) The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.

- `encryption_policy` - (Required) The SAML 2.0 encryption policy for browser-based SSO. Required for SAML 2.0 connections.

- `incoming_bindings` - (Optional) The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections.

- `message_customizations` - (Optional) The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported.

- `protocol` - (Required) The browser-based SSO protocol to use.

- `require_signed_authn_requests` - (Optional) Require AuthN requests to be signed when received via the POST or Redirect bindings.

- `sign_assertions` - (Optional) Always sign the SAML Assertion.

- `sign_response_as_required` - (Optional) Sign SAML Response as required by the associated binding and encryption policy. Applicable to SAML2.0 only and is defaulted to true. It can be set to false only on SAML2.0 connections when signAssertions is set to true.

- `slo_service_endpoints` - (Optional) A list of possible endpoints to send SLO requests and responses.

- `sp_saml_identity_mapping` - (Optional) Process in which users authenticated by the IdP are associated with user accounts local to the SP.

- `sp_ws_fed_identity_mapping` - (Optional) Process in which users authenticated by the IdP are associated with user accounts local to the SP for WS-Federation connection types.

- `sso_service_endpoints` - (Required) A list of possible endpoints to send assertions to.

- `url_whitelist_entries` - (Optional) For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled.

- `ws_fed_token_type` - (Optional) The WS-Federation Token Type to use.

- `ws_trust_version` - (Optional) The WS-Trust version for a WS-Federation connection. The default version is WSTRUST12.

### EncryptionPolicy

The `encryption_policy` block - Defines what to encrypt in the browser-based SSO profile.

- `encrypt_assertion` - Whether the outgoing SAML assertion will be encrypted.
- `encrypted_attributes` - The list of outgoing SAML assertion attributes that will be encrypted. The `encrypt_assertion` property takes precedence over this.
- `encrypt_slo_subject_name_id` - Encrypt the name-identifier attribute in outbound SLO messages. This can be set if the name id is encrypted.
- `slo_subject_name_id_encrypted` - Allow the encryption of the name-identifier attribute for inbound SLO messages. This can be set if SP initiated SLO is enabled.

### SpWsTrust

The `sp_ws_trust` block - Ws-Trust STS provides security-token validation and creation to extend SSO access to identity-enabled Web Services

- `abort_if_not_fulfilled_from_request` - (Optional) If the attribute contract cannot be fulfilled using data from the Request, abort the transaction.

- `attribute_contract` - (Required) A set of user attributes that the IdP sends in the token.

- `default_token_type` - (Optional) The default token type when a web service client (WSC) does not specify in the token request which token type the STS should issue. Defaults to SAML 2.0.

- `encrypt_saml2_assertion` - (Optional) When selected, the STS encrypts the SAML 2.0 assertion. Applicable only to SAML 2.0 security token.  This option does not apply to OAuth assertion profiles.

- `generate_key` - (Optional) When selected, the STS generates a symmetric key to be used in conjunction with the "Holder of Key" (HoK) designation for the assertion's Subject Confirmation Method.  This option does not apply to OAuth assertion profiles.

- `message_customizations` - (Optional) The message customizations for WS-Trust. Depending on server settings, connection type, and protocol this may or may not be supported.

- `minutes_after` - (Optional) The amount of time after the SAML token was issued during which it is to be considered valid. The default value is 30.

- `minutes_before` - (Optional) The amount of time before the SAML token was issued during which it is to be considered valid. The default value is 5.

- `o_auth_assertion_profiles` - (Optional) When selected, four additional token-type requests become available.

- `partner_service_ids` - (Required) The partner service identifiers.

- `request_contract_ref` - (Optional) Request Contract to be used to map attribute values into the security token.

- `token_processor_mappings` - (Required) A list of token processors to validate incoming tokens.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The idp sp connections ID (`connection_id`)

## Import

IDP SP Connections can be imported using the id, e.g.

```
terraform import pingfederate_idp_sp_connection.demo 123
```
