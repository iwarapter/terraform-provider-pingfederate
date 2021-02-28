package idpConnectors

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpConnectorsAPI interface {
	GetIdpConnectorDescriptors() (output *models.SaasPluginDescriptors, resp *http.Response, err error)
	GetIdpConnectorDescriptorsWithContext(ctx context.Context) (output *models.SaasPluginDescriptors, resp *http.Response, err error)

	GetIdpConnectorDescriptorById(input *GetIdpConnectorDescriptorByIdInput) (output *models.SaasPluginDescriptor, resp *http.Response, err error)
	GetIdpConnectorDescriptorByIdWithContext(ctx context.Context, input *GetIdpConnectorDescriptorByIdInput) (output *models.SaasPluginDescriptor, resp *http.Response, err error)
}
