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

func SendMessage(Info string) error {
	msg := tgbotapi.NewMessage(USERID, Info)
	//msg1 := tgbotapi.NewMessage(types.GROUPID, Info)
	msg.ParseMode = "markdown"
	_, err := Bot.Send(msg)
	//_, err = bot.Send(msg1)
	if err != nil {
		return fmt.Errorf("sendmessage: message sending failed: %v", err)
	}
	return nil
}
