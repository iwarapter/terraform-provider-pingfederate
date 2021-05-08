package models

import "encoding/json"

//AccessTokenAttribute - An attribute for an Access Token's attribute contract.
type AccessTokenAttribute struct {
	Name *string `json:"name,omitempty"`
}

//AccessTokenAttributeContract - A set of attributes exposed by an Access Token Manager.
type AccessTokenAttributeContract struct {
	CoreAttributes          *[]*AccessTokenAttribute `json:"coreAttributes,omitempty"`
	DefaultSubjectAttribute *string                  `json:"defaultSubjectAttribute,omitempty"`
	ExtendedAttributes      *[]*AccessTokenAttribute `json:"extendedAttributes,omitempty"`
	Inherited               *bool                    `json:"inherited,omitempty"`
}

//AccessTokenManagementSettings - General access token management settings.
type AccessTokenManagementSettings struct {
	DefaultAccessTokenManagerRef *ResourceLink `json:"defaultAccessTokenManagerRef,omitempty"`
}

//AccessTokenManager - An OAuth access token management plugin instance.
type AccessTokenManager struct {
	AccessControlSettings     *AtmAccessControlSettings     `json:"accessControlSettings,omitempty"`
	AttributeContract         *AccessTokenAttributeContract `json:"attributeContract,omitempty"`
	Configuration             *PluginConfiguration          `json:"configuration,omitempty"`
	Id                        *string                       `json:"id,omitempty"`
	Name                      *string                       `json:"name,omitempty"`
	ParentRef                 *ResourceLink                 `json:"parentRef,omitempty"`
	PluginDescriptorRef       *ResourceLink                 `json:"pluginDescriptorRef,omitempty"`
	SelectionSettings         *AtmSelectionSettings         `json:"selectionSettings,omitempty"`
	SessionValidationSettings *SessionValidationSettings    `json:"sessionValidationSettings,omitempty"`
}

//AccessTokenManagerDescriptor - An OAuth access token management plugin descriptor.
type AccessTokenManagerDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//AccessTokenManagerDescriptors - A collection of OAuth access token management plugin descriptors.
type AccessTokenManagerDescriptors struct {
	Items *[]*AccessTokenManagerDescriptor `json:"items,omitempty"`
}

//AccessTokenManagerMapping - A mapping in a connection that defines how access tokens are created.
type AccessTokenManagerMapping struct {
	AccessTokenManagerRef        *ResourceLink                         `json:"accessTokenManagerRef,omitempty"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//AccessTokenManagers - A collection of OAuth access token management plugin instances.
type AccessTokenManagers struct {
	Items *[]*AccessTokenManager `json:"items,omitempty"`
}

//AccessTokenMapping - The Access Token Attribute Mapping.
type AccessTokenMapping struct {
	AccessTokenManagerRef        *ResourceLink                         `json:"accessTokenManagerRef,omitempty"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Context                      *AccessTokenMappingContext            `json:"context,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//AccessTokenMappingContext - The Access Token Attribute Mapping.
type AccessTokenMappingContext struct {
	ContextRef *ResourceLink `json:"contextRef,omitempty"`
	Type       *string       `json:"type,omitempty"`
}

//AccessTokenMappings - A collection of Access Token Attribute Mapping items.
type AccessTokenMappings []*AccessTokenMapping

//AccountManagementSettings - Account management settings.
type AccountManagementSettings struct {
	AccountStatusAlgorithm     *string `json:"accountStatusAlgorithm,omitempty"`
	AccountStatusAttributeName *string `json:"accountStatusAttributeName,omitempty"`
	DefaultStatus              *bool   `json:"defaultStatus,omitempty"`
	FlagComparisonStatus       *bool   `json:"flagComparisonStatus,omitempty"`
	FlagComparisonValue        *string `json:"flagComparisonValue,omitempty"`
}

//Action - A read-only plugin action that either represents a file download or an arbitrary invocation performed by the plugin.
type Action struct {
	Description   *string       `json:"description,omitempty"`
	Download      *bool         `json:"download,omitempty"`
	Id            *string       `json:"id,omitempty"`
	InvocationRef *ResourceLink `json:"invocationRef,omitempty"`
	Name          *string       `json:"name,omitempty"`
}

//ActionDescriptor - Describes an arbitrary action that is available for a plugin.
type ActionDescriptor struct {
	Description         *string `json:"description,omitempty"`
	Download            *bool   `json:"download,omitempty"`
	DownloadContentType *string `json:"downloadContentType,omitempty"`
	DownloadFileName    *string `json:"downloadFileName,omitempty"`
	Name                *string `json:"name,omitempty"`
}

//ActionResult - The result for non-download plugin actions.
type ActionResult struct {
	Message *string `json:"message,omitempty"`
}

//Actions - A read-only list of available actions for this plugin instance.
type Actions struct {
	Items *[]*Action `json:"items,omitempty"`
}

//AdditionalAllowedEntitiesConfiguration - Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.
type AdditionalAllowedEntitiesConfiguration struct {
	AdditionalAllowedEntities *[]*Entity `json:"additionalAllowedEntities,omitempty"`
	AllowAdditionalEntities   *bool      `json:"allowAdditionalEntities,omitempty"`
	AllowAllEntities          *bool      `json:"allowAllEntities,omitempty"`
}

//AdministrativeAccount - A PingFederate administrator account.
type AdministrativeAccount struct {
	Active            *bool      `json:"active,omitempty"`
	Auditor           *bool      `json:"auditor,omitempty"`
	Department        *string    `json:"department,omitempty"`
	Description       *string    `json:"description,omitempty"`
	EmailAddress      *string    `json:"emailAddress,omitempty"`
	EncryptedPassword *string    `json:"encryptedPassword,omitempty"`
	Password          *string    `json:"password,omitempty"`
	PhoneNumber       *string    `json:"phoneNumber,omitempty"`
	Roles             *[]*string `json:"roles,omitempty"`
	Username          *string    `json:"username,omitempty"`
}

//AdministrativeAccounts - PingFederate administrator accounts.
type AdministrativeAccounts struct {
	Items *[]*AdministrativeAccount `json:"items,omitempty"`
}

//AlternativeLoginHintTokenIssuer - JSON Web Key Set Settings.
type AlternativeLoginHintTokenIssuer struct {
	Issuer  *string `json:"issuer,omitempty"`
	Jwks    *string `json:"jwks,omitempty"`
	JwksURL *string `json:"jwksURL,omitempty"`
}

//ApcMappingPolicyAction - An authentication policy contract selection action.
type ApcMappingPolicyAction struct {
	AttributeMapping                *AttributeMapping `json:"attributeMapping,omitempty"`
	AuthenticationPolicyContractRef *ResourceLink     `json:"authenticationPolicyContractRef,omitempty"`
	Context                         *string           `json:"context,omitempty"`
	Type                            *string           `json:"type,omitempty"`
}

//ApcToPersistentGrantMapping - An authentication policy contract mapping into an OAuth persistent grant.
type ApcToPersistentGrantMapping struct {
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	AuthenticationPolicyContractRef *ResourceLink                         `json:"authenticationPolicyContractRef,omitempty"`
	Id                              *string                               `json:"id,omitempty"`
	IssuanceCriteria                *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//ApcToPersistentGrantMappings - A collection of OAuth authentication policy contract to persistent grant mapping items.
type ApcToPersistentGrantMappings struct {
	Items *[]*ApcToPersistentGrantMapping `json:"items,omitempty"`
}

//ApcToSpAdapterMapping - The Authentication Policy Contract (APC) to SP Adapter Mapping.
type ApcToSpAdapterMapping struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                 *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	DefaultTargetResource            *string                               `json:"defaultTargetResource,omitempty"`
	Id                               *string                               `json:"id,omitempty"`
	IssuanceCriteria                 *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	LicenseConnectionGroupAssignment *string                               `json:"licenseConnectionGroupAssignment,omitempty"`
	SourceId                         *string                               `json:"sourceId,omitempty"`
	TargetId                         *string                               `json:"targetId,omitempty"`
}

//ApcToSpAdapterMappings - A collection of Authentication Policy Contract (APC) to SP Adapter Mappings.
type ApcToSpAdapterMappings struct {
	Items *[]*ApcToSpAdapterMapping `json:"items,omitempty"`
}

//ApiResult - Details on the result of the operation.
type ApiResult struct {
	DeveloperMessage *string             `json:"developerMessage,omitempty"`
	Message          *string             `json:"message,omitempty"`
	ResultId         *string             `json:"resultId,omitempty"`
	ValidationErrors *[]*ValidationError `json:"validationErrors,omitempty"`
}

//ApplicationSessionPolicy - Session controls for user facing PingFederate application endpoints, such as the profile management endpoint.
type ApplicationSessionPolicy struct {
	IdleTimeoutMins *int `json:"idleTimeoutMins,omitempty"`
	MaxTimeoutMins  *int `json:"maxTimeoutMins,omitempty"`
}

//ArtifactResolverLocation - The remote party URLs to resolve the artifact.
type ArtifactResolverLocation struct {
	Index *int    `json:"index,omitempty"`
	Url   *string `json:"url,omitempty"`
}

//ArtifactSettings - The settings for an Artifact binding.
type ArtifactSettings struct {
	Lifetime          *int                         `json:"lifetime,omitempty"`
	ResolverLocations *[]*ArtifactResolverLocation `json:"resolverLocations,omitempty"`
	SourceId          *string                      `json:"sourceId,omitempty"`
}

//AssertionLifetime - The timeframe of validity before and after the issuance of the assertion.
type AssertionLifetime struct {
	MinutesAfter  *int `json:"minutesAfter,omitempty"`
	MinutesBefore *int `json:"minutesBefore,omitempty"`
}

//AtmAccessControlSettings - Access control settings for an access token management plugin instance.
type AtmAccessControlSettings struct {
	AllowedClients  *[]*ResourceLink `json:"allowedClients,omitempty"`
	Inherited       *bool            `json:"inherited,omitempty"`
	RestrictClients *bool            `json:"restrictClients,omitempty"`
}

//AtmSelectionSettings - Selection settings for an access token management plugin instance.
type AtmSelectionSettings struct {
	Inherited    *bool      `json:"inherited,omitempty"`
	ResourceUris *[]*string `json:"resourceUris,omitempty"`
}

//AttributeFulfillmentValue - Defines how an attribute in an attribute contract should be populated.
type AttributeFulfillmentValue struct {
	Source *SourceTypeIdKey `json:"source,omitempty"`
	Value  *string          `json:"value,omitempty"`
}

//AttributeMapping - A list of mappings from attribute sources to attribute targets.
type AttributeMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//AttributeQueryNameMapping - The attribute query name mappings between the SP and the IdP.
type AttributeQueryNameMapping struct {
	LocalName  *string `json:"localName,omitempty"`
	RemoteName *string `json:"remoteName,omitempty"`
}

//AttributeRule - Authentication policy rules using attributes from the previous authentication source. Each rule is evaluated to determine the next action in the policy.
type AttributeRule struct {
	AttributeName *string `json:"attributeName,omitempty"`
	Condition     *string `json:"condition,omitempty"`
	ExpectedValue *string `json:"expectedValue,omitempty"`
	Result        *string `json:"result,omitempty"`
}

//AttributeRules - A collection of attribute rules
type AttributeRules struct {
	FallbackToSuccess *bool             `json:"fallbackToSuccess,omitempty"`
	Items             *[]*AttributeRule `json:"items,omitempty"`
}

//AttributeSource - The configured settings to look up attributes from an associated data store.
type AttributeSource struct {
	LdapAttributeSource
	CustomAttributeSource
	JdbcAttributeSource
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	DataStoreRef                 *ResourceLink                         `json:"dataStoreRef,omitempty"`
	Description                  *string                               `json:"description,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	Type                         *string                               `json:"type,omitempty"`
}

//AuthenticationPoliciesSettings - The settings for the authentication policies.
type AuthenticationPoliciesSettings struct {
	EnableIdpAuthnSelection *bool `json:"enableIdpAuthnSelection,omitempty"`
	EnableSpAuthnSelection  *bool `json:"enableSpAuthnSelection,omitempty"`
}

//AuthenticationPolicy - An authentication policy.
type AuthenticationPolicy struct {
	AuthnSelectionTrees          *[]*AuthenticationPolicyTree `json:"authnSelectionTrees,omitempty"`
	DefaultAuthenticationSources *[]*AuthenticationSource     `json:"defaultAuthenticationSources,omitempty"`
	FailIfNoSelection            *bool                        `json:"failIfNoSelection,omitempty"`
	TrackedHttpParameters        *[]*string                   `json:"trackedHttpParameters,omitempty"`
}

//AuthenticationPolicyContract - Authentication Policy Contracts carry user attributes from the identity provider to the service provider.
type AuthenticationPolicyContract struct {
	CoreAttributes     *[]*AuthenticationPolicyContractAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*AuthenticationPolicyContractAttribute `json:"extendedAttributes,omitempty"`
	Id                 *string                                   `json:"id,omitempty"`
	Name               *string                                   `json:"name,omitempty"`
}

//AuthenticationPolicyContractAssertionMapping - The Authentication Policy Contract Assertion Mapping.
type AuthenticationPolicyContractAssertionMapping struct {
	AbortSsoTransactionAsFailSafe   *bool                                 `json:"abortSsoTransactionAsFailSafe,omitempty"`
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	AuthenticationPolicyContractRef *ResourceLink                         `json:"authenticationPolicyContractRef,omitempty"`
	IssuanceCriteria                *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictVirtualEntityIds        *bool                                 `json:"restrictVirtualEntityIds,omitempty"`
	RestrictedVirtualEntityIds      *[]*string                            `json:"restrictedVirtualEntityIds,omitempty"`
}

//AuthenticationPolicyContractAttribute - An attribute for the Authentication Policy Contract.
type AuthenticationPolicyContractAttribute struct {
	Name *string `json:"name,omitempty"`
}

//AuthenticationPolicyContractMapping - An Authentication Policy Contract mapping into IdP Connection.
type AuthenticationPolicyContractMapping struct {
	AttributeContractFulfillment    map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	AuthenticationPolicyContractRef *ResourceLink                         `json:"authenticationPolicyContractRef,omitempty"`
	IssuanceCriteria                *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictVirtualServerIds        *bool                                 `json:"restrictVirtualServerIds,omitempty"`
	RestrictedVirtualServerIds      *[]*string                            `json:"restrictedVirtualServerIds,omitempty"`
}

//AuthenticationPolicyContracts - A collection of Authentication Policy Contracts
type AuthenticationPolicyContracts struct {
	Items *[]*AuthenticationPolicyContract `json:"items,omitempty"`
}

//AuthenticationPolicyFragment - An authentication policy fragment.
type AuthenticationPolicyFragment struct {
	Description *string                       `json:"description,omitempty"`
	Id          *string                       `json:"id,omitempty"`
	Inputs      *ResourceLink                 `json:"inputs,omitempty"`
	Name        *string                       `json:"name,omitempty"`
	Outputs     *ResourceLink                 `json:"outputs,omitempty"`
	RootNode    *AuthenticationPolicyTreeNode `json:"rootNode,omitempty"`
}

//AuthenticationPolicyFragments - A collection of Authentication Policy Fragments
type AuthenticationPolicyFragments struct {
	Items *[]*AuthenticationPolicyFragment `json:"items,omitempty"`
}

//AuthenticationPolicyTree - An authentication policy tree.
type AuthenticationPolicyTree struct {
	AuthenticationApiApplicationRef *ResourceLink                 `json:"authenticationApiApplicationRef,omitempty"`
	Description                     *string                       `json:"description,omitempty"`
	Enabled                         *bool                         `json:"enabled,omitempty"`
	Name                            *string                       `json:"name,omitempty"`
	RootNode                        *AuthenticationPolicyTreeNode `json:"rootNode,omitempty"`
}

//AuthenticationPolicyTreeNode - An authentication policy tree node.
type AuthenticationPolicyTreeNode struct {
	Action   *PolicyAction                    `json:"action,omitempty"`
	Children *[]*AuthenticationPolicyTreeNode `json:"children,omitempty"`
}

//AuthenticationSelector - An Authentication Selector instance.
type AuthenticationSelector struct {
	AttributeContract   *AuthenticationSelectorAttributeContract `json:"attributeContract,omitempty"`
	Configuration       *PluginConfiguration                     `json:"configuration,omitempty"`
	Id                  *string                                  `json:"id,omitempty"`
	Name                *string                                  `json:"name,omitempty"`
	PluginDescriptorRef *ResourceLink                            `json:"pluginDescriptorRef,omitempty"`
}

//AuthenticationSelectorAttribute - An attribute for the Authentication Selector attribute contract.
type AuthenticationSelectorAttribute struct {
	Name *string `json:"name,omitempty"`
}

//AuthenticationSelectorAttributeContract - A set of attributes exposed by an Authentication Selector.
type AuthenticationSelectorAttributeContract struct {
	ExtendedAttributes *[]*AuthenticationSelectorAttribute `json:"extendedAttributes,omitempty"`
}

//AuthenticationSelectorDescriptor - An Authentication Selector descriptor.
type AuthenticationSelectorDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//AuthenticationSelectorDescriptors - A collection of Authentication Selector descriptors.
type AuthenticationSelectorDescriptors struct {
	Items *[]*AuthenticationSelectorDescriptor `json:"items,omitempty"`
}

//AuthenticationSelectors - A collection of Authentication Selector instances.
type AuthenticationSelectors struct {
	Items *[]*AuthenticationSelector `json:"items,omitempty"`
}

//AuthenticationSessionPolicies - A collection of authentication session policies.
type AuthenticationSessionPolicies struct {
	Items *[]*AuthenticationSessionPolicy `json:"items,omitempty"`
}

//AuthenticationSessionPolicy - The session policy for a specified authentication source.
type AuthenticationSessionPolicy struct {
	AuthenticationSource  *AuthenticationSource `json:"authenticationSource,omitempty"`
	AuthnContextSensitive *bool                 `json:"authnContextSensitive,omitempty"`
	EnableSessions        *bool                 `json:"enableSessions,omitempty"`
	Id                    *string               `json:"id,omitempty"`
	IdleTimeoutMins       *int                  `json:"idleTimeoutMins,omitempty"`
	MaxTimeoutMins        *int                  `json:"maxTimeoutMins,omitempty"`
	Persistent            *bool                 `json:"persistent,omitempty"`
	TimeoutDisplayUnit    *string               `json:"timeoutDisplayUnit,omitempty"`
}

//AuthenticationSource - An authentication source (IdP adapter or IdP connection).
type AuthenticationSource struct {
	SourceRef *ResourceLink `json:"sourceRef,omitempty"`
	Type      *string       `json:"type,omitempty"`
}

//AuthnApiApplication - Authentication API Application.
type AuthnApiApplication struct {
	AdditionalAllowedOrigins *[]*string `json:"additionalAllowedOrigins,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Id                       *string    `json:"id,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	Url                      *string    `json:"url,omitempty"`
}

//AuthnApiApplications - A collection of Authentication API Application items.
type AuthnApiApplications struct {
	Items *[]*AuthnApiApplication `json:"items,omitempty"`
}

//AuthnApiSettings - Authentication API Application Settings.
type AuthnApiSettings struct {
	ApiEnabled            *bool         `json:"apiEnabled,omitempty"`
	DefaultApplicationRef *ResourceLink `json:"defaultApplicationRef,omitempty"`
	EnableApiDescriptions *bool         `json:"enableApiDescriptions,omitempty"`
}

//AuthnContextMapping - The authentication context mapping between local and remote values.
type AuthnContextMapping struct {
	Local  *string `json:"local,omitempty"`
	Remote *string `json:"remote,omitempty"`
}

//AuthnSelectorPolicyAction - An authentication selector selection action.
type AuthnSelectorPolicyAction struct {
	AuthenticationSelectorRef *ResourceLink `json:"authenticationSelectorRef,omitempty"`
	Context                   *string       `json:"context,omitempty"`
	Type                      *string       `json:"type,omitempty"`
}

//AuthnSourcePolicyAction - An authentication source selection action.
type AuthnSourcePolicyAction struct {
	AttributeRules       *AttributeRules            `json:"attributeRules,omitempty"`
	AuthenticationSource *AuthenticationSource      `json:"authenticationSource,omitempty"`
	Context              *string                    `json:"context,omitempty"`
	InputUserIdMapping   *AttributeFulfillmentValue `json:"inputUserIdMapping,omitempty"`
	Type                 *string                    `json:"type,omitempty"`
	UserIdAuthenticated  *bool                      `json:"userIdAuthenticated,omitempty"`
}

//AuthorizationServerSettings - Authorization Server Settings attributes.
type AuthorizationServerSettings struct {
	AdminWebServicePcvRef                  *ResourceLink            `json:"adminWebServicePcvRef,omitempty"`
	AllowUnidentifiedClientExtensionGrants *bool                    `json:"allowUnidentifiedClientExtensionGrants,omitempty"`
	AllowUnidentifiedClientROCreds         *bool                    `json:"allowUnidentifiedClientROCreds,omitempty"`
	AllowedOrigins                         *[]*string               `json:"allowedOrigins,omitempty"`
	ApprovedScopesAttribute                *string                  `json:"approvedScopesAttribute,omitempty"`
	AtmIdForOAuthGrantManagement           *string                  `json:"atmIdForOAuthGrantManagement,omitempty"`
	AuthorizationCodeEntropy               *int                     `json:"authorizationCodeEntropy,omitempty"`
	AuthorizationCodeTimeout               *int                     `json:"authorizationCodeTimeout,omitempty"`
	BypassActivationCodeConfirmation       *bool                    `json:"bypassActivationCodeConfirmation,omitempty"`
	BypassAuthorizationForApprovedGrants   *bool                    `json:"bypassAuthorizationForApprovedGrants,omitempty"`
	DefaultScopeDescription                *string                  `json:"defaultScopeDescription,omitempty"`
	DevicePollingInterval                  *int                     `json:"devicePollingInterval,omitempty"`
	ExclusiveScopeGroups                   *[]*ScopeGroupEntry      `json:"exclusiveScopeGroups,omitempty"`
	ExclusiveScopes                        *[]*ScopeEntry           `json:"exclusiveScopes,omitempty"`
	ParReferenceLength                     *int                     `json:"parReferenceLength,omitempty"`
	ParReferenceTimeout                    *int                     `json:"parReferenceTimeout,omitempty"`
	ParStatus                              *string                  `json:"parStatus,omitempty"`
	PendingAuthorizationTimeout            *int                     `json:"pendingAuthorizationTimeout,omitempty"`
	PersistentGrantContract                *PersistentGrantContract `json:"persistentGrantContract,omitempty"`
	PersistentGrantIdleTimeout             *int                     `json:"persistentGrantIdleTimeout,omitempty"`
	PersistentGrantIdleTimeoutTimeUnit     *string                  `json:"persistentGrantIdleTimeoutTimeUnit,omitempty"`
	PersistentGrantLifetime                *int                     `json:"persistentGrantLifetime,omitempty"`
	PersistentGrantLifetimeUnit            *string                  `json:"persistentGrantLifetimeUnit,omitempty"`
	PersistentGrantReuseGrantTypes         []*string                `json:"persistentGrantReuseGrantTypes,omitempty"`
	RefreshRollingInterval                 *int                     `json:"refreshRollingInterval,omitempty"`
	RefreshTokenLength                     *int                     `json:"refreshTokenLength,omitempty"`
	RegisteredAuthorizationPath            *string                  `json:"registeredAuthorizationPath,omitempty"`
	RollRefreshTokenValues                 *bool                    `json:"rollRefreshTokenValues,omitempty"`
	ScopeForOAuthGrantManagement           *string                  `json:"scopeForOAuthGrantManagement,omitempty"`
	ScopeGroups                            *[]*ScopeGroupEntry      `json:"scopeGroups,omitempty"`
	Scopes                                 *[]*ScopeEntry           `json:"scopes,omitempty"`
	TokenEndpointBaseUrl                   *string                  `json:"tokenEndpointBaseUrl,omitempty"`
	TrackUserSessionsForLogout             *bool                    `json:"trackUserSessionsForLogout,omitempty"`
	UserAuthorizationConsentAdapter        *string                  `json:"userAuthorizationConsentAdapter,omitempty"`
	UserAuthorizationConsentPageSetting    *string                  `json:"userAuthorizationConsentPageSetting,omitempty"`
	UserAuthorizationUrl                   *string                  `json:"userAuthorizationUrl,omitempty"`
}

//BackChannelAuth - The SOAP authentication methods when sending or receiving a message using SOAP back channel.
type BackChannelAuth struct {
	OutboundBackChannelAuth
	InboundBackChannelAuth
	DigitalSignature     *bool                        `json:"digitalSignature,omitempty"`
	HttpBasicCredentials *UsernamePasswordCredentials `json:"httpBasicCredentials,omitempty"`
	Type                 *string                      `json:"type,omitempty"`
}

//BaseDefaultValueLocalIdentityField - Holds fields that are shared by all default value type fields.
type BaseDefaultValueLocalIdentityField struct {
	CheckboxLocalIdentityField
	DateLocalIdentityField
	TextLocalIdentityField
	DropDownLocalIdentityField
	HiddenLocalIdentityField
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	DefaultValue          *string          `json:"defaultValue,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//BaseProviderRole - Base Provider Role.
type BaseProviderRole struct {
	IdpRole
	SpRole
	Enable        *bool `json:"enable,omitempty"`
	EnableSaml10  *bool `json:"enableSaml10,omitempty"`
	EnableSaml11  *bool `json:"enableSaml11,omitempty"`
	EnableWsFed   *bool `json:"enableWsFed,omitempty"`
	EnableWsTrust *bool `json:"enableWsTrust,omitempty"`
}

//BaseSelectionFieldDescriptor - Holds fields that are shared by all selection-type field descriptors.
type BaseSelectionFieldDescriptor struct {
	RadioGroupFieldDescriptor
	SelectFieldDescriptor
	Advanced     *bool           `json:"advanced,omitempty"`
	DefaultValue *string         `json:"defaultValue,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Label        *string         `json:"label,omitempty"`
	Name         *string         `json:"name,omitempty"`
	OptionValues *[]*OptionValue `json:"optionValues,omitempty"`
	Required     *bool           `json:"required,omitempty"`
	Type         *string         `json:"type,omitempty"`
}

//BaseSelectionLocalIdentityField - Holds fields that are shared by all selection-type fields.
type BaseSelectionLocalIdentityField struct {
	CheckboxGroupLocalIdentityField
	DropDownLocalIdentityField
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	Options               *[]*string       `json:"options,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//BinaryLdapAttributeSettings - Binary settings for a LDAP attribute.
type BinaryLdapAttributeSettings struct {
	BinaryEncoding *string `json:"binaryEncoding,omitempty"`
}

//BulkConfig - Model describing a series of configuration operations.
type BulkConfig struct {
	Metadata   *BulkConfigMetadata `json:"metadata,omitempty"`
	Operations *[]*ConfigOperation `json:"operations,omitempty"`
}

//BulkConfigMetadata - Model describing how bulk configuration data was generated.
type BulkConfigMetadata struct {
	PfVersion *string `json:"pfVersion,omitempty"`
}

//CSRResponse - Represents a CSR response file.
type CSRResponse struct {
	FileData *string `json:"fileData,omitempty"`
}

//CaptchaSettings - Settings for CAPTCHA.
type CaptchaSettings struct {
	EncryptedSecretKey *string `json:"encryptedSecretKey,omitempty"`
	SecretKey          *string `json:"secretKey,omitempty"`
	SiteKey            *string `json:"siteKey,omitempty"`
}

//CertView - Certificate details.
type CertView struct {
	CryptoProvider          *string    `json:"cryptoProvider,omitempty"`
	Expires                 *string    `json:"expires,omitempty"`
	Id                      *string    `json:"id,omitempty"`
	IssuerDN                *string    `json:"issuerDN,omitempty"`
	KeyAlgorithm            *string    `json:"keyAlgorithm,omitempty"`
	KeySize                 *int       `json:"keySize,omitempty"`
	SerialNumber            *string    `json:"serialNumber,omitempty"`
	Sha1Fingerprint         *string    `json:"sha1Fingerprint,omitempty"`
	Sha256Fingerprint       *string    `json:"sha256Fingerprint,omitempty"`
	SignatureAlgorithm      *string    `json:"signatureAlgorithm,omitempty"`
	Status                  *string    `json:"status,omitempty"`
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames,omitempty"`
	SubjectDN               *string    `json:"subjectDN,omitempty"`
	ValidFrom               *string    `json:"validFrom,omitempty"`
	Version                 *int       `json:"version,omitempty"`
}

//CertViews - A collection of Certificate items.
type CertViews struct {
	Items *[]*CertView `json:"items,omitempty"`
}

//CertificateExpirationNotificationSettings - Notification settings for certificate expiration events.
type CertificateExpirationNotificationSettings struct {
	EmailAddress             *string       `json:"emailAddress,omitempty"`
	FinalWarningPeriod       *int          `json:"finalWarningPeriod,omitempty"`
	InitialWarningPeriod     *int          `json:"initialWarningPeriod,omitempty"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty"`
}

//CertificateRevocationSettings - Certificate revocation settings.
type CertificateRevocationSettings struct {
	CrlSettings   *CrlSettings   `json:"crlSettings,omitempty"`
	OcspSettings  *OcspSettings  `json:"ocspSettings,omitempty"`
	ProxySettings *ProxySettings `json:"proxySettings,omitempty"`
}

//ChangeDetectionSettings - Setting to detect changes to a user or a group.
type ChangeDetectionSettings struct {
	ChangedUsersAlgorithm  *string `json:"changedUsersAlgorithm,omitempty"`
	GroupObjectClass       *string `json:"groupObjectClass,omitempty"`
	TimeStampAttributeName *string `json:"timeStampAttributeName,omitempty"`
	UserObjectClass        *string `json:"userObjectClass,omitempty"`
	UsnAttributeName       *string `json:"usnAttributeName,omitempty"`
}

//Channel - A channel is a combination of a source data store and a provisioning target. It include settings of a source data store, managing provisioning threads and mapping of attributes.
type Channel struct {
	Active           *bool                    `json:"active,omitempty"`
	AttributeMapping *[]*SaasAttributeMapping `json:"attributeMapping,omitempty"`
	ChannelSource    *ChannelSource           `json:"channelSource,omitempty"`
	MaxThreads       *int                     `json:"maxThreads,omitempty"`
	Name             *string                  `json:"name,omitempty"`
	Timeout          *int                     `json:"timeout,omitempty"`
}

//ChannelSource - The source data source and LDAP settings.
type ChannelSource struct {
	AccountManagementSettings *AccountManagementSettings `json:"accountManagementSettings,omitempty"`
	BaseDn                    *string                    `json:"baseDn,omitempty"`
	ChangeDetectionSettings   *ChangeDetectionSettings   `json:"changeDetectionSettings,omitempty"`
	DataSource                *ResourceLink              `json:"dataSource,omitempty"`
	GroupMembershipDetection  *GroupMembershipDetection  `json:"groupMembershipDetection,omitempty"`
	GroupSourceLocation       *ChannelSourceLocation     `json:"groupSourceLocation,omitempty"`
	GuidAttributeName         *string                    `json:"guidAttributeName,omitempty"`
	GuidBinary                *bool                      `json:"guidBinary,omitempty"`
	UserSourceLocation        *ChannelSourceLocation     `json:"userSourceLocation,omitempty"`
}

//ChannelSourceLocation - The location settings that includes a DN and a LDAP filter.
type ChannelSourceLocation struct {
	Filter       *string `json:"filter,omitempty"`
	GroupDN      *string `json:"groupDN,omitempty"`
	NestedSearch *bool   `json:"nestedSearch,omitempty"`
}

//CheckBoxFieldDescriptor - A boolean field typically rendered as a checkbox in a configuration UI.
type CheckBoxFieldDescriptor struct {
	Advanced     *bool   `json:"advanced,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//CheckboxGroupLocalIdentityField - A checkbox group selection type field.
type CheckboxGroupLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	Options               *[]*string       `json:"options,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//CheckboxLocalIdentityField - A checkbox selection type field.
type CheckboxLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	DefaultValue          *string          `json:"defaultValue,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//CibaServerPolicySettings - Settings for the CIBA request policy configuration.
type CibaServerPolicySettings struct {
	DefaultRequestPolicyRef *ResourceLink `json:"defaultRequestPolicyRef,omitempty"`
}

//Client - OAuth client.
type Client struct {
	AllowAuthenticationApiInit               *bool                       `json:"allowAuthenticationApiInit,omitempty"`
	BypassActivationCodeConfirmationOverride *bool                       `json:"bypassActivationCodeConfirmationOverride,omitempty"`
	BypassApprovalPage                       *bool                       `json:"bypassApprovalPage,omitempty"`
	CibaDeliveryMode                         *string                     `json:"cibaDeliveryMode,omitempty"`
	CibaNotificationEndpoint                 *string                     `json:"cibaNotificationEndpoint,omitempty"`
	CibaPollingInterval                      *int                        `json:"cibaPollingInterval,omitempty"`
	CibaRequestObjectSigningAlgorithm        *string                     `json:"cibaRequestObjectSigningAlgorithm,omitempty"`
	CibaRequireSignedRequests                *bool                       `json:"cibaRequireSignedRequests,omitempty"`
	CibaUserCodeSupported                    *bool                       `json:"cibaUserCodeSupported,omitempty"`
	ClientAuth                               *ClientAuth                 `json:"clientAuth,omitempty"`
	ClientId                                 *string                     `json:"clientId,omitempty"`
	DefaultAccessTokenManagerRef             *ResourceLink               `json:"defaultAccessTokenManagerRef,omitempty"`
	Description                              *string                     `json:"description,omitempty"`
	DeviceFlowSettingType                    *string                     `json:"deviceFlowSettingType,omitempty"`
	DevicePollingIntervalOverride            *int                        `json:"devicePollingIntervalOverride,omitempty"`
	Enabled                                  *bool                       `json:"enabled,omitempty"`
	ExclusiveScopes                          *[]*string                  `json:"exclusiveScopes,omitempty"`
	ExtendedParameters                       map[string]*ParameterValues `json:"extendedParameters,omitempty"`
	GrantTypes                               *[]*string                  `json:"grantTypes,omitempty"`
	JwksSettings                             *JwksSettings               `json:"jwksSettings,omitempty"`
	LogoUrl                                  *string                     `json:"logoUrl,omitempty"`
	Name                                     *string                     `json:"name,omitempty"`
	OidcPolicy                               *ClientOIDCPolicy           `json:"oidcPolicy,omitempty"`
	PendingAuthorizationTimeoutOverride      *int                        `json:"pendingAuthorizationTimeoutOverride,omitempty"`
	PersistentGrantExpirationTime            *int                        `json:"persistentGrantExpirationTime,omitempty"`
	PersistentGrantExpirationTimeUnit        *string                     `json:"persistentGrantExpirationTimeUnit,omitempty"`
	PersistentGrantExpirationType            *string                     `json:"persistentGrantExpirationType,omitempty"`
	PersistentGrantIdleTimeout               *int                        `json:"persistentGrantIdleTimeout,omitempty"`
	PersistentGrantIdleTimeoutTimeUnit       *string                     `json:"persistentGrantIdleTimeoutTimeUnit,omitempty"`
	PersistentGrantIdleTimeoutType           *string                     `json:"persistentGrantIdleTimeoutType,omitempty"`
	RedirectUris                             *[]*string                  `json:"redirectUris,omitempty"`
	RefreshRolling                           *string                     `json:"refreshRolling,omitempty"`
	RequestObjectSigningAlgorithm            *string                     `json:"requestObjectSigningAlgorithm,omitempty"`
	RequestPolicyRef                         *ResourceLink               `json:"requestPolicyRef,omitempty"`
	RequireProofKeyForCodeExchange           *bool                       `json:"requireProofKeyForCodeExchange,omitempty"`
	RequirePushedAuthorizationRequests       *bool                       `json:"requirePushedAuthorizationRequests,omitempty"`
	RequireSignedRequests                    *bool                       `json:"requireSignedRequests,omitempty"`
	RestrictScopes                           *bool                       `json:"restrictScopes,omitempty"`
	RestrictToDefaultAccessTokenManager      *bool                       `json:"restrictToDefaultAccessTokenManager,omitempty"`
	RestrictedResponseTypes                  *[]*string                  `json:"restrictedResponseTypes,omitempty"`
	RestrictedScopes                         *[]*string                  `json:"restrictedScopes,omitempty"`
	TokenExchangeProcessorPolicyRef          *ResourceLink               `json:"tokenExchangeProcessorPolicyRef,omitempty"`
	UserAuthorizationUrlOverride             *string                     `json:"userAuthorizationUrlOverride,omitempty"`
	ValidateUsingAllEligibleAtms             *bool                       `json:"validateUsingAllEligibleAtms,omitempty"`
}

//ClientAuth - Client Authentication.
type ClientAuth struct {
	ClientCertIssuerDn                *string `json:"clientCertIssuerDn,omitempty"`
	ClientCertSubjectDn               *string `json:"clientCertSubjectDn,omitempty"`
	EncryptedSecret                   *string `json:"encryptedSecret,omitempty"`
	EnforceReplayPrevention           *bool   `json:"enforceReplayPrevention,omitempty"`
	Secret                            *string `json:"secret,omitempty"`
	TokenEndpointAuthSigningAlgorithm *string `json:"tokenEndpointAuthSigningAlgorithm,omitempty"`
	Type                              *string `json:"type,omitempty"`
}

//ClientMetadata - The client metadata.
type ClientMetadata struct {
	Description *string `json:"description,omitempty"`
	MultiValued *bool   `json:"multiValued,omitempty"`
	Parameter   *string `json:"parameter,omitempty"`
}

//ClientOIDCPolicy - OAuth Client Open ID Connect Policy.
type ClientOIDCPolicy struct {
	GrantAccessSessionRevocationApi        *bool         `json:"grantAccessSessionRevocationApi,omitempty"`
	GrantAccessSessionSessionManagementApi *bool         `json:"grantAccessSessionSessionManagementApi,omitempty"`
	IdTokenContentEncryptionAlgorithm      *string       `json:"idTokenContentEncryptionAlgorithm,omitempty"`
	IdTokenEncryptionAlgorithm             *string       `json:"idTokenEncryptionAlgorithm,omitempty"`
	IdTokenSigningAlgorithm                *string       `json:"idTokenSigningAlgorithm,omitempty"`
	LogoutUris                             *[]*string    `json:"logoutUris,omitempty"`
	PairwiseIdentifierUserType             *bool         `json:"pairwiseIdentifierUserType,omitempty"`
	PingAccessLogoutCapable                *bool         `json:"pingAccessLogoutCapable,omitempty"`
	PolicyGroup                            *ResourceLink `json:"policyGroup,omitempty"`
	SectorIdentifierUri                    *string       `json:"sectorIdentifierUri,omitempty"`
}

//ClientRegistrationOIDCPolicy - Client Registration Open ID Connect Policy settings.
type ClientRegistrationOIDCPolicy struct {
	IdTokenContentEncryptionAlgorithm *string       `json:"idTokenContentEncryptionAlgorithm,omitempty"`
	IdTokenEncryptionAlgorithm        *string       `json:"idTokenEncryptionAlgorithm,omitempty"`
	IdTokenSigningAlgorithm           *string       `json:"idTokenSigningAlgorithm,omitempty"`
	PolicyGroup                       *ResourceLink `json:"policyGroup,omitempty"`
}

//ClientRegistrationPolicies - A collection of client registration policy plugin instances.
type ClientRegistrationPolicies struct {
	Items *[]*ClientRegistrationPolicy `json:"items,omitempty"`
}

//ClientRegistrationPolicy - A client registration policy plugin instance.
type ClientRegistrationPolicy struct {
	Configuration       *PluginConfiguration `json:"configuration,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	ParentRef           *ResourceLink        `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink        `json:"pluginDescriptorRef,omitempty"`
}

//ClientRegistrationPolicyDescriptor - A client registration policy plugin descriptor.
type ClientRegistrationPolicyDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//ClientRegistrationPolicyDescriptors - A collection of client registration policy plugin descriptors.
type ClientRegistrationPolicyDescriptors struct {
	Items *[]*ClientRegistrationPolicyDescriptor `json:"items,omitempty"`
}

//ClientSecret - Client Secret.
type ClientSecret struct {
	EncryptedSecret *string `json:"encryptedSecret,omitempty"`
	Secret          *string `json:"secret,omitempty"`
}

//ClientSettings - The client settings.
type ClientSettings struct {
	ClientMetadata            *[]*ClientMetadata         `json:"clientMetadata,omitempty"`
	DynamicClientRegistration *DynamicClientRegistration `json:"dynamicClientRegistration,omitempty"`
}

//Clients - A collection of OAuth client items.
type Clients struct {
	Items *[]*Client `json:"items,omitempty"`
}

//ClusterNode - Describes a node in a clustered deployment of PingFederate.
type ClusterNode struct {
	Address   *string `json:"address,omitempty"`
	Index     *int    `json:"index,omitempty"`
	Mode      *string `json:"mode,omitempty"`
	NodeGroup *string `json:"nodeGroup,omitempty"`
	NodeTags  *string `json:"nodeTags,omitempty"`
	Version   *string `json:"version,omitempty"`
}

//ClusterStatus - Information on cluster nodes and replication status.
type ClusterStatus struct {
	LastConfigUpdateTime *string         `json:"lastConfigUpdateTime,omitempty"`
	LastReplicationTime  *string         `json:"lastReplicationTime,omitempty"`
	MixedMode            *bool           `json:"mixedMode,omitempty"`
	Nodes                *[]*ClusterNode `json:"nodes,omitempty"`
	ReplicationRequired  *bool           `json:"replicationRequired,omitempty"`
}

//ConditionalIssuanceCriteriaEntry - An issuance criterion that checks a source attribute against a particular condition and the expected value. If the condition is true then this issuance criterion passes, otherwise the criterion fails.
type ConditionalIssuanceCriteriaEntry struct {
	AttributeName *string          `json:"attributeName,omitempty"`
	Condition     *string          `json:"condition,omitempty"`
	ErrorResult   *string          `json:"errorResult,omitempty"`
	Source        *SourceTypeIdKey `json:"source,omitempty"`
	Value         *string          `json:"value,omitempty"`
}

//ConfigField - A plugin configuration field value.
type ConfigField struct {
	EncryptedValue *string `json:"encryptedValue,omitempty"`
	Inherited      *bool   `json:"inherited,omitempty"`
	Name           *string `json:"name,omitempty"`
	Value          *string `json:"value,omitempty"`
}

//ConfigOperation - Model describing a list of configuration operations for a given resource type.
type ConfigOperation struct {
	ItemIds       *[]*string     `json:"itemIds,omitempty"`
	Items         *[]interface{} `json:"items,omitempty"`
	OperationType *string        `json:"operationType,omitempty"`
	ResourceType  *string        `json:"resourceType,omitempty"`
	SubResource   *string        `json:"subResource,omitempty"`
}

//ConfigRow - A row of configuration values for a plugin configuration table.
type ConfigRow struct {
	DefaultRow *bool           `json:"defaultRow,omitempty"`
	Fields     *[]*ConfigField `json:"fields,omitempty"`
}

//ConfigStoreBundle - List of all configuration settings in a bundle.
type ConfigStoreBundle struct {
	Items *[]*ConfigStoreSetting `json:"items,omitempty"`
}

//ConfigStoreSetting - Single configuration setting.
type ConfigStoreSetting struct {
	Id          *string            `json:"id,omitempty"`
	ListValue   *[]*string         `json:"listValue,omitempty"`
	MapValue    map[string]*string `json:"mapValue,omitempty"`
	StringValue *string            `json:"stringValue,omitempty"`
	Type        *string            `json:"type,omitempty"`
}

//ConfigTable - A plugin configuration table populated with values.
type ConfigTable struct {
	Inherited *bool         `json:"inherited,omitempty"`
	Name      *string       `json:"name,omitempty"`
	Rows      *[]*ConfigRow `json:"rows,omitempty"`
}

//Connection - Settings shared by SP-side and IdP-side connections.
type Connection struct {
	IdpConnection
	SpConnection
	Active                                 *bool                                   `json:"active,omitempty"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfiguration `json:"additionalAllowedEntitiesConfiguration,omitempty"`
	BaseUrl                                *string                                 `json:"baseUrl,omitempty"`
	ContactInfo                            *ContactInfo                            `json:"contactInfo,omitempty"`
	Credentials                            *ConnectionCredentials                  `json:"credentials,omitempty"`
	DefaultVirtualEntityId                 *string                                 `json:"defaultVirtualEntityId,omitempty"`
	EntityId                               *string                                 `json:"entityId,omitempty"`
	ExtendedProperties                     map[string]*ParameterValues             `json:"extendedProperties,omitempty"`
	Id                                     *string                                 `json:"id,omitempty"`
	LicenseConnectionGroup                 *string                                 `json:"licenseConnectionGroup,omitempty"`
	LoggingMode                            *string                                 `json:"loggingMode,omitempty"`
	MetadataReloadSettings                 *ConnectionMetadataUrl                  `json:"metadataReloadSettings,omitempty"`
	Name                                   *string                                 `json:"name,omitempty"`
	Type                                   *string                                 `json:"type,omitempty"`
	VirtualEntityIds                       *[]*string                              `json:"virtualEntityIds,omitempty"`
}

//ConnectionCert - A certificate used for signature verification or XML encryption.
type ConnectionCert struct {
	ActiveVerificationCert    *bool     `json:"activeVerificationCert,omitempty"`
	CertView                  *CertView `json:"certView,omitempty"`
	EncryptionCert            *bool     `json:"encryptionCert,omitempty"`
	PrimaryVerificationCert   *bool     `json:"primaryVerificationCert,omitempty"`
	SecondaryVerificationCert *bool     `json:"secondaryVerificationCert,omitempty"`
	X509File                  *X509File `json:"x509File,omitempty"`
}

//ConnectionCerts - The certificates used for signature verification and XML encryption.
type ConnectionCerts struct {
	Items *[]*ConnectionCert `json:"items,omitempty"`
}

//ConnectionCredentials - The certificates and settings for encryption, signing, and signature verification.
type ConnectionCredentials struct {
	BlockEncryptionAlgorithm      *string                  `json:"blockEncryptionAlgorithm,omitempty"`
	Certs                         *[]*ConnectionCert       `json:"certs,omitempty"`
	DecryptionKeyPairRef          *ResourceLink            `json:"decryptionKeyPairRef,omitempty"`
	InboundBackChannelAuth        *InboundBackChannelAuth  `json:"inboundBackChannelAuth,omitempty"`
	KeyTransportAlgorithm         *string                  `json:"keyTransportAlgorithm,omitempty"`
	OutboundBackChannelAuth       *OutboundBackChannelAuth `json:"outboundBackChannelAuth,omitempty"`
	SecondaryDecryptionKeyPairRef *ResourceLink            `json:"secondaryDecryptionKeyPairRef,omitempty"`
	SigningSettings               *SigningSettings         `json:"signingSettings,omitempty"`
	VerificationIssuerDN          *string                  `json:"verificationIssuerDN,omitempty"`
	VerificationSubjectDN         *string                  `json:"verificationSubjectDN,omitempty"`
}

//ConnectionGroupLicenseView - Connection group license information.
type ConnectionGroupLicenseView struct {
	ConnectionCount *int    `json:"connectionCount,omitempty"`
	EndDate         *string `json:"endDate,omitempty"`
	Name            *string `json:"name,omitempty"`
	StartDate       *string `json:"startDate,omitempty"`
}

//ConnectionMetadataUrl - Configuration settings to enable automatic reload of partner's metadata.
type ConnectionMetadataUrl struct {
	EnableAutoMetadataUpdate *bool         `json:"enableAutoMetadataUpdate,omitempty"`
	MetadataUrlRef           *ResourceLink `json:"metadataUrlRef,omitempty"`
}

//ContactInfo - Contact information.
type ContactInfo struct {
	Company   *string `json:"company,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Phone     *string `json:"phone,omitempty"`
}

//ContinuePolicyAction - The continue selection action.
type ContinuePolicyAction struct {
	Context *string `json:"context,omitempty"`
	Type    *string `json:"type,omitempty"`
}

//ConvertMetadataRequest - A request for converting SAML connection metadata into a JSON connection.
type ConvertMetadataRequest struct {
	ConnectionType          *string     `json:"connectionType,omitempty"`
	ExpectedEntityId        *string     `json:"expectedEntityId,omitempty"`
	ExpectedProtocol        *string     `json:"expectedProtocol,omitempty"`
	SamlMetadata            *string     `json:"samlMetadata,omitempty"`
	TemplateConnection      *Connection `json:"templateConnection,omitempty"`
	VerificationCertificate *string     `json:"verificationCertificate,omitempty"`
}

//ConvertMetadataResponse - A response from converting SAML connection metadata into a JSON connection. It includes the converted connection and the authenticity information of the metadata.
type ConvertMetadataResponse struct {
	CertExpiration   *string     `json:"certExpiration,omitempty"`
	CertSerialNumber *string     `json:"certSerialNumber,omitempty"`
	CertSubjectDn    *string     `json:"certSubjectDn,omitempty"`
	CertTrustStatus  *string     `json:"certTrustStatus,omitempty"`
	Connection       *Connection `json:"connection,omitempty"`
	SignatureStatus  *string     `json:"signatureStatus,omitempty"`
}

//CrlSettings - CRL settings.
type CrlSettings struct {
	NextRetryMinsWhenNextUpdateInPast *int  `json:"nextRetryMinsWhenNextUpdateInPast,omitempty"`
	NextRetryMinsWhenResolveFailed    *int  `json:"nextRetryMinsWhenResolveFailed,omitempty"`
	TreatNonRetrievableCrlAsRevoked   *bool `json:"treatNonRetrievableCrlAsRevoked,omitempty"`
	VerifyCrlSignature                *bool `json:"verifyCrlSignature,omitempty"`
}

//CustomAttributeSource - The configured settings used to look up attributes from a custom data store.
type CustomAttributeSource struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	DataStoreRef                 *ResourceLink                         `json:"dataStoreRef,omitempty"`
	Description                  *string                               `json:"description,omitempty"`
	FilterFields                 *[]*FieldEntry                        `json:"filterFields,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	Type                         *string                               `json:"type,omitempty"`
}

//CustomDataStore - A custom data store.
type CustomDataStore struct {
	Configuration       *PluginConfiguration `json:"configuration,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	MaskAttributeValues *bool                `json:"maskAttributeValues,omitempty"`
	Name                *string              `json:"name,omitempty"`
	ParentRef           *ResourceLink        `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink        `json:"pluginDescriptorRef,omitempty"`
	Type                *string              `json:"type,omitempty"`
}

//CustomDataStoreDescriptor - A custom data store descriptor.
type CustomDataStoreDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//CustomDataStoreDescriptors - A collection of custom data store descriptors.
type CustomDataStoreDescriptors struct {
	Items *[]*CustomDataStoreDescriptor `json:"items,omitempty"`
}

//DataSourceTag
type DataSourceTag struct {
	DefaultSource *bool      `json:"defaultSource,omitempty"`
	Tags          *string    `json:"tags,omitempty"`
	TagsHashSet   *[]*string `json:"tagsHashSet,omitempty"`
}

//DataStore - The set of attributes used to configure a data store.
type DataStore struct {
	JdbcDataStore
	LdapDataStore
	CustomDataStore
	Id                  *string `json:"id,omitempty"`
	MaskAttributeValues *bool   `json:"maskAttributeValues,omitempty"`
	Type                *string `json:"type,omitempty"`
}

//DataStoreAttribute - The data store attribute.
type DataStoreAttribute struct {
	LdapDataStoreAttribute
	Metadata map[string]*string `json:"metadata,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Type     *string            `json:"type,omitempty"`
}

//DataStoreConfig - Local identity profile data store.
type DataStoreConfig struct {
	LdapDataStoreConfig
	DataStoreMapping map[string]*DataStoreAttribute `json:"dataStoreMapping,omitempty"`
	DataStoreRef     *ResourceLink                  `json:"dataStoreRef,omitempty"`
	Type             *string                        `json:"type,omitempty"`
}

//DataStores - A collection of data stores.
type DataStores struct {
	Items    *[]interface{}     `json:"-"`
	RawItems *[]json.RawMessage `json:"items,omitempty"`
}

//DateLocalIdentityField - A date type field.
type DateLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	DefaultValue          *string          `json:"defaultValue,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//DecryptionKeys - Decryption keys used to decrypt message content received from the partner.
type DecryptionKeys struct {
	PrimaryKeyRef       *ResourceLink `json:"primaryKeyRef,omitempty"`
	SecondaryKeyPairRef *ResourceLink `json:"secondaryKeyPairRef,omitempty"`
}

//DecryptionPolicy - Defines what to decrypt in the browser-based SSO profile.
type DecryptionPolicy struct {
	AssertionEncrypted        *bool `json:"assertionEncrypted,omitempty"`
	AttributesEncrypted       *bool `json:"attributesEncrypted,omitempty"`
	SloEncryptSubjectNameID   *bool `json:"sloEncryptSubjectNameID,omitempty"`
	SloSubjectNameIDEncrypted *bool `json:"sloSubjectNameIDEncrypted,omitempty"`
	SubjectNameIdEncrypted    *bool `json:"subjectNameIdEncrypted,omitempty"`
}

//DonePolicyAction - The done selection action.
type DonePolicyAction struct {
	Context *string `json:"context,omitempty"`
	Type    *string `json:"type,omitempty"`
}

//DropDownLocalIdentityField - A dropdown selection type field.
type DropDownLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	DefaultValue          *string          `json:"defaultValue,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	Options               *[]*string       `json:"options,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//DynamicClientRegistration - Dynamic client registration settings.
type DynamicClientRegistration struct {
	AllowClientDelete                        *bool                         `json:"allowClientDelete,omitempty"`
	AllowedExclusiveScopes                   *[]*string                    `json:"allowedExclusiveScopes,omitempty"`
	BypassActivationCodeConfirmationOverride *bool                         `json:"bypassActivationCodeConfirmationOverride,omitempty"`
	CibaPollingInterval                      *int                          `json:"cibaPollingInterval,omitempty"`
	CibaRequireSignedRequests                *bool                         `json:"cibaRequireSignedRequests,omitempty"`
	ClientCertIssuerRef                      *ResourceLink                 `json:"clientCertIssuerRef,omitempty"`
	ClientCertIssuerType                     *string                       `json:"clientCertIssuerType,omitempty"`
	DefaultAccessTokenManagerRef             *ResourceLink                 `json:"defaultAccessTokenManagerRef,omitempty"`
	DeviceFlowSettingType                    *string                       `json:"deviceFlowSettingType,omitempty"`
	DevicePollingIntervalOverride            *int                          `json:"devicePollingIntervalOverride,omitempty"`
	DisableRegistrationAccessTokens          *bool                         `json:"disableRegistrationAccessTokens,omitempty"`
	EnforceReplayPrevention                  *bool                         `json:"enforceReplayPrevention,omitempty"`
	InitialAccessTokenScope                  *string                       `json:"initialAccessTokenScope,omitempty"`
	OidcPolicy                               *ClientRegistrationOIDCPolicy `json:"oidcPolicy,omitempty"`
	PendingAuthorizationTimeoutOverride      *int                          `json:"pendingAuthorizationTimeoutOverride,omitempty"`
	PersistentGrantExpirationTime            *int                          `json:"persistentGrantExpirationTime,omitempty"`
	PersistentGrantExpirationTimeUnit        *string                       `json:"persistentGrantExpirationTimeUnit,omitempty"`
	PersistentGrantExpirationType            *string                       `json:"persistentGrantExpirationType,omitempty"`
	PersistentGrantIdleTimeout               *int                          `json:"persistentGrantIdleTimeout,omitempty"`
	PersistentGrantIdleTimeoutTimeUnit       *string                       `json:"persistentGrantIdleTimeoutTimeUnit,omitempty"`
	PersistentGrantIdleTimeoutType           *string                       `json:"persistentGrantIdleTimeoutType,omitempty"`
	PolicyRefs                               *[]*ResourceLink              `json:"policyRefs,omitempty"`
	RefreshRolling                           *string                       `json:"refreshRolling,omitempty"`
	RequestPolicyRef                         *ResourceLink                 `json:"requestPolicyRef,omitempty"`
	RequireProofKeyForCodeExchange           *bool                         `json:"requireProofKeyForCodeExchange,omitempty"`
	RequireSignedRequests                    *bool                         `json:"requireSignedRequests,omitempty"`
	RestrictCommonScopes                     *bool                         `json:"restrictCommonScopes,omitempty"`
	RestrictToDefaultAccessTokenManager      *bool                         `json:"restrictToDefaultAccessTokenManager,omitempty"`
	RestrictedCommonScopes                   *[]*string                    `json:"restrictedCommonScopes,omitempty"`
	RotateClientSecret                       *bool                         `json:"rotateClientSecret,omitempty"`
	RotateRegistrationAccessToken            *bool                         `json:"rotateRegistrationAccessToken,omitempty"`
	TokenExchangeProcessorPolicyRef          *ResourceLink                 `json:"tokenExchangeProcessorPolicyRef,omitempty"`
	UserAuthorizationUrlOverride             *string                       `json:"userAuthorizationUrlOverride,omitempty"`
}

//EmailLocalIdentityField - An email type field.
type EmailLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//EmailServerSettings - Email server configuration settings.
type EmailServerSettings struct {
	EmailServer              *string `json:"emailServer,omitempty"`
	EnableUtf8MessageHeaders *bool   `json:"enableUtf8MessageHeaders,omitempty"`
	EncryptedPassword        *string `json:"encryptedPassword,omitempty"`
	Password                 *string `json:"password,omitempty"`
	Port                     *int    `json:"port,omitempty"`
	RetryAttempts            *int    `json:"retryAttempts,omitempty"`
	RetryDelay               *int    `json:"retryDelay,omitempty"`
	SourceAddr               *string `json:"sourceAddr,omitempty"`
	SslPort                  *int    `json:"sslPort,omitempty"`
	Timeout                  *int    `json:"timeout,omitempty"`
	UseDebugging             *bool   `json:"useDebugging,omitempty"`
	UseSSL                   *bool   `json:"useSSL,omitempty"`
	UseTLS                   *bool   `json:"useTLS,omitempty"`
	Username                 *string `json:"username,omitempty"`
	VerifyHostname           *bool   `json:"verifyHostname,omitempty"`
}

//EmailVerificationConfig - A local identity email verification configuration.
type EmailVerificationConfig struct {
	EmailVerificationEnabled             *bool         `json:"emailVerificationEnabled,omitempty"`
	EmailVerificationErrorTemplateName   *string       `json:"emailVerificationErrorTemplateName,omitempty"`
	EmailVerificationSentTemplateName    *string       `json:"emailVerificationSentTemplateName,omitempty"`
	EmailVerificationSuccessTemplateName *string       `json:"emailVerificationSuccessTemplateName,omitempty"`
	FieldForEmailToVerify                *string       `json:"fieldForEmailToVerify,omitempty"`
	FieldStoringVerificationStatus       *string       `json:"fieldStoringVerificationStatus,omitempty"`
	NotificationPublisherRef             *ResourceLink `json:"notificationPublisherRef,omitempty"`
	OtlTimeToLive                        *int          `json:"otlTimeToLive,omitempty"`
	VerifyEmailTemplateName              *string       `json:"verifyEmailTemplateName,omitempty"`
}

//EncryptionPolicy - Defines what to encrypt in the browser-based SSO profile.
type EncryptionPolicy struct {
	EncryptAssertion          *bool      `json:"encryptAssertion,omitempty"`
	EncryptSloSubjectNameId   *bool      `json:"encryptSloSubjectNameId,omitempty"`
	EncryptedAttributes       *[]*string `json:"encryptedAttributes,omitempty"`
	SloSubjectNameIDEncrypted *bool      `json:"sloSubjectNameIDEncrypted,omitempty"`
}

//Entity
type Entity struct {
	EntityDescription *string `json:"entityDescription,omitempty"`
	EntityId          *string `json:"entityId,omitempty"`
}

//ExportMetadataRequest - The request for exporting a SAML connection's metadata file for a partner.
type ExportMetadataRequest struct {
	ConnectionId            *string          `json:"connectionId,omitempty"`
	ConnectionType          *string          `json:"connectionType,omitempty"`
	SigningSettings         *SigningSettings `json:"signingSettings,omitempty"`
	UseSecondaryPortForSoap *bool            `json:"useSecondaryPortForSoap,omitempty"`
	VirtualHostName         *string          `json:"virtualHostName,omitempty"`
	VirtualServerId         *string          `json:"virtualServerId,omitempty"`
}

//ExpressionIssuanceCriteriaEntry - An issuance criterion that uses a Boolean return value from an OGNL expression to determine whether or not it passes.
type ExpressionIssuanceCriteriaEntry struct {
	ErrorResult *string `json:"errorResult,omitempty"`
	Expression  *string `json:"expression,omitempty"`
}

//ExtendedProperties - A collection of Extended Properties definitions.
type ExtendedProperties struct {
	Items *[]*ExtendedProperty `json:"items,omitempty"`
}

//ExtendedProperty - Extended Property definition that allows to store additional information about IdP/SP Connections and OAuth Clients.
type ExtendedProperty struct {
	Description *string `json:"description,omitempty"`
	MultiValued *bool   `json:"multiValued,omitempty"`
	Name        *string `json:"name,omitempty"`
}

//FederationInfo - Federation Info.
type FederationInfo struct {
	AutoConnectEntityId *string `json:"autoConnectEntityId,omitempty"`
	BaseUrl             *string `json:"baseUrl,omitempty"`
	Saml1xIssuerId      *string `json:"saml1xIssuerId,omitempty"`
	Saml1xSourceId      *string `json:"saml1xSourceId,omitempty"`
	Saml2EntityId       *string `json:"saml2EntityId,omitempty"`
	WsfedRealm          *string `json:"wsfedRealm,omitempty"`
}

//FieldConfig - A local identity profile field configuration.
type FieldConfig struct {
	Fields                    *[]*LocalIdentityField `json:"fields,omitempty"`
	StripSpaceFromUniqueField *bool                  `json:"stripSpaceFromUniqueField,omitempty"`
}

//FieldDescriptor - Describes a plugin configuration field.
type FieldDescriptor struct {
	RadioGroupFieldDescriptor
	SelectFieldDescriptor
	CheckBoxFieldDescriptor
	UploadFileFieldDescriptor
	TextAreaFieldDescriptor
	TextFieldDescriptor
	HashedTextFieldDescriptor
	Advanced     *bool   `json:"advanced,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//FieldEntry - A simple name value pair to represent a field entry.
type FieldEntry struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

//FragmentPolicyAction - A authentication policy fragment selection action.
type FragmentPolicyAction struct {
	AttributeRules  *AttributeRules   `json:"attributeRules,omitempty"`
	Context         *string           `json:"context,omitempty"`
	Fragment        *ResourceLink     `json:"fragment,omitempty"`
	FragmentMapping *AttributeMapping `json:"fragmentMapping,omitempty"`
	Type            *string           `json:"type,omitempty"`
}

//GlobalAuthenticationSessionPolicy - The global policy for authentication sessions.
type GlobalAuthenticationSessionPolicy struct {
	EnableSessions             *bool   `json:"enableSessions,omitempty"`
	HashUniqueUserKeyAttribute *bool   `json:"hashUniqueUserKeyAttribute,omitempty"`
	IdleTimeoutDisplayUnit     *string `json:"idleTimeoutDisplayUnit,omitempty"`
	IdleTimeoutMins            *int    `json:"idleTimeoutMins,omitempty"`
	MaxTimeoutDisplayUnit      *string `json:"maxTimeoutDisplayUnit,omitempty"`
	MaxTimeoutMins             *int    `json:"maxTimeoutMins,omitempty"`
	PersistentSessions         *bool   `json:"persistentSessions,omitempty"`
}

//GroupMembershipDetection - Settings to detect group memberships.
type GroupMembershipDetection struct {
	GroupMemberAttributeName   *string `json:"groupMemberAttributeName,omitempty"`
	MemberOfGroupAttributeName *string `json:"memberOfGroupAttributeName,omitempty"`
}

//HashedTextFieldDescriptor - A text field that will contain a secure salted hash.
type HashedTextFieldDescriptor struct {
	Advanced     *bool   `json:"advanced,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Size         *int    `json:"size,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//HiddenLocalIdentityField - A hidden selection type field.
type HiddenLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//IdentityHintAttribute - An attribute for the ciba request policy's identity hint attribute contract.
type IdentityHintAttribute struct {
	Name *string `json:"name,omitempty"`
}

//IdentityHintContract - A set of attributes exposed by request policy contract.
type IdentityHintContract struct {
	CoreAttributes     *[]*IdentityHintAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*IdentityHintAttribute `json:"extendedAttributes,omitempty"`
}

//IdpAdapter - An IdP adapter instance.
type IdpAdapter struct {
	AttributeContract   *IdpAdapterAttributeContract `json:"attributeContract,omitempty"`
	AttributeMapping    *IdpAdapterContractMapping   `json:"attributeMapping,omitempty"`
	AuthnCtxClassRef    *string                      `json:"authnCtxClassRef,omitempty"`
	Configuration       *PluginConfiguration         `json:"configuration,omitempty"`
	Id                  *string                      `json:"id,omitempty"`
	Name                *string                      `json:"name,omitempty"`
	ParentRef           *ResourceLink                `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink                `json:"pluginDescriptorRef,omitempty"`
}

//IdpAdapterAssertionMapping - The IdP Adapter Assertion Mapping.
type IdpAdapterAssertionMapping struct {
	AbortSsoTransactionAsFailSafe *bool                                 `json:"abortSsoTransactionAsFailSafe,omitempty"`
	AdapterOverrideSettings       *IdpAdapter                           `json:"adapterOverrideSettings,omitempty"`
	AttributeContractFulfillment  map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources              *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IdpAdapterRef                 *ResourceLink                         `json:"idpAdapterRef,omitempty"`
	IssuanceCriteria              *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictVirtualEntityIds      *bool                                 `json:"restrictVirtualEntityIds,omitempty"`
	RestrictedVirtualEntityIds    *[]*string                            `json:"restrictedVirtualEntityIds,omitempty"`
}

//IdpAdapterAttribute - An attribute for the IdP adapter attribute contract.
type IdpAdapterAttribute struct {
	Masked    *bool   `json:"masked,omitempty"`
	Name      *string `json:"name,omitempty"`
	Pseudonym *bool   `json:"pseudonym,omitempty"`
}

//IdpAdapterAttributeContract - A set of attributes exposed by an IdP adapter.
type IdpAdapterAttributeContract struct {
	CoreAttributes         *[]*IdpAdapterAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes     *[]*IdpAdapterAttribute `json:"extendedAttributes,omitempty"`
	Inherited              *bool                   `json:"inherited,omitempty"`
	MaskOgnlValues         *bool                   `json:"maskOgnlValues,omitempty"`
	UniqueUserKeyAttribute *string                 `json:"uniqueUserKeyAttribute,omitempty"`
}

//IdpAdapterContractMapping
type IdpAdapterContractMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Inherited                    *bool                                 `json:"inherited,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//IdpAdapterDescriptor - An IdP adapter descriptor.
type IdpAdapterDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//IdpAdapterDescriptors - A collection of IdP adapter descriptors.
type IdpAdapterDescriptors struct {
	Items *[]*IdpAdapterDescriptor `json:"items,omitempty"`
}

//IdpAdapterMapping - The OAuth IdP Adapter Mapping.
type IdpAdapterMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	IdpAdapterRef                *ResourceLink                         `json:"idpAdapterRef,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//IdpAdapterMappings - A collection of OAuth IdP adapter mapping items.
type IdpAdapterMappings struct {
	Items *[]*IdpAdapterMapping `json:"items,omitempty"`
}

//IdpAdapters - A collection of IdP adapters.
type IdpAdapters struct {
	Items *[]*IdpAdapter `json:"items,omitempty"`
}

//IdpAttributeQuery - The attribute query profile supports local applications in requesting user attributes from an attribute authority.
type IdpAttributeQuery struct {
	NameMappings *[]*AttributeQueryNameMapping `json:"nameMappings,omitempty"`
	Policy       *IdpAttributeQueryPolicy      `json:"policy,omitempty"`
	Url          *string                       `json:"url,omitempty"`
}

//IdpAttributeQueryPolicy - The attribute query profile's security policy.
type IdpAttributeQueryPolicy struct {
	EncryptNameId             *bool `json:"encryptNameId,omitempty"`
	MaskAttributeValues       *bool `json:"maskAttributeValues,omitempty"`
	RequireEncryptedAssertion *bool `json:"requireEncryptedAssertion,omitempty"`
	RequireSignedAssertion    *bool `json:"requireSignedAssertion,omitempty"`
	RequireSignedResponse     *bool `json:"requireSignedResponse,omitempty"`
	SignAttributeQuery        *bool `json:"signAttributeQuery,omitempty"`
}

//IdpBrowserSso - The settings used to enable secure browser-based SSO to resources at your site.
type IdpBrowserSso struct {
	AdapterMappings                      *[]*SpAdapterMapping                    `json:"adapterMappings,omitempty"`
	AlwaysSignArtifactResponse           *bool                                   `json:"alwaysSignArtifactResponse,omitempty"`
	Artifact                             *ArtifactSettings                       `json:"artifact,omitempty"`
	AssertionsSigned                     *bool                                   `json:"assertionsSigned,omitempty"`
	AttributeContract                    *IdpBrowserSsoAttributeContract         `json:"attributeContract,omitempty"`
	AuthenticationPolicyContractMappings *[]*AuthenticationPolicyContractMapping `json:"authenticationPolicyContractMappings,omitempty"`
	AuthnContextMappings                 *[]*AuthnContextMapping                 `json:"authnContextMappings,omitempty"`
	DecryptionPolicy                     *DecryptionPolicy                       `json:"decryptionPolicy,omitempty"`
	DefaultTargetUrl                     *string                                 `json:"defaultTargetUrl,omitempty"`
	EnabledProfiles                      *[]*string                              `json:"enabledProfiles,omitempty"`
	IdpIdentityMapping                   *string                                 `json:"idpIdentityMapping,omitempty"`
	IncomingBindings                     *[]*string                              `json:"incomingBindings,omitempty"`
	MessageCustomizations                *[]*ProtocolMessageCustomization        `json:"messageCustomizations,omitempty"`
	OauthAuthenticationPolicyContractRef *ResourceLink                           `json:"oauthAuthenticationPolicyContractRef,omitempty"`
	OidcProviderSettings                 *OIDCProviderSettings                   `json:"oidcProviderSettings,omitempty"`
	Protocol                             *string                                 `json:"protocol,omitempty"`
	SignAuthnRequests                    *bool                                   `json:"signAuthnRequests,omitempty"`
	SloServiceEndpoints                  *[]*SloServiceEndpoint                  `json:"sloServiceEndpoints,omitempty"`
	SsoOAuthMapping                      *SsoOAuthMapping                        `json:"ssoOAuthMapping,omitempty"`
	SsoServiceEndpoints                  *[]*IdpSsoServiceEndpoint               `json:"ssoServiceEndpoints,omitempty"`
	UrlWhitelistEntries                  *[]*UrlWhitelistEntry                   `json:"urlWhitelistEntries,omitempty"`
}

//IdpBrowserSsoAttribute - An attribute for the IdP Browser SSO attribute contract.
type IdpBrowserSsoAttribute struct {
	Masked *bool   `json:"masked,omitempty"`
	Name   *string `json:"name,omitempty"`
}

//IdpBrowserSsoAttributeContract - A set of user attributes that the IdP sends in the SAML assertion.
type IdpBrowserSsoAttributeContract struct {
	CoreAttributes     *[]*IdpBrowserSsoAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*IdpBrowserSsoAttribute `json:"extendedAttributes,omitempty"`
}

//IdpConnection - The set of attributes used to configure an IdP connection.
type IdpConnection struct {
	Active                                 *bool                                   `json:"active,omitempty"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfiguration `json:"additionalAllowedEntitiesConfiguration,omitempty"`
	AttributeQuery                         *IdpAttributeQuery                      `json:"attributeQuery,omitempty"`
	BaseUrl                                *string                                 `json:"baseUrl,omitempty"`
	ContactInfo                            *ContactInfo                            `json:"contactInfo,omitempty"`
	Credentials                            *ConnectionCredentials                  `json:"credentials,omitempty"`
	DefaultVirtualEntityId                 *string                                 `json:"defaultVirtualEntityId,omitempty"`
	EntityId                               *string                                 `json:"entityId,omitempty"`
	ErrorPageMsgId                         *string                                 `json:"errorPageMsgId,omitempty"`
	ExtendedProperties                     map[string]*ParameterValues             `json:"extendedProperties,omitempty"`
	Id                                     *string                                 `json:"id,omitempty"`
	IdpBrowserSso                          *IdpBrowserSso                          `json:"idpBrowserSso,omitempty"`
	IdpOAuthGrantAttributeMapping          *IdpOAuthGrantAttributeMapping          `json:"idpOAuthGrantAttributeMapping,omitempty"`
	LicenseConnectionGroup                 *string                                 `json:"licenseConnectionGroup,omitempty"`
	LoggingMode                            *string                                 `json:"loggingMode,omitempty"`
	MetadataReloadSettings                 *ConnectionMetadataUrl                  `json:"metadataReloadSettings,omitempty"`
	Name                                   *string                                 `json:"name,omitempty"`
	OidcClientCredentials                  *OIDCClientCredentials                  `json:"oidcClientCredentials,omitempty"`
	Type                                   *string                                 `json:"type,omitempty"`
	VirtualEntityIds                       *[]*string                              `json:"virtualEntityIds,omitempty"`
	WsTrust                                *IdpWsTrust                             `json:"wsTrust,omitempty"`
}

//IdpConnections - A collection of IdP connections.
type IdpConnections struct {
	Items *[]*IdpConnection `json:"items,omitempty"`
}

//IdpDefaultUrl - IDP Default URL settings.
type IdpDefaultUrl struct {
	ConfirmIdpSlo    *bool   `json:"confirmIdpSlo,omitempty"`
	IdpErrorMsg      *string `json:"idpErrorMsg,omitempty"`
	IdpSloSuccessUrl *string `json:"idpSloSuccessUrl,omitempty"`
}

//IdpOAuthAttributeContract - A set of user attributes that the IdP sends in the OAuth Assertion Grant.
type IdpOAuthAttributeContract struct {
	CoreAttributes     *[]*IdpBrowserSsoAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*IdpBrowserSsoAttribute `json:"extendedAttributes,omitempty"`
}

//IdpOAuthGrantAttributeMapping - The OAuth Assertion Grant settings used to map from your IdP.
type IdpOAuthGrantAttributeMapping struct {
	AccessTokenManagerMappings *[]*AccessTokenManagerMapping `json:"accessTokenManagerMappings,omitempty"`
	IdpOAuthAttributeContract  *IdpOAuthAttributeContract    `json:"idpOAuthAttributeContract,omitempty"`
}

//IdpRole - This property has been deprecated and is no longer used. All Roles and protocols are always enabled.
type IdpRole struct {
	Enable                     *bool          `json:"enable,omitempty"`
	EnableOutboundProvisioning *bool          `json:"enableOutboundProvisioning,omitempty"`
	EnableSaml10               *bool          `json:"enableSaml10,omitempty"`
	EnableSaml11               *bool          `json:"enableSaml11,omitempty"`
	EnableWsFed                *bool          `json:"enableWsFed,omitempty"`
	EnableWsTrust              *bool          `json:"enableWsTrust,omitempty"`
	Saml20Profile              *SAML20Profile `json:"saml20Profile,omitempty"`
}

//IdpSsoServiceEndpoint - The settings that define an endpoint to an IdP SSO service.
type IdpSsoServiceEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	Url     *string `json:"url,omitempty"`
}

//IdpToSpAdapterMapping - The IdP-to-SP Adapter Mapping.
type IdpToSpAdapterMapping struct {
	ApplicationIconUrl               *string                               `json:"applicationIconUrl,omitempty"`
	ApplicationName                  *string                               `json:"applicationName,omitempty"`
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                 *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	DefaultTargetResource            *string                               `json:"defaultTargetResource,omitempty"`
	Id                               *string                               `json:"id,omitempty"`
	IssuanceCriteria                 *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	LicenseConnectionGroupAssignment *string                               `json:"licenseConnectionGroupAssignment,omitempty"`
	SourceId                         *string                               `json:"sourceId,omitempty"`
	TargetId                         *string                               `json:"targetId,omitempty"`
}

//IdpToSpAdapterMappings - A collection of IdP-to-SP Adapter Mappings.
type IdpToSpAdapterMappings struct {
	Items *[]*IdpToSpAdapterMapping `json:"items,omitempty"`
}

//IdpTokenProcessorMapping - The IdP Token Processor Mapping.
type IdpTokenProcessorMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IdpTokenProcessorRef         *ResourceLink                         `json:"idpTokenProcessorRef,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictedVirtualEntityIds   *[]*string                            `json:"restrictedVirtualEntityIds,omitempty"`
}

//IdpWsTrust - Ws-Trust STS provides validation of incoming tokens which enable SSO access to Web Services. It also allows generation of local tokens for Web Services.
type IdpWsTrust struct {
	AttributeContract      *IdpWsTrustAttributeContract `json:"attributeContract,omitempty"`
	GenerateLocalToken     *bool                        `json:"generateLocalToken,omitempty"`
	TokenGeneratorMappings *[]*SpTokenGeneratorMapping  `json:"tokenGeneratorMappings,omitempty"`
}

//IdpWsTrustAttribute - An attribute for the Ws-Trust attribute contract.
type IdpWsTrustAttribute struct {
	Masked *bool   `json:"masked,omitempty"`
	Name   *string `json:"name,omitempty"`
}

//IdpWsTrustAttributeContract - A set of user attributes that this server will receive in the token.
type IdpWsTrustAttributeContract struct {
	CoreAttributes     *[]*IdpWsTrustAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*IdpWsTrustAttribute `json:"extendedAttributes,omitempty"`
}

//InboundBackChannelAuth
type InboundBackChannelAuth struct {
	Certs                 *[]*ConnectionCert           `json:"certs,omitempty"`
	DigitalSignature      *bool                        `json:"digitalSignature,omitempty"`
	HttpBasicCredentials  *UsernamePasswordCredentials `json:"httpBasicCredentials,omitempty"`
	RequireSsl            *bool                        `json:"requireSsl,omitempty"`
	Type                  *string                      `json:"type,omitempty"`
	VerificationIssuerDN  *string                      `json:"verificationIssuerDN,omitempty"`
	VerificationSubjectDN *string                      `json:"verificationSubjectDN,omitempty"`
}

//IncomingProxySettings - Incoming Proxy Settings.
type IncomingProxySettings struct {
	ClientCertChainSSLHeaderName  *string `json:"clientCertChainSSLHeaderName,omitempty"`
	ClientCertSSLHeaderName       *string `json:"clientCertSSLHeaderName,omitempty"`
	ForwardedHostHeaderIndex      *string `json:"forwardedHostHeaderIndex,omitempty"`
	ForwardedHostHeaderName       *string `json:"forwardedHostHeaderName,omitempty"`
	ForwardedIpAddressHeaderIndex *string `json:"forwardedIpAddressHeaderIndex,omitempty"`
	ForwardedIpAddressHeaderName  *string `json:"forwardedIpAddressHeaderName,omitempty"`
	ProxyTerminatesHttpsConns     *bool   `json:"proxyTerminatesHttpsConns,omitempty"`
}

//IssuanceCriteria - A list of criteria that determines whether a transaction (usually a SSO transaction) is continued. All criteria must pass in order for the transaction to continue.
type IssuanceCriteria struct {
	ConditionalCriteria *[]*ConditionalIssuanceCriteriaEntry `json:"conditionalCriteria,omitempty"`
	ExpressionCriteria  *[]*ExpressionIssuanceCriteriaEntry  `json:"expressionCriteria,omitempty"`
}

//JdbcAttributeSource - The configured settings used to look up attributes from a JDBC data store.
type JdbcAttributeSource struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	DataStoreRef                 *ResourceLink                         `json:"dataStoreRef,omitempty"`
	Description                  *string                               `json:"description,omitempty"`
	Filter                       *string                               `json:"filter,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	Schema                       *string                               `json:"schema,omitempty"`
	Table                        *string                               `json:"table,omitempty"`
	Type                         *string                               `json:"type,omitempty"`
}

//JdbcDataStore - A JDBC data store.
type JdbcDataStore struct {
	AllowMultiValueAttributes *bool             `json:"allowMultiValueAttributes,omitempty"`
	BlockingTimeout           *int              `json:"blockingTimeout,omitempty"`
	ConnectionUrl             *string           `json:"connectionUrl,omitempty"`
	ConnectionUrlTags         *[]*JdbcTagConfig `json:"connectionUrlTags,omitempty"`
	DriverClass               *string           `json:"driverClass,omitempty"`
	EncryptedPassword         *string           `json:"encryptedPassword,omitempty"`
	Id                        *string           `json:"id,omitempty"`
	IdleTimeout               *int              `json:"idleTimeout,omitempty"`
	MaskAttributeValues       *bool             `json:"maskAttributeValues,omitempty"`
	MaxPoolSize               *int              `json:"maxPoolSize,omitempty"`
	MinPoolSize               *int              `json:"minPoolSize,omitempty"`
	Name                      *string           `json:"name,omitempty"`
	Password                  *string           `json:"password,omitempty"`
	Type                      *string           `json:"type,omitempty"`
	UserName                  *string           `json:"userName,omitempty"`
	ValidateConnectionSql     *string           `json:"validateConnectionSql,omitempty"`
}

//JdbcTagConfig - A JDBC data store's connection URLs and tags configuration. This is required if no default JDBC database location is specified.
type JdbcTagConfig struct {
	ConnectionUrl *string `json:"connectionUrl,omitempty"`
	DefaultSource *bool   `json:"defaultSource,omitempty"`
	Tags          *string `json:"tags,omitempty"`
}

//JwksSettings - JSON Web Key Set Settings.
type JwksSettings struct {
	Jwks    *string `json:"jwks,omitempty"`
	JwksUrl *string `json:"jwksUrl,omitempty"`
}

//KerberosRealm
type KerberosRealm struct {
	Id                              *string    `json:"id,omitempty"`
	KerberosEncryptedPassword       *string    `json:"kerberosEncryptedPassword,omitempty"`
	KerberosPassword                *string    `json:"kerberosPassword,omitempty"`
	KerberosRealmName               *string    `json:"kerberosRealmName,omitempty"`
	KerberosUsername                *string    `json:"kerberosUsername,omitempty"`
	KeyDistributionCenters          *[]*string `json:"keyDistributionCenters,omitempty"`
	SuppressDomainNameConcatenation *bool      `json:"suppressDomainNameConcatenation,omitempty"`
}

//KerberosRealms - A collection of Kerberos Realms.
type KerberosRealms struct {
	Items *[]*KerberosRealm `json:"items,omitempty"`
}

//KerberosRealmsSettings - Settings for all of the Kerberos Realms.
type KerberosRealmsSettings struct {
	DebugLogOutput *bool   `json:"debugLogOutput,omitempty"`
	ForceTcp       *bool   `json:"forceTcp,omitempty"`
	KdcRetries     *string `json:"kdcRetries,omitempty"`
	KdcTimeout     *string `json:"kdcTimeout,omitempty"`
}

//KeyAlgorithm - Details for a key algorithm.
type KeyAlgorithm struct {
	DefaultKeySize            *int       `json:"defaultKeySize,omitempty"`
	DefaultSignatureAlgorithm *string    `json:"defaultSignatureAlgorithm,omitempty"`
	KeySizes                  *[]*int    `json:"keySizes,omitempty"`
	Name                      *string    `json:"name,omitempty"`
	SignatureAlgorithms       *[]*string `json:"signatureAlgorithms,omitempty"`
}

//KeyAlgorithms - A collection of key algorithms.
type KeyAlgorithms struct {
	Items *[]*KeyAlgorithm `json:"items,omitempty"`
}

//KeyPairRotationSettings - Key Pair Rotation Details
type KeyPairRotationSettings struct {
	ActivationBufferDays *int    `json:"activationBufferDays,omitempty"`
	CreationBufferDays   *int    `json:"creationBufferDays,omitempty"`
	Id                   *string `json:"id,omitempty"`
	KeyAlgorithm         *string `json:"keyAlgorithm,omitempty"`
	KeySize              *int    `json:"keySize,omitempty"`
	SignatureAlgorithm   *string `json:"signatureAlgorithm,omitempty"`
	ValidDays            *int    `json:"validDays,omitempty"`
}

//KeyPairView - Key pair details.
type KeyPairView struct {
	CryptoProvider          *string                  `json:"cryptoProvider,omitempty"`
	Expires                 *string                  `json:"expires,omitempty"`
	Id                      *string                  `json:"id,omitempty"`
	IssuerDN                *string                  `json:"issuerDN,omitempty"`
	KeyAlgorithm            *string                  `json:"keyAlgorithm,omitempty"`
	KeySize                 *int                     `json:"keySize,omitempty"`
	RotationSettings        *KeyPairRotationSettings `json:"rotationSettings,omitempty"`
	SerialNumber            *string                  `json:"serialNumber,omitempty"`
	Sha1Fingerprint         *string                  `json:"sha1Fingerprint,omitempty"`
	Sha256Fingerprint       *string                  `json:"sha256Fingerprint,omitempty"`
	SignatureAlgorithm      *string                  `json:"signatureAlgorithm,omitempty"`
	Status                  *string                  `json:"status,omitempty"`
	SubjectAlternativeNames *[]*string               `json:"subjectAlternativeNames,omitempty"`
	SubjectDN               *string                  `json:"subjectDN,omitempty"`
	ValidFrom               *string                  `json:"validFrom,omitempty"`
	Version                 *int                     `json:"version,omitempty"`
}

//KeyPairViews - A collection of KeyPairView items.
type KeyPairViews struct {
	Items *[]*KeyPairView `json:"items,omitempty"`
}

//LdapAttributeSource - The configured settings used to look up attributes from a LDAP data store.
type LdapAttributeSource struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue   `json:"attributeContractFulfillment,omitempty"`
	BaseDn                       *string                                 `json:"baseDn,omitempty"`
	BinaryAttributeSettings      map[string]*BinaryLdapAttributeSettings `json:"binaryAttributeSettings,omitempty"`
	DataStoreRef                 *ResourceLink                           `json:"dataStoreRef,omitempty"`
	Description                  *string                                 `json:"description,omitempty"`
	Id                           *string                                 `json:"id,omitempty"`
	MemberOfNestedGroup          *bool                                   `json:"memberOfNestedGroup,omitempty"`
	SearchFilter                 *string                                 `json:"searchFilter,omitempty"`
	SearchScope                  *string                                 `json:"searchScope,omitempty"`
	Type                         *string                                 `json:"type,omitempty"`
}

//LdapDataStore - A LDAP data store.
type LdapDataStore struct {
	BinaryAttributes     *[]*string        `json:"binaryAttributes,omitempty"`
	BindAnonymously      *bool             `json:"bindAnonymously,omitempty"`
	ConnectionTimeout    *int              `json:"connectionTimeout,omitempty"`
	CreateIfNecessary    *bool             `json:"createIfNecessary,omitempty"`
	DnsTtl               *int              `json:"dnsTtl,omitempty"`
	EncryptedPassword    *string           `json:"encryptedPassword,omitempty"`
	FollowLDAPReferrals  *bool             `json:"followLDAPReferrals,omitempty"`
	Hostnames            *[]*string        `json:"hostnames,omitempty"`
	HostnamesTags        *[]*LdapTagConfig `json:"hostnamesTags,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	LdapDnsSrvPrefix     *string           `json:"ldapDnsSrvPrefix,omitempty"`
	LdapType             *string           `json:"ldapType,omitempty"`
	LdapsDnsSrvPrefix    *string           `json:"ldapsDnsSrvPrefix,omitempty"`
	MaskAttributeValues  *bool             `json:"maskAttributeValues,omitempty"`
	MaxConnections       *int              `json:"maxConnections,omitempty"`
	MaxWait              *int              `json:"maxWait,omitempty"`
	MinConnections       *int              `json:"minConnections,omitempty"`
	Name                 *string           `json:"name,omitempty"`
	Password             *string           `json:"password,omitempty"`
	ReadTimeout          *int              `json:"readTimeout,omitempty"`
	TestOnBorrow         *bool             `json:"testOnBorrow,omitempty"`
	TestOnReturn         *bool             `json:"testOnReturn,omitempty"`
	TimeBetweenEvictions *int              `json:"timeBetweenEvictions,omitempty"`
	Type                 *string           `json:"type,omitempty"`
	UseDnsSrvRecords     *bool             `json:"useDnsSrvRecords,omitempty"`
	UseSsl               *bool             `json:"useSsl,omitempty"`
	UserDN               *string           `json:"userDN,omitempty"`
	VerifyHost           *bool             `json:"verifyHost,omitempty"`
}

//LdapDataStoreAttribute - LDAP data store attribute.
type LdapDataStoreAttribute struct {
	Metadata map[string]*string `json:"metadata,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Type     *string            `json:"type,omitempty"`
}

//LdapDataStoreConfig - LDAP data store configuration.
type LdapDataStoreConfig struct {
	AuxiliaryObjectClasses *[]*string                     `json:"auxiliaryObjectClasses,omitempty"`
	BaseDn                 *string                        `json:"baseDn,omitempty"`
	CreatePattern          *string                        `json:"createPattern,omitempty"`
	DataStoreMapping       map[string]*DataStoreAttribute `json:"dataStoreMapping,omitempty"`
	DataStoreRef           *ResourceLink                  `json:"dataStoreRef,omitempty"`
	ObjectClass            *string                        `json:"objectClass,omitempty"`
	Type                   *string                        `json:"type,omitempty"`
}

//LdapTagConfig - An LDAP data store's hostnames and tags configuration. This is required if no default hostname is specified.
type LdapTagConfig struct {
	DefaultSource *bool      `json:"defaultSource,omitempty"`
	Hostnames     *[]*string `json:"hostnames,omitempty"`
	Tags          *string    `json:"tags,omitempty"`
}

//LicenseAgreementInfo - PingFederate License Agreement information.
type LicenseAgreementInfo struct {
	Accepted            *bool   `json:"accepted,omitempty"`
	LicenseAgreementUrl *string `json:"licenseAgreementUrl,omitempty"`
}

//LicenseEventNotificationSettings - Notification settings for licensing events.
type LicenseEventNotificationSettings struct {
	EmailAddress             *string       `json:"emailAddress,omitempty"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty"`
}

//LicenseFeatureView - PingFederate license feature details.
type LicenseFeatureView struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

//LicenseFile - License to import.
type LicenseFile struct {
	FileData *string `json:"fileData,omitempty"`
}

//LicenseView - PingFederate License details.
type LicenseView struct {
	BridgeMode          *bool                          `json:"bridgeMode,omitempty"`
	EnforcementType     *string                        `json:"enforcementType,omitempty"`
	ExpirationDate      *string                        `json:"expirationDate,omitempty"`
	Features            *[]*LicenseFeatureView         `json:"features,omitempty"`
	GracePeriod         *int                           `json:"gracePeriod,omitempty"`
	Id                  *string                        `json:"id,omitempty"`
	IssueDate           *string                        `json:"issueDate,omitempty"`
	LicenseGroups       *[]*ConnectionGroupLicenseView `json:"licenseGroups,omitempty"`
	MaxConnections      *int                           `json:"maxConnections,omitempty"`
	Name                *string                        `json:"name,omitempty"`
	NodeLimit           *int                           `json:"nodeLimit,omitempty"`
	OauthEnabled        *bool                          `json:"oauthEnabled,omitempty"`
	Organization        *string                        `json:"organization,omitempty"`
	Product             *string                        `json:"product,omitempty"`
	ProvisioningEnabled *bool                          `json:"provisioningEnabled,omitempty"`
	Tier                *string                        `json:"tier,omitempty"`
	UsedConnections     *int                           `json:"usedConnections,omitempty"`
	Version             *string                        `json:"version,omitempty"`
	WsTrustEnabled      *bool                          `json:"wsTrustEnabled,omitempty"`
}

//LocalIdentityAuthSource - An authentication source name.
type LocalIdentityAuthSource struct {
	Id     *string `json:"id,omitempty"`
	Source *string `json:"source,omitempty"`
}

//LocalIdentityAuthSourceUpdatePolicy - Settings to determine whether to store attributes that came from third-party authentication sources.
type LocalIdentityAuthSourceUpdatePolicy struct {
	RetainAttributes *bool    `json:"retainAttributes,omitempty"`
	StoreAttributes  *bool    `json:"storeAttributes,omitempty"`
	UpdateAttributes *bool    `json:"updateAttributes,omitempty"`
	UpdateInterval   *float32 `json:"updateInterval,omitempty"`
}

//LocalIdentityField - Local identity profile fields.
type LocalIdentityField struct {
	CheckboxGroupLocalIdentityField
	CheckboxLocalIdentityField
	DateLocalIdentityField
	TextLocalIdentityField
	DropDownLocalIdentityField
	EmailLocalIdentityField
	PhoneLocalIdentityField
	HiddenLocalIdentityField
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//LocalIdentityMappingPolicyAction - A local identity profile selection action.
type LocalIdentityMappingPolicyAction struct {
	Context                  *string           `json:"context,omitempty"`
	InboundMapping           *AttributeMapping `json:"inboundMapping,omitempty"`
	LocalIdentityRef         *ResourceLink     `json:"localIdentityRef,omitempty"`
	OutboundAttributeMapping *AttributeMapping `json:"outboundAttributeMapping,omitempty"`
	Type                     *string           `json:"type,omitempty"`
}

//LocalIdentityProfile - A local identity profile.
type LocalIdentityProfile struct {
	ApcId                   *ResourceLink                        `json:"apcId,omitempty"`
	AuthSourceUpdatePolicy  *LocalIdentityAuthSourceUpdatePolicy `json:"authSourceUpdatePolicy,omitempty"`
	AuthSources             *[]*LocalIdentityAuthSource          `json:"authSources,omitempty"`
	DataStoreConfig         *DataStoreConfig                     `json:"dataStoreConfig,omitempty"`
	EmailVerificationConfig *EmailVerificationConfig             `json:"emailVerificationConfig,omitempty"`
	FieldConfig             *FieldConfig                         `json:"fieldConfig,omitempty"`
	Id                      *string                              `json:"id,omitempty"`
	Name                    *string                              `json:"name,omitempty"`
	ProfileConfig           *ProfileConfig                       `json:"profileConfig,omitempty"`
	ProfileEnabled          *bool                                `json:"profileEnabled,omitempty"`
	RegistrationConfig      *RegistrationConfig                  `json:"registrationConfig,omitempty"`
	RegistrationEnabled     *bool                                `json:"registrationEnabled,omitempty"`
}

//LocalIdentityProfiles - A collection of local identity profiles.
type LocalIdentityProfiles struct {
	Items *[]*LocalIdentityProfile `json:"items,omitempty"`
}

//MetadataEventNotificationSettings - Notification settings for metadata update events.
type MetadataEventNotificationSettings struct {
	EmailAddress             *string       `json:"emailAddress,omitempty"`
	NotificationPublisherRef *ResourceLink `json:"notificationPublisherRef,omitempty"`
}

//MetadataLifetimeSettings - Metadata lifetime settings.
type MetadataLifetimeSettings struct {
	CacheDuration *int `json:"cacheDuration,omitempty"`
	ReloadDelay   *int `json:"reloadDelay,omitempty"`
}

//MetadataSigningSettings - Metadata signing settings. If metadata is not signed, this model will be empty.
type MetadataSigningSettings struct {
	SignatureAlgorithm *string       `json:"signatureAlgorithm,omitempty"`
	SigningKeyRef      *ResourceLink `json:"signingKeyRef,omitempty"`
}

//MetadataUrl - Metadata URL and corresponding Signature Verification Certificate.
type MetadataUrl struct {
	CertView          *CertView `json:"certView,omitempty"`
	Id                *string   `json:"id,omitempty"`
	Name              *string   `json:"name,omitempty"`
	Url               *string   `json:"url,omitempty"`
	ValidateSignature *bool     `json:"validateSignature,omitempty"`
	X509File          *X509File `json:"x509File,omitempty"`
}

//MetadataUrls
type MetadataUrls struct {
	Items *[]*MetadataUrl `json:"items,omitempty"`
}

//NewKeyPairSettings - Settings for creating a new key pair.
type NewKeyPairSettings struct {
	City                    *string    `json:"city,omitempty"`
	CommonName              *string    `json:"commonName,omitempty"`
	Country                 *string    `json:"country,omitempty"`
	CryptoProvider          *string    `json:"cryptoProvider,omitempty"`
	Id                      *string    `json:"id,omitempty"`
	KeyAlgorithm            *string    `json:"keyAlgorithm,omitempty"`
	KeySize                 *int       `json:"keySize,omitempty"`
	Organization            *string    `json:"organization,omitempty"`
	OrganizationUnit        *string    `json:"organizationUnit,omitempty"`
	SignatureAlgorithm      *string    `json:"signatureAlgorithm,omitempty"`
	State                   *string    `json:"state,omitempty"`
	SubjectAlternativeNames *[]*string `json:"subjectAlternativeNames,omitempty"`
	ValidDays               *int       `json:"validDays,omitempty"`
}

//NotificationPublisher - A notification publisher plugin instance.
type NotificationPublisher struct {
	Configuration       *PluginConfiguration `json:"configuration,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	ParentRef           *ResourceLink        `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink        `json:"pluginDescriptorRef,omitempty"`
}

//NotificationPublisherDescriptor - A notification publisher plugin descriptor.
type NotificationPublisherDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//NotificationPublisherDescriptors - A collection of notification publisher plugin descriptors.
type NotificationPublisherDescriptors struct {
	Items *[]*NotificationPublisherDescriptor `json:"items,omitempty"`
}

//NotificationPublishers - A collection of notification publisher plugin instances.
type NotificationPublishers struct {
	Items *[]*NotificationPublisher `json:"items,omitempty"`
}

//NotificationPublishersSettings - General notification publisher settings.
type NotificationPublishersSettings struct {
	DefaultNotificationPublisherRef *ResourceLink `json:"defaultNotificationPublisherRef,omitempty"`
}

//NotificationSettings - Settings for notifications relating to licensing and certificate expiration.
type NotificationSettings struct {
	AccountChangesNotificationPublisherRef *ResourceLink                              `json:"accountChangesNotificationPublisherRef,omitempty"`
	CertificateExpirations                 *CertificateExpirationNotificationSettings `json:"certificateExpirations,omitempty"`
	LicenseEvents                          *LicenseEventNotificationSettings          `json:"licenseEvents,omitempty"`
	MetadataNotificationSettings           *MetadataEventNotificationSettings         `json:"metadataNotificationSettings,omitempty"`
	NotifyAdminUserPasswordChanges         *bool                                      `json:"notifyAdminUserPasswordChanges,omitempty"`
}

//OAuthOidcKeysSettings - Setting for OAuth/OpenID Connect signing and decryption key settings.
type OAuthOidcKeysSettings struct {
	P256ActiveCertRef                 *ResourceLink `json:"p256ActiveCertRef,omitempty"`
	P256DecryptionActiveCertRef       *ResourceLink `json:"p256DecryptionActiveCertRef,omitempty"`
	P256DecryptionPreviousCertRef     *ResourceLink `json:"p256DecryptionPreviousCertRef,omitempty"`
	P256DecryptionPublishX5cParameter *bool         `json:"p256DecryptionPublishX5cParameter,omitempty"`
	P256PreviousCertRef               *ResourceLink `json:"p256PreviousCertRef,omitempty"`
	P256PublishX5cParameter           *bool         `json:"p256PublishX5cParameter,omitempty"`
	P384ActiveCertRef                 *ResourceLink `json:"p384ActiveCertRef,omitempty"`
	P384DecryptionActiveCertRef       *ResourceLink `json:"p384DecryptionActiveCertRef,omitempty"`
	P384DecryptionPreviousCertRef     *ResourceLink `json:"p384DecryptionPreviousCertRef,omitempty"`
	P384DecryptionPublishX5cParameter *bool         `json:"p384DecryptionPublishX5cParameter,omitempty"`
	P384PreviousCertRef               *ResourceLink `json:"p384PreviousCertRef,omitempty"`
	P384PublishX5cParameter           *bool         `json:"p384PublishX5cParameter,omitempty"`
	P521ActiveCertRef                 *ResourceLink `json:"p521ActiveCertRef,omitempty"`
	P521DecryptionActiveCertRef       *ResourceLink `json:"p521DecryptionActiveCertRef,omitempty"`
	P521DecryptionPreviousCertRef     *ResourceLink `json:"p521DecryptionPreviousCertRef,omitempty"`
	P521DecryptionPublishX5cParameter *bool         `json:"p521DecryptionPublishX5cParameter,omitempty"`
	P521PreviousCertRef               *ResourceLink `json:"p521PreviousCertRef,omitempty"`
	P521PublishX5cParameter           *bool         `json:"p521PublishX5cParameter,omitempty"`
	RsaActiveCertRef                  *ResourceLink `json:"rsaActiveCertRef,omitempty"`
	RsaDecryptionActiveCertRef        *ResourceLink `json:"rsaDecryptionActiveCertRef,omitempty"`
	RsaDecryptionPreviousCertRef      *ResourceLink `json:"rsaDecryptionPreviousCertRef,omitempty"`
	RsaDecryptionPublishX5cParameter  *bool         `json:"rsaDecryptionPublishX5cParameter,omitempty"`
	RsaPreviousCertRef                *ResourceLink `json:"rsaPreviousCertRef,omitempty"`
	RsaPublishX5cParameter            *bool         `json:"rsaPublishX5cParameter,omitempty"`
	StaticJwksEnabled                 *bool         `json:"staticJwksEnabled,omitempty"`
}

//OAuthRole - This property has been deprecated and is no longer used. OAuth and OpenID Connect are always enabled.
type OAuthRole struct {
	EnableOauth         *bool `json:"enableOauth,omitempty"`
	EnableOpenIdConnect *bool `json:"enableOpenIdConnect,omitempty"`
}

//OIDCClientCredentials - The OpenID Connect Client Credentials settings. This is required for an OIDC Connection.
type OIDCClientCredentials struct {
	ClientId        *string `json:"clientId,omitempty"`
	ClientSecret    *string `json:"clientSecret,omitempty"`
	EncryptedSecret *string `json:"encryptedSecret,omitempty"`
}

//OIDCProviderSettings - The OpenID Provider settings.
type OIDCProviderSettings struct {
	AuthenticationScheme           *string                  `json:"authenticationScheme,omitempty"`
	AuthenticationSigningAlgorithm *string                  `json:"authenticationSigningAlgorithm,omitempty"`
	AuthorizationEndpoint          *string                  `json:"authorizationEndpoint,omitempty"`
	JwksURL                        *string                  `json:"jwksURL,omitempty"`
	LoginType                      *string                  `json:"loginType,omitempty"`
	RequestParameters              *[]*OIDCRequestParameter `json:"requestParameters,omitempty"`
	RequestSigningAlgorithm        *string                  `json:"requestSigningAlgorithm,omitempty"`
	Scopes                         *string                  `json:"scopes,omitempty"`
	TokenEndpoint                  *string                  `json:"tokenEndpoint,omitempty"`
	UserInfoEndpoint               *string                  `json:"userInfoEndpoint,omitempty"`
}

//OIDCRequestParameter - An OIDC custom request parameter.
type OIDCRequestParameter struct {
	ApplicationEndpointOverride *bool   `json:"applicationEndpointOverride,omitempty"`
	Name                        *string `json:"name,omitempty"`
	Value                       *string `json:"value,omitempty"`
}

//OIDCSessionSettings - Settings relating to OpenID Connect session management.
type OIDCSessionSettings struct {
	RevokeUserSessionOnLogout  *bool `json:"revokeUserSessionOnLogout,omitempty"`
	SessionRevocationLifetime  *int  `json:"sessionRevocationLifetime,omitempty"`
	TrackUserSessionsForLogout *bool `json:"trackUserSessionsForLogout,omitempty"`
}

//OcspSettings - OCSP settings.
type OcspSettings struct {
	ActionOnResponderUnavailable *string       `json:"actionOnResponderUnavailable,omitempty"`
	ActionOnStatusUnknown        *string       `json:"actionOnStatusUnknown,omitempty"`
	ActionOnUnsuccessfulResponse *string       `json:"actionOnUnsuccessfulResponse,omitempty"`
	CurrentUpdateGracePeriod     *int          `json:"currentUpdateGracePeriod,omitempty"`
	NextUpdateGracePeriod        *int          `json:"nextUpdateGracePeriod,omitempty"`
	RequesterAddNonce            *bool         `json:"requesterAddNonce,omitempty"`
	ResponderCertReference       *ResourceLink `json:"responderCertReference,omitempty"`
	ResponderTimeout             *int          `json:"responderTimeout,omitempty"`
	ResponderUrl                 *string       `json:"responderUrl,omitempty"`
	ResponseCachePeriod          *int          `json:"responseCachePeriod,omitempty"`
}

//OpenIdConnectAttribute - An attribute for the OpenID Connect returned to OAuth clients.
type OpenIdConnectAttribute struct {
	IncludeInIdToken  *bool   `json:"includeInIdToken,omitempty"`
	IncludeInUserInfo *bool   `json:"includeInUserInfo,omitempty"`
	Name              *string `json:"name,omitempty"`
}

//OpenIdConnectAttributeContract - A set of attributes that will be returned to OAuth clients in response to requests received at the PingFederate UserInfo endpoint.
type OpenIdConnectAttributeContract struct {
	CoreAttributes     *[]*OpenIdConnectAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*OpenIdConnectAttribute `json:"extendedAttributes,omitempty"`
}

//OpenIdConnectPolicies - A collection of OpenID Connect policies.
type OpenIdConnectPolicies struct {
	Items *[]*OpenIdConnectPolicy `json:"items,omitempty"`
}

//OpenIdConnectPolicy - The set of attributes used to configure an OpenID Connect policy.
type OpenIdConnectPolicy struct {
	AccessTokenManagerRef       *ResourceLink                   `json:"accessTokenManagerRef,omitempty"`
	AttributeContract           *OpenIdConnectAttributeContract `json:"attributeContract,omitempty"`
	AttributeMapping            *AttributeMapping               `json:"attributeMapping,omitempty"`
	Id                          *string                         `json:"id,omitempty"`
	IdTokenLifetime             *int                            `json:"idTokenLifetime,omitempty"`
	IncludeSHashInIdToken       *bool                           `json:"includeSHashInIdToken,omitempty"`
	IncludeSriInIdToken         *bool                           `json:"includeSriInIdToken,omitempty"`
	IncludeUserInfoInIdToken    *bool                           `json:"includeUserInfoInIdToken,omitempty"`
	Name                        *string                         `json:"name,omitempty"`
	ReturnIdTokenOnRefreshGrant *bool                           `json:"returnIdTokenOnRefreshGrant,omitempty"`
	ScopeAttributeMappings      map[string]*ParameterValues     `json:"scopeAttributeMappings,omitempty"`
}

//OpenIdConnectSettings - Settings for the OpenID Connect configuration.
type OpenIdConnectSettings struct {
	DefaultPolicyRef *ResourceLink        `json:"defaultPolicyRef,omitempty"`
	SessionSettings  *OIDCSessionSettings `json:"sessionSettings,omitempty"`
}

//OptionValue - An option name and value associated with a selection field.
type OptionValue struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

//OutOfBandAuthAttribute - An attribute for the out of band authenticator plugin instance attribute contract.
type OutOfBandAuthAttribute struct {
	Name *string `json:"name,omitempty"`
}

//OutOfBandAuthAttributeContract - A set of attributes exposed by an out of band authenticator plugin instance.
type OutOfBandAuthAttributeContract struct {
	CoreAttributes     *[]*OutOfBandAuthAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*OutOfBandAuthAttribute `json:"extendedAttributes,omitempty"`
}

//OutOfBandAuthPluginDescriptor - An out of band authenticator plugin descriptor.
type OutOfBandAuthPluginDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//OutOfBandAuthPluginDescriptors
type OutOfBandAuthPluginDescriptors struct {
	Items *[]*OutOfBandAuthPluginDescriptor `json:"items,omitempty"`
}

//OutOfBandAuthenticator - An out of band authenticator plugin instance.
type OutOfBandAuthenticator struct {
	AttributeContract   *OutOfBandAuthAttributeContract `json:"attributeContract,omitempty"`
	Configuration       *PluginConfiguration            `json:"configuration,omitempty"`
	Id                  *string                         `json:"id,omitempty"`
	Name                *string                         `json:"name,omitempty"`
	ParentRef           *ResourceLink                   `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink                   `json:"pluginDescriptorRef,omitempty"`
}

//OutOfBandAuthenticators - A collection of out of band authenticator plugin instances.
type OutOfBandAuthenticators struct {
	Items *[]*OutOfBandAuthenticator `json:"items,omitempty"`
}

//OutboundBackChannelAuth
type OutboundBackChannelAuth struct {
	DigitalSignature     *bool                        `json:"digitalSignature,omitempty"`
	HttpBasicCredentials *UsernamePasswordCredentials `json:"httpBasicCredentials,omitempty"`
	SslAuthKeyPairRef    *ResourceLink                `json:"sslAuthKeyPairRef,omitempty"`
	Type                 *string                      `json:"type,omitempty"`
	ValidatePartnerCert  *bool                        `json:"validatePartnerCert,omitempty"`
}

//OutboundProvision - Outbound Provisioning allows an IdP to create and maintain user accounts at standards-based partner sites using SCIM as well as select-proprietary provisioning partner sites that are protocol-enabled.
type OutboundProvision struct {
	Channels       *[]*Channel     `json:"channels,omitempty"`
	CustomSchema   *Schema         `json:"customSchema,omitempty"`
	TargetSettings *[]*ConfigField `json:"targetSettings,omitempty"`
	Type           *string         `json:"type,omitempty"`
}

//OutboundProvisionDatabase - The settings for database used internally to facilitate outbound provisioning. The database stores state of synchronization between the source data store and the target data store.
type OutboundProvisionDatabase struct {
	DataStoreRef             *ResourceLink `json:"dataStoreRef,omitempty"`
	SynchronizationFrequency *int          `json:"synchronizationFrequency,omitempty"`
}

//P14EKeyPairView - PingOne for Enterprise connection key pair details.
type P14EKeyPairView struct {
	CreationTime              *string   `json:"creationTime,omitempty"`
	CurrentAuthenticationKey  *bool     `json:"currentAuthenticationKey,omitempty"`
	KeyPairView               *CertView `json:"keyPairView,omitempty"`
	PreviousAuthenticationKey *bool     `json:"previousAuthenticationKey,omitempty"`
}

//P14EKeysView - The collection of PingOne for Enterprise connection key pair details.
type P14EKeysView struct {
	KeyPairs *[]*P14EKeyPairView `json:"keyPairs,omitempty"`
}

//PKCS12ExportSettings - Settings for exporting a PKCS12 file from the system.
type PKCS12ExportSettings struct {
	Password *string `json:"password,omitempty"`
}

//PKCS12File - Represents the contents of a PKCS12 file.
type PKCS12File struct {
	CryptoProvider    *string `json:"cryptoProvider,omitempty"`
	EncryptedPassword *string `json:"encryptedPassword,omitempty"`
	FileData          *string `json:"fileData,omitempty"`
	Id                *string `json:"id,omitempty"`
	Password          *string `json:"password,omitempty"`
}

//ParameterValues - Parameter Values.
type ParameterValues struct {
	Values *[]*string `json:"values,omitempty"`
}

//PasswordCredentialValidator - The set of attributes used to configure a password credential validator
type PasswordCredentialValidator struct {
	AttributeContract   *PasswordCredentialValidatorAttributeContract `json:"attributeContract,omitempty"`
	Configuration       *PluginConfiguration                          `json:"configuration,omitempty"`
	Id                  *string                                       `json:"id,omitempty"`
	Name                *string                                       `json:"name,omitempty"`
	ParentRef           *ResourceLink                                 `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink                                 `json:"pluginDescriptorRef,omitempty"`
}

//PasswordCredentialValidatorAttribute - An attribute for the password credential validator attribute contract.
type PasswordCredentialValidatorAttribute struct {
	Name *string `json:"name,omitempty"`
}

//PasswordCredentialValidatorAttributeContract
type PasswordCredentialValidatorAttributeContract struct {
	CoreAttributes     *[]*PasswordCredentialValidatorAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*PasswordCredentialValidatorAttribute `json:"extendedAttributes,omitempty"`
	Inherited          *bool                                    `json:"inherited,omitempty"`
}

//PasswordCredentialValidatorDescriptor - A password credential validator descriptor.
type PasswordCredentialValidatorDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//PasswordCredentialValidatorDescriptors - A collection of password credential validator descriptors.
type PasswordCredentialValidatorDescriptors struct {
	Items *[]*PasswordCredentialValidatorDescriptor `json:"items,omitempty"`
}

//PasswordCredentialValidators - A collection of password credential validators.
type PasswordCredentialValidators struct {
	Items *[]*PasswordCredentialValidator `json:"items,omitempty"`
}

//Pattern
type Pattern struct {
	Flags   *int    `json:"flags,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
}

//PersistentGrantAttribute - A persistent grant contract attribute.
type PersistentGrantAttribute struct {
	Name *string `json:"name,omitempty"`
}

//PersistentGrantContract
type PersistentGrantContract struct {
	CoreAttributes     *[]*PersistentGrantAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*PersistentGrantAttribute `json:"extendedAttributes,omitempty"`
}

//PhoneLocalIdentityField - A phone type field.
type PhoneLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//PingOneConnection - PingOne connection.
type PingOneConnection struct {
	Active                           *bool   `json:"active,omitempty"`
	CreationDate                     *string `json:"creationDate,omitempty"`
	Credential                       *string `json:"credential,omitempty"`
	CredentialId                     *string `json:"credentialId,omitempty"`
	Description                      *string `json:"description,omitempty"`
	EncryptedCredential              *string `json:"encryptedCredential,omitempty"`
	EnvironmentId                    *string `json:"environmentId,omitempty"`
	Id                               *string `json:"id,omitempty"`
	Name                             *string `json:"name,omitempty"`
	OrganizationName                 *string `json:"organizationName,omitempty"`
	PingOneAuthenticationApiEndpoint *string `json:"pingOneAuthenticationApiEndpoint,omitempty"`
	PingOneConnectionId              *string `json:"pingOneConnectionId,omitempty"`
	PingOneManagementApiEndpoint     *string `json:"pingOneManagementApiEndpoint,omitempty"`
	Region                           *string `json:"region,omitempty"`
}

//PingOneConnections - A collection of PingOne connections.
type PingOneConnections struct {
	Items *[]*PingOneConnection `json:"items,omitempty"`
}

//PingOneCredentialStatus - PingOne credential Status
type PingOneCredentialStatus struct {
	PingOneCredentialStatus *string `json:"pingOneCredentialStatus,omitempty"`
}

//PingOneEnvironment
type PingOneEnvironment struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

//PingOneEnvironments - A collection of PingOne Environments.
type PingOneEnvironments struct {
	Items *[]*PingOneEnvironment `json:"items,omitempty"`
}

//PingOneForEnterpriseSettings - PingOne for Enterprise Settings
type PingOneForEnterpriseSettings struct {
	CompanyName                      *string       `json:"companyName,omitempty"`
	ConnectedToPingOneForEnterprise  *bool         `json:"connectedToPingOneForEnterprise,omitempty"`
	CurrentAuthnKeyCreationTime      *string       `json:"currentAuthnKeyCreationTime,omitempty"`
	EnableAdminConsoleSso            *bool         `json:"enableAdminConsoleSso,omitempty"`
	EnableMonitoring                 *bool         `json:"enableMonitoring,omitempty"`
	IdentityRepositoryUpdateRequired *bool         `json:"identityRepositoryUpdateRequired,omitempty"`
	PingOneSsoConnection             *ResourceLink `json:"pingOneSsoConnection,omitempty"`
	PreviousAuthnKeyCreationTime     *string       `json:"previousAuthnKeyCreationTime,omitempty"`
}

//PluginConfigDescriptor - Defines the configuration fields available for a plugin.
type PluginConfigDescriptor struct {
	ActionDescriptors *[]*ActionDescriptor `json:"actionDescriptors,omitempty"`
	Description       *string              `json:"description,omitempty"`
	Fields            *[]*FieldDescriptor  `json:"fields,omitempty"`
	Tables            *[]*TableDescriptor  `json:"tables,omitempty"`
}

//PluginConfiguration - Configuration settings for a plugin instance.
type PluginConfiguration struct {
	Fields *[]*ConfigField `json:"fields,omitempty"`
	Tables *[]*ConfigTable `json:"tables,omitempty"`
}

//PluginDescriptor - Defines a plugin type, including available configuration parameters.
type PluginDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//PluginInstance - A plugin instance.
type PluginInstance struct {
	Configuration       *PluginConfiguration `json:"configuration,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	ParentRef           *ResourceLink        `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink        `json:"pluginDescriptorRef,omitempty"`
}

//PolicyAction - An authentication policy selection action.
type PolicyAction struct {
	ApcMappingPolicyAction
	LocalIdentityMappingPolicyAction
	AuthnSelectorPolicyAction
	AuthnSourcePolicyAction
	ContinuePolicyAction
	RestartPolicyAction
	DonePolicyAction
	FragmentPolicyAction
	Context *string `json:"context,omitempty"`
	Type    *string `json:"type,omitempty"`
}

//ProcessorPolicyToGeneratorMapping - A Token Exchange Processor policy to Token Generator Mapping.
type ProcessorPolicyToGeneratorMapping struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                 *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Id                               *string                               `json:"id,omitempty"`
	IssuanceCriteria                 *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	LicenseConnectionGroupAssignment *string                               `json:"licenseConnectionGroupAssignment,omitempty"`
	SourceId                         *string                               `json:"sourceId,omitempty"`
	TargetId                         *string                               `json:"targetId,omitempty"`
}

//ProcessorPolicyToGeneratorMappings
type ProcessorPolicyToGeneratorMappings struct {
	Items *[]*ProcessorPolicyToGeneratorMapping `json:"items,omitempty"`
}

//ProfileConfig - A local identity profile management configuration.
type ProfileConfig struct {
	DeleteIdentityEnabled *bool   `json:"deleteIdentityEnabled,omitempty"`
	TemplateName          *string `json:"templateName,omitempty"`
}

//ProtocolMessageCustomization - The message customization that will be executed on outgoing PingFederate messages.
type ProtocolMessageCustomization struct {
	ContextName       *string `json:"contextName,omitempty"`
	MessageExpression *string `json:"messageExpression,omitempty"`
}

//ProxySettings - Proxy settings.
type ProxySettings struct {
	Host *string `json:"host,omitempty"`
	Port *int    `json:"port,omitempty"`
}

//RadioGroupFieldDescriptor - A selection-type field intended to be rendered as a group of radio buttons in a UI.
type RadioGroupFieldDescriptor struct {
	Advanced     *bool           `json:"advanced,omitempty"`
	DefaultValue *string         `json:"defaultValue,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Label        *string         `json:"label,omitempty"`
	Name         *string         `json:"name,omitempty"`
	OptionValues *[]*OptionValue `json:"optionValues,omitempty"`
	Required     *bool           `json:"required,omitempty"`
	Type         *string         `json:"type,omitempty"`
}

//RedirectValidationLocalSettings - Settings for local redirect validation.
type RedirectValidationLocalSettings struct {
	EnableInErrorResourceValidation               *bool                                        `json:"enableInErrorResourceValidation,omitempty"`
	EnableTargetResourceValidationForIdpDiscovery *bool                                        `json:"enableTargetResourceValidationForIdpDiscovery,omitempty"`
	EnableTargetResourceValidationForSLO          *bool                                        `json:"enableTargetResourceValidationForSLO,omitempty"`
	EnableTargetResourceValidationForSSO          *bool                                        `json:"enableTargetResourceValidationForSSO,omitempty"`
	WhiteList                                     *[]*RedirectValidationSettingsWhitelistEntry `json:"whiteList,omitempty"`
}

//RedirectValidationPartnerSettings - Settings for redirection at a partner site.
type RedirectValidationPartnerSettings struct {
	EnableWreplyValidationSLO *bool `json:"enableWreplyValidationSLO,omitempty"`
}

//RedirectValidationSettings - Settings for redirect validation for SSO, SLO and IdP discovery.
type RedirectValidationSettings struct {
	RedirectValidationLocalSettings   *RedirectValidationLocalSettings   `json:"redirectValidationLocalSettings,omitempty"`
	RedirectValidationPartnerSettings *RedirectValidationPartnerSettings `json:"redirectValidationPartnerSettings,omitempty"`
}

//RedirectValidationSettingsWhitelistEntry - Whitelist entry for valid target resource.
type RedirectValidationSettingsWhitelistEntry struct {
	AllowQueryAndFragment *bool   `json:"allowQueryAndFragment,omitempty"`
	IdpDiscovery          *bool   `json:"idpDiscovery,omitempty"`
	InErrorResource       *bool   `json:"inErrorResource,omitempty"`
	RequireHttps          *bool   `json:"requireHttps,omitempty"`
	TargetResourceSLO     *bool   `json:"targetResourceSLO,omitempty"`
	TargetResourceSSO     *bool   `json:"targetResourceSSO,omitempty"`
	ValidDomain           *string `json:"validDomain,omitempty"`
	ValidPath             *string `json:"validPath,omitempty"`
}

//RegistrationConfig - A local identity profile registration configuration.
type RegistrationConfig struct {
	CaptchaEnabled                      *bool         `json:"captchaEnabled,omitempty"`
	CreateAuthnSessionAfterRegistration *bool         `json:"createAuthnSessionAfterRegistration,omitempty"`
	ExecuteWorkflow                     *string       `json:"executeWorkflow,omitempty"`
	RegistrationWorkflow                *ResourceLink `json:"registrationWorkflow,omitempty"`
	TemplateName                        *string       `json:"templateName,omitempty"`
	ThisIsMyDeviceEnabled               *bool         `json:"thisIsMyDeviceEnabled,omitempty"`
	UsernameField                       *string       `json:"usernameField,omitempty"`
}

//RequestPolicies - A collection of CIBA request policies.
type RequestPolicies struct {
	Items *[]*RequestPolicy `json:"items,omitempty"`
}

//RequestPolicy - The set of attributes used to configure a CIBA request policy.
type RequestPolicy struct {
	AllowUnsignedLoginHintToken      *bool                               `json:"allowUnsignedLoginHintToken,omitempty"`
	AlternativeLoginHintTokenIssuers *[]*AlternativeLoginHintTokenIssuer `json:"alternativeLoginHintTokenIssuers,omitempty"`
	AuthenticatorRef                 *ResourceLink                       `json:"authenticatorRef,omitempty"`
	Id                               *string                             `json:"id,omitempty"`
	IdentityHintContract             *IdentityHintContract               `json:"identityHintContract,omitempty"`
	IdentityHintContractFulfillment  *AttributeMapping                   `json:"identityHintContractFulfillment,omitempty"`
	IdentityHintMapping              *AttributeMapping                   `json:"identityHintMapping,omitempty"`
	Name                             *string                             `json:"name,omitempty"`
	RequireTokenForIdentityHint      *bool                               `json:"requireTokenForIdentityHint,omitempty"`
	TransactionLifetime              *int                                `json:"transactionLifetime,omitempty"`
	UserCodePcvRef                   *ResourceLink                       `json:"userCodePcvRef,omitempty"`
}

//ResourceCategoryInfo - A model containing information on a category of resource in the administrative API.
type ResourceCategoryInfo struct {
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

//ResourceLink - A reference to a resource.
type ResourceLink struct {
	Id       *string `json:"id,omitempty"`
	Location *string `json:"location,omitempty"`
}

//ResourceOwnerCredentialsMapping - The OAuth Resource Owner Credentials Mapping.
type ResourceOwnerCredentialsMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Id                           *string                               `json:"id,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	PasswordValidatorRef         *ResourceLink                         `json:"passwordValidatorRef,omitempty"`
}

//ResourceOwnerCredentialsMappings - A collection of OAuth Resource Owner Credentials mapping items.
type ResourceOwnerCredentialsMappings struct {
	Items *[]*ResourceOwnerCredentialsMapping `json:"items,omitempty"`
}

//ResourceUsage - An API model representing a reference to an API resource.
type ResourceUsage struct {
	CategoryId *string       `json:"categoryId,omitempty"`
	Id         *string       `json:"id,omitempty"`
	Name       *string       `json:"name,omitempty"`
	Ref        *ResourceLink `json:"ref,omitempty"`
	Type       *string       `json:"type,omitempty"`
}

//ResourceUsages - A collection of resource usages.
type ResourceUsages struct {
	Categories *[]*ResourceCategoryInfo `json:"categories,omitempty"`
	Items      *[]*ResourceUsage        `json:"items,omitempty"`
}

//RestartPolicyAction - The restart selection action.
type RestartPolicyAction struct {
	Context *string `json:"context,omitempty"`
	Type    *string `json:"type,omitempty"`
}

//RolesAndProtocols - This property has been deprecated and is no longer used. All Roles and protocols are always enabled.
type RolesAndProtocols struct {
	EnableIdpDiscovery *bool      `json:"enableIdpDiscovery,omitempty"`
	IdpRole            *IdpRole   `json:"idpRole,omitempty"`
	OauthRole          *OAuthRole `json:"oauthRole,omitempty"`
	SpRole             *SpRole    `json:"spRole,omitempty"`
}

//SAML20Profile - SAML 2.0 Profile.
type SAML20Profile struct {
	SpSAML20Profile
	Enable            *bool `json:"enable,omitempty"`
	EnableAutoConnect *bool `json:"enableAutoConnect,omitempty"`
}

//SaasAttributeMapping - Settings to map the source record attributes to target attributes.
type SaasAttributeMapping struct {
	FieldName     *string                 `json:"fieldName,omitempty"`
	SaasFieldInfo *SaasFieldConfiguration `json:"saasFieldInfo,omitempty"`
}

//SaasFieldConfiguration - The settings that represent how attribute values from source data store will be mapped into Fields specified by the service provider.
type SaasFieldConfiguration struct {
	AttributeNames *[]*string `json:"attributeNames,omitempty"`
	CharacterCase  *string    `json:"characterCase,omitempty"`
	CreateOnly     *bool      `json:"createOnly,omitempty"`
	DefaultValue   *string    `json:"defaultValue,omitempty"`
	Expression     *string    `json:"expression,omitempty"`
	Masked         *bool      `json:"masked,omitempty"`
	Parser         *string    `json:"parser,omitempty"`
	Trim           *bool      `json:"trim,omitempty"`
}

//SaasPluginDescriptor - A SaaS Plugin.
type SaasPluginDescriptor struct {
	ConfigDescriptor               *PluginConfigDescriptor           `json:"configDescriptor,omitempty"`
	Description                    *string                           `json:"description,omitempty"`
	Id                             *string                           `json:"id,omitempty"`
	SaasPluginFieldInfoDescriptors *[]*SaasPluginFieldInfoDescriptor `json:"saasPluginFieldInfoDescriptors,omitempty"`
}

//SaasPluginDescriptors - A collection of SaaS plugins.
type SaasPluginDescriptors struct {
	Items *[]*SaasPluginDescriptor `json:"items,omitempty"`
}

//SaasPluginFieldInfoDescriptor - A Saas Plugin Field configuration.
type SaasPluginFieldInfoDescriptor struct {
	AttributeGroup       *bool                     `json:"attributeGroup,omitempty"`
	Code                 *string                   `json:"code,omitempty"`
	DefaultValue         *string                   `json:"defaultValue,omitempty"`
	DsLdapMap            *bool                     `json:"dsLdapMap,omitempty"`
	Label                *string                   `json:"label,omitempty"`
	MaxLength            *int                      `json:"maxLength,omitempty"`
	MinLength            *int                      `json:"minLength,omitempty"`
	MultiValue           *bool                     `json:"multiValue,omitempty"`
	Notes                *[]*string                `json:"notes,omitempty"`
	Options              *[]*SaasPluginFieldOption `json:"options,omitempty"`
	Pattern              *Pattern                  `json:"pattern,omitempty"`
	PersistForMembership *bool                     `json:"persistForMembership,omitempty"`
	Required             *bool                     `json:"required,omitempty"`
	Unique               *bool                     `json:"unique,omitempty"`
}

//SaasPluginFieldOption - A plugin configuration field value.
type SaasPluginFieldOption struct {
	Code  *string `json:"code,omitempty"`
	Label *string `json:"label,omitempty"`
}

//Schema - Custom SCIM Attributes configuration.
type Schema struct {
	Attributes *[]*SchemaAttribute `json:"attributes,omitempty"`
	Namespace  *string             `json:"namespace,omitempty"`
}

//SchemaAttribute - A custom SCIM attribute.
type SchemaAttribute struct {
	MultiValued   *bool      `json:"multiValued,omitempty"`
	Name          *string    `json:"name,omitempty"`
	SubAttributes *[]*string `json:"subAttributes,omitempty"`
	Types         *[]*string `json:"types,omitempty"`
}

//ScopeEntry - A scope name and its description.
type ScopeEntry struct {
	Description *string `json:"description,omitempty"`
	Dynamic     *bool   `json:"dynamic,omitempty"`
	Name        *string `json:"name,omitempty"`
}

//ScopeGroupEntry - A scope group name and its description.
type ScopeGroupEntry struct {
	Description *string    `json:"description,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Scopes      *[]*string `json:"scopes,omitempty"`
}

//SelectFieldDescriptor - A selection field that is intended to be rendered as a dropdown list of options.
type SelectFieldDescriptor struct {
	Advanced     *bool           `json:"advanced,omitempty"`
	DefaultValue *string         `json:"defaultValue,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Label        *string         `json:"label,omitempty"`
	Name         *string         `json:"name,omitempty"`
	OptionValues *[]*OptionValue `json:"optionValues,omitempty"`
	Required     *bool           `json:"required,omitempty"`
	Type         *string         `json:"type,omitempty"`
}

//ServerSettings - Server configuration settings.
type ServerSettings struct {
	CaptchaSettings   *CaptchaSettings      `json:"captchaSettings,omitempty"`
	ContactInfo       *ContactInfo          `json:"contactInfo,omitempty"`
	EmailServer       *EmailServerSettings  `json:"emailServer,omitempty"`
	FederationInfo    *FederationInfo       `json:"federationInfo,omitempty"`
	Notifications     *NotificationSettings `json:"notifications,omitempty"`
	RolesAndProtocols *RolesAndProtocols    `json:"rolesAndProtocols,omitempty"`
}

//ServiceAssociation - A model representing an association between a PingFederate component (typically a plugin) and a list of PingOne services.
type ServiceAssociation struct {
	ComponentName *string    `json:"componentName,omitempty"`
	Configured    *bool      `json:"configured,omitempty"`
	ServiceNames  *[]*string `json:"serviceNames,omitempty"`
}

//ServiceAssociations - A list of installed components that consume PingOne services.
type ServiceAssociations struct {
	Items *[]*ServiceAssociation `json:"items,omitempty"`
}

//ServiceAuthentication - Service Authentication Settings.
type ServiceAuthentication struct {
	AttributeQuery       *ServiceModel `json:"attributeQuery,omitempty"`
	ConnectionManagement *ServiceModel `json:"connectionManagement,omitempty"`
	Jmx                  *ServiceModel `json:"jmx,omitempty"`
	SsoDirectoryService  *ServiceModel `json:"ssoDirectoryService,omitempty"`
}

//ServiceModel - Service Model.
type ServiceModel struct {
	EncryptedSharedSecret *string `json:"encryptedSharedSecret,omitempty"`
	Id                    *string `json:"id,omitempty"`
	SharedSecret          *string `json:"sharedSecret,omitempty"`
}

//SessionSettings - General settings related to session management.
type SessionSettings struct {
	RevokeUserSessionOnLogout     *bool `json:"revokeUserSessionOnLogout,omitempty"`
	SessionRevocationLifetime     *int  `json:"sessionRevocationLifetime,omitempty"`
	TrackAdapterSessionsForLogout *bool `json:"trackAdapterSessionsForLogout,omitempty"`
}

//SessionValidationSettings - Session validation settings for an access token management plugin instance.
type SessionValidationSettings struct {
	CheckSessionRevocationStatus *bool `json:"checkSessionRevocationStatus,omitempty"`
	CheckValidAuthnSession       *bool `json:"checkValidAuthnSession,omitempty"`
	IncludeSessionId             *bool `json:"includeSessionId,omitempty"`
	Inherited                    *bool `json:"inherited,omitempty"`
	UpdateAuthnSessionActivity   *bool `json:"updateAuthnSessionActivity,omitempty"`
}

//SigningSettings - Settings related to signing messages sent to this partner.
type SigningSettings struct {
	Algorithm                *string       `json:"algorithm,omitempty"`
	IncludeCertInSignature   *bool         `json:"includeCertInSignature,omitempty"`
	IncludeRawKeyInSignature *bool         `json:"includeRawKeyInSignature,omitempty"`
	SigningKeyPairRef        *ResourceLink `json:"signingKeyPairRef,omitempty"`
}

//SloServiceEndpoint - Where SLO logout messages are sent. Only applicable for SAML 2.0.
type SloServiceEndpoint struct {
	Binding     *string `json:"binding,omitempty"`
	ResponseUrl *string `json:"responseUrl,omitempty"`
	Url         *string `json:"url,omitempty"`
}

//SourceTypeIdKey - A key that is meant to reference a source from which an attribute can be retrieved. This model is usually paired with a value which, depending on the SourceType, can be a hardcoded value or a reference to an attribute name specific to that SourceType. Not all values are applicable - a validation error will be returned for incorrect values.<br>For each SourceType, the value should be:<br>ACCOUNT_LINK - If account linking was enabled for the browser SSO, the value must be 'Local User ID', unless it has been overridden in PingFederate's server configuration.<br>ADAPTER - The value is one of the attributes of the IdP Adapter.<br>ASSERTION - The value is one of the attributes coming from the SAML assertion.<br>AUTHENTICATION_POLICY_CONTRACT - The value is one of the attributes coming from an authentication policy contract.<br>LOCAL_IDENTITY_PROFILE - The value is one of the fields coming from a local identity profile.<br>CONTEXT - The value must be one of the following ['TargetResource' or 'OAuthScopes' or 'ClientId' or 'AuthenticationCtx' or 'ClientIp' or 'Locale' or 'StsBasicAuthUsername' or 'StsSSLClientCertSubjectDN' or 'StsSSLClientCertChain' or 'VirtualServerId' or 'AuthenticatingAuthority' or 'DefaultPersistentGrantLifetime']<br>CLAIMS - Attributes provided by the OIDC Provider.<br>CUSTOM_DATA_STORE - The value is one of the attributes returned by this custom data store.<br>EXPRESSION - The value is an OGNL expression.<br>EXTENDED_CLIENT_METADATA - The value is from an OAuth extended client metadata parameter. This source type is deprecated and has been replaced by EXTENDED_PROPERTIES.<br>EXTENDED_PROPERTIES - The value is from an OAuth Client's extended property.<br>IDP_CONNECTION - The value is one of the attributes passed in by the IdP connection.<br>JDBC_DATA_STORE - The value is one of the column names returned from the JDBC attribute source.<br>LDAP_DATA_STORE - The value is one of the LDAP attributes supported by your LDAP data store.<br>MAPPED_ATTRIBUTES - The value is the name of one of the mapped attributes that is defined in the associated attribute mapping.<br>OAUTH_PERSISTENT_GRANT - The value is one of the attributes from the persistent grant.<br>PASSWORD_CREDENTIAL_VALIDATOR - The value is one of the attributes of the PCV.<br>NO_MAPPING - A placeholder value to indicate that an attribute currently has no mapped source.TEXT - A hardcoded value that is used to populate the corresponding attribute.<br>TOKEN - The value is one of the token attributes.<br>REQUEST - The value is from the request context such as the CIBA identity hint contract or the request contract for Ws-Trust.<br>TRACKED_HTTP_PARAMS - The value is from the original request parameters.<br>SUBJECT_TOKEN - The value is one of the OAuth 2.0 Token exchange subject_token attributes.<br>ACTOR_TOKEN - The value is one of the OAuth 2.0 Token exchange actor_token attributes.<br>TOKEN_EXCHANGE_PROCESSOR_POLICY - The value is one of the attributes coming from a Token Exchange Processor policy.<br>FRAGMENT - The value is one of the attributes coming from an authentication policy fragment.<br>INPUTS - The value is one of the attributes coming from an attribute defined in the input authentication policy contract for an authentication policy fragment.
type SourceTypeIdKey struct {
	Id   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

//SpAdapter - An SP adapter instance.
type SpAdapter struct {
	AttributeContract     *SpAdapterAttributeContract     `json:"attributeContract,omitempty"`
	Configuration         *PluginConfiguration            `json:"configuration,omitempty"`
	Id                    *string                         `json:"id,omitempty"`
	Name                  *string                         `json:"name,omitempty"`
	ParentRef             *ResourceLink                   `json:"parentRef,omitempty"`
	PluginDescriptorRef   *ResourceLink                   `json:"pluginDescriptorRef,omitempty"`
	TargetApplicationInfo *SpAdapterTargetApplicationInfo `json:"targetApplicationInfo,omitempty"`
}

//SpAdapterAttribute - An attribute for the SP adapter attribute contract.
type SpAdapterAttribute struct {
	Name *string `json:"name,omitempty"`
}

//SpAdapterAttributeContract - A set of attributes exposed by an SP adapter.
type SpAdapterAttributeContract struct {
	CoreAttributes     *[]*SpAdapterAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*SpAdapterAttribute `json:"extendedAttributes,omitempty"`
	Inherited          *bool                  `json:"inherited,omitempty"`
}

//SpAdapterDescriptor - An SP adapter descriptor.
type SpAdapterDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//SpAdapterDescriptors - A collection of SP adapter descriptors.
type SpAdapterDescriptors struct {
	Items *[]*SpAdapterDescriptor `json:"items,omitempty"`
}

//SpAdapterMapping - A mapping to a SP adapter.
type SpAdapterMapping struct {
	AdapterOverrideSettings      *SpAdapter                            `json:"adapterOverrideSettings,omitempty"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictVirtualEntityIds     *bool                                 `json:"restrictVirtualEntityIds,omitempty"`
	RestrictedVirtualEntityIds   *[]*string                            `json:"restrictedVirtualEntityIds,omitempty"`
	SpAdapterRef                 *ResourceLink                         `json:"spAdapterRef,omitempty"`
}

//SpAdapterTargetApplicationInfo - Target Application Information exposed by an SP adapter.
type SpAdapterTargetApplicationInfo struct {
	ApplicationIconUrl *string `json:"applicationIconUrl,omitempty"`
	ApplicationName    *string `json:"applicationName,omitempty"`
	Inherited          *bool   `json:"inherited,omitempty"`
}

//SpAdapterUrlMapping - SP Adapter URL Mapping
type SpAdapterUrlMapping struct {
	AdapterRef *ResourceLink `json:"adapterRef,omitempty"`
	Url        *string       `json:"url,omitempty"`
}

//SpAdapterUrlMappings
type SpAdapterUrlMappings struct {
	Items *[]*SpAdapterUrlMapping `json:"items,omitempty"`
}

//SpAdapters - A collection of SP adapters.
type SpAdapters struct {
	Items *[]*SpAdapter `json:"items,omitempty"`
}

//SpAttributeQuery - The attribute query profile supports SPs in requesting user attributes.
type SpAttributeQuery struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	Attributes                   *[]*string                            `json:"attributes,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	Policy                       *SpAttributeQueryPolicy               `json:"policy,omitempty"`
}

//SpAttributeQueryPolicy - The attribute query profile's security policy.
type SpAttributeQueryPolicy struct {
	EncryptAssertion            *bool `json:"encryptAssertion,omitempty"`
	RequireEncryptedNameId      *bool `json:"requireEncryptedNameId,omitempty"`
	RequireSignedAttributeQuery *bool `json:"requireSignedAttributeQuery,omitempty"`
	SignAssertion               *bool `json:"signAssertion,omitempty"`
	SignResponse                *bool `json:"signResponse,omitempty"`
}

//SpBrowserSso - The SAML settings used to enable secure browser-based SSO to resources at your partner's site.
type SpBrowserSso struct {
	AdapterMappings                               *[]*IdpAdapterAssertionMapping                   `json:"adapterMappings,omitempty"`
	AlwaysSignArtifactResponse                    *bool                                            `json:"alwaysSignArtifactResponse,omitempty"`
	Artifact                                      *ArtifactSettings                                `json:"artifact,omitempty"`
	AssertionLifetime                             *AssertionLifetime                               `json:"assertionLifetime,omitempty"`
	AttributeContract                             *SpBrowserSsoAttributeContract                   `json:"attributeContract,omitempty"`
	AuthenticationPolicyContractAssertionMappings *[]*AuthenticationPolicyContractAssertionMapping `json:"authenticationPolicyContractAssertionMappings,omitempty"`
	DefaultTargetUrl                              *string                                          `json:"defaultTargetUrl,omitempty"`
	EnabledProfiles                               *[]*string                                       `json:"enabledProfiles,omitempty"`
	EncryptionPolicy                              *EncryptionPolicy                                `json:"encryptionPolicy,omitempty"`
	IncomingBindings                              *[]*string                                       `json:"incomingBindings,omitempty"`
	MessageCustomizations                         *[]*ProtocolMessageCustomization                 `json:"messageCustomizations,omitempty"`
	Protocol                                      *string                                          `json:"protocol,omitempty"`
	RequireSignedAuthnRequests                    *bool                                            `json:"requireSignedAuthnRequests,omitempty"`
	SignAssertions                                *bool                                            `json:"signAssertions,omitempty"`
	SignResponseAsRequired                        *bool                                            `json:"signResponseAsRequired,omitempty"`
	SloServiceEndpoints                           *[]*SloServiceEndpoint                           `json:"sloServiceEndpoints,omitempty"`
	SpSamlIdentityMapping                         *string                                          `json:"spSamlIdentityMapping,omitempty"`
	SpWsFedIdentityMapping                        *string                                          `json:"spWsFedIdentityMapping,omitempty"`
	SsoServiceEndpoints                           *[]*SpSsoServiceEndpoint                         `json:"ssoServiceEndpoints,omitempty"`
	UrlWhitelistEntries                           *[]*UrlWhitelistEntry                            `json:"urlWhitelistEntries,omitempty"`
	WsFedTokenType                                *string                                          `json:"wsFedTokenType,omitempty"`
	WsTrustVersion                                *string                                          `json:"wsTrustVersion,omitempty"`
}

//SpBrowserSsoAttribute - An attribute for the SP Browser SSO attribute contract.
type SpBrowserSsoAttribute struct {
	Name       *string `json:"name,omitempty"`
	NameFormat *string `json:"nameFormat,omitempty"`
}

//SpBrowserSsoAttributeContract - A set of user attributes that the IdP sends in the SAML assertion.
type SpBrowserSsoAttributeContract struct {
	CoreAttributes     *[]*SpBrowserSsoAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*SpBrowserSsoAttribute `json:"extendedAttributes,omitempty"`
}

//SpConnection - The set of attributes used to configure an SP connection.
type SpConnection struct {
	Active                                 *bool                                   `json:"active,omitempty"`
	AdditionalAllowedEntitiesConfiguration *AdditionalAllowedEntitiesConfiguration `json:"additionalAllowedEntitiesConfiguration,omitempty"`
	ApplicationIconUrl                     *string                                 `json:"applicationIconUrl,omitempty"`
	ApplicationName                        *string                                 `json:"applicationName,omitempty"`
	AttributeQuery                         *SpAttributeQuery                       `json:"attributeQuery,omitempty"`
	BaseUrl                                *string                                 `json:"baseUrl,omitempty"`
	ConnectionTargetType                   *string                                 `json:"connectionTargetType,omitempty"`
	ContactInfo                            *ContactInfo                            `json:"contactInfo,omitempty"`
	Credentials                            *ConnectionCredentials                  `json:"credentials,omitempty"`
	DefaultVirtualEntityId                 *string                                 `json:"defaultVirtualEntityId,omitempty"`
	EntityId                               *string                                 `json:"entityId,omitempty"`
	ExtendedProperties                     map[string]*ParameterValues             `json:"extendedProperties,omitempty"`
	Id                                     *string                                 `json:"id,omitempty"`
	LicenseConnectionGroup                 *string                                 `json:"licenseConnectionGroup,omitempty"`
	LoggingMode                            *string                                 `json:"loggingMode,omitempty"`
	MetadataReloadSettings                 *ConnectionMetadataUrl                  `json:"metadataReloadSettings,omitempty"`
	Name                                   *string                                 `json:"name,omitempty"`
	OutboundProvision                      *OutboundProvision                      `json:"outboundProvision,omitempty"`
	SpBrowserSso                           *SpBrowserSso                           `json:"spBrowserSso,omitempty"`
	Type                                   *string                                 `json:"type,omitempty"`
	VirtualEntityIds                       *[]*string                              `json:"virtualEntityIds,omitempty"`
	WsTrust                                *SpWsTrust                              `json:"wsTrust,omitempty"`
}

//SpConnections - A collection of SP connections.
type SpConnections struct {
	Items *[]*SpConnection `json:"items,omitempty"`
}

//SpDefaultUrls - SP Default URLs.
type SpDefaultUrls struct {
	ConfirmSlo    *bool   `json:"confirmSlo,omitempty"`
	SloSuccessUrl *string `json:"sloSuccessUrl,omitempty"`
	SsoSuccessUrl *string `json:"ssoSuccessUrl,omitempty"`
}

//SpRole - This property has been deprecated and is no longer used. All Roles and protocols are always enabled.
type SpRole struct {
	Enable                    *bool            `json:"enable,omitempty"`
	EnableInboundProvisioning *bool            `json:"enableInboundProvisioning,omitempty"`
	EnableOpenIDConnect       *bool            `json:"enableOpenIDConnect,omitempty"`
	EnableSaml10              *bool            `json:"enableSaml10,omitempty"`
	EnableSaml11              *bool            `json:"enableSaml11,omitempty"`
	EnableWsFed               *bool            `json:"enableWsFed,omitempty"`
	EnableWsTrust             *bool            `json:"enableWsTrust,omitempty"`
	Saml20Profile             *SpSAML20Profile `json:"saml20Profile,omitempty"`
}

//SpSAML20Profile - SP SAML 2.0 Profile.
type SpSAML20Profile struct {
	Enable            *bool `json:"enable,omitempty"`
	EnableAutoConnect *bool `json:"enableAutoConnect,omitempty"`
	EnableXASP        *bool `json:"enableXASP,omitempty"`
}

//SpSsoServiceEndpoint - The settings that define a service endpoint to a SP SSO service.
type SpSsoServiceEndpoint struct {
	Binding   *string `json:"binding,omitempty"`
	Index     *int    `json:"index,omitempty"`
	IsDefault *bool   `json:"isDefault,omitempty"`
	Url       *string `json:"url,omitempty"`
}

//SpTokenGeneratorMapping - The SP Token Generator Mapping.
type SpTokenGeneratorMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	DefaultMapping               *bool                                 `json:"defaultMapping,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	RestrictedVirtualEntityIds   *[]*string                            `json:"restrictedVirtualEntityIds,omitempty"`
	SpTokenGeneratorRef          *ResourceLink                         `json:"spTokenGeneratorRef,omitempty"`
}

//SpUrlMapping - SP URL mapping
type SpUrlMapping struct {
	Ref  *ResourceLink `json:"ref,omitempty"`
	Type *string       `json:"type,omitempty"`
	Url  *string       `json:"url,omitempty"`
}

//SpUrlMappings
type SpUrlMappings struct {
	Items *[]*SpUrlMapping `json:"items,omitempty"`
}

//SpWsTrust - Ws-Trust STS provides security-token validation and creation to extend SSO access to identity-enabled Web Services
type SpWsTrust struct {
	AbortIfNotFulfilledFromRequest *bool                            `json:"abortIfNotFulfilledFromRequest,omitempty"`
	AttributeContract              *SpWsTrustAttributeContract      `json:"attributeContract,omitempty"`
	DefaultTokenType               *string                          `json:"defaultTokenType,omitempty"`
	EncryptSaml2Assertion          *bool                            `json:"encryptSaml2Assertion,omitempty"`
	GenerateKey                    *bool                            `json:"generateKey,omitempty"`
	MessageCustomizations          *[]*ProtocolMessageCustomization `json:"messageCustomizations,omitempty"`
	MinutesAfter                   *int                             `json:"minutesAfter,omitempty"`
	MinutesBefore                  *int                             `json:"minutesBefore,omitempty"`
	OAuthAssertionProfiles         *bool                            `json:"oAuthAssertionProfiles,omitempty"`
	PartnerServiceIds              *[]*string                       `json:"partnerServiceIds,omitempty"`
	RequestContractRef             *ResourceLink                    `json:"requestContractRef,omitempty"`
	TokenProcessorMappings         *[]*IdpTokenProcessorMapping     `json:"tokenProcessorMappings,omitempty"`
}

//SpWsTrustAttribute - An attribute for the Ws-Trust attribute contract.
type SpWsTrustAttribute struct {
	Name      *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

//SpWsTrustAttributeContract - A set of user attributes that this server will send in the token.
type SpWsTrustAttributeContract struct {
	CoreAttributes     *[]*SpWsTrustAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*SpWsTrustAttribute `json:"extendedAttributes,omitempty"`
}

//SslServerSettings - Settings for the SSL Server certificate configuration.
type SslServerSettings struct {
	ActiveAdminConsoleCerts  *[]*ResourceLink `json:"activeAdminConsoleCerts,omitempty"`
	ActiveRuntimeServerCerts *[]*ResourceLink `json:"activeRuntimeServerCerts,omitempty"`
	AdminConsoleCertRef      *ResourceLink    `json:"adminConsoleCertRef,omitempty"`
	RuntimeServerCertRef     *ResourceLink    `json:"runtimeServerCertRef,omitempty"`
}

//SsoOAuthMapping - IdP Browser SSO OAuth Attribute Mapping
type SsoOAuthMapping struct {
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
}

//StsRequestParametersContract - A Security Token Service request parameter contract.
type StsRequestParametersContract struct {
	Id         *string    `json:"id,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Parameters *[]*string `json:"parameters,omitempty"`
}

//StsRequestParametersContracts - A Collection of STS Request Parameters Contracts
type StsRequestParametersContracts struct {
	Items *[]*StsRequestParametersContract `json:"items,omitempty"`
}

//SystemKey - A system key.
type SystemKey struct {
	CreationDate     *string `json:"creationDate,omitempty"`
	EncryptedKeyData *string `json:"encryptedKeyData,omitempty"`
	KeyData          *string `json:"keyData,omitempty"`
}

//SystemKeys - Secrets that are used in cryptographic operations to generate and consume internal tokens
type SystemKeys struct {
	Current  *SystemKey `json:"current,omitempty"`
	Pending  *SystemKey `json:"pending,omitempty"`
	Previous *SystemKey `json:"previous,omitempty"`
}

//TableDescriptor - Defines a plugin configuration table.
type TableDescriptor struct {
	Columns           *[]*FieldDescriptor `json:"columns,omitempty"`
	Description       *string             `json:"description,omitempty"`
	Label             *string             `json:"label,omitempty"`
	Name              *string             `json:"name,omitempty"`
	RequireDefaultRow *bool               `json:"requireDefaultRow,omitempty"`
}

//TextAreaFieldDescriptor - A field intended to be rendered as a text box in a UI.
type TextAreaFieldDescriptor struct {
	Advanced     *bool   `json:"advanced,omitempty"`
	Columns      *int    `json:"columns,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Rows         *int    `json:"rows,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//TextFieldDescriptor - A text field.
type TextFieldDescriptor struct {
	Advanced     *bool   `json:"advanced,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Encrypted    *bool   `json:"encrypted,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Size         *int    `json:"size,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//TextLocalIdentityField - A text type field.
type TextLocalIdentityField struct {
	Attributes            map[string]*bool `json:"attributes,omitempty"`
	DefaultValue          *string          `json:"defaultValue,omitempty"`
	Id                    *string          `json:"id,omitempty"`
	Label                 *string          `json:"label,omitempty"`
	ProfilePageField      *bool            `json:"profilePageField,omitempty"`
	RegistrationPageField *bool            `json:"registrationPageField,omitempty"`
	Type                  *string          `json:"type,omitempty"`
}

//TokenExchangeGeneratorGroup - The set of attributes used to configure a OAuth 2.0 Token Exchange Generator group.
type TokenExchangeGeneratorGroup struct {
	GeneratorMappings *[]*TokenExchangeGeneratorMapping `json:"generatorMappings,omitempty"`
	Id                *string                           `json:"id,omitempty"`
	Name              *string                           `json:"name,omitempty"`
	ResourceUris      *[]*string                        `json:"resourceUris,omitempty"`
}

//TokenExchangeGeneratorGroups - A collection of OAuth 2.0 Token Exchange Generator groups.
type TokenExchangeGeneratorGroups struct {
	Items *[]*TokenExchangeGeneratorGroup `json:"items,omitempty"`
}

//TokenExchangeGeneratorMapping - A Token Generator mapping into an OAuth 2.0 Token Exchange requested token type.
type TokenExchangeGeneratorMapping struct {
	DefaultMapping     *bool         `json:"defaultMapping,omitempty"`
	RequestedTokenType *string       `json:"requestedTokenType,omitempty"`
	TokenGenerator     *ResourceLink `json:"tokenGenerator,omitempty"`
}

//TokenExchangeGeneratorSettings - Settings for the OAuth Token Exchange Generator Groups.
type TokenExchangeGeneratorSettings struct {
	DefaultGeneratorGroupRef *ResourceLink `json:"defaultGeneratorGroupRef,omitempty"`
}

//TokenExchangeProcessorAttribute - An attribute for the OAuth 2.0 Token Exchange Processor policy attribute contract.
type TokenExchangeProcessorAttribute struct {
	Name *string `json:"name,omitempty"`
}

//TokenExchangeProcessorAttributeContract - A set of attributes exposed by an OAuth 2.0 Token Exchange Processor policy.
type TokenExchangeProcessorAttributeContract struct {
	CoreAttributes     *[]*TokenExchangeProcessorAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*TokenExchangeProcessorAttribute `json:"extendedAttributes,omitempty"`
}

//TokenExchangeProcessorMapping - A Token Processor(s) mapping into an OAuth 2.0 Token Exchange Processor policy.
type TokenExchangeProcessorMapping struct {
	ActorTokenProcessor          *ResourceLink                         `json:"actorTokenProcessor,omitempty"`
	ActorTokenType               *string                               `json:"actorTokenType,omitempty"`
	AttributeContractFulfillment map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources             *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	IssuanceCriteria             *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	SubjectTokenProcessor        *ResourceLink                         `json:"subjectTokenProcessor,omitempty"`
	SubjectTokenType             *string                               `json:"subjectTokenType,omitempty"`
}

//TokenExchangeProcessorPolicies - A collection of OAuth 2.0 Token Exchange Processor policies.
type TokenExchangeProcessorPolicies struct {
	Items *[]*TokenExchangeProcessorPolicy `json:"items,omitempty"`
}

//TokenExchangeProcessorPolicy - The set of attributes used to configure a OAuth 2.0 Token Exchange processor policy.
type TokenExchangeProcessorPolicy struct {
	ActorTokenRequired *bool                                    `json:"actorTokenRequired,omitempty"`
	AttributeContract  *TokenExchangeProcessorAttributeContract `json:"attributeContract,omitempty"`
	Id                 *string                                  `json:"id,omitempty"`
	Name               *string                                  `json:"name,omitempty"`
	ProcessorMappings  *[]*TokenExchangeProcessorMapping        `json:"processorMappings,omitempty"`
}

//TokenExchangeProcessorSettings - Settings for the OAuth Token Exchange Processor Policy configuration.
type TokenExchangeProcessorSettings struct {
	DefaultProcessorPolicyRef *ResourceLink `json:"defaultProcessorPolicyRef,omitempty"`
}

//TokenGenerator - A token generator instance.
type TokenGenerator struct {
	AttributeContract   *TokenGeneratorAttributeContract `json:"attributeContract,omitempty"`
	Configuration       *PluginConfiguration             `json:"configuration,omitempty"`
	Id                  *string                          `json:"id,omitempty"`
	Name                *string                          `json:"name,omitempty"`
	ParentRef           *ResourceLink                    `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink                    `json:"pluginDescriptorRef,omitempty"`
}

//TokenGeneratorAttribute - An attribute for the token generator attribute contract.
type TokenGeneratorAttribute struct {
	Name *string `json:"name,omitempty"`
}

//TokenGeneratorAttributeContract - A set of attributes exposed by a token generator.
type TokenGeneratorAttributeContract struct {
	CoreAttributes     *[]*TokenGeneratorAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*TokenGeneratorAttribute `json:"extendedAttributes,omitempty"`
	Inherited          *bool                       `json:"inherited,omitempty"`
}

//TokenGeneratorDescriptor - A token generator descriptor.
type TokenGeneratorDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//TokenGeneratorDescriptors - A collection of token generator descriptors.
type TokenGeneratorDescriptors struct {
	Items *[]*TokenGeneratorDescriptor `json:"items,omitempty"`
}

//TokenGenerators - A collection of token generators.
type TokenGenerators struct {
	Items *[]*TokenGenerator `json:"items,omitempty"`
}

//TokenProcessor - A token processor instance.
type TokenProcessor struct {
	AttributeContract   *TokenProcessorAttributeContract `json:"attributeContract,omitempty"`
	Configuration       *PluginConfiguration             `json:"configuration,omitempty"`
	Id                  *string                          `json:"id,omitempty"`
	Name                *string                          `json:"name,omitempty"`
	ParentRef           *ResourceLink                    `json:"parentRef,omitempty"`
	PluginDescriptorRef *ResourceLink                    `json:"pluginDescriptorRef,omitempty"`
}

//TokenProcessorAttribute - An attribute for the token processor attribute contract.
type TokenProcessorAttribute struct {
	Masked *bool   `json:"masked,omitempty"`
	Name   *string `json:"name,omitempty"`
}

//TokenProcessorAttributeContract - A set of attributes exposed by a token processor.
type TokenProcessorAttributeContract struct {
	CoreAttributes     *[]*TokenProcessorAttribute `json:"coreAttributes,omitempty"`
	ExtendedAttributes *[]*TokenProcessorAttribute `json:"extendedAttributes,omitempty"`
	Inherited          *bool                       `json:"inherited,omitempty"`
	MaskOgnlValues     *bool                       `json:"maskOgnlValues,omitempty"`
}

//TokenProcessorDescriptor - A token processor descriptor.
type TokenProcessorDescriptor struct {
	AttributeContract        *[]*string              `json:"attributeContract,omitempty"`
	ClassName                *string                 `json:"className,omitempty"`
	ConfigDescriptor         *PluginConfigDescriptor `json:"configDescriptor,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	SupportsExtendedContract *bool                   `json:"supportsExtendedContract,omitempty"`
}

//TokenProcessorDescriptors - A collection of token processor descriptors.
type TokenProcessorDescriptors struct {
	Items *[]*TokenProcessorDescriptor `json:"items,omitempty"`
}

//TokenProcessors - A collection of token processors.
type TokenProcessors struct {
	Items *[]*TokenProcessor `json:"items,omitempty"`
}

//TokenToTokenMapping - A Token Processor to Token Generator Mapping.
type TokenToTokenMapping struct {
	AttributeContractFulfillment     map[string]*AttributeFulfillmentValue `json:"attributeContractFulfillment,omitempty"`
	AttributeSources                 *[]*AttributeSource                   `json:"attributeSources,omitempty"`
	DefaultTargetResource            *string                               `json:"defaultTargetResource,omitempty"`
	Id                               *string                               `json:"id,omitempty"`
	IssuanceCriteria                 *IssuanceCriteria                     `json:"issuanceCriteria,omitempty"`
	LicenseConnectionGroupAssignment *string                               `json:"licenseConnectionGroupAssignment,omitempty"`
	SourceId                         *string                               `json:"sourceId,omitempty"`
	TargetId                         *string                               `json:"targetId,omitempty"`
}

//TokenToTokenMappings
type TokenToTokenMappings struct {
	Items *[]*TokenToTokenMapping `json:"items,omitempty"`
}

//UploadFileFieldDescriptor - A field which allows the user to upload a file.
type UploadFileFieldDescriptor struct {
	Advanced     *bool   `json:"advanced,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	Description  *string `json:"description,omitempty"`
	Label        *string `json:"label,omitempty"`
	Name         *string `json:"name,omitempty"`
	Required     *bool   `json:"required,omitempty"`
	Type         *string `json:"type,omitempty"`
}

//UrlWhitelistEntry - Url domain and path to be used as whitelist in WS-Federation connection
type UrlWhitelistEntry struct {
	AllowQueryAndFragment *bool   `json:"allowQueryAndFragment,omitempty"`
	RequireHttps          *bool   `json:"requireHttps,omitempty"`
	ValidDomain           *string `json:"validDomain,omitempty"`
	ValidPath             *string `json:"validPath,omitempty"`
}

//UserCredentials - Credentials for an administrator account.
type UserCredentials struct {
	CurrentPassword *string `json:"currentPassword,omitempty"`
	NewPassword     *string `json:"newPassword,omitempty"`
}

//UsernamePasswordCredentials - Username and password credentials.
type UsernamePasswordCredentials struct {
	EncryptedPassword *string `json:"encryptedPassword,omitempty"`
	Password          *string `json:"password,omitempty"`
	Username          *string `json:"username,omitempty"`
}

//ValidationError - A data input validation error.
type ValidationError struct {
	DeveloperMessage *string `json:"developerMessage,omitempty"`
	ErrorId          *string `json:"errorId,omitempty"`
	FieldPath        *string `json:"fieldPath,omitempty"`
	Message          *string `json:"message,omitempty"`
}

//Version - Server version.
type Version struct {
	Version *string `json:"version,omitempty"`
}

//VirtualHostNameSettings - Settings for virtual host names.
type VirtualHostNameSettings struct {
	VirtualHostNames *[]*string `json:"virtualHostNames,omitempty"`
}

//X509File - Encoded certificate data.
type X509File struct {
	CryptoProvider *string `json:"cryptoProvider,omitempty"`
	FileData       *string `json:"fileData,omitempty"`
	Id             *string `json:"id,omitempty"`
}
