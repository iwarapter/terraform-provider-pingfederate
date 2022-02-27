package framework

import "github.com/hashicorp/terraform-plugin-framework/types"

type AccessTokenAttributeData struct {
	Name        types.String `tfsdk:"name"`
	MultiValued types.Bool   `tfsdk:"multi_valued"`
}

type AccessTokenAttributeContractData struct {
	CoreAttributes          *[]*AccessTokenAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes      *[]*AccessTokenAttributeData `tfsdk:"extended_attributes"`
	Inherited               types.Bool                   `tfsdk:"inherited"`
	DefaultSubjectAttribute types.String                 `tfsdk:"default_subject_attribute"`
}

type AccessTokenManagementSettingsData struct {
	DefaultAccessTokenManagerRef types.String `tfsdk:"default_access_token_manager_ref"`
}

type AccessTokenManagerData struct {
	Name                      types.String                      `tfsdk:"name"`
	PluginDescriptorRef       types.String                      `tfsdk:"plugin_descriptor_ref"`
	ParentRef                 types.String                      `tfsdk:"parent_ref"`
	Configuration             *PluginConfigurationData          `tfsdk:"configuration"`
	Id                        types.String                      `tfsdk:"id"`
	AttributeContract         *AccessTokenAttributeContractData `tfsdk:"attribute_contract"`
	SelectionSettings         *AtmSelectionSettingsData         `tfsdk:"selection_settings"`
	AccessControlSettings     *AtmAccessControlSettingsData     `tfsdk:"access_control_settings"`
	SessionValidationSettings *SessionValidationSettingsData    `tfsdk:"session_validation_settings"`
}

type AccessTokenManagerMappingData struct {
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	AccessTokenManagerRef        types.String                              `tfsdk:"access_token_manager_ref"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
}

type AccessTokenMappingData struct {
	Context                      *AccessTokenMappingContextData            `tfsdk:"context"`
	AccessTokenManagerRef        types.String                              `tfsdk:"access_token_manager_ref"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	Id                           types.String                              `tfsdk:"id"`
}

type AccessTokenMappingContextData struct {
	Type       types.String `tfsdk:"type"`
	ContextRef types.String `tfsdk:"context_ref"`
}

type AccountManagementSettingsData struct {
	FlagComparisonValue        types.String `tfsdk:"flag_comparison_value"`
	FlagComparisonStatus       types.Bool   `tfsdk:"flag_comparison_status"`
	DefaultStatus              types.Bool   `tfsdk:"default_status"`
	AccountStatusAttributeName types.String `tfsdk:"account_status_attribute_name"`
	AccountStatusAlgorithm     types.String `tfsdk:"account_status_algorithm"`
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
	AllowAdditionalEntities   types.Bool     `tfsdk:"allow_additional_entities"`
	AllowAllEntities          types.Bool     `tfsdk:"allow_all_entities"`
	AdditionalAllowedEntities *[]*EntityData `tfsdk:"additional_allowed_entities"`
}

type AdditionalKeySetData struct {
	Id          types.String         `tfsdk:"id"`
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	SigningKeys *SigningKeysData     `tfsdk:"signing_keys"`
	Issuers     *[]*ResourceLinkData `tfsdk:"issuers"`
}

type AdditionalKeySetsData struct {
	Items *[]*AdditionalKeySetData `tfsdk:"items"`
}

type AdministrativeAccountData struct {
	Description       types.String   `tfsdk:"description"`
	Department        types.String   `tfsdk:"department"`
	Auditor           types.Bool     `tfsdk:"auditor"`
	PhoneNumber       types.String   `tfsdk:"phone_number"`
	EmailAddress      types.String   `tfsdk:"email_address"`
	Roles             []types.String `tfsdk:"roles"`
	Username          types.String   `tfsdk:"username"`
	Password          types.String   `tfsdk:"password"`
	EncryptedPassword types.String   `tfsdk:"encrypted_password"`
	Active            types.Bool     `tfsdk:"active"`
}

type AdministrativeAccountsData struct {
	Items *[]*AdministrativeAccountData `tfsdk:"items"`
}

type AlternativeLoginHintTokenIssuerData struct {
	Issuer  types.String `tfsdk:"issuer"`
	JwksURL types.String `tfsdk:"jwks_url"`
	Jwks    types.String `tfsdk:"jwks"`
}

type ApcMappingPolicyActionData struct {
	AttributeMapping                *AttributeMappingData `tfsdk:"attribute_mapping"`
	Type                            types.String          `tfsdk:"type"`
	Context                         types.String          `tfsdk:"context"`
	AuthenticationPolicyContractRef types.String          `tfsdk:"authentication_policy_contract_ref"`
}

type ApcToPersistentGrantMappingData struct {
	Id                              types.String                              `tfsdk:"id"`
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type ApcToPersistentGrantMappingsData struct {
	Items *[]*ApcToPersistentGrantMappingData `tfsdk:"items"`
}

type ApcToSpAdapterMappingData struct {
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
	Id                               types.String                              `tfsdk:"id"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
}

type ApcToSpAdapterMappingsData struct {
	Items *[]*ApcToSpAdapterMappingData `tfsdk:"items"`
}

type ApiResultData struct {
	DeveloperMessage types.String            `tfsdk:"developer_message"`
	ValidationErrors *[]*ValidationErrorData `tfsdk:"validation_errors"`
	ResultId         types.String            `tfsdk:"result_id"`
	Message          types.String            `tfsdk:"message"`
}

type ApplicationSessionPolicyData struct {
	IdleTimeoutMins types.Number `tfsdk:"idle_timeout_mins"`
	MaxTimeoutMins  types.Number `tfsdk:"max_timeout_mins"`
}

type ArtifactResolverLocationData struct {
	Index types.Number `tfsdk:"index"`
	Url   types.String `tfsdk:"url"`
}

type ArtifactSettingsData struct {
	ResolverLocations *[]*ArtifactResolverLocationData `tfsdk:"resolver_locations"`
	SourceId          types.String                     `tfsdk:"source_id"`
	Lifetime          types.Number                     `tfsdk:"lifetime"`
}

type AssertionLifetimeData struct {
	MinutesBefore types.Number `tfsdk:"minutes_before"`
	MinutesAfter  types.Number `tfsdk:"minutes_after"`
}

type AtmAccessControlSettingsData struct {
	Inherited       types.Bool           `tfsdk:"inherited"`
	RestrictClients types.Bool           `tfsdk:"restrict_clients"`
	AllowedClients  *[]*ResourceLinkData `tfsdk:"allowed_clients"`
}

type AtmSelectionSettingsData struct {
	Inherited    types.Bool     `tfsdk:"inherited"`
	ResourceUris []types.String `tfsdk:"resource_uris"`
}

type AttributeFulfillmentValueData struct {
	Source *SourceTypeIdKeyData `tfsdk:"source"`
	Value  types.String         `tfsdk:"value"`
}

type AttributeMappingData struct {
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
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
	Type                         types.String                              `tfsdk:"type"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	Id                           types.String                              `tfsdk:"id"`
	Description                  types.String                              `tfsdk:"description"`
}

type AuthenticationPoliciesSettingsData struct {
	EnableSpAuthnSelection  types.Bool `tfsdk:"enable_sp_authn_selection"`
	EnableIdpAuthnSelection types.Bool `tfsdk:"enable_idp_authn_selection"`
}

type AuthenticationPolicyData struct {
	TrackedHttpParameters        []types.String                   `tfsdk:"tracked_http_parameters"`
	FailIfNoSelection            types.Bool                       `tfsdk:"fail_if_no_selection"`
	AuthnSelectionTrees          *[]*AuthenticationPolicyTreeData `tfsdk:"authn_selection_trees"`
	DefaultAuthenticationSources *[]*AuthenticationSourceData     `tfsdk:"default_authentication_sources"`
}

type AuthenticationPolicyContractData struct {
	Id                 types.String                                  `tfsdk:"id"`
	Name               types.String                                  `tfsdk:"name"`
	CoreAttributes     *[]*AuthenticationPolicyContractAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes []types.String                                `tfsdk:"extended_attributes"`
}

type AuthenticationPolicyContractAssertionMappingData struct {
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	RestrictVirtualEntityIds        types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds      []types.String                            `tfsdk:"restricted_virtual_entity_ids"`
	AbortSsoTransactionAsFailSafe   types.Bool                                `tfsdk:"abort_sso_transaction_as_fail_safe"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type AuthenticationPolicyContractAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type AuthenticationPolicyContractMappingData struct {
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	AuthenticationPolicyContractRef types.String                              `tfsdk:"authentication_policy_contract_ref"`
	RestrictVirtualServerIds        types.Bool                                `tfsdk:"restrict_virtual_server_ids"`
	RestrictedVirtualServerIds      []types.String                            `tfsdk:"restricted_virtual_server_ids"`
	JdbcAttributeSources            []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources            []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources          []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
}

type AuthenticationPolicyContractsData struct {
	Items *[]*AuthenticationPolicyContractData `tfsdk:"items"`
}

type AuthenticationPolicyFragmentData struct {
	Id          types.String                      `tfsdk:"id"`
	Name        types.String                      `tfsdk:"name"`
	Description types.String                      `tfsdk:"description"`
	RootNode    *AuthenticationPolicyTreeNodeData `tfsdk:"root_node"`
	Inputs      types.String                      `tfsdk:"inputs"`
	Outputs     types.String                      `tfsdk:"outputs"`
}

type AuthenticationPolicyFragmentsData struct {
	Items *[]*AuthenticationPolicyFragmentData `tfsdk:"items"`
}

type AuthenticationPolicyTreeData struct {
	Description                     types.String                      `tfsdk:"description"`
	AuthenticationApiApplicationRef types.String                      `tfsdk:"authentication_api_application_ref"`
	Enabled                         types.Bool                        `tfsdk:"enabled"`
	RootNode                        *AuthenticationPolicyTreeNodeData `tfsdk:"root_node"`
	Id                              types.String                      `tfsdk:"id"`
	Name                            types.String                      `tfsdk:"name"`
}

type AuthenticationPolicyTreeNodeData struct {
	Action   *PolicyActionData                    `tfsdk:"action"`
	Children *[]*AuthenticationPolicyTreeNodeData `tfsdk:"children"`
}

type AuthenticationSelectorData struct {
	Id                  types.String                                 `tfsdk:"id"`
	Name                types.String                                 `tfsdk:"name"`
	PluginDescriptorRef types.String                                 `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String                                 `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData                     `tfsdk:"configuration"`
	AttributeContract   *AuthenticationSelectorAttributeContractData `tfsdk:"attribute_contract"`
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
	AuthnContextSensitive types.Bool                `tfsdk:"authn_context_sensitive"`
	Id                    types.String              `tfsdk:"id"`
	AuthenticationSource  *AuthenticationSourceData `tfsdk:"authentication_source"`
	EnableSessions        types.Bool                `tfsdk:"enable_sessions"`
	Persistent            types.Bool                `tfsdk:"persistent"`
	IdleTimeoutMins       types.Number              `tfsdk:"idle_timeout_mins"`
	MaxTimeoutMins        types.Number              `tfsdk:"max_timeout_mins"`
	TimeoutDisplayUnit    types.String              `tfsdk:"timeout_display_unit"`
}

type AuthenticationSourceData struct {
	Type      types.String `tfsdk:"type"`
	SourceRef types.String `tfsdk:"source_ref"`
}

type AuthnApiApplicationData struct {
	Id                           types.String   `tfsdk:"id"`
	Name                         types.String   `tfsdk:"name"`
	Url                          types.String   `tfsdk:"url"`
	Description                  types.String   `tfsdk:"description"`
	AdditionalAllowedOrigins     []types.String `tfsdk:"additional_allowed_origins"`
	ClientForRedirectlessModeRef types.String   `tfsdk:"client_for_redirectless_mode_ref"`
}

type AuthnApiApplicationsData struct {
	Items *[]*AuthnApiApplicationData `tfsdk:"items"`
}

type AuthnApiSettingsData struct {
	RestrictAccessToRedirectlessMode types.Bool   `tfsdk:"restrict_access_to_redirectless_mode"`
	IncludeRequestContext            types.Bool   `tfsdk:"include_request_context"`
	ApiEnabled                       types.Bool   `tfsdk:"api_enabled"`
	DefaultApplicationRef            types.String `tfsdk:"default_application_ref"`
	EnableApiDescriptions            types.Bool   `tfsdk:"enable_api_descriptions"`
}

type AuthnContextMappingData struct {
	Local  types.String `tfsdk:"local"`
	Remote types.String `tfsdk:"remote"`
}

type AuthnSelectorPolicyActionData struct {
	Type                      types.String `tfsdk:"type"`
	Context                   types.String `tfsdk:"context"`
	AuthenticationSelectorRef types.String `tfsdk:"authentication_selector_ref"`
}

type AuthnSourcePolicyActionData struct {
	Type                 types.String                   `tfsdk:"type"`
	Context              types.String                   `tfsdk:"context"`
	AttributeRules       *AttributeRulesData            `tfsdk:"attribute_rules"`
	AuthenticationSource *AuthenticationSourceData      `tfsdk:"authentication_source"`
	InputUserIdMapping   *AttributeFulfillmentValueData `tfsdk:"input_user_id_mapping"`
	UserIdAuthenticated  types.Bool                     `tfsdk:"user_id_authenticated"`
}

type AuthorizationServerSettingsData struct {
	PersistentGrantIdleTimeoutTimeUnit     types.String                 `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	PersistentGrantContract                *PersistentGrantContractData `tfsdk:"persistent_grant_contract"`
	PendingAuthorizationTimeout            types.Number                 `tfsdk:"pending_authorization_timeout"`
	PersistentGrantIdleTimeout             types.Number                 `tfsdk:"persistent_grant_idle_timeout"`
	AuthorizationCodeTimeout               types.Number                 `tfsdk:"authorization_code_timeout"`
	AuthorizationCodeEntropy               types.Number                 `tfsdk:"authorization_code_entropy"`
	IncludeIssuerInAuthorizationResponse   types.Bool                   `tfsdk:"include_issuer_in_authorization_response"`
	RefreshRollingInterval                 types.Number                 `tfsdk:"refresh_rolling_interval"`
	ParReferenceLength                     types.Number                 `tfsdk:"par_reference_length"`
	ExclusiveScopeGroups                   *[]*ScopeGroupEntryData      `tfsdk:"exclusive_scope_groups"`
	AdminWebServicePcvRef                  types.String                 `tfsdk:"admin_web_service_pcv_ref"`
	UserAuthorizationUrl                   types.String                 `tfsdk:"user_authorization_url"`
	BypassAuthorizationForApprovedGrants   types.Bool                   `tfsdk:"bypass_authorization_for_approved_grants"`
	TokenEndpointBaseUrl                   types.String                 `tfsdk:"token_endpoint_base_url"`
	PersistentGrantLifetime                types.Number                 `tfsdk:"persistent_grant_lifetime"`
	PersistentGrantReuseGrantTypes         []types.String               `tfsdk:"persistent_grant_reuse_grant_types"`
	AtmIdForOAuthGrantManagement           types.String                 `tfsdk:"atm_id_for_o_auth_grant_management"`
	ScopeGroups                            *[]*ScopeGroupEntryData      `tfsdk:"scope_groups"`
	Scopes                                 *[]*ScopeEntryData           `tfsdk:"scopes"`
	ExclusiveScopes                        *[]*ScopeEntryData           `tfsdk:"exclusive_scopes"`
	TrackUserSessionsForLogout             types.Bool                   `tfsdk:"track_user_sessions_for_logout"`
	ScopeForOAuthGrantManagement           types.String                 `tfsdk:"scope_for_o_auth_grant_management"`
	RegisteredAuthorizationPath            types.String                 `tfsdk:"registered_authorization_path"`
	DevicePollingInterval                  types.Number                 `tfsdk:"device_polling_interval"`
	ParReferenceTimeout                    types.Number                 `tfsdk:"par_reference_timeout"`
	DefaultScopeDescription                types.String                 `tfsdk:"default_scope_description"`
	AllowUnidentifiedClientExtensionGrants types.Bool                   `tfsdk:"allow_unidentified_client_extension_grants"`
	BypassActivationCodeConfirmation       types.Bool                   `tfsdk:"bypass_activation_code_confirmation"`
	RefreshTokenLength                     types.Number                 `tfsdk:"refresh_token_length"`
	AllowUnidentifiedClientROCreds         types.Bool                   `tfsdk:"allow_unidentified_client_ro_creds"`
	UserAuthorizationConsentAdapter        types.String                 `tfsdk:"user_authorization_consent_adapter"`
	ApprovedScopesAttribute                types.String                 `tfsdk:"approved_scopes_attribute"`
	ParStatus                              types.String                 `tfsdk:"par_status"`
	DisallowPlainPKCE                      types.Bool                   `tfsdk:"disallow_plain_pkce"`
	RollRefreshTokenValues                 types.Bool                   `tfsdk:"roll_refresh_token_values"`
	AllowedOrigins                         []types.String               `tfsdk:"allowed_origins"`
	UserAuthorizationConsentPageSetting    types.String                 `tfsdk:"user_authorization_consent_page_setting"`
	PersistentGrantLifetimeUnit            types.String                 `tfsdk:"persistent_grant_lifetime_unit"`
}

type BackChannelAuthData struct {
	Type                 types.String                     `tfsdk:"type"`
	HttpBasicCredentials *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	DigitalSignature     types.Bool                       `tfsdk:"digital_signature"`
}

type BaseDefaultValueLocalIdentityFieldData struct {
	Type                  types.String          `tfsdk:"type"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
}

type BaseProviderRoleData struct {
	Enable        types.Bool `tfsdk:"enable"`
	EnableSaml11  types.Bool `tfsdk:"enable_saml11"`
	EnableSaml10  types.Bool `tfsdk:"enable_saml10"`
	EnableWsFed   types.Bool `tfsdk:"enable_ws_fed"`
	EnableWsTrust types.Bool `tfsdk:"enable_ws_trust"`
}

type BaseSelectionLocalIdentityFieldData struct {
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	Options               []types.String        `tfsdk:"options"`
}

type BinaryLdapAttributeSettingsData struct {
	BinaryEncoding types.String `tfsdk:"binary_encoding"`
}

type CSRResponseData struct {
	FileData types.String `tfsdk:"file_data"`
}

type CaptchaSettingsData struct {
	SiteKey            types.String `tfsdk:"site_key"`
	SecretKey          types.String `tfsdk:"secret_key"`
	EncryptedSecretKey types.String `tfsdk:"encrypted_secret_key"`
}

type CertViewData struct {
	KeySize                 types.Number   `tfsdk:"key_size"`
	Sha256Fingerprint       types.String   `tfsdk:"sha256fingerprint"`
	SubjectAlternativeNames []types.String `tfsdk:"subject_alternative_names"`
	Sha1Fingerprint         types.String   `tfsdk:"sha1fingerprint"`
	CryptoProvider          types.String   `tfsdk:"crypto_provider"`
	SerialNumber            types.String   `tfsdk:"serial_number"`
	SubjectDN               types.String   `tfsdk:"subject_dn"`
	Expires                 types.String   `tfsdk:"expires"`
	SignatureAlgorithm      types.String   `tfsdk:"signature_algorithm"`
	Id                      types.String   `tfsdk:"id"`
	IssuerDN                types.String   `tfsdk:"issuer_dn"`
	ValidFrom               types.String   `tfsdk:"valid_from"`
	KeyAlgorithm            types.String   `tfsdk:"key_algorithm"`
	Version                 types.Number   `tfsdk:"version"`
	Status                  types.String   `tfsdk:"status"`
}

type CertViewsData struct {
	Items *[]*CertViewData `tfsdk:"items"`
}

type CertificateExpirationNotificationSettingsData struct {
	EmailAddress             types.String `tfsdk:"email_address"`
	InitialWarningPeriod     types.Number `tfsdk:"initial_warning_period"`
	FinalWarningPeriod       types.Number `tfsdk:"final_warning_period"`
	NotificationPublisherRef types.String `tfsdk:"notification_publisher_ref"`
}

type CertificateRevocationSettingsData struct {
	OcspSettings  *OcspSettingsData  `tfsdk:"ocsp_settings"`
	CrlSettings   *CrlSettingsData   `tfsdk:"crl_settings"`
	ProxySettings *ProxySettingsData `tfsdk:"proxy_settings"`
}

type ChangeDetectionSettingsData struct {
	TimeStampAttributeName types.String `tfsdk:"time_stamp_attribute_name"`
	UserObjectClass        types.String `tfsdk:"user_object_class"`
	GroupObjectClass       types.String `tfsdk:"group_object_class"`
	ChangedUsersAlgorithm  types.String `tfsdk:"changed_users_algorithm"`
	UsnAttributeName       types.String `tfsdk:"usn_attribute_name"`
}

type ChannelData struct {
	Timeout          types.Number                 `tfsdk:"timeout"`
	Active           types.Bool                   `tfsdk:"active"`
	ChannelSource    *ChannelSourceData           `tfsdk:"channel_source"`
	AttributeMapping *[]*SaasAttributeMappingData `tfsdk:"attribute_mapping"`
	Name             types.String                 `tfsdk:"name"`
	MaxThreads       types.Number                 `tfsdk:"max_threads"`
}

type ChannelSourceData struct {
	ChangeDetectionSettings   *ChangeDetectionSettingsData   `tfsdk:"change_detection_settings"`
	GroupMembershipDetection  *GroupMembershipDetectionData  `tfsdk:"group_membership_detection"`
	AccountManagementSettings *AccountManagementSettingsData `tfsdk:"account_management_settings"`
	BaseDn                    types.String                   `tfsdk:"base_dn"`
	GuidAttributeName         types.String                   `tfsdk:"guid_attribute_name"`
	GuidBinary                types.Bool                     `tfsdk:"guid_binary"`
	GroupSourceLocation       *ChannelSourceLocationData     `tfsdk:"group_source_location"`
	DataSource                types.String                   `tfsdk:"data_source"`
	UserSourceLocation        *ChannelSourceLocationData     `tfsdk:"user_source_location"`
}

type ChannelSourceLocationData struct {
	GroupDN      types.String `tfsdk:"group_dn"`
	Filter       types.String `tfsdk:"filter"`
	NestedSearch types.Bool   `tfsdk:"nested_search"`
}

type CheckboxGroupLocalIdentityFieldData struct {
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	Options               []types.String        `tfsdk:"options"`
}

type CheckboxLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
}

type CibaServerPolicySettingsData struct {
	DefaultRequestPolicyRef types.String `tfsdk:"default_request_policy_ref"`
}

type ClientData struct {
	BypassActivationCodeConfirmationOverride types.Bool                      `tfsdk:"bypass_activation_code_confirmation_override"`
	CibaPollingInterval                      types.Number                    `tfsdk:"ciba_polling_interval"`
	RequestPolicyRef                         types.String                    `tfsdk:"request_policy_ref"`
	PersistentGrantExpirationType            types.String                    `tfsdk:"persistent_grant_expiration_type"`
	RestrictScopes                           types.Bool                      `tfsdk:"restrict_scopes"`
	RestrictedScopes                         []types.String                  `tfsdk:"restricted_scopes"`
	RestrictedResponseTypes                  []types.String                  `tfsdk:"restricted_response_types"`
	DevicePollingIntervalOverride            types.Number                    `tfsdk:"device_polling_interval_override"`
	RequireProofKeyForCodeExchange           types.Bool                      `tfsdk:"require_proof_key_for_code_exchange"`
	CibaRequestObjectSigningAlgorithm        types.String                    `tfsdk:"ciba_request_object_signing_algorithm"`
	CibaUserCodeSupported                    types.Bool                      `tfsdk:"ciba_user_code_supported"`
	PersistentGrantIdleTimeoutType           types.String                    `tfsdk:"persistent_grant_idle_timeout_type"`
	PersistentGrantReuseType                 types.String                    `tfsdk:"persistent_grant_reuse_type"`
	RequireSignedRequests                    types.Bool                      `tfsdk:"require_signed_requests"`
	OidcPolicy                               *ClientOIDCPolicyData           `tfsdk:"oidc_policy"`
	PersistentGrantIdleTimeoutTimeUnit       types.String                    `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	RequestObjectSigningAlgorithm            types.String                    `tfsdk:"request_object_signing_algorithm"`
	ClientAuth                               *ClientAuthData                 `tfsdk:"client_auth"`
	Enabled                                  types.Bool                      `tfsdk:"enabled"`
	Description                              types.String                    `tfsdk:"description"`
	ValidateUsingAllEligibleAtms             types.Bool                      `tfsdk:"validate_using_all_eligible_atms"`
	PersistentGrantIdleTimeout               types.Number                    `tfsdk:"persistent_grant_idle_timeout"`
	CibaDeliveryMode                         types.String                    `tfsdk:"ciba_delivery_mode"`
	RedirectUris                             []types.String                  `tfsdk:"redirect_uris"`
	RestrictToDefaultAccessTokenManager      types.Bool                      `tfsdk:"restrict_to_default_access_token_manager"`
	RequirePushedAuthorizationRequests       types.Bool                      `tfsdk:"require_pushed_authorization_requests"`
	PendingAuthorizationTimeoutOverride      types.Number                    `tfsdk:"pending_authorization_timeout_override"`
	Name                                     types.String                    `tfsdk:"name"`
	RefreshRolling                           types.String                    `tfsdk:"refresh_rolling"`
	RefreshTokenRollingIntervalType          types.String                    `tfsdk:"refresh_token_rolling_interval_type"`
	PersistentGrantReuseGrantTypes           []types.String                  `tfsdk:"persistent_grant_reuse_grant_types"`
	CibaNotificationEndpoint                 types.String                    `tfsdk:"ciba_notification_endpoint"`
	GrantTypes                               []types.String                  `tfsdk:"grant_types"`
	RefreshTokenRollingInterval              types.Number                    `tfsdk:"refresh_token_rolling_interval"`
	ExclusiveScopes                          []types.String                  `tfsdk:"exclusive_scopes"`
	JwksSettings                             *JwksSettingsData               `tfsdk:"jwks_settings"`
	TokenExchangeProcessorPolicyRef          types.String                    `tfsdk:"token_exchange_processor_policy_ref"`
	PersistentGrantExpirationTimeUnit        types.String                    `tfsdk:"persistent_grant_expiration_time_unit"`
	BypassApprovalPage                       types.Bool                      `tfsdk:"bypass_approval_page"`
	DeviceFlowSettingType                    types.String                    `tfsdk:"device_flow_setting_type"`
	UserAuthorizationUrlOverride             types.String                    `tfsdk:"user_authorization_url_override"`
	AllowAuthenticationApiInit               types.Bool                      `tfsdk:"allow_authentication_api_init"`
	ExtendedParameters                       map[string]*ParameterValuesData `tfsdk:"extended_parameters"`
	CibaRequireSignedRequests                types.Bool                      `tfsdk:"ciba_require_signed_requests"`
	ClientId                                 types.String                    `tfsdk:"client_id"`
	LogoUrl                                  types.String                    `tfsdk:"logo_url"`
	DefaultAccessTokenManagerRef             types.String                    `tfsdk:"default_access_token_manager_ref"`
	PersistentGrantExpirationTime            types.Number                    `tfsdk:"persistent_grant_expiration_time"`
}

type ClientAuthData struct {
	EnforceReplayPrevention           types.Bool   `tfsdk:"enforce_replay_prevention"`
	TokenEndpointAuthSigningAlgorithm types.String `tfsdk:"token_endpoint_auth_signing_algorithm"`
	Type                              types.String `tfsdk:"type"`
	Secret                            types.String `tfsdk:"secret"`
	EncryptedSecret                   types.String `tfsdk:"encrypted_secret"`
	ClientCertIssuerDn                types.String `tfsdk:"client_cert_issuer_dn"`
	ClientCertSubjectDn               types.String `tfsdk:"client_cert_subject_dn"`
}

type ClientMetadataData struct {
	Parameter   types.String `tfsdk:"parameter"`
	Description types.String `tfsdk:"description"`
	MultiValued types.Bool   `tfsdk:"multi_valued"`
}

type ClientOIDCPolicyData struct {
	IdTokenEncryptionAlgorithm             types.String   `tfsdk:"id_token_encryption_algorithm"`
	IdTokenContentEncryptionAlgorithm      types.String   `tfsdk:"id_token_content_encryption_algorithm"`
	PolicyGroup                            types.String   `tfsdk:"policy_group"`
	SectorIdentifierUri                    types.String   `tfsdk:"sector_identifier_uri"`
	IdTokenSigningAlgorithm                types.String   `tfsdk:"id_token_signing_algorithm"`
	GrantAccessSessionRevocationApi        types.Bool     `tfsdk:"grant_access_session_revocation_api"`
	GrantAccessSessionSessionManagementApi types.Bool     `tfsdk:"grant_access_session_session_management_api"`
	PingAccessLogoutCapable                types.Bool     `tfsdk:"ping_access_logout_capable"`
	LogoutUris                             []types.String `tfsdk:"logout_uris"`
	PairwiseIdentifierUserType             types.Bool     `tfsdk:"pairwise_identifier_user_type"`
}

type ClientRegistrationOIDCPolicyData struct {
	IdTokenSigningAlgorithm           types.String `tfsdk:"id_token_signing_algorithm"`
	IdTokenEncryptionAlgorithm        types.String `tfsdk:"id_token_encryption_algorithm"`
	IdTokenContentEncryptionAlgorithm types.String `tfsdk:"id_token_content_encryption_algorithm"`
	PolicyGroup                       types.String `tfsdk:"policy_group"`
}

type ClientRegistrationPoliciesData struct {
	Items *[]*ClientRegistrationPolicyData `tfsdk:"items"`
}

type ClientRegistrationPolicyData struct {
	Name                types.String             `tfsdk:"name"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
}

type ClientSecretData struct {
	Secret          types.String `tfsdk:"secret"`
	EncryptedSecret types.String `tfsdk:"encrypted_secret"`
}

type ClientSettingsData struct {
	ClientMetadata            *[]*ClientMetadataData         `tfsdk:"client_metadata"`
	DynamicClientRegistration *DynamicClientRegistrationData `tfsdk:"dynamic_client_registration"`
}

type ClientsData struct {
	Items *[]*ClientData `tfsdk:"items"`
}

type ClusterNodeData struct {
	Address   types.String `tfsdk:"address"`
	Index     types.Number `tfsdk:"index"`
	Mode      types.String `tfsdk:"mode"`
	NodeGroup types.String `tfsdk:"node_group"`
	Version   types.String `tfsdk:"version"`
	NodeTags  types.String `tfsdk:"node_tags"`
}

type ClusterStatusData struct {
	Nodes                *[]*ClusterNodeData `tfsdk:"nodes"`
	LastConfigUpdateTime types.String        `tfsdk:"last_config_update_time"`
	LastReplicationTime  types.String        `tfsdk:"last_replication_time"`
	ReplicationRequired  types.Bool          `tfsdk:"replication_required"`
	MixedMode            types.Bool          `tfsdk:"mixed_mode"`
}

type ConditionalIssuanceCriteriaEntryData struct {
	Source        *SourceTypeIdKeyData `tfsdk:"source"`
	AttributeName types.String         `tfsdk:"attribute_name"`
	Condition     types.String         `tfsdk:"condition"`
	Value         types.String         `tfsdk:"value"`
	ErrorResult   types.String         `tfsdk:"error_result"`
}

type ConfigFieldData struct {
	Value          types.String `tfsdk:"value"`
	EncryptedValue types.String `tfsdk:"encrypted_value"`
	Inherited      types.Bool   `tfsdk:"inherited"`
	Name           types.String `tfsdk:"name"`
}

type ConfigRowData struct {
	Fields     *[]*ConfigFieldData `tfsdk:"fields"`
	DefaultRow types.Bool          `tfsdk:"default_row"`
}

type ConfigStoreBundleData struct {
	Items *[]*ConfigStoreSettingData `tfsdk:"items"`
}

type ConfigStoreSettingData struct {
	ListValue   []types.String          `tfsdk:"list_value"`
	MapValue    map[string]types.String `tfsdk:"map_value"`
	Type        types.String            `tfsdk:"type"`
	Id          types.String            `tfsdk:"id"`
	StringValue types.String            `tfsdk:"string_value"`
}

type ConfigTableData struct {
	Inherited types.Bool        `tfsdk:"inherited"`
	Name      types.String      `tfsdk:"name"`
	Rows      *[]*ConfigRowData `tfsdk:"rows"`
}

type ConfigurationEncryptionKeyData struct {
	KeyId        types.String `tfsdk:"key_id"`
	CreationDate types.String `tfsdk:"creation_date"`
}

type ConfigurationEncryptionKeysData struct {
	Items *[]*ConfigurationEncryptionKeyData `tfsdk:"items"`
}

type ConnectionData struct {
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	Type                                   types.String                                `tfsdk:"type"`
	Id                                     types.String                                `tfsdk:"id"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	Active                                 types.Bool                                  `tfsdk:"active"`
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	Name                                   types.String                                `tfsdk:"name"`
	VirtualEntityIds                       []types.String                              `tfsdk:"virtual_entity_ids"`
}

type ConnectionCertData struct {
	CertView                  *CertViewData `tfsdk:"cert_view"`
	X509File                  *X509FileData `tfsdk:"x509file"`
	ActiveVerificationCert    types.Bool    `tfsdk:"active_verification_cert"`
	PrimaryVerificationCert   types.Bool    `tfsdk:"primary_verification_cert"`
	SecondaryVerificationCert types.Bool    `tfsdk:"secondary_verification_cert"`
	EncryptionCert            types.Bool    `tfsdk:"encryption_cert"`
}

type ConnectionCertsData struct {
	Items *[]*ConnectionCertData `tfsdk:"items"`
}

type ConnectionCredentialsData struct {
	VerificationIssuerDN          types.String                 `tfsdk:"verification_issuer_dn"`
	SigningSettings               *SigningSettingsData         `tfsdk:"signing_settings"`
	OutboundBackChannelAuth       *OutboundBackChannelAuthData `tfsdk:"outbound_back_channel_auth"`
	VerificationSubjectDN         types.String                 `tfsdk:"verification_subject_dn"`
	Certs                         *[]*ConnectionCertData       `tfsdk:"certs"`
	BlockEncryptionAlgorithm      types.String                 `tfsdk:"block_encryption_algorithm"`
	KeyTransportAlgorithm         types.String                 `tfsdk:"key_transport_algorithm"`
	DecryptionKeyPairRef          types.String                 `tfsdk:"decryption_key_pair_ref"`
	SecondaryDecryptionKeyPairRef types.String                 `tfsdk:"secondary_decryption_key_pair_ref"`
	InboundBackChannelAuth        *InboundBackChannelAuthData  `tfsdk:"inbound_back_channel_auth"`
}

type ConnectionGroupLicenseViewData struct {
	Name            types.String `tfsdk:"name"`
	ConnectionCount types.Number `tfsdk:"connection_count"`
	StartDate       types.String `tfsdk:"start_date"`
	EndDate         types.String `tfsdk:"end_date"`
}

type ConnectionMetadataUrlData struct {
	MetadataUrlRef           types.String `tfsdk:"metadata_url_ref"`
	EnableAutoMetadataUpdate types.Bool   `tfsdk:"enable_auto_metadata_update"`
}

type ContactInfoData struct {
	Email     types.String `tfsdk:"email"`
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
	Phone     types.String `tfsdk:"phone"`
	Company   types.String `tfsdk:"company"`
}

type ContinuePolicyActionData struct {
	Type    types.String `tfsdk:"type"`
	Context types.String `tfsdk:"context"`
}

type ConvertMetadataRequestData struct {
	ConnectionType          types.String    `tfsdk:"connection_type"`
	ExpectedProtocol        types.String    `tfsdk:"expected_protocol"`
	ExpectedEntityId        types.String    `tfsdk:"expected_entity_id"`
	SamlMetadata            types.String    `tfsdk:"saml_metadata"`
	VerificationCertificate types.String    `tfsdk:"verification_certificate"`
	TemplateConnection      *ConnectionData `tfsdk:"template_connection"`
}

type ConvertMetadataResponseData struct {
	SignatureStatus  types.String    `tfsdk:"signature_status"`
	CertTrustStatus  types.String    `tfsdk:"cert_trust_status"`
	CertSubjectDn    types.String    `tfsdk:"cert_subject_dn"`
	CertSerialNumber types.String    `tfsdk:"cert_serial_number"`
	CertExpiration   types.String    `tfsdk:"cert_expiration"`
	Connection       *ConnectionData `tfsdk:"connection"`
}

type CrlSettingsData struct {
	TreatNonRetrievableCrlAsRevoked   types.Bool   `tfsdk:"treat_non_retrievable_crl_as_revoked"`
	VerifyCrlSignature                types.Bool   `tfsdk:"verify_crl_signature"`
	NextRetryMinsWhenResolveFailed    types.Number `tfsdk:"next_retry_mins_when_resolve_failed"`
	NextRetryMinsWhenNextUpdateInPast types.Number `tfsdk:"next_retry_mins_when_next_update_in_past"`
}

type CustomAttributeSourceData struct {
	Description                  types.String                              `tfsdk:"description"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	FilterFields                 *[]*FieldEntryData                        `tfsdk:"filter_fields"`
	Id                           types.String                              `tfsdk:"id"`
}

type CustomDataStoreData struct {
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Name                types.String             `tfsdk:"name"`
	Type                types.String             `tfsdk:"type"`
	Id                  types.String             `tfsdk:"id"`
	MaskAttributeValues types.Bool               `tfsdk:"mask_attribute_values"`
}

type DataStoreData struct {
	Type                types.String `tfsdk:"type"`
	Id                  types.String `tfsdk:"id"`
	MaskAttributeValues types.Bool   `tfsdk:"mask_attribute_values"`
}

type DataStoreAttributeData struct {
	Type     types.String            `tfsdk:"type"`
	Name     types.String            `tfsdk:"name"`
	Metadata map[string]types.String `tfsdk:"metadata"`
}

type DataStoreConfigData struct {
	DataStoreRef     types.String                       `tfsdk:"data_store_ref"`
	DataStoreMapping map[string]*DataStoreAttributeData `tfsdk:"data_store_mapping"`
	Type             types.String                       `tfsdk:"type"`
}

type DataStoresData struct {
	Items *[]*DataStoreData `tfsdk:"items"`
}

type DateLocalIdentityFieldData struct {
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	DefaultValue          types.String          `tfsdk:"default_value"`
}

type DecryptionKeysData struct {
	PrimaryKeyRef       types.String `tfsdk:"primary_key_ref"`
	SecondaryKeyPairRef types.String `tfsdk:"secondary_key_pair_ref"`
}

type DecryptionPolicyData struct {
	AttributesEncrypted       types.Bool `tfsdk:"attributes_encrypted"`
	SubjectNameIdEncrypted    types.Bool `tfsdk:"subject_name_id_encrypted"`
	SloEncryptSubjectNameID   types.Bool `tfsdk:"slo_encrypt_subject_name_id"`
	SloSubjectNameIDEncrypted types.Bool `tfsdk:"slo_subject_name_id_encrypted"`
	AssertionEncrypted        types.Bool `tfsdk:"assertion_encrypted"`
}

type DonePolicyActionData struct {
	Type    types.String `tfsdk:"type"`
	Context types.String `tfsdk:"context"`
}

type DropDownLocalIdentityFieldData struct {
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	Options               []types.String        `tfsdk:"options"`
	Id                    types.String          `tfsdk:"id"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Label                 types.String          `tfsdk:"label"`
}

type DynamicClientRegistrationData struct {
	RestrictToDefaultAccessTokenManager      types.Bool                        `tfsdk:"restrict_to_default_access_token_manager"`
	RefreshTokenRollingIntervalType          types.String                      `tfsdk:"refresh_token_rolling_interval_type"`
	DeviceFlowSettingType                    types.String                      `tfsdk:"device_flow_setting_type"`
	UserAuthorizationUrlOverride             types.String                      `tfsdk:"user_authorization_url_override"`
	RestrictCommonScopes                     types.Bool                        `tfsdk:"restrict_common_scopes"`
	RequireSignedRequests                    types.Bool                        `tfsdk:"require_signed_requests"`
	PersistentGrantExpirationTimeUnit        types.String                      `tfsdk:"persistent_grant_expiration_time_unit"`
	PersistentGrantIdleTimeoutType           types.String                      `tfsdk:"persistent_grant_idle_timeout_type"`
	PersistentGrantIdleTimeoutTimeUnit       types.String                      `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	OidcPolicy                               *ClientRegistrationOIDCPolicyData `tfsdk:"oidc_policy"`
	PolicyRefs                               *[]*ResourceLinkData              `tfsdk:"policy_refs"`
	AllowClientDelete                        types.Bool                        `tfsdk:"allow_client_delete"`
	InitialAccessTokenScope                  types.String                      `tfsdk:"initial_access_token_scope"`
	PersistentGrantExpirationType            types.String                      `tfsdk:"persistent_grant_expiration_type"`
	RotateRegistrationAccessToken            types.Bool                        `tfsdk:"rotate_registration_access_token"`
	TokenExchangeProcessorPolicyRef          types.String                      `tfsdk:"token_exchange_processor_policy_ref"`
	RotateClientSecret                       types.Bool                        `tfsdk:"rotate_client_secret"`
	DisableRegistrationAccessTokens          types.Bool                        `tfsdk:"disable_registration_access_tokens"`
	PendingAuthorizationTimeoutOverride      types.Number                      `tfsdk:"pending_authorization_timeout_override"`
	DefaultAccessTokenManagerRef             types.String                      `tfsdk:"default_access_token_manager_ref"`
	ClientCertIssuerRef                      types.String                      `tfsdk:"client_cert_issuer_ref"`
	CibaPollingInterval                      types.Number                      `tfsdk:"ciba_polling_interval"`
	EnforceReplayPrevention                  types.Bool                        `tfsdk:"enforce_replay_prevention"`
	ClientCertIssuerType                     types.String                      `tfsdk:"client_cert_issuer_type"`
	PersistentGrantIdleTimeout               types.Number                      `tfsdk:"persistent_grant_idle_timeout"`
	RefreshRolling                           types.String                      `tfsdk:"refresh_rolling"`
	CibaRequireSignedRequests                types.Bool                        `tfsdk:"ciba_require_signed_requests"`
	AllowedExclusiveScopes                   []types.String                    `tfsdk:"allowed_exclusive_scopes"`
	PersistentGrantExpirationTime            types.Number                      `tfsdk:"persistent_grant_expiration_time"`
	DevicePollingIntervalOverride            types.Number                      `tfsdk:"device_polling_interval_override"`
	BypassActivationCodeConfirmationOverride types.Bool                        `tfsdk:"bypass_activation_code_confirmation_override"`
	RequireProofKeyForCodeExchange           types.Bool                        `tfsdk:"require_proof_key_for_code_exchange"`
	RequestPolicyRef                         types.String                      `tfsdk:"request_policy_ref"`
	RestrictedCommonScopes                   []types.String                    `tfsdk:"restricted_common_scopes"`
	RefreshTokenRollingInterval              types.Number                      `tfsdk:"refresh_token_rolling_interval"`
}

type EmailLocalIdentityFieldData struct {
	Type                  types.String          `tfsdk:"type"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
}

type EmailServerSettingsData struct {
	SourceAddr               types.String `tfsdk:"source_addr"`
	VerifyHostname           types.Bool   `tfsdk:"verify_hostname"`
	EnableUtf8MessageHeaders types.Bool   `tfsdk:"enable_utf8message_headers"`
	EncryptedPassword        types.String `tfsdk:"encrypted_password"`
	EmailServer              types.String `tfsdk:"email_server"`
	Timeout                  types.Number `tfsdk:"timeout"`
	UseTLS                   types.Bool   `tfsdk:"use_tls"`
	RetryAttempts            types.Number `tfsdk:"retry_attempts"`
	UseDebugging             types.Bool   `tfsdk:"use_debugging"`
	Password                 types.String `tfsdk:"password"`
	UseSSL                   types.Bool   `tfsdk:"use_ssl"`
	Username                 types.String `tfsdk:"username"`
	Port                     types.Number `tfsdk:"port"`
	SslPort                  types.Number `tfsdk:"ssl_port"`
	RetryDelay               types.Number `tfsdk:"retry_delay"`
}

type EmailVerificationConfigData struct {
	EmailVerificationEnabled             types.Bool   `tfsdk:"email_verification_enabled"`
	FieldStoringVerificationStatus       types.String `tfsdk:"field_storing_verification_status"`
	EmailVerificationErrorTemplateName   types.String `tfsdk:"email_verification_error_template_name"`
	EmailVerificationOtpTemplateName     types.String `tfsdk:"email_verification_otp_template_name"`
	OtlTimeToLive                        types.Number `tfsdk:"otl_time_to_live"`
	FieldForEmailToVerify                types.String `tfsdk:"field_for_email_to_verify"`
	AllowedOtpCharacterSet               types.String `tfsdk:"allowed_otp_character_set"`
	NotificationPublisherRef             types.String `tfsdk:"notification_publisher_ref"`
	RequireVerifiedEmailTemplateName     types.String `tfsdk:"require_verified_email_template_name"`
	VerifyEmailTemplateName              types.String `tfsdk:"verify_email_template_name"`
	EmailVerificationSentTemplateName    types.String `tfsdk:"email_verification_sent_template_name"`
	EmailVerificationType                types.String `tfsdk:"email_verification_type"`
	OtpRetryAttempts                     types.Number `tfsdk:"otp_retry_attempts"`
	EmailVerificationSuccessTemplateName types.String `tfsdk:"email_verification_success_template_name"`
	OtpLength                            types.Number `tfsdk:"otp_length"`
	OtpTimeToLive                        types.Number `tfsdk:"otp_time_to_live"`
	RequireVerifiedEmail                 types.Bool   `tfsdk:"require_verified_email"`
}

type EncryptionPolicyData struct {
	EncryptAssertion          types.Bool     `tfsdk:"encrypt_assertion"`
	EncryptedAttributes       []types.String `tfsdk:"encrypted_attributes"`
	EncryptSloSubjectNameId   types.Bool     `tfsdk:"encrypt_slo_subject_name_id"`
	SloSubjectNameIDEncrypted types.Bool     `tfsdk:"slo_subject_name_id_encrypted"`
}

type EntityData struct {
	EntityId          types.String `tfsdk:"entity_id"`
	EntityDescription types.String `tfsdk:"entity_description"`
}

type ExportMetadataRequestData struct {
	ConnectionId            types.String         `tfsdk:"connection_id"`
	VirtualServerId         types.String         `tfsdk:"virtual_server_id"`
	SigningSettings         *SigningSettingsData `tfsdk:"signing_settings"`
	UseSecondaryPortForSoap types.Bool           `tfsdk:"use_secondary_port_for_soap"`
	VirtualHostName         types.String         `tfsdk:"virtual_host_name"`
	ConnectionType          types.String         `tfsdk:"connection_type"`
}

type ExpressionIssuanceCriteriaEntryData struct {
	Expression  types.String `tfsdk:"expression"`
	ErrorResult types.String `tfsdk:"error_result"`
}

type ExtendedPropertiesData struct {
	Items *[]*ExtendedPropertyData `tfsdk:"items"`
}

type ExtendedPropertyData struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	MultiValued types.Bool   `tfsdk:"multi_valued"`
}

type FederationInfoData struct {
	Saml1xSourceId      types.String `tfsdk:"saml1x_source_id"`
	WsfedRealm          types.String `tfsdk:"wsfed_realm"`
	BaseUrl             types.String `tfsdk:"base_url"`
	Saml2EntityId       types.String `tfsdk:"saml2entity_id"`
	AutoConnectEntityId types.String `tfsdk:"auto_connect_entity_id"`
	Saml1xIssuerId      types.String `tfsdk:"saml1x_issuer_id"`
}

type FieldConfigData struct {
	Fields                    *[]*LocalIdentityFieldData `tfsdk:"fields"`
	StripSpaceFromUniqueField types.Bool                 `tfsdk:"strip_space_from_unique_field"`
}

type FieldEntryData struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}

type FragmentPolicyActionData struct {
	Type            types.String          `tfsdk:"type"`
	Context         types.String          `tfsdk:"context"`
	AttributeRules  *AttributeRulesData   `tfsdk:"attribute_rules"`
	Fragment        types.String          `tfsdk:"fragment"`
	FragmentMapping *AttributeMappingData `tfsdk:"fragment_mapping"`
}

type GlobalAuthenticationSessionPolicyData struct {
	IdleTimeoutMins            types.Number `tfsdk:"idle_timeout_mins"`
	IdleTimeoutDisplayUnit     types.String `tfsdk:"idle_timeout_display_unit"`
	MaxTimeoutMins             types.Number `tfsdk:"max_timeout_mins"`
	MaxTimeoutDisplayUnit      types.String `tfsdk:"max_timeout_display_unit"`
	EnableSessions             types.Bool   `tfsdk:"enable_sessions"`
	PersistentSessions         types.Bool   `tfsdk:"persistent_sessions"`
	HashUniqueUserKeyAttribute types.Bool   `tfsdk:"hash_unique_user_key_attribute"`
}

type GroupMembershipDetectionData struct {
	MemberOfGroupAttributeName types.String `tfsdk:"member_of_group_attribute_name"`
	GroupMemberAttributeName   types.String `tfsdk:"group_member_attribute_name"`
}

type HiddenLocalIdentityFieldData struct {
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	Id                    types.String          `tfsdk:"id"`
}

type IdentityHintAttributeData struct {
	Name types.String `tfsdk:"name"`
}

type IdentityHintContractData struct {
	CoreAttributes     *[]*IdentityHintAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdentityHintAttributeData `tfsdk:"extended_attributes"`
}

type IdpAdapterData struct {
	Id                  types.String                     `tfsdk:"id"`
	Name                types.String                     `tfsdk:"name"`
	PluginDescriptorRef types.String                     `tfsdk:"plugin_descriptor_ref"`
	AuthnCtxClassRef    types.String                     `tfsdk:"authn_ctx_class_ref"`
	AttributeMapping    *IdpAdapterContractMappingData   `tfsdk:"attribute_mapping"`
	AttributeContract   *IdpAdapterAttributeContractData `tfsdk:"attribute_contract"`
	ParentRef           types.String                     `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData         `tfsdk:"configuration"`
}

type IdpAdapterAssertionMappingData struct {
	AttributeContractFulfillment  map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria              *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	IdpAdapterRef                 types.String                              `tfsdk:"idp_adapter_ref"`
	RestrictVirtualEntityIds      types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds    []types.String                            `tfsdk:"restricted_virtual_entity_ids"`
	AdapterOverrideSettings       *IdpAdapterData                           `tfsdk:"adapter_override_settings"`
	AbortSsoTransactionAsFailSafe types.Bool                                `tfsdk:"abort_sso_transaction_as_fail_safe"`
	JdbcAttributeSources          []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources          []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources        []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
}

type IdpAdapterAttributeData struct {
	Name      types.String `tfsdk:"name"`
	Pseudonym types.Bool   `tfsdk:"pseudonym"`
	Masked    types.Bool   `tfsdk:"masked"`
}

type IdpAdapterAttributeContractData struct {
	MaskOgnlValues         types.Bool                  `tfsdk:"mask_ognl_values"`
	Inherited              types.Bool                  `tfsdk:"inherited"`
	CoreAttributes         *[]*IdpAdapterAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes     *[]*IdpAdapterAttributeData `tfsdk:"extended_attributes"`
	UniqueUserKeyAttribute types.String                `tfsdk:"unique_user_key_attribute"`
}

type IdpAdapterContractMappingData struct {
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	Inherited                    types.Bool                                `tfsdk:"inherited"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
}

type IdpAdapterMappingData struct {
	Id                           types.String                              `tfsdk:"id"`
	IdpAdapterRef                types.String                              `tfsdk:"idp_adapter_ref"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type IdpAdapterMappingsData struct {
	Items *[]*IdpAdapterMappingData `tfsdk:"items"`
}

type IdpAdaptersData struct {
	Items *[]*IdpAdapterData `tfsdk:"items"`
}

type IdpAttributeQueryData struct {
	Url          types.String                      `tfsdk:"url"`
	NameMappings *[]*AttributeQueryNameMappingData `tfsdk:"name_mappings"`
	Policy       *IdpAttributeQueryPolicyData      `tfsdk:"policy"`
}

type IdpAttributeQueryPolicyData struct {
	RequireSignedResponse     types.Bool `tfsdk:"require_signed_response"`
	RequireSignedAssertion    types.Bool `tfsdk:"require_signed_assertion"`
	RequireEncryptedAssertion types.Bool `tfsdk:"require_encrypted_assertion"`
	SignAttributeQuery        types.Bool `tfsdk:"sign_attribute_query"`
	EncryptNameId             types.Bool `tfsdk:"encrypt_name_id"`
	MaskAttributeValues       types.Bool `tfsdk:"mask_attribute_values"`
}

type IdpBrowserSsoData struct {
	Protocol                             types.String                                `tfsdk:"protocol"`
	AttributeContract                    *IdpBrowserSsoAttributeContractData         `tfsdk:"attribute_contract"`
	AdapterMappings                      *[]*SpAdapterMappingData                    `tfsdk:"adapter_mappings"`
	SignAuthnRequests                    types.Bool                                  `tfsdk:"sign_authn_requests"`
	EnabledProfiles                      []types.String                              `tfsdk:"enabled_profiles"`
	IncomingBindings                     []types.String                              `tfsdk:"incoming_bindings"`
	UrlWhitelistEntries                  *[]*UrlWhitelistEntryData                   `tfsdk:"url_whitelist_entries"`
	Artifact                             *ArtifactSettingsData                       `tfsdk:"artifact"`
	SloServiceEndpoints                  *[]*SloServiceEndpointData                  `tfsdk:"slo_service_endpoints"`
	AlwaysSignArtifactResponse           types.Bool                                  `tfsdk:"always_sign_artifact_response"`
	AuthnContextMappings                 *[]*AuthnContextMappingData                 `tfsdk:"authn_context_mappings"`
	IdpIdentityMapping                   types.String                                `tfsdk:"idp_identity_mapping"`
	MessageCustomizations                *[]*ProtocolMessageCustomizationData        `tfsdk:"message_customizations"`
	SsoServiceEndpoints                  *[]*IdpSsoServiceEndpointData               `tfsdk:"sso_service_endpoints"`
	AssertionsSigned                     types.Bool                                  `tfsdk:"assertions_signed"`
	OidcProviderSettings                 *OIDCProviderSettingsData                   `tfsdk:"oidc_provider_settings"`
	DefaultTargetUrl                     types.String                                `tfsdk:"default_target_url"`
	DecryptionPolicy                     *DecryptionPolicyData                       `tfsdk:"decryption_policy"`
	AuthenticationPolicyContractMappings *[]*AuthenticationPolicyContractMappingData `tfsdk:"authentication_policy_contract_mappings"`
	SsoOAuthMapping                      *SsoOAuthMappingData                        `tfsdk:"sso_o_auth_mapping"`
	OauthAuthenticationPolicyContractRef types.String                                `tfsdk:"oauth_authentication_policy_contract_ref"`
}

type IdpBrowserSsoAttributeData struct {
	Name   types.String `tfsdk:"name"`
	Masked types.Bool   `tfsdk:"masked"`
}

type IdpBrowserSsoAttributeContractData struct {
	CoreAttributes     *[]*IdpBrowserSsoAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*IdpBrowserSsoAttributeData `tfsdk:"extended_attributes"`
}

type IdpConnectionData struct {
	Active                                 types.Bool                                  `tfsdk:"active"`
	VirtualEntityIds                       []types.String                              `tfsdk:"virtual_entity_ids"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	OidcClientCredentials                  *OIDCClientCredentialsData                  `tfsdk:"oidc_client_credentials"`
	Type                                   types.String                                `tfsdk:"type"`
	ErrorPageMsgId                         types.String                                `tfsdk:"error_page_msg_id"`
	IdpOAuthGrantAttributeMapping          *IdpOAuthGrantAttributeMappingData          `tfsdk:"idp_o_auth_grant_attribute_mapping"`
	WsTrust                                *IdpWsTrustData                             `tfsdk:"ws_trust"`
	IdpBrowserSso                          *IdpBrowserSsoData                          `tfsdk:"idp_browser_sso"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	Name                                   types.String                                `tfsdk:"name"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	Id                                     types.String                                `tfsdk:"id"`
	AttributeQuery                         *IdpAttributeQueryData                      `tfsdk:"attribute_query"`
}

type IdpConnectionsData struct {
	Items *[]*IdpConnectionData `tfsdk:"items"`
}

type IdpDefaultUrlData struct {
	ConfirmIdpSlo    types.Bool   `tfsdk:"confirm_idp_slo"`
	IdpSloSuccessUrl types.String `tfsdk:"idp_slo_success_url"`
	IdpErrorMsg      types.String `tfsdk:"idp_error_msg"`
}

type IdpOAuthAttributeContractData struct {
	ExtendedAttributes *[]*IdpBrowserSsoAttributeData `tfsdk:"extended_attributes"`
	CoreAttributes     *[]*IdpBrowserSsoAttributeData `tfsdk:"core_attributes"`
}

type IdpOAuthGrantAttributeMappingData struct {
	AccessTokenManagerMappings *[]*AccessTokenManagerMappingData `tfsdk:"access_token_manager_mappings"`
	IdpOAuthAttributeContract  *IdpOAuthAttributeContractData    `tfsdk:"idp_o_auth_attribute_contract"`
}

type IdpRoleData struct {
	Enable                     types.Bool         `tfsdk:"enable"`
	Saml20Profile              *SAML20ProfileData `tfsdk:"saml20profile"`
	EnableOutboundProvisioning types.Bool         `tfsdk:"enable_outbound_provisioning"`
	EnableSaml11               types.Bool         `tfsdk:"enable_saml11"`
	EnableSaml10               types.Bool         `tfsdk:"enable_saml10"`
	EnableWsFed                types.Bool         `tfsdk:"enable_ws_fed"`
	EnableWsTrust              types.Bool         `tfsdk:"enable_ws_trust"`
}

type IdpSsoServiceEndpointData struct {
	Url     types.String `tfsdk:"url"`
	Binding types.String `tfsdk:"binding"`
}

type IdpToSpAdapterMappingData struct {
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	ApplicationName                  types.String                              `tfsdk:"application_name"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
	Id                               types.String                              `tfsdk:"id"`
	ApplicationIconUrl               types.String                              `tfsdk:"application_icon_url"`
}

type IdpToSpAdapterMappingsData struct {
	Items *[]*IdpToSpAdapterMappingData `tfsdk:"items"`
}

type IdpTokenProcessorMappingData struct {
	IdpTokenProcessorRef         types.String                              `tfsdk:"idp_token_processor_ref"`
	RestrictedVirtualEntityIds   []types.String                            `tfsdk:"restricted_virtual_entity_ids"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
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
	VerificationSubjectDN types.String                     `tfsdk:"verification_subject_dn"`
	Type                  types.String                     `tfsdk:"type"`
	HttpBasicCredentials  *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	DigitalSignature      types.Bool                       `tfsdk:"digital_signature"`
	VerificationIssuerDN  types.String                     `tfsdk:"verification_issuer_dn"`
	Certs                 *[]*ConnectionCertData           `tfsdk:"certs"`
	RequireSsl            types.Bool                       `tfsdk:"require_ssl"`
}

type IncomingProxySettingsData struct {
	ForwardedIpAddressHeaderName  types.String `tfsdk:"forwarded_ip_address_header_name"`
	ForwardedIpAddressHeaderIndex types.String `tfsdk:"forwarded_ip_address_header_index"`
	ForwardedHostHeaderName       types.String `tfsdk:"forwarded_host_header_name"`
	ForwardedHostHeaderIndex      types.String `tfsdk:"forwarded_host_header_index"`
	ClientCertSSLHeaderName       types.String `tfsdk:"client_cert_ssl_header_name"`
	ClientCertChainSSLHeaderName  types.String `tfsdk:"client_cert_chain_ssl_header_name"`
	ProxyTerminatesHttpsConns     types.Bool   `tfsdk:"proxy_terminates_https_conns"`
}

type IssuanceCriteriaData struct {
	ConditionalCriteria *[]*ConditionalIssuanceCriteriaEntryData `tfsdk:"conditional_criteria"`
	ExpressionCriteria  *[]*ExpressionIssuanceCriteriaEntryData  `tfsdk:"expression_criteria"`
}

type IssuerData struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Host        types.String `tfsdk:"host"`
	Path        types.String `tfsdk:"path"`
	Id          types.String `tfsdk:"id"`
}

type IssuersData struct {
	Items *[]*IssuerData `tfsdk:"items"`
}

type JdbcAttributeSourceData struct {
	Schema                       types.String                              `tfsdk:"schema"`
	Table                        types.String                              `tfsdk:"table"`
	ColumnNames                  []types.String                            `tfsdk:"column_names"`
	Filter                       types.String                              `tfsdk:"filter"`
	Id                           types.String                              `tfsdk:"id"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	DataStoreRef                 types.String                              `tfsdk:"data_store_ref"`
	Description                  types.String                              `tfsdk:"description"`
}

type JdbcDataStoreData struct {
	IdleTimeout               types.Number          `tfsdk:"idle_timeout"`
	ConnectionUrlTags         *[]*JdbcTagConfigData `tfsdk:"connection_url_tags"`
	UserName                  types.String          `tfsdk:"user_name"`
	MaxPoolSize               types.Number          `tfsdk:"max_pool_size"`
	MinPoolSize               types.Number          `tfsdk:"min_pool_size"`
	EncryptedPassword         types.String          `tfsdk:"encrypted_password"`
	AllowMultiValueAttributes types.Bool            `tfsdk:"allow_multi_value_attributes"`
	Type                      types.String          `tfsdk:"type"`
	Id                        types.String          `tfsdk:"id"`
	MaskAttributeValues       types.Bool            `tfsdk:"mask_attribute_values"`
	ValidateConnectionSql     types.String          `tfsdk:"validate_connection_sql"`
	BlockingTimeout           types.Number          `tfsdk:"blocking_timeout"`
	ConnectionUrl             types.String          `tfsdk:"connection_url"`
	Name                      types.String          `tfsdk:"name"`
	DriverClass               types.String          `tfsdk:"driver_class"`
	Password                  types.String          `tfsdk:"password"`
}

type JdbcTagConfigData struct {
	ConnectionUrl types.String `tfsdk:"connection_url"`
	Tags          types.String `tfsdk:"tags"`
	DefaultSource types.Bool   `tfsdk:"default_source"`
}

type JwksSettingsData struct {
	JwksUrl types.String `tfsdk:"jwks_url"`
	Jwks    types.String `tfsdk:"jwks"`
}

type KerberosKeySetData struct {
	EncryptedKeySet types.String `tfsdk:"encrypted_key_set"`
	DeactivatedAt   types.String `tfsdk:"deactivated_at"`
}

type KerberosRealmData struct {
	Id                                 types.String           `tfsdk:"id"`
	KeyDistributionCenters             []types.String         `tfsdk:"key_distribution_centers"`
	KerberosPassword                   types.String           `tfsdk:"kerberos_password"`
	KeySets                            *[]*KerberosKeySetData `tfsdk:"key_sets"`
	RetainPreviousKeysOnPasswordChange types.Bool             `tfsdk:"retain_previous_keys_on_password_change"`
	SuppressDomainNameConcatenation    types.Bool             `tfsdk:"suppress_domain_name_concatenation"`
	KerberosRealmName                  types.String           `tfsdk:"kerberos_realm_name"`
	KerberosUsername                   types.String           `tfsdk:"kerberos_username"`
	KerberosEncryptedPassword          types.String           `tfsdk:"kerberos_encrypted_password"`
}

type KerberosRealmsData struct {
	Items *[]*KerberosRealmData `tfsdk:"items"`
}

type KerberosRealmsSettingsData struct {
	ForceTcp                  types.Bool   `tfsdk:"force_tcp"`
	KdcRetries                types.String `tfsdk:"kdc_retries"`
	DebugLogOutput            types.Bool   `tfsdk:"debug_log_output"`
	KdcTimeout                types.String `tfsdk:"kdc_timeout"`
	KeySetRetentionPeriodMins types.Number `tfsdk:"key_set_retention_period_mins"`
}

type KeyAlgorithmData struct {
	Name                      types.String   `tfsdk:"name"`
	KeySizes                  []types.Number `tfsdk:"key_sizes"`
	DefaultKeySize            types.Number   `tfsdk:"default_key_size"`
	SignatureAlgorithms       []types.String `tfsdk:"signature_algorithms"`
	DefaultSignatureAlgorithm types.String   `tfsdk:"default_signature_algorithm"`
}

type KeyAlgorithmsData struct {
	Items *[]*KeyAlgorithmData `tfsdk:"items"`
}

type KeyPairExportSettingsData struct {
	Password types.String `tfsdk:"password"`
}

type KeyPairFileData struct {
	Id                types.String `tfsdk:"id"`
	FileData          types.String `tfsdk:"file_data"`
	Format            types.String `tfsdk:"format"`
	Password          types.String `tfsdk:"password"`
	EncryptedPassword types.String `tfsdk:"encrypted_password"`
	CryptoProvider    types.String `tfsdk:"crypto_provider"`
}

type KeyPairRotationSettingsData struct {
	Id                   types.String `tfsdk:"id"`
	CreationBufferDays   types.Number `tfsdk:"creation_buffer_days"`
	ActivationBufferDays types.Number `tfsdk:"activation_buffer_days"`
	ValidDays            types.Number `tfsdk:"valid_days"`
	KeyAlgorithm         types.String `tfsdk:"key_algorithm"`
	KeySize              types.Number `tfsdk:"key_size"`
	SignatureAlgorithm   types.String `tfsdk:"signature_algorithm"`
}

type KeyPairViewData struct {
	KeySize                 types.Number                 `tfsdk:"key_size"`
	SerialNumber            types.String                 `tfsdk:"serial_number"`
	IssuerDN                types.String                 `tfsdk:"issuer_dn"`
	Status                  types.String                 `tfsdk:"status"`
	Id                      types.String                 `tfsdk:"id"`
	Expires                 types.String                 `tfsdk:"expires"`
	KeyAlgorithm            types.String                 `tfsdk:"key_algorithm"`
	Sha1Fingerprint         types.String                 `tfsdk:"sha1fingerprint"`
	Sha256Fingerprint       types.String                 `tfsdk:"sha256fingerprint"`
	CryptoProvider          types.String                 `tfsdk:"crypto_provider"`
	SubjectDN               types.String                 `tfsdk:"subject_dn"`
	SubjectAlternativeNames []types.String               `tfsdk:"subject_alternative_names"`
	ValidFrom               types.String                 `tfsdk:"valid_from"`
	SignatureAlgorithm      types.String                 `tfsdk:"signature_algorithm"`
	Version                 types.Number                 `tfsdk:"version"`
	RotationSettings        *KeyPairRotationSettingsData `tfsdk:"rotation_settings"`
}

type KeyPairViewsData struct {
	Items *[]*KeyPairViewData `tfsdk:"items"`
}

type LdapAttributeSourceData struct {
	BinaryAttributeSettings      map[string]*BinaryLdapAttributeSettingsData `tfsdk:"binary_attribute_settings"`
	Id                           types.String                                `tfsdk:"id"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData   `tfsdk:"attribute_contract_fulfillment"`
	BaseDn                       types.String                                `tfsdk:"base_dn"`
	SearchScope                  types.String                                `tfsdk:"search_scope"`
	SearchAttributes             []types.String                              `tfsdk:"search_attributes"`
	DataStoreRef                 types.String                                `tfsdk:"data_store_ref"`
	Description                  types.String                                `tfsdk:"description"`
	MemberOfNestedGroup          types.Bool                                  `tfsdk:"member_of_nested_group"`
	SearchFilter                 types.String                                `tfsdk:"search_filter"`
}

type LdapDataStoreData struct {
	MaxConnections       types.Number          `tfsdk:"max_connections"`
	Name                 types.String          `tfsdk:"name"`
	UserDN               types.String          `tfsdk:"user_dn"`
	FollowLDAPReferrals  types.Bool            `tfsdk:"follow_ldap_referrals"`
	TestOnBorrow         types.Bool            `tfsdk:"test_on_borrow"`
	MaxWait              types.Number          `tfsdk:"max_wait"`
	HostnamesTags        *[]*LdapTagConfigData `tfsdk:"hostnames_tags"`
	BindAnonymously      types.Bool            `tfsdk:"bind_anonymously"`
	TimeBetweenEvictions types.Number          `tfsdk:"time_between_evictions"`
	EncryptedPassword    types.String          `tfsdk:"encrypted_password"`
	BinaryAttributes     []types.String        `tfsdk:"binary_attributes"`
	CreateIfNecessary    types.Bool            `tfsdk:"create_if_necessary"`
	DnsTtl               types.Number          `tfsdk:"dns_ttl"`
	Hostnames            []types.String        `tfsdk:"hostnames"`
	LdapDnsSrvPrefix     types.String          `tfsdk:"ldap_dns_srv_prefix"`
	Password             types.String          `tfsdk:"password"`
	LdapsDnsSrvPrefix    types.String          `tfsdk:"ldaps_dns_srv_prefix"`
	UseSsl               types.Bool            `tfsdk:"use_ssl"`
	MinConnections       types.Number          `tfsdk:"min_connections"`
	TestOnReturn         types.Bool            `tfsdk:"test_on_return"`
	Type                 types.String          `tfsdk:"type"`
	VerifyHost           types.Bool            `tfsdk:"verify_host"`
	ConnectionTimeout    types.Number          `tfsdk:"connection_timeout"`
	LdapType             types.String          `tfsdk:"ldap_type"`
	UseDnsSrvRecords     types.Bool            `tfsdk:"use_dns_srv_records"`
	ReadTimeout          types.Number          `tfsdk:"read_timeout"`
	Id                   types.String          `tfsdk:"id"`
	MaskAttributeValues  types.Bool            `tfsdk:"mask_attribute_values"`
}

type LdapDataStoreAttributeData struct {
	Metadata map[string]types.String `tfsdk:"metadata"`
	Type     types.String            `tfsdk:"type"`
	Name     types.String            `tfsdk:"name"`
}

type LdapDataStoreConfigData struct {
	CreatePattern          types.String                       `tfsdk:"create_pattern"`
	ObjectClass            types.String                       `tfsdk:"object_class"`
	AuxiliaryObjectClasses []types.String                     `tfsdk:"auxiliary_object_classes"`
	DataStoreRef           types.String                       `tfsdk:"data_store_ref"`
	DataStoreMapping       map[string]*DataStoreAttributeData `tfsdk:"data_store_mapping"`
	Type                   types.String                       `tfsdk:"type"`
	BaseDn                 types.String                       `tfsdk:"base_dn"`
}

type LdapTagConfigData struct {
	Hostnames     []types.String `tfsdk:"hostnames"`
	Tags          types.String   `tfsdk:"tags"`
	DefaultSource types.Bool     `tfsdk:"default_source"`
}

type LicenseAgreementInfoData struct {
	LicenseAgreementUrl types.String `tfsdk:"license_agreement_url"`
	Accepted            types.Bool   `tfsdk:"accepted"`
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
	OauthEnabled        types.Bool                         `tfsdk:"oauth_enabled"`
	BridgeMode          types.Bool                         `tfsdk:"bridge_mode"`
	Features            *[]*LicenseFeatureViewData         `tfsdk:"features"`
	Tier                types.String                       `tfsdk:"tier"`
	ExpirationDate      types.String                       `tfsdk:"expiration_date"`
	NodeLimit           types.Number                       `tfsdk:"node_limit"`
	Product             types.String                       `tfsdk:"product"`
	WsTrustEnabled      types.Bool                         `tfsdk:"ws_trust_enabled"`
	ProvisioningEnabled types.Bool                         `tfsdk:"provisioning_enabled"`
	Name                types.String                       `tfsdk:"name"`
	UsedConnections     types.Number                       `tfsdk:"used_connections"`
	IssueDate           types.String                       `tfsdk:"issue_date"`
	Id                  types.String                       `tfsdk:"id"`
	Organization        types.String                       `tfsdk:"organization"`
	GracePeriod         types.Number                       `tfsdk:"grace_period"`
	LicenseGroups       *[]*ConnectionGroupLicenseViewData `tfsdk:"license_groups"`
	MaxConnections      types.Number                       `tfsdk:"max_connections"`
	EnforcementType     types.String                       `tfsdk:"enforcement_type"`
	Version             types.String                       `tfsdk:"version"`
}

type LocalIdentityAuthSourceData struct {
	Source types.String `tfsdk:"source"`
	Id     types.String `tfsdk:"id"`
}

type LocalIdentityAuthSourceUpdatePolicyData struct {
	StoreAttributes  types.Bool   `tfsdk:"store_attributes"`
	RetainAttributes types.Bool   `tfsdk:"retain_attributes"`
	UpdateAttributes types.Bool   `tfsdk:"update_attributes"`
	UpdateInterval   types.Number `tfsdk:"update_interval"`
}

type LocalIdentityFieldData struct {
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	Id                    types.String          `tfsdk:"id"`
}

type LocalIdentityMappingPolicyActionData struct {
	OutboundAttributeMapping *AttributeMappingData `tfsdk:"outbound_attribute_mapping"`
	Type                     types.String          `tfsdk:"type"`
	Context                  types.String          `tfsdk:"context"`
	LocalIdentityRef         types.String          `tfsdk:"local_identity_ref"`
	InboundMapping           *AttributeMappingData `tfsdk:"inbound_mapping"`
}

type LocalIdentityProfileData struct {
	ProfileEnabled          types.Bool                               `tfsdk:"profile_enabled"`
	Id                      types.String                             `tfsdk:"id"`
	AuthSourceUpdatePolicy  *LocalIdentityAuthSourceUpdatePolicyData `tfsdk:"auth_source_update_policy"`
	RegistrationEnabled     types.Bool                               `tfsdk:"registration_enabled"`
	RegistrationConfig      *RegistrationConfigData                  `tfsdk:"registration_config"`
	ProfileConfig           *ProfileConfigData                       `tfsdk:"profile_config"`
	EmailVerificationConfig *EmailVerificationConfigData             `tfsdk:"email_verification_config"`
	DataStoreConfig         *DataStoreConfigData                     `tfsdk:"data_store_config"`
	Name                    types.String                             `tfsdk:"name"`
	ApcId                   types.String                             `tfsdk:"apc_id"`
	AuthSources             *[]*LocalIdentityAuthSourceData          `tfsdk:"auth_sources"`
	FieldConfig             *FieldConfigData                         `tfsdk:"field_config"`
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
	SigningKeyRef      types.String `tfsdk:"signing_key_ref"`
	SignatureAlgorithm types.String `tfsdk:"signature_algorithm"`
}

type MetadataUrlData struct {
	Id                types.String  `tfsdk:"id"`
	Name              types.String  `tfsdk:"name"`
	Url               types.String  `tfsdk:"url"`
	CertView          *CertViewData `tfsdk:"cert_view"`
	X509File          *X509FileData `tfsdk:"x509file"`
	ValidateSignature types.Bool    `tfsdk:"validate_signature"`
}

type MetadataUrlsData struct {
	Items *[]*MetadataUrlData `tfsdk:"items"`
}

type NewKeyPairSettingsData struct {
	KeySize                 types.Number   `tfsdk:"key_size"`
	SignatureAlgorithm      types.String   `tfsdk:"signature_algorithm"`
	CryptoProvider          types.String   `tfsdk:"crypto_provider"`
	CommonName              types.String   `tfsdk:"common_name"`
	Organization            types.String   `tfsdk:"organization"`
	Country                 types.String   `tfsdk:"country"`
	KeyAlgorithm            types.String   `tfsdk:"key_algorithm"`
	State                   types.String   `tfsdk:"state"`
	ValidDays               types.Number   `tfsdk:"valid_days"`
	Id                      types.String   `tfsdk:"id"`
	SubjectAlternativeNames []types.String `tfsdk:"subject_alternative_names"`
	OrganizationUnit        types.String   `tfsdk:"organization_unit"`
	City                    types.String   `tfsdk:"city"`
}

type NotificationPublisherData struct {
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
}

type NotificationPublishersData struct {
	Items *[]*NotificationPublisherData `tfsdk:"items"`
}

type NotificationPublishersSettingsData struct {
	DefaultNotificationPublisherRef types.String `tfsdk:"default_notification_publisher_ref"`
}

type NotificationSettingsData struct {
	LicenseEvents                          *LicenseEventNotificationSettingsData          `tfsdk:"license_events"`
	CertificateExpirations                 *CertificateExpirationNotificationSettingsData `tfsdk:"certificate_expirations"`
	NotifyAdminUserPasswordChanges         types.Bool                                     `tfsdk:"notify_admin_user_password_changes"`
	AccountChangesNotificationPublisherRef types.String                                   `tfsdk:"account_changes_notification_publisher_ref"`
	MetadataNotificationSettings           *MetadataEventNotificationSettingsData         `tfsdk:"metadata_notification_settings"`
}

type OAuthOidcKeysSettingsData struct {
	RsaDecryptionActiveCertRef        types.String `tfsdk:"rsa_decryption_active_cert_ref"`
	RsaDecryptionPublishX5cParameter  types.Bool   `tfsdk:"rsa_decryption_publish_x5c_parameter"`
	StaticJwksEnabled                 types.Bool   `tfsdk:"static_jwks_enabled"`
	P521ActiveCertRef                 types.String `tfsdk:"p521active_cert_ref"`
	P521PreviousCertRef               types.String `tfsdk:"p521previous_cert_ref"`
	P521DecryptionActiveCertRef       types.String `tfsdk:"p521decryption_active_cert_ref"`
	P521DecryptionPublishX5cParameter types.Bool   `tfsdk:"p521decryption_publish_x5c_parameter"`
	P384DecryptionPublishX5cParameter types.Bool   `tfsdk:"p384decryption_publish_x5c_parameter"`
	RsaDecryptionPreviousCertRef      types.String `tfsdk:"rsa_decryption_previous_cert_ref"`
	P256ActiveCertRef                 types.String `tfsdk:"p256active_cert_ref"`
	P384PreviousCertRef               types.String `tfsdk:"p384previous_cert_ref"`
	RsaPreviousCertRef                types.String `tfsdk:"rsa_previous_cert_ref"`
	P256DecryptionPreviousCertRef     types.String `tfsdk:"p256decryption_previous_cert_ref"`
	P384DecryptionPreviousCertRef     types.String `tfsdk:"p384decryption_previous_cert_ref"`
	P256PreviousCertRef               types.String `tfsdk:"p256previous_cert_ref"`
	P256PublishX5cParameter           types.Bool   `tfsdk:"p256publish_x5c_parameter"`
	P521PublishX5cParameter           types.Bool   `tfsdk:"p521publish_x5c_parameter"`
	P384DecryptionActiveCertRef       types.String `tfsdk:"p384decryption_active_cert_ref"`
	P521DecryptionPreviousCertRef     types.String `tfsdk:"p521decryption_previous_cert_ref"`
	P256DecryptionPublishX5cParameter types.Bool   `tfsdk:"p256decryption_publish_x5c_parameter"`
	P384ActiveCertRef                 types.String `tfsdk:"p384active_cert_ref"`
	P384PublishX5cParameter           types.Bool   `tfsdk:"p384publish_x5c_parameter"`
	RsaActiveCertRef                  types.String `tfsdk:"rsa_active_cert_ref"`
	RsaPublishX5cParameter            types.Bool   `tfsdk:"rsa_publish_x5c_parameter"`
	P256DecryptionActiveCertRef       types.String `tfsdk:"p256decryption_active_cert_ref"`
}

type OAuthRoleData struct {
	EnableOpenIdConnect types.Bool `tfsdk:"enable_open_id_connect"`
	EnableOauth         types.Bool `tfsdk:"enable_oauth"`
}

type OIDCClientCredentialsData struct {
	ClientId        types.String `tfsdk:"client_id"`
	ClientSecret    types.String `tfsdk:"client_secret"`
	EncryptedSecret types.String `tfsdk:"encrypted_secret"`
}

type OIDCProviderSettingsData struct {
	Scopes                         types.String                 `tfsdk:"scopes"`
	AuthorizationEndpoint          types.String                 `tfsdk:"authorization_endpoint"`
	RequestSigningAlgorithm        types.String                 `tfsdk:"request_signing_algorithm"`
	EnablePKCE                     types.Bool                   `tfsdk:"enable_pkce"`
	TokenEndpoint                  types.String                 `tfsdk:"token_endpoint"`
	JwksURL                        types.String                 `tfsdk:"jwks_url"`
	RequestParameters              *[]*OIDCRequestParameterData `tfsdk:"request_parameters"`
	LoginType                      types.String                 `tfsdk:"login_type"`
	AuthenticationScheme           types.String                 `tfsdk:"authentication_scheme"`
	AuthenticationSigningAlgorithm types.String                 `tfsdk:"authentication_signing_algorithm"`
	UserInfoEndpoint               types.String                 `tfsdk:"user_info_endpoint"`
}

type OIDCRequestParameterData struct {
	Name                        types.String                   `tfsdk:"name"`
	AttributeValue              *AttributeFulfillmentValueData `tfsdk:"attribute_value"`
	Value                       types.String                   `tfsdk:"value"`
	ApplicationEndpointOverride types.Bool                     `tfsdk:"application_endpoint_override"`
}

type OIDCSessionSettingsData struct {
	TrackUserSessionsForLogout types.Bool   `tfsdk:"track_user_sessions_for_logout"`
	RevokeUserSessionOnLogout  types.Bool   `tfsdk:"revoke_user_session_on_logout"`
	SessionRevocationLifetime  types.Number `tfsdk:"session_revocation_lifetime"`
}

type OcspSettingsData struct {
	ResponseCachePeriod          types.Number `tfsdk:"response_cache_period"`
	ResponderTimeout             types.Number `tfsdk:"responder_timeout"`
	ActionOnResponderUnavailable types.String `tfsdk:"action_on_responder_unavailable"`
	ActionOnUnsuccessfulResponse types.String `tfsdk:"action_on_unsuccessful_response"`
	ResponderCertReference       types.String `tfsdk:"responder_cert_reference"`
	ResponderUrl                 types.String `tfsdk:"responder_url"`
	CurrentUpdateGracePeriod     types.Number `tfsdk:"current_update_grace_period"`
	NextUpdateGracePeriod        types.Number `tfsdk:"next_update_grace_period"`
	ActionOnStatusUnknown        types.String `tfsdk:"action_on_status_unknown"`
	RequesterAddNonce            types.Bool   `tfsdk:"requester_add_nonce"`
}

type OpenIdConnectAttributeData struct {
	Name              types.String `tfsdk:"name"`
	IncludeInIdToken  types.Bool   `tfsdk:"include_in_id_token"`
	IncludeInUserInfo types.Bool   `tfsdk:"include_in_user_info"`
	MultiValued       types.Bool   `tfsdk:"multi_valued"`
}

type OpenIdConnectAttributeContractData struct {
	CoreAttributes     *[]*OpenIdConnectAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*OpenIdConnectAttributeData `tfsdk:"extended_attributes"`
}

type OpenIdConnectPoliciesData struct {
	Items *[]*OpenIdConnectPolicyData `tfsdk:"items"`
}

type OpenIdConnectPolicyData struct {
	ScopeAttributeMappings      map[string]*ParameterValuesData     `tfsdk:"scope_attribute_mappings"`
	Id                          types.String                        `tfsdk:"id"`
	Name                        types.String                        `tfsdk:"name"`
	AccessTokenManagerRef       types.String                        `tfsdk:"access_token_manager_ref"`
	IncludeUserInfoInIdToken    types.Bool                          `tfsdk:"include_user_info_in_id_token"`
	IncludeSHashInIdToken       types.Bool                          `tfsdk:"include_s_hash_in_id_token"`
	ReissueIdTokenInHybridFlow  types.Bool                          `tfsdk:"reissue_id_token_in_hybrid_flow"`
	AttributeContract           *OpenIdConnectAttributeContractData `tfsdk:"attribute_contract"`
	IdTokenLifetime             types.Number                        `tfsdk:"id_token_lifetime"`
	IncludeSriInIdToken         types.Bool                          `tfsdk:"include_sri_in_id_token"`
	ReturnIdTokenOnRefreshGrant types.Bool                          `tfsdk:"return_id_token_on_refresh_grant"`
	AttributeMapping            *AttributeMappingData               `tfsdk:"attribute_mapping"`
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
	Id                  types.String                        `tfsdk:"id"`
	AttributeContract   *OutOfBandAuthAttributeContractData `tfsdk:"attribute_contract"`
	Name                types.String                        `tfsdk:"name"`
	PluginDescriptorRef types.String                        `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String                        `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData            `tfsdk:"configuration"`
}

type OutOfBandAuthenticatorsData struct {
	Items *[]*OutOfBandAuthenticatorData `tfsdk:"items"`
}

type OutboundBackChannelAuthData struct {
	Type                 types.String                     `tfsdk:"type"`
	HttpBasicCredentials *UsernamePasswordCredentialsData `tfsdk:"http_basic_credentials"`
	DigitalSignature     types.Bool                       `tfsdk:"digital_signature"`
	SslAuthKeyPairRef    types.String                     `tfsdk:"ssl_auth_key_pair_ref"`
	ValidatePartnerCert  types.Bool                       `tfsdk:"validate_partner_cert"`
}

type OutboundProvisionData struct {
	Type           types.String        `tfsdk:"type"`
	TargetSettings *[]*ConfigFieldData `tfsdk:"target_settings"`
	CustomSchema   *SchemaData         `tfsdk:"custom_schema"`
	Channels       *[]*ChannelData     `tfsdk:"channels"`
}

type OutboundProvisionDatabaseData struct {
	DataStoreRef             types.String `tfsdk:"data_store_ref"`
	SynchronizationFrequency types.Number `tfsdk:"synchronization_frequency"`
}

type P14EKeyPairViewData struct {
	CurrentAuthenticationKey  types.Bool    `tfsdk:"current_authentication_key"`
	PreviousAuthenticationKey types.Bool    `tfsdk:"previous_authentication_key"`
	KeyPairView               *CertViewData `tfsdk:"key_pair_view"`
	CreationTime              types.String  `tfsdk:"creation_time"`
}

type P14EKeysViewData struct {
	KeyPairs *[]*P14EKeyPairViewData `tfsdk:"key_pairs"`
}

type ParameterValuesData struct {
	Values []types.String `tfsdk:"values"`
}

type PasswordCredentialValidatorData struct {
	PluginDescriptorRef types.String                                      `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String                                      `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData                          `tfsdk:"configuration"`
	AttributeContract   *PasswordCredentialValidatorAttributeContractData `tfsdk:"attribute_contract"`
	Id                  types.String                                      `tfsdk:"id"`
	Name                types.String                                      `tfsdk:"name"`
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

type PatternData struct {
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
	Type                  types.String          `tfsdk:"type"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
}

type PingOneConnectionData struct {
	Description                      types.String `tfsdk:"description"`
	EncryptedCredential              types.String `tfsdk:"encrypted_credential"`
	CredentialId                     types.String `tfsdk:"credential_id"`
	OrganizationName                 types.String `tfsdk:"organization_name"`
	Region                           types.String `tfsdk:"region"`
	Name                             types.String `tfsdk:"name"`
	Credential                       types.String `tfsdk:"credential"`
	PingOneConnectionId              types.String `tfsdk:"ping_one_connection_id"`
	PingOneManagementApiEndpoint     types.String `tfsdk:"ping_one_management_api_endpoint"`
	PingOneAuthenticationApiEndpoint types.String `tfsdk:"ping_one_authentication_api_endpoint"`
	Id                               types.String `tfsdk:"id"`
	CreationDate                     types.String `tfsdk:"creation_date"`
	Active                           types.Bool   `tfsdk:"active"`
	EnvironmentId                    types.String `tfsdk:"environment_id"`
}

type PingOneConnectionsData struct {
	Items *[]*PingOneConnectionData `tfsdk:"items"`
}

type PingOneCredentialStatusData struct {
	PingOneCredentialStatus types.String `tfsdk:"ping_one_credential_status"`
}

type PingOneEnvironmentData struct {
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
	Id   types.String `tfsdk:"id"`
}

type PingOneEnvironmentsData struct {
	Items *[]*PingOneEnvironmentData `tfsdk:"items"`
}

type PingOneForEnterpriseSettingsData struct {
	CurrentAuthnKeyCreationTime      types.String `tfsdk:"current_authn_key_creation_time"`
	PreviousAuthnKeyCreationTime     types.String `tfsdk:"previous_authn_key_creation_time"`
	IdentityRepositoryUpdateRequired types.Bool   `tfsdk:"identity_repository_update_required"`
	ConnectedToPingOneForEnterprise  types.Bool   `tfsdk:"connected_to_ping_one_for_enterprise"`
	PingOneSsoConnection             types.String `tfsdk:"ping_one_sso_connection"`
	CompanyName                      types.String `tfsdk:"company_name"`
	EnableAdminConsoleSso            types.Bool   `tfsdk:"enable_admin_console_sso"`
	EnableMonitoring                 types.Bool   `tfsdk:"enable_monitoring"`
}

type PingOneLdapGatewayDataStoreData struct {
	LdapType             types.String   `tfsdk:"ldap_type"`
	PingOneEnvironmentId types.String   `tfsdk:"ping_one_environment_id"`
	UseSsl               types.Bool     `tfsdk:"use_ssl"`
	Name                 types.String   `tfsdk:"name"`
	Id                   types.String   `tfsdk:"id"`
	MaskAttributeValues  types.Bool     `tfsdk:"mask_attribute_values"`
	PingOneLdapGatewayId types.String   `tfsdk:"ping_one_ldap_gateway_id"`
	BinaryAttributes     []types.String `tfsdk:"binary_attributes"`
	Type                 types.String   `tfsdk:"type"`
	PingOneConnectionRef types.String   `tfsdk:"ping_one_connection_ref"`
}

type PluginConfigurationData struct {
	Tables *[]*ConfigTableData `tfsdk:"tables"`
	Fields *[]*ConfigFieldData `tfsdk:"fields"`
}

type PluginInstanceData struct {
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
}

type PolicyActionData struct {
	Type    types.String `tfsdk:"type"`
	Context types.String `tfsdk:"context"`
}

type ProcessorPolicyToGeneratorMappingData struct {
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	Id                               types.String                              `tfsdk:"id"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
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

type RedirectValidationLocalSettingsData struct {
	WhiteList                                     *[]*RedirectValidationSettingsWhitelistEntryData `tfsdk:"white_list"`
	EnableTargetResourceValidationForSSO          types.Bool                                       `tfsdk:"enable_target_resource_validation_for_sso"`
	EnableTargetResourceValidationForSLO          types.Bool                                       `tfsdk:"enable_target_resource_validation_for_slo"`
	EnableTargetResourceValidationForIdpDiscovery types.Bool                                       `tfsdk:"enable_target_resource_validation_for_idp_discovery"`
	EnableInErrorResourceValidation               types.Bool                                       `tfsdk:"enable_in_error_resource_validation"`
}

type RedirectValidationPartnerSettingsData struct {
	EnableWreplyValidationSLO types.Bool `tfsdk:"enable_wreply_validation_slo"`
}

type RedirectValidationSettingsData struct {
	RedirectValidationLocalSettings   *RedirectValidationLocalSettingsData   `tfsdk:"redirect_validation_local_settings"`
	RedirectValidationPartnerSettings *RedirectValidationPartnerSettingsData `tfsdk:"redirect_validation_partner_settings"`
}

type RedirectValidationSettingsWhitelistEntryData struct {
	TargetResourceSSO     types.Bool   `tfsdk:"target_resource_sso"`
	TargetResourceSLO     types.Bool   `tfsdk:"target_resource_slo"`
	InErrorResource       types.Bool   `tfsdk:"in_error_resource"`
	IdpDiscovery          types.Bool   `tfsdk:"idp_discovery"`
	ValidDomain           types.String `tfsdk:"valid_domain"`
	ValidPath             types.String `tfsdk:"valid_path"`
	AllowQueryAndFragment types.Bool   `tfsdk:"allow_query_and_fragment"`
	RequireHttps          types.Bool   `tfsdk:"require_https"`
}

type RegistrationConfigData struct {
	ExecuteWorkflow                     types.String `tfsdk:"execute_workflow"`
	CaptchaEnabled                      types.Bool   `tfsdk:"captcha_enabled"`
	TemplateName                        types.String `tfsdk:"template_name"`
	CreateAuthnSessionAfterRegistration types.Bool   `tfsdk:"create_authn_session_after_registration"`
	UsernameField                       types.String `tfsdk:"username_field"`
	ThisIsMyDeviceEnabled               types.Bool   `tfsdk:"this_is_my_device_enabled"`
	RegistrationWorkflow                types.String `tfsdk:"registration_workflow"`
}

type RequestPoliciesData struct {
	Items *[]*RequestPolicyData `tfsdk:"items"`
}

type RequestPolicyData struct {
	Name                             types.String                            `tfsdk:"name"`
	AuthenticatorRef                 types.String                            `tfsdk:"authenticator_ref"`
	AllowUnsignedLoginHintToken      types.Bool                              `tfsdk:"allow_unsigned_login_hint_token"`
	RequireTokenForIdentityHint      types.Bool                              `tfsdk:"require_token_for_identity_hint"`
	AlternativeLoginHintTokenIssuers *[]*AlternativeLoginHintTokenIssuerData `tfsdk:"alternative_login_hint_token_issuers"`
	IdentityHintContract             *IdentityHintContractData               `tfsdk:"identity_hint_contract"`
	IdentityHintContractFulfillment  *AttributeMappingData                   `tfsdk:"identity_hint_contract_fulfillment"`
	Id                               types.String                            `tfsdk:"id"`
	TransactionLifetime              types.Number                            `tfsdk:"transaction_lifetime"`
	IdentityHintMapping              *AttributeMappingData                   `tfsdk:"identity_hint_mapping"`
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
	Id                           types.String                              `tfsdk:"id"`
	PasswordValidatorRef         types.String                              `tfsdk:"password_validator_ref"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
}

type ResourceOwnerCredentialsMappingsData struct {
	Items *[]*ResourceOwnerCredentialsMappingData `tfsdk:"items"`
}

type ResourceUsageData struct {
	Ref        types.String `tfsdk:"ref"`
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	CategoryId types.String `tfsdk:"category_id"`
	Type       types.String `tfsdk:"type"`
}

type ResourceUsagesData struct {
	Categories *[]*ResourceCategoryInfoData `tfsdk:"categories"`
	Items      *[]*ResourceUsageData        `tfsdk:"items"`
}

type RestartPolicyActionData struct {
	Type    types.String `tfsdk:"type"`
	Context types.String `tfsdk:"context"`
}

type RolesAndProtocolsData struct {
	OauthRole          *OAuthRoleData `tfsdk:"oauth_role"`
	IdpRole            *IdpRoleData   `tfsdk:"idp_role"`
	SpRole             *SpRoleData    `tfsdk:"sp_role"`
	EnableIdpDiscovery types.Bool     `tfsdk:"enable_idp_discovery"`
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
	Trim           types.Bool     `tfsdk:"trim"`
	CharacterCase  types.String   `tfsdk:"character_case"`
	Parser         types.String   `tfsdk:"parser"`
	Masked         types.Bool     `tfsdk:"masked"`
	AttributeNames []types.String `tfsdk:"attribute_names"`
	DefaultValue   types.String   `tfsdk:"default_value"`
	Expression     types.String   `tfsdk:"expression"`
	CreateOnly     types.Bool     `tfsdk:"create_only"`
}

type SaasPluginFieldInfoDescriptorData struct {
	MinLength            types.Number                  `tfsdk:"min_length"`
	MaxLength            types.Number                  `tfsdk:"max_length"`
	PersistForMembership types.Bool                    `tfsdk:"persist_for_membership"`
	Unique               types.Bool                    `tfsdk:"unique"`
	DefaultValue         types.String                  `tfsdk:"default_value"`
	DsLdapMap            types.Bool                    `tfsdk:"ds_ldap_map"`
	Options              *[]*SaasPluginFieldOptionData `tfsdk:"options"`
	Required             types.Bool                    `tfsdk:"required"`
	MultiValue           types.Bool                    `tfsdk:"multi_value"`
	Pattern              *PatternData                  `tfsdk:"pattern"`
	AttributeGroup       types.Bool                    `tfsdk:"attribute_group"`
	Label                types.String                  `tfsdk:"label"`
	Notes                []types.String                `tfsdk:"notes"`
	Code                 types.String                  `tfsdk:"code"`
}

type SaasPluginFieldOptionData struct {
	Code  types.String `tfsdk:"code"`
	Label types.String `tfsdk:"label"`
}

type SchemaData struct {
	Namespace  types.String            `tfsdk:"namespace"`
	Attributes *[]*SchemaAttributeData `tfsdk:"attributes"`
}

type SchemaAttributeData struct {
	Name          types.String   `tfsdk:"name"`
	MultiValued   types.Bool     `tfsdk:"multi_valued"`
	Types         []types.String `tfsdk:"types"`
	SubAttributes []types.String `tfsdk:"sub_attributes"`
}

type ScopeEntryData struct {
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Dynamic     types.Bool   `tfsdk:"dynamic"`
}

type ScopeGroupEntryData struct {
	Name        types.String   `tfsdk:"name"`
	Description types.String   `tfsdk:"description"`
	Scopes      []types.String `tfsdk:"scopes"`
}

type SecretManagerData struct {
	PluginDescriptorRef types.String             `tfsdk:"plugin_descriptor_ref"`
	ParentRef           types.String             `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData `tfsdk:"configuration"`
	Id                  types.String             `tfsdk:"id"`
	Name                types.String             `tfsdk:"name"`
}

type SecretManagersData struct {
	Items *[]*SecretManagerData `tfsdk:"items"`
}

type ServerSettingsData struct {
	FederationInfo    *FederationInfoData       `tfsdk:"federation_info"`
	EmailServer       *EmailServerSettingsData  `tfsdk:"email_server"`
	CaptchaSettings   *CaptchaSettingsData      `tfsdk:"captcha_settings"`
	ContactInfo       *ContactInfoData          `tfsdk:"contact_info"`
	Notifications     *NotificationSettingsData `tfsdk:"notifications"`
	RolesAndProtocols *RolesAndProtocolsData    `tfsdk:"roles_and_protocols"`
}

type ServiceAssociationData struct {
	Configured    types.Bool     `tfsdk:"configured"`
	ComponentName types.String   `tfsdk:"component_name"`
	ServiceNames  []types.String `tfsdk:"service_names"`
}

type ServiceAssociationsData struct {
	Items *[]*ServiceAssociationData `tfsdk:"items"`
}

type ServiceAuthenticationData struct {
	AttributeQuery       *ServiceModelData `tfsdk:"attribute_query"`
	Jmx                  *ServiceModelData `tfsdk:"jmx"`
	ConnectionManagement *ServiceModelData `tfsdk:"connection_management"`
	SsoDirectoryService  *ServiceModelData `tfsdk:"sso_directory_service"`
}

type ServiceModelData struct {
	Id                    types.String `tfsdk:"id"`
	SharedSecret          types.String `tfsdk:"shared_secret"`
	EncryptedSharedSecret types.String `tfsdk:"encrypted_shared_secret"`
}

type SessionSettingsData struct {
	TrackAdapterSessionsForLogout types.Bool   `tfsdk:"track_adapter_sessions_for_logout"`
	RevokeUserSessionOnLogout     types.Bool   `tfsdk:"revoke_user_session_on_logout"`
	SessionRevocationLifetime     types.Number `tfsdk:"session_revocation_lifetime"`
}

type SessionValidationSettingsData struct {
	UpdateAuthnSessionActivity   types.Bool `tfsdk:"update_authn_session_activity"`
	Inherited                    types.Bool `tfsdk:"inherited"`
	IncludeSessionId             types.Bool `tfsdk:"include_session_id"`
	CheckValidAuthnSession       types.Bool `tfsdk:"check_valid_authn_session"`
	CheckSessionRevocationStatus types.Bool `tfsdk:"check_session_revocation_status"`
}

type SigningKeysData struct {
	P384PublishX5cParameter types.Bool   `tfsdk:"p384publish_x5c_parameter"`
	RsaActiveCertRef        types.String `tfsdk:"rsa_active_cert_ref"`
	P256PublishX5cParameter types.Bool   `tfsdk:"p256publish_x5c_parameter"`
	P384ActiveCertRef       types.String `tfsdk:"p384active_cert_ref"`
	P384PreviousCertRef     types.String `tfsdk:"p384previous_cert_ref"`
	P521ActiveCertRef       types.String `tfsdk:"p521active_cert_ref"`
	P521PreviousCertRef     types.String `tfsdk:"p521previous_cert_ref"`
	P521PublishX5cParameter types.Bool   `tfsdk:"p521publish_x5c_parameter"`
	P256ActiveCertRef       types.String `tfsdk:"p256active_cert_ref"`
	P256PreviousCertRef     types.String `tfsdk:"p256previous_cert_ref"`
	RsaPreviousCertRef      types.String `tfsdk:"rsa_previous_cert_ref"`
	RsaPublishX5cParameter  types.Bool   `tfsdk:"rsa_publish_x5c_parameter"`
}

type SigningSettingsData struct {
	SigningKeyPairRef        types.String `tfsdk:"signing_key_pair_ref"`
	Algorithm                types.String `tfsdk:"algorithm"`
	IncludeCertInSignature   types.Bool   `tfsdk:"include_cert_in_signature"`
	IncludeRawKeyInSignature types.Bool   `tfsdk:"include_raw_key_in_signature"`
}

type SloServiceEndpointData struct {
	Binding     types.String `tfsdk:"binding"`
	Url         types.String `tfsdk:"url"`
	ResponseUrl types.String `tfsdk:"response_url"`
}

type SourceTypeIdKeyData struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type SpAdapterData struct {
	Name                  types.String                        `tfsdk:"name"`
	PluginDescriptorRef   types.String                        `tfsdk:"plugin_descriptor_ref"`
	ParentRef             types.String                        `tfsdk:"parent_ref"`
	Configuration         *PluginConfigurationData            `tfsdk:"configuration"`
	AttributeContract     *SpAdapterAttributeContractData     `tfsdk:"attribute_contract"`
	TargetApplicationInfo *SpAdapterTargetApplicationInfoData `tfsdk:"target_application_info"`
	Id                    types.String                        `tfsdk:"id"`
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
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SpAdapterRef                 types.String                              `tfsdk:"sp_adapter_ref"`
	RestrictVirtualEntityIds     types.Bool                                `tfsdk:"restrict_virtual_entity_ids"`
	RestrictedVirtualEntityIds   []types.String                            `tfsdk:"restricted_virtual_entity_ids"`
	AdapterOverrideSettings      *SpAdapterData                            `tfsdk:"adapter_override_settings"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
}

type SpAdapterTargetApplicationInfoData struct {
	ApplicationName    types.String `tfsdk:"application_name"`
	ApplicationIconUrl types.String `tfsdk:"application_icon_url"`
	Inherited          types.Bool   `tfsdk:"inherited"`
}

type SpAdapterUrlMappingData struct {
	Url        types.String `tfsdk:"url"`
	AdapterRef types.String `tfsdk:"adapter_ref"`
}

type SpAdapterUrlMappingsData struct {
	Items *[]*SpAdapterUrlMappingData `tfsdk:"items"`
}

type SpAdaptersData struct {
	Items *[]*SpAdapterData `tfsdk:"items"`
}

type SpAttributeQueryData struct {
	Attributes                   []types.String                            `tfsdk:"attributes"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	Policy                       *SpAttributeQueryPolicyData               `tfsdk:"policy"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
}

type SpAttributeQueryPolicyData struct {
	RequireEncryptedNameId      types.Bool `tfsdk:"require_encrypted_name_id"`
	SignResponse                types.Bool `tfsdk:"sign_response"`
	SignAssertion               types.Bool `tfsdk:"sign_assertion"`
	EncryptAssertion            types.Bool `tfsdk:"encrypt_assertion"`
	RequireSignedAttributeQuery types.Bool `tfsdk:"require_signed_attribute_query"`
}

type SpBrowserSsoData struct {
	Artifact                                      *ArtifactSettingsData                                `tfsdk:"artifact"`
	SloServiceEndpoints                           *[]*SloServiceEndpointData                           `tfsdk:"slo_service_endpoints"`
	DefaultTargetUrl                              types.String                                         `tfsdk:"default_target_url"`
	SsoServiceEndpoints                           *[]*SpSsoServiceEndpointData                         `tfsdk:"sso_service_endpoints"`
	SpSamlIdentityMapping                         types.String                                         `tfsdk:"sp_saml_identity_mapping"`
	Protocol                                      types.String                                         `tfsdk:"protocol"`
	WsFedTokenType                                types.String                                         `tfsdk:"ws_fed_token_type"`
	EnabledProfiles                               []types.String                                       `tfsdk:"enabled_profiles"`
	AttributeContract                             *SpBrowserSsoAttributeContractData                   `tfsdk:"attribute_contract"`
	SpWsFedIdentityMapping                        types.String                                         `tfsdk:"sp_ws_fed_identity_mapping"`
	SignAssertions                                types.Bool                                           `tfsdk:"sign_assertions"`
	RequireSignedAuthnRequests                    types.Bool                                           `tfsdk:"require_signed_authn_requests"`
	WsTrustVersion                                types.String                                         `tfsdk:"ws_trust_version"`
	UrlWhitelistEntries                           *[]*UrlWhitelistEntryData                            `tfsdk:"url_whitelist_entries"`
	EncryptionPolicy                              *EncryptionPolicyData                                `tfsdk:"encryption_policy"`
	SignResponseAsRequired                        types.Bool                                           `tfsdk:"sign_response_as_required"`
	AssertionLifetime                             *AssertionLifetimeData                               `tfsdk:"assertion_lifetime"`
	IncomingBindings                              []types.String                                       `tfsdk:"incoming_bindings"`
	MessageCustomizations                         *[]*ProtocolMessageCustomizationData                 `tfsdk:"message_customizations"`
	AlwaysSignArtifactResponse                    types.Bool                                           `tfsdk:"always_sign_artifact_response"`
	AdapterMappings                               *[]*IdpAdapterAssertionMappingData                   `tfsdk:"adapter_mappings"`
	AuthenticationPolicyContractAssertionMappings *[]*AuthenticationPolicyContractAssertionMappingData `tfsdk:"authentication_policy_contract_assertion_mappings"`
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
	BaseUrl                                types.String                                `tfsdk:"base_url"`
	SpBrowserSso                           *SpBrowserSsoData                           `tfsdk:"sp_browser_sso"`
	Credentials                            *ConnectionCredentialsData                  `tfsdk:"credentials"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfigurationData `tfsdk:"additional_allowed_entities_configuration"`
	ExtendedProperties                     map[string]*ParameterValuesData             `tfsdk:"extended_properties"`
	Type                                   types.String                                `tfsdk:"type"`
	AttributeQuery                         *SpAttributeQueryData                       `tfsdk:"attribute_query"`
	ApplicationName                        types.String                                `tfsdk:"application_name"`
	ApplicationIconUrl                     types.String                                `tfsdk:"application_icon_url"`
	VirtualEntityIds                       []types.String                              `tfsdk:"virtual_entity_ids"`
	MetadataReloadSettings                 *ConnectionMetadataUrlData                  `tfsdk:"metadata_reload_settings"`
	LicenseConnectionGroup                 types.String                                `tfsdk:"license_connection_group"`
	LoggingMode                            types.String                                `tfsdk:"logging_mode"`
	OutboundProvision                      *OutboundProvisionData                      `tfsdk:"outbound_provision"`
	EntityId                               types.String                                `tfsdk:"entity_id"`
	Name                                   types.String                                `tfsdk:"name"`
	DefaultVirtualEntityId                 types.String                                `tfsdk:"default_virtual_entity_id"`
	ContactInfo                            *ContactInfoData                            `tfsdk:"contact_info"`
	Id                                     types.String                                `tfsdk:"id"`
	ConnectionTargetType                   types.String                                `tfsdk:"connection_target_type"`
	WsTrust                                *SpWsTrustData                              `tfsdk:"ws_trust"`
}

type SpConnectionsData struct {
	Items *[]*SpConnectionData `tfsdk:"items"`
}

type SpDefaultUrlsData struct {
	SsoSuccessUrl types.String `tfsdk:"sso_success_url"`
	ConfirmSlo    types.Bool   `tfsdk:"confirm_slo"`
	SloSuccessUrl types.String `tfsdk:"slo_success_url"`
}

type SpRoleData struct {
	EnableWsFed               types.Bool           `tfsdk:"enable_ws_fed"`
	EnableWsTrust             types.Bool           `tfsdk:"enable_ws_trust"`
	Enable                    types.Bool           `tfsdk:"enable"`
	Saml20Profile             *SpSAML20ProfileData `tfsdk:"saml20profile"`
	EnableOpenIDConnect       types.Bool           `tfsdk:"enable_open_id_connect"`
	EnableInboundProvisioning types.Bool           `tfsdk:"enable_inbound_provisioning"`
	EnableSaml11              types.Bool           `tfsdk:"enable_saml11"`
	EnableSaml10              types.Bool           `tfsdk:"enable_saml10"`
}

type SpSAML20ProfileData struct {
	Enable            types.Bool `tfsdk:"enable"`
	EnableAutoConnect types.Bool `tfsdk:"enable_auto_connect"`
	EnableXASP        types.Bool `tfsdk:"enable_xasp"`
}

type SpSsoServiceEndpointData struct {
	IsDefault types.Bool   `tfsdk:"is_default"`
	Index     types.Number `tfsdk:"index"`
	Binding   types.String `tfsdk:"binding"`
	Url       types.String `tfsdk:"url"`
}

type SpTokenGeneratorMappingData struct {
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SpTokenGeneratorRef          types.String                              `tfsdk:"sp_token_generator_ref"`
	RestrictedVirtualEntityIds   []types.String                            `tfsdk:"restricted_virtual_entity_ids"`
	DefaultMapping               types.Bool                                `tfsdk:"default_mapping"`
}

type SpUrlMappingData struct {
	Type types.String `tfsdk:"type"`
	Ref  types.String `tfsdk:"ref"`
	Url  types.String `tfsdk:"url"`
}

type SpUrlMappingsData struct {
	Items *[]*SpUrlMappingData `tfsdk:"items"`
}

type SpWsTrustData struct {
	MinutesAfter                   types.Number                         `tfsdk:"minutes_after"`
	AttributeContract              *SpWsTrustAttributeContractData      `tfsdk:"attribute_contract"`
	AbortIfNotFulfilledFromRequest types.Bool                           `tfsdk:"abort_if_not_fulfilled_from_request"`
	RequestContractRef             types.String                         `tfsdk:"request_contract_ref"`
	MessageCustomizations          *[]*ProtocolMessageCustomizationData `tfsdk:"message_customizations"`
	PartnerServiceIds              []types.String                       `tfsdk:"partner_service_ids"`
	OAuthAssertionProfiles         types.Bool                           `tfsdk:"o_auth_assertion_profiles"`
	EncryptSaml2Assertion          types.Bool                           `tfsdk:"encrypt_saml2assertion"`
	TokenProcessorMappings         *[]*IdpTokenProcessorMappingData     `tfsdk:"token_processor_mappings"`
	DefaultTokenType               types.String                         `tfsdk:"default_token_type"`
	GenerateKey                    types.Bool                           `tfsdk:"generate_key"`
	MinutesBefore                  types.Number                         `tfsdk:"minutes_before"`
}

type SpWsTrustAttributeData struct {
	Name      types.String `tfsdk:"name"`
	Namespace types.String `tfsdk:"namespace"`
}

type SpWsTrustAttributeContractData struct {
	ExtendedAttributes *[]*SpWsTrustAttributeData `tfsdk:"extended_attributes"`
	CoreAttributes     *[]*SpWsTrustAttributeData `tfsdk:"core_attributes"`
}

type SslServerSettingsData struct {
	ActiveRuntimeServerCerts *[]*ResourceLinkData `tfsdk:"active_runtime_server_certs"`
	ActiveAdminConsoleCerts  *[]*ResourceLinkData `tfsdk:"active_admin_console_certs"`
	RuntimeServerCertRef     types.String         `tfsdk:"runtime_server_cert_ref"`
	AdminConsoleCertRef      types.String         `tfsdk:"admin_console_cert_ref"`
}

type SsoOAuthMappingData struct {
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
}

type StsRequestParametersContractData struct {
	Id         types.String   `tfsdk:"id"`
	Name       types.String   `tfsdk:"name"`
	Parameters []types.String `tfsdk:"parameters"`
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
	Previous *SystemKeyData `tfsdk:"previous"`
	Pending  *SystemKeyData `tfsdk:"pending"`
}

type TextLocalIdentityFieldData struct {
	Attributes            map[string]types.Bool `tfsdk:"attributes"`
	Type                  types.String          `tfsdk:"type"`
	DefaultValue          types.String          `tfsdk:"default_value"`
	Id                    types.String          `tfsdk:"id"`
	Label                 types.String          `tfsdk:"label"`
	RegistrationPageField types.Bool            `tfsdk:"registration_page_field"`
	ProfilePageField      types.Bool            `tfsdk:"profile_page_field"`
}

type TokenExchangeGeneratorGroupData struct {
	Name              types.String                          `tfsdk:"name"`
	ResourceUris      []types.String                        `tfsdk:"resource_uris"`
	GeneratorMappings *[]*TokenExchangeGeneratorMappingData `tfsdk:"generator_mappings"`
	Id                types.String                          `tfsdk:"id"`
}

type TokenExchangeGeneratorGroupsData struct {
	Items *[]*TokenExchangeGeneratorGroupData `tfsdk:"items"`
}

type TokenExchangeGeneratorMappingData struct {
	RequestedTokenType types.String `tfsdk:"requested_token_type"`
	TokenGenerator     types.String `tfsdk:"token_generator"`
	DefaultMapping     types.Bool   `tfsdk:"default_mapping"`
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
	JdbcAttributeSources         []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources         []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources       []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria             *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SubjectTokenType             types.String                              `tfsdk:"subject_token_type"`
	SubjectTokenProcessor        types.String                              `tfsdk:"subject_token_processor"`
	ActorTokenType               types.String                              `tfsdk:"actor_token_type"`
	ActorTokenProcessor          types.String                              `tfsdk:"actor_token_processor"`
}

type TokenExchangeProcessorPoliciesData struct {
	Items *[]*TokenExchangeProcessorPolicyData `tfsdk:"items"`
}

type TokenExchangeProcessorPolicyData struct {
	Id                 types.String                                 `tfsdk:"id"`
	Name               types.String                                 `tfsdk:"name"`
	ActorTokenRequired types.Bool                                   `tfsdk:"actor_token_required"`
	AttributeContract  *TokenExchangeProcessorAttributeContractData `tfsdk:"attribute_contract"`
	ProcessorMappings  *[]*TokenExchangeProcessorMappingData        `tfsdk:"processor_mappings"`
}

type TokenExchangeProcessorSettingsData struct {
	DefaultProcessorPolicyRef types.String `tfsdk:"default_processor_policy_ref"`
}

type TokenGeneratorData struct {
	ParentRef           types.String                         `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData             `tfsdk:"configuration"`
	Id                  types.String                         `tfsdk:"id"`
	Name                types.String                         `tfsdk:"name"`
	PluginDescriptorRef types.String                         `tfsdk:"plugin_descriptor_ref"`
	AttributeContract   *TokenGeneratorAttributeContractData `tfsdk:"attribute_contract"`
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
	ParentRef           types.String                         `tfsdk:"parent_ref"`
	Configuration       *PluginConfigurationData             `tfsdk:"configuration"`
	AttributeContract   *TokenProcessorAttributeContractData `tfsdk:"attribute_contract"`
	Id                  types.String                         `tfsdk:"id"`
	Name                types.String                         `tfsdk:"name"`
	PluginDescriptorRef types.String                         `tfsdk:"plugin_descriptor_ref"`
}

type TokenProcessorAttributeData struct {
	Name   types.String `tfsdk:"name"`
	Masked types.Bool   `tfsdk:"masked"`
}

type TokenProcessorAttributeContractData struct {
	CoreAttributes     *[]*TokenProcessorAttributeData `tfsdk:"core_attributes"`
	ExtendedAttributes *[]*TokenProcessorAttributeData `tfsdk:"extended_attributes"`
	MaskOgnlValues     types.Bool                      `tfsdk:"mask_ognl_values"`
	Inherited          types.Bool                      `tfsdk:"inherited"`
}

type TokenProcessorsData struct {
	Items *[]*TokenProcessorData `tfsdk:"items"`
}

type TokenToTokenMappingData struct {
	Id                               types.String                              `tfsdk:"id"`
	DefaultTargetResource            types.String                              `tfsdk:"default_target_resource"`
	LicenseConnectionGroupAssignment types.String                              `tfsdk:"license_connection_group_assignment"`
	JdbcAttributeSources             []JdbcAttributeSourceData                 `tfsdk:"jdbc_attribute_sources"`
	LdapAttributeSources             []LdapAttributeSourceData                 `tfsdk:"ldap_attribute_sources"`
	CustomAttributeSources           []CustomAttributeSourceData               `tfsdk:"custom_attribute_sources"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValueData `tfsdk:"attribute_contract_fulfillment"`
	IssuanceCriteria                 *IssuanceCriteriaData                     `tfsdk:"issuance_criteria"`
	SourceId                         types.String                              `tfsdk:"source_id"`
	TargetId                         types.String                              `tfsdk:"target_id"`
}

type TokenToTokenMappingsData struct {
	Items *[]*TokenToTokenMappingData `tfsdk:"items"`
}

type UrlWhitelistEntryData struct {
	ValidPath             types.String `tfsdk:"valid_path"`
	AllowQueryAndFragment types.Bool   `tfsdk:"allow_query_and_fragment"`
	RequireHttps          types.Bool   `tfsdk:"require_https"`
	ValidDomain           types.String `tfsdk:"valid_domain"`
}

type UserCredentialsData struct {
	CurrentPassword types.String `tfsdk:"current_password"`
	NewPassword     types.String `tfsdk:"new_password"`
}

type UsernamePasswordCredentialsData struct {
	Username          types.String `tfsdk:"username"`
	Password          types.String `tfsdk:"password"`
	EncryptedPassword types.String `tfsdk:"encrypted_password"`
}

type ValidationErrorData struct {
	DeveloperMessage types.String `tfsdk:"developer_message"`
	FieldPath        types.String `tfsdk:"field_path"`
	ErrorId          types.String `tfsdk:"error_id"`
	Message          types.String `tfsdk:"message"`
}

type VersionData struct {
	Version types.String `tfsdk:"version"`
}

type VirtualHostNameSettingsData struct {
	VirtualHostNames []types.String `tfsdk:"virtual_host_names"`
}

type X509FileData struct {
	FileData       types.String `tfsdk:"file_data"`
	CryptoProvider types.String `tfsdk:"crypto_provider"`
	Id             types.String `tfsdk:"id"`
}
