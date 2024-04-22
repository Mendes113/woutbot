package bot

import (
	"fmt"
	"log"
	"regexp"


	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// WebhookHandler é responsável por lidar com as atualizações recebidas pelo webhook
func WebhookHandler(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
    log.Println("Iniciando WebhookHandler")
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
    echoMessage(bot, chatID, message)

    // Verifica se a mensagem contém padrões de treino
    if isWorkoutPattern(message) {
        log.Println("Padrão de treino encontrado")
        sendMessage(bot, chatID, "Padrão de treino encontrado")

        // Divide a mensagem em linhas para processar cada conjunto de exercícios
        processWorkoutMessages(bot, chatID, message)
        MakeTrainAndSave(chatID, message)
    }
}

func echoMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
    msg := tgbotapi.NewMessage(chatID, message)
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
    msg := tgbotapi.NewMessage(chatID, text)
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }
}

func processWorkoutMessages(bot *tgbotapi.BotAPI, chatID int64, message string) {
    setsRegex := regexp.MustCompile(`\n{1,}`) // Expressão regular para duas ou mais quebras de linha
    sets := setsRegex.Split(message, -1) // Dividir a mensagem em conjuntos de treino

    var totalWorkload int // Variável para armazenar o workload total do treino
    countsets := 0
    for _, set := range sets {
        if isWorkoutPattern(set) {
            log.Println("Processando set:", set)
            train := MakeTrain(chatID, set)
            totalWorkload += calculateWorkloadFromMessage(set)
            countsets++
            log.Println("Workload Do set:", totalWorkload)
            log.Println("Sets:", countsets)
            if train != "" {
                // Atualiza o workload total do treino

                log.Println("Workload Do set:", totalWorkload)
                sendMessage(bot, chatID, train)


                // Envia uma mensagem com o workload total do treino após o loop
                sendMessage(bot, chatID, fmt.Sprintf("Workload total do treino: %d", totalWorkload))
            }
        }
    }

}



// Calcula o workload total a partir de uma mensagem de treino
func calculateWorkloadFromMessage(message string) int {
    treino := makeWorkoutFromMessage(message)
    return calculateTotalWorkload(treino)
}





// Função para verificar se uma linha contém um padrão de treino
func isWorkoutPattern(line string) bool {
    // Padrão é opcionalmente nome_do_exercicio seguido por número_de_série número_de_repetições peso
    // Exemplo: triceps 1 10 10
    isWorkout := regexp.MustCompile(`(?:[a-zA-Z]+\s+)?\d+\s+\d+\s+\d+`)
    // next line

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

