package slashcommand

import (
	"discordbot/enum/res"

	"github.com/bwmarrin/discordgo"
)

func setMessageCommand() {
	command := discordgo.ApplicationCommand{
		Name:        "set",
		Description: "set message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "key",
				Description: "React Key",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
			{
				Name:        "value",
				Description: "React Value",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
			},
		},
	}

	reg := slashCommandRegistry{command: &command, commandHandleFunc: setMessageCommandFunc}

	slashCommand.rCommand(reg)
}

func setMessageCommandFunc(c context) {
	if slashCommand.messageService.InsertMessage(c.commandOptionArgMap["key"], c.commandOptionArgMap["value"], c.interactionCreate.GuildID) {
		c.session.InteractionRespond(c.interactionCreate.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: res.OK.GetMsg(),
			},
		})
	}

	c.session.InteractionRespond(c.interactionCreate.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: res.FAIL.GetMsg(),
		},
	})
}
