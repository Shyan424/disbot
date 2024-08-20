package slashcommand

import (
	"discordbot/service"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func leaderboard() slashCommandRegistry {
	command := discordgo.ApplicationCommand{
		Name:        "top",
		Description: "top 10 message",
	}

	return slashCommandRegistry{
		command:           &command,
		commandHandleFunc: leaderboardFunc,
	}
}

func leaderboardFunc(context context) {
	top := service.GetLeaderboardService().GetTopLeader(context.interactionCreate.GuildID, 10)
	content := strings.Builder{}

	for i, v := range top {
		content.WriteString(strconv.Itoa(i))
		content.WriteString(". ")
		content.WriteString(v)
		content.WriteString("\n")
	}

	err := context.session.InteractionRespond(context.interactionCreate.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content.String(),
		},
	})

	if err != nil {
		log.Err(err).Msg("All Key Error")
	}
}
