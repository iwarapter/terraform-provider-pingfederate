package connectionMetadata

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConnectionMetadataAPI interface {
	Export(input *ExportInput) (result *string, resp *http.Response, err error)
	Convert(input *ConvertInput) (result *models.ConvertMetadataResponse, resp *http.Response, err error)
}
