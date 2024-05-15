package handler

import "gopkg.in/telebot.v3"

func (h *Handler) ExitFromGroupHandler(c telebot.Context) error {
	chatId := uint64(c.Chat().ID)

	if err := h.service.LogIn(); err != nil {
		return c.Send(err.Error())
	}

	sub, err := h.service.GetSubscriber(chatId)
	if err != nil {
		return c.Send(ErrInternal.Error())
	}

	if sub.GroupId == nil {
		return c.Send(ErrNotAvailable.Error())
	}

	if err = h.service.LeaveFromGroup(*sub.GroupId, sub.ID); err != nil {
		return c.Send(ErrInternal.Error())
	}

	return c.Send("✅ You have left the group", &telebot.ReplyMarkup{RemoveKeyboard: true})

}
