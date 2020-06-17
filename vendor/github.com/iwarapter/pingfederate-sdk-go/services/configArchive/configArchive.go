package configArchive

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
	ServiceName = "ConfigArchive"
)

type ConfigArchiveService struct {
	*client.PfClient
}

// New creates a new instance of the ConfigArchiveService client.
func New(cfg *config.Config) *ConfigArchiveService {

	return &ConfigArchiveService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ConfigArchive operation
func (c *ConfigArchiveService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//ImportConfigArchive - Import a configuration archive.
//RequestType: POST
//Input: input *ImportConfigArchiveInput
func (s *ConfigArchiveService) ImportConfigArchive(input *ImportConfigArchiveInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/configArchive/import"
	op := &request.Operation{
		Name:       "ImportConfigArchive",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}

	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ExportConfigArchive - Export a configuration archive.
//RequestType: GET
//Input:
func (s *ConfigArchiveService) ExportConfigArchive() (resp *http.Response, err error) {
	path := "/configArchive/export"
	op := &request.Operation{
		Name:       "ExportConfigArchive",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, nil)

	if req.Send() == nil {
		return req.HTTPResponse, nil
	}
	return req.HTTPResponse, req.Error
}

type ImportConfigArchiveInput struct {
	ForceImport string

	Body []byte
}
