package oauthResourceOwnerCredentialsMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthResourceOwnerCredentialsMappingsAPI interface {
	GetResourceOwnerCredentialsMappings() (output *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error)
	GetResourceOwnerCredentialsMappingsWithContext(ctx context.Context) (output *models.ResourceOwnerCredentialsMappings, resp *http.Response, err error)

	CreateResourceOwnerCredentialsMapping(input *CreateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	CreateResourceOwnerCredentialsMappingWithContext(ctx context.Context, input *CreateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)

	GetResourceOwnerCredentialsMapping(input *GetResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	GetResourceOwnerCredentialsMappingWithContext(ctx context.Context, input *GetResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)

	UpdateResourceOwnerCredentialsMapping(input *UpdateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)
	UpdateResourceOwnerCredentialsMappingWithContext(ctx context.Context, input *UpdateResourceOwnerCredentialsMappingInput) (output *models.ResourceOwnerCredentialsMapping, resp *http.Response, err error)

	DeleteResourceOwnerCredentialsMapping(input *DeleteResourceOwnerCredentialsMappingInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteResourceOwnerCredentialsMappingWithContext(ctx context.Context, input *DeleteResourceOwnerCredentialsMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
