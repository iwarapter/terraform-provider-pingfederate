package oauthAuthenticationPolicyContractMappings

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type OauthAuthenticationPolicyContractMappingsAPI interface {
	GetApcMappings() (output *models.ApcToPersistentGrantMappings, resp *http.Response, err error)
	GetApcMappingsWithContext(ctx context.Context) (output *models.ApcToPersistentGrantMappings, resp *http.Response, err error)

	CreateApcMapping(input *CreateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	CreateApcMappingWithContext(ctx context.Context, input *CreateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)

	GetApcMapping(input *GetApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	GetApcMappingWithContext(ctx context.Context, input *GetApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)

	UpdateApcMapping(input *UpdateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)
	UpdateApcMappingWithContext(ctx context.Context, input *UpdateApcMappingInput) (output *models.ApcToPersistentGrantMapping, resp *http.Response, err error)

	DeleteApcMapping(input *DeleteApcMappingInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteApcMappingWithContext(ctx context.Context, input *DeleteApcMappingInput) (output *models.ApiResult, resp *http.Response, err error)
}
