package framework

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
		result.AuthenticationPolicyContractRef = types.StringValue(*in.AuthenticationPolicyContractRef.Id)
	} else {
		result.AuthenticationPolicyContractRef = types.StringNull()
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		result.IssuanceCriteria = flattenIssuanceCriteria(in.IssuanceCriteria)
	}

	return &result
}

func flattenApplicationSessionPolicy(in *pf.ApplicationSessionPolicy) *ApplicationSessionPolicyData {
	result := ApplicationSessionPolicyData{}
	if in.IdleTimeoutMins != nil {
		result.IdleTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.IdleTimeoutMins)))
	} else {
		result.IdleTimeoutMins = types.NumberNull()
	}
	if in.MaxTimeoutMins != nil {
		result.MaxTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.MaxTimeoutMins)))
	} else {
		result.MaxTimeoutMins = types.NumberNull()
	}

	return &result
}

func flattenAuthenticationPolicyContract(in *pf.AuthenticationPolicyContract) *AuthenticationPolicyContractData {
	result := AuthenticationPolicyContractData{}
	if in.CoreAttributes != nil {
		var values []attr.Value
		for _, s := range *in.CoreAttributes {
			values = append(values, types.StringValue(*s.Name))
		}
		result.CoreAttributes = types.SetValueMust(types.StringType, values)
	}

	if in.ExtendedAttributes != nil {
		var values []attr.Value
		for _, s := range *in.ExtendedAttributes {
			values = append(values, types.StringValue(*s.Name))
		}
		result.ExtendedAttributes = types.SetValueMust(types.StringType, values)
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}

	return &result
}

func flattenAuthenticationSessionPolicy(in *pf.AuthenticationSessionPolicy) *AuthenticationSessionPolicyData {
	result := AuthenticationSessionPolicyData{}
	if in.AuthenticationSource != nil {
		result.AuthenticationSource = flattenAuthenticationSource(in.AuthenticationSource)
	}
	if in.AuthnContextSensitive != nil {
		result.AuthnContextSensitive = types.BoolValue(*in.AuthnContextSensitive)
	} else {
		result.AuthnContextSensitive = types.BoolNull()
	}
	if in.EnableSessions != nil {
		result.EnableSessions = types.BoolValue(*in.EnableSessions)
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.IdleTimeoutMins != nil {
		result.IdleTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.IdleTimeoutMins)))
	} else {
		result.IdleTimeoutMins = types.NumberNull()
	}
	if in.MaxTimeoutMins != nil {
		result.MaxTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.MaxTimeoutMins)))
	} else {
		result.MaxTimeoutMins = types.NumberNull()
	}
	if in.Persistent != nil {
		result.Persistent = types.BoolValue(*in.Persistent)
	} else {
		result.Persistent = types.BoolNull()
	}
	if in.TimeoutDisplayUnit != nil {
		result.TimeoutDisplayUnit = types.StringValue(*in.TimeoutDisplayUnit)
	} else {
		result.TimeoutDisplayUnit = types.StringNull()
	}

	return &result
}

func flattenAuthorizationServerSettings(in *pf.AuthorizationServerSettings) *AuthorizationServerSettingsData {
	result := AuthorizationServerSettingsData{}
	if in.ActivationCodeCheckMode != nil {
		result.ActivationCodeCheckMode = types.StringValue(*in.ActivationCodeCheckMode)
	} else {
		result.ActivationCodeCheckMode = types.StringNull()
	}
	if in.AdminWebServicePcvRef != nil && in.AdminWebServicePcvRef.Id != nil && *in.AdminWebServicePcvRef.Id != "" {
		result.AdminWebServicePcvRef = types.StringValue(*in.AdminWebServicePcvRef.Id)
	} else {
		result.AdminWebServicePcvRef = types.StringNull()
	}
	if in.AllowUnidentifiedClientExtensionGrants != nil {
		result.AllowUnidentifiedClientExtensionGrants = types.BoolValue(*in.AllowUnidentifiedClientExtensionGrants)
	} else {
		result.AllowUnidentifiedClientExtensionGrants = types.BoolNull()
	}
	if in.AllowUnidentifiedClientROCreds != nil {
		result.AllowUnidentifiedClientROCreds = types.BoolValue(*in.AllowUnidentifiedClientROCreds)
	} else {
		result.AllowUnidentifiedClientROCreds = types.BoolNull()
	}
	if in.AllowedOrigins != nil {
		result.AllowedOrigins = flattenStringList(*in.AllowedOrigins)
	} else {
		result.AllowedOrigins = types.ListNull(types.StringType)
	}
	if in.ApprovedScopesAttribute != nil {
		result.ApprovedScopesAttribute = types.StringValue(*in.ApprovedScopesAttribute)
	} else {
		result.ApprovedScopesAttribute = types.StringNull()
	}
	if in.AtmIdForOAuthGrantManagement != nil {
		result.AtmIdForOAuthGrantManagement = types.StringValue(*in.AtmIdForOAuthGrantManagement)
	} else {
		result.AtmIdForOAuthGrantManagement = types.StringNull()
	}
	if in.AuthorizationCodeEntropy != nil {
		result.AuthorizationCodeEntropy = types.NumberValue(big.NewFloat(float64(*in.AuthorizationCodeEntropy)))
	}
	if in.AuthorizationCodeTimeout != nil {
		result.AuthorizationCodeTimeout = types.NumberValue(big.NewFloat(float64(*in.AuthorizationCodeTimeout)))
	}
	if in.BypassActivationCodeConfirmation != nil {
		result.BypassActivationCodeConfirmation = types.BoolValue(*in.BypassActivationCodeConfirmation)
	}
	if in.BypassAuthorizationForApprovedGrants != nil {
		result.BypassAuthorizationForApprovedGrants = types.BoolValue(*in.BypassAuthorizationForApprovedGrants)
	} else {
		result.BypassAuthorizationForApprovedGrants = types.BoolNull()
	}
	if in.ClientSecretRetentionPeriod != nil {
		result.ClientSecretRetentionPeriod = types.NumberValue(big.NewFloat(float64(*in.ClientSecretRetentionPeriod)))
	} else {
		result.ClientSecretRetentionPeriod = types.NumberNull()
	}
	if in.DefaultScopeDescription != nil {
		result.DefaultScopeDescription = types.StringValue(*in.DefaultScopeDescription)
	}
	if in.DevicePollingInterval != nil {
		result.DevicePollingInterval = types.NumberValue(big.NewFloat(float64(*in.DevicePollingInterval)))
	}
	if in.DisallowPlainPKCE != nil {
		result.DisallowPlainPKCE = types.BoolValue(*in.DisallowPlainPKCE)
	} else {
		result.DisallowPlainPKCE = types.BoolNull()
	}
	if in.ExclusiveScopeGroups != nil {
		result.ExclusiveScopeGroups = flattenScopeGroupEntrys(in.ExclusiveScopeGroups)
	}
	if in.ExclusiveScopes != nil {
		result.ExclusiveScopes = flattenScopeEntrys(in.ExclusiveScopes)
	}
	if in.IncludeIssuerInAuthorizationResponse != nil {
		result.IncludeIssuerInAuthorizationResponse = types.BoolValue(*in.IncludeIssuerInAuthorizationResponse)
	} else {
		result.IncludeIssuerInAuthorizationResponse = types.BoolNull()
	}
	if in.JwtSecuredAuthorizationResponseModeLifetime != nil {
		result.JwtSecuredAuthorizationResponseModeLifetime = types.NumberValue(big.NewFloat(float64(*in.JwtSecuredAuthorizationResponseModeLifetime)))
	} else {
		result.JwtSecuredAuthorizationResponseModeLifetime = types.NumberNull()
	}
	if in.ParReferenceLength != nil {
		result.ParReferenceLength = types.NumberValue(big.NewFloat(float64(*in.ParReferenceLength)))
	} else {
		result.ParReferenceLength = types.NumberNull()
	}
	if in.ParReferenceTimeout != nil {
		result.ParReferenceTimeout = types.NumberValue(big.NewFloat(float64(*in.ParReferenceTimeout)))
	} else {
		result.ParReferenceTimeout = types.NumberNull()
	}
	if in.ParStatus != nil {
		result.ParStatus = types.StringValue(*in.ParStatus)
	} else {
		result.ParStatus = types.StringNull()
	}
	if in.PendingAuthorizationTimeout != nil {
		result.PendingAuthorizationTimeout = types.NumberValue(big.NewFloat(float64(*in.PendingAuthorizationTimeout)))
	}
	if in.PersistentGrantContract != nil {
		result.PersistentGrantContract = flattenPersistentGrantContract(in.PersistentGrantContract)
	}
	if in.PersistentGrantIdleTimeout != nil {
		result.PersistentGrantIdleTimeout = types.NumberValue(big.NewFloat(float64(*in.PersistentGrantIdleTimeout)))
	} else {
		result.PersistentGrantIdleTimeout = types.NumberNull()
	}
	if in.PersistentGrantIdleTimeoutTimeUnit != nil {
		result.PersistentGrantIdleTimeoutTimeUnit = types.StringValue(*in.PersistentGrantIdleTimeoutTimeUnit)
	} else {
		result.PersistentGrantIdleTimeoutTimeUnit = types.StringNull()
	}
	if in.PersistentGrantLifetime != nil {
		result.PersistentGrantLifetime = types.NumberValue(big.NewFloat(float64(*in.PersistentGrantLifetime)))
	} else {
		result.PersistentGrantLifetime = types.NumberNull()
	}
	if in.PersistentGrantLifetimeUnit != nil {
		result.PersistentGrantLifetimeUnit = types.StringValue(*in.PersistentGrantLifetimeUnit)
	} else {
		result.PersistentGrantLifetimeUnit = types.StringNull()
	}
	if in.PersistentGrantReuseGrantTypes != nil {
		result.PersistentGrantReuseGrantTypes = flattenStringList(*in.PersistentGrantReuseGrantTypes)
	} else {
		result.PersistentGrantReuseGrantTypes = types.ListNull(types.StringType)
	}
	if in.RefreshRollingInterval != nil {
		result.RefreshRollingInterval = types.NumberValue(big.NewFloat(float64(*in.RefreshRollingInterval)))
	}
	if in.RefreshTokenLength != nil {
		result.RefreshTokenLength = types.NumberValue(big.NewFloat(float64(*in.RefreshTokenLength)))
	}
	if in.RefreshTokenRollingGracePeriod != nil {
		result.RefreshTokenRollingGracePeriod = types.NumberValue(big.NewFloat(float64(*in.RefreshTokenRollingGracePeriod)))
	} else {
		result.RefreshTokenRollingGracePeriod = types.NumberNull()
	}
	if in.RegisteredAuthorizationPath != nil {
		result.RegisteredAuthorizationPath = types.StringValue(*in.RegisteredAuthorizationPath)
	}
	if in.RollRefreshTokenValues != nil {
		result.RollRefreshTokenValues = types.BoolValue(*in.RollRefreshTokenValues)
	} else {
		result.RollRefreshTokenValues = types.BoolNull()
	}
	if in.ScopeForOAuthGrantManagement != nil {
		result.ScopeForOAuthGrantManagement = types.StringValue(*in.ScopeForOAuthGrantManagement)
	} else {
		result.ScopeForOAuthGrantManagement = types.StringNull()
	}
	if in.ScopeGroups != nil {
		result.ScopeGroups = flattenScopeGroupEntrys(in.ScopeGroups)
	}
	if in.Scopes != nil {
		result.Scopes = flattenScopeEntrys(in.Scopes)
	}
	if in.TokenEndpointBaseUrl != nil {
		result.TokenEndpointBaseUrl = types.StringValue(*in.TokenEndpointBaseUrl)
	} else {
		result.TokenEndpointBaseUrl = types.StringNull()
	}
	if in.TrackUserSessionsForLogout != nil {
		result.TrackUserSessionsForLogout = types.BoolValue(*in.TrackUserSessionsForLogout)
	} else {
		result.TrackUserSessionsForLogout = types.BoolNull()
	}
	if in.UserAuthorizationConsentAdapter != nil {
		result.UserAuthorizationConsentAdapter = types.StringValue(*in.UserAuthorizationConsentAdapter)
	} else {
		result.UserAuthorizationConsentAdapter = types.StringNull()
	}
	if in.UserAuthorizationConsentPageSetting != nil {
		result.UserAuthorizationConsentPageSetting = types.StringValue(*in.UserAuthorizationConsentPageSetting)
	} else {
		result.UserAuthorizationConsentPageSetting = types.StringNull()
	}
	if in.UserAuthorizationUrl != nil {
		result.UserAuthorizationUrl = types.StringValue(*in.UserAuthorizationUrl)
	} else {
		result.UserAuthorizationUrl = types.StringNull()
	}

	return &result
}

func flattenClient(in *pf.Client) *ClientData {
	result := ClientData{}
	if in.AllowAuthenticationApiInit != nil {
		result.AllowAuthenticationApiInit = types.BoolValue(*in.AllowAuthenticationApiInit)
	} else {
		result.AllowAuthenticationApiInit = types.BoolNull()
	}
	if in.BypassActivationCodeConfirmationOverride != nil {
		result.BypassActivationCodeConfirmationOverride = types.BoolValue(*in.BypassActivationCodeConfirmationOverride)
	} else {
		result.BypassActivationCodeConfirmationOverride = types.BoolNull()
	}
	if in.BypassApprovalPage != nil {
		result.BypassApprovalPage = types.BoolValue(*in.BypassApprovalPage)
	} else {
		result.BypassApprovalPage = types.BoolNull()
	}
	if in.CibaDeliveryMode != nil {
		result.CibaDeliveryMode = types.StringValue(*in.CibaDeliveryMode)
	} else {
		result.CibaDeliveryMode = types.StringNull()
	}
	if in.CibaNotificationEndpoint != nil {
		result.CibaNotificationEndpoint = types.StringValue(*in.CibaNotificationEndpoint)
	} else {
		result.CibaNotificationEndpoint = types.StringNull()
	}
	if in.CibaPollingInterval != nil {
		result.CibaPollingInterval = types.NumberValue(big.NewFloat(float64(*in.CibaPollingInterval)))
	} else {
		result.CibaPollingInterval = types.NumberNull()
	}
	if in.CibaRequestObjectSigningAlgorithm != nil {
		result.CibaRequestObjectSigningAlgorithm = types.StringValue(*in.CibaRequestObjectSigningAlgorithm)
	} else {
		result.CibaRequestObjectSigningAlgorithm = types.StringNull()
	}
	if in.CibaRequireSignedRequests != nil {
		result.CibaRequireSignedRequests = types.BoolValue(*in.CibaRequireSignedRequests)
	} else {
		result.CibaRequireSignedRequests = types.BoolNull()
	}
	if in.CibaUserCodeSupported != nil {
		result.CibaUserCodeSupported = types.BoolValue(*in.CibaUserCodeSupported)
	} else {
		result.CibaUserCodeSupported = types.BoolNull()
	}
	if in.ClientAuth != nil && *in.ClientAuth.Type != "NONE" {
		result.ClientAuth = flattenClientAuth(in.ClientAuth)
	}
	if in.ClientId != nil {
		result.ClientId = types.StringValue(*in.ClientId)
	}
	if in.ClientSecretChangedTime != nil {
		result.ClientSecretChangedTime = types.StringValue(*in.ClientSecretChangedTime)
	} else {
		result.ClientSecretChangedTime = types.StringNull()
	}
	if in.ClientSecretRetentionPeriod != nil {
		result.ClientSecretRetentionPeriod = types.NumberValue(big.NewFloat(float64(*in.ClientSecretRetentionPeriod)))
	} else {
		result.ClientSecretRetentionPeriod = types.NumberNull()
	}
	if in.ClientSecretRetentionPeriodType != nil {
		result.ClientSecretRetentionPeriodType = types.StringValue(*in.ClientSecretRetentionPeriodType)
	} else {
		result.ClientSecretRetentionPeriodType = types.StringNull()
	}
	if in.DefaultAccessTokenManagerRef != nil && in.DefaultAccessTokenManagerRef.Id != nil && *in.DefaultAccessTokenManagerRef.Id != "" {
		result.DefaultAccessTokenManagerRef = types.StringValue(*in.DefaultAccessTokenManagerRef.Id)
	} else {
		result.DefaultAccessTokenManagerRef = types.StringNull()
	}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	} else {
		result.Description = types.StringNull()
	}
	if in.DeviceFlowSettingType != nil {
		result.DeviceFlowSettingType = types.StringValue(*in.DeviceFlowSettingType)
	} else {
		result.DeviceFlowSettingType = types.StringNull()
	}
	if in.DevicePollingIntervalOverride != nil {
		result.DevicePollingIntervalOverride = types.NumberValue(big.NewFloat(float64(*in.DevicePollingIntervalOverride)))
	} else {
		result.DevicePollingIntervalOverride = types.NumberNull()
	}
	if in.Enabled != nil {
		result.Enabled = types.BoolValue(*in.Enabled)
	} else {
		result.Enabled = types.BoolNull()
	}
	if in.ExclusiveScopes != nil {
		result.ExclusiveScopes = flattenStringList(*in.ExclusiveScopes)
	} else {
		result.ExclusiveScopes = types.ListNull(types.StringType)
	}
	if in.ExtendedParameters != nil {
		result.ExtendedParameters = flattenMapParameterValuess(in.ExtendedParameters)
	}
	if in.GrantTypes != nil {
		result.GrantTypes = flattenStringList(*in.GrantTypes)
	} else {
		result.GrantTypes = types.ListNull(types.StringType)
	}
	if in.ClientId != nil {
		result.Id = types.StringValue(*in.ClientId)
	}
	if in.JwksSettings != nil {
		result.JwksSettings = flattenJwksSettings(in.JwksSettings)
	}
	if in.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm != nil {
		result.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm = types.StringValue(*in.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm)
	} else {
		result.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm = types.StringNull()
	}
	if in.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm != nil {
		result.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm = types.StringValue(*in.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm)
	} else {
		result.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm = types.StringNull()
	}
	if in.JwtSecuredAuthorizationResponseModeSigningAlgorithm != nil {
		result.JwtSecuredAuthorizationResponseModeSigningAlgorithm = types.StringValue(*in.JwtSecuredAuthorizationResponseModeSigningAlgorithm)
	} else {
		result.JwtSecuredAuthorizationResponseModeSigningAlgorithm = types.StringNull()
	}
	if in.LogoUrl != nil {
		result.LogoUrl = types.StringValue(*in.LogoUrl)
	} else {
		result.LogoUrl = types.StringNull()
	}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}
	if in.OidcPolicy != nil {
		result.OidcPolicy = flattenClientOIDCPolicy(in.OidcPolicy)
	}
	if in.PendingAuthorizationTimeoutOverride != nil {
		result.PendingAuthorizationTimeoutOverride = types.NumberValue(big.NewFloat(float64(*in.PendingAuthorizationTimeoutOverride)))
	} else {
		result.PendingAuthorizationTimeoutOverride = types.NumberNull()
	}
	if in.PersistentGrantExpirationTime != nil {
		result.PersistentGrantExpirationTime = types.NumberValue(big.NewFloat(float64(*in.PersistentGrantExpirationTime)))
	} else {
		result.PersistentGrantExpirationTime = types.NumberNull()
	}
	if in.PersistentGrantExpirationTimeUnit != nil {
		result.PersistentGrantExpirationTimeUnit = types.StringValue(*in.PersistentGrantExpirationTimeUnit)
	} else {
		result.PersistentGrantExpirationTimeUnit = types.StringValue("DAYS")
	}
	if in.PersistentGrantExpirationType != nil {
		result.PersistentGrantExpirationType = types.StringValue(*in.PersistentGrantExpirationType)
	} else {
		result.PersistentGrantExpirationType = types.StringNull()
	}
	if in.PersistentGrantIdleTimeout != nil {
		result.PersistentGrantIdleTimeout = types.NumberValue(big.NewFloat(float64(*in.PersistentGrantIdleTimeout)))
	} else {
		result.PersistentGrantIdleTimeout = types.NumberNull()
	}
	if in.PersistentGrantIdleTimeoutTimeUnit != nil {
		result.PersistentGrantIdleTimeoutTimeUnit = types.StringValue(*in.PersistentGrantIdleTimeoutTimeUnit)
	} else {
		result.PersistentGrantIdleTimeoutTimeUnit = types.StringNull()
	}
	if in.PersistentGrantIdleTimeoutType != nil {
		result.PersistentGrantIdleTimeoutType = types.StringValue(*in.PersistentGrantIdleTimeoutType)
	} else {
		result.PersistentGrantIdleTimeoutType = types.StringNull()
	}
	if in.PersistentGrantReuseGrantTypes != nil {
		result.PersistentGrantReuseGrantTypes = flattenStringList(*in.PersistentGrantReuseGrantTypes)
	} else {
		result.PersistentGrantReuseGrantTypes = types.ListNull(types.StringType)
	}
	if in.PersistentGrantReuseType != nil {
		result.PersistentGrantReuseType = types.StringValue(*in.PersistentGrantReuseType)
	} else {
		result.PersistentGrantReuseType = types.StringNull()
	}
	if in.RedirectUris != nil {
		result.RedirectUris = flattenStringSet(*in.RedirectUris)
	} else {
		result.RedirectUris = types.SetNull(types.StringType)
	}
	if in.RefreshRolling != nil {
		result.RefreshRolling = types.StringValue(*in.RefreshRolling)
	} else {
		result.RefreshRolling = types.StringNull()
	}
	if in.RefreshTokenRollingGracePeriod != nil {
		result.RefreshTokenRollingGracePeriod = types.NumberValue(big.NewFloat(float64(*in.RefreshTokenRollingGracePeriod)))
	} else {
		result.RefreshTokenRollingGracePeriod = types.NumberNull()
	}
	if in.RefreshTokenRollingGracePeriodType != nil {
		result.RefreshTokenRollingGracePeriodType = types.StringValue(*in.RefreshTokenRollingGracePeriodType)
	} else {
		result.RefreshTokenRollingGracePeriodType = types.StringNull()
	}
	if in.RefreshTokenRollingInterval != nil {
		result.RefreshTokenRollingInterval = types.NumberValue(big.NewFloat(float64(*in.RefreshTokenRollingInterval)))
	} else {
		result.RefreshTokenRollingInterval = types.NumberNull()
	}
	if in.RefreshTokenRollingIntervalType != nil {
		result.RefreshTokenRollingIntervalType = types.StringValue(*in.RefreshTokenRollingIntervalType)
	} else {
		result.RefreshTokenRollingIntervalType = types.StringNull()
	}
	if in.RequestObjectSigningAlgorithm != nil {
		result.RequestObjectSigningAlgorithm = types.StringValue(*in.RequestObjectSigningAlgorithm)
	} else {
		result.RequestObjectSigningAlgorithm = types.StringNull()
	}
	if in.RequestPolicyRef != nil && in.RequestPolicyRef.Id != nil && *in.RequestPolicyRef.Id != "" {
		result.RequestPolicyRef = types.StringValue(*in.RequestPolicyRef.Id)
	} else {
		result.RequestPolicyRef = types.StringNull()
	}
	if in.RequireJwtSecuredAuthorizationResponseMode != nil {
		result.RequireJwtSecuredAuthorizationResponseMode = types.BoolValue(*in.RequireJwtSecuredAuthorizationResponseMode)
	} else {
		result.RequireJwtSecuredAuthorizationResponseMode = types.BoolNull()
	}
	if in.RequireProofKeyForCodeExchange != nil {
		result.RequireProofKeyForCodeExchange = types.BoolValue(*in.RequireProofKeyForCodeExchange)
	} else {
		result.RequireProofKeyForCodeExchange = types.BoolNull()
	}
	if in.RequirePushedAuthorizationRequests != nil {
		result.RequirePushedAuthorizationRequests = types.BoolValue(*in.RequirePushedAuthorizationRequests)
	} else {
		result.RequirePushedAuthorizationRequests = types.BoolNull()
	}
	if in.RequireSignedRequests != nil {
		result.RequireSignedRequests = types.BoolValue(*in.RequireSignedRequests)
	} else {
		result.RequireSignedRequests = types.BoolNull()
	}
	if in.RestrictScopes != nil {
		result.RestrictScopes = types.BoolValue(*in.RestrictScopes)
	} else {
		result.RestrictScopes = types.BoolNull()
	}
	if in.RestrictToDefaultAccessTokenManager != nil {
		result.RestrictToDefaultAccessTokenManager = types.BoolValue(*in.RestrictToDefaultAccessTokenManager)
	} else {
		result.RestrictToDefaultAccessTokenManager = types.BoolNull()
	}
	if in.RestrictedResponseTypes != nil {
		result.RestrictedResponseTypes = flattenStringList(*in.RestrictedResponseTypes)
	} else {
		result.RestrictedResponseTypes = types.ListNull(types.StringType)
	}
	if in.RestrictedScopes != nil {
		result.RestrictedScopes = flattenStringSet(*in.RestrictedScopes)
	} else {
		result.RestrictedScopes = types.SetNull(types.StringType)
	}
	if in.TokenExchangeProcessorPolicyRef != nil && in.TokenExchangeProcessorPolicyRef.Id != nil && *in.TokenExchangeProcessorPolicyRef.Id != "" {
		result.TokenExchangeProcessorPolicyRef = types.StringValue(*in.TokenExchangeProcessorPolicyRef.Id)
	} else {
		result.TokenExchangeProcessorPolicyRef = types.StringNull()
	}
	if in.TokenIntrospectionContentEncryptionAlgorithm != nil {
		result.TokenIntrospectionContentEncryptionAlgorithm = types.StringValue(*in.TokenIntrospectionContentEncryptionAlgorithm)
	} else {
		result.TokenIntrospectionContentEncryptionAlgorithm = types.StringNull()
	}
	if in.TokenIntrospectionEncryptionAlgorithm != nil {
		result.TokenIntrospectionEncryptionAlgorithm = types.StringValue(*in.TokenIntrospectionEncryptionAlgorithm)
	} else {
		result.TokenIntrospectionEncryptionAlgorithm = types.StringNull()
	}
	if in.TokenIntrospectionSigningAlgorithm != nil {
		result.TokenIntrospectionSigningAlgorithm = types.StringValue(*in.TokenIntrospectionSigningAlgorithm)
	} else {
		result.TokenIntrospectionSigningAlgorithm = types.StringNull()
	}
	if in.UserAuthorizationUrlOverride != nil {
		result.UserAuthorizationUrlOverride = types.StringValue(*in.UserAuthorizationUrlOverride)
	} else {
		result.UserAuthorizationUrlOverride = types.StringNull()
	}
	if in.ValidateUsingAllEligibleAtms != nil {
		result.ValidateUsingAllEligibleAtms = types.BoolValue(*in.ValidateUsingAllEligibleAtms)
	} else {
		result.ValidateUsingAllEligibleAtms = types.BoolNull()
	}

	return &result
}

func flattenGlobalAuthenticationSessionPolicy(in *pf.GlobalAuthenticationSessionPolicy) *GlobalAuthenticationSessionPolicyData {
	result := GlobalAuthenticationSessionPolicyData{}
	if in.EnableSessions != nil {
		result.EnableSessions = types.BoolValue(*in.EnableSessions)
	}
	if in.HashUniqueUserKeyAttribute != nil {
		result.HashUniqueUserKeyAttribute = types.BoolValue(*in.HashUniqueUserKeyAttribute)
	} else {
		result.HashUniqueUserKeyAttribute = types.BoolNull()
	}
	if in.IdleTimeoutDisplayUnit != nil {
		result.IdleTimeoutDisplayUnit = types.StringValue(*in.IdleTimeoutDisplayUnit)
	} else {
		result.IdleTimeoutDisplayUnit = types.StringNull()
	}
	if in.IdleTimeoutMins != nil {
		result.IdleTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.IdleTimeoutMins)))
	} else {
		result.IdleTimeoutMins = types.NumberNull()
	}
	if in.MaxTimeoutDisplayUnit != nil {
		result.MaxTimeoutDisplayUnit = types.StringValue(*in.MaxTimeoutDisplayUnit)
	} else {
		result.MaxTimeoutDisplayUnit = types.StringNull()
	}
	if in.MaxTimeoutMins != nil {
		result.MaxTimeoutMins = types.NumberValue(big.NewFloat(float64(*in.MaxTimeoutMins)))
	} else {
		result.MaxTimeoutMins = types.NumberNull()
	}
	if in.PersistentSessions != nil {
		result.PersistentSessions = types.BoolValue(*in.PersistentSessions)
	} else {
		result.PersistentSessions = types.BoolNull()
	}

	return &result
}

func flattenMetadataUrl(in *pf.MetadataUrl) *MetadataUrlData {
	result := MetadataUrlData{}
	if in.CertView != nil {
		flat := flattenCertView(in.CertView)
		result.CertView, _ = types.ObjectValueFrom(context.Background(), certViewAttrTypes, *flat)
	} else {
		result.CertView = types.ObjectNull(certViewAttrTypes)
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}
	if in.Url != nil {
		result.Url = types.StringValue(*in.Url)
	}
	if in.ValidateSignature != nil {
		result.ValidateSignature = types.BoolValue(*in.ValidateSignature)
	} else {
		result.ValidateSignature = types.BoolNull()
	}
	if in.X509File != nil {
		result.X509File = flattenX509File(in.X509File)
	}

	return &result
}

func flattenRedirectValidationSettings(in *pf.RedirectValidationSettings) *RedirectValidationSettingsData {
	result := RedirectValidationSettingsData{}
	if in.RedirectValidationLocalSettings != nil {
		result.RedirectValidationLocalSettings = flattenRedirectValidationLocalSettings(in.RedirectValidationLocalSettings)
	}
	if in.RedirectValidationPartnerSettings != nil {
		result.RedirectValidationPartnerSettings = flattenRedirectValidationPartnerSettings(in.RedirectValidationPartnerSettings)
	}

	return &result
}

func flattenSessionSettings(in *pf.SessionSettings) *SessionSettingsData {
	result := SessionSettingsData{}
	if in.RevokeUserSessionOnLogout != nil {
		result.RevokeUserSessionOnLogout = types.BoolValue(*in.RevokeUserSessionOnLogout)
	} else {
		result.RevokeUserSessionOnLogout = types.BoolNull()
	}
	if in.SessionRevocationLifetime != nil {
		result.SessionRevocationLifetime = types.NumberValue(big.NewFloat(float64(*in.SessionRevocationLifetime)))
	} else {
		result.SessionRevocationLifetime = types.NumberNull()
	}
	if in.TrackAdapterSessionsForLogout != nil {
		result.TrackAdapterSessionsForLogout = types.BoolValue(*in.TrackAdapterSessionsForLogout)
	} else {
		result.TrackAdapterSessionsForLogout = types.BoolNull()
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
		result.Value = types.StringValue(*in.Value)
	}

	return &result
}

func flattenAuthenticationSource(in *pf.AuthenticationSource) *AuthenticationSourceData {
	result := AuthenticationSourceData{}
	if in.SourceRef != nil && in.SourceRef.Id != nil && *in.SourceRef.Id != "" {
		result.SourceRef = types.StringValue(*in.SourceRef.Id)
	} else {
		result.SourceRef = types.StringNull()
	}
	if in.Type != nil {
		result.Type = types.StringValue(*in.Type)
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
		result.BinaryEncoding = types.StringValue(*in.BinaryEncoding)
	} else {
		result.BinaryEncoding = types.StringNull()
	}

	return &result
}

func flattenCertView(in *pf.CertView) *CertViewData {
	result := CertViewData{}
	if in.CryptoProvider != nil {
		result.CryptoProvider = types.StringValue(*in.CryptoProvider)
	} else {
		result.CryptoProvider = types.StringNull()
	}
	if in.Expires != nil {
		result.Expires = types.StringValue(*in.Expires)
	} else {
		result.Expires = types.StringNull()
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.IssuerDN != nil {
		result.IssuerDN = types.StringValue(*in.IssuerDN)
	} else {
		result.IssuerDN = types.StringNull()
	}
	if in.KeyAlgorithm != nil {
		result.KeyAlgorithm = types.StringValue(*in.KeyAlgorithm)
	} else {
		result.KeyAlgorithm = types.StringNull()
	}
	if in.KeySize != nil {
		result.KeySize = types.NumberValue(big.NewFloat(float64(*in.KeySize)))
	} else {
		result.KeySize = types.NumberNull()
	}
	if in.SerialNumber != nil {
		result.SerialNumber = types.StringValue(*in.SerialNumber)
	} else {
		result.SerialNumber = types.StringNull()
	}
	if in.Sha1Fingerprint != nil {
		result.Sha1Fingerprint = types.StringValue(*in.Sha1Fingerprint)
	} else {
		result.Sha1Fingerprint = types.StringNull()
	}
	if in.Sha256Fingerprint != nil {
		result.Sha256Fingerprint = types.StringValue(*in.Sha256Fingerprint)
	} else {
		result.Sha256Fingerprint = types.StringNull()
	}
	if in.SignatureAlgorithm != nil {
		result.SignatureAlgorithm = types.StringValue(*in.SignatureAlgorithm)
	} else {
		result.SignatureAlgorithm = types.StringNull()
	}
	if in.Status != nil {
		result.Status = types.StringValue(*in.Status)
	} else {
		result.Status = types.StringNull()
	}
	if in.SubjectAlternativeNames != nil {
		result.SubjectAlternativeNames = flattenStringList(*in.SubjectAlternativeNames)
	} else {
		result.SubjectAlternativeNames = types.ListNull(types.StringType)
	}
	if in.SubjectDN != nil {
		result.SubjectDN = types.StringValue(*in.SubjectDN)
	} else {
		result.SubjectDN = types.StringNull()
	}
	if in.ValidFrom != nil {
		result.ValidFrom = types.StringValue(*in.ValidFrom)
	} else {
		result.ValidFrom = types.StringNull()
	}
	if in.Version != nil {
		result.Version = types.NumberValue(big.NewFloat(float64(*in.Version)))
	} else {
		result.Version = types.NumberNull()
	}

	return &result
}

func flattenClientAuth(in *pf.ClientAuth) *ClientAuthData {
	result := ClientAuthData{}
	if in.ClientCertIssuerDn != nil {
		result.ClientCertIssuerDn = types.StringValue(*in.ClientCertIssuerDn)
	} else {
		result.ClientCertIssuerDn = types.StringNull()
	}
	if in.ClientCertSubjectDn != nil {
		result.ClientCertSubjectDn = types.StringValue(*in.ClientCertSubjectDn)
	} else {
		result.ClientCertSubjectDn = types.StringNull()
	}
	if in.EncryptedSecret != nil {
		result.EncryptedSecret = types.StringValue(*in.EncryptedSecret)
	} else {
		result.EncryptedSecret = types.StringNull()
	}
	if in.EnforceReplayPrevention != nil {
		result.EnforceReplayPrevention = types.BoolValue(*in.EnforceReplayPrevention)
	} else {
		result.EnforceReplayPrevention = types.BoolNull()
	}
	if in.Secret != nil {
		result.Secret = types.StringValue(*in.Secret)
	} else {
		result.Secret = types.StringNull()
	}
	if in.TokenEndpointAuthSigningAlgorithm != nil {
		result.TokenEndpointAuthSigningAlgorithm = types.StringValue(*in.TokenEndpointAuthSigningAlgorithm)
	} else {
		result.TokenEndpointAuthSigningAlgorithm = types.StringNull()
	}
	if in.Type != nil {
		result.Type = types.StringValue(*in.Type)
	} else {
		result.Type = types.StringNull()
	}

	return &result
}

func flattenClientOIDCPolicy(in *pf.ClientOIDCPolicy) *ClientOIDCPolicyData {
	result := ClientOIDCPolicyData{}
	if in.GrantAccessSessionRevocationApi != nil {
		result.GrantAccessSessionRevocationApi = types.BoolValue(*in.GrantAccessSessionRevocationApi)
	} else {
		result.GrantAccessSessionRevocationApi = types.BoolNull()
	}
	if in.GrantAccessSessionSessionManagementApi != nil {
		result.GrantAccessSessionSessionManagementApi = types.BoolValue(*in.GrantAccessSessionSessionManagementApi)
	} else {
		result.GrantAccessSessionSessionManagementApi = types.BoolNull()
	}
	if in.IdTokenContentEncryptionAlgorithm != nil {
		result.IdTokenContentEncryptionAlgorithm = types.StringValue(*in.IdTokenContentEncryptionAlgorithm)
	} else {
		result.IdTokenContentEncryptionAlgorithm = types.StringNull()
	}
	if in.IdTokenEncryptionAlgorithm != nil {
		result.IdTokenEncryptionAlgorithm = types.StringValue(*in.IdTokenEncryptionAlgorithm)
	} else {
		result.IdTokenEncryptionAlgorithm = types.StringNull()
	}
	if in.IdTokenSigningAlgorithm != nil {
		result.IdTokenSigningAlgorithm = types.StringValue(*in.IdTokenSigningAlgorithm)
	} else {
		result.IdTokenSigningAlgorithm = types.StringNull()
	}
	if in.LogoutUris != nil {
		result.LogoutUris = flattenStringList(*in.LogoutUris)
	} else {
		result.LogoutUris = types.ListNull(types.StringType)
	}
	if in.PairwiseIdentifierUserType != nil {
		result.PairwiseIdentifierUserType = types.BoolValue(*in.PairwiseIdentifierUserType)
	} else {
		result.PairwiseIdentifierUserType = types.BoolNull()
	}
	if in.PingAccessLogoutCapable != nil {
		result.PingAccessLogoutCapable = types.BoolValue(*in.PingAccessLogoutCapable)
	} else {
		result.PingAccessLogoutCapable = types.BoolNull()
	}
	if in.PolicyGroup != nil && in.PolicyGroup.Id != nil && *in.PolicyGroup.Id != "" {
		result.PolicyGroup = types.StringValue(*in.PolicyGroup.Id)
	} else {
		result.PolicyGroup = types.StringNull()
	}
	if in.SectorIdentifierUri != nil {
		result.SectorIdentifierUri = types.StringValue(*in.SectorIdentifierUri)
	} else {
		result.SectorIdentifierUri = types.StringNull()
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
		result.AttributeName = types.StringValue(*in.AttributeName)
	}
	if in.Condition != nil {
		result.Condition = types.StringValue(*in.Condition)
	}
	if in.ErrorResult != nil {
		result.ErrorResult = types.StringValue(*in.ErrorResult)
	} else {
		result.ErrorResult = types.StringNull()
	}
	if in.Source != nil {
		result.Source = flattenSourceTypeIdKey(in.Source)
	}
	if in.Value != nil {
		result.Value = types.StringValue(*in.Value)
	}

	return &result
}

func flattenCustomAttributeSource(in *pf.CustomAttributeSource) *CustomAttributeSourceData {
	result := CustomAttributeSourceData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.StringValue(*in.DataStoreRef.Id)
	} else {
		result.DataStoreRef = types.StringNull()
	}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	} else {
		result.Description = types.StringNull()
	}
	if in.FilterFields != nil {
		result.FilterFields = flattenFieldEntrys(in.FilterFields)
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
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
		result.ErrorResult = types.StringValue(*in.ErrorResult)
	} else {
		result.ErrorResult = types.StringNull()
	}
	if in.Expression != nil {
		result.Expression = types.StringValue(*in.Expression)
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
		result.Name = types.StringValue(*in.Name)
	}
	if in.Value != nil {
		result.Value = types.StringValue(*in.Value)
	} else {
		result.Value = types.StringNull()
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
	} else {
		result.ColumnNames = types.ListNull(types.StringType)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.StringValue(*in.DataStoreRef.Id)
	} else {
		result.DataStoreRef = types.StringNull()
	}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	} else {
		result.Description = types.StringNull()
	}
	if in.Filter != nil {
		result.Filter = types.StringValue(*in.Filter)
	} else {
		result.Filter = types.StringNull()
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.Schema != nil {
		result.Schema = types.StringValue(*in.Schema)
	} else {
		result.Schema = types.StringNull()
	}
	if in.Table != nil {
		result.Table = types.StringValue(*in.Table)
	} else {
		result.Table = types.StringNull()
	}

	return &result
}

func flattenJwksSettings(in *pf.JwksSettings) *JwksSettingsData {
	result := JwksSettingsData{}
	if in.Jwks != nil {
		result.Jwks = types.StringValue(*in.Jwks)
	} else {
		result.Jwks = types.StringNull()
	}
	if in.JwksUrl != nil {
		result.JwksUrl = types.StringValue(*in.JwksUrl)
	} else {
		result.JwksUrl = types.StringNull()
	}

	return &result
}

func flattenLdapAttributeSource(in *pf.LdapAttributeSource) *LdapAttributeSourceData {
	result := LdapAttributeSourceData{}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.BaseDn != nil {
		result.BaseDn = types.StringValue(*in.BaseDn)
	} else {
		result.BaseDn = types.StringNull()
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = flattenMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
	}
	if in.DataStoreRef != nil && in.DataStoreRef.Id != nil && *in.DataStoreRef.Id != "" {
		result.DataStoreRef = types.StringValue(*in.DataStoreRef.Id)
	} else {
		result.DataStoreRef = types.StringNull()
	}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	} else {
		result.Description = types.StringNull()
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.MemberOfNestedGroup != nil {
		result.MemberOfNestedGroup = types.BoolValue(*in.MemberOfNestedGroup)
	} else {
		result.MemberOfNestedGroup = types.BoolNull()
	}
	if in.SearchAttributes != nil {
		result.SearchAttributes = flattenStringList(*in.SearchAttributes)
	} else {
		result.SearchAttributes = types.ListNull(types.StringType)
	}
	if in.SearchFilter != nil {
		result.SearchFilter = types.StringValue(*in.SearchFilter)
	} else {
		result.SearchFilter = types.StringNull()
	}
	if in.SearchScope != nil {
		result.SearchScope = types.StringValue(*in.SearchScope)
	} else {
		result.SearchScope = types.StringNull()
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
	} else {
		result.Values = types.ListNull(types.StringType)
	}

	return &result
}

func flattenPersistentGrantAttributes(in *[]*pf.PersistentGrantAttribute) *[]*PersistentGrantAttributeData {
	results := []*PersistentGrantAttributeData{}
	for _, data := range *in {
		results = append(results, flattenPersistentGrantAttribute(data))
	}
	return &results
}
func flattenPersistentGrantAttribute(in *pf.PersistentGrantAttribute) *PersistentGrantAttributeData {
	result := PersistentGrantAttributeData{}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}

	return &result
}

func flattenPersistentGrantContract(in *pf.PersistentGrantContract) *PersistentGrantContractData {
	result := PersistentGrantContractData{}
	if in.CoreAttributes != nil {
		result.CoreAttributes = flattenPersistentGrantAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		result.ExtendedAttributes = flattenPersistentGrantAttributes(in.ExtendedAttributes)
	}

	return &result
}

func flattenRedirectValidationLocalSettings(in *pf.RedirectValidationLocalSettings) *RedirectValidationLocalSettingsData {
	result := RedirectValidationLocalSettingsData{}
	if in.EnableInErrorResourceValidation != nil {
		result.EnableInErrorResourceValidation = types.BoolValue(*in.EnableInErrorResourceValidation)
	} else {
		result.EnableInErrorResourceValidation = types.BoolNull()
	}
	if in.EnableTargetResourceValidationForIdpDiscovery != nil {
		result.EnableTargetResourceValidationForIdpDiscovery = types.BoolValue(*in.EnableTargetResourceValidationForIdpDiscovery)
	} else {
		result.EnableTargetResourceValidationForIdpDiscovery = types.BoolNull()
	}
	if in.EnableTargetResourceValidationForSLO != nil {
		result.EnableTargetResourceValidationForSLO = types.BoolValue(*in.EnableTargetResourceValidationForSLO)
	} else {
		result.EnableTargetResourceValidationForSLO = types.BoolNull()
	}
	if in.EnableTargetResourceValidationForSSO != nil {
		result.EnableTargetResourceValidationForSSO = types.BoolValue(*in.EnableTargetResourceValidationForSSO)
	} else {
		result.EnableTargetResourceValidationForSSO = types.BoolNull()
	}
	if in.WhiteList != nil {
		result.WhiteList = flattenRedirectValidationSettingsWhitelistEntrys(in.WhiteList)
	}

	return &result
}

func flattenRedirectValidationPartnerSettings(in *pf.RedirectValidationPartnerSettings) *RedirectValidationPartnerSettingsData {
	result := RedirectValidationPartnerSettingsData{}
	if in.EnableWreplyValidationSLO != nil {
		result.EnableWreplyValidationSLO = types.BoolValue(*in.EnableWreplyValidationSLO)
	} else {
		result.EnableWreplyValidationSLO = types.BoolNull()
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
		result.AllowQueryAndFragment = types.BoolValue(*in.AllowQueryAndFragment)
	} else {
		result.AllowQueryAndFragment = types.BoolNull()
	}
	if in.IdpDiscovery != nil {
		result.IdpDiscovery = types.BoolValue(*in.IdpDiscovery)
	} else {
		result.IdpDiscovery = types.BoolNull()
	}
	if in.InErrorResource != nil {
		result.InErrorResource = types.BoolValue(*in.InErrorResource)
	} else {
		result.InErrorResource = types.BoolNull()
	}
	if in.RequireHttps != nil {
		result.RequireHttps = types.BoolValue(*in.RequireHttps)
	} else {
		result.RequireHttps = types.BoolNull()
	}
	if in.TargetResourceSLO != nil {
		result.TargetResourceSLO = types.BoolValue(*in.TargetResourceSLO)
	} else {
		result.TargetResourceSLO = types.BoolNull()
	}
	if in.TargetResourceSSO != nil {
		result.TargetResourceSSO = types.BoolValue(*in.TargetResourceSSO)
	} else {
		result.TargetResourceSSO = types.BoolNull()
	}
	if in.ValidDomain != nil {
		result.ValidDomain = types.StringValue(*in.ValidDomain)
	}
	if in.ValidPath != nil {
		result.ValidPath = types.StringValue(*in.ValidPath)
	} else {
		result.ValidPath = types.StringNull()
	}

	return &result
}

func flattenScopeEntrys(in *[]*pf.ScopeEntry) *[]*ScopeEntryData {
	results := []*ScopeEntryData{}
	for _, data := range *in {
		results = append(results, flattenScopeEntry(data))
	}
	return &results
}
func flattenScopeEntry(in *pf.ScopeEntry) *ScopeEntryData {
	result := ScopeEntryData{}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	}
	if in.Dynamic != nil {
		result.Dynamic = types.BoolValue(*in.Dynamic)
	} else {
		result.Dynamic = types.BoolNull()
	}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}

	return &result
}

func flattenScopeGroupEntrys(in *[]*pf.ScopeGroupEntry) *[]*ScopeGroupEntryData {
	results := []*ScopeGroupEntryData{}
	for _, data := range *in {
		results = append(results, flattenScopeGroupEntry(data))
	}
	return &results
}
func flattenScopeGroupEntry(in *pf.ScopeGroupEntry) *ScopeGroupEntryData {
	result := ScopeGroupEntryData{}
	if in.Description != nil {
		result.Description = types.StringValue(*in.Description)
	}
	if in.Name != nil {
		result.Name = types.StringValue(*in.Name)
	}
	if in.Scopes != nil {
		result.Scopes = flattenStringList(*in.Scopes)
	} else {
		result.Scopes = types.ListNull(types.StringType)
	}

	return &result
}

func flattenSourceTypeIdKey(in *pf.SourceTypeIdKey) *SourceTypeIdKeyData {
	result := SourceTypeIdKeyData{}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
	}
	if in.Type != nil {
		result.Type = types.StringValue(*in.Type)
	}

	return &result
}

func flattenX509File(in *pf.X509File) *X509FileData {
	result := X509FileData{}
	if in.CryptoProvider != nil {
		result.CryptoProvider = types.StringValue(*in.CryptoProvider)
	} else {
		result.CryptoProvider = types.StringNull()
	}
	if in.FileData != nil {
		result.FileData = types.StringValue(*in.FileData)
	}
	if in.Id != nil {
		result.Id = types.StringValue(*in.Id)
	} else {
		result.Id = types.StringNull()
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

func flattenStringList(in []*string) types.List {
	var values []attr.Value
	for _, s := range in {
		values = append(values, types.StringValue(*s))
	}
	return types.ListValueMust(types.StringType, values)
}

func flattenStringSet(in []*string) types.Set {
	var values []attr.Value
	for _, s := range in {
		values = append(values, types.StringValue(*s))
	}
	return types.SetValueMust(types.StringType, values)
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
