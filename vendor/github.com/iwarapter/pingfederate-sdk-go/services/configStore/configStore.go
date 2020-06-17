package configStore

import (
	"net/http"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client/metadata"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/request"
)

const (
	// ServiceName - The name of service.
	ServiceName = "ConfigStore"
)

type ConfigStoreService struct {
	*client.PfClient
}

// New creates a new instance of the ConfigStoreService client.
func New(cfg *config.Config) *ConfigStoreService {

	return &ConfigStoreService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a ConfigStore operation
func (c *ConfigStoreService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSetting - Get a single setting from a bundle.
//RequestType: GET
//Input: input *GetSettingInput
func (s *ConfigStoreService) GetSetting(input *GetSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetSetting",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConfigStoreSetting{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSetting - Create or update a setting/bundle.
//RequestType: PUT
//Input: input *UpdateSettingInput
func (s *ConfigStoreService) UpdateSetting(input *UpdateSettingInput) (output *models.ConfigStoreSetting, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateSetting",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.ConfigStoreSetting{}
	req := s.newRequest(op, input.Body, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteSetting - Delete a setting.
//RequestType: DELETE
//Input: input *DeleteSettingInput
func (s *ConfigStoreService) DeleteSetting(input *DeleteSettingInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/configStore/{bundle}/{id}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteSetting",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetSettings - Get all settings from a bundle.
//RequestType: GET
//Input: input *GetSettingsInput
func (s *ConfigStoreService) GetSettings(input *GetSettingsInput) (output *models.ConfigStoreBundle, resp *http.Response, err error) {
	path := "/configStore/{bundle}"
	path = strings.Replace(path, "{bundle}", input.Bundle, -1)

	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.ConfigStoreBundle{}
	req := s.newRequest(op, nil, output)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type DeleteSettingInput struct {
	Bundle string
	Id     string
}

type GetSettingInput struct {
	Bundle string
	Id     string
}

type GetSettingsInput struct {
	Bundle string
}

type UpdateSettingInput struct {
	Body   models.ConfigStoreSetting
	Bundle string
	Id     string
}
