package bot

import (
	"log"



	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// WebhookHandler é responsável por lidar com as atualizações recebidas pelo webhook
func WebhookHandler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
    for update := range updates {
        if update.Message != nil {
            if update.Message.IsCommand() {
                // Lide com mensagens de comando aqui
                sendWelcomeMessage(bot, update.Message.Chat.ID)
            } else {
                // Lide com outras mensagens aqui, se necessário
            }
        } else if update.CallbackQuery != nil {
            // Lide com queries de callback aqui
            handleCallback(bot, update.CallbackQuery)
        }
    }
}

func sendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
    msg := tgbotapi.NewMessage(chatID, "Olá! Eu sou o WoutBot. Como posso te ajudar?")

    // Crie uma linha de botões
    row := []tgbotapi.KeyboardButton{
        tgbotapi.NewKeyboardButton("Comandos"),
        tgbotapi.NewKeyboardButton("Sobre"),
    }

    // Adicione o teclado de resposta à mensagem
    msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(row...))

    // Envie a mensagem com os botões de teclado de resposta
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }


}

