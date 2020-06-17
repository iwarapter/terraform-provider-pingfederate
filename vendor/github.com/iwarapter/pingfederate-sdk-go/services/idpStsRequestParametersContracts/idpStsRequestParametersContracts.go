package idpStsRequestParametersContracts

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
	ServiceName = "IdpStsRequestParametersContracts"
)

type IdpStsRequestParametersContractsService struct {
	*client.PfClient
}

// New creates a new instance of the IdpStsRequestParametersContractsService client.
func New(cfg *config.Config) *IdpStsRequestParametersContractsService {

	return &IdpStsRequestParametersContractsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a IdpStsRequestParametersContracts operation
func (c *IdpStsRequestParametersContractsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetStsRequestParamContracts - Get the list of STS Request Parameters Contracts.
//RequestType: GET
//Input:
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContracts() (output *models.StsRequestParametersContracts, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
	op := &request.Operation{
		Name:       "GetStsRequestParamContracts",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.StsRequestParametersContracts{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateStsRequestParamContract - Create a new STS Request Parameters Contract.
//RequestType: POST
//Input: input *CreateStsRequestParamContractInput
func (s *IdpStsRequestParametersContractsService) CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (output *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts"
	op := &request.Operation{
		Name:       "CreateStsRequestParamContract",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.StsRequestParametersContract{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetStsRequestParamContractById - Get a STS Request Parameters Contract.
//RequestType: GET
//Input: input *GetStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetStsRequestParamContractById",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.StsRequestParametersContract{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateStsRequestParamContractById - Update a STS Request Parameters Contract.
//RequestType: PUT
//Input: input *UpdateStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateStsRequestParamContractById",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.StsRequestParametersContract{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteStsRequestParamContractById - Delete a STS Request Parameters Contract.
//RequestType: DELETE
//Input: input *DeleteStsRequestParamContractByIdInput
func (s *IdpStsRequestParametersContractsService) DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/idp/stsRequestParametersContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteStsRequestParamContractById",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateStsRequestParamContractInput struct {
	Body models.StsRequestParametersContract
}

type DeleteStsRequestParamContractByIdInput struct {
	Id string
}

type GetStsRequestParamContractByIdInput struct {
	Id string
}

type UpdateStsRequestParamContractByIdInput struct {
	Body models.StsRequestParametersContract
	Id   string
}
