package cluster

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ClusterService struct {
	Client *client.PfClient
}

// New creates a new instance of the ClusterService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *ClusterService {

	return &ClusterService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetClusterStatus - Get information on the current status of the cluster.
//RequestType: GET
//Input:
func (s *ClusterService) GetClusterStatus() (result *models.ClusterStatus, resp *http.Response, err error) {
	path := "/cluster/status"
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

//StartReplication - Replicate configuration updates to all nodes in the cluster.
//RequestType: POST
//Input:
func (s *ClusterService) StartReplication() (result *models.ApiResult, resp *http.Response, err error) {
	path := "/cluster/replicate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
