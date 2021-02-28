package version

import (
	"context"
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
	ServiceName = "Version"
)

type VersionService struct {
	*client.PfClient
}

// New creates a new instance of the VersionService client.
func New(cfg *config.Config) *VersionService {

	return &VersionService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Version operation
func (c *VersionService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetVersion - Gets the server version.
//RequestType: GET
//Input:
func (s *VersionService) GetVersion() (output *models.Version, resp *http.Response, err error) {
	return s.GetVersionWithContext(context.Background())
}

//GetVersionWithContext - Gets the server version.
//RequestType: GET
//Input: ctx context.Context,
func (s *VersionService) GetVersionWithContext(ctx context.Context) (output *models.Version, resp *http.Response, err error) {
	path := "/version"
	op := &request.Operation{
		Name:       "GetVersion",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Version{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
