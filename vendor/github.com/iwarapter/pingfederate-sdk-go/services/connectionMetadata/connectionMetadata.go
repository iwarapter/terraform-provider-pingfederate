package connectionMetadata

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
	ServiceName = "ConnectionMetadata"
)

type ConnectionMetadataService struct {
	*client.PfClient
}

// New creates a new instance of the ConnectionMetadataService client.
func New(cfg *config.Config) *ConnectionMetadataService {

	return &ConnectionMetadataService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ConnectionMetadata operation
func (c *ConnectionMetadataService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//Export - Export a connection's SAML metadata that can be given to a partner.
//RequestType: POST
//Input: input *ExportInput
func (s *ConnectionMetadataService) Export(input *ExportInput) (output *string, resp *http.Response, err error) {
	path := "/connectionMetadata/export"
	op := &request.Operation{
		Name:       "Export",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = pingfederate.String("")
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//Convert - Convert a partner's SAML metadata into a JSON representation.
//RequestType: POST
//Input: input *ConvertInput
func (s *ConnectionMetadataService) Convert(input *ConvertInput) (output *models.ConvertMetadataResponse, resp *http.Response, err error) {
	path := "/connectionMetadata/convert"
	op := &request.Operation{
		Name:       "Convert",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ConvertMetadataResponse{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type ConvertInput struct {
	Body models.ConvertMetadataRequest
}

type ExportInput struct {
	Body models.ExportMetadataRequest
}
