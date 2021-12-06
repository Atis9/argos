package handlers

import "github.com/bwmarrin/discordgo"

func AddHandlers(s *discordgo.Session) {
	s.AddHandler(pingPong)
}
