package pingfederate

import (
	"context"
	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateAuthenticationPoliciesResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthenticationPoliciesResourceCreate,
		ReadContext:   resourcePingFederateAuthenticationPoliciesResourceRead,
		UpdateContext: resourcePingFederateAuthenticationPoliciesResourceUpdate,
		DeleteContext: resourcePingFederateAuthenticationPoliciesResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthenticationPoliciesResourceSchema(),
	}
}

func resourcePingFederateAuthenticationPoliciesResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"fail_if_no_selection": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"tracked_http_parameters": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"authn_selection_trees": resourceAuthenticationPolicyTreeSchema(),
		"default_authentication_sources": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
						ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
							v := value.(string)
							switch v {
							case
								"IDP_ADAPTER",
								"IDP_CONNECTION":
								return nil
							}
							return diag.Errorf("must be either 'IDP_ADAPTER' or 'IDP_CONNECTION' not %s", v)
						},
					},
					"source_ref": resourceLinkSchema(),
				},
			},
		},
	}
}

func resourcePingFederateAuthenticationPoliciesResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//svc := m.(pfClient).AuthenticationPolicyContracts
	//input := authenticationPolicyContracts.CreateAuthenticationPolicyContractInput{
	//	Body: *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
	//}
	//result, _, err := svc.CreateAuthenticationPolicyContract(&input)
	//if err != nil {
	//	return diag.Errorf("unable to create AuthenticationPolicyContracts: %s", err)
	//}
	//d.SetId(*result.Id)
	//return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//svc := m.(pfClient).AuthenticationPolicyContracts
	//input := authenticationPolicyContracts.GetAuthenticationPolicyContractInput{
	//	Id: d.Id(),
	//}
	//result, _, err := svc.GetAuthenticationPolicyContract(&input)
	//if err != nil {
	//	return diag.Errorf("unable to read AuthenticationPolicyContracts: %s", err)
	//}
	//return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//svc := m.(pfClient).AuthenticationPolicyContracts
	//input := authenticationPolicyContracts.UpdateAuthenticationPolicyContractInput{
	//	Id:   d.Id(),
	//	Body: *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
	//}
	//result, _, err := svc.UpdateAuthenticationPolicyContract(&input)
	//if err != nil {
	//	return diag.Errorf("unable to update AuthenticationPolicyContracts: %s", err)
	//}
	//
	//return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//svc := m.(pfClient).AuthenticationPolicyContracts
	//input := authenticationPolicyContracts.DeleteAuthenticationPolicyContractInput{
	//	Id: d.Id(),
	//}
	//_, _, err := svc.DeleteAuthenticationPolicyContract(&input)
	//if err != nil {
	//	return diag.Errorf("unable to delete AuthenticationPolicyContracts: %s", err)
	//}
	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPolicyContract) diag.Diagnostics {
	//var diags diag.Diagnostics
	//setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	//if rv.ExtendedAttributes != nil && len(*rv.ExtendedAttributes) > 0 {
	//	if err := d.Set("extended_attributes", flattenAuthenticationPolicyContractAttribute(*rv.ExtendedAttributes)); err != nil {
	//		diags = append(diags, diag.FromErr(err)...)
	//
	//	}
	//}
	//if rv.CoreAttributes != nil && len(*rv.CoreAttributes) > 0 {
	//	if err := d.Set("core_attributes", flattenAuthenticationPolicyContractAttribute(*rv.CoreAttributes)); err != nil {
	//		diags = append(diags, diag.FromErr(err)...)
	//
	//	}
	//}

	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceReadData(d *schema.ResourceData) *pf.AuthenticationPolicyContract {
	//core := expandAuthenticationPolicyContractAttribute(d.Get("core_attributes").(*schema.Set).List())
	//contract := &pf.AuthenticationPolicyContract{
	//	Name:           String(d.Get("name").(string)),
	//	CoreAttributes: core,
	//}
	//
	//if _, ok := d.GetOk("extended_attributes"); ok {
	//	contract.ExtendedAttributes = expandAuthenticationPolicyContractAttribute(d.Get("extended_attributes").(*schema.Set).List())
	//}
	//
	//return contract
	return nil
}
