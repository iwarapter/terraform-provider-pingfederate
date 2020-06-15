package idpSpConnections

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpSpConnectionsService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpSpConnectionsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpSpConnectionsService {

	return &IdpSpConnectionsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetConnections - Get list of SP connections.
//RequestType: GET
//Input: input *GetConnectionsInput
func (s *IdpSpConnectionsService) GetConnections(input *GetConnectionsInput) (result *models.SpConnections, resp *http.Response, err error) {
	path := "/idp/spConnections"
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

//CreateConnection - Create a new SP connection.
//RequestType: POST
//Input: input *CreateConnectionInput
func (s *IdpSpConnectionsService) CreateConnection(input *CreateConnectionInput) (result *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections"
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

//GetConnection - Find SP connection by ID.
//RequestType: GET
//Input: input *GetConnectionInput
func (s *IdpSpConnectionsService) GetConnection(input *GetConnectionInput) (result *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
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

//UpdateConnection - Update an SP connection.
//RequestType: PUT
//Input: input *UpdateConnectionInput
func (s *IdpSpConnectionsService) UpdateConnection(input *UpdateConnectionInput) (result *models.SpConnection, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
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

//DeleteConnection - Delete an SP connection.
//RequestType: DELETE
//Input: input *DeleteConnectionInput
func (s *IdpSpConnectionsService) DeleteConnection(input *DeleteConnectionInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}"
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

//GetSigningSettings - Get the SP connection's signature settings.
//RequestType: GET
//Input: input *GetSigningSettingsInput
func (s *IdpSpConnectionsService) GetSigningSettings(input *GetSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/signingSettings"
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

//UpdateSigningSettings - Update the SP connection's signature settings.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *IdpSpConnectionsService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (result *models.SigningSettings, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/signingSettings"
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

//AddConnectionCert - Add a new SP connection certificate.
//RequestType: POST
//Input: input *AddConnectionCertInput
func (s *IdpSpConnectionsService) AddConnectionCert(input *AddConnectionCertInput) (result *models.ConnectionCert, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
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

//GetConnectionCerts - Get the SP connection's certificates.
//RequestType: GET
//Input: input *GetConnectionCertsInput
func (s *IdpSpConnectionsService) GetConnectionCerts(input *GetConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
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

//UpdateConnectionCerts - Update the SP connection's certificates.
//RequestType: PUT
//Input: input *UpdateConnectionCertsInput
func (s *IdpSpConnectionsService) UpdateConnectionCerts(input *UpdateConnectionCertsInput) (result *models.ConnectionCerts, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/certs"
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

//GetDecryptionKeys - Get the decryption keys of an SP connection.
//RequestType: GET
//Input: input *GetDecryptionKeysInput
func (s *IdpSpConnectionsService) GetDecryptionKeys(input *GetDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/decryptionKeys"
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

//UpdateDecryptionKeys - Updating the SP connection's decryption keys.
//RequestType: PUT
//Input: input *UpdateDecryptionKeysInput
func (s *IdpSpConnectionsService) UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (result *models.DecryptionKeys, resp *http.Response, err error) {
	path := "/idp/spConnections/{id}/credentials/decryptionKeys"
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
	Body models.SpConnection

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
	Body models.SpConnection
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
