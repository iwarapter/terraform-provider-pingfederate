package virtualHostNames

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type VirtualHostNamesAPI interface {
	GetVirtualHostNamesSettings() (output *models.VirtualHostNameSettings, resp *http.Response, err error)
	UpdateVirtualHostNamesSettings(input *UpdateVirtualHostNamesSettingsInput) (output *models.VirtualHostNameSettings, resp *http.Response, err error)
}
