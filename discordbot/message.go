package discordbot

import "github.com/bwmarrin/discordgo"

func messageCreate(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	// 該訊息是bot發送的就往下執行
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

}
