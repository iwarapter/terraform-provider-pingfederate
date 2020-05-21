package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type LocalIdentityIdentityProfilesService service

//GetIdentityProfiles - Get the list of configured local identity profiles.
//RequestType: GET
//Input: input *GetIdentityProfilesInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfiles(input *GetIdentityProfilesInput) (result *LocalIdentityProfiles, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.Page != "" {
		q.Set("page", input.Page)
	}
	if input.NumberPerPage != "" {
		q.Set("numberPerPage", input.NumberPerPage)
	}
	if input.Filter != "" {
		q.Set("filter", input.Filter)
	}
	rel.RawQuery = q.Encode()
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

//CreateIdentityProfile - Create a new local identity profile.
//RequestType: POST
//Input: input *CreateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) CreateIdentityProfile(input *CreateIdentityProfileInput) (result *LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetIdentityProfile - Get the local identity profile by ID.
//RequestType: GET
//Input: input *GetIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfile(input *GetIdentityProfileInput) (result *LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
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

//UpdateIdentityProfile - Update the local identity profile by ID.
//RequestType: PUT
//Input: input *UpdateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) UpdateIdentityProfile(input *UpdateIdentityProfileInput) (result *LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteIdentityProfile - Delete the local identity profile by ID.
//RequestType: DELETE
//Input: input *DeleteIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) DeleteIdentityProfile(input *DeleteIdentityProfileInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
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
