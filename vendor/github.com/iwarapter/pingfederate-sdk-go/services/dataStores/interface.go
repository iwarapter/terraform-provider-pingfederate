package dataStores

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type DataStoresAPI interface {
	GetCustomDataStoreDescriptors() (output *models.CustomDataStoreDescriptors, resp *http.Response, err error)
	GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (output *models.CustomDataStoreDescriptor, resp *http.Response, err error)
	GetDataStores() (output *models.DataStores, resp *http.Response, err error)
	CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	CreateLdapDataStore(input *CreateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	CreateCustomDataStore(input *CreateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	GetDataStore(input *GetDataStoreInput) (output *models.DataStore, resp *http.Response, err error)
	DeleteDataStore(input *DeleteDataStoreInput) (output *models.ApiResult, resp *http.Response, err error)
	GetJdbcDataStore(input *GetJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	GetLdapDataStore(input *GetLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	GetCustomDataStore(input *GetCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (output *models.JdbcDataStore, resp *http.Response, err error)
	UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (output *models.LdapDataStore, resp *http.Response, err error)
	UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (output *models.CustomDataStore, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
