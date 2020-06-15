package keyPairsOauthOpenIdConnect

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsOauthOpenIdConnectService struct {
	Client *client.PfClient
}

// New creates a new instance of the KeyPairsOauthOpenIdConnectService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *KeyPairsOauthOpenIdConnectService {

	return &KeyPairsOauthOpenIdConnectService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetOauthOidcKeysSettings - Retrieve OAuth/Open ID Connect key settings.
//RequestType: GET
//Input:
func (s *KeyPairsOauthOpenIdConnectService) GetOauthOidcKeysSettings() (result *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateOAuthOidcKeysSettings - Update OAuth/Open ID Connect key settings.
//RequestType: PUT
//Input: input *UpdateOAuthOidcKeysSettingsInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (result *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateOAuthOidcKeysSettingsInput struct {
	Body models.OAuthOidcKeysSettings
}
