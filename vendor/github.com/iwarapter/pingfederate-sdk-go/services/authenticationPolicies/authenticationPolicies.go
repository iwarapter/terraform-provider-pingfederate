package authenticationPolicies

import (
	"context"
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
	ServiceName = "AuthenticationPolicies"
)

type AuthenticationPoliciesService struct {
	*client.PfClient
}

// New creates a new instance of the AuthenticationPoliciesService client.
func New(cfg *config.Config) *AuthenticationPoliciesService {

	return &AuthenticationPoliciesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a AuthenticationPolicies operation
func (c *AuthenticationPoliciesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get the authentication policies settings.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetSettings() (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	return s.GetSettingsWithContext(context.Background())
}

//GetSettingsWithContext - Get the authentication policies settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *AuthenticationPoliciesService) GetSettingsWithContext(ctx context.Context) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPoliciesSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Set the authentication policies settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *AuthenticationPoliciesService) UpdateSettings(input *UpdateSettingsInput) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Set the authentication policies settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *AuthenticationPoliciesService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.AuthenticationPoliciesSettings, resp *http.Response, err error) {
	path := "/authenticationPolicies/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPoliciesSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetDefaultAuthenticationPolicy - Get the default configured authentication policy.
//RequestType: GET
//Input:
func (s *AuthenticationPoliciesService) GetDefaultAuthenticationPolicy() (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	return s.GetDefaultAuthenticationPolicyWithContext(context.Background())
}

//GetDefaultAuthenticationPolicyWithContext - Get the default configured authentication policy.
//RequestType: GET
//Input: ctx context.Context,
func (s *AuthenticationPoliciesService) GetDefaultAuthenticationPolicyWithContext(ctx context.Context) (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
	op := &request.Operation{
		Name:       "GetDefaultAuthenticationPolicy",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicy{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateDefaultAuthenticationPolicy - Set the default authentication policy.
//RequestType: PUT
//Input: input *UpdateDefaultAuthenticationPolicyInput
func (s *AuthenticationPoliciesService) UpdateDefaultAuthenticationPolicy(input *UpdateDefaultAuthenticationPolicyInput) (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	return s.UpdateDefaultAuthenticationPolicyWithContext(context.Background(), input)
}

//UpdateDefaultAuthenticationPolicyWithContext - Set the default authentication policy.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateDefaultAuthenticationPolicyInput
func (s *AuthenticationPoliciesService) UpdateDefaultAuthenticationPolicyWithContext(ctx context.Context, input *UpdateDefaultAuthenticationPolicyInput) (output *models.AuthenticationPolicy, resp *http.Response, err error) {
	path := "/authenticationPolicies/default"
	op := &request.Operation{
		Name:       "UpdateDefaultAuthenticationPolicy",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicy{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetFragments - Get all of the authentication policies fragments.
//RequestType: GET
//Input: input *GetFragmentsInput
func (s *AuthenticationPoliciesService) GetFragments(input *GetFragmentsInput) (output *models.AuthenticationPolicyFragments, resp *http.Response, err error) {
	return s.GetFragmentsWithContext(context.Background(), input)
}

//GetFragmentsWithContext - Get all of the authentication policies fragments.
//RequestType: GET
//Input: ctx context.Context, input *GetFragmentsInput
func (s *AuthenticationPoliciesService) GetFragmentsWithContext(ctx context.Context, input *GetFragmentsInput) (output *models.AuthenticationPolicyFragments, resp *http.Response, err error) {
	path := "/authenticationPolicies/fragments"
	op := &request.Operation{
		Name:       "GetFragments",
		HTTPMethod: "GET",
		HTTPPath:   path,
		QueryParams: map[string]string{
			"page":          input.Page,
			"numberPerPage": input.NumberPerPage,
			"filter":        input.Filter,
		},
	}
	output = &models.AuthenticationPolicyFragments{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateFragment - Create an authentication policy fragment.
//RequestType: POST
//Input: input *CreateFragmentInput
func (s *AuthenticationPoliciesService) CreateFragment(input *CreateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	return s.CreateFragmentWithContext(context.Background(), input)
}

//CreateFragmentWithContext - Create an authentication policy fragment.
//RequestType: POST
//Input: ctx context.Context, input *CreateFragmentInput
func (s *AuthenticationPoliciesService) CreateFragmentWithContext(ctx context.Context, input *CreateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	path := "/authenticationPolicies/fragments"
	op := &request.Operation{
		Name:       "CreateFragment",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyFragment{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetFragment - Get an authentication policy fragment by ID.
//RequestType: GET
//Input: input *GetFragmentInput
func (s *AuthenticationPoliciesService) GetFragment(input *GetFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	return s.GetFragmentWithContext(context.Background(), input)
}

//GetFragmentWithContext - Get an authentication policy fragment by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetFragmentInput
func (s *AuthenticationPoliciesService) GetFragmentWithContext(ctx context.Context, input *GetFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	path := "/authenticationPolicies/fragments/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetFragment",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyFragment{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateFragment - Update an authentication policy fragment.
//RequestType: PUT
//Input: input *UpdateFragmentInput
func (s *AuthenticationPoliciesService) UpdateFragment(input *UpdateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	return s.UpdateFragmentWithContext(context.Background(), input)
}

//UpdateFragmentWithContext - Update an authentication policy fragment.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateFragmentInput
func (s *AuthenticationPoliciesService) UpdateFragmentWithContext(ctx context.Context, input *UpdateFragmentInput) (output *models.AuthenticationPolicyFragment, resp *http.Response, err error) {
	path := "/authenticationPolicies/fragments/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateFragment",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.AuthenticationPolicyFragment{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteFragment - Delete an authentication policy fragment.
//RequestType: DELETE
//Input: input *DeleteFragmentInput
func (s *AuthenticationPoliciesService) DeleteFragment(input *DeleteFragmentInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteFragmentWithContext(context.Background(), input)
}

//DeleteFragmentWithContext - Delete an authentication policy fragment.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteFragmentInput
func (s *AuthenticationPoliciesService) DeleteFragmentWithContext(ctx context.Context, input *DeleteFragmentInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/authenticationPolicies/fragments/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteFragment",
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

type CreateFragmentInput struct {
	Body models.AuthenticationPolicyFragment

	BypassExternalValidation *bool
}

type DeleteFragmentInput struct {
	Id string
}

type GetFragmentInput struct {
	Id string
}

type GetFragmentsInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateDefaultAuthenticationPolicyInput struct {
	Body models.AuthenticationPolicy

	BypassExternalValidation *bool
}

type UpdateFragmentInput struct {
	Body models.AuthenticationPolicyFragment
	Id   string

	BypassExternalValidation *bool
}

type UpdateSettingsInput struct {
	Body models.AuthenticationPoliciesSettings
}
