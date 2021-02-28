package authenticationPolicyContracts

import (
	"context"
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
	ServiceName = "AuthenticationPolicyContracts"
)

type AuthenticationPolicyContractsService struct {
	*client.PfClient
}

// New creates a new instance of the AuthenticationPolicyContractsService client.
func New(cfg *config.Config) *AuthenticationPolicyContractsService {

	return &AuthenticationPolicyContractsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationPolicyContracts operation
func (c *AuthenticationPolicyContractsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAuthenticationPolicyContracts - Gets the Authentication Policy Contracts.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractsInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContracts(input *GetAuthenticationPolicyContractsInput) (output *models.AuthenticationPolicyContracts, resp *http.Response, err error) {
	return s.GetAuthenticationPolicyContractsWithContext(context.Background(), input)
}

//GetAuthenticationPolicyContractsWithContext - Gets the Authentication Policy Contracts.
//RequestType: GET
//Input: ctx context.Context, input *GetAuthenticationPolicyContractsInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContractsWithContext(ctx context.Context, input *GetAuthenticationPolicyContractsInput) (output *models.AuthenticationPolicyContracts, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
	op := &request.Operation{
		Name:       "GetAuthenticationPolicyContracts",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.AuthenticationPolicyContracts{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateAuthenticationPolicyContract - Create a new Authentication Policy Contract.
//RequestType: POST
//Input: input *CreateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) CreateAuthenticationPolicyContract(input *CreateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	return s.CreateAuthenticationPolicyContractWithContext(context.Background(), input)
}

//CreateAuthenticationPolicyContractWithContext - Create a new Authentication Policy Contract.
//RequestType: POST
//Input: ctx context.Context, input *CreateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) CreateAuthenticationPolicyContractWithContext(ctx context.Context, input *CreateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts"
	op := &request.Operation{
		Name:       "CreateAuthenticationPolicyContract",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyContract{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAuthenticationPolicyContract - Gets the Authentication Policy Contract by ID.
//RequestType: GET
//Input: input *GetAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContract(input *GetAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	return s.GetAuthenticationPolicyContractWithContext(context.Background(), input)
}

//GetAuthenticationPolicyContractWithContext - Gets the Authentication Policy Contract by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) GetAuthenticationPolicyContractWithContext(ctx context.Context, input *GetAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetAuthenticationPolicyContract",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyContract{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAuthenticationPolicyContract - Update an Authentication Policy Contract by ID.
//RequestType: PUT
//Input: input *UpdateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) UpdateAuthenticationPolicyContract(input *UpdateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	return s.UpdateAuthenticationPolicyContractWithContext(context.Background(), input)
}

//UpdateAuthenticationPolicyContractWithContext - Update an Authentication Policy Contract by ID.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) UpdateAuthenticationPolicyContractWithContext(ctx context.Context, input *UpdateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateAuthenticationPolicyContract",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyContract{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteAuthenticationPolicyContract - Delete an Authentication Policy Contract.
//RequestType: DELETE
//Input: input *DeleteAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) DeleteAuthenticationPolicyContract(input *DeleteAuthenticationPolicyContractInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteAuthenticationPolicyContractWithContext(context.Background(), input)
}

//DeleteAuthenticationPolicyContractWithContext - Delete an Authentication Policy Contract.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteAuthenticationPolicyContractInput
func (s *AuthenticationPolicyContractsService) DeleteAuthenticationPolicyContractWithContext(ctx context.Context, input *DeleteAuthenticationPolicyContractInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationPolicyContracts/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteAuthenticationPolicyContract",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateAuthenticationPolicyContractInput struct {
	Body models.AuthenticationPolicyContract
}

type DeleteAuthenticationPolicyContractInput struct {
	Id string
}

type GetAuthenticationPolicyContractInput struct {
	Id string
}

type GetAuthenticationPolicyContractsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateAuthenticationPolicyContractInput struct {
	Body models.AuthenticationPolicyContract
	Id   string
}
