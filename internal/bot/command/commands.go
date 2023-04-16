package command

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Name    string
	Handler Handler
	Command *discordgo.ApplicationCommand
}

type Handler func(*discordgo.Session, *discordgo.InteractionCreate)

func GetSolanaData(name string) Command {
	command := &discordgo.ApplicationCommand{
		Name:        name,
		Description: "Shows market data about solana, defaults to USD if no currency is specified",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "currency",
				Description: "currency to show the data in",
				Required:    false,
			},
		},
	}

	handler := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		currency := "usd"

		if len(i.ApplicationCommandData().Options) > 0 {
			currency = i.ApplicationCommandData().Options[0].StringValue()
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:  discordgo.MessageFlagsEphemeral,
				Embeds: []*discordgo.MessageEmbed{createEmbedStatistiks(currency).MessageEmbed},
			},
		})
	}

	return Command{
		Name:    name,
		Handler: handler,
		Command: command,
	}
}
