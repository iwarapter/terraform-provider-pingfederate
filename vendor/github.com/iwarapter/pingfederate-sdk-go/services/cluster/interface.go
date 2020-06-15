package cluster

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ClusterAPI interface {
	GetClusterStatus() (result *models.ClusterStatus, resp *http.Response, err error)
	StartReplication() (result *models.ApiResult, resp *http.Response, err error)
}
