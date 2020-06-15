package metadataUrls

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type MetadataUrlsService struct {
	Client *client.PfClient
}

// New creates a new instance of the MetadataUrlsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *MetadataUrlsService {

	return &MetadataUrlsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetMetadataUrls - Get a list of Metadata URLs
//RequestType: GET
//Input:
func (s *MetadataUrlsService) GetMetadataUrls() (result *models.MetadataUrls, resp *http.Response, err error) {
	path := "/metadataUrls"
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

//AddMetadataUrl - Add a new Metadata URL.
//RequestType: POST
//Input: input *AddMetadataUrlInput
func (s *MetadataUrlsService) AddMetadataUrl(input *AddMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls"
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

//GetMetadataUrl - Get a Metadata URL by ID.
//RequestType: GET
//Input: input *GetMetadataUrlInput
func (s *MetadataUrlsService) GetMetadataUrl(input *GetMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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

//UpdateMetadataUrl - Update a Metadata URL by ID.
//RequestType: PUT
//Input: input *UpdateMetadataUrlInput
func (s *MetadataUrlsService) UpdateMetadataUrl(input *UpdateMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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

//DeleteMetadataUrl - Delete a Metadata URL by ID.
//RequestType: DELETE
//Input: input *DeleteMetadataUrlInput
func (s *MetadataUrlsService) DeleteMetadataUrl(input *DeleteMetadataUrlInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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

type AddMetadataUrlInput struct {
	Body models.MetadataUrl
}

type DeleteMetadataUrlInput struct {
	Id string
}

type GetMetadataUrlInput struct {
	Id string
}

type UpdateMetadataUrlInput struct {
	Body models.MetadataUrl
	Id   string
}
