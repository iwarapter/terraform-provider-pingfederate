package localIdentityIdentityProfiles

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type LocalIdentityIdentityProfilesAPI interface {
	GetIdentityProfiles(input *GetIdentityProfilesInput) (output *models.LocalIdentityProfiles, resp *http.Response, err error)
	CreateIdentityProfile(input *CreateIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error)
	GetIdentityProfile(input *GetIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error)
	UpdateIdentityProfile(input *UpdateIdentityProfileInput) (output *models.LocalIdentityProfile, resp *http.Response, err error)
	DeleteIdentityProfile(input *DeleteIdentityProfileInput) (output *models.ApiResult, resp *http.Response, err error)
}
