package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateOpenIdConnectPolicyResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateOpenIdConnectPolicyResourceCreate,
		Read:   resourcePingFederateOpenIdConnectPolicyResourceRead,
		Update: resourcePingFederateOpenIdConnectPolicyResourceUpdate,
		Delete: resourcePingFederateOpenIdConnectPolicyResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingFederateOpenIdConnectPolicyResourceSchema(),
	}
}

func resourcePingFederateOpenIdConnectPolicyResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
		"policy_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"access_token_manager_ref": resourceRequiredLinkSchema(),
		"id_token_lifetime": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"include_sri_in_id_token": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"include_user_in_id_token": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"include_shash_in_id_token": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"return_id_token_on_refresh_grant": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"attribute_contract": resourceOpenIdConnectAttributeContract(),
		"attribute_mapping": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceAttributeMapping(),
		},
		//"scope_attribute_mappings": {
		//	Type:     schema.TypeMap,
		//	Optional: true,
		//	Elem: &schema.Schema{
		//		Type: schema.TypeList,
		//		Elem: &schema.Schema{
		//			Type: schema.TypeString,
		//		},
		//	},
		//},
	}
}

func resourcePingFederateOpenIdConnectPolicyResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthOpenIdConnect
	input := pf.CreatePolicyInput{
		Body:                     *resourcePingFederateOpenIdConnectPolicyResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	input.Body.Id = input.Body.Name
	result, _, err := svc.CreatePolicy(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthOpenIdConnect
	input := pf.GetPolicyInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetPolicy(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthOpenIdConnect
	input := pf.UpdatePolicyInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateOpenIdConnectPolicyResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.UpdatePolicy(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthOpenIdConnect
	input := pf.DeletePolicyInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeletePolicy(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func resourcePingFederateOpenIdConnectPolicyResourceReadResult(d *schema.ResourceData, rv *pf.OpenIdConnectPolicy) (err error) {
	setResourceDataString(d, "name", rv.Name)
	setResourceDataString(d, "policy_id", rv.Id)
	if rv.AccessTokenManagerRef != nil {
		if err = d.Set("access_token_manager_ref", flattenResourceLink(rv.AccessTokenManagerRef)); err != nil {
			return err
		}
	}
	setResourceDataInt(d, "id_token_lifetime", rv.IdTokenLifetime)
	setResourceDataBool(d, "include_sri_in_id_token", rv.IncludeSriInIdToken)
	setResourceDataBool(d, "include_user_in_id_token", rv.IncludeUserInfoInIdToken)
	setResourceDataBool(d, "include_shash_in_id_token", rv.IncludeSHashInIdToken)
	setResourceDataBool(d, "return_id_token_on_refresh_grant", rv.ReturnIdTokenOnRefreshGrant)
	if rv.AttributeContract != nil {
		if err = d.Set("attribute_contract", flattenOpenIdConnectAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}
	if rv.AttributeMapping != nil {
		if err = d.Set("attribute_mapping", flattenAttributeMapping(rv.AttributeMapping)); err != nil {
			return err
		}
	}
	//if rv.ScopeAttributeMappings != nil {
	//	if err = d.Set("scope_attribute_mappings", flattenScopeAttributeMappings(rv.ScopeAttributeMappings)); err != nil {
	//		return err
	//	}
	//}
	return nil
}

func resourcePingFederateOpenIdConnectPolicyResourceReadData(d *schema.ResourceData) *pf.OpenIdConnectPolicy {
	policy := &pf.OpenIdConnectPolicy{
		Id:                    String(d.Get("policy_id").(string)),
		Name:                  String(d.Get("name").(string)),
		AccessTokenManagerRef: expandResourceLink(d.Get("access_token_manager_ref").([]interface{})),
		AttributeContract:     expandOpenIdConnectAttributeContract(d.Get("attribute_contract").([]interface{})),
		AttributeMapping:      expandAttributeMapping(d.Get("attribute_mapping").([]interface{})),
	}

	if v, ok := d.GetOk("id_token_lifetime"); ok {
		policy.IdTokenLifetime = Int(v.(int))
	}
	if v, ok := d.GetOk("include_sri_in_id_token"); ok {
		policy.IncludeSriInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOk("include_user_in_id_token"); ok {
		policy.IncludeUserInfoInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOk("include_shash_in_id_token"); ok {
		policy.IncludeSHashInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOk("return_id_token_on_refresh_grant"); ok {
		policy.ReturnIdTokenOnRefreshGrant = Bool(v.(bool))
	}
	//if v, ok := d.GetOk("scope_attribute_mappings"); ok {
	//	policy.ScopeAttributeMappings = expandScopeAttributeMappings(v.(map[string]interface{}))
	//}
	return policy
}
