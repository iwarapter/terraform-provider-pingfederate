package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthResourceOwnerCredentialsMappings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for OAuth Resource Owner Credentials within PingFederate.",
		CreateContext: resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceCreate,
		ReadContext:   resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceRead,
		UpdateContext: resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceUpdate,
		DeleteContext: resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceSchema(),
	}
}

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"password_validator_ref": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "Reference to the associated Source Password Validator Instance.",
			Elem:        resourceLinkResource(),
		},
		"ldap_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of ldap configured data stores to look up attributes from.",
			Elem:        resourceLdapAttributeSource(),
		},
		"jdbc_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of jdbc configured data stores to look up attributes from.",
			Elem:        resourceJdbcAttributeSource(),
		},
		"custom_attribute_source": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "A list of custom configured data stores to look up attributes from.",
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

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthResourceOwnerCredentialsMappings
	input := oauthResourceOwnerCredentialsMappings.CreateResourceOwnerCredentialsMappingInput{
		Body:                     *resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreateResourceOwnerCredentialsMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create OAuthResourceOwnerCredentialsMappings: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthResourceOwnerCredentialsMappings
	input := oauthResourceOwnerCredentialsMappings.GetResourceOwnerCredentialsMappingInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetResourceOwnerCredentialsMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OAuthResourceOwnerCredentialsMappings: %s", err)
	}
	return resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthResourceOwnerCredentialsMappings
	input := oauthResourceOwnerCredentialsMappings.UpdateResourceOwnerCredentialsMappingInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateResourceOwnerCredentialsMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OAuthResourceOwnerCredentialsMappings: %s", err)
	}

	return resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadResult(d, result)
}

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthResourceOwnerCredentialsMappings
	input := oauthResourceOwnerCredentialsMappings.DeleteResourceOwnerCredentialsMappingInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteResourceOwnerCredentialsMappingWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete OAuthResourceOwnerCredentialsMappings: %s", err)
	}
	return nil
}

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadResult(d *schema.ResourceData, rv *pf.ResourceOwnerCredentialsMapping) diag.Diagnostics {
	var diags diag.Diagnostics
	if rv.PasswordValidatorRef != nil {
		if err := d.Set("password_validator_ref", flattenResourceLink(rv.PasswordValidatorRef)); err != nil {
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

func resourcePingFederateOauthResourceOwnerCredentialsMappingsResourceReadData(d *schema.ResourceData) *pf.ResourceOwnerCredentialsMapping {
	result := &pf.ResourceOwnerCredentialsMapping{
		PasswordValidatorRef:         expandResourceLink(d.Get("password_validator_ref").([]interface{})[0].(map[string]interface{})),
		AttributeContractFulfillment: expandMapOfAttributeFulfillmentValue(d.Get("attribute_contract_fulfillment").(*schema.Set).List()),
		AttributeSources:             &[]*pf.AttributeSource{},
	}
	result.Id = result.PasswordValidatorRef.Id
	if v, ok := d.GetOk("issuance_criteria"); ok {
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
	return result
}
