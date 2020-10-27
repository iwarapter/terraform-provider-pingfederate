## 0.0.7 (Unreleased)

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
