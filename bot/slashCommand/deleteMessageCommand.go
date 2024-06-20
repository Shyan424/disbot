package slashcommand

import (
	"discordbot/enum/res"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var deleteInfoMap = make(map[string]deleteInfo)

type deleteInfo struct {
	message    *discordgo.Interaction
	expireTime time.Time
}

func deleteMessageComand() {
	command := discordgo.ApplicationCommand{
		Name:        "delete",
		Description: "delete message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "key",
				Description: "React Key",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	reg := slashCommandRegistry{
		command:             &command,
		commandHandleFunc:   deleteMessageCommandFunc,
		componentId:         "delMsg",
		componentHandleFunc: deleteMessageComponentFunc,
	}

	slashCommand.rCommand(reg)
}

func deleteMessageCommandFunc(c context) {
	key := c.commandOptionArgMap["key"]
	messages := slashCommand.messageService.GetAllValueByKeyAndGuild(key, c.interactionCreate.GuildID)

	if len(messages) == 0 {
		c.session.InteractionRespond(c.interactionCreate.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: res.FAIL.GetMsg(),
			},
		})

		return
	}

	var content strings.Builder
	optioins := make([]discordgo.SelectMenuOption, len(messages))

	for i, message := range messages {
		content.WriteString(strconv.Itoa(i))
		content.WriteString(" ")
		content.WriteString(message.Value)
		content.WriteString("\n")

		optioins[i] = discordgo.SelectMenuOption{
			Label: strconv.Itoa(i),
			Value: key + "_" + message.Id,
		}
	}

	err := c.session.InteractionRespond(c.interactionCreate.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content.String(),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							CustomID: "delMsg",
							Options:  optioins,
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Err(err).Msgf("Delete key %s fail", key)
	}
	if err == nil {
		deleteInfoMap[key] = deleteInfo{message: c.interactionCreate.Interaction, expireTime: time.Now().Add(5 * time.Minute)}
	}
}

func deleteMessageComponentFunc(c context) {
	values := strings.Split(c.componentArgs[0], "_")
	info := getDeleteInteractionMessage(values[0])

	if info.message == nil {
		c.session.ChannelMessageSend(c.interactionCreate.ChannelID, res.FAIL.GetMsg())
		return
	}

	if isDeleteInfoExprie(info) {
		content := res.EXPIRED.GetMsg()
		c.session.InteractionResponseEdit(info.message, &discordgo.WebhookEdit{
			Components: &[]discordgo.MessageComponent{},
			Content:    &content,
		})

		return
	}

	if slashCommand.messageService.DeleteMessageByIdAndKeyAndGuildId(values[1], values[0], c.interactionCreate.GuildID) {
		deOk := res.OK.GetMsg()

		c.session.InteractionResponseEdit(info.message, &discordgo.WebhookEdit{
			Components: &[]discordgo.MessageComponent{},
			Content:    &deOk,
		})

		return
	}

	content := res.FAIL.GetMsg()
	c.session.InteractionResponseEdit(info.message, &discordgo.WebhookEdit{
		Components: &[]discordgo.MessageComponent{},
		Content:    &content,
	})
}

func getDeleteInteractionMessage(key string) deleteInfo {
	defer delete(deleteInfoMap, key)

	if info, ok := deleteInfoMap[key]; ok {
		return info
	}

	return deleteInfo{}
}

func isDeleteInfoExprie(info deleteInfo) bool {
	return info.expireTime.Before(time.Now())
}
