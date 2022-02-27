package sdkv2provider

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

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

func resourceKeypairResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.",
		},
		"crypto_provider": {
			Type:             schema.TypeString,
			Optional:         true,
			ForceNew:         true,
			Description:      "Cryptographic Provider.  This is only applicable if Hybrid HSM mode is true.",
			ValidateDiagFunc: validateCryptoProvider,
		},
		"file_data": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "Base64 encoded PKCS12 file data. New line characters should be omitted or encoded in this value.",
			ValidateFunc:  validation.StringIsBase64,
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"password": {
			Type:          schema.TypeString,
			Sensitive:     true,
			Optional:      true,
			ForceNew:      true,
			Description:   "Password for the PKCS12 file.",
			ConflictsWith: []string{"city", "common_name", "country", "key_algorithm", "key_size", "organization", "organization_unit", "state", "valid_days"},
			RequiredWith:  []string{"file_data", "password"},
		},
		"city": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "City.",
			ConflictsWith: []string{"file_data", "password"},
		},
		"common_name": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "Common name for key pair subject.",
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"common_name", "country", "key_algorithm", "organization", "valid_days"},
		},
		"country": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "Country.",
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"common_name", "country", "key_algorithm", "organization", "valid_days"},
		},
		"key_algorithm": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Default:       "RSA",
			Description:   "Key generation algorithm. Supported algorithms are available through the /keyPairs/keyAlgorithms endpoint.",
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"common_name", "country", "key_algorithm", "organization", "valid_days"},
		},
		"key_size": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			Description:  "Key size, in bits. If this property is unset, the default size for the key algorithm will be used. Supported key sizes are available through the /keyPairs/keyAlgorithms endpoint.",
			RequiredWith: []string{"common_name", "country", "key_algorithm", "organization", "valid_days"},
		},
		"organization": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "Organization.",
			ConflictsWith: []string{"file_data", "password"},
			RequiredWith:  []string{"common_name", "country", "key_algorithm", "organization", "valid_days"},
		},
		"organization_unit": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "Organization unit.",
			ConflictsWith: []string{"file_data", "password"},
		},
		"state": {
			Type:          schema.TypeString,
			Optional:      true,
			ForceNew:      true,
			Description:   "State.",
			ConflictsWith: []string{"file_data", "password"},
		},
		"subject_alternative_names": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The subject alternative names (SAN).",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
			ConflictsWith: []string{"file_data", "password"},
		},
		"valid_days": {
			Type:          schema.TypeInt,
			Optional:      true,
			ForceNew:      true,
			Description:   "Number of days the key pair will be valid for.",
			ConflictsWith: []string{"file_data", "password"},
		},
		"expires": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The end date up until which the item is valid, in ISO 8601 format (UTC).",
		},
		"issuer_dn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The issuer's distinguished name.",
		},
		"sha256_fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SHA-256 fingerprint in Hex encoding.",
		},
		"serial_number": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The serial number assigned by the CA.",
		},
		"sha1_fingerprint": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "SHA-1 fingerprint in Hex encoding.",
		},
		"signature_algorithm": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Signature algorithm. If this property is unset, the default signature algorithm for the key algorithm will be used. Supported signature algorithms are available through the /keyPairs/keyAlgorithms endpoint.",
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Status of the item.",
		},
		"subject_dn": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The subject's distinguished name.",
		},
		"valid_from": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The start date from which the item is valid, in ISO 8601 format (UTC).",
		},
		"version": {
			Type:        schema.TypeInt,
			Computed:    true,
			Description: "The X.509 version to which the item conforms.",
		},
	}
}

func resourceKeypairCsrResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"keypair_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "ID of the key pair.",
		},
		"file_data": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The CSR response file data in PKCS7 format or as an X.509 certificate. PEM encoding (with or without the header and footer lines) is required. New line characters should be omitted or encoded in this value.",
		},
	}
}

func resourceAuthenticationPolicyTreeResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Description: "A description for the authentication policy.",
				Optional:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Whether or not this authentication policy tree is enabled. Default is true.",
				Optional:    true,
				Default:     true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The authentication policy name. Name is unique.",
				Optional:    true,
			},
			"authentication_api_application_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Authentication API Application Id to be used in this policy branch. If the value is not specified, no Authentication API Application will be used.",
				Elem:        resourceLinkResource(),
			},
			"root_node": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A node inside the authentication policy tree.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": resourcePolicyActionSchema(),
						"children": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The nodes inside the authentication policy tree node.",
							Elem:        resourceAuthenticationPolicyTreeNodeSchemaBuilder(10),
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
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The nodes inside the authentication policy tree node.",
				Elem:        resourceAuthenticationPolicyTreeNodeSchemaBuilder(depth - 1),
			},
		},
	}
	return r
}

func resourcePolicyActionSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		Description: "The result action.",
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The authentication selection type.",
					ValidateFunc: validation.StringInSlice([]string{
						"APC_MAPPING",
						"LOCAL_IDENTITY_MAPPING",
						"AUTHN_SELECTOR",
						"AUTHN_SOURCE",
						"DONE",
						"CONTINUE",
						"RESTART",
						"FRAGMENT",
					}, false),
				},
				"context": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The result context.",
				},
				"authentication_selector_ref": {
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
				},
				"attribute_mapping": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Description: "Contract fulfillment with the authentication policy contract's default values, and additional attributes retrieved from local data stores.",
					Elem:        resourceAttributeMapping(),
				},
				"authentication_policy_contract_ref": {
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
				},
				"local_identity_ref": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem:        resourceLinkResource(),
					Description: "Reference to the associated local identity profile.",
				},
				"inbound_mapping": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem:        resourceAttributeMapping(),
					Description: "Inbound mappings into the local identity profile fields.",
				},
				"outbound_attribute_mapping": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem:        resourceAttributeMapping(),
					Description: "Authentication policy contract mappings associated with this local Identity profile.",
				},
				"attribute_rules":       resourceAttributeRulesSchema(),
				"authentication_source": resourceAuthenticationSourceSchema(),
				"input_user_id_mapping": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "The input user ID mapping.",
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
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem:        resourceAttributeMapping(),
					Description: "The fragment mapping for attributes to be passed into the authentication fragment.",
				},
				"fragment": {
					Type:        schema.TypeList,
					Optional:    true,
					MaxItems:    1,
					Elem:        resourceLinkResource(),
					Description: "Reference to the associated authentication fragment.",
				},
			},
		},
	}
}

func resourceAttributeRulesSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "The authentication policy rules.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"fallback_to_success": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "When all the rules fail, you may choose to default to the general success action or fail. Default to success.",
				},
				"items": resourceAttributeRuleSchema(),
			},
		},
	}
}

func resourceAttributeRuleSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Authentication policy rules using attributes from the previous authentication source. Each rule is evaluated to determine the next action in the policy.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"attribute_name": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The name of the attribute to use in this attribute rule.",
				},
				"expected_value": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The expected value of this attribute rule.",
				},
				"result": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The result of this attribute rule.",
				},
				"condition": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "The condition that will be applied to the attribute's expected value.",
					ValidateFunc: validation.StringInSlice([]string{
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
					}, false),
				},
			},
		},
	}
}

func resourceAuthenticationSourceSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		MaxItems:    1,
		Optional:    true,
		Description: "The associated authentication source.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:     schema.TypeString,
					Required: true,
					ValidateFunc: validation.StringInSlice([]string{
						"IDP_ADAPTER",
						"IDP_CONNECTION",
					}, false),
				},
				"source_ref": resourceLinkSchema(),
			},
		},
	}
}

func resourceLinkResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the resource.",
			},
			"location": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A read-only URL that references the resource. If the resource is not currently URL-accessible, this property will be null.",
			},
		},
	}
}

func resourceLinkSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem:     resourceLinkResource(),
	}
}

func resourcePluginDescriptorRefSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Description: "Reference to the plugin descriptor for this instance. The plugin descriptor cannot be modified once the instance is created.\nNote: Ignored when specifying a connection's adapter override.",
		Elem:        resourceLinkResource(),
	}
}

func resourceRequiredLinkSchema() *schema.Schema {
	s := resourceLinkSchema()
	s.Required = true
	s.Optional = false
	return s
}

func resourcePluginConfiguration() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Description: "Plugin instance configuration.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"tables": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "List of configuration tables.",
					Elem:        resourceConfigTable(),
				},
				"fields": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "List of configuration fields.",
					Elem:        resourceConfigField(),
				},
				"sensitive_fields": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "List of sensitive configuration fields.",
					Elem:        resourceSensitiveConfigField(),
				},
			},
		},
	}
}

func resourceConfigTable() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the table.",
			},
			"rows": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of table rows.",
				Elem:        resourceConfigRow(),
			},
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this table is inherited from its parent instance. If true, the rows become read-only. The default value is false.",
			},
		},
	}
}

func resourceConfigRow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			//Requires https://github.com/hashicorp/terraform-plugin-sdk/issues/261
			"default_row": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this row is the default.",
			},
			"fields": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of configuration fields.",
				Elem:        resourceConfigField(),
			},
			"sensitive_fields": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of sensitive configuration fields.",
				Elem:        resourceSensitiveConfigField(),
			},
		},
	}
}

func resourceSensitiveConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the configuration field.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The value for the configuration field. For encrypted or hashed fields, GETs will not return this attribute. To update an encrypted or hashed field, specify the new value in this attribute.",
			},
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this field is inherited from its parent instance. If true, the value/encrypted value properties become read-only. The default value is false.",
			},
		},
	}
}

func resourceConfigField() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the configuration field.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value for the configuration field.",
			},
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this field is inherited from its parent instance. If true, the value/encrypted value properties become read-only. The default value is false.",
			},
		},
	}
}

func resourcePasswordCredentialValidatorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
			"core_attributes": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of read-only attributes that are automatically populated by the password credential validator descriptor.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extended_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				MinItems:    1,
				Description: "A list of additional attributes that can be returned by the password credential validator. The extended attributes are only used if the adapter supports them.",
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
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
			"core_attributes": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of read-only attributes that are automatically populated by the SP adapter descriptor.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"extended_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of additional attributes that can be returned by the SP adapter. The extended attributes are only used if the adapter supports them.",
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
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Specifies Whether target application information is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
			"application_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The application name.",
			},
			"application_icon_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The application icon URL.",
			},
		},
	}
}

func resourceIdpAdapterAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
			"core_attributes": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "A list of IdP adapter attributes that correspond to the attributes exposed by the IdP adapter type.",
				Elem:        resourceIdpAdapterAttribute(),
			},
			"mask_ognl_values": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false.",
			},
			"extended_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of additional attributes that can be returned by the IdP adapter. The extended attributes are only used if the adapter supports them.",
				Elem:        resourceIdpAdapterAttribute(),
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
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured ldap data stores to look up attributes from.",
				Elem:        resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured jdbc data stores to look up attributes from.",
				Elem:        resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured custom data stores to look up attributes from.",
				Elem:        resourceCustomAttributeSource(),
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of mappings from attribute names to their fulfillment values.",
				Elem:        resourceAttributeFulfillmentValue(),
			},
			"issuance_criteria": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
				Elem:        resourceIssuanceCriteria(),
			},
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this attribute mapping is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
		},
	}
}

func resourceLdapAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Reference to the associated data store.",
				Elem:        resourceLinkResource(),
			},
			"base_dn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The base DN to search from. If not specified, the search will start at the LDAP's root.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.",
			},
			"search_scope": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Determines the node depth of the query.",
				ValidateFunc: validation.StringInSlice([]string{"OBJECT", "ONE_LEVEL", "SUBTREE"}, false),
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.<br>Note: Required for APC-to-SP Adapter Mappings",
			},
			"search_filter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The LDAP filter that will be used to lookup the objects from the directory.",
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings",
				Elem:        resourceAttributeFulfillmentValue(),
			},
			"binary_attribute_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The advanced settings for binary LDAP attributes.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"member_of_nested_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set this to true to return transitive group memberships for the 'memberOf' attribute.  This only applies for Active Directory data sources.  All other data sources will be set to false.",
			},
		},
	}
}

func resourceJdbcAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Reference to the associated data store.",
				Elem:        resourceLinkResource(),
			},
			"schema": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID that defines this attribute source. Only alphanumeric characters allowed.<br>Note: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.",
			},
			"table": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the database table. The name is used to construct the SQL query to retrieve data from the data store.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Lists the table structure that stores information within a database. Some databases, such as Oracle, require a schema for a JDBC query. Other databases, such as MySQL, do not require a schema.",
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings",
				Elem:        resourceAttributeFulfillmentValue(),
			},
			"filter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The JDBC WHERE clause used to query your data store to locate a user record.",
			},
		},
	}
}

func resourceCustomAttributeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data_store_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Reference to the associated data store.",
				Elem:        resourceLinkResource(),
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID that defines this attribute source. Only alphanumeric characters allowed.\nNote: Required for OpenID Connect policy attribute sources, OAuth IdP adapter mappings, OAuth access token mappings and APC-to-SP Adapter Mappings. IdP Connections will ignore this property since it only allows one attribute source to be defined per mapping. IdP-to-SP Adapter Mappings can contain multiple attribute sources.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of this attribute source. The description needs to be unique amongst the attribute sources for the mapping.\nNote: Required for APC-to-SP Adapter Mappings",
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of mappings from attribute names to their fulfillment values. This field is only valid for the SP Connection's Browser SSO mappings",
				Elem:        resourceAttributeFulfillmentValue(),
			},
			"filter_fields": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of fields that can be used to filter a request to the custom data store.",
				Elem:        resourceFieldEntry(),
			},
		},
	}
}

func resourceFieldEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this field.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value of this field. Whether or not the value is required will be determined by plugin validation checks.",
			},
		},
	}
}

func resourceAttributeFulfillmentValue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Contract key value for the source.",
			},
			"source": resourceSourceTypeIdKey(),
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The value for this attribute.",
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
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Description: "The attribute value source.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"type": {
					Type:        schema.TypeString,
					Required:    true,
					Description: "A key that is meant to reference a source from which an attribute can be retrieved. This model is usually paired with a value which, depending on the SourceType, can be a hardcoded value or a reference to an attribute name specific to that SourceType. Not all values are applicable - a validation error will be returned for incorrect values.",
					ValidateFunc: validation.StringInSlice([]string{
						"TOKEN_EXCHANGE_PROCESSOR_POLICY",
						"ACCOUNT_LINK",
						"ADAPTER",
						"ASSERTION",
						"CONTEXT",
						"CUSTOM_DATA_STORE",
						"EXPRESSION",
						"JDBC_DATA_STORE",
						"LDAP_DATA_STORE",
						"MAPPED_ATTRIBUTES",
						"NO_MAPPING",
						"TEXT",
						"TOKEN",
						"REQUEST",
						"OAUTH_PERSISTENT_GRANT",
						"SUBJECT_TOKEN",
						"ACTOR_TOKEN",
						"PASSWORD_CREDENTIAL_VALIDATOR",
						"IDP_CONNECTION",
						"AUTHENTICATION_POLICY_CONTRACT",
						"CLAIMS",
						"LOCAL_IDENTITY_PROFILE",
						"EXTENDED_CLIENT_METADATA",
						"EXTENDED_PROPERTIES",
						"TRACKED_HTTP_PARAMS",
						"FRAGMENT",
						"INPUTS",
					}, false),
				},
				"id": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The attribute source ID that refers to the attribute source that this key references. In some resources, the ID is optional and will be ignored. In these cases the ID should be omitted. If the source type is not an attribute source then the ID can be omitted.",
				},
			},
		},
	}
}

func resourceIssuanceCriteria() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"conditional_criteria": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of conditional issuance criteria where existing attributes must satisfy their conditions against expected values in order for the transaction to continue.",
				Elem:        resourceConditionalIssuanceCriteriaEntry(),
			},
			"expression_criteria": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of expression issuance criteria where the OGNL expressions must evaluate to true in order for the transaction to continue.",
				Elem:        resourceExpressionIssuanceCriteriaEntry(),
			},
		},
	}
}

func resourceConditionalIssuanceCriteriaEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"source": resourceSourceTypeIdKey(),
			"attribute_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the attribute to use in this issuance criterion.",
			},
			"condition": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The condition that will be applied to the source attribute's value and the expected value.",
				ValidateFunc: validation.StringInSlice([]string{
					"EQUALS", "EQUALS_CASE_INSENSITIVE", "EQUALS_DN", "NOT_EQUAL", "NOT_EQUAL_CASE_INSENSITIVE", "NOT_EQUAL_DN", "MULTIVALUE_CONTAINS", "MULTIVALUE_CONTAINS_CASE_INSENSITIVE", "MULTIVALUE_CONTAINS_DN", "MULTIVALUE_DOES_NOT_CONTAIN", "MULTIVALUE_DOES_NOT_CONTAIN_CASE_INSENSITIVE", "MULTIVALUE_DOES_NOT_CONTAIN_DN",
				}, false),
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The expected value of this issuance criterion.",
			},
			"error_result": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.",
			},
		},
	}
}

func resourceExpressionIssuanceCriteriaEntry() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The OGNL expression to evaluate.",
			},
			"error_result": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The error result to return if this issuance criterion fails. This error result will show up in the PingFederate server logs.",
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
				Type:        schema.TypeSet,
				Optional:    true,
				MinItems:    1,
				Description: "A list of additional attributes that can be returned by the Authentication Selector. The extended attributes are only used if the Authentication Selector supports them.",
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
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of read-only attributes (for example, sub) that are automatically populated by PingFederate.",
				Elem:        resourceOpenIdConnectAttribute(),
			},
			"extended_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of additional attributes.",
				Elem:        resourceOpenIdConnectAttribute(),
			},
		},
	}
}

func resourceAttributeMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ldap_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of ldap configured data stores to look up attributes from.",
				Elem:        resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of jdbc configured data stores to look up attributes from.",
				Elem:        resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of custom configured data stores to look up attributes from.",
				Elem:        resourceCustomAttributeSource(),
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of mappings from attribute names to their fulfillment values.",
				Elem:        resourceAttributeFulfillmentValue(),
			},
			"issuance_criteria": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
				Elem:        resourceIssuanceCriteria(),
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The extended property name.",
			},
			"values": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A List of parameter values.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
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
				Type:        schema.TypeList,
				Elem:        resourceAttributeQueryNameMapping(),
				Description: "The attribute name mappings between the SP and the IdP.",
				Optional:    true,
			},
			"policy": {
				Type:        schema.TypeList,
				Elem:        resourceIdpAttributeQueryPolicy(),
				Description: "The attribute query profile's security policy.",
				Optional:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "The URL at your IdP partner's site where attribute queries are to be sent.",
				Required:    true,
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
				Computed: true,
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
				Optional: true,
				Computed: true,
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
			"request_contract_ref": resourceLinkSchema(),
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
				Type:        schema.TypeList,
				Elem:        resourceConnectionCert(),
				Description: "The certificate used for signature verification and XML encryption.",
				Optional:    true,
			},
			"digital_signature": {
				Type:        schema.TypeBool,
				Description: "If incoming or outgoing messages must be signed.",
				Optional:    true,
			},
			"http_basic_credentials": {
				Type:        schema.TypeList,
				Elem:        resourceUsernamePasswordCredentials(),
				Description: "The credentials to use when you authenticate with the SOAP endpoint.",
				Optional:    true,
			},
			"require_ssl": {
				Type:        schema.TypeBool,
				Description: "Incoming HTTP transmissions must use a secure channel.",
				Optional:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "",
				Optional:    true,
			},
			"verification_issuer_dn": {
				Type:        schema.TypeString,
				Description: "If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field.",
				Optional:    true,
			},
			"verification_subject_dn": {
				Type:        schema.TypeString,
				Description: "If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array.",
				Optional:    true,
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
				Type:        schema.TypeString,
				Description: "Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.",
				Optional:    true,
			},
			"file_data": {
				Type:        schema.TypeString,
				Description: "The certificate data in PEM format. New line characters should be omitted or encoded in this value.",
				Required:    true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					eq := strings.ReplaceAll(old, "\n", "") == strings.ReplaceAll(new, "\n", "")
					return eq
				},
			},
			"id": {
				Type:        schema.TypeString,
				Description: "The persistent, unique ID for the certificate. It can be any combination of [a-z0-9._-]. This property is system-assigned if not specified.",
				Optional:    true,
				Computed:    true,
			},
		},
	}
}

//AdditionalAllowedEntitiesConfiguration - Additional allowed entities or issuers configuration. Currently only used in OIDC IdP (RP) connection.
func resourceAdditionalAllowedEntitiesConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_allowed_entities": {
				Type:        schema.TypeList,
				Elem:        resourceEntity(),
				Optional:    true,
				Description: "Set to true to configure additional entities or issuers to be accepted during entity or issuer validation.",
			},
			"allow_additional_entities": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "An array of additional allowed entities or issuers to be accepted during entity or issuer validation.",
			},
			"allow_all_entities": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set to true to accept any entity or issuer during entity or issuer validation. (Not Recommended)",
			},
		},
	}
}

//Entity
func resourceEntity() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Entity description.",
			},
			"entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique entity identifier.",
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
				Type:        schema.TypeBool,
				Description: "Indicates whether this is an active signature verification certificate.",
				Optional:    true,
			},
			"cert_view": {
				Type:        schema.TypeList,
				Elem:        resourceCertView(),
				Description: "Certificate details. This property is read-only and is always ignored on a POST or PUT.",
				Optional:    true,
				Computed:    true,
			},
			"encryption_cert": {
				Type:        schema.TypeBool,
				Description: "Indicates whether to use this cert to encrypt outgoing assertions. Only one certificate in the collection can have this flag set.",
				Optional:    true,
			},
			"primary_verification_cert": {
				Type:        schema.TypeBool,
				Description: "Indicates whether this is the primary signature verification certificate. Only one certificate in the collection can have this flag set.",
				Optional:    true,
			},
			"secondary_verification_cert": {
				Type:        schema.TypeBool,
				Description: "Indicates whether this is the secondary signature verification certificate. Only one certificate in the collection can have this flag set.",
				Optional:    true,
			},
			"x509_file": {
				Type:        schema.TypeList,
				Elem:        resourceX509File(),
				Description: "The certificate data. This property must always be supplied on a POST or PUT.",
				Required:    true,
			},
		},
	}
}

//IdpBrowserSso - The settings used to enable secure browser-based SSO to resources at your site.
func resourceIdpBrowserSso() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_mappings": {
				Type:        schema.TypeList,
				Elem:        resourceSpAdapterMapping(),
				Description: "A list of adapters that map to incoming assertions.",
				Optional:    true,
			},
			"artifact": {
				Type:        schema.TypeList,
				Elem:        resourceArtifactSettings(),
				Description: "The settings for an artifact binding.",
				Optional:    true,
			},
			"assertions_signed": {
				Type:        schema.TypeBool,
				Description: "Specify whether the incoming SAML assertions are signed rather than the entire SAML response being signed.",
				Optional:    true,
			},
			"attribute_contract": {
				Type:        schema.TypeList,
				Elem:        resourceIdpBrowserSsoAttributeContract(),
				Description: "The list of attributes that the IdP sends in the assertion.",
				Optional:    true,
			},
			"authentication_policy_contract_mappings": {
				Type:        schema.TypeList,
				Elem:        resourceAuthenticationPolicyContractMapping(),
				Description: "A list of Authentication Policy Contracts that map to incoming assertions.",
				Optional:    true,
			},
			"authn_context_mappings": {
				Type:        schema.TypeList,
				Elem:        resourceAuthnContextMapping(),
				Description: "A list of authentication context mappings between local and remote values. Applicable for SAML 2.0 and OIDC protocol connections.",
				Optional:    true,
			},
			"decryption_policy": {
				Type:        schema.TypeList,
				Elem:        resourceDecryptionPolicy(),
				Description: "The SAML 2.0 decryption policy for browser-based SSO.",
				Optional:    true,
			},
			"default_target_url": {
				Type:        schema.TypeString,
				Description: "The default target URL for this connection. If defined, this overrides the default URL.",
				Optional:    true,
			},
			"enabled_profiles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.",
				Optional:    true,
			},
			"idp_identity_mapping": {
				Type:        schema.TypeString,
				Description: "Defines the process in which users authenticated by the IdP are associated with user accounts local to the SP.",
				Required:    true,
			},
			"incoming_bindings": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections.",
				Optional:    true,
			},
			"message_customizations": {
				Type:        schema.TypeList,
				Elem:        resourceProtocolMessageCustomization(),
				Description: "The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported.",
				Optional:    true,
			},
			"oauth_authentication_policy_contract_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The Authentication policy contract to map into for OAuth. The policy contract can subsequently be mapped into the OAuth persistent grant.",
				Elem:        resourceLinkResource(),
			},
			"oidc_provider_settings": {
				Type:        schema.TypeList,
				Elem:        resourceOIDCProviderSettings(),
				Description: "The OpenID Provider configuration settings. Required for an OIDC connection.",
				Optional:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "The browser-based SSO protocol to use.",
				Required:    true,
			},
			"sign_authn_requests": {
				Type:        schema.TypeBool,
				Description: "Determines whether SAML authentication requests should be signed.",
				Optional:    true,
			},
			"slo_service_endpoints": {
				Type:        schema.TypeList,
				Elem:        resourceSloServiceEndpoint(),
				Description: "A list of possible endpoints to send SLO requests and responses.",
				Optional:    true,
			},
			"sso_o_auth_mapping": {
				Type:        schema.TypeList,
				Elem:        resourceSsoOAuthMapping(),
				Description: "Direct mapping from the IdP connection to the OAuth persistent grant.",
				Optional:    true,
			},
			"sso_service_endpoints": {
				Type:        schema.TypeList,
				Elem:        resourceIdpSsoServiceEndpoint(),
				Description: "The IdP SSO endpoints that define where to send your authentication requests. Only required for SP initiated SSO. This is required for SAML x.x and WS-FED Connections.",
				Optional:    true,
			},
			"url_whitelist_entries": {
				Type:        schema.TypeList,
				Elem:        resourceUrlWhitelistEntry(),
				Description: "For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled.",
				Optional:    true,
			},
		},
	}
}

//ContactInfo - Contact information.
func resourceContactInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"company": {
				Type:        schema.TypeString,
				Description: "Company name.",
				Optional:    true,
			},
			"email": {
				Type:        schema.TypeString,
				Description: "Contact email address.",
				Optional:    true,
			},
			"first_name": {
				Type:        schema.TypeString,
				Description: "Contact first name.",
				Optional:    true,
			},
			"last_name": {
				Type:        schema.TypeString,
				Description: "Contact last name.",
				Optional:    true,
			},
			"phone": {
				Type:        schema.TypeString,
				Description: "Contact phone number.",
				Optional:    true,
			},
		},
	}
}

//SpAdapterMapping - A mapping to a SP adapter.
func resourceSpAdapterMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_override_settings": {
				Type:        schema.TypeList,
				Elem:        resourceSpAdapter(),
				Description: "Connection specific overridden adapter instance for mapping.",
				Optional:    true,
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        resourceAttributeFulfillmentValue(),
				Description: "A list of mappings from attribute names to their fulfillment values.",
			},
			"ldap_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of ldap configured data stores to look up attributes from.",
				Elem:        resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of jdbc configured data stores to look up attributes from.",
				Elem:        resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of custom configured data stores to look up attributes from.",
				Elem:        resourceCustomAttributeSource(),
			},
			"issuance_criteria": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        resourceIssuanceCriteria(),
				Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
				Optional:    true,
			},
			"restrict_virtual_entity_ids": {
				Type:        schema.TypeBool,
				Description: "Restricts this mapping to specific virtual entity IDs.",
				Optional:    true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of virtual server IDs that this mapping is restricted to.",
				Optional:    true,
			},
			"sp_adapter_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Reference to the associated SP adapter.",
				Elem:        resourceLinkResource(),
			},
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
				Type:        schema.TypeList,
				Elem:        resourceIdpAdapterAttributeContract(),
				Description: "The list of attributes that the IdP adapter provides.",
				Optional:    true,
			},
			"attribute_mapping": {
				Type:        schema.TypeList,
				Elem:        resourceIdpAdapterContractMapping(),
				Description: "The attributes mapping from attribute sources to attribute targets.",
				Optional:    true,
			},
			"authn_ctx_class_ref": {
				Type:        schema.TypeString,
				Description: "The fixed value that indicates how the user was authenticated.",
				Optional:    true,
			},
			"configuration": resourcePluginConfiguration(),
			"id": {
				Type:        schema.TypeString,
				Description: "The ID of the plugin instance. The ID cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The plugin instance name. The name cannot be modified once the instance is created.<br>Note: Ignored when specifying a connection's adapter override.",
				Required:    true,
			},
			//"parent_ref": {
			//	Type:        schema.TypeList,
			//	Optional:    true,
			//	MaxItems:    1,
			//	Description: "The reference to this plugin's parent instance. The parent reference is only accepted if the plugin type supports parent instances. Note: This parent reference is required if this plugin instance is used as an overriding plugin (e.g. connection adapter overrides)",
			//	Elem:        resourceLinkResource(),
			//},
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
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Encrypt the assertion.",
			},
			"require_encrypted_name_id": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Require an encrypted name identifier.",
			},
			"require_signed_attribute_query": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Require signed attribute query.",
			},
			"sign_assertion": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sign the assertion.",
			},
			"sign_response": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Sign the response.",
			},
		},
	}
}

//SpBrowserSso - The SAML settings used to enable secure browser-based SSO to resources at your partner's site.
func resourceSpBrowserSso() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"adapter_mappings": {
				Type:        schema.TypeList,
				Elem:        resourceIdpAdapterAssertionMapping(),
				Description: "A list of adapters that map to outgoing assertions.",
				Optional:    true,
			},
			"artifact": {
				Type:        schema.TypeList,
				Elem:        resourceArtifactSettings(),
				Description: "The settings for an artifact binding.",
				Optional:    true,
			},
			"assertion_lifetime": {
				Type:        schema.TypeList,
				Elem:        resourceAssertionLifetime(),
				Description: "The timeframe of validity before and after the issuance of the assertion.",
				Optional:    true,
			},
			"attribute_contract": {
				Type:        schema.TypeList,
				Elem:        resourceSpBrowserSsoAttributeContract(),
				Description: "A set of user attributes that the IdP sends in the SAML assertion.",
				Optional:    true,
			},
			"authentication_policy_contract_assertion_mappings": {
				Type:        schema.TypeList,
				Elem:        resourceAuthenticationPolicyContractAssertionMapping(),
				Description: "A list of authentication policy contracts that map to outgoing assertions.",
				Optional:    true,
			},
			"default_target_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Default Target URL for SAML1.x connections. For SP connections, this default URL represents the destination on the SP where the user will be directed. For IdP connections, entering a URL in the Default Target URL field overrides the SP Default URL SSO setting.",
			},
			"enabled_profiles": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The profiles that are enabled for browser-based SSO. SAML 2.0 supports all profiles whereas SAML 1.x IdP connections support both IdP and SP (non-standard) initiated SSO. This is required for SAMLx.x Connections.",
				Optional:    true,
			},
			"encryption_policy": {
				Type:        schema.TypeList,
				Elem:        resourceEncryptionPolicy(),
				Description: "The SAML 2.0 encryption policy for browser-based SSO. Required for SAML 2.0 connections.",
				Optional:    true,
			},
			"incoming_bindings": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The SAML bindings that are enabled for browser-based SSO. This is required for SAML 2.0 connections when the enabled profiles contain the SP-initiated SSO profile or either SLO profile. For SAML 1.x based connections, it is not used for SP Connections and it is optional for IdP Connections.",
				Optional:    true,
			},
			"message_customizations": {
				Type:        schema.TypeList,
				Elem:        resourceProtocolMessageCustomization(),
				Description: "The message customizations for browser-based SSO. Depending on server settings, connection type, and protocol this may or may not be supported.",
				Optional:    true,
			},
			"protocol": {
				Type:        schema.TypeString,
				Description: "The browser-based SSO protocol to use.",
				Required:    true,
			},
			"require_signed_authn_requests": {
				Type:        schema.TypeBool,
				Description: "Require AuthN requests to be signed when received via the POST or Redirect bindings.",
				Optional:    true,
			},
			"sign_assertions": {
				Type:        schema.TypeBool,
				Description: "Always sign the SAML Assertion.",
				Optional:    true,
			},
			"sign_response_as_required": {
				Type:        schema.TypeBool,
				Description: "Sign SAML Response as required by the associated binding and encryption policy. Applicable to SAML2.0 only and is defaulted to true. It can be set to false only on SAML2.0 connections when signAssertions is set to true.",
				Optional:    true,
			},
			"slo_service_endpoints": {
				Type:        schema.TypeList,
				Elem:        resourceSloServiceEndpoint(),
				Description: "A list of possible endpoints to send SLO requests and responses.",
				Optional:    true,
			},
			"sp_saml_identity_mapping": {
				Type:        schema.TypeString,
				Description: "Process in which users authenticated by the IdP are associated with user accounts local to the SP.",
				Optional:    true,
			},
			"sp_ws_fed_identity_mapping": {
				Type:        schema.TypeString,
				Description: "Process in which users authenticated by the IdP are associated with user accounts local to the SP for WS-Federation connection types.",
				Optional:    true,
			},
			"sso_service_endpoints": {
				Type:        schema.TypeList,
				Elem:        resourceSpSsoServiceEndpoint(),
				Description: "A list of possible endpoints to send assertions to.",
				Required:    true,
			},
			"url_whitelist_entries": {
				Type:        schema.TypeList,
				Elem:        resourceUrlWhitelistEntry(),
				Description: "For WS-Federation connections, a whitelist of additional allowed domains and paths used to validate wreply for SLO, if enabled.",
				Optional:    true,
			},
			"ws_fed_token_type": {
				Type:        schema.TypeString,
				Description: "The WS-Federation Token Type to use.",
				Optional:    true,
			},
			"ws_trust_version": {
				Type:        schema.TypeString,
				Description: "The WS-Trust version for a WS-Federation connection. The default version is WSTRUST12.",
				Optional:    true,
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
				Type:        schema.TypeString,
				Description: "For GET requests, this field contains the encrypted password, if one exists.  For POST and PUT requests, if you wish to reuse the existing password, this field should be passed back unchanged.",
				Optional:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "User password.  To update the password, specify the plaintext value in this field.  This field will not be populated for GET requests.",
				Optional:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "The username.",
				Optional:    true,
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
				Optional: true,
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
				Type:        schema.TypeString,
				Description: "The local attribute name.",
				Required:    true,
			},
			"remote_name": {
				Type:        schema.TypeString,
				Description: "The remote attribute name as defined by the attribute authority.",
				Required:    true,
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
				MaxItems: 1,
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
				MaxItems: 1,
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
				Optional: true,
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
				MaxItems: 1,
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
				MaxItems: 1,
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
			"ssl_auth_key_pair_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The ID of the key pair used to authenticate with your partner's SOAP endpoint. The ID of the key pair is also known as the alias and can be found by viewing the corresponding certificate under 'SSL Server Certificates' in the PingFederate Administrative Console.",
				Elem:        resourceLinkResource(),
			},
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
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of ldap configured data stores to look up attributes from.",
				Elem:        resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of jdbc configured data stores to look up attributes from.",
				Elem:        resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of custom configured data stores to look up attributes from.",
				Elem:        resourceCustomAttributeSource(),
			},
			"attributes": {
				Type:        schema.TypeList,
				Description: "The list of attributes that may be returned to the SP in the response to an attribute request.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"issuance_criteria": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        resourceIssuanceCriteria(),
				Optional:    true,
				Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
			},
			"policy": {
				Type:        schema.TypeList,
				Elem:        resourceSpAttributeQueryPolicy(),
				Optional:    true,
				Description: "The attribute query profile's security policy.",
			},
		},
	}
}

//CertView - Certificate details.
func resourceCertView() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"crypto_provider": {
				Type:        schema.TypeString,
				Description: "Cryptographic Provider. This is only applicable if Hybrid HSM mode is true.",
				Optional:    true,
			},
			"expires": {
				Type:        schema.TypeString,
				Description: "The end date up until which the item is valid, in ISO 8601 format (UTC).",
				Optional:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Description: "The persistent, unique ID for the certificate.",
				Optional:    true,
			},
			"issuer_dn": {
				Type:        schema.TypeString,
				Description: "The issuer's distinguished name.",
				Optional:    true,
			},
			"key_algorithm": {
				Type:        schema.TypeString,
				Description: "The public key algorithm.",
				Optional:    true,
			},
			"key_size": {
				Type:        schema.TypeInt,
				Description: "The public key size.",
				Optional:    true,
			},
			"serial_number": {
				Type:        schema.TypeString,
				Description: "The serial number assigned by the CA.",
				Optional:    true,
			},
			"sha1_fingerprint": {
				Type:        schema.TypeString,
				Description: "SHA-1 fingerprint in Hex encoding.",
				Optional:    true,
			},
			"sha256_fingerprint": {
				Type:        schema.TypeString,
				Description: "SHA-256 fingerprint in Hex encoding.",
				Optional:    true,
			},
			"signature_algorithm": {
				Type:        schema.TypeString,
				Description: "The signature algorithm.",
				Optional:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Status of the item.",
				Optional:    true,
			},
			"subject_alternative_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The subject alternative names (SAN).",
				Optional:    true,
			},
			"subject_dn": {
				Type:        schema.TypeString,
				Description: "The subject's distinguished name.",
				Optional:    true,
			},
			"valid_from": {
				Type:        schema.TypeString,
				Description: "The start date from which the item is valid, in ISO 8601 format (UTC).",
				Optional:    true,
			},
			"version": {
				Type:        schema.TypeInt,
				Description: "The X.509 version to which the item conforms.",
				Optional:    true,
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
				MaxItems: 1,
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
				Type:        schema.TypeBool,
				Description: "If set to true, SSO transaction will be aborted as a fail-safe when the data-store's attribute mappings fail to complete the attribute contract. Otherwise, the attribute contract with default values is used. By default, this value is false.",
				Optional:    true,
			},
			"adapter_override_settings": {
				Type:        schema.TypeList,
				Elem:        resourceIdpAdapter(),
				Description: "Connection specific configuration overrides for the mapped adapter instance.",
				Optional:    true,
			},
			"attribute_contract_fulfillment": {
				Type:        schema.TypeSet,
				Elem:        resourceAttributeFulfillmentValue(),
				Description: "A list of mappings from attribute names to their fulfillment values.",
				Required:    true,
			},
			"ldap_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured ldap data stores to look up attributes from.",
				Elem:        resourceLdapAttributeSource(),
			},
			"jdbc_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured jdbc data stores to look up attributes from.",
				Elem:        resourceJdbcAttributeSource(),
			},
			"custom_attribute_source": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of configured custom data stores to look up attributes from.",
				Elem:        resourceCustomAttributeSource(),
			},
			"idp_adapter_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Reference to the associated IdP adapter.\nNote: This is ignored if adapter overrides for this mapping exists. In this case, the override's parent adapter reference is used.",
				Elem:        resourceLinkResource(),
			},
			"issuance_criteria": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        resourceIssuanceCriteria(),
				Description: "The issuance criteria that this transaction must meet before the corresponding attribute contract is fulfilled.",
				Optional:    true,
			},
			"restrict_virtual_entity_ids": {
				Type:        schema.TypeBool,
				Description: "Restricts this mapping to specific virtual entity IDs.",
				Optional:    true,
			},
			"restricted_virtual_entity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of virtual server IDs that this mapping is restricted to.",
				Optional:    true,
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
				MaxItems: 1,
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
				MaxItems: 1,
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
				Type:        schema.TypeString,
				Description: "The name of this attribute.",
				Required:    true,
			},
			"name_format": {
				Type:        schema.TypeString,
				Description: "The SAML Name Format for the attribute.",
				Required:    true,
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
				Type:        schema.TypeBool,
				Description: "Encrypt the name identifier.",
				Optional:    true,
			},
			"mask_attribute_values": {
				Type:        schema.TypeBool,
				Description: "Mask attributes in log files.",
				Optional:    true,
			},
			"require_encrypted_assertion": {
				Type:        schema.TypeBool,
				Description: "Require encrypted assertion.",
				Optional:    true,
			},
			"require_signed_assertion": {
				Type:        schema.TypeBool,
				Description: "Require signed assertion.",
				Optional:    true,
			},
			"require_signed_response": {
				Type:        schema.TypeBool,
				Description: "Require signed response.",
				Optional:    true,
			},
			"sign_attribute_query": {
				Type:        schema.TypeBool,
				Description: "Sign the attribute query.",
				Optional:    true,
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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the plugin instance. The ID cannot be modified once the instance is created.\nNote: Ignored when specifying a connection's adapter override.",
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
				Type:        schema.TypeString,
				Description: "The algorithm used to sign messages sent to this partner. The default is SHA1withDSA for DSA certs, SHA256withRSA for RSA certs, and SHA256withECDSA for EC certs. For RSA certs, SHA1withRSA, SHA384withRSA, and SHA512withRSA are also supported. For EC certs, SHA384withECDSA and SHA512withECDSA are also supported. If the connection is WS-Federation with JWT token type, then the possible values are RSA SHA256, RSA SHA384, RSA SHA512, ECDSA SHA256, ECDSA SHA384, ECDSA SHA512",
				Optional:    true,
			},
			"include_cert_in_signature": {
				Type:        schema.TypeBool,
				Description: "Determines whether the signing certificate is included in the signature <KeyInfo> element.",
				Optional:    true,
			},
			"include_raw_key_in_signature": {
				Type:        schema.TypeBool,
				Description: "Determines whether the <KeyValue> element with the raw public key is included in the signature <KeyInfo> element.",
				Optional:    true,
			},
			"signing_key_pair_ref": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Settings related to the manner in which messages sent to the partner are digitally signed. Required for SP Connections.",
				Elem:        resourceLinkResource(),
			},
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
				Type:     schema.TypeSet,
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
				Type:        schema.TypeString,
				Description: "The algorithm used to encrypt assertions sent to this partner. AES_128, AES_256, AES_128_GCM, AES_192_GCM, AES_256_GCM and Triple_DES are also supported. Default is AES_128",
				Optional:    true,
			},
			"certs": {
				Type:        schema.TypeList,
				Elem:        resourceConnectionCert(),
				Description: "The certificates used for signature verification and XML encryption.",
				Optional:    true,
			},
			"decryption_key_pair_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The ID of the primary key pair used to decrypt message content received from this partner. The ID of the key pair is also known as the alias and can be found by viewing the corresponding certificate under 'Signing & Decryption Keys & Certificates' in the PingFederate Administrative Console.",
				Elem:        resourceLinkResource(),
			},
			"inbound_back_channel_auth": {
				Type:        schema.TypeList,
				Elem:        resourceInboundBackChannelAuth(),
				Description: "The SOAP authentication method(s) to use when you receive a message using SOAP back channel.",
				Optional:    true,
			},
			"key_transport_algorithm": {
				Type:        schema.TypeString,
				Description: "The algorithm used to transport keys to this partner. RSA_OAEP and RSA_v15 are supported. Default is RSA_OAEP",
				Optional:    true,
			},
			"outbound_back_channel_auth": {
				Type:        schema.TypeList,
				Elem:        resourceOutboundBackChannelAuth(),
				Description: "The SOAP authentication method(s) to use when you send a message using SOAP back channel.",
				Optional:    true,
			},
			"secondary_decryption_key_pair_ref": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The ID of the secondary key pair used to decrypt message content received from this partner.",
				Elem:        resourceLinkResource(),
			},
			"signing_settings": {
				Type:        schema.TypeList,
				Elem:        resourceSigningSettings(),
				Description: "Settings related to the manner in which messages sent to the partner are digitally signed. Required for SP Connections.",
				Optional:    true,
			},
			"verification_issuer_dn": {
				Type:        schema.TypeString,
				Description: "If a verification Subject DN is provided, you can optionally restrict the issuer to a specific trusted CA by specifying its DN in this field.",
				Optional:    true,
			},
			"verification_subject_dn": {
				Type:        schema.TypeString,
				Description: "If this property is set, the verification trust model is Anchored. The verification certificate must be signed by a trusted CA and included in the incoming message, and the subject DN of the expected certificate is specified in this property. If this property is not set, then a primary verification certificate must be specified in the certs array.",
				Optional:    true,
			},
		},
	}
}

//SpSsoServiceEndpoint - The settings that define a service endpoint to a SP SSO service.
func resourceSpSsoServiceEndpoint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binding": {
				Type:        schema.TypeString,
				Description: "The binding of this endpoint, if applicable - usually only required for SAML 2.0 endpoints.  Supported bindings are Artifact and POST.",
				Required:    true,
			},
			"index": {
				Type:        schema.TypeInt,
				Description: "The priority of the endpoint.",
				Required:    true,
			},
			"is_default": {
				Type:        schema.TypeBool,
				Description: "Whether or not this endpoint is the default endpoint. Defaults to false.",
				Optional:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "The absolute or relative URL of the endpoint. A relative URL can be specified if a base URL for the connection has been defined.",
				Required:    true,
			},
		},
	}
}

//CrlSettings - CRL settings.
func resourceCrlSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"next_retry_mins_when_next_update_in_past": {
				Type:        schema.TypeInt,
				Description: "Next retry on next update expiration in minutes. This value defaults to `60`.",
				Optional:    true,
				Default:     60,
			},
			"next_retry_mins_when_resolve_failed": {
				Type:        schema.TypeInt,
				Description: "Next retry on resolution failure in minutes. This value defaults to `1440`.",
				Optional:    true,
				Default:     1440,
			},
			"treat_non_retrievable_crl_as_revoked": {
				Type:        schema.TypeBool,
				Description: "Treat non retrievable CRL as revoked. This setting defaults to disabled.",
				Optional:    true,
				Default:     false,
			},
			"verify_crl_signature": {
				Type:        schema.TypeBool,
				Description: "Verify CRL signature. This setting defaults to enabled.",
				Optional:    true,
				Default:     true,
			},
		},
	}
}

//OcspSettings - OCSP settings.
func resourceOcspSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_on_responder_unavailable": {
				Type:        schema.TypeString,
				Description: "Action on responder unavailable. This value defaults to `CONTINUE`.",
				Optional:    true,
				Default:     "CONTINUE",
			},
			"action_on_status_unknown": {
				Type:        schema.TypeString,
				Description: "Action on status unknown. This value defaults to `FAIL`.",
				Optional:    true,
				Default:     "FAIL",
			},
			"action_on_unsuccessful_response": {
				Type:        schema.TypeString,
				Description: "Action on unsuccessful response. This value defaults to `FAIL`.",
				Optional:    true,
				Default:     "FAIL",
			},
			"current_update_grace_period": {
				Type:        schema.TypeInt,
				Description: "Current update grace period in minutes. This value defaults to `5`.",
				Optional:    true,
				Default:     5,
			},
			"next_update_grace_period": {
				Type:        schema.TypeInt,
				Description: "Next update grace period in minutes. This value defaults to `5`.",
				Optional:    true,
				Default:     5,
			},
			"requester_add_nonce": {
				Type:        schema.TypeBool,
				Description: "Do not allow responder to use cached responses. This setting defaults to disabled.",
				Optional:    true,
				Default:     false,
			},
			"responder_cert_reference": {
				Type:        schema.TypeList,
				Description: "Resource link to OCSP responder signature verification certificate. A previously selected certificate will be deselected if this attribute is not defined.",
				Optional:    true,
				MaxItems:    1,
				Elem:        resourceLinkResource(),
			},
			"responder_timeout": {
				Type:        schema.TypeInt,
				Description: "Responder connection timeout in seconds. This value defaults to `5`.",
				Optional:    true,
				Default:     5,
			},
			"responder_url": {
				Type:        schema.TypeString,
				Description: "Responder URL address. This field is required if OCSP revocation is enabled.",
				Required:    true,
			},
			"response_cache_period": {
				Type:        schema.TypeInt,
				Description: "Response cache period in hours. This value defaults to `48`.",
				Optional:    true,
				Default:     48,
			},
		},
	}
}

//ProxySettings - Proxy settings.
func resourceProxySettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Description: "Host name.",
				Optional:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Port number.",
				Optional:    true,
			},
		},
	}
}

//ClientMetadata - The client metadata.
func resourceClientMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Description: "The metadata description.",
				Optional:    true,
			},
			"multi_valued": {
				Type:        schema.TypeBool,
				Description: "If the field should allow multiple values.",
				Optional:    true,
			},
			"parameter": {
				Type:        schema.TypeString,
				Description: "The metadata name.",
				Required:    true,
			},
		},
	}
}

//DynamicClientRegistration - Dynamic client registration settings.
func resourceDynamicClientRegistration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_client_delete": {
				Type:        schema.TypeBool,
				Description: "Allow client deletion from dynamic client management.",
				Optional:    true,
			},
			"allowed_exclusive_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The exclusive scopes to allow.",
				Optional:    true,
			},
			"bypass_activation_code_confirmation_override": {
				Type:        schema.TypeBool,
				Description: "Indicates if the Activation Code Confirmation page should be bypassed if 'verification_url_complete' is used by the end user to authorize a device.",
				Optional:    true,
			},
			"ciba_polling_interval": {
				Type:        schema.TypeInt,
				Description: "The minimum amount of time in seconds that the Client must wait between polling requests to the token endpoint. The default is 3 seconds.",
				Optional:    true,
				Default:     3,
			},
			"ciba_require_signed_requests": {
				Type:        schema.TypeBool,
				Description: "Determines whether CIBA signed requests are required for this client.",
				Optional:    true,
				Default:     false,
			},
			"client_cert_issuer_ref": resourceLinkSchema(),
			"client_cert_issuer_type": {
				Type:        schema.TypeString,
				Description: "Client TLS Certificate Issuer Type.",
				Optional:    true,
				Default:     "NONE",
			},
			"default_access_token_manager_ref": resourceLinkSchema(),
			"device_flow_setting_type": {
				Type:        schema.TypeString,
				Description: "Allows an administrator to override the Device Authorization Settings set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
				Optional:    true,
				Default:     "SERVER_DEFAULT",
			},
			"device_polling_interval_override": {
				Type:        schema.TypeInt,
				Description: "The amount of time client should wait between polling requests, in seconds.",
				Optional:    true,
			},
			"disable_registration_access_tokens": {
				Type:        schema.TypeBool,
				Description: "Disable registration access tokens. Local standards may mandate different registration access token requirements. If applicable, implement custom validation and enforcement rules using the DynamicClientRegistrationPlugin interface from the PingFederate SDK, configure the client registration policies (policyRefs), and set this property (disableRegistrationAccessTokens) to true. CAUTION: When the disableRegistrationAccessTokens property is set to true, all clients, not just the ones created using the Dynamic Client Registration protocol, are vulnerable to unrestricted retrievals, updates (including modifications to the client authentication scheme and redirect URIs), and deletes at the /as/clients.oauth2 endpoint unless one or more client registration policies are in place to protect against unauthorized attempts.",
				Optional:    true,
			},
			"enforce_replay_prevention": {
				Type:        schema.TypeBool,
				Description: "Enforce replay prevention.",
				Optional:    true,
				Default:     false,
			},
			"initial_access_token_scope": {
				Type:        schema.TypeString,
				Description: "The initial access token to prevent unwanted client registrations.",
				Optional:    true,
			},
			"oidc_policy": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Elem:        resourceClientRegistrationOIDCPolicy(),
				Description: "Open ID Connect Policy settings.  This is included in the message only when OIDC is enabled.",
				Optional:    true,
			},
			"pending_authorization_timeout_override": {
				Type:        schema.TypeInt,
				Description: "The 'device_code' and 'user_code' timeout, in seconds.",
				Optional:    true,
			},
			"persistent_grant_expiration_time": {
				Type:        schema.TypeInt,
				Description: "The persistent grant expiration time.",
				Optional:    true,
			},
			"persistent_grant_expiration_time_unit": {
				Type:        schema.TypeString,
				Description: "The persistent grant expiration time unit.",
				Optional:    true,
			},
			"persistent_grant_expiration_type": {
				Type:        schema.TypeString,
				Description: "Allows an administrator to override the Persistent Grant Lifetime set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
				Optional:    true,
				Default:     "SERVER_DEFAULT",
			},
			"persistent_grant_idle_timeout": {
				Type:        schema.TypeInt,
				Description: "The persistent grant idle timeout.",
				Optional:    true,
			},
			"persistent_grant_idle_timeout_time_unit": {
				Type:        schema.TypeString,
				Description: "The persistent grant idle timeout time unit.",
				Optional:    true,
			},
			"persistent_grant_idle_timeout_type": {
				Type:        schema.TypeString,
				Description: "Allows an administrator to override the Persistent Grant Idle Timeout set globally for the OAuth AS. Defaults to SERVER_DEFAULT.",
				Optional:    true,
				Default:     "SERVER_DEFAULT",
			},
			"policy_refs": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The client registration policies.",
				Optional:    true,
			},
			"refresh_rolling": {
				Type:        schema.TypeString,
				Description: "Use ROLL or DONT_ROLL to override the Roll Refresh Token Values setting on the Authorization Server Settings. SERVER_DEFAULT will default to the Roll Refresh Token Values setting on the Authorization Server Setting screen. Defaults to SERVER_DEFAULT.",
				Optional:    true,
				Default:     "SERVER_DEFAULT",
			},
			"refresh_token_rolling_interval": {
				Type:        schema.TypeInt,
				Description: "The minimum interval to roll refresh tokens, in hours. This value will override the Refresh Token Rolling Interval Value on the Authorization Server Settings.",
				Optional:    true,
			},
			"refresh_token_rolling_interval_type": {
				Type:        schema.TypeString,
				Description: "Use OVERRIDE_SERVER_DEFAULT to override the Refresh Token Rolling Interval value on the Authorization Server Settings. SERVER_DEFAULT will default to the Refresh Token Rolling Interval value on the Authorization Server Setting. Defaults to SERVER_DEFAULT.",
				Optional:    true,
				Default:     "SERVER_DEFAULT",
			},
			"request_policy_ref": resourceLinkSchema(),
			"require_proof_key_for_code_exchange": {
				Type:        schema.TypeBool,
				Description: "Determines whether Proof Key for Code Exchange (PKCE) is required for the dynamically created client.",
				Optional:    true,
				Default:     false,
			},
			"require_signed_requests": {
				Type:        schema.TypeBool,
				Description: "Require signed requests.",
				Optional:    true,
				Default:     false,
			},
			"restrict_common_scopes": {
				Type:        schema.TypeBool,
				Description: "Restrict common scopes.",
				Optional:    true,
				Default:     false,
			},
			"restrict_to_default_access_token_manager": {
				Type:        schema.TypeBool,
				Description: "Determines whether the client is restricted to using only its default access token manager. The default is false.",
				Optional:    true,
				Default:     false,
			},
			"restricted_common_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The common scopes to restrict.",
				Optional:    true,
			},
			"rotate_client_secret": {
				Type:        schema.TypeBool,
				Description: "Rotate registration access token on dynamic client management requests.",
				Optional:    true,
				Default:     true,
			},
			"rotate_registration_access_token": {
				Type:        schema.TypeBool,
				Description: "Rotate client secret on dynamic client management requests.",
				Optional:    true,
				Default:     true,
			},
			"token_exchange_processor_policy_ref": resourceLinkSchema(),
			"user_authorization_url_override": {
				Type:        schema.TypeString,
				Description: "The URL is used as 'verification_url' and 'verification_url_complete' values in a Device Authorization request.",
				Optional:    true,
			},
		},
	}
}

//ClientRegistrationOIDCPolicy - Client Registration Open ID Connect Policy settings.
func resourceClientRegistrationOIDCPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id_token_content_encryption_algorithm": {
				Type:         schema.TypeString,
				Description:  "The JSON Web Encryption [JWE] content encryption algorithm for the ID Token.<br>AES_128_CBC_HMAC_SHA_256 - Composite AES-CBC-128 HMAC-SHA-256<br>AES_192_CBC_HMAC_SHA_384 - Composite AES-CBC-192 HMAC-SHA-384<br>AES_256_CBC_HMAC_SHA_512 - Composite AES-CBC-256 HMAC-SHA-512<br>AES-GCM-128 - AES_128_GCM<br>AES_192_GCM - AES-GCM-192<br>AES_256_GCM - AES-GCM-256",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{`AES_128_CBC_HMAC_SHA_256`, `AES_192_CBC_HMAC_SHA_384`, `AES_256_CBC_HMAC_SHA_512`, `AES_128_GCM`, `AES_192_GCM`, `AES_256_GCM`}, false),
			},
			"id_token_encryption_algorithm": {
				Type:         schema.TypeString,
				Description:  "The JSON Web Encryption [JWE] encryption algorithm used to encrypt the content encryption key for the ID Token.<br>DIR - Direct Encryption with symmetric key<br>A128KW - AES-128 Key Wrap<br>A192KW - AES-192 Key Wrap<br>A256KW - AES-256 Key Wrap<br>A128GCMKW - AES-GCM-128 key encryption<br>A192GCMKW - AES-GCM-192 key encryption<br>A256GCMKW - AES-GCM-256 key encryption<br>ECDH_ES - ECDH-ES<br>ECDH_ES_A128KW - ECDH-ES with AES-128 Key Wrap<br>ECDH_ES_A192KW - ECDH-ES with AES-192 Key Wrap<br>ECDH_ES_A256KW - ECDH-ES with AES-256 Key Wrap<br>RSA_OAEP - RSAES OAEP<br>",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{`DIR`, `A128KW`, `A192KW`, `A256KW`, `A128GCMKW`, `A192GCMKW`, `A256GCMKW`, `ECDH_ES`, `ECDH_ES_A128KW`, `ECDH_ES_A192KW`, `ECDH_ES_A256KW`, `RSA_OAEP`}, false),
			},
			"id_token_signing_algorithm": {
				Type:         schema.TypeString,
				Description:  "The JSON Web Signature [JWS] algorithm required for the ID Token.<br>NONE - No signing algorithm<br>HS256 - HMAC using SHA-256<br>HS384 - HMAC using SHA-384<br>HS512 - HMAC using SHA-512<br>RS256 - RSA using SHA-256<br>RS384 - RSA using SHA-384<br>RS512 - RSA using SHA-512<br>ES256 - ECDSA using P256 Curve and SHA-256<br>ES384 - ECDSA using P384 Curve and SHA-384<br>ES512 - ECDSA using P521 Curve and SHA-512<br>PS256 - RSASSA-PSS using SHA-256 and MGF1 padding with SHA-256<br>PS384 - RSASSA-PSS using SHA-384 and MGF1 padding with SHA-384<br>PS512 - RSASSA-PSS using SHA-512 and MGF1 padding with SHA-512<br>A null value will represent the default algorithm which is RS256.<br>RSASSA-PSS is only supported with SafeNet Luna, Thales nCipher or Java 11",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{`NONE`, `HS256`, `HS384`, `HS512`, `RS256`, `RS384`, `RS512`, `ES256`, `ES384`, `ES512`, `PS256`, `PS384`, `PS512`}, false),
			},
			"policy_group": resourceRequiredLinkSchema(),
		},
	}
}

func resourceTokenProcessorAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of this attribute.",
				Required:    true,
			},
			"masked": {
				Type:        schema.TypeBool,
				Description: "Specifies whether this attribute is masked in PingFederate logs. Defaults to `false`.",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceTokenProcessorAttributeContract() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"inherited": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether this attribute contract is inherited from its parent instance. If true, the rest of the properties in this model become read-only. The default value is false.",
			},
			"mask_ognl_values": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not all OGNL expressions used to fulfill an outgoing assertion contract should be masked in the logs. Defaults to false.",
			},
			"core_attributes": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A list of token processor attributes that correspond to the attributes exposed by the token processor type.",
				Elem:        resourceTokenProcessorAttribute(),
			},
			"extended_attributes": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of additional attributes that can be returned by the token processor. The extended attributes are only used if the token processor supports them.",
				Elem:        resourceTokenProcessorAttribute(),
			},
		},
	}
}
