package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateSpAdapterResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateSpAdapterResourceCreate,
		Read:   resourcePingFederateSpAdapterResourceRead,
		Update: resourcePingFederateSpAdapterResourceUpdate,
		Delete: resourcePingFederateSpAdapterResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: resourcePingFederateSpAdapterResourceSchema(),
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

func resourcePingFederateSpAdapterResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).SpAdapters
	input := pf.CreateSpAdapterInput{
		Body: *resourcePingFederateSpAdapterResourceReadData(d),
	}
	result, _, err := svc.CreateSpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).SpAdapters
	input := pf.GetSpAdapterInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetSpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).SpAdapters
	input := pf.UpdateSpAdapterInput{
		Id:   d.Id(),
		Body: *resourcePingFederateSpAdapterResourceReadData(d),
	}
	result, _, err := svc.UpdateSpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return resourcePingFederateSpAdapterResourceReadResult(d, result, svc)
}

func resourcePingFederateSpAdapterResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).SpAdapters
	input := pf.DeleteSpAdapterInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteSpAdapter(&input)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func resourcePingFederateSpAdapterResourceReadResult(d *schema.ResourceData, rv *pf.SpAdapter, svc *pf.SpAdaptersService) (err error) {
	desc, _, err := svc.GetSpAdapterDescriptorsById(&pf.GetSpAdapterDescriptorsByIdInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return err
	}
	setResourceDataString(d, "name", rv.Name)
	setResourceDataString(d, "sp_adapter_id", rv.Id)
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
	}

	if rv.AttributeContract != nil {
		if err = d.Set("attribute_contract", flattenSpAdapterAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}

	if rv.TargetApplicationInfo != nil {
		if err = d.Set("target_application_info", flattenSpAdapterTargetApplicationInfo(rv.TargetApplicationInfo)); err != nil {
			return err
		}
	}
	return nil
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
