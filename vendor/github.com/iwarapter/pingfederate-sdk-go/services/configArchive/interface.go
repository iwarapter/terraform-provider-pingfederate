package configArchive

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigArchiveAPI interface {
	ImportConfigArchive(input *ImportConfigArchiveInput) (output *models.ApiResult, resp *http.Response, err error)
	ImportConfigArchiveWithContext(ctx context.Context, input *ImportConfigArchiveInput) (output *models.ApiResult, resp *http.Response, err error)

	ExportConfigArchive() (resp *http.Response, err error)
	ExportConfigArchiveWithContext(ctx context.Context) (resp *http.Response, err error)
}
