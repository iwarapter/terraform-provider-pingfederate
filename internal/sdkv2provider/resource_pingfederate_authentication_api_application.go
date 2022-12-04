package sdkv2provider

import (
	"context"
	"regexp"

	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationApi"
)

func resourcePingFederateAuthnApiApplicationResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides an authentication API application.",
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
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the Authentication API application. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.",
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				v := value.(string)
				r, _ := regexp.Compile(`^[a-zA-Z0-9._-]+$`)
				if !r.MatchString(v) {
					return diag.Errorf("the app_id can only contain alphanumeric characters, dash, dot and underscore.")
				}
				return nil
			},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The Authentication API Application Name. Name must be unique.",
		},
		"url": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The Authentication API Application redirect URL.",
		},
		"additional_allowed_origins": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			Description: "The domain in the redirect URL is always whitelisted. This field contains a list of additional allowed origin URL's for cross-origin resource sharing.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The Authentication API Application description.",
		},
	}
}

func resourcePingFederateAuthnApiApplicationResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.CreateApplicationInput{
		Body: *resourcePingFederateAuthnApiApplicationResourceReadData(d),
	}
	result, _, err := svc.CreateApplicationWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create AuthnApiApplication: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.GetApplicationInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetApplicationWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read AuthnApiApplication: %s", err)
	}
	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.UpdateApplicationInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthnApiApplicationResourceReadData(d),
	}
	result, _, err := svc.UpdateApplicationWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update AuthnApiApplication: %s", err)
	}

	return resourcePingFederateAuthnApiApplicationResourceReadResult(d, result)
}

func resourcePingFederateAuthnApiApplicationResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationApi
	input := authenticationApi.DeleteApplicationInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteApplicationWithContext(ctx, &input)
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
