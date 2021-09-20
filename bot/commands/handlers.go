package commands

import (
	"github.com/Joh4nnesHartl/solpricebot/client"
	"github.com/Joh4nnesHartl/solpricebot/log"
	"github.com/bwmarrin/discordgo"
)

func handleDefault(s *discordgo.Session, m *discordgo.MessageCreate, words []string) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, createEmbedStatistiks("usd").MessageEmbed)
	if err != nil {
		log.Errorf("failed sending response: %s", err.Error())
	}
}

func handleSpecific(s *discordgo.Session, m *discordgo.MessageCreate, words []string) {
	if !client.IsSupportedCurrency(words[1]) {
		return
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, createEmbedStatistiks(words[1]).MessageEmbed)
	if err != nil {
		log.Errorf("failed sending response: %s", err.Error())
	}
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate, words []string) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, createEmbedHelp().MessageEmbed)
	if err != nil {
		log.Errorf("failed sending response: %s", err.Error())
	}
}
