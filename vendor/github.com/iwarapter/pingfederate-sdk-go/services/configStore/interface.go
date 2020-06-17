package configStore

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigStoreAPI interface {
	GetSetting(input *GetSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)
	UpdateSetting(input *UpdateSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)
	DeleteSetting(input *DeleteSettingInput) (output *models.ApiResult, resp *http.Response, err error)
	GetSettings(input *GetSettingsInput) (output *models.ConfigStoreBundle, resp *http.Response, err error)
}
