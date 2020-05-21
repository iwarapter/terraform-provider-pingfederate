package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type CertificatesCaService service

//GetTrustedCAs - Get list of trusted certificate authorities.
//RequestType: GET
//Input:
func (s *CertificatesCaService) GetTrustedCAs() (result *CertViews, resp *http.Response, err error) {
	path := "/certificates/ca"
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

//ImportTrustedCA - Import a new trusted certificate authority.
//RequestType: POST
//Input: input *ImportTrustedCAInput
func (s *CertificatesCaService) ImportTrustedCA(input *ImportTrustedCAInput) (result *CertView, resp *http.Response, err error) {
	path := "/certificates/ca/import"
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

//GetTrustedCert - Retrieve details of a trusted certificate authority.
//RequestType: GET
//Input: input *GetTrustedCertInput
func (s *CertificatesCaService) GetTrustedCert(input *GetTrustedCertInput) (result *CertView, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
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

//DeleteTrustedCA - Delete a trusted certificate authority.
//RequestType: DELETE
//Input: input *DeleteTrustedCAInput
func (s *CertificatesCaService) DeleteTrustedCA(input *DeleteTrustedCAInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
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

//ExportCertificateFile - Download the certificate from a given trusted certificate authority.
//RequestType: GET
//Input: input *ExportCertificateFileInput
func (s *CertificatesCaService) ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error) {
	path := "/certificates/ca/{id}/file"
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
