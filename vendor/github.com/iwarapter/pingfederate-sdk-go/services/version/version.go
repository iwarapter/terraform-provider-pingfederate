package version

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type VersionService struct {
	Client *client.PfClient
}

// New creates a new instance of the VersionService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *VersionService {

	return &VersionService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetVersion - Gets the server version.
//RequestType: GET
//Input:
func (s *VersionService) GetVersion() (result *models.Version, resp *http.Response, err error) {
	path := "/version"
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
