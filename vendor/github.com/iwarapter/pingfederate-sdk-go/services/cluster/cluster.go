package cluster

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "Cluster"
)

type ClusterService struct {
	*client.PfClient
}

// New creates a new instance of the ClusterService client.
func New(cfg *config.Config) *ClusterService {

	return &ClusterService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a Cluster operation
func (c *ClusterService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetClusterStatus - Get information on the current status of the cluster.
//RequestType: GET
//Input:
func (s *ClusterService) GetClusterStatus() (output *models.ClusterStatus, resp *http.Response, err error) {
	path := "/cluster/status"
	op := &request.Operation{
		Name:       "GetClusterStatus",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ClusterStatus{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//StartReplication - Replicate configuration updates to all nodes in the cluster.
//RequestType: POST
//Input:
func (s *ClusterService) StartReplication() (output *models.ApiResult, resp *http.Response, err error) {
	path := "/cluster/replicate"
	op := &request.Operation{
		Name:       "StartReplication",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}
