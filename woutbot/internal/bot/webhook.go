package bot

import (
    "log"

    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupWebhook(bot *tgbotapi.BotAPI) {
    // Obtenha o canal de atualizações do bot
    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := bot.GetUpdatesChan(u)
    if err != nil {
        log.Fatal(err)
    }
    
    // Chame o WebhookHandler para lidar com as atualizações recebidas
    WebhookHandler(bot, updates)
}

