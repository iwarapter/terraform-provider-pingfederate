package configArchive

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigArchiveService struct {
	Client *client.PfClient
}

// New creates a new instance of the ConfigArchiveService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ConfigArchiveService {

	return &ConfigArchiveService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//ImportConfigArchive - Import a configuration archive.
//RequestType: POST
//Input: input *ImportConfigArchiveInput
func (s *ConfigArchiveService) ImportConfigArchive(input *ImportConfigArchiveInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/configArchive/import"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	q := rel.Query()
	if input.ForceImport != "" {
		q.Set("forceImport", input.ForceImport)
	}
	rel.RawQuery = q.Encode()
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

//ExportConfigArchive - Export a configuration archive.
//RequestType: GET
//Input:
func (s *ConfigArchiveService) ExportConfigArchive() (resp *http.Response, err error) {
	path := "/configArchive/export"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.Client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type ImportConfigArchiveInput struct {
	ForceImport string

	Body []byte
}
