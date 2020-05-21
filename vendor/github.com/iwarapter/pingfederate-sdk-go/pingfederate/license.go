package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type LicenseService service

//GetLicenseAgreement - Get license agreement link.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicenseAgreement() (result *LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
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

//UpdateLicenseAgreement - Accept license agreement.
//RequestType: PUT
//Input: input *UpdateLicenseAgreementInput
func (s *LicenseService) UpdateLicenseAgreement(input *UpdateLicenseAgreementInput) (result *LicenseAgreementInfo, resp *http.Response, err error) {
	path := "/license/agreement"
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

//GetLicense - Get a license summary.
//RequestType: GET
//Input:
func (s *LicenseService) GetLicense() (result *LicenseView, resp *http.Response, err error) {
	path := "/license"
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

//UpdateLicense - Import a license.
//RequestType: PUT
//Input: input *UpdateLicenseInput
func (s *LicenseService) UpdateLicense(input *UpdateLicenseInput) (result *LicenseView, resp *http.Response, err error) {
	path := "/license"
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
