package main

import (
	"github.com/istrel/bot/internal/service/product"
	"log"
	"os"
	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	godotenv.Load()
	token :=os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	 updates, err := bot.GetUpdatesChan(u)
	 if err != nil {
		log.Panic(err)
	 }

	 productService := product.NewService()

	for update := range updates {
		if update.Message == nil { 
			continue
		}

		if update.Message.Command() == "help"{
			helpCommand(bot, update.Message)
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case"list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

	func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
		msg:=tgbotapi.NewMessage(inputMessage.Chat.ID, 
			"/help - help\n" + "list - list products")

		bot.Send(msg)
	}

	func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productSerice *product.Serice) {
		outputMsg := "Here all products: \n\n"
		
		products := productSerice
		
		
		for _, p:= range products {
			outputMsg +=p.Title
			outputMsg +="\n"
		}
		msg:=tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

		bot.Send(msg)
	}

	func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
		log.Printf("[%s] %s",inputMessage.From.UserName, inputMessage.Text)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote:"+inputMessage.Text)

		bot.Send(msg)
	}

