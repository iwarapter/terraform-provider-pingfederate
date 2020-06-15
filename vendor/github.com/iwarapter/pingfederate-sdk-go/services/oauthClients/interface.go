package oauthClients

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthClientsAPI interface {
	GetClients(input *GetClientsInput) (result *models.Clients, resp *http.Response, err error)
	CreateClient(input *CreateClientInput) (result *models.Client, resp *http.Response, err error)
	GetClient(input *GetClientInput) (result *models.Client, resp *http.Response, err error)
	UpdateClient(input *UpdateClientInput) (result *models.Client, resp *http.Response, err error)
	DeleteClient(input *DeleteClientInput) (result *models.ApiResult, resp *http.Response, err error)
	GetClientSecret(input *GetClientSecretInput) (result *models.ClientSecret, resp *http.Response, err error)
	UpdateClientSecret(input *UpdateClientSecretInput) (result *models.ClientSecret, resp *http.Response, err error)
}
