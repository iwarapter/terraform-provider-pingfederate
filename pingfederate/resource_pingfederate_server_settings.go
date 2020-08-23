package pingfederate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"
)

func resourcePingFederateServerSettingsResource() *schema.Resource {
	return &schema.Resource{
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
						Type:     schema.TypeString,
						Required: true,
					},
					"saml2_entity_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"saml1x_issuer_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"saml1x_source_id": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"wsfed_realm": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"roles_and_protocols": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
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

func resourcePingFederateServerSettingsResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	input := serverSettings.UpdateServerSettingsInput{
		Body: *resourcePingFederateServerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateServerSettings(&input)
	if err != nil {
		return diag.Errorf("unable to create ServerSettings: %s", err)
	}
	d.SetId("server_settings")
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	result, _, err := svc.GetServerSettings()
	if err != nil {
		return diag.Errorf("unable to read ServerSettings: %s", err)
	}
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ServerSettings
	input := serverSettings.UpdateServerSettingsInput{
		Body: *resourcePingFederateServerSettingsResourceReadData(d),
	}
	result, _, err := svc.UpdateServerSettings(&input)
	if err != nil {
		return diag.Errorf("unable to update ServerSettings: %s", err)
	}
	return resourcePingFederateServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateServerSettingsResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

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
