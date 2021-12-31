package pingfederate

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func resourcePingFederateKeypairResourceReadData(d *schema.ResourceData) *pf.NewKeyPairSettings {
	settings := pf.NewKeyPairSettings{
		CommonName:   String(d.Get("common_name").(string)),
		Country:      String(d.Get("country").(string)),
		KeyAlgorithm: String(d.Get("key_algorithm").(string)),
		KeySize:      Int(d.Get("key_size").(int)),
		Organization: String(d.Get("organization").(string)),
		ValidDays:    Int(d.Get("valid_days").(int)),
	}
	if val, ok := d.GetOk("city"); ok {
		settings.City = String(val.(string))
	}
	if val, ok := d.GetOk("organization_unit"); ok {
		settings.OrganizationUnit = String(val.(string))
	}
	if val, ok := d.GetOk("state"); ok {
		settings.State = String(val.(string))
	}
	if val, ok := d.GetOk("crypto_provider"); ok {
		settings.CryptoProvider = String(val.(string))
	}
	if val, ok := d.GetOk("subject_alternative_names"); ok {
		sans := expandStringList(val.(*schema.Set).List())
		settings.SubjectAlternativeNames = &sans
	}

	return &settings
}

func expandExtendedProperties(in []interface{}) *[]*pf.ExtendedProperty {
	var propertyList []*pf.ExtendedProperty
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ExtendedProperty{
			Name:        String(l["name"].(string)),
			Description: String(l["description"].(string)),
			MultiValued: Bool(l["multi_valued"].(bool)),
		}
		propertyList = append(propertyList, s)
	}
	return &propertyList
}

func expandScopes(in []interface{}) *[]*pf.ScopeEntry {
	var scopeList []*pf.ScopeEntry
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ScopeEntry{
			Name:        String(l["name"].(string)),
			Description: String(l["description"].(string)),
			Dynamic:     Bool(l["dynamic"].(bool)),
		}
		scopeList = append(scopeList, s)
	}
	return &scopeList
}

func expandScopeGroups(in []interface{}) *[]*pf.ScopeGroupEntry {
	var scopeGroupList []*pf.ScopeGroupEntry
	for _, raw := range in {
		l := raw.(map[string]interface{})
		s := &pf.ScopeGroupEntry{
			Name:        String(l["name"].(string)),
			Description: String(l["description"].(string)),
		}
		var scopes []*string
		for _, scope := range l["scopes"].([]interface{}) {
			scopes = append(scopes, String(scope.(string)))
		}
		s.Scopes = &scopes
		scopeGroupList = append(scopeGroupList, s)
	}
	return &scopeGroupList
}

func expandClientAuth(in []interface{}) *pf.ClientAuth {
	ca := &pf.ClientAuth{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["client_cert_issuer_dn"]; ok && val.(string) != "" {
			ca.ClientCertIssuerDn = String(val.(string))
		}
		if val, ok := l["client_cert_subject_dn"]; ok && val.(string) != "" {
			ca.ClientCertSubjectDn = String(val.(string))
		}
		if val, ok := l["enforce_replay_prevention"]; ok {
			ca.EnforceReplayPrevention = Bool(val.(bool))
		}
		if val, ok := l["secret"]; ok {
			ca.Secret = String(val.(string))
		}
		if val, ok := l["token_endpoint_auth_signing_algorithm"]; ok && val.(string) != "" {
			ca.TokenEndpointAuthSigningAlgorithm = String(val.(string))
		}
		ca.Type = String(l["type"].(string))
	}
	return ca
}

func expandJwksSettings(in []interface{}) *pf.JwksSettings {
	ca := &pf.JwksSettings{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["jwks"]; ok {
			ca.Jwks = String(val.(string))
		}
		if val, ok := l["jwks_url"]; ok {
			ca.JwksUrl = String(val.(string))
		}
	}
	return ca
}

//func expandResourceLink(in []interface{}) *pf.ResourceLink {
//	ca := &pf.ResourceLink{}
//	for _, raw := range in {
//		l := raw.(map[string]interface{})
//		if val, ok := l["id"]; ok {
//			ca.Id = String(val.(string))
//		}
//		if val, ok := l["location"]; ok && val.(string) != "" {
//			ca.Location = String(val.(string))
//		}
//	}
//	return ca
//}

func expandClientOIDCPolicy(in []interface{}) *pf.ClientOIDCPolicy {
	ca := &pf.ClientOIDCPolicy{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["grant_access_session_revocation_api"]; ok {
			ca.GrantAccessSessionRevocationApi = Bool(val.(bool))
		}
		if val, ok := l["id_token_signing_algorithm"]; ok && val.(string) != "" {
			ca.IdTokenSigningAlgorithm = String(val.(string))
		}
		if val, ok := l["id_token_encryption_algorithm"]; ok && val.(string) != "" {
			ca.IdTokenEncryptionAlgorithm = String(val.(string))
		}
		if val, ok := l["id_token_content_encryption_algorithm"]; ok && val.(string) != "" {
			ca.IdTokenContentEncryptionAlgorithm = String(val.(string))
		}
		if val, ok := l["sector_identifier_uri"]; ok && val.(string) != "" {
			ca.SectorIdentifierUri = String(val.(string))
		}
		if val, ok := l["logout_uris"]; ok {
			str := expandStringList(val.([]interface{}))
			ca.LogoutUris = &str
		}
		if val, ok := l["ping_access_logout_capable"]; ok {
			ca.PingAccessLogoutCapable = Bool(val.(bool))
		}
		if val, ok := l["pairwise_identifier_user_type"]; ok {
			ca.PairwiseIdentifierUserType = Bool(val.(bool))
		}
		if val, ok := l["policy_group"]; ok && len(val.([]interface{})) > 0 {
			ca.PolicyGroup = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
	}
	return ca
}

func expandConfigFields(in []interface{}) *[]*pf.ConfigField {
	var configFields []*pf.ConfigField
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if _, ok := l["encrypted_value"]; !ok {
			s := &pf.ConfigField{
				Name: String(l["name"].(string)),
			}
			if val, ok := l["value"]; ok {
				s.Value = String(val.(string))
			}
			if val, ok := l["inherited"]; ok {
				s.Inherited = Bool(val.(bool))
			}
			configFields = append(configFields, s)
		}
	}
	return &configFields
}

func expandSensitiveConfigFields(in []interface{}) *[]*pf.ConfigField {
	var configFields []*pf.ConfigField
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["value"]; ok && val.(string) != "" {
			s := &pf.ConfigField{
				Name: String(l["name"].(string)),
			}
			if val, ok := l["value"]; ok {
				s.Value = String(val.(string))
			}
			if val, ok := l["inherited"]; ok {
				s.Inherited = Bool(val.(bool))
			}
			configFields = append(configFields, s)
		}
	}
	return &configFields
}

//func expandConfigRow(in []interface{}) *[]*pf.ConfigRow {
//	configRows := []*pf.ConfigRow{}
//	for _, raw := range in {
//		l := raw.(map[string]interface{})
//		row := &pf.ConfigRow{}
//		if val, ok := l["fields"]; ok {
//			row.Fields = expandConfigFields(val.(*schema.Set).List())
//		}
//		if val, ok := l["sensitive_fields"]; ok {
//			fields := expandSensitiveConfigFields(val.(*schema.Set).List())
//			*row.Fields = append(*row.Fields, *fields...)
//		}
//		configRows = append(configRows, row)
//	}
//	return &configRows
//}

//func expandConfigTable(in []interface{}) *[]*pf.ConfigTable {
//	var configTables []*pf.ConfigTable
//	for _, raw := range in {
//		l := raw.(map[string]interface{})
//		s := &pf.ConfigTable{
//			Name: String(l["name"].(string)),
//		}
//		if val, ok := l["rows"]; ok {
//			s.Rows = expandConfigRow(val.([]interface{}))
//		}
//		if val, ok := l["inherited"]; ok {
//			s.Inherited = Bool(val.(bool))
//		}
//		configTables = append(configTables, s)
//	}
//	return &configTables
//}

func expandPluginConfiguration(in []interface{}) *pf.PluginConfiguration {
	config := &pf.PluginConfiguration{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["tables"]; ok && len(val.([]interface{})) > 0 {
			config.Tables = expandConfigTableList(val.([]interface{}))
		}
		if val, ok := l["fields"]; ok && len(val.(*schema.Set).List()) > 0 {
			config.Fields = expandConfigFields(val.(*schema.Set).List())
		}
		if val, ok := l["sensitive_fields"]; ok && len(val.(*schema.Set).List()) > 0 {
			fields := expandSensitiveConfigFields(val.(*schema.Set).List())
			*config.Fields = append(*config.Fields, *fields...)
		}
	}
	return config
}

func expandAccessTokenAttributeContract(in []interface{}) *pf.AccessTokenAttributeContract {
	pgc := &pf.AccessTokenAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		var atr []*pf.AccessTokenAttribute
		for _, exAtr := range l["extended_attributes"].(*schema.Set).List() {
			atr = append(atr, &pf.AccessTokenAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func expandAuthenticationPolicyContractAttribute(in []interface{}) *[]*pf.AuthenticationPolicyContractAttribute {
	var contractList []*pf.AuthenticationPolicyContractAttribute
	for _, raw := range in {
		c := &pf.AuthenticationPolicyContractAttribute{
			Name: String(raw.(string)),
		}
		contractList = append(contractList, c)
	}
	return &contractList
}

func expandPasswordCredentialValidatorAttribute(in []interface{}) *[]*pf.PasswordCredentialValidatorAttribute {
	contractList := []*pf.PasswordCredentialValidatorAttribute{}
	for _, raw := range in {
		c := &pf.PasswordCredentialValidatorAttribute{
			Name: String(raw.(string)),
		}
		contractList = append(contractList, c)
	}
	return &contractList
}

func expandJdbcTagConfigs(in []interface{}) *[]*pf.JdbcTagConfig {
	var tags []*pf.JdbcTagConfig
	for _, raw := range in {
		l := raw.(map[string]interface{})
		f := &pf.JdbcTagConfig{}
		if v, ok := l["connection_url"]; ok {
			f.ConnectionUrl = String(v.(string))
		}
		if v, ok := l["tags"]; ok {
			f.Tags = String(v.(string))
		}
		if v, ok := l["default_source"]; ok {
			f.DefaultSource = Bool(v.(bool))
		}
		tags = append(tags, f)
	}
	return &tags
}

func expandAuthenticationSelectorAttributeContract(in []interface{}) *pf.AuthenticationSelectorAttributeContract {
	pgc := &pf.AuthenticationSelectorAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		var atr []*pf.AuthenticationSelectorAttribute
		for _, exAtr := range l["extended_attributes"].(*schema.Set).List() {
			atr = append(atr, &pf.AuthenticationSelectorAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func expandAccessTokenMappingContext(in []interface{}) *pf.AccessTokenMappingContext {
	pgc := &pf.AccessTokenMappingContext{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["type"]; ok {
			pgc.Type = String(val.(string))
		}
		if val, ok := l["context_ref"]; ok && len(val.([]interface{})) > 0 {
			pgc.ContextRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
		}
	}
	return pgc
}

func expandOpenIdConnectAttributes(in []interface{}) *[]*pf.OpenIdConnectAttribute {
	attributes := &[]*pf.OpenIdConnectAttribute{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		c := &pf.OpenIdConnectAttribute{}
		if val, ok := l["name"]; ok {
			c.Name = String(val.(string))
		}
		if isOverrideDefaultDelivery(l) {
			// because the terraform-plugin-sdk will not distinguish between an unset bool and false we need to check if our computed value is true
			// to prevent attempting to override with false,false as this is incorrect.
			if val, ok := l["include_in_id_token"]; ok {
				c.IncludeInIdToken = Bool(val.(bool))
			}
			if val, ok := l["include_in_user_info"]; ok {
				c.IncludeInUserInfo = Bool(val.(bool))
			}
		}
		*attributes = append(*attributes, c)
	}
	return attributes
}

func isOverrideDefaultDelivery(d map[string]interface{}) bool {
	if val, ok := d["include_in_id_token"]; ok && val.(bool) {
		return true
	}
	if val, ok := d["include_in_user_info"]; ok && val.(bool) {
		return true
	}
	return false
}

func expandOpenIdConnectAttributeContract(in []interface{}) *pf.OpenIdConnectAttributeContract {
	iac := &pf.OpenIdConnectAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["extended_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			iac.ExtendedAttributes = expandOpenIdConnectAttributes(v.(*schema.Set).List())
		}
		if v, ok := l["core_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			iac.CoreAttributes = expandOpenIdConnectAttributes(v.(*schema.Set).List())
		}
	}
	return iac
}

func expandAttributeMapping(in []interface{}) *pf.AttributeMapping {
	iac := &pf.AttributeMapping{AttributeSources: &[]*pf.AttributeSource{}}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["attribute_contract_fulfillment"]; ok {
			iac.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(v.(*schema.Set).List())
		}
		if v, ok := l["issuance_criteria"]; ok && len(v.([]interface{})) > 0 {
			iac.IssuanceCriteria = expandIssuanceCriteria(v.([]interface{})[0].(map[string]interface{}))
		}

		if v, ok := l["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
			*iac.AttributeSources = append(*iac.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
		}
		if v, ok := l["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
			*iac.AttributeSources = append(*iac.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
		}
		if v, ok := l["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
			*iac.AttributeSources = append(*iac.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
		}

	}
	return iac
}

func expandAuthenticationSources(in []interface{}) *[]*pf.AuthenticationSource {
	sources := &[]*pf.AuthenticationSource{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		*sources = append(*sources, expandAuthenticationSource(l))
	}
	return sources
}

func expandAuthenticationSource(in map[string]interface{}) *pf.AuthenticationSource {
	src := &pf.AuthenticationSource{}
	if v, ok := in["type"]; ok {
		src.Type = String(v.(string))
	}
	if v, ok := in["source_ref"]; ok {
		src.SourceRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
	}
	return src
}

func expandAuthenticationPolicyTrees(in []interface{}) *[]*pf.AuthenticationPolicyTree {
	trees := &[]*pf.AuthenticationPolicyTree{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		src := &pf.AuthenticationPolicyTree{}
		if v, ok := l["name"]; ok {
			src.Name = String(v.(string))
		}
		if v, ok := l["description"]; ok && v.(string) != "" {
			src.Description = String(v.(string))
		}
		if v, ok := l["authentication_api_application_ref"]; ok && len(v.([]interface{})) > 0 {
			src.AuthenticationApiApplicationRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["enabled"]; ok {
			src.Enabled = Bool(v.(bool))
		}
		if v, ok := l["root_node"]; ok && len(v.([]interface{})) > 0 {
			m := v.([]interface{})[0].(map[string]interface{})
			src.RootNode = expandAuthenticationPolicyTreeNode(m)
		}
		*trees = append(*trees, src)
	}
	return trees
}

func expandAuthenticationPolicyTreeNodes(in []interface{}) *[]*pf.AuthenticationPolicyTreeNode {
	nodes := &[]*pf.AuthenticationPolicyTreeNode{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		*nodes = append(*nodes, expandAuthenticationPolicyTreeNode(l))
	}
	return nodes
}

func expandAuthenticationPolicyTreeNode(in map[string]interface{}) *pf.AuthenticationPolicyTreeNode {
	node := &pf.AuthenticationPolicyTreeNode{}
	if v, ok := in["action"]; ok {
		node.Action = expandPolicyAction(v.([]interface{}))
	}
	if v, ok := in["children"]; ok && len(v.([]interface{})) > 0 {
		node.Children = expandAuthenticationPolicyTreeNodes(v.([]interface{}))
	}
	return node
}

func expandPolicyAction(in []interface{}) *pf.PolicyAction {
	action := &pf.PolicyAction{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["type"]; ok {
			action.Type = String(v.(string))
		}
		if v, ok := l["context"]; ok && v.(string) != "" {
			action.Context = String(v.(string))
		}
		if v, ok := l["attribute_mapping"]; ok && len(v.([]interface{})) > 0 {
			action.AttributeMapping = expandAttributeMapping(v.([]interface{}))
		}
		if v, ok := l["authentication_policy_contract_ref"]; ok && len(v.([]interface{})) > 0 {
			action.AuthenticationPolicyContractRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["authentication_selector_ref"]; ok && len(v.([]interface{})) > 0 {
			action.AuthenticationSelectorRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["local_identity_ref"]; ok && len(v.([]interface{})) > 0 {
			action.LocalIdentityRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["inbound_mapping"]; ok && len(v.([]interface{})) > 0 {
			action.InboundMapping = expandAttributeMapping(v.([]interface{}))
		}
		if v, ok := l["outbound_attribute_mapping"]; ok && len(v.([]interface{})) > 0 {
			action.OutboundAttributeMapping = expandAttributeMapping(v.([]interface{}))
		}
		if v, ok := l["authentication_source"]; ok && len(v.([]interface{})) > 0 {
			action.AuthenticationSource = expandAuthenticationSource(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["input_user_id_mapping"]; ok && len(v.([]interface{})) > 0 {
			action.InputUserIdMapping = expandAttributeFulfillmentValue(v.([]interface{})[0].(map[string]interface{}))
		}
		if v, ok := l["attribute_rules"]; ok && len(v.([]interface{})) > 0 {
			action.AuthnSourcePolicyAction.AttributeRules = expandAttributeRules(v.([]interface{}))
		}
		if v, ok := l["fragment_mapping"]; ok && len(v.([]interface{})) > 0 {
			action.FragmentMapping = expandAttributeMapping(v.([]interface{}))
		}
		if v, ok := l["fragment"]; ok && len(v.([]interface{})) > 0 {
			action.Fragment = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		}
	}
	return action
}

func expandAttributeRules(in []interface{}) *pf.AttributeRules {
	rules := &pf.AttributeRules{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["fallback_to_success"]; ok {
			rules.FallbackToSuccess = Bool(v.(bool))
		}
		if v, ok := l["items"]; ok && v.(*schema.Set).Len() > 0 {
			rules.Items = expandAttributeRuleSlice(v.(*schema.Set).List())
		}
	}
	return rules
}

func expandAttributeRuleSlice(in []interface{}) *[]*pf.AttributeRule {
	rules := &[]*pf.AttributeRule{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		r := &pf.AttributeRule{}
		if v, ok := l["attribute_name"]; ok && v.(string) != "" {
			r.AttributeName = String(v.(string))
		}
		if v, ok := l["expected_value"]; ok && v.(string) != "" {
			r.ExpectedValue = String(v.(string))
		}
		if v, ok := l["result"]; ok && v.(string) != "" {
			r.Result = String(v.(string))
		}
		if v, ok := l["condition"]; ok && v.(string) != "" {
			r.Condition = String(v.(string))
		}
		*rules = append(*rules, r)
	}
	return rules
}

func expandMapOfParameterValues(in []interface{}) map[string]*pf.ParameterValues {
	ca := map[string]*pf.ParameterValues{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["key_name"]; ok {
			ca[v.(string)] = expandParameterValues(l)
		}
	}
	return ca
}

//func expandParameterValues(in map[string]interface{}) *pf.ParameterValues {
//	ca := &pf.ParameterValues{}
//	if v, ok := in["values"]; ok && v != "" {
//		strs := expandStringList(v.(*schema.Set).List())
//		ca.Values = &strs
//	}
//	return ca
//}

func expandRolesAndProtcols(in []interface{}) *pf.RolesAndProtocols {
	roles := &pf.RolesAndProtocols{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable_idp_discovery"]; ok {
			roles.EnableIdpDiscovery = Bool(v.(bool)) //expandAttributeRuleSlice(v.(*schema.Set).List())
		}
		if v, ok := l["oauth_role"]; ok {
			roles.OauthRole = expandOauthRole(v.([]interface{}))
		}
		if v, ok := l["idp_role"]; ok {
			roles.IdpRole = expandIdpRole(v.([]interface{}))
		}
		if v, ok := l["sp_role"]; ok {
			roles.SpRole = expandSpRole(v.([]interface{}))
		}
	}
	return roles
}

func expandOauthRole(in []interface{}) *pf.OAuthRole {
	roles := &pf.OAuthRole{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable_oauth"]; ok {
			roles.EnableOauth = Bool(v.(bool))
		}
		if v, ok := l["enable_openid_connect"]; ok {
			roles.EnableOpenIdConnect = Bool(v.(bool))
		}
	}
	return roles
}

func expandIdpRole(in []interface{}) *pf.IdpRole {
	roles := &pf.IdpRole{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable"]; ok {
			roles.Enable = Bool(v.(bool))
		}
		if v, ok := l["saml20_profile"]; ok {
			roles.Saml20Profile = expandSaml20Profile(v.([]interface{}))
		}
		if v, ok := l["enable_outbound_provisioning"]; ok {
			roles.EnableOutboundProvisioning = Bool(v.(bool))
		}
		if v, ok := l["enable_saml11"]; ok {
			roles.EnableSaml11 = Bool(v.(bool))
		}
		if v, ok := l["enable_saml10"]; ok {
			roles.EnableSaml10 = Bool(v.(bool))
		}
		if v, ok := l["enable_ws_fed"]; ok {
			roles.EnableWsFed = Bool(v.(bool))
		}
		if v, ok := l["enable_ws_trust"]; ok {
			roles.EnableWsTrust = Bool(v.(bool))
		}
	}
	return roles
}

func expandSpRole(in []interface{}) *pf.SpRole {
	roles := &pf.SpRole{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable"]; ok {
			roles.Enable = Bool(v.(bool))
		}
		if v, ok := l["saml20_profile"]; ok {
			roles.Saml20Profile = expandSpSaml20Profile(v.([]interface{}))
		}
		if v, ok := l["enable_inbound_provisioning"]; ok {
			roles.EnableInboundProvisioning = Bool(v.(bool))
		}
		if v, ok := l["enable_saml11"]; ok {
			roles.EnableSaml11 = Bool(v.(bool))
		}
		if v, ok := l["enable_saml10"]; ok {
			roles.EnableSaml10 = Bool(v.(bool))
		}
		if v, ok := l["enable_ws_fed"]; ok {
			roles.EnableWsFed = Bool(v.(bool))
		}
		if v, ok := l["enable_ws_trust"]; ok {
			roles.EnableWsTrust = Bool(v.(bool))
		}
		if v, ok := l["enable_openid_connect"]; ok {
			roles.EnableOpenIDConnect = Bool(v.(bool))
		}
	}
	return roles
}

func expandSpSaml20Profile(in []interface{}) *pf.SpSAML20Profile {
	roles := &pf.SpSAML20Profile{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable"]; ok {
			roles.Enable = Bool(v.(bool))
		}
		if v, ok := l["enable_xasp"]; ok {
			roles.EnableXASP = Bool(v.(bool))
		}
		if v, ok := l["enable_auto_connect"]; ok {
			roles.EnableAutoConnect = Bool(v.(bool))
		}
	}
	return roles
}

func expandSaml20Profile(in []interface{}) *pf.SAML20Profile {
	roles := &pf.SAML20Profile{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["enable"]; ok {
			roles.Enable = Bool(v.(bool))
		}
		if v, ok := l["enable_auto_connect"]; ok {
			roles.EnableAutoConnect = Bool(v.(bool))
		}
	}
	return roles
}

func expandFederationInfo(in []interface{}) *pf.FederationInfo {
	info := &pf.FederationInfo{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["base_url"]; ok {
			info.BaseUrl = String(v.(string))
		}
		if v, ok := l["saml2_entity_id"]; ok {
			info.Saml2EntityId = String(v.(string))
		}
		if v, ok := l["saml1x_issuer_id"]; ok {
			info.Saml1xIssuerId = String(v.(string))
		}
		if v, ok := l["saml1x_source_id"]; ok {
			info.Saml1xSourceId = String(v.(string))
		}
		if v, ok := l["wsfed_realm"]; ok {
			info.WsfedRealm = String(v.(string))
		}
	}
	return info
}

func expandMapOfAttributeFulfillmentValue(in []interface{}) map[string]*pf.AttributeFulfillmentValue {
	ca := map[string]*pf.AttributeFulfillmentValue{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["key_name"]; ok {
			ca[v.(string)] = expandAttributeFulfillmentValue(l)
		}
	}
	return ca
}

func expandPersistentGrantContract(in []interface{}) *pf.PersistentGrantContract {
	pgc := &pf.PersistentGrantContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		atr := make([]*pf.PersistentGrantAttribute, 0)
		for _, exAtr := range l["extended_attributes"].([]interface{}) {
			atr = append(atr, &pf.PersistentGrantAttribute{Name: String(exAtr.(string))})
		}
		pgc.ExtendedAttributes = &atr
	}
	return pgc
}

func expandPasswordCredentialValidatorAttributeContract(in []interface{}) *pf.PasswordCredentialValidatorAttributeContract {
	pgc := &pf.PasswordCredentialValidatorAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["inherited"]; ok {
			pgc.Inherited = Bool(val.(bool))
		}
		if v, ok := l["extended_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			pgc.ExtendedAttributes = expandPasswordCredentialValidatorAttribute(v.(*schema.Set).List())
		}
		if v, ok := l["core_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			pgc.CoreAttributes = expandPasswordCredentialValidatorAttribute(v.(*schema.Set).List())
		}
	}
	return pgc
}

func expandIdpAdapterAttribute(in map[string]interface{}) *pf.IdpAdapterAttribute {
	var result pf.IdpAdapterAttribute
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["pseudonym"]; ok {
		result.Pseudonym = Bool(val.(bool))
	}
	if val, ok := in["masked"]; ok {
		result.Masked = Bool(val.(bool))
	}
	return &result
}

func expandSourceTypeIdKey(in map[string]interface{}) *pf.SourceTypeIdKey {
	var result pf.SourceTypeIdKey
	if val, ok := in["type"]; ok {
		result.Type = String(val.(string))
	}
	if val, ok := in["id"]; ok && val.(string) != "" {
		result.Id = String(val.(string))
	}
	return &result
}

//func expandPluginConfiguration(in map[string]interface{}) *pf.PluginConfiguration {
//	var result pf.PluginConfiguration
//	if val, ok := in["tables"]; ok {
//		result.Tables = expandConfigTableList(val.([]interface{}))
//	}
//	if val, ok := in["fields"]; ok {
//		result.Fields = expandConfigFieldList(val.([]interface{}))
//	}
//	return &result
//}

func expandChannelSourceLocation(in map[string]interface{}) *pf.ChannelSourceLocation {
	var result pf.ChannelSourceLocation
	if val, ok := in["group_dn"]; ok {
		result.GroupDN = String(val.(string))
	}
	if val, ok := in["filter"]; ok {
		result.Filter = String(val.(string))
	}
	if val, ok := in["nested_search"]; ok {
		result.NestedSearch = Bool(val.(bool))
	}
	return &result
}

func expandConnectionCredentials(in map[string]interface{}) *pf.ConnectionCredentials {
	var result pf.ConnectionCredentials
	if val, ok := in["signing_settings"]; ok && len(val.([]interface{})) > 0 {
		result.SigningSettings = expandSigningSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["decryption_key_pair_ref"]; ok && len(val.([]interface{})) > 0 {
		result.DecryptionKeyPairRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["verification_issuer_dn"]; ok && val.(string) != "" {
		result.VerificationIssuerDN = String(val.(string))
	}
	if val, ok := in["key_transport_algorithm"]; ok && val.(string) != "" {
		result.KeyTransportAlgorithm = String(val.(string))
	}
	if val, ok := in["block_encryption_algorithm"]; ok && val.(string) != "" {
		result.BlockEncryptionAlgorithm = String(val.(string))
	}
	if val, ok := in["secondary_decryption_key_pair_ref"]; ok && len(val.([]interface{})) > 0 {
		result.SecondaryDecryptionKeyPairRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["outbound_back_channel_auth"]; ok && len(val.([]interface{})) > 0 {
		result.OutboundBackChannelAuth = expandOutboundBackChannelAuth(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["inbound_back_channel_auth"]; ok && len(val.([]interface{})) > 0 {
		result.InboundBackChannelAuth = expandInboundBackChannelAuth(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["verification_subject_dn"]; ok && val.(string) != "" {
		result.VerificationSubjectDN = String(val.(string))
	}
	if val, ok := in["certs"]; ok && len(val.([]interface{})) > 0 {
		result.Certs = expandConnectionCertList(val.([]interface{}))
	}
	return &result
}

func expandChannel(in map[string]interface{}) *pf.Channel {
	var result pf.Channel
	if val, ok := in["active"]; ok {
		result.Active = Bool(val.(bool))
	}
	if val, ok := in["channel_source"]; ok && len(val.([]interface{})) > 0 {
		result.ChannelSource = expandChannelSource(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["attribute_mapping"]; ok {
		result.AttributeMapping = expandSaasAttributeMappingList(val.([]interface{}))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["max_threads"]; ok {
		result.MaxThreads = Int(val.(int))
	}
	if val, ok := in["timeout"]; ok {
		result.Timeout = Int(val.(int))
	}
	return &result
}

func expandSpAdapterAttributeContract(in map[string]interface{}) *pf.SpAdapterAttributeContract {
	var result pf.SpAdapterAttributeContract
	if val, ok := in["core_attributes"]; ok && len(val.(*schema.Set).List()) > 0 {
		result.CoreAttributes = expandSpAdapterAttributeList(val.(*schema.Set).List())
	}
	if val, ok := in["extended_attributes"]; ok && len(val.(*schema.Set).List()) > 0 {
		result.ExtendedAttributes = expandSpAdapterAttributeList(val.(*schema.Set).List())
	}
	if val, ok := in["inherited"]; ok {
		result.Inherited = Bool(val.(bool))
	}
	return &result
}

func expandSaasAttributeMapping(in map[string]interface{}) *pf.SaasAttributeMapping {
	var result pf.SaasAttributeMapping
	if val, ok := in["field_name"]; ok {
		result.FieldName = String(val.(string))
	}
	if val, ok := in["saas_field_info"]; ok && len(val.([]interface{})) > 0 {
		result.SaasFieldInfo = expandSaasFieldConfiguration(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandSaasFieldConfiguration(in map[string]interface{}) *pf.SaasFieldConfiguration {
	var result pf.SaasFieldConfiguration
	if val, ok := in["create_only"]; ok {
		result.CreateOnly = Bool(val.(bool))
	}
	if val, ok := in["trim"]; ok {
		result.Trim = Bool(val.(bool))
	}
	if val, ok := in["character_case"]; ok {
		result.CharacterCase = String(val.(string))
	}
	if val, ok := in["parser"]; ok {
		result.Parser = String(val.(string))
	}
	if val, ok := in["masked"]; ok {
		result.Masked = Bool(val.(bool))
	}
	if val, ok := in["attribute_names"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.AttributeNames = &strs
	}
	if val, ok := in["default_value"]; ok {
		result.DefaultValue = String(val.(string))
	}
	if val, ok := in["expression"]; ok {
		result.Expression = String(val.(string))
	}
	return &result
}

func expandFieldEntry(in map[string]interface{}) *pf.FieldEntry {
	var result pf.FieldEntry
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["value"]; ok {
		result.Value = String(val.(string))
	}
	return &result
}

func expandAdditionalAllowedEntitiesConfiguration(in map[string]interface{}) *pf.AdditionalAllowedEntitiesConfiguration {
	var result pf.AdditionalAllowedEntitiesConfiguration
	if val, ok := in["additional_allowed_entities"]; ok {
		result.AdditionalAllowedEntities = expandEntityList(val.([]interface{}))
	}
	if val, ok := in["allow_additional_entities"]; ok {
		result.AllowAdditionalEntities = Bool(val.(bool))
	}
	if val, ok := in["allow_all_entities"]; ok {
		result.AllowAllEntities = Bool(val.(bool))
	}
	return &result
}

func expandArtifactResolverLocation(in map[string]interface{}) *pf.ArtifactResolverLocation {
	var result pf.ArtifactResolverLocation
	if val, ok := in["url"]; ok {
		result.Url = String(val.(string))
	}
	if val, ok := in["index"]; ok {
		result.Index = Int(val.(int))
	}
	return &result
}

func expandSpBrowserSsoAttributeContract(in map[string]interface{}) *pf.SpBrowserSsoAttributeContract {
	var result pf.SpBrowserSsoAttributeContract
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandSpBrowserSsoAttributeList(val.([]interface{}))
	}
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandSpBrowserSsoAttributeList(val.(*schema.Set).List())
	}
	return &result
}

func expandUsernamePasswordCredentials(in map[string]interface{}) *pf.UsernamePasswordCredentials {
	var result pf.UsernamePasswordCredentials
	if val, ok := in["username"]; ok {
		result.Username = String(val.(string))
	}
	if val, ok := in["password"]; ok {
		result.Password = String(val.(string))
	}
	if val, ok := in["encrypted_password"]; ok {
		result.EncryptedPassword = String(val.(string))
	}
	return &result
}

func expandConnectionCert(in map[string]interface{}) *pf.ConnectionCert {
	var result pf.ConnectionCert
	if val, ok := in["cert_view"]; ok && len(val.([]interface{})) > 0 {
		result.CertView = expandCertView(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["x509_file"]; ok && len(val.([]interface{})) > 0 {
		result.X509File = expandX509File(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["active_verification_cert"]; ok {
		result.ActiveVerificationCert = Bool(val.(bool))
	}
	if val, ok := in["primary_verification_cert"]; ok {
		result.PrimaryVerificationCert = Bool(val.(bool))
	}
	if val, ok := in["secondary_verification_cert"]; ok {
		result.SecondaryVerificationCert = Bool(val.(bool))
	}
	if val, ok := in["encryption_cert"]; ok {
		result.EncryptionCert = Bool(val.(bool))
	}
	return &result
}

func expandIdpBrowserSso(in map[string]interface{}) *pf.IdpBrowserSso {
	var result pf.IdpBrowserSso
	if val, ok := in["incoming_bindings"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.IncomingBindings = &strs
	}
	if val, ok := in["attribute_contract"]; ok && len(val.([]interface{})) > 0 {
		result.AttributeContract = expandIdpBrowserSsoAttributeContract(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["message_customizations"]; ok {
		result.MessageCustomizations = expandProtocolMessageCustomizationList(val.([]interface{}))
	}
	if val, ok := in["protocol"]; ok {
		result.Protocol = String(val.(string))
	}
	if val, ok := in["url_whitelist_entries"]; ok {
		result.UrlWhitelistEntries = expandUrlWhitelistEntryList(val.([]interface{}))
	}
	if val, ok := in["slo_service_endpoints"]; ok {
		result.SloServiceEndpoints = expandSloServiceEndpointList(val.([]interface{}))
	}
	if val, ok := in["sign_authn_requests"]; ok {
		result.SignAuthnRequests = Bool(val.(bool))
	}
	if val, ok := in["idp_identity_mapping"]; ok {
		result.IdpIdentityMapping = String(val.(string))
	}
	if val, ok := in["artifact"]; ok && len(val.([]interface{})) > 0 {
		result.Artifact = expandArtifactSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["adapter_mappings"]; ok {
		result.AdapterMappings = expandSpAdapterMappingList(val.([]interface{}))
	}
	if val, ok := in["authentication_policy_contract_mappings"]; ok {
		result.AuthenticationPolicyContractMappings = expandAuthenticationPolicyContractMappingList(val.([]interface{}))
	}
	if val, ok := in["sso_o_auth_mapping"]; ok && len(val.([]interface{})) > 0 {
		result.SsoOAuthMapping = expandSsoOAuthMapping(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["oauth_authentication_policy_contract_ref"]; ok {
		result.OauthAuthenticationPolicyContractRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["assertions_signed"]; ok {
		result.AssertionsSigned = Bool(val.(bool))
	}
	if val, ok := in["decryption_policy"]; ok && len(val.([]interface{})) > 0 {
		result.DecryptionPolicy = expandDecryptionPolicy(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["oidc_provider_settings"]; ok && len(val.([]interface{})) > 0 {
		result.OidcProviderSettings = expandOIDCProviderSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["enabled_profiles"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.EnabledProfiles = &strs
	}
	if val, ok := in["sso_service_endpoints"]; ok {
		result.SsoServiceEndpoints = expandIdpSsoServiceEndpointList(val.([]interface{}))
	}
	if val, ok := in["default_target_url"]; ok {
		result.DefaultTargetUrl = String(val.(string))
	}
	if val, ok := in["authn_context_mappings"]; ok {
		result.AuthnContextMappings = expandAuthnContextMappingList(val.([]interface{}))
	}
	return &result
}

func expandContactInfo(in map[string]interface{}) *pf.ContactInfo {
	var result pf.ContactInfo
	if val, ok := in["first_name"]; ok {
		result.FirstName = String(val.(string))
	}
	if val, ok := in["last_name"]; ok {
		result.LastName = String(val.(string))
	}
	if val, ok := in["phone"]; ok {
		result.Phone = String(val.(string))
	}
	if val, ok := in["company"]; ok {
		result.Company = String(val.(string))
	}
	if val, ok := in["email"]; ok {
		result.Email = String(val.(string))
	}
	return &result
}

func expandIdpBrowserSsoAttributeContract(in map[string]interface{}) *pf.IdpBrowserSsoAttributeContract {
	var result pf.IdpBrowserSsoAttributeContract
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandIdpBrowserSsoAttributeList(val.([]interface{}))
	}
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandIdpBrowserSsoAttributeList(val.([]interface{}))
	}
	return &result
}

func expandSpSsoServiceEndpoint(in map[string]interface{}) *pf.SpSsoServiceEndpoint {
	var result pf.SpSsoServiceEndpoint
	if val, ok := in["binding"]; ok {
		result.Binding = String(val.(string))
	}
	if val, ok := in["url"]; ok {
		result.Url = String(val.(string))
	}
	if val, ok := in["is_default"]; ok {
		result.IsDefault = Bool(val.(bool))
	}
	if val, ok := in["index"]; ok {
		result.Index = Int(val.(int))
	}
	return &result
}

func expandAttributeFulfillmentValue(in map[string]interface{}) *pf.AttributeFulfillmentValue {
	var result pf.AttributeFulfillmentValue
	if val, ok := in["source"]; ok {
		result.Source = expandSourceTypeIdKey(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["value"]; ok && val != "" {
		result.Value = String(val.(string))
	}
	return &result
}

func expandSpWsTrust(in map[string]interface{}) *pf.SpWsTrust {
	var result pf.SpWsTrust
	if val, ok := in["partner_service_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.PartnerServiceIds = &strs
	}
	if val, ok := in["generate_key"]; ok {
		result.GenerateKey = Bool(val.(bool))
	}
	if val, ok := in["token_processor_mappings"]; ok && len(val.([]interface{})) > 0 {
		result.TokenProcessorMappings = expandIdpTokenProcessorMappingList(val.([]interface{}))
	}
	if val, ok := in["abort_if_not_fulfilled_from_request"]; ok {
		result.AbortIfNotFulfilledFromRequest = Bool(val.(bool))
	}
	if val, ok := in["attribute_contract"]; ok && len(val.([]interface{})) > 0 {
		result.AttributeContract = expandSpWsTrustAttributeContract(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["request_contract_ref"]; ok && len(val.([]interface{})) > 0 {
		result.RequestContractRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["message_customizations"]; ok && len(val.([]interface{})) > 0 {
		result.MessageCustomizations = expandProtocolMessageCustomizationList(val.([]interface{}))
	}
	if val, ok := in["o_auth_assertion_profiles"]; ok {
		result.OAuthAssertionProfiles = Bool(val.(bool))
	}
	if val, ok := in["default_token_type"]; ok {
		result.DefaultTokenType = String(val.(string))
	}
	if val, ok := in["minutes_before"]; ok {
		result.MinutesBefore = Int(val.(int))
	}
	if val, ok := in["minutes_after"]; ok {
		result.MinutesAfter = Int(val.(int))
	}
	if val, ok := in["encrypt_saml2_assertion"]; ok {
		result.EncryptSaml2Assertion = Bool(val.(bool))
	}
	return &result
}

func expandSpAdapterMapping(in map[string]interface{}) *pf.SpAdapterMapping {
	result := pf.SpAdapterMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["sp_adapter_ref"]; ok {
		result.SpAdapterRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["restrict_virtual_entity_ids"]; ok {
		result.RestrictVirtualEntityIds = Bool(val.(bool))
	}
	if val, ok := in["restricted_virtual_entity_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualEntityIds = &strs
	}
	if val, ok := in["adapter_override_settings"]; ok && len(val.([]interface{})) > 0 {
		result.AdapterOverrideSettings = expandSpAdapter(val.([]interface{})[0].(map[string]interface{}))
	}
	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	return &result
}

func expandIdpWsTrustAttribute(in map[string]interface{}) *pf.IdpWsTrustAttribute {
	var result pf.IdpWsTrustAttribute
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["masked"]; ok {
		result.Masked = Bool(val.(bool))
	}
	return &result
}

func expandInboundBackChannelAuth(in map[string]interface{}) *pf.InboundBackChannelAuth {
	var result pf.InboundBackChannelAuth
	if val, ok := in["http_basic_credentials"]; ok && len(val.([]interface{})) > 0 {
		result.HttpBasicCredentials = expandUsernamePasswordCredentials(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["digital_signature"]; ok {
		result.DigitalSignature = Bool(val.(bool))
	}
	if val, ok := in["verification_subject_dn"]; ok && val.(string) != "" {
		result.VerificationSubjectDN = String(val.(string))
	}
	if val, ok := in["verification_issuer_dn"]; ok && val.(string) != "" {
		result.VerificationIssuerDN = String(val.(string))
	}
	if val, ok := in["certs"]; ok && len(val.([]interface{})) > 0 {
		result.Certs = expandConnectionCertList(val.([]interface{}))
	}
	if val, ok := in["require_ssl"]; ok {
		result.RequireSsl = Bool(val.(bool))
	}
	if val, ok := in["type"]; ok {
		result.Type = String(val.(string))
	}
	return &result
}

func expandSpWsTrustAttributeContract(in map[string]interface{}) *pf.SpWsTrustAttributeContract {
	var result pf.SpWsTrustAttributeContract
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandSpWsTrustAttributeList(val.([]interface{}))
	}
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandSpWsTrustAttributeList(val.([]interface{}))
	}
	return &result
}

func expandIdpSsoServiceEndpoint(in map[string]interface{}) *pf.IdpSsoServiceEndpoint {
	var result pf.IdpSsoServiceEndpoint
	if val, ok := in["binding"]; ok {
		result.Binding = String(val.(string))
	}
	if val, ok := in["url"]; ok {
		result.Url = String(val.(string))
	}
	return &result
}

func expandIdpOAuthAttributeContract(in map[string]interface{}) *pf.IdpOAuthAttributeContract {
	var result pf.IdpOAuthAttributeContract
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandIdpBrowserSsoAttributeList(val.([]interface{}))
	}
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandIdpBrowserSsoAttributeList(val.([]interface{}))
	}
	return &result
}

func expandOutboundBackChannelAuth(in map[string]interface{}) *pf.OutboundBackChannelAuth {
	var result pf.OutboundBackChannelAuth
	if val, ok := in["validate_partner_cert"]; ok {
		result.ValidatePartnerCert = Bool(val.(bool))
	}
	if val, ok := in["type"]; ok {
		result.Type = String(val.(string))
	}
	if val, ok := in["http_basic_credentials"]; ok && len(val.([]interface{})) > 0 {
		result.HttpBasicCredentials = expandUsernamePasswordCredentials(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["digital_signature"]; ok {
		result.DigitalSignature = Bool(val.(bool))
	}
	if val, ok := in["ssl_auth_key_pair_ref"]; ok && len(val.([]interface{})) > 0 {
		result.SslAuthKeyPairRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandSpWsTrustAttribute(in map[string]interface{}) *pf.SpWsTrustAttribute {
	var result pf.SpWsTrustAttribute
	if val, ok := in["namespace"]; ok {
		result.Namespace = String(val.(string))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	return &result
}

func expandSloServiceEndpoint(in map[string]interface{}) *pf.SloServiceEndpoint {
	var result pf.SloServiceEndpoint
	if val, ok := in["binding"]; ok {
		result.Binding = String(val.(string))
	}
	if val, ok := in["url"]; ok {
		result.Url = String(val.(string))
	}
	if val, ok := in["response_url"]; ok {
		result.ResponseUrl = String(val.(string))
	}
	return &result
}

func expandAssertionLifetime(in map[string]interface{}) *pf.AssertionLifetime {
	var result pf.AssertionLifetime
	if val, ok := in["minutes_before"]; ok {
		result.MinutesBefore = Int(val.(int))
	}
	if val, ok := in["minutes_after"]; ok {
		result.MinutesAfter = Int(val.(int))
	}
	return &result
}

func expandOIDCProviderSettings(in map[string]interface{}) *pf.OIDCProviderSettings {
	var result pf.OIDCProviderSettings
	if val, ok := in["authorization_endpoint"]; ok {
		result.AuthorizationEndpoint = String(val.(string))
	}
	if val, ok := in["login_type"]; ok {
		result.LoginType = String(val.(string))
	}
	if val, ok := in["authentication_scheme"]; ok {
		result.AuthenticationScheme = String(val.(string))
	}
	if val, ok := in["authentication_signing_algorithm"]; ok {
		result.AuthenticationSigningAlgorithm = String(val.(string))
	}
	if val, ok := in["request_signing_algorithm"]; ok {
		result.RequestSigningAlgorithm = String(val.(string))
	}
	if val, ok := in["request_parameters"]; ok {
		result.RequestParameters = expandOIDCRequestParameterList(val.([]interface{}))
	}
	if val, ok := in["scopes"]; ok {
		result.Scopes = String(val.(string))
	}
	if val, ok := in["user_info_endpoint"]; ok {
		result.UserInfoEndpoint = String(val.(string))
	}
	if val, ok := in["jwks_url"]; ok {
		result.JwksURL = String(val.(string))
	}
	if val, ok := in["token_endpoint"]; ok {
		result.TokenEndpoint = String(val.(string))
	}
	return &result
}

func expandIdpBrowserSsoAttribute(in map[string]interface{}) *pf.IdpBrowserSsoAttribute {
	var result pf.IdpBrowserSsoAttribute
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["masked"]; ok {
		result.Masked = Bool(val.(bool))
	}
	return &result
}

func expandX509File(in map[string]interface{}) *pf.X509File {
	var result pf.X509File
	if val, ok := in["id"]; ok && val.(string) != "" {
		result.Id = String(val.(string))
	}
	if val, ok := in["file_data"]; ok {
		result.FileData = String(val.(string))
	}
	if val, ok := in["crypto_provider"]; ok && val.(string) != "" {
		result.CryptoProvider = String(val.(string))
	}
	return &result
}

func expandIdpTokenProcessorMapping(in map[string]interface{}) *pf.IdpTokenProcessorMapping {
	result := pf.IdpTokenProcessorMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["idp_token_processor_ref"]; ok {
		result.IdpTokenProcessorRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["restricted_virtual_entity_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualEntityIds = &strs
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandIdpAdapterContractMapping(in map[string]interface{}) *pf.IdpAdapterContractMapping {
	result := pf.IdpAdapterContractMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["inherited"]; ok {
		result.Inherited = Bool(val.(bool))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	return &result
}

func expandGroupMembershipDetection(in map[string]interface{}) *pf.GroupMembershipDetection {
	var result pf.GroupMembershipDetection
	if val, ok := in["group_member_attribute_name"]; ok {
		result.GroupMemberAttributeName = String(val.(string))
	}
	if val, ok := in["member_of_group_attribute_name"]; ok {
		result.MemberOfGroupAttributeName = String(val.(string))
	}
	return &result
}

func expandEntity(in map[string]interface{}) *pf.Entity {
	var result pf.Entity
	if val, ok := in["entity_id"]; ok {
		result.EntityId = String(val.(string))
	}
	if val, ok := in["entity_description"]; ok {
		result.EntityDescription = String(val.(string))
	}
	return &result
}

func expandSpAdapter(in map[string]interface{}) *pf.SpAdapter {
	var result pf.SpAdapter
	if val, ok := in["configuration"]; ok {
		result.Configuration = expandPluginConfiguration(val.([]interface{}))
	}
	if val, ok := in["attribute_contract"]; ok {
		result.AttributeContract = expandSpAdapterAttributeContract(val.(map[string]interface{}))
	}
	if val, ok := in["target_application_info"]; ok {
		result.TargetApplicationInfo = expandSpAdapterTargetApplicationInfo(val.(map[string]interface{}))
	}
	if val, ok := in["id"]; ok {
		result.Id = String(val.(string))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["plugin_descriptor_ref"]; ok {
		result.PluginDescriptorRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["parent_ref"]; ok {
		result.ParentRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandIdpWsTrust(in map[string]interface{}) *pf.IdpWsTrust {
	var result pf.IdpWsTrust
	if val, ok := in["attribute_contract"]; ok && len(val.([]interface{})) > 0 {
		result.AttributeContract = expandIdpWsTrustAttributeContract(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["generate_local_token"]; ok {
		result.GenerateLocalToken = Bool(val.(bool))
	}
	if val, ok := in["token_generator_mappings"]; ok {
		result.TokenGeneratorMappings = expandSpTokenGeneratorMappingList(val.([]interface{}))
	}
	return &result
}

func expandConfigTable(in map[string]interface{}) *pf.ConfigTable {
	var result pf.ConfigTable
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["rows"]; ok {
		result.Rows = expandConfigRowList(val.([]interface{}))
	}
	if val, ok := in["inherited"]; ok {
		result.Inherited = Bool(val.(bool))
	}
	return &result
}

func expandOIDCRequestParameter(in map[string]interface{}) *pf.OIDCRequestParameter {
	var result pf.OIDCRequestParameter
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["value"]; ok {
		result.Value = String(val.(string))
	}
	if val, ok := in["application_endpoint_override"]; ok {
		result.ApplicationEndpointOverride = Bool(val.(bool))
	}
	return &result
}

func expandSigningSettings(in map[string]interface{}) *pf.SigningSettings {
	var result pf.SigningSettings
	if val, ok := in["signing_key_pair_ref"]; ok {
		result.SigningKeyPairRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["algorithm"]; ok {
		result.Algorithm = String(val.(string))
	}
	if val, ok := in["include_cert_in_signature"]; ok {
		result.IncludeCertInSignature = Bool(val.(bool))
	}
	if val, ok := in["include_raw_key_in_signature"]; ok {
		result.IncludeRawKeyInSignature = Bool(val.(bool))
	}
	return &result
}

//func expandConfigField(in map[string]interface{}) *pf.ConfigField {
//	var result pf.ConfigField
//	if val, ok := in["name"]; ok {
//		result.Name = String(val.(string))
//	}
//	if val, ok := in["value"]; ok {
//		result.Value = String(val.(string))
//	}
//	if val, ok := in["encrypted_value"]; ok {
//		result.EncryptedValue = String(val.(string))
//	}
//	if val, ok := in["inherited"]; ok {
//		result.Inherited = Bool(val.(bool))
//	}
//	return &result
//}

func expandIdpOAuthGrantAttributeMapping(in map[string]interface{}) *pf.IdpOAuthGrantAttributeMapping {
	var result pf.IdpOAuthGrantAttributeMapping
	if val, ok := in["idp_o_auth_attribute_contract"]; ok && len(val.([]interface{})) > 0 {
		result.IdpOAuthAttributeContract = expandIdpOAuthAttributeContract(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["access_token_manager_mappings"]; ok {
		result.AccessTokenManagerMappings = expandAccessTokenManagerMappingList(val.([]interface{}))
	}
	return &result
}

func expandSpAttributeQuery(in map[string]interface{}) *pf.SpAttributeQuery {
	result := pf.SpAttributeQuery{
		AttributeSources: &[]*pf.AttributeSource{},
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 && val.([]interface{})[0] != nil {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["policy"]; ok && len(val.([]interface{})) > 0 {
		result.Policy = expandSpAttributeQueryPolicy(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["attributes"]; ok && len(val.([]interface{})) > 0 {
		strs := expandStringList(val.([]interface{}))
		result.Attributes = &strs
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	return &result
}

func expandSsoOAuthMapping(in map[string]interface{}) *pf.SsoOAuthMapping {
	result := pf.SsoOAuthMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	return &result
}

func expandIdpWsTrustAttributeContract(in map[string]interface{}) *pf.IdpWsTrustAttributeContract {
	var result pf.IdpWsTrustAttributeContract
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandIdpWsTrustAttributeList(val.([]interface{}))
	}
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandIdpWsTrustAttributeList(val.([]interface{}))
	}
	return &result
}

func expandChannelSource(in map[string]interface{}) *pf.ChannelSource {
	var result pf.ChannelSource
	if val, ok := in["data_source"]; ok && len(val.([]interface{})) > 0 {
		result.DataSource = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["group_membership_detection"]; ok && len(val.([]interface{})) > 0 {
		result.GroupMembershipDetection = expandGroupMembershipDetection(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["base_dn"]; ok {
		result.BaseDn = String(val.(string))
	}
	if val, ok := in["group_source_location"]; ok && len(val.([]interface{})) > 0 {
		result.GroupSourceLocation = expandChannelSourceLocation(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["guid_attribute_name"]; ok {
		result.GuidAttributeName = String(val.(string))
	}
	if val, ok := in["guid_binary"]; ok {
		result.GuidBinary = Bool(val.(bool))
	}
	if val, ok := in["change_detection_settings"]; ok && len(val.([]interface{})) > 0 {
		result.ChangeDetectionSettings = expandChangeDetectionSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["account_management_settings"]; ok && len(val.([]interface{})) > 0 {
		result.AccountManagementSettings = expandAccountManagementSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["user_source_location"]; ok && len(val.([]interface{})) > 0 {
		result.UserSourceLocation = expandChannelSourceLocation(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandAttributeQueryNameMapping(in map[string]interface{}) *pf.AttributeQueryNameMapping {
	var result pf.AttributeQueryNameMapping
	if val, ok := in["local_name"]; ok {
		result.LocalName = String(val.(string))
	}
	if val, ok := in["remote_name"]; ok {
		result.RemoteName = String(val.(string))
	}
	return &result
}

func expandSpBrowserSsoAttribute(in map[string]interface{}) *pf.SpBrowserSsoAttribute {
	var result pf.SpBrowserSsoAttribute
	if val, ok := in["name_format"]; ok {
		result.NameFormat = String(val.(string))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	return &result
}

func expandDecryptionPolicy(in map[string]interface{}) *pf.DecryptionPolicy {
	var result pf.DecryptionPolicy
	if val, ok := in["assertion_encrypted"]; ok {
		result.AssertionEncrypted = Bool(val.(bool))
	}
	if val, ok := in["attributes_encrypted"]; ok {
		result.AttributesEncrypted = Bool(val.(bool))
	}
	if val, ok := in["subject_name_id_encrypted"]; ok {
		result.SubjectNameIdEncrypted = Bool(val.(bool))
	}
	if val, ok := in["slo_encrypt_subject_name_id"]; ok {
		result.SloEncryptSubjectNameID = Bool(val.(bool))
	}
	if val, ok := in["slo_subject_name_id_encrypted"]; ok {
		result.SloSubjectNameIDEncrypted = Bool(val.(bool))
	}
	return &result
}

func expandParameterValues(in map[string]interface{}) *pf.ParameterValues {
	var result pf.ParameterValues
	if val, ok := in["values"]; ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.Values = &strs
	}
	return &result
}

func expandSpAttributeQueryPolicy(in map[string]interface{}) *pf.SpAttributeQueryPolicy {
	var result pf.SpAttributeQueryPolicy
	if val, ok := in["sign_assertion"]; ok {
		result.SignAssertion = Bool(val.(bool))
	}
	if val, ok := in["encrypt_assertion"]; ok {
		result.EncryptAssertion = Bool(val.(bool))
	}
	if val, ok := in["require_signed_attribute_query"]; ok {
		result.RequireSignedAttributeQuery = Bool(val.(bool))
	}
	if val, ok := in["require_encrypted_name_id"]; ok {
		result.RequireEncryptedNameId = Bool(val.(bool))
	}
	if val, ok := in["sign_response"]; ok {
		result.SignResponse = Bool(val.(bool))
	}
	return &result
}

func expandJdbcAttributeSource(in map[string]interface{}) *pf.AttributeSource {
	src := &pf.AttributeSource{Type: String("JDBC")}
	iac := &pf.JdbcAttributeSource{Type: String("JDBC")}
	if v, ok := in["data_store_ref"]; ok && len(v.([]interface{})) > 0 {
		iac.DataStoreRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		src.DataStoreRef = iac.DataStoreRef
	}
	if v, ok := in["schema"]; ok {
		iac.Schema = String(v.(string))
	}
	if v, ok := in["id"]; ok {
		iac.Id = String(v.(string))
		src.Id = iac.Id
	}
	if v, ok := in["table"]; ok {
		iac.Table = String(v.(string))
	}
	if v, ok := in["description"]; ok {
		iac.Description = String(v.(string))
		src.Description = iac.Description
	}
	if v, ok := in["filter"]; ok {
		iac.Filter = String(v.(string))
	}
	if v, ok := in["attribute_contract_fulfillment"]; ok {
		iac.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(v.(*schema.Set).List())
		src.AttributeContractFulfillment = iac.AttributeContractFulfillment
	}
	src.JdbcAttributeSource = *iac
	return src
}

func expandIssuanceCriteria(in map[string]interface{}) *pf.IssuanceCriteria {
	var result pf.IssuanceCriteria
	if val, ok := in["conditional_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.ConditionalCriteria = expandConditionalIssuanceCriteriaEntryList(val.([]interface{}))
	}
	if val, ok := in["expression_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.ExpressionCriteria = expandExpressionIssuanceCriteriaEntryList(val.([]interface{}))
	}
	return &result
}

func expandSpAdapterTargetApplicationInfo(in map[string]interface{}) *pf.SpAdapterTargetApplicationInfo {
	var result pf.SpAdapterTargetApplicationInfo
	if val, ok := in["application_name"]; ok {
		result.ApplicationName = String(val.(string))
	}
	if val, ok := in["application_icon_url"]; ok {
		result.ApplicationIconUrl = String(val.(string))
	}
	if val, ok := in["inherited"]; ok {
		result.Inherited = Bool(val.(bool))
	}
	return &result
}

func expandExpressionIssuanceCriteriaEntry(in map[string]interface{}) *pf.ExpressionIssuanceCriteriaEntry {
	var result pf.ExpressionIssuanceCriteriaEntry
	if val, ok := in["error_result"]; ok {
		result.ErrorResult = String(val.(string))
	}
	if val, ok := in["expression"]; ok {
		result.Expression = String(val.(string))
	}
	return &result
}

func expandUrlWhitelistEntry(in map[string]interface{}) *pf.UrlWhitelistEntry {
	var result pf.UrlWhitelistEntry
	if val, ok := in["allow_query_and_fragment"]; ok {
		result.AllowQueryAndFragment = Bool(val.(bool))
	}
	if val, ok := in["require_https"]; ok {
		result.RequireHttps = Bool(val.(bool))
	}
	if val, ok := in["valid_domain"]; ok {
		result.ValidDomain = String(val.(string))
	}
	if val, ok := in["valid_path"]; ok {
		result.ValidPath = String(val.(string))
	}
	return &result
}

func expandResourceLink(in map[string]interface{}) *pf.ResourceLink {
	var result pf.ResourceLink
	if val, ok := in["id"]; ok {
		result.Id = String(val.(string))
	}
	if val, ok := in["location"]; ok && val.(string) != "" {
		result.Location = String(val.(string))
	}
	return &result
}

func expandSpTokenGeneratorMapping(in map[string]interface{}) *pf.SpTokenGeneratorMapping {
	result := pf.SpTokenGeneratorMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["sp_token_generator_ref"]; ok {
		result.SpTokenGeneratorRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["restricted_virtual_entity_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualEntityIds = &strs
	}
	if val, ok := in["default_mapping"]; ok {
		result.DefaultMapping = Bool(val.(bool))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	return &result
}

func expandAccessTokenManagerMapping(in map[string]interface{}) *pf.AccessTokenManagerMapping {
	result := pf.AccessTokenManagerMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["access_token_manager_ref"]; ok {
		result.AccessTokenManagerRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandSchemaAttribute(in map[string]interface{}) *pf.SchemaAttribute {
	var result pf.SchemaAttribute
	if val, ok := in["types"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.Types = &strs
	}
	if val, ok := in["sub_attributes"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.SubAttributes = &strs
	}
	if val, ok := in["multi_valued"]; ok {
		result.MultiValued = Bool(val.(bool))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	return &result
}

func expandIdpAdapterAssertionMapping(in map[string]interface{}) *pf.IdpAdapterAssertionMapping {
	result := pf.IdpAdapterAssertionMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["abort_sso_transaction_as_fail_safe"]; ok {
		result.AbortSsoTransactionAsFailSafe = Bool(val.(bool))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["idp_adapter_ref"]; ok && len(val.([]interface{})) > 0 {
		result.IdpAdapterRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["restrict_virtual_entity_ids"]; ok {
		result.RestrictVirtualEntityIds = Bool(val.(bool))
	}
	if val, ok := in["restricted_virtual_entity_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualEntityIds = &strs
	}
	if val, ok := in["adapter_override_settings"]; ok && len(val.([]interface{})) > 0 {
		result.AdapterOverrideSettings = expandIdpAdapter(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandAuthenticationPolicyContractAssertionMapping(in map[string]interface{}) *pf.AuthenticationPolicyContractAssertionMapping {
	result := pf.AuthenticationPolicyContractAssertionMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["restrict_virtual_entity_ids"]; ok {
		result.RestrictVirtualEntityIds = Bool(val.(bool))
	}
	if val, ok := in["restricted_virtual_entity_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualEntityIds = &strs
	}
	if val, ok := in["abort_sso_transaction_as_fail_safe"]; ok {
		result.AbortSsoTransactionAsFailSafe = Bool(val.(bool))
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["authentication_policy_contract_ref"]; ok && len(val.([]interface{})) > 0 {
		result.AuthenticationPolicyContractRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandOutboundProvision(in map[string]interface{}) *pf.OutboundProvision {
	result := pf.OutboundProvision{
		TargetSettings: &[]*pf.ConfigField{},
	}
	if val, ok := in["type"]; ok {
		result.Type = String(val.(string))
	}
	if val, ok := in["target_settings"]; ok {
		fields := expandConfigFields(val.(*schema.Set).List())
		*result.TargetSettings = append(*result.TargetSettings, *fields...)
	}
	if val, ok := in["sensitive_target_settings"]; ok {
		fields := expandSensitiveConfigFields(val.(*schema.Set).List())
		*result.TargetSettings = append(*result.TargetSettings, *fields...)
	}
	if val, ok := in["custom_schema"]; ok && len(val.([]interface{})) > 0 {
		result.CustomSchema = expandSchema(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["channels"]; ok {
		result.Channels = expandChannelList(val.([]interface{}))
	}
	return &result
}

func expandAuthnContextMapping(in map[string]interface{}) *pf.AuthnContextMapping {
	var result pf.AuthnContextMapping
	if val, ok := in["remote"]; ok {
		result.Remote = String(val.(string))
	}
	if val, ok := in["local"]; ok {
		result.Local = String(val.(string))
	}
	return &result
}

func expandIdpAttributeQueryPolicy(in map[string]interface{}) *pf.IdpAttributeQueryPolicy {
	var result pf.IdpAttributeQueryPolicy
	if val, ok := in["mask_attribute_values"]; ok {
		result.MaskAttributeValues = Bool(val.(bool))
	}
	if val, ok := in["require_signed_response"]; ok {
		result.RequireSignedResponse = Bool(val.(bool))
	}
	if val, ok := in["require_signed_assertion"]; ok {
		result.RequireSignedAssertion = Bool(val.(bool))
	}
	if val, ok := in["require_encrypted_assertion"]; ok {
		result.RequireEncryptedAssertion = Bool(val.(bool))
	}
	if val, ok := in["sign_attribute_query"]; ok {
		result.SignAttributeQuery = Bool(val.(bool))
	}
	if val, ok := in["encrypt_name_id"]; ok {
		result.EncryptNameId = Bool(val.(bool))
	}
	return &result
}

func expandLdapAttributeSource(in map[string]interface{}) *pf.AttributeSource {
	src := &pf.AttributeSource{Type: String("LDAP")}
	iac := &pf.LdapAttributeSource{Type: String("LDAP")}
	if v, ok := in["data_store_ref"]; ok && len(v.([]interface{})) > 0 {
		iac.DataStoreRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		src.DataStoreRef = iac.DataStoreRef
	}
	if v, ok := in["base_dn"]; ok {
		iac.BaseDn = String(v.(string))
	}
	if v, ok := in["id"]; ok {
		iac.Id = String(v.(string))
		src.Id = iac.Id
	}
	if v, ok := in["search_scope"]; ok {
		iac.SearchScope = String(v.(string))
	}
	if v, ok := in["description"]; ok {
		iac.Description = String(v.(string))
		src.Description = iac.Description
	}
	if v, ok := in["search_filter"]; ok {
		iac.SearchFilter = String(v.(string))
	}
	if v, ok := in["attribute_contract_fulfillment"]; ok {
		iac.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(v.(*schema.Set).List())
		src.AttributeContractFulfillment = iac.AttributeContractFulfillment
	}
	if v, ok := in["binary_attribute_settings"]; ok {
		ca := map[string]*pf.BinaryLdapAttributeSettings{}
		for key, val := range v.(map[string]interface{}) {
			ca[key] = &pf.BinaryLdapAttributeSettings{BinaryEncoding: String(val.(string))}
		}
		iac.BinaryAttributeSettings = ca
	}
	if v, ok := in["member_of_nested_group"]; ok {
		iac.MemberOfNestedGroup = Bool(v.(bool))
	}
	src.LdapAttributeSource = *iac
	return src
}

func expandChangeDetectionSettings(in map[string]interface{}) *pf.ChangeDetectionSettings {
	var result pf.ChangeDetectionSettings
	if val, ok := in["usn_attribute_name"]; ok {
		result.UsnAttributeName = String(val.(string))
	}
	if val, ok := in["time_stamp_attribute_name"]; ok {
		result.TimeStampAttributeName = String(val.(string))
	}
	if val, ok := in["user_object_class"]; ok {
		result.UserObjectClass = String(val.(string))
	}
	if val, ok := in["group_object_class"]; ok {
		result.GroupObjectClass = String(val.(string))
	}
	if val, ok := in["changed_users_algorithm"]; ok {
		result.ChangedUsersAlgorithm = String(val.(string))
	}
	return &result
}

func expandIdpAdapterAttributeContract(in map[string]interface{}) *pf.IdpAdapterAttributeContract {
	var result pf.IdpAdapterAttributeContract
	if val, ok := in["core_attributes"]; ok {
		result.CoreAttributes = expandIdpAdapterAttributeList(val.(*schema.Set).List())
	}
	if val, ok := in["extended_attributes"]; ok {
		result.ExtendedAttributes = expandIdpAdapterAttributeList(val.(*schema.Set).List())
	}
	if val, ok := in["mask_ognl_values"]; ok {
		result.MaskOgnlValues = Bool(val.(bool))
	}
	if val, ok := in["inherited"]; ok {
		result.Inherited = Bool(val.(bool))
	}
	return &result
}

func expandSchema(in map[string]interface{}) *pf.Schema {
	var result pf.Schema
	if val, ok := in["namespace"]; ok {
		result.Namespace = String(val.(string))
	}
	if val, ok := in["attributes"]; ok {
		result.Attributes = expandSchemaAttributeList(val.([]interface{}))
	}
	return &result
}

func expandCertView(in map[string]interface{}) *pf.CertView {
	var result pf.CertView
	if val, ok := in["version"]; ok {
		result.Version = Int(val.(int))
	}
	if val, ok := in["crypto_provider"]; ok && val.(string) != "" {
		result.CryptoProvider = String(val.(string))
	}
	if val, ok := in["id"]; ok {
		result.Id = String(val.(string))
	}
	if val, ok := in["serial_number"]; ok {
		result.SerialNumber = String(val.(string))
	}
	if val, ok := in["issuer_dn"]; ok {
		result.IssuerDN = String(val.(string))
	}
	if val, ok := in["subject_dn"]; ok {
		result.SubjectDN = String(val.(string))
	}
	if val, ok := in["signature_algorithm"]; ok {
		result.SignatureAlgorithm = String(val.(string))
	}
	if val, ok := in["sha1_fingerprint"]; ok {
		result.Sha1Fingerprint = String(val.(string))
	}
	if val, ok := in["expires"]; ok {
		result.Expires = String(val.(string))
	}
	if val, ok := in["key_size"]; ok {
		result.KeySize = Int(val.(int))
	}
	if val, ok := in["sha256_fingerprint"]; ok {
		result.Sha256Fingerprint = String(val.(string))
	}
	if val, ok := in["status"]; ok {
		result.Status = String(val.(string))
	}
	if val, ok := in["subject_alternative_names"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.SubjectAlternativeNames = &strs
	}
	if val, ok := in["valid_from"]; ok {
		result.ValidFrom = String(val.(string))
	}
	if val, ok := in["key_algorithm"]; ok {
		result.KeyAlgorithm = String(val.(string))
	}
	return &result
}

func expandSpBrowserSso(in map[string]interface{}) *pf.SpBrowserSso {
	var result pf.SpBrowserSso
	if val, ok := in["enabled_profiles"]; ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.EnabledProfiles = &strs
	}
	if val, ok := in["default_target_url"]; ok && val.(string) != "" {
		result.DefaultTargetUrl = String(val.(string))
	}
	if val, ok := in["sign_assertions"]; ok {
		result.SignAssertions = Bool(val.(bool))
	}
	if val, ok := in["attribute_contract"]; ok && len(val.([]interface{})) > 0 {
		result.AttributeContract = expandSpBrowserSsoAttributeContract(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["authentication_policy_contract_assertion_mappings"]; ok {
		result.AuthenticationPolicyContractAssertionMappings = expandAuthenticationPolicyContractAssertionMappingList(val.([]interface{}))
	}
	if val, ok := in["assertion_lifetime"]; ok && len(val.([]interface{})) > 0 {
		result.AssertionLifetime = expandAssertionLifetime(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["ws_trust_version"]; ok && val.(string) != "" {
		result.WsTrustVersion = String(val.(string))
	}
	if val, ok := in["sp_ws_fed_identity_mapping"]; ok && val.(string) != "" {
		result.SpWsFedIdentityMapping = String(val.(string))
	}
	if val, ok := in["adapter_mappings"]; ok {
		result.AdapterMappings = expandIdpAdapterAssertionMappingList(val.([]interface{}))
	}
	if val, ok := in["message_customizations"]; ok {
		result.MessageCustomizations = expandProtocolMessageCustomizationList(val.([]interface{}))
	}
	if val, ok := in["url_whitelist_entries"]; ok {
		result.UrlWhitelistEntries = expandUrlWhitelistEntryList(val.([]interface{}))
	}
	if val, ok := in["artifact"]; ok && len(val.([]interface{})) > 0 {
		result.Artifact = expandArtifactSettings(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["sso_service_endpoints"]; ok {
		result.SsoServiceEndpoints = expandSpSsoServiceEndpointList(val.([]interface{}))
	}
	if val, ok := in["sp_saml_identity_mapping"]; ok && val.(string) != "" {
		result.SpSamlIdentityMapping = String(val.(string))
	}
	if val, ok := in["sign_response_as_required"]; ok {
		result.SignResponseAsRequired = Bool(val.(bool))
	}
	if val, ok := in["require_signed_authn_requests"]; ok {
		result.RequireSignedAuthnRequests = Bool(val.(bool))
	}
	if val, ok := in["encryption_policy"]; ok && len(val.([]interface{})) > 0 {
		result.EncryptionPolicy = expandEncryptionPolicy(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["incoming_bindings"]; ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.IncomingBindings = &strs
	}
	if val, ok := in["ws_fed_token_type"]; ok && val.(string) != "" {
		result.WsFedTokenType = String(val.(string))
	}
	if val, ok := in["slo_service_endpoints"]; ok {
		result.SloServiceEndpoints = expandSloServiceEndpointList(val.([]interface{}))
	}
	if val, ok := in["protocol"]; ok {
		result.Protocol = String(val.(string))
	}
	return &result
}

func expandIdpAdapter(in map[string]interface{}) *pf.IdpAdapter {
	var result pf.IdpAdapter
	if val, ok := in["attribute_contract"]; ok {
		result.AttributeContract = expandIdpAdapterAttributeContract(val.(map[string]interface{}))
	}
	if val, ok := in["id"]; ok {
		result.Id = String(val.(string))
	}
	if val, ok := in["name"]; ok {
		result.Name = String(val.(string))
	}
	if val, ok := in["plugin_descriptor_ref"]; ok {
		result.PluginDescriptorRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["parent_ref"]; ok {
		result.ParentRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["configuration"]; ok {
		result.Configuration = expandPluginConfiguration(val.([]interface{}))
	}
	if val, ok := in["authn_ctx_class_ref"]; ok {
		result.AuthnCtxClassRef = String(val.(string))
	}
	if val, ok := in["attribute_mapping"]; ok {
		result.AttributeMapping = expandIdpAdapterContractMapping(val.(map[string]interface{}))
	}
	return &result
}

func expandConfigRow(in map[string]interface{}) *pf.ConfigRow {
	var result pf.ConfigRow
	if val, ok := in["fields"]; ok {
		result.Fields = expandConfigFields(val.(*schema.Set).List())
	}
	if val, ok := in["sensitive_fields"]; ok {
		fields := expandSensitiveConfigFields(val.(*schema.Set).List())
		*result.Fields = append(*result.Fields, *fields...)
	}
	//Requires https://github.com/hashicorp/terraform-plugin-sdk/issues/261
	if val, ok := in["default_row"]; ok && val.(bool) {
		//set only if true
		result.DefaultRow = Bool(val.(bool))
	}
	return &result
}

func expandEncryptionPolicy(in map[string]interface{}) *pf.EncryptionPolicy {
	var result pf.EncryptionPolicy
	if val, ok := in["slo_subject_name_id_encrypted"]; ok {
		result.SloSubjectNameIDEncrypted = Bool(val.(bool))
	}
	if val, ok := in["encrypt_assertion"]; ok {
		result.EncryptAssertion = Bool(val.(bool))
	}
	if val, ok := in["encrypted_attributes"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.EncryptedAttributes = &strs
	}
	if val, ok := in["encrypt_slo_subject_name_id"]; ok {
		result.EncryptSloSubjectNameId = Bool(val.(bool))
	}
	return &result
}

func expandProtocolMessageCustomization(in map[string]interface{}) *pf.ProtocolMessageCustomization {
	var result pf.ProtocolMessageCustomization
	if val, ok := in["message_expression"]; ok {
		result.MessageExpression = String(val.(string))
	}
	if val, ok := in["context_name"]; ok {
		result.ContextName = String(val.(string))
	}
	return &result
}

func expandAccountManagementSettings(in map[string]interface{}) *pf.AccountManagementSettings {
	var result pf.AccountManagementSettings
	if val, ok := in["account_status_attribute_name"]; ok {
		result.AccountStatusAttributeName = String(val.(string))
	}
	if val, ok := in["account_status_algorithm"]; ok {
		result.AccountStatusAlgorithm = String(val.(string))
	}
	if val, ok := in["flag_comparison_value"]; ok {
		result.FlagComparisonValue = String(val.(string))
	}
	if val, ok := in["flag_comparison_status"]; ok {
		result.FlagComparisonStatus = Bool(val.(bool))
	}
	if val, ok := in["default_status"]; ok {
		result.DefaultStatus = Bool(val.(bool))
	}
	return &result
}

func expandAuthenticationPolicyContractMapping(in map[string]interface{}) *pf.AuthenticationPolicyContractMapping {
	result := pf.AuthenticationPolicyContractMapping{AttributeSources: &[]*pf.AttributeSource{}}
	if val, ok := in["authentication_policy_contract_ref"]; ok {
		result.AuthenticationPolicyContractRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["restrict_virtual_server_ids"]; ok {
		result.RestrictVirtualServerIds = Bool(val.(bool))
	}
	if val, ok := in["restricted_virtual_server_ids"]; ok {
		strs := expandStringList(val.([]interface{}))
		result.RestrictedVirtualServerIds = &strs
	}

	if v, ok := in["ldap_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandLdapAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["jdbc_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandJdbcAttributeSourceList(v.([]interface{}))...)
	}
	if v, ok := in["custom_attribute_source"]; ok && len(v.([]interface{})) > 0 {
		*result.AttributeSources = append(*result.AttributeSources, *expandCustomAttributeSourceList(v.([]interface{}))...)
	}
	if val, ok := in["attribute_contract_fulfillment"]; ok {
		result.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(val.(*schema.Set).List())
	}
	if val, ok := in["issuance_criteria"]; ok && len(val.([]interface{})) > 0 {
		result.IssuanceCriteria = expandIssuanceCriteria(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandOIDCClientCredentials(in map[string]interface{}) *pf.OIDCClientCredentials {
	var result pf.OIDCClientCredentials
	if val, ok := in["client_id"]; ok {
		result.ClientId = String(val.(string))
	}
	if val, ok := in["client_secret"]; ok {
		result.ClientSecret = String(val.(string))
	}
	if val, ok := in["encrypted_secret"]; ok {
		result.EncryptedSecret = String(val.(string))
	}
	return &result
}

func expandIdpAttributeQuery(in map[string]interface{}) *pf.IdpAttributeQuery {
	var result pf.IdpAttributeQuery
	if val, ok := in["name_mappings"]; ok {
		result.NameMappings = expandAttributeQueryNameMappingList(val.([]interface{}))
	}
	if val, ok := in["policy"]; ok && len(val.([]interface{})) > 0 {
		result.Policy = expandIdpAttributeQueryPolicy(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["url"]; ok {
		result.Url = String(val.(string))
	}
	return &result
}

func expandCustomAttributeSource(in map[string]interface{}) *pf.AttributeSource {
	src := &pf.AttributeSource{Type: String("CUSTOM")}
	iac := &pf.CustomAttributeSource{Type: String("CUSTOM")}
	if v, ok := in["data_store_ref"]; ok && len(v.([]interface{})) > 0 {
		iac.DataStoreRef = expandResourceLink(v.([]interface{})[0].(map[string]interface{}))
		src.DataStoreRef = iac.DataStoreRef
	}
	if v, ok := in["id"]; ok {
		iac.Id = String(v.(string))
		src.Id = iac.Id
	}
	if v, ok := in["description"]; ok {
		iac.Description = String(v.(string))
		src.Description = iac.Description
	}
	if v, ok := in["filter_fields"]; ok {
		iac.FilterFields = expandFieldEntryList(v.([]interface{}))
	}
	if v, ok := in["attribute_contract_fulfillment"]; ok {
		iac.AttributeContractFulfillment = expandMapOfAttributeFulfillmentValue(v.(*schema.Set).List())
		src.AttributeContractFulfillment = iac.AttributeContractFulfillment
	}
	src.CustomAttributeSource = *iac
	return src
}

func expandArtifactSettings(in map[string]interface{}) *pf.ArtifactSettings {
	var result pf.ArtifactSettings
	if val, ok := in["resolver_locations"]; ok {
		result.ResolverLocations = expandArtifactResolverLocationList(val.([]interface{}))
	}
	if val, ok := in["source_id"]; ok {
		result.SourceId = String(val.(string))
	}
	if val, ok := in["lifetime"]; ok {
		result.Lifetime = Int(val.(int))
	}
	return &result
}

func expandConditionalIssuanceCriteriaEntry(in map[string]interface{}) *pf.ConditionalIssuanceCriteriaEntry {
	var result pf.ConditionalIssuanceCriteriaEntry
	if val, ok := in["source"]; ok && len(val.([]interface{})) > 0 {
		result.Source = expandSourceTypeIdKey(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["attribute_name"]; ok {
		result.AttributeName = String(val.(string))
	}
	if val, ok := in["condition"]; ok {
		result.Condition = String(val.(string))
	}
	if val, ok := in["value"]; ok {
		result.Value = String(val.(string))
	}
	if val, ok := in["error_result"]; ok {
		result.ErrorResult = String(val.(string))
	}
	return &result
}

func expandConnectionMetadataUrl(in map[string]interface{}) *pf.ConnectionMetadataUrl {
	var result pf.ConnectionMetadataUrl
	if val, ok := in["metadata_url_ref"]; ok {
		result.MetadataUrlRef = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["enable_auto_metadata_update"]; ok {
		result.EnableAutoMetadataUpdate = Bool(val.(bool))
	}
	return &result
}
func expandOIDCRequestParameterList(in []interface{}) *[]*pf.OIDCRequestParameter {
	var result []*pf.OIDCRequestParameter
	for _, v := range in {
		result = append(result, expandOIDCRequestParameter(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpWsTrustAttributeList(in []interface{}) *[]*pf.IdpWsTrustAttribute {
	var result []*pf.IdpWsTrustAttribute
	for _, v := range in {
		result = append(result, expandIdpWsTrustAttribute(v.(map[string]interface{})))
	}
	return &result
}
func expandSpSsoServiceEndpointList(in []interface{}) *[]*pf.SpSsoServiceEndpoint {
	var result []*pf.SpSsoServiceEndpoint
	for _, v := range in {
		result = append(result, expandSpSsoServiceEndpoint(v.(map[string]interface{})))
	}
	return &result
}

func expandSpAdapterAttributeList(in []interface{}) *[]*pf.SpAdapterAttribute {
	var contractList []*pf.SpAdapterAttribute
	for _, raw := range in {
		c := &pf.SpAdapterAttribute{
			Name: String(raw.(string)),
		}
		contractList = append(contractList, c)
	}
	return &contractList
}

func expandProtocolMessageCustomizationList(in []interface{}) *[]*pf.ProtocolMessageCustomization {
	var result []*pf.ProtocolMessageCustomization
	for _, v := range in {
		result = append(result, expandProtocolMessageCustomization(v.(map[string]interface{})))
	}
	return &result
}
func expandSpWsTrustAttributeList(in []interface{}) *[]*pf.SpWsTrustAttribute {
	var result []*pf.SpWsTrustAttribute
	for _, v := range in {
		result = append(result, expandSpWsTrustAttribute(v.(map[string]interface{})))
	}
	return &result
}
func expandChannelList(in []interface{}) *[]*pf.Channel {
	var result []*pf.Channel
	for _, v := range in {
		result = append(result, expandChannel(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpAdapterAttributeList(in []interface{}) *[]*pf.IdpAdapterAttribute {
	var result []*pf.IdpAdapterAttribute
	for _, v := range in {
		result = append(result, expandIdpAdapterAttribute(v.(map[string]interface{})))
	}
	return &result
}

//func expandConfigFieldList(in []interface{}) *[]*pf.ConfigField {
//	var result []*pf.ConfigField
//	for _, v := range in {
//		result = append(result, expandConfigField(v.(map[string]interface{})))
//	}
//	return &result
//}
func expandUrlWhitelistEntryList(in []interface{}) *[]*pf.UrlWhitelistEntry {
	var result []*pf.UrlWhitelistEntry
	for _, v := range in {
		result = append(result, expandUrlWhitelistEntry(v.(map[string]interface{})))
	}
	return &result
}
func expandSloServiceEndpointList(in []interface{}) *[]*pf.SloServiceEndpoint {
	var result []*pf.SloServiceEndpoint
	for _, v := range in {
		result = append(result, expandSloServiceEndpoint(v.(map[string]interface{})))
	}
	return &result
}
func expandAuthenticationPolicyContractMappingList(in []interface{}) *[]*pf.AuthenticationPolicyContractMapping {
	var result []*pf.AuthenticationPolicyContractMapping
	for _, v := range in {
		result = append(result, expandAuthenticationPolicyContractMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandAccessTokenManagerMappingList(in []interface{}) *[]*pf.AccessTokenManagerMapping {
	var result []*pf.AccessTokenManagerMapping
	for _, v := range in {
		result = append(result, expandAccessTokenManagerMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandConditionalIssuanceCriteriaEntryList(in []interface{}) *[]*pf.ConditionalIssuanceCriteriaEntry {
	var result []*pf.ConditionalIssuanceCriteriaEntry
	for _, v := range in {
		result = append(result, expandConditionalIssuanceCriteriaEntry(v.(map[string]interface{})))
	}
	return &result
}
func expandConfigTableList(in []interface{}) *[]*pf.ConfigTable {
	var result []*pf.ConfigTable
	for _, v := range in {
		result = append(result, expandConfigTable(v.(map[string]interface{})))
	}
	return &result
}
func expandAuthnContextMappingList(in []interface{}) *[]*pf.AuthnContextMapping {
	var result []*pf.AuthnContextMapping
	for _, v := range in {
		result = append(result, expandAuthnContextMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpBrowserSsoAttributeList(in []interface{}) *[]*pf.IdpBrowserSsoAttribute {
	var result []*pf.IdpBrowserSsoAttribute
	for _, v := range in {
		result = append(result, expandIdpBrowserSsoAttribute(v.(map[string]interface{})))
	}
	return &result
}
func expandSchemaAttributeList(in []interface{}) *[]*pf.SchemaAttribute {
	var result []*pf.SchemaAttribute
	for _, v := range in {
		result = append(result, expandSchemaAttribute(v.(map[string]interface{})))
	}
	return &result
}
func expandAttributeQueryNameMappingList(in []interface{}) *[]*pf.AttributeQueryNameMapping {
	var result []*pf.AttributeQueryNameMapping
	for _, v := range in {
		result = append(result, expandAttributeQueryNameMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandConnectionCertList(in []interface{}) *[]*pf.ConnectionCert {
	var result []*pf.ConnectionCert
	for _, v := range in {
		result = append(result, expandConnectionCert(v.(map[string]interface{})))
	}
	return &result
}
func expandSpAdapterMappingList(in []interface{}) *[]*pf.SpAdapterMapping {
	var result []*pf.SpAdapterMapping
	for _, v := range in {
		result = append(result, expandSpAdapterMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandExpressionIssuanceCriteriaEntryList(in []interface{}) *[]*pf.ExpressionIssuanceCriteriaEntry {
	var result []*pf.ExpressionIssuanceCriteriaEntry
	for _, v := range in {
		result = append(result, expandExpressionIssuanceCriteriaEntry(v.(map[string]interface{})))
	}
	return &result
}
func expandAuthenticationPolicyContractAssertionMappingList(in []interface{}) *[]*pf.AuthenticationPolicyContractAssertionMapping {
	var result []*pf.AuthenticationPolicyContractAssertionMapping
	for _, v := range in {
		result = append(result, expandAuthenticationPolicyContractAssertionMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandFieldEntryList(in []interface{}) *[]*pf.FieldEntry {
	var result []*pf.FieldEntry
	for _, v := range in {
		result = append(result, expandFieldEntry(v.(map[string]interface{})))
	}
	return &result
}
func expandSaasAttributeMappingList(in []interface{}) *[]*pf.SaasAttributeMapping {
	var result []*pf.SaasAttributeMapping
	for _, v := range in {
		result = append(result, expandSaasAttributeMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandEntityList(in []interface{}) *[]*pf.Entity {
	var result []*pf.Entity
	for _, v := range in {
		result = append(result, expandEntity(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpSsoServiceEndpointList(in []interface{}) *[]*pf.IdpSsoServiceEndpoint {
	var result []*pf.IdpSsoServiceEndpoint
	for _, v := range in {
		result = append(result, expandIdpSsoServiceEndpoint(v.(map[string]interface{})))
	}
	return &result
}

//func expandAttributeSourceList(in []interface{}) *[]*pf.AttributeSource {
//	var result []*pf.AttributeSource
//	for _, v := range in {
//		result = append(result, expandAttributeSource(v.(map[string]interface{})))
//	}
//	return &result
//}

func expandLdapAttributeSourceList(in []interface{}) *[]*pf.AttributeSource {
	var result []*pf.AttributeSource
	for _, v := range in {
		result = append(result, expandLdapAttributeSource(v.(map[string]interface{})))
	}
	return &result
}

func expandJdbcAttributeSourceList(in []interface{}) *[]*pf.AttributeSource {
	var result []*pf.AttributeSource
	for _, v := range in {
		result = append(result, expandJdbcAttributeSource(v.(map[string]interface{})))
	}
	return &result
}

func expandCustomAttributeSourceList(in []interface{}) *[]*pf.AttributeSource {
	var result []*pf.AttributeSource
	for _, v := range in {
		result = append(result, expandCustomAttributeSource(v.(map[string]interface{})))
	}
	return &result
}

func expandConfigRowList(in []interface{}) *[]*pf.ConfigRow {
	var result []*pf.ConfigRow
	for _, v := range in {
		result = append(result, expandConfigRow(v.(map[string]interface{})))
	}
	return &result
}

func expandSpBrowserSsoAttributeList(in []interface{}) *[]*pf.SpBrowserSsoAttribute {
	var result []*pf.SpBrowserSsoAttribute
	for _, v := range in {
		result = append(result, expandSpBrowserSsoAttribute(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpAdapterAssertionMappingList(in []interface{}) *[]*pf.IdpAdapterAssertionMapping {
	var result []*pf.IdpAdapterAssertionMapping
	for _, v := range in {
		result = append(result, expandIdpAdapterAssertionMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandArtifactResolverLocationList(in []interface{}) *[]*pf.ArtifactResolverLocation {
	var result []*pf.ArtifactResolverLocation
	for _, v := range in {
		result = append(result, expandArtifactResolverLocation(v.(map[string]interface{})))
	}
	return &result
}
func expandIdpTokenProcessorMappingList(in []interface{}) *[]*pf.IdpTokenProcessorMapping {
	var result []*pf.IdpTokenProcessorMapping
	for _, v := range in {
		result = append(result, expandIdpTokenProcessorMapping(v.(map[string]interface{})))
	}
	return &result
}
func expandSpTokenGeneratorMappingList(in []interface{}) *[]*pf.SpTokenGeneratorMapping {
	var result []*pf.SpTokenGeneratorMapping
	for _, v := range in {
		result = append(result, expandSpTokenGeneratorMapping(v.(map[string]interface{})))
	}
	return &result
}

func expandLdapTagConfigList(in []interface{}) *[]*pf.LdapTagConfig {
	var result []*pf.LdapTagConfig
	for _, v := range in {
		result = append(result, expandLdapTagConfig(v.(map[string]interface{})))
	}
	return &result
}

func expandLdapTagConfig(in map[string]interface{}) *pf.LdapTagConfig {
	var result pf.LdapTagConfig
	if val, ok := in["hostnames"]; ok {
		strs := expandStringList(val.(*schema.Set).List())
		result.Hostnames = &strs
	}
	if val, ok := in["tags"]; ok {
		result.Tags = String(val.(string))
	}
	if val, ok := in["default_source"]; ok {
		result.DefaultSource = Bool(val.(bool))
	}
	return &result
}

func expandSelectionSettings(in []interface{}) *pf.AtmSelectionSettings {
	settings := &pf.AtmSelectionSettings{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["inherited"]; ok {
			settings.Inherited = Bool(val.(bool))
		}
		if v, ok := l["resource_uris"]; ok && len(v.([]interface{})) > 0 {
			sans := expandStringList(v.([]interface{}))
			settings.ResourceUris = &sans
		}
	}
	return settings
}

func expandSessionValidationSettings(in []interface{}) *pf.SessionValidationSettings {
	settings := &pf.SessionValidationSettings{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["inherited"]; ok {
			settings.Inherited = Bool(val.(bool))
		}
		if val, ok := l["check_valid_authn_session"]; ok {
			settings.CheckValidAuthnSession = Bool(val.(bool))
		}
		if val, ok := l["include_session_id"]; ok {
			settings.IncludeSessionId = Bool(val.(bool))
		}
		if val, ok := l["check_session_revocation_status"]; ok {
			settings.CheckSessionRevocationStatus = Bool(val.(bool))
		}
		if val, ok := l["update_authn_session_activity"]; ok {
			settings.UpdateAuthnSessionActivity = Bool(val.(bool))
		}
	}
	return settings
}

func expandAccessControlSettings(in []interface{}) *pf.AtmAccessControlSettings {
	settings := &pf.AtmAccessControlSettings{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if val, ok := l["inherited"]; ok {
			settings.Inherited = Bool(val.(bool))
		}
		if val, ok := l["restrict_clients"]; ok {
			settings.RestrictClients = Bool(val.(bool))
		}
		if val, ok := l["allowed_clients"]; ok {
			var clients []*pf.ResourceLink
			for _, s := range val.([]interface{}) {
				clients = append(clients, &pf.ResourceLink{Id: String(s.(string))})
			}
			settings.AllowedClients = &clients
		}
	}
	return settings
}

func expandOcspSettings(in map[string]interface{}) *pf.OcspSettings {
	var result pf.OcspSettings
	if val, ok := in["requester_add_nonce"]; ok {
		result.RequesterAddNonce = Bool(val.(bool))
	}
	if val, ok := in["responder_url"]; ok {
		result.ResponderUrl = String(val.(string))
	}
	if val, ok := in["current_update_grace_period"]; ok {
		result.CurrentUpdateGracePeriod = Int(val.(int))
	}
	if val, ok := in["next_update_grace_period"]; ok {
		result.NextUpdateGracePeriod = Int(val.(int))
	}
	if val, ok := in["responder_timeout"]; ok {
		result.ResponderTimeout = Int(val.(int))
	}
	if val, ok := in["action_on_status_unknown"]; ok {
		result.ActionOnStatusUnknown = String(val.(string))
	}
	if val, ok := in["action_on_unsuccessful_response"]; ok {
		result.ActionOnUnsuccessfulResponse = String(val.(string))
	}
	if val, ok := in["responder_cert_reference"]; ok && len(val.([]interface{})) > 0 {
		result.ResponderCertReference = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	if val, ok := in["response_cache_period"]; ok {
		result.ResponseCachePeriod = Int(val.(int))
	}
	if val, ok := in["action_on_responder_unavailable"]; ok {
		result.ActionOnResponderUnavailable = String(val.(string))
	}
	return &result
}

func expandCrlSettings(in map[string]interface{}) *pf.CrlSettings {
	var result pf.CrlSettings
	if val, ok := in["treat_non_retrievable_crl_as_revoked"]; ok {
		result.TreatNonRetrievableCrlAsRevoked = Bool(val.(bool))
	}
	if val, ok := in["verify_crl_signature"]; ok {
		result.VerifyCrlSignature = Bool(val.(bool))
	}
	if val, ok := in["next_retry_mins_when_resolve_failed"]; ok {
		result.NextRetryMinsWhenResolveFailed = Int(val.(int))
	}
	if val, ok := in["next_retry_mins_when_next_update_in_past"]; ok {
		result.NextRetryMinsWhenNextUpdateInPast = Int(val.(int))
	}
	return &result
}

func expandProxySettings(in map[string]interface{}) *pf.ProxySettings {
	var result pf.ProxySettings
	if val, ok := in["host"]; ok {
		result.Host = String(val.(string))
	}
	if val, ok := in["port"]; ok {
		result.Port = Int(val.(int))
	}
	return &result
}

func expandClientRegistrationOIDCPolicy(in map[string]interface{}) *pf.ClientRegistrationOIDCPolicy {
	var result pf.ClientRegistrationOIDCPolicy
	if val, ok := in["id_token_signing_algorithm"]; ok {
		result.IdTokenSigningAlgorithm = String(val.(string))
	}
	if val, ok := in["id_token_encryption_algorithm"]; ok {
		result.IdTokenEncryptionAlgorithm = String(val.(string))
	}
	if val, ok := in["id_token_content_encryption_algorithm"]; ok {
		result.IdTokenContentEncryptionAlgorithm = String(val.(string))
	}
	if val, ok := in["policy_group"]; ok && len(val.([]interface{})) > 0 {
		result.PolicyGroup = expandResourceLink(val.([]interface{})[0].(map[string]interface{}))
	}
	return &result
}

func expandResourceLinkList(configured []interface{}) *[]*pf.ResourceLink {
	vs := make([]*pf.ResourceLink, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, &pf.ResourceLink{Id: String(val)})
		}
	}
	return &vs
}

func expandClientMetadataList(in []interface{}) *[]*pf.ClientMetadata {
	var result []*pf.ClientMetadata
	for _, v := range in {
		result = append(result, expandClientMetadata(v.(map[string]interface{})))
	}
	return &result
}

func expandClientMetadata(in map[string]interface{}) *pf.ClientMetadata {
	var result pf.ClientMetadata
	if val, ok := in["parameter"]; ok {
		result.Parameter = String(val.(string))
	}
	if val, ok := in["description"]; ok {
		result.Description = String(val.(string))
	}
	if val, ok := in["multi_valued"]; ok {
		result.MultiValued = Bool(val.(bool))
	}
	return &result
}

func expandTokenProcessorAttributes(in []interface{}) *[]*pf.TokenProcessorAttribute {
	attributes := &[]*pf.TokenProcessorAttribute{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		c := &pf.TokenProcessorAttribute{}
		if val, ok := l["name"]; ok {
			c.Name = String(val.(string))
		}
		if val, ok := l["masked"]; ok {
			c.Masked = Bool(val.(bool))
		}
		*attributes = append(*attributes, c)
	}
	return attributes
}

func expandTokenProcessorAttributeContract(in []interface{}) *pf.TokenProcessorAttributeContract {
	tpac := &pf.TokenProcessorAttributeContract{}
	for _, raw := range in {
		l := raw.(map[string]interface{})
		if v, ok := l["extended_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			tpac.ExtendedAttributes = expandTokenProcessorAttributes(v.(*schema.Set).List())
		}
		if v, ok := l["core_attributes"]; ok && len(v.(*schema.Set).List()) > 0 {
			tpac.CoreAttributes = expandTokenProcessorAttributes(v.(*schema.Set).List())
		}
		if v, ok := l["mask_ognl_values"]; ok {
			tpac.MaskOgnlValues = Bool(v.(bool))
		}
		if v, ok := l["inherited"]; ok {
			tpac.Inherited = Bool(v.(bool))
		}
	}
	return tpac
}
