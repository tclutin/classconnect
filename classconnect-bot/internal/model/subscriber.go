package model

type CreateTelegramSubscriberRequest struct {
	ChatId uint64 `json:"chat_id"`
}

type EnableNotificationSubscriberRequest struct {
	Notification bool `json:"notification"`
}

type SubscriberRequest struct {
	SubID uint64 `json:"sub_id"`
}

type SubscriberResponse struct {
	ID                  uint64  `json:"id"`
	GroupId             *uint64 `json:"group_id"`
	TgChatId            *uint64 `json:"tg_chat_id"`
	DeviceId            *uint64 `json:"device_id"`
	NotificationEnabled bool    `json:"notification_enabled"`
}
