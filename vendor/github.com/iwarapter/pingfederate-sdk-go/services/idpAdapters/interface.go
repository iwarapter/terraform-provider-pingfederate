package idpAdapters

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpAdaptersAPI interface {
	GetIdpAdapterDescriptors() (output *models.IdpAdapterDescriptors, resp *http.Response, err error)
	GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (output *models.IdpAdapterDescriptor, resp *http.Response, err error)
	GetIdpAdapters(input *GetIdpAdaptersInput) (output *models.IdpAdapters, resp *http.Response, err error)
	CreateIdpAdapter(input *CreateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	GetIdpAdapter(input *GetIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	UpdateIdpAdapter(input *UpdateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	DeleteIdpAdapter(input *DeleteIdpAdapterInput) (output *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
