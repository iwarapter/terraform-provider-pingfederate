package bulk

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type BulkService struct {
	Client *client.PfClient
}

// New creates a new instance of the BulkService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *BulkService {

	return &BulkService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//ExportConfiguration - Export all API resources to a JSON file.
//RequestType: GET
//Input: input *ExportConfigurationInput
func (s *BulkService) ExportConfiguration(input *ExportConfigurationInput) (result *models.BulkConfig, resp *http.Response, err error) {
	path := "/bulk/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	q := rel.Query()
	if input.IncludeExternalResources != "" {
		q.Set("includeExternalResources", input.IncludeExternalResources)
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

//ImportConfiguration - Import configuration for a PingFederate deployment from a JSON file.
//RequestType: POST
//Input: input *ImportConfigurationInput
func (s *BulkService) ImportConfiguration(input *ImportConfigurationInput) (resp *http.Response, err error) {
	path := "/bulk/import"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	q := rel.Query()
	if input.FailFast != "" {
		q.Set("failFast", input.FailFast)
	}
	rel.RawQuery = q.Encode()
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, err
	}
	if input.BypassExternalValidation != nil {
		req.Header.Add("X-BypassExternalValidation", fmt.Sprintf("%v", *input.BypassExternalValidation))
	}

	resp, err = s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type ExportConfigurationInput struct {
	IncludeExternalResources string
}

type ImportConfigurationInput struct {
	FailFast string

	Body models.BulkConfig

	BypassExternalValidation *bool
}
