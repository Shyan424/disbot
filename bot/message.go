package bot

import (
	"discordbot/service"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var backMessageService service.BackMessageService

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	// 該訊息是bot發送的就往下執行
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	inputMessage := messageCreate.Content
	var outputMessage string
	if inputMessage[:1] != "!" {
		outputMessage = backMessageService.GetRandomValue(inputMessage)
	} else if strings.HasPrefix(inputMessage, "!set") {
		message := toBackMessage(inputMessage)
		outputMessage = setBackMessage(message)
	}

	session.ChannelMessageSend(messageCreate.ChannelID, outputMessage)
}

type backMessage struct {
	key   string
	value string
}

func toBackMessage(inputMessage string) *backMessage {
	insertStr := strings.Split(inputMessage, " ")
	key := strings.Join(insertStr[1:len(insertStr)-1], " ")
	value := insertStr[len(insertStr)-1]

	return &backMessage{key: key, value: value}
}

func setBackMessage(message *backMessage) string {
	var outputMessage string
	ok := backMessageService.AddValue(message.key, message.value)
	if ok {
		outputMessage = "OK啦"
	} else {
		outputMessage = "???"
	}

	return outputMessage
}
