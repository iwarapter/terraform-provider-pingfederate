package bulk

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type BulkAPI interface {
	ExportConfiguration(input *ExportConfigurationInput) (result *models.BulkConfig, resp *http.Response, err error)
	ImportConfiguration(input *ImportConfigurationInput) (resp *http.Response, err error)
}
