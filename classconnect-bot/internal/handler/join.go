package handler

import (
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func (h *Handler) JoinToGroupHandler(c telebot.Context) error {
	chatId := uint64(c.Chat().ID)

	if err := h.service.LogIn(); err != nil {
		return err
	}

	sub, err := h.service.GetSubscriber(chatId)
	if err != nil {
		return c.Send(ErrInternal.Error())
	}

	if sub.GroupId != nil {
		return c.Send(ErrNotAvailable.Error())
	}

	elements := strings.Split(c.Text(), " ")

	if len(elements) != 3 {
		message := "⚠️ Incorrect command format. Please use the following format:\n\n/join <code> <group_id>\n\nFor example: /join CqRF 23"
		return c.Send(message)
	}

	if len(elements[1]) != 4 {
		message := "⚠️ The code should be exactly 4 characters long. Please use the following format:\n\n/join <code> <group_id>\n\nFor example: /join CqRF 23"
		return c.Send(message)
	}

	groupId, err := strconv.ParseUint(elements[2], 10, 64)
	if err != nil {
		message := "⚠️ Please enter the group number (group_id) as a number.\n\nExample Usage:\n/join <code> <group_id>\n\nFor example: /join CqRF 23"
		return c.Send(message)
	}

	if err = h.service.JoinToGroup(groupId, sub.ID, elements[1]); err != nil {
		message := "⚠️ The group code you entered is incorrect or group_id\n\nExample Usage:\n/join <code> <group_id>\n\nFor example: /join CqRF 23"
		return c.Send(message)
	}

	menu := h.createGroupMenu()

	message := "✅ You have successfully joined the group!"
	return c.Send(message, &menu)
}
