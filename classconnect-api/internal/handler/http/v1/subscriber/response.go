package subscriber

import "github.com/tclutin/classconnect-api/internal/domain/subscriber"

type SubscriberResponse struct {
	ID                  uint64  `json:"id"`
	GroupId             *uint64 `json:"group_id"`
	TgChatId            *uint64 `json:"tg_chat_id"`
	DeviceId            *uint64 `json:"device_id"`
	NotificationEnabled bool    `json:"notification_enabled"`
}

func ConvertSubscriberEntityToResponse(sub subscriber.Subscriber) SubscriberResponse {
	return SubscriberResponse{
		ID:                  sub.ID,
		GroupId:             sub.GroupId,
		TgChatId:            sub.TgChatId,
		DeviceId:            sub.DeviceId,
		NotificationEnabled: sub.NotificationEnabled,
	}
}
