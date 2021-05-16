package pingfederate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func issuanceCriteriaShouldFlatten(in *pf.IssuanceCriteria) bool {
	if in.ExpressionCriteria != nil && len(*in.ExpressionCriteria) > 0 {
		return true
	}
	if in.ConditionalCriteria != nil && len(*in.ConditionalCriteria) > 0 {
		return true
	}
	return false
}

func openIdAttributeContractShouldFlatten(in *pf.OpenIdConnectAttributeContract) bool {
	if in.CoreAttributes != nil && len(*in.CoreAttributes) > 0 {
		return true
	}
	if in.ExtendedAttributes != nil && len(*in.ExtendedAttributes) > 0 {
		return true
	}
	return false
}

func spAdapterAttributeContractShouldFlatten(in *pf.SpAdapterAttributeContract) bool {
	if in.CoreAttributes != nil && len(*in.CoreAttributes) > 0 {
		return true
	}
	if in.ExtendedAttributes != nil && len(*in.ExtendedAttributes) > 0 {
		return true
	}
	return false
}

func maskPluginConfigurationFromDescriptor(desc *pf.PluginConfigDescriptor, origConf, conf *pf.PluginConfiguration) []interface{} {

	if origConf.Fields != nil {
		for _, f := range *desc.Fields {
			if *f.Type == "HASHED_TEXT" || ((*f).Encrypted != nil && *f.Encrypted) {
				for _, i := range *conf.Fields {
					if *i.Name == *f.Name {
						s, _ := getConfigFieldValueByName(*i.Name, origConf.Fields)
						i.Value = String(s)
					}
				}
			}
		}
	}

	if origConf.Tables != nil {
		for _, dt := range *desc.Tables {
			for _, dc := range *dt.Columns {
				if *dc.Type == "HASHED_TEXT" || ((*dc).Encrypted != nil && *dc.Encrypted) {
					for ctIndex, ct := range *conf.Tables {
						for crIndex, cr := range *ct.Rows {
							for _, f := range *cr.Fields {
								if *f.Name == *dc.Name {
									val, _ := getConfigFieldValueByName(*f.Name, (*(*origConf.Tables)[ctIndex].Rows)[crIndex].Fields)
									f.Value = &val
								}
							}
						}
					}
				}
			}
		}
	}

	return flattenPluginConfiguration(conf)
}

func getConfigFieldValueByName(name string, fields *[]*pf.ConfigField) (string, error) {
	for _, f := range *fields {
		if *f.Name == name {
			return *f.Value, nil
		}
	}
	return "", nil
}

//func expandPluginConfigurationWithDescriptor(in []interface{}, desc *pf.PluginConfigDescriptor) *pf.PluginConfiguration {
//	log.Printf("[INFO] Expanding config with descriptor")
//	config := expandPluginConfiguration(in)
//	log.Printf("[INFO] We have %d fields before", len(*config.Fields))
//	for _, descriptor := range *desc.Fields {
//		log.Printf("[INFO] Checking field %s", *descriptor.Name)
//		if descriptor.DefaultValue != nil {
//			if !hasField(*descriptor.Name, config) {
//				log.Printf("[INFO] Field %s is required, default is %s", *descriptor.Name, *descriptor.DefaultValue)
//				*config.Fields = append(*config.Fields, &pf.ConfigField{Name: descriptor.Name, Value: descriptor.DefaultValue})
//			}
//		}
//	}
//	log.Printf("[INFO] We have %d fields after", len(*config.Fields))
//	return config
//}

func validateConfiguration(d *schema.ResourceDiff, desc *pf.PluginConfigDescriptor) error {
	var diags diag.Diagnostics
	config := expandPluginConfiguration(d.Get("configuration").([]interface{}))
	for _, descriptor := range *desc.Fields {
		if descriptor.Required != nil && *descriptor.Required {
			if !hasField(*descriptor.Name, config) {
				if descriptor.DefaultValue != nil {
					diags = append(diags, diag.FromErr(fmt.Errorf("the field '%s' is required, its default value is '%s'", *descriptor.Name, *descriptor.DefaultValue))...)
				} else {
					diags = append(diags, diag.FromErr(fmt.Errorf("the field '%s' is required", *descriptor.Name))...)
				}
			}
		}
	}
	if diags.HasError() {
		msgs := []string{
			"configuration validation failed against the class descriptor definition",
		}
		for _, diagnostic := range diags {
			msgs = append(msgs, diagnostic.Summary)
		}
		return fmt.Errorf(strings.Join(msgs, "\n"))
	}
	return nil
}

func hasField(name string, c *pf.PluginConfiguration) bool {
	for _, field := range *c.Fields {
		if *field.Name == name {
			return true
		}
	}
	return false
}
