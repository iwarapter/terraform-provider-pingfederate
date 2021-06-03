package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/iwarapter/pingfederate-sdk-go/services/incomingProxySettings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateIncomingProxySettingsResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateIncomingProxySettingsResourceCreate,
		ReadContext:   resourcePingFederateIncomingProxySettingsResourceRead,
		UpdateContext: resourcePingFederateIncomingProxySettingsResourceUpdate,
		DeleteContext: resourcePingFederateIncomingProxySettingsResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateIncomingProxySettingsResourceImport,
		},

		Schema: map[string]*schema.Schema{
			"client_cert_chain_sslheader_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"client_cert_sslheader_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarded_host_header_index": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateHeaderIndex(),
				//Default:          "LAST",
			},
			"forwarded_host_header_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarded_ip_address_header_index": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validateHeaderIndex(),
				//Default:          "LAST",
			},
			"forwarded_ip_address_header_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy_terminates_https_conns": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func validateHeaderIndex() func(value interface{}, path cty.Path) diag.Diagnostics {
	return func(value interface{}, path cty.Path) diag.Diagnostics {
		v := value.(string)
		switch v {
		case
			"LAST",
			"FIRST":
			return nil
		}
		return diag.Errorf("must be either 'LAST' or 'FIRST' not %s", v)
	}
}

func resourcePingFederateIncomingProxySettingsResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("IncomingProxySettings")
	return resourcePingFederateIncomingProxySettingsResourceUpdate(ctx, d, m)
}

func resourcePingFederateIncomingProxySettingsResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if !m.(pfClient).IsPF10_1orGreater() {
		return diag.Errorf("incoming_proxy_settings is only available from PF 10.1+")
	}

	svc := m.(pfClient).IncomingProxySettings
	result, _, err := svc.GetIncomingProxySettingsWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read IncomingProxySettings: %s", err)
	}
	return resourcePingFederateIncomingProxySettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
}

func resourcePingFederateIncomingProxySettingsResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	settings := &pf.IncomingProxySettings{}
	if v, ok := d.GetOk("client_cert_chain_sslheader_name"); ok {
		settings.ClientCertChainSSLHeaderName = String(v.(string))
	}
	if v, ok := d.GetOk("client_cert_sslheader_name"); ok {
		settings.ClientCertSSLHeaderName = String(v.(string))
	}
	if v, ok := d.GetOk("forwarded_host_header_index"); ok {
		settings.ForwardedHostHeaderIndex = String(v.(string))
	}
	if v, ok := d.GetOk("forwarded_host_header_name"); ok {
		settings.ForwardedHostHeaderName = String(v.(string))
	}
	if v, ok := d.GetOk("forwarded_ip_address_header_index"); ok {
		settings.ForwardedIpAddressHeaderIndex = String(v.(string))
	}
	if v, ok := d.GetOk("forwarded_ip_address_header_name"); ok {
		settings.ForwardedIpAddressHeaderName = String(v.(string))
	}
	if v, ok := d.GetOk("proxy_terminates_https_conns"); ok {
		settings.ProxyTerminatesHttpsConns = Bool(v.(bool))
	}

	svc := m.(pfClient).IncomingProxySettings
	input := &incomingProxySettings.UpdateIncomingProxySettingsInput{
		Body: *settings,
	}

	result, _, err := svc.UpdateIncomingProxySettingsWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to update IncomingProxySettings: %s", err)
	}
	return resourcePingFederateIncomingProxySettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
}

func resourcePingFederateIncomingProxySettingsResourceDelete(_ context.Context, _ *schema.ResourceData, _ interface{}) diag.Diagnostics {
	//resource cannot be deleted just not tracked by terraform anymore
	return nil
}

func resourcePingFederateIncomingProxySettingsResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).IncomingProxySettings
	result, _, err := svc.GetIncomingProxySettingsWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	resourcePingFederateIncomingProxySettingsResourceReadResult(d, result, m.(pfClient).IsPF10())
	return []*schema.ResourceData{d}, nil
}

func resourcePingFederateIncomingProxySettingsResourceReadResult(d *schema.ResourceData, rv *pf.IncomingProxySettings, isPF10 bool) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "client_cert_chain_sslheader_name", rv.ClientCertChainSSLHeaderName, &diags)
	setResourceDataStringWithDiagnostic(d, "client_cert_sslheader_name", rv.ClientCertSSLHeaderName, &diags)
	setResourceDataStringWithDiagnostic(d, "forwarded_host_header_index", rv.ForwardedHostHeaderIndex, &diags)
	setResourceDataStringWithDiagnostic(d, "forwarded_host_header_name", rv.ForwardedHostHeaderName, &diags)
	setResourceDataStringWithDiagnostic(d, "forwarded_ip_address_header_index", rv.ForwardedIpAddressHeaderIndex, &diags)
	setResourceDataStringWithDiagnostic(d, "forwarded_ip_address_header_name", rv.ForwardedIpAddressHeaderName, &diags)
	setResourceDataBoolWithDiagnostic(d, "proxy_terminates_https_conns", rv.ProxyTerminatesHttpsConns, &diags)

	return diags
}
