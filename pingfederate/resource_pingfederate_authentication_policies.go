package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"
)

func resourcePingFederateAuthenticationPoliciesResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages the PingFederate instance authentication policy tree.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply clears all policies.`,
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
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Fail if policy finds no authentication source.",
		},
		"tracked_http_parameters": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The HTTP request parameters to track and make available to authentication sources, selectors, and contract mappings throughout the authentication policy.",
		},
		"authn_selection_trees": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The list of authentication policy trees.",
			Elem:        resourceAuthenticationPolicyTreeResource(),
		},
		"default_authentication_sources": {
			Type:        schema.TypeSet,
			Description: "The default authentication sources.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"type": {
						Type:     schema.TypeString,
						Required: true,
						ValidateFunc: validation.StringInSlice([]string{
							"IDP_ADAPTER",
							"IDP_CONNECTION",
						}, false),
					},
					"source_ref": resourceLinkSchema(),
				},
			},
		},
	}
}

func resourcePingFederateAuthenticationPoliciesResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{
		Body:                     *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateDefaultAuthenticationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create AuthenticationPolicies: %s", err)
	}
	d.SetId("default_authentication_policies")
	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	result, _, err := svc.GetDefaultAuthenticationPolicyWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read AuthenticationPolicies: %s", err)
	}
	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{
		Body:                     *resourcePingFederateAuthenticationPoliciesResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateDefaultAuthenticationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationPolicies: %s", err)
	}

	return resourcePingFederateAuthenticationPoliciesResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesResourceDelete(ctx context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateDefaultAuthenticationPolicyInput{
		Body: pf.AuthenticationPolicy{},
	}
	_, _, err := svc.UpdateDefaultAuthenticationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to reset AuthenticationPolicies: %s", err)
	}
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
		AuthnSelectionTrees:          expandAuthenticationPolicyTrees(d.Get("authn_selection_trees").([]interface{})),
	}
	return policy
}
