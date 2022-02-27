package sdkv2provider

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientRegistrationPolicies"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateOAuthClientRegistrationPolicyResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for OAuth Client Registration Policy within PingFederate.",
		CreateContext: resourcePingFederateOAuthClientRegistrationPolicyResourceCreate,
		ReadContext:   resourcePingFederateOAuthClientRegistrationPolicyResourceRead,
		UpdateContext: resourcePingFederateOAuthClientRegistrationPolicyResourceUpdate,
		DeleteContext: resourcePingFederateOAuthClientRegistrationPolicyResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateOAuthClientRegistrationPolicyResourceSchema(),
	}
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"policy_id": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			Description:  "The persistent, unique ID for the plugin instance. It can be any combination of [a-zA-Z0-9._-]. This property is system-assigned if not specified.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9._-]{1,35}$`), "the policy_id can only contain alphanumeric characters, dash, dot and underscore."),
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The plugin instance name. The name cannot be modified once the instance is created.",
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
	}
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClientRegistrationPolicies
	ds := resourcePingFederateOAuthClientRegistrationPolicyResourceReadData(d)
	input := oauthClientRegistrationPolicies.CreateDynamicClientRegistrationPolicyInput{
		Body: *ds,
	}
	store, _, err := svc.CreateDynamicClientRegistrationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create Client Registration Policy: %s", err)
	}
	d.SetId(*store.Id)
	return resourcePingFederateOAuthClientRegistrationPolicyResourceReadResult(d, store, svc)
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClientRegistrationPolicies
	input := oauthClientRegistrationPolicies.GetDynamicClientRegistrationPolicyInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetDynamicClientRegistrationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read Client Registration Policy: %s", err)
	}
	return resourcePingFederateOAuthClientRegistrationPolicyResourceReadResult(d, result, svc)
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).OauthClientRegistrationPolicies
	ds := resourcePingFederateOAuthClientRegistrationPolicyResourceReadData(d)
	input := oauthClientRegistrationPolicies.UpdateDynamicClientRegistrationPolicyInput{
		Id:   d.Id(),
		Body: *ds,
	}
	store, _, err := svc.UpdateDynamicClientRegistrationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update Client Registration Policy: %s", err)
	}
	return resourcePingFederateOAuthClientRegistrationPolicyResourceReadResult(d, store, svc)
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).OauthClientRegistrationPolicies
	input := oauthClientRegistrationPolicies.DeleteDynamicClientRegistrationPolicyInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteDynamicClientRegistrationPolicyWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete Client Registration Policy: %s", err)
	}
	return nil
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceReadResult(d *schema.ResourceData, rv *pf.ClientRegistrationPolicy, svc oauthClientRegistrationPolicies.OauthClientRegistrationPoliciesAPI) diag.Diagnostics {
	desc, _, err := svc.GetDynamicClientRegistrationDescriptor(&oauthClientRegistrationPolicies.GetDynamicClientRegistrationDescriptorInput{Id: *rv.PluginDescriptorRef.Id})
	if err != nil {
		return diag.Errorf("unable to retrieve Client Registration Policy descriptor: %s", err)
	}
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "policy_id", rv.Id, &diags)
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

	return diags
}

func resourcePingFederateOAuthClientRegistrationPolicyResourceReadData(d *schema.ResourceData) *pf.ClientRegistrationPolicy {
	ds := &pf.ClientRegistrationPolicy{
		Name:                String(d.Get("name").(string)),
		PluginDescriptorRef: expandResourceLink(d.Get("plugin_descriptor_ref").([]interface{})[0].(map[string]interface{})),
		Configuration:       expandPluginConfiguration(d.Get("configuration").([]interface{})),
	}
	if v, ok := d.GetOk("policy_id"); ok {
		ds.Id = String(v.(string))
	}
	if v, ok := d.GetOk("parent_ref"); ok && len(v.([]interface{})) > 0 {
		ds.ParentRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	return ds
}
