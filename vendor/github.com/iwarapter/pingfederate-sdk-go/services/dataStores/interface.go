package dataStores

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type DataStoresAPI interface {
	GetCustomDataStoreDescriptors() (output *models.CustomDataStoreDescriptors, resp *http.Response, err error)
	GetCustomDataStoreDescriptorsWithContext(ctx context.Context) (output *models.CustomDataStoreDescriptors, resp *http.Response, err error)

	GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (output *models.CustomDataStoreDescriptor, resp *http.Response, err error)
	GetCustomDataStoreDescriptorWithContext(ctx context.Context, input *GetCustomDataStoreDescriptorInput) (output *models.CustomDataStoreDescriptor, resp *http.Response, err error)

	GetDataStores() (output *models.DataStores, resp *http.Response, err error)
	GetDataStoresWithContext(ctx context.Context) (output *models.DataStores, resp *http.Response, err error)

	CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	CreateJdbcDataStoreWithContext(ctx context.Context, input *CreateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)

	CreateLdapDataStore(input *CreateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	CreateLdapDataStoreWithContext(ctx context.Context, input *CreateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)

	CreateCustomDataStore(input *CreateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	CreateCustomDataStoreWithContext(ctx context.Context, input *CreateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)

	GetDataStore(input *GetDataStoreInput) (output *models.DataStore, resp *http.Response, err error)
	GetDataStoreWithContext(ctx context.Context, input *GetDataStoreInput) (output *models.DataStore, resp *http.Response, err error)

	DeleteDataStore(input *DeleteDataStoreInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteDataStoreWithContext(ctx context.Context, input *DeleteDataStoreInput) (output *models.ApiResult, resp *http.Response, err error)

	GetJdbcDataStore(input *GetJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	GetJdbcDataStoreWithContext(ctx context.Context, input *GetJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)

	GetLdapDataStore(input *GetLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	GetLdapDataStoreWithContext(ctx context.Context, input *GetLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)

	GetCustomDataStore(input *GetCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	GetCustomDataStoreWithContext(ctx context.Context, input *GetCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)

	UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	UpdateJdbcDataStoreWithContext(ctx context.Context, input *UpdateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)

	UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	UpdateLdapDataStoreWithContext(ctx context.Context, input *UpdateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)

	UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	UpdateCustomDataStoreWithContext(ctx context.Context, input *UpdateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)

	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetActionsWithContext(ctx context.Context, input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)

	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	GetActionWithContext(ctx context.Context, input *GetActionInput) (output *models.Action, resp *http.Response, err error)

	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
	InvokeActionWithContext(ctx context.Context, input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
