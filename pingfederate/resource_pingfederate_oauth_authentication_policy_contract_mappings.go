package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceCreate,
		ReadContext:   resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceRead,
		UpdateContext: resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceUpdate,
		DeleteContext: resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceSchema(),
	}
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authentication_policy_contract_ref": resourceLinkSchema(),
		"ldap_attribute_source": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceLdapAttributeSource(),
		},
		"jdbc_attribute_source": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceJdbcAttributeSource(),
		},
		"custom_attribute_source": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     resourceCustomAttributeSource(),
		},
		"attribute_contract_fulfillment": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     resourceAttributeFulfillmentValue(),
		},
		"issuance_criteria": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceIssuanceCriteria(),
		},
	}
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.CreateApcMappingInput{
		Body: *resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d),
	}
	result, _, err := svc.CreateApcMapping(&input)
	if err != nil {
		return diag.Errorf("unable to create OauthAuthenticationPolicyContractMappings: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.GetApcMappingInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApcMapping(&input)
	if err != nil {
		return diag.Errorf("unable to read OauthAuthenticationPolicyContractMappings: %s", err)
	}
	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.UpdateApcMappingInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d),
	}
	result, _, err := svc.UpdateApcMapping(&input)
	if err != nil {
		return diag.Errorf("unable to update OauthAuthenticationPolicyContractMappings: %s", err)
	}

	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteApcMapping(&input)
	if err != nil {
		return diag.Errorf("unable to delete OauthAuthenticationPolicyContractMappings: %s", err)
	}
	return nil
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d *schema.ResourceData, rv *pf.ApcToPersistentGrantMapping) diag.Diagnostics {
	var diags diag.Diagnostics

	if rv.AuthenticationPolicyContractRef != nil {
		if err := d.Set("authentication_policy_contract_ref", flattenResourceLink(rv.AuthenticationPolicyContractRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.AttributeContractFulfillment != nil {
		if err := d.Set("attribute_contract_fulfillment", flattenMapOfAttributeFulfillmentValue(rv.AttributeContractFulfillment)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.IssuanceCriteria != nil && (rv.IssuanceCriteria.ExpressionCriteria != nil && rv.IssuanceCriteria.ConditionalCriteria != nil) {
		if err := d.Set("issuance_criteria", flattenIssuanceCriteria(rv.IssuanceCriteria)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if err := flattenAttributeSources(d, rv.AttributeSources); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d *schema.ResourceData) *pf.ApcToPersistentGrantMapping {
	mapping := &pf.ApcToPersistentGrantMapping{
		Id:                              String(d.Id()),
		AuthenticationPolicyContractRef: expandResourceLink(d.Get("authentication_policy_contract_ref").([]interface{})),
		AttributeContractFulfillment:    expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
		AttributeSources:                &[]*pf.AttributeSource{},
	}
	if v, ok := d.GetOk("issuance_criteria"); ok {
		mapping.IssuanceCriteria = expandIssuanceCriteria(v.([]interface{}))
	}
	if v, ok := d.GetOk("ldap_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandLdapAttributeSource(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("jdbc_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandJdbcAttributeSource(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("custom_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandCustomAttributeSource(v.([]interface{}))...)
	}
	return mapping
}
