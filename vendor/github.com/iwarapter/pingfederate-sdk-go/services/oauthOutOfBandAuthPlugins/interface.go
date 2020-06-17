package oauthOutOfBandAuthPlugins

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOutOfBandAuthPluginsAPI interface {
	GetOOBAuthPluginDescriptors() (output *models.OutOfBandAuthPluginDescriptors, resp *http.Response, err error)
	GetOOBAuthPluginDescriptor(input *GetOOBAuthPluginDescriptorInput) (output *models.OutOfBandAuthPluginDescriptor, resp *http.Response, err error)
	GetOOBAuthenticators() (output *models.OutOfBandAuthenticators, resp *http.Response, err error)
	CreateOOBAuthenticator(input *CreateOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error)
	GetOOBAuthenticator(input *GetOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error)
	UpdateOOBAuthenticator(input *UpdateOOBAuthenticatorInput) (output *models.OutOfBandAuthenticator, resp *http.Response, err error)
	DeleteOOBAuthenticator(input *DeleteOOBAuthenticatorInput) (output *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
