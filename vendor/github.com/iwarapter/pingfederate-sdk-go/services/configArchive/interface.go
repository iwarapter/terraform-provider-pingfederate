package configArchive

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConfigArchiveAPI interface {
	ImportConfigArchive(input *ImportConfigArchiveInput) (output *models.ApiResult, resp *http.Response, err error)
	ExportConfigArchive() (resp *http.Response, err error)
}
