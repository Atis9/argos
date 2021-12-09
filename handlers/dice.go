package handlers

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func dice(s *discordgo.Session, m *discordgo.MessageCreate) {
	smc := SessionAndMessageCreate{
		session: s,
		message: m,
	}

	if smc.isSelf() || !smc.isMention() || !smc.containKeyword("dice") {
		return
	}

	diceString := strings.Split(smc.message.Content, " ")[2]
	rollDiceResult := rollDice(diceString)
	message := "Result: " + strconv.Itoa(rollDiceResult)
	s.ChannelMessageSendReply(m.ChannelID, message, m.Reference())
}

func rollDice(dice string) int {
	array := strings.Split(dice, "d")
	n, err := strconv.Atoi(array[0])
	if err != nil {
		return 0
	}
	size, err := strconv.Atoi(array[1])
	if err != nil {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	max := n*size + 1 - n
	result := rand.Intn(max) + n
	return result
}
