package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	selectors = []selector{}
)

func init() {
	selectors = append(selectors, defaultSelector)
	selectors = append(selectors, specificSelector)
	selectors = append(selectors, helpSelector)
}

type selector func(words []string) (bool, handler)

type handler func(*discordgo.Session, *discordgo.MessageCreate, []string)

func GetHandler(words []string) handler {
	for _, s := range selectors {
		if ok, h := s(words); ok {
			return h
		}
	}

	return handleHelp
}
