package spTargetUrlMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTargetUrlMappingsAPI interface {
	GetUrlMappings() (output *models.SpUrlMappings, resp *http.Response, err error)
	UpdateUrlMappings(input *UpdateUrlMappingsInput) (output *models.SpUrlMappings, resp *http.Response, err error)
}
