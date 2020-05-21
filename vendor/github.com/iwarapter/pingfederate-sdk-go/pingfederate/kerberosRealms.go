package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type KerberosRealmsService service

//GetKerberosRealmSettings - Gets the Kerberos Realms Settings.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealmSettings() (result *KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateSettings - Set/Update the Kerberos Realms Settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *KerberosRealmsService) UpdateSettings(input *UpdateSettingsInput) (result *KerberosRealmsSettings, resp *http.Response, err error) {
	path := "/kerberos/realms/settings"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetKerberosRealms - Gets the Kerberos Realms.
//RequestType: GET
//Input:
func (s *KerberosRealmsService) GetKerberosRealms() (result *KerberosRealms, resp *http.Response, err error) {
	path := "/kerberos/realms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//CreateKerberosRealm - Create a new Kerberos Realm.
//RequestType: POST
//Input: input *CreateKerberosRealmInput
func (s *KerberosRealmsService) CreateKerberosRealm(input *CreateKerberosRealmInput) (result *KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//GetKerberosRealm - Find a Kerberos Realm by ID.
//RequestType: GET
//Input: input *GetKerberosRealmInput
func (s *KerberosRealmsService) GetKerberosRealm(input *GetKerberosRealmInput) (result *KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateKerberosRealm - Update a Kerberos Realm by ID.
//RequestType: PUT
//Input: input *UpdateKerberosRealmInput
func (s *KerberosRealmsService) UpdateKerberosRealm(input *UpdateKerberosRealmInput) (result *KerberosRealm, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//DeleteKerberosRealm - Delete a Kerberos Realm.
//RequestType: DELETE
//Input: input *DeleteKerberosRealmInput
func (s *KerberosRealmsService) DeleteKerberosRealm(input *DeleteKerberosRealmInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/kerberos/realms/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
