package spAuthenticationPolicyContractMappings

import (
	"fmt"
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
	ServiceName = "SpAuthenticationPolicyContractMappings"
)

type SpAuthenticationPolicyContractMappingsService struct {
	*client.PfClient
}

// New creates a new instance of the SpAuthenticationPolicyContractMappingsService client.
func New(cfg *config.Config) *SpAuthenticationPolicyContractMappingsService {

	return &SpAuthenticationPolicyContractMappingsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a SpAuthenticationPolicyContractMappings operation
func (c *SpAuthenticationPolicyContractMappingsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetApcToSpAdapterMappings - Get the list of APC-to-SP Adapter Mappings.
//RequestType: GET
//Input:
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappings() (output *models.ApcToSpAdapterMappings, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
	op := &request.Operation{
		Name:       "GetApcToSpAdapterMappings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ApcToSpAdapterMappings{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateApcToSpAdapterMapping - Create a new APC-to-SP Adapter Mapping.
//RequestType: POST
//Input: input *CreateApcToSpAdapterMappingInput
func (s *SpAuthenticationPolicyContractMappingsService) CreateApcToSpAdapterMapping(input *CreateApcToSpAdapterMappingInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings"
	op := &request.Operation{
		Name:       "CreateApcToSpAdapterMapping",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ApcToSpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetApcToSpAdapterMappingById - Get an APC-to-SP Adapter Mapping.
//RequestType: GET
//Input: input *GetApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) GetApcToSpAdapterMappingById(input *GetApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetApcToSpAdapterMappingById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ApcToSpAdapterMapping{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateApcToSpAdapterMappingById - Update an APC-to-SP Adapter Mapping.
//RequestType: PUT
//Input: input *UpdateApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) UpdateApcToSpAdapterMappingById(input *UpdateApcToSpAdapterMappingByIdInput) (output *models.ApcToSpAdapterMapping, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateApcToSpAdapterMappingById",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ApcToSpAdapterMapping{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteApcToSpAdapterMappingById - Delete an APC-to-SP Adapter Mapping.
//RequestType: DELETE
//Input: input *DeleteApcToSpAdapterMappingByIdInput
func (s *SpAuthenticationPolicyContractMappingsService) DeleteApcToSpAdapterMappingById(input *DeleteApcToSpAdapterMappingByIdInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/sp/authenticationPolicyContractMappings/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteApcToSpAdapterMappingById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateApcToSpAdapterMappingInput struct {
	Body models.ApcToSpAdapterMapping

	BypassExternalValidation *bool
}

type DeleteApcToSpAdapterMappingByIdInput struct {
	Id string
}

type GetApcToSpAdapterMappingByIdInput struct {
	Id string
}

type UpdateApcToSpAdapterMappingByIdInput struct {
	Body models.ApcToSpAdapterMapping
	Id   string

	BypassExternalValidation *bool
}
