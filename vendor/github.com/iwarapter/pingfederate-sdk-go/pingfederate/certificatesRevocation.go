package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type CertificatesRevocationService service

//GetRevocationSettings - Get certificate revocation settings.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetRevocationSettings() (result *CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
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

//UpdateRevocationSettings - Update certificate revocation settings.
//RequestType: PUT
//Input: input *UpdateRevocationSettingsInput
func (s *CertificatesRevocationService) UpdateRevocationSettings(input *UpdateRevocationSettingsInput) (result *CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
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

//GetOcspCertificates - Get the list of available OCSP responder signature verification certificates.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetOcspCertificates() (result *CertViews, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
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

//ImportOcspCertificate - Import an OCSP responder signature verification certificate.
//RequestType: POST
//Input: input *ImportOcspCertificateInput
func (s *CertificatesRevocationService) ImportOcspCertificate(input *ImportOcspCertificateInput) (result *CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
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

//GetOcspCertificateById - Get an OCSP responder signature verification certificate by ID.
//RequestType: GET
//Input: input *GetOcspCertificateByIdInput
func (s *CertificatesRevocationService) GetOcspCertificateById(input *GetOcspCertificateByIdInput) (result *CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
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

//DeleteOcspCertificateById - Delete an OCSP responder signature verification certificate by ID.
//RequestType: DELETE
//Input: input *DeleteOcspCertificateByIdInput
func (s *CertificatesRevocationService) DeleteOcspCertificateById(input *DeleteOcspCertificateByIdInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
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
