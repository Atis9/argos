package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"atis.dev/argos/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	client := getClient(token)

	handlers.AddHandlers(client)
	openClient(client)
	runClient(client)
}

func getClient(token string) *discordgo.Session {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Cannot create the session: %v", err)
		return nil
	}

	return client
}

func openClient(client *discordgo.Session) {
	err := client.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}
}

func runClient(client *discordgo.Session) {
	client.UpdateGameStatus(0, "Argos")
	log.Println("Bot is now running")

	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
	log.Println("Gracefully shutdowning")
}
