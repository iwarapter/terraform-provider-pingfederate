package pingfederate

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/iwarapter/pingfederate-sdk-go/services/administrativeAccounts"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationApi"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicies"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationPolicyContracts"
	"github.com/iwarapter/pingfederate-sdk-go/services/authenticationSelectors"
	"github.com/iwarapter/pingfederate-sdk-go/services/bulk"
	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesCa"
	"github.com/iwarapter/pingfederate-sdk-go/services/certificatesRevocation"
	"github.com/iwarapter/pingfederate-sdk-go/services/cluster"
	"github.com/iwarapter/pingfederate-sdk-go/services/configArchive"
	"github.com/iwarapter/pingfederate-sdk-go/services/configStore"
	"github.com/iwarapter/pingfederate-sdk-go/services/connectionMetadata"
	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"
	"github.com/iwarapter/pingfederate-sdk-go/services/extendedProperties"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpAdapters"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpConnectors"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpDefaultUrls"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpSpConnections"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpStsRequestParametersContracts"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpToSpAdapterMapping"
	"github.com/iwarapter/pingfederate-sdk-go/services/idpTokenProcessors"
	"github.com/iwarapter/pingfederate-sdk-go/services/kerberosRealms"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairs"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsOauthOpenIdConnect"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSigning"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslClient"
	"github.com/iwarapter/pingfederate-sdk-go/services/keyPairsSslServer"
	"github.com/iwarapter/pingfederate-sdk-go/services/license"
	"github.com/iwarapter/pingfederate-sdk-go/services/localIdentityIdentityProfiles"
	"github.com/iwarapter/pingfederate-sdk-go/services/metadataUrls"
	"github.com/iwarapter/pingfederate-sdk-go/services/notificationPublishers"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenManagers"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAccessTokenMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthServerSettings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthAuthenticationPolicyContractMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthCibaServerPolicy"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientRegistrationPolicies"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClientSettings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthClients"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthIdpAdapterMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthOpenIdConnect"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthOutOfBandAuthPlugins"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthResourceOwnerCredentialsMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthTokenExchangeGenerator"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthTokenExchangeProcessor"
	"github.com/iwarapter/pingfederate-sdk-go/services/oauthTokenExchangeTokenGeneratorMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/passwordCredentialValidators"
	"github.com/iwarapter/pingfederate-sdk-go/services/redirectValidation"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"
	"github.com/iwarapter/pingfederate-sdk-go/services/session"
	"github.com/iwarapter/pingfederate-sdk-go/services/spAdapters"
	"github.com/iwarapter/pingfederate-sdk-go/services/spAuthenticationPolicyContractMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/spDefaultUrls"
	"github.com/iwarapter/pingfederate-sdk-go/services/spIdpConnections"
	"github.com/iwarapter/pingfederate-sdk-go/services/spTargetUrlMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/spTokenGenerators"
	"github.com/iwarapter/pingfederate-sdk-go/services/tokenProcessorToTokenGeneratorMappings"
	"github.com/iwarapter/pingfederate-sdk-go/services/version"
	"github.com/iwarapter/pingfederate-sdk-go/services/virtualHostNames"
)

type Config struct {
	Username string
	Password string
	Context  string
	BaseURL  string
}

type pfClient struct {
	AdministrativeAccounts                    administrativeAccounts.AdministrativeAccountsAPI
	AuthenticationApi                         authenticationApi.AuthenticationApiAPI
	AuthenticationPolicies                    authenticationPolicies.AuthenticationPoliciesAPI
	AuthenticationPolicyContracts             authenticationPolicyContracts.AuthenticationPolicyContractsAPI
	AuthenticationSelectors                   authenticationSelectors.AuthenticationSelectorsAPI
	Bulk                                      bulk.BulkAPI
	CertificatesCa                            certificatesCa.CertificatesCaAPI
	CertificatesRevocation                    certificatesRevocation.CertificatesRevocationAPI
	Cluster                                   cluster.ClusterAPI
	ConfigArchive                             configArchive.ConfigArchiveAPI
	ConfigStore                               configStore.ConfigStoreAPI
	ConnectionMetadata                        connectionMetadata.ConnectionMetadataAPI
	DataStores                                dataStores.DataStoresAPI
	ExtendedProperties                        extendedProperties.ExtendedPropertiesAPI
	IdpAdapters                               idpAdapters.IdpAdaptersAPI
	IdpConnectors                             idpConnectors.IdpConnectorsAPI
	IdpDefaultUrls                            idpDefaultUrls.IdpDefaultUrlsAPI
	IdpSpConnections                          idpSpConnections.IdpSpConnectionsAPI
	IdpStsRequestParametersContracts          idpStsRequestParametersContracts.IdpStsRequestParametersContractsAPI
	IdpToSpAdapterMapping                     idpToSpAdapterMapping.IdpToSpAdapterMappingAPI
	IdpTokenProcessors                        idpTokenProcessors.IdpTokenProcessorsAPI
	KerberosRealms                            kerberosRealms.KerberosRealmsAPI
	KeyPairs                                  keyPairs.KeyPairsAPI
	KeyPairsOauthOpenIdConnect                keyPairsOauthOpenIdConnect.KeyPairsOauthOpenIdConnectAPI
	KeyPairsSigning                           keyPairsSigning.KeyPairsSigningAPI
	KeyPairsSslClient                         keyPairsSslClient.KeyPairsSslClientAPI
	KeyPairsSslServer                         keyPairsSslServer.KeyPairsSslServerAPI
	License                                   license.LicenseAPI
	LocalIdentityIdentityProfiles             localIdentityIdentityProfiles.LocalIdentityIdentityProfilesAPI
	MetadataUrls                              metadataUrls.MetadataUrlsAPI
	NotificationPublishers                    notificationPublishers.NotificationPublishersAPI
	OauthAccessTokenManagers                  oauthAccessTokenManagers.OauthAccessTokenManagersAPI
	OauthAccessTokenMappings                  oauthAccessTokenMappings.OauthAccessTokenMappingsAPI
	OauthAuthServerSettings                   oauthAuthServerSettings.OauthAuthServerSettingsAPI
	OauthAuthenticationPolicyContractMappings oauthAuthenticationPolicyContractMappings.OauthAuthenticationPolicyContractMappingsAPI
	OauthCibaServerPolicy                     oauthCibaServerPolicy.OauthCibaServerPolicyAPI
	OauthClientRegistrationPolicies           oauthClientRegistrationPolicies.OauthClientRegistrationPoliciesAPI
	OauthClientSettings                       oauthClientSettings.OauthClientSettingsAPI
	OauthClients                              oauthClients.OauthClientsAPI
	OauthIdpAdapterMappings                   oauthIdpAdapterMappings.OauthIdpAdapterMappingsAPI
	OauthOpenIdConnect                        oauthOpenIdConnect.OauthOpenIdConnectAPI
	OauthOutOfBandAuthPlugins                 oauthOutOfBandAuthPlugins.OauthOutOfBandAuthPluginsAPI
	OauthResourceOwnerCredentialsMappings     oauthResourceOwnerCredentialsMappings.OauthResourceOwnerCredentialsMappingsAPI
	OauthTokenExchangeGenerator               oauthTokenExchangeGenerator.OauthTokenExchangeGeneratorAPI
	OauthTokenExchangeProcessor               oauthTokenExchangeProcessor.OauthTokenExchangeProcessorAPI
	OauthTokenExchangeTokenGeneratorMappings  oauthTokenExchangeTokenGeneratorMappings.OauthTokenExchangeTokenGeneratorMappingsAPI
	PasswordCredentialValidators              passwordCredentialValidators.PasswordCredentialValidatorsAPI
	RedirectValidation                        redirectValidation.RedirectValidationAPI
	ServerSettings                            serverSettings.ServerSettingsAPI
	Session                                   session.SessionAPI
	SpAdapters                                spAdapters.SpAdaptersAPI
	SpAuthenticationPolicyContractMappings    spAuthenticationPolicyContractMappings.SpAuthenticationPolicyContractMappingsAPI
	SpDefaultUrls                             spDefaultUrls.SpDefaultUrlsAPI
	SpIdpConnections                          spIdpConnections.SpIdpConnectionsAPI
	SpTargetUrlMappings                       spTargetUrlMappings.SpTargetUrlMappingsAPI
	SpTokenGenerators                         spTokenGenerators.SpTokenGeneratorsAPI
	TokenProcessorToTokenGeneratorMappings    tokenProcessorToTokenGeneratorMappings.TokenProcessorToTokenGeneratorMappingsAPI
	Version                                   version.VersionAPI
	VirtualHostNames                          virtualHostNames.VirtualHostNamesAPI
}

// Client configures and returns a fully initialized PAClient
func (c *Config) Client() (interface{}, diag.Diagnostics) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)

	client := pfClient{
		AdministrativeAccounts:           administrativeAccounts.New(c.Username, c.Password, url, c.Context, nil),
		AuthenticationApi:                authenticationApi.New(c.Username, c.Password, url, c.Context, nil),
		AuthenticationPolicies:           authenticationPolicies.New(c.Username, c.Password, url, c.Context, nil),
		AuthenticationPolicyContracts:    authenticationPolicyContracts.New(c.Username, c.Password, url, c.Context, nil),
		AuthenticationSelectors:          authenticationSelectors.New(c.Username, c.Password, url, c.Context, nil),
		Bulk:                             bulk.New(c.Username, c.Password, url, c.Context, nil),
		CertificatesCa:                   certificatesCa.New(c.Username, c.Password, url, c.Context, nil),
		CertificatesRevocation:           certificatesRevocation.New(c.Username, c.Password, url, c.Context, nil),
		Cluster:                          cluster.New(c.Username, c.Password, url, c.Context, nil),
		ConfigArchive:                    configArchive.New(c.Username, c.Password, url, c.Context, nil),
		ConfigStore:                      configStore.New(c.Username, c.Password, url, c.Context, nil),
		ConnectionMetadata:               connectionMetadata.New(c.Username, c.Password, url, c.Context, nil),
		DataStores:                       dataStores.New(c.Username, c.Password, url, c.Context, nil),
		ExtendedProperties:               extendedProperties.New(c.Username, c.Password, url, c.Context, nil),
		IdpAdapters:                      idpAdapters.New(c.Username, c.Password, url, c.Context, nil),
		IdpConnectors:                    idpConnectors.New(c.Username, c.Password, url, c.Context, nil),
		IdpDefaultUrls:                   idpDefaultUrls.New(c.Username, c.Password, url, c.Context, nil),
		IdpSpConnections:                 idpSpConnections.New(c.Username, c.Password, url, c.Context, nil),
		IdpStsRequestParametersContracts: idpStsRequestParametersContracts.New(c.Username, c.Password, url, c.Context, nil),
		IdpToSpAdapterMapping:            idpToSpAdapterMapping.New(c.Username, c.Password, url, c.Context, nil),
		IdpTokenProcessors:               idpTokenProcessors.New(c.Username, c.Password, url, c.Context, nil),
		KerberosRealms:                   kerberosRealms.New(c.Username, c.Password, url, c.Context, nil),
		KeyPairs:                         keyPairs.New(c.Username, c.Password, url, c.Context, nil),
		KeyPairsOauthOpenIdConnect:       keyPairsOauthOpenIdConnect.New(c.Username, c.Password, url, c.Context, nil),
		KeyPairsSigning:                  keyPairsSigning.New(c.Username, c.Password, url, c.Context, nil),
		KeyPairsSslClient:                keyPairsSslClient.New(c.Username, c.Password, url, c.Context, nil),
		KeyPairsSslServer:                keyPairsSslServer.New(c.Username, c.Password, url, c.Context, nil),
		License:                          license.New(c.Username, c.Password, url, c.Context, nil),
		LocalIdentityIdentityProfiles:    localIdentityIdentityProfiles.New(c.Username, c.Password, url, c.Context, nil),
		MetadataUrls:                     metadataUrls.New(c.Username, c.Password, url, c.Context, nil),
		NotificationPublishers:           notificationPublishers.New(c.Username, c.Password, url, c.Context, nil),
		OauthAccessTokenManagers:         oauthAccessTokenManagers.New(c.Username, c.Password, url, c.Context, nil),
		OauthAccessTokenMappings:         oauthAccessTokenMappings.New(c.Username, c.Password, url, c.Context, nil),
		OauthAuthServerSettings:          oauthAuthServerSettings.New(c.Username, c.Password, url, c.Context, nil),
		OauthAuthenticationPolicyContractMappings: oauthAuthenticationPolicyContractMappings.New(c.Username, c.Password, url, c.Context, nil),
		OauthCibaServerPolicy:                     oauthCibaServerPolicy.New(c.Username, c.Password, url, c.Context, nil),
		OauthClientRegistrationPolicies:           oauthClientRegistrationPolicies.New(c.Username, c.Password, url, c.Context, nil),
		OauthClientSettings:                       oauthClientSettings.New(c.Username, c.Password, url, c.Context, nil),
		OauthClients:                              oauthClients.New(c.Username, c.Password, url, c.Context, nil),
		OauthIdpAdapterMappings:                   oauthIdpAdapterMappings.New(c.Username, c.Password, url, c.Context, nil),
		OauthOpenIdConnect:                        oauthOpenIdConnect.New(c.Username, c.Password, url, c.Context, nil),
		OauthOutOfBandAuthPlugins:                 oauthOutOfBandAuthPlugins.New(c.Username, c.Password, url, c.Context, nil),
		OauthResourceOwnerCredentialsMappings:     oauthResourceOwnerCredentialsMappings.New(c.Username, c.Password, url, c.Context, nil),
		OauthTokenExchangeGenerator:               oauthTokenExchangeGenerator.New(c.Username, c.Password, url, c.Context, nil),
		OauthTokenExchangeProcessor:               oauthTokenExchangeProcessor.New(c.Username, c.Password, url, c.Context, nil),
		OauthTokenExchangeTokenGeneratorMappings:  oauthTokenExchangeTokenGeneratorMappings.New(c.Username, c.Password, url, c.Context, nil),
		PasswordCredentialValidators:              passwordCredentialValidators.New(c.Username, c.Password, url, c.Context, nil),
		RedirectValidation:                        redirectValidation.New(c.Username, c.Password, url, c.Context, nil),
		ServerSettings:                            serverSettings.New(c.Username, c.Password, url, c.Context, nil),
		Session:                                   session.New(c.Username, c.Password, url, c.Context, nil),
		SpAdapters:                                spAdapters.New(c.Username, c.Password, url, c.Context, nil),
		SpAuthenticationPolicyContractMappings:    spAuthenticationPolicyContractMappings.New(c.Username, c.Password, url, c.Context, nil),
		SpDefaultUrls:                             spDefaultUrls.New(c.Username, c.Password, url, c.Context, nil),
		SpIdpConnections:                          spIdpConnections.New(c.Username, c.Password, url, c.Context, nil),
		SpTargetUrlMappings:                       spTargetUrlMappings.New(c.Username, c.Password, url, c.Context, nil),
		SpTokenGenerators:                         spTokenGenerators.New(c.Username, c.Password, url, c.Context, nil),
		TokenProcessorToTokenGeneratorMappings:    tokenProcessorToTokenGeneratorMappings.New(c.Username, c.Password, url, c.Context, nil),
		Version:                                   version.New(c.Username, c.Password, url, c.Context, nil),
		VirtualHostNames:                          virtualHostNames.New(c.Username, c.Password, url, c.Context, nil),
	}
	//if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
	//	client.LogDebug = true
	//}
	return client, nil
}
