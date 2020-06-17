package idpConnectors

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "IdpConnectors"
)

type IdpConnectorsService struct {
	*client.PfClient
}

// New creates a new instance of the IdpConnectorsService client.
func New(cfg *config.Config) *IdpConnectorsService {

	return &IdpConnectorsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpConnectors operation
func (c *IdpConnectorsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIdpConnectorDescriptors - Get the list of available IdP connector descriptors.
//RequestType: GET
//Input:
func (s *IdpConnectorsService) GetIdpConnectorDescriptors() (output *models.SaasPluginDescriptors, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors"
	op := &request.Operation{
		Name:       "GetIdpConnectorDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SaasPluginDescriptors{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdpConnectorDescriptorById - Get the list of available connector descriptors.
//RequestType: GET
//Input: input *GetIdpConnectorDescriptorByIdInput
func (s *IdpConnectorsService) GetIdpConnectorDescriptorById(input *GetIdpConnectorDescriptorByIdInput) (output *models.SaasPluginDescriptor, resp *http.Response, err error) {
	path := "/idp/connectors/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdpConnectorDescriptorById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SaasPluginDescriptor{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type GetIdpConnectorDescriptorByIdInput struct {
	Id string
}
