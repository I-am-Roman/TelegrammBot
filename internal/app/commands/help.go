package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "How can i help you?\n"+
		"/help\t - help\n"+
		"/list\t - list command\n"+
		"/get\t  - get all users\n"+
		"/new\t  - create user [id age first-name second-name email]\n"+
		"/delete\t - delete by email [email]\n")
	c.bot.Send(msg)
}

// func init() {
// 	registeredCommands["help"] = (*Commander).Help
// }
