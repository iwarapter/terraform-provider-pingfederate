package certificatesRevocation

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesRevocationService struct {
	Client *client.PfClient
}

// New creates a new instance of the CertificatesRevocationService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *CertificatesRevocationService {

	return &CertificatesRevocationService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetRevocationSettings - Get certificate revocation settings.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetRevocationSettings() (result *models.CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
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

//UpdateRevocationSettings - Update certificate revocation settings.
//RequestType: PUT
//Input: input *UpdateRevocationSettingsInput
func (s *CertificatesRevocationService) UpdateRevocationSettings(input *UpdateRevocationSettingsInput) (result *models.CertificateRevocationSettings, resp *http.Response, err error) {
	path := "/certificates/revocation/settings"
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

//GetOcspCertificates - Get the list of available OCSP responder signature verification certificates.
//RequestType: GET
//Input:
func (s *CertificatesRevocationService) GetOcspCertificates() (result *models.CertViews, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
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

//ImportOcspCertificate - Import an OCSP responder signature verification certificate.
//RequestType: POST
//Input: input *ImportOcspCertificateInput
func (s *CertificatesRevocationService) ImportOcspCertificate(input *ImportOcspCertificateInput) (result *models.CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates"
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

//GetOcspCertificateById - Get an OCSP responder signature verification certificate by ID.
//RequestType: GET
//Input: input *GetOcspCertificateByIdInput
func (s *CertificatesRevocationService) GetOcspCertificateById(input *GetOcspCertificateByIdInput) (result *models.CertView, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
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

//DeleteOcspCertificateById - Delete an OCSP responder signature verification certificate by ID.
//RequestType: DELETE
//Input: input *DeleteOcspCertificateByIdInput
func (s *CertificatesRevocationService) DeleteOcspCertificateById(input *DeleteOcspCertificateByIdInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/certificates/revocation/ocspCertificates/{id}"
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

type DeleteOcspCertificateByIdInput struct {
	Id string
}

type GetOcspCertificateByIdInput struct {
	Id string
}

type ImportOcspCertificateInput struct {
	Body models.X509File
}

type UpdateRevocationSettingsInput struct {
	Body models.CertificateRevocationSettings
}
