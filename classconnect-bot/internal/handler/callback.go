package handler

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

const (
	JoinToGroupConfirm = "join_group_confirm"
	JoinToGroup        = "join_group"
)

func (h *Handler) CallbackHandler(c telebot.Context) error {
	data := c.Callback().Data

	if err := h.service.LogIn(); err != nil {
		return err
	}

	if strings.Contains(data, JoinToGroupConfirm) {
		groupIdstr := strings.Split(data, "_")[3]

		message := fmt.Sprintf("🚀 To join the group, please use the following command format:\n\n/join <code> <group_id>\n\nFor example: /join CqRv %s\n\nYou can obtain the <code> from the group owner.", groupIdstr)

		return c.Send(message)
	}

	if strings.Contains(data, JoinToGroup) {
		groupIdstr := strings.Split(data, "_")[2]

		groupId, err := strconv.ParseUint(groupIdstr, 10, 64)
		if err != nil {
			return c.Send(err.Error())
		}

		group, err := h.service.GetGroupById(groupId)
		if err != nil {
			return c.Send(err.Error())
		}

		joinBtn := telebot.InlineButton{
			Text:   "Join",
			Unique: fmt.Sprintf("join_group_confirm_%d", group.ID),
		}

		menu := &telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{{joinBtn}},
		}

		var builder strings.Builder
		builder.WriteString("🌟🌟🌟 Group Information 🌟🌟🌟\n\n")
		builder.WriteString(fmt.Sprintf("🆔 ID: %d\n", group.ID))
		builder.WriteString(fmt.Sprintf("🏷️ Name: %s\n", group.Name))
		builder.WriteString(fmt.Sprintf("👥 Members Count: %d\n", group.MembersCount))
		builder.WriteString(fmt.Sprintf("🕒 Created At: %s\n", group.CreatedAt.Format("2006-01-02")))
		builder.WriteString("\n🎉🎉🎉 Enjoy your time here 🎉🎉🎉")

		return c.Send(builder.String(), menu)
	}

	return nil
}
