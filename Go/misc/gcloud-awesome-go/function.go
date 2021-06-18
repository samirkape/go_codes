package mybot

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

// HandleTelegramWebHook parses a POST request from telegram and responds with appropriate actions.
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {
	message := ReceiveMessage{}
	chatID := 0
	msgText := ""

	// Parse incoming request
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Println(err)
			return
		}
		r.Body.Close()
	}

	if message.Message.Chat.ID > 0 {
		log.Println(message.Message.Chat.ID, message.Message.Text)
		chatID = message.Message.Chat.ID
		msgText = message.Message.Text
	}

	// check if the received response is command
	ExecuteCommand(msgText, chatID)
}

func ExecuteCommand(msgText string, chatID int) {

	// Connect to database and fetch a document from collection list
	client := GetDbClient()
	defer client.Disconnect(context.Background())
	colls := ListCollections(client, DbName)
	switch msgText {
	case CMDStart:
		SendMessage("Hello, press command button to start", chatID)
	case CMDListCategories:
		SendMessage("Hold on", chatID)
		SendMessage(ListToMsg(colls), chatID)
		SendMessage("Done!", chatID)
	case CMDListPackages:
		SendMessage("Reply with catergory number", chatID)
	// TODO
	// case "/search":
	// 	SendMessage("Reply with search term", chatID)
	default:
		CheckReply(msgText, chatID, client, DbName, colls)
	}
}

func CheckReply(msgText string, chatID int, client *mongo.Client, DbName string, colls []string) {
	// check if it is unhandled scommand
	if strings.HasPrefix(msgText, "/") {
		SendMessage("Invalid command, try numeric input", chatID)
		return
	}
	// validate package number reply for any alphabet
	pattern := regexp.MustCompile(`.*[a-zA-Z]+.*`)
	msgCharIdx := pattern.FindStringIndex(msgText)
	if msgCharIdx != nil {
		SendMessage("Invalid response. Non numeric input", chatID)
		return
	}
	categoryIdx := strings.Split(msgText, ",")
	if len(categoryIdx) > 0 {
		for _, e := range categoryIdx {
			index, err := strconv.Atoi(e)
			if err != nil {
				log.Println("Unable to convert msg to integer index")
				return
			}
			if index > len(colls) {
				ErrMsg := fmt.Sprintf("Invalid response. Number should not exceed %d", len(colls)-1)
				SendMessage(ErrMsg, chatID)
				return
			}
			pkgs, err := FindDoc(client, DbName, colls[index], "")
			for _, pkg := range pkgs {
				SendMessage(PackageToMsg(pkg), chatID)
			}
		}
	}
}

func PackageToMsg(input SplitLink) string {
	msgString := strings.Builder{}
	msgString.WriteString(fmt.Sprintf("*Name*: %s\n\n", input.Name))
	msgString.WriteString(fmt.Sprintf("*URL*: %s \n\n", input.URL))
	if input.Info != "" {
		msgString.WriteString(fmt.Sprintf("*Description*: _%s_ \n\n", input.Info))
	}
	return msgString.String()
}

// Convert slice of strings to single string
func ListToMsg(list []string) string {
	msg := strings.Builder{}
	for i, pkg := range list {
		msg.WriteString(fmt.Sprint(i) + ". " + pkg[3:] + string("\n"))
	}
	return msg.String()
}
