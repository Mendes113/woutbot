package bot

import (
	"log"
	"regexp"
	"strings"

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
                readMessageFromChat(bot, update.Message.Chat.ID, update.Message.Text)

            }
        } else if update.CallbackQuery != nil {

        }
    }
}


func readMessageFromChat(bot *tgbotapi.BotAPI, chatID int64, message string) {
    log.Println("Mensagem recebida:", message)

    // Replica a mensagem que o usuário mandou
    msg := tgbotapi.NewMessage(chatID, message)
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }

    // Verifica se a mensagem contém padrões de treino
    if isWorkoutPattern(message) {
        log.Print("Padrão de treino encontrado")
        msg := tgbotapi.NewMessage(chatID, "Padrão de treino encontrado")
        _, err := bot.Send(msg)
        if err != nil {
            log.Println("Erro ao enviar mensagem:", err)
        }

        // Divide a mensagem em linhas para processar cada conjunto de exercícios
        lines := strings.Split(message, "\n")
        for _, line := range lines {
            if isWorkoutPattern(line) {
                train := MakeTrain(chatID, line)
                if train != "" {
                    msg := tgbotapi.NewMessage(chatID, train)
                    _, err := bot.Send(msg)
                    if err != nil {
                        log.Println("Erro ao enviar mensagem:", err)
                    }
                }
            }
        }
    }
}

// Função para verificar se uma linha contém um padrão de treino
func isWorkoutPattern(line string) bool {
    // Padrão é opcionalmente nome_do_exercicio seguido por número_de_série número_de_repetições peso
    // Exemplo: triceps 1 10 10
    isWorkout := regexp.MustCompile(`(?:[a-zA-Z]+\s+)?\d+\s+\d+\s+\d+`)
    return isWorkout.MatchString(line)
}



// func identifyEndTraining(message string) bool {
//     // Padrão para identificar o fim do treino
//     endTrainingPattern := regexp.MustCompile(`acabar\s+treino|encerrar\s+treino|fim\s+treino| fim`)

//     // Verifica se a mensagem corresponde ao padrão de encerramento do treino
//     return endTrainingPattern.MatchString(message)
// }





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

