package pingfederate

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/kerberosRealms"
)

func resourcePingFederateKerberosRealmResource() *schema.Resource {
	return &schema.Resource{
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
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"kerberos_realm_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key_distribution_centers": setOfString(),
		"kerberos_username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"kerberos_password": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"kerberos_encrypted_password": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"suppress_domain_name_concatenation": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}

func resourcePingFederateKerberosRealmResourceCreate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.CreateKerberosRealmInput{
		Body: *resourcePingFederateKerberosRealmResourceReadData(d),
	}
	result, _, err := svc.CreateKerberosRealm(&input)
	if err != nil {
		return diag.Errorf("unable to create KerberosRealm: %s", err)
	}
	d.SetId(*result.Id)
	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.GetKerberosRealmInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKerberosRealm(&input)
	if err != nil {
		return diag.Errorf("unable to read KerberosRealm: %s", err)
	}
	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceUpdate(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.UpdateKerberosRealmInput{
		Id:   d.Id(),
		Body: *resourcePingFederateKerberosRealmResourceReadData(d),
	}
	result, _, err := svc.UpdateKerberosRealm(&input)
	if err != nil {
		return diag.Errorf("unable to update KerberosRealm: %s", err)
	}

	return resourcePingFederateKerberosRealmResourceReadResult(d, result)
}

func resourcePingFederateKerberosRealmResourceDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KerberosRealms
	input := kerberosRealms.DeleteKerberosRealmInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKerberosRealm(&input)
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
