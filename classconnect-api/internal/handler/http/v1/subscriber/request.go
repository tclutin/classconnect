package subscriber

type CreateTelegramSubscriberRequest struct {
	ChatId uint64 `json:"chat_id" binding:"required"`
}

type CreateDeviceSubscriberRequest struct {
	DeviceId uint64 `json:"device_id" binding:"required"`
}

type EnableNotificationSubscriberRequest struct {
	Notification bool `json:"notification" binding:"required"`
}
