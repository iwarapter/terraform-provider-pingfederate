package pingfederate

import (
	"fmt"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthAccessTokenManagersResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePingFederateOauthAccessTokenManagersResourceCreate,
		Read:   resourcePingFederateOauthAccessTokenManagersResourceRead,
		Update: resourcePingFederateOauthAccessTokenManagersResourceUpdate,
		Delete: resourcePingFederateOauthAccessTokenManagersResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourcePingFederateOauthAccessTokenManagersResourceSchema(),
	}
}

func resourcePingFederateOauthAccessTokenManagersResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"instance_id": {
			Type:     schema.TypeString,
			Required: true,
			//ValidateFunc:       "message": "The plugin ID must be less than 33 characters, contain no spaces, and be alphanumeric.",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"plugin_descriptor_ref": resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
		"parent_ref":            resourceLinkSchema(),
		"attribute_contract": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"core_attributes": {
						Type:     schema.TypeList,
						Computed: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"extended_attributes": {
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
		"selection_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"inherited": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"resource_uris": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		"session_validation_settings": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"inherited": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"check_valid_authn_session": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"check_session_revocation_status": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"update_authn_session_activity": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
	}
}

func resourcePingFederateOauthAccessTokenManagersResourceCreate(d *schema.ResourceData, m interface{}) error {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.CreateTokenManagerInput{
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d, svc),
	}
	result, _, err := svc.CreateTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.SetId(*result.Id)
	err = resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
	if err != nil {
		return err
	}
	return setOauthAccessTokenManagerAsDefaultIfNoneSet(*result.Id, svc)
}

func setOauthAccessTokenManagerAsDefaultIfNoneSet(id string, svc oauthAccessTokenManagers.OauthAccessTokenManagersAPI) error {
	result, _, err := svc.GetSettings()
	if err != nil {
		return fmt.Errorf("unable to read the OauthAccessTokenManager settings %s", err)
	}
	if result.DefaultAccessTokenManagerRef != nil {
		_, _, err = svc.UpdateSettings(&oauthAccessTokenManagers.UpdateSettingsInput{
			Body: pf.AccessTokenManagementSettings{
				DefaultAccessTokenManagerRef: &pf.ResourceLink{Id: String(id)},
			},
		})
		if err != nil {
			return fmt.Errorf("unable to update the OauthAccessTokenManager settings %s", err)
		}
	}
	return nil
}

func resourcePingFederateOauthAccessTokenManagersResourceRead(d *schema.ResourceData, m interface{}) error {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.GetTokenManagerInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceUpdate(d *schema.ResourceData, m interface{}) error {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.UpdateTokenManagerInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d, svc),
	}
	result, _, err := svc.UpdateTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceDelete(d *schema.ResourceData, m interface{}) error {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.DeleteTokenManagerInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteTokenManager(&input)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func resourcePingFederateOauthAccessTokenManagersResourceReadResult(d *schema.ResourceData, rv *pf.AccessTokenManager, svc oauthAccessTokenManagers.OauthAccessTokenManagersAPI) (err error) {
	desc, _, err := svc.GetTokenManagerDescriptor(&oauthAccessTokenManagers.GetTokenManagerDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
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

func resourcePingFederateOauthAccessTokenManagersResourceReadData(d *schema.ResourceData, svc oauthAccessTokenManagers.OauthAccessTokenManagersAPI) *pf.AccessTokenManager {
	//desc, _, err := svc.GetTokenManagerDescriptor(&pf.GetTokenManagerDescriptorInput{Id: *expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})).Id})
	//if err != nil {
	//	//TODO
	//}
	atm := &pf.AccessTokenManager{
		Name:                String(d.Get("name").(string)),
		Id:                  String(d.Get("instance_id").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		//Configuration:       expandPluginConfigurationWithDescriptor(d.Get("configuration").([]interface{}), desc.ConfigDescriptor),
		Configuration:     expandPluginConfiguration(d.Get("configuration").([]interface{})),
		AttributeContract: expandAccessTokenAttributeContract(d.Get("attribute_contract").([]interface{})),
	}
	return atm
}
