package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"atis.dev/argos/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

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
		fmt.Println("error creating Discord session,", err)
		return nil
	}

	return client
}

func openClient(client *discordgo.Session) {
	err := client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func runClient(client *discordgo.Session) {
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	client.Close()
}
