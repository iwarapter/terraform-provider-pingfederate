package spTargetUrlMappings

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
	ServiceName = "SpTargetUrlMappings"
)

type SpTargetUrlMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the SpTargetUrlMappingsService client.
func New(cfg *config.Config) *SpTargetUrlMappingsService {

	return &SpTargetUrlMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpTargetUrlMappings operation
func (c *SpTargetUrlMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetUrlMappings - List the mappings between URLs and adapter or connection instances.
//RequestType: GET
//Input:
func (s *SpTargetUrlMappingsService) GetUrlMappings() (output *models.SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
	op := &request.Operation{
		Name:       "GetUrlMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.SpUrlMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateUrlMappings - Update the mappings between URLs and adapters or connections instances.
//RequestType: PUT
//Input: input *UpdateUrlMappingsInput
func (s *SpTargetUrlMappingsService) UpdateUrlMappings(input *UpdateUrlMappingsInput) (output *models.SpUrlMappings, resp *http.Response, err error) {
	path := "/sp/targetUrlMappings"
	op := &request.Operation{
		Name:       "UpdateUrlMappings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.SpUrlMappings{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateUrlMappingsInput struct {
	Body models.SpUrlMappings
}
