package slashcommand

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func allKey() {
	command := discordgo.ApplicationCommand{
		Name:        "all_key",
		Description: "get all key",
	}

	slashCommand.rCommand(slashCommandRegistry{
		command:           &command,
		commandHandleFunc: allMessageFunc,
	})
}

func allMessageFunc(c context) {
	keySlice := slashCommand.messageService.GetAllKeyByGuildId(c.interactionCreate.GuildID)
	content := ""

	if len(keySlice) > 0 {
		content = strings.Join(keySlice, ", ")
	} else {
		content = "沒有這種東西"
	}

	err := c.session.InteractionRespond(c.interactionCreate.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})

	if err != nil {
		log.Err(err).Msg("All Key Error")
	}
}
