package oauthClients

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientsAPI interface {
	GetClients(input *GetClientsInput) (output *models.Clients, resp *http.Response, err error)
	CreateClient(input *CreateClientInput) (output *models.Client, resp *http.Response, err error)
	GetClient(input *GetClientInput) (output *models.Client, resp *http.Response, err error)
	UpdateClient(input *UpdateClientInput) (output *models.Client, resp *http.Response, err error)
	DeleteClient(input *DeleteClientInput) (output *models.ApiResult, resp *http.Response, err error)
	GetClientSecret(input *GetClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)
	UpdateClientSecret(input *UpdateClientSecretInput) (output *models.ClientSecret, resp *http.Response, err error)
}
