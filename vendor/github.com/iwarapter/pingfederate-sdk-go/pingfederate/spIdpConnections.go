package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type SpIdpConnectionsService service

//GetConnections - Get list of IdP connections.
//RequestType: GET
//Input: input *GetConnectionsInput
func (s *SpIdpConnectionsService) GetConnections(input *GetConnectionsInput) (result *IdpConnections, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
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

//CreateConnection - Create a new IdP connection.
//RequestType: POST
//Input: input *CreateConnectionInput
func (s *SpIdpConnectionsService) CreateConnection(input *CreateConnectionInput) (result *IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetConnection - Find IdP connection by ID.
//RequestType: GET
//Input: input *GetConnectionInput
func (s *SpIdpConnectionsService) GetConnection(input *GetConnectionInput) (result *IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateConnection - Update an IdP connection.
//RequestType: PUT
//Input: input *UpdateConnectionInput
func (s *SpIdpConnectionsService) UpdateConnection(input *UpdateConnectionInput) (result *IdpConnection, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteConnection - Delete an IdP connection.
//RequestType: DELETE
//Input: input *DeleteConnectionInput
func (s *SpIdpConnectionsService) DeleteConnection(input *DeleteConnectionInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetSigningSettings - Get the IdP connection's signature settings.
//RequestType: GET
//Input: input *GetSigningSettingsInput
func (s *SpIdpConnectionsService) GetSigningSettings(input *GetSigningSettingsInput) (result *SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateSigningSettings - Update the IdP connection's signature settings.
//RequestType: PUT
//Input: input *UpdateSigningSettingsInput
func (s *SpIdpConnectionsService) UpdateSigningSettings(input *UpdateSigningSettingsInput) (result *SigningSettings, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/signingSettings"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//AddConnectionCert - Add a new IdP connection certificate.
//RequestType: POST
//Input: input *AddConnectionCertInput
func (s *SpIdpConnectionsService) AddConnectionCert(input *AddConnectionCertInput) (result *ConnectionCert, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetConnectionCerts - Get the IdP connection's certificates.
//RequestType: GET
//Input: input *GetConnectionCertsInput
func (s *SpIdpConnectionsService) GetConnectionCerts(input *GetConnectionCertsInput) (result *ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateConnectionCerts - Update the IdP connection's certificates.
//RequestType: PUT
//Input: input *UpdateConnectionCertsInput
func (s *SpIdpConnectionsService) UpdateConnectionCerts(input *UpdateConnectionCertsInput) (result *ConnectionCerts, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/certs"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//GetDecryptionKeys - Get the decryption keys of an IdP connection.
//RequestType: GET
//Input: input *GetDecryptionKeysInput
func (s *SpIdpConnectionsService) GetDecryptionKeys(input *GetDecryptionKeysInput) (result *DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

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

//UpdateDecryptionKeys - Updating the IdP connection's decryption keys.
//RequestType: PUT
//Input: input *UpdateDecryptionKeysInput
func (s *SpIdpConnectionsService) UpdateDecryptionKeys(input *UpdateDecryptionKeysInput) (result *DecryptionKeys, resp *http.Response, err error) {
	path := "/sp/idpConnections/{id}/credentials/decryptionKeys"
	path = strings.Replace(path, "{id}", input.Id, -1)

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
