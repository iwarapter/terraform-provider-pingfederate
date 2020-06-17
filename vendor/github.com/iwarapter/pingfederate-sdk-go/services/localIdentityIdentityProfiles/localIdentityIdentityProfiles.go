package localIdentityIdentityProfiles

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
	ServiceName = "LocalIdentityIdentityProfiles"
)

type LocalIdentityIdentityProfilesService struct {
	*client.PfClient
}

// New creates a new instance of the LocalIdentityIdentityProfilesService client.
func New(cfg *config.Config) *LocalIdentityIdentityProfilesService {

	return &LocalIdentityIdentityProfilesService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a LocalIdentityIdentityProfiles operation
func (c *LocalIdentityIdentityProfilesService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetIdentityProfiles - Get the list of configured local identity profiles.
//RequestType: GET
//Input: input *GetIdentityProfilesInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfiles(input *GetIdentityProfilesInput) (output *models.LocalIdentityProfiles, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	op := &request.Operation{
		Name:       "GetIdentityProfiles",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LocalIdentityProfiles{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateIdentityProfile - Create a new local identity profile.
//RequestType: POST
//Input: input *CreateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) CreateIdentityProfile(input *CreateIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	op := &request.Operation{
		Name:       "CreateIdentityProfile",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.LocalIdentityProfile{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetIdentityProfile - Get the local identity profile by ID.
//RequestType: GET
//Input: input *GetIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfile(input *GetIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetIdentityProfile",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.LocalIdentityProfile{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateIdentityProfile - Update the local identity profile by ID.
//RequestType: PUT
//Input: input *UpdateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) UpdateIdentityProfile(input *UpdateIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateIdentityProfile",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.LocalIdentityProfile{}
	req := s.newRequest(op, input.Body, output)
	if input.BypassExternalValidation != nil {
		req.HTTPRequest.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}
	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteIdentityProfile - Delete the local identity profile by ID.
//RequestType: DELETE
//Input: input *DeleteIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) DeleteIdentityProfile(input *DeleteIdentityProfileInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteIdentityProfile",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateIdentityProfileInput struct {
	Body models.LocalIdentityProfile

	BypassExternalValidation *bool
}

type DeleteIdentityProfileInput struct {
	Id string
}

type GetIdentityProfileInput struct {
	Id string
}

type GetIdentityProfilesInput struct {
	Page          string
	NumberPerPage string
	Filter        string
}

type UpdateIdentityProfileInput struct {
	Body models.LocalIdentityProfile
	Id   string

	BypassExternalValidation *bool
}
