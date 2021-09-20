package events

import (
	"github.com/bwmarrin/discordgo"
)

func AssignSession(bot *discordgo.Session) {
	bot.AddHandler(messageCreateHandler)
}
