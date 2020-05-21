package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type BulkService service

//ExportConfiguration - Export all API resources to a JSON file.
//RequestType: GET
//Input: input *ExportConfigurationInput
func (s *BulkService) ExportConfiguration(input *ExportConfigurationInput) (result *BulkConfig, resp *http.Response, err error) {
	path := "/bulk/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.IncludeExternalResources != "" {
		q.Set("includeExternalResources", input.IncludeExternalResources)
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

//ImportConfiguration - Import configuration for a PingFederate deployment from a JSON file.
//RequestType: POST
//Input: input *ImportConfigurationInput
func (s *BulkService) ImportConfiguration(input *ImportConfigurationInput) (resp *http.Response, err error) {
	path := "/bulk/import"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.FailFast != "" {
		q.Set("failFast", input.FailFast)
	}
	rel.RawQuery = q.Encode()
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}
