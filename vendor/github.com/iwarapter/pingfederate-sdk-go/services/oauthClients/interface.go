package oauthClients

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientsAPI interface {
	GetClients(input *GetClientsInput) (output *models.Clients, resp *http.Response, err error)
	GetClientsWithContext(ctx context.Context, input *GetClientsInput) (output *models.Clients, resp *http.Response, err error)

	CreateClient(input *CreateClientInput) (output *models.Client, resp *http.Response, err error)
	CreateClientWithContext(ctx context.Context, input *CreateClientInput) (output *models.Client, resp *http.Response, err error)

	GetClient(input *GetClientInput) (output *models.Client, resp *http.Response, err error)
	GetClientWithContext(ctx context.Context, input *GetClientInput) (output *models.Client, resp *http.Response, err error)

	UpdateClient(input *UpdateClientInput) (output *models.Client, resp *http.Response, err error)
	UpdateClientWithContext(ctx context.Context, input *UpdateClientInput) (output *models.Client, resp *http.Response, err error)

	DeleteClient(input *DeleteClientInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteClientWithContext(ctx context.Context, input *DeleteClientInput) (output *models.ApiResult, resp *http.Response, err error)

	GetClientSecret(input *GetClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)
	GetClientSecretWithContext(ctx context.Context, input *GetClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)

	UpdateClientSecret(input *UpdateClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)
	UpdateClientSecretWithContext(ctx context.Context, input *UpdateClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)
}
