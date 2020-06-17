package keyPairsOauthOpenIdConnect

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "KeyPairsOauthOpenIdConnect"
)

type KeyPairsOauthOpenIdConnectService struct {
	*client.PfClient
}

// New creates a new instance of the KeyPairsOauthOpenIdConnectService client.
func New(cfg *config.Config) *KeyPairsOauthOpenIdConnectService {

	return &KeyPairsOauthOpenIdConnectService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a KeyPairsOauthOpenIdConnect operation
func (c *KeyPairsOauthOpenIdConnectService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetOauthOidcKeysSettings - Retrieve OAuth/Open ID Connect key settings.
//RequestType: GET
//Input:
func (s *KeyPairsOauthOpenIdConnectService) GetOauthOidcKeysSettings() (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	op := &request.Operation{
		Name:       "GetOauthOidcKeysSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OAuthOidcKeysSettings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOAuthOidcKeysSettings - Update OAuth/Open ID Connect key settings.
//RequestType: PUT
//Input: input *UpdateOAuthOidcKeysSettingsInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	op := &request.Operation{
		Name:       "UpdateOAuthOidcKeysSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OAuthOidcKeysSettings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateOAuthOidcKeysSettingsInput struct {
	Body models.OAuthOidcKeysSettings
}
