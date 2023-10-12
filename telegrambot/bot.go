package telegrambot

import (
	"encoding/json"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/patrickmn/go-cache"
)

type Config struct {
	ChatID   int
	BotToken string
}

var bot *tgbotapi.BotAPI
var messageId int
var chatID int64
var config Config

func loadConfig(filename string) (Config, error) {

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}

func init() {

	config, configErr := loadConfig("./config.json")
	if configErr != nil {
		panic(configErr)
	}

	var err error
	messageId = -1

	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		panic(err)
	}
}

// SendMessage sends a message using the Telegram bot.
func Inform(message string) {
	// Implement the logic for sending a message here

	// print(message)

	var msg tgbotapi.Chattable // Declare the msg variable outside the if-else blocks

	if messageId == -1 {
		msg = tgbotapi.NewMessage(int64(config.ChatID), message)
	} else {
		msg = tgbotapi.NewEditMessageText(int64(config.ChatID), messageId, message)
	}

	msgId, err := bot.Send(msg)

	if err != nil {
		panic(err)
	}

	if messageId == -1 {
		messageId = msgId.MessageID
	}

}

func SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)

}
