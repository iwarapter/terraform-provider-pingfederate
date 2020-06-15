package idpConnectors

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpConnectorsAPI interface {
	GetIdpConnectorDescriptors() (result *models.SaasPluginDescriptors, resp *http.Response, err error)
	GetIdpConnectorDescriptorById(input *GetIdpConnectorDescriptorByIdInput) (result *models.SaasPluginDescriptor, resp *http.Response, err error)
}
