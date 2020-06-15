package notificationPublishers

import (
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type NotificationPublishersAPI interface {
	GetSettings() (result *models.NotificationPublishersSettings, resp *http.Response, err error)
	UpdateSettings(input *UpdateSettingsInput) (result *models.NotificationPublishersSettings, resp *http.Response, err error)
	GetNotificationPublisherPluginDescriptors() (result *models.NotificationPublisherDescriptors, resp *http.Response, err error)
	GetNotificationPublisherPluginDescriptor(input *GetNotificationPublisherPluginDescriptorInput) (result *models.NotificationPublisherDescriptor, resp *http.Response, err error)
	GetNotificationPublishers() (result *models.NotificationPublishers, resp *http.Response, err error)
	CreateNotificationPublisher(input *CreateNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error)
	GetNotificationPublisher(input *GetNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error)
	UpdateNotificationPublisher(input *UpdateNotificationPublisherInput) (result *models.NotificationPublisher, resp *http.Response, err error)
	DeleteNotificationPublisher(input *DeleteNotificationPublisherInput) (result *models.ApiResult, resp *http.Response, err error)
	GetActions(input *GetActionsInput) (result *models.Actions, resp *http.Response, err error)
	GetAction(input *GetActionInput) (result *models.Action, resp *http.Response, err error)
	InvokeAction(input *InvokeActionInput) (result *models.ActionResult, resp *http.Response, err error)
}
