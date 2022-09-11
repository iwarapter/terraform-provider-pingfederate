package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func expandApcToPersistentGrantMapping(in ApcToPersistentGrantMappingData) *pf.ApcToPersistentGrantMapping {
	var result pf.ApcToPersistentGrantMapping
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if result.AttributeSources == nil {
		result.AttributeSources = &[]*pf.AttributeSource{}
	}
	if in.LdapAttributeSources != nil {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSources(in.LdapAttributeSources)...)
	}
	if in.JdbcAttributeSources != nil {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSources(in.JdbcAttributeSources)...)
	}
	if in.CustomAttributeSources != nil {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSources(in.CustomAttributeSources)...)
	}
	if !in.AuthenticationPolicyContractRef.Unknown && !in.AuthenticationPolicyContractRef.Null {
		result.AuthenticationPolicyContractRef = &pf.ResourceLink{Id: String(in.AuthenticationPolicyContractRef.Value)}
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if in.IssuanceCriteria != nil {
		result.IssuanceCriteria = expandIssuanceCriteria(*in.IssuanceCriteria)
	}

	return &result
}

func expandAuthenticationPolicyContract(in AuthenticationPolicyContractData) *pf.AuthenticationPolicyContract {
	var result pf.AuthenticationPolicyContract
	result.CoreAttributes = &[]*pf.AuthenticationPolicyContractAttribute{
		{
			Name: String("subject"),
		},
	}
	if in.ExtendedAttributes != nil && len(in.ExtendedAttributes) > 0 {
		attrs := []*pf.AuthenticationPolicyContractAttribute{}
		for _, data := range in.ExtendedAttributes {
			attrs = append(attrs, &pf.AuthenticationPolicyContractAttribute{Name: String(data.Value)})
		}
		result.ExtendedAttributes = &attrs
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
	}

	return &result
}

func expandClient(in ClientData) *pf.Client {
	var result pf.Client
	if !in.AllowAuthenticationApiInit.Unknown && !in.AllowAuthenticationApiInit.Null {
		result.AllowAuthenticationApiInit = Bool(in.AllowAuthenticationApiInit.Value)
	}
	if !in.BypassActivationCodeConfirmationOverride.Unknown && !in.BypassActivationCodeConfirmationOverride.Null {
		result.BypassActivationCodeConfirmationOverride = Bool(in.BypassActivationCodeConfirmationOverride.Value)
	}
	if !in.BypassApprovalPage.Unknown && !in.BypassApprovalPage.Null {
		result.BypassApprovalPage = Bool(in.BypassApprovalPage.Value)
	}
	if !in.CibaDeliveryMode.Unknown && !in.CibaDeliveryMode.Null {
		result.CibaDeliveryMode = String(in.CibaDeliveryMode.Value)
	}
	if !in.CibaNotificationEndpoint.Unknown && !in.CibaNotificationEndpoint.Null {
		result.CibaNotificationEndpoint = String(in.CibaNotificationEndpoint.Value)
	}
	if !in.CibaPollingInterval.Unknown && !in.CibaPollingInterval.Null {
		i64, _ := in.CibaPollingInterval.Value.Int64()
		result.CibaPollingInterval = Int(int(i64))
	}
	if !in.CibaRequestObjectSigningAlgorithm.Unknown && !in.CibaRequestObjectSigningAlgorithm.Null {
		result.CibaRequestObjectSigningAlgorithm = String(in.CibaRequestObjectSigningAlgorithm.Value)
	}
	if !in.CibaRequireSignedRequests.Unknown && !in.CibaRequireSignedRequests.Null {
		result.CibaRequireSignedRequests = Bool(in.CibaRequireSignedRequests.Value)
	}
	if !in.CibaUserCodeSupported.Unknown && !in.CibaUserCodeSupported.Null {
		result.CibaUserCodeSupported = Bool(in.CibaUserCodeSupported.Value)
	}
	if in.ClientAuth != nil {
		result.ClientAuth = expandClientAuth(*in.ClientAuth)
	}
	if !in.ClientId.Unknown && !in.ClientId.Null {
		result.ClientId = String(in.ClientId.Value)
	}
	if !in.DefaultAccessTokenManagerRef.Unknown && !in.DefaultAccessTokenManagerRef.Null {
		result.DefaultAccessTokenManagerRef = &pf.ResourceLink{Id: String(in.DefaultAccessTokenManagerRef.Value)}
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if !in.DeviceFlowSettingType.Unknown && !in.DeviceFlowSettingType.Null {
		result.DeviceFlowSettingType = String(in.DeviceFlowSettingType.Value)
	}
	if !in.DevicePollingIntervalOverride.Unknown && !in.DevicePollingIntervalOverride.Null {
		i64, _ := in.DevicePollingIntervalOverride.Value.Int64()
		result.DevicePollingIntervalOverride = Int(int(i64))
	}
	if !in.Enabled.Unknown && !in.Enabled.Null {
		result.Enabled = Bool(in.Enabled.Value)
	}
	if in.ExclusiveScopes != nil {
		result.ExclusiveScopes = expandStringList(in.ExclusiveScopes)
	}
	if in.ExtendedParameters != nil {
		result.ExtendedParameters = expandMapParameterValuess(in.ExtendedParameters)
	}
	if in.GrantTypes != nil {
		result.GrantTypes = expandStringList(in.GrantTypes)
	}
	if in.JwksSettings != nil {
		result.JwksSettings = expandJwksSettings(*in.JwksSettings)
	}
	if !in.LogoUrl.Unknown && !in.LogoUrl.Null {
		result.LogoUrl = String(in.LogoUrl.Value)
	}
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
	}
	if in.OidcPolicy != nil {
		result.OidcPolicy = expandClientOIDCPolicy(*in.OidcPolicy)
	}
	if !in.PendingAuthorizationTimeoutOverride.Unknown && !in.PendingAuthorizationTimeoutOverride.Null {
		i64, _ := in.PendingAuthorizationTimeoutOverride.Value.Int64()
		result.PendingAuthorizationTimeoutOverride = Int(int(i64))
	}
	if !in.PersistentGrantExpirationTime.Unknown && !in.PersistentGrantExpirationTime.Null {
		i64, _ := in.PersistentGrantExpirationTime.Value.Int64()
		result.PersistentGrantExpirationTime = Int(int(i64))
	}
	if !in.PersistentGrantExpirationTimeUnit.Unknown && !in.PersistentGrantExpirationTimeUnit.Null {
		result.PersistentGrantExpirationTimeUnit = String(in.PersistentGrantExpirationTimeUnit.Value)
	}
	if !in.PersistentGrantExpirationType.Unknown && !in.PersistentGrantExpirationType.Null {
		result.PersistentGrantExpirationType = String(in.PersistentGrantExpirationType.Value)
	}
	if !in.PersistentGrantIdleTimeout.Unknown && !in.PersistentGrantIdleTimeout.Null {
		i64, _ := in.PersistentGrantIdleTimeout.Value.Int64()
		result.PersistentGrantIdleTimeout = Int(int(i64))
	}
	if !in.PersistentGrantIdleTimeoutTimeUnit.Unknown && !in.PersistentGrantIdleTimeoutTimeUnit.Null {
		result.PersistentGrantIdleTimeoutTimeUnit = String(in.PersistentGrantIdleTimeoutTimeUnit.Value)
	}
	if !in.PersistentGrantIdleTimeoutType.Unknown && !in.PersistentGrantIdleTimeoutType.Null {
		result.PersistentGrantIdleTimeoutType = String(in.PersistentGrantIdleTimeoutType.Value)
	}
	if in.PersistentGrantReuseGrantTypes != nil {
		result.PersistentGrantReuseGrantTypes = expandStringList(in.PersistentGrantReuseGrantTypes)
	}
	if !in.PersistentGrantReuseType.Unknown && !in.PersistentGrantReuseType.Null {
		result.PersistentGrantReuseType = String(in.PersistentGrantReuseType.Value)
	}
	if in.RedirectUris != nil {
		result.RedirectUris = expandStringList(in.RedirectUris)
	}
	if !in.RefreshRolling.Unknown && !in.RefreshRolling.Null {
		result.RefreshRolling = String(in.RefreshRolling.Value)
	}
	if !in.RefreshTokenRollingInterval.Unknown && !in.RefreshTokenRollingInterval.Null {
		i64, _ := in.RefreshTokenRollingInterval.Value.Int64()
		result.RefreshTokenRollingInterval = Int(int(i64))
	}
	if !in.RefreshTokenRollingIntervalType.Unknown && !in.RefreshTokenRollingIntervalType.Null {
		result.RefreshTokenRollingIntervalType = String(in.RefreshTokenRollingIntervalType.Value)
	}
	if !in.RequestObjectSigningAlgorithm.Unknown && !in.RequestObjectSigningAlgorithm.Null {
		result.RequestObjectSigningAlgorithm = String(in.RequestObjectSigningAlgorithm.Value)
	}
	if !in.RequestPolicyRef.Unknown && !in.RequestPolicyRef.Null {
		result.RequestPolicyRef = &pf.ResourceLink{Id: String(in.RequestPolicyRef.Value)}
	}
	if !in.RequireProofKeyForCodeExchange.Unknown && !in.RequireProofKeyForCodeExchange.Null {
		result.RequireProofKeyForCodeExchange = Bool(in.RequireProofKeyForCodeExchange.Value)
	}
	if !in.RequirePushedAuthorizationRequests.Unknown && !in.RequirePushedAuthorizationRequests.Null {
		result.RequirePushedAuthorizationRequests = Bool(in.RequirePushedAuthorizationRequests.Value)
	}
	if !in.RequireSignedRequests.Unknown && !in.RequireSignedRequests.Null {
		result.RequireSignedRequests = Bool(in.RequireSignedRequests.Value)
	}
	if !in.RestrictScopes.Unknown && !in.RestrictScopes.Null {
		result.RestrictScopes = Bool(in.RestrictScopes.Value)
	}
	if !in.RestrictToDefaultAccessTokenManager.Unknown && !in.RestrictToDefaultAccessTokenManager.Null {
		result.RestrictToDefaultAccessTokenManager = Bool(in.RestrictToDefaultAccessTokenManager.Value)
	}
	if in.RestrictedResponseTypes != nil {
		result.RestrictedResponseTypes = expandStringList(in.RestrictedResponseTypes)
	}
	if in.RestrictedScopes != nil {
		result.RestrictedScopes = expandStringList(in.RestrictedScopes)
	}
	if !in.TokenExchangeProcessorPolicyRef.Unknown && !in.TokenExchangeProcessorPolicyRef.Null {
		result.TokenExchangeProcessorPolicyRef = &pf.ResourceLink{Id: String(in.TokenExchangeProcessorPolicyRef.Value)}
	}
	if !in.UserAuthorizationUrlOverride.Unknown && !in.UserAuthorizationUrlOverride.Null {
		result.UserAuthorizationUrlOverride = String(in.UserAuthorizationUrlOverride.Value)
	}
	if !in.ValidateUsingAllEligibleAtms.Unknown && !in.ValidateUsingAllEligibleAtms.Null {
		result.ValidateUsingAllEligibleAtms = Bool(in.ValidateUsingAllEligibleAtms.Value)
	}

	return &result
}

func expandRedirectValidationSettings(in RedirectValidationSettingsData) *pf.RedirectValidationSettings {
	var result pf.RedirectValidationSettings
	if in.RedirectValidationLocalSettings != nil {
		result.RedirectValidationLocalSettings = expandRedirectValidationLocalSettings(*in.RedirectValidationLocalSettings)
	}
	if in.RedirectValidationPartnerSettings != nil {
		result.RedirectValidationPartnerSettings = expandRedirectValidationPartnerSettings(*in.RedirectValidationPartnerSettings)
	}

	return &result
}

func expandMapAttributeFulfillmentValues(in map[string]*AttributeFulfillmentValueData) map[string]*pf.AttributeFulfillmentValue {
	results := map[string]*pf.AttributeFulfillmentValue{}
	for s, data := range in {
		results[s] = expandAttributeFulfillmentValue(*data)
	}
	return results
}
func expandAttributeFulfillmentValue(in AttributeFulfillmentValueData) *pf.AttributeFulfillmentValue {
	var result pf.AttributeFulfillmentValue
	if in.Source != nil {
		result.Source = expandSourceTypeIdKey(*in.Source)
	}
	if !in.Value.Unknown && !in.Value.Null {
		result.Value = String(in.Value.Value)
	}

	return &result
}

func expandMapBinaryLdapAttributeSettingss(in map[string]*BinaryLdapAttributeSettingsData) map[string]*pf.BinaryLdapAttributeSettings {
	results := map[string]*pf.BinaryLdapAttributeSettings{}
	for s, data := range in {
		results[s] = expandBinaryLdapAttributeSettings(*data)
	}
	return results
}
func expandBinaryLdapAttributeSettings(in BinaryLdapAttributeSettingsData) *pf.BinaryLdapAttributeSettings {
	var result pf.BinaryLdapAttributeSettings
	if !in.BinaryEncoding.Unknown && !in.BinaryEncoding.Null {
		result.BinaryEncoding = String(in.BinaryEncoding.Value)
	}

	return &result
}

func expandClientAuth(in ClientAuthData) *pf.ClientAuth {
	var result pf.ClientAuth
	if !in.ClientCertIssuerDn.Unknown && !in.ClientCertIssuerDn.Null {
		result.ClientCertIssuerDn = String(in.ClientCertIssuerDn.Value)
	}
	if !in.ClientCertSubjectDn.Unknown && !in.ClientCertSubjectDn.Null {
		result.ClientCertSubjectDn = String(in.ClientCertSubjectDn.Value)
	}
	if !in.EncryptedSecret.Unknown && !in.EncryptedSecret.Null {
		result.EncryptedSecret = String(in.EncryptedSecret.Value)
	}
	if !in.EnforceReplayPrevention.Unknown && !in.EnforceReplayPrevention.Null {
		result.EnforceReplayPrevention = Bool(in.EnforceReplayPrevention.Value)
	}
	if !in.Secret.Unknown && !in.Secret.Null {
		result.Secret = String(in.Secret.Value)
	}
	if !in.TokenEndpointAuthSigningAlgorithm.Unknown && !in.TokenEndpointAuthSigningAlgorithm.Null {
		result.TokenEndpointAuthSigningAlgorithm = String(in.TokenEndpointAuthSigningAlgorithm.Value)
	}
	if !in.Type.Unknown && !in.Type.Null {
		result.Type = String(in.Type.Value)
	}

	return &result
}

func expandClientOIDCPolicy(in ClientOIDCPolicyData) *pf.ClientOIDCPolicy {
	var result pf.ClientOIDCPolicy
	if !in.GrantAccessSessionRevocationApi.Unknown && !in.GrantAccessSessionRevocationApi.Null {
		result.GrantAccessSessionRevocationApi = Bool(in.GrantAccessSessionRevocationApi.Value)
	}
	if !in.GrantAccessSessionSessionManagementApi.Unknown && !in.GrantAccessSessionSessionManagementApi.Null {
		result.GrantAccessSessionSessionManagementApi = Bool(in.GrantAccessSessionSessionManagementApi.Value)
	}
	if !in.IdTokenContentEncryptionAlgorithm.Unknown && !in.IdTokenContentEncryptionAlgorithm.Null {
		result.IdTokenContentEncryptionAlgorithm = String(in.IdTokenContentEncryptionAlgorithm.Value)
	}
	if !in.IdTokenEncryptionAlgorithm.Unknown && !in.IdTokenEncryptionAlgorithm.Null {
		result.IdTokenEncryptionAlgorithm = String(in.IdTokenEncryptionAlgorithm.Value)
	}
	if !in.IdTokenSigningAlgorithm.Unknown && !in.IdTokenSigningAlgorithm.Null {
		result.IdTokenSigningAlgorithm = String(in.IdTokenSigningAlgorithm.Value)
	}
	if in.LogoutUris != nil {
		result.LogoutUris = expandStringList(in.LogoutUris)
	}
	if !in.PairwiseIdentifierUserType.Unknown && !in.PairwiseIdentifierUserType.Null {
		result.PairwiseIdentifierUserType = Bool(in.PairwiseIdentifierUserType.Value)
	}
	if !in.PingAccessLogoutCapable.Unknown && !in.PingAccessLogoutCapable.Null {
		result.PingAccessLogoutCapable = Bool(in.PingAccessLogoutCapable.Value)
	}
	if !in.PolicyGroup.Unknown && !in.PolicyGroup.Null {
		result.PolicyGroup = &pf.ResourceLink{Id: String(in.PolicyGroup.Value)}
	}
	if !in.SectorIdentifierUri.Unknown && !in.SectorIdentifierUri.Null {
		result.SectorIdentifierUri = String(in.SectorIdentifierUri.Value)
	}

	return &result
}

func expandConditionalIssuanceCriteriaEntrys(in *[]*ConditionalIssuanceCriteriaEntryData) *[]*pf.ConditionalIssuanceCriteriaEntry {
	results := []*pf.ConditionalIssuanceCriteriaEntry{}
	for _, data := range *in {
		results = append(results, expandConditionalIssuanceCriteriaEntry(*data))
	}
	return &results
}
func expandConditionalIssuanceCriteriaEntry(in ConditionalIssuanceCriteriaEntryData) *pf.ConditionalIssuanceCriteriaEntry {
	var result pf.ConditionalIssuanceCriteriaEntry
	if !in.AttributeName.Unknown && !in.AttributeName.Null {
		result.AttributeName = String(in.AttributeName.Value)
	}
	if !in.Condition.Unknown && !in.Condition.Null {
		result.Condition = String(in.Condition.Value)
	}
	if !in.ErrorResult.Unknown && !in.ErrorResult.Null {
		result.ErrorResult = String(in.ErrorResult.Value)
	}
	if in.Source != nil {
		result.Source = expandSourceTypeIdKey(*in.Source)
	}
	if !in.Value.Unknown && !in.Value.Null {
		result.Value = String(in.Value.Value)
	}

	return &result
}

func expandCustomAttributeSource(in CustomAttributeSourceData) *pf.CustomAttributeSource {
	var result pf.CustomAttributeSource
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if in.FilterFields != nil {
		result.FilterFields = expandFieldEntrys(in.FilterFields)
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}

	return &result
}

func expandExpressionIssuanceCriteriaEntrys(in *[]*ExpressionIssuanceCriteriaEntryData) *[]*pf.ExpressionIssuanceCriteriaEntry {
	results := []*pf.ExpressionIssuanceCriteriaEntry{}
	for _, data := range *in {
		results = append(results, expandExpressionIssuanceCriteriaEntry(*data))
	}
	return &results
}
func expandExpressionIssuanceCriteriaEntry(in ExpressionIssuanceCriteriaEntryData) *pf.ExpressionIssuanceCriteriaEntry {
	var result pf.ExpressionIssuanceCriteriaEntry
	if !in.ErrorResult.Unknown && !in.ErrorResult.Null {
		result.ErrorResult = String(in.ErrorResult.Value)
	}
	if !in.Expression.Unknown && !in.Expression.Null {
		result.Expression = String(in.Expression.Value)
	}

	return &result
}

func expandFieldEntrys(in *[]*FieldEntryData) *[]*pf.FieldEntry {
	results := []*pf.FieldEntry{}
	for _, data := range *in {
		results = append(results, expandFieldEntry(*data))
	}
	return &results
}
func expandFieldEntry(in FieldEntryData) *pf.FieldEntry {
	var result pf.FieldEntry
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
	}
	if !in.Value.Unknown && !in.Value.Null {
		result.Value = String(in.Value.Value)
	}

	return &result
}

func expandIssuanceCriteria(in IssuanceCriteriaData) *pf.IssuanceCriteria {
	var result pf.IssuanceCriteria
	if in.ConditionalCriteria != nil {
		result.ConditionalCriteria = expandConditionalIssuanceCriteriaEntrys(in.ConditionalCriteria)
	}
	if in.ExpressionCriteria != nil {
		result.ExpressionCriteria = expandExpressionIssuanceCriteriaEntrys(in.ExpressionCriteria)
	}

	return &result
}

func expandJdbcAttributeSource(in JdbcAttributeSourceData) *pf.JdbcAttributeSource {
	var result pf.JdbcAttributeSource
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.ColumnNames != nil {
		result.ColumnNames = expandStringList(in.ColumnNames)
	}
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if !in.Filter.Unknown && !in.Filter.Null {
		result.Filter = String(in.Filter.Value)
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.Schema.Unknown && !in.Schema.Null {
		result.Schema = String(in.Schema.Value)
	}
	if !in.Table.Unknown && !in.Table.Null {
		result.Table = String(in.Table.Value)
	}

	return &result
}

func expandJwksSettings(in JwksSettingsData) *pf.JwksSettings {
	var result pf.JwksSettings
	if !in.Jwks.Unknown && !in.Jwks.Null {
		result.Jwks = String(in.Jwks.Value)
	}
	if !in.JwksUrl.Unknown && !in.JwksUrl.Null {
		result.JwksUrl = String(in.JwksUrl.Value)
	}

	return &result
}

func expandLdapAttributeSource(in LdapAttributeSourceData) *pf.LdapAttributeSource {
	var result pf.LdapAttributeSource
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.BaseDn.Unknown && !in.BaseDn.Null {
		result.BaseDn = String(in.BaseDn.Value)
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = expandMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
	}
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.MemberOfNestedGroup.Unknown && !in.MemberOfNestedGroup.Null {
		result.MemberOfNestedGroup = Bool(in.MemberOfNestedGroup.Value)
	}
	if in.SearchAttributes != nil {
		result.SearchAttributes = expandStringList(in.SearchAttributes)
	}
	if !in.SearchFilter.Unknown && !in.SearchFilter.Null {
		result.SearchFilter = String(in.SearchFilter.Value)
	}
	if !in.SearchScope.Unknown && !in.SearchScope.Null {
		result.SearchScope = String(in.SearchScope.Value)
	}

	return &result
}

func expandMapParameterValuess(in map[string]*ParameterValuesData) map[string]*pf.ParameterValues {
	results := map[string]*pf.ParameterValues{}
	for s, data := range in {
		results[s] = expandParameterValues(*data)
	}
	return results
}
func expandParameterValues(in ParameterValuesData) *pf.ParameterValues {
	var result pf.ParameterValues
	if in.Values != nil {
		result.Values = expandStringList(in.Values)
	}

	return &result
}

func expandRedirectValidationLocalSettings(in RedirectValidationLocalSettingsData) *pf.RedirectValidationLocalSettings {
	var result pf.RedirectValidationLocalSettings
	if !in.EnableInErrorResourceValidation.Unknown && !in.EnableInErrorResourceValidation.Null {
		result.EnableInErrorResourceValidation = Bool(in.EnableInErrorResourceValidation.Value)
	}
	if !in.EnableTargetResourceValidationForIdpDiscovery.Unknown && !in.EnableTargetResourceValidationForIdpDiscovery.Null {
		result.EnableTargetResourceValidationForIdpDiscovery = Bool(in.EnableTargetResourceValidationForIdpDiscovery.Value)
	}
	if !in.EnableTargetResourceValidationForSLO.Unknown && !in.EnableTargetResourceValidationForSLO.Null {
		result.EnableTargetResourceValidationForSLO = Bool(in.EnableTargetResourceValidationForSLO.Value)
	}
	if !in.EnableTargetResourceValidationForSSO.Unknown && !in.EnableTargetResourceValidationForSSO.Null {
		result.EnableTargetResourceValidationForSSO = Bool(in.EnableTargetResourceValidationForSSO.Value)
	}
	if in.WhiteList != nil {
		result.WhiteList = expandRedirectValidationSettingsWhitelistEntrys(in.WhiteList)
	}

	return &result
}

func expandRedirectValidationPartnerSettings(in RedirectValidationPartnerSettingsData) *pf.RedirectValidationPartnerSettings {
	var result pf.RedirectValidationPartnerSettings
	if !in.EnableWreplyValidationSLO.Unknown && !in.EnableWreplyValidationSLO.Null {
		result.EnableWreplyValidationSLO = Bool(in.EnableWreplyValidationSLO.Value)
	}

	return &result
}

func expandRedirectValidationSettingsWhitelistEntrys(in *[]*RedirectValidationSettingsWhitelistEntryData) *[]*pf.RedirectValidationSettingsWhitelistEntry {
	results := []*pf.RedirectValidationSettingsWhitelistEntry{}
	for _, data := range *in {
		results = append(results, expandRedirectValidationSettingsWhitelistEntry(*data))
	}
	return &results
}
func expandRedirectValidationSettingsWhitelistEntry(in RedirectValidationSettingsWhitelistEntryData) *pf.RedirectValidationSettingsWhitelistEntry {
	var result pf.RedirectValidationSettingsWhitelistEntry
	if !in.AllowQueryAndFragment.Unknown && !in.AllowQueryAndFragment.Null {
		result.AllowQueryAndFragment = Bool(in.AllowQueryAndFragment.Value)
	}
	if !in.IdpDiscovery.Unknown && !in.IdpDiscovery.Null {
		result.IdpDiscovery = Bool(in.IdpDiscovery.Value)
	}
	if !in.InErrorResource.Unknown && !in.InErrorResource.Null {
		result.InErrorResource = Bool(in.InErrorResource.Value)
	}
	if !in.RequireHttps.Unknown && !in.RequireHttps.Null {
		result.RequireHttps = Bool(in.RequireHttps.Value)
	}
	if !in.TargetResourceSLO.Unknown && !in.TargetResourceSLO.Null {
		result.TargetResourceSLO = Bool(in.TargetResourceSLO.Value)
	}
	if !in.TargetResourceSSO.Unknown && !in.TargetResourceSSO.Null {
		result.TargetResourceSSO = Bool(in.TargetResourceSSO.Value)
	}
	if !in.ValidDomain.Unknown && !in.ValidDomain.Null {
		result.ValidDomain = String(in.ValidDomain.Value)
	}
	if !in.ValidPath.Unknown && !in.ValidPath.Null {
		result.ValidPath = String(in.ValidPath.Value)
	}

	return &result
}

func expandSourceTypeIdKey(in SourceTypeIdKeyData) *pf.SourceTypeIdKey {
	var result pf.SourceTypeIdKey
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.Type.Unknown && !in.Type.Null {
		result.Type = String(in.Type.Value)
	}

	return &result
}

func expandJdbcAttributeSources(in []JdbcAttributeSourceData) *[]*pf.AttributeSource {
	results := []*pf.AttributeSource{}
	for _, data := range in {
		source := &pf.AttributeSource{
			Type:                String("JDBC"),
			JdbcAttributeSource: *expandJdbcAttributeSource(data),
		}
		source.JdbcAttributeSource.Type = String("JDBC")
		source.DataStoreRef = source.JdbcAttributeSource.DataStoreRef
		source.AttributeContractFulfillment = source.JdbcAttributeSource.AttributeContractFulfillment
		source.Id = source.JdbcAttributeSource.Id
		source.Description = source.JdbcAttributeSource.Description

		results = append(results, source)
	}
	return &results
}

func expandLdapAttributeSources(in []LdapAttributeSourceData) *[]*pf.AttributeSource {
	results := []*pf.AttributeSource{}
	for _, data := range in {
		source := &pf.AttributeSource{
			Type:                String("LDAP"),
			LdapAttributeSource: *expandLdapAttributeSource(data),
		}
		source.LdapAttributeSource.Type = String("LDAP")
		source.DataStoreRef = source.LdapAttributeSource.DataStoreRef
		source.AttributeContractFulfillment = source.LdapAttributeSource.AttributeContractFulfillment
		source.Id = source.LdapAttributeSource.Id
		source.Description = source.LdapAttributeSource.Description
		results = append(results, source)
	}
	return &results
}

func expandCustomAttributeSources(in []CustomAttributeSourceData) *[]*pf.AttributeSource {
	results := []*pf.AttributeSource{}
	for _, data := range in {
		source := &pf.AttributeSource{
			Type:                  String("CUSTOM"),
			CustomAttributeSource: *expandCustomAttributeSource(data),
		}
		source.CustomAttributeSource.Type = String("CUSTOM")
		source.DataStoreRef = source.CustomAttributeSource.DataStoreRef
		source.AttributeContractFulfillment = source.CustomAttributeSource.AttributeContractFulfillment
		source.Id = source.CustomAttributeSource.Id
		source.Description = source.CustomAttributeSource.Description
		results = append(results, source)
	}
	return &results
}

func expandStringList(in []types.String) *[]*string {
	results := []*string{}
	for _, s := range in {
		results = append(results, String(s.Value))
	}
	return &results
}
