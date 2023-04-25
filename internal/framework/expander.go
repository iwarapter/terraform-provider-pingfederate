package framework

import (
	"context"

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
	if !in.AuthenticationPolicyContractRef.IsUnknown() && !in.AuthenticationPolicyContractRef.IsNull() {
		result.AuthenticationPolicyContractRef = &pf.ResourceLink{Id: String(in.AuthenticationPolicyContractRef.ValueString())}
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if in.IssuanceCriteria != nil {
		result.IssuanceCriteria = expandIssuanceCriteria(*in.IssuanceCriteria)
	}

	return &result
}

func expandApplicationSessionPolicy(in ApplicationSessionPolicyData) *pf.ApplicationSessionPolicy {
	var result pf.ApplicationSessionPolicy
	if !in.IdleTimeoutMins.IsUnknown() && !in.IdleTimeoutMins.IsNull() {
		i64, _ := in.IdleTimeoutMins.ValueBigFloat().Int64()
		result.IdleTimeoutMins = Int(int(i64))
	}
	if !in.MaxTimeoutMins.IsUnknown() && !in.MaxTimeoutMins.IsNull() {
		i64, _ := in.MaxTimeoutMins.ValueBigFloat().Int64()
		result.MaxTimeoutMins = Int(int(i64))
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
	if !in.ExtendedAttributes.IsUnknown() && !in.ExtendedAttributes.IsNull() {
		results := []*string{}
		diags := in.ExtendedAttributes.ElementsAs(context.Background(), &results, false)
		diags.HasError()

		attrs := []*pf.AuthenticationPolicyContractAttribute{}
		for _, data := range results {
			attrs = append(attrs, &pf.AuthenticationPolicyContractAttribute{Name: data})
		}
		result.ExtendedAttributes = &attrs
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.Name.IsUnknown() && !in.Name.IsNull() {
		result.Name = String(in.Name.ValueString())
	}

	return &result
}

func expandAuthenticationSessionPolicy(in AuthenticationSessionPolicyData) *pf.AuthenticationSessionPolicy {
	var result pf.AuthenticationSessionPolicy
	if in.AuthenticationSource != nil {
		result.AuthenticationSource = expandAuthenticationSource(*in.AuthenticationSource)
	}
	if !in.AuthnContextSensitive.IsUnknown() && !in.AuthnContextSensitive.IsNull() {
		result.AuthnContextSensitive = Bool(in.AuthnContextSensitive.ValueBool())
	}
	if !in.EnableSessions.IsUnknown() && !in.EnableSessions.IsNull() {
		result.EnableSessions = Bool(in.EnableSessions.ValueBool())
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.IdleTimeoutMins.IsUnknown() && !in.IdleTimeoutMins.IsNull() {
		i64, _ := in.IdleTimeoutMins.ValueBigFloat().Int64()
		result.IdleTimeoutMins = Int(int(i64))
	}
	if !in.MaxTimeoutMins.IsUnknown() && !in.MaxTimeoutMins.IsNull() {
		i64, _ := in.MaxTimeoutMins.ValueBigFloat().Int64()
		result.MaxTimeoutMins = Int(int(i64))
	}
	if !in.Persistent.IsUnknown() && !in.Persistent.IsNull() {
		result.Persistent = Bool(in.Persistent.ValueBool())
	}
	if !in.TimeoutDisplayUnit.IsUnknown() && !in.TimeoutDisplayUnit.IsNull() {
		result.TimeoutDisplayUnit = String(in.TimeoutDisplayUnit.ValueString())
	}

	return &result
}

func expandClient(in ClientData) *pf.Client {
	var result pf.Client
	if !in.AllowAuthenticationApiInit.IsUnknown() && !in.AllowAuthenticationApiInit.IsNull() {
		result.AllowAuthenticationApiInit = Bool(in.AllowAuthenticationApiInit.ValueBool())
	}
	if !in.BypassActivationCodeConfirmationOverride.IsUnknown() && !in.BypassActivationCodeConfirmationOverride.IsNull() {
		result.BypassActivationCodeConfirmationOverride = Bool(in.BypassActivationCodeConfirmationOverride.ValueBool())
	}
	if !in.BypassApprovalPage.IsUnknown() && !in.BypassApprovalPage.IsNull() {
		result.BypassApprovalPage = Bool(in.BypassApprovalPage.ValueBool())
	}
	if !in.CibaDeliveryMode.IsUnknown() && !in.CibaDeliveryMode.IsNull() {
		result.CibaDeliveryMode = String(in.CibaDeliveryMode.ValueString())
	}
	if !in.CibaNotificationEndpoint.IsUnknown() && !in.CibaNotificationEndpoint.IsNull() {
		result.CibaNotificationEndpoint = String(in.CibaNotificationEndpoint.ValueString())
	}
	if !in.CibaPollingInterval.IsUnknown() && !in.CibaPollingInterval.IsNull() {
		i64, _ := in.CibaPollingInterval.ValueBigFloat().Int64()
		result.CibaPollingInterval = Int(int(i64))
	}
	if !in.CibaRequestObjectSigningAlgorithm.IsUnknown() && !in.CibaRequestObjectSigningAlgorithm.IsNull() {
		result.CibaRequestObjectSigningAlgorithm = String(in.CibaRequestObjectSigningAlgorithm.ValueString())
	}
	if !in.CibaRequireSignedRequests.IsUnknown() && !in.CibaRequireSignedRequests.IsNull() {
		result.CibaRequireSignedRequests = Bool(in.CibaRequireSignedRequests.ValueBool())
	}
	if !in.CibaUserCodeSupported.IsUnknown() && !in.CibaUserCodeSupported.IsNull() {
		result.CibaUserCodeSupported = Bool(in.CibaUserCodeSupported.ValueBool())
	}
	if in.ClientAuth != nil {
		result.ClientAuth = expandClientAuth(*in.ClientAuth)
	}
	if !in.ClientId.IsUnknown() && !in.ClientId.IsNull() {
		result.ClientId = String(in.ClientId.ValueString())
	}
	if !in.ClientSecretChangedTime.IsUnknown() && !in.ClientSecretChangedTime.IsNull() {
		result.ClientSecretChangedTime = String(in.ClientSecretChangedTime.ValueString())
	}
	if !in.ClientSecretRetentionPeriod.IsUnknown() && !in.ClientSecretRetentionPeriod.IsNull() {
		i64, _ := in.ClientSecretRetentionPeriod.ValueBigFloat().Int64()
		result.ClientSecretRetentionPeriod = Int(int(i64))
	}
	if !in.ClientSecretRetentionPeriodType.IsUnknown() && !in.ClientSecretRetentionPeriodType.IsNull() {
		result.ClientSecretRetentionPeriodType = String(in.ClientSecretRetentionPeriodType.ValueString())
	}
	if !in.DefaultAccessTokenManagerRef.IsUnknown() && !in.DefaultAccessTokenManagerRef.IsNull() {
		result.DefaultAccessTokenManagerRef = &pf.ResourceLink{Id: String(in.DefaultAccessTokenManagerRef.ValueString())}
	}
	if !in.Description.IsUnknown() && !in.Description.IsNull() {
		result.Description = String(in.Description.ValueString())
	}
	if !in.DeviceFlowSettingType.IsUnknown() && !in.DeviceFlowSettingType.IsNull() {
		result.DeviceFlowSettingType = String(in.DeviceFlowSettingType.ValueString())
	}
	if !in.DevicePollingIntervalOverride.IsUnknown() && !in.DevicePollingIntervalOverride.IsNull() {
		i64, _ := in.DevicePollingIntervalOverride.ValueBigFloat().Int64()
		result.DevicePollingIntervalOverride = Int(int(i64))
	}
	if !in.Enabled.IsUnknown() && !in.Enabled.IsNull() {
		result.Enabled = Bool(in.Enabled.ValueBool())
	}
	if !in.ExclusiveScopes.IsUnknown() && !in.ExclusiveScopes.IsNull() {
		result.ExclusiveScopes = expandStringList(in.ExclusiveScopes)
	}
	if in.ExtendedParameters != nil {
		result.ExtendedParameters = expandMapParameterValuess(in.ExtendedParameters)
	}
	if !in.GrantTypes.IsUnknown() && !in.GrantTypes.IsNull() {
		result.GrantTypes = expandStringList(in.GrantTypes)
	}
	if in.JwksSettings != nil {
		result.JwksSettings = expandJwksSettings(*in.JwksSettings)
	}
	if !in.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm.IsUnknown() && !in.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm.IsNull() {
		result.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm = String(in.JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm.ValueString())
	}
	if !in.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm.IsUnknown() && !in.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm.IsNull() {
		result.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm = String(in.JwtSecuredAuthorizationResponseModeEncryptionAlgorithm.ValueString())
	}
	if !in.JwtSecuredAuthorizationResponseModeSigningAlgorithm.IsUnknown() && !in.JwtSecuredAuthorizationResponseModeSigningAlgorithm.IsNull() {
		result.JwtSecuredAuthorizationResponseModeSigningAlgorithm = String(in.JwtSecuredAuthorizationResponseModeSigningAlgorithm.ValueString())
	}
	if !in.LogoUrl.IsUnknown() && !in.LogoUrl.IsNull() {
		result.LogoUrl = String(in.LogoUrl.ValueString())
	}
	if !in.Name.IsUnknown() && !in.Name.IsNull() {
		result.Name = String(in.Name.ValueString())
	}
	if in.OidcPolicy != nil {
		result.OidcPolicy = expandClientOIDCPolicy(*in.OidcPolicy)
	}
	if !in.PendingAuthorizationTimeoutOverride.IsUnknown() && !in.PendingAuthorizationTimeoutOverride.IsNull() {
		i64, _ := in.PendingAuthorizationTimeoutOverride.ValueBigFloat().Int64()
		result.PendingAuthorizationTimeoutOverride = Int(int(i64))
	}
	if !in.PersistentGrantExpirationTime.IsUnknown() && !in.PersistentGrantExpirationTime.IsNull() {
		i64, _ := in.PersistentGrantExpirationTime.ValueBigFloat().Int64()
		result.PersistentGrantExpirationTime = Int(int(i64))
	}
	if !in.PersistentGrantExpirationTimeUnit.IsUnknown() && !in.PersistentGrantExpirationTimeUnit.IsNull() {
		result.PersistentGrantExpirationTimeUnit = String(in.PersistentGrantExpirationTimeUnit.ValueString())
	}
	if !in.PersistentGrantExpirationType.IsUnknown() && !in.PersistentGrantExpirationType.IsNull() {
		result.PersistentGrantExpirationType = String(in.PersistentGrantExpirationType.ValueString())
	}
	if !in.PersistentGrantIdleTimeout.IsUnknown() && !in.PersistentGrantIdleTimeout.IsNull() {
		i64, _ := in.PersistentGrantIdleTimeout.ValueBigFloat().Int64()
		result.PersistentGrantIdleTimeout = Int(int(i64))
	}
	if !in.PersistentGrantIdleTimeoutTimeUnit.IsUnknown() && !in.PersistentGrantIdleTimeoutTimeUnit.IsNull() {
		result.PersistentGrantIdleTimeoutTimeUnit = String(in.PersistentGrantIdleTimeoutTimeUnit.ValueString())
	}
	if !in.PersistentGrantIdleTimeoutType.IsUnknown() && !in.PersistentGrantIdleTimeoutType.IsNull() {
		result.PersistentGrantIdleTimeoutType = String(in.PersistentGrantIdleTimeoutType.ValueString())
	}
	if !in.PersistentGrantReuseGrantTypes.IsUnknown() && !in.PersistentGrantReuseGrantTypes.IsNull() {
		result.PersistentGrantReuseGrantTypes = expandStringList(in.PersistentGrantReuseGrantTypes)
	}
	if !in.PersistentGrantReuseType.IsUnknown() && !in.PersistentGrantReuseType.IsNull() {
		result.PersistentGrantReuseType = String(in.PersistentGrantReuseType.ValueString())
	}
	if !in.RedirectUris.IsUnknown() && !in.RedirectUris.IsNull() {
		result.RedirectUris = expandStringSet(in.RedirectUris)
	}
	if !in.RefreshRolling.IsUnknown() && !in.RefreshRolling.IsNull() {
		result.RefreshRolling = String(in.RefreshRolling.ValueString())
	}
	if !in.RefreshTokenRollingGracePeriod.IsUnknown() && !in.RefreshTokenRollingGracePeriod.IsNull() {
		i64, _ := in.RefreshTokenRollingGracePeriod.ValueBigFloat().Int64()
		result.RefreshTokenRollingGracePeriod = Int(int(i64))
	}
	if !in.RefreshTokenRollingGracePeriodType.IsUnknown() && !in.RefreshTokenRollingGracePeriodType.IsNull() {
		result.RefreshTokenRollingGracePeriodType = String(in.RefreshTokenRollingGracePeriodType.ValueString())
	}
	if !in.RefreshTokenRollingInterval.IsUnknown() && !in.RefreshTokenRollingInterval.IsNull() {
		i64, _ := in.RefreshTokenRollingInterval.ValueBigFloat().Int64()
		result.RefreshTokenRollingInterval = Int(int(i64))
	}
	if !in.RefreshTokenRollingIntervalType.IsUnknown() && !in.RefreshTokenRollingIntervalType.IsNull() {
		result.RefreshTokenRollingIntervalType = String(in.RefreshTokenRollingIntervalType.ValueString())
	}
	if !in.RequestObjectSigningAlgorithm.IsUnknown() && !in.RequestObjectSigningAlgorithm.IsNull() {
		result.RequestObjectSigningAlgorithm = String(in.RequestObjectSigningAlgorithm.ValueString())
	}
	if !in.RequestPolicyRef.IsUnknown() && !in.RequestPolicyRef.IsNull() {
		result.RequestPolicyRef = &pf.ResourceLink{Id: String(in.RequestPolicyRef.ValueString())}
	}
	if !in.RequireJwtSecuredAuthorizationResponseMode.IsUnknown() && !in.RequireJwtSecuredAuthorizationResponseMode.IsNull() {
		result.RequireJwtSecuredAuthorizationResponseMode = Bool(in.RequireJwtSecuredAuthorizationResponseMode.ValueBool())
	}
	if !in.RequireProofKeyForCodeExchange.IsUnknown() && !in.RequireProofKeyForCodeExchange.IsNull() {
		result.RequireProofKeyForCodeExchange = Bool(in.RequireProofKeyForCodeExchange.ValueBool())
	}
	if !in.RequirePushedAuthorizationRequests.IsUnknown() && !in.RequirePushedAuthorizationRequests.IsNull() {
		result.RequirePushedAuthorizationRequests = Bool(in.RequirePushedAuthorizationRequests.ValueBool())
	}
	if !in.RequireSignedRequests.IsUnknown() && !in.RequireSignedRequests.IsNull() {
		result.RequireSignedRequests = Bool(in.RequireSignedRequests.ValueBool())
	}
	if !in.RestrictScopes.IsUnknown() && !in.RestrictScopes.IsNull() {
		result.RestrictScopes = Bool(in.RestrictScopes.ValueBool())
	}
	if !in.RestrictToDefaultAccessTokenManager.IsUnknown() && !in.RestrictToDefaultAccessTokenManager.IsNull() {
		result.RestrictToDefaultAccessTokenManager = Bool(in.RestrictToDefaultAccessTokenManager.ValueBool())
	}
	if !in.RestrictedResponseTypes.IsUnknown() && !in.RestrictedResponseTypes.IsNull() {
		result.RestrictedResponseTypes = expandStringList(in.RestrictedResponseTypes)
	}
	if !in.RestrictedScopes.IsUnknown() && !in.RestrictedScopes.IsNull() {
		result.RestrictedScopes = expandStringSet(in.RestrictedScopes)
	}
	if !in.TokenExchangeProcessorPolicyRef.IsUnknown() && !in.TokenExchangeProcessorPolicyRef.IsNull() {
		result.TokenExchangeProcessorPolicyRef = &pf.ResourceLink{Id: String(in.TokenExchangeProcessorPolicyRef.ValueString())}
	}
	if !in.TokenIntrospectionContentEncryptionAlgorithm.IsUnknown() && !in.TokenIntrospectionContentEncryptionAlgorithm.IsNull() {
		result.TokenIntrospectionContentEncryptionAlgorithm = String(in.TokenIntrospectionContentEncryptionAlgorithm.ValueString())
	}
	if !in.TokenIntrospectionEncryptionAlgorithm.IsUnknown() && !in.TokenIntrospectionEncryptionAlgorithm.IsNull() {
		result.TokenIntrospectionEncryptionAlgorithm = String(in.TokenIntrospectionEncryptionAlgorithm.ValueString())
	}
	if !in.TokenIntrospectionSigningAlgorithm.IsUnknown() && !in.TokenIntrospectionSigningAlgorithm.IsNull() {
		result.TokenIntrospectionSigningAlgorithm = String(in.TokenIntrospectionSigningAlgorithm.ValueString())
	}
	if !in.UserAuthorizationUrlOverride.IsUnknown() && !in.UserAuthorizationUrlOverride.IsNull() {
		result.UserAuthorizationUrlOverride = String(in.UserAuthorizationUrlOverride.ValueString())
	}
	if !in.ValidateUsingAllEligibleAtms.IsUnknown() && !in.ValidateUsingAllEligibleAtms.IsNull() {
		result.ValidateUsingAllEligibleAtms = Bool(in.ValidateUsingAllEligibleAtms.ValueBool())
	}

	return &result
}

func expandGlobalAuthenticationSessionPolicy(in GlobalAuthenticationSessionPolicyData) *pf.GlobalAuthenticationSessionPolicy {
	var result pf.GlobalAuthenticationSessionPolicy
	if !in.EnableSessions.IsUnknown() && !in.EnableSessions.IsNull() {
		result.EnableSessions = Bool(in.EnableSessions.ValueBool())
	}
	if !in.HashUniqueUserKeyAttribute.IsUnknown() && !in.HashUniqueUserKeyAttribute.IsNull() {
		result.HashUniqueUserKeyAttribute = Bool(in.HashUniqueUserKeyAttribute.ValueBool())
	}
	if !in.IdleTimeoutDisplayUnit.IsUnknown() && !in.IdleTimeoutDisplayUnit.IsNull() {
		result.IdleTimeoutDisplayUnit = String(in.IdleTimeoutDisplayUnit.ValueString())
	}
	if !in.IdleTimeoutMins.IsUnknown() && !in.IdleTimeoutMins.IsNull() {
		i64, _ := in.IdleTimeoutMins.ValueBigFloat().Int64()
		result.IdleTimeoutMins = Int(int(i64))
	}
	if !in.MaxTimeoutDisplayUnit.IsUnknown() && !in.MaxTimeoutDisplayUnit.IsNull() {
		result.MaxTimeoutDisplayUnit = String(in.MaxTimeoutDisplayUnit.ValueString())
	}
	if !in.MaxTimeoutMins.IsUnknown() && !in.MaxTimeoutMins.IsNull() {
		i64, _ := in.MaxTimeoutMins.ValueBigFloat().Int64()
		result.MaxTimeoutMins = Int(int(i64))
	}
	if !in.PersistentSessions.IsUnknown() && !in.PersistentSessions.IsNull() {
		result.PersistentSessions = Bool(in.PersistentSessions.ValueBool())
	}

	return &result
}

func expandMetadataUrl(in MetadataUrlData) *pf.MetadataUrl {
	var result pf.MetadataUrl
	if in.CertView != nil {
		result.CertView = expandCertView(*in.CertView)
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.Name.IsUnknown() && !in.Name.IsNull() {
		result.Name = String(in.Name.ValueString())
	}
	if !in.Url.IsUnknown() && !in.Url.IsNull() {
		result.Url = String(in.Url.ValueString())
	}
	if !in.ValidateSignature.IsUnknown() && !in.ValidateSignature.IsNull() {
		result.ValidateSignature = Bool(in.ValidateSignature.ValueBool())
	}
	if in.X509File != nil {
		result.X509File = expandX509File(*in.X509File)
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

func expandSessionSettings(in SessionSettingsData) *pf.SessionSettings {
	var result pf.SessionSettings
	if !in.RevokeUserSessionOnLogout.IsUnknown() && !in.RevokeUserSessionOnLogout.IsNull() {
		result.RevokeUserSessionOnLogout = Bool(in.RevokeUserSessionOnLogout.ValueBool())
	}
	if !in.SessionRevocationLifetime.IsUnknown() && !in.SessionRevocationLifetime.IsNull() {
		i64, _ := in.SessionRevocationLifetime.ValueBigFloat().Int64()
		result.SessionRevocationLifetime = Int(int(i64))
	}
	if !in.TrackAdapterSessionsForLogout.IsUnknown() && !in.TrackAdapterSessionsForLogout.IsNull() {
		result.TrackAdapterSessionsForLogout = Bool(in.TrackAdapterSessionsForLogout.ValueBool())
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
	if !in.Value.IsUnknown() && !in.Value.IsNull() {
		result.Value = String(in.Value.ValueString())
	}

	return &result
}

func expandAuthenticationSource(in AuthenticationSourceData) *pf.AuthenticationSource {
	var result pf.AuthenticationSource
	if !in.SourceRef.IsUnknown() && !in.SourceRef.IsNull() {
		result.SourceRef = &pf.ResourceLink{Id: String(in.SourceRef.ValueString())}
	}
	if !in.Type.IsUnknown() && !in.Type.IsNull() {
		result.Type = String(in.Type.ValueString())
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
	if !in.BinaryEncoding.IsUnknown() && !in.BinaryEncoding.IsNull() {
		result.BinaryEncoding = String(in.BinaryEncoding.ValueString())
	}

	return &result
}

func expandCertView(in CertViewData) *pf.CertView {
	var result pf.CertView
	if !in.CryptoProvider.IsUnknown() && !in.CryptoProvider.IsNull() {
		result.CryptoProvider = String(in.CryptoProvider.ValueString())
	}
	if !in.Expires.IsUnknown() && !in.Expires.IsNull() {
		result.Expires = String(in.Expires.ValueString())
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.IssuerDN.IsUnknown() && !in.IssuerDN.IsNull() {
		result.IssuerDN = String(in.IssuerDN.ValueString())
	}
	if !in.KeyAlgorithm.IsUnknown() && !in.KeyAlgorithm.IsNull() {
		result.KeyAlgorithm = String(in.KeyAlgorithm.ValueString())
	}
	if !in.KeySize.IsUnknown() && !in.KeySize.IsNull() {
		i64, _ := in.KeySize.ValueBigFloat().Int64()
		result.KeySize = Int(int(i64))
	}
	if !in.SerialNumber.IsUnknown() && !in.SerialNumber.IsNull() {
		result.SerialNumber = String(in.SerialNumber.ValueString())
	}
	if !in.Sha1Fingerprint.IsUnknown() && !in.Sha1Fingerprint.IsNull() {
		result.Sha1Fingerprint = String(in.Sha1Fingerprint.ValueString())
	}
	if !in.Sha256Fingerprint.IsUnknown() && !in.Sha256Fingerprint.IsNull() {
		result.Sha256Fingerprint = String(in.Sha256Fingerprint.ValueString())
	}
	if !in.SignatureAlgorithm.IsUnknown() && !in.SignatureAlgorithm.IsNull() {
		result.SignatureAlgorithm = String(in.SignatureAlgorithm.ValueString())
	}
	if !in.Status.IsUnknown() && !in.Status.IsNull() {
		result.Status = String(in.Status.ValueString())
	}
	if !in.SubjectAlternativeNames.IsUnknown() && !in.SubjectAlternativeNames.IsNull() {
		result.SubjectAlternativeNames = expandStringList(in.SubjectAlternativeNames)
	}
	if !in.SubjectDN.IsUnknown() && !in.SubjectDN.IsNull() {
		result.SubjectDN = String(in.SubjectDN.ValueString())
	}
	if !in.ValidFrom.IsUnknown() && !in.ValidFrom.IsNull() {
		result.ValidFrom = String(in.ValidFrom.ValueString())
	}
	if !in.Version.IsUnknown() && !in.Version.IsNull() {
		i64, _ := in.Version.ValueBigFloat().Int64()
		result.Version = Int(int(i64))
	}

	return &result
}

func expandClientAuth(in ClientAuthData) *pf.ClientAuth {
	var result pf.ClientAuth
	if !in.ClientCertIssuerDn.IsUnknown() && !in.ClientCertIssuerDn.IsNull() {
		result.ClientCertIssuerDn = String(in.ClientCertIssuerDn.ValueString())
	}
	if !in.ClientCertSubjectDn.IsUnknown() && !in.ClientCertSubjectDn.IsNull() {
		result.ClientCertSubjectDn = String(in.ClientCertSubjectDn.ValueString())
	}
	if !in.EncryptedSecret.IsUnknown() && !in.EncryptedSecret.IsNull() {
		result.EncryptedSecret = String(in.EncryptedSecret.ValueString())
	}
	if !in.EnforceReplayPrevention.IsUnknown() && !in.EnforceReplayPrevention.IsNull() {
		result.EnforceReplayPrevention = Bool(in.EnforceReplayPrevention.ValueBool())
	}
	if !in.Secret.IsUnknown() && !in.Secret.IsNull() {
		result.Secret = String(in.Secret.ValueString())
	}
	if !in.TokenEndpointAuthSigningAlgorithm.IsUnknown() && !in.TokenEndpointAuthSigningAlgorithm.IsNull() {
		result.TokenEndpointAuthSigningAlgorithm = String(in.TokenEndpointAuthSigningAlgorithm.ValueString())
	}
	if !in.Type.IsUnknown() && !in.Type.IsNull() {
		result.Type = String(in.Type.ValueString())
	}

	return &result
}

func expandClientOIDCPolicy(in ClientOIDCPolicyData) *pf.ClientOIDCPolicy {
	var result pf.ClientOIDCPolicy
	if !in.GrantAccessSessionRevocationApi.IsUnknown() && !in.GrantAccessSessionRevocationApi.IsNull() {
		result.GrantAccessSessionRevocationApi = Bool(in.GrantAccessSessionRevocationApi.ValueBool())
	}
	if !in.GrantAccessSessionSessionManagementApi.IsUnknown() && !in.GrantAccessSessionSessionManagementApi.IsNull() {
		result.GrantAccessSessionSessionManagementApi = Bool(in.GrantAccessSessionSessionManagementApi.ValueBool())
	}
	if !in.IdTokenContentEncryptionAlgorithm.IsUnknown() && !in.IdTokenContentEncryptionAlgorithm.IsNull() {
		result.IdTokenContentEncryptionAlgorithm = String(in.IdTokenContentEncryptionAlgorithm.ValueString())
	}
	if !in.IdTokenEncryptionAlgorithm.IsUnknown() && !in.IdTokenEncryptionAlgorithm.IsNull() {
		result.IdTokenEncryptionAlgorithm = String(in.IdTokenEncryptionAlgorithm.ValueString())
	}
	if !in.IdTokenSigningAlgorithm.IsUnknown() && !in.IdTokenSigningAlgorithm.IsNull() {
		result.IdTokenSigningAlgorithm = String(in.IdTokenSigningAlgorithm.ValueString())
	}
	if !in.LogoutUris.IsUnknown() && !in.LogoutUris.IsNull() {
		result.LogoutUris = expandStringList(in.LogoutUris)
	}
	if !in.PairwiseIdentifierUserType.IsUnknown() && !in.PairwiseIdentifierUserType.IsNull() {
		result.PairwiseIdentifierUserType = Bool(in.PairwiseIdentifierUserType.ValueBool())
	}
	if !in.PingAccessLogoutCapable.IsUnknown() && !in.PingAccessLogoutCapable.IsNull() {
		result.PingAccessLogoutCapable = Bool(in.PingAccessLogoutCapable.ValueBool())
	}
	if !in.PolicyGroup.IsUnknown() && !in.PolicyGroup.IsNull() {
		result.PolicyGroup = &pf.ResourceLink{Id: String(in.PolicyGroup.ValueString())}
	}
	if !in.SectorIdentifierUri.IsUnknown() && !in.SectorIdentifierUri.IsNull() {
		result.SectorIdentifierUri = String(in.SectorIdentifierUri.ValueString())
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
	if !in.AttributeName.IsUnknown() && !in.AttributeName.IsNull() {
		result.AttributeName = String(in.AttributeName.ValueString())
	}
	if !in.Condition.IsUnknown() && !in.Condition.IsNull() {
		result.Condition = String(in.Condition.ValueString())
	}
	if !in.ErrorResult.IsUnknown() && !in.ErrorResult.IsNull() {
		result.ErrorResult = String(in.ErrorResult.ValueString())
	}
	if in.Source != nil {
		result.Source = expandSourceTypeIdKey(*in.Source)
	}
	if !in.Value.IsUnknown() && !in.Value.IsNull() {
		result.Value = String(in.Value.ValueString())
	}

	return &result
}

func expandCustomAttributeSource(in CustomAttributeSourceData) *pf.CustomAttributeSource {
	var result pf.CustomAttributeSource
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.DataStoreRef.IsUnknown() && !in.DataStoreRef.IsNull() {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.ValueString())}
	}
	if !in.Description.IsUnknown() && !in.Description.IsNull() {
		result.Description = String(in.Description.ValueString())
	}
	if in.FilterFields != nil {
		result.FilterFields = expandFieldEntrys(in.FilterFields)
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
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
	if !in.ErrorResult.IsUnknown() && !in.ErrorResult.IsNull() {
		result.ErrorResult = String(in.ErrorResult.ValueString())
	}
	if !in.Expression.IsUnknown() && !in.Expression.IsNull() {
		result.Expression = String(in.Expression.ValueString())
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
	if !in.Name.IsUnknown() && !in.Name.IsNull() {
		result.Name = String(in.Name.ValueString())
	}
	if !in.Value.IsUnknown() && !in.Value.IsNull() {
		result.Value = String(in.Value.ValueString())
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
	if !in.ColumnNames.IsUnknown() && !in.ColumnNames.IsNull() {
		result.ColumnNames = expandStringList(in.ColumnNames)
	}
	if !in.DataStoreRef.IsUnknown() && !in.DataStoreRef.IsNull() {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.ValueString())}
	}
	if !in.Description.IsUnknown() && !in.Description.IsNull() {
		result.Description = String(in.Description.ValueString())
	}
	if !in.Filter.IsUnknown() && !in.Filter.IsNull() {
		result.Filter = String(in.Filter.ValueString())
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.Schema.IsUnknown() && !in.Schema.IsNull() {
		result.Schema = String(in.Schema.ValueString())
	}
	if !in.Table.IsUnknown() && !in.Table.IsNull() {
		result.Table = String(in.Table.ValueString())
	}

	return &result
}

func expandJwksSettings(in JwksSettingsData) *pf.JwksSettings {
	var result pf.JwksSettings
	if !in.Jwks.IsUnknown() && !in.Jwks.IsNull() {
		result.Jwks = String(in.Jwks.ValueString())
	}
	if !in.JwksUrl.IsUnknown() && !in.JwksUrl.IsNull() {
		result.JwksUrl = String(in.JwksUrl.ValueString())
	}

	return &result
}

func expandLdapAttributeSource(in LdapAttributeSourceData) *pf.LdapAttributeSource {
	var result pf.LdapAttributeSource
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.BaseDn.IsUnknown() && !in.BaseDn.IsNull() {
		result.BaseDn = String(in.BaseDn.ValueString())
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = expandMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
	}
	if !in.DataStoreRef.IsUnknown() && !in.DataStoreRef.IsNull() {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.ValueString())}
	}
	if !in.Description.IsUnknown() && !in.Description.IsNull() {
		result.Description = String(in.Description.ValueString())
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.MemberOfNestedGroup.IsUnknown() && !in.MemberOfNestedGroup.IsNull() {
		result.MemberOfNestedGroup = Bool(in.MemberOfNestedGroup.ValueBool())
	}
	if !in.SearchAttributes.IsUnknown() && !in.SearchAttributes.IsNull() {
		result.SearchAttributes = expandStringList(in.SearchAttributes)
	}
	if !in.SearchFilter.IsUnknown() && !in.SearchFilter.IsNull() {
		result.SearchFilter = String(in.SearchFilter.ValueString())
	}
	if !in.SearchScope.IsUnknown() && !in.SearchScope.IsNull() {
		result.SearchScope = String(in.SearchScope.ValueString())
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
	if !in.Values.IsUnknown() && !in.Values.IsNull() {
		result.Values = expandStringList(in.Values)
	}

	return &result
}

func expandRedirectValidationLocalSettings(in RedirectValidationLocalSettingsData) *pf.RedirectValidationLocalSettings {
	var result pf.RedirectValidationLocalSettings
	if !in.EnableInErrorResourceValidation.IsUnknown() && !in.EnableInErrorResourceValidation.IsNull() {
		result.EnableInErrorResourceValidation = Bool(in.EnableInErrorResourceValidation.ValueBool())
	}
	if !in.EnableTargetResourceValidationForIdpDiscovery.IsUnknown() && !in.EnableTargetResourceValidationForIdpDiscovery.IsNull() {
		result.EnableTargetResourceValidationForIdpDiscovery = Bool(in.EnableTargetResourceValidationForIdpDiscovery.ValueBool())
	}
	if !in.EnableTargetResourceValidationForSLO.IsUnknown() && !in.EnableTargetResourceValidationForSLO.IsNull() {
		result.EnableTargetResourceValidationForSLO = Bool(in.EnableTargetResourceValidationForSLO.ValueBool())
	}
	if !in.EnableTargetResourceValidationForSSO.IsUnknown() && !in.EnableTargetResourceValidationForSSO.IsNull() {
		result.EnableTargetResourceValidationForSSO = Bool(in.EnableTargetResourceValidationForSSO.ValueBool())
	}
	if in.WhiteList != nil {
		result.WhiteList = expandRedirectValidationSettingsWhitelistEntrys(in.WhiteList)
	}

	return &result
}

func expandRedirectValidationPartnerSettings(in RedirectValidationPartnerSettingsData) *pf.RedirectValidationPartnerSettings {
	var result pf.RedirectValidationPartnerSettings
	if !in.EnableWreplyValidationSLO.IsUnknown() && !in.EnableWreplyValidationSLO.IsNull() {
		result.EnableWreplyValidationSLO = Bool(in.EnableWreplyValidationSLO.ValueBool())
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
	if !in.AllowQueryAndFragment.IsUnknown() && !in.AllowQueryAndFragment.IsNull() {
		result.AllowQueryAndFragment = Bool(in.AllowQueryAndFragment.ValueBool())
	}
	if !in.IdpDiscovery.IsUnknown() && !in.IdpDiscovery.IsNull() {
		result.IdpDiscovery = Bool(in.IdpDiscovery.ValueBool())
	}
	if !in.InErrorResource.IsUnknown() && !in.InErrorResource.IsNull() {
		result.InErrorResource = Bool(in.InErrorResource.ValueBool())
	}
	if !in.RequireHttps.IsUnknown() && !in.RequireHttps.IsNull() {
		result.RequireHttps = Bool(in.RequireHttps.ValueBool())
	}
	if !in.TargetResourceSLO.IsUnknown() && !in.TargetResourceSLO.IsNull() {
		result.TargetResourceSLO = Bool(in.TargetResourceSLO.ValueBool())
	}
	if !in.TargetResourceSSO.IsUnknown() && !in.TargetResourceSSO.IsNull() {
		result.TargetResourceSSO = Bool(in.TargetResourceSSO.ValueBool())
	}
	if !in.ValidDomain.IsUnknown() && !in.ValidDomain.IsNull() {
		result.ValidDomain = String(in.ValidDomain.ValueString())
	}
	if !in.ValidPath.IsUnknown() && !in.ValidPath.IsNull() {
		result.ValidPath = String(in.ValidPath.ValueString())
	}

	return &result
}

func expandSourceTypeIdKey(in SourceTypeIdKeyData) *pf.SourceTypeIdKey {
	var result pf.SourceTypeIdKey
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
	}
	if !in.Type.IsUnknown() && !in.Type.IsNull() {
		result.Type = String(in.Type.ValueString())
	}

	return &result
}

func expandX509File(in X509FileData) *pf.X509File {
	var result pf.X509File
	if !in.CryptoProvider.IsUnknown() && !in.CryptoProvider.IsNull() {
		result.CryptoProvider = String(in.CryptoProvider.ValueString())
	}
	if !in.FileData.IsUnknown() && !in.FileData.IsNull() {
		result.FileData = String(in.FileData.ValueString())
	}
	if !in.Id.IsUnknown() && !in.Id.IsNull() {
		result.Id = String(in.Id.ValueString())
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

func expandStringSet(in types.Set) *[]*string {
	results := []*string{}
	_ = in.ElementsAs(context.Background(), &results, false)
	return &results
}

func expandStringList(in types.List) *[]*string {
	results := []*string{}
	_ = in.ElementsAs(context.Background(), &results, false)
	return &results
}
