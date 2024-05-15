package subscriber

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/internal/domain/subscriber"
	"github.com/tclutin/classconnect-api/internal/handler/http/middleware"
	resp "github.com/tclutin/classconnect-api/pkg/http"
	"log/slog"
	"net/http"
	"strconv"
)

const (
	layerSubscriberHandler = "handler.subscriber."
)

type Service interface {
	CreateDeviceSubscriber(ctx context.Context, deviceId uint64) error
	CreateTelegramSubscriber(ctx context.Context, deviceId uint64) error
	EnableNotificationSubscriber(ctx context.Context, subId uint64, isNotification bool) error
	GetSubscriberByDeviceId(ctx context.Context, deviceId uint64) (subscriber.Subscriber, error)
	GetSubscriberByChatId(ctx context.Context, chatId uint64) (subscriber.Subscriber, error)
}

type Handler struct {
	service Service
	logger  *slog.Logger
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitAPI(router *gin.RouterGroup, auth *auth.Service) {
	subscriberGroup := router.Group("/subscribers", middleware.AuthMiddleware(auth))
	{
		subscriberGroup.POST("/device", h.CreateDeviceSubscriber)
		subscriberGroup.POST("/telegram", h.CreateTelegramSubscriber)
		subscriberGroup.GET("/device/:deviceId", h.GetSubscriberByDeviceId)
		subscriberGroup.GET("/telegram/:chatId", h.GetSubscriberChatId)
		subscriberGroup.PATCH("/:subscriberId", h.EnableNotification)
	}
}

func (h *Handler) CreateTelegramSubscriber(c *gin.Context) {
	var request CreateTelegramSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	err := h.service.CreateTelegramSubscriber(c.Request.Context(), request.ChatId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, resp.NewAPIResponse("Successfully"))
}

func (h *Handler) CreateDeviceSubscriber(c *gin.Context) {
	var request CreateDeviceSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	err := h.service.CreateDeviceSubscriber(c.Request.Context(), request.DeviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, resp.NewAPIResponse("Successfully"))
}

func (h *Handler) EnableNotification(c *gin.Context) {
	var request EnableNotificationSubscriberRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	id, err := strconv.Atoi(c.Param("subscriberId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	err = h.service.EnableNotificationSubscriber(c.Request.Context(), uint64(id), request.Notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.NewAPIResponse("Successfully"))
}

func (h *Handler) GetSubscriberByDeviceId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("deviceId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	sub, err := h.service.GetSubscriberByDeviceId(c.Request.Context(), uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.NewAPIResponse("Successfully"))
		return
	}

	c.JSON(http.StatusOK, ConvertSubscriberEntityToResponse(sub))
}

func (h *Handler) GetSubscriberChatId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("chatId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	sub, err := h.service.GetSubscriberByChatId(c.Request.Context(), uint64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, ConvertSubscriberEntityToResponse(sub))
}
