package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tclutin/classconnect-api/internal/domain/auth"
	"github.com/tclutin/classconnect-api/internal/handler/http/middleware"
	resp "github.com/tclutin/classconnect-api/pkg/http"
	"log/slog"
	"net/http"
)

const (
	layerAuthHandler = "handler.auth."
)

type Service interface {
	LogIn(ctx context.Context, dto auth.LoginDTO) (string, error)
	SignUp(ctx context.Context, dto auth.SignupDTO) (string, error)
	GetUserByUsernameWithDetail(ctx context.Context, username string) (auth.UserDetailDTO, error)
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
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/signup", h.SignUp)
		authGroup.POST("/login", h.LogIn)
		authGroup.GET("/me", middleware.AuthMiddleware(auth), h.Me)
	}
}

func (h *Handler) LogIn(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	token, err := h.service.LogIn(c.Request.Context(), request.ToDTO())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, TokenResponse{AccessToken: token})
}

func (h *Handler) SignUp(c *gin.Context) {
	var request SignupRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	token, err := h.service.SignUp(c.Request.Context(), request.ToDTO())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, TokenResponse{AccessToken: token})
}

func (h *Handler) Me(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.NewAPIErrorResponse("username not found in context"))
		return
	}

	userDetail, err := h.service.GetUserByUsernameWithDetail(c.Request.Context(), username.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.NewAPIErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, ConvertUserDetailDTOToResponse(userDetail))
}
