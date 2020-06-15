package idpAdapters

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpAdaptersAPI interface {
	GetIdpAdapterDescriptors() (result *models.IdpAdapterDescriptors, resp *http.Response, err error)
	GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (result *models.IdpAdapterDescriptor, resp *http.Response, err error)
	GetIdpAdapters(input *GetIdpAdaptersInput) (result *models.IdpAdapters, resp *http.Response, err error)
	CreateIdpAdapter(input *CreateIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error)
	GetIdpAdapter(input *GetIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error)
	UpdateIdpAdapter(input *UpdateIdpAdapterInput) (result *models.IdpAdapter, resp *http.Response, err error)
	DeleteIdpAdapter(input *DeleteIdpAdapterInput) (result *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error)
}
