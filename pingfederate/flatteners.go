package pingfederate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

//func flattenSpAdapterAttribute(in *pf.SpAdapterAttribute) []map[string]interface{} {
//	m := make([]map[string]interface{}, 0, 1)
//	s := make(map[string]interface{})
//	if in.Name != nil {
//		s["name"] = *in.Name
//	}
//	return append(m, s)
//}

func flattenSpAdapterAttributes(in *[]*pf.SpAdapterAttribute) *schema.Set {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, *v.Name)
	}
	return schema.NewSet(schema.HashString, m)
}

func flattenSpAdapterAttributeContract(in *pf.SpAdapterAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Inherited != nil {
		s["inherited"] = *in.Inherited
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenSpAdapterAttributes(in.ExtendedAttributes)
	}
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenSpAdapterAttributes(in.CoreAttributes)
	}
	m = append(m, s)
	return m
}

func flattenSpAdapterTargetApplicationInfo(in *pf.SpAdapterTargetApplicationInfo) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Inherited != nil {
		s["inherited"] = *in.Inherited
	}
	if in.ApplicationName != nil {
		s["application_name"] = *in.ApplicationName
	}
	if in.ApplicationIconUrl != nil {
		s["application_icon_url"] = *in.ApplicationIconUrl
	}
	m = append(m, s)
	return m
}

func flattenIdpAdapterContractMapping(in *pf.IdpAdapterContractMapping) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Inherited != nil {
		s["inherited"] = *in.Inherited
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && (in.IssuanceCriteria.ExpressionCriteria != nil && in.IssuanceCriteria.ConditionalCriteria != nil) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}

	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	m = append(m, s)
	return m
}

func flattenLdapAttributeSource(in *pf.AttributeSource) map[string]interface{} {
	s := make(map[string]interface{})
	if in.DataStoreRef != nil {
		s["data_store_ref"] = flattenResourceLink(in.DataStoreRef)
	}
	if in.BaseDn != nil {
		s["base_dn"] = *in.BaseDn
	}
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.SearchScope != nil {
		s["search_scope"] = *in.SearchScope
	}
	if in.Description != nil {
		s["description"] = *in.Description
	}
	if in.SearchFilter != nil {
		s["search_filter"] = *in.SearchFilter
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.BinaryAttributeSettings != nil {
		attributes := map[string]string{}
		for s2 := range in.BinaryAttributeSettings {
			attributes[s2] = *(*in.BinaryAttributeSettings[s2]).BinaryEncoding
		}
		s["binary_attribute_settings"] = attributes
	}
	if in.MemberOfNestedGroup != nil {
		s["member_of_nested_group"] = *in.MemberOfNestedGroup
	}
	return s
}

func flattenJdbcAttributeSource(in *pf.AttributeSource) map[string]interface{} {
	s := make(map[string]interface{})
	if in.DataStoreRef != nil {
		s["data_store_ref"] = flattenResourceLink(in.DataStoreRef)
	}
	if in.Schema != nil {
		s["schema"] = *in.Schema
	}
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Table != nil {
		s["table"] = *in.Table
	}
	if in.Description != nil {
		s["description"] = *in.Description
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.Filter != nil {
		s["filter"] = *in.Filter
	}
	return s
}

func flattenCustomAttributeSource(in *pf.AttributeSource) map[string]interface{} {
	s := make(map[string]interface{})
	if in.DataStoreRef != nil {
		s["data_store_ref"] = flattenResourceLink(in.DataStoreRef)
	}
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Description != nil {
		s["description"] = *in.Description
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.FilterFields != nil {
		s["filter_fields"] = flattenFieldEntry(in.FilterFields)
	}
	return s
}

func flattenFieldEntry(in *[]*pf.FieldEntry) []interface{} {
	var m []interface{}
	for _, v := range *in {
		s := make(map[string]interface{})
		if v.Name != nil {
			s["name"] = *v.Name
		}
		if v.Value != nil {
			s["value"] = *v.Value
		}
		m = append(m, s)
	}
	return m
}

func flattenMapOfAttributeFulfillmentValue(in map[string]*pf.AttributeFulfillmentValue) *schema.Set {
	m := make([]interface{}, 0, len(in))
	for s2 := range in {
		s := flattenAttributeFulfillmentValue(in[s2])
		s["key_name"] = s2
		m = append(m, s)
	}
	return schema.NewSet(attributeFulfillmentValueHash, m)
}

func flattenAttributeFulfillmentValue(in *pf.AttributeFulfillmentValue) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Source != nil {
		s["source"] = flattenSourceTypeIdKey(in.Source)
	}
	if in.Value != nil {
		s["value"] = *in.Value
	}
	return s
}

func flattenSourceTypeIdKey(in *pf.SourceTypeIdKey) []interface{} {
	m := make([]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Type != nil {
		s["type"] = *in.Type
	}
	m = append(m, s)
	return m
}

func flattenIssuanceCriteria(in *pf.IssuanceCriteria) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ConditionalCriteria != nil && len(*in.ConditionalCriteria) > 0 {
		s["conditional_criteria"] = flattenConditionalIssuanceCriteriaEntrys(in.ConditionalCriteria)
	}
	if in.ExpressionCriteria != nil && len(*in.ExpressionCriteria) > 0 {
		s["expression_criteria"] = flattenExpressionIssuanceCriteriaEntrys(in.ExpressionCriteria)
	}
	m = append(m, s)
	return m
}

func flattenConditionalIssuanceCriteriaEntry(in *pf.ConditionalIssuanceCriteriaEntry) map[string]interface{} {
	//m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ErrorResult != nil {
		s["error_result"] = *in.ErrorResult
	}
	if in.Source != nil {
		s["source"] = flattenSourceTypeIdKey(in.Source)
	}
	if in.AttributeName != nil {
		s["attribute_name"] = *in.AttributeName
	}
	if in.Condition != nil {
		s["condition"] = *in.Condition
	}
	if in.Value != nil {
		s["value"] = *in.Value
	}
	return s
}

//
//func flattenConditionalIssuanceCriteriaEntry(in []*pf.ConditionalIssuanceCriteriaEntry) []map[string]interface{} {
//	m := make([]map[string]interface{}, 0, len(in))
//	for _, v := range in {
//		s := make(map[string]interface{})
//		if v.Source != nil {
//			s["source"] = flattenSourceTypeIdKey(v.Source)
//		}
//		if v.AttributeName != nil {
//			s["attribute_name"] = *v.AttributeName
//		}
//		if v.Condition != nil {
//			s["condition"] = *v.Condition
//		}
//		if v.Value != nil {
//			s["value"] = *v.Value
//		}
//		if v.ErrorResult != nil {
//			s["error_result"] = *v.ErrorResult
//		}
//		m = append(m, s)
//	}
//	return m
//}

func flattenExpressionIssuanceCriteriaEntry(in *pf.ExpressionIssuanceCriteriaEntry) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Expression != nil {
		s["expression"] = *in.Expression
	}
	if in.ErrorResult != nil {
		s["error_result"] = *in.ErrorResult
	}
	return s
}

//func flattenExpressionIssuanceCriteriaEntry(in []*pf.ExpressionIssuanceCriteriaEntry) []map[string]interface{} {
//	m := make([]map[string]interface{}, 0, len(in))
//	for _, v := range in {
//		s := make(map[string]interface{})
//		if v.Expression != nil {
//			s["expression"] = *v.Expression
//		}
//		if v.ErrorResult != nil {
//			s["error_result"] = *v.ErrorResult
//		}
//		m = append(m, s)
//	}
//	return m
//}

// Takes list of pointers to strings. Expand to an array
// of raw strings and returns a []interface{}
// to keep compatibility w/ schema.NewSetschema.NewSet
func flattenStringList(list []*string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, *v)
	}
	return vs
}

func flattenScopes(in []*pf.ScopeEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		s["description"] = *v.Description
		m = append(m, s)
	}
	return m
}

func flattenScopeGroups(in []*pf.ScopeGroupEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		s["description"] = *v.Description
		s["scopes"] = flattenStringList(*v.Scopes)
		m = append(m, s)
	}
	return m
}

func flattenPersistentGrantContract(in *pf.PersistentGrantContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["extended_attributes"] = flattenPersistentGrantAttributes(*in.ExtendedAttributes)
	s["core_attributes"] = flattenPersistentGrantAttributes(*in.CoreAttributes)
	m = append(m, s)
	return m
}

func flattenPersistentGrantAttributes(in []*pf.PersistentGrantAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func flattenClientAuth(orig, in *pf.ClientAuth) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ClientCertIssuerDn != nil {
		s["client_cert_issuer_dn"] = *in.ClientCertIssuerDn
	}
	if in.ClientCertSubjectDn != nil {
		s["client_cert_subject_dn"] = *in.ClientCertSubjectDn
	}
	if in.EnforceReplayPrevention != nil {
		s["enforce_replay_prevention"] = *in.EnforceReplayPrevention
	}
	if orig.Secret != nil {
		s["secret"] = *orig.Secret
	}
	if in.TokenEndpointAuthSigningAlgorithm != nil {
		s["token_endpoint_auth_signing_algorithm"] = *in.TokenEndpointAuthSigningAlgorithm
	}
	s["type"] = *in.Type
	m = append(m, s)
	return m
}

func flattenJwksSettings(in *pf.JwksSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Jwks != nil {
		s["jwks"] = *in.Jwks
	}
	if in.JwksUrl != nil {
		s["jwks_url"] = *in.JwksUrl
	}
	m = append(m, s)
	return m
}

func flattenResourceLink(in *pf.ResourceLink) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Location != nil {
		s["location"] = *in.Location
	}
	m = append(m, s)
	return m
}

func flattenClientOIDCPolicy(in *pf.ClientOIDCPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.GrantAccessSessionRevocationApi != nil {
		s["grant_access_session_revocation_api"] = *in.GrantAccessSessionRevocationApi
	}
	if in.IdTokenSigningAlgorithm != nil {
		s["id_token_signing_algorithm"] = *in.IdTokenSigningAlgorithm
	}
	if in.LogoutUris != nil && len(*in.LogoutUris) > 0 {
		s["logout_uris"] = flattenStringList(*in.LogoutUris)
	}
	if in.PingAccessLogoutCapable != nil {
		s["ping_access_logout_capable"] = *in.PingAccessLogoutCapable
	}
	if in.PairwiseIdentifierUserType != nil {
		s["pairwise_identifier_user_type"] = *in.PairwiseIdentifierUserType
	}
	if in.PolicyGroup != nil {
		s["policy_group"] = flattenResourceLink(in.PolicyGroup)
	}
	if in.IdTokenEncryptionAlgorithm != nil {
		s["id_token_encryption_algorithm"] = *in.IdTokenEncryptionAlgorithm
	}
	if in.IdTokenContentEncryptionAlgorithm != nil {
		s["id_token_content_encryption_algorithm"] = *in.IdTokenContentEncryptionAlgorithm
	}
	if in.SectorIdentifierUri != nil {
		s["sector_identifier_uri"] = *in.SectorIdentifierUri
	}

	m = append(m, s)
	return m
}

func flattenConfigField(in []*pf.ConfigField) *schema.Set {
	var m []interface{}
	for _, v := range in {
		if v.EncryptedValue != nil {
			continue
		}
		s := make(map[string]interface{})
		s["name"] = *v.Name
		//We check if the Encrypted value is set, if its not we can update the value as a normal password field
		//will not return the value so we need to not overwrite it, which unfortunely means we cannot track password changes
		//this is a limitation of ping federate.
		if v.Value != nil && v.EncryptedValue == nil {
			s["value"] = *v.Value
		}
		// if v.EncryptedValue != nil && *v.EncryptedValue != "" {
		// 	s["encrypted_value"] = *v.EncryptedValue
		// }
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return schema.NewSet(configFieldHash, m)
}

func flattenSensitiveConfigField(in []*pf.ConfigField) *schema.Set {
	var m []interface{}
	for _, v := range in {
		if v.EncryptedValue == nil {
			continue
		}
		s := make(map[string]interface{})
		s["name"] = *v.Name
		//We check if the Encrypted value is set, if its not we can update the value as a normal password field
		//will not return the value so we need to not overwrite it, which unfortunely means we cannot track password changes
		//this is a limitation of ping federate.
		//if v.Value != nil && v.EncryptedValue == nil {
		if v.Value != nil {
			s["value"] = *v.Value
		}
		// if v.EncryptedValue != nil && *v.EncryptedValue != "" {
		// 	s["encrypted_value"] = *v.EncryptedValue
		// }
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return schema.NewSet(configFieldHash, m)
}

func flattenConfigRow(in []*pf.ConfigRow) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["fields"] = flattenConfigField(*v.Fields)
		s["sensitive_fields"] = flattenSensitiveConfigField(*v.Fields)
		m = append(m, s)
	}
	return m
}

func flattenConfigTable(in []*pf.ConfigTable) []interface{} {
	var m []interface{}
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = *v.Name
		if v.Rows != nil {
			s["rows"] = flattenConfigRow(*v.Rows)
		}
		if v.Inherited != nil {
			s["inherited"] = *v.Inherited
		}
		m = append(m, s)
	}
	return m
}

func flattenPluginConfiguration(in *pf.PluginConfiguration) []interface{} {
	s := make(map[string]interface{})
	if in.Tables != nil {
		s["tables"] = flattenConfigTable(*in.Tables)
	}
	if in.Fields != nil {
		s["fields"] = flattenConfigField(*in.Fields)
	}
	if in.Fields != nil {
		s["sensitive_fields"] = flattenSensitiveConfigField(*in.Fields)
	}
	return []interface{}{s}
}

func flattenAccessTokenAttributeContract(in *pf.AccessTokenAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenAccessTokenAttributes(*in.ExtendedAttributes)
	}
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenAccessTokenAttributes(*in.CoreAttributes)
	}
	m = append(m, s)
	return m
}

func flattenAccessTokenAttributes(in []*pf.AccessTokenAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func flattenAuthenticationPolicyContractAttribute(in []*pf.AuthenticationPolicyContractAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func flattenPasswordCredentialValidatorAttribute(in []*pf.PasswordCredentialValidatorAttribute) *schema.Set {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return schema.NewSet(schema.HashString, m)
}

func flattenPasswordCredentialValidatorAttributeContract(in *pf.PasswordCredentialValidatorAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Inherited != nil {
		s["inherited"] = *in.Inherited
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenPasswordCredentialValidatorAttribute(*in.ExtendedAttributes)
	}
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenPasswordCredentialValidatorAttribute(*in.CoreAttributes)
	}
	m = append(m, s)
	return m
}

func flattenJdbcTagConfigs(in *[]*pf.JdbcTagConfig) *schema.Set {
	var m []interface{}
	for _, v := range *in {
		s := make(map[string]interface{})
		if v.ConnectionUrl != nil {
			s["connection_url"] = *v.ConnectionUrl
		}
		if v.Tags != nil {
			s["tags"] = *v.Tags
		}
		if v.DefaultSource != nil {
			s["default_source"] = *v.DefaultSource
		}
		m = append(m, s)
	}
	return schema.NewSet(jdbcTagConfigHash, m)
}

func flattenAuthenticationSelectorAttributeContract(in *pf.AuthenticationSelectorAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["extended_attributes"] = flattenAuthenticationSelectorAttributes(*in.ExtendedAttributes)
	m = append(m, s)
	return m
}

func flattenAuthenticationSelectorAttributes(in []*pf.AuthenticationSelectorAttribute) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, *v.Name)
	}
	return m
}

func flattenAccessTokenMappingContext(in *pf.AccessTokenMappingContext) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["type"] = in.Type
	if in.ContextRef != nil {
		s["context_ref"] = flattenResourceLink(in.ContextRef)
	}
	m = append(m, s)
	return m
}

//func flattenAttributeSources(d *schema.ResourceData, rv *[]*pf.AttributeSource) error {
//	if *rv != nil && len(*rv) > 0 {
//		var ldapAttributes []interface{}
//		var jdbcAttributes []interface{}
//		var customAttributes []interface{}
//		for _, v := range *rv {
//			switch *v.Type {
//			case "LDAP":
//				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(&v.LdapAttributeSource))
//			case "JDBC":
//				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
//			case "CUSTOM":
//				customAttributes = append(customAttributes, flattenCustomAttributeSource(&v.CustomAttributeSource))
//			}
//		}
//		if len(ldapAttributes) > 0 {
//			if err := d.Set("ldap_attribute_source", ldapAttributes); err != nil {
//				return err
//			}
//		}
//		if len(jdbcAttributes) > 0 {
//			if err := d.Set("jdbc_attribute_source", jdbcAttributes); err != nil {
//				return err
//			}
//		}
//		if len(customAttributes) > 0 {
//			if err := d.Set("custom_attribute_source", customAttributes); err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}

func flattenOpenIdConnectAttributes(in []*pf.OpenIdConnectAttribute) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		s["name"] = v.Name
		if v.IncludeInUserInfo != nil {
			s["include_in_user_info"] = v.IncludeInUserInfo
		}
		if v.IncludeInIdToken != nil {
			s["include_in_id_token"] = v.IncludeInIdToken
		}
		m = append(m, s)
	}
	return m
}

func flattenOpenIdConnectAttributeContract(in *pf.OpenIdConnectAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ExtendedAttributes != nil && len(*in.ExtendedAttributes) > 0 {
		s["extended_attributes"] = flattenOpenIdConnectAttributes(*in.ExtendedAttributes)
	}
	if in.CoreAttributes != nil && len(*in.CoreAttributes) > 0 {
		s["core_attributes"] = flattenOpenIdConnectAttributes(*in.CoreAttributes)
	}
	m = append(m, s)
	return m
}

func flattenAttributeMapping(in *pf.AttributeMapping) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}

	if in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	m = append(m, s)
	return m
}

func flattenAuthenticationSources(in []*pf.AuthenticationSource) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, flattenAuthenticationSource(v)...)
	}
	return m
}

func flattenAuthenticationSource(in *pf.AuthenticationSource) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	s["type"] = in.Type
	s["source_ref"] = flattenResourceLink(in.SourceRef)
	m = append(m, s)
	return m
}

func flattenAuthenticationPolicyTrees(in []*pf.AuthenticationPolicyTree) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		s := make(map[string]interface{})
		if v.Name != nil {
			s["name"] = *v.Name
		}
		if v.Description != nil {
			s["description"] = *v.Description
		}
		if v.Enabled != nil {
			s["enabled"] = *v.Enabled
		}
		if v.AuthenticationApiApplicationRef != nil {
			s["authentication_api_application_ref"] = flattenResourceLink(v.AuthenticationApiApplicationRef)
		}
		if v.RootNode != nil {
			s["root_node"] = flattenAuthenticationPolicyTreeNodes([]*pf.AuthenticationPolicyTreeNode{v.RootNode})
		}
		m = append(m, s)
	}
	return m
}

func flattenAuthenticationPolicyTreeNodes(in []*pf.AuthenticationPolicyTreeNode) []interface{} {
	m := make([]interface{}, 0, len(in))
	for _, v := range in {
		m = append(m, flattenAuthenticationPolicyTreeNode(v))
	}
	return m
}

func flattenAuthenticationPolicyTreeNode(in *pf.AuthenticationPolicyTreeNode) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Action != nil {
		s["action"] = flattenPolicyAction(in.Action)
	}
	if in.Children != nil {
		s["children"] = flattenAuthenticationPolicyTreeNodes(*in.Children)
	}
	return s
}

func flattenPolicyAction(in *pf.PolicyAction) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Type != nil {
		s["type"] = *in.Type
	}
	if in.Context != nil {
		s["context"] = *in.Context
	}
	if in.AttributeMapping != nil {
		s["attribute_mapping"] = flattenAttributeMapping(in.AttributeMapping)
	}
	if in.AuthenticationPolicyContractRef != nil {
		s["authentication_policy_contract_ref"] = flattenResourceLink(in.AuthenticationPolicyContractRef)
	}
	if in.AuthenticationSelectorRef != nil {
		s["authentication_selector_ref"] = flattenResourceLink(in.AuthenticationSelectorRef)
	}
	if in.LocalIdentityRef != nil {
		s["local_identity_ref"] = flattenResourceLink(in.LocalIdentityRef)
	}
	if in.InboundMapping != nil {
		s["inbound_mapping"] = flattenAttributeMapping(in.InboundMapping)
	}
	if in.OutboundAttributeMapping != nil {
		s["outbound_attribute_mapping"] = flattenAttributeMapping(in.OutboundAttributeMapping)
	}
	if in.AuthenticationSource != nil {
		s["authentication_source"] = flattenAuthenticationSource(in.AuthenticationSource)
	}
	if in.InputUserIdMapping != nil {
		m := make([]interface{}, 0, 1)
		s["input_user_id_mapping"] = append(m, flattenAttributeFulfillmentValue(in.InputUserIdMapping))
	}
	if in.AttributeRules != nil {
		s["attribute_rules"] = flattenAttributeRules(in.AttributeRules)
	}
	return append(m, s)
}

func flattenAttributeRules(in *pf.AttributeRules) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.FallbackToSuccess != nil {
		s["fallback_to_success"] = *in.FallbackToSuccess
	}
	if in.Items != nil {
		s["items"] = flattenAttributeRuleSlice(in.Items)
	}
	return append(m, s)
}

func flattenAttributeRuleSlice(in *[]*pf.AttributeRule) *schema.Set {
	m := make([]interface{}, 0, len(*in))
	for _, r := range *in {
		s := make(map[string]interface{})
		if r.Condition != nil {
			s["condition"] = *r.Condition
		}
		if r.AttributeName != nil {
			s["attribute_name"] = *r.AttributeName
		}
		if r.ExpectedValue != nil {
			s["expected_value"] = *r.ExpectedValue
		}
		if r.Result != nil {
			s["result"] = *r.Result
		}
		m = append(m, s)
	}
	return schema.NewSet(attributeRuleSliceHash, m)
}

func flattenMapOfParameterValues(in map[string]*pf.ParameterValues) *schema.Set {
	m := make([]interface{}, 0, len(in))
	for s2 := range in {
		s := flattenParameterValues(in[s2])
		s["key_name"] = s2
		m = append(m, s)
	}
	return schema.NewSet(scopeAttributeMappingsHash, m)
}

func flattenParameterValues(in *pf.ParameterValues) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Values != nil {
		s["values"] = *in.Values
	}
	return s
}

func flattenRolesAndProtocols(in *pf.RolesAndProtocols) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.EnableIdpDiscovery != nil {
		s["enable_idp_discovery"] = *in.EnableIdpDiscovery
	}
	if in.OauthRole != nil {
		s["oauth_role"] = flattenOauthRole(in.OauthRole)
	}
	if in.IdpRole != nil {
		s["idp_role"] = flattenIdpRole(in.IdpRole)
	}
	if in.SpRole != nil {
		s["sp_role"] = flattenSpRole(in.SpRole)
	}
	m = append(m, s)
	return m
}

func flattenOauthRole(in *pf.OAuthRole) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.EnableOauth != nil {
		s["enable_oauth"] = *in.EnableOauth
	}
	if in.EnableOpenIdConnect != nil {
		s["enable_openid_connect"] = *in.EnableOpenIdConnect
	}
	m = append(m, s)
	return m
}

func flattenIdpRole(in *pf.IdpRole) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Enable != nil {
		s["enable"] = *in.Enable
	}
	if in.Saml20Profile != nil {
		s["saml20_profile"] = flattenSaml20Profile(in.Saml20Profile)
	}
	if in.EnableOutboundProvisioning != nil {
		s["enable_outbound_provisioning"] = *in.EnableOutboundProvisioning
	}
	if in.EnableSaml11 != nil {
		s["enable_saml11"] = *in.EnableSaml11
	}
	if in.EnableSaml10 != nil {
		s["enable_saml10"] = *in.EnableSaml10
	}
	if in.EnableWsFed != nil {
		s["enable_ws_fed"] = *in.EnableWsFed
	}
	if in.EnableWsTrust != nil {
		s["enable_ws_trust"] = *in.EnableWsTrust
	}
	m = append(m, s)
	return m
}

func flattenSpRole(in *pf.SpRole) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Enable != nil {
		s["enable"] = *in.Enable
	}
	if in.Saml20Profile != nil {
		s["saml20_profile"] = flattenSpSAML20Profile(in.Saml20Profile)
	}
	if in.EnableInboundProvisioning != nil {
		s["enable_inbound_provisioning"] = *in.EnableInboundProvisioning
	}
	if in.EnableSaml11 != nil {
		s["enable_saml11"] = *in.EnableSaml11
	}
	if in.EnableSaml10 != nil {
		s["enable_saml10"] = *in.EnableSaml10
	}
	if in.EnableWsFed != nil {
		s["enable_ws_fed"] = *in.EnableWsFed
	}
	if in.EnableWsTrust != nil {
		s["enable_ws_trust"] = *in.EnableWsTrust
	}
	if in.EnableOpenIDConnect != nil {
		s["enable_openid_connect"] = *in.EnableOpenIDConnect
	}
	m = append(m, s)
	return m
}

func flattenSaml20Profile(in *pf.SAML20Profile) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Enable != nil {
		s["enable"] = *in.Enable
	}
	if in.EnableAutoConnect != nil {
		s["enable_auto_connect"] = *in.EnableAutoConnect
	}
	m = append(m, s)
	return m
}

func flattenSpSAML20Profile(in *pf.SpSAML20Profile) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Enable != nil {
		s["enable"] = *in.Enable
	}
	if in.EnableAutoConnect != nil {
		s["enable_auto_connect"] = *in.EnableAutoConnect
	}
	if in.EnableXASP != nil {
		s["enable_xasp"] = *in.EnableXASP
	}
	m = append(m, s)
	return m
}

func flattenFederationInfo(in *pf.FederationInfo) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.BaseUrl != nil {
		s["base_url"] = *in.BaseUrl
	}
	if in.Saml2EntityId != nil {
		s["saml2_entity_id"] = *in.Saml2EntityId
	}
	if in.Saml1xIssuerId != nil {
		s["saml1x_issuer_id"] = *in.Saml1xIssuerId
	}
	if in.Saml1xSourceId != nil {
		s["saml1x_source_id"] = *in.Saml1xSourceId
	}
	if in.WsfedRealm != nil {
		s["wsfed_realm"] = *in.WsfedRealm
	}
	m = append(m, s)
	return m
}

func flattenIdpWsTrustAttribute(in *pf.IdpWsTrustAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.Masked != nil {
		s["masked"] = *in.Masked
	}
	return s
}

func flattenSpWsTrustAttributeContract(in *pf.SpWsTrustAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenSpWsTrustAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenSpWsTrustAttributes(in.ExtendedAttributes)
	}
	return append(m, s)
}
func flattenIdpAttributeQuery(in *pf.IdpAttributeQuery) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Url != nil {
		s["url"] = *in.Url
	}
	if in.NameMappings != nil {
		s["name_mappings"] = flattenAttributeQueryNameMappings(in.NameMappings)
	}
	if in.Policy != nil {
		s["policy"] = flattenIdpAttributeQueryPolicy(in.Policy)
	}
	return append(m, s)
}

func flattenChannelSourceLocation(in *pf.ChannelSourceLocation) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.GroupDN != nil {
		s["group_dn"] = *in.GroupDN
	}
	if in.Filter != nil {
		s["filter"] = *in.Filter
	}
	if in.NestedSearch != nil {
		s["nested_search"] = *in.NestedSearch
	}
	return append(m, s)
}
func flattenIdpAdapterAssertionMapping(in *pf.IdpAdapterAssertionMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.IdpAdapterRef != nil {
		s["idp_adapter_ref"] = flattenResourceLink(in.IdpAdapterRef)
	}
	if in.RestrictVirtualEntityIds != nil {
		s["restrict_virtual_entity_ids"] = *in.RestrictVirtualEntityIds
	}
	if in.RestrictedVirtualEntityIds != nil {
		s["restricted_virtual_entity_ids"] = *in.RestrictedVirtualEntityIds
	}
	if in.AdapterOverrideSettings != nil {
		s["adapter_override_settings"] = flattenIdpAdapter(in.AdapterOverrideSettings)
	}
	if in.AbortSsoTransactionAsFailSafe != nil {
		s["abort_sso_transaction_as_fail_safe"] = *in.AbortSsoTransactionAsFailSafe
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	return s
}

func flattenOIDCRequestParameter(in *pf.OIDCRequestParameter) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.Value != nil {
		s["value"] = *in.Value
	}
	if in.ApplicationEndpointOverride != nil {
		s["application_endpoint_override"] = *in.ApplicationEndpointOverride
	}
	return s
}
func flattenConnectionCredentials(in *pf.ConnectionCredentials) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.OutboundBackChannelAuth != nil {
		s["outbound_back_channel_auth"] = flattenOutboundBackChannelAuth(in.OutboundBackChannelAuth)
	}
	if in.VerificationSubjectDN != nil {
		s["verification_subject_dn"] = *in.VerificationSubjectDN
	}
	if in.KeyTransportAlgorithm != nil {
		s["key_transport_algorithm"] = *in.KeyTransportAlgorithm
	}
	if in.DecryptionKeyPairRef != nil {
		s["decryption_key_pair_ref"] = flattenResourceLink(in.DecryptionKeyPairRef)
	}
	if in.SigningSettings != nil {
		s["signing_settings"] = flattenSigningSettings(in.SigningSettings)
	}
	if in.SecondaryDecryptionKeyPairRef != nil {
		s["secondary_decryption_key_pair_ref"] = flattenResourceLink(in.SecondaryDecryptionKeyPairRef)
	}
	if in.InboundBackChannelAuth != nil {
		s["inbound_back_channel_auth"] = flattenInboundBackChannelAuth(in.InboundBackChannelAuth)
	}
	if in.VerificationIssuerDN != nil {
		s["verification_issuer_dn"] = *in.VerificationIssuerDN
	}
	if in.Certs != nil {
		s["certs"] = flattenConnectionCerts(in.Certs)
	}
	if in.BlockEncryptionAlgorithm != nil {
		s["block_encryption_algorithm"] = *in.BlockEncryptionAlgorithm
	}
	return append(m, s)
}
func flattenSpBrowserSso(in *pf.SpBrowserSso) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.SpSamlIdentityMapping != nil {
		s["sp_saml_identity_mapping"] = *in.SpSamlIdentityMapping
	}
	if in.RequireSignedAuthnRequests != nil {
		s["require_signed_authn_requests"] = *in.RequireSignedAuthnRequests
	}
	if in.MessageCustomizations != nil {
		s["message_customizations"] = flattenProtocolMessageCustomizations(in.MessageCustomizations)
	}
	if in.AssertionLifetime != nil {
		s["assertion_lifetime"] = flattenAssertionLifetime(in.AssertionLifetime)
	}
	if in.UrlWhitelistEntries != nil {
		s["url_whitelist_entries"] = flattenUrlWhitelistEntrys(in.UrlWhitelistEntries)
	}
	if in.SloServiceEndpoints != nil {
		s["slo_service_endpoints"] = flattenSloServiceEndpoints(in.SloServiceEndpoints)
	}
	if in.SpWsFedIdentityMapping != nil {
		s["sp_ws_fed_identity_mapping"] = *in.SpWsFedIdentityMapping
	}
	if in.SignResponseAsRequired != nil {
		s["sign_response_as_required"] = *in.SignResponseAsRequired
	}
	if in.SignAssertions != nil {
		s["sign_assertions"] = *in.SignAssertions
	}
	if in.AdapterMappings != nil {
		s["adapter_mappings"] = flattenIdpAdapterAssertionMappings(in.AdapterMappings)
	}
	if in.EnabledProfiles != nil {
		s["enabled_profiles"] = *in.EnabledProfiles
	}
	if in.DefaultTargetUrl != nil {
		s["default_target_url"] = *in.DefaultTargetUrl
	}
	if in.SsoServiceEndpoints != nil {
		s["sso_service_endpoints"] = flattenSpSsoServiceEndpoints(in.SsoServiceEndpoints)
	}
	if in.WsTrustVersion != nil {
		s["ws_trust_version"] = *in.WsTrustVersion
	}
	if in.IncomingBindings != nil {
		s["incoming_bindings"] = *in.IncomingBindings
	}
	if in.Artifact != nil {
		s["artifact"] = flattenArtifactSettings(in.Artifact)
	}
	if in.EncryptionPolicy != nil {
		s["encryption_policy"] = flattenEncryptionPolicy(in.EncryptionPolicy)
	}
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenSpBrowserSsoAttributeContract(in.AttributeContract)
	}
	if in.AuthenticationPolicyContractAssertionMappings != nil {
		s["authentication_policy_contract_assertion_mappings"] = flattenAuthenticationPolicyContractAssertionMappings(in.AuthenticationPolicyContractAssertionMappings)
	}
	if in.Protocol != nil {
		s["protocol"] = *in.Protocol
	}
	if in.WsFedTokenType != nil {
		s["ws_fed_token_type"] = *in.WsFedTokenType
	}
	return append(m, s)
}
func flattenAttributeQueryNameMapping(in *pf.AttributeQueryNameMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.LocalName != nil {
		s["local_name"] = *in.LocalName
	}
	if in.RemoteName != nil {
		s["remote_name"] = *in.RemoteName
	}
	return s
}
func flattenSchema(in *pf.Schema) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Namespace != nil {
		s["namespace"] = *in.Namespace
	}
	if in.Attributes != nil {
		s["attributes"] = flattenSchemaAttributes(in.Attributes)
	}
	return append(m, s)
}
func flattenOutboundProvision(in *pf.OutboundProvision) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Type != nil {
		s["type"] = *in.Type
	}
	if in.TargetSettings != nil {
		s["target_settings"] = flattenConfigField(*in.TargetSettings)
		s["sensitive_target_settings"] = flattenSensitiveConfigField(*in.TargetSettings)
	}
	if in.CustomSchema != nil {
		s["custom_schema"] = flattenSchema(in.CustomSchema)
	}
	if in.Channels != nil {
		s["channels"] = flattenChannels(in.Channels)
	}
	return append(m, s)
}
func flattenX509File(in *pf.X509File) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.FileData != nil {
		s["file_data"] = *in.FileData
	}
	if in.CryptoProvider != nil {
		s["crypto_provider"] = *in.CryptoProvider
	}
	return append(m, s)
}
func flattenUrlWhitelistEntry(in *pf.UrlWhitelistEntry) map[string]interface{} {
	s := make(map[string]interface{})
	if in.ValidDomain != nil {
		s["valid_domain"] = *in.ValidDomain
	}
	if in.ValidPath != nil {
		s["valid_path"] = *in.ValidPath
	}
	if in.AllowQueryAndFragment != nil {
		s["allow_query_and_fragment"] = *in.AllowQueryAndFragment
	}
	if in.RequireHttps != nil {
		s["require_https"] = *in.RequireHttps
	}
	return s
}
func flattenIdpOAuthGrantAttributeMapping(in *pf.IdpOAuthGrantAttributeMapping) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AccessTokenManagerMappings != nil {
		s["access_token_manager_mappings"] = flattenAccessTokenManagerMappings(in.AccessTokenManagerMappings)
	}
	if in.IdpOAuthAttributeContract != nil {
		s["idp_o_auth_attribute_contract"] = flattenIdpOAuthAttributeContract(in.IdpOAuthAttributeContract)
	}
	return append(m, s)
}

func flattenIdpSsoServiceEndpoint(in *pf.IdpSsoServiceEndpoint) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Binding != nil {
		s["binding"] = *in.Binding
	}
	if in.Url != nil {
		s["url"] = *in.Url
	}
	return s
}

func flattenSloServiceEndpoint(in *pf.SloServiceEndpoint) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Binding != nil {
		s["binding"] = *in.Binding
	}
	if in.Url != nil {
		s["url"] = *in.Url
	}
	if in.ResponseUrl != nil {
		s["response_url"] = *in.ResponseUrl
	}
	return s
}
func flattenSaasFieldConfiguration(in *pf.SaasFieldConfiguration) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Expression != nil {
		s["expression"] = *in.Expression
	}
	if in.CreateOnly != nil {
		s["create_only"] = *in.CreateOnly
	}
	if in.Trim != nil {
		s["trim"] = *in.Trim
	}
	if in.CharacterCase != nil {
		s["character_case"] = *in.CharacterCase
	}
	if in.Parser != nil {
		s["parser"] = *in.Parser
	}
	if in.Masked != nil {
		s["masked"] = *in.Masked
	}
	if in.AttributeNames != nil {
		s["attribute_names"] = *in.AttributeNames
	}
	if in.DefaultValue != nil {
		s["default_value"] = *in.DefaultValue
	}
	return append(m, s)
}
func flattenSpBrowserSsoAttribute(in *pf.SpBrowserSsoAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.NameFormat != nil {
		s["name_format"] = *in.NameFormat
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	return s
}
func flattenSchemaAttribute(in *pf.SchemaAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Types != nil {
		s["types"] = *in.Types
	}
	if in.SubAttributes != nil {
		s["sub_attributes"] = *in.SubAttributes
	}
	if in.MultiValued != nil {
		s["multi_valued"] = *in.MultiValued
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	return s
}
func flattenIdpBrowserSso(in *pf.IdpBrowserSso) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.IncomingBindings != nil {
		s["incoming_bindings"] = *in.IncomingBindings
	}
	if in.UrlWhitelistEntries != nil {
		s["url_whitelist_entries"] = flattenUrlWhitelistEntrys(in.UrlWhitelistEntries)
	}
	if in.SloServiceEndpoints != nil {
		s["slo_service_endpoints"] = flattenSloServiceEndpoints(in.SloServiceEndpoints)
	}
	if in.DecryptionPolicy != nil {
		s["decryption_policy"] = flattenDecryptionPolicy(in.DecryptionPolicy)
	}
	if in.AdapterMappings != nil {
		s["adapter_mappings"] = flattenSpAdapterMappings(in.AdapterMappings)
	}
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenIdpBrowserSsoAttributeContract(in.AttributeContract)
	}
	if in.SsoOAuthMapping != nil {
		s["sso_o_auth_mapping"] = flattenSsoOAuthMapping(in.SsoOAuthMapping)
	}
	if in.Protocol != nil {
		s["protocol"] = *in.Protocol
	}
	if in.Artifact != nil {
		s["artifact"] = flattenArtifactSettings(in.Artifact)
	}
	if in.DefaultTargetUrl != nil {
		s["default_target_url"] = *in.DefaultTargetUrl
	}
	if in.AssertionsSigned != nil {
		s["assertions_signed"] = *in.AssertionsSigned
	}
	if in.SignAuthnRequests != nil {
		s["sign_authn_requests"] = *in.SignAuthnRequests
	}
	if in.EnabledProfiles != nil {
		s["enabled_profiles"] = *in.EnabledProfiles
	}
	if in.SsoServiceEndpoints != nil {
		s["sso_service_endpoints"] = flattenIdpSsoServiceEndpoints(in.SsoServiceEndpoints)
	}
	if in.AuthnContextMappings != nil {
		s["authn_context_mappings"] = flattenAuthnContextMappings(in.AuthnContextMappings)
	}
	if in.IdpIdentityMapping != nil {
		s["idp_identity_mapping"] = *in.IdpIdentityMapping
	}
	if in.AuthenticationPolicyContractMappings != nil {
		s["authentication_policy_contract_mappings"] = flattenAuthenticationPolicyContractMappings(in.AuthenticationPolicyContractMappings)
	}
	if in.OidcProviderSettings != nil {
		s["oidc_provider_settings"] = flattenOIDCProviderSettings(in.OidcProviderSettings)
	}
	if in.MessageCustomizations != nil {
		s["message_customizations"] = flattenProtocolMessageCustomizations(in.MessageCustomizations)
	}
	if in.OauthAuthenticationPolicyContractRef != nil {
		s["oauth_authentication_policy_contract_ref"] = flattenResourceLink(in.OauthAuthenticationPolicyContractRef)
	}
	return append(m, s)
}
func flattenSpAdapter(in *pf.SpAdapter) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.PluginDescriptorRef != nil {
		s["plugin_descriptor_ref"] = flattenResourceLink(in.PluginDescriptorRef)
	}
	if in.ParentRef != nil {
		s["parent_ref"] = flattenResourceLink(in.ParentRef)
	}
	if in.Configuration != nil {
		s["configuration"] = flattenPluginConfiguration(in.Configuration)
	}
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenSpAdapterAttributeContract(in.AttributeContract)
	}
	if in.TargetApplicationInfo != nil {
		s["target_application_info"] = flattenSpAdapterTargetApplicationInfo(in.TargetApplicationInfo)
	}
	return append(m, s)
}
func flattenChannelSource(in *pf.ChannelSource) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.BaseDn != nil {
		s["base_dn"] = *in.BaseDn
	}
	if in.UserSourceLocation != nil {
		s["user_source_location"] = flattenChannelSourceLocation(in.UserSourceLocation)
	}
	if in.DataSource != nil {
		s["data_source"] = flattenResourceLink(in.DataSource)
	}
	if in.GuidAttributeName != nil {
		s["guid_attribute_name"] = *in.GuidAttributeName
	}
	if in.GuidBinary != nil {
		s["guid_binary"] = *in.GuidBinary
	}
	if in.GroupMembershipDetection != nil {
		s["group_membership_detection"] = flattenGroupMembershipDetection(in.GroupMembershipDetection)
	}
	if in.ChangeDetectionSettings != nil {
		s["change_detection_settings"] = flattenChangeDetectionSettings(in.ChangeDetectionSettings)
	}
	if in.AccountManagementSettings != nil {
		s["account_management_settings"] = flattenAccountManagementSettings(in.AccountManagementSettings)
	}
	if in.GroupSourceLocation != nil {
		s["group_source_location"] = flattenChannelSourceLocation(in.GroupSourceLocation)
	}
	return append(m, s)
}
func flattenSpTokenGeneratorMapping(in *pf.SpTokenGeneratorMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.SpTokenGeneratorRef != nil {
		s["sp_token_generator_ref"] = flattenResourceLink(in.SpTokenGeneratorRef)
	}
	if in.RestrictedVirtualEntityIds != nil {
		s["restricted_virtual_entity_ids"] = *in.RestrictedVirtualEntityIds
	}
	if in.DefaultMapping != nil {
		s["default_mapping"] = *in.DefaultMapping
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	return s
}

func flattenIdpAttributeQueryPolicy(in *pf.IdpAttributeQueryPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.RequireSignedResponse != nil {
		s["require_signed_response"] = *in.RequireSignedResponse
	}
	if in.RequireSignedAssertion != nil {
		s["require_signed_assertion"] = *in.RequireSignedAssertion
	}
	if in.RequireEncryptedAssertion != nil {
		s["require_encrypted_assertion"] = *in.RequireEncryptedAssertion
	}
	if in.SignAttributeQuery != nil {
		s["sign_attribute_query"] = *in.SignAttributeQuery
	}
	if in.EncryptNameId != nil {
		s["encrypt_name_id"] = *in.EncryptNameId
	}
	if in.MaskAttributeValues != nil {
		s["mask_attribute_values"] = *in.MaskAttributeValues
	}
	return append(m, s)
}
func flattenConnectionCert(in *pf.ConnectionCert) map[string]interface{} {
	s := make(map[string]interface{})
	if in.CertView != nil {
		s["cert_view"] = flattenCertView(in.CertView)
	}
	if in.X509File != nil {
		s["x509_file"] = flattenX509File(in.X509File)
	}
	if in.ActiveVerificationCert != nil {
		s["active_verification_cert"] = *in.ActiveVerificationCert
	}
	if in.PrimaryVerificationCert != nil {
		s["primary_verification_cert"] = *in.PrimaryVerificationCert
	}
	if in.SecondaryVerificationCert != nil {
		s["secondary_verification_cert"] = *in.SecondaryVerificationCert
	}
	if in.EncryptionCert != nil {
		s["encryption_cert"] = *in.EncryptionCert
	}
	return s
}
func flattenSpAttributeQueryPolicy(in *pf.SpAttributeQueryPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.RequireSignedAttributeQuery != nil {
		s["require_signed_attribute_query"] = *in.RequireSignedAttributeQuery
	}
	if in.RequireEncryptedNameId != nil {
		s["require_encrypted_name_id"] = *in.RequireEncryptedNameId
	}
	if in.SignResponse != nil {
		s["sign_response"] = *in.SignResponse
	}
	if in.SignAssertion != nil {
		s["sign_assertion"] = *in.SignAssertion
	}
	if in.EncryptAssertion != nil {
		s["encrypt_assertion"] = *in.EncryptAssertion
	}
	return append(m, s)
}

func flattenIdpWsTrustAttributeContract(in *pf.IdpWsTrustAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenIdpWsTrustAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenIdpWsTrustAttributes(in.ExtendedAttributes)
	}
	return append(m, s)
}
func flattenArtifactSettings(in *pf.ArtifactSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Lifetime != nil {
		s["lifetime"] = *in.Lifetime
	}
	if in.ResolverLocations != nil {
		s["resolver_locations"] = flattenArtifactResolverLocations(in.ResolverLocations)
	}
	if in.SourceId != nil {
		s["source_id"] = *in.SourceId
	}
	return append(m, s)
}
func flattenContactInfo(in *pf.ContactInfo) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Company != nil {
		s["company"] = *in.Company
	}
	if in.Email != nil {
		s["email"] = *in.Email
	}
	if in.FirstName != nil {
		s["first_name"] = *in.FirstName
	}
	if in.LastName != nil {
		s["last_name"] = *in.LastName
	}
	if in.Phone != nil {
		s["phone"] = *in.Phone
	}
	return append(m, s)
}

func flattenChangeDetectionSettings(in *pf.ChangeDetectionSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.UserObjectClass != nil {
		s["user_object_class"] = *in.UserObjectClass
	}
	if in.GroupObjectClass != nil {
		s["group_object_class"] = *in.GroupObjectClass
	}
	if in.ChangedUsersAlgorithm != nil {
		s["changed_users_algorithm"] = *in.ChangedUsersAlgorithm
	}
	if in.UsnAttributeName != nil {
		s["usn_attribute_name"] = *in.UsnAttributeName
	}
	if in.TimeStampAttributeName != nil {
		s["time_stamp_attribute_name"] = *in.TimeStampAttributeName
	}
	return append(m, s)
}
func flattenSsoOAuthMapping(in *pf.SsoOAuthMapping) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	return append(m, s)
}
func flattenIdpAdapterAttributeContract(in *pf.IdpAdapterAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Inherited != nil {
		s["inherited"] = *in.Inherited
	}
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenIdpAdapterAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenIdpAdapterAttributes(in.ExtendedAttributes)
	}
	if in.MaskOgnlValues != nil {
		s["mask_ognl_values"] = *in.MaskOgnlValues
	}
	return append(m, s)
}
func flattenSpAdapterMapping(in *pf.SpAdapterMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.SpAdapterRef != nil {
		s["sp_adapter_ref"] = flattenResourceLink(in.SpAdapterRef)
	}
	if in.RestrictVirtualEntityIds != nil {
		s["restrict_virtual_entity_ids"] = *in.RestrictVirtualEntityIds
	}
	if in.RestrictedVirtualEntityIds != nil {
		s["restricted_virtual_entity_ids"] = *in.RestrictedVirtualEntityIds
	}
	if in.AdapterOverrideSettings != nil {
		s["adapter_override_settings"] = flattenSpAdapter(in.AdapterOverrideSettings)
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	return s
}
func flattenOIDCClientCredentials(in *pf.OIDCClientCredentials) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.ClientSecret != nil {
		s["client_secret"] = *in.ClientSecret
	}
	if in.EncryptedSecret != nil {
		s["encrypted_secret"] = *in.EncryptedSecret
	}
	if in.ClientId != nil {
		s["client_id"] = *in.ClientId
	}
	return append(m, s)
}
func flattenUsernamePasswordCredentials(in *pf.UsernamePasswordCredentials) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Password != nil {
		s["password"] = *in.Password
	}
	if in.EncryptedPassword != nil {
		s["encrypted_password"] = *in.EncryptedPassword
	}
	if in.Username != nil {
		s["username"] = *in.Username
	}
	return append(m, s)
}
func flattenConnectionMetadataUrl(in *pf.ConnectionMetadataUrl) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.MetadataUrlRef != nil {
		s["metadata_url_ref"] = flattenResourceLink(in.MetadataUrlRef)
	}
	if in.EnableAutoMetadataUpdate != nil {
		s["enable_auto_metadata_update"] = *in.EnableAutoMetadataUpdate
	}
	return append(m, s)
}
func flattenEncryptionPolicy(in *pf.EncryptionPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.EncryptSloSubjectNameId != nil {
		s["encrypt_slo_subject_name_id"] = *in.EncryptSloSubjectNameId
	}
	if in.SloSubjectNameIDEncrypted != nil {
		s["slo_subject_name_id_encrypted"] = *in.SloSubjectNameIDEncrypted
	}
	if in.EncryptAssertion != nil {
		s["encrypt_assertion"] = *in.EncryptAssertion
	}
	if in.EncryptedAttributes != nil {
		s["encrypted_attributes"] = *in.EncryptedAttributes
	}
	return append(m, s)
}

func flattenIdpWsTrust(in *pf.IdpWsTrust) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenIdpWsTrustAttributeContract(in.AttributeContract)
	}
	if in.GenerateLocalToken != nil {
		s["generate_local_token"] = *in.GenerateLocalToken
	}
	if in.TokenGeneratorMappings != nil {
		s["token_generator_mappings"] = flattenSpTokenGeneratorMappings(in.TokenGeneratorMappings)
	}
	return append(m, s)
}
func flattenSpSsoServiceEndpoint(in *pf.SpSsoServiceEndpoint) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Binding != nil {
		s["binding"] = *in.Binding
	}
	if in.Url != nil {
		s["url"] = *in.Url
	}
	if in.IsDefault != nil {
		s["is_default"] = *in.IsDefault
	}
	if in.Index != nil {
		s["index"] = *in.Index
	}
	return s
}

func flattenInboundBackChannelAuth(in *pf.InboundBackChannelAuth) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.RequireSsl != nil {
		s["require_ssl"] = *in.RequireSsl
	}
	if in.Type != nil {
		s["type"] = *in.Type
	}
	if in.HttpBasicCredentials != nil {
		s["http_basic_credentials"] = flattenUsernamePasswordCredentials(in.HttpBasicCredentials)
	}
	if in.DigitalSignature != nil {
		s["digital_signature"] = *in.DigitalSignature
	}
	if in.VerificationSubjectDN != nil {
		s["verification_subject_dn"] = *in.VerificationSubjectDN
	}
	if in.VerificationIssuerDN != nil {
		s["verification_issuer_dn"] = *in.VerificationIssuerDN
	}
	if in.Certs != nil {
		s["certs"] = flattenConnectionCerts(in.Certs)
	}
	return append(m, s)
}

func flattenCertView(in *pf.CertView) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.SubjectAlternativeNames != nil {
		s["subject_alternative_names"] = *in.SubjectAlternativeNames
	}
	if in.ValidFrom != nil {
		s["valid_from"] = *in.ValidFrom
	}
	if in.KeyAlgorithm != nil {
		s["key_algorithm"] = *in.KeyAlgorithm
	}
	if in.SignatureAlgorithm != nil {
		s["signature_algorithm"] = *in.SignatureAlgorithm
	}
	if in.Sha256Fingerprint != nil {
		s["sha256_fingerprint"] = *in.Sha256Fingerprint
	}
	if in.SubjectDN != nil {
		s["subject_dn"] = *in.SubjectDN
	}
	if in.CryptoProvider != nil {
		s["crypto_provider"] = *in.CryptoProvider
	}
	if in.IssuerDN != nil {
		s["issuer_dn"] = *in.IssuerDN
	}
	if in.Expires != nil {
		s["expires"] = *in.Expires
	}
	if in.KeySize != nil {
		s["key_size"] = *in.KeySize
	}
	if in.Sha1Fingerprint != nil {
		s["sha1_fingerprint"] = *in.Sha1Fingerprint
	}
	if in.Id != nil {
		s["id"] = *in.Id
	}
	if in.Version != nil {
		s["version"] = *in.Version
	}
	if in.Status != nil {
		s["status"] = *in.Status
	}
	if in.SerialNumber != nil {
		s["serial_number"] = *in.SerialNumber
	}
	return append(m, s)
}
func flattenAuthenticationPolicyContractAssertionMapping(in *pf.AuthenticationPolicyContractAssertionMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.AuthenticationPolicyContractRef != nil {
		s["authentication_policy_contract_ref"] = flattenResourceLink(in.AuthenticationPolicyContractRef)
	}
	if in.RestrictVirtualEntityIds != nil {
		s["restrict_virtual_entity_ids"] = *in.RestrictVirtualEntityIds
	}
	if in.RestrictedVirtualEntityIds != nil {
		s["restricted_virtual_entity_ids"] = *in.RestrictedVirtualEntityIds
	}
	if in.AbortSsoTransactionAsFailSafe != nil {
		s["abort_sso_transaction_as_fail_safe"] = *in.AbortSsoTransactionAsFailSafe
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	return s
}
func flattenDecryptionPolicy(in *pf.DecryptionPolicy) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.SubjectNameIdEncrypted != nil {
		s["subject_name_id_encrypted"] = *in.SubjectNameIdEncrypted
	}
	if in.SloEncryptSubjectNameID != nil {
		s["slo_encrypt_subject_name_id"] = *in.SloEncryptSubjectNameID
	}
	if in.SloSubjectNameIDEncrypted != nil {
		s["slo_subject_name_id_encrypted"] = *in.SloSubjectNameIDEncrypted
	}
	if in.AssertionEncrypted != nil {
		s["assertion_encrypted"] = *in.AssertionEncrypted
	}
	if in.AttributesEncrypted != nil {
		s["attributes_encrypted"] = *in.AttributesEncrypted
	}
	return append(m, s)
}
func flattenIdpAdapter(in *pf.IdpAdapter) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.PluginDescriptorRef != nil {
		s["plugin_descriptor_ref"] = flattenResourceLink(in.PluginDescriptorRef)
	}
	if in.ParentRef != nil {
		s["parent_ref"] = flattenResourceLink(in.ParentRef)
	}
	if in.Configuration != nil {
		s["configuration"] = flattenPluginConfiguration(in.Configuration)
	}
	if in.AuthnCtxClassRef != nil {
		s["authn_ctx_class_ref"] = *in.AuthnCtxClassRef
	}
	if in.AttributeMapping != nil {
		s["attribute_mapping"] = flattenIdpAdapterContractMapping(in.AttributeMapping)
	}
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenIdpAdapterAttributeContract(in.AttributeContract)
	}
	if in.Id != nil {
		s["id"] = *in.Id
	}
	return append(m, s)
}
func flattenAssertionLifetime(in *pf.AssertionLifetime) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.MinutesBefore != nil {
		s["minutes_before"] = *in.MinutesBefore
	}
	if in.MinutesAfter != nil {
		s["minutes_after"] = *in.MinutesAfter
	}
	return append(m, s)
}
func flattenGroupMembershipDetection(in *pf.GroupMembershipDetection) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.MemberOfGroupAttributeName != nil {
		s["member_of_group_attribute_name"] = *in.MemberOfGroupAttributeName
	}
	if in.GroupMemberAttributeName != nil {
		s["group_member_attribute_name"] = *in.GroupMemberAttributeName
	}
	return append(m, s)
}
func flattenChannel(in *pf.Channel) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.MaxThreads != nil {
		s["max_threads"] = *in.MaxThreads
	}
	if in.Timeout != nil {
		s["timeout"] = *in.Timeout
	}
	if in.Active != nil {
		s["active"] = *in.Active
	}
	if in.ChannelSource != nil {
		s["channel_source"] = flattenChannelSource(in.ChannelSource)
	}
	if in.AttributeMapping != nil {
		s["attribute_mapping"] = flattenSaasAttributeMappings(in.AttributeMapping)
	}
	return s
}
func flattenIdpBrowserSsoAttribute(in *pf.IdpBrowserSsoAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Masked != nil {
		s["masked"] = *in.Masked
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	return s
}
func flattenAccountManagementSettings(in *pf.AccountManagementSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AccountStatusAlgorithm != nil {
		s["account_status_algorithm"] = *in.AccountStatusAlgorithm
	}
	if in.FlagComparisonValue != nil {
		s["flag_comparison_value"] = *in.FlagComparisonValue
	}
	if in.FlagComparisonStatus != nil {
		s["flag_comparison_status"] = *in.FlagComparisonStatus
	}
	if in.DefaultStatus != nil {
		s["default_status"] = *in.DefaultStatus
	}
	if in.AccountStatusAttributeName != nil {
		s["account_status_attribute_name"] = *in.AccountStatusAttributeName
	}
	return append(m, s)
}
func flattenIdpAdapterAttribute(in *pf.IdpAdapterAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Name != nil {
		s["name"] = *in.Name
	}
	if in.Pseudonym != nil {
		s["pseudonym"] = *in.Pseudonym
	}
	if in.Masked != nil {
		s["masked"] = *in.Masked
	}
	return s
}

func flattenSpWsTrust(in *pf.SpWsTrust) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.RequestContractRef != nil {
		s["request_contract_ref"] = flattenResourceLink(in.RequestContractRef)
	}
	if in.MessageCustomizations != nil {
		s["message_customizations"] = flattenProtocolMessageCustomizations(in.MessageCustomizations)
	}
	if in.PartnerServiceIds != nil {
		s["partner_service_ids"] = *in.PartnerServiceIds
	}
	if in.MinutesBefore != nil {
		s["minutes_before"] = *in.MinutesBefore
	}
	if in.AttributeContract != nil {
		s["attribute_contract"] = flattenSpWsTrustAttributeContract(in.AttributeContract)
	}
	if in.TokenProcessorMappings != nil {
		s["token_processor_mappings"] = flattenIdpTokenProcessorMappings(in.TokenProcessorMappings)
	}
	if in.AbortIfNotFulfilledFromRequest != nil {
		s["abort_if_not_fulfilled_from_request"] = *in.AbortIfNotFulfilledFromRequest
	}
	if in.OAuthAssertionProfiles != nil {
		s["o_auth_assertion_profiles"] = *in.OAuthAssertionProfiles
	}
	if in.DefaultTokenType != nil {
		s["default_token_type"] = *in.DefaultTokenType
	}
	if in.GenerateKey != nil {
		s["generate_key"] = *in.GenerateKey
	}
	if in.MinutesAfter != nil {
		s["minutes_after"] = *in.MinutesAfter
	}
	if in.EncryptSaml2Assertion != nil {
		s["encrypt_saml2_assertion"] = *in.EncryptSaml2Assertion
	}
	return append(m, s)
}

func flattenAuthnContextMapping(in *pf.AuthnContextMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Local != nil {
		s["local"] = *in.Local
	}
	if in.Remote != nil {
		s["remote"] = *in.Remote
	}
	return s
}
func flattenSaasAttributeMapping(in *pf.SaasAttributeMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.SaasFieldInfo != nil {
		s["saas_field_info"] = flattenSaasFieldConfiguration(in.SaasFieldInfo)
	}
	if in.FieldName != nil {
		s["field_name"] = *in.FieldName
	}
	return s
}
func flattenIdpOAuthAttributeContract(in *pf.IdpOAuthAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenIdpBrowserSsoAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenIdpBrowserSsoAttributes(in.ExtendedAttributes)
	}
	return append(m, s)
}

func flattenOIDCProviderSettings(in *pf.OIDCProviderSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.LoginType != nil {
		s["login_type"] = *in.LoginType
	}
	if in.AuthenticationSigningAlgorithm != nil {
		s["authentication_signing_algorithm"] = *in.AuthenticationSigningAlgorithm
	}
	if in.RequestSigningAlgorithm != nil {
		s["request_signing_algorithm"] = *in.RequestSigningAlgorithm
	}
	if in.TokenEndpoint != nil {
		s["token_endpoint"] = *in.TokenEndpoint
	}
	if in.JwksURL != nil {
		s["jwks_url"] = *in.JwksURL
	}
	if in.Scopes != nil {
		s["scopes"] = *in.Scopes
	}
	if in.AuthenticationScheme != nil {
		s["authentication_scheme"] = *in.AuthenticationScheme
	}
	if in.UserInfoEndpoint != nil {
		s["user_info_endpoint"] = *in.UserInfoEndpoint
	}
	if in.RequestParameters != nil {
		s["request_parameters"] = flattenOIDCRequestParameters(in.RequestParameters)
	}
	if in.AuthorizationEndpoint != nil {
		s["authorization_endpoint"] = *in.AuthorizationEndpoint
	}
	return append(m, s)
}
func flattenSpAttributeQuery(in *pf.SpAttributeQuery) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.Policy != nil {
		s["policy"] = flattenSpAttributeQueryPolicy(in.Policy)
	}
	if in.Attributes != nil && len(*in.Attributes) > 0 {
		s["attributes"] = *in.Attributes
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	return append(m, s)
}
func flattenAdditionalAllowedEntitiesConfiguration(in *pf.AdditionalAllowedEntitiesConfiguration) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.AllowAllEntities != nil {
		s["allow_all_entities"] = *in.AllowAllEntities
	}
	if in.AdditionalAllowedEntities != nil {
		s["additional_allowed_entities"] = flattenEntitys(in.AdditionalAllowedEntities)
	}
	if in.AllowAdditionalEntities != nil {
		s["allow_additional_entities"] = *in.AllowAdditionalEntities
	}
	return append(m, s)
}
func flattenSigningSettings(in *pf.SigningSettings) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.SigningKeyPairRef != nil {
		s["signing_key_pair_ref"] = flattenResourceLink(in.SigningKeyPairRef)
	}
	if in.Algorithm != nil {
		s["algorithm"] = *in.Algorithm
	}
	if in.IncludeCertInSignature != nil {
		s["include_cert_in_signature"] = *in.IncludeCertInSignature
	}
	if in.IncludeRawKeyInSignature != nil {
		s["include_raw_key_in_signature"] = *in.IncludeRawKeyInSignature
	}
	return append(m, s)
}
func flattenEntity(in *pf.Entity) map[string]interface{} {
	s := make(map[string]interface{})
	if in.EntityId != nil {
		s["entity_id"] = *in.EntityId
	}
	if in.EntityDescription != nil {
		s["entity_description"] = *in.EntityDescription
	}
	return s
}

func flattenIdpTokenProcessorMapping(in *pf.IdpTokenProcessorMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.IdpTokenProcessorRef != nil {
		s["idp_token_processor_ref"] = flattenResourceLink(in.IdpTokenProcessorRef)
	}
	if in.RestrictedVirtualEntityIds != nil {
		s["restricted_virtual_entity_ids"] = *in.RestrictedVirtualEntityIds
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	return s
}
func flattenArtifactResolverLocation(in *pf.ArtifactResolverLocation) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Index != nil {
		s["index"] = *in.Index
	}
	if in.Url != nil {
		s["url"] = *in.Url
	}
	return s
}

func flattenSpWsTrustAttribute(in *pf.SpWsTrustAttribute) map[string]interface{} {
	s := make(map[string]interface{})
	if in.Namespace != nil {
		s["namespace"] = *in.Namespace
	}
	if in.Name != nil {
		s["name"] = *in.Name
	}
	return s
}
func flattenSpBrowserSsoAttributeContract(in *pf.SpBrowserSsoAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenSpBrowserSsoAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenSpBrowserSsoAttributes(in.ExtendedAttributes)
	}
	return append(m, s)
}
func flattenProtocolMessageCustomization(in *pf.ProtocolMessageCustomization) map[string]interface{} {
	s := make(map[string]interface{})
	if in.ContextName != nil {
		s["context_name"] = *in.ContextName
	}
	if in.MessageExpression != nil {
		s["message_expression"] = *in.MessageExpression
	}
	return s
}
func flattenIdpBrowserSsoAttributeContract(in *pf.IdpBrowserSsoAttributeContract) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.CoreAttributes != nil {
		s["core_attributes"] = flattenIdpBrowserSsoAttributes(in.CoreAttributes)
	}
	if in.ExtendedAttributes != nil {
		s["extended_attributes"] = flattenIdpBrowserSsoAttributes(in.ExtendedAttributes)
	}
	return append(m, s)
}

func flattenOutboundBackChannelAuth(in *pf.OutboundBackChannelAuth) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, 1)
	s := make(map[string]interface{})
	if in.Type != nil {
		s["type"] = *in.Type
	}
	if in.HttpBasicCredentials != nil {
		s["http_basic_credentials"] = flattenUsernamePasswordCredentials(in.HttpBasicCredentials)
	}
	if in.DigitalSignature != nil {
		s["digital_signature"] = *in.DigitalSignature
	}
	if in.SslAuthKeyPairRef != nil {
		s["ssl_auth_key_pair_ref"] = flattenResourceLink(in.SslAuthKeyPairRef)
	}
	if in.ValidatePartnerCert != nil {
		s["validate_partner_cert"] = *in.ValidatePartnerCert
	}
	return append(m, s)
}

func flattenAuthenticationPolicyContractMapping(in *pf.AuthenticationPolicyContractMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.AuthenticationPolicyContractRef != nil {
		s["authentication_policy_contract_ref"] = flattenResourceLink(in.AuthenticationPolicyContractRef)
	}
	if in.RestrictVirtualServerIds != nil {
		s["restrict_virtual_server_ids"] = *in.RestrictVirtualServerIds
	}
	if in.RestrictedVirtualServerIds != nil {
		s["restricted_virtual_server_ids"] = *in.RestrictedVirtualServerIds
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	return s
}
func flattenAccessTokenManagerMapping(in *pf.AccessTokenManagerMapping) map[string]interface{} {
	s := make(map[string]interface{})
	if in.AttributeContractFulfillment != nil {
		s["attribute_contract_fulfillment"] = flattenMapOfAttributeFulfillmentValue(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		s["issuance_criteria"] = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.AccessTokenManagerRef != nil {
		s["access_token_manager_ref"] = flattenResourceLink(in.AccessTokenManagerRef)
	}
	if *in.AttributeSources != nil && len(*in.AttributeSources) > 0 {
		var ldapAttributes []interface{}
		var jdbcAttributes []interface{}
		var customAttributes []interface{}
		for _, v := range *in.AttributeSources {
			switch *v.Type {
			case "LDAP":
				ldapAttributes = append(ldapAttributes, flattenLdapAttributeSource(v))
			case "JDBC":
				jdbcAttributes = append(jdbcAttributes, flattenJdbcAttributeSource(v))
			case "CUSTOM":
				customAttributes = append(customAttributes, flattenCustomAttributeSource(v))
			}
		}
		if len(ldapAttributes) > 0 {
			s["ldap_attribute_source"] = ldapAttributes
		}
		if len(jdbcAttributes) > 0 {
			s["jdbc_attribute_source"] = jdbcAttributes
		}
		if len(customAttributes) > 0 {
			s["custom_attribute_source"] = customAttributes
		}
	}
	return s
}
func flattenConnectionCerts(in *[]*pf.ConnectionCert) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenConnectionCert(v))
	}
	return m
}
func flattenSloServiceEndpoints(in *[]*pf.SloServiceEndpoint) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSloServiceEndpoint(v))
	}
	return m
}
func flattenSchemaAttributes(in *[]*pf.SchemaAttribute) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSchemaAttribute(v))
	}
	return m
}
func flattenAuthenticationPolicyContractMappings(in *[]*pf.AuthenticationPolicyContractMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenAuthenticationPolicyContractMapping(v))
	}
	return m
}
func flattenIdpWsTrustAttributes(in *[]*pf.IdpWsTrustAttribute) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpWsTrustAttribute(v))
	}
	return m
}

func flattenEntitys(in *[]*pf.Entity) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenEntity(v))
	}
	return m
}
func flattenExpressionIssuanceCriteriaEntrys(in *[]*pf.ExpressionIssuanceCriteriaEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenExpressionIssuanceCriteriaEntry(v))
	}
	return m
}
func flattenUrlWhitelistEntrys(in *[]*pf.UrlWhitelistEntry) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenUrlWhitelistEntry(v))
	}
	return m
}
func flattenIdpAdapterAssertionMappings(in *[]*pf.IdpAdapterAssertionMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpAdapterAssertionMapping(v))
	}
	return m
}
func flattenAuthenticationPolicyContractAssertionMappings(in *[]*pf.AuthenticationPolicyContractAssertionMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenAuthenticationPolicyContractAssertionMapping(v))
	}
	return m
}
func flattenIdpAdapterAttributes(in *[]*pf.IdpAdapterAttribute) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpAdapterAttribute(v))
	}
	return m
}
func flattenOIDCRequestParameters(in *[]*pf.OIDCRequestParameter) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenOIDCRequestParameter(v))
	}
	return m
}

func flattenAttributeQueryNameMappings(in *[]*pf.AttributeQueryNameMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenAttributeQueryNameMapping(v))
	}
	return m
}
func flattenChannels(in *[]*pf.Channel) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenChannel(v))
	}
	return m
}
func flattenArtifactResolverLocations(in *[]*pf.ArtifactResolverLocation) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenArtifactResolverLocation(v))
	}
	return m
}
func flattenSaasAttributeMappings(in *[]*pf.SaasAttributeMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSaasAttributeMapping(v))
	}
	return m
}
func flattenIdpBrowserSsoAttributes(in *[]*pf.IdpBrowserSsoAttribute) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpBrowserSsoAttribute(v))
	}
	return m
}
func flattenConditionalIssuanceCriteriaEntrys(in *[]*pf.ConditionalIssuanceCriteriaEntry) []map[string]interface{} {
	m := make([]map[string]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenConditionalIssuanceCriteriaEntry(v))
	}
	return m
}

func flattenLdapAttributeSources(in *[]*pf.AttributeSource) []interface{} {
	var m []interface{}
	for _, v := range *in {
		if *v.Type == "LDAP" {
			m = append(m, flattenLdapAttributeSource(v))
		}
	}
	return m
}
func flattenJdbcAttributeSources(in *[]*pf.AttributeSource) []interface{} {
	var m []interface{}
	for _, v := range *in {
		if *v.Type == "JDBC" {
			m = append(m, flattenJdbcAttributeSource(v))
		}
	}
	return m
}

func flattenCustomAttributeSources(in *[]*pf.AttributeSource) []interface{} {
	var m []interface{}
	for _, v := range *in {
		if *v.Type == "CUSTOM" {
			m = append(m, flattenCustomAttributeSource(v))
		}
	}
	return m
}

func flattenProtocolMessageCustomizations(in *[]*pf.ProtocolMessageCustomization) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenProtocolMessageCustomization(v))
	}
	return m
}
func flattenSpSsoServiceEndpoints(in *[]*pf.SpSsoServiceEndpoint) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSpSsoServiceEndpoint(v))
	}
	return m
}
func flattenAuthnContextMappings(in *[]*pf.AuthnContextMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenAuthnContextMapping(v))
	}
	return m
}

//func flattenConfigTables(in *[]*pf.ConfigTable) []interface{} {
//	m := make([]interface{}, 0, len(*in))
//	for _, v := range *in {
//		m = append(m, flattenConfigTable(v))
//	}
//	return m
//}
func flattenIdpSsoServiceEndpoints(in *[]*pf.IdpSsoServiceEndpoint) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpSsoServiceEndpoint(v))
	}
	return m
}
func flattenSpTokenGeneratorMappings(in *[]*pf.SpTokenGeneratorMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSpTokenGeneratorMapping(v))
	}
	return m
}

//func flattenFieldEntrys(in *[]*pf.FieldEntry) []interface{} {
//	m := make([]interface{}, 0, len(*in))
//	for _, v := range *in {
//		m = append(m, flattenFieldEntry(v))
//	}
//	return m
//}
func flattenAccessTokenManagerMappings(in *[]*pf.AccessTokenManagerMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenAccessTokenManagerMapping(v))
	}
	return m
}

func flattenSpAdapterMappings(in *[]*pf.SpAdapterMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSpAdapterMapping(v))
	}
	return m
}
func flattenIdpTokenProcessorMappings(in *[]*pf.IdpTokenProcessorMapping) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenIdpTokenProcessorMapping(v))
	}
	return m
}
func flattenSpWsTrustAttributes(in *[]*pf.SpWsTrustAttribute) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSpWsTrustAttribute(v))
	}
	return m
}

//func flattenConfigFields(in *[]*pf.ConfigField) []interface{} {
//	m := make([]interface{}, 0, len(*in))
//	for _, v := range *in {
//		m = append(m, flattenConfigField(v))
//	}
//	return m
//}
func flattenSpBrowserSsoAttributes(in *[]*pf.SpBrowserSsoAttribute) []interface{} {
	m := make([]interface{}, 0, len(*in))
	for _, v := range *in {
		m = append(m, flattenSpBrowserSsoAttribute(v))
	}
	return m
}
