package pingfederate

import (
	"context"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"
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
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{
		Body: *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
	}
	result, _, err := svc.UpdateDefaultAuthenticationPolicy(&input)
	if err != nil {
		return diag.Errorf("unable to create AuthenticationPolicies: %s", err)
	}
	d.SetId("default_authentication_policies")
	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	result, _, err := svc.GetDefaultAuthenticationPolicy()
	if err != nil {
		return diag.Errorf("unable to read AuthenticationPolicies: %s", err)
	}
	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{
		Body: *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
	}
	result, _, err := svc.UpdateDefaultAuthenticationPolicy(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationPolicies: %s", err)
	}

	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPolicy) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "fail_if_no_selection", rv.FailIfNoSelection, &diags)

	if rv.DefaultAuthenticationSources != nil {
		if err := d.Set("default_authentication_sources", flattenAuthenticationSources(*rv.DefaultAuthenticationSources)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.TrackedHttpParameters != nil {
		if err := d.Set("tracked_http_parameters", flattenStringList(*rv.TrackedHttpParameters)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AuthnSelectionTrees != nil {
		if err := d.Set("authn_selection_trees", flattenAuthenticationPolicyTrees(*rv.AuthnSelectionTrees)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return nil
}

func resourcePingFederateAuthenticationPoliciesResourceReadData(d *schema.ResourceData) *pf.AuthenticationPolicy {
	strs := expandStringList(d.Get("tracked_http_parameters").(*schema.Set).List())
	policy := &pf.AuthenticationPolicy{
		FailIfNoSelection:            Bool(d.Get("fail_if_no_selection").(bool)),
		TrackedHttpParameters:        &strs,
		DefaultAuthenticationSources: expandAuthenticationSources(d.Get("default_authentication_sources").(*schema.Set).List()),
		AuthnSelectionTrees:          expandAuthenticationPolicyTrees(d.Get("authn_selection_trees").(*schema.Set).List()),
	}
	return policy
}
