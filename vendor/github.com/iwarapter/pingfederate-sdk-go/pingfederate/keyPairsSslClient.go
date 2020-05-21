package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type KeyPairsSslClientService service

//GetKeyPairs - Get list of key pairs.
//RequestType: GET
//Input:
func (s *KeyPairsSslClientService) GetKeyPairs() (result *KeyPairViews, resp *http.Response, err error) {
	path := "/keyPairs/sslClient"
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
func (s *KeyPairsSslClientService) ImportKeyPair(input *ImportKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/import"
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
func (s *KeyPairsSslClientService) CreateKeyPair(input *CreateKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/generate"
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
func (s *KeyPairsSslClientService) GetKeyPair(input *GetKeyPairInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}"
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
func (s *KeyPairsSslClientService) DeleteKeyPair(input *DeleteKeyPairInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}"
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
func (s *KeyPairsSslClientService) ExportCsr(input *ExportCsrInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}/csr"
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
func (s *KeyPairsSslClientService) ImportCsrResponse(input *ImportCsrResponseInput) (result *KeyPairView, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}/csr"
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
func (s *KeyPairsSslClientService) ExportPKCS12File(input *ExportPKCS12FileInput) (resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}/pkcs12"
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
func (s *KeyPairsSslClientService) ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error) {
	path := "/keyPairs/sslClient/{id}/certificate"
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
