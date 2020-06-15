package version

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type VersionAPI interface {
	GetVersion() (result *models.Version, resp *http.Response, err error)
}
