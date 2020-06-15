package metadataUrls

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type MetadataUrlsAPI interface {
	GetMetadataUrls() (result *models.MetadataUrls, resp *http.Response, err error)
	AddMetadataUrl(input *AddMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error)
	GetMetadataUrl(input *GetMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error)
	UpdateMetadataUrl(input *UpdateMetadataUrlInput) (result *models.MetadataUrl, resp *http.Response, err error)
	DeleteMetadataUrl(input *DeleteMetadataUrlInput) (result *models.ApiResult, resp *http.Response, err error)
}
