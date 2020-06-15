package configStore

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigStoreAPI interface {
	GetSetting(input *GetSettingInput) (result *models.ConfigStoreSetting, resp *http.Response, err error)
	UpdateSetting(input *UpdateSettingInput) (result *models.ConfigStoreSetting, resp *http.Response, err error)
	DeleteSetting(input *DeleteSettingInput) (result *models.ApiResult, resp *http.Response, err error)
	GetSettings(input *GetSettingsInput) (result *models.ConfigStoreBundle, resp *http.Response, err error)
}
