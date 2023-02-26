package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	my_var := strings.Split(args, " ")
	name := my_var[1]
	fmt.Println("my_var - ", my_var)
	index, err := strconv.Atoi(my_var[0])

	if err != nil {
		log.Println("WRONG ARSG", args)
		return
	}

	err = c.productService.Edit(index, name)

	if err != nil {
		log.Print("Fail to delete product", args)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Edition Success!")
	// msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)
}
