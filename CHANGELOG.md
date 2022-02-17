## 0.0.24 (Unreleased)

FEATURES:
- Add support for `pingfederate_idp_token_processor`. (#160)

BREAKING CHANGES:
- Behaviour change for `pingfederate_idp_sp_connection` the `ws_trust.attribute_contract.core_attributes` is computed and can no longer be set. (#159)

BUG FIXES:

- Remove unnecessary retry on OAuth client creation. (#165)
- Fixed issue `incoming_proxy_settings` didn't run on 11+ . (#162)
- Fixed issue `ws_trust` on `pingfederate_idp_sp_connection` when `request_contract_ref` is not provided. (#158)
- Ensure `logging_module` on `pingfederate_idp_sp_connection` has a default (`STANDARD`). (#159)

## 0.0.23 (December 27th, 2021)

NOTES:

* Regression testing against 9.3.3 and 10.0.x has been removed as PingIdentity no longer provider working container images.
* Tested against newly released PingFederate 11.0

BUG FIXES:

- Authentication policy failed marshalling of `attribute_rules` in PingFederate API. (#148)
- Change `ws_trust.request_contract_ref` on `pingfederate_idp_sp_connection` to no longer be required. (#155)

## 0.0.22 (October 23rd, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

- Add support for `pingfederate_oauth_client_settings` (singleton). (#145)
- Add support for `pingfederate_oauth_client_registration_policy` to support DCR settings. (#145)
- Add support for `pingfederate_oauth_access_token_manager_settings`. (#143)

BUG FIXES:

- Authentication policy tree default enabled status regression has been fixed back to `true`. (#139)

## 0.0.21 (September 13th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

- Added more cross functional PingFederate version testing, 9.3 -> 10.3.

BUG FIXES:

- Don't set configuration default rows unless true (allows PF9.3 to work). (#135)

## 0.0.20 (September 8th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

- Add support for `pingfederate_certificates_revocation_settings`. (#130)
- Add support for `pingfederate_pingone_connection`. (#115)

BUG FIXES:

- Fixed incorrect validation on AttributeContractFulfillment. (#129)

## 0.0.19 (August 13th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

- Add data source to return PingFederate version. (#116)
- Generated provider docs for all resources and data-sources. (#114)
- Added functional testing against PingFederate 10.3.x

BUG FIXES:

- Fixed issue with `pingfederate_idp_sp_connection` block `outbound_provision` not correctly handling `sensitive_target_settings`.
- Fixed issue with `pingfederate_idp_sp_connection` block `sp_browser_sso.attribute_contract.extended_attributes` ordering. (#117)

DEPRECATIONS:

- The `pingfederate_oauth_auth_server_settings` attribute `approved_scope_attribute` has been deprecated please use `approved_scopes_attribute` to correctly align with the Admin API.
- The `pingfederate_oauth_openid_connect_policy` attribute `include_user_in_id_token` has been deprecated please use `include_user_info_in_id_token` to correctly align with the Admin API.
- The `pingfederate_server_settings` block `roles_and_protocols` has been deprecated in PingFederate 10.1.

## 0.0.18 (July 15th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

- Add support for PAR oauth server settings. (#108)
- Add support configuring IDP Adapter `instance_id`. (#106)

BUG FIXES:

- Add missing field to oauth client `restrict_to_default_access_token_manager`. (#109)
- Fix broken field to oauth server settings `admin_web_service_pcv_ref`. (#105)

## 0.0.17 (June 3rd, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BREAKING CHANGES:

Attribute contracts for the following resources no longer require read-only `core_attributes` and they can no longer be defined, this affects the following resources.

- `pingfederate_sp_adapter`
- `pingfederate_oauth_openid_connect_policy`
- `pingfederate_password_credential_validator`

FEATURES:

- Add support for extended properties `pingfederate_extended_properties`.
- Add support for incoming proxy settings `pingfederate_incoming_proxy_settings` (#85).
- Add support for openid connect keypairs `pingfederate_keypairs_oauth_openid_connect` (#86).
- Enhance oauth access token managers with support for `parent_ref`, `access_control_settings`, `selection_settings` and `session_validation_settings` (#99).

BUG FIXES:

- Fixed issue with openid connect policy core attributes not being computed. (#94)

## 0.0.16 (May 9th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x
- Add ability to specify `default_row` for configuration table rows, the default value is `false`, be sure to review any configuration blocks with rows before upgrading.
- This main test version for the provider is now 10.2, the functional tests cover all versions of 10.x still.

FEATURES:

- Add custom data store resource/data sources. (#82)
- Add authentication policy fragment (PF10.2+). (#80)
- Add support for dynamic scopes (oauth_auth_server_settings). (#82) (Thanks to @mosersil for this contribution)
- Add support setting authentication policy contract id. (#81)
- Add support setting data stores id. (#89)

BUG FIXES:

- Add workaround for PingFederate bug with race conditions deleting authentication policy contracts and sp connections. (#91)

## 0.0.15 (Apr 25th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x
- Some additional documentation cleanup. (#62)
- Added ability to handle certificate formatting differences for `idp_sp_connection`  `credentials -> certs -> x509_file` block.

FEATURES:

- Add jdbc/ldap data store data sources. (#71)

BUG FIXES:

- Add workaround for PingFederate bug with race conditions deleting data stores and sp connections. (#77)
- Fix crash with empty `contact_info` block on `idp_sp_connection`. (#70 )

## 0.0.14 (Apr 13th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BUG FIXES:

* Fix issue with imported keypairs and `key_size` fields. ([#68](https://github.com/iwarapter/terraform-provider-pingfederate/issues/68))

## 0.0.13 (Apr 12th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BUG FIXES:

* Fix `idp_sp_connections` credentials.certs to correctly compute the `cert_view` block.
* Add workaround for PingFederate bug with race conditions deleting data stores and sp connections. ([#66](https://github.com/iwarapter/terraform-provider-pingfederate/issues/66))
* Add workaround for PingFederate bug with race conditions when managing `certificate_ca`. ([#63](https://github.com/iwarapter/terraform-provider-pingfederate/issues/63))

## 0.0.12 (Mar 11th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `pingfederate_oauth_resource_owner_credentials_mappings` ([#55](https://github.com/iwarapter/terraform-provider-pingfederate/issues/55))

BUG FIXES:

* Ensure certain reference fields force new resource if changed.
    * `plugin_descriptor_ref` on `pingfederate_authentication_selector`, `pingfederate_oauth_access_token_manager`,`resource_pingfederate_idp_adapter`,`resource_pingfederate_sp_adapter` and `pingfederate_password_credential_validator`
    * `access_token_manager_ref` / `context_ref` on `resource_pingfederate_oauth_access_token_mappings`

## 0.0.11 (Feb 28th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* Add support for `darwin/arm64` release. ([#52](https://github.com/iwarapter/terraform-provider-pingfederate/issues/52))

BUG FIXES:

* resource/pingfederate_authentication_policies: Increase depth of nested policies.
* resource/pingfederate_jdbc_data_store: Fix idempotency on `encrypted_password` field. ([#51](https://github.com/iwarapter/terraform-provider-pingfederate/issues/51))

## 0.0.10 (Feb 11th, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BUG FIXES:

* Ensure OIDC Policy overrides for default delivery is correctly set.

## 0.0.9 (Feb 1st, 2021)

NOTES:

* This is an early release experimental build for PingFederate 10.x
* Added documentation for the following resources:
  - `pingfederate_authentication_api_application`
  - `pingfederate_authentication_selector`
  - `pingfederate_jdbc_data_store`
  - `pingfederate_ldap_data_store`
  - `pingfederate_oauth_openid_connect_policy`

BUG FIXES:

* Fix #42
* Fix authentication selector not ForceNew when name or plugin changes.
* Add hostname tags for ldap_data_store.
* Fix oauth_access_token_mappings mishandling issuance_criteria
* Add retry on oauth client creation to handle flakey PF API.
* Add retry on oauth openid connect policy creation to handle flakey PF API.

## 0.0.8 (Nov 17th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x
* The `bypass_external_validation` which several resources used has been migrated to provider level configuration.
This is a breaking change as the attribute was removed from the following affected resources:
    - `resource_pingfederate_authentication_policies`
    - `resource_pingfederate_idp_adapter`
    - `resource_pingfederate_idp_sp_connection`
    - `resource_pingfederate_jdbc_data_store`
    - `resource_pingfederate_ldap_data_store`
    - `resource_pingfederate_oauth_access_token_mappings`
    - `resource_pingfederate_oauth_openid_connect_policy`
    - `resource_pingfederate_sp_authentication_policy_contract_mapping`
    - `resource_pingfederate_sp_idp_connection`

BUG FIXES:

* Fix issues with importing several resources.
* Add configuration validation for the provider block for any initial connection issues.

## 0.0.7 (Nov 9th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BUG FIXES:

* resource/pingfederate_authentication_selector: Fix handling of the attribute contract.

## 0.0.6 (Oct 15th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

BUG FIXES:

* resource/pingfederate_oauth_client: Was missing several configuration fields and didnt support `NONE` client auth types.
* Fixed issue with root boolean values not being set to `false` on several resources
* Changed descriptor validation logic to soft fail if the role isnt enabled (but could be as part of the apply).

## 0.0.5 (Sept 14th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `pingfederate_authentication_api_application`
* **New Resource:** `pingfederate_authentication_api_settings`
* **New Resource:** `pingfederate_idp_sp_connection`
* **New Resource:** `pingfederate_kerberos_realm`
* **New Resource:** `pingfederate_keypair_signing`
* **New Resource:** `pingfederate_sp_idp_connection`

## 0.0.4 (Aug 24th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `resource_pingfederate_certificates_ca`

## 0.0.3 (Aug 24th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `resource_pingfederate_notification_publisher`

## 0.0.2 (Aug 6th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `pingfederate_authentication_policies_settings`
* **New Resource:** `pingfederate_authentication_policies`
* **New Resource:** `pingfederate_server_settings`

## 0.0.1 (June 24th, 2020)

NOTES:

* This is an early release experimental build for PingFederate 10.x

FEATURES:

* **New Resource:** `pingfederate_authentication_policy_contract`
* **New Resource:** `pingfederate_authentication_selector`
* **New Resource:** `pingfederate_jdbc_data_store`
* **New Resource:** `pingfederate_ldap_data_store`
* **New Resource:** `pingfederate_idp_adapter`
* **New Resource:** `pingfederate_oauth_auth_server_settings`
* **New Resource:** `pingfederate_oauth_authentication_policy_contract_mapping`
* **New Resource:** `pingfederate_oauth_client`
* **New Resource:** `pingfederate_oauth_access_token_manager`
* **New Resource:** `pingfederate_oauth_access_token_mappings`
* **New Resource:** `pingfederate_oauth_openid_connect_policy`
* **New Resource:** `pingfederate_sp_adapter`
* **New Resource:** `pingfederate_sp_authentication_policy_contract_mapping`
* **New Resource:** `pingfederate_password_credential_validator`
