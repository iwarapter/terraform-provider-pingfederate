package connectionMetadata

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ConnectionMetadataAPI interface {
	Export(input *ExportInput) (output *string, resp *http.Response, err error)
	Convert(input *ConvertInput) (output *models.ConvertMetadataResponse, resp *http.Response, err error)
}
