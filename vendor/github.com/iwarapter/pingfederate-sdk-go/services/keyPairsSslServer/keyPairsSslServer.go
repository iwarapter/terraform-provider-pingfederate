package keyPairsSslServer

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSslServerService struct {
	Client *client.PfClient
}

// New creates a new instance of the KeyPairsSslServerService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *KeyPairsSslServerService {

	return &KeyPairsSslServerService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetKeyPairs() (result *models.KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/sslServer"
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

//ImportKeyPair - Import a new key pair.
//RequestType: POST
//Input: input *ImportKeyPairInput
func (s *KeyPairsSslServerService) ImportKeyPair(input *ImportKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/import"
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

//CreateKeyPair - Generate a new key pair.
//RequestType: POST
//Input: input *CreateKeyPairInput
func (s *KeyPairsSslServerService) CreateKeyPair(input *CreateKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/generate"
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

//GetKeyPair - Retrieve details of a key pair.
//RequestType: GET
//Input: input *GetKeyPairInput
func (s *KeyPairsSslServerService) GetKeyPair(input *GetKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
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

//DeleteKeyPair - Delete a key pair.
//RequestType: DELETE
//Input: input *DeleteKeyPairInput
func (s *KeyPairsSslServerService) DeleteKeyPair(input *DeleteKeyPairInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
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

//ExportCsr - Generate a new certificate signing request (CSR) for this key pair.
//RequestType: GET
//Input: input *ExportCsrInput
func (s *KeyPairsSslServerService) ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
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

//ImportCsrResponse - Import a CSR response for this key pair.
//RequestType: POST
//Input: input *ImportCsrResponseInput
func (s *KeyPairsSslServerService) ImportCsrResponse(input *ImportCsrResponseInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
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

//ExportPKCS12File - Download the key pair in PKCS12 format.
//RequestType: POST
//Input: input *ExportPKCS12FileInput
func (s *KeyPairsSslServerService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/pkcs12"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, err
	}

	resp, err = s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//ExportCertificateFile - Download the certificate from a given key pair.
//RequestType: GET
//Input: input *ExportCertificateFileInput
func (s *KeyPairsSslServerService) ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/certificate"
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

//GetSettings - Get the SSL Server Certificate Settings.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetSettings() (result *models.SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
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

//UpdateSettings - Update the SSL Server Certificate Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KeyPairsSslServerService) UpdateSettings(input *UpdateSettingsInput) (result *models.SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
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

type CreateKeyPairInput struct {
	Body models.NewKeyPairSettings
}

type DeleteKeyPairInput struct {
	Id string
}

type ExportCertificateFileInput struct {
	Id string
}

type ExportCsrInput struct {
	Id string
}

type ExportPKCS12FileInput struct {
	Body models.PKCS12ExportSettings
	Id   string
}

type GetKeyPairInput struct {
	Id string
}

type ImportCsrResponseInput struct {
	Body models.CSRResponse
	Id   string
}

type ImportKeyPairInput struct {
	Body models.PKCS12File
}

type UpdateSettingsInput struct {
	Body models.SslServerSettings
}
