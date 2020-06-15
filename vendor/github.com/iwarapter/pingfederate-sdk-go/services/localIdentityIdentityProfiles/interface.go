package localIdentityIdentityProfiles

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LocalIdentityIdentityProfilesAPI interface {
	GetIdentityProfiles(input *GetIdentityProfilesInput) (result *models.LocalIdentityProfiles, resp *http.Response, err error)
	CreateIdentityProfile(input *CreateIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error)
	GetIdentityProfile(input *GetIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error)
	UpdateIdentityProfile(input *UpdateIdentityProfileInput) (result *models.LocalIdentityProfile, resp *http.Response, err error)
	DeleteIdentityProfile(input *DeleteIdentityProfileInput) (result *models.ApiResult, resp *http.Response, err error)
}
