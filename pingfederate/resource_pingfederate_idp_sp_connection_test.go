package pingfederate

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iwarapter/pingfederate-sdk-go/services/idpSpConnections"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

func init() {
	resource.AddTestSweepers("idp_sp_connection", &resource.Sweeper{
		Name:         "idp_sp_connection",
		Dependencies: []string{},
		F: func(r string) error {
			svc := idpSpConnections.New(cfg)
			results, _, err := svc.GetConnections(&idpSpConnections.GetConnectionsInput{Filter: "acc_test"})
			if err != nil {
				return fmt.Errorf("unable to list idp sp connections %s", err)
			}
			for _, item := range *results.Items {
				_, _, err := svc.DeleteConnection(&idpSpConnections.DeleteConnectionInput{Id: *item.Id})
				if err != nil {
					return fmt.Errorf("unable to sweep idp sp connection %s because %s", *item.Id, err)
				}
			}
			return nil
		},
	})
}

func TestAccPingFederateIdpSpConnection(t *testing.T) {
	resourceName := "pingfederate_idp_sp_connection.demo"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingFederateIdpSpConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingFederateIdpSpConnectionConfig("foo"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpSpConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "credentials.0.certs.0.cert_view.0.id"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.expires", "2038-01-17T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.issuer_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.serial_number", "143266978916655856878034712317230054538369994"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.sha1_fingerprint", "8DA7F965EC5EFC37910F1C6E59FDC1CC6A6EDE16"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.sha256_fingerprint", "8ECDE6884F3D87B1125BA31AC3FCB13D7016DE7F57CC904FE1CB97C6AE98196E"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.subject_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.valid_from", "2015-05-26T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.version", "3"),
				),
			},
			{
				Config: testAccPingFederateIdpSpConnectionConfig("bar"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingFederateIdpSpConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "credentials.0.certs.0.cert_view.0.id"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.expires", "2038-01-17T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.issuer_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.key_algorithm", "RSA"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.key_size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.serial_number", "143266978916655856878034712317230054538369994"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.sha1_fingerprint", "8DA7F965EC5EFC37910F1C6E59FDC1CC6A6EDE16"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.sha256_fingerprint", "8ECDE6884F3D87B1125BA31AC3FCB13D7016DE7F57CC904FE1CB97C6AE98196E"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.signature_algorithm", "SHA256withRSA"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.status", "VALID"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.subject_dn", "CN=Amazon Root CA 1, O=Amazon, C=US"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.valid_from", "2015-05-26T00:00:00.000Z"),
					resource.TestCheckResourceAttr(resourceName, "credentials.0.certs.0.cert_view.0.version", "3"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingFederateIdpSpConnectionDestroy(s *terraform.State) error {
	return nil
}

func testAccPingFederateIdpSpConnectionConfig(configUpdate string) string {
	return fmt.Sprintf(`
resource "pingfederate_idp_sp_connection" "demo" {
  name = "acc_test_foo"
  entity_id = "foo"
  active = true
  logging_mode = "STANDARD"
  credentials {
	certs {
	  x509_file {
		file_data = file("test_cases/amazon_root_ca1.pem")
	  }
	}
    inbound_back_channel_auth {
      type = "INBOUND"
      digital_signature = false
      require_ssl = false
      verification_subject_dn = "cn=%s"
    }
  }
  attribute_query {
    jdbc_attribute_source {
      filter = "*"
      description = "foo"
      schema = "INFORMATION_SCHEMA"
      table = "ADMINISTRABLE_ROLE_AUTHORIZATIONS"
      id = "foo"
      data_store_ref {
        id = "ProvisionerDS"
      }
    }

    attribute_contract_fulfillment {
      key_name = "foo"
      source {
        type = "JDBC_DATA_STORE"
        id = "foo"
      }
      value = "GRANTEE"
    }

    attributes = ["foo"]
    policy {
      sign_response = false
      sign_assertion = false
      encrypt_assertion = false
      require_signed_attribute_query = false
      require_encrypted_name_id = false
    }
  }
}
`, configUpdate)
}

func testAccCheckPingFederateIdpSpConnectionExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" || rs.Primary.ID == "0" {
			return fmt.Errorf("No rule ID is set")
		}

		conn := testAccProvider.Meta().(pfClient).IdpSpConnections
		result, _, err := conn.GetConnection(&idpSpConnections.GetConnectionInput{Id: rs.Primary.ID})

		if err != nil {
			return fmt.Errorf("Error: IdpSpConnection (%s) not found", n)
		}

		if *result.Name != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Error: IdpSpConnection response (%s) didnt match state (%s)", *result.Name, rs.Primary.Attributes["name"])
		}

		return nil
	}
}

func Test_resourcePingFederateIdpSpConnectionResourceReadData(t *testing.T) {
	cases := []struct {
		Resource pf.SpConnection
	}{
		{
			Resource: pf.SpConnection{
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
				ApplicationIconUrl: String("foo"),
				ApplicationName:    String("foo"),
				AttributeQuery: &pf.SpAttributeQuery{
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
					Attributes: &[]*string{String("foo")},
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
					Policy: &pf.SpAttributeQueryPolicy{
						EncryptAssertion:            Bool(true),
						RequireEncryptedNameId:      Bool(true),
						RequireSignedAttributeQuery: Bool(true),
						SignAssertion:               Bool(true),
						SignResponse:                Bool(true),
					},
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
				ExtendedProperties: map[string]*pf.ParameterValues{
					"foo": {
						Values: &[]*string{String("foo")},
					},
				},
				Id:                     String("foo"),
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
				OutboundProvision: &pf.OutboundProvision{
					Channels: &[]*pf.Channel{
						{
							Active: Bool(true),
							AttributeMapping: &[]*pf.SaasAttributeMapping{
								{
									FieldName: String("foo"),
									SaasFieldInfo: &pf.SaasFieldConfiguration{
										AttributeNames: &[]*string{String("foo")},
										CharacterCase:  String("foo"),
										CreateOnly:     Bool(true),
										DefaultValue:   String("foo"),
										Expression:     String("foo"),
										Masked:         Bool(true),
										Parser:         String("foo"),
										Trim:           Bool(true),
									},
								},
							},
							ChannelSource: &pf.ChannelSource{
								AccountManagementSettings: &pf.AccountManagementSettings{
									AccountStatusAlgorithm:     String("foo"),
									AccountStatusAttributeName: String("foo"),
									DefaultStatus:              Bool(true),
									FlagComparisonStatus:       Bool(true),
									FlagComparisonValue:        String("foo"),
								},
								BaseDn: String("foo"),
								ChangeDetectionSettings: &pf.ChangeDetectionSettings{
									ChangedUsersAlgorithm:  String("foo"),
									GroupObjectClass:       String("foo"),
									TimeStampAttributeName: String("foo"),
									UserObjectClass:        String("foo"),
									UsnAttributeName:       String("foo"),
								},
								DataSource: &pf.ResourceLink{
									Id:       String("foo"),
									Location: String("foo"),
								},
								GroupMembershipDetection: &pf.GroupMembershipDetection{
									GroupMemberAttributeName:   String("foo"),
									MemberOfGroupAttributeName: String("foo"),
								},
								GroupSourceLocation: &pf.ChannelSourceLocation{
									Filter:       String("foo"),
									GroupDN:      String("foo"),
									NestedSearch: Bool(true),
								},
								GuidAttributeName: String("foo"),
								GuidBinary:        Bool(true),
								UserSourceLocation: &pf.ChannelSourceLocation{
									Filter:       String("foo"),
									GroupDN:      String("foo"),
									NestedSearch: Bool(true),
								},
							},
							MaxThreads: Int(5),
							Name:       String("foo"),
							Timeout:    Int(5),
						},
					},
					CustomSchema: &pf.Schema{
						Attributes: &[]*pf.SchemaAttribute{
							{
								MultiValued:   Bool(true),
								Name:          String("foo"),
								SubAttributes: &[]*string{String("foo")},
								Types:         &[]*string{String("foo")},
							},
						},
						Namespace: String("foo"),
					},
					TargetSettings: &[]*pf.ConfigField{
						{
							//EncryptedValue: String("foo"),
							Inherited: Bool(true),
							Name:      String("foo"),
							Value:     String("foo"),
						},
					},
					Type: String("foo"),
				},
				SpBrowserSso: &pf.SpBrowserSso{
					AdapterMappings: &[]*pf.IdpAdapterAssertionMapping{
						{
							AbortSsoTransactionAsFailSafe: Bool(true),
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
							IdpAdapterRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
							RestrictVirtualEntityIds:   Bool(true),
							RestrictedVirtualEntityIds: &[]*string{String("foo")},
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
					AssertionLifetime: &pf.AssertionLifetime{
						MinutesAfter:  Int(1),
						MinutesBefore: Int(1),
					},
					AttributeContract: &pf.SpBrowserSsoAttributeContract{
						CoreAttributes: &[]*pf.SpBrowserSsoAttribute{
							{
								Name:       String("foo"),
								NameFormat: String("foo"),
							},
						},
						ExtendedAttributes: &[]*pf.SpBrowserSsoAttribute{
							{
								Name:       String("foo"),
								NameFormat: String("foo"),
							},
						},
					},
					AuthenticationPolicyContractAssertionMappings: &[]*pf.AuthenticationPolicyContractAssertionMapping{
						{
							AbortSsoTransactionAsFailSafe: Bool(true),
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
							AuthenticationPolicyContractRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
							RestrictVirtualEntityIds:   Bool(true),
							RestrictedVirtualEntityIds: &[]*string{String("foo")},
						},
					},
					DefaultTargetUrl: String("foo"),
					EnabledProfiles:  &[]*string{String("foo")},
					EncryptionPolicy: &pf.EncryptionPolicy{
						EncryptAssertion:          Bool(true),
						EncryptSloSubjectNameId:   Bool(true),
						EncryptedAttributes:       &[]*string{String("foo")},
						SloSubjectNameIDEncrypted: Bool(true),
					},
					IncomingBindings: &[]*string{String("foo")},
					MessageCustomizations: &[]*pf.ProtocolMessageCustomization{
						{
							ContextName:       String("foo"),
							MessageExpression: String("foo"),
						},
					},
					Protocol:                   String("foo"),
					RequireSignedAuthnRequests: Bool(true),
					SignAssertions:             Bool(true),
					SignResponseAsRequired:     Bool(true),
					SloServiceEndpoints: &[]*pf.SloServiceEndpoint{
						{
							Binding:     String("foo"),
							ResponseUrl: String("foo"),
							Url:         String("foo"),
						},
					},
					SpSamlIdentityMapping:  String("foo"),
					SpWsFedIdentityMapping: String("foo"),
					SsoServiceEndpoints: &[]*pf.SpSsoServiceEndpoint{
						{
							Binding:   String("foo"),
							Index:     Int(1),
							IsDefault: Bool(true),
							Url:       String("foo"),
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
					WsFedTokenType: String("foo"),
					WsTrustVersion: String("foo"),
				},
				Type:             String("SP"),
				VirtualEntityIds: &[]*string{String("foo")},
				WsTrust: &pf.SpWsTrust{
					AbortIfNotFulfilledFromRequest: Bool(true),
					AttributeContract: &pf.SpWsTrustAttributeContract{
						CoreAttributes: &[]*pf.SpWsTrustAttribute{
							{
								Name:      String("foo"),
								Namespace: String("foo"),
							},
						},
						ExtendedAttributes: &[]*pf.SpWsTrustAttribute{
							{
								Name:      String("foo"),
								Namespace: String("foo"),
							},
						},
					},
					DefaultTokenType:      String("foo"),
					EncryptSaml2Assertion: Bool(true),
					GenerateKey:           Bool(true),
					MessageCustomizations: &[]*pf.ProtocolMessageCustomization{
						{
							ContextName:       String("foo"),
							MessageExpression: String("foo"),
						},
					},
					MinutesAfter:           Int(1),
					MinutesBefore:          Int(1),
					OAuthAssertionProfiles: Bool(true),
					PartnerServiceIds:      &[]*string{String("foo")},
					RequestContractRef: &pf.ResourceLink{
						Id:       String("foo"),
						Location: String("foo"),
					},
					TokenProcessorMappings: &[]*pf.IdpTokenProcessorMapping{
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
							IdpTokenProcessorRef: &pf.ResourceLink{
								Id:       String("foo"),
								Location: String("foo"),
							},
							RestrictedVirtualEntityIds: &[]*string{String("foo")},
						},
					},
				},
			},
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("tc:%v", i), func(t *testing.T) {

			resourceSchema := resourcePingFederateIdpSpConnectionResourceSchema()
			resourceLocalData := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{})
			resourcePingFederateIdpSpConnectionResourceReadResult(resourceLocalData, &tc.Resource)

			assert.Equal(t, tc.Resource, *resourcePingFederateIdpSpConnectionResourceReadData(resourceLocalData))
		})
	}
}
