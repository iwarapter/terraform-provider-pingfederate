package pingfederate

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
)

func resourcePingFederateOauthAccessTokenManagersResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateOauthAccessTokenManagersResourceCreate,
		Read:   resourcePingFederateOauthAccessTokenManagersResourceRead,
		Update: resourcePingFederateOauthAccessTokenManagersResourceUpdate,
		Delete: resourcePingFederateOauthAccessTokenManagersResourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ValidateFunc:       "message": "The plugin ID must be less than 33 characters, contain no spaces, and be alphanumeric.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"plugin_descriptor_ref": resourceLinkSchema(),
			"configuration":         resourcePluginConfiguration(),
			"parent_ref":            resourceLinkSchema(),
			"attribute_contract": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"extended_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MinItems: 1,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"selection_settings": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inherited": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"resource_uris": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"session_validation_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inherited": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"check_valid_authn_session": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"check_session_revocation_status": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
						"update_authn_session_activity": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePingFederateOauthAccessTokenManagersResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenManagers
	input := pf.CreateTokenManagerInput{
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d),
	}
	result, _, err := svc.CreateTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenManagers
	input := pf.GetTokenManagerInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenManagers
	input := pf.UpdateTokenManagerInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d),
	}
	result, _, err := svc.UpdateTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(*pf.PfClient).OauthAccessTokenManagers
	input := pf.DeleteTokenManagerInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateOauthAccessTokenManagersResourceReadResult(d *schema.ResourceData, rv *pf.AccessTokenManager, svc *pf.OauthAccessTokenManagersService) (err error) {
	desc, _, err := svc.GetTokenManagerDescriptor(&pf.GetTokenManagerDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return err
	}

	setResourceDataString(d, "name", rv.Name)
	setResourceDataString(d, "instance_id", rv.Id)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
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
		if err = d.Set("attribute_contract", flattenAccessTokenAttributeContract(rv.AttributeContract)); err != nil {
			return err
		}
	}
	return nil
}

func resourcePingFederateOauthAccessTokenManagersResourceReadData(d *schema.ResourceData) *pf.AccessTokenManager {
	atm := &pf.AccessTokenManager{
		Name:                String(d.Get("name").(string)),
		Id:                  String(d.Get("instance_id").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
		AttributeContract:   expandAccessTokenAttributeContract(d.Get("attribute_contract").([]interface{})),
	}
	return atm
}
