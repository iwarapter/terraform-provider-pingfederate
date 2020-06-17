package extendedProperties

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "ExtendedProperties"
)

type ExtendedPropertiesService struct {
	*client.PfClient
}

// New creates a new instance of the ExtendedPropertiesService client.
func New(cfg *config.Config) *ExtendedPropertiesService {

	return &ExtendedPropertiesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ExtendedProperties operation
func (c *ExtendedPropertiesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetExtendedProperties - Get the defined Extended Properties.
//RequestType: GET
//Input:
func (s *ExtendedPropertiesService) GetExtendedProperties() (output *models.ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
	op := &request.Operation{
		Name:       "GetExtendedProperties",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ExtendedProperties{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateExtendedProperties - Update the Extended Properties.
//RequestType: PUT
//Input: input *UpdateExtendedPropertiesInput
func (s *ExtendedPropertiesService) UpdateExtendedProperties(input *UpdateExtendedPropertiesInput) (output *models.ExtendedProperties, resp *http.Response, err error) {
	path := "/extendedProperties"
	op := &request.Operation{
		Name:       "UpdateExtendedProperties",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ExtendedProperties{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateExtendedPropertiesInput struct {
	Body models.ExtendedProperties
}
