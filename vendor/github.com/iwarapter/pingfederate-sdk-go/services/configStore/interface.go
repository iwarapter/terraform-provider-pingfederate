package configStore

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigStoreAPI interface {
	GetSetting(input *GetSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)
	GetSettingWithContext(ctx context.Context, input *GetSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)

	UpdateSetting(input *UpdateSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)
	UpdateSettingWithContext(ctx context.Context, input *UpdateSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error)

	DeleteSetting(input *DeleteSettingInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteSettingWithContext(ctx context.Context, input *DeleteSettingInput) (output *models.ApiResult, resp *http.Response, err error)

	GetSettings(input *GetSettingsInput) (output *models.ConfigStoreBundle, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context, input *GetSettingsInput) (output *models.ConfigStoreBundle, resp *http.Response, err error)
}
