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
		Description:   "Provides configuration for OAuth Authentication Policy Contract Mappings within PingFederate.",
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
		"authentication_policy_contract_ref": {
			Type:        schema.TypeList,
			Required:    true,
			ForceNew:    true,
			MaxItems:    1,
			Description: "Reference to the associated authentication policy contract. The reference cannot be changed after the mapping has been created.",
			Elem:        resourceLinkResource(),
		},
		"ldap_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of configured ldap data stores to look up attributes from.",
			Elem:        resourceLdapAttributeSource(),
		},
		"jdbc_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of configured jdbc data stores to look up attributes from.",
			Elem:        resourceJdbcAttributeSource(),
		},
		"custom_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of configured custom data stores to look up attributes from.",
			Elem:        resourceCustomAttributeSource(),
		},
		"attribute_contract_fulfillment": {
			Type:        schema.TypeSet,
			Required:    true,
			Description: "A list of mappings from attribute names to their fulfillment values.",
			Elem:        resourceAttributeFulfillmentValue(),
		},
		"issuance_criteria": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
			Elem:        resourceIssuanceCriteria(),
		},
	}
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.CreateApcMappingInput{
		Body: *resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d),
	}
	result, _, err := svc.CreateApcMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create OauthAuthenticationPolicyContractMappings: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.GetApcMappingInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApcMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OauthAuthenticationPolicyContractMappings: %s", err)
	}
	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.UpdateApcMappingInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d),
	}
	result, _, err := svc.UpdateApcMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OauthAuthenticationPolicyContractMappings: %s", err)
	}

	return resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAuthenticationPolicyContractMappings
	input := oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteApcMappingWithContext(ctx, &input)
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
	if rv.IssuanceCriteria != nil && issuanceCriteriaShouldFlatten(rv.IssuanceCriteria) {
		if err := d.Set("issuance_criteria", flattenIssuanceCriteria(rv.IssuanceCriteria)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AttributeSources != nil {
		if m := flattenLdapAttributeSources(rv.AttributeSources); len(m) > 0 {
			if err := d.Set("ldap_attribute_source", m); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
		if m := flattenJdbcAttributeSources(rv.AttributeSources); len(m) > 0 {
			if err := d.Set("jdbc_attribute_source", m); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
		if m := flattenCustomAttributeSources(rv.AttributeSources); len(m) > 0 {
			if err := d.Set("custom_attribute_source", m); err != nil {
				diags = append(diags, diag.FromErr(err)...)
			}
		}
	}
	return diags
}

func resourcePingFederateOauthAuthenticationPolicyContractMappingsResourceReadData(d *schema.ResourceData) *pf.ApcToPersistentGrantMapping {
	mapping := &pf.ApcToPersistentGrantMapping{
		Id:                              String(d.Id()),
		AuthenticationPolicyContractRef: expandResourceLink(d.Get("authentication_policy_contract_ref").([]interface{})[0].(map[string]interface{})),
		AttributeContractFulfillment:    expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
		AttributeSources:                &[]*pf.AttributeSource{},
	}
	if v, ok := d.GetOk("issuance_criteria"); ok && len(v.([]interface{})) > 0 {
		mapping.IssuanceCriteria = expandIssuanceCriteria(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("ldap_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("jdbc_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("custom_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*mapping.AttributeSources = append(*mapping.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	return mapping
}
