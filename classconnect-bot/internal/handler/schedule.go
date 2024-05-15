package handler

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"strings"
)

func (h *Handler) GetScheduleForDayHandler(c telebot.Context) error {
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

	subjects, err := h.service.GetScheduleForDay(sub.ID)
	if err != nil {
		return c.Send(ErrInternal.Error())
	}

	if len(subjects) == 0 {
		return c.Send("⚠️ No subjects scheduled for today")
	}

	var message strings.Builder
	message.WriteString("📚 Today's Subjects 📚\n\n")

	for _, subject := range subjects {
		message.WriteString(fmt.Sprintf("📖 Subject: %s\n", subject.Name))
		message.WriteString(fmt.Sprintf("📍 Location: %s\n", subject.Cabinet))
		message.WriteString(fmt.Sprintf("👨‍🏫 Teacher: %s\n", subject.Teacher))
		message.WriteString(fmt.Sprintf("⏰ Time: %s - %s\n", subject.StartTime, subject.EndTime))
		message.WriteString(fmt.Sprintf("📝 Description: %s\n\n", subject.Description))
	}

	return c.Send(message.String())
}
