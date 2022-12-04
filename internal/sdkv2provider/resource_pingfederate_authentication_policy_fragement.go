package sdkv2provider

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateAuthenticationPolicyFragmentResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthenticationPolicyFragmentResourceCreate,
		ReadContext:   resourcePingFederateAuthenticationPolicyFragmentResourceRead,
		UpdateContext: resourcePingFederateAuthenticationPolicyFragmentResourceUpdate,
		DeleteContext: resourcePingFederateAuthenticationPolicyFragmentResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthenticationPolicyFragmentResourceSchema(),
	}
}

func resourcePingFederateAuthenticationPolicyFragmentResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"policy_fragment_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The authentication policy fragment ID. ID is unique.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The authentication policy fragment name. Name is unique.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A description for the authentication policy.",
		},
		"inputs": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The reference to the authentication policy contract to use as the attribute inputs for this authentication policy fragment.",
			Elem:        resourceLinkResource(),
		},
		"outputs": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The reference to the authentication policy contract to use as the attribute outputs for this authentication policy fragment.",
			Elem:        resourceLinkResource(),
		},
		"root_node": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "A node inside the authentication policy tree.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"action": resourcePolicyActionSchema(),
					"children": {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "The nodes inside the authentication policy tree node.",
						Elem:        resourceAuthenticationPolicyTreeNodeSchemaBuilder(10),
					},
				},
			},
		},
	}
}

func resourcePingFederateAuthenticationPolicyFragmentResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.CreateFragmentInput{
		Body:                     *resourcePingFederateAuthenticationPolicyFragmentResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreateFragmentWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create AuthenticationPolicyFragment: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthenticationPolicyFragmentResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyFragmentResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !m.(pfClient).IsPF10_2orGreater() {
		return diag.Errorf("authentication_policy_fragment is only available from PF 10.2+")
	}
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.GetFragmentInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetFragmentWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read AuthenticationPolicyFragment: %s", err)
	}
	return resourcePingFederateAuthenticationPolicyFragmentResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyFragmentResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateFragmentInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateAuthenticationPolicyFragmentResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateFragmentWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationPolicyFragment: %s", err)
	}

	return resourcePingFederateAuthenticationPolicyFragmentResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyFragmentResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.DeleteFragmentInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteFragmentWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete AuthenticationPolicyFragment: %s", err)
	}
	return nil
}

func resourcePingFederateAuthenticationPolicyFragmentResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPolicyFragment) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "policy_fragment_id", rv.Id, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "description", rv.Description, &diags)
	if rv.Inputs != nil {
		if err := d.Set("inputs", flattenResourceLink(rv.Inputs)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Outputs != nil {
		if err := d.Set("outputs", flattenResourceLink(rv.Outputs)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.RootNode != nil {
		if err := d.Set("root_node", flattenAuthenticationPolicyTreeNodes([]*pf.AuthenticationPolicyTreeNode{rv.RootNode})); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return nil
}

func resourcePingFederateAuthenticationPolicyFragmentResourceReadData(d *schema.ResourceData) *pf.AuthenticationPolicyFragment {
	policy := &pf.AuthenticationPolicyFragment{
		Name:     String(d.Get("name").(string)),
		Inputs:   expandResourceLink(d.Get("inputs").([]interface{})[0].(map[string]interface{})),
		Outputs:  expandResourceLink(d.Get("outputs").([]interface{})[0].(map[string]interface{})),
		RootNode: expandAuthenticationPolicyTreeNode(d.Get("root_node").([]interface{})[0].(map[string]interface{})),
	}
	if v, ok := d.GetOk("description"); ok {
		policy.Description = String(v.(string))
	}
	if v, ok := d.GetOk("policy_fragment_id"); ok {
		policy.Id = String(v.(string))
	}
	return policy
}
