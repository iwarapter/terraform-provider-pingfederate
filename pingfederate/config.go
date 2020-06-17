package pingfederate

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

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
	/* #nosec */
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	url, _ := url.Parse(c.BaseURL)

	cfg := config.NewConfig().WithEndpoint(url.String() + c.Context).WithUsername(c.Username).WithPassword(c.Password)

	if os.Getenv("TF_LOG") == "DEBUG" || os.Getenv("TF_LOG") == "TRACE" {
		cfg.WithDebug(true)
	}

	client := pfClient{
		AdministrativeAccounts:           administrativeAccounts.New(cfg),
		AuthenticationApi:                authenticationApi.New(cfg),
		AuthenticationPolicies:           authenticationPolicies.New(cfg),
		AuthenticationPolicyContracts:    authenticationPolicyContracts.New(cfg),
		AuthenticationSelectors:          authenticationSelectors.New(cfg),
		Bulk:                             bulk.New(cfg),
		CertificatesCa:                   certificatesCa.New(cfg),
		CertificatesRevocation:           certificatesRevocation.New(cfg),
		Cluster:                          cluster.New(cfg),
		ConfigArchive:                    configArchive.New(cfg),
		ConfigStore:                      configStore.New(cfg),
		ConnectionMetadata:               connectionMetadata.New(cfg),
		DataStores:                       dataStores.New(cfg),
		ExtendedProperties:               extendedProperties.New(cfg),
		IdpAdapters:                      idpAdapters.New(cfg),
		IdpConnectors:                    idpConnectors.New(cfg),
		IdpDefaultUrls:                   idpDefaultUrls.New(cfg),
		IdpSpConnections:                 idpSpConnections.New(cfg),
		IdpStsRequestParametersContracts: idpStsRequestParametersContracts.New(cfg),
		IdpToSpAdapterMapping:            idpToSpAdapterMapping.New(cfg),
		IdpTokenProcessors:               idpTokenProcessors.New(cfg),
		KerberosRealms:                   kerberosRealms.New(cfg),
		KeyPairs:                         keyPairs.New(cfg),
		KeyPairsOauthOpenIdConnect:       keyPairsOauthOpenIdConnect.New(cfg),
		KeyPairsSigning:                  keyPairsSigning.New(cfg),
		KeyPairsSslClient:                keyPairsSslClient.New(cfg),
		KeyPairsSslServer:                keyPairsSslServer.New(cfg),
		License:                          license.New(cfg),
		LocalIdentityIdentityProfiles:    localIdentityIdentityProfiles.New(cfg),
		MetadataUrls:                     metadataUrls.New(cfg),
		NotificationPublishers:           notificationPublishers.New(cfg),
		OauthAccessTokenManagers:         oauthAccessTokenManagers.New(cfg),
		OauthAccessTokenMappings:         oauthAccessTokenMappings.New(cfg),
		OauthAuthServerSettings:          oauthAuthServerSettings.New(cfg),
		OauthAuthenticationPolicyContractMappings: oauthAuthenticationPolicyContractMappings.New(cfg),
		OauthCibaServerPolicy:                     oauthCibaServerPolicy.New(cfg),
		OauthClientRegistrationPolicies:           oauthClientRegistrationPolicies.New(cfg),
		OauthClientSettings:                       oauthClientSettings.New(cfg),
		OauthClients:                              oauthClients.New(cfg),
		OauthIdpAdapterMappings:                   oauthIdpAdapterMappings.New(cfg),
		OauthOpenIdConnect:                        oauthOpenIdConnect.New(cfg),
		OauthOutOfBandAuthPlugins:                 oauthOutOfBandAuthPlugins.New(cfg),
		OauthResourceOwnerCredentialsMappings:     oauthResourceOwnerCredentialsMappings.New(cfg),
		OauthTokenExchangeGenerator:               oauthTokenExchangeGenerator.New(cfg),
		OauthTokenExchangeProcessor:               oauthTokenExchangeProcessor.New(cfg),
		OauthTokenExchangeTokenGeneratorMappings:  oauthTokenExchangeTokenGeneratorMappings.New(cfg),
		PasswordCredentialValidators:              passwordCredentialValidators.New(cfg),
		RedirectValidation:                        redirectValidation.New(cfg),
		ServerSettings:                            serverSettings.New(cfg),
		Session:                                   session.New(cfg),
		SpAdapters:                                spAdapters.New(cfg),
		SpAuthenticationPolicyContractMappings:    spAuthenticationPolicyContractMappings.New(cfg),
		SpDefaultUrls:                             spDefaultUrls.New(cfg),
		SpIdpConnections:                          spIdpConnections.New(cfg),
		SpTargetUrlMappings:                       spTargetUrlMappings.New(cfg),
		SpTokenGenerators:                         spTokenGenerators.New(cfg),
		TokenProcessorToTokenGeneratorMappings:    tokenProcessorToTokenGeneratorMappings.New(cfg),
		Version:                                   version.New(cfg),
		VirtualHostNames:                          virtualHostNames.New(cfg),
	}
	return client, nil
}
