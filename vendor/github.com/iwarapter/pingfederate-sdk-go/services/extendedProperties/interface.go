package extendedProperties

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type ExtendedPropertiesAPI interface {
	GetExtendedProperties() (output *models.ExtendedProperties, resp *http.Response, err error)
	GetExtendedPropertiesWithContext(ctx context.Context) (output *models.ExtendedProperties, resp *http.Response, err error)

	UpdateExtendedProperties(input *UpdateExtendedPropertiesInput) (output *models.ExtendedProperties, resp *http.Response, err error)
	UpdateExtendedPropertiesWithContext(ctx context.Context, input *UpdateExtendedPropertiesInput) (output *models.ExtendedProperties, resp *http.Response, err error)
}
