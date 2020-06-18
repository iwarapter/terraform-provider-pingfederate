package pingfederate

import (
	"context"
	"fmt"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationSelectors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateAuthenticationSelectorResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateAuthenticationSelectorResourceCreate,
		ReadContext:   resourcePingFederateAuthenticationSelectorResourceRead,
		UpdateContext: resourcePingFederateAuthenticationSelectorResourceUpdate,
		DeleteContext: resourcePingFederateAuthenticationSelectorResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateAuthenticationSelectorResourceSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(pfClient).AuthenticationSelectors
			if className, ok := d.GetOk("plugin_descriptor_ref.0.id"); ok {
				desc, _, err := svc.GetAuthenticationSelectorDescriptorsById(&authenticationSelectors.GetAuthenticationSelectorDescriptorsByIdInput{Id: className.(string)})
				if err != nil {
					descs, _, err := svc.GetAuthenticationSelectorDescriptors()
					if err == nil && descs != nil {
						list := func(in *[]*pf.AuthenticationSelectorDescriptor) string {
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

func resourcePingFederateAuthenticationSelectorResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
		"attribute_contract": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem:     resourceAuthenticationSelectorAttributeContract(),
		},
	}
}

func resourcePingFederateAuthenticationSelectorResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationSelectors
	input := authenticationSelectors.CreateAuthenticationSelectorInput{
		Body: *resourcePingFederateAuthenticationSelectorResourceReadData(d),
	}
	result, _, err := svc.CreateAuthenticationSelector(&input)
	if err != nil {
		return diag.Errorf("unable to create AuthenticationSelectors: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationSelectors
	input := authenticationSelectors.GetAuthenticationSelectorInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthenticationSelector(&input)
	if err != nil {
		return diag.Errorf("unable to read AuthenticationSelectors: %s", err)
	}
	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationSelectors
	input := authenticationSelectors.UpdateAuthenticationSelectorInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthenticationSelectorResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationSelector(&input)
	if err != nil {
		return diag.Errorf("unable to update AuthenticationSelectors: %s", err)
	}

	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).AuthenticationSelectors
	input := authenticationSelectors.DeleteAuthenticationSelectorInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteAuthenticationSelector(&input)
	if err != nil {
		return diag.Errorf("unable to delete AuthenticationSelectors: %s", err)
	}
	return nil
}

func resourcePingFederateAuthenticationSelectorResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationSelector, svc authenticationSelectors.AuthenticationSelectorsAPI) diag.Diagnostics {
	desc, _, err := svc.GetAuthenticationSelectorDescriptorsById(&authenticationSelectors.GetAuthenticationSelectorDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve AuthenticationSelectors descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.AttributeContract != nil && rv.AttributeContract.ExtendedAttributes != nil && len(*rv.AttributeContract.ExtendedAttributes) > 0 {
		if err = d.Set("extended_attributes", flattenAuthenticationSelectorAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.Configuration != nil {
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	return nil
}

func resourcePingFederateAuthenticationSelectorResourceReadData(d *schema.ResourceData) *pf.AuthenticationSelector {
	selector := &pf.AuthenticationSelector{
		Name:                String(d.Get("name").(string)),
		Id:                  String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("attribute_contract"); ok {
		selector.AttributeContract = expandAuthenticationSelectorAttributeContract(v.([]interface{}))
	}

	return selector
}
