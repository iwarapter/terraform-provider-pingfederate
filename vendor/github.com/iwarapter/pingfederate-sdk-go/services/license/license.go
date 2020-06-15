package license

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LicenseService struct {
	Client *client.PfClient
}

// New creates a new instance of the LicenseService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *LicenseService {

	return &LicenseService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetLicenseAgreement - Get license agreement link.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicenseAgreement() (result *models.LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
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

//UpdateLicenseAgreement - Accept license agreement.
//RequestType: PUT
//Input: input *UpdateLicenseAgreementInput
func (s *LicenseService) UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (result *models.LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
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

//GetLicense - Get a license summary.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicense() (result *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
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

//UpdateLicense - Import a license.
//RequestType: PUT
//Input: input *UpdateLicenseInput
func (s *LicenseService) UpdateLicense(input *UpdateLicenseInput) (result *models.LicenseView, resp *http.Response, err error) {
	path := "/license"
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

type UpdateLicenseInput struct {
	Body models.LicenseFile
}

type UpdateLicenseAgreementInput struct {
	Body models.LicenseAgreementInfo
}
