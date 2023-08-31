package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var err error
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		client.ApplicationCommandCreate(
			client.State.User.ID,
			"",
			&discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping-Pong",
			},
		)
	})

	client.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.ApplicationCommandData().Name == "ping" {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "pong",
				},
			})
		}
	})

	err = client.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}

	client.UpdateGameStatus(0, "Argos")
	log.Println("Bot is now running")

	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
