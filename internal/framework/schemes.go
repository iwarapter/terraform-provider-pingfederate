package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func resourceApcToPersistentGrantMapping() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `An authentication policy contract mapping into an OAuth persistent grant.`,
		Attributes: map[string]tfsdk.Attribute{
			"attribute_contract_fulfillment": {
				Description: `A list of mappings from attribute names to their fulfillment values.`,
				Required:    true,
				Attributes:  mapAttributeFulfillmentValue(),
			},
			"jdbc_attribute_sources": {
				Description: `The configured settings used to look up attributes from a JDBC data store.`,
				Optional:    true,
				Attributes:  listJdbcAttributeSource(),
			},
			"ldap_attribute_sources": {
				Description: `The configured settings used to look up attributes from a LDAP data store.`,
				Optional:    true,
				Attributes:  listLdapAttributeSource(),
			},
			"custom_attribute_sources": {
				Description: `The configured settings used to look up attributes from a custom data store.`,
				Optional:    true,
				Attributes:  listCustomAttributeSource(),
			},
			"authentication_policy_contract_ref": {
				Description: `Reference to the associated authentication policy contract. The reference cannot be changed after the mapping has been created.`,
				Required:    true,
				Type:        types.StringType,
			},
			"id": {
				Description: `The ID of the authentication policy contract to persistent grant mapping.`,
				Required:    true,
				Type:        types.StringType,
			},
			"issuance_criteria": {
				Description: `The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.`,
				Optional:    true,
				Attributes:  singleIssuanceCriteria(),
			},
		},
	}
}

func resourceAuthenticationPolicyContract() tfsdk.Schema {
	return tfsdk.Schema{
		Description: `Authentication Policy Contracts carry user attributes from the identity provider to the service provider.`,
		Attributes: map[string]tfsdk.Attribute{
			"core_attributes": {
				Description: `A list of read-only assertion attributes (for example, subject) that are automatically populated by PingFederate.`,
				Computed:    true,
				Attributes:  listAuthenticationPolicyContractAttribute(),
			},
			"extended_attributes": {
				Description: `A list of additional attributes as needed.`,
				Optional:    true,
				Type:        types.ListType{ElemType: types.StringType},
			},
			"id": {
				Description: `The persistent, unique ID for the authentication policy contract. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.`,
				Optional:    true,
				Type:        types.StringType,
			},
			"name": {
				Description: `The Authentication Policy Contract Name. Name is unique.`,
				Required:    true,
				Type:        types.StringType,
			},
		},
	}
}

func listJdbcAttributeSource() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  mapAttributeFulfillmentValue(),
		},
		"column_names": {
			Description: `A list of column names used to construct the SQL query to retrieve data from the specified table in the datastore.`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Type:        types.StringType,
		},
		"filter": {
			Description: `The JDBC WHERE clause used to query your data store to locate a user record.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"schema": {
			Description: `Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"table": {
			Description: `The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func listLdapAttributeSource() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  mapAttributeFulfillmentValue(),
		},
		"base_dn": {
			Description: `The base DN to search from. If not specified, the search will start at the LDAP's root.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"binary_attribute_settings": {
			Description: `The advanced settings for binary LDAP attributes.`,
			Optional:    true,
			Attributes:  mapBinaryLdapAttributeSettings(),
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Type:        types.StringType,
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"member_of_nested_group": {
			Description: `Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.`,
			Optional:    true,
			Type:        types.BoolType,
		},
		"search_attributes": {
			Description: `A list of LDAP attributes returned from search and available for mapping.`,
			Optional:    true,
			Type: types.ListType{
				ElemType: types.StringType,
			},
		},
		"search_filter": {
			Description: `The LDAP filter that will be used to lookup the objects from the directory.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"search_scope": {
			Description: `Determines the node depth of the query.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func listCustomAttributeSource() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"attribute_contract_fulfillment": {
			Description: `A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings`,
			Optional:    true,
			Attributes:  mapAttributeFulfillmentValue(),
		},
		"data_store_ref": {
			Description: `Reference to the associated data store.`,
			Required:    true,
			Type:        types.StringType,
		},
		"description": {
			Description: `The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings`,
			Optional:    true,
			Type:        types.StringType,
		},
		"filter_fields": {
			Description: `The list of fields that can be used to filter a request to the custom data store.`,
			Optional:    true,
			Attributes:  listFieldEntry(),
		},
		"id": {
			Description: `The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func listConditionalIssuanceCriteriaEntry() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"attribute_name": {
			Description: `The name of the attribute to use in this issuance criterion.`,
			Required:    true,
			Type:        types.StringType,
		},
		"condition": {
			Description: `The condition that will be applied to the source attribute's value and the expected value.`,
			Required:    true,
			Type:        types.StringType,
		},
		"error_result": {
			Description: `The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"source": {
			Description: `The source of the attribute.`,
			Required:    true,
			Attributes:  singleSourceTypeIdKey(),
		},
		"value": {
			Description: `The expected value of this issuance criterion.`,
			Required:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func listExpressionIssuanceCriteriaEntry() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"error_result": {
			Description: `The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"expression": {
			Description: `The OGNL expression to evaluate.`,
			Required:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func listAuthenticationPolicyContractAttribute() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"name": {
			Description: `The name of this attribute.`,
			Required:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}

func singleSourceTypeIdKey() tfsdk.NestedAttributes {
	return tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"id": {
			Description: `The attribute source ID that refers to the attribute source that this key references. In some resources, the ID is optional and will be ignored. In these cases the ID should be omitted. If the source type is not an attribute source then the ID can be omitted.`,
			Optional:    true,
			Type:        types.StringType,
		},
		"type": {
			Description: `The source type of this key.`,
			Required:    true,
			Type:        types.StringType,
		},
	})
}

func mapBinaryLdapAttributeSettings() tfsdk.NestedAttributes {
	return tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
		"binary_encoding": {
			Description: `Get the encoding type for this attribute. If not specified, the default is BASE64.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}, tfsdk.MapNestedAttributesOptions{})
}

func singleIssuanceCriteria() tfsdk.NestedAttributes {
	return tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
		"conditional_criteria": {
			Description: `A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue.`,
			Optional:    true,
			Attributes:  listConditionalIssuanceCriteriaEntry(),
		},
		"expression_criteria": {
			Description: `A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue.`,
			Optional:    true,
			Attributes:  listExpressionIssuanceCriteriaEntry(),
		},
	})
}

func mapAttributeFulfillmentValue() tfsdk.NestedAttributes {
	return tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
		"source": {
			Description: `The attribute value source.`,
			Required:    true,
			Attributes:  singleSourceTypeIdKey(),
		},
		"value": {
			Description: `The value for this attribute.`,
			Required:    true,
			Type:        types.StringType,
		},
	}, tfsdk.MapNestedAttributesOptions{})
}

func listFieldEntry() tfsdk.NestedAttributes {
	return tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
		"name": {
			Description: `The name of this field.`,
			Required:    true,
			Type:        types.StringType,
		},
		"value": {
			Description: `The value of this field. Whether or not the value is required will be determined by plugin validation checks.`,
			Optional:    true,
			Type:        types.StringType,
		},
	}, tfsdk.ListNestedAttributesOptions{})
}
