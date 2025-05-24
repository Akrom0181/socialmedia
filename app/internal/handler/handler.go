package handler

import (
	"log/slog"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
	"gitlab.saidoff.uz/company/muslim-administration/reading/back/internal/config"
	"gitlab.saidoff.uz/company/muslim-administration/reading/back/internal/service"
)

type Handler struct {
	logger  *slog.Logger
	service service.I
	cfg     *config.Config
	redis   service.RedisI
}

func NewHandler(logger *slog.Logger, service service.I, cfg *config.Config, redis service.RedisI) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
		cfg:     cfg,
		redis:   redis,
	}
}

func (h *Handler) Register(router *router.Router[*core.RequestEvent]) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/password-reset-otp-confirm", h.PasswordResetOTPConfirmHandler)
		}

	}
}
