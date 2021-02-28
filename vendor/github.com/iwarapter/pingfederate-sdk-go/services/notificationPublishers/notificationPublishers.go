package notificationPublishers

import (
	"context"
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
	ServiceName = "NotificationPublishers"
)

type NotificationPublishersService struct {
	*client.PfClient
}

// New creates a new instance of the NotificationPublishersService client.
func New(cfg *config.Config) *NotificationPublishersService {

	return &NotificationPublishersService{PfClient: client.New(
		*cfg,
		metadata.ClientInfo{
			ServiceName: ServiceName,
			Endpoint:    *cfg.Endpoint,
			APIVersion:  pingfederate.SDKVersion,
		},
	)}
}

// newRequest creates a new request for a NotificationPublishers operation
func (c *NotificationPublishersService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	return req
}

//GetSettings - Get general notification publisher settings.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetSettings() (output *models.NotificationPublishersSettings, resp *http.Response, err error) {
	return s.GetSettingsWithContext(context.Background())
}

//GetSettingsWithContext - Get general notification publisher settings.
//RequestType: GET
//Input: ctx context.Context,
func (s *NotificationPublishersService) GetSettingsWithContext(ctx context.Context) (output *models.NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
	op := &request.Operation{
		Name:       "GetSettings",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationPublishersSettings{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateSettings - Update general notification publisher settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *NotificationPublishersService) UpdateSettings(input *UpdateSettingsInput) (output *models.NotificationPublishersSettings, resp *http.Response, err error) {
	return s.UpdateSettingsWithContext(context.Background(), input)
}

//UpdateSettingsWithContext - Update general notification publisher settings.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateSettingsInput
func (s *NotificationPublishersService) UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
	op := &request.Operation{
		Name:       "UpdateSettings",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.NotificationPublishersSettings{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetNotificationPublisherPluginDescriptors - Get the list of available Notification Publisher Plugin descriptors.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptors() (output *models.NotificationPublisherDescriptors, resp *http.Response, err error) {
	return s.GetNotificationPublisherPluginDescriptorsWithContext(context.Background())
}

//GetNotificationPublisherPluginDescriptorsWithContext - Get the list of available Notification Publisher Plugin descriptors.
//RequestType: GET
//Input: ctx context.Context,
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptorsWithContext(ctx context.Context) (output *models.NotificationPublisherDescriptors, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors"
	op := &request.Operation{
		Name:       "GetNotificationPublisherPluginDescriptors",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationPublisherDescriptors{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetNotificationPublisherPluginDescriptor - Get the description of a notification publisher plugin descriptor.
//RequestType: GET
//Input: input *GetNotificationPublisherPluginDescriptorInput
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptor(input *GetNotificationPublisherPluginDescriptorInput) (output *models.NotificationPublisherDescriptor, resp *http.Response, err error) {
	return s.GetNotificationPublisherPluginDescriptorWithContext(context.Background(), input)
}

//GetNotificationPublisherPluginDescriptorWithContext - Get the description of a notification publisher plugin descriptor.
//RequestType: GET
//Input: ctx context.Context, input *GetNotificationPublisherPluginDescriptorInput
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptorWithContext(ctx context.Context, input *GetNotificationPublisherPluginDescriptorInput) (output *models.NotificationPublisherDescriptor, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetNotificationPublisherPluginDescriptor",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationPublisherDescriptor{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetNotificationPublishers - Get a list of notification publisher plugin instances.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublishers() (output *models.NotificationPublishers, resp *http.Response, err error) {
	return s.GetNotificationPublishersWithContext(context.Background())
}

//GetNotificationPublishersWithContext - Get a list of notification publisher plugin instances.
//RequestType: GET
//Input: ctx context.Context,
func (s *NotificationPublishersService) GetNotificationPublishersWithContext(ctx context.Context) (output *models.NotificationPublishers, resp *http.Response, err error) {
	path := "/notificationPublishers"
	op := &request.Operation{
		Name:       "GetNotificationPublishers",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationPublishers{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//CreateNotificationPublisher - Create a notification publisher plugin instance.
//RequestType: POST
//Input: input *CreateNotificationPublisherInput
func (s *NotificationPublishersService) CreateNotificationPublisher(input *CreateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	return s.CreateNotificationPublisherWithContext(context.Background(), input)
}

//CreateNotificationPublisherWithContext - Create a notification publisher plugin instance.
//RequestType: POST
//Input: ctx context.Context, input *CreateNotificationPublisherInput
func (s *NotificationPublishersService) CreateNotificationPublisherWithContext(ctx context.Context, input *CreateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers"
	op := &request.Operation{
		Name:       "CreateNotificationPublisher",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.NotificationPublisher{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetNotificationPublisher - Get a specific notification publisher plugin instance.
//RequestType: GET
//Input: input *GetNotificationPublisherInput
func (s *NotificationPublishersService) GetNotificationPublisher(input *GetNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	return s.GetNotificationPublisherWithContext(context.Background(), input)
}

//GetNotificationPublisherWithContext - Get a specific notification publisher plugin instance.
//RequestType: GET
//Input: ctx context.Context, input *GetNotificationPublisherInput
func (s *NotificationPublishersService) GetNotificationPublisherWithContext(ctx context.Context, input *GetNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetNotificationPublisher",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.NotificationPublisher{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//UpdateNotificationPublisher - Update a notification publisher plugin instance.
//RequestType: PUT
//Input: input *UpdateNotificationPublisherInput
func (s *NotificationPublishersService) UpdateNotificationPublisher(input *UpdateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	return s.UpdateNotificationPublisherWithContext(context.Background(), input)
}

//UpdateNotificationPublisherWithContext - Update a notification publisher plugin instance.
//RequestType: PUT
//Input: ctx context.Context, input *UpdateNotificationPublisherInput
func (s *NotificationPublishersService) UpdateNotificationPublisherWithContext(ctx context.Context, input *UpdateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "UpdateNotificationPublisher",
		HTTPMethod: "PUT",
		HTTPPath:   path,
	}
	output = &models.NotificationPublisher{}
	req := s.newRequest(op, input.Body, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//DeleteNotificationPublisher - Delete a notification publisher plugin instance.
//RequestType: DELETE
//Input: input *DeleteNotificationPublisherInput
func (s *NotificationPublishersService) DeleteNotificationPublisher(input *DeleteNotificationPublisherInput) (output *models.ApiResult, resp *http.Response, err error) {
	return s.DeleteNotificationPublisherWithContext(context.Background(), input)
}

//DeleteNotificationPublisherWithContext - Delete a notification publisher plugin instance.
//RequestType: DELETE
//Input: ctx context.Context, input *DeleteNotificationPublisherInput
func (s *NotificationPublishersService) DeleteNotificationPublisherWithContext(ctx context.Context, input *DeleteNotificationPublisherInput) (output *models.ApiResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "DeleteNotificationPublisher",
		HTTPMethod: "DELETE",
		HTTPPath:   path,
	}

	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetActions - List the actions for a notification publisher plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *NotificationPublishersService) GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	return s.GetActionsWithContext(context.Background(), input)
}

//GetActionsWithContext - List the actions for a notification publisher plugin instance.
//RequestType: GET
//Input: ctx context.Context, input *GetActionsInput
func (s *NotificationPublishersService) GetActionsWithContext(ctx context.Context, input *GetActionsInput) (output *models.Actions, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions"
	path = strings.Replace(path, "{id}", input.Id, -1)

	op := &request.Operation{
		Name:       "GetActions",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Actions{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//GetAction - Find an notification publisher plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *NotificationPublishersService) GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	return s.GetActionWithContext(context.Background(), input)
}

//GetActionWithContext - Find an notification publisher plugin instance's action by ID.
//RequestType: GET
//Input: ctx context.Context, input *GetActionInput
func (s *NotificationPublishersService) GetActionWithContext(ctx context.Context, input *GetActionInput) (output *models.Action, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "GetAction",
		HTTPMethod: "GET",
		HTTPPath:   path,
	}
	output = &models.Action{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

//InvokeAction - Invokes an action for notification publisher plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *NotificationPublishersService) InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	return s.InvokeActionWithContext(context.Background(), input)
}

//InvokeActionWithContext - Invokes an action for notification publisher plugin instance.
//RequestType: POST
//Input: ctx context.Context, input *InvokeActionInput
func (s *NotificationPublishersService) InvokeActionWithContext(ctx context.Context, input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	op := &request.Operation{
		Name:       "InvokeAction",
		HTTPMethod: "POST",
		HTTPPath:   path,
	}
	output = &models.ActionResult{}
	req := s.newRequest(op, nil, output)
	req.HTTPRequest = req.HTTPRequest.WithContext(ctx)

	if req.Send() == nil {
		return output, req.HTTPResponse, nil
	}
	return nil, req.HTTPResponse, req.Error
}

type CreateNotificationPublisherInput struct {
	Body models.NotificationPublisher
}

type DeleteNotificationPublisherInput struct {
	Id string
}

type GetActionInput struct {
	Id       string
	ActionId string
}

type GetActionsInput struct {
	Id string
}

type GetNotificationPublisherInput struct {
	Id string
}

type GetNotificationPublisherPluginDescriptorInput struct {
	Id string
}

type InvokeActionInput struct {
	Id       string
	ActionId string
}

type UpdateNotificationPublisherInput struct {
	Body models.NotificationPublisher
	Id   string
}

type UpdateSettingsInput struct {
	Body models.NotificationPublishersSettings
}
