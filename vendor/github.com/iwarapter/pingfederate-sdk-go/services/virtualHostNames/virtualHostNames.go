package virtualHostNames

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type VirtualHostNamesService struct {
	Client *client.PfClient
}

// New creates a new instance of the VirtualHostNamesService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *VirtualHostNamesService {

	return &VirtualHostNamesService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetVirtualHostNamesSettings - Retrieve virtual host names settings.
//RequestType: GET
//Input:
func (s *VirtualHostNamesService) GetVirtualHostNamesSettings() (result *models.VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
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

//UpdateVirtualHostNamesSettings - Update virtual host names settings.
//RequestType: PUT
//Input: input *UpdateVirtualHostNamesSettingsInput
func (s *VirtualHostNamesService) UpdateVirtualHostNamesSettings(input *UpdateVirtualHostNamesSettingsInput) (result *models.VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
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

type UpdateVirtualHostNamesSettingsInput struct {
	Body models.VirtualHostNameSettings
}
