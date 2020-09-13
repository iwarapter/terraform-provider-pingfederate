package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationApi"
)

func resourcePingFederateAuthnApiApplicationResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthnApiApplicationResourceCreate,
		ReadContext:   resourcePingFederateAuthnApiApplicationResourceRead,
		UpdateContext: resourcePingFederateAuthnApiApplicationResourceUpdate,
		DeleteContext: resourcePingFederateAuthnApiApplicationResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthnApiApplicationResourceSchema(),
	}
}

func resourcePingFederateAuthnApiApplicationResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"additional_allowed_origins": setOfString(),
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourcePingFederateAuthnApiApplicationResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.CreateApplicationInput{
		Body: *resourcePingFederateAuthnApiApplicationResourceReadData(d),
	}
	result, _, err := svc.CreateApplication(&input)
	if err != nil {
		return diag.Errorf("unable to create AuthnApiApplication: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.GetApplicationInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApplication(&input)
	if err != nil {
		return diag.Errorf("unable to read AuthnApiApplication: %s", err)
	}
	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.UpdateApplicationInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthnApiApplicationResourceReadData(d),
	}
	result, _, err := svc.UpdateApplication(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthnApiApplication: %s", err)
	}

	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.DeleteApplicationInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteApplication(&input)
	if err != nil {
		return diag.Errorf("unable to delete AuthnApiApplication: %s", err)
	}
	return nil
}

func resourcePingFederateAuthnApiApplicationResourceReadResult(d *schema.ResourceData, rv *pf.AuthnApiApplication) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "app_id", rv.Id, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "url", rv.Url, &diags)
	setResourceDataStringWithDiagnostic(d, "description", rv.Description, &diags)
	if rv.AdditionalAllowedOrigins != nil {
		if err := d.Set("additional_allowed_origins", rv.AdditionalAllowedOrigins); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateAuthnApiApplicationResourceReadData(d *schema.ResourceData) *pf.AuthnApiApplication {
	result := pf.AuthnApiApplication{
		Name: String(d.Get("name").(string)),
		Url:  String(d.Get("url").(string)),
	}
	if val, ok := d.GetOk("app_id"); ok {
		result.Id = String(val.(string))
	}
	if val, ok := d.GetOk("description"); ok {
		result.Description = String(val.(string))
	}
	if val, ok := d.GetOk("additional_allowed_origins"); ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.AdditionalAllowedOrigins = &strs
	}
	return &result
}
