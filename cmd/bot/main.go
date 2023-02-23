package main

import (
	"fmt"
	"log"
	"os"

	"githab.com/telegrammbot/bot/internal/app/commands"
	"githab.com/telegrammbot/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("TOKEN")

	//  зашружает из  .env данные
	godotenv.Load()

	// передавать можно токен в командной строке
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// выводит все сообщения, которые пришли
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// бот обращается и просит update (свежие сообщение)
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	some_sub := commands.NewCommander(bot, productService)
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "help":
				some_sub.Help(update.Message)
			case "list":
				some_sub.List(update.Message)
			default:
				some_sub.Default(update.Message)

			}

		}
	}
}
