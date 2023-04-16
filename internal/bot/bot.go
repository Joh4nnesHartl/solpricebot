package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/top-gg/go-dbl"

	"github.com/Joh4nnesHartl/solpricebot/internal/bot/command"
	"github.com/Joh4nnesHartl/solpricebot/internal/bot/events"
	"github.com/Joh4nnesHartl/solpricebot/internal/bot/shards"
	"github.com/Joh4nnesHartl/solpricebot/internal/client"
	"github.com/Joh4nnesHartl/solpricebot/pkg/log"
)

const (
	botID = 888052145734172732
)

func NewSessionManager(token string) (*shards.Manager, error) {
	manager, err := shards.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	commands := []command.Command{
		command.GetSolanaData("sol"),
	}

	handlers := make(map[string]command.Handler, len(commands))
	for _, c := range commands {
		handlers[c.Name] = c.Handler
	}

	manager.AddHandler(events.OnConnect())
	manager.AddHandler(events.OnInteractionCreate(handlers))

	// TODO: Move this to a better place

	err = manager.Start()
	if err != nil {
		return nil, err
	}

	for _, cmd := range commands {
		_, err := manager.Gateway.ApplicationCommandCreate(manager.Shards[0].Session.State.User.ID, "", cmd.Command)
		if err != nil {
			return nil, err
		}
	}

	manager.RegisterIntent(discordgo.IntentGuildMessages)

	return manager, nil
}

func NewDblClient(token string) (*dbl.Client, error) {
	return dbl.NewClient(token)
}

func StartStatusUpdater(manager *shards.Manager) {
	go func() {
		for range client.Changed {
			coinData, _ := client.GetSolMarketData("usd")

			status := fmt.Sprintf("$%.2f", coinData.CurrentPrice)
			for _, s := range manager.Shards {
				s.Session.UpdateStatusComplex(discordgo.UpdateStatusData{
					Activities: []*discordgo.Activity{
						{
							Name: status,
							Type: discordgo.ActivityTypeWatching,
						},
					},
				})

				log.Info(fmt.Sprintf("[SHARD: %d] updated bot status to: %s", s.ID, status))
			}
		}
	}()
}

func StartDBLupdater(dblClient *dbl.Client, manager *shards.Manager) {
	time.AfterFunc(20*time.Second, func() { updateDBLdata(dblClient, manager) })

	go func() {
		for range time.Tick(30 * time.Minute) {
			err := updateDBLdata(dblClient, manager)
			if err != nil {
				log.Error(fmt.Sprintf("error updating top.gg data: %s", err.Error()))
			}
		}
	}()
}

func updateDBLdata(dblClient *dbl.Client, manager *shards.Manager) error {
	serverCountPerShard := []int{}
	totalServerCount := 0
	var shardsToServerCountStr strings.Builder

	for _, shard := range manager.Shards {
		shardsToServerCountStr.WriteString(fmt.Sprintf(" [Shard: %d] | [Guids: %d]", shard.ID, shard.GuildCount()))
		serverCountPerShard = append(serverCountPerShard, shard.GuildCount())
		totalServerCount += shard.GuildCount()
	}

	log.Info(fmt.Sprintf("Updating server top.gg server count: (Total Server Count: %d)%s", totalServerCount, shardsToServerCountStr.String()))
	return dblClient.PostBotStats(fmt.Sprint(botID), &dbl.BotStatsPayload{
		Shards:     serverCountPerShard,
		ShardCount: manager.ShardCount,
	})
}
