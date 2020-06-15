package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateAuthenticationSelectorResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateAuthenticationSelectorResourceCreate,
		Read:   resourcePingFederateAuthenticationSelectorResourceRead,
		Update: resourcePingFederateAuthenticationSelectorResourceUpdate,
		Delete: resourcePingFederateAuthenticationSelectorResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingFederateAuthenticationSelectorResourceSchema(),
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

func resourcePingFederateAuthenticationSelectorResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationSelectors
	input := pf.CreateAuthenticationSelectorInput{
		Body: *resourcePingFederateAuthenticationSelectorResourceReadData(d),
	}
	result, _, err := svc.CreateAuthenticationSelector(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationSelectors
	input := pf.GetAuthenticationSelectorInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthenticationSelector(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationSelectors
	input := pf.UpdateAuthenticationSelectorInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthenticationSelectorResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationSelector(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return resourcePingFederateAuthenticationSelectorResourceReadResult(d, result, svc)
}

func resourcePingFederateAuthenticationSelectorResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationSelectors
	input := pf.DeleteAuthenticationSelectorInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteAuthenticationSelector(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateAuthenticationSelectorResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationSelector, svc *pf.AuthenticationSelectorsService) (err error) {
	desc, _, err := svc.GetAuthenticationSelectorDescriptorsById(&pf.GetAuthenticationSelectorDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return err
	}

	setResourceDataString(d, "name", rv.Name)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			return err
		}
	}
	if rv.AttributeContract != nil && rv.AttributeContract.ExtendedAttributes != nil && len(*rv.AttributeContract.ExtendedAttributes) > 0 {
		if err = d.Set("extended_attributes", flattenAuthenticationSelectorAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}
	if rv.Configuration != nil {
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			return err
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
