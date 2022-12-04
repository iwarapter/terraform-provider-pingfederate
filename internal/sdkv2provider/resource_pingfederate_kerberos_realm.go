package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/kerberosRealms"
)

func resourcePingFederateKerberosRealmResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for Kerberos Realms within PingFederate.",
		CreateContext: resourcePingFederateKerberosRealmResourceCreate,
		ReadContext:   resourcePingFederateKerberosRealmResourceRead,
		UpdateContext: resourcePingFederateKerberosRealmResourceUpdate,
		DeleteContext: resourcePingFederateKerberosRealmResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateKerberosRealmResourceSchema(),
	}
}

func resourcePingFederateKerberosRealmResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"realm_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the Kerberos Realm. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.",
		},
		"kerberos_realm_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The Domain/Realm name used for display in UI screens.",
		},
		"key_distribution_centers": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "The Domain Controller/Key Distribution Center Host Action Names.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"kerberos_username": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The Domain/Realm username.",
		},
		"kerberos_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Sensitive:   true,
			Description: "The Domain/Realm password. GETs will not return this attribute. To update this field, specify the new value in this attribute.",
		},
		"kerberos_encrypted_password": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			Description: "For GET requests, this field contains the encrypted Domain/Realm password, if one exists. For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged.",
		},
		"suppress_domain_name_concatenation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Controls whether the KDC hostnames and the realm name are concatenated in the auto-generated krb5.conf file. Default is false.",
		},
	}
}

func resourcePingFederateKerberosRealmResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.CreateKerberosRealmInput{
		Body: *resourcePingFederateKerberosRealmResourceReadData(d),
	}
	result, _, err := svc.CreateKerberosRealmWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to create KerberosRealm: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.GetKerberosRealmInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKerberosRealmWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read KerberosRealm: %s", err)
	}
	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.UpdateKerberosRealmInput{
		Id:   d.Id(),
		Body: *resourcePingFederateKerberosRealmResourceReadData(d),
	}
	result, _, err := svc.UpdateKerberosRealmWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update KerberosRealm: %s", err)
	}

	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.DeleteKerberosRealmInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKerberosRealmWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete KerberosRealm: %s", err)
	}
	return nil
}

func resourcePingFederateKerberosRealmResourceReadResult(d *schema.ResourceData, rv *pf.KerberosRealm) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "realm_id", rv.Id, &diags)
	setResourceDataStringWithDiagnostic(d, "kerberos_realm_name", rv.KerberosRealmName, &diags)
	setResourceDataStringWithDiagnostic(d, "kerberos_username", rv.KerberosUsername, &diags)
	setResourceDataStringWithDiagnostic(d, "kerberos_password", rv.KerberosPassword, &diags)
	setResourceDataStringWithDiagnostic(d, "kerberos_encrypted_password", rv.KerberosEncryptedPassword, &diags)
	setResourceDataBoolWithDiagnostic(d, "suppress_domain_name_concatenation", rv.SuppressDomainNameConcatenation, &diags)
	if rv.KeyDistributionCenters != nil {
		if err := d.Set("key_distribution_centers", rv.KeyDistributionCenters); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}

func resourcePingFederateKerberosRealmResourceReadData(d *schema.ResourceData) *pf.KerberosRealm {
	result := pf.KerberosRealm{
		KerberosRealmName:               String(d.Get("kerberos_realm_name").(string)),
		KerberosUsername:                String(d.Get("kerberos_username").(string)),
		SuppressDomainNameConcatenation: Bool(d.Get("suppress_domain_name_concatenation").(bool)),
	}
	if val, ok := d.GetOk("realm_id"); ok {
		result.Id = String(val.(string))
	}
	if val, ok := d.GetOk("kerberos_password"); ok {
		result.KerberosPassword = String(val.(string))
	}
	if val, ok := d.GetOk("kerberos_encrypted_password"); ok {
		result.KerberosEncryptedPassword = String(val.(string))
	}
	if val, ok := d.GetOk("key_distribution_centers"); ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.KeyDistributionCenters = &strs
	}
	return &result
}
