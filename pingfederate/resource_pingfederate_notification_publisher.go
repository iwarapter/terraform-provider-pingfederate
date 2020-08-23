package pingfederate

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingfederate-sdk-go/services/notificationPublishers"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateNotificationPublisherResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateNotificationPublisherResourceCreate,
		ReadContext:   resourcePingFederateNotificationPublisherResourceRead,
		UpdateContext: resourcePingFederateNotificationPublisherResourceUpdate,
		DeleteContext: resourcePingFederateNotificationPublisherResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateNotificationPublisherResourceSchema(),
		CustomizeDiff: func(ctx context.Context, d *schema.ResourceDiff, m interface{}) error {
			svc := m.(pfClient).NotificationPublishers
			if className, ok := d.GetOk("plugin_descriptor_ref.0.id"); ok {
				desc, _, err := svc.GetNotificationPublisherPluginDescriptor(&notificationPublishers.GetNotificationPublisherPluginDescriptorInput{Id: className.(string)})
				if err != nil {
					descs, _, err := svc.GetNotificationPublisherPluginDescriptors()
					if err == nil && descs != nil {
						list := func(in *[]*pf.NotificationPublisherDescriptor) string {
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

func resourcePingFederateNotificationPublisherResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"publisher_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
		"parent_ref":            resourceLinkSchema(),
		"configuration":         resourcePluginConfiguration(),
	}
}

func resourcePingFederateNotificationPublisherResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).NotificationPublishers
	input := notificationPublishers.CreateNotificationPublisherInput{
		Body: *resourcePingFederateNotificationPublisherResourceReadData(d),
	}
	result, _, err := svc.CreateNotificationPublisher(&input)
	if err != nil {
		return diag.Errorf("unable to create NotificationPublishers: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateNotificationPublisherResourceReadResult(d, result, svc)
}

func resourcePingFederateNotificationPublisherResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).NotificationPublishers
	input := notificationPublishers.GetNotificationPublisherInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetNotificationPublisher(&input)
	if err != nil {
		return diag.Errorf("unable to read NotificationPublishers: %s", err)
	}
	return resourcePingFederateNotificationPublisherResourceReadResult(d, result, svc)
}

func resourcePingFederateNotificationPublisherResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).NotificationPublishers
	input := notificationPublishers.UpdateNotificationPublisherInput{
		Id:   d.Id(),
		Body: *resourcePingFederateNotificationPublisherResourceReadData(d),
	}
	result, _, err := svc.UpdateNotificationPublisher(&input)
	if err != nil {
		return diag.Errorf("unable to update NotificationPublishers: %s", err)
	}

	return resourcePingFederateNotificationPublisherResourceReadResult(d, result, svc)
}

func resourcePingFederateNotificationPublisherResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).NotificationPublishers
	input := notificationPublishers.DeleteNotificationPublisherInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteNotificationPublisher(&input)
	if err != nil {
		return diag.Errorf("unable to delete NotificationPublishers: %s", err)
	}
	return nil
}

func resourcePingFederateNotificationPublisherResourceReadResult(d *schema.ResourceData, rv *pf.NotificationPublisher, svc notificationPublishers.NotificationPublishersAPI) diag.Diagnostics {
	desc, _, err := svc.GetNotificationPublisherPluginDescriptor(&notificationPublishers.GetNotificationPublisherPluginDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve NotificationPublishers descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringithDiagnostic(d, "publisher_id", rv.Id, &diags)
	if rv.PluginDescriptorRef != nil {
		if err = d.Set("plugin_descriptor_ref", flattenResourceLink(rv.PluginDescriptorRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.ParentRef != nil {
		if err = d.Set("parent_ref", flattenResourceLink(rv.ParentRef)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}

	if rv.Configuration != nil {
		orig := expandPluginConfiguration(d.Get("configuration").([]interface{}))

		if err = d.Set("configuration", maskPluginConfigurationFromDescriptor(desc.ConfigDescriptor, orig, rv.Configuration)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateNotificationPublisherResourceReadData(d *schema.ResourceData) *pf.NotificationPublisher {
	validator := &pf.NotificationPublisher{
		Id:                  String(d.Get("publisher_id").(string)),
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}

	if v, ok := d.GetOk("parent_ref"); ok {
		validator.ParentRef = expandResourceLink(v.([]interface{}))
	}
	return validator
}
