package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederatePasswordCredentialValidatorResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederatePasswordCredentialValidatorResourceCreate,
		Read:   resourcePingFederatePasswordCredentialValidatorResourceRead,
		Update: resourcePingFederatePasswordCredentialValidatorResourceUpdate,
		Delete: resourcePingFederatePasswordCredentialValidatorResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederatePasswordCredentialValidatorResourceSchema(),
	}
}

func resourcePingFederatePasswordCredentialValidatorResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourceLinkSchema(),
		"parent_ref":            resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
		"attribute_contract": &schema.Schema{
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem:     resourcePasswordCredentialValidatorAttributeContract(),
		},
	}
}

func resourcePingFederatePasswordCredentialValidatorResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).PasswordCredentialValidators
	input := pf.CreatePasswordCredentialValidatorInput{
		Body: *resourcePingFederatePasswordCredentialValidatorResourceReadData(d),
	}
	input.Body.Id = input.Body.Name
	result, _, err := svc.CreatePasswordCredentialValidator(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederatePasswordCredentialValidatorResourceReadResult(d, result, svc)
}

func resourcePingFederatePasswordCredentialValidatorResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).PasswordCredentialValidators
	input := pf.GetPasswordCredentialValidatorInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetPasswordCredentialValidator(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourcePingFederatePasswordCredentialValidatorResourceReadResult(d, result, svc)
}

func resourcePingFederatePasswordCredentialValidatorResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).PasswordCredentialValidators
	input := pf.UpdatePasswordCredentialValidatorInput{
		Id:   d.Id(),
		Body: *resourcePingFederatePasswordCredentialValidatorResourceReadData(d),
	}
	result, _, err := svc.UpdatePasswordCredentialValidator(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return resourcePingFederatePasswordCredentialValidatorResourceReadResult(d, result, svc)
}

func resourcePingFederatePasswordCredentialValidatorResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).PasswordCredentialValidators
	input := pf.DeletePasswordCredentialValidatorInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeletePasswordCredentialValidator(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func resourcePingFederatePasswordCredentialValidatorResourceReadResult(d *schema.ResourceData, rv *pf.PasswordCredentialValidator, svc *pf.PasswordCredentialValidatorsService) (err error) {
	desc, _, err := svc.GetPasswordCredentialValidatorDescriptor(&pf.GetPasswordCredentialValidatorDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return err
	}

	setResourceDataString(d, "name", rv.Name)
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
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			return err
		}
		//if err = d.Set("configuration", flattenPluginConfiguration(rv.Configuration)); err != nil {
		//	return err
		//}
	}
	if rv.AttributeContract != nil {
		if err = d.Set("attribute_contract", flattenPasswordCredentialValidatorAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}
	// if rv.ExtendedAttributes != nil && len(*rv.ExtendedAttributes) > 0 {
	// 	if err = d.Set("extended_attributes", flattenPasswordCredentialValidatorAttribute(*rv.ExtendedAttributes)); err != nil {
	// 		return err
	// 	}
	// }
	// if rv.CoreAttributes != nil && len(*rv.CoreAttributes) > 0 {
	// 	if err = d.Set("core_attributes", flattenPasswordCredentialValidatorAttribute(*rv.CoreAttributes)); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func resourcePingFederatePasswordCredentialValidatorResourceReadData(d *schema.ResourceData) *pf.PasswordCredentialValidator {
	validator := &pf.PasswordCredentialValidator{
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("parent_ref"); ok {
		validator.ParentRef = expandResourceLink(v.([]interface{}))
	}

	if v, ok := d.GetOk("attribute_contract"); ok {
		validator.AttributeContract = expandPasswordCredentialValidatorAttributeContract(v.([]interface{}))
	}

	return validator
}
