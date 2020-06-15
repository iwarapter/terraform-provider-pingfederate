package localIdentityIdentityProfiles

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LocalIdentityIdentityProfilesService struct {
	Client *client.PfClient
}

// New creates a new instance of the LocalIdentityIdentityProfilesService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *LocalIdentityIdentityProfilesService {

	return &LocalIdentityIdentityProfilesService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetIdentityProfiles - Get the list of configured local identity profiles.
//RequestType: GET
//Input: input *GetIdentityProfilesInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfiles(input *GetIdentityProfilesInput) (result *models.LocalIdentityProfiles, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
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

//CreateIdentityProfile - Create a new local identity profile.
//RequestType: POST
//Input: input *CreateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) CreateIdentityProfile(input *CreateIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetIdentityProfile - Get the local identity profile by ID.
//RequestType: GET
//Input: input *GetIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) GetIdentityProfile(input *GetIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
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

//UpdateIdentityProfile - Update the local identity profile by ID.
//RequestType: PUT
//Input: input *UpdateIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) UpdateIdentityProfile(input *UpdateIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteIdentityProfile - Delete the local identity profile by ID.
//RequestType: DELETE
//Input: input *DeleteIdentityProfileInput
func (s *LocalIdentityIdentityProfilesService) DeleteIdentityProfile(input *DeleteIdentityProfileInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/localIdentity/identityProfiles/{id}"
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
