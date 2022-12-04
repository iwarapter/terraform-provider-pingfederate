package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
)

func resourcePingFederateKeypairSslServerSettingsResource() *schema.Resource {
	return &schema.Resource{
		Description: `Manages the PingFederate instance Ssl Server Settings.

-> This resource manages a singleton within PingFederate and as such you should ONLY ever declare one of this resource type. Deleting this resource simply stops tracking changes.`,
		CreateContext: resourcePingFederateKeypairSslServerSettingsResourceCreate,
		ReadContext:   resourcePingFederateKeypairSslServerSettingsResourceRead,
		UpdateContext: resourcePingFederateKeypairSslServerSettingsResourceUpdate,
		DeleteContext: resourcePingFederateKeypairSslServerSettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederateKeypairSslServerSettingsResourceSchema(),
	}
}

func resourcePingFederateKeypairSslServerSettingsResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"runtime_server_cert": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Reference to the default SSL Server Certificate Key pair active for Runtime Server.",
		},
		"admin_server_cert": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Reference to the default SSL Server Certificate Key pair active for PF Administrator Console.",
		},
		"active_runtime_server_certs": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "The active SSL Server Certificate Key pairs for Runtime Server.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"active_admin_server_certs": {
			Type:        schema.TypeSet,
			Required:    true,
			MinItems:    1,
			Description: "The active SSL Server Certificate Key pairs for PF Administrator Console.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func resourcePingFederateKeypairSslServerSettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer

	input := &keyPairsSslServer.UpdateSettingsInput{Body: *resourcePingFederateKeypairSslServerSettingsResourceReadData(d)}
	result, _, err := svc.UpdateSettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to create SslServerSettings: %s", err)
	}

	d.SetId("ssl_server_settings")
	return resourcePingFederateKeypairSslServerSettingsResourceReadResult(d, result)

}

func resourcePingFederateKeypairSslServerSettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer
	result, _, err := svc.GetSettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read SslServerSettings: %s", err)
	}
	return resourcePingFederateKeypairSslServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateKeypairSslServerSettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSslServer

	input := &keyPairsSslServer.UpdateSettingsInput{Body: *resourcePingFederateKeypairSslServerSettingsResourceReadData(d)}
	result, _, err := svc.UpdateSettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to create SslServerSettings: %s", err)
	}
	return resourcePingFederateKeypairSslServerSettingsResourceReadResult(d, result)
}

func resourcePingFederateKeypairSslServerSettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	return nil
}

func resourcePingFederateKeypairSslServerSettingsResourceReadResult(d *schema.ResourceData, rv *pf.SslServerSettings) diag.Diagnostics {
	var diags diag.Diagnostics
	if err := d.Set("admin_server_cert", rv.AdminConsoleCertRef.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	if err := d.Set("runtime_server_cert", rv.RuntimeServerCertRef.Id); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	admins := []string{}
	for _, link := range *rv.ActiveAdminConsoleCerts {
		admins = append(admins, *link.Id)
	}
	if err := d.Set("active_admin_server_certs", admins); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	runtimes := []string{}
	for _, link := range *rv.ActiveRuntimeServerCerts {
		runtimes = append(runtimes, *link.Id)
	}
	if err := d.Set("active_runtime_server_certs", runtimes); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	return diags
}

func resourcePingFederateKeypairSslServerSettingsResourceReadData(d *schema.ResourceData) *pf.SslServerSettings {
	settings := &pf.SslServerSettings{
		AdminConsoleCertRef:      &pf.ResourceLink{Id: String(d.Get("admin_server_cert").(string))},
		RuntimeServerCertRef:     &pf.ResourceLink{Id: String(d.Get("runtime_server_cert").(string))},
		ActiveRuntimeServerCerts: &[]*pf.ResourceLink{},
		ActiveAdminConsoleCerts:  &[]*pf.ResourceLink{},
	}
	admins := d.Get("active_admin_server_certs")
	for _, v := range admins.(*schema.Set).List() {
		*settings.ActiveAdminConsoleCerts = append(*settings.ActiveAdminConsoleCerts, &pf.ResourceLink{Id: String(v.(string))})
	}
	runtimes := d.Get("active_runtime_server_certs")
	for _, v := range runtimes.(*schema.Set).List() {
		*settings.ActiveRuntimeServerCerts = append(*settings.ActiveRuntimeServerCerts, &pf.ResourceLink{Id: String(v.(string))})
	}
	return settings
}
