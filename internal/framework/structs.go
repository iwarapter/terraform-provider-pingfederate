package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessTokenAttributeData struct {
	MultiValued types.Bool   `tfsdk:"multi_valued"`
	Name        types.String `tfsdk:"name"`
}

type AccessTokenAttributeContractData struct {
	CoreAttributes          *[]*AccessTokenAttributeData `tfsdk:"core_attributes"`
	DefaultSubjectAttribute types.String                 `tfsdk:"default_subject_attribute"`
	ExtendedAttributes      *[]*AccessTokenAttributeData `tfsdk:"extended_attributes"`
	Inherited               types.Bool                   `tfsdk:"inherited"`
}

type AccessTokenManagementSettingsData struct {
	DefaultAccessTokenManagerRef types.String `tfsdk:"default_access_token_manager_ref"`
}

type AccessTokenManagerData struct {
	AccessControlSettings     *AtmAccessControlSettingsData     `tfsdk:"access_control_settings"`
	AttributeContract         *AccessTokenAttributeContractData `tfsdk:"attribute_contract"`
	Configuration             *PluginConfigurationData          `tfsdk:"configuration"`
	Id                        types.String                      `tfsdk:"id"`
	Name                      types.String                      `tfsdk:"name"`
	ParentRef                 types.String                      `tfsdk:"parent_ref"`
	PluginDescriptorRef       types.String                      `tfsdk:"plugin_descriptor_ref"`
	SelectionSettings         *AtmSelectionSettingsData         `tfsdk:"selection_settings"`
	SequenceNumber            types.Number                      `tfsdk:"sequence_number"`
	SessionValidationSettings *SessionValidationSettingsData    `tfsdk:"session_validation_settings"`
}

type AccessTokenManagerMappingData struct {
	AccessTokenManagerRef        types.String                              `tfsdk:"access_token_manager_ref"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type AccessTokenMappingData struct {
	AccessTokenManagerRef        types.String                              `tfsdk:"access_token_manager_ref"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Context                      *AccessTokenMappingContextData            `tfsdk:"context"`
	Id                           types.String                              `tfsdk:"id"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type AccessTokenMappingContextData struct {
	ContextRef types.String `tfsdk:"context_ref"`
	Type       types.String `tfsdk:"type"`
}

type AccountManagementSettingsData struct {
	AccountStatusAlgorithm     types.String `tfsdk:"account_status_algorithm"`
	AccountStatusAttributeName types.String `tfsdk:"account_status_attribute_name"`
	DefaultStatus              types.Bool   `tfsdk:"default_status"`
	FlagComparisonStatus       types.Bool   `tfsdk:"flag_comparison_status"`
	FlagComparisonValue        types.String `tfsdk:"flag_comparison_value"`
}

type ActionDescriptorData struct {
	Description         types.String            `tfsdk:"description"`
	Download            types.Bool              `tfsdk:"download"`
	DownloadContentType types.String            `tfsdk:"download_content_type"`
	DownloadFileName    types.String            `tfsdk:"download_file_name"`
	Name                types.String            `tfsdk:"name"`
	Parameters          *[]*FieldDescriptorData `tfsdk:"parameters"`
}

type ActionOptionsData struct {
	Parameters *[]*ActionParameterData `tfsdk:"parameters"`
}

type ActionParameterData struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ActionResultData struct {
	Message types.String `tfsdk:"message"`
}

type AdditionalAllowedEntitiesConfigurationData struct {
	AdditionalAllowedEntities *[]*EntityData `tfsdk:"additional_allowed_entities"`
	AllowAdditionalEntities   types.Bool     `tfsdk:"allow_additional_entities"`
	AllowAllEntities          types.Bool     `tfsdk:"allow_all_entities"`
}

type AdditionalKeySetData struct {
	Description types.String         `tfsdk:"description"`
	Id          types.String         `tfsdk:"id"`
	Issuers     *[]*ResourceLinkData `tfsdk:"issuers"`
	Name        types.String         `tfsdk:"name"`
	SigningKeys *SigningKeysData     `tfsdk:"signing_keys"`
}

type AdditionalKeySetsData struct {
	Items *[]*AdditionalKeySetData `tfsdk:"items"`
}

type AdministrativeAccountData struct {
	Active            types.Bool   `tfsdk:"active"`
	Auditor           types.Bool   `tfsdk:"auditor"`
	Department        types.String `tfsdk:"department"`
	Description       types.String `tfsdk:"description"`
	EmailAddress      types.String `tfsdk:"email_address"`
	EncryptedPassword types.String `tfsdk:"encrypted_password"`
	Password          types.String `tfsdk:"password"`
	PhoneNumber       types.String `tfsdk:"phone_number"`
	Roles             types.List   `tfsdk:"roles"`
	Username          types.String `tfsdk:"username"`
}

type AdministrativeAccountsData struct {
	Items *[]*AdministrativeAccountData `tfsdk:"items"`
}

type AlternativeLoginHintTokenIssuerData struct {
	Issuer  types.String `tfsdk:"issuer"`
	Jwks    types.String `tfsdk:"jwks"`
	JwksURL types.String `tfsdk:"jwks_url"`
}

type ApcMappingPolicyActionData struct {
	AttributeMapping                *AttributeMappingData `tfsdk:"attribute_mapping"`
	AuthenticationPolicyContractRef types.String          `tfsdk:"authentication_policy_contract_ref"`
	Context                         types.String          `tfsdk:"context"`
	Type                            types.String          `tfsdk:"type"`
}

type ApcToPersistentGrantMappingData struct {
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	Id                              types.String                              `tfsdk:"id"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type ApcToPersistentGrantMappingsData struct {
	Items *[]*ApcToPersistentGrantMappingData `tfsdk:"items"`
}

type ApcToSpAdapterMappingData struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
	Id                               types.String                              `tfsdk:"id"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
}

type ApcToSpAdapterMappingsData struct {
	Items *[]*ApcToSpAdapterMappingData `tfsdk:"items"`
}

type ApiResponseData struct {
}

type ApiResultData struct {
	DeveloperMessage types.String            `tfsdk:"developer_message"`
	Message          types.String            `tfsdk:"message"`
	ResultId         types.String            `tfsdk:"result_id"`
	ValidationErrors *[]*ValidationErrorData `tfsdk:"validation_errors"`
}

type ApplicationSessionPolicyData struct {
	Id              types.String `tfsdk:"id"`
	IdleTimeoutMins types.Number `tfsdk:"idle_timeout_mins"`
	MaxTimeoutMins  types.Number `tfsdk:"max_timeout_mins"`
}

type ArtifactResolverLocationData struct {
	Index types.Number `tfsdk:"index"`
	Url   types.String `tfsdk:"url"`
}

type ArtifactSettingsData struct {
	Lifetime          types.Number                     `tfsdk:"lifetime"`
	ResolverLocations *[]*ArtifactResolverLocationData `tfsdk:"resolver_locations"`
	SourceId          types.String                     `tfsdk:"source_id"`
}

type AssertionLifetimeData struct {
	MinutesAfter  types.Number `tfsdk:"minutes_after"`
	MinutesBefore types.Number `tfsdk:"minutes_before"`
}

type AtmAccessControlSettingsData struct {
	AllowedClients  *[]*ResourceLinkData `tfsdk:"allowed_clients"`
	Inherited       types.Bool           `tfsdk:"inherited"`
	RestrictClients types.Bool           `tfsdk:"restrict_clients"`
}

type AtmSelectionSettingsData struct {
	Inherited    types.Bool `tfsdk:"inherited"`
	ResourceUris types.List `tfsdk:"resource_uris"`
}

type AttributeData struct {
	Name types.String `tfsdk:"name"`
}

type AttributeFulfillmentValueData struct {
	Source *SourceTypeIdKeyData `tfsdk:"source"`
	Value  types.String         `tfsdk:"value"`
}

type AttributeMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type AttributeQueryNameMappingData struct {
	LocalName  types.String `tfsdk:"local_name"`
	RemoteName types.String `tfsdk:"remote_name"`
}

type AttributeRuleData struct {
	AttributeName types.String `tfsdk:"attribute_name"`
	Condition     types.String `tfsdk:"condition"`
	ExpectedValue types.String `tfsdk:"expected_value"`
	Result        types.String `tfsdk:"result"`
}

type AttributeRulesData struct {
	FallbackToSuccess types.Bool            `tfsdk:"fallback_to_success"`
	Items             *[]*AttributeRuleData `tfsdk:"items"`
}

type AttributeSourceData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	Description                  types.String                              `tfsdk:"description"`
	Id                           types.String                              `tfsdk:"id"`
	Type                         types.String                              `tfsdk:"type"`
}

type AuthenticationPoliciesSettingsData struct {
	EnableIdpAuthnSelection types.Bool `tfsdk:"enable_idp_authn_selection"`
	EnableSpAuthnSelection  types.Bool `tfsdk:"enable_sp_authn_selection"`
}

type AuthenticationPolicyData struct {
	AuthnSelectionTrees          *[]*AuthenticationPolicyTreeData `tfsdk:"authn_selection_trees"`
	DefaultAuthenticationSources *[]*AuthenticationSourceData     `tfsdk:"default_authentication_sources"`
	FailIfNoSelection            types.Bool                       `tfsdk:"fail_if_no_selection"`
	TrackedHttpParameters        types.List                       `tfsdk:"tracked_http_parameters"`
}

type AuthenticationPolicyContractData struct {
	CoreAttributes     types.Set    `tfsdk:"core_attributes"`
	ExtendedAttributes types.Set    `tfsdk:"extended_attributes"`
	Id                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
}

type AuthenticationPolicyContractAssertionMappingData struct {
	AbortSsoTransactionAsFailSafe   types.Bool                                `tfsdk:"abort_sso_transaction_as_fail_safe"`
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictVirtualEntityIds        types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds      types.List                                `tfsdk:"restricted_virtual_entity_ids"`
}

type AuthenticationPolicyContractAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type AuthenticationPolicyContractMappingData struct {
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictVirtualServerIds        types.Bool                                `tfsdk:"restrict_virtual_server_ids"`
	RestrictedVirtualServerIds      types.List                                `tfsdk:"restricted_virtual_server_ids"`
}

type AuthenticationPolicyContractsData struct {
	Items *[]*AuthenticationPolicyContractData `tfsdk:"items"`
}

type AuthenticationPolicyFragmentData struct {
	Description types.String                      `tfsdk:"description"`
	Id          types.String                      `tfsdk:"id"`
	Inputs      types.String                      `tfsdk:"inputs"`
	Name        types.String                      `tfsdk:"name"`
	Outputs     types.String                      `tfsdk:"outputs"`
	RootNode    *AuthenticationPolicyTreeNodeData `tfsdk:"root_node"`
}

type AuthenticationPolicyFragmentsData struct {
	Items *[]*AuthenticationPolicyFragmentData `tfsdk:"items"`
}

type AuthenticationPolicyTreeData struct {
	AuthenticationApiApplicationRef types.String                      `tfsdk:"authentication_api_application_ref"`
	Description                     types.String                      `tfsdk:"description"`
	Enabled                         types.Bool                        `tfsdk:"enabled"`
	HandleFailuresLocally           types.Bool                        `tfsdk:"handle_failures_locally"`
	Id                              types.String                      `tfsdk:"id"`
	Name                            types.String                      `tfsdk:"name"`
	RootNode                        *AuthenticationPolicyTreeNodeData `tfsdk:"root_node"`
}

type AuthenticationPolicyTreeNodeData struct {
	Action   *PolicyActionData                    `tfsdk:"action"`
	Children *[]*AuthenticationPolicyTreeNodeData `tfsdk:"children"`
}

type AuthenticationSelectorData struct {
	AttributeContract   *AuthenticationSelectorAttributeContractData `tfsdk:"attribute_contract"`
	Configuration       *PluginConfigurationData                     `tfsdk:"configuration"`
	Id                  types.String                                 `tfsdk:"id"`
	Name                types.String                                 `tfsdk:"name"`
	ParentRef           types.String                                 `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                                 `tfsdk:"plugin_descriptor_ref"`
}

type AuthenticationSelectorAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type AuthenticationSelectorAttributeContractData struct {
	ExtendedAttributes *[]*AuthenticationSelectorAttributeData `tfsdk:"extended_attributes"`
}

type AuthenticationSelectorsData struct {
	Items *[]*AuthenticationSelectorData `tfsdk:"items"`
}

type AuthenticationSessionPoliciesData struct {
	Items *[]*AuthenticationSessionPolicyData `tfsdk:"items"`
}

type AuthenticationSessionPolicyData struct {
	AuthenticationSource  *AuthenticationSourceData `tfsdk:"authentication_source"`
	AuthnContextSensitive types.Bool                `tfsdk:"authn_context_sensitive"`
	EnableSessions        types.Bool                `tfsdk:"enable_sessions"`
	Id                    types.String              `tfsdk:"id"`
	IdleTimeoutMins       types.Number              `tfsdk:"idle_timeout_mins"`
	MaxTimeoutMins        types.Number              `tfsdk:"max_timeout_mins"`
	Persistent            types.Bool                `tfsdk:"persistent"`
	TimeoutDisplayUnit    types.String              `tfsdk:"timeout_display_unit"`
}

type AuthenticationSourceData struct {
	SourceRef types.String `tfsdk:"source_ref"`
	Type      types.String `tfsdk:"type"`
}

type AuthnApiApplicationData struct {
	AdditionalAllowedOrigins     types.List   `tfsdk:"additional_allowed_origins"`
	ClientForRedirectlessModeRef types.String `tfsdk:"client_for_redirectless_mode_ref"`
	Description                  types.String `tfsdk:"description"`
	Id                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	Url                          types.String `tfsdk:"url"`
}

type AuthnApiApplicationsData struct {
	Items *[]*AuthnApiApplicationData `tfsdk:"items"`
}

type AuthnApiSettingsData struct {
	ApiEnabled                       types.Bool   `tfsdk:"api_enabled"`
	DefaultApplicationRef            types.String `tfsdk:"default_application_ref"`
	EnableApiDescriptions            types.Bool   `tfsdk:"enable_api_descriptions"`
	IncludeRequestContext            types.Bool   `tfsdk:"include_request_context"`
	RestrictAccessToRedirectlessMode types.Bool   `tfsdk:"restrict_access_to_redirectless_mode"`
}

type AuthnContextMappingData struct {
	Local  types.String `tfsdk:"local"`
	Remote types.String `tfsdk:"remote"`
}

type AuthnSelectorPolicyActionData struct {
	AuthenticationSelectorRef types.String `tfsdk:"authentication_selector_ref"`
	Context                   types.String `tfsdk:"context"`
	Type                      types.String `tfsdk:"type"`
}

type AuthnSourcePolicyActionData struct {
	AttributeRules       *AttributeRulesData            `tfsdk:"attribute_rules"`
	AuthenticationSource *AuthenticationSourceData      `tfsdk:"authentication_source"`
	Context              types.String                   `tfsdk:"context"`
	InputUserIdMapping   *AttributeFulfillmentValueData `tfsdk:"input_user_id_mapping"`
	Type                 types.String                   `tfsdk:"type"`
	UserIdAuthenticated  types.Bool                     `tfsdk:"user_id_authenticated"`
}

type AuthorizationServerSettingsData struct {
	ActivationCodeCheckMode                     types.String                 `tfsdk:"activation_code_check_mode"`
	AdminWebServicePcvRef                       types.String                 `tfsdk:"admin_web_service_pcv_ref"`
	AllowUnidentifiedClientExtensionGrants      types.Bool                   `tfsdk:"allow_unidentified_client_extension_grants"`
	AllowUnidentifiedClientROCreds              types.Bool                   `tfsdk:"allow_unidentified_client_ro_creds"`
	AllowedOrigins                              types.List                   `tfsdk:"allowed_origins"`
	ApprovedScopesAttribute                     types.String                 `tfsdk:"approved_scopes_attribute"`
	AtmIdForOAuthGrantManagement                types.String                 `tfsdk:"atm_id_for_o_auth_grant_management"`
	AuthorizationCodeEntropy                    types.Number                 `tfsdk:"authorization_code_entropy"`
	AuthorizationCodeTimeout                    types.Number                 `tfsdk:"authorization_code_timeout"`
	BypassActivationCodeConfirmation            types.Bool                   `tfsdk:"bypass_activation_code_confirmation"`
	BypassAuthorizationForApprovedGrants        types.Bool                   `tfsdk:"bypass_authorization_for_approved_grants"`
	ClientSecretRetentionPeriod                 types.Number                 `tfsdk:"client_secret_retention_period"`
	DefaultScopeDescription                     types.String                 `tfsdk:"default_scope_description"`
	DevicePollingInterval                       types.Number                 `tfsdk:"device_polling_interval"`
	DisallowPlainPKCE                           types.Bool                   `tfsdk:"disallow_plain_pkce"`
	ExclusiveScopeGroups                        *[]*ScopeGroupEntryData      `tfsdk:"exclusive_scope_groups"`
	ExclusiveScopes                             *[]*ScopeEntryData           `tfsdk:"exclusive_scopes"`
	Id                                          types.String                 `tfsdk:"id"`
	IncludeIssuerInAuthorizationResponse        types.Bool                   `tfsdk:"include_issuer_in_authorization_response"`
	JwtSecuredAuthorizationResponseModeLifetime types.Number                 `tfsdk:"jwt_secured_authorization_response_mode_lifetime"`
	ParReferenceLength                          types.Number                 `tfsdk:"par_reference_length"`
	ParReferenceTimeout                         types.Number                 `tfsdk:"par_reference_timeout"`
	ParStatus                                   types.String                 `tfsdk:"par_status"`
	PendingAuthorizationTimeout                 types.Number                 `tfsdk:"pending_authorization_timeout"`
	PersistentGrantContract                     *PersistentGrantContractData `tfsdk:"persistent_grant_contract"`
	PersistentGrantIdleTimeout                  types.Number                 `tfsdk:"persistent_grant_idle_timeout"`
	PersistentGrantIdleTimeoutTimeUnit          types.String                 `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	PersistentGrantLifetime                     types.Number                 `tfsdk:"persistent_grant_lifetime"`
	PersistentGrantLifetimeUnit                 types.String                 `tfsdk:"persistent_grant_lifetime_unit"`
	PersistentGrantReuseGrantTypes              types.List                   `tfsdk:"persistent_grant_reuse_grant_types"`
	RefreshRollingInterval                      types.Number                 `tfsdk:"refresh_rolling_interval"`
	RefreshTokenLength                          types.Number                 `tfsdk:"refresh_token_length"`
	RefreshTokenRollingGracePeriod              types.Number                 `tfsdk:"refresh_token_rolling_grace_period"`
	RegisteredAuthorizationPath                 types.String                 `tfsdk:"registered_authorization_path"`
	RollRefreshTokenValues                      types.Bool                   `tfsdk:"roll_refresh_token_values"`
	ScopeForOAuthGrantManagement                types.String                 `tfsdk:"scope_for_o_auth_grant_management"`
	ScopeGroups                                 *[]*ScopeGroupEntryData      `tfsdk:"scope_groups"`
	Scopes                                      *[]*ScopeEntryData           `tfsdk:"scopes"`
	TokenEndpointBaseUrl                        types.String                 `tfsdk:"token_endpoint_base_url"`
	TrackUserSessionsForLogout                  types.Bool                   `tfsdk:"track_user_sessions_for_logout"`
	UserAuthorizationConsentAdapter             types.String                 `tfsdk:"user_authorization_consent_adapter"`
	UserAuthorizationConsentPageSetting         types.String                 `tfsdk:"user_authorization_consent_page_setting"`
	UserAuthorizationUrl                        types.String                 `tfsdk:"user_authorization_url"`
}

type BackChannelAuthData struct {
	DigitalSignature     types.Bool                       `tfsdk:"digital_signature"`
	HttpBasicCredentials *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	Type                 types.String                     `tfsdk:"type"`
}

type BaseDefaultValueLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type BaseProviderRoleData struct {
	Enable        types.Bool `tfsdk:"enable"`
	EnableSaml10  types.Bool `tfsdk:"enable_saml10"`
	EnableSaml11  types.Bool `tfsdk:"enable_saml11"`
	EnableWsFed   types.Bool `tfsdk:"enable_ws_fed"`
	EnableWsTrust types.Bool `tfsdk:"enable_ws_trust"`
}

type BaseSelectionLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	Options               types.List            `tfsdk:"options"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type BaseSigningSettingsData struct {
	Algorithm                types.String `tfsdk:"algorithm"`
	IncludeCertInSignature   types.Bool   `tfsdk:"include_cert_in_signature"`
	IncludeRawKeyInSignature types.Bool   `tfsdk:"include_raw_key_in_signature"`
	SigningKeyPairRef        types.String `tfsdk:"signing_key_pair_ref"`
}

type BinaryLdapAttributeSettingsData struct {
	BinaryEncoding types.String `tfsdk:"binary_encoding"`
}

type CSRResponseData struct {
	FileData types.String `tfsdk:"file_data"`
}

type CaptchaSettingsData struct {
	EncryptedSecretKey types.String `tfsdk:"encrypted_secret_key"`
	SecretKey          types.String `tfsdk:"secret_key"`
	SiteKey            types.String `tfsdk:"site_key"`
}

type CertViewData struct {
	CryptoProvider          types.String `tfsdk:"crypto_provider"`
	Expires                 types.String `tfsdk:"expires"`
	Id                      types.String `tfsdk:"id"`
	IssuerDN                types.String `tfsdk:"issuer_dn"`
	KeyAlgorithm            types.String `tfsdk:"key_algorithm"`
	KeySize                 types.Number `tfsdk:"key_size"`
	SerialNumber            types.String `tfsdk:"serial_number"`
	Sha1Fingerprint         types.String `tfsdk:"sha1fingerprint"`
	Sha256Fingerprint       types.String `tfsdk:"sha256fingerprint"`
	SignatureAlgorithm      types.String `tfsdk:"signature_algorithm"`
	Status                  types.String `tfsdk:"status"`
	SubjectAlternativeNames types.List   `tfsdk:"subject_alternative_names"`
	SubjectDN               types.String `tfsdk:"subject_dn"`
	ValidFrom               types.String `tfsdk:"valid_from"`
	Version                 types.Number `tfsdk:"version"`
}

type CertViewsData struct {
	Items *[]*CertViewData `tfsdk:"items"`
}

type CertificateExpirationNotificationSettingsData struct {
	EmailAddress             types.String `tfsdk:"email_address"`
	FinalWarningPeriod       types.Number `tfsdk:"final_warning_period"`
	InitialWarningPeriod     types.Number `tfsdk:"initial_warning_period"`
	NotificationPublisherRef types.String `tfsdk:"notification_publisher_ref"`
}

type CertificateRevocationSettingsData struct {
	CrlSettings   *CrlSettingsData   `tfsdk:"crl_settings"`
	OcspSettings  *OcspSettingsData  `tfsdk:"ocsp_settings"`
	ProxySettings *ProxySettingsData `tfsdk:"proxy_settings"`
}

type ChangeDetectionSettingsData struct {
	ChangedUsersAlgorithm  types.String `tfsdk:"changed_users_algorithm"`
	GroupObjectClass       types.String `tfsdk:"group_object_class"`
	TimeStampAttributeName types.String `tfsdk:"time_stamp_attribute_name"`
	UserObjectClass        types.String `tfsdk:"user_object_class"`
	UsnAttributeName       types.String `tfsdk:"usn_attribute_name"`
}

type ChannelData struct {
	Active           types.Bool                   `tfsdk:"active"`
	AttributeMapping *[]*SaasAttributeMappingData `tfsdk:"attribute_mapping"`
	ChannelSource    *ChannelSourceData           `tfsdk:"channel_source"`
	MaxThreads       types.Number                 `tfsdk:"max_threads"`
	Name             types.String                 `tfsdk:"name"`
	Timeout          types.Number                 `tfsdk:"timeout"`
}

type ChannelSourceData struct {
	AccountManagementSettings *AccountManagementSettingsData `tfsdk:"account_management_settings"`
	BaseDn                    types.String                   `tfsdk:"base_dn"`
	ChangeDetectionSettings   *ChangeDetectionSettingsData   `tfsdk:"change_detection_settings"`
	DataSource                types.String                   `tfsdk:"data_source"`
	GroupMembershipDetection  *GroupMembershipDetectionData  `tfsdk:"group_membership_detection"`
	GroupSourceLocation       *ChannelSourceLocationData     `tfsdk:"group_source_location"`
	GuidAttributeName         types.String                   `tfsdk:"guid_attribute_name"`
	GuidBinary                types.Bool                     `tfsdk:"guid_binary"`
	UserSourceLocation        *ChannelSourceLocationData     `tfsdk:"user_source_location"`
}

type ChannelSourceLocationData struct {
	Filter       types.String `tfsdk:"filter"`
	GroupDN      types.String `tfsdk:"group_dn"`
	NestedSearch types.Bool   `tfsdk:"nested_search"`
}

type CheckboxGroupLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	Options               types.List            `tfsdk:"options"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type CheckboxLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type CibaServerPolicySettingsData struct {
	DefaultRequestPolicyRef types.String `tfsdk:"default_request_policy_ref"`
}

type ClientData struct {
	AllowAuthenticationApiInit                                    types.Bool                      `tfsdk:"allow_authentication_api_init"`
	BypassActivationCodeConfirmationOverride                      types.Bool                      `tfsdk:"bypass_activation_code_confirmation_override"`
	BypassApprovalPage                                            types.Bool                      `tfsdk:"bypass_approval_page"`
	CibaDeliveryMode                                              types.String                    `tfsdk:"ciba_delivery_mode"`
	CibaNotificationEndpoint                                      types.String                    `tfsdk:"ciba_notification_endpoint"`
	CibaPollingInterval                                           types.Number                    `tfsdk:"ciba_polling_interval"`
	CibaRequestObjectSigningAlgorithm                             types.String                    `tfsdk:"ciba_request_object_signing_algorithm"`
	CibaRequireSignedRequests                                     types.Bool                      `tfsdk:"ciba_require_signed_requests"`
	CibaUserCodeSupported                                         types.Bool                      `tfsdk:"ciba_user_code_supported"`
	ClientAuth                                                    *ClientAuthData                 `tfsdk:"client_auth"`
	ClientId                                                      types.String                    `tfsdk:"client_id"`
	ClientSecretChangedTime                                       types.String                    `tfsdk:"client_secret_changed_time"`
	ClientSecretRetentionPeriod                                   types.Number                    `tfsdk:"client_secret_retention_period"`
	ClientSecretRetentionPeriodType                               types.String                    `tfsdk:"client_secret_retention_period_type"`
	DefaultAccessTokenManagerRef                                  types.String                    `tfsdk:"default_access_token_manager_ref"`
	Description                                                   types.String                    `tfsdk:"description"`
	DeviceFlowSettingType                                         types.String                    `tfsdk:"device_flow_setting_type"`
	DevicePollingIntervalOverride                                 types.Number                    `tfsdk:"device_polling_interval_override"`
	Enabled                                                       types.Bool                      `tfsdk:"enabled"`
	ExclusiveScopes                                               types.List                      `tfsdk:"exclusive_scopes"`
	ExtendedParameters                                            map[string]*ParameterValuesData `tfsdk:"extended_parameters"`
	GrantTypes                                                    types.List                      `tfsdk:"grant_types"`
	Id                                                            types.String                    `tfsdk:"id"`
	JwksSettings                                                  *JwksSettingsData               `tfsdk:"jwks_settings"`
	JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm types.String                    `tfsdk:"jwt_secured_authorization_response_mode_content_encryption_algorithm"`
	JwtSecuredAuthorizationResponseModeEncryptionAlgorithm        types.String                    `tfsdk:"jwt_secured_authorization_response_mode_encryption_algorithm"`
	JwtSecuredAuthorizationResponseModeSigningAlgorithm           types.String                    `tfsdk:"jwt_secured_authorization_response_mode_signing_algorithm"`
	LogoUrl                                                       types.String                    `tfsdk:"logo_url"`
	Name                                                          types.String                    `tfsdk:"name"`
	OidcPolicy                                                    *ClientOIDCPolicyData           `tfsdk:"oidc_policy"`
	PendingAuthorizationTimeoutOverride                           types.Number                    `tfsdk:"pending_authorization_timeout_override"`
	PersistentGrantExpirationTime                                 types.Number                    `tfsdk:"persistent_grant_expiration_time"`
	PersistentGrantExpirationTimeUnit                             types.String                    `tfsdk:"persistent_grant_expiration_time_unit"`
	PersistentGrantExpirationType                                 types.String                    `tfsdk:"persistent_grant_expiration_type"`
	PersistentGrantIdleTimeout                                    types.Number                    `tfsdk:"persistent_grant_idle_timeout"`
	PersistentGrantIdleTimeoutTimeUnit                            types.String                    `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	PersistentGrantIdleTimeoutType                                types.String                    `tfsdk:"persistent_grant_idle_timeout_type"`
	PersistentGrantReuseGrantTypes                                types.List                      `tfsdk:"persistent_grant_reuse_grant_types"`
	PersistentGrantReuseType                                      types.String                    `tfsdk:"persistent_grant_reuse_type"`
	RedirectUris                                                  types.Set                       `tfsdk:"redirect_uris"`
	RefreshRolling                                                types.String                    `tfsdk:"refresh_rolling"`
	RefreshTokenRollingGracePeriod                                types.Number                    `tfsdk:"refresh_token_rolling_grace_period"`
	RefreshTokenRollingGracePeriodType                            types.String                    `tfsdk:"refresh_token_rolling_grace_period_type"`
	RefreshTokenRollingInterval                                   types.Number                    `tfsdk:"refresh_token_rolling_interval"`
	RefreshTokenRollingIntervalType                               types.String                    `tfsdk:"refresh_token_rolling_interval_type"`
	RequestObjectSigningAlgorithm                                 types.String                    `tfsdk:"request_object_signing_algorithm"`
	RequestPolicyRef                                              types.String                    `tfsdk:"request_policy_ref"`
	RequireJwtSecuredAuthorizationResponseMode                    types.Bool                      `tfsdk:"require_jwt_secured_authorization_response_mode"`
	RequireProofKeyForCodeExchange                                types.Bool                      `tfsdk:"require_proof_key_for_code_exchange"`
	RequirePushedAuthorizationRequests                            types.Bool                      `tfsdk:"require_pushed_authorization_requests"`
	RequireSignedRequests                                         types.Bool                      `tfsdk:"require_signed_requests"`
	RestrictScopes                                                types.Bool                      `tfsdk:"restrict_scopes"`
	RestrictToDefaultAccessTokenManager                           types.Bool                      `tfsdk:"restrict_to_default_access_token_manager"`
	RestrictedResponseTypes                                       types.List                      `tfsdk:"restricted_response_types"`
	RestrictedScopes                                              types.Set                       `tfsdk:"restricted_scopes"`
	TokenExchangeProcessorPolicyRef                               types.String                    `tfsdk:"token_exchange_processor_policy_ref"`
	TokenIntrospectionContentEncryptionAlgorithm                  types.String                    `tfsdk:"token_introspection_content_encryption_algorithm"`
	TokenIntrospectionEncryptionAlgorithm                         types.String                    `tfsdk:"token_introspection_encryption_algorithm"`
	TokenIntrospectionSigningAlgorithm                            types.String                    `tfsdk:"token_introspection_signing_algorithm"`
	UserAuthorizationUrlOverride                                  types.String                    `tfsdk:"user_authorization_url_override"`
	ValidateUsingAllEligibleAtms                                  types.Bool                      `tfsdk:"validate_using_all_eligible_atms"`
}

type ClientAuthData struct {
	ClientCertIssuerDn                types.String `tfsdk:"client_cert_issuer_dn"`
	ClientCertSubjectDn               types.String `tfsdk:"client_cert_subject_dn"`
	EncryptedSecret                   types.String `tfsdk:"encrypted_secret"`
	EnforceReplayPrevention           types.Bool   `tfsdk:"enforce_replay_prevention"`
	Secret                            types.String `tfsdk:"secret"`
	TokenEndpointAuthSigningAlgorithm types.String `tfsdk:"token_endpoint_auth_signing_algorithm"`
	Type                              types.String `tfsdk:"type"`
}

type ClientMetadataData struct {
	Description types.String `tfsdk:"description"`
	MultiValued types.Bool   `tfsdk:"multi_valued"`
	Parameter   types.String `tfsdk:"parameter"`
}

type ClientOIDCPolicyData struct {
	GrantAccessSessionRevocationApi        types.Bool   `tfsdk:"grant_access_session_revocation_api"`
	GrantAccessSessionSessionManagementApi types.Bool   `tfsdk:"grant_access_session_session_management_api"`
	IdTokenContentEncryptionAlgorithm      types.String `tfsdk:"id_token_content_encryption_algorithm"`
	IdTokenEncryptionAlgorithm             types.String `tfsdk:"id_token_encryption_algorithm"`
	IdTokenSigningAlgorithm                types.String `tfsdk:"id_token_signing_algorithm"`
	LogoutUris                             types.List   `tfsdk:"logout_uris"`
	PairwiseIdentifierUserType             types.Bool   `tfsdk:"pairwise_identifier_user_type"`
	PingAccessLogoutCapable                types.Bool   `tfsdk:"ping_access_logout_capable"`
	PolicyGroup                            types.String `tfsdk:"policy_group"`
	SectorIdentifierUri                    types.String `tfsdk:"sector_identifier_uri"`
}

type ClientRegistrationOIDCPolicyData struct {
	IdTokenContentEncryptionAlgorithm types.String `tfsdk:"id_token_content_encryption_algorithm"`
	IdTokenEncryptionAlgorithm        types.String `tfsdk:"id_token_encryption_algorithm"`
	IdTokenSigningAlgorithm           types.String `tfsdk:"id_token_signing_algorithm"`
	PolicyGroup                       types.String `tfsdk:"policy_group"`
}

type ClientRegistrationPoliciesData struct {
	Items *[]*ClientRegistrationPolicyData `tfsdk:"items"`
}

type ClientRegistrationPolicyData struct {
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
}

type ClientSecretData struct {
	EncryptedSecret  types.String            `tfsdk:"encrypted_secret"`
	SecondarySecrets *[]*SecondarySecretData `tfsdk:"secondary_secrets"`
	Secret           types.String            `tfsdk:"secret"`
}

type ClientSettingsData struct {
	ClientMetadata            *[]*ClientMetadataData         `tfsdk:"client_metadata"`
	DynamicClientRegistration *DynamicClientRegistrationData `tfsdk:"dynamic_client_registration"`
}

type ClientsData struct {
	Items *[]*ClientData `tfsdk:"items"`
}

type ClusterNodeData struct {
	Address                types.String `tfsdk:"address"`
	ConfigurationTimestamp types.String `tfsdk:"configuration_timestamp"`
	Index                  types.Number `tfsdk:"index"`
	Mode                   types.String `tfsdk:"mode"`
	NodeGroup              types.String `tfsdk:"node_group"`
	NodeTags               types.String `tfsdk:"node_tags"`
	ReplicationStatus      types.String `tfsdk:"replication_status"`
	Version                types.String `tfsdk:"version"`
}

type ClusterStatusData struct {
	LastConfigUpdateTime types.String        `tfsdk:"last_config_update_time"`
	LastReplicationTime  types.String        `tfsdk:"last_replication_time"`
	MixedMode            types.Bool          `tfsdk:"mixed_mode"`
	Nodes                *[]*ClusterNodeData `tfsdk:"nodes"`
	ReplicationRequired  types.Bool          `tfsdk:"replication_required"`
}

type ConditionalIssuanceCriteriaEntryData struct {
	AttributeName types.String         `tfsdk:"attribute_name"`
	Condition     types.String         `tfsdk:"condition"`
	ErrorResult   types.String         `tfsdk:"error_result"`
	Source        *SourceTypeIdKeyData `tfsdk:"source"`
	Value         types.String         `tfsdk:"value"`
}

type ConfigFieldData struct {
	EncryptedValue types.String `tfsdk:"encrypted_value"`
	Inherited      types.Bool   `tfsdk:"inherited"`
	Name           types.String `tfsdk:"name"`
	Value          types.String `tfsdk:"value"`
}

type ConfigRowData struct {
	DefaultRow types.Bool          `tfsdk:"default_row"`
	Fields     *[]*ConfigFieldData `tfsdk:"fields"`
}

type ConfigStoreBundleData struct {
	Items *[]*ConfigStoreSettingData `tfsdk:"items"`
}

type ConfigStoreSettingData struct {
	Id          types.String            `tfsdk:"id"`
	ListValue   types.List              `tfsdk:"list_value"`
	MapValue    map[string]types.String `tfsdk:"map_value"`
	StringValue types.String            `tfsdk:"string_value"`
	Type        types.String            `tfsdk:"type"`
}

type ConfigTableData struct {
	Inherited types.Bool        `tfsdk:"inherited"`
	Name      types.String      `tfsdk:"name"`
	Rows      *[]*ConfigRowData `tfsdk:"rows"`
}

type ConfigurationEncryptionKeyData struct {
	CreationDate types.String `tfsdk:"creation_date"`
	KeyId        types.String `tfsdk:"key_id"`
}

type ConfigurationEncryptionKeysData struct {
	Items *[]*ConfigurationEncryptionKeyData `tfsdk:"items"`
}

type ConnectionData struct {
	Active                                 types.Bool                                  `tfsdk:"active"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	Id                                     types.String                                `tfsdk:"id"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	Name                                   types.String                                `tfsdk:"name"`
	Type                                   types.String                                `tfsdk:"type"`
	VirtualEntityIds                       types.List                                  `tfsdk:"virtual_entity_ids"`
}

type ConnectionCertData struct {
	ActiveVerificationCert    types.Bool    `tfsdk:"active_verification_cert"`
	CertView                  *CertViewData `tfsdk:"cert_view"`
	EncryptionCert            types.Bool    `tfsdk:"encryption_cert"`
	PrimaryVerificationCert   types.Bool    `tfsdk:"primary_verification_cert"`
	SecondaryVerificationCert types.Bool    `tfsdk:"secondary_verification_cert"`
	X509File                  *X509FileData `tfsdk:"x509file"`
}

type ConnectionCertsData struct {
	Items *[]*ConnectionCertData `tfsdk:"items"`
}

type ConnectionCredentialsData struct {
	BlockEncryptionAlgorithm      types.String                 `tfsdk:"block_encryption_algorithm"`
	Certs                         *[]*ConnectionCertData       `tfsdk:"certs"`
	DecryptionKeyPairRef          types.String                 `tfsdk:"decryption_key_pair_ref"`
	InboundBackChannelAuth        *InboundBackChannelAuthData  `tfsdk:"inbound_back_channel_auth"`
	KeyTransportAlgorithm         types.String                 `tfsdk:"key_transport_algorithm"`
	OutboundBackChannelAuth       *OutboundBackChannelAuthData `tfsdk:"outbound_back_channel_auth"`
	SecondaryDecryptionKeyPairRef types.String                 `tfsdk:"secondary_decryption_key_pair_ref"`
	SigningSettings               *SigningSettingsData         `tfsdk:"signing_settings"`
	VerificationIssuerDN          types.String                 `tfsdk:"verification_issuer_dn"`
	VerificationSubjectDN         types.String                 `tfsdk:"verification_subject_dn"`
}

type ConnectionGroupLicenseViewData struct {
	ConnectionCount types.Number `tfsdk:"connection_count"`
	EndDate         types.String `tfsdk:"end_date"`
	Name            types.String `tfsdk:"name"`
	StartDate       types.String `tfsdk:"start_date"`
}

type ConnectionMetadataUrlData struct {
	EnableAutoMetadataUpdate types.Bool   `tfsdk:"enable_auto_metadata_update"`
	MetadataUrlRef           types.String `tfsdk:"metadata_url_ref"`
}

type ContactInfoData struct {
	Company   types.String `tfsdk:"company"`
	Email     types.String `tfsdk:"email"`
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
	Phone     types.String `tfsdk:"phone"`
}

type ContinuePolicyActionData struct {
	Context types.String `tfsdk:"context"`
	Type    types.String `tfsdk:"type"`
}

type ConvertMetadataRequestData struct {
	ConnectionType          types.String    `tfsdk:"connection_type"`
	ExpectedEntityId        types.String    `tfsdk:"expected_entity_id"`
	ExpectedProtocol        types.String    `tfsdk:"expected_protocol"`
	SamlMetadata            types.String    `tfsdk:"saml_metadata"`
	TemplateConnection      *ConnectionData `tfsdk:"template_connection"`
	VerificationCertificate types.String    `tfsdk:"verification_certificate"`
}

type ConvertMetadataResponseData struct {
	CertExpiration   types.String    `tfsdk:"cert_expiration"`
	CertSerialNumber types.String    `tfsdk:"cert_serial_number"`
	CertSubjectDn    types.String    `tfsdk:"cert_subject_dn"`
	CertTrustStatus  types.String    `tfsdk:"cert_trust_status"`
	Connection       *ConnectionData `tfsdk:"connection"`
	SignatureStatus  types.String    `tfsdk:"signature_status"`
}

type CrlSettingsData struct {
	NextRetryMinsWhenNextUpdateInPast types.Number `tfsdk:"next_retry_mins_when_next_update_in_past"`
	NextRetryMinsWhenResolveFailed    types.Number `tfsdk:"next_retry_mins_when_resolve_failed"`
	TreatNonRetrievableCrlAsRevoked   types.Bool   `tfsdk:"treat_non_retrievable_crl_as_revoked"`
	VerifyCrlSignature                types.Bool   `tfsdk:"verify_crl_signature"`
}

type CustomAttributeSourceData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	Description                  types.String                              `tfsdk:"description"`
	FilterFields                 *[]*FieldEntryData                        `tfsdk:"filter_fields"`
	Id                           types.String                              `tfsdk:"id"`
}

type CustomDataStoreData struct {
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	MaskAttributeValues types.Bool               `tfsdk:"mask_attribute_values"`
	Name                types.String             `tfsdk:"name"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	Type                types.String             `tfsdk:"type"`
}

type DataStoreData struct {
	Id                  types.String `tfsdk:"id"`
	MaskAttributeValues types.Bool   `tfsdk:"mask_attribute_values"`
	Type                types.String `tfsdk:"type"`
}

type DataStoreAttributeData struct {
	Metadata map[string]types.String `tfsdk:"metadata"`
	Name     types.String            `tfsdk:"name"`
	Type     types.String            `tfsdk:"type"`
}

type DataStoreConfigData struct {
	DataStoreMapping map[string]*DataStoreAttributeData `tfsdk:"data_store_mapping"`
	DataStoreRef     types.String                       `tfsdk:"data_store_ref"`
	Type             types.String                       `tfsdk:"type"`
}

type DataStoreRepositoryData struct {
	DataStoreRef                  types.String                              `tfsdk:"data_store_ref"`
	JitRepositoryAttributeMapping map[string]*AttributeFulfillmentValueData `tfsdk:"jit_repository_attribute_mapping"`
	Type                          types.String                              `tfsdk:"type"`
}

type DataStoresData struct {
	Items *[]*DataStoreData `tfsdk:"items"`
}

type DateLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type DecryptionKeysData struct {
	PrimaryKeyRef       types.String `tfsdk:"primary_key_ref"`
	SecondaryKeyPairRef types.String `tfsdk:"secondary_key_pair_ref"`
}

type DecryptionPolicyData struct {
	AssertionEncrypted        types.Bool `tfsdk:"assertion_encrypted"`
	AttributesEncrypted       types.Bool `tfsdk:"attributes_encrypted"`
	SloEncryptSubjectNameID   types.Bool `tfsdk:"slo_encrypt_subject_name_id"`
	SloSubjectNameIDEncrypted types.Bool `tfsdk:"slo_subject_name_id_encrypted"`
	SubjectNameIdEncrypted    types.Bool `tfsdk:"subject_name_id_encrypted"`
}

type DonePolicyActionData struct {
	Context types.String `tfsdk:"context"`
	Type    types.String `tfsdk:"type"`
}

type DropDownLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	Options               types.List            `tfsdk:"options"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type DynamicClientRegistrationData struct {
	AllowClientDelete                          types.Bool                        `tfsdk:"allow_client_delete"`
	AllowedExclusiveScopes                     types.List                        `tfsdk:"allowed_exclusive_scopes"`
	BypassActivationCodeConfirmationOverride   types.Bool                        `tfsdk:"bypass_activation_code_confirmation_override"`
	CibaPollingInterval                        types.Number                      `tfsdk:"ciba_polling_interval"`
	CibaRequireSignedRequests                  types.Bool                        `tfsdk:"ciba_require_signed_requests"`
	ClientCertIssuerRef                        types.String                      `tfsdk:"client_cert_issuer_ref"`
	ClientCertIssuerType                       types.String                      `tfsdk:"client_cert_issuer_type"`
	ClientSecretRetentionPeriodOverride        types.Number                      `tfsdk:"client_secret_retention_period_override"`
	ClientSecretRetentionPeriodType            types.String                      `tfsdk:"client_secret_retention_period_type"`
	DefaultAccessTokenManagerRef               types.String                      `tfsdk:"default_access_token_manager_ref"`
	DeviceFlowSettingType                      types.String                      `tfsdk:"device_flow_setting_type"`
	DevicePollingIntervalOverride              types.Number                      `tfsdk:"device_polling_interval_override"`
	DisableRegistrationAccessTokens            types.Bool                        `tfsdk:"disable_registration_access_tokens"`
	EnforceReplayPrevention                    types.Bool                        `tfsdk:"enforce_replay_prevention"`
	InitialAccessTokenScope                    types.String                      `tfsdk:"initial_access_token_scope"`
	OidcPolicy                                 *ClientRegistrationOIDCPolicyData `tfsdk:"oidc_policy"`
	PendingAuthorizationTimeoutOverride        types.Number                      `tfsdk:"pending_authorization_timeout_override"`
	PersistentGrantExpirationTime              types.Number                      `tfsdk:"persistent_grant_expiration_time"`
	PersistentGrantExpirationTimeUnit          types.String                      `tfsdk:"persistent_grant_expiration_time_unit"`
	PersistentGrantExpirationType              types.String                      `tfsdk:"persistent_grant_expiration_type"`
	PersistentGrantIdleTimeout                 types.Number                      `tfsdk:"persistent_grant_idle_timeout"`
	PersistentGrantIdleTimeoutTimeUnit         types.String                      `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	PersistentGrantIdleTimeoutType             types.String                      `tfsdk:"persistent_grant_idle_timeout_type"`
	PolicyRefs                                 *[]*ResourceLinkData              `tfsdk:"policy_refs"`
	RefreshRolling                             types.String                      `tfsdk:"refresh_rolling"`
	RefreshTokenRollingGracePeriod             types.Number                      `tfsdk:"refresh_token_rolling_grace_period"`
	RefreshTokenRollingGracePeriodType         types.String                      `tfsdk:"refresh_token_rolling_grace_period_type"`
	RefreshTokenRollingInterval                types.Number                      `tfsdk:"refresh_token_rolling_interval"`
	RefreshTokenRollingIntervalType            types.String                      `tfsdk:"refresh_token_rolling_interval_type"`
	RequestPolicyRef                           types.String                      `tfsdk:"request_policy_ref"`
	RequireJwtSecuredAuthorizationResponseMode types.Bool                        `tfsdk:"require_jwt_secured_authorization_response_mode"`
	RequireProofKeyForCodeExchange             types.Bool                        `tfsdk:"require_proof_key_for_code_exchange"`
	RequireSignedRequests                      types.Bool                        `tfsdk:"require_signed_requests"`
	RestrictCommonScopes                       types.Bool                        `tfsdk:"restrict_common_scopes"`
	RestrictToDefaultAccessTokenManager        types.Bool                        `tfsdk:"restrict_to_default_access_token_manager"`
	RestrictedCommonScopes                     types.List                        `tfsdk:"restricted_common_scopes"`
	RetainClientSecret                         types.Bool                        `tfsdk:"retain_client_secret"`
	RotateClientSecret                         types.Bool                        `tfsdk:"rotate_client_secret"`
	RotateRegistrationAccessToken              types.Bool                        `tfsdk:"rotate_registration_access_token"`
	TokenExchangeProcessorPolicyRef            types.String                      `tfsdk:"token_exchange_processor_policy_ref"`
	UserAuthorizationUrlOverride               types.String                      `tfsdk:"user_authorization_url_override"`
}

type EmailLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type EmailServerSettingsData struct {
	EmailServer              types.String `tfsdk:"email_server"`
	EnableUtf8MessageHeaders types.Bool   `tfsdk:"enable_utf8message_headers"`
	EncryptedPassword        types.String `tfsdk:"encrypted_password"`
	Password                 types.String `tfsdk:"password"`
	Port                     types.Number `tfsdk:"port"`
	RetryAttempts            types.Number `tfsdk:"retry_attempts"`
	RetryDelay               types.Number `tfsdk:"retry_delay"`
	SourceAddr               types.String `tfsdk:"source_addr"`
	SslPort                  types.Number `tfsdk:"ssl_port"`
	Timeout                  types.Number `tfsdk:"timeout"`
	UseDebugging             types.Bool   `tfsdk:"use_debugging"`
	UseSSL                   types.Bool   `tfsdk:"use_ssl"`
	UseTLS                   types.Bool   `tfsdk:"use_tls"`
	Username                 types.String `tfsdk:"username"`
	VerifyHostname           types.Bool   `tfsdk:"verify_hostname"`
}

type EmailVerificationConfigData struct {
	AllowedOtpCharacterSet               types.String `tfsdk:"allowed_otp_character_set"`
	EmailVerificationEnabled             types.Bool   `tfsdk:"email_verification_enabled"`
	EmailVerificationErrorTemplateName   types.String `tfsdk:"email_verification_error_template_name"`
	EmailVerificationOtpTemplateName     types.String `tfsdk:"email_verification_otp_template_name"`
	EmailVerificationSentTemplateName    types.String `tfsdk:"email_verification_sent_template_name"`
	EmailVerificationSuccessTemplateName types.String `tfsdk:"email_verification_success_template_name"`
	EmailVerificationType                types.String `tfsdk:"email_verification_type"`
	FieldForEmailToVerify                types.String `tfsdk:"field_for_email_to_verify"`
	FieldStoringVerificationStatus       types.String `tfsdk:"field_storing_verification_status"`
	NotificationPublisherRef             types.String `tfsdk:"notification_publisher_ref"`
	OtlTimeToLive                        types.Number `tfsdk:"otl_time_to_live"`
	OtpLength                            types.Number `tfsdk:"otp_length"`
	OtpRetryAttempts                     types.Number `tfsdk:"otp_retry_attempts"`
	OtpTimeToLive                        types.Number `tfsdk:"otp_time_to_live"`
	RequireVerifiedEmail                 types.Bool   `tfsdk:"require_verified_email"`
	RequireVerifiedEmailTemplateName     types.String `tfsdk:"require_verified_email_template_name"`
	VerifyEmailTemplateName              types.String `tfsdk:"verify_email_template_name"`
}

type EncryptionPolicyData struct {
	EncryptAssertion          types.Bool `tfsdk:"encrypt_assertion"`
	EncryptSloSubjectNameId   types.Bool `tfsdk:"encrypt_slo_subject_name_id"`
	EncryptedAttributes       types.List `tfsdk:"encrypted_attributes"`
	SloSubjectNameIDEncrypted types.Bool `tfsdk:"slo_subject_name_id_encrypted"`
}

type EntityData struct {
	EntityDescription types.String `tfsdk:"entity_description"`
	EntityId          types.String `tfsdk:"entity_id"`
}

type ExportMetadataRequestData struct {
	ConnectionId            types.String             `tfsdk:"connection_id"`
	ConnectionType          types.String             `tfsdk:"connection_type"`
	SigningSettings         *BaseSigningSettingsData `tfsdk:"signing_settings"`
	UseSecondaryPortForSoap types.Bool               `tfsdk:"use_secondary_port_for_soap"`
	VirtualHostName         types.String             `tfsdk:"virtual_host_name"`
	VirtualServerId         types.String             `tfsdk:"virtual_server_id"`
}

type ExpressionIssuanceCriteriaEntryData struct {
	ErrorResult types.String `tfsdk:"error_result"`
	Expression  types.String `tfsdk:"expression"`
}

type ExtendedPropertiesData struct {
	Items *[]*ExtendedPropertyData `tfsdk:"items"`
}

type ExtendedPropertyData struct {
	Description types.String `tfsdk:"description"`
	MultiValued types.Bool   `tfsdk:"multi_valued"`
	Name        types.String `tfsdk:"name"`
}

type FederationInfoData struct {
	AutoConnectEntityId types.String `tfsdk:"auto_connect_entity_id"`
	BaseUrl             types.String `tfsdk:"base_url"`
	Saml1xIssuerId      types.String `tfsdk:"saml1x_issuer_id"`
	Saml1xSourceId      types.String `tfsdk:"saml1x_source_id"`
	Saml2EntityId       types.String `tfsdk:"saml2entity_id"`
	WsfedRealm          types.String `tfsdk:"wsfed_realm"`
}

type FieldConfigData struct {
	Fields                    *[]*LocalIdentityFieldData `tfsdk:"fields"`
	StripSpaceFromUniqueField types.Bool                 `tfsdk:"strip_space_from_unique_field"`
}

type FieldDescriptorData struct {
	Advanced     types.Bool   `tfsdk:"advanced"`
	DefaultValue types.String `tfsdk:"default_value"`
	Description  types.String `tfsdk:"description"`
	Label        types.String `tfsdk:"label"`
	Name         types.String `tfsdk:"name"`
	Required     types.Bool   `tfsdk:"required"`
	Type         types.String `tfsdk:"type"`
}

type FieldEntryData struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type FragmentPolicyActionData struct {
	AttributeRules  *AttributeRulesData   `tfsdk:"attribute_rules"`
	Context         types.String          `tfsdk:"context"`
	Fragment        types.String          `tfsdk:"fragment"`
	FragmentMapping *AttributeMappingData `tfsdk:"fragment_mapping"`
	Type            types.String          `tfsdk:"type"`
}

type GeneralSettingsData struct {
	DatastoreValidationIntervalSecs         types.Number `tfsdk:"datastore_validation_interval_secs"`
	DisableAutomaticConnectionValidation    types.Bool   `tfsdk:"disable_automatic_connection_validation"`
	IdpConnectionTransactionLoggingOverride types.String `tfsdk:"idp_connection_transaction_logging_override"`
	RequestHeaderForCorrelationId           types.String `tfsdk:"request_header_for_correlation_id"`
	SpConnectionTransactionLoggingOverride  types.String `tfsdk:"sp_connection_transaction_logging_override"`
}

type GlobalAuthenticationSessionPolicyData struct {
	EnableSessions             types.Bool   `tfsdk:"enable_sessions"`
	HashUniqueUserKeyAttribute types.Bool   `tfsdk:"hash_unique_user_key_attribute"`
	Id                         types.String `tfsdk:"id"`
	IdleTimeoutDisplayUnit     types.String `tfsdk:"idle_timeout_display_unit"`
	IdleTimeoutMins            types.Number `tfsdk:"idle_timeout_mins"`
	MaxTimeoutDisplayUnit      types.String `tfsdk:"max_timeout_display_unit"`
	MaxTimeoutMins             types.Number `tfsdk:"max_timeout_mins"`
	PersistentSessions         types.Bool   `tfsdk:"persistent_sessions"`
}

type GroupAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type GroupMembershipDetectionData struct {
	GroupMemberAttributeName   types.String `tfsdk:"group_member_attribute_name"`
	MemberOfGroupAttributeName types.String `tfsdk:"member_of_group_attribute_name"`
}

type GroupsData struct {
	ReadGroups  *ReadGroupsData  `tfsdk:"read_groups"`
	WriteGroups *WriteGroupsData `tfsdk:"write_groups"`
}

type HiddenLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type IdentityHintAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type IdentityHintContractData struct {
	CoreAttributes     *[]*IdentityHintAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdentityHintAttributeData `tfsdk:"extended_attributes"`
}

type IdentityStoreInboundProvisioningUserRepositoryData struct {
	PluginDescriptorRef types.String `tfsdk:"plugin_descriptor_ref"`
	Type                types.String `tfsdk:"type"`
}

type IdentityStoreProvisionerData struct {
	AttributeContract      *IdentityStoreProvisionerAttributeContractData      `tfsdk:"attribute_contract"`
	Configuration          *PluginConfigurationData                            `tfsdk:"configuration"`
	GroupAttributeContract *IdentityStoreProvisionerGroupAttributeContractData `tfsdk:"group_attribute_contract"`
	Id                     types.String                                        `tfsdk:"id"`
	Name                   types.String                                        `tfsdk:"name"`
	ParentRef              types.String                                        `tfsdk:"parent_ref"`
	PluginDescriptorRef    types.String                                        `tfsdk:"plugin_descriptor_ref"`
}

type IdentityStoreProvisionerAttributeContractData struct {
	CoreAttributes     *[]*AttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*AttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool        `tfsdk:"inherited"`
}

type IdentityStoreProvisionerDescriptorData struct {
	AttributeContract             types.List                  `tfsdk:"attribute_contract"`
	ClassName                     types.String                `tfsdk:"class_name"`
	ConfigDescriptor              *PluginConfigDescriptorData `tfsdk:"config_descriptor"`
	GroupAttributeContract        types.List                  `tfsdk:"group_attribute_contract"`
	Id                            types.String                `tfsdk:"id"`
	Name                          types.String                `tfsdk:"name"`
	SupportsExtendedContract      types.Bool                  `tfsdk:"supports_extended_contract"`
	SupportsGroupExtendedContract types.Bool                  `tfsdk:"supports_group_extended_contract"`
}

type IdentityStoreProvisionerDescriptorsData struct {
	Items *[]*IdentityStoreProvisionerDescriptorData `tfsdk:"items"`
}

type IdentityStoreProvisionerGroupAttributeContractData struct {
	CoreAttributes     *[]*GroupAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*GroupAttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool             `tfsdk:"inherited"`
}

type IdentityStoreProvisionersData struct {
	Items *[]*IdentityStoreProvisionerData `tfsdk:"items"`
}

type IdpAdapterData struct {
	AttributeContract   *IdpAdapterAttributeContractData `tfsdk:"attribute_contract"`
	AttributeMapping    *IdpAdapterContractMappingData   `tfsdk:"attribute_mapping"`
	AuthnCtxClassRef    types.String                     `tfsdk:"authn_ctx_class_ref"`
	Configuration       *PluginConfigurationData         `tfsdk:"configuration"`
	Id                  types.String                     `tfsdk:"id"`
	Name                types.String                     `tfsdk:"name"`
	ParentRef           types.String                     `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                     `tfsdk:"plugin_descriptor_ref"`
}

type IdpAdapterAssertionMappingData struct {
	AbortSsoTransactionAsFailSafe types.Bool                                `tfsdk:"abort_sso_transaction_as_fail_safe"`
	AdapterOverrideSettings       *IdpAdapterData                           `tfsdk:"adapter_override_settings"`
	AttributeContractFulfillment  map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources          []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources          []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources        []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IdpAdapterRef                 types.String                              `tfsdk:"idp_adapter_ref"`
	IssuanceCriteria              *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictVirtualEntityIds      types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds    types.List                                `tfsdk:"restricted_virtual_entity_ids"`
}

type IdpAdapterAttributeData struct {
	Masked    types.Bool   `tfsdk:"masked"`
	Name      types.String `tfsdk:"name"`
	Pseudonym types.Bool   `tfsdk:"pseudonym"`
}

type IdpAdapterAttributeContractData struct {
	CoreAttributes         *[]*IdpAdapterAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes     *[]*IdpAdapterAttributeData `tfsdk:"extended_attributes"`
	Inherited              types.Bool                  `tfsdk:"inherited"`
	MaskOgnlValues         types.Bool                  `tfsdk:"mask_ognl_values"`
	UniqueUserKeyAttribute types.String                `tfsdk:"unique_user_key_attribute"`
}

type IdpAdapterContractMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Inherited                    types.Bool                                `tfsdk:"inherited"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type IdpAdapterMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Id                           types.String                              `tfsdk:"id"`
	IdpAdapterRef                types.String                              `tfsdk:"idp_adapter_ref"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type IdpAdapterMappingsData struct {
	Items *[]*IdpAdapterMappingData `tfsdk:"items"`
}

type IdpAdaptersData struct {
	Items *[]*IdpAdapterData `tfsdk:"items"`
}

type IdpAttributeQueryData struct {
	NameMappings *[]*AttributeQueryNameMappingData `tfsdk:"name_mappings"`
	Policy       *IdpAttributeQueryPolicyData      `tfsdk:"policy"`
	Url          types.String                      `tfsdk:"url"`
}

type IdpAttributeQueryPolicyData struct {
	EncryptNameId             types.Bool `tfsdk:"encrypt_name_id"`
	MaskAttributeValues       types.Bool `tfsdk:"mask_attribute_values"`
	RequireEncryptedAssertion types.Bool `tfsdk:"require_encrypted_assertion"`
	RequireSignedAssertion    types.Bool `tfsdk:"require_signed_assertion"`
	RequireSignedResponse     types.Bool `tfsdk:"require_signed_response"`
	SignAttributeQuery        types.Bool `tfsdk:"sign_attribute_query"`
}

type IdpBrowserSsoData struct {
	AdapterMappings                      *[]*SpAdapterMappingData                    `tfsdk:"adapter_mappings"`
	AlwaysSignArtifactResponse           types.Bool                                  `tfsdk:"always_sign_artifact_response"`
	Artifact                             *ArtifactSettingsData                       `tfsdk:"artifact"`
	AssertionsSigned                     types.Bool                                  `tfsdk:"assertions_signed"`
	AttributeContract                    *IdpBrowserSsoAttributeContractData         `tfsdk:"attribute_contract"`
	AuthenticationPolicyContractMappings *[]*AuthenticationPolicyContractMappingData `tfsdk:"authentication_policy_contract_mappings"`
	AuthnContextMappings                 *[]*AuthnContextMappingData                 `tfsdk:"authn_context_mappings"`
	DecryptionPolicy                     *DecryptionPolicyData                       `tfsdk:"decryption_policy"`
	DefaultTargetUrl                     types.String                                `tfsdk:"default_target_url"`
	EnabledProfiles                      types.List                                  `tfsdk:"enabled_profiles"`
	IdpIdentityMapping                   types.String                                `tfsdk:"idp_identity_mapping"`
	IncomingBindings                     types.List                                  `tfsdk:"incoming_bindings"`
	JitProvisioning                      *JitProvisioningData                        `tfsdk:"jit_provisioning"`
	MessageCustomizations                *[]*ProtocolMessageCustomizationData        `tfsdk:"message_customizations"`
	OauthAuthenticationPolicyContractRef types.String                                `tfsdk:"oauth_authentication_policy_contract_ref"`
	OidcProviderSettings                 *OIDCProviderSettingsData                   `tfsdk:"oidc_provider_settings"`
	Protocol                             types.String                                `tfsdk:"protocol"`
	SignAuthnRequests                    types.Bool                                  `tfsdk:"sign_authn_requests"`
	SloServiceEndpoints                  *[]*SloServiceEndpointData                  `tfsdk:"slo_service_endpoints"`
	SsoOAuthMapping                      *SsoOAuthMappingData                        `tfsdk:"sso_o_auth_mapping"`
	SsoServiceEndpoints                  *[]*IdpSsoServiceEndpointData               `tfsdk:"sso_service_endpoints"`
	UrlWhitelistEntries                  *[]*UrlWhitelistEntryData                   `tfsdk:"url_whitelist_entries"`
}

type IdpBrowserSsoAttributeData struct {
	Masked types.Bool   `tfsdk:"masked"`
	Name   types.String `tfsdk:"name"`
}

type IdpBrowserSsoAttributeContractData struct {
	CoreAttributes     *[]*IdpBrowserSsoAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdpBrowserSsoAttributeData `tfsdk:"extended_attributes"`
}

type IdpConnectionData struct {
	Active                                 types.Bool                                  `tfsdk:"active"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	AttributeQuery                         *IdpAttributeQueryData                      `tfsdk:"attribute_query"`
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	ErrorPageMsgId                         types.String                                `tfsdk:"error_page_msg_id"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	Id                                     types.String                                `tfsdk:"id"`
	IdpBrowserSso                          *IdpBrowserSsoData                          `tfsdk:"idp_browser_sso"`
	IdpOAuthGrantAttributeMapping          *IdpOAuthGrantAttributeMappingData          `tfsdk:"idp_o_auth_grant_attribute_mapping"`
	InboundProvisioning                    *IdpInboundProvisioningData                 `tfsdk:"inbound_provisioning"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	Name                                   types.String                                `tfsdk:"name"`
	OidcClientCredentials                  *OIDCClientCredentialsData                  `tfsdk:"oidc_client_credentials"`
	Type                                   types.String                                `tfsdk:"type"`
	VirtualEntityIds                       types.List                                  `tfsdk:"virtual_entity_ids"`
	WsTrust                                *IdpWsTrustData                             `tfsdk:"ws_trust"`
}

type IdpConnectionsData struct {
	Items *[]*IdpConnectionData `tfsdk:"items"`
}

type IdpDefaultUrlData struct {
	ConfirmIdpSlo    types.Bool   `tfsdk:"confirm_idp_slo"`
	IdpErrorMsg      types.String `tfsdk:"idp_error_msg"`
	IdpSloSuccessUrl types.String `tfsdk:"idp_slo_success_url"`
}

type IdpInboundProvisioningData struct {
	ActionOnDelete types.String                           `tfsdk:"action_on_delete"`
	CustomSchema   *SchemaData                            `tfsdk:"custom_schema"`
	GroupSupport   types.Bool                             `tfsdk:"group_support"`
	Groups         *GroupsData                            `tfsdk:"groups"`
	UserRepository *InboundProvisioningUserRepositoryData `tfsdk:"user_repository"`
	Users          *UsersData                             `tfsdk:"users"`
}

type IdpInboundProvisioningAttributeData struct {
	Masked types.Bool   `tfsdk:"masked"`
	Name   types.String `tfsdk:"name"`
}

type IdpInboundProvisioningAttributeContractData struct {
	CoreAttributes     *[]*IdpInboundProvisioningAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdpInboundProvisioningAttributeData `tfsdk:"extended_attributes"`
}

type IdpOAuthAttributeContractData struct {
	CoreAttributes     *[]*IdpBrowserSsoAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdpBrowserSsoAttributeData `tfsdk:"extended_attributes"`
}

type IdpOAuthGrantAttributeMappingData struct {
	AccessTokenManagerMappings *[]*AccessTokenManagerMappingData `tfsdk:"access_token_manager_mappings"`
	IdpOAuthAttributeContract  *IdpOAuthAttributeContractData    `tfsdk:"idp_o_auth_attribute_contract"`
}

type IdpRoleData struct {
	Enable                     types.Bool         `tfsdk:"enable"`
	EnableOutboundProvisioning types.Bool         `tfsdk:"enable_outbound_provisioning"`
	EnableSaml10               types.Bool         `tfsdk:"enable_saml10"`
	EnableSaml11               types.Bool         `tfsdk:"enable_saml11"`
	EnableWsFed                types.Bool         `tfsdk:"enable_ws_fed"`
	EnableWsTrust              types.Bool         `tfsdk:"enable_ws_trust"`
	Saml20Profile              *SAML20ProfileData `tfsdk:"saml20profile"`
}

type IdpSsoServiceEndpointData struct {
	Binding types.String `tfsdk:"binding"`
	Url     types.String `tfsdk:"url"`
}

type IdpToSpAdapterMappingData struct {
	ApplicationIconUrl               types.String                              `tfsdk:"application_icon_url"`
	ApplicationName                  types.String                              `tfsdk:"application_name"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
	Id                               types.String                              `tfsdk:"id"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
}

type IdpToSpAdapterMappingsData struct {
	Items *[]*IdpToSpAdapterMappingData `tfsdk:"items"`
}

type IdpTokenProcessorMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IdpTokenProcessorRef         types.String                              `tfsdk:"idp_token_processor_ref"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictedVirtualEntityIds   types.List                                `tfsdk:"restricted_virtual_entity_ids"`
}

type IdpWsTrustData struct {
	AttributeContract      *IdpWsTrustAttributeContractData `tfsdk:"attribute_contract"`
	GenerateLocalToken     types.Bool                       `tfsdk:"generate_local_token"`
	TokenGeneratorMappings *[]*SpTokenGeneratorMappingData  `tfsdk:"token_generator_mappings"`
}

type IdpWsTrustAttributeData struct {
	Masked types.Bool   `tfsdk:"masked"`
	Name   types.String `tfsdk:"name"`
}

type IdpWsTrustAttributeContractData struct {
	CoreAttributes     *[]*IdpWsTrustAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdpWsTrustAttributeData `tfsdk:"extended_attributes"`
}

type InboundBackChannelAuthData struct {
	Certs                 *[]*ConnectionCertData           `tfsdk:"certs"`
	DigitalSignature      types.Bool                       `tfsdk:"digital_signature"`
	HttpBasicCredentials  *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	RequireSsl            types.Bool                       `tfsdk:"require_ssl"`
	Type                  types.String                     `tfsdk:"type"`
	VerificationIssuerDN  types.String                     `tfsdk:"verification_issuer_dn"`
	VerificationSubjectDN types.String                     `tfsdk:"verification_subject_dn"`
}

type InboundProvisioningUserRepositoryData struct {
	Type types.String `tfsdk:"type"`
}

type IncomingProxySettingsData struct {
	ClientCertChainSSLHeaderName  types.String `tfsdk:"client_cert_chain_ssl_header_name"`
	ClientCertSSLHeaderName       types.String `tfsdk:"client_cert_ssl_header_name"`
	ForwardedHostHeaderIndex      types.String `tfsdk:"forwarded_host_header_index"`
	ForwardedHostHeaderName       types.String `tfsdk:"forwarded_host_header_name"`
	ForwardedIpAddressHeaderIndex types.String `tfsdk:"forwarded_ip_address_header_index"`
	ForwardedIpAddressHeaderName  types.String `tfsdk:"forwarded_ip_address_header_name"`
	ProxyTerminatesHttpsConns     types.Bool   `tfsdk:"proxy_terminates_https_conns"`
}

type IssuanceCriteriaData struct {
	ConditionalCriteria *[]*ConditionalIssuanceCriteriaEntryData `tfsdk:"conditional_criteria"`
	ExpressionCriteria  *[]*ExpressionIssuanceCriteriaEntryData  `tfsdk:"expression_criteria"`
}

type IssuerData struct {
	Description types.String `tfsdk:"description"`
	Host        types.String `tfsdk:"host"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Path        types.String `tfsdk:"path"`
}

type IssuerCertData struct {
	Active   types.Bool    `tfsdk:"active"`
	CertView *CertViewData `tfsdk:"cert_view"`
	X509File *X509FileData `tfsdk:"x509file"`
}

type IssuerCertsData struct {
	Items *[]*IssuerCertData `tfsdk:"items"`
}

type IssuersData struct {
	Items *[]*IssuerData `tfsdk:"items"`
}

type JdbcAttributeSourceData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	ColumnNames                  types.List                                `tfsdk:"column_names"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	Description                  types.String                              `tfsdk:"description"`
	Filter                       types.String                              `tfsdk:"filter"`
	Id                           types.String                              `tfsdk:"id"`
	Schema                       types.String                              `tfsdk:"schema"`
	Table                        types.String                              `tfsdk:"table"`
}

type JdbcDataStoreData struct {
	AllowMultiValueAttributes types.Bool            `tfsdk:"allow_multi_value_attributes"`
	BlockingTimeout           types.Number          `tfsdk:"blocking_timeout"`
	ConnectionUrl             types.String          `tfsdk:"connection_url"`
	ConnectionUrlTags         *[]*JdbcTagConfigData `tfsdk:"connection_url_tags"`
	DriverClass               types.String          `tfsdk:"driver_class"`
	EncryptedPassword         types.String          `tfsdk:"encrypted_password"`
	Id                        types.String          `tfsdk:"id"`
	IdleTimeout               types.Number          `tfsdk:"idle_timeout"`
	MaskAttributeValues       types.Bool            `tfsdk:"mask_attribute_values"`
	MaxPoolSize               types.Number          `tfsdk:"max_pool_size"`
	MinPoolSize               types.Number          `tfsdk:"min_pool_size"`
	Name                      types.String          `tfsdk:"name"`
	Password                  types.String          `tfsdk:"password"`
	Type                      types.String          `tfsdk:"type"`
	UserName                  types.String          `tfsdk:"user_name"`
	ValidateConnectionSql     types.String          `tfsdk:"validate_connection_sql"`
}

type JdbcDataStoreRepositoryData struct {
	DataStoreRef                  types.String                              `tfsdk:"data_store_ref"`
	JitRepositoryAttributeMapping map[string]*AttributeFulfillmentValueData `tfsdk:"jit_repository_attribute_mapping"`
	SqlMethod                     *SqlMethodData                            `tfsdk:"sql_method"`
	Type                          types.String                              `tfsdk:"type"`
}

type JdbcTagConfigData struct {
	ConnectionUrl types.String `tfsdk:"connection_url"`
	DefaultSource types.Bool   `tfsdk:"default_source"`
	Tags          types.String `tfsdk:"tags"`
}

type JitProvisioningData struct {
	ErrorHandling  types.String                       `tfsdk:"error_handling"`
	EventTrigger   types.String                       `tfsdk:"event_trigger"`
	UserAttributes *JitProvisioningUserAttributesData `tfsdk:"user_attributes"`
	UserRepository *DataStoreRepositoryData           `tfsdk:"user_repository"`
}

type JitProvisioningUserAttributesData struct {
	AttributeContract *[]*IdpBrowserSsoAttributeData `tfsdk:"attribute_contract"`
	DoAttributeQuery  types.Bool                     `tfsdk:"do_attribute_query"`
}

type JwksSettingsData struct {
	Jwks    types.String `tfsdk:"jwks"`
	JwksUrl types.String `tfsdk:"jwks_url"`
}

type KerberosKeySetData struct {
	DeactivatedAt   types.String `tfsdk:"deactivated_at"`
	EncryptedKeySet types.String `tfsdk:"encrypted_key_set"`
}

type KerberosRealmData struct {
	ConnectionType                     types.String           `tfsdk:"connection_type"`
	Id                                 types.String           `tfsdk:"id"`
	KerberosEncryptedPassword          types.String           `tfsdk:"kerberos_encrypted_password"`
	KerberosPassword                   types.String           `tfsdk:"kerberos_password"`
	KerberosRealmName                  types.String           `tfsdk:"kerberos_realm_name"`
	KerberosUsername                   types.String           `tfsdk:"kerberos_username"`
	KeyDistributionCenters             types.List             `tfsdk:"key_distribution_centers"`
	KeySets                            *[]*KerberosKeySetData `tfsdk:"key_sets"`
	LdapGatewayDataStoreRef            types.String           `tfsdk:"ldap_gateway_data_store_ref"`
	RetainPreviousKeysOnPasswordChange types.Bool             `tfsdk:"retain_previous_keys_on_password_change"`
	SuppressDomainNameConcatenation    types.Bool             `tfsdk:"suppress_domain_name_concatenation"`
}

type KerberosRealmsData struct {
	Items *[]*KerberosRealmData `tfsdk:"items"`
}

type KerberosRealmsSettingsData struct {
	DebugLogOutput            types.Bool   `tfsdk:"debug_log_output"`
	ForceTcp                  types.Bool   `tfsdk:"force_tcp"`
	KdcRetries                types.String `tfsdk:"kdc_retries"`
	KdcTimeout                types.String `tfsdk:"kdc_timeout"`
	KeySetRetentionPeriodMins types.Number `tfsdk:"key_set_retention_period_mins"`
}

type KeyAlgorithmData struct {
	DefaultKeySize            types.Number `tfsdk:"default_key_size"`
	DefaultSignatureAlgorithm types.String `tfsdk:"default_signature_algorithm"`
	KeySizes                  types.List   `tfsdk:"key_sizes"`
	Name                      types.String `tfsdk:"name"`
	SignatureAlgorithms       types.List   `tfsdk:"signature_algorithms"`
}

type KeyAlgorithmsData struct {
	Items *[]*KeyAlgorithmData `tfsdk:"items"`
}

type KeyPairExportSettingsData struct {
	Password types.String `tfsdk:"password"`
}

type KeyPairFileData struct {
	CryptoProvider    types.String `tfsdk:"crypto_provider"`
	EncryptedPassword types.String `tfsdk:"encrypted_password"`
	FileData          types.String `tfsdk:"file_data"`
	Format            types.String `tfsdk:"format"`
	Id                types.String `tfsdk:"id"`
	Password          types.String `tfsdk:"password"`
}

type KeyPairRotationSettingsData struct {
	ActivationBufferDays types.Number `tfsdk:"activation_buffer_days"`
	CreationBufferDays   types.Number `tfsdk:"creation_buffer_days"`
	Id                   types.String `tfsdk:"id"`
	KeyAlgorithm         types.String `tfsdk:"key_algorithm"`
	KeySize              types.Number `tfsdk:"key_size"`
	SignatureAlgorithm   types.String `tfsdk:"signature_algorithm"`
	ValidDays            types.Number `tfsdk:"valid_days"`
}

type KeyPairViewData struct {
	CryptoProvider          types.String                 `tfsdk:"crypto_provider"`
	Expires                 types.String                 `tfsdk:"expires"`
	Id                      types.String                 `tfsdk:"id"`
	IssuerDN                types.String                 `tfsdk:"issuer_dn"`
	KeyAlgorithm            types.String                 `tfsdk:"key_algorithm"`
	KeySize                 types.Number                 `tfsdk:"key_size"`
	RotationSettings        *KeyPairRotationSettingsData `tfsdk:"rotation_settings"`
	SerialNumber            types.String                 `tfsdk:"serial_number"`
	Sha1Fingerprint         types.String                 `tfsdk:"sha1fingerprint"`
	Sha256Fingerprint       types.String                 `tfsdk:"sha256fingerprint"`
	SignatureAlgorithm      types.String                 `tfsdk:"signature_algorithm"`
	Status                  types.String                 `tfsdk:"status"`
	SubjectAlternativeNames types.List                   `tfsdk:"subject_alternative_names"`
	SubjectDN               types.String                 `tfsdk:"subject_dn"`
	ValidFrom               types.String                 `tfsdk:"valid_from"`
	Version                 types.Number                 `tfsdk:"version"`
}

type KeyPairViewsData struct {
	Items *[]*KeyPairViewData `tfsdk:"items"`
}

type LdapAttributeSourceData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData   `tfsdk:"attribute_contract_fulfillment"`
	BaseDn                       types.String                                `tfsdk:"base_dn"`
	BinaryAttributeSettings      map[string]*BinaryLdapAttributeSettingsData `tfsdk:"binary_attribute_settings"`
	DataStoreRef                 types.String                                `tfsdk:"data_store_ref"`
	Description                  types.String                                `tfsdk:"description"`
	Id                           types.String                                `tfsdk:"id"`
	MemberOfNestedGroup          types.Bool                                  `tfsdk:"member_of_nested_group"`
	SearchAttributes             types.List                                  `tfsdk:"search_attributes"`
	SearchFilter                 types.String                                `tfsdk:"search_filter"`
	SearchScope                  types.String                                `tfsdk:"search_scope"`
}

type LdapDataStoreData struct {
	BinaryAttributes     types.List            `tfsdk:"binary_attributes"`
	BindAnonymously      types.Bool            `tfsdk:"bind_anonymously"`
	ConnectionTimeout    types.Number          `tfsdk:"connection_timeout"`
	CreateIfNecessary    types.Bool            `tfsdk:"create_if_necessary"`
	DnsTtl               types.Number          `tfsdk:"dns_ttl"`
	EncryptedPassword    types.String          `tfsdk:"encrypted_password"`
	FollowLDAPReferrals  types.Bool            `tfsdk:"follow_ldap_referrals"`
	Hostnames            types.List            `tfsdk:"hostnames"`
	HostnamesTags        *[]*LdapTagConfigData `tfsdk:"hostnames_tags"`
	Id                   types.String          `tfsdk:"id"`
	LdapDnsSrvPrefix     types.String          `tfsdk:"ldap_dns_srv_prefix"`
	LdapType             types.String          `tfsdk:"ldap_type"`
	LdapsDnsSrvPrefix    types.String          `tfsdk:"ldaps_dns_srv_prefix"`
	MaskAttributeValues  types.Bool            `tfsdk:"mask_attribute_values"`
	MaxConnections       types.Number          `tfsdk:"max_connections"`
	MaxWait              types.Number          `tfsdk:"max_wait"`
	MinConnections       types.Number          `tfsdk:"min_connections"`
	Name                 types.String          `tfsdk:"name"`
	Password             types.String          `tfsdk:"password"`
	ReadTimeout          types.Number          `tfsdk:"read_timeout"`
	TestOnBorrow         types.Bool            `tfsdk:"test_on_borrow"`
	TestOnReturn         types.Bool            `tfsdk:"test_on_return"`
	TimeBetweenEvictions types.Number          `tfsdk:"time_between_evictions"`
	Type                 types.String          `tfsdk:"type"`
	UseDnsSrvRecords     types.Bool            `tfsdk:"use_dns_srv_records"`
	UseSsl               types.Bool            `tfsdk:"use_ssl"`
	UserDN               types.String          `tfsdk:"user_dn"`
	VerifyHost           types.Bool            `tfsdk:"verify_host"`
}

type LdapDataStoreAttributeData struct {
	Metadata map[string]types.String `tfsdk:"metadata"`
	Name     types.String            `tfsdk:"name"`
	Type     types.String            `tfsdk:"type"`
}

type LdapDataStoreConfigData struct {
	AuxiliaryObjectClasses types.List                         `tfsdk:"auxiliary_object_classes"`
	BaseDn                 types.String                       `tfsdk:"base_dn"`
	CreatePattern          types.String                       `tfsdk:"create_pattern"`
	DataStoreMapping       map[string]*DataStoreAttributeData `tfsdk:"data_store_mapping"`
	DataStoreRef           types.String                       `tfsdk:"data_store_ref"`
	ObjectClass            types.String                       `tfsdk:"object_class"`
	Type                   types.String                       `tfsdk:"type"`
}

type LdapDataStoreRepositoryData struct {
	BaseDn                        types.String                              `tfsdk:"base_dn"`
	DataStoreRef                  types.String                              `tfsdk:"data_store_ref"`
	JitRepositoryAttributeMapping map[string]*AttributeFulfillmentValueData `tfsdk:"jit_repository_attribute_mapping"`
	Type                          types.String                              `tfsdk:"type"`
	UniqueUserIdFilter            types.String                              `tfsdk:"unique_user_id_filter"`
}

type LdapInboundProvisioningUserRepositoryData struct {
	BaseDn              types.String `tfsdk:"base_dn"`
	DataStoreRef        types.String `tfsdk:"data_store_ref"`
	Type                types.String `tfsdk:"type"`
	UniqueGroupIdFilter types.String `tfsdk:"unique_group_id_filter"`
	UniqueUserIdFilter  types.String `tfsdk:"unique_user_id_filter"`
}

type LdapTagConfigData struct {
	DefaultSource types.Bool   `tfsdk:"default_source"`
	Hostnames     types.List   `tfsdk:"hostnames"`
	Tags          types.String `tfsdk:"tags"`
}

type LicenseAgreementInfoData struct {
	Accepted            types.Bool   `tfsdk:"accepted"`
	LicenseAgreementUrl types.String `tfsdk:"license_agreement_url"`
}

type LicenseEventNotificationSettingsData struct {
	EmailAddress             types.String `tfsdk:"email_address"`
	NotificationPublisherRef types.String `tfsdk:"notification_publisher_ref"`
}

type LicenseFeatureViewData struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type LicenseFileData struct {
	FileData types.String `tfsdk:"file_data"`
}

type LicenseViewData struct {
	BridgeMode          types.Bool                         `tfsdk:"bridge_mode"`
	EnforcementType     types.String                       `tfsdk:"enforcement_type"`
	ExpirationDate      types.String                       `tfsdk:"expiration_date"`
	Features            *[]*LicenseFeatureViewData         `tfsdk:"features"`
	GracePeriod         types.Number                       `tfsdk:"grace_period"`
	Id                  types.String                       `tfsdk:"id"`
	IssueDate           types.String                       `tfsdk:"issue_date"`
	LicenseGroups       *[]*ConnectionGroupLicenseViewData `tfsdk:"license_groups"`
	MaxConnections      types.Number                       `tfsdk:"max_connections"`
	Name                types.String                       `tfsdk:"name"`
	NodeLimit           types.Number                       `tfsdk:"node_limit"`
	OauthEnabled        types.Bool                         `tfsdk:"oauth_enabled"`
	Organization        types.String                       `tfsdk:"organization"`
	Product             types.String                       `tfsdk:"product"`
	ProvisioningEnabled types.Bool                         `tfsdk:"provisioning_enabled"`
	Tier                types.String                       `tfsdk:"tier"`
	UsedConnections     types.Number                       `tfsdk:"used_connections"`
	Version             types.String                       `tfsdk:"version"`
	WsTrustEnabled      types.Bool                         `tfsdk:"ws_trust_enabled"`
}

type LocalIdentityAuthSourceData struct {
	Id     types.String `tfsdk:"id"`
	Source types.String `tfsdk:"source"`
}

type LocalIdentityAuthSourceUpdatePolicyData struct {
	RetainAttributes types.Bool   `tfsdk:"retain_attributes"`
	StoreAttributes  types.Bool   `tfsdk:"store_attributes"`
	UpdateAttributes types.Bool   `tfsdk:"update_attributes"`
	UpdateInterval   types.Number `tfsdk:"update_interval"`
}

type LocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type LocalIdentityMappingPolicyActionData struct {
	Context                  types.String          `tfsdk:"context"`
	InboundMapping           *AttributeMappingData `tfsdk:"inbound_mapping"`
	LocalIdentityRef         types.String          `tfsdk:"local_identity_ref"`
	OutboundAttributeMapping *AttributeMappingData `tfsdk:"outbound_attribute_mapping"`
	Type                     types.String          `tfsdk:"type"`
}

type LocalIdentityProfileData struct {
	ApcId                   types.String                             `tfsdk:"apc_id"`
	AuthSourceUpdatePolicy  *LocalIdentityAuthSourceUpdatePolicyData `tfsdk:"auth_source_update_policy"`
	AuthSources             *[]*LocalIdentityAuthSourceData          `tfsdk:"auth_sources"`
	DataStoreConfig         *DataStoreConfigData                     `tfsdk:"data_store_config"`
	EmailVerificationConfig *EmailVerificationConfigData             `tfsdk:"email_verification_config"`
	FieldConfig             *FieldConfigData                         `tfsdk:"field_config"`
	Id                      types.String                             `tfsdk:"id"`
	Name                    types.String                             `tfsdk:"name"`
	ProfileConfig           *ProfileConfigData                       `tfsdk:"profile_config"`
	ProfileEnabled          types.Bool                               `tfsdk:"profile_enabled"`
	RegistrationConfig      *RegistrationConfigData                  `tfsdk:"registration_config"`
	RegistrationEnabled     types.Bool                               `tfsdk:"registration_enabled"`
}

type LocalIdentityProfilesData struct {
	Items *[]*LocalIdentityProfileData `tfsdk:"items"`
}

type MetadataEventNotificationSettingsData struct {
	EmailAddress             types.String `tfsdk:"email_address"`
	NotificationPublisherRef types.String `tfsdk:"notification_publisher_ref"`
}

type MetadataLifetimeSettingsData struct {
	CacheDuration types.Number `tfsdk:"cache_duration"`
	ReloadDelay   types.Number `tfsdk:"reload_delay"`
}

type MetadataSigningSettingsData struct {
	SignatureAlgorithm types.String `tfsdk:"signature_algorithm"`
	SigningKeyRef      types.String `tfsdk:"signing_key_ref"`
}

type MetadataUrlData struct {
	CertView          types.Object  `tfsdk:"cert_view"`
	Id                types.String  `tfsdk:"id"`
	Name              types.String  `tfsdk:"name"`
	Url               types.String  `tfsdk:"url"`
	ValidateSignature types.Bool    `tfsdk:"validate_signature"`
	X509File          *X509FileData `tfsdk:"x509file"`
}

type MetadataUrlsData struct {
	Items *[]*MetadataUrlData `tfsdk:"items"`
}

type MoveItemRequestData struct {
	Location types.String `tfsdk:"location"`
	MoveToId types.String `tfsdk:"move_to_id"`
}

type NewKeyPairSettingsData struct {
	City                    types.String `tfsdk:"city"`
	CommonName              types.String `tfsdk:"common_name"`
	Country                 types.String `tfsdk:"country"`
	CryptoProvider          types.String `tfsdk:"crypto_provider"`
	Id                      types.String `tfsdk:"id"`
	KeyAlgorithm            types.String `tfsdk:"key_algorithm"`
	KeySize                 types.Number `tfsdk:"key_size"`
	Organization            types.String `tfsdk:"organization"`
	OrganizationUnit        types.String `tfsdk:"organization_unit"`
	SignatureAlgorithm      types.String `tfsdk:"signature_algorithm"`
	State                   types.String `tfsdk:"state"`
	SubjectAlternativeNames types.List   `tfsdk:"subject_alternative_names"`
	ValidDays               types.Number `tfsdk:"valid_days"`
}

type NotificationPublisherData struct {
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
}

type NotificationPublishersData struct {
	Items *[]*NotificationPublisherData `tfsdk:"items"`
}

type NotificationPublishersSettingsData struct {
	DefaultNotificationPublisherRef types.String `tfsdk:"default_notification_publisher_ref"`
}

type NotificationSettingsData struct {
	AccountChangesNotificationPublisherRef types.String                                   `tfsdk:"account_changes_notification_publisher_ref"`
	CertificateExpirations                 *CertificateExpirationNotificationSettingsData `tfsdk:"certificate_expirations"`
	LicenseEvents                          *LicenseEventNotificationSettingsData          `tfsdk:"license_events"`
	MetadataNotificationSettings           *MetadataEventNotificationSettingsData         `tfsdk:"metadata_notification_settings"`
	NotifyAdminUserPasswordChanges         types.Bool                                     `tfsdk:"notify_admin_user_password_changes"`
}

type OAuthOidcKeysSettingsData struct {
	P256ActiveCertRef                 types.String `tfsdk:"p256active_cert_ref"`
	P256DecryptionActiveCertRef       types.String `tfsdk:"p256decryption_active_cert_ref"`
	P256DecryptionPreviousCertRef     types.String `tfsdk:"p256decryption_previous_cert_ref"`
	P256DecryptionPublishX5cParameter types.Bool   `tfsdk:"p256decryption_publish_x5c_parameter"`
	P256PreviousCertRef               types.String `tfsdk:"p256previous_cert_ref"`
	P256PublishX5cParameter           types.Bool   `tfsdk:"p256publish_x5c_parameter"`
	P384ActiveCertRef                 types.String `tfsdk:"p384active_cert_ref"`
	P384DecryptionActiveCertRef       types.String `tfsdk:"p384decryption_active_cert_ref"`
	P384DecryptionPreviousCertRef     types.String `tfsdk:"p384decryption_previous_cert_ref"`
	P384DecryptionPublishX5cParameter types.Bool   `tfsdk:"p384decryption_publish_x5c_parameter"`
	P384PreviousCertRef               types.String `tfsdk:"p384previous_cert_ref"`
	P384PublishX5cParameter           types.Bool   `tfsdk:"p384publish_x5c_parameter"`
	P521ActiveCertRef                 types.String `tfsdk:"p521active_cert_ref"`
	P521DecryptionActiveCertRef       types.String `tfsdk:"p521decryption_active_cert_ref"`
	P521DecryptionPreviousCertRef     types.String `tfsdk:"p521decryption_previous_cert_ref"`
	P521DecryptionPublishX5cParameter types.Bool   `tfsdk:"p521decryption_publish_x5c_parameter"`
	P521PreviousCertRef               types.String `tfsdk:"p521previous_cert_ref"`
	P521PublishX5cParameter           types.Bool   `tfsdk:"p521publish_x5c_parameter"`
	RsaActiveCertRef                  types.String `tfsdk:"rsa_active_cert_ref"`
	RsaDecryptionActiveCertRef        types.String `tfsdk:"rsa_decryption_active_cert_ref"`
	RsaDecryptionPreviousCertRef      types.String `tfsdk:"rsa_decryption_previous_cert_ref"`
	RsaDecryptionPublishX5cParameter  types.Bool   `tfsdk:"rsa_decryption_publish_x5c_parameter"`
	RsaPreviousCertRef                types.String `tfsdk:"rsa_previous_cert_ref"`
	RsaPublishX5cParameter            types.Bool   `tfsdk:"rsa_publish_x5c_parameter"`
	StaticJwksEnabled                 types.Bool   `tfsdk:"static_jwks_enabled"`
}

type OAuthRoleData struct {
	EnableOauth         types.Bool `tfsdk:"enable_oauth"`
	EnableOpenIdConnect types.Bool `tfsdk:"enable_open_id_connect"`
}

type OIDCClientCredentialsData struct {
	ClientId        types.String `tfsdk:"client_id"`
	ClientSecret    types.String `tfsdk:"client_secret"`
	EncryptedSecret types.String `tfsdk:"encrypted_secret"`
}

type OIDCProviderSettingsData struct {
	AuthenticationScheme           types.String                 `tfsdk:"authentication_scheme"`
	AuthenticationSigningAlgorithm types.String                 `tfsdk:"authentication_signing_algorithm"`
	AuthorizationEndpoint          types.String                 `tfsdk:"authorization_endpoint"`
	EnablePKCE                     types.Bool                   `tfsdk:"enable_pkce"`
	JwksURL                        types.String                 `tfsdk:"jwks_url"`
	LoginType                      types.String                 `tfsdk:"login_type"`
	RequestParameters              *[]*OIDCRequestParameterData `tfsdk:"request_parameters"`
	RequestSigningAlgorithm        types.String                 `tfsdk:"request_signing_algorithm"`
	Scopes                         types.String                 `tfsdk:"scopes"`
	TokenEndpoint                  types.String                 `tfsdk:"token_endpoint"`
	UserInfoEndpoint               types.String                 `tfsdk:"user_info_endpoint"`
}

type OIDCRequestParameterData struct {
	ApplicationEndpointOverride types.Bool                     `tfsdk:"application_endpoint_override"`
	AttributeValue              *AttributeFulfillmentValueData `tfsdk:"attribute_value"`
	Name                        types.String                   `tfsdk:"name"`
	Value                       types.String                   `tfsdk:"value"`
}

type OIDCSessionSettingsData struct {
	RevokeUserSessionOnLogout  types.Bool   `tfsdk:"revoke_user_session_on_logout"`
	SessionRevocationLifetime  types.Number `tfsdk:"session_revocation_lifetime"`
	TrackUserSessionsForLogout types.Bool   `tfsdk:"track_user_sessions_for_logout"`
}

type OcspSettingsData struct {
	ActionOnResponderUnavailable types.String `tfsdk:"action_on_responder_unavailable"`
	ActionOnStatusUnknown        types.String `tfsdk:"action_on_status_unknown"`
	ActionOnUnsuccessfulResponse types.String `tfsdk:"action_on_unsuccessful_response"`
	CurrentUpdateGracePeriod     types.Number `tfsdk:"current_update_grace_period"`
	NextUpdateGracePeriod        types.Number `tfsdk:"next_update_grace_period"`
	RequesterAddNonce            types.Bool   `tfsdk:"requester_add_nonce"`
	ResponderCertReference       types.String `tfsdk:"responder_cert_reference"`
	ResponderTimeout             types.Number `tfsdk:"responder_timeout"`
	ResponderUrl                 types.String `tfsdk:"responder_url"`
	ResponseCachePeriod          types.Number `tfsdk:"response_cache_period"`
}

type OpenIdConnectAttributeData struct {
	IncludeInIdToken  types.Bool   `tfsdk:"include_in_id_token"`
	IncludeInUserInfo types.Bool   `tfsdk:"include_in_user_info"`
	MultiValued       types.Bool   `tfsdk:"multi_valued"`
	Name              types.String `tfsdk:"name"`
}

type OpenIdConnectAttributeContractData struct {
	CoreAttributes     *[]*OpenIdConnectAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*OpenIdConnectAttributeData `tfsdk:"extended_attributes"`
}

type OpenIdConnectPoliciesData struct {
	Items *[]*OpenIdConnectPolicyData `tfsdk:"items"`
}

type OpenIdConnectPolicyData struct {
	AccessTokenManagerRef       types.String                        `tfsdk:"access_token_manager_ref"`
	AttributeContract           *OpenIdConnectAttributeContractData `tfsdk:"attribute_contract"`
	AttributeMapping            *AttributeMappingData               `tfsdk:"attribute_mapping"`
	Id                          types.String                        `tfsdk:"id"`
	IdTokenLifetime             types.Number                        `tfsdk:"id_token_lifetime"`
	IncludeSHashInIdToken       types.Bool                          `tfsdk:"include_s_hash_in_id_token"`
	IncludeSriInIdToken         types.Bool                          `tfsdk:"include_sri_in_id_token"`
	IncludeUserInfoInIdToken    types.Bool                          `tfsdk:"include_user_info_in_id_token"`
	Name                        types.String                        `tfsdk:"name"`
	ReissueIdTokenInHybridFlow  types.Bool                          `tfsdk:"reissue_id_token_in_hybrid_flow"`
	ReturnIdTokenOnRefreshGrant types.Bool                          `tfsdk:"return_id_token_on_refresh_grant"`
	ScopeAttributeMappings      map[string]*ParameterValuesData     `tfsdk:"scope_attribute_mappings"`
}

type OpenIdConnectSettingsData struct {
	DefaultPolicyRef types.String             `tfsdk:"default_policy_ref"`
	SessionSettings  *OIDCSessionSettingsData `tfsdk:"session_settings"`
}

type OptionValueData struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type OutOfBandAuthAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type OutOfBandAuthAttributeContractData struct {
	CoreAttributes     *[]*OutOfBandAuthAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*OutOfBandAuthAttributeData `tfsdk:"extended_attributes"`
}

type OutOfBandAuthenticatorData struct {
	AttributeContract   *OutOfBandAuthAttributeContractData `tfsdk:"attribute_contract"`
	Configuration       *PluginConfigurationData            `tfsdk:"configuration"`
	Id                  types.String                        `tfsdk:"id"`
	Name                types.String                        `tfsdk:"name"`
	ParentRef           types.String                        `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                        `tfsdk:"plugin_descriptor_ref"`
}

type OutOfBandAuthenticatorsData struct {
	Items *[]*OutOfBandAuthenticatorData `tfsdk:"items"`
}

type OutboundBackChannelAuthData struct {
	DigitalSignature     types.Bool                       `tfsdk:"digital_signature"`
	HttpBasicCredentials *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	SslAuthKeyPairRef    types.String                     `tfsdk:"ssl_auth_key_pair_ref"`
	Type                 types.String                     `tfsdk:"type"`
	ValidatePartnerCert  types.Bool                       `tfsdk:"validate_partner_cert"`
}

type OutboundProvisionData struct {
	Channels       *[]*ChannelData     `tfsdk:"channels"`
	CustomSchema   *SchemaData         `tfsdk:"custom_schema"`
	TargetSettings *[]*ConfigFieldData `tfsdk:"target_settings"`
	Type           types.String        `tfsdk:"type"`
}

type OutboundProvisionDatabaseData struct {
	DataStoreRef             types.String `tfsdk:"data_store_ref"`
	SynchronizationFrequency types.Number `tfsdk:"synchronization_frequency"`
}

type P14EKeyPairViewData struct {
	CreationTime              types.String  `tfsdk:"creation_time"`
	CurrentAuthenticationKey  types.Bool    `tfsdk:"current_authentication_key"`
	KeyPairView               *CertViewData `tfsdk:"key_pair_view"`
	PreviousAuthenticationKey types.Bool    `tfsdk:"previous_authentication_key"`
}

type P14EKeysViewData struct {
	KeyPairs *[]*P14EKeyPairViewData `tfsdk:"key_pairs"`
}

type ParameterValuesData struct {
	Values types.List `tfsdk:"values"`
}

type PasswordCredentialValidatorData struct {
	AttributeContract   *PasswordCredentialValidatorAttributeContractData `tfsdk:"attribute_contract"`
	Configuration       *PluginConfigurationData                          `tfsdk:"configuration"`
	Id                  types.String                                      `tfsdk:"id"`
	Name                types.String                                      `tfsdk:"name"`
	ParentRef           types.String                                      `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                                      `tfsdk:"plugin_descriptor_ref"`
}

type PasswordCredentialValidatorAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type PasswordCredentialValidatorAttributeContractData struct {
	CoreAttributes     *[]*PasswordCredentialValidatorAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*PasswordCredentialValidatorAttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool                                   `tfsdk:"inherited"`
}

type PasswordCredentialValidatorsData struct {
	Items *[]*PasswordCredentialValidatorData `tfsdk:"items"`
}

type PersistentGrantAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type PersistentGrantContractData struct {
	CoreAttributes     *[]*PersistentGrantAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*PersistentGrantAttributeData `tfsdk:"extended_attributes"`
}

type PhoneLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type PingOneConnectionData struct {
	Active                           types.Bool   `tfsdk:"active"`
	CreationDate                     types.String `tfsdk:"creation_date"`
	Credential                       types.String `tfsdk:"credential"`
	CredentialId                     types.String `tfsdk:"credential_id"`
	Description                      types.String `tfsdk:"description"`
	EncryptedCredential              types.String `tfsdk:"encrypted_credential"`
	EnvironmentId                    types.String `tfsdk:"environment_id"`
	Id                               types.String `tfsdk:"id"`
	Name                             types.String `tfsdk:"name"`
	OrganizationName                 types.String `tfsdk:"organization_name"`
	PingOneAuthenticationApiEndpoint types.String `tfsdk:"ping_one_authentication_api_endpoint"`
	PingOneConnectionId              types.String `tfsdk:"ping_one_connection_id"`
	PingOneManagementApiEndpoint     types.String `tfsdk:"ping_one_management_api_endpoint"`
	Region                           types.String `tfsdk:"region"`
}

type PingOneConnectionsData struct {
	Items *[]*PingOneConnectionData `tfsdk:"items"`
}

type PingOneCredentialStatusData struct {
	PingOneCredentialStatus types.String `tfsdk:"ping_one_credential_status"`
}

type PingOneEnvironmentData struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type PingOneEnvironmentsData struct {
	Items *[]*PingOneEnvironmentData `tfsdk:"items"`
}

type PingOneForEnterpriseSettingsData struct {
	CompanyName                      types.String `tfsdk:"company_name"`
	ConnectedToPingOneForEnterprise  types.Bool   `tfsdk:"connected_to_ping_one_for_enterprise"`
	CurrentAuthnKeyCreationTime      types.String `tfsdk:"current_authn_key_creation_time"`
	EnableAdminConsoleSso            types.Bool   `tfsdk:"enable_admin_console_sso"`
	EnableMonitoring                 types.Bool   `tfsdk:"enable_monitoring"`
	IdentityRepositoryUpdateRequired types.Bool   `tfsdk:"identity_repository_update_required"`
	PingOneSsoConnection             types.String `tfsdk:"ping_one_sso_connection"`
	PreviousAuthnKeyCreationTime     types.String `tfsdk:"previous_authn_key_creation_time"`
}

type PingOneLdapGatewayDataStoreData struct {
	BinaryAttributes     types.List   `tfsdk:"binary_attributes"`
	Id                   types.String `tfsdk:"id"`
	LdapType             types.String `tfsdk:"ldap_type"`
	MaskAttributeValues  types.Bool   `tfsdk:"mask_attribute_values"`
	Name                 types.String `tfsdk:"name"`
	PingOneConnectionRef types.String `tfsdk:"ping_one_connection_ref"`
	PingOneEnvironmentId types.String `tfsdk:"ping_one_environment_id"`
	PingOneLdapGatewayId types.String `tfsdk:"ping_one_ldap_gateway_id"`
	Type                 types.String `tfsdk:"type"`
	UseSsl               types.Bool   `tfsdk:"use_ssl"`
}

type PluginConfigDescriptorData struct {
	ActionDescriptors *[]*ActionDescriptorData `tfsdk:"action_descriptors"`
	Description       types.String             `tfsdk:"description"`
	Fields            *[]*FieldDescriptorData  `tfsdk:"fields"`
	Tables            *[]*TableDescriptorData  `tfsdk:"tables"`
}

type PluginConfigurationData struct {
	Fields *[]*ConfigFieldData `tfsdk:"fields"`
	Tables *[]*ConfigTableData `tfsdk:"tables"`
}

type PluginInstanceData struct {
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
}

type PolicyActionData struct {
	Context types.String `tfsdk:"context"`
	Type    types.String `tfsdk:"type"`
}

type ProcessorPolicyToGeneratorMappingData struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Id                               types.String                              `tfsdk:"id"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
}

type ProcessorPolicyToGeneratorMappingsData struct {
	Items *[]*ProcessorPolicyToGeneratorMappingData `tfsdk:"items"`
}

type ProfileConfigData struct {
	DeleteIdentityEnabled types.Bool   `tfsdk:"delete_identity_enabled"`
	TemplateName          types.String `tfsdk:"template_name"`
}

type ProtocolMessageCustomizationData struct {
	ContextName       types.String `tfsdk:"context_name"`
	MessageExpression types.String `tfsdk:"message_expression"`
}

type ProxySettingsData struct {
	Host types.String `tfsdk:"host"`
	Port types.Number `tfsdk:"port"`
}

type ReadGroupsData struct {
	AttributeContract    *IdpInboundProvisioningAttributeContractData `tfsdk:"attribute_contract"`
	AttributeFulfillment map[string]*AttributeFulfillmentValueData    `tfsdk:"attribute_fulfillment"`
	Attributes           *[]*AttributeData                            `tfsdk:"attributes"`
}

type ReadUsersData struct {
	AttributeContract    *IdpInboundProvisioningAttributeContractData `tfsdk:"attribute_contract"`
	AttributeFulfillment map[string]*AttributeFulfillmentValueData    `tfsdk:"attribute_fulfillment"`
	Attributes           *[]*AttributeData                            `tfsdk:"attributes"`
}

type RedirectValidationLocalSettingsData struct {
	EnableInErrorResourceValidation               types.Bool                                       `tfsdk:"enable_in_error_resource_validation"`
	EnableTargetResourceValidationForIdpDiscovery types.Bool                                       `tfsdk:"enable_target_resource_validation_for_idp_discovery"`
	EnableTargetResourceValidationForSLO          types.Bool                                       `tfsdk:"enable_target_resource_validation_for_slo"`
	EnableTargetResourceValidationForSSO          types.Bool                                       `tfsdk:"enable_target_resource_validation_for_sso"`
	WhiteList                                     *[]*RedirectValidationSettingsWhitelistEntryData `tfsdk:"white_list"`
}

type RedirectValidationPartnerSettingsData struct {
	EnableWreplyValidationSLO types.Bool `tfsdk:"enable_wreply_validation_slo"`
}

type RedirectValidationSettingsData struct {
	Id                                types.String                           `tfsdk:"id"`
	RedirectValidationLocalSettings   *RedirectValidationLocalSettingsData   `tfsdk:"redirect_validation_local_settings"`
	RedirectValidationPartnerSettings *RedirectValidationPartnerSettingsData `tfsdk:"redirect_validation_partner_settings"`
}

type RedirectValidationSettingsWhitelistEntryData struct {
	AllowQueryAndFragment types.Bool   `tfsdk:"allow_query_and_fragment"`
	IdpDiscovery          types.Bool   `tfsdk:"idp_discovery"`
	InErrorResource       types.Bool   `tfsdk:"in_error_resource"`
	RequireHttps          types.Bool   `tfsdk:"require_https"`
	TargetResourceSLO     types.Bool   `tfsdk:"target_resource_slo"`
	TargetResourceSSO     types.Bool   `tfsdk:"target_resource_sso"`
	ValidDomain           types.String `tfsdk:"valid_domain"`
	ValidPath             types.String `tfsdk:"valid_path"`
}

type RegistrationConfigData struct {
	CaptchaEnabled                      types.Bool   `tfsdk:"captcha_enabled"`
	CreateAuthnSessionAfterRegistration types.Bool   `tfsdk:"create_authn_session_after_registration"`
	ExecuteWorkflow                     types.String `tfsdk:"execute_workflow"`
	RegistrationWorkflow                types.String `tfsdk:"registration_workflow"`
	TemplateName                        types.String `tfsdk:"template_name"`
	ThisIsMyDeviceEnabled               types.Bool   `tfsdk:"this_is_my_device_enabled"`
	UsernameField                       types.String `tfsdk:"username_field"`
}

type RequestPoliciesData struct {
	Items *[]*RequestPolicyData `tfsdk:"items"`
}

type RequestPolicyData struct {
	AllowUnsignedLoginHintToken      types.Bool                              `tfsdk:"allow_unsigned_login_hint_token"`
	AlternativeLoginHintTokenIssuers *[]*AlternativeLoginHintTokenIssuerData `tfsdk:"alternative_login_hint_token_issuers"`
	AuthenticatorRef                 types.String                            `tfsdk:"authenticator_ref"`
	Id                               types.String                            `tfsdk:"id"`
	IdentityHintContract             *IdentityHintContractData               `tfsdk:"identity_hint_contract"`
	IdentityHintContractFulfillment  *AttributeMappingData                   `tfsdk:"identity_hint_contract_fulfillment"`
	IdentityHintMapping              *AttributeMappingData                   `tfsdk:"identity_hint_mapping"`
	Name                             types.String                            `tfsdk:"name"`
	RequireTokenForIdentityHint      types.Bool                              `tfsdk:"require_token_for_identity_hint"`
	TransactionLifetime              types.Number                            `tfsdk:"transaction_lifetime"`
	UserCodePcvRef                   types.String                            `tfsdk:"user_code_pcv_ref"`
}

type ResourceCategoryInfoData struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResourceLinkData struct {
	Id       types.String `tfsdk:"id"`
	Location types.String `tfsdk:"location"`
}

type ResourceOwnerCredentialsMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Id                           types.String                              `tfsdk:"id"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	PasswordValidatorRef         types.String                              `tfsdk:"password_validator_ref"`
}

type ResourceOwnerCredentialsMappingsData struct {
	Items *[]*ResourceOwnerCredentialsMappingData `tfsdk:"items"`
}

type ResourceUsageData struct {
	CategoryId types.String `tfsdk:"category_id"`
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Ref        types.String `tfsdk:"ref"`
	Type       types.String `tfsdk:"type"`
}

type ResourceUsagesData struct {
	Categories *[]*ResourceCategoryInfoData `tfsdk:"categories"`
	Items      *[]*ResourceUsageData        `tfsdk:"items"`
}

type RestartPolicyActionData struct {
	Context types.String `tfsdk:"context"`
	Type    types.String `tfsdk:"type"`
}

type RolesAndProtocolsData struct {
	EnableIdpDiscovery types.Bool     `tfsdk:"enable_idp_discovery"`
	IdpRole            *IdpRoleData   `tfsdk:"idp_role"`
	OauthRole          *OAuthRoleData `tfsdk:"oauth_role"`
	SpRole             *SpRoleData    `tfsdk:"sp_role"`
}

type SAML20ProfileData struct {
	Enable            types.Bool `tfsdk:"enable"`
	EnableAutoConnect types.Bool `tfsdk:"enable_auto_connect"`
}

type SaasAttributeMappingData struct {
	FieldName     types.String                `tfsdk:"field_name"`
	SaasFieldInfo *SaasFieldConfigurationData `tfsdk:"saas_field_info"`
}

type SaasFieldConfigurationData struct {
	AttributeNames types.List   `tfsdk:"attribute_names"`
	CharacterCase  types.String `tfsdk:"character_case"`
	CreateOnly     types.Bool   `tfsdk:"create_only"`
	DefaultValue   types.String `tfsdk:"default_value"`
	Expression     types.String `tfsdk:"expression"`
	Masked         types.Bool   `tfsdk:"masked"`
	Parser         types.String `tfsdk:"parser"`
	Trim           types.Bool   `tfsdk:"trim"`
}

type SaasPluginFieldInfoDescriptorData struct {
	AttributeGroup       types.Bool                    `tfsdk:"attribute_group"`
	Code                 types.String                  `tfsdk:"code"`
	DefaultValue         types.String                  `tfsdk:"default_value"`
	DsLdapMap            types.Bool                    `tfsdk:"ds_ldap_map"`
	Label                types.String                  `tfsdk:"label"`
	MaxLength            types.Number                  `tfsdk:"max_length"`
	MinLength            types.Number                  `tfsdk:"min_length"`
	MultiValue           types.Bool                    `tfsdk:"multi_value"`
	Notes                types.List                    `tfsdk:"notes"`
	Options              *[]*SaasPluginFieldOptionData `tfsdk:"options"`
	Pattern              types.String                  `tfsdk:"pattern"`
	PersistForMembership types.Bool                    `tfsdk:"persist_for_membership"`
	Required             types.Bool                    `tfsdk:"required"`
	Unique               types.Bool                    `tfsdk:"unique"`
}

type SaasPluginFieldOptionData struct {
	Code  types.String `tfsdk:"code"`
	Label types.String `tfsdk:"label"`
}

type SchemaData struct {
	Attributes *[]*SchemaAttributeData `tfsdk:"attributes"`
	Namespace  types.String            `tfsdk:"namespace"`
}

type SchemaAttributeData struct {
	MultiValued   types.Bool   `tfsdk:"multi_valued"`
	Name          types.String `tfsdk:"name"`
	SubAttributes types.List   `tfsdk:"sub_attributes"`
	Types         types.List   `tfsdk:"types"`
}

type ScopeEntryData struct {
	Description types.String `tfsdk:"description"`
	Dynamic     types.Bool   `tfsdk:"dynamic"`
	Name        types.String `tfsdk:"name"`
}

type ScopeGroupEntryData struct {
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	Scopes      types.List   `tfsdk:"scopes"`
}

type SecondarySecretData struct {
	EncryptedSecret types.String `tfsdk:"encrypted_secret"`
	ExpiryTime      types.String `tfsdk:"expiry_time"`
	Secret          types.String `tfsdk:"secret"`
}

type SecretManagerData struct {
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
}

type SecretManagersData struct {
	Items *[]*SecretManagerData `tfsdk:"items"`
}

type ServerSettingsData struct {
	CaptchaSettings   *CaptchaSettingsData      `tfsdk:"captcha_settings"`
	ContactInfo       *ContactInfoData          `tfsdk:"contact_info"`
	EmailServer       *EmailServerSettingsData  `tfsdk:"email_server"`
	FederationInfo    *FederationInfoData       `tfsdk:"federation_info"`
	Notifications     *NotificationSettingsData `tfsdk:"notifications"`
	RolesAndProtocols *RolesAndProtocolsData    `tfsdk:"roles_and_protocols"`
}

type ServiceAssociationData struct {
	ComponentName types.String `tfsdk:"component_name"`
	Configured    types.Bool   `tfsdk:"configured"`
	ServiceNames  types.List   `tfsdk:"service_names"`
}

type ServiceAssociationsData struct {
	Items *[]*ServiceAssociationData `tfsdk:"items"`
}

type ServiceAuthenticationData struct {
	AttributeQuery       *ServiceModelData `tfsdk:"attribute_query"`
	ConnectionManagement *ServiceModelData `tfsdk:"connection_management"`
	Jmx                  *ServiceModelData `tfsdk:"jmx"`
	SsoDirectoryService  *ServiceModelData `tfsdk:"sso_directory_service"`
}

type ServiceModelData struct {
	EncryptedSharedSecret types.String `tfsdk:"encrypted_shared_secret"`
	Id                    types.String `tfsdk:"id"`
	SharedSecret          types.String `tfsdk:"shared_secret"`
}

type SessionSettingsData struct {
	Id                            types.String `tfsdk:"id"`
	RevokeUserSessionOnLogout     types.Bool   `tfsdk:"revoke_user_session_on_logout"`
	SessionRevocationLifetime     types.Number `tfsdk:"session_revocation_lifetime"`
	TrackAdapterSessionsForLogout types.Bool   `tfsdk:"track_adapter_sessions_for_logout"`
}

type SessionValidationSettingsData struct {
	CheckSessionRevocationStatus types.Bool `tfsdk:"check_session_revocation_status"`
	CheckValidAuthnSession       types.Bool `tfsdk:"check_valid_authn_session"`
	IncludeSessionId             types.Bool `tfsdk:"include_session_id"`
	Inherited                    types.Bool `tfsdk:"inherited"`
	UpdateAuthnSessionActivity   types.Bool `tfsdk:"update_authn_session_activity"`
}

type SigningKeysData struct {
	P256ActiveCertRef       types.String `tfsdk:"p256active_cert_ref"`
	P256PreviousCertRef     types.String `tfsdk:"p256previous_cert_ref"`
	P256PublishX5cParameter types.Bool   `tfsdk:"p256publish_x5c_parameter"`
	P384ActiveCertRef       types.String `tfsdk:"p384active_cert_ref"`
	P384PreviousCertRef     types.String `tfsdk:"p384previous_cert_ref"`
	P384PublishX5cParameter types.Bool   `tfsdk:"p384publish_x5c_parameter"`
	P521ActiveCertRef       types.String `tfsdk:"p521active_cert_ref"`
	P521PreviousCertRef     types.String `tfsdk:"p521previous_cert_ref"`
	P521PublishX5cParameter types.Bool   `tfsdk:"p521publish_x5c_parameter"`
	RsaActiveCertRef        types.String `tfsdk:"rsa_active_cert_ref"`
	RsaPreviousCertRef      types.String `tfsdk:"rsa_previous_cert_ref"`
	RsaPublishX5cParameter  types.Bool   `tfsdk:"rsa_publish_x5c_parameter"`
}

type SigningSettingsData struct {
	Algorithm                     types.String         `tfsdk:"algorithm"`
	AlternativeSigningKeyPairRefs *[]*ResourceLinkData `tfsdk:"alternative_signing_key_pair_refs"`
	IncludeCertInSignature        types.Bool           `tfsdk:"include_cert_in_signature"`
	IncludeRawKeyInSignature      types.Bool           `tfsdk:"include_raw_key_in_signature"`
	SigningKeyPairRef             types.String         `tfsdk:"signing_key_pair_ref"`
}

type SloServiceEndpointData struct {
	Binding     types.String `tfsdk:"binding"`
	ResponseUrl types.String `tfsdk:"response_url"`
	Url         types.String `tfsdk:"url"`
}

type SourceTypeIdKeyData struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type SpAdapterData struct {
	AttributeContract     *SpAdapterAttributeContractData     `tfsdk:"attribute_contract"`
	Configuration         *PluginConfigurationData            `tfsdk:"configuration"`
	Id                    types.String                        `tfsdk:"id"`
	Name                  types.String                        `tfsdk:"name"`
	ParentRef             types.String                        `tfsdk:"parent_ref"`
	PluginDescriptorRef   types.String                        `tfsdk:"plugin_descriptor_ref"`
	TargetApplicationInfo *SpAdapterTargetApplicationInfoData `tfsdk:"target_application_info"`
}

type SpAdapterAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type SpAdapterAttributeContractData struct {
	CoreAttributes     *[]*SpAdapterAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*SpAdapterAttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool                 `tfsdk:"inherited"`
}

type SpAdapterMappingData struct {
	AdapterOverrideSettings      *SpAdapterData                            `tfsdk:"adapter_override_settings"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictVirtualEntityIds     types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds   types.List                                `tfsdk:"restricted_virtual_entity_ids"`
	SpAdapterRef                 types.String                              `tfsdk:"sp_adapter_ref"`
}

type SpAdapterTargetApplicationInfoData struct {
	ApplicationIconUrl types.String `tfsdk:"application_icon_url"`
	ApplicationName    types.String `tfsdk:"application_name"`
	Inherited          types.Bool   `tfsdk:"inherited"`
}

type SpAdapterUrlMappingData struct {
	AdapterRef types.String `tfsdk:"adapter_ref"`
	Url        types.String `tfsdk:"url"`
}

type SpAdapterUrlMappingsData struct {
	Items *[]*SpAdapterUrlMappingData `tfsdk:"items"`
}

type SpAdaptersData struct {
	Items *[]*SpAdapterData `tfsdk:"items"`
}

type SpAttributeQueryData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	Attributes                   types.List                                `tfsdk:"attributes"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	Policy                       *SpAttributeQueryPolicyData               `tfsdk:"policy"`
}

type SpAttributeQueryPolicyData struct {
	EncryptAssertion            types.Bool `tfsdk:"encrypt_assertion"`
	RequireEncryptedNameId      types.Bool `tfsdk:"require_encrypted_name_id"`
	RequireSignedAttributeQuery types.Bool `tfsdk:"require_signed_attribute_query"`
	SignAssertion               types.Bool `tfsdk:"sign_assertion"`
	SignResponse                types.Bool `tfsdk:"sign_response"`
}

type SpBrowserSsoData struct {
	AdapterMappings                               *[]*IdpAdapterAssertionMappingData                   `tfsdk:"adapter_mappings"`
	AlwaysSignArtifactResponse                    types.Bool                                           `tfsdk:"always_sign_artifact_response"`
	Artifact                                      *ArtifactSettingsData                                `tfsdk:"artifact"`
	AssertionLifetime                             *AssertionLifetimeData                               `tfsdk:"assertion_lifetime"`
	AttributeContract                             *SpBrowserSsoAttributeContractData                   `tfsdk:"attribute_contract"`
	AuthenticationPolicyContractAssertionMappings *[]*AuthenticationPolicyContractAssertionMappingData `tfsdk:"authentication_policy_contract_assertion_mappings"`
	DefaultTargetUrl                              types.String                                         `tfsdk:"default_target_url"`
	EnabledProfiles                               types.List                                           `tfsdk:"enabled_profiles"`
	EncryptionPolicy                              *EncryptionPolicyData                                `tfsdk:"encryption_policy"`
	IncomingBindings                              types.List                                           `tfsdk:"incoming_bindings"`
	MessageCustomizations                         *[]*ProtocolMessageCustomizationData                 `tfsdk:"message_customizations"`
	Protocol                                      types.String                                         `tfsdk:"protocol"`
	RequireSignedAuthnRequests                    types.Bool                                           `tfsdk:"require_signed_authn_requests"`
	SignAssertions                                types.Bool                                           `tfsdk:"sign_assertions"`
	SignResponseAsRequired                        types.Bool                                           `tfsdk:"sign_response_as_required"`
	SloServiceEndpoints                           *[]*SloServiceEndpointData                           `tfsdk:"slo_service_endpoints"`
	SpSamlIdentityMapping                         types.String                                         `tfsdk:"sp_saml_identity_mapping"`
	SpWsFedIdentityMapping                        types.String                                         `tfsdk:"sp_ws_fed_identity_mapping"`
	SsoServiceEndpoints                           *[]*SpSsoServiceEndpointData                         `tfsdk:"sso_service_endpoints"`
	UrlWhitelistEntries                           *[]*UrlWhitelistEntryData                            `tfsdk:"url_whitelist_entries"`
	WsFedTokenType                                types.String                                         `tfsdk:"ws_fed_token_type"`
	WsTrustVersion                                types.String                                         `tfsdk:"ws_trust_version"`
}

type SpBrowserSsoAttributeData struct {
	Name       types.String `tfsdk:"name"`
	NameFormat types.String `tfsdk:"name_format"`
}

type SpBrowserSsoAttributeContractData struct {
	CoreAttributes     *[]*SpBrowserSsoAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*SpBrowserSsoAttributeData `tfsdk:"extended_attributes"`
}

type SpConnectionData struct {
	Active                                 types.Bool                                  `tfsdk:"active"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	ApplicationIconUrl                     types.String                                `tfsdk:"application_icon_url"`
	ApplicationName                        types.String                                `tfsdk:"application_name"`
	AttributeQuery                         *SpAttributeQueryData                       `tfsdk:"attribute_query"`
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	ConnectionTargetType                   types.String                                `tfsdk:"connection_target_type"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	Id                                     types.String                                `tfsdk:"id"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	Name                                   types.String                                `tfsdk:"name"`
	OutboundProvision                      *OutboundProvisionData                      `tfsdk:"outbound_provision"`
	SpBrowserSso                           *SpBrowserSsoData                           `tfsdk:"sp_browser_sso"`
	Type                                   types.String                                `tfsdk:"type"`
	VirtualEntityIds                       types.List                                  `tfsdk:"virtual_entity_ids"`
	WsTrust                                *SpWsTrustData                              `tfsdk:"ws_trust"`
}

type SpConnectionsData struct {
	Items *[]*SpConnectionData `tfsdk:"items"`
}

type SpDefaultUrlsData struct {
	ConfirmSlo    types.Bool   `tfsdk:"confirm_slo"`
	SloSuccessUrl types.String `tfsdk:"slo_success_url"`
	SsoSuccessUrl types.String `tfsdk:"sso_success_url"`
}

type SpRoleData struct {
	Enable                    types.Bool           `tfsdk:"enable"`
	EnableInboundProvisioning types.Bool           `tfsdk:"enable_inbound_provisioning"`
	EnableOpenIDConnect       types.Bool           `tfsdk:"enable_open_id_connect"`
	EnableSaml10              types.Bool           `tfsdk:"enable_saml10"`
	EnableSaml11              types.Bool           `tfsdk:"enable_saml11"`
	EnableWsFed               types.Bool           `tfsdk:"enable_ws_fed"`
	EnableWsTrust             types.Bool           `tfsdk:"enable_ws_trust"`
	Saml20Profile             *SpSAML20ProfileData `tfsdk:"saml20profile"`
}

type SpSAML20ProfileData struct {
	Enable            types.Bool `tfsdk:"enable"`
	EnableAutoConnect types.Bool `tfsdk:"enable_auto_connect"`
	EnableXASP        types.Bool `tfsdk:"enable_xasp"`
}

type SpSsoServiceEndpointData struct {
	Binding   types.String `tfsdk:"binding"`
	Index     types.Number `tfsdk:"index"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	Url       types.String `tfsdk:"url"`
}

type SpTokenGeneratorMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	DefaultMapping               types.Bool                                `tfsdk:"default_mapping"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	RestrictedVirtualEntityIds   types.List                                `tfsdk:"restricted_virtual_entity_ids"`
	SpTokenGeneratorRef          types.String                              `tfsdk:"sp_token_generator_ref"`
}

type SpUrlMappingData struct {
	Ref  types.String `tfsdk:"ref"`
	Type types.String `tfsdk:"type"`
	Url  types.String `tfsdk:"url"`
}

type SpUrlMappingsData struct {
	Items *[]*SpUrlMappingData `tfsdk:"items"`
}

type SpWsTrustData struct {
	AbortIfNotFulfilledFromRequest types.Bool                           `tfsdk:"abort_if_not_fulfilled_from_request"`
	AttributeContract              *SpWsTrustAttributeContractData      `tfsdk:"attribute_contract"`
	DefaultTokenType               types.String                         `tfsdk:"default_token_type"`
	EncryptSaml2Assertion          types.Bool                           `tfsdk:"encrypt_saml2assertion"`
	GenerateKey                    types.Bool                           `tfsdk:"generate_key"`
	MessageCustomizations          *[]*ProtocolMessageCustomizationData `tfsdk:"message_customizations"`
	MinutesAfter                   types.Number                         `tfsdk:"minutes_after"`
	MinutesBefore                  types.Number                         `tfsdk:"minutes_before"`
	OAuthAssertionProfiles         types.Bool                           `tfsdk:"o_auth_assertion_profiles"`
	PartnerServiceIds              types.List                           `tfsdk:"partner_service_ids"`
	RequestContractRef             types.String                         `tfsdk:"request_contract_ref"`
	TokenProcessorMappings         *[]*IdpTokenProcessorMappingData     `tfsdk:"token_processor_mappings"`
}

type SpWsTrustAttributeData struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type SpWsTrustAttributeContractData struct {
	CoreAttributes     *[]*SpWsTrustAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*SpWsTrustAttributeData `tfsdk:"extended_attributes"`
}

type SqlMethodData struct {
	StoredProcedure *StoredProcedureData `tfsdk:"stored_procedure"`
	Table           *TableData           `tfsdk:"table"`
}

type SslServerSettingsData struct {
	ActiveAdminConsoleCerts  *[]*ResourceLinkData `tfsdk:"active_admin_console_certs"`
	ActiveRuntimeServerCerts *[]*ResourceLinkData `tfsdk:"active_runtime_server_certs"`
	AdminConsoleCertRef      types.String         `tfsdk:"admin_console_cert_ref"`
	RuntimeServerCertRef     types.String         `tfsdk:"runtime_server_cert_ref"`
}

type SsoOAuthMappingData struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type StoredProcedureData struct {
	Schema          types.String `tfsdk:"schema"`
	StoredProcedure types.String `tfsdk:"stored_procedure"`
}

type StsRequestParametersContractData struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Parameters types.List   `tfsdk:"parameters"`
}

type StsRequestParametersContractsData struct {
	Items *[]*StsRequestParametersContractData `tfsdk:"items"`
}

type SystemKeyData struct {
	CreationDate     types.String `tfsdk:"creation_date"`
	EncryptedKeyData types.String `tfsdk:"encrypted_key_data"`
	KeyData          types.String `tfsdk:"key_data"`
}

type SystemKeysData struct {
	Current  *SystemKeyData `tfsdk:"current"`
	Pending  *SystemKeyData `tfsdk:"pending"`
	Previous *SystemKeyData `tfsdk:"previous"`
}

type TableData struct {
	Schema         types.String `tfsdk:"schema"`
	TableName      types.String `tfsdk:"table_name"`
	UniqueIdColumn types.String `tfsdk:"unique_id_column"`
}

type TableDescriptorData struct {
	Columns           *[]*FieldDescriptorData `tfsdk:"columns"`
	Description       types.String            `tfsdk:"description"`
	Label             types.String            `tfsdk:"label"`
	Name              types.String            `tfsdk:"name"`
	RequireDefaultRow types.Bool              `tfsdk:"require_default_row"`
}

type TextLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	Type                  types.String          `tfsdk:"type"`
}

type TokenExchangeGeneratorGroupData struct {
	GeneratorMappings *[]*TokenExchangeGeneratorMappingData `tfsdk:"generator_mappings"`
	Id                types.String                          `tfsdk:"id"`
	Name              types.String                          `tfsdk:"name"`
	ResourceUris      types.List                            `tfsdk:"resource_uris"`
}

type TokenExchangeGeneratorGroupsData struct {
	Items *[]*TokenExchangeGeneratorGroupData `tfsdk:"items"`
}

type TokenExchangeGeneratorMappingData struct {
	DefaultMapping     types.Bool   `tfsdk:"default_mapping"`
	RequestedTokenType types.String `tfsdk:"requested_token_type"`
	TokenGenerator     types.String `tfsdk:"token_generator"`
}

type TokenExchangeGeneratorSettingsData struct {
	DefaultGeneratorGroupRef types.String `tfsdk:"default_generator_group_ref"`
}

type TokenExchangeProcessorAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type TokenExchangeProcessorAttributeContractData struct {
	CoreAttributes     *[]*TokenExchangeProcessorAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*TokenExchangeProcessorAttributeData `tfsdk:"extended_attributes"`
}

type TokenExchangeProcessorMappingData struct {
	ActorTokenProcessor          types.String                              `tfsdk:"actor_token_processor"`
	ActorTokenType               types.String                              `tfsdk:"actor_token_type"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SubjectTokenProcessor        types.String                              `tfsdk:"subject_token_processor"`
	SubjectTokenType             types.String                              `tfsdk:"subject_token_type"`
}

type TokenExchangeProcessorPoliciesData struct {
	Items *[]*TokenExchangeProcessorPolicyData `tfsdk:"items"`
}

type TokenExchangeProcessorPolicyData struct {
	ActorTokenRequired types.Bool                                   `tfsdk:"actor_token_required"`
	AttributeContract  *TokenExchangeProcessorAttributeContractData `tfsdk:"attribute_contract"`
	Id                 types.String                                 `tfsdk:"id"`
	Name               types.String                                 `tfsdk:"name"`
	ProcessorMappings  *[]*TokenExchangeProcessorMappingData        `tfsdk:"processor_mappings"`
}

type TokenExchangeProcessorSettingsData struct {
	DefaultProcessorPolicyRef types.String `tfsdk:"default_processor_policy_ref"`
}

type TokenGeneratorData struct {
	AttributeContract   *TokenGeneratorAttributeContractData `tfsdk:"attribute_contract"`
	Configuration       *PluginConfigurationData             `tfsdk:"configuration"`
	Id                  types.String                         `tfsdk:"id"`
	Name                types.String                         `tfsdk:"name"`
	ParentRef           types.String                         `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                         `tfsdk:"plugin_descriptor_ref"`
}

type TokenGeneratorAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type TokenGeneratorAttributeContractData struct {
	CoreAttributes     *[]*TokenGeneratorAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*TokenGeneratorAttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool                      `tfsdk:"inherited"`
}

type TokenGeneratorsData struct {
	Items *[]*TokenGeneratorData `tfsdk:"items"`
}

type TokenProcessorData struct {
	AttributeContract   *TokenProcessorAttributeContractData `tfsdk:"attribute_contract"`
	Configuration       *PluginConfigurationData             `tfsdk:"configuration"`
	Id                  types.String                         `tfsdk:"id"`
	Name                types.String                         `tfsdk:"name"`
	ParentRef           types.String                         `tfsdk:"parent_ref"`
	PluginDescriptorRef types.String                         `tfsdk:"plugin_descriptor_ref"`
}

type TokenProcessorAttributeData struct {
	Masked types.Bool   `tfsdk:"masked"`
	Name   types.String `tfsdk:"name"`
}

type TokenProcessorAttributeContractData struct {
	CoreAttributes     *[]*TokenProcessorAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*TokenProcessorAttributeData `tfsdk:"extended_attributes"`
	Inherited          types.Bool                      `tfsdk:"inherited"`
	MaskOgnlValues     types.Bool                      `tfsdk:"mask_ognl_values"`
}

type TokenProcessorsData struct {
	Items *[]*TokenProcessorData `tfsdk:"items"`
}

type TokenToTokenMappingData struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
	Id                               types.String                              `tfsdk:"id"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
}

type TokenToTokenMappingsData struct {
	Items *[]*TokenToTokenMappingData `tfsdk:"items"`
}

type UrlWhitelistEntryData struct {
	AllowQueryAndFragment types.Bool   `tfsdk:"allow_query_and_fragment"`
	RequireHttps          types.Bool   `tfsdk:"require_https"`
	ValidDomain           types.String `tfsdk:"valid_domain"`
	ValidPath             types.String `tfsdk:"valid_path"`
}

type UserCredentialsData struct {
	CurrentPassword types.String `tfsdk:"current_password"`
	NewPassword     types.String `tfsdk:"new_password"`
}

type UsernamePasswordCredentialsData struct {
	EncryptedPassword types.String `tfsdk:"encrypted_password"`
	Password          types.String `tfsdk:"password"`
	Username          types.String `tfsdk:"username"`
}

type UsersData struct {
	ReadUsers  *ReadUsersData  `tfsdk:"read_users"`
	WriteUsers *WriteUsersData `tfsdk:"write_users"`
}

type ValidationErrorData struct {
	DeveloperMessage types.String `tfsdk:"developer_message"`
	ErrorId          types.String `tfsdk:"error_id"`
	FieldPath        types.String `tfsdk:"field_path"`
	Message          types.String `tfsdk:"message"`
}

type VersionData struct {
	Version types.String `tfsdk:"version"`
}

type VirtualHostNameSettingsData struct {
	VirtualHostNames types.List `tfsdk:"virtual_host_names"`
}

type WriteGroupsData struct {
	AttributeFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_fulfillment"`
}

type WriteUsersData struct {
	AttributeFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_fulfillment"`
}

type WsTrustStsSettingsData struct {
	BasicAuthnEnabled      types.Bool                          `tfsdk:"basic_authn_enabled"`
	ClientCertAuthnEnabled types.Bool                          `tfsdk:"client_cert_authn_enabled"`
	IssuerCerts            *[]*ResourceLinkData                `tfsdk:"issuer_certs"`
	RestrictByIssuerCert   types.Bool                          `tfsdk:"restrict_by_issuer_cert"`
	RestrictBySubjectDn    types.Bool                          `tfsdk:"restrict_by_subject_dn"`
	SubjectDns             types.List                          `tfsdk:"subject_dns"`
	Users                  *[]*UsernamePasswordCredentialsData `tfsdk:"users"`
}

type X509FileData struct {
	CryptoProvider types.String `tfsdk:"crypto_provider"`
	FileData       types.String `tfsdk:"file_data"`
	Id             types.String `tfsdk:"id"`
}
