package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpTokenProcessors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateIdpTokenProcessorResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for IDP Token Processor within PingFederate.",
		CreateContext: resourcePingFederateIdpTokenProcessorResourceCreate,
		ReadContext:   resourcePingFederateIdpTokenProcessorResourceRead,
		UpdateContext: resourcePingFederateIdpTokenProcessorResourceUpdate,
		DeleteContext: resourcePingFederateIdpTokenProcessorResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateIdpTokenProcessorResourceSchema(),
	}
}

func resourcePingFederateIdpTokenProcessorResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"processor_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the connection. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The plugin instance name.",
		},
		"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
		"parent_ref": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances.\nNote: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)",
			Elem:        resourceLinkResource(),
		},
		"configuration": resourcePluginConfiguration(),
		"attribute_contract": {
			Type:        schema.TypeList,
			Optional:    true,
			Computed:    true,
			MaxItems:    1,
			Description: "The list of attributes that the Authentication Selector provides.",
			Elem:        resourceTokenProcessorAttributeContract(),
		},
	}
}

func resourcePingFederateIdpTokenProcessorResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpTokenProcessors
	body, err := resourcePingFederateIdpTokenProcessorResourceReadData(d, svc)
	if err != nil {
		return diag.Errorf("unable to create IdpTokenProcessors create request: %s", err)
	}
	input := idpTokenProcessors.CreateTokenProcessorInput{
		Body: *body,
	}
	result, _, err := svc.CreateTokenProcessorWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create IdpTokenProcessors: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpTokenProcessorResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpTokenProcessorResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpTokenProcessors
	input := idpTokenProcessors.GetTokenProcessorInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTokenProcessorWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read IdpTokenProcessors: %s", err)
	}
	return resourcePingFederateIdpTokenProcessorResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpTokenProcessorResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).IdpTokenProcessors
	body, err := resourcePingFederateIdpTokenProcessorResourceReadData(d, svc)
	if err != nil {
		return diag.Errorf("unable to create IdpTokenProcessors update request: %s", err)
	}
	input := idpTokenProcessors.UpdateTokenProcessorInput{
		Id:   d.Id(),
		Body: *body,
	}
	result, _, err := svc.UpdateTokenProcessorWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update IdpTokenProcessors: %s", err)
	}

	return resourcePingFederateIdpTokenProcessorResourceReadResult(d, result, svc)
}

func resourcePingFederateIdpTokenProcessorResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).IdpTokenProcessors
	input := idpTokenProcessors.DeleteTokenProcessorInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteTokenProcessorWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete IdpTokenProcessors: %s", err)
	}
	return nil
}

func resourcePingFederateIdpTokenProcessorResourceReadResult(d *schema.ResourceData, rv *pf.TokenProcessor, svc idpTokenProcessors.IdpTokenProcessorsAPI) diag.Diagnostics {
	desc, _, err := svc.GetTokenProcessorDescriptorsById(&idpTokenProcessors.GetTokenProcessorDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve IdpTokenProcessors descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "processor_id", rv.Id, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
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

		if err := d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.AttributeContract != nil {
		if err := d.Set("attribute_contract", flattenTokenProcessorAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateIdpTokenProcessorResourceReadData(d *schema.ResourceData, svc idpTokenProcessors.IdpTokenProcessorsAPI) (*pf.TokenProcessor, error) {
	processor := pf.TokenProcessor{
		Id:                  String(d.Get("processor_id").(string)),
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})[0].(map[string]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
		AttributeContract:   &pf.TokenProcessorAttributeContract{CoreAttributes: &[]*pf.TokenProcessorAttribute{}},
	}
	desc, _, err := svc.GetTokenProcessorDescriptorsById(&idpTokenProcessors.GetTokenProcessorDescriptorsByIdInput{Id: *processor.PluginDescriptorRef.Id})
	if err != nil {
		return nil, err
	}
	if v, ok := d.GetOk("parent_ref"); ok && len(v.([]interface{})) > 0 {
		processor.ParentRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := d.GetOk("attribute_contract"); ok {
		processor.AttributeContract = expandTokenProcessorAttributeContract(v.([]interface{}))
	}
	if len(*processor.AttributeContract.CoreAttributes) == 0 {
		for _, s := range *desc.AttributeContract {
			*processor.AttributeContract.CoreAttributes = append(*processor.AttributeContract.CoreAttributes, &pf.TokenProcessorAttribute{Name: String(*s)})
		}
	}

	return &processor, nil
}
