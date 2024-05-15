package subscriber

import "context"

const (
	layerSubscriberService = "service.subscriber."
)

type Repository interface {
	CreateTelegramSubscriber(ctx context.Context, chatId uint64) error
	GetSubscriberByChatId(ctx context.Context, chatId uint64) (Subscriber, error)
	CreateDeviceSubscriber(ctx context.Context, deviceId uint64) error
	GetSubscriberByDeviceId(ctx context.Context, deviceId uint64) (Subscriber, error)
	GetSubscriberById(ctx context.Context, id uint64) (Subscriber, error)
	UpdateSubscriber(ctx context.Context, sub Subscriber) error
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateDeviceSubscriber(ctx context.Context, deviceId uint64) error {
	err := s.repository.CreateDeviceSubscriber(ctx, deviceId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateTelegramSubscriber(ctx context.Context, chatId uint64) error {
	err := s.repository.CreateTelegramSubscriber(ctx, chatId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetSubscriberByDeviceId(ctx context.Context, deviceId uint64) (Subscriber, error) {
	sub, err := s.repository.GetSubscriberByDeviceId(ctx, deviceId)
	if err != nil {
		return Subscriber{}, err
	}

	return sub, nil
}

func (s *Service) GetSubscriberByChatId(ctx context.Context, chatId uint64) (Subscriber, error) {
	sub, err := s.repository.GetSubscriberByChatId(ctx, chatId)
	if err != nil {
		return Subscriber{}, err
	}

	return sub, nil
}

func (s *Service) GetSubscriberById(ctx context.Context, id uint64) (Subscriber, error) {
	sub, err := s.repository.GetSubscriberById(ctx, id)
	if err != nil {
		return Subscriber{}, err
	}

	return sub, nil
}

func (s *Service) EnableNotificationSubscriber(ctx context.Context, id uint64, isNotification bool) error {
	byId, err := s.repository.GetSubscriberById(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	sub := Subscriber{
		ID:                  byId.ID,
		GroupId:             byId.GroupId,
		TgChatId:            byId.TgChatId,
		DeviceId:            byId.DeviceId,
		NotificationEnabled: isNotification,
	}

	err = s.repository.UpdateSubscriber(ctx, sub)
	if err != nil {
		return err
	}

	return nil
}
