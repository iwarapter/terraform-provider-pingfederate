package pingfederate

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOauthAccessTokenManagersResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateOauthAccessTokenManagersResourceCreate,
		ReadContext:   resourcePingFederateOauthAccessTokenManagersResourceRead,
		UpdateContext: resourcePingFederateOauthAccessTokenManagersResourceUpdate,
		DeleteContext: resourcePingFederateOauthAccessTokenManagersResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateOauthAccessTokenManagersResourceSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(pfClient).OauthAccessTokenManagers
			if className, ok := d.GetOk("plugin_descriptor_ref.0.id"); ok {
				desc, resp, err := svc.GetTokenManagerDescriptorWithContext(ctx, &oauthAccessTokenManagers.GetTokenManagerDescriptorInput{Id: className.(string)})
				if resp != nil && resp.StatusCode == http.StatusForbidden {
					log.Printf("[WARN] Unable to query OAuthTokenManagerDescriptor, OAuth 2.0 authorization server role enabled")
					return nil
				}
				if err != nil {
					descs, _, err := svc.GetTokenManagerDescriptorsWithContext(ctx)
					if err == nil && descs != nil {
						list := func(in *[]*pf.AccessTokenManagerDescriptor) string {
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
		"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
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
						Type:     schema.TypeSet,
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

func resourcePingFederateOauthAccessTokenManagersResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.CreateTokenManagerInput{
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d, svc),
	}
	result, _, err := svc.CreateTokenManagerWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OauthAccessTokenManagers: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.GetTokenManagerInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetTokenManagerWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read OauthAccessTokenManagers: %s", err)
	}
	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.UpdateTokenManagerInput{
		Id:   d.Id(),
		Body: *resourcePingFederateOauthAccessTokenManagersResourceReadData(d, svc),
	}
	result, _, err := svc.UpdateTokenManagerWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update OauthAccessTokenManagers: %s", err)
	}

	return resourcePingFederateOauthAccessTokenManagersResourceReadResult(d, result, svc)
}

func resourcePingFederateOauthAccessTokenManagersResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthAccessTokenManagers
	input := oauthAccessTokenManagers.DeleteTokenManagerInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteTokenManagerWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete OauthAccessTokenManagers: %s", err)
	}
	return nil
}

func resourcePingFederateOauthAccessTokenManagersResourceReadResult(d *schema.ResourceData, rv *pf.AccessTokenManager, svc oauthAccessTokenManagers.OauthAccessTokenManagersAPI) diag.Diagnostics {
	desc, _, err := svc.GetTokenManagerDescriptor(&oauthAccessTokenManagers.GetTokenManagerDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve oauthAccessTokenManagers descriptor: %s", err)

	}
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "instance_id", rv.Id, &diags)
	if rv.PluginDescriptorRef != nil {
		if err := d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
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
		if err := d.Set("attribute_contract", flattenAccessTokenAttributeContract(rv.AttributeContract)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateOauthAccessTokenManagersResourceReadData(d *schema.ResourceData, svc oauthAccessTokenManagers.OauthAccessTokenManagersAPI) *pf.AccessTokenManager {
	//desc, _, err := svc.GetTokenManagerDescriptor(&pf.GetTokenManagerDescriptorInput{Id: *expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})).Id})
	//if err != nil {
	//	//TODO
	//}
	atm := &pf.AccessTokenManager{
		Name:                String(d.Get("name").(string)),
		Id:                  String(d.Get("instance_id").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})[0].(map[string]interface{})),
		//Configuration:       expandPluginConfigurationWithDescriptor(d.Get("configuration").([]interface{}), desc.ConfigDescriptor),
		Configuration:     expandPluginConfiguration(d.Get("configuration").([]interface{})),
		AttributeContract: expandAccessTokenAttributeContract(d.Get("attribute_contract").([]interface{})),
	}
	return atm
}
