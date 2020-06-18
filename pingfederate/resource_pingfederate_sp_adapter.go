package pingfederate

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingfederate-sdk-go/services/spAdapters"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateSpAdapterResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateSpAdapterResourceCreate,
		ReadContext:   resourcePingFederateSpAdapterResourceRead,
		UpdateContext: resourcePingFederateSpAdapterResourceUpdate,
		DeleteContext: resourcePingFederateSpAdapterResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateSpAdapterResourceSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(pfClient).SpAdapters
			if className, ok := d.GetOk("plugin_descriptor_ref.0.id"); ok {
				desc, _, err := svc.GetSpAdapterDescriptorsById(&spAdapters.GetSpAdapterDescriptorsByIdInput{Id: className.(string)})
				if err != nil {
					descs, _, err := svc.GetSpAdapterDescriptors()
					if err == nil && descs != nil {
						list := func(in *[]*pf.SpAdapterDescriptor) string {
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

func resourcePingFederateSpAdapterResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"sp_adapter_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourceRequiredLinkSchema(),
		"parent_ref":            resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
		"attribute_contract": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem:     resourceSpAdapterAttributeContract(),
		},
		"target_application_info": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceSpAdapterTargetApplicationInfo(),
		},
	}
}

func resourcePingFederateSpAdapterResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAdapters
	input := spAdapters.CreateSpAdapterInput{
		Body: *resourcePingFederateSpAdapterResourceReadData(d),
	}
	result, _, err := svc.CreateSpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to create SpAdapters: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAdapters
	input := spAdapters.GetSpAdapterInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to read SpAdapters: %s", err)
	}
	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAdapters
	input := spAdapters.UpdateSpAdapterInput{
		Id:   d.Id(),
		Body: *resourcePingFederateSpAdapterResourceReadData(d),
	}
	result, _, err := svc.UpdateSpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to update SpAdapters: %s", err)
	}

	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).SpAdapters
	input := spAdapters.DeleteSpAdapterInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteSpAdapter(&input)
	if err != nil {
		return diag.Errorf("unable to delete SpAdapters: %s", err)
	}
	return nil
}

func resourcePingFederateSpAdapterResourceReadResult(d *schema.ResourceData, rv *pf.SpAdapter, svc spAdapters.SpAdaptersAPI) diag.Diagnostics {
	desc, _, err := svc.GetSpAdapterDescriptorsById(&spAdapters.GetSpAdapterDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve SpAdapters descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringithDiagnostic(d, "sp_adapter_id", rv.Id, &diags)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.ParentRef != nil {
		if err = d.Set("parent_ref", flattenResourceLink(rv.ParentRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.Configuration != nil {
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AttributeContract != nil {
		if err = d.Set("attribute_contract", flattenSpAdapterAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.TargetApplicationInfo != nil {
		if err = d.Set("target_application_info", flattenSpAdapterTargetApplicationInfo(rv.TargetApplicationInfo)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateSpAdapterResourceReadData(d *schema.ResourceData) *pf.SpAdapter {
	validator := &pf.SpAdapter{
		Id:                  String(d.Get("sp_adapter_id").(string)),
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("parent_ref"); ok {
		validator.ParentRef = expandResourceLink(v.([]interface{}))
	}
	if v, ok := d.GetOk("attribute_contract"); ok {
		validator.AttributeContract = expandSpAdapterAttributeContract(v.([]interface{}))
	}
	if v, ok := d.GetOk("target_application_info"); ok {
		validator.TargetApplicationInfo = expandSpAdapterTargetApplicationInfo(v.([]interface{}))
	}

	return validator
}
