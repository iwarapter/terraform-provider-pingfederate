package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConfigArchiveService service

//ImportConfigArchive - Import a configuration archive.
//RequestType: POST
//Input: input *ImportConfigArchiveInput
func (s *ConfigArchiveService) ImportConfigArchive(input *ImportConfigArchiveInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/configArchive/import"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	q := rel.Query()
	if input.ForceImport != "" {
		q.Set("forceImport", input.ForceImport)
	}
	rel.RawQuery = q.Encode()
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

//ExportConfigArchive - Export a configuration archive.
//RequestType: GET
//Input:
func (s *ConfigArchiveService) ExportConfigArchive() (resp *http.Response, err error) {
	path := "/configArchive/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}
