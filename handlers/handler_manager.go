package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type SessionAndMessageCreate struct {
	session *discordgo.Session
	message *discordgo.MessageCreate
}

func AddHandlers(s *discordgo.Session) {
	s.AddHandler(pingPong)
	s.AddHandler(dice)
}

func (smc *SessionAndMessageCreate) isMention() bool {
	content := smc.message.ContentWithMentionsReplaced()
	selfUsername := smc.session.State.User.Username
	return strings.HasPrefix(content, "@"+selfUsername)
}

func (smc *SessionAndMessageCreate) containKeyword(keyword string) bool {
	content := smc.message.ContentWithMentionsReplaced()
	return strings.Contains(content, keyword)
}

func (smc *SessionAndMessageCreate) isSelf() bool {
	return smc.message.Author.ID == smc.session.State.User.ID
}
