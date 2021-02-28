package idpStsRequestParametersContracts

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpStsRequestParametersContractsAPI interface {
	GetStsRequestParamContracts() (output *models.StsRequestParametersContracts, resp *http.Response, err error)
	GetStsRequestParamContractsWithContext(ctx context.Context) (output *models.StsRequestParametersContracts, resp *http.Response, err error)

	CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	CreateStsRequestParamContractWithContext(ctx context.Context, input *CreateStsRequestParamContractInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)

	GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	GetStsRequestParamContractByIdWithContext(ctx context.Context, input *GetStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)

	UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	UpdateStsRequestParamContractByIdWithContext(ctx context.Context, input *UpdateStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)

	DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteStsRequestParamContractByIdWithContext(ctx context.Context, input *DeleteStsRequestParamContractByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
