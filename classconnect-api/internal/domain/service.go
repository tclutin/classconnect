package domain

import (
	"github.com/tclutin/classconnect-api/internal/config"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/internal/domain/group"
	"github.com/tclutin/classconnect-api/internal/domain/schedule"
	"github.com/tclutin/classconnect-api/internal/domain/subscriber"
	"github.com/tclutin/classconnect-api/internal/repository/postgres"
)

type Services struct {
	Auth       *auth.Service
	Group      *group.Service
	Subscriber *subscriber.Service
	Schedule   *schedule.Service
}

func NewServices(config *config.Config, repositories *postgres.Repositories) *Services {
	authService := auth.NewService(config, repositories.User)
	groupService := group.NewService(repositories.Group, repositories.User, repositories.Subscriber, repositories.Schedule)
	subscriberService := subscriber.NewService(repositories.Subscriber)
	scheduleService := schedule.NewService(repositories.Schedule, repositories.User, repositories.Group, repositories.Subscriber)
	return &Services{
		Auth:       authService,
		Group:      groupService,
		Subscriber: subscriberService,
		Schedule:   scheduleService,
	}
}
