package events

import (
	"fmt"

	"github.com/Joh4nnesHartl/solpricebot/internal/bot/command"
	"github.com/Joh4nnesHartl/solpricebot/pkg/log"
	"github.com/bwmarrin/discordgo"
)

func OnConnect() func(*discordgo.Session, *discordgo.Connect) {
	return func(s *discordgo.Session, _ *discordgo.Connect) {
		log.Info(fmt.Sprintf("Shard #%v connected", s.ShardID))
	}
}

func OnInteractionCreate(commandHandlers map[string]command.Handler) func(*discordgo.Session, *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			handler(s, i)
		} else {
			log.Error(fmt.Sprintf("Unknown command: %v", i.ApplicationCommandData().Name))
		}
	}
}
