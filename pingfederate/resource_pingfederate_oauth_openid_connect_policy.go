package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthOpenIdConnect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOpenIdConnectPolicyResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateOpenIdConnectPolicyResourceCreate,
		ReadContext:   resourcePingFederateOpenIdConnectPolicyResourceRead,
		UpdateContext: resourcePingFederateOpenIdConnectPolicyResourceUpdate,
		DeleteContext: resourcePingFederateOpenIdConnectPolicyResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateOpenIdConnectPolicyResourceSchema(),
		Timeouts: &schema.ResourceTimeout{
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourcePingFederateOpenIdConnectPolicyResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"policy_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
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
		"attribute_contract": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem:     resourceOpenIdConnectAttributeContract(),
		},
		"attribute_mapping": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceAttributeMapping(),
		},
		"scope_attribute_mappings": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     resourceParameterValues(),
		},
	}
}

func resourcePingFederateOpenIdConnectPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthOpenIdConnect
	input := oauthOpenIdConnect.CreatePolicyInput{
		Body:                     *resourcePingFederateOpenIdConnectPolicyResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.CreatePolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create OauthOpenIdConnectPolicy: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthOpenIdConnect
	input := oauthOpenIdConnect.GetPolicyInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OauthOpenIdConnectPolicy: %s", err)
	}
	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthOpenIdConnect
	input := oauthOpenIdConnect.UpdatePolicyInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateOpenIdConnectPolicyResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdatePolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OauthOpenIdConnectPolicy: %s", err)
	}

	return resourcePingFederateOpenIdConnectPolicyResourceReadResult(d, result)
}

func resourcePingFederateOpenIdConnectPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthOpenIdConnect
	input := oauthOpenIdConnect.DeletePolicyInput{
		Id: d.Id(),
	}
	err := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {

		_, resp, err := svc.DeletePolicyWithContext(ctx, &input)
		if resp != nil && resp.StatusCode == http.StatusUnprocessableEntity {
			return resource.RetryableError(fmt.Errorf("unable to delete with retry OauthOpenIdConnectPolicy: %s", err))
		}
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("unable to delete OauthOpenIdConnectPolicy: %s", err))
		}
		return nil
	})
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourcePingFederateOpenIdConnectPolicyResourceReadResult(d *schema.ResourceData, rv *pf.OpenIdConnectPolicy) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "policy_id", rv.Id, &diags)
	if rv.AccessTokenManagerRef != nil {
		if err := d.Set("access_token_manager_ref", flattenResourceLink(rv.AccessTokenManagerRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)

		}
	}
	setResourceDataIntWithDiagnostic(d, "id_token_lifetime", rv.IdTokenLifetime, &diags)
	setResourceDataBoolWithDiagnostic(d, "include_sri_in_id_token", rv.IncludeSriInIdToken, &diags)
	setResourceDataBoolWithDiagnostic(d, "include_user_in_id_token", rv.IncludeUserInfoInIdToken, &diags)
	setResourceDataBoolWithDiagnostic(d, "include_shash_in_id_token", rv.IncludeSHashInIdToken, &diags)
	setResourceDataBoolWithDiagnostic(d, "return_id_token_on_refresh_grant", rv.ReturnIdTokenOnRefreshGrant, &diags)
	if rv.AttributeContract != nil && openIdAttributeContractShouldFlatten(rv.AttributeContract) {
		if err := d.Set("attribute_contract", flattenOpenIdConnectAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.AttributeMapping != nil {
		if err := d.Set("attribute_mapping", flattenAttributeMapping(rv.AttributeMapping)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.ScopeAttributeMappings != nil {
		if err := d.Set("scope_attribute_mappings", flattenMapOfParameterValues(rv.ScopeAttributeMappings)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateOpenIdConnectPolicyResourceReadData(d *schema.ResourceData) *pf.OpenIdConnectPolicy {
	policy := &pf.OpenIdConnectPolicy{
		Id:                    String(d.Get("policy_id").(string)),
		Name:                  String(d.Get("name").(string)),
		AccessTokenManagerRef: expandResourceLink(d.Get("access_token_manager_ref").([]interface{})[0].(map[string]interface{})),
		AttributeMapping:      expandAttributeMapping(d.Get("attribute_mapping").([]interface{})),
	}

	if v, ok := d.GetOk("attribute_contract"); ok {
		policy.AttributeContract = expandOpenIdConnectAttributeContract(v.([]interface{}))
	}
	if v, ok := d.GetOk("id_token_lifetime"); ok {
		policy.IdTokenLifetime = Int(v.(int))
	}
	if v, ok := d.GetOkExists("include_sri_in_id_token"); ok {
		policy.IncludeSriInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("include_user_in_id_token"); ok {
		policy.IncludeUserInfoInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("include_shash_in_id_token"); ok {
		policy.IncludeSHashInIdToken = Bool(v.(bool))
	}
	if v, ok := d.GetOkExists("return_id_token_on_refresh_grant"); ok {
		policy.ReturnIdTokenOnRefreshGrant = Bool(v.(bool))
	}
	if v, ok := d.GetOk("scope_attribute_mappings"); ok {
		policy.ScopeAttributeMappings = expandMapOfParameterValues(v.(*schema.Set).List())
	}
	return policy
}
