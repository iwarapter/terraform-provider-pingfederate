package spTargetUrlMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpTargetUrlMappingsAPI interface {
	GetUrlMappings() (output *models.SpUrlMappings, resp *http.Response, err error)
	GetUrlMappingsWithContext(ctx context.Context) (output *models.SpUrlMappings, resp *http.Response, err error)

	UpdateUrlMappings(input *UpdateUrlMappingsInput) (output *models.SpUrlMappings, resp *http.Response, err error)
	UpdateUrlMappingsWithContext(ctx context.Context, input *UpdateUrlMappingsInput) (output *models.SpUrlMappings, resp *http.Response, err error)
}
