package spTargetUrlMappings

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTargetUrlMappingsAPI interface {
	GetUrlMappings() (result *models.SpUrlMappings, resp *http.Response, err error)
	UpdateUrlMappings(input *UpdateUrlMappingsInput) (result *models.SpUrlMappings, resp *http.Response, err error)
}
