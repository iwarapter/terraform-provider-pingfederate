package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type KeyPairsOauthOpenIdConnectService service

//GetOauthOidcKeysSettings - Retrieve OAuth/Open ID Connect key settings.
//RequestType: GET
//Input:
func (s *KeyPairsOauthOpenIdConnectService) GetOauthOidcKeysSettings() (result *OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateOAuthOidcKeysSettings - Update OAuth/Open ID Connect key settings.
//RequestType: PUT
//Input: input *UpdateOAuthOidcKeysSettingsInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (result *OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
