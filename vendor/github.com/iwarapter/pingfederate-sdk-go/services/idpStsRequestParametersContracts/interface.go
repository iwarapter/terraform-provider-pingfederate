package idpStsRequestParametersContracts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type IdpStsRequestParametersContractsAPI interface {
	GetStsRequestParamContracts() (result *models.StsRequestParametersContracts, resp *http.Response, err error)
	CreateStsRequestParamContract(input *CreateStsRequestParamContractInput) (result *models.StsRequestParametersContract, resp *http.Response, err error)
	GetStsRequestParamContractById(input *GetStsRequestParamContractByIdInput) (result *models.StsRequestParametersContract, resp *http.Response, err error)
	UpdateStsRequestParamContractById(input *UpdateStsRequestParamContractByIdInput) (result *models.StsRequestParametersContract, resp *http.Response, err error)
	DeleteStsRequestParamContractById(input *DeleteStsRequestParamContractByIdInput) (result *models.ApiResult, resp *http.Response, err error)
}
