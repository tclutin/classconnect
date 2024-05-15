package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/tclutin/classconnect-api/internal/domain"
	"github.com/tclutin/classconnect-api/internal/handler/http/v1/auth"
	"github.com/tclutin/classconnect-api/internal/handler/http/v1/group"
	"github.com/tclutin/classconnect-api/internal/handler/http/v1/schedule"
	"github.com/tclutin/classconnect-api/internal/handler/http/v1/subscriber"
	"log/slog"
)

type Handler struct {
	services *domain.Services
	logger   *slog.Logger
}

func NewHandler(services *domain.Services, logger *slog.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitAPI(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		auth.NewHandler(h.services.Auth, h.logger).InitAPI(v1, h.services.Auth)
		group.NewHandler(h.services.Group, h.logger).InitAPI(v1, h.services.Auth)
		subscriber.NewHandler(h.services.Subscriber, h.logger).InitAPI(v1, h.services.Auth)
		schedule.NewHandler(h.services.Schedule, h.logger).InitAPI(v1, h.services.Auth)
	}
}
