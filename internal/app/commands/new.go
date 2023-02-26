package commands

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	api "githab.com/telegrammbot/bot/api/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()
	my_var := strings.Split(args, " ")
	id := my_var[0]
	age := my_var[1]
	first_name := my_var[2]
	second_name := my_var[3]
	email := my_var[4]

	print(id, age, first_name, second_name, email)

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
	respond, err := grpc.CreateUser(ctx, &api.User{Id: id, Age: age, FirstName: first_name, LastName: second_name, Email: email})
	fmt.Println("Respond - ", respond)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"Creation Success")
	// msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)

}
