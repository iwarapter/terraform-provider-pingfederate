package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
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

func (r *pingfederateOAuthClientResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return resourceClient(), nil
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
		//resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), &originalSecret)...)
		resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("client_auth").AtName("secret"), &originalSecret)...)
		//resp.Diagnostics.Append(resp.State.SetAttribute(ctx, tftypes.NewAttributePath().WithAttributeName("client_auth").WithAttributeName("secret"), originalSecret)...)
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
					ExclusiveScopes:                          clientDataV0.ExclusiveScopes,
					GrantTypes:                               clientDataV0.GrantTypes,
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
					PersistentGrantReuseGrantTypes:               nil,
					PersistentGrantReuseType:                     types.StringNull(),
					RedirectUris:                                 clientDataV0.RedirectUris,
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
					RestrictedResponseTypes:                      clientDataV0.RestrictedResponseTypes,
					RestrictedScopes:                             clientDataV0.RestrictedScopes,
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
						LogoutUris:                             clientDataV0.OidcPolicy[0].LogoutUris,
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
					clientDataV1.ExtendedParameters[property.KeyName.ValueString()] = &ParameterValuesData{Values: property.Values}
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, &clientDataV1)...)
				if resp.Diagnostics.HasError() {
					return
				}
			},
		},
	}
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

func resourceClientV0() tfsdk.Schema {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Required: true,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
			},
			"client_id": {
				Type:     types.StringType,
				Required: true,
			},
			"enabled": {
				Type:     types.BoolType,
				Optional: true,
			},
			"grant_types": {
				Type: types.SetType{
					ElemType: types.StringType,
				},
				Required: true,
			},
			"bypass_approval_page": {
				Type:     types.BoolType,
				Optional: true,
			},
			"description": {
				Type:     types.StringType,
				Optional: true,
			},
			"exclusive_scopes": {
				Type: types.SetType{
					ElemType: types.StringType,
				},
				Optional: true,
			},
			"logo_url": {
				Type:     types.StringType,
				Optional: true,
			},
			"persistent_grant_expiration_time": {
				Type:     types.NumberType,
				Optional: true,
			},
			"persistent_grant_expiration_time_unit": {
				Type:     types.StringType,
				Optional: true,
			},
			"persistent_grant_expiration_type": {
				Type:     types.StringType,
				Optional: true,
			},
			"redirect_uris": {
				Type: types.SetType{
					ElemType: types.StringType,
				},
				Optional: true,
			},
			"refresh_rolling": {
				Type:     types.StringType,
				Optional: true,
			},
			"require_signed_requests": {
				Type:     types.BoolType,
				Optional: true,
			},
			"restrict_scopes": {
				Type:     types.BoolType,
				Optional: true,
			},
			"restricted_response_types": {
				Type: types.SetType{
					ElemType: types.StringType,
				},
				Optional: true,
			},
			"restricted_scopes": {
				Type: types.SetType{
					ElemType: types.StringType,
				},
				Optional: true,
			},
			"validate_using_all_eligible_atms": {
				Type:     types.BoolType,
				Optional: true,
			},
			"client_auth": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"client_cert_issuer_dn": {
						Type:     types.StringType,
						Optional: true,
					},
					"client_cert_subject_dn": {
						Type:     types.StringType,
						Optional: true,
					},
					"enforce_replay_prevention": {
						Type:     types.BoolType,
						Optional: true,
					},
					//TODO do we enable Secret/EncryptedSecret??
					"secret": {
						Type:      types.StringType,
						Optional:  true,
						Sensitive: true,
					},
					"type": {
						Type:     types.StringType,
						Required: true,
					},
					"token_endpoint_auth_signing_algorithm": {
						Type:     types.StringType,
						Optional: true,
					},
				}),
			},
			"default_access_token_manager_ref": {
				Optional:   true,
				Attributes: tfsdk.ListNestedAttributes(legacyResourceLinkSchema()),
			},
			"oidc_policy": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"grant_access_session_revocation_api": {
						Type:     types.BoolType,
						Optional: true,
					},
					"pairwise_identifier_user_type": {
						Type:     types.BoolType,
						Optional: true,
					},
					"id_token_signing_algorithm": {
						Type:     types.StringType,
						Optional: true,
					},
					"id_token_encryption_algorithm": {
						Type:     types.StringType,
						Optional: true,
					},
					"id_token_content_encryption_algorithm": {
						Type:     types.StringType,
						Optional: true,
					},
					"sector_identifier_uri": {
						Type:     types.StringType,
						Optional: true,
					},
					"logout_uris": {
						Type:     types.ListType{ElemType: types.StringType},
						Optional: true,
					},
					"ping_access_logout_capable": {
						Type:     types.BoolType,
						Optional: true,
					},
					"policy_group": {
						Optional:   true,
						Attributes: tfsdk.ListNestedAttributes(legacyResourceLinkSchema()),
					},
				}),
			},
			"jwks_settings": {
				Optional: true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"jwks": {
						Type:     types.StringType,
						Optional: true,
					},
					"jwks_url": {
						Type:     types.StringType,
						Optional: true,
					},
				}),
			},
			"ciba_delivery_mode": {
				Type:     types.StringType,
				Optional: true,
			},
			"ciba_notification_endpoint": {
				Type:     types.StringType,
				Optional: true,
			},
			"ciba_polling_interval": {
				Type:     types.NumberType,
				Optional: true,
			},
			"ciba_request_object_signing_algorithm": {
				Type:     types.StringType,
				Optional: true,
			},
			"ciba_require_signed_requests": {
				Type:     types.BoolType,
				Optional: true,
			},
			"ciba_user_code_supported": {
				Type:     types.BoolType,
				Optional: true,
			},
			"bypass_activation_code_confirmation_override": {
				Type:     types.BoolType,
				Optional: true,
			},
			"device_flow_setting_type": {
				Type:     types.StringType,
				Optional: true,
			},
			"device_polling_interval_override": {
				Type:     types.NumberType,
				Optional: true,
			},
			"extended_properties": {
				Attributes: tfsdk.SetNestedAttributes(legacyResourceParameterValues()),
				Optional:   true,
			},
			"pending_authorization_timeout_override": {
				Type:     types.NumberType,
				Optional: true,
			},
			"persistent_grant_idle_timeout": {
				Type:        types.NumberType,
				Description: "The persistent grant idle timeout.",
				Optional:    true,
			},
			"persistent_grant_idle_timeout_time_unit": {
				Type:     types.StringType,
				Optional: true,
			},
			"persistent_grant_idle_timeout_type": {
				Type:     types.StringType,
				Optional: true,
			},
			"request_object_signing_algorithm": {
				Type:     types.StringType,
				Optional: true,
			},
			"request_policy_ref": {
				Optional:   true,
				Attributes: tfsdk.ListNestedAttributes(legacyResourceLinkSchema()),
			},
			"require_proof_key_for_code_exchange": {
				Type:     types.BoolType,
				Optional: true,
			},
			"require_pushed_authorization_requests": {
				Type:     types.BoolType,
				Optional: true,
			},
			"token_exchange_processor_policy_ref": {
				Optional:   true,
				Attributes: tfsdk.ListNestedAttributes(legacyResourceLinkSchema()),
			},
			"user_authorization_url_override": {
				Type:     types.StringType,
				Optional: true,
			},
			"restrict_to_default_access_token_manager": {
				Type:     types.BoolType,
				Optional: true,
			},
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
