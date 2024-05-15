package postgres

import (
	"context"
	"github.com/tclutin/classconnect-api/internal/domain/subscriber"
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"log/slog"
)

const (
	layerSubscriberRepository = "repository.subscriber."
)

type SubscriberRepository struct {
	db     postgresql.Client
	logger *slog.Logger
}

func NewSubscriberRepository(client postgresql.Client, logger *slog.Logger) *SubscriberRepository {
	return &SubscriberRepository{
		db:     client,
		logger: logger,
	}
}

func (s *SubscriberRepository) CreateTelegramSubscriber(ctx context.Context, chatId uint64) error {
	sql := `INSERT INTO public.subscribers (telegram_chat_id) VALUES ($1)`

	_, err := s.db.Exec(ctx, sql, chatId)

	return err
}

func (s *SubscriberRepository) CreateDeviceSubscriber(ctx context.Context, deviceId uint64) error {
	sql := `INSERT INTO public.subscribers (device_id) VALUES ($1)`

	_, err := s.db.Exec(ctx, sql, deviceId)

	return err
}

func (s *SubscriberRepository) GetSubscriberByDeviceId(ctx context.Context, deviceId uint64) (subscriber.Subscriber, error) {
	sql := `SELECT * FROM public.subscribers WHERE device_id = $1`

	row := s.db.QueryRow(ctx, sql, deviceId)

	var getSubscriber subscriber.Subscriber

	err := row.Scan(
		&getSubscriber.ID,
		&getSubscriber.GroupId,
		&getSubscriber.TgChatId,
		&getSubscriber.DeviceId,
		&getSubscriber.NotificationEnabled)

	if err != nil {
		return subscriber.Subscriber{}, err
	}

	return getSubscriber, nil
}

func (s *SubscriberRepository) GetSubscriberByChatId(ctx context.Context, chatId uint64) (subscriber.Subscriber, error) {
	sql := `SELECT * FROM public.subscribers WHERE telegram_chat_id = $1`

	row := s.db.QueryRow(ctx, sql, chatId)

	var getSubscriber subscriber.Subscriber

	err := row.Scan(
		&getSubscriber.ID,
		&getSubscriber.GroupId,
		&getSubscriber.TgChatId,
		&getSubscriber.DeviceId,
		&getSubscriber.NotificationEnabled)

	if err != nil {
		return subscriber.Subscriber{}, err
	}

	return getSubscriber, nil
}

func (s *SubscriberRepository) GetSubscriberById(ctx context.Context, id uint64) (subscriber.Subscriber, error) {
	sql := `SELECT * FROM public.subscribers WHERE id = $1`

	row := s.db.QueryRow(ctx, sql, id)

	var getSubscriber subscriber.Subscriber

	err := row.Scan(
		&getSubscriber.ID,
		&getSubscriber.GroupId,
		&getSubscriber.TgChatId,
		&getSubscriber.DeviceId,
		&getSubscriber.NotificationEnabled)

	if err != nil {
		return subscriber.Subscriber{}, err
	}

	return getSubscriber, nil
}

func (s *SubscriberRepository) UpdateSubscriber(ctx context.Context, sub subscriber.Subscriber) error {
	sql := `UPDATE public.subscribers SET group_id = $1,
                              telegram_chat_id = $2,
                              device_id = $3,
                              notification_enabled = $4 WHERE id = $5`

	_, err := s.db.Exec(ctx, sql,
		sub.GroupId,
		sub.TgChatId,
		sub.DeviceId,
		sub.NotificationEnabled,
		sub.ID)

	return err
}
