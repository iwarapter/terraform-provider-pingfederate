package metadataUrls

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
	ServiceName = "MetadataUrls"
)

type MetadataUrlsService struct {
	*client.PfClient
}

// New creates a new instance of the MetadataUrlsService client.
func New(cfg *config.Config) *MetadataUrlsService {

	return &MetadataUrlsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a MetadataUrls operation
func (c *MetadataUrlsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetMetadataUrls - Get a list of Metadata URLs
//RequestType: GET
//Input:
func (s *MetadataUrlsService) GetMetadataUrls() (output *models.MetadataUrls, resp *http.Response, err error) {
	path := "/metadataUrls"
	op := &request.Operation{
		Name:       "GetMetadataUrls",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.MetadataUrls{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddMetadataUrl - Add a new Metadata URL.
//RequestType: POST
//Input: input *AddMetadataUrlInput
func (s *MetadataUrlsService) AddMetadataUrl(input *AddMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls"
	op := &request.Operation{
		Name:       "AddMetadataUrl",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.MetadataUrl{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetMetadataUrl - Get a Metadata URL by ID.
//RequestType: GET
//Input: input *GetMetadataUrlInput
func (s *MetadataUrlsService) GetMetadataUrl(input *GetMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetMetadataUrl",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.MetadataUrl{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateMetadataUrl - Update a Metadata URL by ID.
//RequestType: PUT
//Input: input *UpdateMetadataUrlInput
func (s *MetadataUrlsService) UpdateMetadataUrl(input *UpdateMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateMetadataUrl",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.MetadataUrl{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteMetadataUrl - Delete a Metadata URL by ID.
//RequestType: DELETE
//Input: input *DeleteMetadataUrlInput
func (s *MetadataUrlsService) DeleteMetadataUrl(input *DeleteMetadataUrlInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteMetadataUrl",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddMetadataUrlInput struct {
	Body models.MetadataUrl
}

type DeleteMetadataUrlInput struct {
	Id string
}

type GetMetadataUrlInput struct {
	Id string
}

type UpdateMetadataUrlInput struct {
	Body models.MetadataUrl
	Id   string
}
