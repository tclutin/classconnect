package postgres

import (
	"github.com/tclutin/classconnect-api/pkg/client/postgresql"
	"log/slog"
)

type Repositories struct {
	User       *UserRepository
	Group      *GroupRepository
	Subscriber *SubscriberRepository
	Schedule   *ScheduleRepository
}

func NewRepositories(client postgresql.Client, logger *slog.Logger) *Repositories {
	return &Repositories{
		User:       NewUserRepository(client, logger),
		Group:      NewGroupRepository(client, logger),
		Subscriber: NewSubscriberRepository(client, logger),
		Schedule:   NewScheduleRepository(client, logger),
	}
}
