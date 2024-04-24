package main

import (
	
	"log"
	"woutbot/config"
	"woutbot/internal/bot"

	// "woutbot/internal/bot/db"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	
	// db.Connect()
	// db.SaveToDB()
	// Obter a chave da API do arquivo de configuração (ou de onde quer que você a esteja obtendo)
	api := config.GetAPI()

	// Crie uma nova instância de BotAPI
	botAPI, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Fatal(err)
	}


	// Configure o webhook passando a instância de BotAPI
	bot.SetupWebhook(botAPI)
	
}
