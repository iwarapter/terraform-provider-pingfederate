package pingfederate

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
)

func resourcePingFederateKeypairSigningResource() *schema.Resource {
	return &schema.Resource{
		Description:   "Provides configuration for Signing KeyPairs within PingFederate.",
		CreateContext: resourcePingFederateKeypairSigningResourceCreate,
		ReadContext:   resourcePingFederateKeypairSigningResourceRead,
		DeleteContext: resourcePingFederateKeypairSigningResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourcePingFederateKeypairSigningResourceImport,
		},
		Schema: resourceKeypairResourceSchema(),
	}
}

func resourcePingFederateKeypairSigningResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	if _, ok := d.GetOk("file_data"); ok {
		input := keyPairsSigning.ImportKeyPairInput{
			Body: pf.KeyPairFile{
				FileData: String(d.Get("file_data").(string)),
				Password: String(d.Get("password").(string)),
			},
		}
		result, _, err := svc.ImportKeyPairWithContext(ctx, &input)
		if err != nil {
			return diag.Errorf("unable to create SigningKeypair: %s", err)
		}

		d.SetId(*result.Id)
		return resourceKeypairResourceReadResult(d, result)
	}
	input := keyPairsSigning.CreateKeyPairInput{
		Body: *resourcePingFederateKeypairResourceReadData(d),
	}
	result, _, err := svc.CreateKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to generate SigningKeypair: %s", err)
	}

	d.SetId(*result.Id)
	return resourceKeypairResourceReadResult(d, result)
}

func resourcePingFederateKeypairSigningResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to read SigningKeypair: %s", err)
	}
	return resourceKeypairResourceReadResult(d, result)
}

func resourcePingFederateKeypairSigningResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsMutexKV.Lock("connection_delete")
	defer awsMutexKV.Unlock("connection_delete")

	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.DeleteKeyPairInput{
		Id: d.Id(),
	}
	_, _, err := svc.DeleteKeyPairWithContext(ctx, &input)
	if err != nil {
		return diag.Errorf("unable to delete SigningKeypair: %s", err)
	}
	return nil
}

func resourcePingFederateKeypairSigningResourceImport(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	svc := m.(pfClient).KeyPairsSigning
	input := keyPairsSigning.GetKeyPairInput{
		Id: d.Id(),
	}
	result, _, err := svc.GetKeyPairWithContext(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve read signing keypair for import %s", err)
	}

	return importKeyPairView(d, result, err)
}

func importKeyPairView(d *schema.ResourceData, result *pf.KeyPairView, err error) ([]*schema.ResourceData, error) {
	diags := resourceKeypairResourceReadResult(d, result)
	//import based on upload
	//TODO unable to properly support upload style imports - https://discuss.hashicorp.com/t/importer-functions-reading-file-config/17624/2

	//import based on generate
	if m, err := extractMatch("CN=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "common_name", String(m), &diags)
	}
	if m, err := extractMatch("OU=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "organization_unit", String(m), &diags)
	}
	if m, err := extractMatch("O=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "organization", String(m), &diags)
	}
	if m, err := extractMatch("L=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "city", String(m), &diags)
	}
	if m, err := extractMatch("ST=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "state", String(m), &diags)
	}
	if m, err := extractMatch("C=([^,]+)", *result.SubjectDN); err == nil {
		setResourceDataStringWithDiagnostic(d, "country", String(m), &diags)
	}
	from, _ := time.Parse(time.RFC3339, *result.ValidFrom)
	expires, _ := time.Parse(time.RFC3339, *result.Expires)
	setResourceDataIntWithDiagnostic(d, "valid_days", Int(int(expires.Sub(from).Hours()/24)), &diags)
	if diags.HasError() {
		return nil, fmt.Errorf("unable to import  %s", err)
	}
	return []*schema.ResourceData{d}, nil
}

func extractMatch(re, source string) (string, error) {
	reg := regexp.MustCompile(re)
	matches := reg.FindStringSubmatch(source)
	if len(matches) == 2 {
		return matches[1], nil
	}
	return "", fmt.Errorf("unable to find match, matches: %v", matches)
}
