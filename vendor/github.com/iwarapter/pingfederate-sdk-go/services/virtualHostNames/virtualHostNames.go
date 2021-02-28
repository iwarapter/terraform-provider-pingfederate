package virtualHostNames

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
	ServiceName = "VirtualHostNames"
)

type VirtualHostNamesService struct {
	*client.PfClient
}

// New creates a new instance of the VirtualHostNamesService client.
func New(cfg *config.Config) *VirtualHostNamesService {

	return &VirtualHostNamesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a VirtualHostNames operation
func (c *VirtualHostNamesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetVirtualHostNamesSettings - Retrieve virtual host names settings.
//RequestType: GET
//Input:
func (s *VirtualHostNamesService) GetVirtualHostNamesSettings() (output *models.VirtualHostNameSettings, resp *http.Response, err error) {
	return s.GetVirtualHostNamesSettingsWithContext(context.Background())
}

//GetVirtualHostNamesSettingsWithContext - Retrieve virtual host names settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *VirtualHostNamesService) GetVirtualHostNamesSettingsWithContext(ctx context.Context) (output *models.VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
	op := &request.Operation{
		Name:       "GetVirtualHostNamesSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.VirtualHostNameSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateVirtualHostNamesSettings - Update virtual host names settings.
//RequestType: PUT
//Input: input *UpdateVirtualHostNamesSettingsInput
func (s *VirtualHostNamesService) UpdateVirtualHostNamesSettings(input *UpdateVirtualHostNamesSettingsInput) (output *models.VirtualHostNameSettings, resp *http.Response, err error) {
	return s.UpdateVirtualHostNamesSettingsWithContext(context.Background(), input)
}

//UpdateVirtualHostNamesSettingsWithContext - Update virtual host names settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateVirtualHostNamesSettingsInput
func (s *VirtualHostNamesService) UpdateVirtualHostNamesSettingsWithContext(ctx context.Context, input *UpdateVirtualHostNamesSettingsInput) (output *models.VirtualHostNameSettings, resp *http.Response, err error) {
	path := "/virtualHostNames"
	op := &request.Operation{
		Name:       "UpdateVirtualHostNamesSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.VirtualHostNameSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type UpdateVirtualHostNamesSettingsInput struct {
	Body models.VirtualHostNameSettings
}
