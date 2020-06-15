package passwordCredentialValidators

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PasswordCredentialValidatorsService struct {
	Client *client.PfClient
}

// New creates a new instance of the PasswordCredentialValidatorsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *PasswordCredentialValidatorsService {

	return &PasswordCredentialValidatorsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetPasswordCredentialValidatorDescriptors - Get a list of available password credential validator descriptors.
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptors() (result *models.PasswordCredentialValidatorDescriptors, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidatorDescriptor - Get the description of a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorDescriptorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (result *models.PasswordCredentialValidatorDescriptor, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidators - Get the list of available password credential validators
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidators() (result *models.PasswordCredentialValidators, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreatePasswordCredentialValidator - Create a new password credential validator instance
//RequestType: POST
//Input: input *CreatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidator - Find a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePasswordCredentialValidator - Update a password credential validator instance.
//RequestType: PUT
//Input: input *UpdatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (result *models.PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePasswordCredentialValidator - Delete a password credential validator instance.
//RequestType: DELETE
//Input: input *DeletePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type CreatePasswordCredentialValidatorInput struct {
	Body models.PasswordCredentialValidator
}

type DeletePasswordCredentialValidatorInput struct {
	Id string
}

type GetPasswordCredentialValidatorInput struct {
	Id string
}

type GetPasswordCredentialValidatorDescriptorInput struct {
	Id string
}

type UpdatePasswordCredentialValidatorInput struct {
	Body models.PasswordCredentialValidator
	Id   string
}
