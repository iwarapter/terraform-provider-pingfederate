package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingfederate-sdk-go/services/spAuthenticationPolicyContractMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateSpAuthenticationPolicyContractMappingResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for SP Authentication Policy Contract Mappings within PingFederate.",
		CreateContext: resourcePingFederateSpAuthenticationPolicyContractMappingResourceCreate,
		ReadContext:   resourcePingFederateSpAuthenticationPolicyContractMappingResourceRead,
		UpdateContext: resourcePingFederateSpAuthenticationPolicyContractMappingResourceUpdate,
		DeleteContext: resourcePingFederateSpAuthenticationPolicyContractMappingResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateSpAuthenticationPolicyContractMappingResourceSchema(),
	}
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"source_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the Authentication Policy Contract.",
		},
		"issuance_criteria": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
			Elem:        resourceIssuanceCriteria(),
		},
		"target_id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The id of the SP Adapter.",
		},
		"default_target_resource": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Default target URL for this APC-to-adapter mapping configuration.",
		},
		"license_connection_group_assignment": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The license connection group.",
		},
	}
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAuthenticationPolicyContractMappings
	input := spAuthenticationPolicyContractMappings.CreateApcToSpAdapterMappingInput{
		Body:                     *resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreateApcToSpAdapterMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create SpAuthenticationPolicyContractMapping: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(d, result)
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAuthenticationPolicyContractMappings
	input := spAuthenticationPolicyContractMappings.GetApcToSpAdapterMappingByIdInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApcToSpAdapterMappingByIdWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read SpAuthenticationPolicyContractMapping: %s", err)
	}
	return resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(d, result)
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAuthenticationPolicyContractMappings
	input := spAuthenticationPolicyContractMappings.UpdateApcToSpAdapterMappingByIdInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateApcToSpAdapterMappingByIdWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update SpAuthenticationPolicyContractMapping: %s", err)
	}

	return resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(d, result)
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAuthenticationPolicyContractMappings
	input := spAuthenticationPolicyContractMappings.DeleteApcToSpAdapterMappingByIdInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteApcToSpAdapterMappingByIdWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete SpAuthenticationPolicyContractMapping: %s", err)
	}
	return nil
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(d *schema.ResourceData, rv *pf.ApcToSpAdapterMapping) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "source_id", rv.SourceId, &diags)
	setResourceDataStringWithDiagnostic(d, "target_id", rv.TargetId, &diags)
	setResourceDataStringWithDiagnostic(d, "default_target_resource", rv.DefaultTargetResource, &diags)
	setResourceDataStringWithDiagnostic(d, "license_connection_group_assignment", rv.LicenseConnectionGroupAssignment, &diags)

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

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(d *schema.ResourceData) *pf.ApcToSpAdapterMapping {
	result := &pf.ApcToSpAdapterMapping{
		Id:                           String(d.Id()),
		SourceId:                     String(d.Get("source_id").(string)),
		TargetId:                     String(d.Get("target_id").(string)),
		AttributeContractFulfillment: expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
		AttributeSources:             &[]*pf.AttributeSource{},
	}
	if v, ok := d.GetOk("issuance_criteria"); ok && len(v.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("ldap_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("jdbc_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("custom_attribute_source"); ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := d.GetOk("default_target_resource"); ok {
		result.DefaultTargetResource = String(v.(string))
	}
	if v, ok := d.GetOk("license_connection_group_assignment"); ok {
		result.LicenseConnectionGroupAssignment = String(v.(string))
	}
	return result
}
