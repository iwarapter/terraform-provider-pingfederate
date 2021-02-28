package administrativeAccounts

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
	ServiceName = "AdministrativeAccounts"
)

type AdministrativeAccountsService struct {
	*client.PfClient
}

// New creates a new instance of the AdministrativeAccountsService client.
func New(cfg *config.Config) *AdministrativeAccountsService {

	return &AdministrativeAccountsService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AdministrativeAccounts operation
func (c *AdministrativeAccountsService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetAccounts - Get all the PingFederate native Administrative Accounts.
//RequestType: GET
//Input:
func (s *AdministrativeAccountsService) GetAccounts() (output *models.AdministrativeAccounts, resp *http.Response, err error) {
	return s.GetAccountsWithContext(context.Background())
}

//GetAccountsWithContext - Get all the PingFederate native Administrative Accounts.
//RequestType: GET
//Input: ctx context.Context,
func (s *AdministrativeAccountsService) GetAccountsWithContext(ctx context.Context) (output *models.AdministrativeAccounts, resp *http.Response, err error) {
	path := "/administrativeAccounts"
	op := &request.Operation{
		Name:       "GetAccounts",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdministrativeAccounts{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//AddAccount - Add a new PingFederate native Administrative Account.
//RequestType: POST
//Input: input *AddAccountInput
func (s *AdministrativeAccountsService) AddAccount(input *AddAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	return s.AddAccountWithContext(context.Background(), input)
}

//AddAccountWithContext - Add a new PingFederate native Administrative Account.
//RequestType: POST
//Input: ctx context.Context, input *AddAccountInput
func (s *AdministrativeAccountsService) AddAccountWithContext(ctx context.Context, input *AddAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts"
	op := &request.Operation{
		Name:       "AddAccount",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AdministrativeAccount{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAccount - Get a PingFederate native Administrative Account.
//RequestType: GET
//Input: input *GetAccountInput
func (s *AdministrativeAccountsService) GetAccount(input *GetAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	return s.GetAccountWithContext(context.Background(), input)
}

//GetAccountWithContext - Get a PingFederate native Administrative Account.
//RequestType: GET
//Input: ctx context.Context, input *GetAccountInput
func (s *AdministrativeAccountsService) GetAccountWithContext(ctx context.Context, input *GetAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

	op := &request.Operation{
		Name:       "GetAccount",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AdministrativeAccount{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateAccount - Update the information for a native Administrative Account.
//RequestType: PUT
//Input: input *UpdateAccountInput
func (s *AdministrativeAccountsService) UpdateAccount(input *UpdateAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	return s.UpdateAccountWithContext(context.Background(), input)
}

//UpdateAccountWithContext - Update the information for a native Administrative Account.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateAccountInput
func (s *AdministrativeAccountsService) UpdateAccountWithContext(ctx context.Context, input *UpdateAccountInput) (output *models.AdministrativeAccount, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

	op := &request.Operation{
		Name:       "UpdateAccount",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AdministrativeAccount{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteAccount - Delete a PingFederate native Administrative Account information.
//RequestType: DELETE
//Input: input *DeleteAccountInput
func (s *AdministrativeAccountsService) DeleteAccount(input *DeleteAccountInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteAccountWithContext(context.Background(), input)
}

//DeleteAccountWithContext - Delete a PingFederate native Administrative Account information.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteAccountInput
func (s *AdministrativeAccountsService) DeleteAccountWithContext(ctx context.Context, input *DeleteAccountInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}"
	path = strings.Replace(path, "{username}", input.Username, -1)

	op := &request.Operation{
		Name:       "DeleteAccount",
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

//ResetPassword - Reset the Password of an existing PingFederate native Administrative Account.
//RequestType: POST
//Input: input *ResetPasswordInput
func (s *AdministrativeAccountsService) ResetPassword(input *ResetPasswordInput) (output *models.UserCredentials, resp *http.Response, err error) {
	return s.ResetPasswordWithContext(context.Background(), input)
}

//ResetPasswordWithContext - Reset the Password of an existing PingFederate native Administrative Account.
//RequestType: POST
//Input: ctx context.Context, input *ResetPasswordInput
func (s *AdministrativeAccountsService) ResetPasswordWithContext(ctx context.Context, input *ResetPasswordInput) (output *models.UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/{username}/resetPassword"
	path = strings.Replace(path, "{username}", input.Username, -1)

	op := &request.Operation{
		Name:       "ResetPassword",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.UserCredentials{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//ChangePassword - Change the Password of current PingFederate native Account.
//RequestType: POST
//Input: input *ChangePasswordInput
func (s *AdministrativeAccountsService) ChangePassword(input *ChangePasswordInput) (output *models.UserCredentials, resp *http.Response, err error) {
	return s.ChangePasswordWithContext(context.Background(), input)
}

//ChangePasswordWithContext - Change the Password of current PingFederate native Account.
//RequestType: POST
//Input: ctx context.Context, input *ChangePasswordInput
func (s *AdministrativeAccountsService) ChangePasswordWithContext(ctx context.Context, input *ChangePasswordInput) (output *models.UserCredentials, resp *http.Response, err error) {
	path := "/administrativeAccounts/changePassword"
	op := &request.Operation{
		Name:       "ChangePassword",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.UserCredentials{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type AddAccountInput struct {
	Body models.AdministrativeAccount
}

type ChangePasswordInput struct {
	Body models.UserCredentials
}

type DeleteAccountInput struct {
	Username string
}

type GetAccountInput struct {
	Username string
}

type ResetPasswordInput struct {
	Body     models.UserCredentials
	Username string
}

type UpdateAccountInput struct {
	Body     models.AdministrativeAccount
	Username string
}
