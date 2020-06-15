package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateAuthenticationPolicyContractResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateAuthenticationPolicyContractResourceCreate,
		Read:   resourcePingFederateAuthenticationPolicyContractResourceRead,
		Update: resourcePingFederateAuthenticationPolicyContractResourceUpdate,
		Delete: resourcePingFederateAuthenticationPolicyContractResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: resourcePingFederateAuthenticationPolicyContractResourceSchema(),
	}
}

func resourcePingFederateAuthenticationPolicyContractResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"core_attributes": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"extended_attributes": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			MinItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourcePingFederateAuthenticationPolicyContractResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationPolicyContracts
	input := pf.CreateAuthenticationPolicyContractInput{
		Body: *resourcePingFederateAuthenticationPolicyContractResourceReadData(d),
	}
	result, _, err := svc.CreateAuthenticationPolicyContract(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.SetId(*result.Id)
	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationPolicyContracts
	input := pf.GetAuthenticationPolicyContractInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetAuthenticationPolicyContract(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationPolicyContracts
	input := pf.UpdateAuthenticationPolicyContractInput{
		Id:   d.Id(),
		Body: *resourcePingFederateAuthenticationPolicyContractResourceReadData(d),
	}
	result, _, err := svc.UpdateAuthenticationPolicyContract(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return resourcePingFederateAuthenticationPolicyContractResourceReadResult(d, result)
}

func resourcePingFederateAuthenticationPolicyContractResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).AuthenticationPolicyContracts
	input := pf.DeleteAuthenticationPolicyContractInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteAuthenticationPolicyContract(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateAuthenticationPolicyContractResourceReadResult(d *schema.ResourceData, rv *pf.AuthenticationPolicyContract) (err error) {
	setResourceDataString(d, "name", rv.Name)
	if rv.ExtendedAttributes != nil && len(*rv.ExtendedAttributes) > 0 {
		if err = d.Set("extended_attributes", flattenAuthenticationPolicyContractAttribute(*rv.ExtendedAttributes)); err != nil {
			return err
		}
	}
	if rv.CoreAttributes != nil && len(*rv.CoreAttributes) > 0 {
		if err = d.Set("core_attributes", flattenAuthenticationPolicyContractAttribute(*rv.CoreAttributes)); err != nil {
			return err
		}
	}

	return nil
}

func resourcePingFederateAuthenticationPolicyContractResourceReadData(d *schema.ResourceData) *pf.AuthenticationPolicyContract {
	core := expandAuthenticationPolicyContractAttribute(d.Get("core_attributes").(*schema.Set).List())
	contract := &pf.AuthenticationPolicyContract{
		Name:           String(d.Get("name").(string)),
		CoreAttributes: core,
	}

	if _, ok := d.GetOk("extended_attributes"); ok {
		contract.ExtendedAttributes = expandAuthenticationPolicyContractAttribute(d.Get("extended_attributes").(*schema.Set).List())
	}

	return contract
}
