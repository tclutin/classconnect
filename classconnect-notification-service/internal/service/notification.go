package service

import (
	"classconnect-notification-service/pkg/client/telegram"
	"log/slog"
)

type NotificationService struct {
	tgClient telegram.Client
	logger   *slog.Logger
}

func NewNotificationService(tgClient telegram.Client, logger *slog.Logger) *NotificationService {
	return &NotificationService{tgClient: tgClient, logger: logger}
}

func (n *NotificationService) Send(chatId uint64, message string) {
	n.logger.Debug("Sending message", slog.Uint64("chatID", chatId), slog.String("message", message))

	if err := n.tgClient.Send(chatId, message); err != nil {
		n.logger.Error("error sending message: ", slog.Any("error", err))
	}
}
