package metadataUrls

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type MetadataUrlsAPI interface {
	GetMetadataUrls() (output *models.MetadataUrls, resp *http.Response, err error)
	GetMetadataUrlsWithContext(ctx context.Context) (output *models.MetadataUrls, resp *http.Response, err error)

	AddMetadataUrl(input *AddMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	AddMetadataUrlWithContext(ctx context.Context, input *AddMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)

	GetMetadataUrl(input *GetMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	GetMetadataUrlWithContext(ctx context.Context, input *GetMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)

	UpdateMetadataUrl(input *UpdateMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)
	UpdateMetadataUrlWithContext(ctx context.Context, input *UpdateMetadataUrlInput) (output *models.MetadataUrl, resp *http.Response, err error)

	DeleteMetadataUrl(input *DeleteMetadataUrlInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteMetadataUrlWithContext(ctx context.Context, input *DeleteMetadataUrlInput) (output *models.ApiResult, resp *http.Response, err error)
}
