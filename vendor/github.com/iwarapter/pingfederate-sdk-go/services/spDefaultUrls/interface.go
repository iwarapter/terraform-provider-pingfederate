package spDefaultUrls

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpDefaultUrlsAPI interface {
	GetDefaultUrls() (result *models.SpDefaultUrls, resp *http.Response, err error)
	UpdateDefaultUrls(input *UpdateDefaultUrlsInput) (result *models.SpDefaultUrls, resp *http.Response, err error)
}
