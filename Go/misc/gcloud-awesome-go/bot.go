package mybot

import (
	"errors"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI

var Token string

const USERID = 1346530914

// bot constructor
func init() {
	Token = os.Getenv("TOKEN")
	bot, err := getTBot()
	Bot = bot
	if err != nil {
		fmt.Println("bot initialization failed")
		os.Exit(-1)
	}
}

// initialize and validate bot
func getTBot() (*tgbotapi.BotAPI, error) {
	if len(Token) == 0 {
		return nil, errors.New("getTBot: could not find bot token")
	}
	bot, err := tgbotapi.NewBotAPI(Token)
	//bot.Debug = true
	if err != nil {
		return nil, fmt.Errorf("getTBot: error initializing bot: %v", err)
	}
	return bot, err
}

func NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: true,
	}
}

func SendMessage(Info string, userid int) error {
	msg := NewMessage(int64(userid), Info)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := Bot.Send(msg)
	if err != nil {
		return fmt.Errorf("sendmessage: message sending failed: %v", err)
	}
	return nil
}
