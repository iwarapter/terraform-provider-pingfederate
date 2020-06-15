package dataStores

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type DataStoresAPI interface {
	GetCustomDataStoreDescriptors() (result *models.CustomDataStoreDescriptors, resp *http.Response, err error)
	GetCustomDataStoreDescriptor(input *GetCustomDataStoreDescriptorInput) (result *models.CustomDataStoreDescriptor, resp *http.Response, err error)
	GetDataStores() (result *models.DataStores, resp *http.Response, err error)
	CreateJdbcDataStore(input *CreateJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error)
	CreateLdapDataStore(input *CreateLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error)
	CreateCustomDataStore(input *CreateCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error)
	GetDataStore(input *GetDataStoreInput) (result *models.DataStore, resp *http.Response, err error)
	DeleteDataStore(input *DeleteDataStoreInput) (result *models.ApiResult, resp *http.Response, err error)
	GetJdbcDataStore(input *GetJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error)
	GetLdapDataStore(input *GetLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error)
	GetCustomDataStore(input *GetCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error)
	UpdateJdbcDataStore(input *UpdateJdbcDataStoreInput) (result *models.JdbcDataStore, resp *http.Response, err error)
	UpdateLdapDataStore(input *UpdateLdapDataStoreInput) (result *models.LdapDataStore, resp *http.Response, err error)
	UpdateCustomDataStore(input *UpdateCustomDataStoreInput) (result *models.CustomDataStore, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error)
}
