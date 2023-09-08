package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	client, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	client.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		s.ApplicationCommandCreate(
			s.State.User.ID,
			"",
			&discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping-Pong",
			},
		)
		s.ApplicationCommandCreate(
			s.State.User.ID,
			"",
			&discordgo.ApplicationCommand{
				Name:        "roll",
				Description: "Roll NdM",
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:        discordgo.ApplicationCommandOptionString,
						Name:        "dice",
						Description: "e.g.) 1d6",
						Required:    true,
					},
				},
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
		} else if i.ApplicationCommandData().Name == "roll" {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			var rollDiceResult string
			if option, ok := optionMap["dice"]; ok {
				str := option.StringValue()
				r, _ := regexp.Compile(`\d+d\d+`)
				if r.MatchString(str) {
					array := strings.Split(str, "d")
					n, err := strconv.Atoi(array[0])
					size, err := strconv.Atoi(array[1])
					if err != nil {
						rollDiceResult = "0"
					} else {
						max := n*size + 1 - n
						result := rand.Intn(max) + n
						rollDiceResult = strconv.Itoa(result)
					}
				} else {
					rollDiceResult = "0"
				}
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: rollDiceResult,
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

	log.Println("Removing commands")
	registeredCommands, err := client.ApplicationCommands(client.State.User.ID, "")
	if err != nil {
		log.Fatalf("Could not fetch registered commands: %v", err)
	}

	for _, v := range registeredCommands {
		err := client.ApplicationCommandDelete(client.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Gracefully shutdowning")
}
