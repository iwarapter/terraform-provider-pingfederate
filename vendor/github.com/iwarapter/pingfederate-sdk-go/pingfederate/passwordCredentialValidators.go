package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type PasswordCredentialValidatorsService service

//GetPasswordCredentialValidatorDescriptors - Get a list of available password credential validator descriptors.
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptors() (result *PasswordCredentialValidatorDescriptors, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidatorDescriptor - Get the description of a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorDescriptorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidatorDescriptor(input *GetPasswordCredentialValidatorDescriptorInput) (result *PasswordCredentialValidatorDescriptor, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidators - Get the list of available password credential validators
//RequestType: GET
//Input:
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidators() (result *PasswordCredentialValidators, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreatePasswordCredentialValidator - Create a new password credential validator instance
//RequestType: POST
//Input: input *CreatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) CreatePasswordCredentialValidator(input *CreatePasswordCredentialValidatorInput) (result *PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetPasswordCredentialValidator - Find a password credential validator by ID.
//RequestType: GET
//Input: input *GetPasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) GetPasswordCredentialValidator(input *GetPasswordCredentialValidatorInput) (result *PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdatePasswordCredentialValidator - Update a password credential validator instance.
//RequestType: PUT
//Input: input *UpdatePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) UpdatePasswordCredentialValidator(input *UpdatePasswordCredentialValidatorInput) (result *PasswordCredentialValidator, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeletePasswordCredentialValidator - Delete a password credential validator instance.
//RequestType: DELETE
//Input: input *DeletePasswordCredentialValidatorInput
func (s *PasswordCredentialValidatorsService) DeletePasswordCredentialValidator(input *DeletePasswordCredentialValidatorInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/passwordCredentialValidators/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
