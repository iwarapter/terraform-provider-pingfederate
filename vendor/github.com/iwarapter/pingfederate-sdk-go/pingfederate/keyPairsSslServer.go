package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type KeyPairsSslServerService service

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetKeyPairs() (result *KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/sslServer"
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

//ImportKeyPair - Import a new key pair.
//RequestType: POST
//Input: input *ImportKeyPairInput
func (s *KeyPairsSslServerService) ImportKeyPair(input *ImportKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/import"
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

//CreateKeyPair - Generate a new key pair.
//RequestType: POST
//Input: input *CreateKeyPairInput
func (s *KeyPairsSslServerService) CreateKeyPair(input *CreateKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/generate"
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

//GetKeyPair - Retrieve details of a key pair.
//RequestType: GET
//Input: input *GetKeyPairInput
func (s *KeyPairsSslServerService) GetKeyPair(input *GetKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
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

//DeleteKeyPair - Delete a key pair.
//RequestType: DELETE
//Input: input *DeleteKeyPairInput
func (s *KeyPairsSslServerService) DeleteKeyPair(input *DeleteKeyPairInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}"
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

//ExportCsr - Generate a new certificate signing request (CSR) for this key pair.
//RequestType: GET
//Input: input *ExportCsrInput
func (s *KeyPairsSslServerService) ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
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

//ImportCsrResponse - Import a CSR response for this key pair.
//RequestType: POST
//Input: input *ImportCsrResponseInput
func (s *KeyPairsSslServerService) ImportCsrResponse(input *ImportCsrResponseInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/csr"
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

//ExportPKCS12File - Download the key pair in PKCS12 format.
//RequestType: POST
//Input: input *ExportPKCS12FileInput
func (s *KeyPairsSslServerService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/sslServer/{id}/pkcs12"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
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

//GetSettings - Get the SSL Server Certificate Settings.
//RequestType: GET
//Input:
func (s *KeyPairsSslServerService) GetSettings() (result *SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
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

//UpdateSettings - Update the SSL Server Certificate Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KeyPairsSslServerService) UpdateSettings(input *UpdateSettingsInput) (result *SslServerSettings, resp *http.Response, err error) {
	path := "/keyPairs/sslServer/settings"
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
