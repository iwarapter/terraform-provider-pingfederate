package framework

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func resourceApcToPersistentGrantMapping() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `An authentication policy contract mapping into an OAuth persistent grant.`,
		Version:     1,
		Attributes: map[string]tfsdk.Attribute{
			"attribute_contract_fulfillment": {
				Description: `A list of mappings from attribute names to their fulfillment values.`,
				Required:    true,
				Attributes:  tfsdk.MapNestedAttributes(mapAttributeFulfillmentValue()),
			},
			"jdbc_attribute_sources": {
				Description: `The configured settings used to look up attributes from a JDBC data store.`,
				Optional:    true,
				Attributes:  tfsdk.ListNestedAttributes(listJdbcAttributeSource()),
			},
			"ldap_attribute_sources": {
				Description: `The configured settings used to look up attributes from a LDAP data store.`,
				Optional:    true,
				Attributes:  tfsdk.ListNestedAttributes(listLdapAttributeSource()),
			},
			"custom_attribute_sources": {
				Description: `The configured settings used to look up attributes from a custom data store.`,
				Optional:    true,
				Attributes:  tfsdk.ListNestedAttributes(listCustomAttributeSource()),
			},
			"authentication_policy_contract_ref": {
				Description: `Reference to the associated authentication policy contract. The reference cannot be changed after the mapping has been created.`,
				Required:    true,
				Type:        types.StringType,
			},
			"id": {
				Description: `The ID of the authentication policy contract to persistent grant mapping.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"issuance_criteria": {
				Description: `The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleIssuanceCriteria()),
			},
		},
	}
}

func resourceAuthenticationPolicyContract() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `Authentication Policy Contracts carry user attributes from the identity provider to the service provider.`,
		Attributes: map[string]tfsdk.Attribute{
			"core_attributes": {
				Description: `A list of read-only assertion attributes (for example, subject) that are automatically populated by PingFederate.`,
				Optional:    true,
				Computed:    true,
				Type: types.SetType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.SetValueMust(types.StringType, []attr.Value{types.StringValue("subject")})),
				},
			},
			"extended_attributes": {
				Description: `A list of additional attributes as needed.`,
				Optional:    true,
				Type: types.SetType{
					ElemType: types.StringType,
				},
			},
			"id": {
				Description: `The persistent, unique ID for the authentication policy contract. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"name": {
				Description: `The Authentication Policy Contract Name. Name is unique.`,
				Required:    true,
				Type:        types.StringType,
			},
		},
	}
}

func resourceClient() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `OAuth client.`,
		Version:     1,
		Attributes: map[string]tfsdk.Attribute{
			"allow_authentication_api_init": {
				Description: `Set to true to allow this client to initiate the authentication API redirectless flow.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"bypass_activation_code_confirmation_override": {
				Description: `Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device. This overrides the 'bypassUseCodeConfirmation' value present in Authorization Server Settings.`,
				Optional:    true,
				Type:        types.BoolType,
			},
			"bypass_approval_page": {
				Description: `Use this setting, for example, when you want to deploy a trusted application and authenticate end users via an IdP adapter or IdP connection.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"ciba_delivery_mode": {
				Description: `The token delivery mode for the client.  The default value is 'POLL'.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"ciba_notification_endpoint": {
				Description: `The endpoint the OP will call after a successful or failed end-user authentication.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"ciba_polling_interval": {
				Description: `The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"ciba_request_object_signing_algorithm": {
				Description: `The JSON Web Signature [JWS] algorithm that must be used to sign the CIBA Request Object. All signing algorithms are allowed if value is not present <br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"ciba_require_signed_requests": {
				Description: `Determines whether CIBA signed requests are required for this client.`,
				Optional:    true,
				Type:        types.BoolType,
			},
			"ciba_user_code_supported": {
				Description: `Determines whether CIBA user code is supported for this client.`,
				Optional:    true,
				Type:        types.BoolType,
			},
			"client_auth": {
				Description: `Client authentication settings.  If this model is null, it indicates that no client authentication will be used.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleClientAuth()),
			},
			"client_id": {
				Description: `A unique identifier the client provides to the Resource Server to identify itself. This identifier is included with every request the client makes. For PUT requests, this field is optional and it will be overridden by the 'id' parameter of the PUT request.`,
				Required:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"client_secret_changed_time": {
				Description: `The time at which the client secret was last changed. This property is read only and is ignored on PUT and POST requests.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
			},
			"client_secret_retention_period": {
				Description: `The length of time in minutes that client secrets will be retained as secondary secrets after secret change. The default value is 0, which will disable secondary client secret retention. This value will override the Client Secret Retention Period value on the Authorization Server Settings.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"client_secret_retention_period_type": {
				Description: `Use OVERRIDE_SERVER_DEFAULT to override the Client Secret Retention Period value on the Authorization Server Settings. SERVER_DEFAULT will default to the Client Secret Retention Period value on the Authorization Server Setting. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"default_access_token_manager_ref": {
				Description: `The default access token manager for this client.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"description": {
				Description: `A description of what the client application does. This description appears when the user is prompted for authorization.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"device_flow_setting_type": {
				Description: `Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"device_polling_interval_override": {
				Description: `The amount of time client should wait between polling requests, in seconds. This overrides the 'devicePollingInterval' value present in Authorization Server Settings.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"enabled": {
				Description: `Specifies whether the client is enabled. The default value is true.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(true)),
				},
			},
			"exclusive_scopes": {
				Description: `The exclusive scopes available for this client.`,
				Optional:    true,
				Computed:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.ListValueMust(types.StringType, []attr.Value{})),
				},
			},
			"extended_parameters": {
				Description: `OAuth Client Metadata can be extended to use custom Client Metadata Parameters. The names of these custom parameters should be defined in /extendedProperties.`,
				Optional:    true,
				Attributes:  tfsdk.MapNestedAttributes(mapParameterValues()),
			},
			"grant_types": {
				Description: `The grant types allowed for this client. The EXTENSION grant type applies to SAML/JWT assertion grants.`,
				Required:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
			},
			"id": {
				Description: `The client_id of the oauth client.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"jwks_settings": {
				Description: `JSON Web Key Set Settings of the OAuth client. Required if private key JWT client authentication or signed requests is enabled.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleJwksSettings()),
			},
			"jwt_secured_authorization_response_mode_content_encryption_algorithm": {
				Description: `The JSON Web Encryption [JWE] content-encryption algorithm for the JWT Secured Authorization Response.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES_128_GCM - AES-GCM-128<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256`,
				Optional:    true,
				Type:        types.StringType,
			},
			"jwt_secured_authorization_response_mode_encryption_algorithm": {
				Description: `The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content-encryption key of the JWT Secured Authorization Response.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256`,
				Optional:    true,
				Type:        types.StringType,
			},
			"jwt_secured_authorization_response_mode_signing_algorithm": {
				Description: `The JSON Web Signature [JWS] algorithm required to sign the JWT Secured Authorization Response.<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11`,
				Optional:    true,
				Type:        types.StringType,
			},
			"logo_url": {
				Description: `The location of the logo used on user-facing OAuth grant authorization and revocation pages.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"name": {
				Description: `A descriptive name for the client instance. This name appears when the user is prompted for authorization.`,
				Required:    true,
				Type:        types.StringType,
			},
			"oidc_policy": {
				Description: `Open ID Connect Policy settings.  This is included in the message only when OIDC is enabled.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleClientOIDCPolicy()),
			},
			"pending_authorization_timeout_override": {
				Description: `The 'device_code' and 'user_code' timeout, in seconds. This overrides the 'pendingAuthorizationTimeout' value present in Authorization Server Settings.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"persistent_grant_expiration_time": {
				Description: `The persistent grant expiration time. -1 indicates an indefinite amount of time.`,
				Optional:    true,
				Computed:    true,
				Type:        types.NumberType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.NumberValue(big.NewFloat(0))),
				},
			},
			"persistent_grant_expiration_time_unit": {
				Description: `The persistent grant expiration time unit.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("DAYS")),
				},
			},
			"persistent_grant_expiration_type": {
				Description: `Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"persistent_grant_idle_timeout": {
				Description: `The persistent grant idle timeout.`,
				Optional:    true,
				Computed:    true,
				Type:        types.NumberType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.NumberValue(big.NewFloat(0))),
				},
			},
			"persistent_grant_idle_timeout_time_unit": {
				Description: `The persistent grant idle timeout time unit.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("DAYS")),
				},
			},
			"persistent_grant_idle_timeout_type": {
				Description: `Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"persistent_grant_reuse_grant_types": {
				Description: `The grant types that the OAuth AS can reuse rather than creating a new grant for each request. This value will override the Reuse Existing Persistent Access Grants for Grant Types on the Authorization Server Settings. Only 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' are valid grant types.`,
				Optional:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
			},
			"persistent_grant_reuse_type": {
				Description: `Allows and administrator to override the Reuse Existing Persistent Access Grants for Grant Types set globally for OAuth AS. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"redirect_uris": {
				Description: `URIs to which the OAuth AS may redirect the resource owner's user agent after authorization is obtained. A redirection URI is used with the Authorization Code and Implicit grant types. Wildcards are allowed. However, for security reasons, make the URL as restrictive as possible.For example: https://*.company.com/* Important: If more than one URI is added or if a single URI uses wildcards, then Authorization Code grant and token requests must contain a specific matching redirect uri parameter.`,
				Optional:    true,
				Computed:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.ListValueMust(types.StringType, []attr.Value{})),
				},
			},
			"refresh_rolling": {
				Description: `Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"refresh_token_rolling_grace_period": {
				Description: `The grace period that a rolled refresh token remains valid in seconds.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"refresh_token_rolling_grace_period_type": {
				Description: `When specified, it overrides the global Refresh Token Grace Period defined in the Authorization Server Settings. The default value is SERVER_DEFAULT`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"refresh_token_rolling_interval": {
				Description: `The minimum interval to roll refresh tokens, in hours. This value will override the Refresh Token Rolling Interval Value on the Authorization Server Settings.`,
				Optional:    true,
				Type:        types.NumberType,
			},
			"refresh_token_rolling_interval_type": {
				Description: `Use OVERRIDE_SERVER_DEFAULT to override the Refresh Token Rolling Interval value on the Authorization Server Settings. SERVER_DEFAULT will default to the Refresh Token Rolling Interval value on the Authorization Server Setting. Defaults to SERVER_DEFAULT.`,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.StringValue("SERVER_DEFAULT")),
				},
			},
			"request_object_signing_algorithm": {
				Description: `The JSON Web Signature [JWS] algorithm that must be used to sign the Request Object. All signing algorithms are allowed if value is not present <br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"request_policy_ref": {
				Description: `The CIBA request policy.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"require_jwt_secured_authorization_response_mode": {
				Description: `Determines whether JWT Secured authorization response mode is required when initiating an authorization request. The default is false.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"require_proof_key_for_code_exchange": {
				Description: `Determines whether Proof Key for Code Exchange (PKCE) is required for this client.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"require_pushed_authorization_requests": {
				Description: `Determines whether pushed authorization requests are required when initiating an authorization request. The default is false.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"require_signed_requests": {
				Description: `Determines whether signed requests are required for this client`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"restrict_scopes": {
				Description: `Restricts this client's access to specific scopes.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"restrict_to_default_access_token_manager": {
				Description: `Determines whether the client is restricted to using only its default access token manager. The default is false.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
			"restricted_response_types": {
				Description: `The response types allowed for this client. If omitted all response types are available to the client.`,
				Optional:    true,
				Computed:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.ListValueMust(types.StringType, []attr.Value{})),
				},
			},
			"restricted_scopes": {
				Description: `The scopes available for this client.`,
				Optional:    true,
				Computed:    true,
				Type: types.ListType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.ListValueMust(types.StringType, []attr.Value{})),
				},
			},
			"token_exchange_processor_policy_ref": {
				Description: `The Token Exchange Processor policy.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"token_introspection_content_encryption_algorithm": {
				Description: `The JSON Web Encryption [JWE] content-encryption algorithm for the Token Introspection Response.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES_128_GCM - AES-GCM-128<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256`,
				Optional:    true,
				Type:        types.StringType,
			},
			"token_introspection_encryption_algorithm": {
				Description: `The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content-encryption key of the Token Introspection Response.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256`,
				Optional:    true,
				Type:        types.StringType,
			},
			"token_introspection_signing_algorithm": {
				Description: `The JSON Web Signature [JWS] algorithm required to sign the Token Introspection Response.<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11`,
				Optional:    true,
				Type:        types.StringType,
			},
			"user_authorization_url_override": {
				Description: `The URL used as 'verification_url' and 'verification_url_complete' values in a Device Authorization request. This property overrides the 'userAuthorizationUrl' value present in Authorization Server Settings.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"validate_using_all_eligible_atms": {
				Description: `Validates token using all eligible access token managers for the client. This setting is ignored if 'restrictToDefaultAccessTokenManager' is set to true.`,
				Optional:    true,
				Computed:    true,
				Type:        types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					Default(types.BoolValue(false)),
				},
			},
		},
	}
}

func resourceRedirectValidationSettings() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `Settings for redirect validation for SSO, SLO and IdP discovery.`,
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description: ``,
				Optional:    true,
				Computed:    true,
				Type:        types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			"redirect_validation_local_settings": {
				Description: `Settings for local redirect validation.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleRedirectValidationLocalSettings()),
			},
			"redirect_validation_partner_settings": {
				Description: `Settings for redirection at a partner site.`,
				Optional:    true,
				Attributes:  tfsdk.SingleNestedAttributes(singleRedirectValidationPartnerSettings()),
			},
		},
	}
}

func listJdbcAttributeSource() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  tfsdk.MapNestedAttributes(mapAttributeFulfillmentValue()),
		},
		"column_names": {
			Description: `A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore.`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("JDBC")),
			},
		},
		"filter": {
			Description: `The JDBC WHERE clause used to query your data store to locate a user record.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"schema": {
			Description: `Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"table": {
			Description: `The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func listLdapAttributeSource() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  tfsdk.MapNestedAttributes(mapAttributeFulfillmentValue()),
		},
		"base_dn": {
			Description: `The base DN to search from. If not specified, the search will start at the LDAP's root.`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("")),
			},
		},
		"binary_attribute_settings": {
			Description: `The advanced settings for binary LDAP attributes.`,
			Optional:    true,
			Computed:    true,
			Attributes:  tfsdk.MapNestedAttributes(mapBinaryLdapAttributeSettings()),
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.MapValueMust(types.ObjectType{AttrTypes: map[string]attr.Type{"binary_encoding": types.StringType}}, map[string]attr.Value{})),
			},
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("JDBC")),
				Default(types.StringValue("JDBC")),
			},
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"member_of_nested_group": {
			Description: `Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"search_attributes": {
			Description: `A list of LDAP attributes returned from search and available for mapping.`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
		"search_filter": {
			Description: `The LDAP filter that will be used to lookup the objects from the directory.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"search_scope": {
			Description: `Determines the node depth of the query.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func listCustomAttributeSource() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  tfsdk.MapNestedAttributes(mapAttributeFulfillmentValue()),
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("JDBC")),
				Default(types.StringValue("JDBC")),
				Default(types.StringValue("JDBC")),
			},
		},
		"filter_fields": {
			Description: `The list of fields that can be used to filter a request to the custom data store.`,
			Optional:    true,
			Attributes:  tfsdk.ListNestedAttributes(listFieldEntry()),
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func mapAttributeFulfillmentValue() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"source": {
			Description: `The attribute value source.`,
			Required:    true,
			Attributes:  tfsdk.SingleNestedAttributes(singleSourceTypeIdKey()),
		},
		"value": {
			Description: `The value for this attribute.`,
			Required:    true,
			Type:        types.StringType,
		},
	}
}

func mapBinaryLdapAttributeSettings() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"binary_encoding": {
			Description: `Get the encoding type for this attribute. If not specified, the default is BASE64.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func singleClientAuth() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"client_cert_issuer_dn": {
			Description: `Client TLS Certificate Issuer DN.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"client_cert_subject_dn": {
			Description: `Client TLS Certificate Subject DN.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"encrypted_secret": {
			Description: `For GET requests, this field contains the encrypted client secret, if one exists.  For POST and PUT requests, if you wish to reuse the existing secret, this field should be passed back unchanged.`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
		},
		"enforce_replay_prevention": {
			Description: `Enforce replay prevention on JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"secret": {
			Description: `Client secret for Basic Authentication.  To update the client secret, specify the plaintext value in this field.  This field will not be populated for GET requests.`,
			Optional:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				resource.UseStateForUnknown(),
			},
		},
		"token_endpoint_auth_signing_algorithm": {
			Description: `The JSON Web Signature [JWS] algorithm that must be used to sign the JSON Web Tokens. This field is applicable only for Private Key JWT Client Authentication. All signing algorithms are allowed if value is not present <br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"type": {
			Description: `Client authentication type.<br>The required field for type SECRET is secret.<br>The required fields for type CERTIFICATE are clientCertIssuerDn and clientCertSubjectDn.<br>The required field for type PRIVATE_KEY_JWT is: either jwks or jwksUrl.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func singleClientOIDCPolicy() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"grant_access_session_revocation_api": {
			Description: `Determines whether this client is allowed to access the Session Revocation API.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"grant_access_session_session_management_api": {
			Description: `Determines whether this client is allowed to access the Session Management API.`,
			Optional:    true,
			Computed:    true,
			Type:        types.BoolType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.BoolValue(false)),
			},
		},
		"id_token_content_encryption_algorithm": {
			Description: `The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES_128_GCM - AES-GCM-128<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256`,
			Optional:    true,
			Type:        types.StringType,
		},
		"id_token_encryption_algorithm": {
			Description: `The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>RSA_OAEP_256 - RSAES OAEP using SHA-256 and MGF1 with SHA-256`,
			Optional:    true,
			Type:        types.StringType,
		},
		"id_token_signing_algorithm": {
			Description: `The JSON Web Signature [JWS] algorithm required for the ID Token.<br>NONE - No signing algorithm<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("RS256")),
			},
		},
		"logout_uris": {
			Description: `A list of client logout URI's which will be invoked when a user logs out through one of PingFederate's SLO endpoints.`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
		"pairwise_identifier_user_type": {
			Description: `Determines whether the subject identifier type is pairwise.`,
			Optional:    true,
			Computed:    true,
			Type:        types.BoolType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.BoolValue(false)),
			},
		},
		"ping_access_logout_capable": {
			Description: `Set this value to true if you wish to enable client application logout, and the client is PingAccess, or its logout endpoints follow the PingAccess path convention.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"policy_group": {
			Description: `The Open ID Connect policy. A null value will represent the default policy group.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"sector_identifier_uri": {
			Description: `The URI references a file with a single JSON array of Redirect URI and JWKS URL values.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func listConditionalIssuanceCriteriaEntry() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"attribute_name": {
			Description: `The name of the attribute to use in this issuance criterion.`,
			Required:    true,
			Type:        types.StringType,
		},
		"condition": {
			Description: `The condition that will be applied to the source attribute's value and the expected value.`,
			Required:    true,
			Type:        types.StringType,
		},
		"error_result": {
			Description: `The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("")),
			},
		},
		"source": {
			Description: `The source of the attribute.`,
			Required:    true,
			Attributes:  tfsdk.SingleNestedAttributes(singleSourceTypeIdKey()),
		},
		"value": {
			Description: `The expected value of this issuance criterion.`,
			Required:    true,
			Type:        types.StringType,
		},
	}
}

func listExpressionIssuanceCriteriaEntry() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"error_result": {
			Description: `The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.`,
			Optional:    true,
			Computed:    true,
			Type:        types.StringType,
			PlanModifiers: tfsdk.AttributePlanModifiers{
				Default(types.StringValue("")),
			},
		},
		"expression": {
			Description: `The OGNL expression to evaluate.`,
			Required:    true,
			Type:        types.StringType,
		},
	}
}

func listFieldEntry() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"name": {
			Description: `The name of this field.`,
			Required:    true,
			Type:        types.StringType,
		},
		"value": {
			Description: `The value of this field. Whether or not the value is required will be determined by plugin validation checks.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func singleIssuanceCriteria() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"conditional_criteria": {
			Description: `A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue.`,
			Optional:    true,
			Attributes:  tfsdk.ListNestedAttributes(listConditionalIssuanceCriteriaEntry()),
		},
		"expression_criteria": {
			Description: `A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue.`,
			Optional:    true,
			Attributes:  tfsdk.ListNestedAttributes(listExpressionIssuanceCriteriaEntry()),
		},
	}
}

func singleJwksSettings() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"jwks": {
			Description: `JSON Web Key Set (JWKS) document of the OAuth client. Either 'jwks' or 'jwksUrl' must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"jwks_url": {
			Description: `JSON Web Key Set (JWKS) URL of the OAuth client. Either 'jwks' or 'jwksUrl' must be provided if private key JWT client authentication or signed requests is enabled.  If the client signs its JWTs using an RSASSA-PSS signing algorithm, PingFederate must either use Java 11 or be integrated with a hardware security module (HSM) to process the digital signatures.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func mapParameterValues() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"values": {
			Description: `A List of values`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

func singleRedirectValidationLocalSettings() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"enable_in_error_resource_validation": {
			Description: `Enable validation for error resource.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"enable_target_resource_validation_for_idp_discovery": {
			Description: `Enable target resource validation for IdP discovery.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"enable_target_resource_validation_for_slo": {
			Description: `Enable target resource validation for SLO.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"enable_target_resource_validation_for_sso": {
			Description: `Enable target resource validation for SSO.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"white_list": {
			Description: `List of URLs that are designated as valid target resources.`,
			Optional:    true,
			Attributes:  tfsdk.ListNestedAttributes(listRedirectValidationSettingsWhitelistEntry()),
		},
	}
}

func singleRedirectValidationPartnerSettings() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"enable_wreply_validation_slo": {
			Description: `Enable wreply validation for SLO.`,
			Optional:    true,
			Type:        types.BoolType,
		},
	}
}

func listRedirectValidationSettingsWhitelistEntry() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"allow_query_and_fragment": {
			Description: `Allow any query parameters and fragment in the resource.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"idp_discovery": {
			Description: `Enable this target resource for IdP discovery validation.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"in_error_resource": {
			Description: `Enable this target resource for in error resource validation.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"require_https": {
			Description: `Require HTTPS for accessing this resource.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"target_resource_slo": {
			Description: `Enable this target resource for SLO redirect validation.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"target_resource_sso": {
			Description: `Enable this target resource for SSO redirect validation.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"valid_domain": {
			Description: `Domain of a valid resource.`,
			Required:    true,
			Type:        types.StringType,
		},
		"valid_path": {
			Description: `Path of a valid resource.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}
}

func singleSourceTypeIdKey() map[string]tfsdk.Attribute {
	return map[string]tfsdk.Attribute{
		"id": {
			Description: `The attribute source ID that refers to the attribute source that this key references. In some resources, the ID is optional and will be ignored. In these cases the ID should be omitted. If the source type is not an attribute source then the ID can be omitted.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"type": {
			Description: `The source type of this key.`,
			Required:    true,
			Type:        types.StringType,
		},
	}
}
