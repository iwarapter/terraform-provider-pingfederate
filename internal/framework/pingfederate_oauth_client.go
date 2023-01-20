package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                 = &pingfederateOAuthClientResource{}
	_ resource.ResourceWithConfigure    = &pingfederateOAuthClientResource{}
	_ resource.ResourceWithImportState  = &pingfederateOAuthClientResource{}
	_ resource.ResourceWithUpgradeState = &pingfederateOAuthClientResource{}
)

type pingfederateOAuthClientResource struct {
	client *pfClient
}

func NewOAuthClientResource() resource.Resource {
	return &pingfederateOAuthClientResource{}
}

func (r *pingfederateOAuthClientResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceClient()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateOAuthClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateOAuthClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_client"
}

func (r *pingfederateOAuthClientResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Skip modification on resource creation
	if req.State.Raw.IsNull() {
		return
	}

	// Skip modification on resource destruction
	if req.Plan.Raw.IsNull() {
		return
	}

	// Skip modification if planning no changes
	if req.Plan.Raw.Equal(req.State.Raw) {
		return
	}

	var originalPlan, plan, state ClientData

	resp.Diagnostics.Append(req.Plan.Get(ctx, &originalPlan)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if r.client.IsVersionLessEqThan(11, 0) {
		plan.ClientSecretChangedTime = types.StringNull()
	}
	if plan.ClientAuth == nil {
		plan.ClientSecretChangedTime = types.StringNull()
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		return
	} else {
		if plan.ClientAuth.Secret.IsNull() {
			plan.ClientSecretChangedTime = types.StringNull()
		}
	}

	// Skip modification if encrypted_secret is known
	if !plan.ClientAuth.EncryptedSecret.IsUnknown() {
		return
	}

	// Copy encrypted_secret prior state to plan
	plan.ClientAuth.EncryptedSecret = state.ClientAuth.EncryptedSecret

	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Undo modification if there are changes outside encrypted_secret
	if !resp.Plan.Raw.Equal(req.State.Raw) {
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &originalPlan)...)
	}
}

func (r *pingfederateOAuthClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.versionRequestModifier(&data)
	body, _, err := r.client.OauthClients.CreateClientWithContext(ctx, &oauthClients.CreateClientInput{
		Body: *expandClient(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OAuthClient, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthClients.GetClientWithContext(ctx, &oauthClients.GetClientInput{Id: data.ClientId.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	r.versionRequestModifier(&data)
	body, _, err := r.client.OauthClients.UpdateClientWithContext(ctx, &oauthClients.UpdateClientInput{
		Body: *expandClient(data),
		Id:   data.ClientId.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update OAuthClient, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenClient(body)))...)
	if data.ClientAuth != nil && !data.ClientAuth.Secret.IsNull() {
		var originalSecret string
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("client_auth").AtName("secret"), originalSecret)...)
	}
}

func (r *pingfederateOAuthClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClientData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.OauthClients.DeleteClientWithContext(ctx, &oauthClients.DeleteClientInput{Id: data.ClientId.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete OAuthClient, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateOAuthClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("client_id"), req, resp)
}

func (r *pingfederateOAuthClientResource) UpgradeState(context.Context) map[int64]resource.StateUpgrader {
	schemaV0 := resourceClientV0()

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var clientDataV0 ClientDataV0
				resp.Diagnostics.Append(req.State.Get(ctx, &clientDataV0)...)
				if resp.Diagnostics.HasError() {
					return
				}

				clientDataV1 := ClientData{
					AllowAuthenticationApiInit:               types.BoolNull(),
					BypassActivationCodeConfirmationOverride: clientDataV0.BypassActivationCodeConfirmationOverride,
					BypassApprovalPage:                       clientDataV0.BypassApprovalPage,
					CibaDeliveryMode:                         clientDataV0.CibaDeliveryMode,
					CibaNotificationEndpoint:                 clientDataV0.CibaNotificationEndpoint,
					CibaPollingInterval:                      clientDataV0.CibaPollingInterval,
					CibaRequestObjectSigningAlgorithm:        clientDataV0.CibaRequestObjectSigningAlgorithm,
					CibaRequireSignedRequests:                clientDataV0.RequireSignedRequests,
					CibaUserCodeSupported:                    clientDataV0.CibaUserCodeSupported,
					ClientId:                                 clientDataV0.ClientId,
					ClientSecretChangedTime:                  types.StringNull(),
					ClientSecretRetentionPeriod:              types.NumberNull(),
					ClientSecretRetentionPeriodType:          types.StringNull(),
					Description:                              clientDataV0.Description,
					DeviceFlowSettingType:                    clientDataV0.DeviceFlowSettingType,
					DevicePollingIntervalOverride:            clientDataV0.DevicePollingIntervalOverride,
					Enabled:                                  clientDataV0.Enabled,
					ExclusiveScopes:                          sliceStringTypeToList(clientDataV0.ExclusiveScopes),
					GrantTypes:                               sliceStringTypeToList(clientDataV0.GrantTypes),
					Id:                                       clientDataV0.ClientId,
					JwtSecuredAuthorizationResponseModeContentEncryptionAlgorithm: types.StringNull(),
					JwtSecuredAuthorizationResponseModeEncryptionAlgorithm:        types.StringNull(),
					JwtSecuredAuthorizationResponseModeSigningAlgorithm:           types.StringNull(),
					LogoUrl:                                      clientDataV0.LogoUrl,
					Name:                                         clientDataV0.Name,
					PendingAuthorizationTimeoutOverride:          clientDataV0.PendingAuthorizationTimeoutOverride,
					PersistentGrantExpirationTime:                clientDataV0.PersistentGrantExpirationTime,
					PersistentGrantExpirationTimeUnit:            clientDataV0.PersistentGrantExpirationTimeUnit,
					PersistentGrantExpirationType:                clientDataV0.PersistentGrantExpirationType,
					PersistentGrantIdleTimeout:                   clientDataV0.PersistentGrantIdleTimeout,
					PersistentGrantIdleTimeoutTimeUnit:           clientDataV0.PersistentGrantIdleTimeoutTimeUnit,
					PersistentGrantIdleTimeoutType:               clientDataV0.PersistentGrantIdleTimeoutType,
					PersistentGrantReuseGrantTypes:               types.ListNull(types.StringType),
					PersistentGrantReuseType:                     types.StringNull(),
					RedirectUris:                                 sliceStringTypeToSet(clientDataV0.RedirectUris),
					RefreshRolling:                               clientDataV0.RefreshRolling,
					RefreshTokenRollingGracePeriod:               types.NumberNull(),
					RefreshTokenRollingGracePeriodType:           types.StringNull(),
					RefreshTokenRollingInterval:                  types.NumberNull(),
					RefreshTokenRollingIntervalType:              types.StringNull(),
					RequestObjectSigningAlgorithm:                clientDataV0.RequestObjectSigningAlgorithm,
					RequireJwtSecuredAuthorizationResponseMode:   types.BoolNull(),
					RequireProofKeyForCodeExchange:               clientDataV0.RequireProofKeyForCodeExchange,
					RequirePushedAuthorizationRequests:           clientDataV0.RequirePushedAuthorizationRequests,
					RequireSignedRequests:                        clientDataV0.RequireSignedRequests,
					RestrictScopes:                               clientDataV0.RestrictScopes,
					RestrictToDefaultAccessTokenManager:          clientDataV0.RestrictToDefaultAccessTokenManager,
					RestrictedResponseTypes:                      sliceStringTypeToList(clientDataV0.RestrictedResponseTypes),
					RestrictedScopes:                             sliceStringTypeToSet(clientDataV0.RestrictedScopes),
					TokenIntrospectionContentEncryptionAlgorithm: types.StringNull(),
					TokenIntrospectionEncryptionAlgorithm:        types.StringNull(),
					TokenIntrospectionSigningAlgorithm:           types.StringNull(),
					UserAuthorizationUrlOverride:                 clientDataV0.UserAuthorizationUrlOverride,
					ValidateUsingAllEligibleAtms:                 clientDataV0.ValidateUsingAllEligibleAtms,
				}
				if len(clientDataV0.ClientAuth) == 1 {
					clientDataV1.ClientAuth = &ClientAuthData{
						ClientCertIssuerDn:                clientDataV0.ClientAuth[0].ClientCertIssuerDn,
						ClientCertSubjectDn:               clientDataV0.ClientAuth[0].ClientCertSubjectDn,
						EnforceReplayPrevention:           clientDataV0.ClientAuth[0].EnforceReplayPrevention,
						Secret:                            clientDataV0.ClientAuth[0].Secret,
						TokenEndpointAuthSigningAlgorithm: clientDataV0.ClientAuth[0].TokenEndpointAuthSigningAlgorithm,
						Type:                              clientDataV0.ClientAuth[0].Type,
					}
				}
				if len(clientDataV0.DefaultAccessTokenManagerRef) == 1 {
					clientDataV1.DefaultAccessTokenManagerRef = clientDataV0.DefaultAccessTokenManagerRef[0].ID
				}
				if len(clientDataV0.RequestPolicyRef) == 1 {
					clientDataV1.RequestPolicyRef = clientDataV0.RequestPolicyRef[0].ID
				}
				if len(clientDataV0.TokenExchangeProcessorPolicyRef) == 1 {
					clientDataV1.TokenExchangeProcessorPolicyRef = clientDataV0.TokenExchangeProcessorPolicyRef[0].ID
				}
				if len(clientDataV0.OidcPolicy) == 1 {
					clientDataV1.OidcPolicy = &ClientOIDCPolicyData{
						GrantAccessSessionRevocationApi:        clientDataV0.OidcPolicy[0].GrantAccessSessionRevocationApi,
						GrantAccessSessionSessionManagementApi: types.BoolNull(),
						IdTokenContentEncryptionAlgorithm:      clientDataV0.OidcPolicy[0].IdTokenContentEncryptionAlgorithm,
						IdTokenEncryptionAlgorithm:             clientDataV0.OidcPolicy[0].IdTokenEncryptionAlgorithm,
						IdTokenSigningAlgorithm:                clientDataV0.OidcPolicy[0].IdTokenSigningAlgorithm,
						LogoutUris:                             sliceStringTypeToList(clientDataV0.OidcPolicy[0].LogoutUris),
						PairwiseIdentifierUserType:             clientDataV0.OidcPolicy[0].PairwiseIdentifierUserType,
						PingAccessLogoutCapable:                clientDataV0.OidcPolicy[0].PingAccessLogoutCapable,
						SectorIdentifierUri:                    clientDataV0.OidcPolicy[0].SectorIdentifierUri,
					}
					if len(clientDataV0.OidcPolicy[0].PolicyGroup) == 1 {
						clientDataV1.OidcPolicy.PolicyGroup = clientDataV0.OidcPolicy[0].PolicyGroup[0].ID
					}
				}
				if len(clientDataV0.JwksSettings) == 1 {
					clientDataV1.JwksSettings = &clientDataV0.JwksSettings[0]
				}
				if len(clientDataV0.ExtendedProperties) > 0 {
					clientDataV1.ExtendedParameters = map[string]*ParameterValuesData{}
				}
				for _, property := range clientDataV0.ExtendedProperties {
					clientDataV1.ExtendedParameters[property.KeyName.ValueString()] = &ParameterValuesData{Values: sliceStringTypeToList(property.Values)}
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, &clientDataV1)...)
			},
		},
	}
}

func sliceStringTypeToList(slice []types.String) types.List {
	if len(slice) == 0 {
		return types.ListNull(types.StringType)
	}
	strs := []attr.Value{}
	for _, value := range slice {
		strs = append(strs, value)
	}
	return types.ListValueMust(types.StringType, strs)
}

func sliceStringTypeToSet(slice []types.String) types.Set {
	if len(slice) == 0 {
		return types.SetNull(types.StringType)
	}
	strs := []attr.Value{}
	for _, value := range slice {
		strs = append(strs, value)
	}
	return types.SetValueMust(types.StringType, strs)
}

// old version of pingfederate dont handle the follow fields so we strip them before marshalling
func (r *pingfederateOAuthClientResource) versionRequestModifier(data *ClientData) {
	if !r.client.IsVersionGreaterEqThan(10, 3) {
		data.RefreshTokenRollingIntervalType = types.StringNull()
	}
	if !r.client.IsVersionGreaterEqThan(11, 0) {
		data.PersistentGrantReuseType = types.StringNull()
		data.RefreshTokenRollingGracePeriodType = types.StringNull()
	}
	if !r.client.IsVersionGreaterEqThan(11, 1) {
		data.ClientSecretRetentionPeriodType = types.StringNull()
		data.RequireJwtSecuredAuthorizationResponseMode = types.BoolNull()
	}
}

// old version of pingfederate dont handle the follow fields so add the defaults back to keep state happy!
func (r *pingfederateOAuthClientResource) versionResponseModifier(data *ClientData) *ClientData {
	if !r.client.IsVersionGreaterEqThan(10, 3) {
		data.RefreshTokenRollingIntervalType = types.StringValue("SERVER_DEFAULT")
	}
	if !r.client.IsVersionGreaterEqThan(11, 0) {
		data.PersistentGrantReuseType = types.StringValue("SERVER_DEFAULT")
		data.RefreshTokenRollingGracePeriodType = types.StringValue("SERVER_DEFAULT")
	}
	if !r.client.IsVersionGreaterEqThan(11, 1) {
		data.ClientSecretRetentionPeriodType = types.StringValue("SERVER_DEFAULT")
		data.RequireJwtSecuredAuthorizationResponseMode = types.BoolValue(false)
	}

	return data
}

func resourceClientV0() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id":                                    schema.StringAttribute{Required: true},
			"name":                                  schema.StringAttribute{Required: true},
			"client_id":                             schema.StringAttribute{Required: true},
			"enabled":                               schema.BoolAttribute{Optional: true},
			"grant_types":                           schema.SetAttribute{ElementType: types.StringType, Required: true},
			"bypass_approval_page":                  schema.BoolAttribute{Optional: true},
			"description":                           schema.StringAttribute{Optional: true},
			"exclusive_scopes":                      schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"logo_url":                              schema.StringAttribute{Optional: true},
			"persistent_grant_expiration_time":      schema.NumberAttribute{Optional: true},
			"persistent_grant_expiration_time_unit": schema.StringAttribute{Optional: true},
			"persistent_grant_expiration_type":      schema.StringAttribute{Optional: true},
			"redirect_uris":                         schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"refresh_rolling":                       schema.StringAttribute{Optional: true},
			"require_signed_requests":               schema.BoolAttribute{Optional: true},
			"restrict_scopes":                       schema.BoolAttribute{Optional: true},
			"restricted_response_types":             schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"restricted_scopes":                     schema.SetAttribute{ElementType: types.StringType, Optional: true},
			"validate_using_all_eligible_atms":      schema.BoolAttribute{Optional: true},
			"client_auth": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"client_cert_issuer_dn":     schema.StringAttribute{Optional: true},
						"client_cert_subject_dn":    schema.StringAttribute{Optional: true},
						"enforce_replay_prevention": schema.BoolAttribute{Optional: true},
						//TODO do we enable Secret/EncryptedSecret??
						"secret":                                schema.StringAttribute{Optional: true, Sensitive: true},
						"type":                                  schema.StringAttribute{Required: true},
						"token_endpoint_auth_signing_algorithm": schema.StringAttribute{Optional: true},
					},
				},
			},
			"default_access_token_manager_ref": resourceLinkSchemaV0(),
			"oidc_policy": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"grant_access_session_revocation_api":   schema.BoolAttribute{Optional: true},
						"pairwise_identifier_user_type":         schema.BoolAttribute{Optional: true},
						"id_token_signing_algorithm":            schema.StringAttribute{Optional: true},
						"id_token_encryption_algorithm":         schema.StringAttribute{Optional: true},
						"id_token_content_encryption_algorithm": schema.StringAttribute{Optional: true},
						"sector_identifier_uri":                 schema.StringAttribute{Optional: true},
						"logout_uris":                           schema.ListAttribute{ElementType: types.StringType, Optional: true},
						"ping_access_logout_capable":            schema.BoolAttribute{Optional: true},
						"policy_group":                          resourceLinkSchemaV0(),
					},
				},
			},
			"jwks_settings": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jwks":     schema.StringAttribute{Optional: true},
						"jwks_url": schema.StringAttribute{Optional: true},
					},
				},
			},
			"ciba_delivery_mode":                           schema.StringAttribute{Optional: true},
			"ciba_notification_endpoint":                   schema.StringAttribute{Optional: true},
			"ciba_polling_interval":                        schema.NumberAttribute{Optional: true},
			"ciba_request_object_signing_algorithm":        schema.StringAttribute{Optional: true},
			"ciba_require_signed_requests":                 schema.BoolAttribute{Optional: true},
			"ciba_user_code_supported":                     schema.BoolAttribute{Optional: true},
			"bypass_activation_code_confirmation_override": schema.BoolAttribute{Optional: true},
			"device_flow_setting_type":                     schema.StringAttribute{Optional: true},
			"device_polling_interval_override":             schema.NumberAttribute{Optional: true},
			"extended_properties":                          legacyResourceParameterValues(),
			"pending_authorization_timeout_override":       schema.NumberAttribute{Optional: true},
			"persistent_grant_idle_timeout":                schema.NumberAttribute{Optional: true},
			"persistent_grant_idle_timeout_time_unit":      schema.StringAttribute{Optional: true},
			"persistent_grant_idle_timeout_type":           schema.StringAttribute{Optional: true},
			"request_object_signing_algorithm":             schema.StringAttribute{Optional: true},
			"request_policy_ref":                           resourceLinkSchemaV0(),
			"require_proof_key_for_code_exchange":          schema.BoolAttribute{Optional: true},
			"require_pushed_authorization_requests":        schema.BoolAttribute{Optional: true},
			"token_exchange_processor_policy_ref":          resourceLinkSchemaV0(),
			"user_authorization_url_override":              schema.StringAttribute{Optional: true},
			"restrict_to_default_access_token_manager":     schema.BoolAttribute{Optional: true},
		},
	}
}

type ResourceLink struct {
	ID       types.String `tfsdk:"id"`
	Location types.String `tfsdk:"location"`
}
type ClientOIDCPolicyDataV0 struct {
	GrantAccessSessionRevocationApi   types.Bool     `tfsdk:"grant_access_session_revocation_api"`
	IdTokenContentEncryptionAlgorithm types.String   `tfsdk:"id_token_content_encryption_algorithm"`
	IdTokenEncryptionAlgorithm        types.String   `tfsdk:"id_token_encryption_algorithm"`
	IdTokenSigningAlgorithm           types.String   `tfsdk:"id_token_signing_algorithm"`
	LogoutUris                        []types.String `tfsdk:"logout_uris"`
	PairwiseIdentifierUserType        types.Bool     `tfsdk:"pairwise_identifier_user_type"`
	PingAccessLogoutCapable           types.Bool     `tfsdk:"ping_access_logout_capable"`
	PolicyGroup                       []ResourceLink `tfsdk:"policy_group"`
	SectorIdentifierUri               types.String   `tfsdk:"sector_identifier_uri"`
}
type ClientAuthDataV0 struct {
	ClientCertIssuerDn                types.String `tfsdk:"client_cert_issuer_dn"`
	ClientCertSubjectDn               types.String `tfsdk:"client_cert_subject_dn"`
	EnforceReplayPrevention           types.Bool   `tfsdk:"enforce_replay_prevention"`
	Secret                            types.String `tfsdk:"secret"`
	TokenEndpointAuthSigningAlgorithm types.String `tfsdk:"token_endpoint_auth_signing_algorithm"`
	Type                              types.String `tfsdk:"type"`
}
type ClientDataV0 struct {
	ID                                       types.String               `tfsdk:"id"`
	Name                                     types.String               `tfsdk:"name"`
	ClientId                                 types.String               `tfsdk:"client_id"`
	Enabled                                  types.Bool                 `tfsdk:"enabled"`
	GrantTypes                               []types.String             `tfsdk:"grant_types"`
	PersistentGrantExpirationTime            types.Number               `tfsdk:"persistent_grant_expiration_time"`
	RequireProofKeyForCodeExchange           types.Bool                 `tfsdk:"require_proof_key_for_code_exchange"`
	RestrictedScopes                         []types.String             `tfsdk:"restricted_scopes"`
	CibaDeliveryMode                         types.String               `tfsdk:"ciba_delivery_mode"`
	PendingAuthorizationTimeoutOverride      types.Number               `tfsdk:"pending_authorization_timeout_override"`
	RequestObjectSigningAlgorithm            types.String               `tfsdk:"request_object_signing_algorithm"`
	RestrictScopes                           types.Bool                 `tfsdk:"restrict_scopes"`
	TokenExchangeProcessorPolicyRef          []ResourceLink             `tfsdk:"token_exchange_processor_policy_ref"`
	CibaPollingInterval                      types.Number               `tfsdk:"ciba_polling_interval"`
	RedirectUris                             []types.String             `tfsdk:"redirect_uris"`
	PersistentGrantIdleTimeout               types.Number               `tfsdk:"persistent_grant_idle_timeout"`
	DevicePollingIntervalOverride            types.Number               `tfsdk:"device_polling_interval_override"`
	DeviceFlowSettingType                    types.String               `tfsdk:"device_flow_setting_type"`
	LogoUrl                                  types.String               `tfsdk:"logo_url"`
	OidcPolicy                               []ClientOIDCPolicyDataV0   `tfsdk:"oidc_policy"`
	PersistentGrantExpirationTimeUnit        types.String               `tfsdk:"persistent_grant_expiration_time_unit"`
	RequirePushedAuthorizationRequests       types.Bool                 `tfsdk:"require_pushed_authorization_requests"`
	UserAuthorizationUrlOverride             types.String               `tfsdk:"user_authorization_url_override"`
	BypassActivationCodeConfirmationOverride types.Bool                 `tfsdk:"bypass_activation_code_confirmation_override"`
	ValidateUsingAllEligibleAtms             types.Bool                 `tfsdk:"validate_using_all_eligible_atms"`
	CibaUserCodeSupported                    types.Bool                 `tfsdk:"ciba_user_code_supported"`
	Description                              types.String               `tfsdk:"description"`
	JwksSettings                             []JwksSettingsData         `tfsdk:"jwks_settings"`
	RefreshRolling                           types.String               `tfsdk:"refresh_rolling"`
	RequestPolicyRef                         []ResourceLink             `tfsdk:"request_policy_ref"`
	RequireSignedRequests                    types.Bool                 `tfsdk:"require_signed_requests"`
	CibaNotificationEndpoint                 types.String               `tfsdk:"ciba_notification_endpoint"`
	CibaRequireSignedRequests                types.Bool                 `tfsdk:"ciba_require_signed_requests"`
	RestrictToDefaultAccessTokenManager      types.Bool                 `tfsdk:"restrict_to_default_access_token_manager"`
	PersistentGrantExpirationType            types.String               `tfsdk:"persistent_grant_expiration_type"`
	PersistentGrantIdleTimeoutTimeUnit       types.String               `tfsdk:"persistent_grant_idle_timeout_time_unit"`
	DefaultAccessTokenManagerRef             []ResourceLink             `tfsdk:"default_access_token_manager_ref"`
	ExclusiveScopes                          []types.String             `tfsdk:"exclusive_scopes"`
	ClientAuth                               []ClientAuthDataV0         `tfsdk:"client_auth"`
	PersistentGrantIdleTimeoutType           types.String               `tfsdk:"persistent_grant_idle_timeout_type"`
	RestrictedResponseTypes                  []types.String             `tfsdk:"restricted_response_types"`
	BypassApprovalPage                       types.Bool                 `tfsdk:"bypass_approval_page"`
	CibaRequestObjectSigningAlgorithm        types.String               `tfsdk:"ciba_request_object_signing_algorithm"`
	ExtendedProperties                       []ExtendedPropertiesDataV0 `tfsdk:"extended_properties"`
}

type ExtendedPropertiesDataV0 struct {
	KeyName types.String   `tfsdk:"key_name"`
	Values  []types.String `tfsdk:"values"`
}
