package authenticationPolicyContracts

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type AuthenticationPolicyContractsAPI interface {
	GetAuthenticationPolicyContracts(input *GetAuthenticationPolicyContractsInput) (result *models.AuthenticationPolicyContracts, resp *http.Response, err error)
	CreateAuthenticationPolicyContract(input *CreateAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error)
	GetAuthenticationPolicyContract(input *GetAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error)
	UpdateAuthenticationPolicyContract(input *UpdateAuthenticationPolicyContractInput) (result *models.AuthenticationPolicyContract, resp *http.Response, err error)
	DeleteAuthenticationPolicyContract(input *DeleteAuthenticationPolicyContractInput) (result *models.ApiResult, resp *http.Response, err error)
}
