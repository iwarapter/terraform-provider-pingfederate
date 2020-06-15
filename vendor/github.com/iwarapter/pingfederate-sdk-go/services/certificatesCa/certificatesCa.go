package certificatesCa

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type CertificatesCaService struct {
	Client *client.PfClient
}

// New creates a new instance of the CertificatesCaService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *CertificatesCaService {

	return &CertificatesCaService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetTrustedCAs - Get list of trusted certificate authorities.
//RequestType: GET
//Input:
func (s *CertificatesCaService) GetTrustedCAs() (result *models.CertViews, resp *http.Response, err error) {
	path := "/certificates/ca"
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

//GetTrustedCert - Retrieve details of a trusted certificate authority.
//RequestType: GET
//Input: input *GetTrustedCertInput
func (s *CertificatesCaService) GetTrustedCert(input *GetTrustedCertInput) (result *models.CertView, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
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

//DeleteTrustedCA - Delete a trusted certificate authority.
//RequestType: DELETE
//Input: input *DeleteTrustedCAInput
func (s *CertificatesCaService) DeleteTrustedCA(input *DeleteTrustedCAInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/certificates/ca/{id}"
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

//ImportTrustedCA - Import a new trusted certificate authority.
//RequestType: POST
//Input: input *ImportTrustedCAInput
func (s *CertificatesCaService) ImportTrustedCA(input *ImportTrustedCAInput) (result *models.CertView, resp *http.Response, err error) {
	path := "/certificates/ca/import"
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

//ExportCertificateFile - Download the certificate from a given trusted certificate authority.
//RequestType: GET
//Input: input *ExportCertificateFileInput
func (s *CertificatesCaService) ExportCertificateFile(input *ExportCertificateFileInput) (result *string, resp *http.Response, err error) {
	path := "/certificates/ca/{id}/file"
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

type DeleteTrustedCAInput struct {
	Id string
}

type ExportCertificateFileInput struct {
	Id string
}

type GetTrustedCertInput struct {
	Id string
}

type ImportTrustedCAInput struct {
	Body models.X509File
}
