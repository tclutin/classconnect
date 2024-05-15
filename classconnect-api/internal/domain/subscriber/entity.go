package subscriber

import "errors"

var (
	ErrNotFound       = errors.New("subscriber not found")
	ErrNotExistsGroup = errors.New("subscriber does not have a group")
)

type Subscriber struct {
	ID                  uint64
	GroupId             *uint64
	TgChatId            *uint64
	DeviceId            *uint64
	NotificationEnabled bool
}
