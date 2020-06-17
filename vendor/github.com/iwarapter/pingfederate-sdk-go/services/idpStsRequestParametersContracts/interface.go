package idpStsRequestParametersContracts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpStsRequestParametersContractsAPI interface {
	GetStsRequestParamContracts() (output *models.StsRequestParametersContracts, resp *http.Response, err error)
	CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (output *models.StsRequestParametersContract, resp *http.Response, err error)
	DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (output *models.ApiResult, resp *http.Response, err error)
}
