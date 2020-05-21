package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
)

type ClusterService service

//GetClusterStatus - Get information on the current status of the cluster.
//RequestType: GET
//Input:
func (s *ClusterService) GetClusterStatus() (result *ClusterStatus, resp *http.Response, err error) {
	path := "/cluster/status"
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

//StartReplication - Replicate configuration updates to all nodes in the cluster.
//RequestType: POST
//Input:
func (s *ClusterService) StartReplication() (result *ApiResult, resp *http.Response, err error) {
	path := "/cluster/replicate"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
