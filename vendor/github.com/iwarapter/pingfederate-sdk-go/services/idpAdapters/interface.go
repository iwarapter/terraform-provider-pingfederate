package idpAdapters

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpAdaptersAPI interface {
	GetIdpAdapterDescriptors() (output *models.IdpAdapterDescriptors, resp *http.Response, err error)
	GetIdpAdapterDescriptorsWithContext(ctx context.Context) (output *models.IdpAdapterDescriptors, resp *http.Response, err error)

	GetIdpAdapterDescriptorsById(input *GetIdpAdapterDescriptorsByIdInput) (output *models.IdpAdapterDescriptor, resp *http.Response, err error)
	GetIdpAdapterDescriptorsByIdWithContext(ctx context.Context, input *GetIdpAdapterDescriptorsByIdInput) (output *models.IdpAdapterDescriptor, resp *http.Response, err error)

	GetIdpAdapters(input *GetIdpAdaptersInput) (output *models.IdpAdapters, resp *http.Response, err error)
	GetIdpAdaptersWithContext(ctx context.Context, input *GetIdpAdaptersInput) (output *models.IdpAdapters, resp *http.Response, err error)

	CreateIdpAdapter(input *CreateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	CreateIdpAdapterWithContext(ctx context.Context, input *CreateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)

	GetIdpAdapter(input *GetIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	GetIdpAdapterWithContext(ctx context.Context, input *GetIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)

	UpdateIdpAdapter(input *UpdateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)
	UpdateIdpAdapterWithContext(ctx context.Context, input *UpdateIdpAdapterInput) (output *models.IdpAdapter, resp *http.Response, err error)

	DeleteIdpAdapter(input *DeleteIdpAdapterInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteIdpAdapterWithContext(ctx context.Context, input *DeleteIdpAdapterInput) (output *models.ApiResult, resp *http.Response, err error)

	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetActionsWithContext(ctx context.Context, input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)

	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	GetActionWithContext(ctx context.Context, input *GetActionInput) (output *models.Action, resp *http.Response, err error)

	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
	InvokeActionWithContext(ctx context.Context, input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
