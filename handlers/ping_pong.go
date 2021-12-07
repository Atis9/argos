package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func pingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if isTrigger(s, m) {
		// s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+"> pong")
		s.ChannelMessageSendReply(m.ChannelID, "pong", m.Reference())
	}
}

func isTrigger(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	content := m.ContentWithMentionsReplaced()
	return strings.HasPrefix(content, "@"+s.State.User.Username) && strings.Contains(content, "ping")
}
