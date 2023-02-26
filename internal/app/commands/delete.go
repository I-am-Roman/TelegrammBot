package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	api "githab.com/telegrammbot/bot/api/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	emal := args
	// if err != nil {
	// 	log.Println("WRONG ARSG", arg)
	// 	return
	// }
	// err = c.productService.Delete(arg)

	// if err != nil {
	// 	log.Print("Fail to delete product", arg)
	// }

	// // берем аргументы из командной строки
	// args := inputMessage.CommandArguments()
	// arg, err := strconv.Atoi(args)
	// if err != nil {
	// 	log.Println("WRONG ARSG", arg)
	// 	return
	// }

	// работа с grpc клиентом
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpc := api.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	respond, err := grpc.DeletUser(ctx, &api.Email{Message: emal})
	fmt.Println("Respond - ", respond)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Deletion Success!")
	// msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)
}
