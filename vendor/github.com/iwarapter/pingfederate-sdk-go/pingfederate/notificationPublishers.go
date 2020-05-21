package pingfederate

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type NotificationPublishersService service

//GetSettings - Get general notification publisher settings.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetSettings() (result *NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
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

//UpdateSettings - Update general notification publisher settings.
//RequestType: PUT
//Input: input *UpdateSettingsInput
func (s *NotificationPublishersService) UpdateSettings(input *UpdateSettingsInput) (result *NotificationPublishersSettings, resp *http.Response, err error) {
	path := "/notificationPublishers/settings"
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

//GetNotificationPublisherPluginDescriptors - Get the list of available Notification Publisher Plugin descriptors.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptors() (result *NotificationPublisherDescriptors, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors"
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

//GetNotificationPublisherPluginDescriptor - Get the description of a notification publisher plugin descriptor.
//RequestType: GET
//Input: input *GetNotificationPublisherPluginDescriptorInput
func (s *NotificationPublishersService) GetNotificationPublisherPluginDescriptor(input *GetNotificationPublisherPluginDescriptorInput) (result *NotificationPublisherDescriptor, resp *http.Response, err error) {
	path := "/notificationPublishers/descriptors/{id}"
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

//GetNotificationPublishers - Get a list of notification publisher plugin instances.
//RequestType: GET
//Input:
func (s *NotificationPublishersService) GetNotificationPublishers() (result *NotificationPublishers, resp *http.Response, err error) {
	path := "/notificationPublishers"
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

//CreateNotificationPublisher - Create a notification publisher plugin instance.
//RequestType: POST
//Input: input *CreateNotificationPublisherInput
func (s *NotificationPublishersService) CreateNotificationPublisher(input *CreateNotificationPublisherInput) (result *NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers"
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

//GetNotificationPublisher - Get a specific notification publisher plugin instance.
//RequestType: GET
//Input: input *GetNotificationPublisherInput
func (s *NotificationPublishersService) GetNotificationPublisher(input *GetNotificationPublisherInput) (result *NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//UpdateNotificationPublisher - Update a notification publisher plugin instance.
//RequestType: PUT
//Input: input *UpdateNotificationPublisherInput
func (s *NotificationPublishersService) UpdateNotificationPublisher(input *UpdateNotificationPublisherInput) (result *NotificationPublisher, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//DeleteNotificationPublisher - Delete a notification publisher plugin instance.
//RequestType: DELETE
//Input: input *DeleteNotificationPublisherInput
func (s *NotificationPublishersService) DeleteNotificationPublisher(input *DeleteNotificationPublisherInput) (result *ApiResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}"
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

//GetActions - List the actions for a notification publisher plugin instance.
//RequestType: GET
//Input: input *GetActionsInput
func (s *NotificationPublishersService) GetActions(input *GetActionsInput) (result *Actions, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions"
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

//GetAction - Find an notification publisher plugin instance's action by ID.
//RequestType: GET
//Input: input *GetActionInput
func (s *NotificationPublishersService) GetAction(input *GetActionInput) (result *Action, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

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

//InvokeAction - Invokes an action for notification publisher plugin instance.
//RequestType: POST
//Input: input *InvokeActionInput
func (s *NotificationPublishersService) InvokeAction(input *InvokeActionInput) (result *ActionResult, resp *http.Response, err error) {
	path := "/notificationPublishers/{id}/actions/{actionId}/invokeAction"
	path = strings.Replace(path, "{id}", input.Id, -1)

	path = strings.Replace(path, "{actionId}", input.ActionId, -1)

	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("POST", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}
