package bot

import (
	"discordbot/enum/res"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	// 該訊息是bot發送的就往下執行
	if messageCreate.Author.ID == session.State.User.ID || messageCreate.Content == "" {
		return
	}

	context := context{
		guildId:     messageCreate.GuildID,
		message:     messageCreate.Content,
		attachments: messageCreate.Attachments,
	}

	outputMessage := context.handleMessage()

	if outputMessage != "" {
		session.ChannelMessageSend(messageCreate.ChannelID, outputMessage)
	}
}

type context struct {
	guildId     string
	message     string
	attachments []*discordgo.MessageAttachment
}

func (c context) handleMessage() string {
	var outputMessage string
	if strings.HasPrefix(c.message, "!") {
		outputMessage = c.handleCommamd()
	} else {
		outputMessage = backMessageService.GetRandomValue(c.message, c.guildId)
	}

	return outputMessage
}

func (c context) handleCommamd() string {
	messages := strings.Split(c.message, " ")
	act := messages[0][1:]

	switch act {
	case "set":
		if len(messages) > 3 || (len(messages) > 2 && len(c.attachments) == 1) {
			return res.FAIL.GetMsg()
		}

		if len(c.attachments) > 0 {
			messages = append(messages, c.attachments[0].URL)
		}

		message := toBackMessage(messages)
		return setBackMessage(message, c.guildId)
	}

	return res.WHAT.GetMsg()
}

type backMessage struct {
	key   string
	value string
}

func toBackMessage(messages []string) *backMessage {
	key := strings.Join(messages[1:len(messages)-1], " ")
	value := messages[len(messages)-1]

	return &backMessage{key: key, value: value}
}

func setBackMessage(message *backMessage, guildId string) string {
	var outputMessage res.Res
	ok := backMessageService.InsertMessage(message.key, message.value, guildId)
	if ok {
		outputMessage = res.OK
	} else {
		outputMessage = res.FAIL
	}

	return outputMessage.GetMsg()
}
