package spIdpConnections

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpIdpConnectionsService struct {
	Client *client.PfClient
}

// New creates a new instance of the SpIdpConnectionsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *SpIdpConnectionsService {

	return &SpIdpConnectionsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetConnections - Get list of IdP connections.
//RequestType: GET
//Input: input *GetConnectionsInput
func (s *SpIdpConnectionsService) GetConnections(input *GetConnectionsInput) (result *models.IdpConnections, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	q := rel.Query()
	if input.EntityId != "" {
		q.Set("entityId", input.EntityId)
	}
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
	}
	rel.RawQuery = q.Encode()
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

//CreateConnection - Create a new IdP connection.
//RequestType: POST
//Input: input *CreateConnectionInput
func (s *SpIdpConnectionsService) CreateConnection(input *CreateConnectionInput) (result *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetConnection - Find IdP connection by ID.
//RequestType: GET
//Input: input *GetConnectionInput
func (s *SpIdpConnectionsService) GetConnection(input *GetConnectionInput) (result *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateConnection - Update an IdP connection.
//RequestType: PUT
//Input: input *UpdateConnectionInput
func (s *SpIdpConnectionsService) UpdateConnection(input *UpdateConnectionInput) (result *models.IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteConnection - Delete an IdP connection.
//RequestType: DELETE
//Input: input *DeleteConnectionInput
func (s *SpIdpConnectionsService) DeleteConnection(input *DeleteConnectionInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSigningSettings - Get the IdP connection's signature settings.
//RequestType: GET
//Input: input *GetSigningSettingsInput
func (s *SpIdpConnectionsService) GetSigningSettings(input *GetSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateSigningSettings - Update the IdP connection's signature settings.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *SpIdpConnectionsService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//AddConnectionCert - Add a new IdP connection certificate.
//RequestType: POST
//Input: input *AddConnectionCertInput
func (s *SpIdpConnectionsService) AddConnectionCert(input *AddConnectionCertInput) (result *models.ConnectionCert, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetConnectionCerts - Get the IdP connection's certificates.
//RequestType: GET
//Input: input *GetConnectionCertsInput
func (s *SpIdpConnectionsService) GetConnectionCerts(input *GetConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateConnectionCerts - Update the IdP connection's certificates.
//RequestType: PUT
//Input: input *UpdateConnectionCertsInput
func (s *SpIdpConnectionsService) UpdateConnectionCerts(input *UpdateConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//GetDecryptionKeys - Get the decryption keys of an IdP connection.
//RequestType: GET
//Input: input *GetDecryptionKeysInput
func (s *SpIdpConnectionsService) GetDecryptionKeys(input *GetDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateDecryptionKeys - Updating the IdP connection's decryption keys.
//RequestType: PUT
//Input: input *UpdateDecryptionKeysInput
func (s *SpIdpConnectionsService) UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

type AddConnectionCertInput struct {
	Body models.ConnectionCert
	Id   string
}

type CreateConnectionInput struct {
	Body models.IdpConnection

	BypassExternalValidation *bool
}

type DeleteConnectionInput struct {
	Id string
}

type GetConnectionInput struct {
	Id string
}

type GetConnectionCertsInput struct {
	Id string
}

type GetConnectionsInput struct {
	EntityId      string
	Page          string
	NumberPerPage string
	Filter        string
}

type GetDecryptionKeysInput struct {
	Id string
}

type GetSigningSettingsInput struct {
	Id string
}

type UpdateConnectionInput struct {
	Body models.IdpConnection
	Id   string

	BypassExternalValidation *bool
}

type UpdateConnectionCertsInput struct {
	Body models.ConnectionCerts
	Id   string
}

type UpdateDecryptionKeysInput struct {
	Body models.DecryptionKeys
	Id   string
}

type UpdateSigningSettingsInput struct {
	Body models.SigningSettings
	Id   string
}
