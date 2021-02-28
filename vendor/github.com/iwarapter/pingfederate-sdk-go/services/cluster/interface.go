package cluster

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ClusterAPI interface {
	GetClusterStatus() (output *models.ClusterStatus, resp *http.Response, err error)
	GetClusterStatusWithContext(ctx context.Context) (output *models.ClusterStatus, resp *http.Response, err error)

	StartReplication() (output *models.ApiResult, resp *http.Response, err error)
	StartReplicationWithContext(ctx context.Context) (output *models.ApiResult, resp *http.Response, err error)
}
