package bulk

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type BulkAPI interface {
	ExportConfiguration(input *ExportConfigurationInput) (output *models.BulkConfig, resp *http.Response, err error)
	ExportConfigurationWithContext(ctx context.Context, input *ExportConfigurationInput) (output *models.BulkConfig, resp *http.Response, err error)

	ImportConfiguration(input *ImportConfigurationInput) (resp *http.Response, err error)
	ImportConfigurationWithContext(ctx context.Context, input *ImportConfigurationInput) (resp *http.Response, err error)
}
