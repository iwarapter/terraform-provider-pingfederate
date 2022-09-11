package framework

import (
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/types"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func flattenApcToPersistentGrantMapping(in *pf.ApcToPersistentGrantMapping) *ApcToPersistentGrantMappingData {
	result := ApcToPersistentGrantMappingData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		result.JdbcAttributeSources = append(result.JdbcAttributeSources, flattenJdbcAttributeSources(in.AttributeSources)...)
	}
	if in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		result.LdapAttributeSources = append(result.LdapAttributeSources, flattenLdapAttributeSources(in.AttributeSources)...)
	}
	if in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		result.CustomAttributeSources = append(result.CustomAttributeSources, flattenCustomAttributeSources(in.AttributeSources)...)
	}
	if in.AuthenticationPolicyContractRef != nil && in.AuthenticationPolicyContractRef.Id != nil && *in.AuthenticationPolicyContractRef.Id != "" {
		result.AuthenticationPolicyContractRef = types.String{Value: *in.AuthenticationPolicyContractRef.Id}
	} else {
		result.AuthenticationPolicyContractRef = types.String{Null: true}
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		result.IssuanceCriteria = flattenIssuanceCriteria(in.IssuanceCriteria)
	}

	return &result
}

func flattenAuthenticationPolicyContract(in *pf.AuthenticationPolicyContract) *AuthenticationPolicyContractData {
	result := AuthenticationPolicyContractData{}
	if in.CoreAttributes != nil {
		attrs := []types.String{}
		for _, attribute := range *in.CoreAttributes {
			attrs = append(attrs, types.String{Value: *attribute.Name})
		}
		result.CoreAttributes = attrs
	}
	if in.ExtendedAttributes != nil {
		attrs := []types.String{}
		for _, attribute := range *in.ExtendedAttributes {
			attrs = append(attrs, types.String{Value: *attribute.Name})
		}
		result.ExtendedAttributes = attrs
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
	}

	return &result
}

func flattenClient(in *pf.Client) *ClientData {
	result := ClientData{}
	if in.AllowAuthenticationApiInit != nil {
		result.AllowAuthenticationApiInit = types.Bool{Value: *in.AllowAuthenticationApiInit}
	} else {
		result.AllowAuthenticationApiInit = types.Bool{Null: true}
	}
	if in.BypassActivationCodeConfirmationOverride != nil {
		result.BypassActivationCodeConfirmationOverride = types.Bool{Value: *in.BypassActivationCodeConfirmationOverride}
	} else {
		result.BypassActivationCodeConfirmationOverride = types.Bool{Null: true}
	}
	if in.BypassApprovalPage != nil {
		result.BypassApprovalPage = types.Bool{Value: *in.BypassApprovalPage}
	} else {
		result.BypassApprovalPage = types.Bool{Null: true}
	}
	if in.CibaDeliveryMode != nil {
		result.CibaDeliveryMode = types.String{Value: *in.CibaDeliveryMode}
	} else {
		result.CibaDeliveryMode = types.String{Null: true}
	}
	if in.CibaNotificationEndpoint != nil {
		result.CibaNotificationEndpoint = types.String{Value: *in.CibaNotificationEndpoint}
	} else {
		result.CibaNotificationEndpoint = types.String{Null: true}
	}
	if in.CibaPollingInterval != nil {
		result.CibaPollingInterval = types.Number{Value: big.NewFloat(float64(*in.CibaPollingInterval))}
	} else {
		result.CibaPollingInterval = types.Number{Null: true}
	}
	if in.CibaRequestObjectSigningAlgorithm != nil {
		result.CibaRequestObjectSigningAlgorithm = types.String{Value: *in.CibaRequestObjectSigningAlgorithm}
	} else {
		result.CibaRequestObjectSigningAlgorithm = types.String{Null: true}
	}
	if in.CibaRequireSignedRequests != nil {
		result.CibaRequireSignedRequests = types.Bool{Value: *in.CibaRequireSignedRequests}
	} else {
		result.CibaRequireSignedRequests = types.Bool{Null: true}
	}
	if in.CibaUserCodeSupported != nil {
		result.CibaUserCodeSupported = types.Bool{Value: *in.CibaUserCodeSupported}
	} else {
		result.CibaUserCodeSupported = types.Bool{Null: true}
	}
	if in.ClientAuth != nil && *in.ClientAuth.Type != "NONE" {
		result.ClientAuth = flattenClientAuth(in.ClientAuth)
	}
	if in.ClientId != nil {
		result.ClientId = types.String{Value: *in.ClientId}
	}
	if in.DefaultAccessTokenManagerRef != nil && in.DefaultAccessTokenManagerRef.Id != nil && *in.DefaultAccessTokenManagerRef.Id != "" {
		result.DefaultAccessTokenManagerRef = types.String{Value: *in.DefaultAccessTokenManagerRef.Id}
	} else {
		result.DefaultAccessTokenManagerRef = types.String{Null: true}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.DeviceFlowSettingType != nil {
		result.DeviceFlowSettingType = types.String{Value: *in.DeviceFlowSettingType}
	} else {
		result.DeviceFlowSettingType = types.String{Null: true}
	}
	if in.DevicePollingIntervalOverride != nil {
		result.DevicePollingIntervalOverride = types.Number{Value: big.NewFloat(float64(*in.DevicePollingIntervalOverride))}
	} else {
		result.DevicePollingIntervalOverride = types.Number{Null: true}
	}
	if in.Enabled != nil {
		result.Enabled = types.Bool{Value: *in.Enabled}
	} else {
		result.Enabled = types.Bool{Null: true}
	}
	if in.ExclusiveScopes != nil {
		result.ExclusiveScopes = flattenStringList(*in.ExclusiveScopes)
	}
	if in.ExtendedParameters != nil {
		result.ExtendedParameters = flattenMapParameterValuess(in.ExtendedParameters)
	}
	if in.GrantTypes != nil {
		result.GrantTypes = flattenStringList(*in.GrantTypes)
	}
	if in.ClientId != nil {
		result.Id = types.String{Value: *in.ClientId}
	}
	if in.JwksSettings != nil {
		result.JwksSettings = flattenJwksSettings(in.JwksSettings)
	}
	if in.LogoUrl != nil {
		result.LogoUrl = types.String{Value: *in.LogoUrl}
	} else {
		result.LogoUrl = types.String{Null: true}
	}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
	}
	if in.OidcPolicy != nil {
		result.OidcPolicy = flattenClientOIDCPolicy(in.OidcPolicy)
	}
	if in.PendingAuthorizationTimeoutOverride != nil {
		result.PendingAuthorizationTimeoutOverride = types.Number{Value: big.NewFloat(float64(*in.PendingAuthorizationTimeoutOverride))}
	} else {
		result.PendingAuthorizationTimeoutOverride = types.Number{Null: true}
	}
	if in.PersistentGrantExpirationTime != nil {
		result.PersistentGrantExpirationTime = types.Number{Value: big.NewFloat(float64(*in.PersistentGrantExpirationTime))}
	} else {
		result.PersistentGrantExpirationTime = types.Number{Null: true}
	}
	if in.PersistentGrantExpirationTimeUnit != nil {
		result.PersistentGrantExpirationTimeUnit = types.String{Value: *in.PersistentGrantExpirationTimeUnit}
	} else {
		result.PersistentGrantExpirationTimeUnit = types.String{Value: "DAYS"}
	}
	if in.PersistentGrantExpirationType != nil {
		result.PersistentGrantExpirationType = types.String{Value: *in.PersistentGrantExpirationType}
	} else {
		result.PersistentGrantExpirationType = types.String{Null: true}
	}
	if in.PersistentGrantIdleTimeout != nil {
		result.PersistentGrantIdleTimeout = types.Number{Value: big.NewFloat(float64(*in.PersistentGrantIdleTimeout))}
	} else {
		result.PersistentGrantIdleTimeout = types.Number{Null: true}
	}
	if in.PersistentGrantIdleTimeoutTimeUnit != nil {
		result.PersistentGrantIdleTimeoutTimeUnit = types.String{Value: *in.PersistentGrantIdleTimeoutTimeUnit}
	} else {
		result.PersistentGrantIdleTimeoutTimeUnit = types.String{Null: true}
	}
	if in.PersistentGrantIdleTimeoutType != nil {
		result.PersistentGrantIdleTimeoutType = types.String{Value: *in.PersistentGrantIdleTimeoutType}
	} else {
		result.PersistentGrantIdleTimeoutType = types.String{Null: true}
	}
	if in.PersistentGrantReuseGrantTypes != nil {
		result.PersistentGrantReuseGrantTypes = flattenStringList(*in.PersistentGrantReuseGrantTypes)
	}
	if in.PersistentGrantReuseType != nil {
		result.PersistentGrantReuseType = types.String{Value: *in.PersistentGrantReuseType}
	} else {
		result.PersistentGrantReuseType = types.String{Null: true}
	}
	if in.RedirectUris != nil {
		result.RedirectUris = flattenStringList(*in.RedirectUris)
	}
	if in.RefreshRolling != nil {
		result.RefreshRolling = types.String{Value: *in.RefreshRolling}
	} else {
		result.RefreshRolling = types.String{Null: true}
	}
	if in.RefreshTokenRollingInterval != nil {
		result.RefreshTokenRollingInterval = types.Number{Value: big.NewFloat(float64(*in.RefreshTokenRollingInterval))}
	} else {
		result.RefreshTokenRollingInterval = types.Number{Null: true}
	}
	if in.RefreshTokenRollingIntervalType != nil {
		result.RefreshTokenRollingIntervalType = types.String{Value: *in.RefreshTokenRollingIntervalType}
	} else {
		result.RefreshTokenRollingIntervalType = types.String{Null: true}
	}
	if in.RequestObjectSigningAlgorithm != nil {
		result.RequestObjectSigningAlgorithm = types.String{Value: *in.RequestObjectSigningAlgorithm}
	} else {
		result.RequestObjectSigningAlgorithm = types.String{Null: true}
	}
	if in.RequestPolicyRef != nil && in.RequestPolicyRef.Id != nil && *in.RequestPolicyRef.Id != "" {
		result.RequestPolicyRef = types.String{Value: *in.RequestPolicyRef.Id}
	} else {
		result.RequestPolicyRef = types.String{Null: true}
	}
	if in.RequireProofKeyForCodeExchange != nil {
		result.RequireProofKeyForCodeExchange = types.Bool{Value: *in.RequireProofKeyForCodeExchange}
	} else {
		result.RequireProofKeyForCodeExchange = types.Bool{Null: true}
	}
	if in.RequirePushedAuthorizationRequests != nil {
		result.RequirePushedAuthorizationRequests = types.Bool{Value: *in.RequirePushedAuthorizationRequests}
	} else {
		result.RequirePushedAuthorizationRequests = types.Bool{Null: true}
	}
	if in.RequireSignedRequests != nil {
		result.RequireSignedRequests = types.Bool{Value: *in.RequireSignedRequests}
	} else {
		result.RequireSignedRequests = types.Bool{Null: true}
	}
	if in.RestrictScopes != nil {
		result.RestrictScopes = types.Bool{Value: *in.RestrictScopes}
	} else {
		result.RestrictScopes = types.Bool{Null: true}
	}
	if in.RestrictToDefaultAccessTokenManager != nil {
		result.RestrictToDefaultAccessTokenManager = types.Bool{Value: *in.RestrictToDefaultAccessTokenManager}
	} else {
		result.RestrictToDefaultAccessTokenManager = types.Bool{Null: true}
	}
	if in.RestrictedResponseTypes != nil {
		result.RestrictedResponseTypes = flattenStringList(*in.RestrictedResponseTypes)
	}
	if in.RestrictedScopes != nil {
		result.RestrictedScopes = flattenStringList(*in.RestrictedScopes)
	}
	if in.TokenExchangeProcessorPolicyRef != nil && in.TokenExchangeProcessorPolicyRef.Id != nil && *in.TokenExchangeProcessorPolicyRef.Id != "" {
		result.TokenExchangeProcessorPolicyRef = types.String{Value: *in.TokenExchangeProcessorPolicyRef.Id}
	} else {
		result.TokenExchangeProcessorPolicyRef = types.String{Null: true}
	}
	if in.UserAuthorizationUrlOverride != nil {
		result.UserAuthorizationUrlOverride = types.String{Value: *in.UserAuthorizationUrlOverride}
	} else {
		result.UserAuthorizationUrlOverride = types.String{Null: true}
	}
	if in.ValidateUsingAllEligibleAtms != nil {
		result.ValidateUsingAllEligibleAtms = types.Bool{Value: *in.ValidateUsingAllEligibleAtms}
	} else {
		result.ValidateUsingAllEligibleAtms = types.Bool{Null: true}
	}

	return &result
}

func flattenRedirectValidationSettings(in *pf.RedirectValidationSettings) *RedirectValidationSettingsData {
	result := RedirectValidationSettingsData{}
	result.Id = types.String{Value: "settings"}
	if in.RedirectValidationLocalSettings != nil {
		result.RedirectValidationLocalSettings = flattenRedirectValidationLocalSettings(in.RedirectValidationLocalSettings)
	}
	if in.RedirectValidationPartnerSettings != nil {
		result.RedirectValidationPartnerSettings = flattenRedirectValidationPartnerSettings(in.RedirectValidationPartnerSettings)
	}

	return &result
}

func flattenMapAttributeFulfillmentValues(in map[string]*pf.AttributeFulfillmentValue) map[string]*AttributeFulfillmentValueData {
	results := map[string]*AttributeFulfillmentValueData{}
	for s, data := range in {
		results[s] = flattenAttributeFulfillmentValue(data)
	}
	return results
}
func flattenAttributeFulfillmentValue(in *pf.AttributeFulfillmentValue) *AttributeFulfillmentValueData {
	result := AttributeFulfillmentValueData{}
	if in.Source != nil {
		result.Source = flattenSourceTypeIdKey(in.Source)
	}
	if in.Value != nil {
		result.Value = types.String{Value: *in.Value}
	}

	return &result
}

func flattenMapBinaryLdapAttributeSettingss(in map[string]*pf.BinaryLdapAttributeSettings) map[string]*BinaryLdapAttributeSettingsData {
	results := map[string]*BinaryLdapAttributeSettingsData{}
	for s, data := range in {
		results[s] = flattenBinaryLdapAttributeSettings(data)
	}
	return results
}
func flattenBinaryLdapAttributeSettings(in *pf.BinaryLdapAttributeSettings) *BinaryLdapAttributeSettingsData {
	result := BinaryLdapAttributeSettingsData{}
	if in.BinaryEncoding != nil {
		result.BinaryEncoding = types.String{Value: *in.BinaryEncoding}
	} else {
		result.BinaryEncoding = types.String{Null: true}
	}

	return &result
}

func flattenClientAuth(in *pf.ClientAuth) *ClientAuthData {
	result := ClientAuthData{}
	if in.ClientCertIssuerDn != nil {
		result.ClientCertIssuerDn = types.String{Value: *in.ClientCertIssuerDn}
	} else {
		result.ClientCertIssuerDn = types.String{Null: true}
	}
	if in.ClientCertSubjectDn != nil {
		result.ClientCertSubjectDn = types.String{Value: *in.ClientCertSubjectDn}
	} else {
		result.ClientCertSubjectDn = types.String{Null: true}
	}
	if in.EncryptedSecret != nil {
		result.EncryptedSecret = types.String{Value: *in.EncryptedSecret}
	} else {
		result.EncryptedSecret = types.String{Null: true}
	}
	if in.EnforceReplayPrevention != nil {
		result.EnforceReplayPrevention = types.Bool{Value: *in.EnforceReplayPrevention}
	} else {
		result.EnforceReplayPrevention = types.Bool{Null: true}
	}
	if in.Secret != nil {
		result.Secret = types.String{Value: *in.Secret}
	} else {
		result.Secret = types.String{Null: true}
	}
	if in.TokenEndpointAuthSigningAlgorithm != nil {
		result.TokenEndpointAuthSigningAlgorithm = types.String{Value: *in.TokenEndpointAuthSigningAlgorithm}
	} else {
		result.TokenEndpointAuthSigningAlgorithm = types.String{Null: true}
	}
	if in.Type != nil {
		result.Type = types.String{Value: *in.Type}
	} else {
		result.Type = types.String{Null: true}
	}

	return &result
}

func flattenClientOIDCPolicy(in *pf.ClientOIDCPolicy) *ClientOIDCPolicyData {
	result := ClientOIDCPolicyData{}
	if in.GrantAccessSessionRevocationApi != nil {
		result.GrantAccessSessionRevocationApi = types.Bool{Value: *in.GrantAccessSessionRevocationApi}
	} else {
		result.GrantAccessSessionRevocationApi = types.Bool{Null: true}
	}
	if in.GrantAccessSessionSessionManagementApi != nil {
		result.GrantAccessSessionSessionManagementApi = types.Bool{Value: *in.GrantAccessSessionSessionManagementApi}
	} else {
		result.GrantAccessSessionSessionManagementApi = types.Bool{Null: true}
	}
	if in.IdTokenContentEncryptionAlgorithm != nil {
		result.IdTokenContentEncryptionAlgorithm = types.String{Value: *in.IdTokenContentEncryptionAlgorithm}
	} else {
		result.IdTokenContentEncryptionAlgorithm = types.String{Null: true}
	}
	if in.IdTokenEncryptionAlgorithm != nil {
		result.IdTokenEncryptionAlgorithm = types.String{Value: *in.IdTokenEncryptionAlgorithm}
	} else {
		result.IdTokenEncryptionAlgorithm = types.String{Null: true}
	}
	if in.IdTokenSigningAlgorithm != nil {
		result.IdTokenSigningAlgorithm = types.String{Value: *in.IdTokenSigningAlgorithm}
	} else {
		result.IdTokenSigningAlgorithm = types.String{Null: true}
	}
	if in.LogoutUris != nil {
		result.LogoutUris = flattenStringList(*in.LogoutUris)
	}
	if in.PairwiseIdentifierUserType != nil {
		result.PairwiseIdentifierUserType = types.Bool{Value: *in.PairwiseIdentifierUserType}
	} else {
		result.PairwiseIdentifierUserType = types.Bool{Null: true}
	}
	if in.PingAccessLogoutCapable != nil {
		result.PingAccessLogoutCapable = types.Bool{Value: *in.PingAccessLogoutCapable}
	} else {
		result.PingAccessLogoutCapable = types.Bool{Null: true}
	}
	if in.PolicyGroup != nil && in.PolicyGroup.Id != nil && *in.PolicyGroup.Id != "" {
		result.PolicyGroup = types.String{Value: *in.PolicyGroup.Id}
	} else {
		result.PolicyGroup = types.String{Null: true}
	}
	if in.SectorIdentifierUri != nil {
		result.SectorIdentifierUri = types.String{Value: *in.SectorIdentifierUri}
	} else {
		result.SectorIdentifierUri = types.String{Null: true}
	}

	return &result
}

func flattenConditionalIssuanceCriteriaEntrys(in *[]*pf.ConditionalIssuanceCriteriaEntry) *[]*ConditionalIssuanceCriteriaEntryData {
	results := []*ConditionalIssuanceCriteriaEntryData{}
	for _, data := range *in {
		results = append(results, flattenConditionalIssuanceCriteriaEntry(data))
	}
	return &results
}
func flattenConditionalIssuanceCriteriaEntry(in *pf.ConditionalIssuanceCriteriaEntry) *ConditionalIssuanceCriteriaEntryData {
	result := ConditionalIssuanceCriteriaEntryData{}
	if in.AttributeName != nil {
		result.AttributeName = types.String{Value: *in.AttributeName}
	}
	if in.Condition != nil {
		result.Condition = types.String{Value: *in.Condition}
	}
	if in.ErrorResult != nil {
		result.ErrorResult = types.String{Value: *in.ErrorResult}
	} else {
		result.ErrorResult = types.String{Null: true}
	}
	if in.Source != nil {
		result.Source = flattenSourceTypeIdKey(in.Source)
	}
	if in.Value != nil {
		result.Value = types.String{Value: *in.Value}
	}

	return &result
}

func flattenCustomAttributeSource(in *pf.CustomAttributeSource) *CustomAttributeSourceData {
	result := CustomAttributeSourceData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	} else {
		result.DataStoreRef = types.String{Null: true}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.FilterFields != nil {
		result.FilterFields = flattenFieldEntrys(in.FilterFields)
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}

	return &result
}

func flattenExpressionIssuanceCriteriaEntrys(in *[]*pf.ExpressionIssuanceCriteriaEntry) *[]*ExpressionIssuanceCriteriaEntryData {
	results := []*ExpressionIssuanceCriteriaEntryData{}
	for _, data := range *in {
		results = append(results, flattenExpressionIssuanceCriteriaEntry(data))
	}
	return &results
}
func flattenExpressionIssuanceCriteriaEntry(in *pf.ExpressionIssuanceCriteriaEntry) *ExpressionIssuanceCriteriaEntryData {
	result := ExpressionIssuanceCriteriaEntryData{}
	if in.ErrorResult != nil {
		result.ErrorResult = types.String{Value: *in.ErrorResult}
	} else {
		result.ErrorResult = types.String{Null: true}
	}
	if in.Expression != nil {
		result.Expression = types.String{Value: *in.Expression}
	}

	return &result
}

func flattenFieldEntrys(in *[]*pf.FieldEntry) *[]*FieldEntryData {
	results := []*FieldEntryData{}
	for _, data := range *in {
		results = append(results, flattenFieldEntry(data))
	}
	return &results
}
func flattenFieldEntry(in *pf.FieldEntry) *FieldEntryData {
	result := FieldEntryData{}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
	}
	if in.Value != nil {
		result.Value = types.String{Value: *in.Value}
	} else {
		result.Value = types.String{Null: true}
	}

	return &result
}

func flattenIssuanceCriteria(in *pf.IssuanceCriteria) *IssuanceCriteriaData {
	result := IssuanceCriteriaData{}
	if in.ConditionalCriteria != nil {
		result.ConditionalCriteria = flattenConditionalIssuanceCriteriaEntrys(in.ConditionalCriteria)
	}
	if in.ExpressionCriteria != nil {
		result.ExpressionCriteria = flattenExpressionIssuanceCriteriaEntrys(in.ExpressionCriteria)
	}

	return &result
}

func flattenJdbcAttributeSource(in *pf.JdbcAttributeSource) *JdbcAttributeSourceData {
	result := JdbcAttributeSourceData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.ColumnNames != nil {
		result.ColumnNames = flattenStringList(*in.ColumnNames)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	} else {
		result.DataStoreRef = types.String{Null: true}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.Filter != nil {
		result.Filter = types.String{Value: *in.Filter}
	} else {
		result.Filter = types.String{Null: true}
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.Schema != nil {
		result.Schema = types.String{Value: *in.Schema}
	} else {
		result.Schema = types.String{Null: true}
	}
	if in.Table != nil {
		result.Table = types.String{Value: *in.Table}
	} else {
		result.Table = types.String{Null: true}
	}

	return &result
}

func flattenJwksSettings(in *pf.JwksSettings) *JwksSettingsData {
	result := JwksSettingsData{}
	if in.Jwks != nil {
		result.Jwks = types.String{Value: *in.Jwks}
	} else {
		result.Jwks = types.String{Null: true}
	}
	if in.JwksUrl != nil {
		result.JwksUrl = types.String{Value: *in.JwksUrl}
	} else {
		result.JwksUrl = types.String{Null: true}
	}

	return &result
}

func flattenLdapAttributeSource(in *pf.LdapAttributeSource) *LdapAttributeSourceData {
	result := LdapAttributeSourceData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.BaseDn != nil {
		result.BaseDn = types.String{Value: *in.BaseDn}
	} else {
		result.BaseDn = types.String{Null: true}
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = flattenMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	} else {
		result.DataStoreRef = types.String{Null: true}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.MemberOfNestedGroup != nil {
		result.MemberOfNestedGroup = types.Bool{Value: *in.MemberOfNestedGroup}
	} else {
		result.MemberOfNestedGroup = types.Bool{Null: true}
	}
	if in.SearchAttributes != nil {
		result.SearchAttributes = flattenStringList(*in.SearchAttributes)
	}
	if in.SearchFilter != nil {
		result.SearchFilter = types.String{Value: *in.SearchFilter}
	} else {
		result.SearchFilter = types.String{Null: true}
	}
	if in.SearchScope != nil {
		result.SearchScope = types.String{Value: *in.SearchScope}
	} else {
		result.SearchScope = types.String{Null: true}
	}

	return &result
}

func flattenMapParameterValuess(in map[string]*pf.ParameterValues) map[string]*ParameterValuesData {
	results := map[string]*ParameterValuesData{}
	for s, data := range in {
		results[s] = flattenParameterValues(data)
	}
	return results
}
func flattenParameterValues(in *pf.ParameterValues) *ParameterValuesData {
	result := ParameterValuesData{}
	if in.Values != nil {
		result.Values = flattenStringList(*in.Values)
	}

	return &result
}

func flattenRedirectValidationLocalSettings(in *pf.RedirectValidationLocalSettings) *RedirectValidationLocalSettingsData {
	result := RedirectValidationLocalSettingsData{}
	if in.EnableInErrorResourceValidation != nil {
		result.EnableInErrorResourceValidation = types.Bool{Value: *in.EnableInErrorResourceValidation}
	} else {
		result.EnableInErrorResourceValidation = types.Bool{Null: true}
	}
	if in.EnableTargetResourceValidationForIdpDiscovery != nil {
		result.EnableTargetResourceValidationForIdpDiscovery = types.Bool{Value: *in.EnableTargetResourceValidationForIdpDiscovery}
	} else {
		result.EnableTargetResourceValidationForIdpDiscovery = types.Bool{Null: true}
	}
	if in.EnableTargetResourceValidationForSLO != nil {
		result.EnableTargetResourceValidationForSLO = types.Bool{Value: *in.EnableTargetResourceValidationForSLO}
	} else {
		result.EnableTargetResourceValidationForSLO = types.Bool{Null: true}
	}
	if in.EnableTargetResourceValidationForSSO != nil {
		result.EnableTargetResourceValidationForSSO = types.Bool{Value: *in.EnableTargetResourceValidationForSSO}
	} else {
		result.EnableTargetResourceValidationForSSO = types.Bool{Null: true}
	}
	if in.WhiteList != nil {
		result.WhiteList = flattenRedirectValidationSettingsWhitelistEntrys(in.WhiteList)
	}

	return &result
}

func flattenRedirectValidationPartnerSettings(in *pf.RedirectValidationPartnerSettings) *RedirectValidationPartnerSettingsData {
	result := RedirectValidationPartnerSettingsData{}
	if in.EnableWreplyValidationSLO != nil {
		result.EnableWreplyValidationSLO = types.Bool{Value: *in.EnableWreplyValidationSLO}
	} else {
		result.EnableWreplyValidationSLO = types.Bool{Null: true}
	}

	return &result
}

func flattenRedirectValidationSettingsWhitelistEntrys(in *[]*pf.RedirectValidationSettingsWhitelistEntry) *[]*RedirectValidationSettingsWhitelistEntryData {
	results := []*RedirectValidationSettingsWhitelistEntryData{}
	for _, data := range *in {
		results = append(results, flattenRedirectValidationSettingsWhitelistEntry(data))
	}
	return &results
}
func flattenRedirectValidationSettingsWhitelistEntry(in *pf.RedirectValidationSettingsWhitelistEntry) *RedirectValidationSettingsWhitelistEntryData {
	result := RedirectValidationSettingsWhitelistEntryData{}
	if in.AllowQueryAndFragment != nil {
		result.AllowQueryAndFragment = types.Bool{Value: *in.AllowQueryAndFragment}
	} else {
		result.AllowQueryAndFragment = types.Bool{Null: true}
	}
	if in.IdpDiscovery != nil {
		result.IdpDiscovery = types.Bool{Value: *in.IdpDiscovery}
	} else {
		result.IdpDiscovery = types.Bool{Null: true}
	}
	if in.InErrorResource != nil {
		result.InErrorResource = types.Bool{Value: *in.InErrorResource}
	} else {
		result.InErrorResource = types.Bool{Null: true}
	}
	if in.RequireHttps != nil {
		result.RequireHttps = types.Bool{Value: *in.RequireHttps}
	} else {
		result.RequireHttps = types.Bool{Null: true}
	}
	if in.TargetResourceSLO != nil {
		result.TargetResourceSLO = types.Bool{Value: *in.TargetResourceSLO}
	} else {
		result.TargetResourceSLO = types.Bool{Null: true}
	}
	if in.TargetResourceSSO != nil {
		result.TargetResourceSSO = types.Bool{Value: *in.TargetResourceSSO}
	} else {
		result.TargetResourceSSO = types.Bool{Null: true}
	}
	if in.ValidDomain != nil {
		result.ValidDomain = types.String{Value: *in.ValidDomain}
	}
	if in.ValidPath != nil {
		result.ValidPath = types.String{Value: *in.ValidPath}
	} else {
		result.ValidPath = types.String{Null: true}
	}

	return &result
}

func flattenSourceTypeIdKey(in *pf.SourceTypeIdKey) *SourceTypeIdKeyData {
	result := SourceTypeIdKeyData{}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.Type != nil {
		result.Type = types.String{Value: *in.Type}
	}

	return &result
}

func flattenJdbcAttributeSources(in *[]*pf.AttributeSource) []JdbcAttributeSourceData {
	results := []JdbcAttributeSourceData{}
	for _, source := range *in {
		if source.Type != nil && *source.Type == "JDBC" {
			source.JdbcAttributeSource.DataStoreRef = source.DataStoreRef
			source.JdbcAttributeSource.AttributeContractFulfillment = source.AttributeContractFulfillment
			source.JdbcAttributeSource.Id = source.Id
			source.JdbcAttributeSource.Description = source.Description
			results = append(results, *flattenJdbcAttributeSource(&source.JdbcAttributeSource))
		}
	}
	return results
}

func flattenLdapAttributeSources(in *[]*pf.AttributeSource) []LdapAttributeSourceData {
	results := []LdapAttributeSourceData{}
	for _, source := range *in {
		if source.Type != nil && *source.Type == "LDAP" {
			results = append(results, *flattenLdapAttributeSource(&source.LdapAttributeSource))
		}
	}
	return results
}

func flattenCustomAttributeSources(in *[]*pf.AttributeSource) []CustomAttributeSourceData {
	results := []CustomAttributeSourceData{}
	for _, source := range *in {
		if source.Type != nil && *source.Type == "CUSTOM" {
			results = append(results, *flattenCustomAttributeSource(&source.CustomAttributeSource))
		}
	}
	return results
}

func flattenStringList(in []*string) []types.String {
	results := []types.String{}
	for _, s := range in {
		results = append(results, types.String{Value: *s})
	}
	return results
}

func issuanceCriteriaShouldFlatten(in *pf.IssuanceCriteria) bool {
	if in.ExpressionCriteria != nil && len(*in.ExpressionCriteria) > 0 {
		return true
	}
	if in.ConditionalCriteria != nil && len(*in.ConditionalCriteria) > 0 {
		return true
	}
	return false
}
