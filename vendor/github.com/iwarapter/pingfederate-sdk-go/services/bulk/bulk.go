package bulk

import (
	"fmt"
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
	ServiceName = "Bulk"
)

type BulkService struct {
	*client.PfClient
}

// New creates a new instance of the BulkService client.
func New(cfg *config.Config) *BulkService {

	return &BulkService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Bulk operation
func (c *BulkService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//ExportConfiguration - Export all API resources to a JSON file.
//RequestType: GET
//Input: input *ExportConfigurationInput
func (s *BulkService) ExportConfiguration(input *ExportConfigurationInput) (output *models.BulkConfig, resp *http.Response, err error) {
	path := "/bulk/export"
	op := &request.Operation{
		Name:       "ExportConfiguration",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.BulkConfig{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ImportConfiguration - Import configuration for a PingFederate deployment from a JSON file.
//RequestType: POST
//Input: input *ImportConfigurationInput
func (s *BulkService) ImportConfiguration(input *ImportConfigurationInput) (resp *http.Response, err error) {
	path := "/bulk/import"
	op := &request.Operation{
		Name:       "ImportConfiguration",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}

	req := s.newRequest(op, input.Body, nil)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

type ExportConfigurationInput struct {
	IncludeExternalResources string
}

type ImportConfigurationInput struct {
	FailFast string

	Body models.BulkConfig

	BypassExternalValidation *bool
}
