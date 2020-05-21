package pingfederate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"runtime"
)

const logReqMsg = `DEBUG: Request %s Details:
---[ REQUEST ]--------------------------------------
%s
-----------------------------------------------------`
const logRespMsg = `DEBUG: Response %s Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`

type PfClient struct {
	Username   string
	Password   string
	BaseURL    *url.URL
	Context    string
	LogDebug   bool
	httpClient *http.Client

	AdministrativeAccounts                    *AdministrativeAccountsService
	AuthenticationApi                         *AuthenticationApiService
	AuthenticationPolicies                    *AuthenticationPoliciesService
	AuthenticationPolicyContracts             *AuthenticationPolicyContractsService
	AuthenticationSelectors                   *AuthenticationSelectorsService
	Bulk                                      *BulkService
	CertificatesCa                            *CertificatesCaService
	CertificatesRevocation                    *CertificatesRevocationService
	Cluster                                   *ClusterService
	ConfigArchive                             *ConfigArchiveService
	ConfigStore                               *ConfigStoreService
	ConnectionMetadata                        *ConnectionMetadataService
	DataStores                                *DataStoresService
	ExtendedProperties                        *ExtendedPropertiesService
	IdpAdapters                               *IdpAdaptersService
	IdpConnectors                             *IdpConnectorsService
	IdpDefaultUrls                            *IdpDefaultUrlsService
	IdpSpConnections                          *IdpSpConnectionsService
	IdpStsRequestParametersContracts          *IdpStsRequestParametersContractsService
	IdpToSpAdapterMapping                     *IdpToSpAdapterMappingService
	IdpTokenProcessors                        *IdpTokenProcessorsService
	KerberosRealms                            *KerberosRealmsService
	KeyPairs                                  *KeyPairsService
	KeyPairsOauthOpenIdConnect                *KeyPairsOauthOpenIdConnectService
	KeyPairsSigning                           *KeyPairsSigningService
	KeyPairsSslClient                         *KeyPairsSslClientService
	KeyPairsSslServer                         *KeyPairsSslServerService
	License                                   *LicenseService
	LocalIdentityIdentityProfiles             *LocalIdentityIdentityProfilesService
	MetadataUrls                              *MetadataUrlsService
	NotificationPublishers                    *NotificationPublishersService
	OauthAccessTokenManagers                  *OauthAccessTokenManagersService
	OauthAccessTokenMappings                  *OauthAccessTokenMappingsService
	OauthAuthServerSettings                   *OauthAuthServerSettingsService
	OauthAuthenticationPolicyContractMappings *OauthAuthenticationPolicyContractMappingsService
	OauthCibaServerPolicy                     *OauthCibaServerPolicyService
	OauthClientRegistrationPolicies           *OauthClientRegistrationPoliciesService
	OauthClientSettings                       *OauthClientSettingsService
	OauthClients                              *OauthClientsService
	OauthIdpAdapterMappings                   *OauthIdpAdapterMappingsService
	OauthOpenIdConnect                        *OauthOpenIdConnectService
	OauthOutOfBandAuthPlugins                 *OauthOutOfBandAuthPluginsService
	OauthResourceOwnerCredentialsMappings     *OauthResourceOwnerCredentialsMappingsService
	OauthTokenExchangeGenerator               *OauthTokenExchangeGeneratorService
	OauthTokenExchangeProcessor               *OauthTokenExchangeProcessorService
	OauthTokenExchangeTokenGeneratorMappings  *OauthTokenExchangeTokenGeneratorMappingsService
	PasswordCredentialValidators              *PasswordCredentialValidatorsService
	RedirectValidation                        *RedirectValidationService
	ServerSettings                            *ServerSettingsService
	Session                                   *SessionService
	SpAdapters                                *SpAdaptersService
	SpAuthenticationPolicyContractMappings    *SpAuthenticationPolicyContractMappingsService
	SpDefaultUrls                             *SpDefaultUrlsService
	SpIdpConnections                          *SpIdpConnectionsService
	SpTargetUrlMappings                       *SpTargetUrlMappingsService
	SpTokenGenerators                         *SpTokenGeneratorsService
	TokenProcessorToTokenGeneratorMappings    *TokenProcessorToTokenGeneratorMappingsService
	Version                                   *VersionService
	VirtualHostNames                          *VirtualHostNamesService
}

type service struct {
	client *PfClient
}

func NewClient(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *PfClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &PfClient{httpClient: httpClient}
	c.Username = username
	c.Password = password
	c.BaseURL = baseUrl
	c.Context = context

	c.AdministrativeAccounts = &AdministrativeAccountsService{client: c}
	c.AuthenticationApi = &AuthenticationApiService{client: c}
	c.AuthenticationPolicies = &AuthenticationPoliciesService{client: c}
	c.AuthenticationPolicyContracts = &AuthenticationPolicyContractsService{client: c}
	c.AuthenticationSelectors = &AuthenticationSelectorsService{client: c}
	c.Bulk = &BulkService{client: c}
	c.CertificatesCa = &CertificatesCaService{client: c}
	c.CertificatesRevocation = &CertificatesRevocationService{client: c}
	c.Cluster = &ClusterService{client: c}
	c.ConfigArchive = &ConfigArchiveService{client: c}
	c.ConfigStore = &ConfigStoreService{client: c}
	c.ConnectionMetadata = &ConnectionMetadataService{client: c}
	c.DataStores = &DataStoresService{client: c}
	c.ExtendedProperties = &ExtendedPropertiesService{client: c}
	c.IdpAdapters = &IdpAdaptersService{client: c}
	c.IdpConnectors = &IdpConnectorsService{client: c}
	c.IdpDefaultUrls = &IdpDefaultUrlsService{client: c}
	c.IdpSpConnections = &IdpSpConnectionsService{client: c}
	c.IdpStsRequestParametersContracts = &IdpStsRequestParametersContractsService{client: c}
	c.IdpToSpAdapterMapping = &IdpToSpAdapterMappingService{client: c}
	c.IdpTokenProcessors = &IdpTokenProcessorsService{client: c}
	c.KerberosRealms = &KerberosRealmsService{client: c}
	c.KeyPairs = &KeyPairsService{client: c}
	c.KeyPairsOauthOpenIdConnect = &KeyPairsOauthOpenIdConnectService{client: c}
	c.KeyPairsSigning = &KeyPairsSigningService{client: c}
	c.KeyPairsSslClient = &KeyPairsSslClientService{client: c}
	c.KeyPairsSslServer = &KeyPairsSslServerService{client: c}
	c.License = &LicenseService{client: c}
	c.LocalIdentityIdentityProfiles = &LocalIdentityIdentityProfilesService{client: c}
	c.MetadataUrls = &MetadataUrlsService{client: c}
	c.NotificationPublishers = &NotificationPublishersService{client: c}
	c.OauthAccessTokenManagers = &OauthAccessTokenManagersService{client: c}
	c.OauthAccessTokenMappings = &OauthAccessTokenMappingsService{client: c}
	c.OauthAuthServerSettings = &OauthAuthServerSettingsService{client: c}
	c.OauthAuthenticationPolicyContractMappings = &OauthAuthenticationPolicyContractMappingsService{client: c}
	c.OauthCibaServerPolicy = &OauthCibaServerPolicyService{client: c}
	c.OauthClientRegistrationPolicies = &OauthClientRegistrationPoliciesService{client: c}
	c.OauthClientSettings = &OauthClientSettingsService{client: c}
	c.OauthClients = &OauthClientsService{client: c}
	c.OauthIdpAdapterMappings = &OauthIdpAdapterMappingsService{client: c}
	c.OauthOpenIdConnect = &OauthOpenIdConnectService{client: c}
	c.OauthOutOfBandAuthPlugins = &OauthOutOfBandAuthPluginsService{client: c}
	c.OauthResourceOwnerCredentialsMappings = &OauthResourceOwnerCredentialsMappingsService{client: c}
	c.OauthTokenExchangeGenerator = &OauthTokenExchangeGeneratorService{client: c}
	c.OauthTokenExchangeProcessor = &OauthTokenExchangeProcessorService{client: c}
	c.OauthTokenExchangeTokenGeneratorMappings = &OauthTokenExchangeTokenGeneratorMappingsService{client: c}
	c.PasswordCredentialValidators = &PasswordCredentialValidatorsService{client: c}
	c.RedirectValidation = &RedirectValidationService{client: c}
	c.ServerSettings = &ServerSettingsService{client: c}
	c.Session = &SessionService{client: c}
	c.SpAdapters = &SpAdaptersService{client: c}
	c.SpAuthenticationPolicyContractMappings = &SpAuthenticationPolicyContractMappingsService{client: c}
	c.SpDefaultUrls = &SpDefaultUrlsService{client: c}
	c.SpIdpConnections = &SpIdpConnectionsService{client: c}
	c.SpTargetUrlMappings = &SpTargetUrlMappingsService{client: c}
	c.SpTokenGenerators = &SpTokenGeneratorsService{client: c}
	c.TokenProcessorToTokenGeneratorMappings = &TokenProcessorToTokenGeneratorMappingsService{client: c}
	c.Version = &VersionService{client: c}
	c.VirtualHostNames = &VirtualHostNamesService{client: c}
	return c
}

func (c *PfClient) newRequest(method string, path *url.URL, body interface{}) (*http.Request, error) {
	u := c.BaseURL.ResolveReference(path)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("X-Xsrf-Header", "pingfederate")
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s (%s; %s; %s)", SDKName, SDKVersion, runtime.Version(), runtime.GOOS, runtime.GOARCH))
	return req, nil
}

func (c *PfClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.LogDebug {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logReqMsg, "pingaccess-sdk-go", string(requestDump))
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if c.LogDebug {
		responseDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf(logRespMsg, "pingaccess-sdk-go", string(responseDump))
	}
	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse. Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := ApiResult{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return &PingFederateError{
				ApiResult: ApiResult{
					Message: String("Unable to parse error response: " + string(data)),
				},
			}
		}
	}

	return &PingFederateError{
		ApiResult: errorResponse,
	}
}

// PingFederateError occurs when PingFederate returns a non 2XX response
type PingFederateError struct {
	ApiResult ApiResult
}

func (r *PingFederateError) Error() (message string) {
	if r.ApiResult.Message != nil {
		message = *r.ApiResult.Message
	}
	if r.ApiResult.ValidationErrors != nil && len(*r.ApiResult.ValidationErrors) > 0 {
		for _, v := range *r.ApiResult.ValidationErrors {
			if v.Message != nil {
				message = fmt.Sprintf("%s\n%s", message, *v.Message)
			}
		}
	}
	return
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
