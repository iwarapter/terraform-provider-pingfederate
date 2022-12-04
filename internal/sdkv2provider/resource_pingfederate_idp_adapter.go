package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/go-cty/cty"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateIdpAdapterResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for IDP Adapters within PingFederate.",
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
				desc, resp, err := svc.GetIdpAdapterDescriptorsByIdWithContext(ctx, &idpAdapters.GetIdpAdapterDescriptorsByIdInput{Id: className.(string)})
				if resp != nil && resp.StatusCode == http.StatusForbidden {
					log.Printf("[WARN] Unable to query IdpAdapterDescriptor, IdP role not enabled")
					return nil
				}
				if err != nil {
					descs, _, err := svc.GetIdpAdapterDescriptorsWithContext(ctx)
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
		"instance_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The ID of the plugin instance. The ID cannot be modified once the instance is created. Note: Ignored when specifying a connection's adapter override.",
			ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
				v := value.(string)
				r, _ := regexp.Compile(`^[a-zA-Z0-9._-]+$`)
				if !r.MatchString(v) || len(v) >= 33 {
					return diag.Errorf("The plugin ID must be less than 33 characters, contain no spaces, and be alphanumeric.")
				}
				return nil
			},
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The plugin instance name. The name cannot be modified once the instance is created. Note: Ignored when specifying a connection's adapter override.",
		},
		"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
		"parent_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances. Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)",
			Elem:        resourceLinkResource(),
		},
		"configuration": resourcePluginConfiguration(),
		"authn_ctx_class_ref": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The fixed value that indicates how the user was authenticated.",
		},
		"attribute_mapping": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The attributes mapping from attribute sources to attribute targets.",
			Elem:        resourceIdpAdapterAttributeMapping(),
		},
		"attribute_contract": {
			Type:        schema.TypeList,
			Required:    true,
			MaxItems:    1,
			Description: "The list of attributes that the IdP adapter provides.",
			Elem:        resourceIdpAdapterAttributeContract(),
		},
	}
}

func resourcePingFederateIdpAdapterResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.CreateIdpAdapterInput{
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	input.Body.Id = input.Body.Name
	result, _, err := svc.CreateIdpAdapterWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create IdpAdapters: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpAdapterResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.GetIdpAdapterInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetIdpAdapterWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read IdpAdapters: %s", err)
	}
	return resourcePingFederateIdpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpAdapterResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.UpdateIdpAdapterInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(m.(pfClient).BypassExternalValidation),
	}
	result, _, err := svc.UpdateIdpAdapterWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update IdpAdapters: %s", err)
	}

	return resourcePingFederateIdpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpAdapterResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpAdapters
	input := idpAdapters.DeleteIdpAdapterInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteIdpAdapterWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete IdpAdapters: %s", err)
	}
	return nil
}

func resourcePingFederateIdpAdapterResourceReadResult(d *schema.ResourceData, rv *pf.IdpAdapter, svc idpAdapters.IdpAdaptersAPI) diag.Diagnostics {
	desc, _, err := svc.GetIdpAdapterDescriptorsById(&idpAdapters.GetIdpAdapterDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve IdpAdapters descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "authn_ctx_class_ref", rv.AuthnCtxClassRef, &diags)
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
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
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
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})[0].(map[string]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("parent_ref"); ok && len(v.([]interface{})) > 0 {
		validator.ParentRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("authn_ctx_class_ref"); ok {
		validator.AuthnCtxClassRef = String(v.(string))
	}
	if v, ok := d.GetOk("attribute_contract"); ok && len(v.([]interface{})) > 0 {
		validator.AttributeContract = expandIdpAdapterAttributeContract(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("attribute_mapping"); ok && len(v.([]interface{})) > 0 {
		validator.AttributeMapping = expandIdpAdapterContractMapping(v.([]interface{})[0].(map[string]interface{}))
	}

	return validator
}
