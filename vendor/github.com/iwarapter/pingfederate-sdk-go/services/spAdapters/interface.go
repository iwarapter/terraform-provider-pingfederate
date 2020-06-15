package spAdapters

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAdaptersAPI interface {
	GetSpAdapterDescriptors() (result *models.SpAdapterDescriptors, resp *http.Response, err error)
	GetSpAdapterDescriptorsById(input *GetSpAdapterDescriptorsByIdInput) (result *models.SpAdapterDescriptor, resp *http.Response, err error)
	GetSpAdapters(input *GetSpAdaptersInput) (result *models.SpAdapters, resp *http.Response, err error)
	CreateSpAdapter(input *CreateSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error)
	GetSpAdapter(input *GetSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error)
	UpdateSpAdapter(input *UpdateSpAdapterInput) (result *models.SpAdapter, resp *http.Response, err error)
	DeleteSpAdapter(input *DeleteSpAdapterInput) (result *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error)
	GetUrlMappings() (result *models.SpAdapterUrlMappings, resp *http.Response, err error)
	UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *models.SpAdapterUrlMappings, resp *http.Response, err error)
}
