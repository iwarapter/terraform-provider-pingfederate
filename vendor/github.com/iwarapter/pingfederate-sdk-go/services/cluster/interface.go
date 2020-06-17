package cluster

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ClusterAPI interface {
	GetClusterStatus() (output *models.ClusterStatus, resp *http.Response, err error)
	StartReplication() (output *models.ApiResult, resp *http.Response, err error)
}
