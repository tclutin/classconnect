package handler

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"strings"
)

func (h *Handler) GetGroupInfoHandler(c telebot.Context) error {
	chatId := uint64(c.Chat().ID)

	if err := h.service.LogIn(); err != nil {
		return err
	}

	sub, err := h.service.GetSubscriber(chatId)
	if err != nil {
		return c.Send(ErrInternal.Error())
	}

	if sub.GroupId == nil {
		return c.Send(ErrNotAvailable.Error())
	}

	group, err := h.service.GetGroupById(*sub.GroupId)
	if err != nil {
		return c.Send(ErrInternal.Error())
	}

	var builder strings.Builder
	builder.WriteString("🌟🌟🌟 Group Information 🌟🌟🌟\n\n")
	builder.WriteString(fmt.Sprintf("🆔 ID: %d\n", group.ID))
	builder.WriteString(fmt.Sprintf("🏷️ Name: %s\n", group.Name))
	builder.WriteString(fmt.Sprintf("🗓️ Schedule: %v\n", group.IsExistsSchedule))
	builder.WriteString(fmt.Sprintf("👥 Members Count: %d\n", group.MembersCount))
	builder.WriteString(fmt.Sprintf("🕒 Created At: %s\n", group.CreatedAt.Format("2006-01-02")))
	builder.WriteString("\n🎉🎉🎉 Enjoy your time here 🎉🎉🎉")

	return c.Send(builder.String())
}
