package pingfederate

//lint:file-ignore SA1019 Ignore deprecated GetOkExists - no current alternative

import (
	"context"

	"github.com/iwarapter/pingfederate-sdk-go/services/extendedProperties"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateExtendedPropertiesResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePingFederateExtendedPropertiesResourceCreate,
		ReadContext:   resourcePingFederateExtendedPropertiesResourceRead,
		UpdateContext: resourcePingFederateExtendedPropertiesResourceUpdate,
		DeleteContext: resourcePingFederateExtendedPropertiesResourceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"property": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"multi_valued": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
		},
	}
}

func resourcePingFederateExtendedPropertiesResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("extended_properties")
	return resourcePingFederateExtendedPropertiesResourceUpdate(ctx, d, m)
}

func resourcePingFederateExtendedPropertiesResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ExtendedProperties
	result, _, err := svc.GetExtendedPropertiesWithContext(ctx)
	if err != nil {
		return diag.Errorf("unable to read ExtendedProperties: %s", err)
	}
	return resourcePingFederateExtendedPropertiesResourceReadResult(d, result)
}

func resourcePingFederateExtendedPropertiesResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	properties := &pf.ExtendedProperties{
		Items: &[]*pf.ExtendedProperty{},
	}

	if v, ok := d.GetOk("property"); ok {
		properties.Items = expandExtendedProperties(v.(*schema.Set).List())
	}

	svc := m.(pfClient).ExtendedProperties
	input := &extendedProperties.UpdateExtendedPropertiesInput{
		Body: *properties,
	}

	result, _, err := svc.UpdateExtendedPropertiesWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to update ExtendedProperties: %s", err)
	}
	return resourcePingFederateExtendedPropertiesResourceReadResult(d, result)
}

func resourcePingFederateExtendedPropertiesResourceDelete(ctx context.Context, _ *schema.ResourceData, m interface{}) diag.Diagnostics {
	svc := m.(pfClient).ExtendedProperties
	input := &extendedProperties.UpdateExtendedPropertiesInput{
		Body: pf.ExtendedProperties{},
	}
	_, _, err := svc.UpdateExtendedPropertiesWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("unable to delete ExtendedProperties: %s", err)
	}
	return nil
}

func resourcePingFederateExtendedPropertiesResourceReadResult(d *schema.ResourceData, rv *pf.ExtendedProperties) diag.Diagnostics {
	var diags diag.Diagnostics
	if rv.Items != nil && len(*rv.Items) > 0 {
		if err := d.Set("property", flattenExtendedProperties(*rv.Items)); err != nil {
			diags = append(diags, diag.FromErr(err)...)
		}
	}
	return diags
}
