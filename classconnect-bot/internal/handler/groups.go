package handler

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

func (h *Handler) ShowGroupsHandler(c telebot.Context) error {
	chatId := uint64(c.Chat().ID)

	if err := h.service.LogIn(); err != nil {
		return c.Send(err.Error())
	}

	sub, err := h.service.GetSubscriber(chatId)
	if err != nil {
		return c.Send(err.Error())
	}

	if sub.GroupId != nil {
		return c.Send(ErrNotAvailable.Error())
	}

	groups, err := h.service.GetGroups()
	if err != nil {
		return c.Send(err.Error())
	}

	var buttons [][]telebot.InlineButton
	for _, group := range groups {
		btn := telebot.InlineButton{
			Text:   fmt.Sprintf("%s (%d)", group.Name, group.ID),
			Unique: fmt.Sprintf("join_group_%d", group.ID),
		}
		buttons = append(buttons, []telebot.InlineButton{btn})
	}

	keyboard := telebot.ReplyMarkup{
		InlineKeyboard: buttons,
	}

	return c.Send("Select the group you want to join", &keyboard)
}
