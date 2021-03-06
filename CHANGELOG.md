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
