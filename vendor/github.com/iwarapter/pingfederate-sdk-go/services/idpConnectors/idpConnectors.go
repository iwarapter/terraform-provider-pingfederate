package idpConnectors

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpConnectorsService struct {
	Client *client.PfClient
}

// New creates a new instance of the IdpConnectorsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *IdpConnectorsService {

	return &IdpConnectorsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetIdpConnectorDescriptors - Get the list of available IdP connector descriptors.
//RequestType: GET
//Input:
func (s *IdpConnectorsService) GetIdpConnectorDescriptors() (result *models.SaasPluginDescriptors, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors"
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

//GetIdpConnectorDescriptorById - Get the list of available connector descriptors.
//RequestType: GET
//Input: input *GetIdpConnectorDescriptorByIdInput
func (s *IdpConnectorsService) GetIdpConnectorDescriptorById(input *GetIdpConnectorDescriptorByIdInput) (result *models.SaasPluginDescriptor, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors/{id}"
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

type GetIdpConnectorDescriptorByIdInput struct {
	Id string
}
