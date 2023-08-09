package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	// 該訊息是bot發送的就往下執行
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	inputMessage := messageCreate.Content
	var outputMessage string
	if strings.HasPrefix(inputMessage, "!") {
		messages := strings.Split(inputMessage, " ")

		if len(messages) < 3 {
			session.ChannelMessageSend(messageCreate.ChannelID, "????")
		}

		act := messages[0][1:]

		switch act {
		case "set":
			message := toBackMessage(inputMessage)
			outputMessage = setBackMessage(message, messageCreate.GuildID)
		}

	} else {
		outputMessage = backMessageService.GetRandomValue(inputMessage, messageCreate.GuildID)
	}

	if outputMessage != "" {
		session.ChannelMessageSend(messageCreate.ChannelID, outputMessage)
	}
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

func setBackMessage(message *backMessage, guildId string) string {
	var outputMessage string
	ok := backMessageService.InsertMessage(message.key, message.value, guildId)
	if ok {
		outputMessage = "OK啦"
	} else {
		outputMessage = "???"
	}

	return outputMessage
}
