package pingfederate

import (
	"fmt"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func TestAccPingFederateSpIdpConnection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateSpIdpConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateSpIdpConnectionConfig("https://foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpIdpConnectionExists("pingfederate_sp_idp_connection.demo"),
					//testAccCheckPingFederateSpIdpConnectionAttributes(),
				),
			},
			{
				Config: testAccPingFederateSpIdpConnectionConfig("https://bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateSpIdpConnectionExists("pingfederate_sp_idp_connection.demo"),
				),
			},
		},
	})
}

func testAccCheckPingFederateSpIdpConnectionDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateSpIdpConnectionConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_sp_idp_connection" "demo" {
  name = "foo"
  entity_id = "foo"
  active = true
  logging_mode = "STANDARD"
  credentials {
    outbound_back_channel_auth {
      type = "OUTBOUND"
      digital_signature = false
      validate_partner_cert = false
    }
  }
  attribute_query {
	url = "%s"
    policy {
      sign_attribute_query = false
	  encrypt_name_id = false
	  require_signed_response = false
	  require_signed_assertion = false
	  require_encrypted_assertion = false
	  mask_attribute_values = false
    }
  }
}
`, configUpdate)
}

func testAccCheckPingFederateSpIdpConnectionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("no rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).SpIdpConnections
		result, _, err := conn.GetConnection(&spIdpConnections.GetConnectionInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("error: SpIdpConnection (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("error: SpIdpConnection response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateSpIdpConnectionResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.IdpConnection
	}{
		{
			Resource: pf.IdpConnection{
				Active: Bool(false),
				AdditionalAllowedEntitiesConfiguration: &pf.AdditionalAllowedEntitiesConfiguration{
					AdditionalAllowedEntities: &[]*pf.Entity{
						{
							EntityDescription: String("foo"),
							EntityId:          String("foo"),
						},
					},
					AllowAdditionalEntities: Bool(true),
					AllowAllEntities:        Bool(true),
				},
				AttributeQuery: &pf.IdpAttributeQuery{
					NameMappings: &[]*pf.AttributeQueryNameMapping{
						{
							LocalName:  String("foo"),
							RemoteName: String("foo"),
						},
					},
					Policy: &pf.IdpAttributeQueryPolicy{
						EncryptNameId:             Bool(true),
						MaskAttributeValues:       Bool(true),
						RequireEncryptedAssertion: Bool(true),
						RequireSignedAssertion:    Bool(true),
						RequireSignedResponse:     Bool(true),
						SignAttributeQuery:        Bool(true),
					},
					Url: String("foo"),
				},
				BaseUrl: String("foo"),
				ContactInfo: &pf.ContactInfo{
					Company:   String("foo"),
					Email:     String("foo"),
					FirstName: String("foo"),
					LastName:  String("foo"),
					Phone:     String("foo"),
				},
				Credentials: &pf.ConnectionCredentials{
					BlockEncryptionAlgorithm: String("foo"),
					Certs: &[]*pf.ConnectionCert{
						{
							ActiveVerificationCert: Bool(true),
							CertView: &pf.CertView{
								CryptoProvider:          String("foo"),
								Expires:                 String("foo"),
								Id:                      String("foo"),
								IssuerDN:                String("foo"),
								KeyAlgorithm:            String("foo"),
								KeySize:                 Int(1024),
								SerialNumber:            String("foo"),
								Sha1Fingerprint:         String("foo"),
								Sha256Fingerprint:       String("foo"),
								SignatureAlgorithm:      String("foo"),
								Status:                  String("foo"),
								SubjectAlternativeNames: &[]*string{String("foo")},
								SubjectDN:               String("foo"),
								ValidFrom:               String("foo"),
								Version:                 Int(1),
							},
							EncryptionCert:            Bool(true),
							PrimaryVerificationCert:   Bool(true),
							SecondaryVerificationCert: Bool(true),
							X509File: &pf.X509File{
								CryptoProvider: String("foo"),
								FileData:       String("foo"),
								Id:             String("foo"),
							},
						},
					},
					DecryptionKeyPairRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					InboundBackChannelAuth: &pf.InboundBackChannelAuth{
						Certs: &[]*pf.ConnectionCert{
							{
								ActiveVerificationCert: Bool(true),
								CertView: &pf.CertView{
									CryptoProvider:          String("foo"),
									Expires:                 String("foo"),
									Id:                      String("foo"),
									IssuerDN:                String("foo"),
									KeyAlgorithm:            String("foo"),
									KeySize:                 Int(1024),
									SerialNumber:            String("foo"),
									Sha1Fingerprint:         String("foo"),
									Sha256Fingerprint:       String("foo"),
									SignatureAlgorithm:      String("foo"),
									Status:                  String("foo"),
									SubjectAlternativeNames: &[]*string{String("foo")},
									SubjectDN:               String("foo"),
									ValidFrom:               String("foo"),
									Version:                 Int(1),
								},
								EncryptionCert:            Bool(true),
								PrimaryVerificationCert:   Bool(true),
								SecondaryVerificationCert: Bool(true),
								X509File: &pf.X509File{
									CryptoProvider: String("foo"),
									FileData:       String("foo"),
									Id:             String("foo"),
								},
							},
						},
						DigitalSignature: Bool(true),
						HttpBasicCredentials: &pf.UsernamePasswordCredentials{
							EncryptedPassword: String("foo"),
							Password:          String("foo"),
							Username:          String("foo"),
						},
						RequireSsl:            Bool(true),
						Type:                  String("foo"),
						VerificationIssuerDN:  String("foo"),
						VerificationSubjectDN: String("foo"),
					},
					KeyTransportAlgorithm: String("foo"),
					OutboundBackChannelAuth: &pf.OutboundBackChannelAuth{
						DigitalSignature: Bool(true),
						HttpBasicCredentials: &pf.UsernamePasswordCredentials{
							EncryptedPassword: String("foo"),
							Password:          String("foo"),
							Username:          String("foo"),
						},
						SslAuthKeyPairRef: &pf.ResourceLink{
							Id:       String("foo"),
							Location: String("foo"),
						},
						Type:                String("foo"),
						ValidatePartnerCert: Bool(true),
					},
					SecondaryDecryptionKeyPairRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					SigningSettings: &pf.SigningSettings{
						Algorithm:                String("foo"),
						IncludeCertInSignature:   Bool(true),
						IncludeRawKeyInSignature: Bool(true),
						SigningKeyPairRef: &pf.ResourceLink{
							Id:       String("foo"),
							Location: String("foo"),
						},
					},
					VerificationIssuerDN:  String("foo"),
					VerificationSubjectDN: String("foo"),
				},
				DefaultVirtualEntityId: String("foo"),
				EntityId:               String("foo"),
				ErrorPageMsgId:         String("foo"),
				ExtendedProperties: map[string]*pf.ParameterValues{
					"foo": {
						Values: &[]*string{String("foo")},
					},
				},
				Id: String("foo"),
				IdpBrowserSso: &pf.IdpBrowserSso{
					AdapterMappings: &[]*pf.SpAdapterMapping{
						{
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{
										Id:   String("foo"),
										Type: String("foo"),
									},
									Value: String("foo"),
								},
							},
							AttributeSources: &[]*pf.AttributeSource{
								{
									LdapAttributeSource: pf.LdapAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										BaseDn: String("foo"),
										BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
											"foo": {BinaryEncoding: String("foo")},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description:         String("foo"),
										Id:                  String("foo"),
										MemberOfNestedGroup: Bool(true),
										SearchFilter:        String("foo"),
										SearchScope:         String("foo"),
										Type:                String("LDAP"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("LDAP"),
								},
								{
									JdbcAttributeSource: pf.JdbcAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										Filter:      String("foo"),
										Id:          String("foo"),
										Schema:      String("foo"),
										Table:       String("foo"),
										Type:        String("JDBC"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("JDBC"),
								},
								{
									CustomAttributeSource: pf.CustomAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										FilterFields: &[]*pf.FieldEntry{
											{
												Name:  String("foo"),
												Value: String("foo"),
											},
										},
										Id:   String("foo"),
										Type: String("CUSTOM"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("CUSTOM"),
								},
							},
							IssuanceCriteria: &pf.IssuanceCriteria{
								ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
									{
										AttributeName: String("foo"),
										Condition:     String("foo"),
										ErrorResult:   String("foo"),
										Source: &pf.SourceTypeIdKey{
											Id:   String("foo"),
											Type: String("foo"),
										},
										Value: String("foo"),
									},
								},
								ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
									{
										ErrorResult: String("foo"),
										Expression:  String("foo"),
									},
								},
							},
							RestrictVirtualEntityIds:   Bool(true),
							RestrictedVirtualEntityIds: &[]*string{String("foo")},
							SpAdapterRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
						},
					},
					Artifact: &pf.ArtifactSettings{
						Lifetime: Int(1),
						ResolverLocations: &[]*pf.ArtifactResolverLocation{
							{
								Index: Int(1),
								Url:   String("foo"),
							},
						},
						SourceId: String("foo"),
					},
					AssertionsSigned: Bool(true),
					AttributeContract: &pf.IdpBrowserSsoAttributeContract{
						CoreAttributes: &[]*pf.IdpBrowserSsoAttribute{
							{
								Name:   String("foo"),
								Masked: Bool(true),
							},
						},
						ExtendedAttributes: &[]*pf.IdpBrowserSsoAttribute{
							{
								Name:   String("foo"),
								Masked: Bool(true),
							},
						},
					},
					AuthenticationPolicyContractMappings: &[]*pf.AuthenticationPolicyContractMapping{
						{
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{
										Id:   String("foo"),
										Type: String("foo"),
									},
									Value: String("foo"),
								},
							},
							AttributeSources: &[]*pf.AttributeSource{
								{
									LdapAttributeSource: pf.LdapAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										BaseDn: String("foo"),
										BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
											"foo": {BinaryEncoding: String("foo")},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description:         String("foo"),
										Id:                  String("foo"),
										MemberOfNestedGroup: Bool(true),
										SearchFilter:        String("foo"),
										SearchScope:         String("foo"),
										Type:                String("LDAP"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("LDAP"),
								},
								{
									JdbcAttributeSource: pf.JdbcAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										Filter:      String("foo"),
										Id:          String("foo"),
										Schema:      String("foo"),
										Table:       String("foo"),
										Type:        String("JDBC"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("JDBC"),
								},
								{
									CustomAttributeSource: pf.CustomAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										FilterFields: &[]*pf.FieldEntry{
											{
												Name:  String("foo"),
												Value: String("foo"),
											},
										},
										Id:   String("foo"),
										Type: String("CUSTOM"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("CUSTOM"),
								},
							},
							AuthenticationPolicyContractRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
							IssuanceCriteria: &pf.IssuanceCriteria{
								ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
									{
										AttributeName: String("foo"),
										Condition:     String("foo"),
										ErrorResult:   String("foo"),
										Source: &pf.SourceTypeIdKey{
											Id:   String("foo"),
											Type: String("foo"),
										},
										Value: String("foo"),
									},
								},
								ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
									{
										ErrorResult: String("foo"),
										Expression:  String("foo"),
									},
								},
							},
							RestrictVirtualServerIds:   Bool(true),
							RestrictedVirtualServerIds: &[]*string{String("foo")},
						},
					},
					AuthnContextMappings: &[]*pf.AuthnContextMapping{
						{
							Local:  String("foo"),
							Remote: String("foo"),
						},
					},
					DecryptionPolicy: &pf.DecryptionPolicy{
						AssertionEncrypted:        Bool(true),
						AttributesEncrypted:       Bool(true),
						SloEncryptSubjectNameID:   Bool(true),
						SloSubjectNameIDEncrypted: Bool(true),
						SubjectNameIdEncrypted:    Bool(true),
					},
					DefaultTargetUrl:   String("foo"),
					EnabledProfiles:    &[]*string{String("foo")},
					IdpIdentityMapping: String("foo"),
					IncomingBindings:   &[]*string{String("foo")},
					MessageCustomizations: &[]*pf.ProtocolMessageCustomization{
						{
							ContextName:       String("foo"),
							MessageExpression: String("foo"),
						},
					},
					OauthAuthenticationPolicyContractRef: &pf.ResourceLink{
						Id: String("foo"),
					},
					OidcProviderSettings: &pf.OIDCProviderSettings{
						AuthenticationScheme:           String("foo"),
						AuthenticationSigningAlgorithm: String("foo"),
						AuthorizationEndpoint:          String("foo"),
						JwksURL:                        String("foo"),
						LoginType:                      String("foo"),
						RequestParameters: &[]*pf.OIDCRequestParameter{
							{
								ApplicationEndpointOverride: Bool(true),
								Name:                        String("foo"),
								Value:                       String("foo"),
							},
						},
						RequestSigningAlgorithm: String("foo"),
						Scopes:                  String("foo"),
						TokenEndpoint:           String("foo"),
						UserInfoEndpoint:        String("foo"),
					},
					Protocol:          String("foo"),
					SignAuthnRequests: Bool(true),
					SloServiceEndpoints: &[]*pf.SloServiceEndpoint{
						{
							Binding:     String("foo"),
							ResponseUrl: String("foo"),
							Url:         String("foo"),
						},
					},
					SsoOAuthMapping: &pf.SsoOAuthMapping{
						AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
							"foo": {
								Source: &pf.SourceTypeIdKey{
									Id:   String("foo"),
									Type: String("foo"),
								},
								Value: String("foo"),
							},
						},
						AttributeSources: &[]*pf.AttributeSource{
							{
								LdapAttributeSource: pf.LdapAttributeSource{
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									BaseDn: String("foo"),
									BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
										"foo": {BinaryEncoding: String("foo")},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description:         String("foo"),
									Id:                  String("foo"),
									MemberOfNestedGroup: Bool(true),
									SearchFilter:        String("foo"),
									SearchScope:         String("foo"),
									Type:                String("LDAP"),
								},
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
										Value:  String("foo"),
									},
								},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description: String("foo"),
								Id:          String("foo"),
								Type:        String("LDAP"),
							},
							{
								JdbcAttributeSource: pf.JdbcAttributeSource{
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Filter:      String("foo"),
									Id:          String("foo"),
									Schema:      String("foo"),
									Table:       String("foo"),
									Type:        String("JDBC"),
								},
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
										Value:  String("foo"),
									},
								},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description: String("foo"),
								Id:          String("foo"),
								Type:        String("JDBC"),
							},
							{
								CustomAttributeSource: pf.CustomAttributeSource{
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									FilterFields: &[]*pf.FieldEntry{
										{
											Name:  String("foo"),
											Value: String("foo"),
										},
									},
									Id:   String("foo"),
									Type: String("CUSTOM"),
								},
								AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
									"foo": {
										Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
										Value:  String("foo"),
									},
								},
								DataStoreRef: &pf.ResourceLink{
									Id: String("foo"),
								},
								Description: String("foo"),
								Id:          String("foo"),
								Type:        String("CUSTOM"),
							},
						},
						IssuanceCriteria: &pf.IssuanceCriteria{
							ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
								{
									AttributeName: String("foo"),
									Condition:     String("foo"),
									ErrorResult:   String("foo"),
									Source: &pf.SourceTypeIdKey{
										Id:   String("foo"),
										Type: String("foo"),
									},
									Value: String("foo"),
								},
							},
							ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
								{
									ErrorResult: String("foo"),
									Expression:  String("foo"),
								},
							},
						},
					},
					SsoServiceEndpoints: &[]*pf.IdpSsoServiceEndpoint{
						{
							Binding: String("foo"),
							Url:     String("foo"),
						},
					},
					UrlWhitelistEntries: &[]*pf.UrlWhitelistEntry{
						{
							AllowQueryAndFragment: Bool(true),
							RequireHttps:          Bool(true),
							ValidDomain:           String("foo"),
							ValidPath:             String("foo"),
						},
					},
				},
				IdpOAuthGrantAttributeMapping: &pf.IdpOAuthGrantAttributeMapping{
					AccessTokenManagerMappings: &[]*pf.AccessTokenManagerMapping{
						{
							AccessTokenManagerRef: &pf.ResourceLink{
								Id: String("foo"),
							},
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{
										Id:   String("foo"),
										Type: String("foo"),
									},
									Value: String("foo"),
								},
							},
							AttributeSources: &[]*pf.AttributeSource{
								{
									LdapAttributeSource: pf.LdapAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										BaseDn: String("foo"),
										BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
											"foo": {BinaryEncoding: String("foo")},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description:         String("foo"),
										Id:                  String("foo"),
										MemberOfNestedGroup: Bool(true),
										SearchFilter:        String("foo"),
										SearchScope:         String("foo"),
										Type:                String("LDAP"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("LDAP"),
								},
								{
									JdbcAttributeSource: pf.JdbcAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										Filter:      String("foo"),
										Id:          String("foo"),
										Schema:      String("foo"),
										Table:       String("foo"),
										Type:        String("JDBC"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("JDBC"),
								},
								{
									CustomAttributeSource: pf.CustomAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										FilterFields: &[]*pf.FieldEntry{
											{
												Name:  String("foo"),
												Value: String("foo"),
											},
										},
										Id:   String("foo"),
										Type: String("CUSTOM"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("CUSTOM"),
								},
							},
							IssuanceCriteria: &pf.IssuanceCriteria{
								ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
									{
										AttributeName: String("foo"),
										Condition:     String("foo"),
										ErrorResult:   String("foo"),
										Source: &pf.SourceTypeIdKey{
											Id:   String("foo"),
											Type: String("foo"),
										},
										Value: String("foo"),
									},
								},
								ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
									{
										ErrorResult: String("foo"),
										Expression:  String("foo"),
									},
								},
							},
						},
					},
					IdpOAuthAttributeContract: &pf.IdpOAuthAttributeContract{
						CoreAttributes: &[]*pf.IdpBrowserSsoAttribute{
							{
								Masked: Bool(true),
								Name:   String("foo"),
							},
						},
						ExtendedAttributes: &[]*pf.IdpBrowserSsoAttribute{
							{
								Masked: Bool(true),
								Name:   String("foo"),
							},
						},
					},
				},
				LicenseConnectionGroup: String("foo"),
				LoggingMode:            String("foo"),
				MetadataReloadSettings: &pf.ConnectionMetadataUrl{
					EnableAutoMetadataUpdate: Bool(true),
					MetadataUrlRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
				},
				Name: String("foo"),
				OidcClientCredentials: &pf.OIDCClientCredentials{
					ClientId:        String("foo"),
					ClientSecret:    String("foo"),
					EncryptedSecret: String("foo"),
				},
				Type:             String("SP"),
				VirtualEntityIds: &[]*string{String("foo")},
				WsTrust: &pf.IdpWsTrust{
					AttributeContract: &pf.IdpWsTrustAttributeContract{
						CoreAttributes: &[]*pf.IdpWsTrustAttribute{
							{
								Name:   String("foo"),
								Masked: Bool(true),
							},
						},
						ExtendedAttributes: &[]*pf.IdpWsTrustAttribute{
							{
								Name:   String("foo"),
								Masked: Bool(true),
							},
						},
					},
					GenerateLocalToken: Bool(true),
					TokenGeneratorMappings: &[]*pf.SpTokenGeneratorMapping{
						{
							AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
								"foo": {
									Source: &pf.SourceTypeIdKey{
										Id:   String("foo"),
										Type: String("foo"),
									},
									Value: String("foo"),
								},
							},
							AttributeSources: &[]*pf.AttributeSource{
								{
									LdapAttributeSource: pf.LdapAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										BaseDn: String("foo"),
										BinaryAttributeSettings: map[string]*pf.BinaryLdapAttributeSettings{
											"foo": {BinaryEncoding: String("foo")},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description:         String("foo"),
										Id:                  String("foo"),
										MemberOfNestedGroup: Bool(true),
										SearchFilter:        String("foo"),
										SearchScope:         String("foo"),
										Type:                String("LDAP"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("LDAP"),
								},
								{
									JdbcAttributeSource: pf.JdbcAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										Filter:      String("foo"),
										Id:          String("foo"),
										Schema:      String("foo"),
										Table:       String("foo"),
										Type:        String("JDBC"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("JDBC"),
								},
								{
									CustomAttributeSource: pf.CustomAttributeSource{
										AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
											"foo": {
												Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
												Value:  String("foo"),
											},
										},
										DataStoreRef: &pf.ResourceLink{
											Id: String("foo"),
										},
										Description: String("foo"),
										FilterFields: &[]*pf.FieldEntry{
											{
												Name:  String("foo"),
												Value: String("foo"),
											},
										},
										Id:   String("foo"),
										Type: String("CUSTOM"),
									},
									AttributeContractFulfillment: map[string]*pf.AttributeFulfillmentValue{
										"foo": {
											Source: &pf.SourceTypeIdKey{Type: String("foo"), Id: String("bar")},
											Value:  String("foo"),
										},
									},
									DataStoreRef: &pf.ResourceLink{
										Id: String("foo"),
									},
									Description: String("foo"),
									Id:          String("foo"),
									Type:        String("CUSTOM"),
								},
							},
							DefaultMapping: Bool(true),
							IssuanceCriteria: &pf.IssuanceCriteria{
								ConditionalCriteria: &[]*pf.ConditionalIssuanceCriteriaEntry{
									{
										AttributeName: String("foo"),
										Condition:     String("foo"),
										ErrorResult:   String("foo"),
										Source: &pf.SourceTypeIdKey{
											Id:   String("foo"),
											Type: String("foo"),
										},
										Value: String("foo"),
									},
								},
								ExpressionCriteria: &[]*pf.ExpressionIssuanceCriteriaEntry{
									{
										ErrorResult: String("foo"),
										Expression:  String("foo"),
									},
								},
							},
							RestrictedVirtualEntityIds: &[]*string{String("foo")},
							SpTokenGeneratorRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateSpIdpConnectionResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateSpIdpConnectionResourceReadResult(resourceLocalData, &tc.Resource)

			if got := *resourcePingFederateSpIdpConnectionResourceReadData(resourceLocalData); !cmp.Equal(got, tc.Resource) {
				t.Errorf("resourcePingFederateSpIdpConnectionResourceReadData() = %v", cmp.Diff(got, tc.Resource))
			}
		})
	}
}
