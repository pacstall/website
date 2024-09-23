package log

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func connect(token string) *discordgo.Session {
	client, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("failed to connect to discord\n%v\n", err)
	}

	return client
}
