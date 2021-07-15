# Resource: pingfederate_oauth_client

Provides an OAuth Client.

## Example Usage

```hcl
resource "pingfederate_oauth_client" "demo" {
  client_id = "demo"
  name      = "demo"

  grant_types = [
    "EXTENSION",
  ]

  client_auth {
    enforce_replay_prevention = false
    secret                    = "super_top_secret"
    type                      = "SECRET"
  }

  default_access_token_manager_ref {
    id = pingfederate_oauth_access_token_manager.demo.id
  }

  oidc_policy {
    grant_access_session_revocation_api = false
    logout_uris                         = ["https://example.com/logout"]
    ping_access_logout_capable          = true
  }
}
```

## Argument Attributes

The following arguments are supported:

- `bypass_activation_code_confirmation_override` - (Optional) Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device. This overrides the 'bypassUseCodeConfirmation' value present in Authorization Server Settings.

- `bypass_approval_page` - (Optional) Use this setting, for example, when you want to deploy a trusted application and authenticate end users via an IdP adapter or IdP connection.

- `ciba_delivery_mode` - (Optional) The token delivery mode for the client.  The default value is 'POLL'.

- `ciba_notification_endpoint` - (Optional) The endpoint the OP will call after a successful or failed end-user authentication.

- `ciba_polling_interval` - (Optional) The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds.

- `ciba_request_object_signing_algorithm` - (Optional) The JSON Web Signature [JWS] algorithm that must be used to sign the CIBA Request Object. All signing algorithms are allowed if value is not present
	 - RS256 - RSA using SHA-256
	 - RS384 - RSA using SHA-384
	 - RS512 - RSA using SHA-512
	 - ES256 - ECDSA using P256 Curve and SHA-256
	 - ES384 - ECDSA using P384 Curve and SHA-384
	 - ES512 - ECDSA using P521 Curve and SHA-512
	 - PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256
	 - PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384
	 - PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512
	 - RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.

- `ciba_require_signed_requests` - (Optional) Determines whether CIBA signed requests are required for this client.

- `ciba_user_code_supported` - (Optional) Determines whether CIBA user code is supported for this client.

- `client_auth` - (Optional) Client authentication settings.  If this model is null, it indicates that no client authentication will be used.

- `client_id` - (Required) A unique identifier the client provides to the Resource Server to identify itself. This identifier is included with every request the client makes. For PUT requests, this field is optional and it will be overridden by the 'id' parameter of the PUT request.

- `default_access_token_manager_ref` - (Optional) The default access token manager for this client.

- `description` - (Optional) A description of what the client application does. This description appears when the user is prompted for authorization.

- `device_flow_setting_type` - (Optional) Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT.

- `device_polling_interval_override` - (Optional) The amount of time client should wait between polling requests, in seconds. This overrides the 'devicePollingInterval' value present in Authorization Server Settings.

- `enabled` - (Optional) Specifies whether the client is enabled. The default value is true.

- `exclusive_scopes` - (Optional) The exclusive scopes available for this client.

- `extended_parameters` - (Optional) OAuth Client Metadata can be extended to use custom Client Metadata Parameters. The names of these custom parameters should be defined in /extendedProperties.

- `grant_types` - (Required) The grant types allowed for this client. The EXTENSION grant type applies to SAML/JWT assertion grants.

- `jwks_settings` - (Optional) JSON Web Key Set Settings of the OAuth client. Required if private key JWT client authentication or signed requests is enabled.

- `logo_url` - (Optional) The location of the logo used on user-facing OAuth grant authorization and revocation pages.

- `name` - (Required) A descriptive name for the client instance. This name appears when the user is prompted for authorization.

- `oidc_policy` - (Optional) Open ID Connect Policy settings.  This is included in the message only when OIDC is enabled.

- `pending_authorization_timeout_override` - (Optional) The 'device_code' and 'user_code' timeout, in seconds. This overrides the 'pendingAuthorizationTimeout' value present in Authorization Server Settings.

- `persistent_grant_expiration_time` - (Optional) The persistent grant expiration time. -1 indicates an indefinite amount of time.

- `persistent_grant_expiration_time_unit` - (Optional) The persistent grant expiration time unit.

- `persistent_grant_expiration_type` - (Optional) Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT.

- `persistent_grant_idle_timeout` - (Optional) The persistent grant idle timeout.

- `persistent_grant_idle_timeout_time_unit` - (Optional) The persistent grant idle timeout time unit.

- `persistent_grant_idle_timeout_type` - (Optional) Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT.

- `redirect_uris` - (Optional) URIs to which the OAuth AS may redirect the resource owner's user agent after authorization is obtained. A redirection URI is used with the Authorization Code and Implicit grant types. Wildcards are allowed. However, for security reasons, make the URL as restrictive as possible.For example: https://*.company.com/* Important: If more than one URI is added or if a single URI uses wildcards, then Authorization Code grant and token requests must contain a specific matching redirect uri parameter.

- `refresh_rolling` - (Optional) Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT.

- `request_object_signing_algorithm` - (Optional) The JSON Web Signature [JWS] algorithm that must be used to sign the Request Object. All signing algorithms are allowed if value is not present
	 - RS256 - RSA using SHA-256
	 - RS384 - RSA using SHA-384
	 - RS512 - RSA using SHA-512
	 - ES256 - ECDSA using P256 Curve and SHA-256
	 - ES384 - ECDSA using P384 Curve and SHA-384
	 - ES512 - ECDSA using P521 Curve and SHA-512
	 - PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256
	 - PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384
	 - PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512
	 - RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.

- `request_policy_ref` - (Optional) The CIBA request policy.

- `require_proof_key_for_code_exchange` - (Optional) Determines whether Proof Key for Code Exchange (PKCE) is required for this client.

- `require_signed_requests` - (Optional) Determines whether signed requests are required for this client

- `restrict_scopes` - (Optional) Restricts this client's access to specific scopes.

- `restricted_response_types` - (Optional) The response types allowed for this client. If omitted all response types are available to the client.

- `restricted_scopes` - (Optional) The scopes available for this client.

- `token_exchange_processor_policy_ref` - (Optional) The Token Exchange Processor policy.

- `user_authorization_url_override` - (Optional) The URL used as 'verification_url' and 'verification_url_complete' values in a Device Authorization request. This property overrides the 'userAuthorizationUrl' value present in Authorization Server Settings.

- `validate_using_all_eligible_atms` - (Optional) Validates token using all eligible access token managers for the client.

- `restrict_to_default_access_token_manager` - (Optional) Determines whether the client is restricted to using only its default access token manager. The default is false.

### client_auth

The `client_auth` block - Client Authentication.

- `client_cert_issuer_dn` - (Optional) Client TLS Certificate Issuer DN.

- `client_cert_subject_dn` - (Optional) Client TLS Certificate Subject DN.

- `enforce_replay_prevention` - (Optional) Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication.

- `secret` - (Optional) Client secret for Basic Authentication.  To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests.

- `token_endpoint_auth_signing_algorithm` - (Optional) The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present

     - RS256 - RSA using SHA-256
	 - RS384 - RSA using SHA-384
	 - RS512 - RSA using SHA-512
	 - ES256 - ECDSA using P256 Curve and SHA-256
	 - ES384 - ECDSA using P384 Curve and SHA-384
	 - ES512 - ECDSA using P521 Curve and SHA-512
	 - PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256
	 - PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384
	 - PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512
	 - RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.

- `type` - (Optional) Client authentication type. The required field for type SECRET is secret. The required fields for type CERTIFICATE are clientCertIssuerDn and clientCertSubjectDn. The required field for type PRIVATE_KEY_JWT is: either jwks or jwksUrl.

### oidc_policy

The `oidc_policy` block - OAuth Client Open ID Connect Policy.

- `grant_access_session_revocation_api` - (Optional) Determines whether this client is allowed to access the Session Revocation API.

- `id_token_content_encryption_algorithm` - (Optional) The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.
	 - AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256
	 - AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384
	 - AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512
	 - AES-GCM-128 - AES_128_GCM
	 - AES_192_GCM - AES-GCM-192
	 - AES_256_GCM - AES-GCM-256

- `id_token_encryption_algorithm` - (Optional) The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.
	 - DIR - Direct Encryption with symmetric key
	 - A128KW - AES-128 Key Wrap
	 - A192KW - AES-192 Key Wrap
	 - A256KW - AES-256 Key Wrap
	 - A128GCMKW - AES-GCM-128 key encryption
	 - A192GCMKW - AES-GCM-192 key encryption
	 - A256GCMKW - AES-GCM-256 key encryption
	 - ECDH_ES - ECDH-ES
	 - ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap
	 - ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap
	 - ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap
	 - RSA_OAEP - RSAES OAEP

- `id_token_signing_algorithm` - (Optional) The JSON Web Signature [JWS] algorithm required for the ID Token.
	 - NONE - No signing algorithm
	 - HS256 - HMAC using SHA-256
	 - HS384 - HMAC using SHA-384
	 - HS512 - HMAC using SHA-512
	 - RS256 - RSA using SHA-256
	 - RS384 - RSA using SHA-384
	 - RS512 - RSA using SHA-512
	 - ES256 - ECDSA using P256 Curve and SHA-256
	 - ES384 - ECDSA using P384 Curve and SHA-384
	 - ES512 - ECDSA using P521 Curve and SHA-512
	 - PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256
	 - PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384
	 - PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512
	 - A null value will represent the default algorithm which is RS256.
	 - RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11

- `logout_uris` - (Optional) A list of client logout URI's which will be invoked when a user logs out through one of PingFederate's SLO endpoints.

- `pairwise_identifier_user_type` - (Optional) Determines whether the subject identifier type is pairwise.

- `ping_access_logout_capable` - (Optional) Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention.

- `policy_group` - (Optional) The Open ID Connect policy. A null value will represent the default policy group.

- `sector_identifier_uri` - (Optional) The URI references a file with a single JSON array of Redirect URI and JWKS URL values.


### jwks_settings

The `jwks_settings` block - JSON Web Key Set Settings.

- `jwks` - (Optional) JSON Web Key Set (JWKS) document of the OAuth client. Either 'jwks' or 'jwksUrl' must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.

- `jwks_url` - (Optional) JSON Web Key Set (JWKS) URL of the OAuth client. Either 'jwks' or 'jwksUrl' must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- [`id`](#id) - The ID of the oauth client (`client_id`).

## Import

OAuth Clients can be imported using the id, e.g.

```
terraform import pingfederate_oauth_client.demo 123
```
