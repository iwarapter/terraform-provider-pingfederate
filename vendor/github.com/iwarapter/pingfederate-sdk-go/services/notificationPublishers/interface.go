package notificationPublishers

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type NotificationPublishersAPI interface {
	GetSettings() (output *models.NotificationPublishersSettings, resp *http.Response, err error)
	GetSettingsWithContext(ctx context.Context) (output *models.NotificationPublishersSettings, resp *http.Response, err error)

	UpdateSettings(input *UpdateSettingsInput) (output *models.NotificationPublishersSettings, resp *http.Response, err error)
	UpdateSettingsWithContext(ctx context.Context, input *UpdateSettingsInput) (output *models.NotificationPublishersSettings, resp *http.Response, err error)

	GetNotificationPublisherPluginDescriptors() (output *models.NotificationPublisherDescriptors, resp *http.Response, err error)
	GetNotificationPublisherPluginDescriptorsWithContext(ctx context.Context) (output *models.NotificationPublisherDescriptors, resp *http.Response, err error)

	GetNotificationPublisherPluginDescriptor(input *GetNotificationPublisherPluginDescriptorInput) (output *models.NotificationPublisherDescriptor, resp *http.Response, err error)
	GetNotificationPublisherPluginDescriptorWithContext(ctx context.Context, input *GetNotificationPublisherPluginDescriptorInput) (output *models.NotificationPublisherDescriptor, resp *http.Response, err error)

	GetNotificationPublishers() (output *models.NotificationPublishers, resp *http.Response, err error)
	GetNotificationPublishersWithContext(ctx context.Context) (output *models.NotificationPublishers, resp *http.Response, err error)

	CreateNotificationPublisher(input *CreateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)
	CreateNotificationPublisherWithContext(ctx context.Context, input *CreateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)

	GetNotificationPublisher(input *GetNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)
	GetNotificationPublisherWithContext(ctx context.Context, input *GetNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)

	UpdateNotificationPublisher(input *UpdateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)
	UpdateNotificationPublisherWithContext(ctx context.Context, input *UpdateNotificationPublisherInput) (output *models.NotificationPublisher, resp *http.Response, err error)

	DeleteNotificationPublisher(input *DeleteNotificationPublisherInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteNotificationPublisherWithContext(ctx context.Context, input *DeleteNotificationPublisherInput) (output *models.ApiResult, resp *http.Response, err error)

	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetActionsWithContext(ctx context.Context, input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)

	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	GetActionWithContext(ctx context.Context, input *GetActionInput) (output *models.Action, resp *http.Response, err error)

	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
	InvokeActionWithContext(ctx context.Context, input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
}
