package oauthOutOfBandAuthPlugins

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthOutOfBandAuthPluginsAPI interface {
	GetOOBAuthPluginDescriptors() (result *models.OutOfBandAuthPluginDescriptors, resp *http.Response, err error)
	GetOOBAuthPluginDescriptor(input *GetOOBAuthPluginDescriptorInput) (result *models.OutOfBandAuthPluginDescriptor, resp *http.Response, err error)
	GetOOBAuthenticators() (result *models.OutOfBandAuthenticators, resp *http.Response, err error)
	CreateOOBAuthenticator(input *CreateOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error)
	GetOOBAuthenticator(input *GetOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error)
	UpdateOOBAuthenticator(input *UpdateOOBAuthenticatorInput) (result *models.OutOfBandAuthenticator, resp *http.Response, err error)
	DeleteOOBAuthenticator(input *DeleteOOBAuthenticatorInput) (result *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error)
}
