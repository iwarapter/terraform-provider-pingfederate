package framework

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	_ resource.ResourceWithSchema      = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	_ resource.ResourceWithConfigure   = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
	_ resource.ResourceWithImportState = &pingfederateOauthAuthenticationPolicyContractMappingResource{}
)

type pingfederateOauthAuthenticationPolicyContractMappingResource struct {
	client *pfClient
}

func NewOauthAuthenticationPolicyContractMappingResource() resource.Resource {
	return &pingfederateOauthAuthenticationPolicyContractMappingResource{}
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = resourceApcToPersistentGrantMapping()
}

// Configure adds the client configured client to the resource.
func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*pfClient)
}

// Metadata returns the resource type name.
func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oauth_authentication_policy_contract_mapping"
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.CreateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.CreateApcMappingInput{
		BypassExternalValidation: Bool(r.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create oauthAccessTokenMapping, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenApcToPersistentGrantMapping(body)))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.GetApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.GetApcMappingInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to get oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenApcToPersistentGrantMapping(body)))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body, _, err := r.client.OauthAuthenticationPolicyContractMappings.UpdateApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.UpdateApcMappingInput{
		BypassExternalValidation: Bool(r.client.BypassExternalValidation),
		Body:                     *expandApcToPersistentGrantMapping(data),
		Id:                       data.Id.ValueString(),
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, r.versionResponseModifier(flattenApcToPersistentGrantMapping(body)))...)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ApcToPersistentGrantMappingData
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.OauthAuthenticationPolicyContractMappings.DeleteApcMappingWithContext(ctx, &oauthAuthenticationPolicyContractMappings.DeleteApcMappingInput{Id: data.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete oauthAccessTokenMapping, got error: %s", err))
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) UpgradeState(context.Context) map[int64]resource.StateUpgrader {
	schemaV0 := resourceApcToPersistentGrantMappingV0()

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schemaV0,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var mappingDataV0 ApcToPersistentGrantMappingDataV0
				resp.Diagnostics.Append(req.State.Get(ctx, &mappingDataV0)...)
				if resp.Diagnostics.HasError() {
					return
				}

				mappingDataV1 := ApcToPersistentGrantMappingData{
					Id: mappingDataV0.Id,
				}
				for _, source := range mappingDataV0.LdapAttributeSources {
					ds := LdapAttributeSourceData{
						BaseDn: source.BaseDn,
						//BinaryAttributeSettings: map[string]*BinaryLdapAttributeSettingsData{},
						DataStoreRef:        source.DataStoreRef[0].ID,
						Description:         source.Description,
						Id:                  source.Id,
						MemberOfNestedGroup: source.MemberOfNestedGroup,
						SearchAttributes:    source.SearchAttributes,
						SearchFilter:        source.SearchFilter,
						SearchScope:         source.SearchScope,
					}
					for k, v := range source.BinaryAttributeSettings {
						ds.BinaryAttributeSettings[k] = &BinaryLdapAttributeSettingsData{BinaryEncoding: v}
					}
					if len(source.AttributeContractFulfillment) > 0 {
						ds.AttributeContractFulfillment = map[string]*AttributeFulfillmentValueData{}
					}
					for _, attr := range source.AttributeContractFulfillment {
						ds.AttributeContractFulfillment[attr.KeyName.ValueString()] = &AttributeFulfillmentValueData{
							Source: &SourceTypeIdKeyData{
								Id:   attr.Source[0].Id,
								Type: attr.Source[0].Type,
							},
							Value: attr.Value,
						}
					}
					mappingDataV1.LdapAttributeSources = append(mappingDataV1.LdapAttributeSources, ds)
				}
				for _, source := range mappingDataV0.JdbcAttributeSources {
					ds := JdbcAttributeSourceData{
						ColumnNames:  source.ColumnNames,
						DataStoreRef: source.DataStoreRef[0].ID,
						Description:  source.Description,
						Filter:       source.Filter,
						Id:           source.Id,
						Schema:       source.Schema,
						Table:        source.Table,
					}
					if len(source.AttributeContractFulfillment) > 0 {
						ds.AttributeContractFulfillment = map[string]*AttributeFulfillmentValueData{}
					}
					for _, attr := range source.AttributeContractFulfillment {
						ds.AttributeContractFulfillment[attr.KeyName.ValueString()] = &AttributeFulfillmentValueData{
							Source: &attr.Source[0],
							Value:  attr.Value,
						}
					}
					mappingDataV1.JdbcAttributeSources = append(mappingDataV1.JdbcAttributeSources, ds)
				}
				for _, source := range mappingDataV0.CustomAttributeSources {
					ds := CustomAttributeSourceData{
						DataStoreRef: source.DataStoreRef[0].ID,
						Description:  source.Description,
						FilterFields: source.FilterFields,
						Id:           source.Id,
					}
					if len(source.AttributeContractFulfillment) > 0 {
						ds.AttributeContractFulfillment = map[string]*AttributeFulfillmentValueData{}
					}
					for _, attr := range source.AttributeContractFulfillment {
						ds.AttributeContractFulfillment[attr.KeyName.ValueString()] = &AttributeFulfillmentValueData{
							Source: &SourceTypeIdKeyData{
								Id:   attr.Source[0].Id,
								Type: attr.Source[0].Type,
							},
							Value: attr.Value,
						}
					}
					mappingDataV1.CustomAttributeSources = append(mappingDataV1.CustomAttributeSources, ds)
				}
				if len(mappingDataV0.AttributeContractFulfillment) > 0 {
					mappingDataV1.AttributeContractFulfillment = map[string]*AttributeFulfillmentValueData{}
				}
				for _, attr := range mappingDataV0.AttributeContractFulfillment {
					src := &SourceTypeIdKeyData{
						Id:   attr.Source[0].Id,
						Type: attr.Source[0].Type,
					}
					if src.Id.ValueString() == "" {
						src.Id = types.StringNull()
					}
					mappingDataV1.AttributeContractFulfillment[attr.KeyName.ValueString()] = &AttributeFulfillmentValueData{
						Source: src,
						Value:  attr.Value,
					}

				}
				if len(mappingDataV0.AuthenticationPolicyContractRef) == 1 {
					mappingDataV1.AuthenticationPolicyContractRef = mappingDataV0.AuthenticationPolicyContractRef[0].ID
				}
				if len(mappingDataV0.IssuanceCriteria) == 1 {
					mappingDataV1.IssuanceCriteria = &IssuanceCriteriaData{
						ExpressionCriteria: mappingDataV0.IssuanceCriteria[0].ExpressionCriteria,
					}
					if mappingDataV0.IssuanceCriteria[0].ConditionalCriteria != nil {
						mappingDataV1.IssuanceCriteria.ConditionalCriteria = &[]*ConditionalIssuanceCriteriaEntryData{}
						for _, data := range *mappingDataV0.IssuanceCriteria[0].ConditionalCriteria {
							d := &ConditionalIssuanceCriteriaEntryData{
								AttributeName: data.AttributeName,
								Condition:     data.Condition,
								ErrorResult:   data.ErrorResult,
								Value:         data.Value,
							}
							if len(data.Source) == 1 {
								d.Source = &SourceTypeIdKeyData{
									Id:   data.Source[0].Id,
									Type: data.Source[0].Type,
								}
							}
							*mappingDataV1.IssuanceCriteria.ConditionalCriteria = append(*mappingDataV1.IssuanceCriteria.ConditionalCriteria, d)
						}
					}
				}
				resp.Diagnostics.Append(resp.State.Set(ctx, &mappingDataV1)...)
			},
		},
	}
}

func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// old version of pingfederate have different defaults which we need to handle
func (r *pingfederateOauthAuthenticationPolicyContractMappingResource) versionResponseModifier(data *ApcToPersistentGrantMappingData) *ApcToPersistentGrantMappingData {
	if r.client.IsVersionLessEqThan(10, 2) {
		for i := range data.LdapAttributeSources {
			if data.LdapAttributeSources[i].BaseDn.IsNull() {
				data.LdapAttributeSources[i].BaseDn = types.StringValue("")
			}
		}
		if data.IssuanceCriteria != nil {
			if data.IssuanceCriteria.ConditionalCriteria != nil {
				for i := range *data.IssuanceCriteria.ConditionalCriteria {
					if (*data.IssuanceCriteria.ConditionalCriteria)[i].ErrorResult.IsNull() {
						(*data.IssuanceCriteria.ConditionalCriteria)[i].ErrorResult = types.StringValue("")
					}
				}
			}
			if data.IssuanceCriteria.ExpressionCriteria != nil {
				for i := range *data.IssuanceCriteria.ExpressionCriteria {
					if (*data.IssuanceCriteria.ExpressionCriteria)[i].ErrorResult.IsNull() {
						(*data.IssuanceCriteria.ExpressionCriteria)[i].ErrorResult = types.StringValue("")
					}
				}
			}
		}
	}
	return data
}

func resourceApcToPersistentGrantMappingV0() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"authentication_policy_contract_ref": resourceLinkSchemaV0(),
			"attribute_contract_fulfillment":     attributeContractFulfilmentSchemaV0(),
			"jdbc_attribute_source": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_contract_fulfillment": attributeContractFulfilmentSchemaV0(),
						"column_names": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
						},
						"data_store_ref": resourceLinkSchemaV0(),
						"description":    schema.StringAttribute{Optional: true},
						"filter":         schema.StringAttribute{Optional: true},
						"id":             schema.StringAttribute{Optional: true},
						"schema":         schema.StringAttribute{Optional: true},
						"table":          schema.StringAttribute{Optional: true},
					},
				},
			},
			"ldap_attribute_source": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_contract_fulfillment": attributeContractFulfilmentSchemaV0(),
						"base_dn":                        schema.StringAttribute{Optional: true},
						"binary_attribute_settings":      schema.MapAttribute{Optional: true, ElementType: types.StringType},
						"data_store_ref":                 resourceLinkSchemaV0(),
						"description":                    schema.StringAttribute{Optional: true},
						"id":                             schema.StringAttribute{Optional: true},
						"member_of_nested_group":         schema.BoolAttribute{Optional: true},
						"search_attributes":              schema.ListAttribute{Optional: true, ElementType: types.StringType},
						"search_filter":                  schema.StringAttribute{Optional: true},
						"search_scope":                   schema.StringAttribute{Optional: true},
					},
				},
			},
			"custom_attribute_source": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attribute_contract_fulfillment": attributeContractFulfilmentSchemaV0(),
						"filter_fields": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name":  schema.StringAttribute{Required: true},
									"value": schema.StringAttribute{Optional: true},
								},
							},
						},
						"data_store_ref": resourceLinkSchemaV0(),
						"description":    schema.StringAttribute{Optional: true},
						"id":             schema.StringAttribute{Optional: true},
					},
				},
			},
			"id": schema.StringAttribute{Optional: true},
			"issuance_criteria": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"conditional_criteria": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"attribute_name": schema.StringAttribute{
										Required: true,
									},
									"condition": schema.StringAttribute{
										Required: true,
									},
									"error_result": schema.StringAttribute{
										Optional: true,
									},
									"source": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: singleSourceTypeIdKey(),
										},
									},
									"value": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
						"expression_criteria": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: listExpressionIssuanceCriteriaEntry(),
							},
						},
					},
				},
			},
		},
	}
}

func attributeContractFulfilmentSchemaV0() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"key_name": schema.StringAttribute{
					Required: true,
				},
				"source": schema.ListNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: singleSourceTypeIdKey(),
					},
					Required: true,
				},
				"value": schema.StringAttribute{
					Required: true,
				},
			},
		},
	}
}

func resourceLinkSchemaV0() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		Optional: true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: legacyResourceLinkSchema(),
		},
	}
}

type ApcToPersistentGrantMappingDataV0 struct {
	Id                              types.String                      `tfsdk:"id"`
	AuthenticationPolicyContractRef []ResourceLink                    `tfsdk:"authentication_policy_contract_ref"`
	JdbcAttributeSources            []JdbcAttributeSourceDataV0       `tfsdk:"jdbc_attribute_source"`
	LdapAttributeSources            []LdapAttributeSourceDataV0       `tfsdk:"ldap_attribute_source"`
	CustomAttributeSources          []CustomAttributeSourceDataV0     `tfsdk:"custom_attribute_source"`
	IssuanceCriteria                []IssuanceCriteriaDataV0          `tfsdk:"issuance_criteria"`
	AttributeContractFulfillment    []AttributeFulfillmentValueDataV0 `tfsdk:"attribute_contract_fulfillment"`
}

type JdbcAttributeSourceDataV0 struct {
	AttributeContractFulfillment []AttributeFulfillmentValueDataV0 `tfsdk:"attribute_contract_fulfillment"`
	ColumnNames                  []types.String                    `tfsdk:"column_names"`
	DataStoreRef                 []ResourceLink                    `tfsdk:"data_store_ref"`
	Description                  types.String                      `tfsdk:"description"`
	Filter                       types.String                      `tfsdk:"filter"`
	Id                           types.String                      `tfsdk:"id"`
	Schema                       types.String                      `tfsdk:"schema"`
	Table                        types.String                      `tfsdk:"table"`
}

type LdapAttributeSourceDataV0 struct {
	AttributeContractFulfillment []AttributeFulfillmentValueDataV0 `tfsdk:"attribute_contract_fulfillment"`
	BaseDn                       types.String                      `tfsdk:"base_dn"`
	BinaryAttributeSettings      map[string]types.String           `tfsdk:"binary_attribute_settings"`
	DataStoreRef                 []ResourceLink                    `tfsdk:"data_store_ref"`
	Description                  types.String                      `tfsdk:"description"`
	Id                           types.String                      `tfsdk:"id"`
	MemberOfNestedGroup          types.Bool                        `tfsdk:"member_of_nested_group"`
	SearchAttributes             []types.String                    `tfsdk:"search_attributes"`
	SearchFilter                 types.String                      `tfsdk:"search_filter"`
	SearchScope                  types.String                      `tfsdk:"search_scope"`
}

type CustomAttributeSourceDataV0 struct {
	AttributeContractFulfillment []AttributeFulfillmentValueDataV0 `tfsdk:"attribute_contract_fulfillment"`
	DataStoreRef                 []ResourceLink                    `tfsdk:"data_store_ref"`
	Description                  types.String                      `tfsdk:"description"`
	FilterFields                 *[]*FieldEntryData                `tfsdk:"filter_fields"`
	Id                           types.String                      `tfsdk:"id"`
}

type AttributeFulfillmentValueDataV0 struct {
	KeyName types.String          `tfsdk:"key_name"`
	Source  []SourceTypeIdKeyData `tfsdk:"source"`
	Value   types.String          `tfsdk:"value"`
}

type IssuanceCriteriaDataV0 struct {
	ConditionalCriteria *[]*ConditionalIssuanceCriteriaEntryDataV0 `tfsdk:"conditional_criteria"`
	ExpressionCriteria  *[]*ExpressionIssuanceCriteriaEntryData    `tfsdk:"expression_criteria"`
}

type ConditionalIssuanceCriteriaEntryDataV0 struct {
	AttributeName types.String          `tfsdk:"attribute_name"`
	Condition     types.String          `tfsdk:"condition"`
	ErrorResult   types.String          `tfsdk:"error_result"`
	Source        []SourceTypeIdKeyData `tfsdk:"source"`
	Value         types.String          `tfsdk:"value"`
}
