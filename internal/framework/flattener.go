package framework

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func flattenApcToPersistentGrantMapping(in *pf.ApcToPersistentGrantMapping) *ApcToPersistentGrantMappingData {
	result := ApcToPersistentGrantMappingData{}
	if in.AuthenticationPolicyContractRef != nil {
		result.AuthenticationPolicyContractRef = types.String{Value: *in.AuthenticationPolicyContractRef.Id}
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
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(in.IssuanceCriteria) {
		result.IssuanceCriteria = flattenIssuanceCriteria(in.IssuanceCriteria)
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	}

	return &result
}

func flattenAuthenticationPolicyContract(in *pf.AuthenticationPolicyContract) *AuthenticationPolicyContractData {
	result := AuthenticationPolicyContractData{}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
	} else {
		result.Name = types.String{Null: true}
	}
	if in.CoreAttributes != nil {
		result.CoreAttributes = flattenAuthenticationPolicyContractAttributes(in.CoreAttributes)
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

func flattenAuthenticationPolicyContractAttributes(in *[]*pf.AuthenticationPolicyContractAttribute) *[]*AuthenticationPolicyContractAttributeData {
	results := []*AuthenticationPolicyContractAttributeData{}
	for _, data := range *in {
		results = append(results, flattenAuthenticationPolicyContractAttribute(data))
	}
	return &results
}
func flattenAuthenticationPolicyContractAttribute(in *pf.AuthenticationPolicyContractAttribute) *AuthenticationPolicyContractAttributeData {
	result := AuthenticationPolicyContractAttributeData{}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
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

func flattenConditionalIssuanceCriteriaEntrys(in *[]*pf.ConditionalIssuanceCriteriaEntry) *[]*ConditionalIssuanceCriteriaEntryData {
	results := []*ConditionalIssuanceCriteriaEntryData{}
	for _, data := range *in {
		results = append(results, flattenConditionalIssuanceCriteriaEntry(data))
	}
	return &results
}
func flattenConditionalIssuanceCriteriaEntry(in *pf.ConditionalIssuanceCriteriaEntry) *ConditionalIssuanceCriteriaEntryData {
	result := ConditionalIssuanceCriteriaEntryData{}
	if in.Source != nil {
		result.Source = flattenSourceTypeIdKey(in.Source)
	}
	if in.AttributeName != nil {
		result.AttributeName = types.String{Value: *in.AttributeName}
	}
	if in.Condition != nil {
		result.Condition = types.String{Value: *in.Condition}
	}
	if in.Value != nil {
		result.Value = types.String{Value: *in.Value}
	}
	if in.ErrorResult != nil {
		result.ErrorResult = types.String{Value: *in.ErrorResult}
	} else {
		result.ErrorResult = types.String{Null: true}
	}

	return &result
}

func flattenCustomAttributeSource(in *pf.CustomAttributeSource) *CustomAttributeSourceData {
	result := CustomAttributeSourceData{}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.DataStoreRef != nil {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	}
	if in.FilterFields != nil {
		result.FilterFields = flattenFieldEntrys(in.FilterFields)
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
	if in.Expression != nil {
		result.Expression = types.String{Value: *in.Expression}
	}
	if in.ErrorResult != nil {
		result.ErrorResult = types.String{Value: *in.ErrorResult}
	} else {
		result.ErrorResult = types.String{Null: true}
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
	if in.Value != nil {
		result.Value = types.String{Value: *in.Value}
	} else {
		result.Value = types.String{Null: true}
	}
	if in.Name != nil {
		result.Name = types.String{Value: *in.Name}
	}

	return &result
}

func flattenIssuanceCriteria(in *pf.IssuanceCriteria) *IssuanceCriteriaData {
	result := IssuanceCriteriaData{}
	if in.ExpressionCriteria != nil {
		result.ExpressionCriteria = flattenExpressionIssuanceCriteriaEntrys(in.ExpressionCriteria)
	}
	if in.ConditionalCriteria != nil {
		result.ConditionalCriteria = flattenConditionalIssuanceCriteriaEntrys(in.ConditionalCriteria)
	}

	return &result
}

func flattenJdbcAttributeSource(in *pf.JdbcAttributeSource) *JdbcAttributeSourceData {
	result := JdbcAttributeSourceData{}
	if in.DataStoreRef != nil {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
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
	if in.ColumnNames != nil {
		result.ColumnNames = flattenStringList(*in.ColumnNames)
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
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}

	return &result
}

func flattenLdapAttributeSource(in *pf.LdapAttributeSource) *LdapAttributeSourceData {
	result := LdapAttributeSourceData{}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
	}
	if in.AttributeContractFulfillment != nil {
		result.AttributeContractFulfillment = flattenMapAttributeFulfillmentValues(in.AttributeContractFulfillment)
	}
	if in.BaseDn != nil {
		result.BaseDn = types.String{Value: *in.BaseDn}
	} else {
		result.BaseDn = types.String{Null: true}
	}
	if in.SearchScope != nil {
		result.SearchScope = types.String{Value: *in.SearchScope}
	} else {
		result.SearchScope = types.String{Null: true}
	}
	if in.SearchAttributes != nil {
		result.SearchAttributes = flattenStringList(*in.SearchAttributes)
	}
	if in.BinaryAttributeSettings != nil {
		result.BinaryAttributeSettings = flattenMapBinaryLdapAttributeSettingss(in.BinaryAttributeSettings)
	}
	if in.DataStoreRef != nil {
		result.DataStoreRef = types.String{Value: *in.DataStoreRef.Id}
	}
	if in.Description != nil {
		result.Description = types.String{Value: *in.Description}
	} else {
		result.Description = types.String{Null: true}
	}
	if in.MemberOfNestedGroup != nil {
		result.MemberOfNestedGroup = types.Bool{Value: *in.MemberOfNestedGroup}
	} else {
		result.MemberOfNestedGroup = types.Bool{Null: true}
	}
	if in.SearchFilter != nil {
		result.SearchFilter = types.String{Value: *in.SearchFilter}
	} else {
		result.SearchFilter = types.String{Null: true}
	}

	return &result
}

func flattenSourceTypeIdKey(in *pf.SourceTypeIdKey) *SourceTypeIdKeyData {
	result := SourceTypeIdKeyData{}
	if in.Type != nil {
		result.Type = types.String{Value: *in.Type}
	}
	if in.Id != nil {
		result.Id = types.String{Value: *in.Id}
	} else {
		result.Id = types.String{Null: true}
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
