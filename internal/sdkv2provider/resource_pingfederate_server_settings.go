package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"
)

func resourcePingFederateServerSettingsResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages the PingFederate instance server settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateServerSettingsResourceCreate,
		ReadContext:   resourcePingFederateServerSettingsResourceRead,
		UpdateContext: resourcePingFederateServerSettingsResourceUpdate,
		DeleteContext: resourcePingFederateServerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateServerSettingsResourceSchema(),
	}
}

func resourcePingFederateServerSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"federation_info": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"base_url": {
						Type:        schema.TypeString,
						Description: "The fully qualified host name, port, and path (if applicable) on which the PingFederate server runs.",
						Optional:    true,
					},
					"saml1x_issuer_id": {
						Type:        schema.TypeString,
						Description: "This ID identifies your federation server for SAML 1.x transactions. As with SAML 2.0, it is usually defined as an organization's URL or a DNS address. The SourceID used for artifact resolution is derived from this ID using SHA1.",
						Optional:    true,
					},
					"saml1x_source_id": {
						Type:        schema.TypeString,
						Description: "If supplied, the Source ID value entered here is used for SAML 1.x, instead of being derived from the SAML 1.x Issuer/Audience.",
						Optional:    true,
					},
					"saml2_entity_id": {
						Type:        schema.TypeString,
						Description: "This ID defines your organization as the entity operating the server for SAML 2.0 transactions. It is usually defined as an organization's URL or a DNS address; for example: pingidentity.com. The SAML SourceID used for artifact resolution is derived from this ID using SHA1.",
						Optional:    true,
					},
					"wsfed_realm": {
						Type:        schema.TypeString,
						Description: "The URI of the realm associated with the PingFederate server. A realm represents a single unit of security administration or trust.",
						Optional:    true,
					},
				},
			},
		},
		"roles_and_protocols": {
			Type:       schema.TypeList,
			Optional:   true,
			Deprecated: "Starting with PingFederate 10.1, roles and protocols are always enabled and no longer configurable through the administrative console and API.",
			MaxItems:   1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enable_idp_discovery": {
						Type:     schema.TypeBool,
						Optional: true,
						Default:  false,
					},
					"oauth_role": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable_oauth": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  true,
								},
								"enable_openid_connect": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
							},
						},
					},
					"idp_role": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  true,
								},
								"saml20_profile": {
									Type:     schema.TypeList,
									MaxItems: 1,
									Optional: true,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"enable": {
												Type:     schema.TypeBool,
												Optional: true,
												Default:  true,
											},
											"enable_auto_connect": {
												Type:     schema.TypeBool,
												Optional: true,
												Default:  false,
											},
										},
									},
								},
								"enable_outbound_provisioning": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
								"enable_saml11": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
								"enable_saml10": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
								"enable_ws_fed": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
								"enable_ws_trust": {
									Type:     schema.TypeBool,
									Optional: true,
									Default:  false,
								},
							},
						},
					},
					"sp_role": {
						Type:     schema.TypeList,
						MaxItems: 1,
						Required: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"saml20_profile": {
									Type:     schema.TypeList,
									MaxItems: 1,
									Optional: true,
									Computed: true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"enable": {
												Type:     schema.TypeBool,
												Optional: true,
											},
											"enable_xasp": {
												Type:     schema.TypeBool,
												Optional: true,
											},
											"enable_auto_connect": {
												Type:     schema.TypeBool,
												Optional: true,
											},
										},
									},
								},
								"enable_inbound_provisioning": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"enable_saml11": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"enable_saml10": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"enable_ws_fed": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"enable_ws_trust": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"enable_openid_connect": {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourcePingFederateServerSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	input := serverSettings.UpdateServerSettingsInput{
		Body: *resourcePingFederateServerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateServerSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create ServerSettings: %s", err)
	}
	d.SetId("server_settings")
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	result, _, err := svc.GetServerSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read ServerSettings: %s", err)
	}
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	input := serverSettings.UpdateServerSettingsInput{
		Body: *resourcePingFederateServerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateServerSettingsWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update ServerSettings: %s", err)
	}
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {

	return nil
}

func resourcePingFederateServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.ServerSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	if rv.RolesAndProtocols != nil {
		if err := d.Set("roles_and_protocols", flattenRolesAndProtocols(rv.RolesAndProtocols)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	if rv.FederationInfo != nil {
		if err := d.Set("federation_info", flattenFederationInfo(rv.FederationInfo)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateServerSettingsResourceReadData(d *schema.ResourceData) *pf.ServerSettings {
	validator := &pf.ServerSettings{
		RolesAndProtocols: expandRolesAndProtcols(d.Get("roles_and_protocols").([]interface{})),
		FederationInfo:    expandFederationInfo(d.Get("federation_info").([]interface{})),
	}
	return validator
}
