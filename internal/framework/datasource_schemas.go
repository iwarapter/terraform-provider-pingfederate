package framework

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func datasourceAuthorizationServerSettings() schema.Schema {
	return schema.Schema{
		Description: `Authorization Server Settings attributes.`,
		Attributes: map[string]schema.Attribute{
			"activation_code_check_mode": schema.StringAttribute{
				Description: `Determines whether the user is prompted to enter or confirm the activation code after authenticating or before. The default is AFTER_AUTHENTICATION.`,
				Computed:    true,
			},
			"admin_web_service_pcv_ref": schema.StringAttribute{
				Description: `The password credential validator reference that is used for authenticating access to the OAuth Administrative Web Service.`,
				Computed:    true,
			},
			"allow_unidentified_client_extension_grants": schema.BoolAttribute{
				Description: `Allow unidentified clients to request extension grants. The default value is false.`,
				Computed:    true,
			},
			"allow_unidentified_client_ro_creds": schema.BoolAttribute{
				Description: `Allow unidentified clients to request resource owner password credentials grants. The default value is false.`,
				Computed:    true,
			},
			"allowed_origins": schema.ListAttribute{
				Description: `The list of allowed origins.`,
				Computed:    true,
				ElementType: types.StringType,
			},
			"approved_scopes_attribute": schema.StringAttribute{
				Description: `Attribute from the external consent adapter's contract, intended for storing approved scopes returned by the external consent page.`,
				Computed:    true,
			},
			"atm_id_for_o_auth_grant_management": schema.StringAttribute{
				Description: `The ID of the Access Token Manager used for OAuth enabled grant management.`,
				Computed:    true,
			},
			"authorization_code_entropy": schema.NumberAttribute{
				Description: `The authorization code entropy, in bytes.`,
				Computed:    true,
			},
			"authorization_code_timeout": schema.NumberAttribute{
				Description: `The authorization code timeout, in seconds.`,
				Computed:    true,
			},
			"bypass_activation_code_confirmation": schema.BoolAttribute{
				Description: `Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device.`,
				Computed:    true,
			},
			"bypass_authorization_for_approved_grants": schema.BoolAttribute{
				Description: `Bypass authorization for previously approved persistent grants. The default value is false.`,
				Computed:    true,
			},
			"client_secret_retention_period": schema.NumberAttribute{
				Description: `The length of time in minutes that client secrets will be retained as secondary secrets after secret change. The default value is 0, which will disable secondary client secret retention.`,
				Computed:    true,
			},
			"default_scope_description": schema.StringAttribute{
				Description: `The default scope description.`,
				Computed:    true,
			},
			"device_polling_interval": schema.NumberAttribute{
				Description: `The amount of time client should wait between polling requests, in seconds.`,
				Computed:    true,
			},
			"disallow_plain_pkce": schema.BoolAttribute{
				Description: `Determines whether PKCE's 'plain' code challenge method will be disallowed. The default value is false.`,
				Computed:    true,
			},
			"exclusive_scope_groups": schema.ListNestedAttribute{
				Description: `The list of exclusive scope groups.`,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: listScopeGroupEntry(),
				},
			},
			"exclusive_scopes": schema.ListNestedAttribute{
				Description: `The list of exclusive scopes.`,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: listScopeEntry(),
				},
			},
			"id": schema.StringAttribute{
				Description: ``,
				Computed:    true,
			},
			"include_issuer_in_authorization_response": schema.BoolAttribute{
				Description: `Determines whether the authorization server's issuer value is added to the authorization response or not. The default value is false.`,
				Computed:    true,
			},
			"jwt_secured_authorization_response_mode_lifetime": schema.NumberAttribute{
				Description: `The lifetime, in seconds, of the JWT Secured authorization response. The default value is 600.`,
				Computed:    true,
			},
			"par_reference_length": schema.NumberAttribute{
				Description: `The entropy of pushed authorization request references, in bytes. The default value is 24.`,
				Computed:    true,
			},
			"par_reference_timeout": schema.NumberAttribute{
				Description: `The timeout, in seconds, of the pushed authorization request reference. The default value is 60.`,
				Computed:    true,
			},
			"par_status": schema.StringAttribute{
				Description: `The status of pushed authorization request support. The default value is ENABLED.`,
				Computed:    true,
			},
			"pending_authorization_timeout": schema.NumberAttribute{
				Description: `The 'device_code' and 'user_code' timeout, in seconds.`,
				Computed:    true,
			},
			"persistent_grant_contract": schema.SingleNestedAttribute{
				Description: `The persistent grant contract defines attributes that are associated with OAuth persistent grants.`,
				Computed:    true,
				Attributes:  singlePersistentGrantContract(),
			},
			"persistent_grant_idle_timeout": schema.NumberAttribute{
				Description: `The persistent grant idle timeout. The default value is 30 (days). -1 indicates an indefinite amount of time.`,
				Computed:    true,
			},
			"persistent_grant_idle_timeout_time_unit": schema.StringAttribute{
				Description: `The persistent grant idle timeout time unit.`,
				Computed:    true,
			},
			"persistent_grant_lifetime": schema.NumberAttribute{
				Description: `The persistent grant lifetime. The default value is indefinite. -1 indicates an indefinite amount of time.`,
				Computed:    true,
			},
			"persistent_grant_lifetime_unit": schema.StringAttribute{
				Description: `The persistent grant lifetime unit.`,
				Computed:    true,
			},
			"persistent_grant_reuse_grant_types": schema.ListAttribute{
				Description: `The grant types that the OAuth AS can reuse rather than creating a new grant for each request. Only 'IMPLICIT' or 'AUTHORIZATION_CODE' or 'RESOURCE_OWNER_CREDENTIALS' are valid grant types.`,
				Computed:    true,
				ElementType: types.StringType,
			},
			"refresh_rolling_interval": schema.NumberAttribute{
				Description: `The minimum interval to roll refresh tokens, in hours.`,
				Computed:    true,
			},
			"refresh_token_length": schema.NumberAttribute{
				Description: `The refresh token length in number of characters.`,
				Computed:    true,
			},
			"refresh_token_rolling_grace_period": schema.NumberAttribute{
				Description: `The grace period that a rolled refresh token remains valid in seconds. The default value is 0.`,
				Computed:    true,
			},
			"registered_authorization_path": schema.StringAttribute{
				Description: `The Registered Authorization Path is concatenated to PingFederate base URL to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request. PingFederate listens to this path if specified`,
				Computed:    true,
			},
			"roll_refresh_token_values": schema.BoolAttribute{
				Description: `The roll refresh token values default policy. The default value is true.`,
				Computed:    true,
			},
			"scope_for_o_auth_grant_management": schema.StringAttribute{
				Description: `The OAuth scope to validate when accessing grant management service.`,
				Computed:    true,
			},
			"scope_groups": schema.ListNestedAttribute{
				Description: `The list of common scope groups.`,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: listScopeGroupEntry(),
				},
			},
			"scopes": schema.ListNestedAttribute{
				Description: `The list of common scopes.`,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: listScopeEntry(),
				},
			},
			"token_endpoint_base_url": schema.StringAttribute{
				Description: `The token endpoint base URL used to validate the 'aud' claim during Private Key JWT Client Authentication.`,
				Computed:    true,
			},
			"track_user_sessions_for_logout": schema.BoolAttribute{
				Description: `Determines whether user sessions are tracked for logout. If this property is not provided on a PUT, the setting is left unchanged.`,
				Computed:    true,
			},
			"user_authorization_consent_adapter": schema.StringAttribute{
				Description: `Adapter ID of the external consent adapter to be used for the consent page user interface.`,
				Computed:    true,
			},
			"user_authorization_consent_page_setting": schema.StringAttribute{
				Description: `User Authorization Consent Page setting to use PingFederate's internal consent page or an external system`,
				Computed:    true,
			},
			"user_authorization_url": schema.StringAttribute{
				Description: `The URL used to generate 'verification_url' and 'verification_url_complete' values in a Device Authorization request`,
				Computed:    true,
			},
		},
	}
}

func datasourceJdbcDataStores() schema.Schema {
	return schema.Schema{
		Description: `List of available JDBC data stores`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: ``,
				Computed:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: ``,
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: listJdbcDataStore(),
				},
			},
		},
	}
}

func listJdbcDataStore() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"allow_multi_value_attributes": schema.BoolAttribute{
			Description: `Indicates that this data store can select more than one record from a column and return the results as a multi-value attribute.`,
			Computed:    true,
		},
		"blocking_timeout": schema.NumberAttribute{
			Description: `The amount of time in milliseconds a request waits to get a connection from the connection pool before it fails. Omitting this attribute will set the value to the connection pool default.`,
			Computed:    true,
		},
		"connection_url": schema.StringAttribute{
			Description: `The default location of the JDBC database. This field is required if no mapping for JDBC database location and tags are specified.`,
			Computed:    true,
		},
		"connection_url_tags": schema.ListNestedAttribute{
			Description: `The set of connection URLs and associated tags for this JDBC data store.`,
			Computed:    true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: listJdbcTagConfig(),
			},
		},
		"driver_class": schema.StringAttribute{
			Description: `The name of the driver class used to communicate with the source database.`,
			Computed:    true,
		},
		"encrypted_password": schema.StringAttribute{
			Description: `The encrypted password needed to access the database. If you do not want to update the stored value, this attribute should be passed back unchanged. Secret Reference may be provided in this field with format 'OBF:MGR:{secretManagerId}:{secretId}'.`,
			Computed:    true,
		},
		"id": schema.StringAttribute{
			Description: `The persistent, unique ID for the data store. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.`,
			Computed:    true,
		},
		"idle_timeout": schema.NumberAttribute{
			Description: `The length of time in minutes the connection can be idle in the pool before it is closed. Omitting this attribute will set the value to the connection pool default.`,
			Computed:    true,
		},
		"mask_attribute_values": schema.BoolAttribute{
			Description: `Whether attribute values should be masked in the log.`,
			Computed:    true,
		},
		"max_pool_size": schema.NumberAttribute{
			Description: `The largest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.`,
			Computed:    true,
		},
		"min_pool_size": schema.NumberAttribute{
			Description: `The smallest number of database connections in the connection pool for the given data store. Omitting this attribute will set the value to the connection pool default.`,
			Computed:    true,
		},
		"name": schema.StringAttribute{
			Description: `The data store name with a unique value across all data sources. Omitting this attribute will set the value to a combination of the connection url and the username.`,
			Computed:    true,
		},
		"password": schema.StringAttribute{
			Description: `The password needed to access the database. GETs will not return this attribute. To update this field, specify the new value in this attribute.`,
			Computed:    true,
		},
		"type": schema.StringAttribute{
			Description: `The data store type.`,
			Computed:    true,
			Validators: []validator.String{
				stringvalidator.OneOf("LDAP", "PING_ONE_LDAP_GATEWAY", "JDBC", "CUSTOM"),
			},
		},
		"user_name": schema.StringAttribute{
			Description: `The name that identifies the user when connecting to the database.`,
			Computed:    true,
		},
		"validate_connection_sql": schema.StringAttribute{
			Description: `A simple SQL statement used by PingFederate at runtime to verify that the database connection is still active and to reconnect if needed.`,
			Computed:    true,
		},
	}
}

func listJdbcTagConfig() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"connection_url": schema.StringAttribute{
			Description: `The location of the JDBC database. `,
			Computed:    true,
		},
		"default_source": schema.BoolAttribute{
			Description: `Whether this is the default connection. Defaults to false if not specified.`,
			Computed:    true,
		},
		"tags": schema.StringAttribute{
			Description: `Tags associated with this data source.`,
			Computed:    true,
		},
	}
}

func listPersistentGrantAttribute() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Description: `The name of this attribute.`,
			Computed:    true,
		},
	}
}

func singlePersistentGrantContract() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"core_attributes": schema.ListNestedAttribute{
			Description: `This is a read-only list of persistent grant attributes and includes USER_KEY and USER_NAME. Changes to this field will be ignored.`,
			Computed:    true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: listPersistentGrantAttribute(),
			},
		},
		"extended_attributes": schema.ListNestedAttribute{
			Description: `A list of additional attributes for the persistent grant contract.`,
			Computed:    true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: listPersistentGrantAttribute(),
			},
		},
	}
}

func listScopeEntry() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Description: `The description of the scope that appears when the user is prompted for authorization.`,
			Computed:    true,
		},
		"dynamic": schema.BoolAttribute{
			Description: `True if the scope is dynamic. (Defaults to false)`,
			Computed:    true,
		},
		"name": schema.StringAttribute{
			Description: `The name of the scope.`,
			Computed:    true,
		},
	}
}

func listScopeGroupEntry() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"description": schema.StringAttribute{
			Description: `The description of the scope group.`,
			Computed:    true,
		},
		"name": schema.StringAttribute{
			Description: `The name of the scope group.`,
			Computed:    true,
		},
		"scopes": schema.ListAttribute{
			Description: `The set of scopes for this scope group.`,
			Computed:    true,
			ElementType: types.StringType,
		},
	}
}
