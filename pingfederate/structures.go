package pingfederate

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"sort"
	"strings"

	"github.com/hashicorp/go-cty/cty"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func hashcodeString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func setOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func requiredSetOfString() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Required: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func resourceKeypairResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"crypto_provider": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			ValidateDiagFunc: validateCryptoProvider,
		},
		"file_data": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"password": {
			Type:          schema.TypeString,
			Sensitive:     true,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"city": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
		},
		"common_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "organization", "organization_unit", "state", "valid_days"},
		},
		"country": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_algorithm": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Default:       "RSA",
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"key_size": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			RequiredWith: []string{"city", "common_name", "country", "key_algorithm", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
		},
		"organization_unit": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
		},
		"state": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
		},
		"subject_alternative_names": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			ConflictsWith: []string{"file_data", "password"},
		},
		"valid_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"file_data", "password"},
		},
		"expires": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"issuer_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha256_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"serial_number": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sha1_fingerprint": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"signature_algorithm": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_cn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject_dn": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"valid_from": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func resourceKeypairCsrResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"file_data": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
}

func resourceAuthenticationPolicyTreeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
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
				"authentication_api_application_ref": resourceLinkSchema(),
				"enabled": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"root_node": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"action": resourcePolicyActionSchema(),
							"children": {
								Type:     schema.TypeList,
								Optional: true,
								Elem:     resourceAuthenticationPolicyTreeNodeSchemaBuilder(10),
							},
						},
					},
				},
			},
		},
	}
}

// https://github.com/hashicorp/terraform-plugin-sdk/issues/112
// the terraform sdk does not support tree like structure
// to work around this limitation we generate a depth limited
// tree structure, allowing the terraform recursive validators to work
// whilst giving hopefully enough 'depth' to handle our use cases
func resourceAuthenticationPolicyTreeNodeSchemaBuilder(depth int) *schema.Resource {
	if depth <= 0 {
		return &schema.Resource{}
	}
	r := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": resourcePolicyActionSchema(),
			"children": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceAuthenticationPolicyTreeNodeSchemaBuilder(depth - 1),
			},
		},
	}
	return r
}

func resourcePolicyActionSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
						v := value.(string)
						switch v {
						case
							"APC_MAPPING",
							"LOCAL_IDENTITY_MAPPING",
							"AUTHN_SELECTOR",
							"AUTHN_SOURCE",
							"DONE",
							"CONTINUE",
							"RESTART",
							"FRAGMENT":
							return nil
						}
						return diag.Errorf("must be either 'APC_MAPPING' or 'LOCAL_IDENTITY_MAPPING' or 'AUTHN_SELECTOR' or 'AUTHN_SOURCE' or 'DONE' or 'CONTINUE' or 'RESTART' or 'FRAGMENT' not %s", v)
					},
				},
				"context": {
					Type:     schema.TypeString,
					Optional: true,
				},
				"authentication_selector_ref": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,

					//ConflictsWith: []string{
					//	"authentication_policy_contract_ref",
					//	"attribute_mapping",
					//	"local_identity_ref",
					//	"inbound_mapping",
					//	"outbound_attribute_mapping",
					//	"attribute_rules",
					//	"authentication_source",
					//	"input_user_id_mapping",
					//},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"id": {
								Type:     schema.TypeString,
								Required: true,
							},
							"location": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
				"attribute_mapping": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem:     resourceAttributeMapping(),
					//ConflictsWith: []string{
					//	"local_identity_ref",
					//	"inbound_mapping",
					//	"outbound_attribute_mapping",
					//	"authentication_selector_ref",
					//	"attribute_rules",
					//	"authentication_source",
					//	"input_user_id_mapping",
					//},
				},
				"authentication_policy_contract_ref": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					//ConflictsWith: []string{
					//	"local_identity_ref",
					//	"inbound_mapping",
					//	"outbound_attribute_mapping",
					//	"attribute_rules",
					//	"authentication_source",
					//	"input_user_id_mapping",
					//},
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"id": {
								Type:     schema.TypeString,
								Required: true,
							},
							"location": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
				"local_identity_ref": resourceLinkSchema(),
				"inbound_mapping": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem:     resourceAttributeMapping(),
				},
				"outbound_attribute_mapping": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem:     resourceAttributeMapping(),
				},
				"attribute_rules":       resourceAttributeRulesSchema(),
				"authentication_source": resourceAuthenticationSourceSchema(),
				"input_user_id_mapping": {
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"source": resourceSourceTypeIdKey(),
							"value": {
								Type:     schema.TypeString,
								Optional: true,
							},
						},
					},
				},
				"fragment_mapping": {
					Type:     schema.TypeList,
					Optional: true,
					MaxItems: 1,
					Elem:     resourceAttributeMapping(),
					//ConflictsWith: []string{
					//	"local_identity_ref",
					//	"inbound_mapping",
					//	"outbound_attribute_mapping",
					//	"authentication_selector_ref",
					//	"attribute_rules",
					//	"authentication_source",
					//	"input_user_id_mapping",
					//},
				},
				"fragment": resourceLinkSchema(),
			},
		},
	}
}

func resourceAttributeRulesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fallback_to_success": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"items": resourceAttributeRuleSchema(),
			},
		},
	}
}

func resourceAttributeRuleSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"attribute_name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"expected_value": {
					Type:     schema.TypeString,
					Required: true,
				},
				"result": {
					Type:     schema.TypeString,
					Required: true,
				},
				"condition": {
					Type:     schema.TypeString,
					Required: true,
					ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
						v := value.(string)
						opts := []string{
							"EQUALS",
							"EQUALS_CASE_INSENSITIVE",
							"EQUALS_DN",
							"NOT_EQUAL",
							"NOT_EQUAL_CASE_INSENSITIVE",
							"NOT_EQUAL_DN",
							"MULTIVALUE_CONTAINS",
							"MULTIVALUE_CONTAINS_CASE_INSENSITIVE",
							"MULTIVALUE_CONTAINS_DN",
							"MULTIVALUE_DOES_NOT_CONTAIN",
							"MULTIVALUE_DOES_NOT_CONTAIN_CASE_INSENSITIVE",
							"MULTIVALUE_DOES_NOT_CONTAIN_DN",
						}
						sort.Strings(opts)

						i := sort.SearchStrings(opts, v)
						if !(i < len(opts) && opts[i] == v) {
							return diag.Errorf("must be one of '%s' not %s", strings.Join(opts, "' or '"), v)
						}
						return nil
					},
				},
			},
		},
	}
}

func resourceAuthenticationSourceSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
						v := value.(string)
						switch v {
						case
							"IDP_ADAPTER",
							"IDP_CONNECTION":
							return nil
						}
						return diag.Errorf("must be either 'IDP_ADAPTER' or 'IDP_CONNECTION' not %s", v)
					},
				},
				"source_ref": resourceLinkSchema(),
			},
		},
	}
}

func resourceLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"location": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func resourcePluginDescriptorRefSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"location": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func resourceRequiredLinkSchema() *schema.Schema {
	s := resourceLinkSchema()
	s.Required = true
	s.Optional = false
	return s
}

func resourceForceNewLinkSchemaRef() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
				"location": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func resourcePluginConfiguration() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"tables": {
					Type:     schema.TypeList,
					Optional: true,
					Elem:     resourceConfigTable(),
				},
				"fields": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem:     resourceConfigField(),
				},
				"sensitive_fields": {
					Type:     schema.TypeSet,
					Optional: true,
					Elem:     resourceSensitiveConfigField(),
				},
			},
		},
	}
}

func resourceConfigTable() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"rows": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceConfigRow(),
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceConfigRow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			//Requires https://github.com/hashicorp/terraform-plugin-sdk/issues/261
			"default_row": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceConfigField(),
			},
			"sensitive_fields": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceSensitiveConfigField(),
			},
		},
	}
}

func resourceSensitiveConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourcePasswordCredentialValidatorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"core_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceSpAdapterAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"core_attributes": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				//MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
func resourceSpAdapterTargetApplicationInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"application_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_icon_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceIdpAdapterAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"core_attributes": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Elem:     resourceIdpAdapterAttribute(),
			},
			"mask_ognl_values": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceIdpAdapterAttribute(),
			},
		},
	}
}

func resourceIdpAdapterAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of this attribute.",
				Required:    true,
			},
			"pseudonym": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this attribute is used to construct a pseudonym for the SP. Defaults to false.",
				Optional:    true,
				Default:     false,
			},
			"masked": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this attribute is masked in PingFederate logs. Defaults to false.",
				Optional:    true,
			},
		},
	}
}

func resourceIdpAdapterAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     resourceIssuanceCriteria(),
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceLdapAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": resourceRequiredLinkSchema(),
			"base_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_scope": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"search_filter": {
				Type:     schema.TypeString,
				Required: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"binary_attribute_settings": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member_of_nested_group": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceJdbcAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": resourceRequiredLinkSchema(),
			"schema": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"table": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"filter": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCustomAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": resourceRequiredLinkSchema(),
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"filter_fields": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     resourceFieldEntry(),
			},
		},
	}
}

func resourceFieldEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAttributeFulfillmentValue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source": resourceSourceTypeIdKey(),
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func attributeFulfillmentValueHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["key_name"].(string))
	//if d, ok := m["value"]; ok && d.(string) != "" {
	//	buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	//}
	//if d, ok := m["source"]; ok && d.(string) != "" {
	//	buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	//}
	return hashcodeString(buf.String())
}

func resourceSourceTypeIdKey() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					//TODO ValidateFunc:
					// ['TOKEN_EXCHANGE_PROCESSOR_POLICY' or 'ACCOUNT_LINK' or 'ADAPTER' or 'ASSERTION' or 'CONTEXT' or 'CUSTOM_DATA_STORE' or 'EXPRESSION' or 'JDBC_DATA_STORE' or 'LDAP_DATA_STORE' or 'MAPPED_ATTRIBUTES' or 'NO_MAPPING' or 'TEXT' or 'TOKEN' or 'REQUEST' or 'OAUTH_PERSISTENT_GRANT' or 'SUBJECT_TOKEN' or 'ACTOR_TOKEN' or 'PASSWORD_CREDENTIAL_VALIDATOR' or 'IDP_CONNECTION' or 'AUTHENTICATION_POLICY_CONTRACT' or 'CLAIMS' or 'LOCAL_IDENTITY_PROFILE' or 'EXTENDED_CLIENT_METADATA' or 'EXTENDED_PROPERTIES' or 'TRACKED_HTTP_PARAMS']
				},
				"id": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func resourceIssuanceCriteria() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conditional_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceConditionalIssuanceCriteriaEntry(),
			},
			"expression_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceExpressionIssuanceCriteriaEntry(),
			},
		},
	}
}

func resourceConditionalIssuanceCriteriaEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source": resourceSourceTypeIdKey(),
			"attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"condition": {
				Type:     schema.TypeString,
				Required: true,
				//TODO ValidateFunc: //['EQUALS' or 'EQUALS_CASE_INSENSITIVE' or 'EQUALS_DN' or 'NOT_EQUAL' or 'NOT_EQUAL_CASE_INSENSITIVE' or 'NOT_EQUAL_DN' or 'MULTIVALUE_CONTAINS' or 'MULTIVALUE_CONTAINS_CASE_INSENSITIVE' or 'MULTIVALUE_CONTAINS_DN' or 'MULTIVALUE_DOES_NOT_CONTAIN' or 'MULTIVALUE_DOES_NOT_CONTAIN_CASE_INSENSITIVE' or 'MULTIVALUE_DOES_NOT_CONTAIN_DN']
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"error_result": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceExpressionIssuanceCriteriaEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type:     schema.TypeString,
				Required: true,
			},
			"error_result": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func jdbcTagConfigHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["connection_url"].(string))
	if d, ok := m["tags"]; ok && d.(string) != "" {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["default_source"]; ok {
		buf.WriteString(fmt.Sprintf("%t-", d.(bool)))
	}
	return hashcodeString(buf.String())
}

func resourceAuthenticationSelectorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceOpenIdConnectAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of this attribute.",
				Required:    true,
			},
			"override_default_delivery": {
				Type:        schema.TypeBool,
				Description: "This is true when either include in id or user info is true and is used to make the resulting API calls correct",
				Computed:    true,
			},
			"include_in_id_token": {
				Type:        schema.TypeBool,
				Description: "Attribute is included in the ID Token.",
				Optional:    true,
				Default:     false,
			},
			"include_in_user_info": {
				Type:        schema.TypeBool,
				Description: "Attribute is included in the User Info.",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceOpenIdConnectAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     resourceOpenIdConnectAttribute(),
			},
			"extended_attributes": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     resourceOpenIdConnectAttribute(),
			},
		},
	}
}

func resourceAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     resourceIssuanceCriteria(),
			},
		},
	}
}

func attributeRuleSliceHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["condition"].(string))
	buf.WriteString(m["attribute_name"].(string))
	buf.WriteString(m["expected_value"].(string))
	buf.WriteString(m["result"].(string))
	return hashcodeString(buf.String())
}

func resourceParameterValues() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"values": setOfString(),
		},
	}
}

func scopeAttributeMappingsHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	buf.WriteString(m["key_name"].(string))
	return hashcodeString(buf.String())
}

//IdpAttributeQuery - The attribute query profile supports local applications in requesting user attributes from an attribute authority.
func resourceIdpAttributeQuery() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceAttributeQueryNameMapping(),
				Optional: true,
			},
			"policy": {
				Type:     schema.TypeList,
				Elem:     resourceIdpAttributeQueryPolicy(),
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//SpWsTrustAttributeContract - A set of user attributes that this server will send in the token.
func resourceSpWsTrustAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceSpWsTrustAttribute(),
				Optional: true,
			},
			"extended_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceSpWsTrustAttribute(),
				Optional: true,
			},
		},
	}
}

//SpWsTrust - Ws-Trust STS provides security-token validation and creation to extend SSO access to identity-enabled Web Services
func resourceSpWsTrust() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"abort_if_not_fulfilled_from_request": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceSpWsTrustAttributeContract(),
				Required: true,
			},
			"default_token_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypt_saml2_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"generate_key": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"message_customizations": {
				Type:     schema.TypeList,
				Elem:     resourceProtocolMessageCustomization(),
				Optional: true,
			},
			"minutes_after": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"minutes_before": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"o_auth_assertion_profiles": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"partner_service_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"request_contract_ref": resourceRequiredLinkSchema(),
			"token_processor_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceIdpTokenProcessorMapping(),
				Required: true,
			},
		},
	}
}

//IdpSsoServiceEndpoint - The settings that define an endpoint to an IdP SSO service.
func resourceIdpSsoServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//InboundBackChannelAuth
func resourceInboundBackChannelAuth() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certs": {
				Type:     schema.TypeList,
				Elem:     resourceConnectionCert(),
				Optional: true,
			},
			"digital_signature": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"http_basic_credentials": {
				Type:     schema.TypeList,
				Elem:     resourceUsernamePasswordCredentials(),
				Optional: true,
			},
			"require_ssl": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"verification_issuer_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"verification_subject_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//EncryptionPolicy - Defines what to encrypt in the browser-based SSO profile.
func resourceEncryptionPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encrypt_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encrypt_slo_subject_name_id": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"encrypted_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"slo_subject_name_id_encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//X509File - Encoded certificate data.
func resourceX509File() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crypto_provider": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_data": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					eq := strings.ReplaceAll(old, "\n", "") == strings.ReplaceAll(new, "\n", "")
					return eq
				},
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

//AdditionalAllowedEntitiesConfiguration - Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.
func resourceAdditionalAllowedEntitiesConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_allowed_entities": {
				Type:     schema.TypeList,
				Elem:     resourceEntity(),
				Optional: true,
			},
			"allow_additional_entities": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_all_entities": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//Entity
func resourceEntity() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//IdpWsTrustAttribute - An attribute for the Ws-Trust attribute contract.
func resourceIdpWsTrustAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//ConnectionCert - A certificate used for signature verification or XML encryption.
func resourceConnectionCert() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active_verification_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cert_view": {
				Type:     schema.TypeList,
				Elem:     resourceCertView(),
				Optional: true,
				Computed: true,
			},
			"encryption_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"primary_verification_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"secondary_verification_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"x509_file": {
				Type:     schema.TypeList,
				Elem:     resourceX509File(),
				Required: true,
			},
		},
	}
}

//IdpBrowserSso - The settings used to enable secure browser-based SSO to resources at your site.
func resourceIdpBrowserSso() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceSpAdapterMapping(),
				Optional: true,
			},
			"artifact": {
				Type:     schema.TypeList,
				Elem:     resourceArtifactSettings(),
				Optional: true,
			},
			"assertions_signed": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceIdpBrowserSsoAttributeContract(),
				Optional: true,
			},
			"authentication_policy_contract_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceAuthenticationPolicyContractMapping(),
				Optional: true,
			},
			"authn_context_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceAuthnContextMapping(),
				Optional: true,
			},
			"decryption_policy": {
				Type:     schema.TypeList,
				Elem:     resourceDecryptionPolicy(),
				Optional: true,
			},
			"default_target_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled_profiles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"idp_identity_mapping": {
				Type:     schema.TypeString,
				Required: true,
			},
			"incoming_bindings": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"message_customizations": {
				Type:     schema.TypeList,
				Elem:     resourceProtocolMessageCustomization(),
				Optional: true,
			},
			"oauth_authentication_policy_contract_ref": resourceRequiredLinkSchema(),
			"oidc_provider_settings": {
				Type:     schema.TypeList,
				Elem:     resourceOIDCProviderSettings(),
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sign_authn_requests": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"slo_service_endpoints": {
				Type:     schema.TypeList,
				Elem:     resourceSloServiceEndpoint(),
				Optional: true,
			},
			"sso_o_auth_mapping": {
				Type:     schema.TypeList,
				Elem:     resourceSsoOAuthMapping(),
				Optional: true,
			},
			"sso_service_endpoints": {
				Type:     schema.TypeList,
				Elem:     resourceIdpSsoServiceEndpoint(),
				Optional: true,
			},
			"url_whitelist_entries": {
				Type:     schema.TypeList,
				Elem:     resourceUrlWhitelistEntry(),
				Optional: true,
			},
		},
	}
}

//ContactInfo - Contact information.
func resourceContactInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"company": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//SpAdapterMapping - A mapping to a SP adapter.
func resourceSpAdapterMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_override_settings": {
				Type:     schema.TypeList,
				Elem:     resourceSpAdapter(),
				Optional: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restrict_virtual_entity_ids": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"sp_adapter_ref": resourceRequiredLinkSchema(),
		},
	}
}

//ChangeDetectionSettings - Setting to detect changes to a user or a group.
func resourceChangeDetectionSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"changed_users_algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_object_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_stamp_attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_object_class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usn_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//Schema - Custom SCIM Attributes configuration.
func resourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attributes": {
				Type:     schema.TypeList,
				Elem:     resourceSchemaAttribute(),
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//IdpAdapter - An IdP adapter instance.
func resourceIdpAdapter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceIdpAdapterAttributeContract(),
				Optional: true,
			},
			"attribute_mapping": {
				Type:     schema.TypeList,
				Elem:     resourceIdpAdapterContractMapping(),
				Optional: true,
			},
			"authn_ctx_class_ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"configuration": resourcePluginConfiguration(),
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_ref":            resourceRequiredLinkSchema(),
			"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
		},
	}
}

//GroupMembershipDetection - Settings to detect group memberships.
func resourceGroupMembershipDetection() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_member_attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"member_of_group_attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//SaasAttributeMapping - Settings to map the source record attributes to target attributes.
func resourceSaasAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"field_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"saas_field_info": {
				Type:     schema.TypeList,
				Elem:     resourceSaasFieldConfiguration(),
				Required: true,
			},
		},
	}
}

//SpAttributeQueryPolicy - The attribute query profile's security policy.
func resourceSpAttributeQueryPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encrypt_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_encrypted_name_id": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_signed_attribute_query": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sign_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sign_response": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//SpBrowserSso - The SAML settings used to enable secure browser-based SSO to resources at your partner's site.
func resourceSpBrowserSso() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_mappings": {
				Type: schema.TypeList,
				Elem: resourceIdpAdapterAssertionMapping(),
				//Required: true,
				Optional: true,
			},
			"artifact": {
				Type:     schema.TypeList,
				Elem:     resourceArtifactSettings(),
				Optional: true,
			},
			"assertion_lifetime": {
				Type: schema.TypeList,
				Elem: resourceAssertionLifetime(),
				//Required: true,
				Optional: true,
			},
			"attribute_contract": {
				Type: schema.TypeList,
				Elem: resourceSpBrowserSsoAttributeContract(),
				//Required: true,
				Optional: true,
			},
			"authentication_policy_contract_assertion_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceAuthenticationPolicyContractAssertionMapping(),
				Optional: true,
			},
			"default_target_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled_profiles": setOfString(),
			"encryption_policy": {
				Type: schema.TypeList,
				Elem: resourceEncryptionPolicy(),
				//Required: true,
				Optional: true,
			},
			"incoming_bindings": setOfString(),
			"message_customizations": {
				Type:     schema.TypeList,
				Elem:     resourceProtocolMessageCustomization(),
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"require_signed_authn_requests": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sign_assertions": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sign_response_as_required": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"slo_service_endpoints": {
				Type:     schema.TypeList,
				Elem:     resourceSloServiceEndpoint(),
				Optional: true,
			},
			"sp_saml_identity_mapping": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sp_ws_fed_identity_mapping": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sso_service_endpoints": {
				Type:     schema.TypeList,
				Elem:     resourceSpSsoServiceEndpoint(),
				Required: true,
			},
			"url_whitelist_entries": {
				Type:     schema.TypeList,
				Elem:     resourceUrlWhitelistEntry(),
				Optional: true,
			},
			"ws_fed_token_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ws_trust_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//OIDCProviderSettings - The OpenID Provider settings.
func resourceOIDCProviderSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"authentication_scheme": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authentication_signing_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authorization_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jwks_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"login_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"request_parameters": {
				Type:     schema.TypeList,
				Elem:     resourceOIDCRequestParameter(),
				Optional: true,
			},
			"request_signing_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scopes": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_info_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//UsernamePasswordCredentials - Username and password credentials.
func resourceUsernamePasswordCredentials() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encrypted_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//DecryptionPolicy - Defines what to decrypt in the browser-based SSO profile.
func resourceDecryptionPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"assertion_encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"attributes_encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"slo_encrypt_subject_name_id": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"slo_subject_name_id_encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"subject_name_id_encrypted": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//SpWsTrustAttribute - An attribute for the Ws-Trust attribute contract.
func resourceSpWsTrustAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//IdpWsTrust - Ws-Trust STS provides validation of incoming tokens which enable SSO access to Web Services. It also allows generation of local tokens for Web Services.
func resourceIdpWsTrust() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceIdpWsTrustAttributeContract(),
				Required: true,
			},
			"generate_local_token": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"token_generator_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceSpTokenGeneratorMapping(),
				Optional: true,
			},
		},
	}
}

//Channel - A channel is a combination of a source data store and a provisioning target. It include settings of a source data store, managing provisioning threads and mapping of attributes.
func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"attribute_mapping": {
				Type:     schema.TypeList,
				Elem:     resourceSaasAttributeMapping(),
				Required: true,
			},
			"channel_source": {
				Type:     schema.TypeList,
				Elem:     resourceChannelSource(),
				Required: true,
			},
			"max_threads": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

//ArtifactResolverLocation - The remote party URLs to resolve the artifact.
func resourceArtifactResolverLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//AttributeQueryNameMapping - The attribute query name mappings between the SP and the IdP.
func resourceAttributeQueryNameMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//IdpWsTrustAttributeContract - A set of user attributes that this server will receive in the token.
func resourceIdpWsTrustAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpWsTrustAttribute(),
				Optional: true,
			},
			"extended_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpWsTrustAttribute(),
				Optional: true,
			},
		},
	}
}

//IdpAdapterContractMapping
func resourceIdpAdapterContractMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"inherited": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
		},
	}
}

//SloServiceEndpoint - Where SLO logout messages are sent. Only applicable for SAML 2.0.
func resourceSloServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": {
				Type:     schema.TypeString,
				Required: true,
			},
			"response_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//IdpOAuthGrantAttributeMapping - The OAuth Assertion Grant settings used to map from your IdP.
func resourceIdpOAuthGrantAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_token_manager_mappings": {
				Type:     schema.TypeList,
				Elem:     resourceAccessTokenManagerMapping(),
				Optional: true,
			},
			"idp_o_auth_attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceIdpOAuthAttributeContract(),
				Optional: true,
			},
		},
	}
}

//AuthenticationPolicyContractMapping - An Authentication Policy Contract mapping into IdP Connection.
func resourceAuthenticationPolicyContractMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"authentication_policy_contract_ref": resourceRequiredLinkSchema(),
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restrict_virtual_server_ids": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restricted_virtual_server_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

//ConnectionMetadataUrl - Configuration settings to enable automatic reload of partner's metadata.
func resourceConnectionMetadataUrl() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_auto_metadata_update": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"metadata_url_ref": resourceRequiredLinkSchema(),
		},
	}
}

//ProtocolMessageCustomization - The message customization that will be executed on outgoing PingFederate messages.
func resourceProtocolMessageCustomization() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"context_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"message_expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//OutboundProvision - Outbound Provisioning allows an IdP to create and maintain user accounts at standards-based partner sites using SCIM as well as select-proprietary provisioning partner sites that are protocol-enabled.
func resourceOutboundProvision() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"channels": {
				Type:     schema.TypeList,
				Elem:     resourceChannel(),
				Required: true,
			},
			"custom_schema": {
				Type:     schema.TypeList,
				Elem:     resourceSchema(),
				Optional: true,
			},
			"target_settings": {
				Type:     schema.TypeSet,
				Elem:     resourceConfigField(),
				Required: true,
			},
			"sensitive_target_settings": {
				Type:     schema.TypeSet,
				Elem:     resourceConfigField(),
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//SpTokenGeneratorMapping - The SP Token Generator Mapping.
func resourceSpTokenGeneratorMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"default_mapping": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"sp_token_generator_ref": resourceRequiredLinkSchema(),
		},
	}
}

//AuthenticationPolicyContractAssertionMapping - The Authentication Policy Contract Assertion Mapping.
func resourceAuthenticationPolicyContractAssertionMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"abort_sso_transaction_as_fail_safe": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"authentication_policy_contract_ref": resourceRequiredLinkSchema(),
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restrict_virtual_entity_ids": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

//OIDCRequestParameter - An OIDC custom request parameter.
func resourceOIDCRequestParameter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application_endpoint_override": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//ChannelSource - The source data source and LDAP settings.
func resourceChannelSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account_management_settings": {
				Type:     schema.TypeList,
				Elem:     resourceAccountManagementSettings(),
				Required: true,
			},
			"base_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"change_detection_settings": {
				Type:     schema.TypeList,
				Elem:     resourceChangeDetectionSettings(),
				Required: true,
			},
			"data_source": resourceRequiredLinkSchema(),
			"group_membership_detection": {
				Type:     schema.TypeList,
				Elem:     resourceGroupMembershipDetection(),
				Required: true,
			},
			"group_source_location": {
				Type:     schema.TypeList,
				Elem:     resourceChannelSourceLocation(),
				Optional: true,
			},
			"guid_attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"guid_binary": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"user_source_location": {
				Type:     schema.TypeList,
				Elem:     resourceChannelSourceLocation(),
				Required: true,
			},
		},
	}
}

//ChannelSourceLocation - The location settings that includes a DN and a LDAP filter.
func resourceChannelSourceLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nested_search": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//OIDCClientCredentials - The OpenID Connect Client Credentials settings. This is required for an OIDC Connection.
func resourceOIDCClientCredentials() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encrypted_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//AuthnContextMapping - The authentication context mapping between local and remote values.
func resourceAuthnContextMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"local": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//OutboundBackChannelAuth
func resourceOutboundBackChannelAuth() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"digital_signature": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"http_basic_credentials": {
				Type:     schema.TypeList,
				Elem:     resourceUsernamePasswordCredentials(),
				Optional: true,
			},
			"ssl_auth_key_pair_ref": resourceLinkSchema(),
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"validate_partner_cert": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//SpAttributeQuery - The attribute query profile supports SPs in requesting user attributes.
func resourceSpAttributeQuery() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     resourceAttributeFulfillmentValue(),
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"policy": {
				Type:     schema.TypeList,
				Elem:     resourceSpAttributeQueryPolicy(),
				Optional: true,
			},
		},
	}
}

//CertView - Certificate details.
func resourceCertView() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crypto_provider": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expires": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"issuer_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sha1_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sha256_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subject_alternative_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"subject_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"valid_from": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

//AccessTokenManagerMapping - A mapping in a connection that defines how access tokens are created.
func resourceAccessTokenManagerMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_token_manager_ref": resourceRequiredLinkSchema(),
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
		},
	}
}

//IdpAdapterAssertionMapping - The IdP Adapter Assertion Mapping.
func resourceIdpAdapterAssertionMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"abort_sso_transaction_as_fail_safe": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"adapter_override_settings": {
				Type:     schema.TypeList,
				Elem:     resourceIdpAdapter(),
				Optional: true,
			},
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"idp_adapter_ref": resourceRequiredLinkSchema(),
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restrict_virtual_entity_ids": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

//IdpOAuthAttributeContract - A set of user attributes that the IdP sends in the OAuth Assertion Grant.
func resourceIdpOAuthAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpBrowserSsoAttribute(),
				Optional: true,
			},
			"extended_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpBrowserSsoAttribute(),
				Optional: true,
			},
		},
	}
}

//IdpBrowserSsoAttributeContract - A set of user attributes that the IdP sends in the SAML assertion.
func resourceIdpBrowserSsoAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpBrowserSsoAttribute(),
				Optional: true,
			},
			"extended_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceIdpBrowserSsoAttribute(),
				Optional: true,
			},
		},
	}
}

//AccountManagementSettings - Account management settings.
func resourceAccountManagementSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account_status_algorithm": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_status_attribute_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flag_comparison_status": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flag_comparison_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//SsoOAuthMapping - IdP Browser SSO OAuth Attribute Mapping
func resourceSsoOAuthMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
		},
	}
}

//IdpTokenProcessorMapping - The IdP Token Processor Mapping.
func resourceIdpTokenProcessorMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract_fulfillment": {
				Type:     schema.TypeSet,
				Elem:     resourceAttributeFulfillmentValue(),
				Required: true,
			},
			"ldap_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceCustomAttributeSource(),
			},
			"idp_token_processor_ref": resourceRequiredLinkSchema(),
			"issuance_criteria": {
				Type:     schema.TypeList,
				Elem:     resourceIssuanceCriteria(),
				Optional: true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

//SpBrowserSsoAttribute - An attribute for the SP Browser SSO attribute contract.
func resourceSpBrowserSsoAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_format": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//ArtifactSettings - The settings for an Artifact binding.
func resourceArtifactSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"lifetime": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"resolver_locations": {
				Type:     schema.TypeList,
				Elem:     resourceArtifactResolverLocation(),
				Required: true,
			},
			"source_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//IdpAttributeQueryPolicy - The attribute query profile's security policy.
func resourceIdpAttributeQueryPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"encrypt_name_id": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mask_attribute_values": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_encrypted_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_signed_assertion": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_signed_response": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sign_attribute_query": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//SpAdapter - An SP adapter instance.
func resourceSpAdapter() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_contract": {
				Type:     schema.TypeList,
				Elem:     resourceSpAdapterAttributeContract(),
				Optional: true,
			},
			"configuration": resourcePluginConfiguration(),
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_ref":            resourceRequiredLinkSchema(),
			"plugin_descriptor_ref": resourcePluginDescriptorRefSchema(),
			"target_application_info": {
				Type:     schema.TypeList,
				Elem:     resourceSpAdapterTargetApplicationInfo(),
				Optional: true,
			},
		},
	}
}

//SchemaAttribute - A custom SCIM attribute.
func resourceSchemaAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"multi_valued": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sub_attributes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

//SigningSettings - Settings related to signing messages sent to this partner.
func resourceSigningSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_cert_in_signature": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_raw_key_in_signature": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"signing_key_pair_ref": resourceRequiredLinkSchema(),
		},
	}
}

//SaasFieldConfiguration - The settings that represent how attribute values from source data store will be mapped into Fields specified by the service provider.
func resourceSaasFieldConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"character_case": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"create_only": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"masked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parser": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trim": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

//SpBrowserSsoAttributeContract - A set of user attributes that the IdP sends in the SAML assertion.
func resourceSpBrowserSsoAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"core_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceSpBrowserSsoAttribute(),
				Optional: true,
			},
			"extended_attributes": {
				Type:     schema.TypeList,
				Elem:     resourceSpBrowserSsoAttribute(),
				Optional: true,
			},
		},
	}
}

//AssertionLifetime - The timeframe of validity before and after the issuance of the assertion.
func resourceAssertionLifetime() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"minutes_after": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"minutes_before": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

//IdpBrowserSsoAttribute - An attribute for the IdP Browser SSO attribute contract.
func resourceIdpBrowserSsoAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"masked": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

//UrlWhitelistEntry - Url domain and path to be used as whitelist in WS-Federation connection
func resourceUrlWhitelistEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_query_and_fragment": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"require_https": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"valid_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"valid_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//ConnectionCredentials - The certificates and settings for encryption, signing, and signature verification.
func resourceConnectionCredentials() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_encryption_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"certs": {
				Type:     schema.TypeList,
				Elem:     resourceConnectionCert(),
				Optional: true,
			},
			"decryption_key_pair_ref": resourceLinkSchema(),
			"inbound_back_channel_auth": {
				Type:     schema.TypeList,
				Elem:     resourceInboundBackChannelAuth(),
				Optional: true,
			},
			"key_transport_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"outbound_back_channel_auth": {
				Type:     schema.TypeList,
				Elem:     resourceOutboundBackChannelAuth(),
				Optional: true,
			},
			"secondary_decryption_key_pair_ref": resourceLinkSchema(),
			"signing_settings": {
				Type:     schema.TypeList,
				Elem:     resourceSigningSettings(),
				Optional: true,
			},
			"verification_issuer_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"verification_subject_dn": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//SpSsoServiceEndpoint - The settings that define a service endpoint to a SP SSO service.
func resourceSpSsoServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": {
				Type:     schema.TypeString,
				Required: true,
			},
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"is_default": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
