package handler

import (
	"github.com/tclutin/classconnect-bot/internal/service"
	"gopkg.in/telebot.v3"
)

type Handler struct {
	bot     *telebot.Bot
	service *service.Service
}

func NewHandler(bot *telebot.Bot, service *service.Service) *Handler {
	return &Handler{
		bot:     bot,
		service: service,
	}
}

func (h *Handler) Init() {
	//Just commands
	h.bot.Handle("/start", h.StartHandler)
	h.bot.Handle("/groups", h.ShowGroupsHandler)
	h.bot.Handle("/join", h.JoinToGroupHandler)

	//Reply menu
	h.bot.Handle("👥 My group", h.GetGroupInfoHandler)
	h.bot.Handle("🗓️ Getting a schedule for today", h.GetScheduleForDayHandler)
	h.bot.Handle("❌ Exit", h.ExitFromGroupHandler)

	//Events
	h.bot.Handle(telebot.OnCallback, h.CallbackHandler)

}
