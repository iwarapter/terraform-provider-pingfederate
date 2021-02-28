package spDefaultUrls

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpDefaultUrlsAPI interface {
	GetDefaultUrls() (output *models.SpDefaultUrls, resp *http.Response, err error)
	GetDefaultUrlsWithContext(ctx context.Context) (output *models.SpDefaultUrls, resp *http.Response, err error)

	UpdateDefaultUrls(input *UpdateDefaultUrlsInput) (output *models.SpDefaultUrls, resp *http.Response, err error)
	UpdateDefaultUrlsWithContext(ctx context.Context, input *UpdateDefaultUrlsInput) (output *models.SpDefaultUrls, resp *http.Response, err error)
}
