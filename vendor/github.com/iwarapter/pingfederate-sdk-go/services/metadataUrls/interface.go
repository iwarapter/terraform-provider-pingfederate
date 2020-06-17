package metadataUrls

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type MetadataUrlsAPI interface {
	GetMetadataUrls() (output *models.MetadataUrls, resp *http.Response, err error)
	AddMetadataUrl(input *AddMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	GetMetadataUrl(input *GetMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	UpdateMetadataUrl(input *UpdateMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	DeleteMetadataUrl(input *DeleteMetadataUrlInput) (output *models.ApiResult, resp *http.Response, err error)
}
