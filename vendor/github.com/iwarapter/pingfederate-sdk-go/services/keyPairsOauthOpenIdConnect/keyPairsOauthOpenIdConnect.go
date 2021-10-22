package keyPairsOauthOpenIdConnect

import (
	"context"
	"net/http"
	"strings"

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

//GetOauthOidcKeysSettings - Retrieve OAuth/OpenID Connect key settings.
//RequestType: GET
//Input:
func (s *KeyPairsOauthOpenIdConnectService) GetOauthOidcKeysSettings() (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	return s.GetOauthOidcKeysSettingsWithContext(context.Background())
}

//GetOauthOidcKeysSettingsWithContext - Retrieve OAuth/OpenID Connect key settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *KeyPairsOauthOpenIdConnectService) GetOauthOidcKeysSettingsWithContext(ctx context.Context) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	op := &request.Operation{
		Name:       "GetOauthOidcKeysSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.OAuthOidcKeysSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateOAuthOidcKeysSettings - Update OAuth/OpenID Connect key settings.
//RequestType: PUT
//Input: input *UpdateOAuthOidcKeysSettingsInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateOAuthOidcKeysSettings(input *UpdateOAuthOidcKeysSettingsInput) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	return s.UpdateOAuthOidcKeysSettingsWithContext(context.Background(), input)
}

//UpdateOAuthOidcKeysSettingsWithContext - Update OAuth/OpenID Connect key settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateOAuthOidcKeysSettingsInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateOAuthOidcKeysSettingsWithContext(ctx context.Context, input *UpdateOAuthOidcKeysSettingsInput) (output *models.OAuthOidcKeysSettings, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect"
	op := &request.Operation{
		Name:       "UpdateOAuthOidcKeysSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.OAuthOidcKeysSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKeySets - Retrieve OAuth/OpenID Connect additional signing key sets.
//RequestType: GET
//Input:
func (s *KeyPairsOauthOpenIdConnectService) GetKeySets() (output *models.AdditionalKeySets, resp *http.Response, err error) {
	return s.GetKeySetsWithContext(context.Background())
}

//GetKeySetsWithContext - Retrieve OAuth/OpenID Connect additional signing key sets.
//RequestType: GET
//Input: ctx context.Context,
func (s *KeyPairsOauthOpenIdConnectService) GetKeySetsWithContext(ctx context.Context) (output *models.AdditionalKeySets, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect/additionalKeySets"
	op := &request.Operation{
		Name:       "GetKeySets",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdditionalKeySets{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateKeySet - Create a new OAuth/OpenID Connect additional signing key set.
//RequestType: POST
//Input: input *CreateKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) CreateKeySet(input *CreateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	return s.CreateKeySetWithContext(context.Background(), input)
}

//CreateKeySetWithContext - Create a new OAuth/OpenID Connect additional signing key set.
//RequestType: POST
//Input: ctx context.Context, input *CreateKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) CreateKeySetWithContext(ctx context.Context, input *CreateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect/additionalKeySets"
	op := &request.Operation{
		Name:       "CreateKeySet",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AdditionalKeySet{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetKeySet - Retrieve an OAuth/OpenID Connect additional signing key set.
//RequestType: GET
//Input: input *GetKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) GetKeySet(input *GetKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	return s.GetKeySetWithContext(context.Background(), input)
}

//GetKeySetWithContext - Retrieve an OAuth/OpenID Connect additional signing key set.
//RequestType: GET
//Input: ctx context.Context, input *GetKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) GetKeySetWithContext(ctx context.Context, input *GetKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect/additionalKeySets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetKeySet",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdditionalKeySet{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateKeySet - Update an existing OAuth/OpenID Connect additional signing key set.
//RequestType: PUT
//Input: input *UpdateKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateKeySet(input *UpdateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	return s.UpdateKeySetWithContext(context.Background(), input)
}

//UpdateKeySetWithContext - Update an existing OAuth/OpenID Connect additional signing key set.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) UpdateKeySetWithContext(ctx context.Context, input *UpdateKeySetInput) (output *models.AdditionalKeySet, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect/additionalKeySets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateKeySet",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AdditionalKeySet{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteKeySet - Delete an existing OAuth/OpenID Connect additional signing key set.
//RequestType: DELETE
//Input: input *DeleteKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) DeleteKeySet(input *DeleteKeySetInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteKeySetWithContext(context.Background(), input)
}

//DeleteKeySetWithContext - Delete an existing OAuth/OpenID Connect additional signing key set.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteKeySetInput
func (s *KeyPairsOauthOpenIdConnectService) DeleteKeySetWithContext(ctx context.Context, input *DeleteKeySetInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/oauthOpenIdConnect/additionalKeySets/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteKeySet",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateKeySetInput struct {
	Body models.AdditionalKeySet
}

type DeleteKeySetInput struct {
	Id string
}

type GetKeySetInput struct {
	Id string
}

type UpdateKeySetInput struct {
	Body models.AdditionalKeySet
	Id   string
}

type UpdateOAuthOidcKeysSettingsInput struct {
	Body models.OAuthOidcKeysSettings
}
