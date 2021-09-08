package pingOneConnections

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type PingOneConnectionsAPI interface {
	GetPingOneConnections() (output *models.PingOneConnections, resp *http.Response, err error)
	GetPingOneConnectionsWithContext(ctx context.Context) (output *models.PingOneConnections, resp *http.Response, err error)

	CreatePingOneConnection(input *CreatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)
	CreatePingOneConnectionWithContext(ctx context.Context, input *CreatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)

	GetPingOneConnection(input *GetPingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)
	GetPingOneConnectionWithContext(ctx context.Context, input *GetPingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)

	UpdatePingOneConnection(input *UpdatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)
	UpdatePingOneConnectionWithContext(ctx context.Context, input *UpdatePingOneConnectionInput) (output *models.PingOneConnection, resp *http.Response, err error)

	DeletePingOneConnection(input *DeletePingOneConnectionInput) (output *models.ApiResult, resp *http.Response, err error)
	DeletePingOneConnectionWithContext(ctx context.Context, input *DeletePingOneConnectionInput) (output *models.ApiResult, resp *http.Response, err error)

	GetCredentialStatus(input *GetCredentialStatusInput) (output *models.PingOneCredentialStatus, resp *http.Response, err error)
	GetCredentialStatusWithContext(ctx context.Context, input *GetCredentialStatusInput) (output *models.PingOneCredentialStatus, resp *http.Response, err error)

	GetPingOneConnectionEnvironments(input *GetPingOneConnectionEnvironmentsInput) (output *models.PingOneEnvironments, resp *http.Response, err error)
	GetPingOneConnectionEnvironmentsWithContext(ctx context.Context, input *GetPingOneConnectionEnvironmentsInput) (output *models.PingOneEnvironments, resp *http.Response, err error)

	GetPingOneConnectionUsages(input *GetPingOneConnectionUsagesInput) (output *models.ResourceUsages, resp *http.Response, err error)
	GetPingOneConnectionUsagesWithContext(ctx context.Context, input *GetPingOneConnectionUsagesInput) (output *models.ResourceUsages, resp *http.Response, err error)

	GetPingOneConnectionAssociations(input *GetPingOneConnectionAssociationsInput) (output *models.ServiceAssociations, resp *http.Response, err error)
	GetPingOneConnectionAssociationsWithContext(ctx context.Context, input *GetPingOneConnectionAssociationsInput) (output *models.ServiceAssociations, resp *http.Response, err error)
}
