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

    // Chame a função para enviar o botão inline
    sendInlineButton(bot, chatID)
}

func sendInlineButton(bot *tgbotapi.BotAPI, chatID int64) {
    // Crie um novo botão inline
    inlineButton := tgbotapi.NewInlineKeyboardButtonData("Novo Treino!", "novo_treino")
    inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(inlineButton))

    // Crie uma nova mensagem com o botão inline
    msg := tgbotapi.NewMessage(chatID, "Pressione Um dos botões abaixo:")
    msg.ReplyMarkup = inlineKeyboard

    // Envie a mensagem com o botão inline
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }
}


func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
    msg := tgbotapi.NewMessage(chatID, text)
    _, err := bot.Send(msg)
    if err != nil {
        log.Println("Erro ao enviar mensagem:", err)
    }
}
// Variável global para armazenar o estado do botão "Novo Treino!"
var novoTreinoPending map[int64]bool

func init() {
    novoTreinoPending = make(map[int64]bool)
}

// HandleCallback lida com uma query de callback de botão inline
func handleCallback(bot *tgbotapi.BotAPI, query *tgbotapi.CallbackQuery) {
    // Aqui você pode processar a query de callback e responder de acordo
    switch query.Data {
    case "novo_treino":
        // Verifica se já está aguardando um novo treino para este usuário
        if _, ok := novoTreinoPending[query.Message.Chat.ID]; ok {
            SendMessage(bot, query.Message.Chat.ID, "Já estamos aguardando seu novo treino.")
            return
        }

        // Marca que estamos aguardando um novo treino para este usuário
        novoTreinoPending[query.Message.Chat.ID] = true

        // Solicite ao usuário que insira o treino
        msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Por favor, digite o nome do novo treino:")
        _, err := bot.Send(msg)
        if err != nil {
            log.Println("Erro ao enviar mensagem:", err)
        }
    default:
        // Caso nenhum botão correspondente seja encontrado
        log.Println("Botão não reconhecido:", query.Data)
    }
}

// waitForTrainingInput aguarda a entrada do usuário para o novo treino
func waitForTrainingInput(bot *tgbotapi.BotAPI, chatID int64) {
    // Obtenha um canal de atualizações para receber as mensagens do usuário
    updates, err := bot.GetUpdatesChan(tgbotapi.UpdateConfig{
        Timeout: 60,
    })
    if err != nil {
        log.Println("Erro ao obter o canal de atualizações:", err)
        return
    }

    // Itere sobre as atualizações para receber as mensagens do usuário
    for update := range updates {
        if update.Message != nil && novoTreinoPending[update.Message.Chat.ID] {
            // Processar o treino inserido pelo usuário
            novoTreino := update.Message.Text
            feedbackTreino := MontarTreino(update.Message.Chat.ID, novoTreino)
            SendMessage(bot, update.Message.Chat.ID, "Treino montado! "+feedbackTreino)
            // Parar de aguardar por novas mensagens e limpar o estado do novo treino
            delete(novoTreinoPending, update.Message.Chat.ID)
            break
        }
    }
}

