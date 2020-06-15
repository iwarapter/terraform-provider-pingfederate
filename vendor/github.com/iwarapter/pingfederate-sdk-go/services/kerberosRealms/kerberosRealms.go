package kerberosRealms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type KerberosRealmsService struct {
	Client *client.PfClient
}

// New creates a new instance of the KerberosRealmsService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *KerberosRealmsService {

	return &KerberosRealmsService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetKerberosRealmSettings - Gets the Kerberos Realms Settings.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealmSettings() (result *models.KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateSettings - Set/Update the Kerberos Realms Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KerberosRealmsService) UpdateSettings(input *UpdateSettingsInput) (result *models.KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetKerberosRealms - Gets the Kerberos Realms.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealms() (result *models.KerberosRealms, resp *http.Response, err error) {
	path := "/kerberos/realms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreateKerberosRealm - Create a new Kerberos Realm.
//RequestType: POST
//Input: input *CreateKerberosRealmInput
func (s *KerberosRealmsService) CreateKerberosRealm(input *CreateKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetKerberosRealm - Find a Kerberos Realm by ID.
//RequestType: GET
//Input: input *GetKerberosRealmInput
func (s *KerberosRealmsService) GetKerberosRealm(input *GetKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateKerberosRealm - Update a Kerberos Realm by ID.
//RequestType: PUT
//Input: input *UpdateKerberosRealmInput
func (s *KerberosRealmsService) UpdateKerberosRealm(input *UpdateKerberosRealmInput) (result *models.KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteKerberosRealm - Delete a Kerberos Realm.
//RequestType: DELETE
//Input: input *DeleteKerberosRealmInput
func (s *KerberosRealmsService) DeleteKerberosRealm(input *DeleteKerberosRealmInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type CreateKerberosRealmInput struct {
	Body models.KerberosRealm
}

type DeleteKerberosRealmInput struct {
	Id string
}

type GetKerberosRealmInput struct {
	Id string
}

type UpdateKerberosRealmInput struct {
	Body models.KerberosRealm
	Id   string
}

type UpdateSettingsInput struct {
	Body models.KerberosRealmsSettings
}
