package bot

import (
	"discordbot/enum/res"
	"discordbot/service"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if session == nil || messageCreate == nil {
		return
	}
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

	if outputMessage != "" {
		service.GetLeaderboardService().AddScore(context.guildId, context.message)
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
	return c.backMessageService.GetRandomValue(c.message, c.guildId)
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
