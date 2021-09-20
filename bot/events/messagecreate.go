package events

import (
	"strings"

	"github.com/Joh4nnesHartl/solpricebot/bot/commands"
	"github.com/bwmarrin/discordgo"
)

// CommandPrefix the prefix
const (
	commandPrefix = "!sol"
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	words := strings.Split(m.Content, " ")
	if len(words) == 0 {
		return
	}

	if words[0] != commandPrefix {
		return
	}

	handler := commands.GetHandler(words)

	handler(s, m, words)
}
