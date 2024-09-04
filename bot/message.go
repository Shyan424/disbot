package bot

import (
	"discordbot/enum/res"
	"discordbot/service"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const COMMAND_PREFIX string = "+"

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	// 該訊息是bot發送的就往下執行
	if messageCreate.Author.ID == session.State.User.ID || messageCreate.Content == "" {
		return
	}

	context := context{
		guildId:            messageCreate.GuildID,
		message:            messageCreate.Content,
		attachments:        messageCreate.Attachments,
		backMessageService: service.GetBackMessageService(),
	}

	outputMessage := context.handleMessage()

	if outputMessage != "" && !strings.HasPrefix(context.message, COMMAND_PREFIX) {
		service.GetLeaderboardService().AddScore(context.guildId, context.message)
	}
	if outputMessage != "" {
		session.ChannelMessageSend(messageCreate.ChannelID, outputMessage)
	}
}

type context struct {
	guildId            string
	message            string
	attachments        []*discordgo.MessageAttachment
	backMessageService service.BackMessageService
}

func (c context) handleMessage() string {
	var outputMessage string
	if strings.HasPrefix(c.message, COMMAND_PREFIX) {
		outputMessage = c.handleCommamd()
	} else {
		outputMessage = c.backMessageService.GetRandomValue(c.message, c.guildId)
	}

	return outputMessage
}

func (c context) handleCommamd() string {
	messages := strings.Split(c.message, " ")
	act := messages[0][1:]

	switch act {
	case "set":
		if len(messages) < 3 && (len(messages) < 2 && len(c.attachments) != 1) {
			return res.FAIL.GetMsg()
		}

		if len(c.attachments) > 0 {
			messages = append(messages, c.attachments[0].URL)
		}

		message := toBackMessage(messages)
		return setBackMessage(message, c)
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

func setBackMessage(message *backMessage, c context) string {
	if c.backMessageService.InsertMessage(message.key, message.value, c.guildId) {
		return res.OK.GetMsg()
	}

	return res.FAIL.GetMsg()
}
