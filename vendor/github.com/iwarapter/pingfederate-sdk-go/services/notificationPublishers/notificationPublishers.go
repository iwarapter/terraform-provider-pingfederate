package notificationPublishers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/client"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type NotificationPublishersService struct {
	Client *client.PfClient
}

// New creates a new instance of the NotificationPublishersService client.
func New(username string, password string, baseUrl *url.URL, context string, httpClient *http.Client) *NotificationPublishersService {

	return &NotificationPublishersService{Client: client.NewClient(username, password, baseUrl, context, httpClient)}
}

//GetSettings - Get general notification publisher settings.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetSettings() (result *models.NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
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

//UpdateSettings - Update general notification publisher settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *NotificationPublishersService) UpdateSettings(input *UpdateSettingsInput) (result *models.NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
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

//GetNotificationPublisherPluginDescriptors - Get the list of available Notification Publisher Plugin descriptors.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptors() (result *models.NotificationPublisherDescriptors, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors"
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

//GetNotificationPublisherPluginDescriptor - Get the description of a notification publisher plugin descriptor.
//RequestType: GET
//Input: input *GetNotificationPublisherPluginDescriptorInput
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptor(input *GetNotificationPublisherPluginDescriptorInput) (result *models.NotificationPublisherDescriptor, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors/{id}"
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

//GetNotificationPublishers - Get a list of notification publisher plugin instances.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublishers() (result *models.NotificationPublishers, resp *http.Response, err error) {
	path := "/notificationPublishers"
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

//CreateNotificationPublisher - Create a notification publisher plugin instance.
//RequestType: POST
//Input: input *CreateNotificationPublisherInput
func (s *NotificationPublishersService) CreateNotificationPublisher(input *CreateNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers"
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

//GetNotificationPublisher - Get a specific notification publisher plugin instance.
//RequestType: GET
//Input: input *GetNotificationPublisherInput
func (s *NotificationPublishersService) GetNotificationPublisher(input *GetNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//UpdateNotificationPublisher - Update a notification publisher plugin instance.
//RequestType: PUT
//Input: input *UpdateNotificationPublisherInput
func (s *NotificationPublishersService) UpdateNotificationPublisher(input *UpdateNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//DeleteNotificationPublisher - Delete a notification publisher plugin instance.
//RequestType: DELETE
//Input: input *DeleteNotificationPublisherInput
func (s *NotificationPublishersService) DeleteNotificationPublisher(input *DeleteNotificationPublisherInput) (result *models.ApiResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//GetActions - List the actions for a notification publisher plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *NotificationPublishersService) GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions"
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

//GetAction - Find an notification publisher plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *NotificationPublishersService) GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

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

//InvokeAction - Invokes an action for notification publisher plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *NotificationPublishersService) InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.Client.Context, path)}
	req, err := s.Client.NewRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.Client.Do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

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
