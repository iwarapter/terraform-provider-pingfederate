package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"
)

func resourcePingFederateAuthenticationPoliciesSettingsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthenticationPoliciesSettingsResourceCreate,
		ReadContext:   resourcePingFederateAuthenticationPoliciesSettingsResourceRead,
		UpdateContext: resourcePingFederateAuthenticationPoliciesSettingsResourceUpdate,
		DeleteContext: resourcePingFederateAuthenticationPoliciesSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthenticationPoliciesSettingsResourceSchema(),
	}
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_idp_authn_selection": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_sp_authn_selection": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateSettingsInput{
		Body: pf.AuthenticationPoliciesSettings{
			EnableIdpAuthnSelection: Bool(d.Get("enable_idp_authn_selection").(bool)),
			EnableSpAuthnSelection:  Bool(d.Get("enable_sp_authn_selection").(bool)),
		},
	}
	result, _, err := svc.UpdateSettings(&input)
	if err != nil {
		return diag.Errorf("unable to set AuthenticationPoliciesSettings: %s", err)
	}
	d.SetId("default_authentication_policies_settings")
	return resourcePingFederateAuthenticationPoliciesSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	result, _, err := svc.GetSettings()
	if err != nil {
		return diag.Errorf("unable to read AuthenticationPoliciesSettings: %s", err)
	}
	return resourcePingFederateAuthenticationPoliciesSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationPolicies
	input := authenticationPolicies.UpdateSettingsInput{
		Body: pf.AuthenticationPoliciesSettings{
			EnableIdpAuthnSelection: Bool(d.Get("enable_idp_authn_selection").(bool)),
			EnableSpAuthnSelection:  Bool(d.Get("enable_sp_authn_selection").(bool)),
		},
	}
	result, _, err := svc.UpdateSettings(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationPolicies: %s", err)
	}

	return resourcePingFederateAuthenticationPoliciesSettingsResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateAuthenticationPoliciesSettingsResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPoliciesSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataBoolWithDiagnostic(d, "enable_idp_authn_selection", rv.EnableIdpAuthnSelection, &diags)
	setResourceDataBoolWithDiagnostic(d, "enable_sp_authn_selection", rv.EnableSpAuthnSelection, &diags)
	return nil
}
