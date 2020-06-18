package pingfederate

import (
	"context"
	"fmt"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateIdpAdapterResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateIdpAdapterResourceCreate,
		ReadContext:   resourcePingFederateIdpAdapterResourceRead,
		UpdateContext: resourcePingFederateIdpAdapterResourceUpdate,
		DeleteContext: resourcePingFederateIdpAdapterResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateIdpAdapterResourceSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(pfClient).IdpAdapters
			if className, ok := d.GetOk("plugin_descriptor_ref.0.id"); ok {
				desc, _, err := svc.GetIdpAdapterDescriptorsById(&idpAdapters.GetIdpAdapterDescriptorsByIdInput{Id: className.(string)})
				if err != nil {
					descs, _, err := svc.GetIdpAdapterDescriptors()
					if err == nil && descs != nil {
						list := func(in *[]*pf.IdpAdapterDescriptor) string {
							var plugins []string
							for _, descriptor := range *in {
								plugins = append(plugins, *descriptor.ClassName)
							}
							return strings.Join(plugins, "\n\t")
						}
						return fmt.Errorf("unable to find plugin_descriptor for %s available plugins:\n\t%s", className.(string), list(descs.Items))
					}
					return fmt.Errorf("unable to find plugin_descriptor for %s", className.(string))
				}
				return validateConfiguration(d, desc.ConfigDescriptor)
			}
			return nil
		},
	}
}

func resourcePingFederateIdpAdapterResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bypass_external_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "External validation will be bypassed when set to true. Default to false.",
			Default:     false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourceRequiredLinkSchema(),
		"parent_ref":            resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
		"authn_ctx_class_ref": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"attribute_mapping": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceIdpAdapterAttributeMapping(),
		},
		"attribute_contract": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem:     resourceIdpAdapterAttributeContract(),
		},
	}
}

func resourcePingFederateIdpAdapterResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.CreateIdpAdapterInput{
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	input.Body.Id = input.Body.Name
	result, _, err := svc.CreateIdpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to create IdpAdapters: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.GetIdpAdapterInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetIdpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to read IdpAdapters: %s", err)
	}
	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.UpdateIdpAdapterInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.UpdateIdpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to update IdpAdapters: %s", err)
	}

	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.DeleteIdpAdapterInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteIdpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to delete IdpAdapters: %s", err)
	}
	return nil
}

func resourcePingFederateIdpAdapterResourceReadResult(d *schema.ResourceData, rv *pf.IdpAdapter) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringithDiagnostic(d, "authn_ctx_class_ref", rv.AuthnCtxClassRef, &diags)
	if rv.PluginDescriptorRef != nil {
		if err := d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.ParentRef != nil {
		if err := d.Set("parent_ref", flattenResourceLink(rv.ParentRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.Configuration != nil {
		if err := d.Set("configuration", flattenPluginConfiguration(rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AttributeContract != nil {
		if err := d.Set("attribute_contract", flattenIdpAdapterAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AttributeMapping != nil {
		if err := d.Set("attribute_mapping", flattenIdpAdapterContractMapping(rv.AttributeMapping)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateIdpAdapterResourceReadData(d *schema.ResourceData) *pf.IdpAdapter {
	validator := &pf.IdpAdapter{
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("parent_ref"); ok {
		validator.ParentRef = expandResourceLink(v.([]interface{}))
	}
	if v, ok := d.GetOk("authn_ctx_class_ref"); ok {
		validator.AuthnCtxClassRef = String(v.(string))
	}
	if v, ok := d.GetOk("attribute_contract"); ok {
		validator.AttributeContract = expandIdpAdapterAttributeContract(v.([]interface{}))
	}
	if v, ok := d.GetOk("attribute_mapping"); ok {
		validator.AttributeMapping = expandIdpAdapterContractMapping(v.([]interface{}))
	}

	return validator
}
