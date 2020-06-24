package pingfederate

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateAuthenticationPolicyContractResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthenticationPolicyContractResourceCreate,
		ReadContext:   resourcePingFederateAuthenticationPolicyContractResourceRead,
		UpdateContext: resourcePingFederateAuthenticationPolicyContractResourceUpdate,
		DeleteContext: resourcePingFederateAuthenticationPolicyContractResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateAuthenticationPolicyContractResourceSchema(),
	}
}

func resourcePingFederateAuthenticationPolicyContractResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"core_attributes": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"extended_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourcePingFederateAuthenticationPolicyContractResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicyContracts
	input := authenticationPolicyContracts.CreateAuthenticationPolicyContractInput{
		Body: *resourcePingFederateAuthenticationPolicyContractResourceReadData(d),
	}
	result, _, err := svc.CreateAuthenticationPolicyContract(&input)
	if err != nil {
		return diag.Errorf("unable to create AuthenticationPolicyContracts: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicyContracts
	input := authenticationPolicyContracts.GetAuthenticationPolicyContractInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthenticationPolicyContract(&input)
	if err != nil {
		return diag.Errorf("unable to read AuthenticationPolicyContracts: %s", err)
	}
	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicyContracts
	input := authenticationPolicyContracts.UpdateAuthenticationPolicyContractInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthenticationPolicyContractResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationPolicyContract(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationPolicyContracts: %s", err)
	}

	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicyContracts
	input := authenticationPolicyContracts.DeleteAuthenticationPolicyContractInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteAuthenticationPolicyContract(&input)
	if err != nil {
		return diag.Errorf("unable to delete AuthenticationPolicyContracts: %s", err)
	}
	return nil
}

func resourcePingFederateAuthenticationPolicyContractResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPolicyContract) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	if rv.ExtendedAttributes != nil && len(*rv.ExtendedAttributes) > 0 {
		if err := d.Set("extended_attributes", flattenAuthenticationPolicyContractAttribute(*rv.ExtendedAttributes)); err != nil {
			diags = append(diags, diag.FromErr(err)...)

		}
	}
	if rv.CoreAttributes != nil && len(*rv.CoreAttributes) > 0 {
		if err := d.Set("core_attributes", flattenAuthenticationPolicyContractAttribute(*rv.CoreAttributes)); err != nil {
			diags = append(diags, diag.FromErr(err)...)

		}
	}

	return nil
}

func resourcePingFederateAuthenticationPolicyContractResourceReadData(d *schema.ResourceData) *pf.AuthenticationPolicyContract {
	contract := &pf.AuthenticationPolicyContract{
		Name: String(d.Get("name").(string)),
		CoreAttributes: &[]*pf.AuthenticationPolicyContractAttribute{
			{
				Name: String("subject"),
			},
		},
	}

	if _, ok := d.GetOk("extended_attributes"); ok {
		contract.ExtendedAttributes = expandAuthenticationPolicyContractAttribute(d.Get("extended_attributes").(*schema.Set).List())
	}

	return contract
}
