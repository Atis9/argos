package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func pingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	smc := SessionAndMessageCreate{
		session: s,
		message: m,
	}

	if smc.isSelf() || !smc.isMention() || !smc.containKeyword("ping") {
		return
	}

	// s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> pong")
	s.ChannelMessageSendReply(m.ChannelID, "pong", m.Reference())
}
