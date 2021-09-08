package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/pingOneConnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederatePingOneConnectionResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for PingOne Connections within PingFederate.",
		CreateContext: resourcePingFederatePingOneConnectionResourceCreate,
		ReadContext:   resourcePingFederatePingOneConnectionResourceRead,
		UpdateContext: resourcePingFederatePingOneConnectionResourceUpdate,
		DeleteContext: resourcePingFederatePingOneConnectionResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: resourcePingFederatePingOneConnectionResourceSchema(),
	}
}

func resourcePingFederatePingOneConnectionResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:        schema.TypeBool,
			Description: "Whether or not this connection is active. Defaults to true.",
			Optional:    true,
			Default:     true,
		},
		"creation_date": {
			Type:        schema.TypeString,
			Description: "The creation date of the PingOne connection. This field is read only.",
			Computed:    true,
		},
		"credential": {
			Type:          schema.TypeString,
			Description:   "The credential for the PingOne connection. To update the credential, specify the plaintext value of the credential in this field. This field will not be populated for GET requests.",
			Optional:      true,
			ConflictsWith: []string{"encrypted_credential"},
		},
		"credential_id": {
			Type:        schema.TypeString,
			Description: "The ID of the PingOne credential. This field is read only.",
			Computed:    true,
		},
		"description": {
			Type:        schema.TypeString,
			Description: "A description for the PingOne connection.",
			Optional:    true,
		},
		"encrypted_credential": {
			Type:          schema.TypeString,
			Description:   "The encrypted credential for the PingOne connection. For POST and PUT requests, if you wish to keep the existing credential, this field should be passed back unchanged.",
			Optional:      true,
			ConflictsWith: []string{"credential"},
		},
		"environment_id": {
			Type:        schema.TypeString,
			Description: "The ID of the environment of the PingOne credential. This field is read only.",
			Computed:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the PingOne connection.",
			Required:    true,
		},
		"organization_name": {
			Type:        schema.TypeString,
			Description: "The name of the organization associated with this PingOne connection. This field is read only.",
			Computed:    true,
		},
		"ping_one_authentication_api_endpoint": {
			Type:        schema.TypeString,
			Description: "The PingOne authentication API endpoint. This field is read only.",
			Computed:    true,
		},
		"ping_one_connection_id": {
			Type:        schema.TypeString,
			Description: "The ID of the PingOne connection. This field is read only.",
			Computed:    true,
		},
		"ping_one_management_api_endpoint": {
			Type:        schema.TypeString,
			Description: "The PingOne management API endpoint. This field is read only.",
			Computed:    true,
		},
		"region": {
			Type:        schema.TypeString,
			Description: "The region of the PingOne connection. This field is read only.",
			Computed:    true,
		},
	}
}

func resourcePingFederatePingOneConnectionResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).PingOneConnections
	input := pingOneConnections.CreatePingOneConnectionInput{
		Body: *resourcePingFederatePingOneConnectionResourceReadData(d),
	}
	result, _, err := svc.CreatePingOneConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*result.Id)
	return resourcePingFederatePingOneConnectionResourceReadResult(d, result)
}

func resourcePingFederatePingOneConnectionResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).PingOneConnections
	input := pingOneConnections.GetPingOneConnectionInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetPingOneConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read PingOneConnections: %s", err)
	}
	return resourcePingFederatePingOneConnectionResourceReadResult(d, result)
}

func resourcePingFederatePingOneConnectionResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).PingOneConnections
	input := pingOneConnections.UpdatePingOneConnectionInput{
		Id:   d.Id(),
		Body: *resourcePingFederatePingOneConnectionResourceReadData(d),
	}
	result, _, err := svc.UpdatePingOneConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to update PingOneConnections: %s", err)
	}

	return resourcePingFederatePingOneConnectionResourceReadResult(d, result)
}

func resourcePingFederatePingOneConnectionResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).PingOneConnections
	input := pingOneConnections.DeletePingOneConnectionInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeletePingOneConnectionWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete PingOneConnections: %s", err)
	}
	return nil
}

func resourcePingFederatePingOneConnectionResourceReadResult(d *schema.ResourceData, rv *pf.PingOneConnection) diag.Diagnostics {
	var diags diag.Diagnostics
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	setResourceDataStringWithDiagnostic(d, "credential_id", rv.CredentialId, &diags)
	setResourceDataStringWithDiagnostic(d, "ping_one_connection_id", rv.PingOneConnectionId, &diags)
	setResourceDataStringWithDiagnostic(d, "ping_one_management_api_endpoint", rv.PingOneManagementApiEndpoint, &diags)
	setResourceDataBoolWithDiagnostic(d, "active", rv.Active, &diags)
	setResourceDataStringWithDiagnostic(d, "credential", rv.Credential, &diags)
	setResourceDataStringWithDiagnostic(d, "region", rv.Region, &diags)
	setResourceDataStringWithDiagnostic(d, "ping_one_authentication_api_endpoint", rv.PingOneAuthenticationApiEndpoint, &diags)
	setResourceDataStringWithDiagnostic(d, "encrypted_credential", rv.EncryptedCredential, &diags)
	setResourceDataStringWithDiagnostic(d, "organization_name", rv.OrganizationName, &diags)
	setResourceDataStringWithDiagnostic(d, "description", rv.Description, &diags)
	setResourceDataStringWithDiagnostic(d, "creation_date", rv.CreationDate, &diags)
	setResourceDataStringWithDiagnostic(d, "environment_id", rv.EnvironmentId, &diags)
	setResourceDataStringWithDiagnostic(d, "name", rv.Name, &diags)
	return diags
}

func resourcePingFederatePingOneConnectionResourceReadData(d *schema.ResourceData) *pf.PingOneConnection {
	conn := &pf.PingOneConnection{
		Name: String(d.Get("name").(string)),
	}
	if v, ok := d.GetOk("active"); ok {
		conn.Active = Bool(v.(bool))
	}
	if v, ok := d.GetOk("description"); ok {
		conn.Description = String(v.(string))
	}
	if v, ok := d.GetOk("credential"); ok {
		conn.Credential = String(v.(string))
	}
	if v, ok := d.GetOk("encrypted_credential"); ok {
		conn.EncryptedCredential = String(v.(string))
	}

	return conn
}
