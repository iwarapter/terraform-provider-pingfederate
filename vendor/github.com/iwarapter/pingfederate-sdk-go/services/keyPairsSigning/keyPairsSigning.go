package keyPairsSigning

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KeyPairsSigningService struct {
	Client *client.PfClient
}

// New creates a new instance of the KeyPairsSigningService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *KeyPairsSigningService {

	return &KeyPairsSigningService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSigningService) GetKeyPairs() (result *models.KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/signing"
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
func (s *KeyPairsSigningService) ImportKeyPair(input *ImportKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/import"
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
func (s *KeyPairsSigningService) CreateKeyPair(input *CreateKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/generate"
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
func (s *KeyPairsSigningService) GetKeyPair(input *GetKeyPairInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}"
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
func (s *KeyPairsSigningService) DeleteKeyPair(input *DeleteKeyPairInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}"
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
func (s *KeyPairsSigningService) ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/csr"
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
func (s *KeyPairsSigningService) ImportCsrResponse(input *ImportCsrResponseInput) (result *models.KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/csr"
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
func (s *KeyPairsSigningService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/pkcs12"
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
func (s *KeyPairsSigningService) ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/certificate"
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

//GetRotationSettings - Retrieve details of rotation settings for a key pair.
//RequestType: GET
//Input: input *GetRotationSettingsInput
func (s *KeyPairsSigningService) GetRotationSettings(input *GetRotationSettingsInput) (result *models.KeyPairRotationSettings, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
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

//UpdateRotationSettings - Add rotation settings to a key pair
//RequestType: PUT
//Input: input *UpdateRotationSettingsInput
func (s *KeyPairsSigningService) UpdateRotationSettings(input *UpdateRotationSettingsInput) (result *models.KeyPairRotationSettings, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
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

//DeleteKeyPairRotationSettings - Delete rotation settings for a signing key pair.
//RequestType: DELETE
//Input: input *DeleteKeyPairRotationSettingsInput
func (s *KeyPairsSigningService) DeleteKeyPairRotationSettings(input *DeleteKeyPairRotationSettingsInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/signing/{id}/rotationSettings"
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

type CreateKeyPairInput struct {
	Body models.NewKeyPairSettings
}

type DeleteKeyPairInput struct {
	Id string
}

type DeleteKeyPairRotationSettingsInput struct {
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

type GetRotationSettingsInput struct {
	Id string
}

type ImportCsrResponseInput struct {
	Body models.CSRResponse
	Id   string
}

type ImportKeyPairInput struct {
	Body models.PKCS12File
}

type UpdateRotationSettingsInput struct {
	Body models.KeyPairRotationSettings
	Id   string
}
