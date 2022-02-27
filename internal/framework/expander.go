package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func expandApcToPersistentGrantMapping(in ApcToPersistentGrantMappingData) *pf.ApcToPersistentGrantMapping {
	var result pf.ApcToPersistentGrantMapping
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.AuthenticationPolicyContractRef.Unknown && !in.AuthenticationPolicyContractRef.Null {
		result.AuthenticationPolicyContractRef = &pf.ResourceLink{Id: String(in.AuthenticationPolicyContractRef.Value)}
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
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil {
		result.IssuanceCriteria = expandIssuanceCriteria(*in.IssuanceCriteria)
	}

	return &result
}

func expandAuthenticationPolicyContract(in AuthenticationPolicyContractData) *pf.AuthenticationPolicyContract {
	result := pf.AuthenticationPolicyContract{
		CoreAttributes: &[]*pf.AuthenticationPolicyContractAttribute{
			{
				Name: String("subject"),
			},
		},
	}

	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
	}
	//if in.CoreAttributes != nil {
	//	result.CoreAttributes = expandAuthenticationPolicyContractAttributes(in.CoreAttributes)
	//}
	if in.ExtendedAttributes != nil && len(in.ExtendedAttributes) > 0 {
		attrs := []*pf.AuthenticationPolicyContractAttribute{}
		for _, data := range in.ExtendedAttributes {
			attrs = append(attrs, &pf.AuthenticationPolicyContractAttribute{Name: String(data.Value)})
		}
		result.ExtendedAttributes = &attrs
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

func expandAuthenticationPolicyContractAttributes(in *[]*AuthenticationPolicyContractAttributeData) *[]*pf.AuthenticationPolicyContractAttribute {
	results := []*pf.AuthenticationPolicyContractAttribute{}
	for _, data := range *in {
		results = append(results, expandAuthenticationPolicyContractAttribute(*data))
	}
	return &results
}
func expandAuthenticationPolicyContractAttribute(in AuthenticationPolicyContractAttributeData) *pf.AuthenticationPolicyContractAttribute {
	var result pf.AuthenticationPolicyContractAttribute
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
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

func expandConditionalIssuanceCriteriaEntrys(in *[]*ConditionalIssuanceCriteriaEntryData) *[]*pf.ConditionalIssuanceCriteriaEntry {
	results := []*pf.ConditionalIssuanceCriteriaEntry{}
	for _, data := range *in {
		results = append(results, expandConditionalIssuanceCriteriaEntry(*data))
	}
	return &results
}
func expandConditionalIssuanceCriteriaEntry(in ConditionalIssuanceCriteriaEntryData) *pf.ConditionalIssuanceCriteriaEntry {
	var result pf.ConditionalIssuanceCriteriaEntry
	if !in.ErrorResult.Unknown && !in.ErrorResult.Null {
		result.ErrorResult = String(in.ErrorResult.Value)
	}
	if in.Source != nil {
		result.Source = expandSourceTypeIdKey(*in.Source)
	}
	if !in.AttributeName.Unknown && !in.AttributeName.Null {
		result.AttributeName = String(in.AttributeName.Value)
	}
	if !in.Condition.Unknown && !in.Condition.Null {
		result.Condition = String(in.Condition.Value)
	}
	if !in.Value.Unknown && !in.Value.Null {
		result.Value = String(in.Value.Value)
	}

	return &result
}

func expandCustomAttributeSource(in CustomAttributeSourceData) *pf.CustomAttributeSource {
	var result pf.CustomAttributeSource
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}
	if in.FilterFields != nil {
		result.FilterFields = expandFieldEntrys(in.FilterFields)
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
	if !in.Value.Unknown && !in.Value.Null {
		result.Value = String(in.Value.Value)
	}
	if !in.Name.Unknown && !in.Name.Null {
		result.Name = String(in.Name.Value)
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
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.Schema.Unknown && !in.Schema.Null {
		result.Schema = String(in.Schema.Value)
	}
	if !in.Table.Unknown && !in.Table.Null {
		result.Table = String(in.Table.Value)
	}
	if in.ColumnNames != nil {
		result.ColumnNames = expandStringList(in.ColumnNames)
	}
	if !in.Filter.Unknown && !in.Filter.Null {
		result.Filter = String(in.Filter.Value)
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}

	return &result
}

func expandLdapAttributeSource(in LdapAttributeSourceData) *pf.LdapAttributeSource {
	var result pf.LdapAttributeSource
	if !in.DataStoreRef.Unknown && !in.DataStoreRef.Null {
		result.DataStoreRef = &pf.ResourceLink{Id: String(in.DataStoreRef.Value)}
	}
	if !in.Description.Unknown && !in.Description.Null {
		result.Description = String(in.Description.Value)
	}
	if !in.MemberOfNestedGroup.Unknown && !in.MemberOfNestedGroup.Null {
		result.MemberOfNestedGroup = Bool(in.MemberOfNestedGroup.Value)
	}
	if !in.SearchFilter.Unknown && !in.SearchFilter.Null {
		result.SearchFilter = String(in.SearchFilter.Value)
	}
	if !in.Id.Unknown && !in.Id.Null {
		result.Id = String(in.Id.Value)
	}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = expandMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if !in.BaseDn.Unknown && !in.BaseDn.Null {
		result.BaseDn = String(in.BaseDn.Value)
	}
	if !in.SearchScope.Unknown && !in.SearchScope.Null {
		result.SearchScope = String(in.SearchScope.Value)
	}
	if in.SearchAttributes != nil {
		result.SearchAttributes = expandStringList(in.SearchAttributes)
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = expandMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
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
