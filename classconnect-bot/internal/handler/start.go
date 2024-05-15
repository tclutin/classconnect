package handler

import (
	"gopkg.in/telebot.v3"
)

func (h *Handler) StartHandler(c telebot.Context) error {
	chatId := uint64(c.Chat().ID)

	if err := h.service.LogIn(); err != nil {
		return err
	}

	sub, _ := h.service.GetSubscriber(chatId)
	if sub.ID == 0 {
		if err := h.service.CreateSubscriber(chatId); err != nil {
			return err
		}
	}

	var keyboard telebot.ReplyMarkup

	if sub.GroupId != nil {
		keyboard = h.createGroupMenu()
		keyboard.ResizeKeyboard = true
	}

	welcome := `🎉 Welcome to ClassConnect! 🎓

    Get ready to embark on your educational journey with us! 🚀 Now you can effortlessly join study groups, receive timely notifications, and keep track of your class schedules.
	
	Join a group to stay updated and never miss out on important study sessions or events! 📚 ClassConnect is here to make your academic life easier and more organized.

	Use the /groups command to view groups

	Simply choose a group to join and unlock all the amazing features our app has to offer! Let's get started! 💪`

	return c.Send(welcome, &keyboard)
}

func (h *Handler) createGroupMenu() telebot.ReplyMarkup {
	return telebot.ReplyMarkup{
		ReplyKeyboard: [][]telebot.ReplyButton{
			{telebot.ReplyButton{Text: "👥 My group"}},
			{telebot.ReplyButton{Text: "🗓️ Getting a schedule for today"}},
			{telebot.ReplyButton{Text: "❌ Exit"}},
		}, ResizeKeyboard: true,
	}
}
