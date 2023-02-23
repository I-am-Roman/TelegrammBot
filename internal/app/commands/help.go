package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "How can i help you?\n"+
		"/help - help\n"+
		"/list - list products")
	c.bot.Send(msg)
}
