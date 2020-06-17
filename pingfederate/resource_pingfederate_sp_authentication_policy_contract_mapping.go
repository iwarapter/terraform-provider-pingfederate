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
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
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
		"source_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"issuance_criteria": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceIssuanceCriteria(),
		},
		"target_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"default_target_resource": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"license_connection_group_assignment": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAuthenticationPolicyContractMappings
	input := spAuthenticationPolicyContractMappings.CreateApcToSpAdapterMappingInput{
		Body:                     *resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.CreateApcToSpAdapterMapping(&input)
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
	result, _, err := svc.GetApcToSpAdapterMappingById(&input)
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
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.UpdateApcToSpAdapterMappingById(&input)
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
	_, _, err := svc.DeleteApcToSpAdapterMappingById(&input)
	if err != nil {
		return diag.Errorf("unable to delete SpAuthenticationPolicyContractMapping: %s", err)
	}
	return nil
}

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadResult(d *schema.ResourceData, rv *pf.ApcToSpAdapterMapping) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "source_id", rv.SourceId, &diags)
	setResourceDataStringithDiagnostic(d, "target_id", rv.TargetId, &diags)
	setResourceDataStringithDiagnostic(d, "default_target_resource", rv.DefaultTargetResource, &diags)
	setResourceDataStringithDiagnostic(d, "license_connection_group_assignment", rv.LicenseConnectionGroupAssignment, &diags)

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

func resourcePingFederateSpAuthenticationPolicyContractMappingResourceReadData(d *schema.ResourceData) *pf.ApcToSpAdapterMapping {
	mapping := &pf.ApcToSpAdapterMapping{
		Id:                           String(d.Id()),
		SourceId:                     String(d.Get("source_id").(string)),
		TargetId:                     String(d.Get("target_id").(string)),
		AttributeContractFulfillment: expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
		AttributeSources:             &[]*pf.AttributeSource{},
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
	if v, ok := d.GetOk("default_target_resource"); ok {
		mapping.DefaultTargetResource = String(v.(string))
	}
	if v, ok := d.GetOk("license_connection_group_assignment"); ok {
		mapping.LicenseConnectionGroupAssignment = String(v.(string))
	}
	return mapping
}
