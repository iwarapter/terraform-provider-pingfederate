package extendedProperties

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ExtendedPropertiesAPI interface {
	GetExtendedProperties() (output *models.ExtendedProperties, resp *http.Response, err error)
	UpdateExtendedProperties(input *UpdateExtendedPropertiesInput) (output *models.ExtendedProperties, resp *http.Response, err error)
}
