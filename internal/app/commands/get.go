package commands

import (
	"context"
	"log"
	"time"

	"flag"

	api "githab.com/telegrammbot/bot/api/bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

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
	respond, err := grpc.GetUsers(ctx, &api.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		respond.String())
	// msg.ReplyToMessageID = update.Message.MessageID
	c.bot.Send(msg)

	// product, err := c.productService.Get(arg)
	// if err != nil {
	// 	log.Print("Fail to get product", arg)
	// }

	// log.Printf("Greeting: %s", r.GetMessage())
}

//GET НАСТРОИТЬ НА РАБОТУ С GRCP

// func (c *Commander) Get(inputMessage *tgbotapi.Message) {
// 	args := inputMessage.CommandArguments()
// 	arg, err := strconv.Atoi(args)
// 	if err != nil {
// 		log.Println("WRONG ARSG", arg)
// 		return
// 	}
// 	product, err := c.productService.Get(arg)
// 	if err != nil {
// 		log.Print("Fail to get product", arg)
// 	}

// 	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
// 		product.Title)
// 	// msg.ReplyToMessageID = update.Message.MessageID
// 	c.bot.Send(msg)

// }

// func init() {
// 	registeredCommands["get"] = (*Commander).Get
// }
