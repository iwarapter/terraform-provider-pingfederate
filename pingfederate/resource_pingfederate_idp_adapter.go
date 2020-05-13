package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateIdpAdapterResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateIdpAdapterResourceCreate,
		Read:   resourcePingFederateIdpAdapterResourceRead,
		Update: resourcePingFederateIdpAdapterResourceUpdate,
		Delete: resourcePingFederateIdpAdapterResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingFederateIdpAdapterResourceSchema(),
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

func resourcePingFederateIdpAdapterResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).IdpAdapters
	input := pf.CreateIdpAdapterInput{
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	input.Body.Id = input.Body.Name
	result, _, err := svc.CreateIdpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).IdpAdapters
	input := pf.GetIdpAdapterInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetIdpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).IdpAdapters
	input := pf.UpdateIdpAdapterInput{
		Id:                       d.Id(),
		Body:                     *resourcePingFederateIdpAdapterResourceReadData(d),
		BypassExternalValidation: Bool(d.Get("bypass_external_validation").(bool)),
	}
	result, _, err := svc.UpdateIdpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return resourcePingFederateIdpAdapterResourceReadResult(d, result)
}

func resourcePingFederateIdpAdapterResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).IdpAdapters
	input := pf.DeleteIdpAdapterInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteIdpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func resourcePingFederateIdpAdapterResourceReadResult(d *schema.ResourceData, rv *pf.IdpAdapter) (err error) {
	setResourceDataString(d, "name", rv.Name)
	setResourceDataString(d, "authn_ctx_class_ref", rv.AuthnCtxClassRef)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			return err
		}
	}

	if rv.ParentRef != nil {
		if err = d.Set("parent_ref", flattenResourceLink(rv.ParentRef)); err != nil {
			return err
		}
	}

	if rv.Configuration != nil {
		if err = d.Set("configuration", flattenPluginConfiguration(rv.Configuration)); err != nil {
			return err
		}
	}

	if rv.AttributeContract != nil {
		if err = d.Set("attribute_contract", flattenIdpAdapterAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}

	if rv.AttributeMapping != nil {
		if err = d.Set("attribute_mapping", flattenIdpAdapterContractMapping(rv.AttributeMapping)); err != nil {
			return err
		}
	}
	return nil
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
