package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MetadataUrlsService service

//GetMetadataUrls - Get a list of Metadata URLs
//RequestType: GET
//Input:
func (s *MetadataUrlsService) GetMetadataUrls() (result *MetadataUrls, resp *http.Response, err error) {
	path := "/metadataUrls"
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

//AddMetadataUrl - Add a new Metadata URL.
//RequestType: POST
//Input: input *AddMetadataUrlInput
func (s *MetadataUrlsService) AddMetadataUrl(input *AddMetadataUrlInput) (result *MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls"
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

//GetMetadataUrl - Get a Metadata URL by ID.
//RequestType: GET
//Input: input *GetMetadataUrlInput
func (s *MetadataUrlsService) GetMetadataUrl(input *GetMetadataUrlInput) (result *MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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

//UpdateMetadataUrl - Update a Metadata URL by ID.
//RequestType: PUT
//Input: input *UpdateMetadataUrlInput
func (s *MetadataUrlsService) UpdateMetadataUrl(input *UpdateMetadataUrlInput) (result *MetadataUrl, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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

//DeleteMetadataUrl - Delete a Metadata URL by ID.
//RequestType: DELETE
//Input: input *DeleteMetadataUrlInput
func (s *MetadataUrlsService) DeleteMetadataUrl(input *DeleteMetadataUrlInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/metadataUrls/{id}"
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
