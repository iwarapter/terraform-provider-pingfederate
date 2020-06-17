package authenticationPolicyContracts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPolicyContractsAPI interface {
	GetAuthenticationPolicyContracts(input *GetAuthenticationPolicyContractsInput) (output *models.AuthenticationPolicyContracts, resp *http.Response, err error)
	CreateAuthenticationPolicyContract(input *CreateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error)
	GetAuthenticationPolicyContract(input *GetAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error)
	UpdateAuthenticationPolicyContract(input *UpdateAuthenticationPolicyContractInput) (output *models.AuthenticationPolicyContract, resp *http.Response, err error)
	DeleteAuthenticationPolicyContract(input *DeleteAuthenticationPolicyContractInput) (output *models.ApiResult, resp *http.Response, err error)
}
