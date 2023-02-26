package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	outputMsg := "Here all products: \n"
	products := c.productService.List()
	for _, product := range products {
		outputMsg = outputMsg + product.Title
		outputMsg += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", "Some data"),
		),
	)

	c.bot.Send(msg)
}

// func init() {
// 	registeredCommands["list"] = (*Commander).List
// }
